//go:build windows
// +build windows

package system

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// NewPlatform returns a Windows implementation of Platform.
func NewPlatform() (Platform, error) {
	return &winPlatform{}, nil
}

// winPlatform uses Win32 APIs to simulate input and prevent sleep.
type winPlatform struct{}

// SimulateMove performs a relative mouse move via SendInput.
func (p *winPlatform) SimulateMove(dx, dy int) error {
	var input INPUT
	input.Type = INPUT_MOUSE
	input.Mi = MOUSEINPUT{
		Dx:      int32(dx),
		Dy:      int32(dy),
		DwFlags: MOUSEEVENTF_MOVE,
	}

	n, _, err := procSendInput.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&input)),
		uintptr(unsafe.Sizeof(input)),
	)
	if n == 0 {
		return fmt.Errorf("SendInput failed: %v", err)
	}
	return nil
}

// SimulateKey is not implemented yet; return nil to avoid failure.
func (p *winPlatform) SimulateKey(code int) error {
	// TODO: implement key simulation if needed.
	return nil
}

// PreventSleep toggles system/display required flags to keep the system awake.
func (p *winPlatform) PreventSleep(enable bool) error {
	var flags uint32
	if enable {
		flags = ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED
	} else {
		flags = ES_CONTINUOUS
	}
	r, _, err := procSetThreadExecutionState.Call(uintptr(flags))
	if r == 0 {
		return fmt.Errorf("SetThreadExecutionState failed: %v", err)
	}
	return nil
}

// CursorPos returns the current cursor position.
func (p *winPlatform) CursorPos() (int, int, error) {
	var pt POINT
	r, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if r == 0 {
		return 0, 0, fmt.Errorf("GetCursorPos failed: %v", err)
	}
	return int(pt.X), int(pt.Y), nil
}

// Win32 bindings.

const (
	INPUT_MOUSE         = 0
	MOUSEEVENTF_MOVE    = 0x0001
	ES_CONTINUOUS       = 0x80000000
	ES_SYSTEM_REQUIRED  = 0x00000001
	ES_DISPLAY_REQUIRED = 0x00000002
)

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

type POINT struct {
	X int32
	Y int32
}

var (
	user32                      = windows.NewLazySystemDLL("user32.dll")
	kernel32                    = windows.NewLazySystemDLL("kernel32.dll")
	procSendInput               = user32.NewProc("SendInput")
	procSetThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")
	procGetCursorPos            = user32.NewProc("GetCursorPos")
)
