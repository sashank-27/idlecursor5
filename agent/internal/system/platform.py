"""
Windows platform implementation for cursor manipulation and sleep prevention.
Uses ctypes to access Win32 APIs.
"""

import ctypes
import logging
import sys

logger = logging.getLogger(__name__)


class Platform:
    """Windows implementation using Win32 APIs."""

    def __init__(self):
        if sys.platform != "win32":
            raise RuntimeError("Windows platform required")

        # Load necessary Windows DLLs
        self.user32 = ctypes.windll.user32
        self.kernel32 = ctypes.windll.kernel32

        # Win32 constants
        self.INPUT_MOUSE = 0
        self.MOUSEEVENTF_MOVE = 0x0001
        self.ES_CONTINUOUS = 0x80000000
        self.ES_SYSTEM_REQUIRED = 0x00000001
        self.ES_DISPLAY_REQUIRED = 0x00000002

    def simulate_move(self, dx, dy):
        """
        Simulate a relative mouse move.

        Args:
            dx: Horizontal offset
            dy: Vertical offset
        """
        # MOUSEINPUT structure
        class MOUSEINPUT(ctypes.Structure):
            _fields_ = [
                ("dx", ctypes.c_long),
                ("dy", ctypes.c_long),
                ("mouseData", ctypes.c_ulong),
                ("dwFlags", ctypes.c_ulong),
                ("time", ctypes.c_ulong),
                ("dwExtraInfo", ctypes.c_void_p),
            ]

        # INPUT structure
        class INPUT(ctypes.Structure):
            _fields_ = [("type", ctypes.c_ulong), ("mi", MOUSEINPUT)]

        # Create input event
        input_event = INPUT()
        input_event.type = self.INPUT_MOUSE
        input_event.mi.dx = int(dx)
        input_event.mi.dy = int(dy)
        input_event.mi.dwFlags = self.MOUSEEVENTF_MOVE

        # Call SendInput
        result = self.user32.SendInput(1, ctypes.byref(input_event), ctypes.sizeof(INPUT))
        if result == 0:
            raise RuntimeError("SendInput failed")

    def prevent_sleep(self, enable):
        """
        Enable or disable sleep prevention.

        Args:
            enable: True to keep system awake, False to allow sleep
        """
        if enable:
            flags = (
                self.ES_CONTINUOUS
                | self.ES_SYSTEM_REQUIRED
                | self.ES_DISPLAY_REQUIRED
            )
        else:
            flags = self.ES_CONTINUOUS

        result = self.kernel32.SetThreadExecutionState(ctypes.c_ulong(flags))
        if result == 0:
            raise RuntimeError("SetThreadExecutionState failed")

    def cursor_pos(self):
        """
        Get current cursor position.

        Returns:
            (x, y) tuple of current cursor coordinates
        """

        class POINT(ctypes.Structure):
            _fields_ = [("x", ctypes.c_long), ("y", ctypes.c_long)]

        pt = POINT()
        result = self.user32.GetCursorPos(ctypes.byref(pt))
        if result == 0:
            raise RuntimeError("GetCursorPos failed")

        return int(pt.x), int(pt.y)
