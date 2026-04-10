"""
Behavior engine that simulates user presence with cursor movements.
Detects user activity and adapts accordingly.
"""

import logging
import random
import threading
import time

logger = logging.getLogger(__name__)


class Engine:
    """Drives simulated presence with human-like micro movements."""

    def __init__(self, platform):
        self.platform = platform

    def run(self, cancel_event, mode, emit):
        """
        Run the behavior engine until cancel_event is set.

        Args:
            cancel_event: threading.Event to signal stop
            mode: Behavior mode (e.g., "Meeting", "Build")
            emit: Callback function(action, meta) for logging
        """
        interval = 0.5  # 500ms between movements

        last_x = 0
        last_y = 0
        have_pos = False
        paused_until = None

        if self.platform:
            try:
                self.platform.prevent_sleep(True)
            except Exception as e:
                emit("prevent_sleep_error", {"error": str(e)})

        emit("engine_start", {"mode": mode})

        ticker = threading.Event()
        last_tick = time.time()

        try:
            while not cancel_event.is_set():
                now = time.time()
                elapsed = now - last_tick

                if elapsed >= interval:
                    last_tick = now

                    # Detect user movement and pause if user is active
                    if self.platform:
                        try:
                            x, y = self.platform.cursor_pos()
                            if have_pos:
                                dist = abs(x - last_x) + abs(y - last_y)
                                if dist > 80:  # User activity threshold
                                    paused_until = time.time() + 15  # Pause for 15 seconds
                                    emit(
                                        "user_active_pause",
                                        {"until": paused_until},
                                    )
                            last_x, last_y, have_pos = x, y, True
                        except Exception as e:
                            logger.debug(f"Could not get cursor position: {e}")

                    # Check if paused
                    if paused_until and time.time() < paused_until:
                        time.sleep(0.05)
                        continue

                    # Generate random movement
                    dx = random.randint(-100, 100)
                    dy = random.randint(-100, 100)

                    if dx == 0 and dy == 0:
                        dx = 20  # Ensure visible movement

                    if self.platform:
                        try:
                            self.platform.simulate_move(dx, dy)
                            # Update position after move
                            try:
                                x, y = self.platform.cursor_pos()
                                last_x, last_y, have_pos = x, y, True
                            except Exception:
                                pass
                        except Exception as e:
                            emit("move_error", {"error": str(e)})
                    else:
                        # Cloud mode - no actual movement
                        pass

                    emit("micro_move", {"mode": mode, "dx": str(dx), "dy": str(dy)})
                else:
                    time.sleep(0.01)  # Small sleep to avoid busy-waiting

        except Exception as e:
            logger.error(f"Engine error: {e}")
            emit("engine_error", {"error": str(e)})
        finally:
            if self.platform:
                try:
                    self.platform.prevent_sleep(False)
                except Exception as e:
                    logger.debug(f"Could not disable sleep prevention: {e}")

            emit("engine_stop", {"reason": "context_cancel"})
