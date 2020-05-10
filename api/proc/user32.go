package proc

import (
	"syscall"
)

var (
	dllUser32 = syscall.NewLazyDLL("user32.dll")

	AppendMenu              = dllUser32.NewProc("AppendMenuW")
	ClientToScreen          = dllUser32.NewProc("ClientToScreen")
	CopyAcceleratorTable    = dllUser32.NewProc("CopyAcceleratorTableW")
	CreateAcceleratorTable  = dllUser32.NewProc("CreateAcceleratorTableW")
	CreateMenu              = dllUser32.NewProc("CreateMenu")
	CreateWindowEx          = dllUser32.NewProc("CreateWindowExW")
	DefWindowProc           = dllUser32.NewProc("DefWindowProcW")
	DeleteMenu              = dllUser32.NewProc("DeleteMenu")
	DestroyAcceleratorTable = dllUser32.NewProc("DestroyAcceleratorTable")
	DestroyMenu             = dllUser32.NewProc("DestroyMenu")
	DestroyWindow           = dllUser32.NewProc("DestroyWindow")
	DispatchMessage         = dllUser32.NewProc("DispatchMessageW")
	DrawMenuBar             = dllUser32.NewProc("DrawMenuBar")
	EnableMenuItem          = dllUser32.NewProc("EnableMenuItem")
	EnableWindow            = dllUser32.NewProc("EnableWindow")
	EnumChildWindows        = dllUser32.NewProc("EnumChildWindows")
	EnumDisplayMonitors     = dllUser32.NewProc("EnumDisplayMonitors")
	GetAncestor             = dllUser32.NewProc("GetAncestor")
	GetClassInfoEx          = dllUser32.NewProc("GetClassInfoExW")
	GetClientRect           = dllUser32.NewProc("GetClientRect")
	GetForegroundWindow     = dllUser32.NewProc("GetForegroundWindow")
	GetMenu                 = dllUser32.NewProc("GetMenu")
	GetMenuInfo             = dllUser32.NewProc("GetMenuInfo")
	GetMenuItemCount        = dllUser32.NewProc("GetMenuItemCount")
	GetMenuItemID           = dllUser32.NewProc("GetMenuItemID")
	GetMenuItemInfo         = dllUser32.NewProc("GetMenuItemInfoW")
	GetMessage              = dllUser32.NewProc("GetMessageW")
	GetMonitorInfo          = dllUser32.NewProc("GetMonitorInfoW")
	GetParent               = dllUser32.NewProc("GetParent")
	GetSubMenu              = dllUser32.NewProc("GetSubMenu")
	GetSystemMetrics        = dllUser32.NewProc("GetSystemMetrics")
	GetWindowDC             = dllUser32.NewProc("GetWindowDC")
	GetWindowLongPtr        = dllUser32.NewProc("SetWindowLongPtrW")
	GetWindowRect           = dllUser32.NewProc("GetWindowRect")
	GetWindowText           = dllUser32.NewProc("GetWindowTextW")
	GetWindowTextLength     = dllUser32.NewProc("GetWindowTextLengthW")
	InvalidateRect          = dllUser32.NewProc("InvalidateRect")
	IsDialogMessage         = dllUser32.NewProc("IsDialogMessageW")
	IsWindowEnabled         = dllUser32.NewProc("IsWindowEnabled")
	LoadCursor              = dllUser32.NewProc("LoadCursorW")
	MessageBox              = dllUser32.NewProc("MessageBoxW")
	PostQuitMessage         = dllUser32.NewProc("PostQuitMessage")
	RegisterClassEx         = dllUser32.NewProc("RegisterClassExW")
	ReleaseDC               = dllUser32.NewProc("ReleaseDC")
	ScreenToClient          = dllUser32.NewProc("ScreenToClient")
	SendMessage             = dllUser32.NewProc("SendMessageW")
	SetFocus                = dllUser32.NewProc("SetFocus")
	SetMenuItemInfo         = dllUser32.NewProc("SetMenuItemInfoW")
	SetRect                 = dllUser32.NewProc("SetRect")
	SetWindowLongPtr        = dllUser32.NewProc("SetWindowLongPtrW")
	SetWindowText           = dllUser32.NewProc("SetWindowTextW")
	ShowWindow              = dllUser32.NewProc("ShowWindow")
	SystemParametersInfo    = dllUser32.NewProc("SystemParametersInfoW")
	TranslateAccelerator    = dllUser32.NewProc("TranslateAcceleratorW")
	TranslateMessage        = dllUser32.NewProc("TranslateMessage")
	UpdateWindow            = dllUser32.NewProc("UpdateWindow")
)
