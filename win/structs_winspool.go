//go:build windows

package win

import (
	"github.com/kjk/windigo/win/co"
)

type DOCINFO struct {
	CbSize   int
	DocName  *uint16
	Output   *uint16
	Datatype *uint16
	Type     co.DOCINFO
}

// Page Range structure for PrintDlgEx
type PRINTPAGERANGE struct {
	FromPage uint32
	ToPage   uint32
}

// PrintDlgEx structure.
type PRINTDLGEXW struct {
	StructSize        uint32          // size of structure in bytes
	HwndOwner         HWND            // caller's window handle
	DevMode           HGLOBAL         // handle to DevMode
	DevNames          HGLOBAL         // handle to DevNames
	DC                HDC             // printer DC/IC or NULL
	Flags             uint32          // PD_ flags
	Flags2            uint32          // reserved
	ExclusionFlags    uint32          // items to exclude from driver pages
	NumPageRanges     uint32          // number of page ranges
	MaxPageRanges     uint32          // max number of page ranges
	PageRanges        *PRINTPAGERANGE // array of page ranges
	MinPage           uint32          // min page number
	MaxPage           uint32          // max page number
	Copies            uint32          // number of copies
	Instance          HINSTANCE       // instance handle
	PrintTemplateName *uint16         // template name for app specific area
	Callback          uintptr         // app callback interface
	NumPropertyPages  uint32          // number of app property pages in lphPropertyPages
	PropertyPages     uintptr         // array of app property page handles HPROPSHEETPAGE   *
	StartPage         uint32          // start page id
	DwResultAction    uint32          // result action if S_OK is returned
}

const CCHDEVICENAME = 32
const CCHFORMNAME = 32

type DEVMODE struct {
	deviceName    [CCHDEVICENAME]uint16
	SpecVersion   uint16
	DriverVersion uint16
	Size          uint16
	DriverExtra   uint16
	Fields        uint32

	/* printer only fields */
	PrinterFields struct {
		Orientation   int16
		PaperSize     int16
		PaperLength   int16
		PaperWidth    int16
		Scale         int16
		Copies        int16
		DefaultSource int16
		PrintQuality  int16
	}

	/* display only fields */
	// DisplayFields  struct {
	//     POINTL dmPosition;
	//     DWORD  dmDisplayOrientation;
	//     DWORD  dmDisplayFixedOutput;
	//   } DUMMYSTRUCTNAME2;
	// } DUMMYUNIONNAME;
	Color       int16
	Duplex      int16
	YResolution int16
	TTOption    int16
	Collate     int16
	formName    [CCHFORMNAME]uint16
	LogPixels   uint16
	BitsPerPel  uint32
	PelsWidth   uint32
	PelsHeight  uint32
	Dummy       struct {
		DisplayFlags uint32
		Nup          uint32
	}
	DisplayFrequency uint32
	ICMMethod        uint32
	ICMIntent        uint32
	MediaType        uint32
	DitherType       uint32
	Reserved1        uint32
	Reserved2        uint32
	PanningWidth     uint32
	PanningHeight    uint32
}

func (dm *DEVMODE) DeviceName() string { return Str.FromNativeSlice(dm.deviceName[:]) }
func (dm *DEVMODE) SetDeviceName(val string) {
	copy(dm.deviceName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dm.deviceName)-1)))
}

func (dm *DEVMODE) FormName() string { return Str.FromNativeSlice(dm.formName[:]) }
func (dm *DEVMODE) SetFormName(val string) {
	copy(dm.formName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(dm.formName)-1)))
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-options
type PRINTER_OPTIONS struct {
	CbSize uint
	Flags  PRINTER_OPTION
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/printer-defaults
type PRINTER_DEFAULTS struct {
	datatype      *uint16 // LPWSTR
	DevMode       uintptr // *DEVMODEW
	DesiredAccess uint32  // ACCESS_MASK
}

func (pd *PRINTER_DEFAULTS) DataType() string { return Str.FromNativePtr(pd.datatype) }
func (pd *PRINTER_DEFAULTS) SetFormName(val string) {
	pd.datatype = Str.ToNativePtr(val)
}
