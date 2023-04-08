//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/errco"
)

// A handle to an accelerator table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#haccel
type HACCEL HANDLE

// ⚠️ You must defer HACCEL.DestroyAcceleratorTable().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList []ACCEL) HACCEL {
	ret, _, err := syscall.SyscallN(proc.CreateAcceleratorTable.Addr(),
		uintptr(unsafe.Pointer(&accelList[0])), uintptr(len(accelList)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HACCEL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyacceleratortablew
func (hAccel HACCEL) CopyAcceleratorTable() []ACCEL {
	szRet, _, _ := syscall.SyscallN(proc.CopyAcceleratorTable.Addr(),
		uintptr(hAccel), 0, 0)
	if szRet == 0 {
		return []ACCEL{}
	}
	accelList := make([]ACCEL, szRet)
	syscall.SyscallN(proc.CopyAcceleratorTable.Addr(),
		uintptr(hAccel), uintptr(unsafe.Pointer(&accelList[0])), szRet)
	return accelList
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
func (hAccel HACCEL) DestroyAcceleratorTable() error {
	ret, _, err := syscall.SyscallN(proc.DestroyAcceleratorTable.Addr(),
		uintptr(hAccel))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}
