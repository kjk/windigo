//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win"
	"github.com/kjk/windigo/win/com/com"
	"github.com/kjk/windigo/win/com/shell/shellco"
	"github.com/kjk/windigo/win/com/shell/shellvt"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-idesktopwallpaper
type IDesktopWallpaper interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-advanceslideshow
	AdvanceSlideshow(direction shellco.DSD)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-enable
	Enable(enable bool)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getbackgroundcolor
	GetBackgroundColor() win.COLORREF

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getmonitordevicepathat
	GetMonitorDevicePathAt(monitorIndex int) string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getmonitordevicepathcount
	GetMonitorDevicePathCount() int

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getmonitorrect
	GetMonitorRECT(monitorId string) win.RECT

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getposition
	GetPosition() shellco.DWPOS

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getslideshowoptions
	GetSlideshowOptions() (opts shellco.DSO, msTransition int)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getstatus
	GetStatus() shellco.DSS

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getwallpaper
	GetWallpaper(monitorId win.StrOpt) string

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-setbackgroundcolor
	SetBackgroundColor(color win.COLORREF)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-setposition
	SetPosition(position shellco.DWPOS)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-setslideshowoptions
	SetSlideshowOptions(opts shellco.DSO, msTransition int)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-setwallpaper
	SetWallpaper(monitorId win.StrOpt, imagePath string)
}

type _IDesktopWallpaper struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IDesktopWallpaper.Release().
//
// Example:
//
//	deskWp := shell.NewIDesktopWallpaper(
//		com.CoCreateInstance(
//			shellco.CLSID_DesktopWallpaper, nil,
//			comco.CLSCTX_LOCAL_SERVER,
//			shellco.IID_IDesktopWallpaper),
//	)
//	defer deskWp.Release()
func NewIDesktopWallpaper(base com.IUnknown) IDesktopWallpaper {
	return &_IDesktopWallpaper{IUnknown: base}
}

func (me *_IDesktopWallpaper) AdvanceSlideshow(direction shellco.DSD) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).AdvanceSlideshow,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(direction))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) Enable(enable bool) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).Enable,
		uintptr(unsafe.Pointer(me.Ptr())),
		util.BoolToUintptr(enable))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetBackgroundColor() win.COLORREF {
	var color win.COLORREF
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetBackgroundColor,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&color)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return color
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetMonitorDevicePathAt(monitorIndex int) string {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetMonitorDevicePathAt,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(monitorIndex), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetMonitorDevicePathCount() int {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetMonitorDevicePathCount,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&count)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(count)
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetMonitorRECT(monitorId string) win.RECT {
	var rc win.RECT
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetMonitorRECT,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(monitorId))),
		uintptr(unsafe.Pointer(&rc)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return rc
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetPosition() shellco.DWPOS {
	var pos shellco.DWPOS
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetPosition,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pos)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return pos
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetSlideshowOptions() (opts shellco.DSO, msTransition int) {
	var transition uint32
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetSlideshowOptions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&opts)), uintptr(unsafe.Pointer(&transition)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return opts, int(transition)
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetStatus() shellco.DSS {
	var status shellco.DSS
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetStatus,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&status)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return status
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) GetWallpaper(monitorId win.StrOpt) string {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).GetWallpaper,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(monitorId.Raw()), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) SetBackgroundColor(color win.COLORREF) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).SetBackgroundColor,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(color))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) SetPosition(position shellco.DWPOS) {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).SetPosition,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(position))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) SetSlideshowOptions(
	opts shellco.DSO,
	msTransition int) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).SetSlideshowOptions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(opts), uintptr(msTransition))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDesktopWallpaper) SetWallpaper(
	monitorId win.StrOpt,
	imagePath string) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDesktopWallpaper)(unsafe.Pointer(*me.Ptr())).SetWallpaper,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(monitorId.Raw()),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(imagePath))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
