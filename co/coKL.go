/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// RegKeyOpenKeyEx() samDesired.
type KEY uint32

const (
	KEY_QUERY_VALUE        KEY = 0x0001
	KEY_SET_VALUE          KEY = 0x0002
	KEY_CREATE_SUB_KEY     KEY = 0x0004
	KEY_ENUMERATE_SUB_KEYS KEY = 0x0008
	KEY_NOTIFY             KEY = 0x0010
	KEY_CREATE_LINK        KEY = 0x0020
	KEY_WOW64_32KEY        KEY = 0x0200
	KEY_WOW64_64KEY        KEY = 0x0100
	KEY_WOW64_RES          KEY = 0x0300
	KEY_READ               KEY = (KEY(STANDARD_RIGHTS_READ) | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) &^ KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_WRITE              KEY = (KEY(STANDARD_RIGHTS_WRITE) | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) &^ KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_EXECUTE            KEY = KEY_READ &^ KEY(ACCESS_RIGHTS_SYNCHRONIZE)
	KEY_ALL_ACCESS         KEY = (KEY(STANDARD_RIGHTS_ALL) | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) &^ KEY(ACCESS_RIGHTS_SYNCHRONIZE)
)

// LITEM mask.
type LIF uint32

const (
	LIF_ITEMINDEX LIF = 0x00000001
	LIF_STATE     LIF = 0x00000002
	LIF_ITEMID    LIF = 0x00000004
	LIF_URL       LIF = 0x00000008
)

// LITEM state.
type LIS uint32

const (
	LIS_FOCUSED       LIS = 0x00000001
	LIS_ENABLED       LIS = 0x00000002
	LIS_VISITED       LIS = 0x00000004
	LIS_HOTTRACK      LIS = 0x00000008
	LIS_DEFAULTCOLORS LIS = 0x00000010
)

// SysLink control messages.
type LM WM

const (
	LM_HITTEST        LM = LM(WM_USER + 0x300)
	LM_GETIDEALHEIGHT LM = LM(WM_USER + 0x301)
	LM_SETITEM        LM = LM(WM_USER + 0x302)
	LM_GETITEM        LM = LM(WM_USER + 0x303)
	LM_GETIDEALSIZE   LM = LM_GETIDEALHEIGHT
)

// LockSetForegroundWindow uLockCode.
type LSFW uint32

const (
	LSFW_LOCK   LSFW = 1
	LSFW_UNLOCK LSFW = 2
)

// ListView current view.
type LV_VIEW uint16

const (
	LV_VIEW_ICON      LV_VIEW = 0x0000
	LV_VIEW_DETAILS   LV_VIEW = 0x0001
	LV_VIEW_SMALLICON LV_VIEW = 0x0002
	LV_VIEW_LIST      LV_VIEW = 0x0003
	LV_VIEW_TILE      LV_VIEW = 0x0004
)

// NMLVCUSTOMDRAW dwItemType.
type LVCDI uint32

const (
	LVCDI_ITEM     LVCDI = 0x00000000
	LVCDI_GROUP    LVCDI = 0x00000001
	LVCDI_TEMSLIST LVCDI = 0x00000002
)

// LVCOLUMN mask.
type LVCF uint32

const (
	LVCF_DEFAULTWIDTH LVCF = 0x0080
	LVCF_FMT          LVCF = 0x0001
	LVCF_IDEALWIDTH   LVCF = 0x0100
	LVCF_IMAGE        LVCF = 0x0010
	LVCF_MINWIDTH     LVCF = 0x0040
	LVCF_ORDER        LVCF = 0x0020
	LVCF_SUBITEM      LVCF = 0x0008
	LVCF_TEXT         LVCF = 0x0004
	LVCF_WIDTH        LVCF = 0x0002
)

// LVFINDINFO flags.
type LVFI uint32

const (
	LVFI_PARAM     LVFI = 0x0001
	LVFI_STRING    LVFI = 0x0002
	LVFI_SUBSTRING LVFI = 0x0004
	LVFI_PARTIAL   LVFI = 0x0008
	LVFI_WRAP      LVFI = 0x0020
	LVFI_NEARESTXY LVFI = 0x0040
)

// NMLVCUSTOMDRAW uAlign.
type LVGA_HEADER uint32

const (
	LVGA_HEADER_LEFT   LVGA_HEADER = 0x00000001
	LVGA_HEADER_CENTER LVGA_HEADER = 0x00000002
	LVGA_HEADER_RIGHT  LVGA_HEADER = 0x00000004
)

// NMLVGETINFOTIP dwFlags.
type LVGIT uint32

const (
	LVGIT_ZERO     LVGIT = 0x0000
	LVGIT_UNFOLDED LVGIT = 0x0001
)

// LVHITTESTINFO flags.
type LVHT uint32

const (
	LVHT_NOWHERE             LVHT = 0x00000001
	LVHT_ONITEMICON          LVHT = 0x00000002
	LVHT_ONITEMLABEL         LVHT = 0x00000004
	LVHT_ONITEMSTATEICON     LVHT = 0x00000008
	LVHT_ONITEM              LVHT = LVHT_ONITEMICON | LVHT_ONITEMLABEL | LVHT_ONITEMSTATEICON
	LVHT_ABOVE               LVHT = 0x00000008
	LVHT_BELOW               LVHT = 0x00000010
	LVHT_TORIGHT             LVHT = 0x00000020
	LVHT_TOLEFT              LVHT = 0x00000040
	LVHT_EX_GROUP_HEADER     LVHT = 0x10000000
	LVHT_EX_GROUP_FOOTER     LVHT = 0x20000000
	LVHT_EX_GROUP_COLLAPSE   LVHT = 0x40000000
	LVHT_EX_GROUP_BACKGROUND LVHT = 0x80000000
	LVHT_EX_GROUP_STATEICON  LVHT = 0x01000000
	LVHT_EX_GROUP_SUBSETLINK LVHT = 0x02000000
	LVHT_EX_GROUP            LVHT = LVHT_EX_GROUP_BACKGROUND | LVHT_EX_GROUP_COLLAPSE | LVHT_EX_GROUP_FOOTER | LVHT_EX_GROUP_HEADER | LVHT_EX_GROUP_STATEICON | LVHT_EX_GROUP_SUBSETLINK
	LVHT_EX_ONCONTENTS       LVHT = 0x04000000
	LVHT_EX_FOOTER           LVHT = 0x08000000
)

// LVITEM mask.
type LVIF uint32

const (
	LVIF_COLFMT      LVIF = 0x00010000
	LVIF_COLUMNS     LVIF = 0x00000200
	LVIF_GROUPID     LVIF = 0x00000100
	LVIF_IMAGE       LVIF = 0x00000002
	LVIF_INDENT      LVIF = 0x00000010
	LVIF_NORECOMPUTE LVIF = 0x00000800
	LVIF_PARAM       LVIF = 0x00000004
	LVIF_STATE       LVIF = 0x00000008
	LVIF_TEXT        LVIF = 0x00000001
)

// ListView item states.
type LVIS uint32

const (
	LVIS_NONE           LVIS = 0
	LVIS_FOCUSED        LVIS = 0x0001
	LVIS_SELECTED       LVIS = 0x0002
	LVIS_CUT            LVIS = 0x0004
	LVIS_DROPHILITED    LVIS = 0x0008
	LVIS_GLOW           LVIS = 0x0010
	LVIS_ACTIVATING     LVIS = 0x0020
	LVIS_OVERLAYMASK    LVIS = 0x0f00
	LVIS_STATEIMAGEMASK LVIS = 0xf000
)

// NMITEMACTIVATE uKeyFlags.
type LVKF uint32

const (
	LVKF_ALT     LVKF = 0x0001
	LVKF_CONTROL LVKF = 0x0002
	LVKF_SHIFT   LVKF = 0x0004
)

// ListView messages.
type LVM WM

const (
	_LVM_FIRST LVM = 0x1000

	LVM_GETBKCOLOR               LVM = _LVM_FIRST + 0
	LVM_SETBKCOLOR               LVM = _LVM_FIRST + 1
	LVM_GETIMAGELIST             LVM = _LVM_FIRST + 2
	LVM_SETIMAGELIST             LVM = _LVM_FIRST + 3
	LVM_GETITEMCOUNT             LVM = _LVM_FIRST + 4
	LVM_DELETEITEM               LVM = _LVM_FIRST + 8
	LVM_DELETEALLITEMS           LVM = _LVM_FIRST + 9
	LVM_GETCALLBACKMASK          LVM = _LVM_FIRST + 10
	LVM_SETCALLBACKMASK          LVM = _LVM_FIRST + 11
	LVM_GETNEXTITEM              LVM = _LVM_FIRST + 12
	LVM_GETITEMRECT              LVM = _LVM_FIRST + 14
	LVM_SETITEMPOSITION          LVM = _LVM_FIRST + 15
	LVM_GETITEMPOSITION          LVM = _LVM_FIRST + 16
	LVM_HITTEST                  LVM = _LVM_FIRST + 18
	LVM_ENSUREVISIBLE            LVM = _LVM_FIRST + 19
	LVM_SCROLL                   LVM = _LVM_FIRST + 20
	LVM_REDRAWITEMS              LVM = _LVM_FIRST + 21
	LVM_ARRANGE                  LVM = _LVM_FIRST + 22
	LVM_GETEDITCONTROL           LVM = _LVM_FIRST + 24
	LVM_DELETECOLUMN             LVM = _LVM_FIRST + 28
	LVM_GETCOLUMNWIDTH           LVM = _LVM_FIRST + 29
	LVM_SETCOLUMNWIDTH           LVM = _LVM_FIRST + 30
	LVM_GETHEADER                LVM = _LVM_FIRST + 31
	LVM_CREATEDRAGIMAGE          LVM = _LVM_FIRST + 33
	LVM_GETVIEWRECT              LVM = _LVM_FIRST + 34
	LVM_GETTEXTCOLOR             LVM = _LVM_FIRST + 35
	LVM_SETTEXTCOLOR             LVM = _LVM_FIRST + 36
	LVM_GETTEXTBKCOLOR           LVM = _LVM_FIRST + 37
	LVM_SETTEXTBKCOLOR           LVM = _LVM_FIRST + 38
	LVM_GETTOPINDEX              LVM = _LVM_FIRST + 39
	LVM_GETCOUNTPERPAGE          LVM = _LVM_FIRST + 40
	LVM_GETORIGIN                LVM = _LVM_FIRST + 41
	LVM_UPDATE                   LVM = _LVM_FIRST + 42
	LVM_SETITEMSTATE             LVM = _LVM_FIRST + 43
	LVM_GETITEMSTATE             LVM = _LVM_FIRST + 44
	LVM_SETITEMCOUNT             LVM = _LVM_FIRST + 47
	LVM_SORTITEMS                LVM = _LVM_FIRST + 48
	LVM_SETITEMPOSITION32        LVM = _LVM_FIRST + 49
	LVM_GETSELECTEDCOUNT         LVM = _LVM_FIRST + 50
	LVM_GETITEMSPACING           LVM = _LVM_FIRST + 51
	LVM_SETICONSPACING           LVM = _LVM_FIRST + 53
	LVM_SETEXTENDEDLISTVIEWSTYLE LVM = _LVM_FIRST + 54
	LVM_GETEXTENDEDLISTVIEWSTYLE LVM = _LVM_FIRST + 55
	LVM_GETSUBITEMRECT           LVM = _LVM_FIRST + 56
	LVM_SUBITEMHITTEST           LVM = _LVM_FIRST + 57
	LVM_SETCOLUMNORDERARRAY      LVM = _LVM_FIRST + 58
	LVM_GETCOLUMNORDERARRAY      LVM = _LVM_FIRST + 59
	LVM_SETHOTITEM               LVM = _LVM_FIRST + 60
	LVM_GETHOTITEM               LVM = _LVM_FIRST + 61
	LVM_SETHOTCURSOR             LVM = _LVM_FIRST + 62
	LVM_GETHOTCURSOR             LVM = _LVM_FIRST + 63
	LVM_APPROXIMATEVIEWRECT      LVM = _LVM_FIRST + 64
	LVM_SETWORKAREAS             LVM = _LVM_FIRST + 65
	LVM_GETSELECTIONMARK         LVM = _LVM_FIRST + 66
	LVM_SETSELECTIONMARK         LVM = _LVM_FIRST + 67
	LVM_GETWORKAREAS             LVM = _LVM_FIRST + 70
	LVM_SETHOVERTIME             LVM = _LVM_FIRST + 71
	LVM_GETHOVERTIME             LVM = _LVM_FIRST + 72
	LVM_GETNUMBEROFWORKAREAS     LVM = _LVM_FIRST + 73
	LVM_SETTOOLTIPS              LVM = _LVM_FIRST + 74
	LVM_GETITEM                  LVM = _LVM_FIRST + 75
	LVM_SETITEM                  LVM = _LVM_FIRST + 76
	LVM_INSERTITEM               LVM = _LVM_FIRST + 77
	LVM_GETTOOLTIPS              LVM = _LVM_FIRST + 78
	LVM_SORTITEMSEX              LVM = _LVM_FIRST + 81
	LVM_FINDITEM                 LVM = _LVM_FIRST + 83
	LVM_GETSTRINGWIDTH           LVM = _LVM_FIRST + 87
	LVM_GETGROUPSTATE            LVM = _LVM_FIRST + 92
	LVM_GETFOCUSEDGROUP          LVM = _LVM_FIRST + 93
	LVM_GETCOLUMN                LVM = _LVM_FIRST + 95
	LVM_SETCOLUMN                LVM = _LVM_FIRST + 96
	LVM_INSERTCOLUMN             LVM = _LVM_FIRST + 97
	LVM_GETGROUPRECT             LVM = _LVM_FIRST + 98
	LVM_GETITEMTEXT              LVM = _LVM_FIRST + 115
	LVM_SETITEMTEXT              LVM = _LVM_FIRST + 116
	LVM_GETISEARCHSTRING         LVM = _LVM_FIRST + 117
	LVM_EDITLABEL                LVM = _LVM_FIRST + 118
	LVM_SETBKIMAGE               LVM = _LVM_FIRST + 138
	LVM_GETBKIMAGE               LVM = _LVM_FIRST + 139
	LVM_SETSELECTEDCOLUMN        LVM = _LVM_FIRST + 140
	LVM_SETVIEW                  LVM = _LVM_FIRST + 142
	LVM_GETVIEW                  LVM = _LVM_FIRST + 143
	LVM_INSERTGROUP              LVM = _LVM_FIRST + 145
	LVM_SETGROUPINFO             LVM = _LVM_FIRST + 147
	LVM_GETGROUPINFO             LVM = _LVM_FIRST + 149
	LVM_REMOVEGROUP              LVM = _LVM_FIRST + 150
	LVM_MOVEGROUP                LVM = _LVM_FIRST + 151
	LVM_GETGROUPCOUNT            LVM = _LVM_FIRST + 152
	LVM_GETGROUPINFOBYINDEX      LVM = _LVM_FIRST + 153
	LVM_MOVEITEMTOGROUP          LVM = _LVM_FIRST + 154
	LVM_SETGROUPMETRICS          LVM = _LVM_FIRST + 155
	LVM_GETGROUPMETRICS          LVM = _LVM_FIRST + 156
	LVM_ENABLEGROUPVIEW          LVM = _LVM_FIRST + 157
	LVM_SORTGROUPS               LVM = _LVM_FIRST + 158
	LVM_INSERTGROUPSORTED        LVM = _LVM_FIRST + 159
	LVM_REMOVEALLGROUPS          LVM = _LVM_FIRST + 160
	LVM_HASGROUP                 LVM = _LVM_FIRST + 161
	LVM_SETTILEVIEWINFO          LVM = _LVM_FIRST + 162
	LVM_GETTILEVIEWINFO          LVM = _LVM_FIRST + 163
	LVM_SETTILEINFO              LVM = _LVM_FIRST + 164
	LVM_GETTILEINFO              LVM = _LVM_FIRST + 165
	LVM_SETINSERTMARK            LVM = _LVM_FIRST + 166
	LVM_GETINSERTMARK            LVM = _LVM_FIRST + 167
	LVM_INSERTMARKHITTEST        LVM = _LVM_FIRST + 168
	LVM_GETINSERTMARKRECT        LVM = _LVM_FIRST + 169
	LVM_SETINSERTMARKCOLOR       LVM = _LVM_FIRST + 170
	LVM_GETINSERTMARKCOLOR       LVM = _LVM_FIRST + 171
	LVM_SETINFOTIP               LVM = _LVM_FIRST + 173
	LVM_GETSELECTEDCOLUMN        LVM = _LVM_FIRST + 174
	LVM_ISGROUPVIEWENABLED       LVM = _LVM_FIRST + 175
	LVM_GETOUTLINECOLOR          LVM = _LVM_FIRST + 176
	LVM_SETOUTLINECOLOR          LVM = _LVM_FIRST + 177
	LVM_CANCELEDITLABEL          LVM = _LVM_FIRST + 179
	LVM_MAPINDEXTOID             LVM = _LVM_FIRST + 180
	LVM_MAPIDTOINDEX             LVM = _LVM_FIRST + 181
	LVM_ISITEMVISIBLE            LVM = _LVM_FIRST + 182
	LVM_GETEMPTYTEXT             LVM = _LVM_FIRST + 204
	LVM_GETFOOTERRECT            LVM = _LVM_FIRST + 205
	LVM_GETFOOTERINFO            LVM = _LVM_FIRST + 206
	LVM_GETFOOTERITEMRECT        LVM = _LVM_FIRST + 207
	LVM_GETFOOTERITEM            LVM = _LVM_FIRST + 208
	LVM_GETITEMINDEXRECT         LVM = _LVM_FIRST + 209
	LVM_SETITEMINDEXSTATE        LVM = _LVM_FIRST + 210
	LVM_GETNEXTITEMINDEX         LVM = _LVM_FIRST + 211
)

// ListView notifications, sent via WM_NOTIFY.
type LVN NM

const (
	_LVN_FIRST LVN = -100

	LVN_ITEMCHANGING        LVN = _LVN_FIRST - 0
	LVN_ITEMCHANGED         LVN = _LVN_FIRST - 1
	LVN_INSERTITEM          LVN = _LVN_FIRST - 2
	LVN_DELETEITEM          LVN = _LVN_FIRST - 3
	LVN_DELETEALLITEMS      LVN = _LVN_FIRST - 4
	LVN_BEGINLABELEDIT      LVN = _LVN_FIRST - 75
	LVN_ENDLABELEDIT        LVN = _LVN_FIRST - 76
	LVN_COLUMNCLICK         LVN = _LVN_FIRST - 8
	LVN_BEGINDRAG           LVN = _LVN_FIRST - 9
	LVN_BEGINRDRAG          LVN = _LVN_FIRST - 11
	LVN_ODCACHEHINT         LVN = _LVN_FIRST - 13
	LVN_ODFINDITEM          LVN = _LVN_FIRST - 79
	LVN_ITEMACTIVATE        LVN = _LVN_FIRST - 14
	LVN_ODSTATECHANGED      LVN = _LVN_FIRST - 15
	LVN_HOTTRACK            LVN = _LVN_FIRST - 21
	LVN_GETDISPINFO         LVN = _LVN_FIRST - 77
	LVN_SETDISPINFO         LVN = _LVN_FIRST - 78
	LVN_KEYDOWN             LVN = _LVN_FIRST - 55
	LVN_MARQUEEBEGIN        LVN = _LVN_FIRST - 56
	LVN_GETINFOTIP          LVN = _LVN_FIRST - 58
	LVN_INCREMENTALSEARCH   LVN = _LVN_FIRST - 63
	LVN_COLUMNDROPDOWN      LVN = _LVN_FIRST - 64
	LVN_COLUMNOVERFLOWCLICK LVN = _LVN_FIRST - 66
	LVN_BEGINSCROLL         LVN = _LVN_FIRST - 80
	LVN_ENDSCROLL           LVN = _LVN_FIRST - 81
	LVN_LINKCLICK           LVN = _LVN_FIRST - 84
	LVN_GETEMPTYMARKUP      LVN = _LVN_FIRST - 87
)

// LVM_GETNEXTITEM item relationship.
type LVNI uint32

const (
	LVNI_ALL           LVNI = 0x0000
	LVNI_FOCUSED       LVNI = 0x0001
	LVNI_SELECTED      LVNI = 0x0002
	LVNI_CUT           LVNI = 0x0004
	LVNI_DROPHILITED   LVNI = 0x0008
	LVNI_STATEMASK     LVNI = LVNI_FOCUSED | LVNI_SELECTED | LVNI_CUT | LVNI_DROPHILITED
	LVNI_VISIBLEORDER  LVNI = 0x0010
	LVNI_PREVIOUS      LVNI = 0x0020
	LVNI_VISIBLEONLY   LVNI = 0x0040
	LVNI_SAMEGROUPONLY LVNI = 0x0080
	LVNI_ABOVE         LVNI = 0x0100
	LVNI_BELOW         LVNI = 0x0200
	LVNI_TOLEFT        LVNI = 0x0400
	LVNI_TORIGHT       LVNI = 0x0800
	LVNI_DIRECTIONMASK LVNI = LVNI_ABOVE | LVNI_BELOW | LVNI_TOLEFT | LVNI_TORIGHT
)

// LVM_GETITEMRECT portion.
type LVIR uint32

const (
	LVIR_BOUNDS       LVIR = 0
	LVIR_ICON         LVIR = 1
	LVIR_LABEL        LVIR = 2
	LVIR_SELECTBOUNDS LVIR = 3
)

// ListView styles.
type LVS WS

const (
	LVS_ALIGNLEFT       LVS = 0x0800
	LVS_ALIGNMASK       LVS = 0x0c00
	LVS_ALIGNTOP        LVS = 0x0000
	LVS_AUTOARRANGE     LVS = 0x0100
	LVS_EDITLABELS      LVS = 0x0200
	LVS_ICON            LVS = 0x0000
	LVS_LIST            LVS = 0x0003
	LVS_NOCOLUMNHEADER  LVS = 0x4000
	LVS_NOLABELWRAP     LVS = 0x0080
	LVS_NOSCROLL        LVS = 0x2000
	LVS_NOSORTHEADER    LVS = 0x8000
	LVS_OWNERDATA       LVS = 0x1000
	LVS_OWNERDRAWFIXED  LVS = 0x0400
	LVS_REPORT          LVS = 0x0001
	LVS_SHAREIMAGELISTS LVS = 0x0040
	LVS_SHOWSELALWAYS   LVS = 0x0008
	LVS_SINGLESEL       LVS = 0x0004
	LVS_SMALLICON       LVS = 0x0002
	LVS_SORTASCENDING   LVS = 0x0010
	LVS_SORTDESCENDING  LVS = 0x0020
	LVS_TYPEMASK        LVS = 0x0003
	LVS_TYPESTYLEMASK   LVS = 0xfc00
)

// ListView extended styles.
type LVS_EX WS_EX

const (
	LVS_EX_NONE                  LVS_EX = 0
	LVS_EX_AUTOAUTOARRANGE       LVS_EX = 0x01000000
	LVS_EX_AUTOCHECKSELECT       LVS_EX = 0x08000000
	LVS_EX_AUTOSIZECOLUMNS       LVS_EX = 0x10000000
	LVS_EX_BORDERSELECT          LVS_EX = 0x00008000
	LVS_EX_CHECKBOXES            LVS_EX = 0x00000004
	LVS_EX_COLUMNOVERFLOW        LVS_EX = 0x80000000
	LVS_EX_COLUMNSNAPPOINTS      LVS_EX = 0x40000000
	LVS_EX_DOUBLEBUFFER          LVS_EX = 0x00010000
	LVS_EX_FLATSB                LVS_EX = 0x00000100
	LVS_EX_FULLROWSELECT         LVS_EX = 0x00000020
	LVS_EX_GRIDLINES             LVS_EX = 0x00000001
	LVS_EX_HEADERDRAGDROP        LVS_EX = 0x00000010
	LVS_EX_HEADERINALLVIEWS      LVS_EX = 0x02000000
	LVS_EX_HIDELABELS            LVS_EX = 0x00020000
	LVS_EX_INFOTIP               LVS_EX = 0x00000400
	LVS_EX_JUSTIFYCOLUMNS        LVS_EX = 0x00200000
	LVS_EX_LABELTIP              LVS_EX = 0x00004000
	LVS_EX_MULTIWORKAREAS        LVS_EX = 0x00002000
	LVS_EX_ONECLICKACTIVATE      LVS_EX = 0x00000040
	LVS_EX_REGIONAL              LVS_EX = 0x00000200
	LVS_EX_SIMPLESELECT          LVS_EX = 0x00100000
	LVS_EX_SINGLEROW             LVS_EX = 0x00040000
	LVS_EX_SNAPTOGRID            LVS_EX = 0x00080000
	LVS_EX_SUBITEMIMAGES         LVS_EX = 0x00000002
	LVS_EX_TRACKSELECT           LVS_EX = 0x00000008
	LVS_EX_TRANSPARENTBKGND      LVS_EX = 0x00400000
	LVS_EX_TRANSPARENTSHADOWTEXT LVS_EX = 0x00800000
	LVS_EX_TWOCLICKACTIVATE      LVS_EX = 0x00000080
	LVS_EX_UNDERLINECOLD         LVS_EX = 0x00001000
	LVS_EX_UNDERLINEHOT          LVS_EX = 0x00000800
)

// LVM_SETIMAGELIST type.
type LVSIL uint32

const (
	LVSIL_NORMAL      LVSIL = 0
	LVSIL_SMALL       LVSIL = 1
	LVSIL_STATE       LVSIL = 2
	LVSIL_GROUPHEADER LVSIL = 3
)

// SysLink control styles.
type LWS WS

const (
	LWS_TRANSPARENT    LWS = 0x0001
	LWS_IGNORERETURN   LWS = 0x0002
	LWS_NOPREFIX       LWS = 0x0004
	LWS_USEVISUALSTYLE LWS = 0x0008
	LWS_USECUSTOMTEXT  LWS = 0x0010
	LWS_RIGHT          LWS = 0x0020
)