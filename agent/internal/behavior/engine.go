package behavior

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"apc/internal/system"
)

// Engine drives simulated presence; it delegates to a Platform implementation.
type Engine struct {
	platform system.Platform
}

func New(p system.Platform) *Engine {
	return &Engine{platform: p}
}

// Run starts a simple loop that emits placeholder actions and performs micro moves.
func (e *Engine) Run(ctx context.Context, mode string, emit func(action string, meta map[string]string)) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var lastX, lastY int
	havePos := false
	pausedUntil := time.Time{}

	if e.platform != nil {
		if err := e.platform.PreventSleep(true); err != nil {
			emit("prevent_sleep_error", map[string]string{"error": err.Error()})
		}
		defer e.platform.PreventSleep(false)
	}

	emit("engine_start", map[string]string{"mode": mode})

	for {
		select {
		case <-ctx.Done():
			emit("engine_stop", map[string]string{"reason": "context_cancel"})
			return
		case <-ticker.C:
			// Detect user movement and pause if user is active.
			if e.platform != nil {
				if x, y, err := e.platform.CursorPos(); err == nil {
					if havePos {
						if dist(x-lastX, y-lastY) > 80 { // user activity threshold
							pausedUntil = time.Now().Add(15 * time.Second)
							emit("user_active_pause", map[string]string{"until": pausedUntil.UTC().Format(time.RFC3339)})
						}
					}
					lastX, lastY, havePos = x, y, true
				}
			}

			if !pausedUntil.IsZero() && time.Now().Before(pausedUntil) {
				continue
			}

			dx := rand.Intn(201) - 100 // range -100..100
			dy := rand.Intn(201) - 100
			if dx == 0 && dy == 0 {
				dx = 20 // ensure visible movement
			}
			if e.platform != nil {
				if err := e.platform.SimulateMove(dx, dy); err != nil {
					emit("move_error", map[string]string{"error": err.Error()})
				} else {
					// update last position after our move
					if x, y, err := e.platform.CursorPos(); err == nil {
						lastX, lastY, havePos = x, y, true
					}
				}
			}
			emit("micro_move", map[string]string{"mode": mode, "dx": intToStr(dx), "dy": intToStr(dy)})
		}
	}
}

func dist(dx, dy int) int {
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func intToStr(v int) string {
	return fmt.Sprintf("%d", v)
}
