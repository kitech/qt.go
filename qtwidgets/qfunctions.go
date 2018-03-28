package qtwidgets

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtWidgets/qsizepolicy.h:68
// index:44
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSizePolicy, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_44(key QSizePolicy_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSizePolicy_PTR() != nil {
		convArg0 = key.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash11QSizePolicyj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qdrawutil.h:147
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawBorderPixmap(QPainter *, const QRect &, const QMargins &, const QPixmap &, const QRect &, const QMargins &, const QTileRules &, QDrawBorderPixmap::DrawingHints)

/*

 */
func QDrawBorderPixmap(painter qtgui.QPainter_ITF /*777 QPainter **/, targetRect qtcore.QRect_ITF, targetMargins qtcore.QMargins_ITF, pixmap qtgui.QPixmap_ITF, sourceRect qtcore.QRect_ITF, sourceMargins qtcore.QMargins_ITF, rules QTileRules_ITF, hints int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetRect != nil && targetRect.QRect_PTR() != nil {
		convArg1 = targetRect.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if targetMargins != nil && targetMargins.QMargins_PTR() != nil {
		convArg2 = targetMargins.QMargins_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg3 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if sourceRect != nil && sourceRect.QRect_PTR() != nil {
		convArg4 = sourceRect.QRect_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if sourceMargins != nil && sourceMargins.QMargins_PTR() != nil {
		convArg5 = sourceMargins.QMargins_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if rules != nil && rules.QTileRules_PTR() != nil {
		convArg6 = rules.QTileRules_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z17qDrawBorderPixmapP8QPainterRK5QRectRK8QMarginsRK7QPixmapS3_S6_RK10QTileRules6QFlagsIN17QDrawBorderPixmap11DrawingHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:159
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void qDrawBorderPixmap(QPainter *, const QRect &, const QMargins &, const QPixmap &)

/*

 */
func QDrawBorderPixmap_1(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRect_ITF, margins qtcore.QMargins_ITF, pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if target != nil && target.QRect_PTR() != nil {
		convArg1 = target.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg2 = margins.QMargins_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg3 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z17qDrawBorderPixmapP8QPainterRK5QRectRK8QMarginsRK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:80
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadePanel(QPainter *, int, int, int, int, const QPalette &, bool, int, const QBrush *)

/*

 */
func QDrawShadePanel(p qtgui.QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg5 = pal.QPalette_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg8 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z15qDrawShadePanelP8QPainteriiiiRK8QPalettebiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, lineWidth, convArg8)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:84
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadePanel(QPainter *, const QRect &, const QPalette &, bool, int, const QBrush *)

/*

 */
func QDrawShadePanel_1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg5 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z15qDrawShadePanelP8QPainterRK5QRectRK8QPalettebiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, lineWidth, convArg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:88
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinButton(QPainter *, int, int, int, int, const QPalette &, bool, const QBrush *)

/*

 */
func QDrawWinButton(p qtgui.QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg5 = pal.QPalette_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg7 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawWinButtonP8QPainteriiiiRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:92
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinButton(QPainter *, const QRect &, const QPalette &, bool, const QBrush *)

/*

 */
func QDrawWinButton_1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg4 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawWinButtonP8QPainterRK5QRectRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:70
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeRect(QPainter *, int, int, int, int, const QPalette &, bool, int, int, const QBrush *)

/*

 */
func QDrawShadeRect(p qtgui.QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg5 = pal.QPalette_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg9 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeRectP8QPainteriiiiRK8QPalettebiiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, lineWidth, midLineWidth, convArg9)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:75
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeRect(QPainter *, const QRect &, const QPalette &, bool, int, int, const QBrush *)

/*

 */
func QDrawShadeRect_1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg6 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeRectP8QPainterRK5QRectRK8QPalettebiiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, lineWidth, midLineWidth, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:62
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeLine(QPainter *, int, int, int, int, const QPalette &, bool, int, int)

/*

 */
func QDrawShadeLine(p qtgui.QPainter_ITF /*777 QPainter **/, x1 int, y1 int, x2 int, y2 int, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg5 = pal.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeLineP8QPainteriiiiRK8QPalettebii", qtrt.FFI_TYPE_POINTER, convArg0, x1, y1, x2, y2, convArg5, sunken, lineWidth, midLineWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:66
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeLine(QPainter *, const QPoint &, const QPoint &, const QPalette &, bool, int, int)

/*

 */
func QDrawShadeLine_1(p qtgui.QPainter_ITF /*777 QPainter **/, p1 qtcore.QPoint_ITF, p2 qtcore.QPoint_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if p1 != nil && p1.QPoint_PTR() != nil {
		convArg1 = p1.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p2 != nil && p2.QPoint_PTR() != nil {
		convArg2 = p2.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg3 = pal.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeLineP8QPainterRK6QPointS3_RK8QPalettebii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, sunken, lineWidth, midLineWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:104
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawPlainRect(QPainter *, int, int, int, int, const QColor &, int, const QBrush *)

/*

 */
func QDrawPlainRect(p qtgui.QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, arg5 qtgui.QColor_ITF, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if arg5 != nil && arg5.QColor_PTR() != nil {
		convArg5 = arg5.QColor_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg7 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawPlainRectP8QPainteriiiiRK6QColoriPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, lineWidth, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:107
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawPlainRect(QPainter *, const QRect &, const QColor &, int, const QBrush *)

/*

 */
func QDrawPlainRect_1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, arg2 qtgui.QColor_ITF, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if arg2 != nil && arg2.QColor_PTR() != nil {
		convArg2 = arg2.QColor_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg4 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawPlainRectP8QPainterRK5QRectRK6QColoriPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, lineWidth, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:96
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinPanel(QPainter *, int, int, int, int, const QPalette &, bool, const QBrush *)

/*

 */
func QDrawWinPanel(p qtgui.QPainter_ITF /*777 QPainter **/, x int, y int, w int, h int, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg5 = pal.QPalette_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg7 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z13qDrawWinPanelP8QPainteriiiiRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:100
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinPanel(QPainter *, const QRect &, const QPalette &, bool, const QBrush *)

/*

 */
func QDrawWinPanel_1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg1 = r.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pal != nil && pal.QPalette_PTR() != nil {
		convArg2 = pal.QPalette_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if fill != nil && fill.QBrush_PTR() != nil {
		convArg4 = fill.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z13qDrawWinPanelP8QPainterRK5QRectRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, convArg4)
	qtrt.ErrPrint(err, rv)
}

//  body block end
