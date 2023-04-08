//go:build windows

package win

import (
	"syscall"

	"github.com/kjk/windigo/internal/proc"
)

// Handle to an OLE block of memory.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
type HTASKMEM HANDLE

// ⚠️ You must defer HTASKMEM.CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemalloc
func CoTaskMemAlloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.SyscallN(proc.CoTaskMemAlloc.Addr(),
		uintptr(numBytes))
	if ret == 0 {
		panic("CoTaskMemAlloc() failed.")
	}
	return HTASKMEM(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (hMem HTASKMEM) CoTaskMemFree() {
	syscall.SyscallN(proc.CoTaskMemFree.Addr(),
		uintptr(hMem))
}

// ⚠️ You must defer CoTaskMemFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemrealloc
func (hMem HTASKMEM) CoTaskMemRealloc(numBytes int) HTASKMEM {
	ret, _, _ := syscall.SyscallN(proc.CoTaskMemRealloc.Addr(),
		uintptr(hMem), uintptr(numBytes))
	if ret == 0 {
		panic("CoTaskMemRealloc() failed.")
	}
	return HTASKMEM(ret)
}
