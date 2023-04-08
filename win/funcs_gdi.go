//go:build windows

package win

import (
	"syscall"

	"github.com/kjk/windigo/internal/proc"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := syscall.SyscallN(proc.GdiFlush.Addr())
	return ret == 0
}
