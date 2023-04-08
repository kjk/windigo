//go:build windows

package wm

import (
	"github.com/kjk/windigo/win"
)

// Raw message parameters to any message: WPARAM and LPARAM.
type Any struct {
	WParam win.WPARAM
	LParam win.LPARAM
}
