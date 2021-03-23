package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to HFILEMAP, providing several operations.
//
// Created with OpenFileMapped().
type FileMapped interface {
	// Unmaps and releases the file resource.
	Close()

	// Returns a slice to the memory-mapped bytes. The FileMapped object must
	// remain open while the slice is being used.
	//
	// If you need to close the file right away, use CopyToBuffer() instead.
	HotSlice() []byte

	// Returns a new []byte with a copy of all data in the file.
	ReadAll() []byte

	// Returns a new []byte with a copy of data, start with offset, and with the
	// given length.
	ReadChunk(offset, length int) []byte

	// Truncates or expands the file, according to the new size. Zero will empty the
	// file.
	//
	// Internally, the file is unmapped, then remapped back into memory.
	Resize(numBytes int) error

	// Retrieves the file size. This value is cached.
	Size() int
}

//------------------------------------------------------------------------------

type _FileMapped struct {
	objFile  _File
	hMap     HFILEMAP
	pMem     HFILEMAPMEM
	sz       int
	readOnly bool // necessary for SetSize()
}

// Opens a memory-mapped file, returning a new high-level FileMapped object.
//
// ⚠️ You must defer Close().
func OpenFileMapped(
	filePath string, behavior co.OPEN_FILEMAP) (FileMapped, error) {

	var mapOpts co.OPEN_FILE
	var readOnly bool

	switch behavior {
	case co.OPEN_FILEMAP_MODE_READ:
		mapOpts = co.OPEN_FILE_READ_EXISTING
		readOnly = true
	case co.OPEN_FILEMAP_MODE_RW:
		mapOpts = co.OPEN_FILE_RW_EXISTING
		readOnly = false
	}

	me := _FileMapped{
		objFile:  _File{},
		hMap:     HFILEMAP(0),
		pMem:     HFILEMAPMEM(0),
		sz:       0,
		readOnly: readOnly,
	}

	if lerr := me.objFile.openFile(filePath, mapOpts); lerr != nil {
		return nil, lerr
	}
	if lerr := me.mapInMemory(); lerr != nil {
		return nil, lerr
	}
	return &me, nil
}

func (me *_FileMapped) Close() {
	if me.pMem != 0 {
		me.pMem.UnmapViewOfFile()
		me.pMem = 0
	}
	if me.hMap != 0 {
		me.hMap.CloseHandle()
		me.hMap = 0
	}
	me.objFile.Close()
	me.sz = 0
}

func (me *_FileMapped) HotSlice() []byte {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(me.pMem), int(me.sz), int(me.sz)}

	return *(*[]byte)(unsafe.Pointer(&sliceMem))
}

func (me *_FileMapped) ReadAll() []byte {
	return me.ReadChunk(0, me.sz)
}

func (me *_FileMapped) ReadChunk(offset, length int) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

func (me *_FileMapped) Resize(numBytes int) error {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if lerr := me.objFile.Resize(numBytes); lerr != nil {
		return lerr
	}
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *_FileMapped) Size() int {
	return me.sz
}

func (me *_FileMapped) mapInMemory() error {
	// Mapping into memory.
	pageFlags := co.PAGE_READWRITE
	if me.readOnly {
		pageFlags = co.PAGE_READONLY
	}

	var lerr error
	if me.hMap, lerr = me.objFile.hFile.CreateFileMapping(
		nil, pageFlags, co.SEC_NONE, 0, ""); lerr != nil {
		return lerr
	}

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_WRITE
	if me.readOnly {
		mapFlags = co.FILE_MAP_READ
	}

	if me.pMem, lerr = me.hMap.MapViewOfFile(mapFlags, 0, 0); lerr != nil {
		return lerr
	}

	// Cache file size.
	me.sz = me.objFile.Size()

	return nil // file mapped successfully
}