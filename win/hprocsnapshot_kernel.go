//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// Handle to a process snapshot.
type HPROCSNAPSHOT HANDLE

// ⚠️ You must defer HPROCSNAPSHOT.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
func CreateToolhelp32Snapshot(
	flags co.TH32CS, processId uint32) (HPROCSNAPSHOT, error) {

	ret, _, err := syscall.SyscallN(proc.CreateToolhelp32Snapshot.Addr(),
		uintptr(flags), uintptr(processId))
	if int(ret) == _INVALID_HANDLE_VALUE {
		return HPROCSNAPSHOT(0), errco.ERROR(err)
	}
	return HPROCSNAPSHOT(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcSnap HPROCSNAPSHOT) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hProcSnap))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumModules().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32firstw
func (hProcSnap HPROCSNAPSHOT) Module32First(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(proc.Module32First.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumModules().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32nextw
func (hProcSnap HPROCSNAPSHOT) Module32Next(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(proc.Module32Next.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumProcesses().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAPSHOT) Process32First(
	buf *PROCESSENTRY32) (bool, error) {

	ret, _, err := syscall.SyscallN(proc.Process32First.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumProcesses().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAPSHOT) Process32Next(
	buf *PROCESSENTRY32) (bool, error) {

	ret, _, err := syscall.SyscallN(proc.Process32Next.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumThreads().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32first
func (hProcSnap HPROCSNAPSHOT) Thread32First(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(proc.Thread32First.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

// This function is rather tricky. Prefer using HPROCSNAPSHOT.EnumThreads().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32next
func (hProcSnap HPROCSNAPSHOT) Thread32Next(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(proc.Thread32Next.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}
