//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/com/comvt"
	"github.com/kjk/windigo/win/com/dshow/dshowvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
type IFilterGraph interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-addfilter
	AddFilter(filter IBaseFilter, name string) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-connectdirect
	ConnectDirect(pinOut, pinIn IPin, mt *AM_MEDIA_TYPE)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-disconnect
	Disconnect(pin IPin)

	// ⚠️ You must defer IEnumFilters.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-enumfilters
	EnumFilters() IEnumFilters

	// ⚠️ You must defer IBaseFilter.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-findfilterbyname
	FindFilterByName(name string) (IBaseFilter, bool)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-reconnect
	Reconnect(pin IPin)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-removefilter
	RemoveFilter(filter IBaseFilter)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifiltergraph-setdefaultsyncsource
	SetDefaultSyncSource()
}

type _IFilterGraph struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFilterGraph.Release().
//
// Example:
//
//	fg := dshow.NewIFilterGraph(
//		com.CoCreateInstance(
//			dshowco.CLSID_FilterGraph, nil,
//			comco.CLSCTX_INPROC_SERVER,
//			dshowco.IID_IFilterGraph),
//	)
//	defer fg.Release()
func NewIFilterGraph(base com.IUnknown) IFilterGraph {
	return &_IFilterGraph{IUnknown: base}
}

func (me *_IFilterGraph) AddFilter(filter IBaseFilter, name string) error {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).AddFilter,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(filter.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(name))))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

func (me *_IFilterGraph) ConnectDirect(pinOut, pinIn IPin, mt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).AddFilter,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pinOut.Ptr())),
		uintptr(unsafe.Pointer(pinIn.Ptr())),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFilterGraph) Disconnect(pin IPin) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).Disconnect,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pin.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFilterGraph) EnumFilters() IEnumFilters {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).EnumFilters,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumFilters(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IFilterGraph) FindFilterByName(name string) (IBaseFilter, bool) {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).FindFilterByName,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIBaseFilter(com.NewIUnknown(ppvQueried)), true
	} else if hr == errco.VFW_E_NOT_FOUND {
		return nil, false
	} else {
		panic(hr)
	}
}

func (me *_IFilterGraph) Reconnect(pin IPin) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).Reconnect,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(pin.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFilterGraph) RemoveFilter(filter IBaseFilter) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).RemoveFilter,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(filter.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IFilterGraph) SetDefaultSyncSource() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IFilterGraph)(unsafe.Pointer(*me.Ptr())).SetDefaultSyncSource,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
