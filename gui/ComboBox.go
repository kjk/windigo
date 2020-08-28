/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Native combo box control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/about-combo-boxes
type ComboBox struct {
	_ControlNativeBase
}

// Adds a single item to the combo box.
func (me *ComboBox) Add(text string) *ComboBox {
	me.sendCbMessage(co.CB_ADDSTRING,
		0, win.LPARAM(unsafe.Pointer(win.StrToPtr(text))))
	return me
}

// Adds many items to the combo box.
func (me *ComboBox) AddMany(texts []string) *ComboBox {
	for _, text := range texts {
		me.Add(text)
	}
	return me
}

// Returns the number of items in the list box of a combo box.
func (me *ComboBox) Count() uint32 {
	return uint32(me.sendCbMessage(co.CB_GETCOUNT, 0, 0))
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) Create(
	parent Window, ctrlId, x, y int32, width uint32,
	exStyles co.WS_EX, styles co.WS, cbStyles co.CBS) *ComboBox {

	x, y, width, _ = _Util.MultiplyDpi(x, y, width, 0)

	me._ControlNativeBase.create(exStyles, "COMBOBOX", "",
		styles|co.WS(cbStyles), x, y, width, 0, parent, ctrlId)
	_globalUiFont.SetOnControl(me)
	return me
}

// Calls CreateWindowEx() with CBS_DROPDOWN.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateEditable(
	parent Window, ctrlId, x, y int32, width uint32) *ComboBox {

	return me.Create(parent, ctrlId, x, y, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWN)
}

// Calls CreateWindowEx() with CBS_DROPDOWN | CBS_SORT.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateEditableSorted(
	parent Window, ctrlId, x, y int32, width uint32) *ComboBox {

	return me.Create(parent, ctrlId, x, y, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWN|co.CBS_SORT)
}

// Calls CreateWindowEx() with CBS_DROPDOWNLIST.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateFixed(
	parent Window, ctrlId, x, y int32, width uint32) *ComboBox {

	return me.Create(parent, ctrlId, x, y, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWNLIST)
}

// Calls CreateWindowEx() with CBS_DROPDOWNLIST | CBS_SORT.
//
// Position and size will be adjusted to the current system DPI.
func (me *ComboBox) CreateFixedSorted(
	parent Window, ctrlId, x, y int32, width uint32) *ComboBox {

	return me.Create(parent, ctrlId, x, y, width, co.WS_EX_NONE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.CBS_DROPDOWNLIST|co.CBS_SORT)
}

// Limits the length of the text the user may type with CB_LIMITTEXT. Pass zero
// to remove the limitation.
//
// Works only if the combo box is editable.
func (me *ComboBox) LimitEditText(numChars uint32) *ComboBox {
	me.sendCbMessage(co.CB_LIMITTEXT, win.WPARAM(numChars), 0)
	return me
}

// Returns the index of the selected item, or -1 if none.
func (me *ComboBox) SelectedIndex() int32 {
	return int32(me.sendCbMessage(co.CB_GETCURSEL, 0, 0))
}

// Sets the index of the selected item, or -1 to clear.
func (me *ComboBox) SelectIndex(index int32) *ComboBox {
	me.sendCbMessage(co.CB_SETCURSEL, win.WPARAM(index), 0)
	return me
}

// Returns the selected text, if any.
func (me *ComboBox) SelectedText() (string, bool) {
	idx := me.SelectedIndex()
	if idx < 0 {
		return "", false
	}
	return me.Text(uint32(idx))
}

// Returns the string at the given index, if any.
//
// In an editable combo box, the text typed by the user can be retrieved with
// Hwnd().GetWindowText().
func (me *ComboBox) Text(index uint32) (string, bool) {
	len := int(me.sendCbMessage(co.CB_GETLBTEXTLEN, win.WPARAM(index), 0))
	if len == -1 {
		return "", false
	} else if len == 0 {
		return "", true
	}

	buf := make([]uint16, len+1)
	me.sendCbMessage(co.CB_GETLBTEXT,
		win.WPARAM(index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf), true
}

// Syntactic sugar.
func (me *ComboBox) sendCbMessage(msg co.CB,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
