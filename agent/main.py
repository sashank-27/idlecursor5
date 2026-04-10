#!/usr/bin/env python3
"""
Activity Presence Controller - Python Agent
Main entry point for the HTTP server.
"""

import os
import signal
import sys
import logging
from dotenv import load_dotenv

from internal.api.server import create_app

# Setup logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

load_dotenv()


def get_env(key, fallback=""):
    """Get environment variable or return fallback."""
    value = os.getenv(key)
    return value if value else fallback


def main():
    # Read configuration from environment
    bind_addr = get_env("APC_BIND", "0.0.0.0:8787")
    pairing_token = os.getenv("APC_PAIRING_TOKEN", "")
    cert_file = os.getenv("APC_CERT_FILE", "")
    key_file = os.getenv("APC_KEY_FILE", "")
    allow_insecure = os.getenv("APC_ALLOW_INSECURE", "false").lower() == "true"
    cloud_mode = os.getenv("APC_CLOUD_MODE", "false").lower() == "true"

    # Parse host and port
    if ":" in bind_addr:
        host, port_str = bind_addr.rsplit(":", 1)
        port = int(port_str)
    else:
        host = bind_addr
        port = 8787

    config = {
        "host": host,
        "port": port,
        "pairing_token": pairing_token.strip(),
        "cert_file": cert_file,
        "key_file": key_file,
        "allow_insecure": allow_insecure,
        "cloud_mode": cloud_mode,
    }

    app = create_app(config)

    logger.info(f"agent listening on {bind_addr} (https if certs set)")

    # Handle graceful shutdown
    def signal_handler(signum, frame):
        logger.info("shutdown requested; draining...")
        sys.exit(0)

    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)

    # Run the server
    if cert_file and key_file:
        logger.info(f"serving HTTPS with cert {cert_file}")
        app.run(
            host=host,
            port=port,
            ssl_context=(cert_file, key_file),
            debug=False,
            threaded=True,
        )
    else:
        if not allow_insecure:
            logger.error(
                "TLS not configured; set APC_CERT_FILE/APC_KEY_FILE "
                "or APC_ALLOW_INSECURE=true for dev HTTP"
            )
            sys.exit(1)

        logger.warning(f"WARNING: serving HTTP without TLS on {bind_addr} (dev mode only)")
        app.run(host=host, port=port, debug=False, threaded=True)


if __name__ == "__main__":
    main()
