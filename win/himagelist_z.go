//go:build windows

package win

import (
	"github.com/kjk/windigo/internal/util"
	"github.com/kjk/windigo/win/co"
)

// This helper method loads icons from the shell, used by Windows Explorer to
// represent the given file extensions, like "mp3".
func (hImg HIMAGELIST) AddIconFromShell(fileExtensions ...string) {
	sz := hImg.GetIconSize()
	isIco16 := sz.Cx == 16 && sz.Cy == 16
	isIco32 := sz.Cx == 32 && sz.Cy == 32
	if !isIco16 && !isIco32 {
		panic("AddIconFromShell can load only 16x16 or 32x32 icons.")
	}

	shgfi := co.SHGFI_USEFILEATTRIBUTES | co.SHGFI_ICON |
		util.Iif(isIco32, co.SHGFI_LARGEICON, co.SHGFI_SMALLICON).(co.SHGFI)

	fi := SHFILEINFO{}
	for _, fileExt := range fileExtensions {
		SHGetFileInfo("*."+fileExt, co.FILE_ATTRIBUTE_NORMAL, &fi, shgfi)
		hImg.AddIcon(fi.HIcon)
		fi.HIcon.DestroyIcon()
	}
}
