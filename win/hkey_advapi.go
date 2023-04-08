//go:build windows

package win

import (
	"sort"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a registry key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hkey
type HKEY HANDLE

// Predefined registry key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/predefined-keys
const (
	HKEY_CLASSES_ROOT        HKEY = 0x8000_0000
	HKEY_CURRENT_USER        HKEY = 0x8000_0001
	HKEY_LOCAL_MACHINE       HKEY = 0x8000_0002
	HKEY_USERS               HKEY = 0x8000_0003
	HKEY_PERFORMANCE_DATA    HKEY = 0x8000_0004
	HKEY_PERFORMANCE_TEXT    HKEY = 0x8000_0050
	HKEY_PERFORMANCE_NLSTEXT HKEY = 0x8000_0060
	HKEY_CURRENT_CONFIG      HKEY = 0x8000_0005
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regclosekey
func (hKey HKEY) RegCloseKey() error {
	ret, _, _ := syscall.SyscallN(proc.RegCloseKey.Addr(),
		uintptr(hKey))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyw
func (hKey HKEY) RegDeleteKey(subKey string) error {
	ret, _, _ := syscall.SyscallN(proc.RegDeleteKey.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// samDesired must be KEY_WOW64_32KEY or KEY_WOW64_64KEY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyexw
func (hKey HKEY) RegDeleteKeyEx(subKey string, samDesired co.KEY) error {
	ret, _, _ := syscall.SyscallN(proc.RegDeleteKeyEx.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(samDesired), 0)
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletekeyvaluew
func (hKey HKEY) RegDeleteKeyValue(subKey, valueName string) error {
	ret, _, _ := syscall.SyscallN(proc.RegDeleteKeyValue.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(valueName))))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regdeletetreew
func (hKey HKEY) RegDeleteTree(subKey string) error {
	ret, _, _ := syscall.SyscallN(proc.RegDeleteTree.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))))
	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// Returns the names of all subkeys within a key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumkeyexw
func (hKey HKEY) RegEnumKeyEx() ([]string, error) {
	keyInfo, err := hKey.RegQueryInfoKey()
	if err != nil {
		return nil, err
	}

	keyNames := make([]string, 0, keyInfo.NumSubKeys)        // key names to be returned
	keyNameBuf := make([]uint16, keyInfo.MaxSubKeyNameLen+1) // to receive the names of the keys
	var keyNameBufLen uint32

	for i := 0; i < int(keyInfo.NumSubKeys); i++ {
		keyNameBufLen = uint32(len(keyNameBuf)) // reset available buffer size

		ret, _, _ := syscall.SyscallN(proc.RegEnumKeyEx.Addr(),
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&keyNameBuf[0])),
			uintptr(unsafe.Pointer(&keyNameBufLen)), // receives the number of chars without null
			0, 0, 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			return nil, wErr
		}

		keyNames = append(keyNames, Str.FromNativeSlice(keyNameBuf))
	}

	Path.Sort(keyNames)
	return keyNames, nil
}

type _ValueEnum struct {
	Name string
	Type co.REG
}

// Returns the names and types of all values within a key.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regenumvaluew
func (hKey HKEY) RegEnumValue() ([]_ValueEnum, error) {
	keyInfo, err := hKey.RegQueryInfoKey()
	if err != nil {
		return nil, err
	}

	values := make([]_ValueEnum, 0, keyInfo.NumValues) // to be returned

	valueNameBuf := make([]uint16, keyInfo.MaxValueNameLen+2) // room to avoid "more data" error
	var valueNameBufLen uint32
	var valueTypeBuf co.REG

	for i := 0; i < int(keyInfo.NumValues); i++ {
		valueNameBufLen = uint32(len(valueNameBuf)) // reset available buffer size

		ret, _, _ := syscall.SyscallN(proc.RegEnumValue.Addr(),
			uintptr(hKey), uintptr(i),
			uintptr(unsafe.Pointer(&valueNameBuf[0])),
			uintptr(unsafe.Pointer(&valueNameBufLen)), // receives the number of chars without null
			0, uintptr(unsafe.Pointer(&valueTypeBuf)), 0, 0)

		if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
			return nil, wErr
		}

		values = append(values, struct {
			Name string
			Type co.REG
		}{
			Name: Str.FromNativeSlice(valueNameBuf),
			Type: valueTypeBuf,
		})
	}

	sort.Slice(values, func(a, b int) bool {
		return strings.ToUpper(values[a].Name) < strings.ToUpper(values[b].Name) // case insensitive
	})
	return values, nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regflushkey
func (hKey HKEY) RegFlushKey() error {
	ret, _, _ := syscall.SyscallN(proc.RegFlushKey.Addr(),
		uintptr(hKey))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// This function is rather tricky. Prefer using HKEY.ReadBinary(),
// HKEY.ReadDword(), HKEY.ReadQword() or HKEY.ReadString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-reggetvaluew
func (hKey HKEY) RegGetValue(
	subKey, value string, flags co.RRF, pdwType *co.REG,
	pData unsafe.Pointer, pDataLen *uint32) error {

	ret, _, _ := syscall.SyscallN(proc.RegGetValue.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(value))),
		uintptr(flags), uintptr(unsafe.Pointer(pdwType)),
		uintptr(pData), uintptr(unsafe.Pointer(pDataLen)))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}

// ⚠️ You must defer HKEY.RegCloseKey().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regopenkeyexw
func (hKey HKEY) RegOpenKeyEx(
	subKey string, ulOptions co.REG_OPTION, samDesired co.KEY) (HKEY, error) {

	var openedKey HKEY
	ret, _, _ := syscall.SyscallN(proc.RegOpenKeyEx.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(ulOptions), uintptr(samDesired),
		uintptr(unsafe.Pointer(&openedKey)))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return HKEY(0), wErr
	}
	return openedKey, nil
}

type _KeyInfo struct {
	Class                 string
	NumSubKeys            uint32
	MaxSubKeyNameLen      uint32
	MaxSubKeyClassLen     uint32
	NumValues             uint32
	MaxValueNameLen       uint32
	MaxValueDataLen       uint32
	SecurityDescriptorLen uint32
	LastWriteTime         time.Time
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regqueryinfokeyw
func (hKey HKEY) RegQueryInfoKey() (_KeyInfo, error) {
	var info _KeyInfo
	var classBuf [_MAX_PATH + 1]uint16 // arbitrary
	classBufLen := uint32(len(classBuf))
	var ft FILETIME

	ret, _, _ := syscall.SyscallN(proc.RegQueryInfoKey.Addr(),
		uintptr(hKey),
		uintptr(unsafe.Pointer(&classBuf[0])), uintptr(unsafe.Pointer(&classBufLen)), 0,
		uintptr(unsafe.Pointer(&info.NumSubKeys)),
		uintptr(unsafe.Pointer(&info.MaxSubKeyNameLen)),
		uintptr(unsafe.Pointer(&info.MaxSubKeyClassLen)),
		uintptr(unsafe.Pointer(&info.NumValues)),
		uintptr(unsafe.Pointer(&info.MaxValueNameLen)),
		uintptr(unsafe.Pointer(&info.MaxValueDataLen)),
		uintptr(unsafe.Pointer(&info.SecurityDescriptorLen)),
		uintptr(unsafe.Pointer(&ft)))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return _KeyInfo{}, wErr
	}

	info.Class = Str.FromNativeSlice(classBuf[:])
	info.LastWriteTime = ft.ToTime()
	return info, nil
}

// This function is rather tricky. Prefer using HKEY.WriteBinary(),
// HKEY.WriteDword(), HKEY.WriteQword() or HKEY.WriteString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winreg/nf-winreg-regsetkeyvaluew
func (hKey HKEY) RegSetKeyValue(
	subKey, valueName string, dwType co.REG,
	pData unsafe.Pointer, dataLen uint32) error {

	ret, _, _ := syscall.SyscallN(proc.RegSetKeyValue.Addr(),
		uintptr(hKey), uintptr(unsafe.Pointer(Str.ToNativePtr(subKey))),
		uintptr(unsafe.Pointer(Str.ToNativePtr(valueName))),
		uintptr(dwType), uintptr(pData), uintptr(dataLen))

	if wErr := errco.ERROR(ret); wErr != errco.SUCCESS {
		return wErr
	}
	return nil
}
