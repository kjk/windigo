//go:build windows

package proc

import "syscall"

var (
	winspool = syscall.NewLazyDLL("winspool.drv")
	sspool   = syscall.NewLazyDLL("spoolss.dll")

	GetDefaultPrinter = winspool.NewProc("GetDefaultPrinterW")
	EnumPrinters      = winspool.NewProc("EnumPrintersW")

	ClosePrinter = sspool.NewProc("ClosePrinter")
)
