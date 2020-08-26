/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// ACCELL fVirt.
type ACCELF uint8

const (
	ACCELF_NONE    ACCELF = 0
	ACCELF_VIRTKEY ACCELF = 1
	ACCELF_SHIFT   ACCELF = 0x04
	ACCELF_CONTROL ACCELF = 0x08
	ACCELF_ALT     ACCELF = 0x10
)

// Composes KEY, STANDARD_RIGHTS; originally has no prefix.
type ACCESS_RIGHTS uint32

const (
	ACCESS_RIGHTS_DELETE       ACCESS_RIGHTS = 0x00010000
	ACCESS_RIGHTS_READ_CONTROL ACCESS_RIGHTS = 0x00020000
	ACCESS_RIGHTS_WRITE_DAC    ACCESS_RIGHTS = 0x00040000
	ACCESS_RIGHTS_WRITE_OWNER  ACCESS_RIGHTS = 0x00080000
	ACCESS_RIGHTS_SYNCHRONIZE  ACCESS_RIGHTS = 0x00100000
)

// NMTVASYNCDRAW dwRetFlags, don't seem to be defined anywhere, values are unconfirmed.
type ADRF uint32

const (
	ADRF_DRAWSYNC     ADRF = 0
	ADRF_DRAWNOTHING  ADRF = 1
	ADRF_DRAWFALLBACK ADRF = 2
	ADRF_DRAWIMAGE    ADRF = 3
)

// WM_APPCOMMAND command.
type APPCOMMAND int16

const (
	APPCOMMAND_BROWSER_BACKWARD                  APPCOMMAND = 1
	APPCOMMAND_BROWSER_FORWARD                   APPCOMMAND = 2
	APPCOMMAND_BROWSER_REFRESH                   APPCOMMAND = 3
	APPCOMMAND_BROWSER_STOP                      APPCOMMAND = 4
	APPCOMMAND_BROWSER_SEARCH                    APPCOMMAND = 5
	APPCOMMAND_BROWSER_FAVORITES                 APPCOMMAND = 6
	APPCOMMAND_BROWSER_HOME                      APPCOMMAND = 7
	APPCOMMAND_VOLUME_MUTE                       APPCOMMAND = 8
	APPCOMMAND_VOLUME_DOWN                       APPCOMMAND = 9
	APPCOMMAND_VOLUME_UP                         APPCOMMAND = 10
	APPCOMMAND_MEDIA_NEXTTRACK                   APPCOMMAND = 11
	APPCOMMAND_MEDIA_PREVIOUSTRACK               APPCOMMAND = 12
	APPCOMMAND_MEDIA_STOP                        APPCOMMAND = 13
	APPCOMMAND_MEDIA_PLAY_PAUSE                  APPCOMMAND = 14
	APPCOMMAND_LAUNCH_MAIL                       APPCOMMAND = 15
	APPCOMMAND_LAUNCH_MEDIA_SELECT               APPCOMMAND = 16
	APPCOMMAND_LAUNCH_APP1                       APPCOMMAND = 17
	APPCOMMAND_LAUNCH_APP2                       APPCOMMAND = 18
	APPCOMMAND_BASS_DOWN                         APPCOMMAND = 19
	APPCOMMAND_BASS_BOOST                        APPCOMMAND = 20
	APPCOMMAND_BASS_UP                           APPCOMMAND = 21
	APPCOMMAND_TREBLE_DOWN                       APPCOMMAND = 22
	APPCOMMAND_TREBLE_UP                         APPCOMMAND = 23
	APPCOMMAND_MICROPHONE_VOLUME_MUTE            APPCOMMAND = 24
	APPCOMMAND_MICROPHONE_VOLUME_DOWN            APPCOMMAND = 25
	APPCOMMAND_MICROPHONE_VOLUME_UP              APPCOMMAND = 26
	APPCOMMAND_HELP                              APPCOMMAND = 27
	APPCOMMAND_FIND                              APPCOMMAND = 28
	APPCOMMAND_NEW                               APPCOMMAND = 29
	APPCOMMAND_OPEN                              APPCOMMAND = 30
	APPCOMMAND_CLOSE                             APPCOMMAND = 31
	APPCOMMAND_SAVE                              APPCOMMAND = 32
	APPCOMMAND_PRINT                             APPCOMMAND = 33
	APPCOMMAND_UNDO                              APPCOMMAND = 34
	APPCOMMAND_REDO                              APPCOMMAND = 35
	APPCOMMAND_COPY                              APPCOMMAND = 36
	APPCOMMAND_CUT                               APPCOMMAND = 37
	APPCOMMAND_PASTE                             APPCOMMAND = 38
	APPCOMMAND_REPLY_TO_MAIL                     APPCOMMAND = 39
	APPCOMMAND_FORWARD_MAIL                      APPCOMMAND = 40
	APPCOMMAND_SEND_MAIL                         APPCOMMAND = 41
	APPCOMMAND_SPELL_CHECK                       APPCOMMAND = 42
	APPCOMMAND_DICTATE_OR_COMMAND_CONTROL_TOGGLE APPCOMMAND = 43
	APPCOMMAND_MIC_ON_OFF_TOGGLE                 APPCOMMAND = 44
	APPCOMMAND_CORRECTION_LIST                   APPCOMMAND = 45
	APPCOMMAND_MEDIA_PLAY                        APPCOMMAND = 46
	APPCOMMAND_MEDIA_PAUSE                       APPCOMMAND = 47
	APPCOMMAND_MEDIA_RECORD                      APPCOMMAND = 48
	APPCOMMAND_MEDIA_FAST_FORWARD                APPCOMMAND = 49
	APPCOMMAND_MEDIA_REWIND                      APPCOMMAND = 50
	APPCOMMAND_MEDIA_CHANNEL_UP                  APPCOMMAND = 51
	APPCOMMAND_MEDIA_CHANNEL_DOWN                APPCOMMAND = 52
	APPCOMMAND_DELETE                            APPCOMMAND = 53
	APPCOMMAND_DWM_FLIP3D                        APPCOMMAND = 54
)

// SetBkMode() mode.
type BKMODE int32

const (
	BKMODE_TRANSPARENT BKMODE = 1
	BKMODE_OPAQUE      BKMODE = 2
)

// Button control messages.
type BM WM

const (
	BM_GETCHECK     BM = 0x00F0
	BM_SETCHECK     BM = 0x00F1
	BM_GETSTATE     BM = 0x00F2
	BM_SETSTATE     BM = 0x00F3
	BM_SETSTYLE     BM = 0x00F4
	BM_CLICK        BM = 0x00F5
	BM_GETIMAGE     BM = 0x00F6
	BM_SETIMAGE     BM = 0x00F7
	BM_SETDONTCLICK BM = 0x00F8
)

// Button control styles.
type BS WS

const (
	BS_PUSHBUTTON      BS = 0x00000000
	BS_DEFPUSHBUTTON   BS = 0x00000001
	BS_CHECKBOX        BS = 0x00000002
	BS_AUTOCHECKBOX    BS = 0x00000003
	BS_RADIOBUTTON     BS = 0x00000004
	BS_3STATE          BS = 0x00000005
	BS_AUTO3STATE      BS = 0x00000006
	BS_GROUPBOX        BS = 0x00000007
	BS_USERBUTTON      BS = 0x00000008
	BS_AUTORADIOBUTTON BS = 0x00000009
	BS_PUSHBOX         BS = 0x0000000A
	BS_OWNERDRAW       BS = 0x0000000B
	BS_TYPEMASK        BS = 0x0000000F
	BS_LEFTTEXT        BS = 0x00000020
	BS_TEXT            BS = 0x00000000
	BS_ICON            BS = 0x00000040
	BS_BITMAP          BS = 0x00000080
	BS_LEFT            BS = 0x00000100
	BS_RIGHT           BS = 0x00000200
	BS_CENTER          BS = 0x00000300
	BS_TOP             BS = 0x00000400
	BS_BOTTOM          BS = 0x00000800
	BS_VCENTER         BS = 0x00000C00
	BS_PUSHLIKE        BS = 0x00001000
	BS_MULTILINE       BS = 0x00002000
	BS_NOTIFY          BS = 0x00004000
	BS_FLAT            BS = 0x00008000
	BS_RIGHTBUTTON     BS = BS_LEFTTEXT
)

// IsDlgButtonChecked() nIDButton.
type BST uint32

const (
	BST_UNCHECKED     BST = 0x0000
	BST_CHECKED       BST = 0x0001
	BST_INDETERMINATE BST = 0x0002
	BST_PUSHED        BST = 0x0004
	BST_FOCUS         BST = 0x0008
)

// Combo box messages.
type CB WM

const (
	CB_GETEDITSEL            CB = 0x0140
	CB_LIMITTEXT             CB = 0x0141
	CB_SETEDITSEL            CB = 0x0142
	CB_ADDSTRING             CB = 0x0143
	CB_DELETESTRING          CB = 0x0144
	CB_DIR                   CB = 0x0145
	CB_GETCOUNT              CB = 0x0146
	CB_GETCURSEL             CB = 0x0147
	CB_GETLBTEXT             CB = 0x0148
	CB_GETLBTEXTLEN          CB = 0x0149
	CB_INSERTSTRING          CB = 0x014A
	CB_RESETCONTENT          CB = 0x014B
	CB_FINDSTRING            CB = 0x014C
	CB_SELECTSTRING          CB = 0x014D
	CB_SETCURSEL             CB = 0x014E
	CB_SHOWDROPDOWN          CB = 0x014F
	CB_GETITEMDATA           CB = 0x0150
	CB_SETITEMDATA           CB = 0x0151
	CB_GETDROPPEDCONTROLRECT CB = 0x0152
	CB_SETITEMHEIGHT         CB = 0x0153
	CB_GETITEMHEIGHT         CB = 0x0154
	CB_SETEXTENDEDUI         CB = 0x0155
	CB_GETEXTENDEDUI         CB = 0x0156
	CB_GETDROPPEDSTATE       CB = 0x0157
	CB_FINDSTRINGEXACT       CB = 0x0158
	CB_SETLOCALE             CB = 0x0159
	CB_GETLOCALE             CB = 0x015A
	CB_GETTOPINDEX           CB = 0x015b
	CB_SETTOPINDEX           CB = 0x015c
	CB_GETHORIZONTALEXTENT   CB = 0x015d
	CB_SETHORIZONTALEXTENT   CB = 0x015e
	CB_GETDROPPEDWIDTH       CB = 0x015f
	CB_SETDROPPEDWIDTH       CB = 0x0160
	CB_INITSTORAGE           CB = 0x0161
	CB_GETCOMBOBOXINFO       CB = 0x0164
	CB_MSGMAX                CB = 0x0165
)

// Combo box styles.
type CBS WS

const (
	CBS_SIMPLE            CBS = 0x0001
	CBS_DROPDOWN          CBS = 0x0002
	CBS_DROPDOWNLIST      CBS = 0x0003
	CBS_OWNERDRAWFIXED    CBS = 0x0010
	CBS_OWNERDRAWVARIABLE CBS = 0x0020
	CBS_AUTOHSCROLL       CBS = 0x0040
	CBS_OEMCONVERT        CBS = 0x0080
	CBS_SORT              CBS = 0x0100
	CBS_HASSTRINGS        CBS = 0x0200
	CBS_NOINTEGRALHEIGHT  CBS = 0x0400
	CBS_DISABLENOSCROLL   CBS = 0x0800
	CBS_UPPERCASE         CBS = 0x2000
	CBS_LOWERCASE         CBS = 0x4000
)

// Common controls shared messages.
type CCM WM

const (
	_CCM_FIRST CCM = 0x2000

	CCM_SETBKCOLOR       CCM = _CCM_FIRST + 1
	CCM_SETCOLORSCHEME   CCM = _CCM_FIRST + 2
	CCM_GETCOLORSCHEME   CCM = _CCM_FIRST + 3
	CCM_GETDROPTARGET    CCM = _CCM_FIRST + 4
	CCM_SETUNICODEFORMAT CCM = _CCM_FIRST + 5
	CCM_GETUNICODEFORMAT CCM = _CCM_FIRST + 6
	CCM_SETVERSION       CCM = _CCM_FIRST + 0x7
	CCM_GETVERSION       CCM = _CCM_FIRST + 0x8
	CCM_SETNOTIFYWINDOW  CCM = _CCM_FIRST + 0x9
	CCM_SETWINDOWTHEME   CCM = _CCM_FIRST + 0xb
	CCM_DPISCALE         CCM = _CCM_FIRST + 0xc
)

// NMCUSTOMDRAW uItemState.
type CDIS uint32

const (
	CDIS_SELECTED         CDIS = 0x0001
	CDIS_GRAYED           CDIS = 0x0002
	CDIS_DISABLED         CDIS = 0x0004
	CDIS_CHECKED          CDIS = 0x0008
	CDIS_FOCUS            CDIS = 0x0010
	CDIS_DEFAULT          CDIS = 0x0020
	CDIS_HOT              CDIS = 0x0040
	CDIS_MARKED           CDIS = 0x0080
	CDIS_INDETERMINATE    CDIS = 0x0100
	CDIS_SHOWKEYBOARDCUES CDIS = 0x0200
	CDIS_NEARHOT          CDIS = 0x0400
	CDIS_OTHERSIDEHOT     CDIS = 0x0800
	CDIS_DROPHILITED      CDIS = 0x1000
)

// NMCUSTOMDRAW dwDrawStage.
type CDDS uint32

const (
	CDDS_PREPAINT      CDDS = 0x00000001
	CDDS_POSTPAINT     CDDS = 0x00000002
	CDDS_PREERASE      CDDS = 0x00000003
	CDDS_POSTERASE     CDDS = 0x00000004
	CDDS_ITEM          CDDS = 0x00010000
	CDDS_ITEMPREPAINT  CDDS = CDDS_ITEM | CDDS_PREPAINT
	CDDS_ITEMPOSTPAINT CDDS = CDDS_ITEM | CDDS_POSTPAINT
	CDDS_ITEMPREERASE  CDDS = CDDS_ITEM | CDDS_PREERASE
	CDDS_ITEMPOSTERASE CDDS = CDDS_ITEM | CDDS_POSTERASE
	CDDS_SUBITEM       CDDS = 0x00020000
)

// NM_CUSTOMDRAW return value.
type CDRF uint32

const (
	CDRF_DODEFAULT         CDRF = 0x00000000
	CDRF_NEWFONT           CDRF = 0x00000002
	CDRF_SKIPDEFAULT       CDRF = 0x00000004
	CDRF_DOERASE           CDRF = 0x00000008
	CDRF_SKIPPOSTPAINT     CDRF = 0x00000100
	CDRF_NOTIFYPOSTPAINT   CDRF = 0x00000010
	CDRF_NOTIFYITEMDRAW    CDRF = 0x00000020
	CDRF_NOTIFYSUBITEMDRAW CDRF = 0x00000020
	CDRF_NOTIFYPOSTERASE   CDRF = 0x00000040
)

// CoCreateInstance() dwClsContext.
type CLSCTX uint32

const (
	CLSCTX_INPROC_SERVER          CLSCTX = 0x1
	CLSCTX_INPROC_HANDLER         CLSCTX = 0x2
	CLSCTX_LOCAL_SERVER           CLSCTX = 0x4
	CLSCTX_INPROC_SERVER16        CLSCTX = 0x8
	CLSCTX_REMOTE_SERVER          CLSCTX = 0x10
	CLSCTX_INPROC_HANDLER16       CLSCTX = 0x20
	CLSCTX_RESERVED1              CLSCTX = 0x40
	CLSCTX_RESERVED2              CLSCTX = 0x80
	CLSCTX_RESERVED3              CLSCTX = 0x100
	CLSCTX_RESERVED4              CLSCTX = 0x200
	CLSCTX_NO_CODE_DOWNLOAD       CLSCTX = 0x400
	CLSCTX_RESERVED5              CLSCTX = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL      CLSCTX = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   CLSCTX = 0x2000
	CLSCTX_NO_FAILURE_LOG         CLSCTX = 0x4000
	CLSCTX_DISABLE_AAA            CLSCTX = 0x8000
	CLSCTX_ENABLE_AAA             CLSCTX = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   CLSCTX = 0x20000
	CLSCTX_ACTIVATE_X86_SERVER    CLSCTX = 0x40000
	CLSCTX_ACTIVATE_32_BIT_SERVER CLSCTX = CLSCTX_ACTIVATE_X86_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER CLSCTX = 0x80000
	CLSCTX_ENABLE_CLOAKING        CLSCTX = 0x100000
	CLSCTX_APPCONTAINER           CLSCTX = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU     CLSCTX = 0x800000
	CLSCTX_RESERVED6              CLSCTX = 0x1000000
	CLSCTX_ACTIVATE_ARM32_SERVER  CLSCTX = 0x2000000
	CLSCTX_PS_DLL                 CLSCTX = 0x80000000
	CLSCTX_ALL                    CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
	CLSCTX_SERVER                 CLSCTX = CLSCTX_INPROC_SERVER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

// CoInitializeEx() dwCoInit.
type COINIT uint32

const (
	COINIT_APARTMENTTHREADED COINIT = 0x2
	COINIT_MULTITHREADED     COINIT = 0x0
	COINIT_DISABLE_OLE1DDE   COINIT = 0x4
	COINIT_SPEED_OVER_MEMORY COINIT = 0x8
)

// System colors.
type COLOR uint32

const (
	COLOR_SCROLLBAR               COLOR = 0
	COLOR_BACKGROUND              COLOR = 1
	COLOR_ACTIVECAPTION           COLOR = 2
	COLOR_INACTIVECAPTION         COLOR = 3
	COLOR_MENU                    COLOR = 4
	COLOR_WINDOW                  COLOR = 5
	COLOR_WINDOWFRAME             COLOR = 6
	COLOR_MENUTEXT                COLOR = 7
	COLOR_WINDOWTEXT              COLOR = 8
	COLOR_CAPTIONTEXT             COLOR = 9
	COLOR_ACTIVEBORDER            COLOR = 10
	COLOR_INACTIVEBORDER          COLOR = 11
	COLOR_APPWORKSPACE            COLOR = 12
	COLOR_HIGHLIGHT               COLOR = 13
	COLOR_HIGHLIGHTTEXT           COLOR = 14
	COLOR_BTNFACE                 COLOR = 15
	COLOR_BTNSHADOW               COLOR = 16
	COLOR_GRAYTEXT                COLOR = 17
	COLOR_BTNTEXT                 COLOR = 18
	COLOR_INACTIVECAPTIONTEXT     COLOR = 19
	COLOR_BTNHIGHLIGHT            COLOR = 20
	COLOR_3DDKSHADOW              COLOR = 21
	COLOR_3DLIGHT                 COLOR = 22
	COLOR_INFOTEXT                COLOR = 23
	COLOR_INFOBK                  COLOR = 24
	COLOR_HOTLIGHT                COLOR = 26
	COLOR_GRADIENTACTIVECAPTION   COLOR = 27
	COLOR_GRADIENTINACTIVECAPTION COLOR = 28
	COLOR_MENUHILIGHT             COLOR = 29
	COLOR_MENUBAR                 COLOR = 30
	COLOR_DESKTOP                 COLOR = COLOR_BACKGROUND
	COLOR_3DFACE                  COLOR = COLOR_BTNFACE
	COLOR_3DSHADOW                COLOR = COLOR_BTNSHADOW
	COLOR_3DHIGHLIGHT             COLOR = COLOR_BTNHIGHLIGHT
	COLOR_3DHILIGHT               COLOR = COLOR_BTNHIGHLIGHT
	COLOR_BTNHILIGHT              COLOR = COLOR_BTNHIGHLIGHT
)

// Window class styles.
type CS uint32

const (
	CS_VREDRAW         CS = 0x0001
	CS_HREDRAW         CS = 0x0002
	CS_DBLCLKS         CS = 0x0008
	CS_OWNDC           CS = 0x0020
	CS_CLASSDC         CS = 0x0040
	CS_PARENTDC        CS = 0x0080
	CS_NOCLOSE         CS = 0x0200
	CS_SAVEBITS        CS = 0x0800
	CS_BYTEALIGNCLIENT CS = 0x1000
	CS_BYTEALIGNWINDOW CS = 0x2000
	CS_GLOBALCLASS     CS = 0x4000
	CS_IME             CS = 0x00010000
	CS_DROPSHADOW      CS = 0x00020000
)
