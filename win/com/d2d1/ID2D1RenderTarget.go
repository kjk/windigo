//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/d2d1/d2d1vt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type ID2D1RenderTarget interface {
	ID2D1Resource

	// ⚠️ You must defer ID2D1RenderTarget.EndDraw().
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-begindraw
	BeginDraw()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-enddraw
	EndDraw() (tag1, tag2 uint64)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-flush
	Flush() (tag1, tag2 uint64)
}

type _ID2D1RenderTarget struct{ ID2D1Resource }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1RenderTarget.Release().
func NewID2D1RenderTarget(base com.IUnknown) ID2D1RenderTarget {
	return &_ID2D1RenderTarget{ID2D1Resource: NewID2D1Resource(base)}
}

func (me *_ID2D1RenderTarget) BeginDraw() {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).BeginDraw,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ID2D1RenderTarget) EndDraw() (tag1, tag2 uint64) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).EndDraw,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}

func (me *_ID2D1RenderTarget) Flush() (tag1, tag2 uint64) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).Flush,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}
