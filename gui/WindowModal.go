/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"unsafe"
	"wingows/co"
	"wingows/gui/wm"
	"wingows/win"
)

// Modal popup window.
// Allows message and notification handling.
type WindowModal struct {
	windowBase
	prevFocusParent win.HWND // child control last focused on parent
	setup           windowModalSetup
}

// Parameters that will be used to create the window.
func (me *WindowModal) Setup() *windowModalSetup {
	if me.windowBase.Hwnd() != 0 {
		panic("Cannot change setup after the window was created.")
	}
	me.setup.initOnce() // guard
	return &me.setup
}

// Creates the modal window and disables the parent. This function will return
// only after the modal is closed.
func (me *WindowModal) Show(parent Window) {
	me.setup.initOnce() // guard
	hInst := parent.Hwnd().GetInstance()
	me.windowBase.registerClass(me.setup.genWndClassEx(hInst))

	me.defaultMessageHandling()

	me.prevFocusParent = win.GetFocus() // currently focused control
	parent.Hwnd().EnableWindow(false)   // https://devblogs.microsoft.com/oldnewthing/20040227-00/?p=40463

	_, _, cx, cy := multiplyByDpi(0, 0, me.setup.Width, me.setup.Height)

	me.windowBase.createWindow("WindowModal", me.setup.ExStyle,
		me.setup.ClassName, me.setup.Title, me.setup.Style,
		0, 0, // initially anchored at zero
		cx, cy, parent, win.HMENU(0), hInst)

	rc := me.windowBase.Hwnd().GetWindowRect()
	rcParent := parent.Hwnd().GetWindowRect() // both rc relative to screen

	me.windowBase.Hwnd().SetWindowPos(co.SWP_HWND(0), // center modal over parent (warning: happens after WM_CREATE processing)
		rcParent.Left+(rcParent.Right-rcParent.Left)/2-(rc.Right-rc.Left)/2,
		rcParent.Top+(rcParent.Bottom-rcParent.Top)/2-(rc.Bottom-rc.Top)/2,
		0, 0, co.SWP_NOZORDER|co.SWP_NOSIZE)

	me.runModalLoop()
}

func (me *WindowModal) defaultMessageHandling() {
	me.windowBase.OnMsg().WmSetFocus(func(p wm.SetFocus) {
		if me.windowBase.Hwnd() == win.GetFocus() {
			// If window receive focus, delegate to first child.
			// This also happens right after the modal is created.
			me.windowBase.Hwnd().
				GetNextDlgTabItem(win.HWND(0), false).
				SetFocus()
		}
	})

	me.windowBase.OnMsg().WmClose(func() {
		me.windowBase.Hwnd().
			GetWindow(co.GW_OWNER).EnableWindow(true) // re-enable parent
		me.windowBase.Hwnd().DestroyWindow() // then destroy modal
		me.prevFocusParent.SetFocus()        // could be on WM_DESTROY too
	})
}

func (me *WindowModal) runModalLoop() {
	msg := win.MSG{}
	for {
		status := msg.GetMessage(win.HWND(0), 0, 0)
		if status == 0 {
			// WM_QUIT was sent, exit modal loop now.
			// https://devblogs.microsoft.com/oldnewthing/20050222-00/?p=36393
			win.PostQuitMessage(int32(msg.WParam))
			break
		}

		// If a child window, will retrieve its top-level parent.
		// If a top-level, use itself.
		if msg.HWnd.GetAncestor(co.GA_ROOT).IsDialogMessage(&msg) {
			// Processed all keyboard actions for child controls.
			if me.hwnd == win.HWND(0) {
				break // our modal was destroyed
			} else {
				continue
			}
		}

		msg.TranslateMessage()
		msg.DispatchMessage()

		if me.hwnd == win.HWND(0) {
			break // our modal was destroyed
		}
	}
}

//------------------------------------------------------------------------------

type windowModalSetup struct {
	wasInit bool // default to false

	ClassName        string      // Optional, defaults to a hash generated by WNDCLASSEX parameters. Passed to RegisterClassEx.
	ClassStyle       co.CS       // Window class style, passed to RegisterClassEx.
	HCursor          win.HCURSOR // Window cursor, passed to RegisterClassEx.
	HBrushBackground win.HBRUSH  // Window background brush, passed to RegisterClassEx.

	Title   string   // The title of the window, passed to CreateWindowEx.
	Width   uint32   // Initial width of the window, passed to CreateWindowEx.
	Height  uint32   // Initial height of the window, passed to CreateWindowEx.
	Style   co.WS    // Window style, passed to CreateWindowEx.
	ExStyle co.WS_EX // Window extended style, passed to CreateWindowEx.
}

func (me *windowModalSetup) initOnce() {
	if !me.wasInit {
		me.wasInit = true

		me.ClassStyle = co.CS_DBLCLKS

		me.Width = 500 // arbitrary dimensions
		me.Height = 400
		me.Style = co.WS_CAPTION | co.WS_SYSMENU | co.WS_CLIPCHILDREN | co.WS_BORDER | co.WS_VISIBLE
		me.ExStyle = co.WS_EX(0)
	}
}

func (me *windowModalSetup) genWndClassEx(hInst win.HINSTANCE) *win.WNDCLASSEX {
	wcx := win.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 {
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = win.HINSTANCE(0).LoadCursor(co.IDC_ARROW)
	}

	if me.HBrushBackground != 0 {
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = win.CreateSysColorBrush(co.COLOR_BTNFACE)
	}

	if me.ClassName == "" {
		me.ClassName = wcx.Hash() // generate hash after all other fields are set
	}
	wcx.LpszClassName = win.StrToPtr(me.ClassName)

	return &wcx
}
