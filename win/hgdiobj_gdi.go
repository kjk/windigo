//go:build windows

package win

import (
	"syscall"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a GDI object.
//
// This type is used as the base type for the specialized GDI objects, being
// rarely used as itself.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hgdiobj
type HGDIOBJ HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hGdiObj HGDIOBJ) DeleteObject() error {
	ret, _, err := syscall.SyscallN(proc.DeleteObject.Addr(),
		uintptr(hGdiObj))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
