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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-gettickcount64
func GetTickCount64() uint64 {
	retLo, retHi, _ := syscall.SyscallN(proc.GetTickCount64.Addr())
	return util.Make64(uint32(retLo), uint32(retHi))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
func VerifyVersionInfo(
	ovi *OSVERSIONINFOEX, typeMask co.VER, conditionMask uint64) (bool, error) {

	ovi.SetDwOsVersionInfoSize() // safety
	cMaskLo, cMaskHi := util.Break64(conditionMask)

	ret, _, err := syscall.SyscallN(proc.VerifyVersionInfo.Addr(),
		uintptr(unsafe.Pointer(ovi)),
		uintptr(typeMask), uintptr(cMaskLo), uintptr(cMaskHi))

	if wErr := errco.ERROR(err); ret == 0 && wErr == errco.OLD_WIN_VERSION {
		return false, nil
	} else if ret == 0 {
		return false, wErr // actual error
	} else {
		return true, nil
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/nf-winnt-versetconditionmask
func VerSetConditionMask(
	conditionMask uint64, typeMask co.VER, condition co.VER_COND) uint64 {

	cMaskLo, cMaskHi := util.Break64(conditionMask)

	retLo, retHi, _ := syscall.SyscallN(proc.VerSetConditionMask.Addr(),
		uintptr(cMaskLo), uintptr(cMaskHi),
		uintptr(typeMask), uintptr(condition))
	return util.Make64(uint32(retLo), uint32(retHi))
}
