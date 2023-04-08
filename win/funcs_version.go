//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/errco"
)

// Automatically allocs the buffer with GetFileVersionInfoSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfow
func GetFileVersionInfo(fileName string) ([]byte, error) {
	visz, errSz := GetFileVersionInfoSize(fileName)
	if errSz != nil {
		return nil, errSz
	}

	buf := make([]byte, visz) // alloc the buffer

	ret, _, err := syscall.SyscallN(proc.GetFileVersionInfo.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))),
		0, uintptr(visz), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}
	return buf, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizew
func GetFileVersionInfoSize(fileName string) (uint32, error) {
	var lpdwHandle uint32
	ret, _, err := syscall.SyscallN(proc.GetFileVersionInfoSize.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(&lpdwHandle)))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return uint32(ret), nil
}

// Returns a pointer to the block and its size, which varies according to the
// data type. Returns false if the block doesn't exist.
//
// This function is rather tricky. Prefer using ResourceInfo.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluew
func VerQueryValue(
	block []byte, subBlock string) (ptr unsafe.Pointer, sz uint32, exists bool) {

	var lplpBuffer uintptr
	var puLen uint32
	ret, _, _ := syscall.SyscallN(proc.VerQueryValue.Addr(),
		uintptr(unsafe.Pointer(&block[0])),
		uintptr(unsafe.Pointer(Str.ToNativePtr(subBlock))),
		uintptr(unsafe.Pointer(&lplpBuffer)), uintptr(unsafe.Pointer(&puLen)))
	if ret == 0 {
		return nil, 0, false
	}
	return unsafe.Pointer(lplpBuffer), puLen, true
}
