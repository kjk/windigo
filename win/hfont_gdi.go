//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a font.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hfont
type HFONT HGDIOBJ

// ⚠️ You must defer HFONT.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(lf *LOGFONT) HFONT {
	ret, _, err := syscall.SyscallN(proc.CreateFontIndirect.Addr(),
		uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hFont HFONT) DeleteObject() error {
	return HGDIOBJ(hFont).DeleteObject()
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getobject
func (hFont HFONT) GetObject(lf *LOGFONT) {
	ret, _, err := syscall.SyscallN(proc.GetObject.Addr(),
		uintptr(hFont), unsafe.Sizeof(*lf), uintptr(unsafe.Pointer(lf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
