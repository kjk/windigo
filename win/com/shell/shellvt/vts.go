//go:build windows

package shellvt

import (
	"github.com/kjk/windigo/win/com/com/comvt"
)

// IDataObject virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-idataobject
type IDataObject struct {
	comvt.IUnknown
	GetData               uintptr
	GetDataHere           uintptr
	QueryGetData          uintptr
	GetCanonicalFormatEtc uintptr
	SetData               uintptr
	EnumFormatEtc         uintptr
	DAdvise               uintptr
	DUnadvise             uintptr
	EnumDAdvise           uintptr
}

// IDesktopWallpaper virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-idesktopwallpaper
type IDesktopWallpaper struct {
	comvt.IUnknown
	SetWallpaper              uintptr
	GetWallpaper              uintptr
	GetMonitorDevicePathAt    uintptr
	GetMonitorDevicePathCount uintptr
	GetMonitorRECT            uintptr
	SetBackgroundColor        uintptr
	GetBackgroundColor        uintptr
	SetPosition               uintptr
	GetPosition               uintptr
	SetSlideshow              uintptr
	GetSlideshow              uintptr
	SetSlideshowOptions       uintptr
	GetSlideshowOptions       uintptr
	AdvanceSlideshow          uintptr
	GetStatus                 uintptr
	Enable                    uintptr
}

// IDropTarget virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-idroptarget
type IDropTarget struct {
	comvt.IUnknown
	DragEnter uintptr
	DragOver  uintptr
	DragLeave uintptr
	Drop      uintptr
}

// IFileDialog virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifiledialog
type IFileDialog struct {
	IModalWindow
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
}

// IFileOpenDialog virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifileopendialog
type IFileOpenDialog struct {
	IFileDialog
	GetResults       uintptr
	GetSelectedItems uintptr
}

// IFileSaveDialog virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ifilesavedialog
type IFileSaveDialog struct {
	IFileDialog
	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

// IModalWindow virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-imodalwindow
type IModalWindow struct {
	comvt.IUnknown
	Show uintptr
}

// IShellItem virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct {
	comvt.IUnknown
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

// IShellItemArray virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct {
	comvt.IUnknown
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}

// IShellLink virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
type IShellLink struct {
	comvt.IUnknown
	GetPath             uintptr
	GetIDList           uintptr
	SetIDList           uintptr
	GetDescription      uintptr
	SetDescription      uintptr
	GetWorkingDirectory uintptr
	SetWorkingDirectory uintptr
	GetArguments        uintptr
	SetArguments        uintptr
	GetHotkey           uintptr
	SetHotkey           uintptr
	GetShowCmd          uintptr
	SetShowCmd          uintptr
	GetIconLocation     uintptr
	SetIconLocation     uintptr
	SetRelativePath     uintptr
	Resolve             uintptr
	SetPath             uintptr
}

// ITaskbarList virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist
type ITaskbarList struct {
	comvt.IUnknown
	HrInit       uintptr
	AddTab       uintptr
	DeleteTab    uintptr
	ActivateTab  uintptr
	SetActiveAlt uintptr
}

// ITaskbarList2 virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist2
type ITaskbarList2 struct {
	ITaskbarList
	MarkFullscreenWindow uintptr
}

// ITaskbarList3 virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct {
	ITaskbarList2
	SetProgressValue      uintptr
	SetProgressState      uintptr
	RegisterTab           uintptr
	UnregisterTab         uintptr
	SetTabOrder           uintptr
	SetTabActive          uintptr
	ThumbBarAddButtons    uintptr
	ThumbBarUpdateButtons uintptr
	ThumbBarSetImageList  uintptr
	SetOverlayIcon        uintptr
	SetThumbnailTooltip   uintptr
	SetThumbnailClip      uintptr
}

// ITaskbarList4 virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist4
type ITaskbarList4 struct {
	ITaskbarList3
	SetTabProperties uintptr
}
