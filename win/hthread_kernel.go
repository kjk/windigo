//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// Handle to a thread.
type HTHREAD HANDLE

// ⚠️ You must defer HTHREAD.CloseHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthread
func GetCurrentThread() HTHREAD {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentThread.Addr())
	return HTHREAD(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.SyscallN(proc.GetCurrentThreadId.Addr())
	return uint32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hThread HTHREAD) CloseHandle() error {
	ret, _, err := syscall.SyscallN(proc.CloseHandle.Addr(),
		uintptr(hThread))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodethread
func (hThread HTHREAD) GetExitCodeThread() (uint32, error) {
	var exitCode uint32
	ret, _, err := syscall.SyscallN(proc.GetExitCodeThread.Addr(),
		uintptr(hThread), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return exitCode, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessidofthread
func (hThread HTHREAD) GetProcessIdOfThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(proc.GetProcessIdOfThread.Addr(),
		uintptr(hThread))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadid
func (hThread HTHREAD) GetThreadId() (uint32, error) {
	ret, _, err := syscall.SyscallN(proc.GetThreadId.Addr(),
		uintptr(hThread))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getthreadtimes
func (hThread HTHREAD) GetThreadTimes() (
	creationTime, exitTime, kernelTime, userTime FILETIME, e error) {

	ret, _, err := syscall.SyscallN(proc.GetThreadTimes.Addr(),
		uintptr(hThread), uintptr(unsafe.Pointer(&creationTime)),
		uintptr(unsafe.Pointer(&exitTime)), uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)))
	if ret == 0 {
		e = errco.ERROR(err)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-resumethread
func (hThread HTHREAD) ResumeThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(proc.ResumeThread.Addr(),
		uintptr(hThread))
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminatethread
func (hThread HTHREAD) TerminateThread(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(proc.TerminateThread.Addr(),
		uintptr(hThread), uintptr(exitCode))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-suspendthread
func (hThread HTHREAD) SuspendThread() (uint32, error) {
	ret, _, err := syscall.SyscallN(proc.SuspendThread.Addr(),
		uintptr(hThread))
	if int(ret) == -1 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
func (hThread HTHREAD) WaitForSingleObject(milliseconds uint32) (co.WAIT, error) {
	ret, _, err := syscall.SyscallN(proc.WaitForSingleObject.Addr(),
		uintptr(hThread), uintptr(milliseconds))
	if co.WAIT(ret) == co.WAIT_FAILED {
		return co.WAIT_FAILED, errco.ERROR(err)
	}
	return co.WAIT(ret), nil
}
