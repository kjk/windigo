//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmenablemmcss
func DwmEnableMMCSS(enable bool) error {
	ret, _, _ := syscall.SyscallN(proc.DwmEnableMMCSS.Addr(),
		util.BoolToUintptr(enable))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmflush
func DwmFlush() error {
	ret, _, _ := syscall.SyscallN(proc.DwmFlush.Addr())
	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmgetcolorizationcolor
func DwmGetColorizationColor() (color COLORREF, isOpaqueBlend bool) {
	bOpaqueBlend := int32(util.BoolToUintptr(isOpaqueBlend)) // BOOL
	ret, _, _ := syscall.SyscallN(proc.DwmGetColorizationColor.Addr(),
		uintptr(unsafe.Pointer(&color)), uintptr(unsafe.Pointer(&bOpaqueBlend)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	isOpaqueBlend = bOpaqueBlend != 0
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmiscompositionenabled
func DwmIsCompositionEnabled() bool {
	var pfEnabled int32 // BOOL
	ret, _, _ := syscall.SyscallN(proc.DwmIsCompositionEnabled.Addr(),
		uintptr(unsafe.Pointer(&pfEnabled)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return pfEnabled != 0
}
