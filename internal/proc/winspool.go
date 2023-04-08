//go:build windows

package proc

import "syscall"

var (
	winspool = syscall.NewLazyDLL("winspool.drv")
	sspool   = syscall.NewLazyDLL("spoolss.dll")

	GetDefaultPrinter  = winspool.NewProc("GetDefaultPrinterW")
	EnumPrinters       = winspool.NewProc("EnumPrintersW")
	DeviceCapabilities = winspool.NewProc("DeviceCapabilitiesW")

	ClosePrinter = sspool.NewProc("ClosePrinter")
)
