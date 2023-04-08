//go:build windows

package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/dshow/dshowco"
	"github.com/kjk/windigo/win/com/dshow/dshowvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter interface {
	com.IPersist

	// Pass -1 for infinite timeout.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-getstate
	GetState(msTimeout int) (dshowco.FILTER_STATE, error)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-pause
	Pause() bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-run
	Run(start time.Duration) bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-stop
	Stop() bool
}

type _IMediaFilter struct{ com.IPersist }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaFilter.Release().
func NewIMediaFilter(base com.IUnknown) IMediaFilter {
	return &_IMediaFilter{IPersist: com.NewIPersist(base)}
}

func (me *_IMediaFilter) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	var state dshowco.FILTER_STATE
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).GetState,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(msTimeout)), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE {
		return state, hr
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Pause() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Pause,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Run(start time.Duration) bool {
	iStart := util.DurationToNano100(start)
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Run,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(iStart))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Stop() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Stop,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
