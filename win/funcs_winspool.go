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
	return uintptr(unsafe.Pointer(v))
}

func fromBuf[T any](v []T) uintptr {
	return uintptr(unsafe.Pointer(&v[0]))
}

func fromUtf8(s string) uintptr {
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

/*
int DeviceCapabilitiesW(
  [in]  LPCWSTR        pDevice,
  [in]  LPCWSTR        pPort,
  [in]  WORD           fwCapability,
  [out] LPWSTR         pOutput,
  [in]  const DEVMODEW *pDevMode
);
*/

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-devicecapabilitiesw

func DeviceCapabilitiesBins(device string, port string) ([]uint16, error) {
	addr := proc.DeviceCapabilities.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(co.DC_BINS), 0, 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	if ret == 0 {
		return nil, nil
	}
	bins := make([]uint16, ret)
	ret, _, _ = syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(co.DC_BINS), fromBuf(bins), 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	return bins, nil
}

func DeviceCapabilitiesBinNames(device string, port string) ([]string, error) {
	var bins []string
	addr := proc.DeviceCapabilities.Addr()
	ret, _, _ := syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(co.DC_BINS), 0, 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	if ret == 0 {
		return nil, nil
	}
	nBins := int(ret)
	binNameSize := 24
	buf := make([]uint16, nBins*binNameSize)
	ret, _, _ = syscall.SyscallN(addr, fromUtf8(device), fromUtf8(port), uintptr(co.DC_BINNAMES), fromBuf(buf), 0)
	if int(ret) < 0 {
		return nil, errco.ERROR(ret)
	}
	n := 0
	for i := 0; i < nBins; i++ {
		s := Str.FromNativeSlice(buf[n : n+binNameSize])
		bins = append(bins, s)
		n += binNameSize
	}
	return bins, nil
}
