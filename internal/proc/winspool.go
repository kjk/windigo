//go:build windows

package proc

import "syscall"

var (
	winspool = syscall.NewLazyDLL("winspool.drv")
	sspool   = syscall.NewLazyDLL("spoolss.dll")

	GetDefaultPrinter  = winspool.NewProc("GetDefaultPrinterW")
	EnumPrinters       = winspool.NewProc("EnumPrintersW")
	DeviceCapabilities = winspool.NewProc("DeviceCapabilitiesW")
	DocumentProperties = winspool.NewProc("DocumentProperties")
	OpenPrinter        = winspool.NewProc("OpenPrinterW")
	OpenPrinter2       = winspool.NewProc("OpenPrinter2W")
	PrinterProperties  = winspool.NewProc("PrinterProperties")
	IsValidDevmode     = winspool.NewProc("IsValidDevmode")
	GetPrinter         = winspool.NewProc("GetPrinterW")

	//GetPrinterData    = winspool.NewProc("GetPrinterDataW")
	//GetPrinterDataE   = winspool.NewProc("GetPrinterDataEx")

	ClosePrinter = sspool.NewProc("ClosePrinter")
)
