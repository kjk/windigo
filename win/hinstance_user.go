//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (hInst HINSTANCE) CreateDialogParam(
	templateName ResId, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) HWND {

	templateNameVal, templateNameBuf := templateName.raw()
	ret, _, err := syscall.SyscallN(proc.CreateDialogParam.Addr(),
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam))
	runtime.KeepAlive(templateNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HWND(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dialogboxparamw
func (hInst HINSTANCE) DialogBoxParam(
	templateName ResId, hwndParent HWND,
	dialogFunc uintptr, dwInitParam LPARAM) uintptr {

	templateNameVal, templateNameBuf := templateName.raw()
	ret, _, err := syscall.SyscallN(proc.DialogBoxParam.Addr(),
		uintptr(hInst), templateNameVal,
		uintptr(hwndParent), dialogFunc, uintptr(dwInitParam))
	runtime.KeepAlive(templateNameBuf)
	if int(ret) == -1 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(
	className *uint16, destBuf *WNDCLASSEX) (ATOM, error) {

	ret, _, err := syscall.SyscallN(proc.GetClassInfoEx.Addr(),
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		return ATOM(0), errco.ERROR(err)
	}
	return ATOM(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (hInst HINSTANCE) LoadAccelerators(tableName ResId) HACCEL {
	tableNameVal, tableNameBuf := tableName.raw()
	ret, _, err := syscall.SyscallN(proc.LoadAccelerators.Addr(),
		uintptr(hInst), tableNameVal)
	runtime.KeepAlive(tableNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HACCEL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(cursorName CursorRes) HCURSOR {
	cursorNameVal, cursorNameBuf := cursorName.raw()
	ret, _, err := syscall.SyscallN(proc.LoadCursor.Addr(),
		uintptr(hInst), cursorNameVal)
	runtime.KeepAlive(cursorNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HCURSOR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(iconName IconRes) HICON {
	iconNameVal, iconNameBuf := iconName.raw()
	ret, _, err := syscall.SyscallN(proc.LoadIcon.Addr(),
		uintptr(hInst), iconNameVal)
	runtime.KeepAlive(iconNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// Returned HGDIOBJ must be cast into HBITMAP, HCURSOR or HICON.
//
// ⚠️ If the object is not being loaded from the application resources, you must
// defer its respective DeleteObject().
//
// Example loading a 16x16 icon resource:
//
//	const MY_ICON_ID int = 101
//
//	hIcon := win.HICON(
//		win.GetModuleHandle(win.StrOptNone()).LoadImage(
//			win.ResIdInt(MY_ICON_ID),
//			co.IMAGE_ICON,
//			16, 16,
//			co.LR_DEFAULTCOLOR,
//		),
//	)
//
// Example loading a bitmap from file:
//
//	hBmp := win.HBITMAP(
//		win.HINSTANCE(0).LoadImage(
//			win.ResIdStr("C:\\Temp\\image.bmp"),
//			co.IMAGE_BITMAP,
//			0, 0,
//			co.LR_LOADFROMFILE,
//		),
//	)
//	defer hBmp.DeleteObject()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
func (hInst HINSTANCE) LoadImage(
	name ResId, imgType co.IMAGE, cx, cy int32, fuLoad co.LR) HGDIOBJ {

	nameVal, nameBuf := name.raw()
	ret, _, err := syscall.SyscallN(proc.LoadImage.Addr(),
		uintptr(hInst), nameVal, uintptr(imgType),
		uintptr(cx), uintptr(cy), uintptr(fuLoad))
	runtime.KeepAlive(nameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGDIOBJ(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(menuName ResId) HMENU {
	menuNameVal, menuNameBuf := menuName.raw()
	ret, _, err := syscall.SyscallN(proc.LoadMenu.Addr(),
		uintptr(hInst), menuNameVal)
	runtime.KeepAlive(menuNameBuf)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HMENU(ret)
}
