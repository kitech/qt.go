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
// index:43
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSizePolicy, uint)
func QHash_43(key *QSizePolicy /*123*/, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash11QSizePolicyj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qdrawutil.h:147
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawBorderPixmap(QPainter *, const QRect &, const QMargins &, const QPixmap &, const QRect &, const QMargins &, const struct QTileRules &, QDrawBorderPixmap::DrawingHints)
func QDrawBorderPixmap(painter *qtgui.QPainter /*777 QPainter **/, targetRect *qtcore.QRect, targetMargins *qtcore.QMargins, pixmap *qtgui.QPixmap, sourceRect *qtcore.QRect, sourceMargins *qtcore.QMargins, rules *QTileRules, hints int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = targetRect.GetCthis()
	var convArg2 = targetMargins.GetCthis()
	var convArg3 = pixmap.GetCthis()
	var convArg4 = sourceRect.GetCthis()
	var convArg5 = sourceMargins.GetCthis()
	var convArg6 = rules.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z17qDrawBorderPixmapP8QPainterRK5QRectRK8QMarginsRK7QPixmapS3_S6_RK10QTileRules6QFlagsIN17QDrawBorderPixmap11DrawingHintEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:159
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void qDrawBorderPixmap(QPainter *, const QRect &, const QMargins &, const QPixmap &)
func QDrawBorderPixmap_1(painter *qtgui.QPainter /*777 QPainter **/, target *qtcore.QRect, margins *qtcore.QMargins, pixmap *qtgui.QPixmap) {
	var convArg0 = painter.GetCthis()
	var convArg1 = target.GetCthis()
	var convArg2 = margins.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z17qDrawBorderPixmapP8QPainterRK5QRectRK8QMarginsRK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:80
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadePanel(QPainter *, int, int, int, int, const QPalette &, _Bool, int, const QBrush *)
func QDrawShadePanel(p *qtgui.QPainter /*777 QPainter **/, x int, y int, w int, h int, pal *qtgui.QPalette, sunken bool, lineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg5 = pal.GetCthis()
	var convArg8 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z15qDrawShadePanelP8QPainteriiiiRK8QPalettebiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, lineWidth, convArg8)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:84
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadePanel(QPainter *, const QRect &, const QPalette &, _Bool, int, const QBrush *)
func QDrawShadePanel_1(p *qtgui.QPainter /*777 QPainter **/, r *qtcore.QRect, pal *qtgui.QPalette, sunken bool, lineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg2 = pal.GetCthis()
	var convArg5 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z15qDrawShadePanelP8QPainterRK5QRectRK8QPalettebiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, lineWidth, convArg5)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:88
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinButton(QPainter *, int, int, int, int, const QPalette &, _Bool, const QBrush *)
func QDrawWinButton(p *qtgui.QPainter /*777 QPainter **/, x int, y int, w int, h int, pal *qtgui.QPalette, sunken bool, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg5 = pal.GetCthis()
	var convArg7 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawWinButtonP8QPainteriiiiRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:92
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinButton(QPainter *, const QRect &, const QPalette &, _Bool, const QBrush *)
func QDrawWinButton_1(p *qtgui.QPainter /*777 QPainter **/, r *qtcore.QRect, pal *qtgui.QPalette, sunken bool, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg2 = pal.GetCthis()
	var convArg4 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawWinButtonP8QPainterRK5QRectRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:70
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeRect(QPainter *, int, int, int, int, const QPalette &, _Bool, int, int, const QBrush *)
func QDrawShadeRect(p *qtgui.QPainter /*777 QPainter **/, x int, y int, w int, h int, pal *qtgui.QPalette, sunken bool, lineWidth int, midLineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg5 = pal.GetCthis()
	var convArg9 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeRectP8QPainteriiiiRK8QPalettebiiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, lineWidth, midLineWidth, convArg9)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:75
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeRect(QPainter *, const QRect &, const QPalette &, _Bool, int, int, const QBrush *)
func QDrawShadeRect_1(p *qtgui.QPainter /*777 QPainter **/, r *qtcore.QRect, pal *qtgui.QPalette, sunken bool, lineWidth int, midLineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg2 = pal.GetCthis()
	var convArg6 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeRectP8QPainterRK5QRectRK8QPalettebiiPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, lineWidth, midLineWidth, convArg6)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:62
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeLine(QPainter *, int, int, int, int, const QPalette &, _Bool, int, int)
func QDrawShadeLine(p *qtgui.QPainter /*777 QPainter **/, x1 int, y1 int, x2 int, y2 int, pal *qtgui.QPalette, sunken bool, lineWidth int, midLineWidth int) {
	var convArg0 = p.GetCthis()
	var convArg5 = pal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeLineP8QPainteriiiiRK8QPalettebii", qtrt.FFI_TYPE_POINTER, convArg0, x1, y1, x2, y2, convArg5, sunken, lineWidth, midLineWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:66
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawShadeLine(QPainter *, const QPoint &, const QPoint &, const QPalette &, _Bool, int, int)
func QDrawShadeLine_1(p *qtgui.QPainter /*777 QPainter **/, p1 *qtcore.QPoint, p2 *qtcore.QPoint, pal *qtgui.QPalette, sunken bool, lineWidth int, midLineWidth int) {
	var convArg0 = p.GetCthis()
	var convArg1 = p1.GetCthis()
	var convArg2 = p2.GetCthis()
	var convArg3 = pal.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawShadeLineP8QPainterRK6QPointS3_RK8QPalettebii", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, sunken, lineWidth, midLineWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:104
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawPlainRect(QPainter *, int, int, int, int, const QColor &, int, const QBrush *)
func QDrawPlainRect(p *qtgui.QPainter /*777 QPainter **/, x int, y int, w int, h int, arg5 *qtgui.QColor, lineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg5 = arg5.GetCthis()
	var convArg7 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawPlainRectP8QPainteriiiiRK6QColoriPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, lineWidth, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:107
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawPlainRect(QPainter *, const QRect &, const QColor &, int, const QBrush *)
func QDrawPlainRect_1(p *qtgui.QPainter /*777 QPainter **/, r *qtcore.QRect, arg2 *qtgui.QColor, lineWidth int, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg2 = arg2.GetCthis()
	var convArg4 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z14qDrawPlainRectP8QPainterRK5QRectRK6QColoriPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, lineWidth, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:96
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinPanel(QPainter *, int, int, int, int, const QPalette &, _Bool, const QBrush *)
func QDrawWinPanel(p *qtgui.QPainter /*777 QPainter **/, x int, y int, w int, h int, pal *qtgui.QPalette, sunken bool, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg5 = pal.GetCthis()
	var convArg7 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z13qDrawWinPanelP8QPainteriiiiRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, x, y, w, h, convArg5, sunken, convArg7)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdrawutil.h:100
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void qDrawWinPanel(QPainter *, const QRect &, const QPalette &, _Bool, const QBrush *)
func QDrawWinPanel_1(p *qtgui.QPainter /*777 QPainter **/, r *qtcore.QRect, pal *qtgui.QPalette, sunken bool, fill *qtgui.QBrush /*777 const QBrush **/) {
	var convArg0 = p.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg2 = pal.GetCthis()
	var convArg4 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z13qDrawWinPanelP8QPainterRK5QRectRK8QPalettebPK6QBrush", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, sunken, convArg4)
	qtrt.ErrPrint(err, rv)
}

//  body block end
