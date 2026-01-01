//go:build !windows
// +build !windows

package system

import "fmt"

// NewPlatform returns an error on non-Windows platforms (not yet implemented).
func NewPlatform() (Platform, error) {
	return nil, fmt.Errorf("platform not implemented for this OS")
}
