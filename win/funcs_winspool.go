//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

type PRINTER_INFO_5W struct {
	PrinterName              *uint16
	PortName                 *uint16
	Attributes               co.PRINTER_ATTRIBUTE
	DeviceNotSelectedTimeout uint32
	TransmissionRetryTimeout uint32
}

func fromPtr[T any](v *T) uintptr {
	if v == nil {
		return 0
	}
	return uintptr(unsafe.Pointer(v))
}

func fromBuf[T any](v []T) uintptr {
	return uintptr(unsafe.Pointer(&v[0]))
}

func fromUtf8(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(Str.ToNativePtr(s)))
}

type PrinterInfo struct {
	PrinterName              string
	PortName                 string
	Attributes               co.PRINTER_ATTRIBUTE
	DeviceNotSelectedTimeout uint32
	TransmissionRetryTimeout uint32
}

func EnumPrinters(dwFlags co.PRINTER_ENUM) ([]PrinterInfo, error) {
	var cbNeeded uint32
	var nPrinters uint32
	addr := proc.EnumPrinters.Addr()
	syscall.SyscallN(addr, uintptr(dwFlags), 0, 5, 0, 0, fromPtr(&cbNeeded), fromPtr(&nPrinters))
	buf := make([]byte, cbNeeded)
	cbBuf := cbNeeded
	res, _, err := syscall.SyscallN(addr, uintptr(dwFlags), 0, 5, fromBuf(buf), uintptr(cbBuf), fromPtr(&cbNeeded), fromPtr(&nPrinters))
	if res == 0 {
		// TODO: errco.ERROR(res) ?
		return nil, errco.ERROR(err)
	}
	if nPrinters == 0 {
		return nil, nil
	}
	nSize := unsafe.Sizeof(PRINTER_INFO_5W{})
	var printers []PrinterInfo
	for i := 0; i < int(nPrinters); i++ {
		pir := (*PRINTER_INFO_5W)(unsafe.Pointer(&buf[int(nSize)*i]))
		var pi PrinterInfo
		pi.PrinterName = Str.FromNativePtr(pir.PrinterName)
		pi.PortName = Str.FromNativePtr(pir.PortName)
		pi.Attributes = pir.Attributes
		pi.DeviceNotSelectedTimeout = pir.DeviceNotSelectedTimeout
		pi.TransmissionRetryTimeout = pir.TransmissionRetryTimeout
		printers = append(printers, pi)
	}
	return printers, nil
}

func GetDefaultPrinter() (string, error) {
	var buf [512 + 1]uint16
	var lenInOut = uint32(len(buf))
	addr := proc.GetDefaultPrinter.Addr()
	ret, _, _ := syscall.SyscallN(addr,
		fromBuf(buf[:]), fromPtr(&lenInOut))
	// TODO: handle ret == ERROR_INSUFFICIENT_BUFFER and ERROR_FILE_NOT_FOUND
	// or return an error
	if ret == 0 {
		return "", errco.ERROR(ret)
	}
	return Str.FromNativeSlice(buf[:]), nil

}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-devicecapabilitiesw

func DeviceCapabilitiesWORD(device string, port string, dc co.DC) ([]uint16, error) {
	addr := proc.DeviceCapabilities.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), 0, 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	if ret == 0 {
		return nil, nil
	}
	buf := make([]uint16, ret)
	ret, _, _ = syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), fromBuf(buf), 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	return buf, nil
}

func DeviceCapabilitiesPaperSize(device string, port string) ([]POINT, error) {
	addr := proc.DeviceCapabilities.Addr()
	dc := co.DC_PAPERSIZE
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), 0, 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	if ret == 0 {
		return nil, nil
	}
	buf := make([]POINT, ret)
	ret, _, _ = syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), fromBuf(buf), 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	return buf, nil
}

func DeviceCapabilitiesBins(device string, port string) ([]uint16, error) {
	return DeviceCapabilitiesWORD(device, port, co.DC_BINS)
}

func DeviceCapabilitiesPapers(device string, port string) ([]uint16, error) {
	return DeviceCapabilitiesWORD(device, port, co.DC_PAPERS)
}

func DeviceCapabilitiesFixedStrings(device string, port string, dc co.DC, strSize int) ([]string, error) {
	var strings []string
	addr := proc.DeviceCapabilities.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), 0, 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	if ret == 0 {
		return nil, nil
	}
	nStrings := int(ret)
	buf := make([]uint16, nStrings*strSize)
	ret, _, _ = syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(dc), fromBuf(buf), 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	n := 0
	for i := 0; i < nStrings; i++ {
		s := Str.FromNativeSlice(buf[n : n+strSize])
		strings = append(strings, s)
		n += strSize
	}
	return strings, nil
}

func DeviceCapabilitiesBinNames(device string, port string) ([]string, error) {
	return DeviceCapabilitiesFixedStrings(device, port, co.DC_BINS, 24)
}

func DeviceCapabilitiesPaperNames(device string, port string) ([]string, error) {
	return DeviceCapabilitiesFixedStrings(device, port, co.DC_PAPERNAMES, 64)
}

func DeviceCapabilitiesPersonality(device string, port string) ([]string, error) {
	return DeviceCapabilitiesFixedStrings(device, port, co.DC_PERSONALITY, 32)
}

func DeviceCapabilitiesOne(device string, port string, flags co.DM) uintptr {
	addr := proc.DeviceCapabilities.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(flags), 0, 0)
	return ret
}

func DeviceCapabilitiesCollate(device string, port string) bool {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_COLLATE))
	return ret == 1
}

func DeviceCapabilitiesColorDevice(device string, port string) bool {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_COLORDEVICE))
	return ret == 1
}

func DeviceCapabilitiesCopies(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_COPIES))
	return int(ret)
}

func DeviceCapabilitiesDriverVersion(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_DRIVER))
	return int(ret)
}

func DeviceCapabilitiesDuplex(device string, port string) bool {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_DUPLEX))
	return ret == 1
}

// Returns the number of bytes required for the device-specific portion of the DEVMODE structure for the printer driver
func DeviceCapabilitiesExtra(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_EXTRA))
	return int(ret)
}

// Returns the dmFields member of the printer driver's DEVMODE structure.
func DeviceCapabilitiesFields(device string, port string) uint16 {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_FIELDS))
	return uint16(ret)
}

// Returns the dmSize member of the printer driver's DEVMODE structure
func DeviceCapabilitiesSize(device string, port string) uint16 {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_SIZE))
	return uint16(ret)
}

// returns 0, 90, 270
func DeviceCapabilitiesOrientation(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_ORIENTATION))
	return int(ret)
}

func DeviceCapabilitiesVersion(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_VERSION))
	return int(ret)
}

func DeviceCapabilitesPrinterMem(device string, port string) int {
	ret := DeviceCapabilitiesOne(device, port, co.DM(co.DC_PRINTERMEM))
	return int(ret)
}

/*
LONG DocumentProperties(

	_In_  HWND     hWnd,
	_In_  HANDLE   hPrinter,
	_In_  LPTSTR   pDeviceName,
	_Out_ PDEVMODE pDevModeOutput,
	_In_  PDEVMODE pDevModeInput,
	_In_  DWORD    fMode

);
*/
// https://learn.microsoft.com/en-us/windows/win32/printdocs/documentproperties
func DocumentProperties(hwnd HWND, hPrinter HANDLE, deviceName string, devModeIn *DEVMODE, fMode uint32) {
	panic("NYI")
}

type HPRINTER HANDLE

// https://learn.microsoft.com/en-us/windows/win32/printdocs/openprinter
func OpenPrinter(printerName string, printerDefaults *PRINTER_DEFAULTS) (HPRINTER, error) {
	addr := proc.OpenPrinter.Addr()
	var hout HPRINTER
	ret, _, err := syscall.SyscallN(addr, fromUtf8(printerName), fromPtr(&hout), fromPtr(printerDefaults))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return hout, nil
}

// Printer option flags that can be passed to OpenPrinter2 for
// controlling whether the cached or non cached handle is used.
type PRINTER_OPTION uint32

const (
	PRINTER_OPTION_NO_CACHE       PRINTER_OPTION = 1 << 0
	PRINTER_OPTION_CACHE          PRINTER_OPTION = 1 << 1
	PRINTER_OPTION_CLIENT_CHANGE  PRINTER_OPTION = 1 << 2
	PRINTER_OPTION_NO_CLIENT_DATA PRINTER_OPTION = 1 << 3
)

// https://learn.microsoft.com/en-us/windows/win32/printdocs/openprinter2
func OpenPrinter2(printerName string, printerDefaults *PRINTER_DEFAULTS, options PRINTER_OPTION) (HPRINTER, error) {
	addr := proc.OpenPrinter2.Addr()
	var hout HPRINTER
	opts := PRINTER_OPTIONS{
		Flags: options,
	}
	opts.CbSize = uint(unsafe.Sizeof(opts))
	ret, _, err := syscall.SyscallN(addr, fromUtf8(printerName), fromPtr(&hout), fromPtr(printerDefaults), fromPtr(&opts))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return hout, nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/closeprinter
func (hprint HPRINTER) ClosePrinter() error {
	addr := proc.ClosePrinter.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hprint))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/printerproperties
func (hprint HPRINTER) PrinterProperties(hwnd HWND) error {
	addr := proc.PrinterProperties.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hprint), uintptr(hwnd))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

/*
BOOL GetPrinter(
  _In_  HANDLE  hPrinter,
  _In_  DWORD   Level,
  _Out_ LPBYTE  pPrinter,
  _In_  DWORD   cbBuf,
  _Out_ LPDWORD pcbNeeded
);
*/

// https://learn.microsoft.com/en-us/windows/win32/printdocs/getprinter
func GetPrinterRaw(hprint HPRINTER, level int) ([]byte, error) {
	addr := proc.GetPrinter.Addr()
	cbNeeded := uint32(0)
	ret, _, err := syscall.SyscallN(addr, uintptr(level), 0, 0, fromPtr(&cbNeeded))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}
	buf := make([]byte, int(cbNeeded))
	ret, _, err = syscall.SyscallN(addr, uintptr(level), fromBuf(buf), uintptr(len(buf)), fromPtr(&cbNeeded))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}
	return buf, nil
}

func (hprint HPRINTER) GetPrinterLevel2() (*PRINTER_INFO_2, error) {
	d, err := GetPrinterRaw(hprint, 2)
	if err != nil {
		return nil, err
	}
	// TODO: not sure if this is valid
	res := (*PRINTER_INFO_2)(unsafe.Pointer(&d[0]))
	return res, nil
}

// https://learn.microsoft.com/en-us/windows/win32/printdocs/isvaliddevmode
func IsValidDevmode(devMode *DEVMODE, devModeSize uint32) bool {
	addr := proc.IsValidDevmode.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromPtr(devMode), uintptr(devModeSize))
	return ret != 0
}

type PSECURITY_DESCRIPTOR uintptr

type PRINTER_INFO_2 struct {
	ServerName         *uint16
	PrinterName        *uint16
	ShareName          *uint16
	PortName           *uint16
	DriverName         *uint16
	Comment            *uint16
	Location           *uint16
	DevMode            *DEVMODE
	SepFile            *uint16
	PrintProcessor     *uint16
	Datatype           *uint16
	Parameters         *uint16
	SecurityDescriptor PSECURITY_DESCRIPTOR
	Attributes         uint32
	Priority           uint32
	DefaultPriority    uint32
	StartTime          uint32
	UntilTime          uint32
	Status             uint32
	NumJobs            uint32
	AveragePPM         uint32
}
