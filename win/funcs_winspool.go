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
	Attributes               uint32
	DeviceNotSelectedTimeout uint32
	TransmissionRetryTimeout uint32
}

func fromUint32Ptr(v *uint32) uintptr {
	return uintptr(unsafe.Pointer(v))
}

func fromBuf(v []byte) uintptr {
	return uintptr(unsafe.Pointer(&v[0]))
}

type PrinterInfo struct {
	PrinterName              string
	PortName                 string
	Attributes               uint32
	DeviceNotSelectedTimeout uint32
	TransmissionRetryTimeout uint32
}

func EnumPrinters(dwFlags co.PRINTER_ENUM) ([]PrinterInfo, error) {
	var cbNeeded uint32
	var nPrinters uint32
	addr := proc.EnumPrinters.Addr()
	syscall.SyscallN(addr, uintptr(dwFlags), 0, 5, 0, 0, fromUint32Ptr(&cbNeeded), fromUint32Ptr(&nPrinters))
	buf := make([]byte, cbNeeded)
	cbBuf := cbNeeded
	res, _, err := syscall.SyscallN(addr, uintptr(dwFlags), 0, 5, fromBuf(buf), uintptr(cbBuf), fromUint32Ptr(&cbNeeded), fromUint32Ptr(&nPrinters))
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
	ret, _, _ := syscall.SyscallN(proc.GetDefaultPrinter.Addr(),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&lenInOut)))
	// TODO: handle ret == ERROR_INSUFFICIENT_BUFFER and ERROR_FILE_NOT_FOUND
	// or return an error
	if ret == 0 {
		return "", errco.ERROR(ret)
	}
	return Str.FromNativeSlice(buf[:]), nil

}
