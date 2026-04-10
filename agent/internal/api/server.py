"""
Flask API server for Activity Presence Controller.
Provides HTTP endpoints for session management, status, and logging.
"""

import json
import logging
import os
import threading
from datetime import datetime, timezone
from threading import Lock
from flask import Flask, request, jsonify, Response, stream_with_context
from flask_cors import CORS
from functools import wraps

from internal.behavior.engine import Engine
from internal.system.platform import Platform

logger = logging.getLogger(__name__)


class APCServer:
    """Manages the HTTP server state and behavior."""

    def __init__(self, config):
        self.config = config
        self.lock = Lock()

        # State
        self.state = {
            "state": "idle",
            "mode": "",
            "userPresent": False,
            "policyLocked": False,
            "nextAction": "",
            "startedAt": None,
        }

        self.logs = []
        self.session_cancel = None
        self.session_thread = None

        # Initialize platform
        if config["cloud_mode"]:
            logger.info("running in cloud mode (cursor movement disabled)")
            self.platform = None
        else:
            try:
                self.platform = Platform()
                logger.info("platform init success: Windows")
            except Exception as e:
                logger.warning(f"platform init warning: {e}")
                self.platform = None

        self.engine = Engine(self.platform)

    def log_action(self, action, meta=None):
        """Log an action with metadata."""
        if meta is None:
            meta = {}
        with self.lock:
            entry = {
                "ts": datetime.now(timezone.utc).isoformat(),
                "action": action,
                "meta": meta,
            }
            self.logs.append(entry)

    def get_snapshot(self):
        """Get current state snapshot."""
        with self.lock:
            return self.state.copy()

    def update_state(self, **kwargs):
        """Update state safely."""
        with self.lock:
            self.state.update(kwargs)

    # =====================
    # HTTP Handlers
    # =====================

    def handle_health(self):
        """GET /health - Liveness check."""
        return jsonify({"status": "ok"})

    def handle_status(self):
        """GET /status - Return current state."""
        return jsonify(self.get_snapshot())

    def handle_start_session(self):
        """POST /session/start - Start an activity session."""
        if request.method != "POST":
            return "", 405

        try:
            data = request.get_json() or {}
        except Exception:
            return jsonify({"error": "invalid payload"}), 400

        with self.lock:
            if self.state["state"] == "active":
                return jsonify({"error": "session already active"}), 409

            if self.state["policyLocked"]:
                return jsonify({"error": "policy locked"}), 403

            mode = data.get("mode", "default")
            randomness = data.get("randomness", 0.5)
            idle_threshold = data.get("idleThresholdSeconds", 15)
            max_duration = data.get("maxDurationMinutes", 0)

            self.state["state"] = "active"
            self.state["mode"] = mode
            self.state["startedAt"] = datetime.now(timezone.utc).isoformat()
            self.state["nextAction"] = "simulate in 5s"

        self.log_action("session_start", {"mode": mode})

        # Start engine in background thread
        cancel_event = threading.Event()
        self.session_cancel = cancel_event

        def run_engine():
            self.engine.run(cancel_event, mode, self.log_action)
            with self.lock:
                self.state["state"] = "idle"
                self.state["nextAction"] = ""

        thread = threading.Thread(target=run_engine, daemon=True)
        self.session_thread = thread
        thread.start()

        return jsonify({"status": "started"}), 200

    def handle_stop_session(self):
        """POST /session/stop - Stop the current session."""
        if request.method != "POST":
            return "", 405

        with self.lock:
            cancel_event = self.session_cancel
            self.session_cancel = None

        if cancel_event:
            cancel_event.set()

        self.update_state(state="idle", nextAction="")
        self.log_action("session_stop", {})

        return jsonify({"status": "stopped"}), 200

    def handle_policy_lock(self):
        """POST /policy/lock - Lock/unlock policy."""
        if request.method != "POST":
            return "", 405

        try:
            data = request.get_json() or {}
        except Exception:
            return jsonify({"error": "invalid payload"}), 400

        locked = data.get("locked", False)

        action = "policy_locked" if locked else "policy_unlocked"
        if locked:
            self.update_state(policyLocked=True, state="paused")
        else:
            self.update_state(policyLocked=False)

        self.log_action(action, {})

        return jsonify({"status": action}), 200

    def handle_logs(self):
        """GET /logs - Return action logs."""
        with self.lock:
            return jsonify({"entries": self.logs})

    def handle_stream(self):
        """GET /stream - Server-sent events stream."""
        def generate():
            # Send initial snapshot
            data = json.dumps(self.get_snapshot())
            yield f"event: status\ndata: {data}\n\n"

            # Stream updates every 2 seconds
            import time

            last_state = None
            while True:
                try:
                    time.sleep(2)
                    current_state = self.get_snapshot()
                    if current_state != last_state:
                        data = json.dumps(current_state)
                        yield f"event: status\ndata: {data}\n\n"
                        last_state = current_state
                except GeneratorExit:
                    break

        response = Response(stream_with_context(generate()), mimetype="text/event-stream")
        response.headers["Cache-Control"] = "no-cache"
        response.headers["Connection"] = "keep-alive"
        response.headers["X-Accel-Buffering"] = "no"
        return response


def _auth_required(f):
    """Decorator for authentication check."""

    @wraps(f)
    def decorated_function(*args, **kwargs):
        server = kwargs.get("server")
        if not server or not server.config["pairing_token"]:
            return f(*args, **kwargs)

        auth_header = request.headers.get("Authorization", "")
        expected = f"Bearer {server.config['pairing_token']}"
        if auth_header != expected:
            return jsonify({"error": "unauthorized"}), 401

        return f(*args, **kwargs)

    return decorated_function


def create_app(config):
    """Create and configure the Flask app."""
    app = Flask(__name__)
    CORS(app)

    server = APCServer(config)

    # Mount routes
    @app.route("/health", methods=["GET"])
    def health():
        return server.handle_health()

    @app.route("/status", methods=["GET"])
    @_auth_required
    def status(server=server):
        return server.handle_status()

    @app.route("/session/start", methods=["POST"])
    @_auth_required
    def start_session(server=server):
        return server.handle_start_session()

    @app.route("/session/stop", methods=["POST"])
    @_auth_required
    def stop_session(server=server):
        return server.handle_stop_session()

    @app.route("/policy/lock", methods=["POST"])
    @_auth_required
    def policy_lock(server=server):
        return server.handle_policy_lock()

    @app.route("/logs", methods=["GET"])
    @_auth_required
    def logs(server=server):
        return server.handle_logs()

    @app.route("/stream", methods=["GET"])
    @_auth_required
    def stream(server=server):
        return server.handle_stream()

    @app.errorhandler(404)
    def not_found(e):
        return jsonify({"error": "not found"}), 404

    @app.errorhandler(500)
    def server_error(e):
        logger.error(f"Internal server error: {e}")
        return jsonify({"error": "internal server error"}), 500

    return app
