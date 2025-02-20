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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService interface {
	com.IUnknown

	// ⚠️ You must defer IUnknown.Release() on the returned object.
	//
	// Example for IMFVideoDisplayControl:
	//
	//	var gs dshow.IMFGetService // initialized somewhere
	//
	//	vdc := dshow.NewIMFVideoDisplayControl(
	//		gs.GetService(
	//			win.NewGuidFromClsid(dshowco.CLSID_MR_VideoRenderService),
	//			win.NewGuidFromIid(dshowco.IID_IMFVideoDisplayControl),
	//		),
	//	)
	//	defer vdc.Release()
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
	GetService(guidService, riid *win.GUID) com.IUnknown
}

type _IMFGetService struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMFGetService.Release().
//
// Example:
//
//	var vmr dshow.IBaseFilter // initialized somewhere
//
//	gs := dshow.NewIMFGetService(
//		vmr.QueryInterface(dshowco.IID_IMFGetService),
//	)
//	defer gs.Release()
func NewIMFGetService(base com.IUnknown) IMFGetService {
	return &_IMFGetService{IUnknown: base}
}

func (me *_IMFGetService) GetService(guidService, riid *win.GUID) com.IUnknown {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMFGetService)(unsafe.Pointer(*me.Ptr())).GetService,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(guidService)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return com.NewIUnknown(ppQueried)
	} else {
		panic(hr)
	}
}
