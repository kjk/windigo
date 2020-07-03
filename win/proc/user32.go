/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package proc

import (
	"syscall"
)

var (
	dllUser32 = syscall.NewLazyDLL("user32.dll")

	AppendMenu                    = dllUser32.NewProc("AppendMenuW")
	BeginDeferWindowPos           = dllUser32.NewProc("BeginDeferWindowPos")
	ClientToScreen                = dllUser32.NewProc("ClientToScreen")
	CopyAcceleratorTable          = dllUser32.NewProc("CopyAcceleratorTableW")
	CreateAcceleratorTable        = dllUser32.NewProc("CreateAcceleratorTableW")
	CreateMenu                    = dllUser32.NewProc("CreateMenu")
	CreateWindowEx                = dllUser32.NewProc("CreateWindowExW")
	DeferWindowPos                = dllUser32.NewProc("DeferWindowPos")
	DefWindowProc                 = dllUser32.NewProc("DefWindowProcW")
	DeleteMenu                    = dllUser32.NewProc("DeleteMenu")
	DestroyAcceleratorTable       = dllUser32.NewProc("DestroyAcceleratorTable")
	DestroyMenu                   = dllUser32.NewProc("DestroyMenu")
	DestroyWindow                 = dllUser32.NewProc("DestroyWindow")
	DispatchMessage               = dllUser32.NewProc("DispatchMessageW")
	DrawMenuBar                   = dllUser32.NewProc("DrawMenuBar")
	EnableMenuItem                = dllUser32.NewProc("EnableMenuItem")
	EnableWindow                  = dllUser32.NewProc("EnableWindow")
	EndDeferWindowPos             = dllUser32.NewProc("EndDeferWindowPos")
	EnumChildWindows              = dllUser32.NewProc("EnumChildWindows")
	EnumDisplayMonitors           = dllUser32.NewProc("EnumDisplayMonitors")
	GetAncestor                   = dllUser32.NewProc("GetAncestor")
	GetClassInfoEx                = dllUser32.NewProc("GetClassInfoExW")
	GetClientRect                 = dllUser32.NewProc("GetClientRect")
	GetCursorPos                  = dllUser32.NewProc("GetCursorPos")
	GetDC                         = dllUser32.NewProc("GetDC")
	GetDlgItem                    = dllUser32.NewProc("GetDlgItem")
	GetDpiForSystem               = dllUser32.NewProc("GetDpiForSystem")
	GetDpiForWindow               = dllUser32.NewProc("GetDpiForWindow")
	GetFocus                      = dllUser32.NewProc("GetFocus")
	GetForegroundWindow           = dllUser32.NewProc("GetForegroundWindow")
	GetMenu                       = dllUser32.NewProc("GetMenu")
	GetMenuInfo                   = dllUser32.NewProc("GetMenuInfo")
	GetMenuItemCount              = dllUser32.NewProc("GetMenuItemCount")
	GetMenuItemID                 = dllUser32.NewProc("GetMenuItemID")
	GetMenuItemInfo               = dllUser32.NewProc("GetMenuItemInfoW")
	GetMessage                    = dllUser32.NewProc("GetMessageW")
	GetMonitorInfo                = dllUser32.NewProc("GetMonitorInfoW")
	GetNextDlgTabItem             = dllUser32.NewProc("GetNextDlgTabItem")
	GetParent                     = dllUser32.NewProc("GetParent")
	GetSubMenu                    = dllUser32.NewProc("GetSubMenu")
	GetSystemMetrics              = dllUser32.NewProc("GetSystemMetrics")
	GetWindow                     = dllUser32.NewProc("GetWindow")
	GetWindowDC                   = dllUser32.NewProc("GetWindowDC")
	GetWindowLongPtr              = dllUser32.NewProc("GetWindowLongPtrW")
	GetWindowRect                 = dllUser32.NewProc("GetWindowRect")
	GetWindowText                 = dllUser32.NewProc("GetWindowTextW")
	GetWindowTextLength           = dllUser32.NewProc("GetWindowTextLengthW")
	InvalidateRect                = dllUser32.NewProc("InvalidateRect")
	IsChild                       = dllUser32.NewProc("IsChild")
	IsDialogMessage               = dllUser32.NewProc("IsDialogMessageW")
	IsDlgButtonChecked            = dllUser32.NewProc("IsDlgButtonChecked")
	IsWindow                      = dllUser32.NewProc("IsWindow")
	IsWindowEnabled               = dllUser32.NewProc("IsWindowEnabled")
	LoadCursor                    = dllUser32.NewProc("LoadCursorW")
	LoadIcon                      = dllUser32.NewProc("LoadIconW")
	LoadMenu                      = dllUser32.NewProc("LoadMenuW")
	MenuItemFromPoint             = dllUser32.NewProc("MeunItemFromPoint")
	MessageBox                    = dllUser32.NewProc("MessageBoxW")
	MonitorFromPoint              = dllUser32.NewProc("MonitorFromPoint")
	PostMessage                   = dllUser32.NewProc("PostMessageW")
	PostQuitMessage               = dllUser32.NewProc("PostQuitMessage")
	RegisterClassEx               = dllUser32.NewProc("RegisterClassExW")
	ReleaseDC                     = dllUser32.NewProc("ReleaseDC")
	ScreenToClient                = dllUser32.NewProc("ScreenToClient")
	SendMessage                   = dllUser32.NewProc("SendMessageW")
	SetFocus                      = dllUser32.NewProc("SetFocus")
	SetForegroundWindow           = dllUser32.NewProc("SetForegroundWindow")
	SetMenuItemInfo               = dllUser32.NewProc("SetMenuItemInfoW")
	SetProcessDPIAware            = dllUser32.NewProc("SetProcessDPIAware")
	SetProcessDpiAwarenessContext = dllUser32.NewProc("SetProcessDpiAwarenessContext")
	SetRect                       = dllUser32.NewProc("SetRect")
	SetWindowLongPtr              = dllUser32.NewProc("SetWindowLongPtrW")
	SetWindowPos                  = dllUser32.NewProc("SetWindowPos")
	SetWindowText                 = dllUser32.NewProc("SetWindowTextW")
	ShowWindow                    = dllUser32.NewProc("ShowWindow")
	SystemParametersInfo          = dllUser32.NewProc("SystemParametersInfoW")
	TranslateAccelerator          = dllUser32.NewProc("TranslateAcceleratorW")
	TranslateMessage              = dllUser32.NewProc("TranslateMessage")
	UpdateWindow                  = dllUser32.NewProc("UpdateWindow")
)
