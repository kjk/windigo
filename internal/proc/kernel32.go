//go:build windows

package proc

import (
	"syscall"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	AllocConsole                    = kernel32.NewProc("AllocConsole")
	AttachConsole                   = kernel32.NewProc("AttachConsole")
	CloseHandle                     = kernel32.NewProc("CloseHandle")
	ConnectNamedPipe                = kernel32.NewProc("ConnectNamedPipe")
	CopyFile                        = kernel32.NewProc("CopyFileW")
	CreateDirectory                 = kernel32.NewProc("CreateDirectoryW")
	CreateFile                      = kernel32.NewProc("CreateFileW")
	CreateFileMappingFromApp        = kernel32.NewProc("CreateFileMappingFromApp")
	CreateNamedPipe                 = kernel32.NewProc("CreateNamedPipeW")
	CreateProcess                   = kernel32.NewProc("CreateProcessW")
	CreateToolhelp32Snapshot        = kernel32.NewProc("CreateToolhelp32Snapshot")
	DeleteFile                      = kernel32.NewProc("DeleteFileW")
	DisconnectNamedPipe             = kernel32.NewProc("DisconnectNamedPipe")
	ExitProcess                     = kernel32.NewProc("ExitProcess")
	ExpandEnvironmentStrings        = kernel32.NewProc("ExpandEnvironmentStringsW")
	FileTimeToSystemTime            = kernel32.NewProc("FileTimeToSystemTime")
	FindClose                       = kernel32.NewProc("FindClose")
	FindFirstFile                   = kernel32.NewProc("FindFirstFileW")
	FindNextFile                    = kernel32.NewProc("FindNextFileW")
	FindResource                    = kernel32.NewProc("FindResourceW")
	FindResourceEx                  = kernel32.NewProc("FindResourceExW")
	FlushViewOfFile                 = kernel32.NewProc("FlushViewOfFile")
	FreeConsole                     = kernel32.NewProc("FreeConsole")
	FreeEnvironmentStrings          = kernel32.NewProc("FreeEnvironmentStringsW")
	FreeLibrary                     = kernel32.NewProc("FreeLibrary")
	GetCommandLine                  = kernel32.NewProc("GetCommandLineW")
	GetConsoleCP                    = kernel32.NewProc("GetConsoleCP")
	GetConsoleTitle                 = kernel32.NewProc("GetConsoleTitleW")
	GetConsoleWindow                = kernel32.NewProc("GetConsoleWindow")
	GetCurrentConsoleFont           = kernel32.NewProc("GetCurrentConsoleFont")
	GetCurrentDirectory             = kernel32.NewProc("GetCurrentDirectoryW")
	GetCurrentProcess               = kernel32.NewProc("GetCurrentProcess")
	GetCurrentProcessId             = kernel32.NewProc("GetCurrentProcessId")
	GetCurrentThread                = kernel32.NewProc("GetCurrentThread")
	GetCurrentThreadId              = kernel32.NewProc("GetCurrentThreadId")
	GetDynamicTimeZoneInformation   = kernel32.NewProc("GetDynamicTimeZoneInformation")
	GetEnvironmentStrings           = kernel32.NewProc("GetEnvironmentStringsW")
	GetExitCodeProcess              = kernel32.NewProc("GetExitCodeProcess")
	GetExitCodeThread               = kernel32.NewProc("GetExitCodeThread")
	GetFileAttributes               = kernel32.NewProc("GetFileAttributesW")
	GetFileSizeEx                   = kernel32.NewProc("GetFileSizeEx")
	GetModuleFileName               = kernel32.NewProc("GetModuleFileNameW")
	GetModuleHandle                 = kernel32.NewProc("GetModuleHandleW")
	GetProcAddress                  = kernel32.NewProc("GetProcAddress")
	GetProcessId                    = kernel32.NewProc("GetProcessId")
	GetProcessIdOfThread            = kernel32.NewProc("GetProcessIdOfThread")
	GetProcessTimes                 = kernel32.NewProc("GetProcessTimes")
	GetStartupInfo                  = kernel32.NewProc("GetStartupInfoW")
	GetStdHandle                    = kernel32.NewProc("GetStdHandle")
	GetSystemInfo                   = kernel32.NewProc("GetSystemInfo")
	GetSystemTime                   = kernel32.NewProc("GetSystemTime")
	GetSystemTimeAsFileTime         = kernel32.NewProc("GetSystemTimeAsFileTime")
	GetSystemTimePreciseAsFileTime  = kernel32.NewProc("GetSystemTimePreciseAsFileTime")
	GetSystemTimes                  = kernel32.NewProc("GetSystemTimes")
	GetThreadId                     = kernel32.NewProc("GetThreadId")
	GetThreadTimes                  = kernel32.NewProc("GetThreadTimes")
	GetTickCount64                  = kernel32.NewProc("GetTickCount64")
	GetTimeZoneInformation          = kernel32.NewProc("GetTimeZoneInformation")
	GetTimeZoneInformationForYear   = kernel32.NewProc("GetTimeZoneInformationForYear")
	GetVolumeInformation            = kernel32.NewProc("GetVolumeInformationW")
	GetWindowsDirectory             = kernel32.NewProc("GetWindowsDirectoryW")
	GlobalAddAtom                   = kernel32.NewProc("GlobalAddAtomW")
	GlobalAlloc                     = kernel32.NewProc("GlobalAlloc")
	GlobalDeleteAtom                = kernel32.NewProc("GlobalDeleteAtom")
	GlobalFlags                     = kernel32.NewProc("GlobalFlags")
	GlobalFree                      = kernel32.NewProc("GlobalFree")
	GlobalGetAtomName               = kernel32.NewProc("GlobalGetAtomNameW")
	GlobalLock                      = kernel32.NewProc("GlobalLock")
	GlobalReAlloc                   = kernel32.NewProc("GlobalReAlloc")
	GlobalSize                      = kernel32.NewProc("GlobalSize")
	GlobalUnlock                    = kernel32.NewProc("GlobalUnlock")
	LoadLibrary                     = kernel32.NewProc("LoadLibraryW")
	LoadResource                    = kernel32.NewProc("LoadResource")
	LockFile                        = kernel32.NewProc("LockFile")
	LockFileEx                      = kernel32.NewProc("LockFileEx")
	LockResource                    = kernel32.NewProc("LockResource")
	MapViewOfFileFromApp            = kernel32.NewProc("MapViewOfFileFromApp")
	Module32First                   = kernel32.NewProc("Module32FirstW")
	Module32Next                    = kernel32.NewProc("Module32NextW")
	MoveFile                        = kernel32.NewProc("MoveFileW")
	MoveFileEx                      = kernel32.NewProc("MoveFileExW")
	MulDiv                          = kernel32.NewProc("MulDiv")
	OpenProcess                     = kernel32.NewProc("OpenProcess")
	Process32First                  = kernel32.NewProc("Process32FirstW")
	Process32Next                   = kernel32.NewProc("Process32NextW")
	QueryPerformanceCounter         = kernel32.NewProc("QueryPerformanceCounter")
	QueryPerformanceFrequency       = kernel32.NewProc("QueryPerformanceFrequency")
	ReadConsole                     = kernel32.NewProc("ReadConsoleW")
	ReadFile                        = kernel32.NewProc("ReadFile")
	ReadProcessMemory               = kernel32.NewProc("ReadProcessMemory")
	RemoveDirectory                 = kernel32.NewProc("RemoveDirectoryW")
	ReplaceFile                     = kernel32.NewProc("ReplaceFileW")
	ResumeThread                    = kernel32.NewProc("ResumeThread")
	SetConsoleCursorInfo            = kernel32.NewProc("SetConsoleCursorInfo")
	SetConsoleCursorPosition        = kernel32.NewProc("SetConsoleCursorPosition")
	SetConsoleDisplayMode           = kernel32.NewProc("SetConsoleDisplayMode")
	SetConsoleMode                  = kernel32.NewProc("SetConsoleMode")
	SetConsoleOutputCP              = kernel32.NewProc("SetConsoleOutputCP")
	SetConsoleScreenBufferSize      = kernel32.NewProc("SetConsoleScreenBufferSize")
	SetConsoleTitle                 = kernel32.NewProc("SetConsoleTitleW")
	SetCurrentDirectory             = kernel32.NewProc("SetCurrentDirectoryW")
	SetEndOfFile                    = kernel32.NewProc("SetEndOfFile")
	SetFileAttributes               = kernel32.NewProc("SetFileAttributesW")
	SetFilePointerEx                = kernel32.NewProc("SetFilePointerEx")
	SetLastError                    = kernel32.NewProc("SetLastError")
	SizeofResource                  = kernel32.NewProc("SizeofResource")
	Sleep                           = kernel32.NewProc("Sleep")
	SuspendThread                   = kernel32.NewProc("SuspendThread")
	SystemTimeToFileTime            = kernel32.NewProc("SystemTimeToFileTime")
	SystemTimeToTzSpecificLocalTime = kernel32.NewProc("SystemTimeToTzSpecificLocalTime")
	TerminateProcess                = kernel32.NewProc("TerminateProcess")
	TerminateThread                 = kernel32.NewProc("TerminateThread")
	Thread32First                   = kernel32.NewProc("Thread32First")
	Thread32Next                    = kernel32.NewProc("Thread32Next")
	TzSpecificLocalTimeToSystemTime = kernel32.NewProc("TzSpecificLocalTimeToSystemTime")
	UnlockFile                      = kernel32.NewProc("UnlockFile")
	UnlockFileEx                    = kernel32.NewProc("UnlockFileEx")
	UnmapViewOfFile                 = kernel32.NewProc("UnmapViewOfFile")
	VerifyVersionInfo               = kernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask             = kernel32.NewProc("VerSetConditionMask")
	WaitForSingleObject             = kernel32.NewProc("WaitForSingleObject")
	WriteConsole                    = kernel32.NewProc("WriteConsoleW")
	WriteFile                       = kernel32.NewProc("WriteFile")
	WriteProcessMemory              = kernel32.NewProc("WriteProcessMemory")
)
