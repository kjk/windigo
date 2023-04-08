//go:build windows

package win

import (
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rc *RECT, style co.WS, hasMenu bool, exStyle co.WS_EX) {
	ret, _, err := syscall.SyscallN(proc.AdjustWindowRectEx.Addr(),
		uintptr(unsafe.Pointer(rc)), uintptr(style),
		util.BoolToUintptr(hasMenu), uintptr(exStyle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-allowsetforegroundwindow
func AllowSetForegroundWindow(processId uint32) {
	ret, _, err := syscall.SyscallN(proc.AllowSetForegroundWindow.Addr(),
		uintptr(processId))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
func BroadcastSystemMessage(
	flags co.BSF,
	recipients co.BSM,
	msg co.WM,
	wParam WPARAM,
	lParam LPARAM) (broadcastSuccessful bool, receivers co.BSM, e error) {

	receivers = recipients

	ret, _, err := syscall.SyscallN(proc.BroadcastSystemMessage.Addr(),
		uintptr(flags), uintptr(unsafe.Pointer(&receivers)),
		uintptr(msg), uintptr(wParam), uintptr(lParam))

	broadcastSuccessful = int(ret) > 1
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
}

// This function creates HCURSOR only. The HICON variation is
// CreateIconFromResourceEx().
//
// ⚠️ You must defer HCURSOR.DestroyCursor().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateCursorFromResourceEx(
	resBits []byte, fmtVersion int,
	cxDesired, cyDesired int,
	flags co.LR) (HCURSOR, error) {

	hIcon, err := CreateIconFromResourceEx(
		resBits, fmtVersion, cxDesired, cyDesired, flags)
	return HCURSOR(hIcon), err
}

// This function creates HICON only. The HCURSOR variation is
// CreateCursorFromResourceEx().
//
// ⚠️ You must defer HICON.DestroyIcon().
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateIconFromResourceEx(
	resBits []byte, fmtVersion int,
	cxDesired, cyDesired int,
	flags co.LR) (HICON, error) {

	ret, _, err := syscall.SyscallN(proc.CreateIconFromResourceEx.Addr(),
		uintptr(unsafe.Pointer(&resBits[0])), uintptr(len(resBits)),
		1, uintptr(fmtVersion), uintptr(cxDesired), uintptr(cyDesired),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), errco.ERROR(err)
	}
	return HICON(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroycarret
func DestroyCaret() error {
	ret, _, err := syscall.SyscallN(proc.DestroyCaret.Addr())
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.SyscallN(proc.DispatchMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-endmenu
func EndMenu() {
	ret, _, err := syscall.SyscallN(proc.EndMenu.Addr())
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevices(
	device StrOpt, flags co.EDD,
	callback func(devNum int, info *DISPLAY_DEVICE) bool) {

	devicePtr := device.Raw()
	devNum := 0

	dide := DISPLAY_DEVICE{}
	dide.SetCb()

	for {
		ret, _, err := syscall.SyscallN(proc.EnumDisplayDevices.Addr(),
			uintptr(devicePtr), uintptr(devNum),
			uintptr(unsafe.Pointer(&dide)), uintptr(flags))

		if ret == 0 {
			if wErr := errco.ERROR(err); wErr != errco.SUCCESS {
				panic(wErr)
			} else {
				break
			}
		}

		if !callback(devNum, &dide) {
			break
		}
		devNum++
	}

	runtime.KeepAlive(devicePtr)
}

// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(hWnd HWND) bool) {
	pPack := &_EnumWindowsPack{f: callback}
	_globalEnumWindowsMutex.Lock()
	if _globalEnumWindowsFuncs == nil { // the set was not initialized yet?
		_globalEnumWindowsFuncs = make(map[*_EnumWindowsPack]struct{}, 1)
	}
	_globalEnumWindowsFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumWindowsMutex.Unlock()

	ret, _, err := syscall.SyscallN(proc.EnumWindows.Addr(),
		_globalEnumWindowsCallback, uintptr(unsafe.Pointer(pPack)))

	_globalEnumWindowsMutex.Lock()
	delete(_globalEnumWindowsFuncs, pPack) // remove from the set
	_globalEnumWindowsMutex.Unlock()

	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _EnumWindowsPack struct{ f func(hWnd HWND) bool }

var (
	_globalEnumWindowsFuncs    map[*_EnumWindowsPack]struct{}
	_globalEnumWindowsMutex    = sync.Mutex{}
	_globalEnumWindowsCallback = syscall.NewCallback(
		func(hWnd HWND, lParam LPARAM) uintptr {
			pPack := (*_EnumWindowsPack)(unsafe.Pointer(lParam))
			return util.BoolToUintptr(pPack.f(hWnd))
		})
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(virtKeyCode co.VK) uint16 {
	ret, _, _ := syscall.SyscallN(proc.GetAsyncKeyState.Addr(),
		uintptr(virtKeyCode))
	return uint16(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcaretpos
func GetCaretPos() RECT {
	var rc RECT
	ret, _, err := syscall.SyscallN(proc.GetCaretPos.Addr(),
		uintptr(unsafe.Pointer(&rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return rc
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdialogbaseunits
func GetDialogBaseUnits() (horz, vert uint16) {
	ret, _, _ := syscall.SyscallN(proc.GetDialogBaseUnits.Addr())
	horz, vert = LOWORD(uint32(ret)), HIWORD(uint32(ret))
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getguithreadinfo
func GetGUIThreadInfo(thread_id uint32, info *GUITHREADINFO) {
	info.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(proc.GetGUIThreadInfo.Addr(),
		uintptr(thread_id), uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getinputstate
func GetInputState() bool {
	ret, _, _ := syscall.SyscallN(proc.GetInputState.Addr())
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagew
func GetMessage(
	msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) (int32, error) {

	ret, _, err := syscall.SyscallN(proc.GetMessage.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax))
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return int32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessageextrainfo
func GetMessageExtraInfo() LPARAM {
	ret, _, _ := syscall.SyscallN(proc.GetMessageExtraInfo.Addr())
	return LPARAM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagepos
func GetMessagePos() POINT {
	ret, _, _ := syscall.SyscallN(proc.GetMessagePos.Addr())
	return POINT{
		X: int32(LOWORD(uint32(ret))),
		Y: int32(HIWORD(uint32(ret))),
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessagetime
func GetMessageTime() time.Duration {
	ret, _, _ := syscall.SyscallN(proc.GetMessageTime.Addr())
	return time.Duration(ret * uintptr(time.Millisecond))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getphysicalcursorpos
func GetPhysicalCursorPos() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetPhysicalCursorPos.Addr(),
		uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getprocessdefaultlayout
func GetProcessDefaultLayout() co.LAYOUT {
	var defaultLayout co.LAYOUT
	ret, _, err := syscall.SyscallN(proc.GetProcessDefaultLayout.Addr(),
		uintptr(unsafe.Pointer(&defaultLayout)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return defaultLayout
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getqueuestatus
func GetQueueStatus(flags co.QS) uint32 {
	ret, _, _ := syscall.SyscallN(proc.GetQueueStatus.Addr(),
		uintptr(flags))
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
func GetSysColor(index co.COLOR) COLORREF {
	ret, _, _ := syscall.SyscallN(proc.GetSysColor.Addr(),
		uintptr(index))
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index co.SM) int32 {
	ret, _, _ := syscall.SyscallN(proc.GetSystemMetrics.Addr(),
		uintptr(index))
	return int32(ret)
}

// Available in Windows 10, version 1607.
//
// 📑 https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetricsfordpi
func GetSystemMetricsForDpi(index co.SM, dpi uint32) int32 {
	ret, _, err := syscall.SyscallN(proc.GetSystemMetricsForDpi.Addr(),
		uintptr(index))
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessage
func InSendMessage() bool {
	ret, _, _ := syscall.SyscallN(proc.InSendMessage.Addr())
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-insendmessageex
func InSendMessageEx() co.ISMEX {
	ret, _, _ := syscall.SyscallN(proc.InSendMessageEx.Addr())
	return co.ISMEX(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isguithread
func IsGUIThread(convertToGuiThread bool) (bool, error) {
	ret, _, _ := syscall.SyscallN(proc.IsGUIThread.Addr(),
		util.BoolToUintptr(convertToGuiThread))
	if convertToGuiThread && errco.ERROR(ret) == errco.NOT_ENOUGH_MEMORY {
		return false, errco.NOT_ENOUGH_MEMORY
	}
	return ret != 0, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-locksetforegroundwindow
func LockSetForegroundWindow(lockCode co.LSFW) {
	ret, _, err := syscall.SyscallN(proc.LockSetForegroundWindow.Addr(),
		uintptr(lockCode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(
	msg *MSG, hWnd HWND,
	msgFilterMin, msgFilterMax co.WM, removeMsg co.PM) bool {

	ret, _, _ := syscall.SyscallN(proc.PeekMessage.Addr(),
		uintptr(unsafe.Pointer(msg)), uintptr(hWnd),
		uintptr(msgFilterMin), uintptr(msgFilterMax), uintptr(removeMsg))
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	syscall.SyscallN(proc.PostQuitMessage.Addr(),
		uintptr(exitCode))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(
	idThread uint32, msg co.WM, wParam WPARAM, lParam LPARAM) error {

	ret, _, err := syscall.SyscallN(proc.PostThreadMessage.Addr(),
		uintptr(idThread), uintptr(msg), uintptr(wParam), uintptr(lParam))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wcx *WNDCLASSEX) (ATOM, error) {
	wcx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(proc.RegisterClassEx.Addr(),
		uintptr(unsafe.Pointer(wcx)))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return ATOM(0), wErr
	} else {
		return ATOM(ret), nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(message string) (co.WM, error) {
	ret, _, err := syscall.SyscallN(proc.RegisterWindowMessage.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(message))))

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return co.WM(0), wErr
	} else {
		return co.WM(ret), nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-replymessage
func ReplyMessage(result uintptr) bool {
	ret, _, _ := syscall.SyscallN(proc.ReplyMessage.Addr(),
		result)
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.SyscallN(proc.TranslateMessage.Addr(),
		uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// Available in Windows 10, version 1703.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value co.DPI_AWARE_CTX) error {
	ret, _, err := syscall.SyscallN(proc.SetProcessDpiAwarenessContext.Addr(),
		uintptr(value))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// Available in Windows Vista.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() {
	ret, _, _ := syscall.SyscallN(proc.SetProcessDPIAware.Addr())
	if ret == 0 {
		panic("SetProcessDPIAware() failed.")
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setmessageextrainfo
func SetMessageExtraInfo(lp LPARAM) LPARAM {
	ret, _, _ := syscall.SyscallN(proc.SetMessageExtraInfo.Addr(),
		uintptr(lp))
	return LPARAM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdefaultlayout
func SetProcessDefaultLayout(defaultLayout co.LAYOUT) {
	ret, _, err := syscall.SyscallN(proc.SetProcessDefaultLayout.Addr(),
		uintptr(defaultLayout))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(className ClassName, hInst HINSTANCE) error {
	classNameVal, classNameBuf := className.raw()
	ret, _, err := syscall.SyscallN(proc.UnregisterClass.Addr(),
		classNameVal, uintptr(hInst))
	runtime.KeepAlive(classNameBuf)

	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		return wErr
	} else {
		return nil
	}
}
