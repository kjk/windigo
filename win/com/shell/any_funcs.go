//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/shell/shellvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func RegisterDragDrop(hWnd win.HWND, dropTarget *shellvt.IDropTarget) {
	ret, _, _ := syscall.SyscallN(proc.RegisterDragDrop.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(&dropTarget)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-revokedragdrop
func RevokeDragDrop(hWnd win.HWND) {
	ret, _, _ := syscall.SyscallN(proc.RevokeDragDrop.Addr(),
		uintptr(hWnd))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
