//go:build windows

package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/com/comvt"
	"github.com/kjk/windigo/win/com/dshow/dshowco"
	"github.com/kjk/windigo/win/com/dshow/dshowvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
type IPin interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-beginflush
	BeginFlush()

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connect
	Connect(mt *AM_MEDIA_TYPE) (IPin, error)

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectedto
	ConnectedTo() (IPin, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-connectionmediatype
	ConnectionMediaType(mt *AM_MEDIA_TYPE) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-disconnect
	Disconnect()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endflush
	EndFlush()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-endofstream
	EndOfStream()

	// ⚠️ You must defer IEnumMediaTypes.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-enummediatypes
	EnumMediaTypes() (IEnumMediaTypes, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-newsegment
	NewSegment(start, stop time.Duration, rate float64)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-queryaccept
	QueryAccept(mt *AM_MEDIA_TYPE) (bool, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-querydirection
	QueryDirection() dshowco.PIN_DIRECTION

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ipin-queryid
	QueryId() string
}

type _IPin struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IPin.Release().
func NewIPin(base com.IUnknown) IPin {
	return &_IPin{IUnknown: base}
}

func (me *_IPin) BeginFlush() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).BeginFlush,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) Connect(mt *AM_MEDIA_TYPE) (IPin, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Connect,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) ConnectedTo() (IPin, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).ConnectedTo,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) ConnectionMediaType(mt *AM_MEDIA_TYPE) error {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).ConnectionMediaType,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

func (me *_IPin) Disconnect() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).Disconnect,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EndFlush() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndFlush,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EndOfStream() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EndOfStream,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) EnumMediaTypes() (IEnumMediaTypes, error) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).EnumMediaTypes,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumMediaTypes(com.NewIUnknown(ppQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IPin) NewSegment(start, stop time.Duration, rate float64) {
	iStart, iStop := util.DurationToNano100(start), util.DurationToNano100(stop)
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).NewSegment,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(iStart), uintptr(iStop), uintptr(rate))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IPin) QueryAccept(mt *AM_MEDIA_TYPE) (bool, error) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).QueryAccept,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK || hr == errco.S_FALSE {
		return hr == errco.S_OK, nil
	} else {
		return false, hr
	}
}

func (me *_IPin) QueryDirection() dshowco.PIN_DIRECTION {
	var pPinDir dshowco.PIN_DIRECTION
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).QueryDirection,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pPinDir)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pPinDir
	} else {
		panic(hr)
	}
}

func (me *_IPin) QueryId() string {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IPin)(unsafe.Pointer(*me.Ptr())).QueryId,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}
