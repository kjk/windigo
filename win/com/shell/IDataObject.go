//go:build windows

package shell

import (
	"github.com/kjk/windigo/win/com/com"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-idataobject
type IDataObject interface {
	com.IUnknown
}

type _IDataObject struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IDataObject.Release().
func NewIDataObject(base com.IUnknown) IDataObject {
	return &_IDataObject{IUnknown: base}
}
