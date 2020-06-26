/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HMENU HANDLE

func (hMenu HMENU) AppendMenu(uFlags co.MF, uIDNewItem uintptr,
	lpNewItem uintptr) {

	ret, _, lerr := syscall.Syscall6(proc.AppendMenu.Addr(), 4,
		uintptr(hMenu), uintptr(uFlags), uIDNewItem, uintptr(lpNewItem),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("AppendMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func CreateMenu() HMENU {
	ret, _, lerr := syscall.Syscall(proc.CreateMenu.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("CreateMenu failed: %d %s",
			lerr, lerr.Error()))
	}
	return HMENU(ret)
}

func (hMenu HMENU) DeleteMenuById(id co.ID) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uintptr(id), uintptr(co.MF_BYCOMMAND))
	if ret == 0 {
		panic(fmt.Sprintf("DeleteMeny by ID failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) DeleteMenuByPos(index uint32) {
	ret, _, lerr := syscall.Syscall(proc.DeleteMenu.Addr(), 3,
		uintptr(hMenu), uintptr(index), uintptr(co.MF_BYPOSITION))
	if ret == 0 {
		panic(fmt.Sprintf("DeleteMeny by pos failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) DestroyMenu() {
	ret, _, lerr := syscall.Syscall(proc.DestroyMenu.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("DestroyMenu failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) EnableMenuItem(uIDEnableItem uint32, uEnable co.MF) {
	syscall.Syscall(proc.EnableMenuItem.Addr(), 3,
		uintptr(hMenu), uintptr(uIDEnableItem), uintptr(uEnable))
}

func (hMenu HMENU) GetMenuItemCount() uint32 {
	ret, _, lerr := syscall.Syscall(proc.GetMenuItemCount.Addr(), 1,
		uintptr(hMenu), 0, 0)
	if int32(ret) == -1 {
		panic(fmt.Sprintf("GetItemCount failed: %d %s",
			lerr, lerr.Error()))
	}
	return uint32(ret)
}

func (hMenu HMENU) GetMenuInfo(mi *MENUINFO) {
	mi.CbSize = uint32(unsafe.Sizeof(*mi)) // safety

	ret, _, lerr := syscall.Syscall(proc.GetMenuInfo.Addr(), 2,
		uintptr(hMenu), uintptr(unsafe.Pointer(mi)), 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuInfo failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetMenuItemID(index uint32) co.ID {
	ret, _, _ := syscall.Syscall(proc.GetMenuItemID.Addr(), 2,
		uintptr(hMenu), uintptr(index), 0)
	return co.ID(ret)
}

func (hMenu HMENU) GetMenuItemInfoById(id co.ID, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(id), 0, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuItemInfo by ID failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetMenuItemInfoByPos(index uint32, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.GetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(index), 1, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("GetMenuItemInfo by pos failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) GetSubMenu(nPos uint32) HMENU {
	ret, _, _ := syscall.Syscall(proc.GetSubMenu.Addr(), 2,
		uintptr(hMenu), uintptr(nPos), 0)
	return HMENU(ret)
}

func (hMenu HMENU) SetMenuItemInfoById(id co.ID, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(id), 0, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo by ID failed: %d %s",
			lerr, lerr.Error()))
	}
}

func (hMenu HMENU) SetMenuItemInfoByPos(index uint32, mii *MENUITEMINFO) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii)) // safety

	ret, _, lerr := syscall.Syscall6(proc.SetMenuItemInfo.Addr(), 4,
		uintptr(hMenu), uintptr(index), 1, uintptr(unsafe.Pointer(mii)),
		0, 0)
	if ret == 0 {
		panic(fmt.Sprintf("SetMenuItemInfo by pos failed: %d %s",
			lerr, lerr.Error()))
	}
}