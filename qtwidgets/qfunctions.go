package qtwidgets

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

func init_unused_10027() {
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
// /usr/include/qt/QtWidgets/qwidget.h:737
// index:111
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QWidget::RenderFlags::enum_type, int)

/*

 */
func Operator_or111(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN7QWidget10RenderFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:871
// index:112
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QStyle::State::enum_type, int)

/*

 */
func Operator_or112(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN6QStyle9StateFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:872
// index:113
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QStyle::SubControls::enum_type, int)

/*

 */
func Operator_or113(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN6QStyle10SubControlEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleoption.h:586
// index:115
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QStyleOptionToolButton::ToolButtonFeatures::enum_type, int)

/*

 */
func Operator_or115(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN22QStyleOptionToolButton17ToolButtonFeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleoption.h:255
// index:119
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QStyleOptionButton::ButtonFeatures::enum_type, int)

/*

 */
func Operator_or119(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN18QStyleOptionButton13ButtonFeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesturerecognizer.h:89
// index:120
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QGestureRecognizer::Result::enum_type, int)

/*

 */
func Operator_or120(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN18QGestureRecognizer10ResultFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleoption.h:150
// index:121
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QStyleOptionFrame::FrameFeatures::enum_type, int)

/*

 */
func Operator_or121(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN17QStyleOptionFrame12FrameFeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:78
// index:122
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QFileIconProvider::Options::enum_type, int)

/*

 */
func Operator_or122(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN17QFileIconProvider6OptionEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:198
// index:129
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QPinchGesture::ChangeFlags::enum_type, int)

/*

 */
func Operator_or129(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QPinchGesture10ChangeFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:234
// index:137
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSizePolicy::ControlTypes::enum_type, int)

/*

 */
func Operator_or137(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QSizePolicy11ControlTypeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:68
// index:53
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSizePolicy, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash53(key QSizePolicy_ITF /*123*/, seed uint) uint {
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
func QDrawBorderPixmap1(painter qtgui.QPainter_ITF /*777 QPainter **/, target qtcore.QRect_ITF, margins qtcore.QMargins_ITF, pixmap qtgui.QPixmap_ITF) {
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
func QDrawShadePanel1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
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
func QDrawWinButton1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
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
func QDrawShadeRect1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
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
func QDrawShadeLine1(p qtgui.QPainter_ITF /*777 QPainter **/, p1 qtcore.QPoint_ITF, p2 qtcore.QPoint_ITF, pal qtgui.QPalette_ITF, sunken bool, lineWidth int, midLineWidth int) {
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
func QDrawPlainRect1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, arg2 qtgui.QColor_ITF, lineWidth int, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
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
func QDrawWinPanel1(p qtgui.QPainter_ITF /*777 QPainter **/, r qtcore.QRect_ITF, pal qtgui.QPalette_ITF, sunken bool, fill qtgui.QBrush_ITF /*777 const QBrush **/) {
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
