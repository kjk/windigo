package consts

type GA uint32 // GetAncestor

const (
	GA_PARENT    GA = 1
	GA_ROOT      GA = 2
	GA_ROOTOWNER GA = 3
)

type GDC int32 // GetDeviceCaps, originally constants have no prefix

const (
	GDC_DRIVERVERSION   GDC = 0
	GDC_TECHNOLOGY      GDC = 2
	GDC_HORZSIZE        GDC = 4
	GDC_VERTSIZE        GDC = 6
	GDC_HORZRES         GDC = 8
	GDC_VERTRES         GDC = 10
	GDC_BITSPIXEL       GDC = 12
	GDC_PLANES          GDC = 14
	GDC_NUMBRUSHES      GDC = 16
	GDC_NUMPENS         GDC = 18
	GDC_NUMMARKERS      GDC = 20
	GDC_NUMFONTS        GDC = 22
	GDC_NUMCOLORS       GDC = 24
	GDC_PDEVICESIZE     GDC = 26
	GDC_CURVECAPS       GDC = 28
	GDC_LINECAPS        GDC = 30
	GDC_POLYGONALCAPS   GDC = 32
	GDC_TEXTCAPS        GDC = 34
	GDC_CLIPCAPS        GDC = 36
	GDC_RASTERCAPS      GDC = 38
	GDC_ASPECTX         GDC = 40
	GDC_ASPECTY         GDC = 42
	GDC_ASPECTXY        GDC = 44
	GDC_LOGPIXELSX      GDC = 88
	GDC_LOGPIXELSY      GDC = 90
	GDC_SIZEPALETTE     GDC = 104
	GDC_NUMRESERVED     GDC = 106
	GDC_COLORRES        GDC = 108
	GDC_PHYSICALWIDTH   GDC = 110
	GDC_PHYSICALHEIGHT  GDC = 111
	GDC_PHYSICALOFFSETX GDC = 112
	GDC_PHYSICALOFFSETY GDC = 113
	GDC_SCALINGFACTORX  GDC = 114
	GDC_SCALINGFACTORY  GDC = 115
	GDC_VREFRESH        GDC = 116
	GDC_DESKTOPVERTRES  GDC = 117
	GDC_DESKTOPHORZRES  GDC = 118
	GDC_BLTALIGNMENT    GDC = 119
	GDC_SHADEBLENDCAPS  GDC = 120
	GDC_COLORMGMTCAPS   GDC = 121
)

type GW uint32 // GetWindow

const (
	GW_HWNDFIRST    GW = 0
	GW_HWNDLAST     GW = 1
	GW_HWNDNEXT     GW = 2
	GW_HWNDPREV     GW = 3
	GW_OWNER        GW = 4
	GW_CHILD        GW = 5
	GW_ENABLEDPOPUP GW = 6
	GW_MAX          GW = 6
)

type GWLP int32 // GetWindowLongPtr offsets

const (
	GWLP_STYLE      GWLP = -16
	GWLP_EXSTYLE    GWLP = -20
	GWLP_WNDPROC    GWLP = -4
	GWLP_HINSTANCE  GWLP = -6
	GWLP_HWNDPARENT GWLP = -8
	GWLP_USERDATA   GWLP = -21
	GWLP_ID         GWLP = -12
)

type HDM WM // list view header message

const (
	hDM_FIRST HDM = 0x1200

	HDM_GETITEMCOUNT HDM = hDM_FIRST + 0
	HDM_INSERTITEM   HDM = hDM_FIRST + 10
	HDM_DELETEITEM   HDM = hDM_FIRST + 2
	HDM_GETITEM      HDM = hDM_FIRST + 11
	HDM_SETITEM      HDM = hDM_FIRST + 12
	HDM_LAYOUT       HDM = hDM_FIRST + 5
)

type ID uint16 // dialog box command ID

const (
	IDOK       ID = 1
	IDCANCEL   ID = 2
	IDABORT    ID = 3
	IDRETRY    ID = 4
	IDIGNORE   ID = 5
	IDYES      ID = 6
	IDNO       ID = 7
	IDCLOSE    ID = 8
	IDHELP     ID = 9
	IDTRYAGAIN ID = 10
	IDCONTINUE ID = 11
	IDTIMEOUT  ID = 32000
)

type IDC uintptr // LoadCursor

const (
	IDC_ARROW       IDC = 32512
	IDC_IBEAM       IDC = 32513
	IDC_WAIT        IDC = 32514
	IDC_CROSS       IDC = 32515
	IDC_UPARROW     IDC = 32516
	IDC_SIZENWSE    IDC = 32642
	IDC_SIZENESW    IDC = 32643
	IDC_SIZEWE      IDC = 32644
	IDC_SIZENS      IDC = 32645
	IDC_SIZEALL     IDC = 32646
	IDC_NO          IDC = 32648
	IDC_HAND        IDC = 32649
	IDC_APPSTARTING IDC = 32650
	IDC_HELP        IDC = 32651
	IDC_PIN         IDC = 32671
	IDC_PERSON      IDC = 32672
)
