/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"encoding/binary"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

// IUnknown is the base to all COM interfaces.
type IUnknown struct {
	lpVtbl uintptr
}

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

// Creates any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *IUnknown) coCreateInstance(clsid *co.GUID, iid *co.GUID) {
	if me.lpVtbl != 0 {
		panic("Trying to CoCreateInstance() an IUnknown already created.")
	}

	if iid == nil {
		iid = &co.Guid_IUnknown // if iid is not passed, assume IUnknown
	}

	clsidFlip := cloneFlipLastUint64(clsid)
	iidFlip := cloneFlipLastUint64(iid)
	retIUnk := &IUnknown{}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(&clsidFlip)), 0,
		uintptr(co.CLSCTX_INPROC_SERVER),
		uintptr(unsafe.Pointer(&iidFlip)), uintptr(unsafe.Pointer(&retIUnk)), 0)

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		panic(lerr.Format("CoCreateInstance failed."))
	}
	*me = *retIUnk
}

// Queries any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func (me *IUnknown) queryInterface(iid *co.GUID) *IUnknown {
	lpVtbl := (*iUnknownVtbl)(unsafe.Pointer(me.lpVtbl))
	iidFlip := cloneFlipLastUint64(iid)
	retIUnk := &IUnknown{}

	ret, _, _ := syscall.Syscall(lpVtbl.AddRef, 3,
		uintptr(unsafe.Pointer(me)), uintptr(unsafe.Pointer(&iidFlip)),
		uintptr(unsafe.Pointer(&retIUnk)))

	lerr := co.ERROR(ret)
	if lerr != co.ERROR_S_OK {
		me.Release() // free resource
		panic(lerr.Format("IUnknown.QueryInterface failed."))
	}
	return retIUnk
}

func (me *IUnknown) AddRef() uint32 {
	lpVtbl := (*iUnknownVtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.AddRef, 1,
		uintptr(unsafe.Pointer(me)), 0, 0)
	return uint32(ret)
}

func (me *IUnknown) Release() uint32 {
	lpVtbl := (*iUnknownVtbl)(unsafe.Pointer(me.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.Release, 1,
		uintptr(unsafe.Pointer(me)), 0, 0)
	return uint32(ret)
}

// Returns a new GUID with the last uint64 member bytes flipped.
// This is better than having a makeGuid() function being called to initialize
// all the GUIDs with correct by order, or even having to flip them manually on
// each global declaration.
func cloneFlipLastUint64(guid *co.GUID) co.GUID {
	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], guid.Data4)
	guidCopy := *guid
	guidCopy.Data4 = binary.LittleEndian.Uint64(buf64[:])
	return guidCopy
}