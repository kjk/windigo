/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// CreateFileMapping() flProtect.
type PAGE uint32

const (
	PAGE_NONE                   PAGE = 0
	PAGE_NOACCESS               PAGE = 0x01
	PAGE_READONLY               PAGE = 0x02
	PAGE_READWRITE              PAGE = 0x04
	PAGE_WRITECOPY              PAGE = 0x08
	PAGE_EXECUTE                PAGE = 0x10
	PAGE_EXECUTE_READ           PAGE = 0x20
	PAGE_EXECUTE_READWRITE      PAGE = 0x40
	PAGE_EXECUTE_WRITECOPY      PAGE = 0x80
	PAGE_GUARD                  PAGE = 0x100
	PAGE_NOCACHE                PAGE = 0x200
	PAGE_WRITECOMBINE           PAGE = 0x400
	PAGE_ENCLAVE_THREAD_CONTROL PAGE = 0x80000000
	PAGE_REVERT_TO_FILE_MAP     PAGE = 0x80000000
	PAGE_TARGETS_NO_UPDATE      PAGE = 0x40000000
	PAGE_TARGETS_INVALID        PAGE = 0x40000000
	PAGE_ENCLAVE_UNVALIDATED    PAGE = 0x20000000
	PAGE_ENCLAVE_DECOMMIT       PAGE = 0x10000000
)

// WM_POWERBROADCAST event.
type PBT uint32

const (
	PBT_APMQUERYSUSPEND       PBT = 0x0000
	PBT_APMQUERYSTANDBY       PBT = 0x0001
	PBT_APMQUERYSUSPENDFAILED PBT = 0x0002
	PBT_APMQUERYSTANDBYFAILED PBT = 0x0003
	PBT_APMSUSPEND            PBT = 0x0004
	PBT_APMSTANDBY            PBT = 0x0005
	PBT_APMRESUMECRITICAL     PBT = 0x0006
	PBT_APMRESUMESUSPEND      PBT = 0x0007
	PBT_APMRESUMESTANDBY      PBT = 0x0008
	PBT_APMBATTERYLOW         PBT = 0x0009
	PBT_APMPOWERSTATUSCHANGE  PBT = 0x000A
	PBT_APMOEMEVENT           PBT = 0x000B
	PBT_APMRESUMEAUTOMATIC    PBT = 0x0012
	PBT_POWERSETTINGCHANGE    PBT = 0x8013
)

// WM_PRINT drawing options.
type PRF uint32

const (
	PRF_CHECKVISIBLE PRF = 0x00000001
	PRF_NONCLIENT    PRF = 0x00000002
	PRF_CLIENT       PRF = 0x00000004
	PRF_ERASEBKGND   PRF = 0x00000008
	PRF_CHILDREN     PRF = 0x00000010
	PRF_OWNED        PRF = 0x00000020
)

// PolyDraw() aj.
type PT uint8

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)

// RegQueryValueEx() lpType.
type REG uint32

const (
	REG_NONE                       REG = 0 // No value type
	REG_SZ                         REG = 1 // Unicode nul terminated string
	REG_EXPAND_SZ                  REG = 2 // Unicode nul terminated string (with environment variable references)
	REG_BINARY                     REG = 3 // Free form binary
	REG_DWORD                      REG = 4 // 32-bit number
	REG_DWORD_LITTLE_ENDIAN        REG = 4 // 32-bit number (same as REG_DWORD)
	REG_DWORD_BIG_ENDIAN           REG = 5 // 32-bit number
	REG_LINK                       REG = 6 // Symbolic Link (unicode)
	REG_MULTI_SZ                   REG = 7 // Multiple Unicode strings
	REG_RESOURCE_LIST              REG = 8 // Resource list in the resource map
	REG_FULL_RESOURCE_DESCRIPTOR   REG = 9 // Resource list in the hardware description
	REG_RESOURCE_REQUIREMENTS_LIST REG = 10
	REG_QWORD                      REG = 11 // 64-bit number
	REG_QWORD_LITTLE_ENDIAN        REG = 11 // 64-bit number (same as REG_QWORD)
)

// RegOpenKeyEx() ulOptions.
type REG_OPTION uint32

const (
	REG_OPTION_NONE           REG_OPTION = 0
	REG_OPTION_RESERVED       REG_OPTION = 0x00000000
	REG_OPTION_NON_VOLATILE   REG_OPTION = 0x00000000
	REG_OPTION_VOLATILE       REG_OPTION = 0x00000001
	REG_OPTION_CREATE_LINK    REG_OPTION = 0x00000002
	REG_OPTION_BACKUP_RESTORE REG_OPTION = 0x00000004
	REG_OPTION_OPEN_LINK      REG_OPTION = 0x00000008
)

// IMAGELISTDRAWPARAMS dwRop.
type ROP uint32

const (
	ROP_SRCCOPY        ROP = 0x00CC0020
	ROP_SRCPAINT       ROP = 0x00EE0086
	ROP_SRCAND         ROP = 0x008800C6
	ROP_SRCINVERT      ROP = 0x00660046
	ROP_SRCERASE       ROP = 0x00440328
	ROP_NOTSRCCOPY     ROP = 0x00330008
	ROP_NOTSRCERASE    ROP = 0x001100A6
	ROP_MERGECOPY      ROP = 0x00C000CA
	ROP_MERGEPAINT     ROP = 0x00BB0226
	ROP_PATCOPY        ROP = 0x00F00021
	ROP_PATPAINT       ROP = 0x00FB0A09
	ROP_PATINVERT      ROP = 0x005A0049
	ROP_DSTINVERT      ROP = 0x00550009
	ROP_BLACKNESS      ROP = 0x00000042
	ROP_WHITENESS      ROP = 0x00FF0062
	ROP_NOMIRRORBITMAP ROP = 0x80000000
	ROP_CAPTUREBLT     ROP = 0x40000000
)
