//go:build windows

package win

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/kjk/windigo/internal/proc"
	"github.com/kjk/windigo/win/co"
	"github.com/kjk/windigo/win/errco"
)

// A handle to a device context (DC).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hdc
type HDC HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortpath
func (hdc HDC) AbortPath() {
	ret, _, err := syscall.SyscallN(proc.AbortPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-anglearc
func (hdc HDC) AngleArc(center POINT, r uint32, startAngle, sweepAngle float32) {
	ret, _, err := syscall.SyscallN(proc.AngleArc.Addr(),
		uintptr(hdc), uintptr(center.X), uintptr(center.Y), uintptr(r),
		uintptr(startAngle), uintptr(sweepAngle))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arc
func (hdc HDC) Arc(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.Arc.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-arcto
func (hdc HDC) ArcTo(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.ArcTo.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HDC.EndPath().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-beginpath
func (hdc HDC) BeginPath() {
	ret, _, err := syscall.SyscallN(proc.BeginPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-bitblt
func (hdc HDC) BitBlt(
	destTopLeft POINT, sz SIZE, hdcSrc HDC, srcTopLeft POINT, rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.BitBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(sz.Cx), uintptr(sz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-canceldc
func (hdc HDC) CancelDC() {
	ret, _, err := syscall.SyscallN(proc.CancelDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-chord
func (hdc HDC) Chord(bound RECT, radialStart, radialEnd POINT) {
	ret, _, err := syscall.SyscallN(proc.Chord.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(radialStart.X), uintptr(radialStart.Y),
		uintptr(radialEnd.X), uintptr(radialEnd.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-closefigure
func (hdc HDC) CloseFigure() {
	ret, _, err := syscall.SyscallN(proc.CloseFigure.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func (hdc HDC) CreateCompatibleBitmap(cx, cy int32) HBITMAP {
	ret, _, err := syscall.SyscallN(proc.CreateCompatibleBitmap.Addr(),
		uintptr(hdc), uintptr(cx), uintptr(cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// ⚠️ You must defer HDC.DeleteDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func (hdc HDC) CreateCompatibleDC() HDC {
	ret, _, err := syscall.SyscallN(proc.CreateCompatibleDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HDC(ret)
}

// ⚠️ You must defer HBITMAP.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func (hdc HDC) CreateDIBSection(
	bmi *BITMAPINFO, usage co.DIB,
	hSection HFILEMAP, offset uint32) (HBITMAP, *byte) {

	var ppvBits *byte
	ret, _, err := syscall.SyscallN(proc.CreateDIBSection.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(bmi)), uintptr(usage),
		uintptr(unsafe.Pointer(&ppvBits)), uintptr(hSection), uintptr(offset))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret), ppvBits
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func (hdc HDC) DeleteDC() {
	ret, _, err := syscall.SyscallN(proc.DeleteDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ellipse
func (hdc HDC) Ellipse(bound RECT) {
	ret, _, err := syscall.SyscallN(proc.Ellipse.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpath
func (hdc HDC) EndPath() {
	ret, _, err := syscall.SyscallN(proc.EndPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillpath
func (hdc HDC) FillPath() {
	ret, _, err := syscall.SyscallN(proc.FillPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-fillrect
func (hdc HDC) FillRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.SyscallN(proc.FillRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-fillrgn
func (hdc HDC) FillRgn(hRgn HRGN, hBrush HBRUSH) {
	ret, _, err := syscall.SyscallN(proc.FillRgn.Addr(),
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-flattenpath
func (hdc HDC) FlattenPath() {
	ret, _, err := syscall.SyscallN(proc.FlattenPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-framergn
func (hdc HDC) FrameRgn(hRgn HRGN, hBrush HBRUSH, w, h int32) {
	ret, _, err := syscall.SyscallN(proc.FrameRgn.Addr(),
		uintptr(hdc), uintptr(hRgn), uintptr(hBrush), uintptr(w), uintptr(h))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getcurrentpositionex
func (hdc HDC) GetCurrentPositionEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetCurrentPositionEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcbrushcolor
func (hdc HDC) GetDCBrushColor() COLORREF {
	ret, _, err := syscall.SyscallN(proc.GetDCBrushColor.Addr(),
		uintptr(hdc))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdcpencolor
func (hdc HDC) GetDCPenColor() COLORREF {
	ret, _, err := syscall.SyscallN(proc.GetDCPenColor.Addr(),
		uintptr(hdc))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
func (hdc HDC) GetDeviceCaps(index co.GDC) int32 {
	ret, _, _ := syscall.SyscallN(proc.GetDeviceCaps.Addr(),
		uintptr(hdc), uintptr(index))
	return int32(ret)
}

// Note that this method fails if bitmapDataBuffer is an ordinary Go slice; it
// must be allocated directly from the OS heap.
//
// Example taking a screenshot:
//
//	cxScreen := win.GetSystemMetrics(co.SM_CXSCREEN)
//	cyScreen := win.GetSystemMetrics(co.SM_CYSCREEN)
//
//	hdcScreen := win.HWND(0).GetDC()
//	defer win.HWND(0).ReleaseDC(hdcScreen)
//
//	hBmp := hdcScreen.CreateCompatibleBitmap(cxScreen, cyScreen)
//	defer hBmp.DeleteObject()
//
//	hdcMem := hdcScreen.CreateCompatibleDC()
//	defer hdcMem.DeleteDC()
//
//	hBmpOld := hdcMem.SelectObjectBitmap(hBmp)
//	defer hdcMem.SelectObjectBitmap(hBmpOld)
//
//	hdcMem.BitBlt(
//		win.POINT{X: 0, Y: 0},
//		win.SIZE{Cx: cxScreen, Cy: cyScreen},
//		hdcScreen,
//		win.POINT{X: 0, Y: 0},
//		co.ROP_SRCCOPY,
//	)
//
//	bi := win.BITMAPINFO{
//		BmiHeader: win.BITMAPINFOHEADER{
//			BiWidth:       cxScreen,
//			BiHeight:      cyScreen,
//			BiPlanes:      1,
//			BiBitCount:    32,
//			BiCompression: co.BI_RGB,
//		},
//	}
//	bi.BmiHeader.SetBiSize()
//
//	bmpObj := win.BITMAP{}
//	hBmp.GetObject(&bmpObj)
//	bmpSize := bmpObj.CalcBitmapSize(bi.BmiHeader.BiBitCount)
//
//	rawMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, bmpSize)
//	defer rawMem.GlobalFree()
//
//	bmpSlice := rawMem.GlobalLock(bmpSize)
//	defer rawMem.GlobalUnlock()
//
//	hdcScreen.GetDIBits(hBmp, 0, int(cyScreen), bmpSlice, &bi, co.DIB_RGB_COLORS)
//
//	bfh := win.BITMAPFILEHEADER{}
//	bfh.SetBfType()
//	bfh.SetBfOffBits(uint32(unsafe.Sizeof(bfh) + unsafe.Sizeof(bi.BmiHeader)))
//	bfh.SetBfSize(bfh.BfOffBits() + uint32(bmpSize))
//
//	fo, _ := win.FileOpen("C:\\Temp\\foo.bmp", co.FILE_OPEN_RW_OPEN_OR_CREATE)
//	defer fo.Close()
//
//	fo.Write(bfh.Serialize())
//	fo.Write(bi.BmiHeader.Serialize())
//	fo.Write(bmpSlice)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
func (hdc HDC) GetDIBits(
	hbm HBITMAP,
	firstScanLine, numScanLines int,
	bitmapDataBuffer []byte, bmi *BITMAPINFO, usage co.DIB) int {

	var dataBufPtr *byte
	if bitmapDataBuffer != nil {
		dataBufPtr = &bitmapDataBuffer[0]
	}

	bmi.BmiHeader.SetBiSize() // safety

	ret, _, err := syscall.SyscallN(proc.GetDIBits.Addr(),
		uintptr(hdc), uintptr(hbm), uintptr(firstScanLine), uintptr(numScanLines),
		uintptr(unsafe.Pointer(dataBufPtr)), uintptr(unsafe.Pointer(bmi)),
		uintptr(usage))

	if wErr := errco.ERROR(ret); wErr == errco.INVALID_PARAMETER {
		panic(wErr)
	} else if ret == 0 {
		panic(errco.ERROR(err))
	}

	return int(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getpolyfillmode
func (hdc HDC) GetPolyFillMode() co.POLYF {
	ret, _, err := syscall.SyscallN(proc.GetPolyFillMode.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func (hdc HDC) GetTextExtentPoint32(text string) SIZE {
	var sz SIZE
	lpString16 := Str.ToNativeSlice(text)
	ret, _, err := syscall.SyscallN(proc.GetTextExtentPoint32.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1), uintptr(unsafe.Pointer(&sz)))
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextfacew
func (hdc HDC) GetTextFace() string {
	var buf [_LF_FACESIZE]uint16
	ret, _, err := syscall.SyscallN(proc.GetTextFace.Addr(),
		uintptr(hdc), uintptr(len(buf)), uintptr(unsafe.Pointer(&buf[0])))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gettextmetricsw
func (hdc HDC) GetTextMetrics(tm *TEXTMETRIC) {
	ret, _, err := syscall.SyscallN(proc.GetTextMetrics.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(tm)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportextex
func (hdc HDC) GetViewportExtEx() SIZE {
	var sz SIZE
	ret, _, err := syscall.SyscallN(proc.GetViewportExtEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getviewportorgex
func (hdc HDC) GetViewportOrgEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetViewportOrgEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindowextex
func (hdc HDC) GetWindowExtEx() SIZE {
	var sz SIZE
	ret, _, err := syscall.SyscallN(proc.GetWindowExtEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&sz)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getwindoworgex
func (hdc HDC) GetWindowOrgEx() POINT {
	var pt POINT
	ret, _, err := syscall.SyscallN(proc.GetWindowOrgEx.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-intersectcliprect
func (hdc HDC) IntersectClipRect(coords RECT) co.REGION {
	ret, _, err := syscall.SyscallN(proc.IntersectClipRect.Addr(),
		uintptr(hdc), uintptr(coords.Left), uintptr(coords.Top),
		uintptr(coords.Right), uintptr((coords.Bottom)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-invertrgn
func (hdc HDC) InvertRgn(hRgn HRGN) {
	ret, _, err := syscall.SyscallN(proc.InvertRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lineto
func (hdc HDC) LineTo(x, y int32) {
	ret, _, err := syscall.SyscallN(proc.LineTo.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-lptodp
func (hdc HDC) LPtoDP(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.LPtoDP.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-maskblt
func (hdc HDC) MaskBlt(
	destTopLeft POINT, sz SIZE, hdcSrc HDC, srcTopLeft POINT,
	hbmMask HBITMAP, maskOffset POINT, rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.MaskBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(sz.Cx), uintptr(sz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(hbmMask), uintptr(maskOffset.X), uintptr(maskOffset.Y),
		uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-movetoex
func (hdc HDC) MoveToEx(x, y int32, pt *POINT) {
	ret, _, err := syscall.SyscallN(proc.MoveToEx.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(pt)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-paintrgn
func (hdc HDC) PaintRgn(hRgn HRGN) {
	ret, _, err := syscall.SyscallN(proc.PaintRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HRGN.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pathtoregion
func (hdc HDC) PathToRegion() HRGN {
	ret, _, err := syscall.SyscallN(proc.PathToRegion.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HRGN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-pie
func (hdc HDC) Pie(bound RECT, endPointRadial1, endPointRadial2 POINT) {
	ret, _, err := syscall.SyscallN(proc.Pie.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(endPointRadial1.X), uintptr(endPointRadial1.Y),
		uintptr(endPointRadial2.X), uintptr(endPointRadial2.Y))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
func (hdc HDC) PolyDraw(pts []POINT, usage []co.PT) {
	if len(pts) != len(usage) {
		panic(fmt.Sprintf("PolyDraw different slice sizes: %d, %d.",
			len(pts), len(usage)))
	}
	ret, _, err := syscall.SyscallN(proc.PolyDraw.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])),
		uintptr(unsafe.Pointer(&usage[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polygon
func (hdc HDC) Polygon(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.Polygon.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polyline
func (hdc HDC) Polyline(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.Polyline.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polylineto
func (hdc HDC) PolylineTo(pts []POINT) {
	ret, _, err := syscall.SyscallN(proc.PolylineTo.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolygon
func (hdc HDC) PolyPolygon(pts [][]POINT) {
	totalPoints := 0
	for _, block := range pts {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)    // flat slice of all points
	blockCount := make([]int32, 0, len(pts)) // lengths of each block of points
	for _, block := range pts {
		flat = append(flat, block...)
		blockCount = append(blockCount, int32(len(block)))
	}

	ret, _, err := syscall.SyscallN(proc.PolyPolygon.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polypolyline
func (hdc HDC) PolyPolyline(pts [][]POINT) {
	totalPoints := 0
	for _, block := range pts {
		totalPoints += len(block)
	}

	flat := make([]POINT, 0, totalPoints)     // flat slice of all points
	blockCount := make([]uint32, 0, len(pts)) // lengths of each block of points
	for _, block := range pts {
		flat = append(flat, block...)
		blockCount = append(blockCount, uint32(len(block)))
	}

	ret, _, err := syscall.SyscallN(proc.PolyPolyline.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(&flat[0])),
		uintptr(unsafe.Pointer(&blockCount[0])), uintptr(len(blockCount)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-ptvisible
func (hdc HDC) PtVisible(x, y int32) bool {
	ret, _, err := syscall.SyscallN(proc.PtVisible.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rectangle
func (hdc HDC) Rectangle(bound RECT) {
	ret, _, err := syscall.SyscallN(proc.Rectangle.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Used together with HDC.SaveDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-restoredc
func (hdc HDC) RestoreDC(savedDC int32) {
	ret, _, err := syscall.SyscallN(proc.RestoreDC.Addr(),
		uintptr(hdc), uintptr(savedDC))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-roundrect
func (hdc HDC) RoundRect(bound RECT, sz SIZE) {
	ret, _, err := syscall.SyscallN(proc.RoundRect.Addr(),
		uintptr(hdc), uintptr(bound.Left), uintptr(bound.Top),
		uintptr(bound.Right), uintptr(bound.Bottom),
		uintptr(sz.Cx), uintptr(sz.Cy))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Used together with HDC.RestoreDC().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-savedc
func (hdc HDC) SaveDC() int32 {
	ret, _, err := syscall.SyscallN(proc.SaveDC.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectclippath
func (hdc HDC) SelectClipPath(mode co.RGN) {
	ret, _, err := syscall.SyscallN(proc.SelectClipPath.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectcliprgn
func (hdc HDC) SelectClipRgn(hRgn HRGN) co.REGION {
	ret, _, err := syscall.SyscallN(proc.SelectClipRgn.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == _REGION_ERROR {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBitmap(hBmp HBITMAP) HBITMAP {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hBmp))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBITMAP(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectBrush(hBrush HBRUSH) HBRUSH {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectFont(hFont HFONT) HFONT {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hFont))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HFONT(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectPen(hPen HPEN) HPEN {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hPen))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HPEN(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func (hdc HDC) SelectObjectRgn(hRgn HRGN) co.REGION {
	ret, _, err := syscall.SyscallN(proc.SelectObject.Addr(),
		uintptr(hdc), uintptr(hRgn))
	if ret == _HGDI_ERROR || ret == _REGION_ERROR {
		panic(errco.ERROR(err))
	}
	return co.REGION(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
func (hdc HDC) SetArcDirection(direction co.AD) co.AD {
	ret, _, err := syscall.SyscallN(proc.SetArcDirection.Addr(),
		uintptr(hdc), uintptr(direction))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.AD(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkcolor
func (hdc HDC) SetBkColor(color COLORREF) COLORREF {
	ret, _, err := syscall.SyscallN(proc.SetBkColor.Addr(),
		uintptr(hdc), uintptr(color))
	if ret == _CLR_INVALID {
		panic(errco.ERROR(err))
	}
	return COLORREF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
func (hdc HDC) SetBkMode(mode co.BKMODE) co.BKMODE {
	ret, _, err := syscall.SyscallN(proc.SetBkMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.BKMODE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
func (hdc HDC) SetPolyFillMode(mode co.POLYF) co.POLYF {
	ret, _, err := syscall.SyscallN(proc.SetPolyFillMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.POLYF(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
func (hdc HDC) SetStretchBltMode(mode co.STRETCH) co.STRETCH {
	ret, _, err := syscall.SyscallN(proc.SetStretchBltMode.Addr(),
		uintptr(hdc), uintptr(mode))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return co.STRETCH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
func (hdc HDC) SetTextAlign(align co.TA) {
	ret, _, err := syscall.SyscallN(proc.SetTextAlign.Addr(),
		uintptr(hdc), uintptr(align))
	if ret == _GDI_ERR {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-stretchblt
func (hdc HDC) StretchBlt(
	destTopLeft POINT, destSz SIZE,
	hdcSrc HDC, srcTopLeft POINT, srcSz SIZE, rop co.ROP) {

	ret, _, err := syscall.SyscallN(proc.StretchBlt.Addr(),
		uintptr(hdc), uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(hdcSrc), uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy), uintptr(rop))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokeandfillpath
func (hdc HDC) StrokeAndFillPath() {
	ret, _, err := syscall.SyscallN(proc.StrokeAndFillPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-strokepath
func (hdc HDC) StrokePath() {
	ret, _, err := syscall.SyscallN(proc.StrokePath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-textoutw
func (hdc HDC) TextOut(x, y int32, text string) {
	lpString16 := Str.ToNativeSlice(text)
	ret, _, err := syscall.SyscallN(proc.TextOut.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y),
		uintptr(unsafe.Pointer(&lpString16[0])),
		uintptr(len(lpString16)-1))
	runtime.KeepAlive(lpString16)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// This method is called from the destination HDC.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-transparentblt
func (hdc HDC) TransparentBlt(
	destTopLeft POINT, destSz SIZE,
	hdcSrc HDC, srcTopLeft POINT, srcSz SIZE,
	colorTransparent COLORREF) {

	ret, _, err := syscall.SyscallN(proc.TransparentBlt.Addr(),
		uintptr(hdc),
		uintptr(destTopLeft.X), uintptr(destTopLeft.Y),
		uintptr(destSz.Cx), uintptr(destSz.Cy),
		uintptr(hdcSrc),
		uintptr(srcTopLeft.X), uintptr(srcTopLeft.Y),
		uintptr(srcSz.Cx), uintptr(srcSz.Cy),
		uintptr(colorTransparent))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-widenpath
func (hdc HDC) WidenPath() {
	ret, _, err := syscall.SyscallN(proc.WidenPath.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-startdocw
func (hdc HDC) StartDoc(docName string, outputFile string) error {
	var res DOCINFO
	res.CbSize = int(unsafe.Sizeof(res))
	res.DocName = Str.ToNativePtr(docName)
	if outputFile != "" {
		res.Output = Str.ToNativePtr(outputFile)
	}
	addr := proc.StartDoc.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc), fromPtr(&res))
	if ret > 0 {
		return nil
	}
	// TODO: what if ret returns 0?
	return errco.ERROR(err)
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-enddoc
func (hdc HDC) EndDoc() error {
	addr := proc.EndDoc.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc))
	if ret > 0 {
		return nil
	}
	// TODO: what if ret returns 0?
	return errco.ERROR(err)
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setmapmode
// returns previous map mode
func (hdc HDC) SetMapMode(mapMode co.MM) (co.MM, error) {
	addr := proc.SetMapMode.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc), uintptr(mapMode))
	if ret == 0 {
		// TODO: is errco>ERROR() correct ?
		return co.MM(0), errco.ERROR(err)
	}
	return co.MM(ret), nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdcw
func CreateDC(driver string, device string, port string, pdm *DEVMODE) (HDC, error) {
	addr := proc.CreateDC.Addr()
	ret, _, err := syscall.SyscallN(addr, fromUtf8(driver), fromUtf8(device), fromUtf8(port), fromPtr(pdm))
	if ret == 0 {
		return HDC(0), errco.ERROR(err)
	}
	return HDC(ret), nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpage
func (hdc HDC) StartPage() error {
	addr := proc.StartPage.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc))
	if int(ret) <= 0 {
		return errco.ERROR(err)
	}
	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-endpage
func (hdc HDC) EndtPage() error {
	addr := proc.EndPage.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc))
	if int(ret) <= 0 {
		return errco.ERROR(err)
	}
	return nil
}

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-abortdoc
func (hdc HDC) AbortDoc() error {
	addr := proc.AbortDoc.Addr()
	ret, _, err := syscall.SyscallN(addr, uintptr(hdc))
	if int(ret) <= 0 {
		return errco.ERROR(err)
	}
	return nil
}
