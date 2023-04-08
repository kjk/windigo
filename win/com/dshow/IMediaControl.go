//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/autom"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/dshow/dshowco"
	"github.com/kjk/windigo/win/com/dshow/dshowvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
type IMediaControl interface {
	autom.IDispatch

	// Pass -1 for infinite timeout.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-getstate
	GetState(msTimeout int) (dshowco.FILTER_STATE, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-pause
	Pause() bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-renderfile
	RenderFile(fileName string)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-run
	Run() bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stop
	Stop()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-imediacontrol-stopwhenready
	StopWhenReady() bool
}

type _IMediaControl struct{ autom.IDispatch }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaControl.Release().
//
// Example:
//
//	var gb dshow.IGraphBuilder // initialized somewhere
//
//	mc := dshow.NewIMediaControl(
//		gb.QueryInterface(dshowco.IID_IMediaControl),
//	)
//	defer mc.Release()
func NewIMediaControl(base com.IUnknown) IMediaControl {
	return &_IMediaControl{IDispatch: autom.NewIDispatch(base)}
}

func (me *_IMediaControl) GetState(
	msTimeout int) (dshowco.FILTER_STATE, error) {

	var state dshowco.FILTER_STATE
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).GetState,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(msTimeout)), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE || hr == errco.VFW_S_CANT_CUE {
		return state, hr
	} else {
		panic(hr)
	}
}

func (me *_IMediaControl) Pause() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Pause,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaControl) RenderFile(fileName string) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).RenderFile,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IMediaControl) Run() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Run,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaControl) Stop() {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).Stop,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IMediaControl) StopWhenReady() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaControl)(unsafe.Pointer(*me.Ptr())).StopWhenReady,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
