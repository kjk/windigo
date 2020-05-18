/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package consts

type LIF uint32 // LITEM mask

const (
	LIF_ITEMINDEX LIF = 0x00000001
	LIF_STATE     LIF = 0x00000002
	LIF_ITEMID    LIF = 0x00000004
	LIF_URL       LIF = 0x00000008
)

type LIS uint32 // LITEM state

const (
	LIS_FOCUSED       LIS = 0x00000001
	LIS_ENABLED       LIS = 0x00000002
	LIS_VISITED       LIS = 0x00000004
	LIS_HOTTRACK      LIS = 0x00000008
	LIS_DEFAULTCOLORS LIS = 0x00000010
)

type LV_VIEW uint16 // list view view

const (
	LV_VIEW_ICON      LV_VIEW = 0x0000
	LV_VIEW_DETAILS   LV_VIEW = 0x0001
	LV_VIEW_SMALLICON LV_VIEW = 0x0002
	LV_VIEW_LIST      LV_VIEW = 0x0003
	LV_VIEW_TILE      LV_VIEW = 0x0004
	LV_VIEW_MAX       LV_VIEW = 0x0004
)

type LVCF uint32 // LVCOLUMN mask

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

type LVFI uint32 // LVFINDINFO

const (
	LVFI_PARAM     LVFI = 0x0001
	LVFI_STRING    LVFI = 0x0002
	LVFI_SUBSTRING LVFI = 0x0004
	LVFI_PARTIAL   LVFI = 0x0008
	LVFI_WRAP      LVFI = 0x0020
	LVFI_NEARESTXY LVFI = 0x0040
)

type LVGIT uint32 // NMLVGETINFOTIP

const (
	LVGIT_ZERO     LVGIT = 0x0000
	LVGIT_UNFOLDED LVGIT = 0x0001
)

type LVIF uint32 // LVITEM mask

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

type LVIS uint32 // list view item state

const (
	LVIS_FOCUSED        LVIS = 0x0001
	LVIS_SELECTED       LVIS = 0x0002
	LVIS_CUT            LVIS = 0x0004
	LVIS_DROPHILITED    LVIS = 0x0008
	LVIS_GLOW           LVIS = 0x0010
	LVIS_ACTIVATING     LVIS = 0x0020
	LVIS_OVERLAYMASK    LVIS = 0x0F00
	LVIS_STATEIMAGEMASK LVIS = 0xF000
)

type LVKF uint32 // NMITEMACTIVATE UKeyFlags

const (
	LVKF_ALT     LVKF = 0x0001
	LVKF_CONTROL LVKF = 0x0002
	LVKF_SHIFT   LVKF = 0x0004
)

type LVM WM // list view message

const (
	lVM_FIRST LVM = 0x1000

	LVM_GETBKCOLOR               LVM = lVM_FIRST + 0
	LVM_SETBKCOLOR               LVM = lVM_FIRST + 1
	LVM_GETIMAGELIST             LVM = lVM_FIRST + 2
	LVM_SETIMAGELIST             LVM = lVM_FIRST + 3
	LVM_GETITEMCOUNT             LVM = lVM_FIRST + 4
	LVM_DELETEITEM               LVM = lVM_FIRST + 8
	LVM_DELETEALLITEMS           LVM = lVM_FIRST + 9
	LVM_GETCOLUMNWIDTH           LVM = lVM_FIRST + 29
	LVM_SETCOLUMNWIDTH           LVM = lVM_FIRST + 30
	LVM_GETHEADER                LVM = lVM_FIRST + 31
	LVM_CREATEDRAGIMAGE          LVM = lVM_FIRST + 33
	LVM_GETVIEWRECT              LVM = lVM_FIRST + 34
	LVM_GETTEXTCOLOR             LVM = lVM_FIRST + 35
	LVM_SETTEXTCOLOR             LVM = lVM_FIRST + 36
	LVM_GETTEXTBKCOLOR           LVM = lVM_FIRST + 37
	LVM_SETTEXTBKCOLOR           LVM = lVM_FIRST + 38
	LVM_GETTOPINDEX              LVM = lVM_FIRST + 39
	LVM_GETCOUNTPERPAGE          LVM = lVM_FIRST + 40
	LVM_GETORIGIN                LVM = lVM_FIRST + 41
	LVM_UPDATE                   LVM = lVM_FIRST + 42
	LVM_SETITEMSTATE             LVM = lVM_FIRST + 43
	LVM_GETITEMSTATE             LVM = lVM_FIRST + 44
	LVM_SORTITEMS                LVM = lVM_FIRST + 48
	LVM_SETITEMPOSITION32        LVM = lVM_FIRST + 49
	LVM_GETSELECTEDCOUNT         LVM = lVM_FIRST + 50
	LVM_SETICONSPACING           LVM = lVM_FIRST + 53
	LVM_SETEXTENDEDLISTVIEWSTYLE LVM = lVM_FIRST + 54
	LVM_GETEXTENDEDLISTVIEWSTYLE LVM = lVM_FIRST + 55
	LVM_APPROXIMATEVIEWRECT      LVM = lVM_FIRST + 64
	LVM_SETTOOLTIPS              LVM = lVM_FIRST + 74
	LVM_GETITEM                  LVM = lVM_FIRST + 75
	LVM_SETITEM                  LVM = lVM_FIRST + 76
	LVM_INSERTITEM               LVM = lVM_FIRST + 77
	LVM_GETTOOLTIPS              LVM = lVM_FIRST + 78
	LVM_GETCOLUMN                LVM = lVM_FIRST + 95
	LVM_SETCOLUMN                LVM = lVM_FIRST + 96
	LVM_INSERTCOLUMN             LVM = lVM_FIRST + 97
	LVM_GETITEMTEXT              LVM = lVM_FIRST + 115
	LVM_SETITEMTEXT              LVM = lVM_FIRST + 116
	LVM_GETISEARCHSTRING         LVM = lVM_FIRST + 117
	LVM_SETBKIMAGE               LVM = lVM_FIRST + 138
	LVM_GETBKIMAGE               LVM = lVM_FIRST + 139
	LVM_SETVIEW                  LVM = lVM_FIRST + 142
	LVM_GETVIEW                  LVM = lVM_FIRST + 143
	LVM_HASGROUP                 LVM = lVM_FIRST + 161
	LVM_ISGROUPVIEWENABLED       LVM = lVM_FIRST + 175
	LVM_ISITEMVISIBLE            LVM = lVM_FIRST + 182
	LVM_GETEMPTYTEXT             LVM = lVM_FIRST + 204
)

type LVN NM // list view notification

const (
	lVN_FIRST LVN = -100

	LVN_ITEMCHANGING        LVN = lVN_FIRST - 0
	LVN_ITEMCHANGED         LVN = lVN_FIRST - 1
	LVN_INSERTITEM          LVN = lVN_FIRST - 2
	LVN_DELETEITEM          LVN = lVN_FIRST - 3
	LVN_DELETEALLITEMS      LVN = lVN_FIRST - 4
	LVN_BEGINLABELEDIT      LVN = lVN_FIRST - 75
	LVN_ENDLABELEDIT        LVN = lVN_FIRST - 76
	LVN_COLUMNCLICK         LVN = lVN_FIRST - 8
	LVN_BEGINDRAG           LVN = lVN_FIRST - 9
	LVN_BEGINRDRAG          LVN = lVN_FIRST - 11
	LVN_ODCACHEHINT         LVN = lVN_FIRST - 13
	LVN_ODFINDITEM          LVN = lVN_FIRST - 79
	LVN_ITEMACTIVATE        LVN = lVN_FIRST - 14
	LVN_ODSTATECHANGED      LVN = lVN_FIRST - 15
	LVN_HOTTRACK            LVN = lVN_FIRST - 21
	LVN_GETDISPINFO         LVN = lVN_FIRST - 77
	LVN_SETDISPINFO         LVN = lVN_FIRST - 78
	LVN_KEYDOWN             LVN = lVN_FIRST - 55
	LVN_MARQUEEBEGIN        LVN = lVN_FIRST - 56
	LVN_GETINFOTIP          LVN = lVN_FIRST - 58
	LVN_INCREMENTALSEARCH   LVN = lVN_FIRST - 63
	LVN_COLUMNDROPDOWN      LVN = lVN_FIRST - 64
	LVN_COLUMNOVERFLOWCLICK LVN = lVN_FIRST - 66
	LVN_BEGINSCROLL         LVN = lVN_FIRST - 80
	LVN_ENDSCROLL           LVN = lVN_FIRST - 81
	LVN_LINKCLICK           LVN = lVN_FIRST - 84
	LVN_GETEMPTYMARKUP      LVN = lVN_FIRST - 87
)

type LVS WS // list view style

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

type LVS_EX WS_EX // list view extended style

const (
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
