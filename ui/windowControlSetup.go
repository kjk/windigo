/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

type windowControlSetup struct {
	wasInit bool // default to false

	ClassName        string      // Optional, defaults to a hash generated by WNDCLASSEX parameters. Passed to RegisterClassEx.
	ClassStyle       c.CS        // Window class style, passed to RegisterClassEx.
	HCursor          api.HCURSOR // Window cursor, passed to RegisterClassEx.
	HBrushBackground api.HBRUSH  // Window background brush, passed to RegisterClassEx.

	Style   c.WS    // Window style, passed to CreateWindowEx.
	ExStyle c.WS_EX // Window extended style, passed to CreateWindowEx. For a border, use WS_EX_CLIENTEDGE
}

func (me *windowControlSetup) initOnce() {
	if !me.wasInit {
		me.wasInit = true

		me.ClassStyle = c.CS_DBLCLKS

		me.Style = c.WS_CHILD | c.WS_VISIBLE | c.WS_CLIPCHILDREN | c.WS_CLIPSIBLINGS
		me.ExStyle = c.WS_EX(0)
	}
}

func (me *windowControlSetup) genWndClassEx(
	hInst api.HINSTANCE) *api.WNDCLASSEX {

	wcx := api.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle

	if me.HCursor != 0 {
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = api.HINSTANCE(0).LoadCursor(c.IDC_ARROW)
	}

	if me.HBrushBackground != 0 {
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = api.NewBrushFromSysColor(c.COLOR_WINDOW)
	}

	if me.ClassName == "" {
		me.ClassName = wcx.Hash() // generate hash after all other fields are set
	}
	wcx.LpszClassName = api.StrToUtf16Ptr(me.ClassName)

	return &wcx
}
