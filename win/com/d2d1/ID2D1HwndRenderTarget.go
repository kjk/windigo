//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/d2d1/d2d1co"
	"github.com/kjk/windigo/win/com/d2d1/d2d1vt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1hwndrendertarget
type ID2D1HwndRenderTarget interface {
	ID2D1RenderTarget

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-checkwindowstate
	CheckWindowState() d2d1co.WINDOW_STATE

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-gethwnd
	GetHwnd() win.HWND

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-resize(constd2d1_size_u)
	Resize(pixelSize SIZE_U)
}

type _ID2D1HwndRenderTarget struct{ ID2D1RenderTarget }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1HwndRenderTarget.Release().
func NewID2D1HwndRenderTarget(base com.IUnknown) ID2D1HwndRenderTarget {
	return &_ID2D1HwndRenderTarget{ID2D1RenderTarget: NewID2D1RenderTarget(base)}
}

func (me *_ID2D1HwndRenderTarget) CheckWindowState() d2d1co.WINDOW_STATE {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1HwndRenderTarget)(unsafe.Pointer(*me.Ptr())).GetHwnd,
		uintptr(unsafe.Pointer(me.Ptr())))
	return d2d1co.WINDOW_STATE(ret)
}

func (me *_ID2D1HwndRenderTarget) GetHwnd() win.HWND {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1HwndRenderTarget)(unsafe.Pointer(*me.Ptr())).GetHwnd,
		uintptr(unsafe.Pointer(me.Ptr())))
	return win.HWND(ret)
}

func (me *_ID2D1HwndRenderTarget) Resize(pixelSize SIZE_U) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1HwndRenderTarget)(unsafe.Pointer(*me.Ptr())).Resize,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pixelSize)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
