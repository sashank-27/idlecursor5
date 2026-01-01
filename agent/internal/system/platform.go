package system

// Platform contains placeholders for OS-specific hooks. Implementations will
// be added per-OS to keep the core behavior portable.
type Platform interface {
	SimulateMove(dx, dy int) error
	SimulateKey(code int) error
	PreventSleep(enable bool) error
	CursorPos() (x int, y int, err error)
}
