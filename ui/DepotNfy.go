/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"
	"wingows/co"
	"wingows/win"
)

type _NfyHash struct { // custom hash for WM_NOTIFY messages
	IdFrom int
	Code   co.NM
}

// Keeps all user common control notification handlers.
type _DepotNfy struct {
	mapNfys map[_NfyHash]func(p WmNotify) uintptr
}

func (me *_DepotNfy) addNfy(idFrom int, code co.NM,
	userFunc func(p WmNotify) uintptr) {

	if me.mapNfys == nil { // guard
		me.mapNfys = make(map[_NfyHash]func(p WmNotify) uintptr, 4) // arbitrary capacity
	}
	me.mapNfys[_NfyHash{IdFrom: idFrom, Code: code}] = userFunc
}

func (me *_DepotNfy) processMessage(msg co.WM, p Wm) (uintptr, bool) {
	if msg == co.WM_NOTIFY {
		pNfy := WmNotify(p)
		hash := _NfyHash{
			IdFrom: int(pNfy.NmHdr().IdFrom),
			Code:   co.NM(pNfy.NmHdr().Code),
		}
		if userFunc, hasNfy := me.mapNfys[hash]; hasNfy {
			return userFunc(pNfy), true
		}
	}

	return 0, false // no user handler found
}

//---------------------------------------------------------- ComboBoxEx CBEN ---

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-beginedit
func (me *_DepotNfy) CbenBeginEdit(comboBoxExId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_BEGINEDIT), func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-deleteitem
func (me *_DepotNfy) CbenDeleteItem(comboBoxExId int, userFunc func(p *win.NMCOMBOBOXEX)) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMCOMBOBOXEX)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-dragbegin
func (me *_DepotNfy) CbenDragBegin(comboBoxExId int, userFunc func(p *win.NMCBEDRAGBEGIN)) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_DRAGBEGIN), func(p WmNotify) uintptr {
		userFunc((*win.NMCBEDRAGBEGIN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-endedit
func (me *_DepotNfy) CbenEndEdit(comboBoxExId int, userFunc func(p *win.NMCBEENDEDIT) bool) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_ENDEDIT), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMCBEENDEDIT)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-getdispinfo
func (me *_DepotNfy) CbenGetDispInfo(comboBoxExId int, userFunc func(p *win.NMCOMBOBOXEX)) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMCOMBOBOXEX)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/cben-insertitem
func (me *_DepotNfy) CbenInsertItem(comboBoxExId int, userFunc func(p *win.NMCOMBOBOXEX)) {
	me.addNfy(comboBoxExId, co.NM(co.CBEN_INSERTITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMCOMBOBOXEX)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------ ComboBoxEx NM ---

func (me *_DepotNfy) CbenSetCursor(comboBoxExId int, userFunc func(p *win.NMMOUSE) int) {
	me.addNfy(comboBoxExId, co.NM(co.NM_SETCURSOR), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam))))
	})
}

//------------------------------------------------------- DateTimePicker DTN ---

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-closeup
func (me *_DepotNfy) DtnCloseUp(dateTimePickerId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_CLOSEUP), func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-datetimechange
func (me *_DepotNfy) DtnDateTimeChange(dateTimePickerId int, userFunc func(p *win.NMDATETIMECHANGE)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_DATETIMECHANGE), func(p WmNotify) uintptr {
		userFunc((*win.NMDATETIMECHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-dropdown
func (me *_DepotNfy) DtnDropDown(dateTimePickerId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_DROPDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-format
func (me *_DepotNfy) DtnFormat(dateTimePickerId int, userFunc func(p *win.NMDATETIMEFORMAT)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_FORMAT), func(p WmNotify) uintptr {
		userFunc((*win.NMDATETIMEFORMAT)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-formatquery
func (me *_DepotNfy) DtnFormatQuery(dateTimePickerId int, userFunc func(p *win.NMDATETIMEFORMATQUERY)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_FORMATQUERY), func(p WmNotify) uintptr {
		userFunc((*win.NMDATETIMEFORMATQUERY)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-userstring
func (me *_DepotNfy) DtnUserString(dateTimePickerId int, userFunc func(p *win.NMDATETIMESTRING)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_USERSTRING), func(p WmNotify) uintptr {
		userFunc((*win.NMDATETIMESTRING)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/dtn-wmkeydown
func (me *_DepotNfy) DtnWmKeyDown(dateTimePickerId int, userFunc func(p *win.NMDATETIMEWMKEYDOWN)) {
	me.addNfy(dateTimePickerId, co.NM(co.DTN_WMKEYDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMDATETIMEWMKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//-------------------------------------------------------- DateTimePicker NM ---

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-date-time
func (me *_DepotNfy) DtnKillFocus(dateTimePickerId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(dateTimePickerId, co.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-date-time-
func (me *_DepotNfy) DtnSetFocus(dateTimePickerId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(dateTimePickerId, co.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- ListView LVN ---

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-begindrag
func (me *_DepotNfy) LvnBeginDrag(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginlabeledit
func (me *_DepotNfy) LvnBeginLabelEdit(listViewId int, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginrdrag
func (me *_DepotNfy) LvnBeginRDrag(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-beginscroll
func (me *_DepotNfy) LvnBeginScroll(listViewId int, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(listViewId, co.NM(co.LVN_BEGINSCROLL), func(p WmNotify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnclick
func (me *_DepotNfy) LvnColumnClick(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columndropdown
func (me *_DepotNfy) LvnColumnDropDown(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNDROPDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-columnoverflowclick
func (me *_DepotNfy) LvnColumnOverflowClick(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_COLUMNOVERFLOWCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteallitems
func (me *_DepotNfy) LvnDeleteAllItems(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_DELETEALLITEMS), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-deleteitem
func (me *_DepotNfy) LvnDeleteItem(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endlabeledit
func (me *_DepotNfy) LvnEndLabelEdit(listViewId int, userFunc func(p *win.NMLVDISPINFO) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-endscroll
func (me *_DepotNfy) LvnEndScroll(listViewId int, userFunc func(p *win.NMLVSCROLL)) {
	me.addNfy(listViewId, co.NM(co.LVN_ENDSCROLL), func(p WmNotify) uintptr {
		userFunc((*win.NMLVSCROLL)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getdispinfo
func (me *_DepotNfy) LvnGetDispInfo(listViewId int, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(listViewId, co.NM(co.LVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getemptymarkup
func (me *_DepotNfy) LvnGetEmptyMarkup(listViewId int, userFunc func(p *win.NMLVEMPTYMARKUP) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_GETEMPTYMARKUP), func(p WmNotify) uintptr {
		if userFunc((*win.NMLVEMPTYMARKUP)(unsafe.Pointer(p.LParam))) {
			return 1
		}
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-getinfotip
func (me *_DepotNfy) LvnGetInfoTip(listViewId int, userFunc func(p *win.NMLVGETINFOTIP)) {
	me.addNfy(listViewId, co.NM(co.LVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*win.NMLVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-hottrack
func (me *_DepotNfy) LvnHotTrack(listViewId int, userFunc func(p *win.NMLISTVIEW) int) {
	me.addNfy(listViewId, co.NM(co.LVN_HOTTRACK), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-incrementalsearch
func (me *_DepotNfy) LvnIncrementalSearch(listViewId int, userFunc func(p *win.NMLVFINDITEM) int) {
	me.addNfy(listViewId, co.NM(co.LVN_INCREMENTALSEARCH), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-insertitem
func (me *_DepotNfy) LvnInsertItem(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_INSERTITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemactivate
func (me *_DepotNfy) LvnItemActivate(listViewId int, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMACTIVATE), func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanged
func (me *_DepotNfy) LvnItemChanged(listViewId int, userFunc func(p *win.NMLISTVIEW)) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-itemchanging
func (me *_DepotNfy) LvnItemChanging(listViewId int, userFunc func(p *win.NMLISTVIEW) bool) {
	me.addNfy(listViewId, co.NM(co.LVN_ITEMCHANGING), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMLISTVIEW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-keydown
func (me *_DepotNfy) LvnKeyDown(listViewId int, userFunc func(p *win.NMLVKEYDOWN)) {
	me.addNfy(listViewId, co.NM(co.LVN_KEYDOWN), func(p WmNotify) uintptr {
		userFunc((*win.NMLVKEYDOWN)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-linkclick
func (me *_DepotNfy) LvnLinkClick(listViewId int, userFunc func(p *win.NMLVLINK)) {
	me.addNfy(listViewId, co.NM(co.LVN_LINKCLICK), func(p WmNotify) uintptr {
		userFunc((*win.NMLVLINK)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-marqueebegin
func (me *_DepotNfy) LvnMarqueeBegin(listViewId int, userFunc func(p *win.NMHDR) uint) {
	me.addNfy(listViewId, co.NM(co.LVN_MARQUEEBEGIN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odcachehint
func (me *_DepotNfy) LvnODCacheHint(listViewId int, userFunc func(p *win.NMLVCACHEHINT)) {
	me.addNfy(listViewId, co.NM(co.LVN_ODCACHEHINT), func(p WmNotify) uintptr {
		userFunc((*win.NMLVCACHEHINT)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odfinditem
func (me *_DepotNfy) LvnODFindItem(listViewId int, userFunc func(p *win.NMLVFINDITEM) int) {
	me.addNfy(listViewId, co.NM(co.LVN_ODFINDITEM), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVFINDITEM)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-odstatechanged
func (me *_DepotNfy) LvnODStateChanged(listViewId int, userFunc func(p *win.NMLVODSTATECHANGE)) {
	me.addNfy(listViewId, co.NM(co.LVN_ODSTATECHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMLVODSTATECHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/lvn-setdispinfo
func (me *_DepotNfy) LvnSetDispInfo(listViewId int, userFunc func(p *win.NMLVDISPINFO)) {
	me.addNfy(listViewId, co.NM(co.LVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMLVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//-------------------------------------------------------------- ListView NM ---

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-list-view
func (me *_DepotNfy) LvnClick(listViewId int, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-list-view
func (me *_DepotNfy) LvnCustomDraw(listViewId int, userFunc func(p *win.NMLVCUSTOMDRAW) co.CDRF) {
	me.addNfy(listViewId, co.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMLVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-list-view
func (me *_DepotNfy) LvnDblClk(listViewId int, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-hover-list-view
func (me *_DepotNfy) LvnHover(listViewId int, userFunc func(p *win.NMHDR) uint) {
	me.addNfy(listViewId, co.NM_HOVER, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-list-view
func (me *_DepotNfy) LvnKillFocus(listViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-list-view
func (me *_DepotNfy) LvnRClick(listViewId int, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-list-view
func (me *_DepotNfy) LvnRDblClk(listViewId int, userFunc func(p *win.NMITEMACTIVATE)) {
	me.addNfy(listViewId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMITEMACTIVATE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-releasedcapture-list-view-
func (me *_DepotNfy) LvnReleasedCapture(listViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_RELEASEDCAPTURE, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-list-view-
func (me *_DepotNfy) LvnReturn(listViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-list-view-
func (me *_DepotNfy) LvnSetFocus(listViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(listViewId, co.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------ StatusBar SBN ---

// https://docs.microsoft.com/en-us/windows/win32/controls/sbn-simplemodechange
func (me *_DepotNfy) SbnSimpleModeChange(statusBarId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(statusBarId, co.NM(co.SBN_SIMPLEMODECHANGE), func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- StatusBar NM ---

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-status-bar
func (me *_DepotNfy) SbnClick(statusBarId int, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-status-bar
func (me *_DepotNfy) SbnDblClk(statusBarId int, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-status-bar
func (me *_DepotNfy) SbnRClick(statusBarId int, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-status-bar
func (me *_DepotNfy) SbnRDblClk(statusBarId int, userFunc func(p *win.NMMOUSE)) {
	me.addNfy(statusBarId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

//------------------------------------------------------------- TreeView TVN ---

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-asyncdraw
func (me *_DepotNfy) TvnAsyncDraw(treeViewId int, userFunc func(p *win.NMTVASYNCDRAW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ASYNCDRAW), func(p WmNotify) uintptr {
		userFunc((*win.NMTVASYNCDRAW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-begindrag
func (me *_DepotNfy) TvnBeginDrag(treeViewId int, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginlabeledit
func (me *_DepotNfy) TvnBeginLabelEdit(treeViewId int, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINLABELEDIT), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-beginrdrag
func (me *_DepotNfy) TvnBeginRDrag(treeViewId int, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_BEGINRDRAG), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-deleteitem
func (me *_DepotNfy) TvnDeleteItem(treeViewId int, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_DELETEITEM), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-endlabeledit
func (me *_DepotNfy) TvnEndLabelEdit(treeViewId int, userFunc func(p *win.NMTVDISPINFO) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ENDLABELEDIT), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getdispinfo
func (me *_DepotNfy) TvnGetDispInfo(treeViewId int, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(treeViewId, co.NM(co.TVN_GETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-getinfotip
func (me *_DepotNfy) TvnGetInfoTip(treeViewId int, userFunc func(p *win.NMTVGETINFOTIP)) {
	me.addNfy(treeViewId, co.NM(co.TVN_GETINFOTIP), func(p WmNotify) uintptr {
		userFunc((*win.NMTVGETINFOTIP)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanged
func (me *_DepotNfy) TvnItemChanged(treeViewId int, userFunc func(p *win.NMTVITEMCHANGE)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemchanging
func (me *_DepotNfy) TvnItemChanging(treeViewId int, userFunc func(p *win.NMTVITEMCHANGE) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMCHANGING), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMTVITEMCHANGE)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanded
func (me *_DepotNfy) TvnItemExpanded(treeViewId int, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMEXPANDED), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-itemexpanding
func (me *_DepotNfy) TvnItemExpanding(treeViewId int, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_ITEMEXPANDING), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-keydown
func (me *_DepotNfy) TvnKeyDown(treeViewId int, userFunc func(p *win.NMTVKEYDOWN) int) {
	me.addNfy(treeViewId, co.NM(co.TVN_KEYDOWN), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTVKEYDOWN)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanged
func (me *_DepotNfy) TvnSelChanged(treeViewId int, userFunc func(p *win.NMTREEVIEW)) {
	me.addNfy(treeViewId, co.NM(co.TVN_SELCHANGED), func(p WmNotify) uintptr {
		userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-selchanging
func (me *_DepotNfy) TvnSelChanging(treeViewId int, userFunc func(p *win.NMTREEVIEW) bool) {
	me.addNfy(treeViewId, co.NM(co.TVN_SELCHANGING), func(p WmNotify) uintptr {
		return _Util.BoolToUintptr(userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-setdispinfo
func (me *_DepotNfy) TvnSetDispInfo(treeViewId int, userFunc func(p *win.NMTVDISPINFO)) {
	me.addNfy(treeViewId, co.NM(co.TVN_SETDISPINFO), func(p WmNotify) uintptr {
		userFunc((*win.NMTVDISPINFO)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/tvn-singleexpand
func (me *_DepotNfy) TvnSingleExpand(treeViewId int, userFunc func(p *win.NMTREEVIEW) co.TVNRET) {
	me.addNfy(treeViewId, co.NM(co.TVN_SINGLEEXPAND), func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTREEVIEW)(unsafe.Pointer(p.LParam))))
	})
}

//--------------------------------------------------------------- TreView NM ---

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-tree-view
func (me *_DepotNfy) TvnClick(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_CLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-customdraw-tree-view
func (me *_DepotNfy) TvnCustomDraw(treeViewId int, userFunc func(p *win.NMTVCUSTOMDRAW) co.CDRF) {
	me.addNfy(treeViewId, co.NM_CUSTOMDRAW, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMTVCUSTOMDRAW)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-tree-view
func (me *_DepotNfy) TvnDblClk(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_DBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-killfocus-tree-view
func (me *_DepotNfy) TvnKillFocus(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_KILLFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-tree-view
func (me *_DepotNfy) TvnRClick(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RCLICK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-tree-view
func (me *_DepotNfy) TvnRDblClk(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RDBLCLK, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-return-tree-view-
func (me *_DepotNfy) TvnReturn(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_RETURN, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setcursor-tree-view-
func (me *_DepotNfy) TvnSetCursor(treeViewId int, userFunc func(p *win.NMMOUSE) int) {
	me.addNfy(treeViewId, co.NM_SETCURSOR, func(p WmNotify) uintptr {
		return uintptr(userFunc((*win.NMMOUSE)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-setfocus-tree-view-
func (me *_DepotNfy) TvnSetFocus(treeViewId int, userFunc func(p *win.NMHDR)) {
	me.addNfy(treeViewId, co.NM_SETFOCUS, func(p WmNotify) uintptr {
		userFunc((*win.NMHDR)(unsafe.Pointer(p.LParam)))
		return 0
	})
}