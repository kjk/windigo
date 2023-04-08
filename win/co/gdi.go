//go:build windows

package co

// SetArcDirection() dir.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
type AD int32

const (
	AD_COUNTERCLOCKWISE AD = 1
	AD_CLOCKWISE        AD = 2
)

// BITMAPINFOHEADER biCompression.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BI uint32

const (
	BI_RGB       BI = 0
	BI_RLE8      BI = 1
	BI_RLE4      BI = 2
	BI_BITFIELDS BI = 3
	BI_JPEG      BI = 4
	BI_PNG       BI = 5
)

// SetBkMode() mode. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
type BKMODE int32

const (
	BKMODE_TRANSPARENT BKMODE = 1
	BKMODE_OPAQUE      BKMODE = 2
)

// LOGBRUSH lbStyle. Originally with BS prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type BRS uint32

const (
	BRS_SOLID         BRS = 0
	BRS_NULL          BRS = 1
	BRS_HOLLOW        BRS = BRS_NULL
	BRS_HATCHED       BRS = 2
	BRS_PATTERN       BRS = 3
	BRS_INDEXED       BRS = 4
	BRS_DIBPATTERN    BRS = 5
	BRS_DIBPATTERNPT  BRS = 6
	BRS_PATTERN8X8    BRS = 7
	BRS_DIBPATTERN8X8 BRS = 8
	BRS_MONOPATTERN   BRS = 9
)

// TEXTMETRIC tmCharSet. Originally with _CHARSET suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type CHARSET uint8

const (
	CHARSET_ANSI        CHARSET = 0
	CHARSET_DEFAULT     CHARSET = 1
	CHARSET_SYMBOL      CHARSET = 2
	CHARSET_SHIFTJIS    CHARSET = 128
	CHARSET_HANGUL      CHARSET = 129
	CHARSET_GB2312      CHARSET = 134
	CHARSET_CHINESEBIG5 CHARSET = 136
	CHARSET_OEM         CHARSET = 255
	CHARSET_JOHAB       CHARSET = 130
	CHARSET_HEBREW      CHARSET = 177
	CHARSET_ARABIC      CHARSET = 178
	CHARSET_GREEK       CHARSET = 161
	CHARSET_TURKISH     CHARSET = 162
	CHARSET_VIETNAMESE  CHARSET = 163
	CHARSET_THAI        CHARSET = 222
	CHARSET_EASTEUROPE  CHARSET = 238
	CHARSET_RUSSIAN     CHARSET = 204
	CHARSET_MAC         CHARSET = 77
	CHARSET_BALTIC      CHARSET = 186
)

// CreateDIBSection() usage.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
type DIB uint32

const (
	DIB_RGB_COLORS DIB = 0 // Color table in RGBs.
	DIB_PAL_COLORS DIB = 1 // Color table in palette indices.
)

// LOGFONT lfWeight.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type FW uint32

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT FW = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD   FW = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD  FW = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK      FW = FW_HEAVY
)

// GetDeviceCaps() index. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
type GDC int32

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

// CreateHatchBrush() iHatch.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
type HS int32

const (
	HS_HORIZONTAL HS = 0 // Pattern: -----
	HS_VERTICAL   HS = 1 // Pattern: |||||
	HS_FDIAGONAL  HS = 2 // Pattern: \\\\\
	HS_BDIAGONAL  HS = 3 // Pattern: /////
	HS_CROSS      HS = 4 // Pattern: +++++
	HS_DIAGCROSS  HS = 5 // Pattern: xxxxx
)

// SetPolyFillMode() mode. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
type POLYF int32

const (
	POLYF_ALTERNATE POLYF = 1
	POLYF_WINDING   POLYF = 2
)

// WM_PRINT drawing options.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-print
type PRF uint32

const (
	PRF_CHECKVISIBLE PRF = 0x0000_0001
	PRF_NONCLIENT    PRF = 0x0000_0002
	PRF_CLIENT       PRF = 0x0000_0004
	PRF_ERASEBKGND   PRF = 0x0000_0008
	PRF_CHILDREN     PRF = 0x0000_0010
	PRF_OWNED        PRF = 0x0000_0020
)

// CreatePen() iStyle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
type PS int32

const (
	PS_SOLID       PS = 0
	PS_DASH        PS = 1
	PS_DOT         PS = 2
	PS_DASHDOT     PS = 3
	PS_DASHDOTDOT  PS = 4
	PS_NULL        PS = 5
	PS_INSIDEFRAME PS = 6
)

// PolyDraw() aj.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
type PT uint8

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)

// SelectObject() return value. Originally with REGION suffix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
type REGION uint32

const (
	REGION_NULL    REGION = 1
	REGION_SIMPLE  REGION = 2
	REGION_COMPLEX REGION = 3
)

// CombineRgn() and SelectClipPath() mode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
type RGN int32

const (
	RGN_AND  RGN = 1
	RGN_OR   RGN = 2
	RGN_XOR  RGN = 3
	RGN_DIFF RGN = 4
	RGN_COPY RGN = 5
)

// BitBlt() rop, IMAGELISTDRAWPARAMS dwRop.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type ROP uint32

const (
	ROP_SRCCOPY        ROP = 0x00cc_0020
	ROP_SRCPAINT       ROP = 0x00ee_0086
	ROP_SRCAND         ROP = 0x0088_00c6
	ROP_SRCINVERT      ROP = 0x0066_0046
	ROP_SRCERASE       ROP = 0x0044_0328
	ROP_NOTSRCCOPY     ROP = 0x0033_0008
	ROP_NOTSRCERASE    ROP = 0x0011_00a6
	ROP_MERGECOPY      ROP = 0x00c0_00ca
	ROP_MERGEPAINT     ROP = 0x00bb_0226
	ROP_PATCOPY        ROP = 0x00f0_0021
	ROP_PATPAINT       ROP = 0x00fb_0a09
	ROP_PATINVERT      ROP = 0x005a_0049
	ROP_DSTINVERT      ROP = 0x0055_0009
	ROP_BLACKNESS      ROP = 0x0000_0042
	ROP_WHITENESS      ROP = 0x00ff_0062
	ROP_NOMIRRORBITMAP ROP = 0x8000_0000
	ROP_CAPTUREBLT     ROP = 0x4000_0000
)

// SetStretchBltMode() mode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
type STRETCH int32

const (
	STRETCH_BLACKONWHITE STRETCH = 1
	STRETCH_WHITEONBLACK STRETCH = 2
	STRETCH_COLORONCOLOR STRETCH = 3
	STRETCH_HALFTONE     STRETCH = 4
	STRETCH_ANDSCANS     STRETCH = STRETCH_BLACKONWHITE
	STRETCH_ORSCANS      STRETCH = STRETCH_WHITEONBLACK
	STRETCH_DELETESCANS  STRETCH = STRETCH_COLORONCOLOR
)

// SetTextAlign() align. Includes values with VTA prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
type TA uint32

const (
	TA_NOUPDATECP TA = 0
	TA_UPDATECP   TA = 1
	TA_LEFT       TA = 0
	TA_RIGHT      TA = 2
	TA_CENTER     TA = 6
	TA_TOP        TA = 0
	TA_BOTTOM     TA = 8
	TA_BASELINE   TA = 24
	TA_RTLREADING TA = 256
)

// TrackPopupMenu() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
type TPM uint32

const (
	TPM_LEFTBUTTON      TPM = 0x0000
	TPM_RIGHTBUTTON     TPM = 0x0002
	TPM_LEFTALIGN       TPM = 0x0000
	TPM_CENTERALIGN     TPM = 0x0004
	TPM_RIGHTALIGN      TPM = 0x0008
	TPM_TOPALIGN        TPM = 0x0000
	TPM_VCENTERALIGN    TPM = 0x0010
	TPM_BOTTOMALIGN     TPM = 0x0020
	TPM_HORIZONTAL      TPM = 0x0000
	TPM_VERTICAL        TPM = 0x0040
	TPM_NONOTIFY        TPM = 0x0080
	TPM_RETURNCMD       TPM = 0x0100
	TPM_RECURSE         TPM = 0x0001
	TPM_HORPOSANIMATION TPM = 0x0400
	TPM_HORNEGANIMATION TPM = 0x0800
	TPM_VERPOSANIMATION TPM = 0x1000
	TPM_VERNEGANIMATION TPM = 0x2000
	TPM_NOANIMATION     TPM = 0x4000
	TPM_LAYOUTRTL       TPM = 0x8000
	TPM_WORKAREA        TPM = 0x1_0000
)

// for StartDoc
type DOCINFO uint32

const (
	DI_APPBANDING            DOCINFO = 0x00000001
	DI_ROPS_READ_DESTINATION DOCINFO = 0x00000002
)

type DMORIENT uint32

const (
	DMORIENT_PORTRAIT  DMORIENT = 1
	DMORIENT_LANDSCAPE DMORIENT = 2
)

type DMPAPER uint32

const (
	DMPAPER_LETTER                        DMPAPER = 1   /* Letter 8 1/2 x 11 in               */
	DMPAPER_LETTERSMALL                   DMPAPER = 2   /* Letter Small 8 1/2 x 11 in         */
	DMPAPER_TABLOID                       DMPAPER = 3   /* Tabloid 11 x 17 in                 */
	DMPAPER_LEDGER                        DMPAPER = 4   /* Ledger 17 x 11 in                  */
	DMPAPER_LEGAL                         DMPAPER = 5   /* Legal 8 1/2 x 14 in                */
	DMPAPER_STATEMENT                     DMPAPER = 6   /* Statement 5 1/2 x 8 1/2 in         */
	DMPAPER_EXECUTIVE                     DMPAPER = 7   /* Executive 7 1/4 x 10 1/2 in        */
	DMPAPER_A3                            DMPAPER = 8   /* A3 297 x 420 mm                    */
	DMPAPER_A4                            DMPAPER = 9   /* A4 210 x 297 mm                    */
	DMPAPER_A4SMALL                       DMPAPER = 10  /* A4 Small 210 x 297 mm              */
	DMPAPER_A5                            DMPAPER = 11  /* A5 148 x 210 mm                    */
	DMPAPER_B4                            DMPAPER = 12  /* B4 (JIS) 250 x 354                 */
	DMPAPER_B5                            DMPAPER = 13  /* B5 (JIS) 182 x 257 mm              */
	DMPAPER_FOLIO                         DMPAPER = 14  /* Folio 8 1/2 x 13 in                */
	DMPAPER_QUARTO                        DMPAPER = 15  /* Quarto 215 x 275 mm                */
	DMPAPER_10X14                         DMPAPER = 16  /* 10x14 in                           */
	DMPAPER_11X17                         DMPAPER = 17  /* 11x17 in                           */
	DMPAPER_NOTE                          DMPAPER = 18  /* Note 8 1/2 x 11 in                 */
	DMPAPER_ENV_9                         DMPAPER = 19  /* Envelope #9 3 7/8 x 8 7/8          */
	DMPAPER_ENV_10                        DMPAPER = 20  /* Envelope #10 4 1/8 x 9 1/2         */
	DMPAPER_ENV_11                        DMPAPER = 21  /* Envelope #11 4 1/2 x 10 3/8        */
	DMPAPER_ENV_12                        DMPAPER = 22  /* Envelope #12 4 \276 x 11           */
	DMPAPER_ENV_14                        DMPAPER = 23  /* Envelope #14 5 x 11 1/2            */
	DMPAPER_CSHEET                        DMPAPER = 24  /* C size sheet                       */
	DMPAPER_DSHEET                        DMPAPER = 25  /* D size sheet                       */
	DMPAPER_ESHEET                        DMPAPER = 26  /* E size sheet                       */
	DMPAPER_ENV_DL                        DMPAPER = 27  /* Envelope DL 110 x 220mm            */
	DMPAPER_ENV_C5                        DMPAPER = 28  /* Envelope C5 162 x 229 mm           */
	DMPAPER_ENV_C3                        DMPAPER = 29  /* Envelope C3  324 x 458 mm          */
	DMPAPER_ENV_C4                        DMPAPER = 30  /* Envelope C4  229 x 324 mm          */
	DMPAPER_ENV_C6                        DMPAPER = 31  /* Envelope C6  114 x 162 mm          */
	DMPAPER_ENV_C65                       DMPAPER = 32  /* Envelope C65 114 x 229 mm          */
	DMPAPER_ENV_B4                        DMPAPER = 33  /* Envelope B4  250 x 353 mm          */
	DMPAPER_ENV_B5                        DMPAPER = 34  /* Envelope B5  176 x 250 mm          */
	DMPAPER_ENV_B6                        DMPAPER = 35  /* Envelope B6  176 x 125 mm          */
	DMPAPER_ENV_ITALY                     DMPAPER = 36  /* Envelope 110 x 230 mm              */
	DMPAPER_ENV_MONARCH                   DMPAPER = 37  /* Envelope Monarch 3.875 x 7.5 in    */
	DMPAPER_ENV_PERSONAL                  DMPAPER = 38  /* 6 3/4 Envelope 3 5/8 x 6 1/2 in    */
	DMPAPER_FANFOLD_US                    DMPAPER = 39  /* US Std Fanfold 14 7/8 x 11 in      */
	DMPAPER_FANFOLD_STD_GERMAN            DMPAPER = 40  /* German Std Fanfold 8 1/2 x 12 in   */
	DMPAPER_FANFOLD_LGL_GERMAN            DMPAPER = 41  /* German Legal Fanfold 8 1/2 x 13 in */
	DMPAPER_ISO_B4                        DMPAPER = 42  /* B4 (ISO) 250 x 353 mm              */
	DMPAPER_JAPANESE_POSTCARD             DMPAPER = 43  /* Japanese Postcard 100 x 148 mm     */
	DMPAPER_9X11                          DMPAPER = 44  /* 9 x 11 in                          */
	DMPAPER_10X11                         DMPAPER = 45  /* 10 x 11 in                         */
	DMPAPER_15X11                         DMPAPER = 46  /* 15 x 11 in                         */
	DMPAPER_ENV_INVITE                    DMPAPER = 47  /* Envelope Invite 220 x 220 mm       */
	DMPAPER_RESERVED_48                   DMPAPER = 48  /* RESERVED--DO NOT USE               */
	DMPAPER_RESERVED_49                   DMPAPER = 49  /* RESERVED--DO NOT USE               */
	DMPAPER_LETTER_EXTRA                  DMPAPER = 50  /* Letter Extra 9 \275 x 12 in        */
	DMPAPER_LEGAL_EXTRA                   DMPAPER = 51  /* Legal Extra 9 \275 x 15 in         */
	DMPAPER_TABLOID_EXTRA                 DMPAPER = 52  /* Tabloid Extra 11.69 x 18 in        */
	DMPAPER_A4_EXTRA                      DMPAPER = 53  /* A4 Extra 9.27 x 12.69 in           */
	DMPAPER_LETTER_TRANSVERSE             DMPAPER = 54  /* Letter Transverse 8 \275 x 11 in   */
	DMPAPER_A4_TRANSVERSE                 DMPAPER = 55  /* A4 Transverse 210 x 297 mm         */
	DMPAPER_LETTER_EXTRA_TRANSVERSE       DMPAPER = 56  /* Letter Extra Transverse 9\275 x 12 in */
	DMPAPER_A_PLUS                        DMPAPER = 57  /* SuperA/SuperA/A4 227 x 356 mm      */
	DMPAPER_B_PLUS                        DMPAPER = 58  /* SuperB/SuperB/A3 305 x 487 mm      */
	DMPAPER_LETTER_PLUS                   DMPAPER = 59  /* Letter Plus 8.5 x 12.69 in         */
	DMPAPER_A4_PLUS                       DMPAPER = 60  /* A4 Plus 210 x 330 mm               */
	DMPAPER_A5_TRANSVERSE                 DMPAPER = 61  /* A5 Transverse 148 x 210 mm         */
	DMPAPER_B5_TRANSVERSE                 DMPAPER = 62  /* B5 (JIS) Transverse 182 x 257 mm   */
	DMPAPER_A3_EXTRA                      DMPAPER = 63  /* A3 Extra 322 x 445 mm              */
	DMPAPER_A5_EXTRA                      DMPAPER = 64  /* A5 Extra 174 x 235 mm              */
	DMPAPER_B5_EXTRA                      DMPAPER = 65  /* B5 (ISO) Extra 201 x 276 mm        */
	DMPAPER_A2                            DMPAPER = 66  /* A2 420 x 594 mm                    */
	DMPAPER_A3_TRANSVERSE                 DMPAPER = 67  /* A3 Transverse 297 x 420 mm         */
	DMPAPER_A3_EXTRA_TRANSVERSE           DMPAPER = 68  /* A3 Extra Transverse 322 x 445 mm   */
	DMPAPER_DBL_JAPANESE_POSTCARD         DMPAPER = 69  /* Japanese Double Postcard 200 x 148 mm */
	DMPAPER_A6                            DMPAPER = 70  /* A6 105 x 148 mm                 */
	DMPAPER_JENV_KAKU2                    DMPAPER = 71  /* Japanese Envelope Kaku #2       */
	DMPAPER_JENV_KAKU3                    DMPAPER = 72  /* Japanese Envelope Kaku #3       */
	DMPAPER_JENV_CHOU3                    DMPAPER = 73  /* Japanese Envelope Chou #3       */
	DMPAPER_JENV_CHOU4                    DMPAPER = 74  /* Japanese Envelope Chou #4       */
	DMPAPER_LETTER_ROTATED                DMPAPER = 75  /* Letter Rotated 11 x 8 1/2 11 in */
	DMPAPER_A3_ROTATED                    DMPAPER = 76  /* A3 Rotated 420 x 297 mm         */
	DMPAPER_A4_ROTATED                    DMPAPER = 77  /* A4 Rotated 297 x 210 mm         */
	DMPAPER_A5_ROTATED                    DMPAPER = 78  /* A5 Rotated 210 x 148 mm         */
	DMPAPER_B4_JIS_ROTATED                DMPAPER = 79  /* B4 (JIS) Rotated 364 x 257 mm   */
	DMPAPER_B5_JIS_ROTATED                DMPAPER = 80  /* B5 (JIS) Rotated 257 x 182 mm   */
	DMPAPER_JAPANESE_POSTCARD_ROTATED     DMPAPER = 81  /* Japanese Postcard Rotated 148 x 100 mm */
	DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED DMPAPER = 82  /* Double Japanese Postcard Rotated 148 x 200 mm */
	DMPAPER_A6_ROTATED                    DMPAPER = 83  /* A6 Rotated 148 x 105 mm         */
	DMPAPER_JENV_KAKU2_ROTATED            DMPAPER = 84  /* Japanese Envelope Kaku #2 Rotated */
	DMPAPER_JENV_KAKU3_ROTATED            DMPAPER = 85  /* Japanese Envelope Kaku #3 Rotated */
	DMPAPER_JENV_CHOU3_ROTATED            DMPAPER = 86  /* Japanese Envelope Chou #3 Rotated */
	DMPAPER_JENV_CHOU4_ROTATED            DMPAPER = 87  /* Japanese Envelope Chou #4 Rotated */
	DMPAPER_B6_JIS                        DMPAPER = 88  /* B6 (JIS) 128 x 182 mm           */
	DMPAPER_B6_JIS_ROTATED                DMPAPER = 89  /* B6 (JIS) Rotated 182 x 128 mm   */
	DMPAPER_12X11                         DMPAPER = 90  /* 12 x 11 in                      */
	DMPAPER_JENV_YOU4                     DMPAPER = 91  /* Japanese Envelope You #4        */
	DMPAPER_JENV_YOU4_ROTATED             DMPAPER = 92  /* Japanese Envelope You #4 Rotated*/
	DMPAPER_P16K                          DMPAPER = 93  /* PRC 16K 146 x 215 mm            */
	DMPAPER_P32K                          DMPAPER = 94  /* PRC 32K 97 x 151 mm             */
	DMPAPER_P32KBIG                       DMPAPER = 95  /* PRC 32K(Big) 97 x 151 mm        */
	DMPAPER_PENV_1                        DMPAPER = 96  /* PRC Envelope #1 102 x 165 mm    */
	DMPAPER_PENV_2                        DMPAPER = 97  /* PRC Envelope #2 102 x 176 mm    */
	DMPAPER_PENV_3                        DMPAPER = 98  /* PRC Envelope #3 125 x 176 mm    */
	DMPAPER_PENV_4                        DMPAPER = 99  /* PRC Envelope #4 110 x 208 mm    */
	DMPAPER_PENV_5                        DMPAPER = 100 /* PRC Envelope #5 110 x 220 mm    */
	DMPAPER_PENV_6                        DMPAPER = 101 /* PRC Envelope #6 120 x 230 mm    */
	DMPAPER_PENV_7                        DMPAPER = 102 /* PRC Envelope #7 160 x 230 mm    */
	DMPAPER_PENV_8                        DMPAPER = 103 /* PRC Envelope #8 120 x 309 mm    */
	DMPAPER_PENV_9                        DMPAPER = 104 /* PRC Envelope #9 229 x 324 mm    */
	DMPAPER_PENV_10                       DMPAPER = 105 /* PRC Envelope #10 324 x 458 mm   */
	DMPAPER_P16K_ROTATED                  DMPAPER = 106 /* PRC 16K Rotated                 */
	DMPAPER_P32K_ROTATED                  DMPAPER = 107 /* PRC 32K Rotated                 */
	DMPAPER_P32KBIG_ROTATED               DMPAPER = 108 /* PRC 32K(Big) Rotated            */
	DMPAPER_PENV_1_ROTATED                DMPAPER = 109 /* PRC Envelope #1 Rotated 165 x 102 mm */
	DMPAPER_PENV_2_ROTATED                DMPAPER = 110 /* PRC Envelope #2 Rotated 176 x 102 mm */
	DMPAPER_PENV_3_ROTATED                DMPAPER = 111 /* PRC Envelope #3 Rotated 176 x 125 mm */
	DMPAPER_PENV_4_ROTATED                DMPAPER = 112 /* PRC Envelope #4 Rotated 208 x 110 mm */
	DMPAPER_PENV_5_ROTATED                DMPAPER = 113 /* PRC Envelope #5 Rotated 220 x 110 mm */
	DMPAPER_PENV_6_ROTATED                DMPAPER = 114 /* PRC Envelope #6 Rotated 230 x 120 mm */
	DMPAPER_PENV_7_ROTATED                DMPAPER = 115 /* PRC Envelope #7 Rotated 230 x 160 mm */
	DMPAPER_PENV_8_ROTATED                DMPAPER = 116 /* PRC Envelope #8 Rotated 309 x 120 mm */
	DMPAPER_PENV_9_ROTATED                DMPAPER = 117 /* PRC Envelope #9 Rotated 324 x 229 mm */
	DMPAPER_PENV_10_ROTATED               DMPAPER = 118 /* PRC Envelope #10 Rotated 458 x 324 mm */
)

type DMBIN uint32

const (
	DMBIN_UPPER         DMBIN = 1
	DMBIN_ONLYONE       DMBIN = 1
	DMBIN_LOWER         DMBIN = 2
	DMBIN_MIDDLE        DMBIN = 3
	DMBIN_MANUAL        DMBIN = 4
	DMBIN_ENVELOPE      DMBIN = 5
	DMBIN_ENVMANUAL     DMBIN = 6
	DMBIN_AUTO          DMBIN = 7
	DMBIN_TRACTOR       DMBIN = 8
	DMBIN_SMALLFMT      DMBIN = 9
	DMBIN_LARGEFMT      DMBIN = 10
	DMBIN_LARGECAPACITY DMBIN = 11
	DMBIN_CASSETTE      DMBIN = 14
	DMBIN_FORMSOURCE    DMBIN = 15
)

type MM int

const (
	MM_TEXT        MM = 1
	MM_LOMETRIC    MM = 2
	MM_HIMETRIC    MM = 3
	MM_LOENGLISH   MM = 4
	MM_HIENGLISH   MM = 5
	MM_TWIPS       MM = 6
	MM_ISOTROPIC   MM = 7
	MM_ANISOTROPIC MM = 8
)
