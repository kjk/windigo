//go:build windows

package win

import (
	"syscall"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(index co.COLOR) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.GetSysColorBrush.Addr(),
		uintptr(index))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}
