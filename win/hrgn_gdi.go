//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hrgn
type HRGN HGDIOBJ

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createellipticrgn
func CreateEllipticRgn(boundTopLeft, boundBottomRight POINT) HRGN {
	ret, _, err := syscall.SyscallN(proc.CreateEllipticRgn.Addr(),
		uintptr(boundTopLeft.X), uintptr(boundTopLeft.Y),
		uintptr(boundBottomRight.X), uintptr(boundBottomRight.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createrectrgnindirect
func CreateRectRgnIndirect(rc *RECT) HRGN {
	ret, _, err := syscall.SyscallN(proc.CreateRectRgnIndirect.Addr(),
		uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createroundrectrgn
func CreateRoundRectRgn(topLeft, bottomRight POINT, szEllipse SIZE) HRGN {
	ret, _, err := syscall.SyscallN(proc.CreateRoundRectRgn.Addr(),
		uintptr(topLeft.X), uintptr(topLeft.Y),
		uintptr(bottomRight.X), uintptr(bottomRight.Y),
		uintptr(szEllipse.Cx), uintptr(szEllipse.Cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func (hRgn HRGN) DeleteObject() error {
	return HGDIOBJ(hRgn).DeleteObject()
}

// Combines the two regions and stores the result in current region.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
func (hRgn HRGN) CombineRgn(hrgnSrc1, hrgnSrc2 HRGN, mode co.RGN) co.REGION {
	ret, _, err := syscall.SyscallN(proc.CombineRgn.Addr(),
		uintptr(hRgn), uintptr(hrgnSrc1), uintptr(hrgnSrc2), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-offsetrgn
func (hRgn HRGN) OffsetRgn(x, y int32) co.REGION {
	ret, _, err := syscall.SyscallN(proc.OffsetRgn.Addr(),
		uintptr(hRgn), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}
