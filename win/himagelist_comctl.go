//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// A handle to an image list.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/image-lists
type HIMAGELIST HANDLE

// Usually flags is ILC_COLOR32.
//
// ⚠️ You must defer HIMAGELIST.Destroy().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageListCreate(
	cx, cy uint32, flags co.ILC, szInitial, szGrow uint32) HIMAGELIST {

	ret, _, err := syscall.SyscallN(proc.ImageList_Create.Addr(),
		uintptr(cx), uintptr(cy), uintptr(flags),
		uintptr(szInitial), uintptr(szGrow))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HIMAGELIST(ret)
}

// If icon was loaded from resource with LoadIcon(), it doesn't need to be
// destroyed, because all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
func (hImg HIMAGELIST) AddIcon(hIcons ...HICON) {
	for _, hIco := range hIcons {
		hImg.ReplaceIcon(-1, hIco)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_destroy
func (hImg HIMAGELIST) Destroy() error {
	// http://www.catch22.net/tuts/win32/system-image-list
	// https://www.autohotkey.com/docs/commands/ListView.htm
	ret, _, err := syscall.SyscallN(proc.ImageList_Destroy.Addr(),
		uintptr(hImg))
	if ret == 0 && errco.ERROR(err) != errco.SUCCESS {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_geticonsize
func (hImg HIMAGELIST) GetIconSize() SIZE {
	var sz SIZE
	ret, _, err := syscall.SyscallN(proc.ImageList_GetIconSize.Addr(),
		uintptr(hImg),
		uintptr(unsafe.Pointer(&sz.Cx)), uintptr(unsafe.Pointer(&sz.Cy)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_getimagecount
func (hImg HIMAGELIST) GetImageCount() uint32 {
	ret, _, _ := syscall.SyscallN(proc.ImageList_GetImageCount.Addr(),
		uintptr(hImg))
	return uint32(ret)
}

// If icon was loaded from resource with LoadIcon(), it doesn't need to be
// destroyed, because all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_replaceicon
func (hImg HIMAGELIST) ReplaceIcon(i int32, hIcon HICON) int32 {
	ret, _, err := syscall.SyscallN(proc.ImageList_ReplaceIcon.Addr(),
		uintptr(hImg), uintptr(i), uintptr(hIcon))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}
