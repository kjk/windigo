//go:build windows

package win

import (
	"syscall"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a cursor.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hcursor
type HCURSOR HANDLE

// ⚠️ You must defer HCURSOR.DestroyCursor().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copycursor
func (hCursor HCURSOR) CopyCursor() HCURSOR {
	return (HCURSOR)(((HICON)(hCursor)).CopyIcon())
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycursor
func (hCursor HCURSOR) DestroyCursor() error {
	ret, _, err := syscall.SyscallN(proc.DestroyCursor.Addr(),
		uintptr(hCursor))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setsystemcursor
func (hCursor HCURSOR) SetSystemCursor(id co.OCR) {
	ret, _, err := syscall.SyscallN(proc.SetSystemCursor.Addr(),
		uintptr(hCursor), uintptr(id))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
