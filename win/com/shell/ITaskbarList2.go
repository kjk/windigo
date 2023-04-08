//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/shell/shellvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 interface {
	ITaskbarList

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist2-markfullscreenwindow
	MarkFullscreenWindow(hwnd win.HWND, fullScreen bool)
}

type _ITaskbarList2 struct{ ITaskbarList }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ITaskbarList2.Release().
//
// Example:
//
//	taskbl := shell.NewITaskbarList2(
//		com.CoCreateInstance(
//			shellco.CLSID_TaskbarList, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			shellco.IID_ITaskbarList2),
//	)
//	defer taskbl.Release()
func NewITaskbarList2(base com.IUnknown) ITaskbarList2 {
	return &_ITaskbarList2{ITaskbarList: NewITaskbarList(base)}
}

func (me *_ITaskbarList2) MarkFullscreenWindow(hwnd win.HWND, fullScreen bool) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.ITaskbarList2)(unsafe.Pointer(*me.Ptr())).MarkFullscreenWindow,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwnd), util.BoolToUintptr(fullScreen))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
