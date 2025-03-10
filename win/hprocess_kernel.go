//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// Handle to a process.
type HPROCESS HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocesses
func EnumProcesses() ([]uint32, error) {
	const UINT32_SZ = unsafe.Sizeof(uint32(0)) // in bytes

	const BLOCK int = 256 // arbitrary
	bufSz := BLOCK

	var processIds []uint32
	var bytesNeeded, numReturned uint32

	for {
		processIds = make([]uint32, bufSz)

		ret, _, err := syscall.SyscallN(proc.EnumProcesses.Addr(),
			uintptr(unsafe.Pointer(&processIds[0])),
			uintptr(len(processIds))*UINT32_SZ, // array size in bytes
			uintptr(unsafe.Pointer(&bytesNeeded)))

		if ret == 0 {
			return nil, errco.ERROR(err)
		}

		numReturned = bytesNeeded / uint32(UINT32_SZ)
		if numReturned < uint32(len(processIds)) { // to break, must have at least 1 element gap
			break
		}

		bufSz += BLOCK // increase buffer size to try again
	}

	return processIds[:numReturned], nil
}

// ⚠️ You must defer HPROCESS.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocess
func GetCurrentProcess() HPROCESS {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentProcess.Addr())
	return HPROCESS(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid
func GetCurrentProcessId() uint32 {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentProcessId.Addr())
	return uint32(ret)
}

// ⚠️ You must defer HPROCESS.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
func OpenProcess(
	desiredAccess co.PROCESS,
	inheritHandle bool, processId uint32) (HPROCESS, error) {

	ret, _, err := syscall.SyscallN(proc.OpenProcess.Addr(),
		uintptr(desiredAccess), util.BoolToUintptr(inheritHandle),
		uintptr(processId))
	if ret == 0 {
		return HPROCESS(0), errco.ERROR(err)
	}
	return HPROCESS(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcess HPROCESS) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hProcess))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocessmodules
func (hProcess HPROCESS) EnumProcessModules() ([]HINSTANCE, error) {
	const HINSTANCE_SZ = unsafe.Sizeof(HINSTANCE(0)) // in bytes

	var bytesNeeded uint32
	ret, _, err := syscall.SyscallN(proc.EnumProcessModules.Addr(),
		uintptr(hProcess), 0, 0, uintptr(unsafe.Pointer(&bytesNeeded)))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}

	hModules := make([]HINSTANCE, uintptr(bytesNeeded)/HINSTANCE_SZ)
	ret, _, err = syscall.SyscallN(proc.EnumProcessModules.Addr(),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&hModules[0])),
		uintptr(len(hModules))*HINSTANCE_SZ, // array size in bytes
		uintptr(unsafe.Pointer(&bytesNeeded)))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}

	return hModules, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
func (hProcess HPROCESS) GetExitCodeProcess() (uint32, error) {
	var exitCode uint32
	ret, _, err := syscall.SyscallN(proc.GetExitCodeProcess.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return exitCode, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	var processName [_MAX_PATH + 1]uint16
	ret, _, err := syscall.SyscallN(proc.GetModuleBaseName.Addr(),
		uintptr(hProcess), uintptr(hModule),
		uintptr(unsafe.Pointer(&processName[0])),
		uintptr(len(processName)))
	if ret == 0 {
		return "", errco.ERROR(err)
	}
	return Str.FromNativeSlice(processName[:]), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() (uint32, error) {
	ret, _, err := syscall.SyscallN(proc.GetProcessId.Addr(),
		uintptr(hProcess))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes() (
	creationTime, exitTime, kernelTime, userTime FILETIME, e error) {

	ret, _, err := syscall.SyscallN(proc.GetProcessTimes.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&creationTime)),
		uintptr(unsafe.Pointer(&exitTime)), uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)))
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
func (hProcess HPROCESS) ReadProcessMemory(
	baseAddress uintptr, buffer []byte) (numBytesRead uint64, e error) {

	ret, _, err := syscall.SyscallN(proc.ReadProcessMemory.Addr(),
		uintptr(hProcess), baseAddress, uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)), uintptr(unsafe.Pointer(&numBytesRead)))
	if ret == 0 {
		numBytesRead, e = 0, errco.ERROR(err)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setuserobjectinformationw
func (hProcess HPROCESS) SetUserObjectInformation(
	index co.UOI, info unsafe.Pointer, infoLen uintptr) error {

	ret, _, err := syscall.SyscallN(proc.SetUserObjectInformation.Addr(),
		uintptr(hProcess), uintptr(index), uintptr(info), uintptr(infoLen))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func (hProcess HPROCESS) TerminateProcess(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(proc.TerminateProcess.Addr(),
		uintptr(hProcess), uintptr(exitCode))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hProcess HPROCESS) WaitForSingleObject(milliseconds uint32) (co.WAIT, error) {
	ret, _, err := syscall.SyscallN(proc.WaitForSingleObject.Addr(),
		uintptr(hProcess), uintptr(milliseconds))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, errco.ERROR(err)
	}
	return co.WAIT(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory
func (hProcess HPROCESS) WriteProcessMemory(
	baseAddress uintptr, data []byte) (numBytesWritten uint64, e error) {

	ret, _, err := syscall.SyscallN(proc.WriteProcessMemory.Addr(),
		uintptr(hProcess), baseAddress, uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)), uintptr(unsafe.Pointer(&numBytesWritten)))
	if ret == 0 {
		numBytesWritten, e = 0, errco.ERROR(err)
	}
	return
}
