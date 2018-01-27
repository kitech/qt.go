package qtwidgets

// /usr/include/qt/QtWidgets/qstyle.h
// #include <qstyle.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyle struct {
	*qtcore.QObject
}

func (this *QStyle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QStyle) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQStyleFromPointer(cthis unsafe.Pointer) *QStyle {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStyle{bcthis0}
}
func (*QStyle) NewFromPointer(cthis unsafe.Pointer) *QStyle {
	return NewQStyleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyle.h:66
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStyle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:73
// index:0
// Public
// void QStyle()
func NewQStyle() *QStyle {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyleC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyle.h:74
// index:0
// Public virtual
// void ~QStyle()
func DeleteQStyle(*QStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:76
// index:0
// Public virtual
// void polish(QWidget *)
func (this *QStyle) Polish(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:79
// index:1
// Public virtual
// void polish(QApplication *)
func (this *QStyle) Polish_1(application *QApplication /*777 QApplication **/) {
	var convArg0 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:82
// index:2
// Public virtual
// void polish(QPalette &)
func (this *QStyle) Polish_2(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishER8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:77
// index:0
// Public virtual
// void unpolish(QWidget *)
func (this *QStyle) Unpolish(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:80
// index:1
// Public virtual
// void unpolish(QApplication *)
func (this *QStyle) Unpolish_1(application *QApplication /*777 QApplication **/) {
	var convArg0 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:84
// index:0
// Public virtual
// QRect itemTextRect(const QFontMetrics &, const QRect &, int, bool, const QString &)
func (this *QStyle) ItemTextRect(fm *qtgui.QFontMetrics, r *qtcore.QRect, flags int, enabled bool, text *qtcore.QString) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = fm.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg4 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, flags, enabled, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:88
// index:0
// Public virtual
// QRect itemPixmapRect(const QRect &, int, const QPixmap &)
func (this *QStyle) ItemPixmapRect(r *qtcore.QRect, flags int, pixmap *qtgui.QPixmap) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = r.GetCthis()
	var convArg2 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, flags, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:90
// index:0
// Public virtual
// void drawItemText(QPainter *, const QRect &, int, const QPalette &, bool, const QString &, QPalette::ColorRole)
func (this *QStyle) DrawItemText(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, flags int, pal *qtgui.QPalette, enabled bool, text *qtcore.QString, textRole int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg3 = pal.GetCthis()
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags, convArg3, enabled, convArg5, textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:94
// index:0
// Public virtual
// void drawItemPixmap(QPainter *, const QRect &, int, const QPixmap &)
func (this *QStyle) DrawItemPixmap(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, alignment int, pixmap *qtgui.QPixmap) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:97
// index:0
// Public virtual
// QPalette standardPalette()
func (this *QStyle) StandardPalette() *qtgui.QPalette /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle15standardPaletteEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:204
// index:0
// Public pure virtual
// void drawPrimitive(QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QStyle) DrawPrimitive(pe int, opt *QStyleOption /*777 const QStyleOption **/, p *qtgui.QPainter /*777 QPainter **/, w *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13drawPrimitiveENS_16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:275
// index:0
// Public pure virtual
// void drawControl(QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QStyle) DrawControl(element int, opt *QStyleOption /*777 const QStyleOption **/, p *qtgui.QPainter /*777 QPainter **/, w *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11drawControlENS_14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:364
// index:0
// Public pure virtual
// QRect subElementRect(QStyle::SubElement, const QStyleOption *, const QWidget *)
func (this *QStyle) SubElementRect(subElement int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subElementRectENS_10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), subElement, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:443
// index:0
// Public pure virtual
// void drawComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *)
func (this *QStyle) DrawComplexControl(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, p *qtgui.QPainter /*777 QPainter **/, widget *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle18drawComplexControlENS_14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:445
// index:0
// Public pure virtual
// QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *)
func (this *QStyle) HitTestComplexControl(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, pt *qtcore.QPoint, widget *QWidget /*777 const QWidget **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = pt.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle21hitTestComplexControlENS_14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:447
// index:0
// Public pure virtual
// QRect subControlRect(QStyle::ComplexControl, const QStyleOptionComplex *, QStyle::SubControl, const QWidget *)
func (this *QStyle) SubControlRect(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, sc int, widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = opt.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subControlRectENS_14ComplexControlEPK19QStyleOptionComplexNS_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), cc, convArg1, sc, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:582
// index:0
// Public pure virtual
// int pixelMetric(QStyle::PixelMetric, const QStyleOption *, const QWidget *)
func (this *QStyle) PixelMetric(metric int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11pixelMetricENS_11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstyle.h:614
// index:0
// Public pure virtual
// QSize sizeFromContents(QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *)
func (this *QStyle) SizeFromContents(ct int, opt *QStyleOption /*777 const QStyleOption **/, contentsSize *qtcore.QSize, w *QWidget /*777 const QWidget **/) *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = opt.GetCthis()
	var convArg2 = contentsSize.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle16sizeFromContentsENS_12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), ct, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:748
// index:0
// Public pure virtual
// int styleHint(QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *)
func (this *QStyle) StyleHint(stylehint int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/, returnData *QStyleHintReturn /*777 QStyleHintReturn **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	var convArg3 = returnData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle9styleHintENS_9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stylehint, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstyle.h:828
// index:0
// Public pure virtual
// QPixmap standardPixmap(QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QStyle) StandardPixmap(standardPixmap int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14standardPixmapENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), standardPixmap, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:831
// index:0
// Public pure virtual
// QIcon standardIcon(QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QStyle) StandardIcon(standardIcon int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QIcon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12standardIconENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), standardIcon, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:834
// index:0
// Public pure virtual
// QPixmap generatedIconPixmap(QIcon::Mode, const QPixmap &, const QStyleOption *)
func (this *QStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt *QStyleOption /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = pixmap.GetCthis()
	var convArg2 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), iconMode, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstyle.h:837
// index:0
// Public static
// QRect visualRect(Qt::LayoutDirection, const QRect &, const QRect &)
func (this *QStyle) VisualRect(direction int, boundingRect *qtcore.QRect, logicalRect *qtcore.QRect) *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle10visualRectEN2Qt15LayoutDirectionERK5QRectS4_", ffiqt.FFI_TYPE_POINTER, direction, boundingRect, logicalRect)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStyle_VisualRect(direction int, boundingRect *qtcore.QRect, logicalRect *qtcore.QRect) *qtcore.QRect /*123*/ {
	var nilthis *QStyle
	rv := nilthis.VisualRect(direction, boundingRect, logicalRect)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:839
// index:0
// Public static
// QPoint visualPos(Qt::LayoutDirection, const QRect &, const QPoint &)
func (this *QStyle) VisualPos(direction int, boundingRect *qtcore.QRect, logicalPos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle9visualPosEN2Qt15LayoutDirectionERK5QRectRK6QPoint", ffiqt.FFI_TYPE_POINTER, direction, boundingRect, logicalPos)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QStyle_VisualPos(direction int, boundingRect *qtcore.QRect, logicalPos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var nilthis *QStyle
	rv := nilthis.VisualPos(direction, boundingRect, logicalPos)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:841
// index:0
// Public static
// int sliderPositionFromValue(int, int, int, int, bool)
func (this *QStyle) SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderPositionFromValueEiiiib", ffiqt.FFI_TYPE_POINTER, min, max, val, space, upsideDown)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStyle_SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) int {
	var nilthis *QStyle
	rv := nilthis.SliderPositionFromValue(min, max, val, space, upsideDown)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:843
// index:0
// Public static
// int sliderValueFromPosition(int, int, int, int, bool)
func (this *QStyle) SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderValueFromPositionEiiiib", ffiqt.FFI_TYPE_POINTER, min, max, pos, space, upsideDown)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStyle_SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) int {
	var nilthis *QStyle
	rv := nilthis.SliderValueFromPosition(min, max, pos, space, upsideDown)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:849
// index:0
// Public pure virtual
// int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *)
func (this *QStyle) LayoutSpacing(control1 int, control2 int, orientation int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg3 = option.GetCthis()
	var convArg4 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstyle.h:856
// index:0
// Public
// const QStyle * proxy()
func (this *QStyle) Proxy() *QStyle /*777 const QStyle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle5proxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QStyle__StateFlag = int

const QStyle__State_None QStyle__StateFlag = 0
const QStyle__State_Enabled QStyle__StateFlag = 1
const QStyle__State_Raised QStyle__StateFlag = 2
const QStyle__State_Sunken QStyle__StateFlag = 4
const QStyle__State_Off QStyle__StateFlag = 8
const QStyle__State_NoChange QStyle__StateFlag = 16
const QStyle__State_On QStyle__StateFlag = 32
const QStyle__State_DownArrow QStyle__StateFlag = 64
const QStyle__State_Horizontal QStyle__StateFlag = 128
const QStyle__State_HasFocus QStyle__StateFlag = 256
const QStyle__State_Top QStyle__StateFlag = 512
const QStyle__State_Bottom QStyle__StateFlag = 1024
const QStyle__State_FocusAtBorder QStyle__StateFlag = 2048
const QStyle__State_AutoRaise QStyle__StateFlag = 4096
const QStyle__State_MouseOver QStyle__StateFlag = 8192
const QStyle__State_UpArrow QStyle__StateFlag = 16384
const QStyle__State_Selected QStyle__StateFlag = 32768
const QStyle__State_Active QStyle__StateFlag = 65536
const QStyle__State_Window QStyle__StateFlag = 131072
const QStyle__State_Open QStyle__StateFlag = 262144
const QStyle__State_Children QStyle__StateFlag = 524288
const QStyle__State_Item QStyle__StateFlag = 1048576
const QStyle__State_Sibling QStyle__StateFlag = 2097152
const QStyle__State_Editing QStyle__StateFlag = 4194304
const QStyle__State_KeyboardFocusChange QStyle__StateFlag = 8388608
const QStyle__State_ReadOnly QStyle__StateFlag = 33554432
const QStyle__State_Small QStyle__StateFlag = 67108864
const QStyle__State_Mini QStyle__StateFlag = 134217728

type QStyle__PrimitiveElement = int

const QStyle__PE_Frame QStyle__PrimitiveElement = 0
const QStyle__PE_FrameDefaultButton QStyle__PrimitiveElement = 1
const QStyle__PE_FrameDockWidget QStyle__PrimitiveElement = 2
const QStyle__PE_FrameFocusRect QStyle__PrimitiveElement = 3
const QStyle__PE_FrameGroupBox QStyle__PrimitiveElement = 4
const QStyle__PE_FrameLineEdit QStyle__PrimitiveElement = 5
const QStyle__PE_FrameMenu QStyle__PrimitiveElement = 6
const QStyle__PE_FrameStatusBar QStyle__PrimitiveElement = 7
const QStyle__PE_FrameStatusBarItem QStyle__PrimitiveElement = 7
const QStyle__PE_FrameTabWidget QStyle__PrimitiveElement = 8
const QStyle__PE_FrameWindow QStyle__PrimitiveElement = 9
const QStyle__PE_FrameButtonBevel QStyle__PrimitiveElement = 10
const QStyle__PE_FrameButtonTool QStyle__PrimitiveElement = 11
const QStyle__PE_FrameTabBarBase QStyle__PrimitiveElement = 12
const QStyle__PE_PanelButtonCommand QStyle__PrimitiveElement = 13
const QStyle__PE_PanelButtonBevel QStyle__PrimitiveElement = 14
const QStyle__PE_PanelButtonTool QStyle__PrimitiveElement = 15
const QStyle__PE_PanelMenuBar QStyle__PrimitiveElement = 16
const QStyle__PE_PanelToolBar QStyle__PrimitiveElement = 17
const QStyle__PE_PanelLineEdit QStyle__PrimitiveElement = 18
const QStyle__PE_IndicatorArrowDown QStyle__PrimitiveElement = 19
const QStyle__PE_IndicatorArrowLeft QStyle__PrimitiveElement = 20
const QStyle__PE_IndicatorArrowRight QStyle__PrimitiveElement = 21
const QStyle__PE_IndicatorArrowUp QStyle__PrimitiveElement = 22
const QStyle__PE_IndicatorBranch QStyle__PrimitiveElement = 23
const QStyle__PE_IndicatorButtonDropDown QStyle__PrimitiveElement = 24
const QStyle__PE_IndicatorViewItemCheck QStyle__PrimitiveElement = 25
const QStyle__PE_IndicatorItemViewItemCheck QStyle__PrimitiveElement = 25
const QStyle__PE_IndicatorCheckBox QStyle__PrimitiveElement = 26
const QStyle__PE_IndicatorDockWidgetResizeHandle QStyle__PrimitiveElement = 27
const QStyle__PE_IndicatorHeaderArrow QStyle__PrimitiveElement = 28
const QStyle__PE_IndicatorMenuCheckMark QStyle__PrimitiveElement = 29
const QStyle__PE_IndicatorProgressChunk QStyle__PrimitiveElement = 30
const QStyle__PE_IndicatorRadioButton QStyle__PrimitiveElement = 31
const QStyle__PE_IndicatorSpinDown QStyle__PrimitiveElement = 32
const QStyle__PE_IndicatorSpinMinus QStyle__PrimitiveElement = 33
const QStyle__PE_IndicatorSpinPlus QStyle__PrimitiveElement = 34
const QStyle__PE_IndicatorSpinUp QStyle__PrimitiveElement = 35
const QStyle__PE_IndicatorToolBarHandle QStyle__PrimitiveElement = 36
const QStyle__PE_IndicatorToolBarSeparator QStyle__PrimitiveElement = 37
const QStyle__PE_PanelTipLabel QStyle__PrimitiveElement = 38
const QStyle__PE_IndicatorTabTear QStyle__PrimitiveElement = 39
const QStyle__PE_IndicatorTabTearLeft QStyle__PrimitiveElement = 39
const QStyle__PE_PanelScrollAreaCorner QStyle__PrimitiveElement = 40
const QStyle__PE_Widget QStyle__PrimitiveElement = 41
const QStyle__PE_IndicatorColumnViewArrow QStyle__PrimitiveElement = 42
const QStyle__PE_IndicatorItemViewItemDrop QStyle__PrimitiveElement = 43
const QStyle__PE_PanelItemViewItem QStyle__PrimitiveElement = 44
const QStyle__PE_PanelItemViewRow QStyle__PrimitiveElement = 45
const QStyle__PE_PanelStatusBar QStyle__PrimitiveElement = 46
const QStyle__PE_IndicatorTabClose QStyle__PrimitiveElement = 47
const QStyle__PE_PanelMenu QStyle__PrimitiveElement = 48
const QStyle__PE_IndicatorTabTearRight QStyle__PrimitiveElement = 49
const QStyle__PE_CustomBase QStyle__PrimitiveElement = 251658240

type QStyle__ControlElement = int

const QStyle__CE_PushButton QStyle__ControlElement = 0
const QStyle__CE_PushButtonBevel QStyle__ControlElement = 1
const QStyle__CE_PushButtonLabel QStyle__ControlElement = 2
const QStyle__CE_CheckBox QStyle__ControlElement = 3
const QStyle__CE_CheckBoxLabel QStyle__ControlElement = 4
const QStyle__CE_RadioButton QStyle__ControlElement = 5
const QStyle__CE_RadioButtonLabel QStyle__ControlElement = 6
const QStyle__CE_TabBarTab QStyle__ControlElement = 7
const QStyle__CE_TabBarTabShape QStyle__ControlElement = 8
const QStyle__CE_TabBarTabLabel QStyle__ControlElement = 9
const QStyle__CE_ProgressBar QStyle__ControlElement = 10
const QStyle__CE_ProgressBarGroove QStyle__ControlElement = 11
const QStyle__CE_ProgressBarContents QStyle__ControlElement = 12
const QStyle__CE_ProgressBarLabel QStyle__ControlElement = 13
const QStyle__CE_MenuItem QStyle__ControlElement = 14
const QStyle__CE_MenuScroller QStyle__ControlElement = 15
const QStyle__CE_MenuVMargin QStyle__ControlElement = 16
const QStyle__CE_MenuHMargin QStyle__ControlElement = 17
const QStyle__CE_MenuTearoff QStyle__ControlElement = 18
const QStyle__CE_MenuEmptyArea QStyle__ControlElement = 19
const QStyle__CE_MenuBarItem QStyle__ControlElement = 20
const QStyle__CE_MenuBarEmptyArea QStyle__ControlElement = 21
const QStyle__CE_ToolButtonLabel QStyle__ControlElement = 22
const QStyle__CE_Header QStyle__ControlElement = 23
const QStyle__CE_HeaderSection QStyle__ControlElement = 24
const QStyle__CE_HeaderLabel QStyle__ControlElement = 25
const QStyle__CE_ToolBoxTab QStyle__ControlElement = 26
const QStyle__CE_SizeGrip QStyle__ControlElement = 27
const QStyle__CE_Splitter QStyle__ControlElement = 28
const QStyle__CE_RubberBand QStyle__ControlElement = 29
const QStyle__CE_DockWidgetTitle QStyle__ControlElement = 30
const QStyle__CE_ScrollBarAddLine QStyle__ControlElement = 31
const QStyle__CE_ScrollBarSubLine QStyle__ControlElement = 32
const QStyle__CE_ScrollBarAddPage QStyle__ControlElement = 33
const QStyle__CE_ScrollBarSubPage QStyle__ControlElement = 34
const QStyle__CE_ScrollBarSlider QStyle__ControlElement = 35
const QStyle__CE_ScrollBarFirst QStyle__ControlElement = 36
const QStyle__CE_ScrollBarLast QStyle__ControlElement = 37
const QStyle__CE_FocusFrame QStyle__ControlElement = 38
const QStyle__CE_ComboBoxLabel QStyle__ControlElement = 39
const QStyle__CE_ToolBar QStyle__ControlElement = 40
const QStyle__CE_ToolBoxTabShape QStyle__ControlElement = 41
const QStyle__CE_ToolBoxTabLabel QStyle__ControlElement = 42
const QStyle__CE_HeaderEmptyArea QStyle__ControlElement = 43
const QStyle__CE_ColumnViewGrip QStyle__ControlElement = 44
const QStyle__CE_ItemViewItem QStyle__ControlElement = 45
const QStyle__CE_ShapedFrame QStyle__ControlElement = 46
const QStyle__CE_CustomBase QStyle__ControlElement = 4026531840

type QStyle__SubElement = int

const QStyle__SE_PushButtonContents QStyle__SubElement = 0
const QStyle__SE_PushButtonFocusRect QStyle__SubElement = 1
const QStyle__SE_CheckBoxIndicator QStyle__SubElement = 2
const QStyle__SE_CheckBoxContents QStyle__SubElement = 3
const QStyle__SE_CheckBoxFocusRect QStyle__SubElement = 4
const QStyle__SE_CheckBoxClickRect QStyle__SubElement = 5
const QStyle__SE_RadioButtonIndicator QStyle__SubElement = 6
const QStyle__SE_RadioButtonContents QStyle__SubElement = 7
const QStyle__SE_RadioButtonFocusRect QStyle__SubElement = 8
const QStyle__SE_RadioButtonClickRect QStyle__SubElement = 9
const QStyle__SE_ComboBoxFocusRect QStyle__SubElement = 10
const QStyle__SE_SliderFocusRect QStyle__SubElement = 11
const QStyle__SE_ProgressBarGroove QStyle__SubElement = 12
const QStyle__SE_ProgressBarContents QStyle__SubElement = 13
const QStyle__SE_ProgressBarLabel QStyle__SubElement = 14
const QStyle__SE_ToolBoxTabContents QStyle__SubElement = 15
const QStyle__SE_HeaderLabel QStyle__SubElement = 16
const QStyle__SE_HeaderArrow QStyle__SubElement = 17
const QStyle__SE_TabWidgetTabBar QStyle__SubElement = 18
const QStyle__SE_TabWidgetTabPane QStyle__SubElement = 19
const QStyle__SE_TabWidgetTabContents QStyle__SubElement = 20
const QStyle__SE_TabWidgetLeftCorner QStyle__SubElement = 21
const QStyle__SE_TabWidgetRightCorner QStyle__SubElement = 22
const QStyle__SE_ViewItemCheckIndicator QStyle__SubElement = 23
const QStyle__SE_ItemViewItemCheckIndicator QStyle__SubElement = 23
const QStyle__SE_TabBarTearIndicator QStyle__SubElement = 24
const QStyle__SE_TabBarTearIndicatorLeft QStyle__SubElement = 24
const QStyle__SE_TreeViewDisclosureItem QStyle__SubElement = 25
const QStyle__SE_LineEditContents QStyle__SubElement = 26
const QStyle__SE_FrameContents QStyle__SubElement = 27
const QStyle__SE_DockWidgetCloseButton QStyle__SubElement = 28
const QStyle__SE_DockWidgetFloatButton QStyle__SubElement = 29
const QStyle__SE_DockWidgetTitleBarText QStyle__SubElement = 30
const QStyle__SE_DockWidgetIcon QStyle__SubElement = 31
const QStyle__SE_CheckBoxLayoutItem QStyle__SubElement = 32
const QStyle__SE_ComboBoxLayoutItem QStyle__SubElement = 33
const QStyle__SE_DateTimeEditLayoutItem QStyle__SubElement = 34
const QStyle__SE_DialogButtonBoxLayoutItem QStyle__SubElement = 35
const QStyle__SE_LabelLayoutItem QStyle__SubElement = 36
const QStyle__SE_ProgressBarLayoutItem QStyle__SubElement = 37
const QStyle__SE_PushButtonLayoutItem QStyle__SubElement = 38
const QStyle__SE_RadioButtonLayoutItem QStyle__SubElement = 39
const QStyle__SE_SliderLayoutItem QStyle__SubElement = 40
const QStyle__SE_SpinBoxLayoutItem QStyle__SubElement = 41
const QStyle__SE_ToolButtonLayoutItem QStyle__SubElement = 42
const QStyle__SE_FrameLayoutItem QStyle__SubElement = 43
const QStyle__SE_GroupBoxLayoutItem QStyle__SubElement = 44
const QStyle__SE_TabWidgetLayoutItem QStyle__SubElement = 45
const QStyle__SE_ItemViewItemDecoration QStyle__SubElement = 46
const QStyle__SE_ItemViewItemText QStyle__SubElement = 47
const QStyle__SE_ItemViewItemFocusRect QStyle__SubElement = 48
const QStyle__SE_TabBarTabLeftButton QStyle__SubElement = 49
const QStyle__SE_TabBarTabRightButton QStyle__SubElement = 50
const QStyle__SE_TabBarTabText QStyle__SubElement = 51
const QStyle__SE_ShapedFrameContents QStyle__SubElement = 52
const QStyle__SE_ToolBarHandle QStyle__SubElement = 53
const QStyle__SE_TabBarScrollLeftButton QStyle__SubElement = 54
const QStyle__SE_TabBarScrollRightButton QStyle__SubElement = 55
const QStyle__SE_TabBarTearIndicatorRight QStyle__SubElement = 56
const QStyle__SE_CustomBase QStyle__SubElement = 4026531840

type QStyle__ComplexControl = int

const QStyle__CC_SpinBox QStyle__ComplexControl = 0
const QStyle__CC_ComboBox QStyle__ComplexControl = 1
const QStyle__CC_ScrollBar QStyle__ComplexControl = 2
const QStyle__CC_Slider QStyle__ComplexControl = 3
const QStyle__CC_ToolButton QStyle__ComplexControl = 4
const QStyle__CC_TitleBar QStyle__ComplexControl = 5
const QStyle__CC_Dial QStyle__ComplexControl = 6
const QStyle__CC_GroupBox QStyle__ComplexControl = 7
const QStyle__CC_MdiControls QStyle__ComplexControl = 8
const QStyle__CC_CustomBase QStyle__ComplexControl = 4026531840

type QStyle__SubControl = int

const QStyle__SC_None QStyle__SubControl = 0
const QStyle__SC_ScrollBarAddLine QStyle__SubControl = 1
const QStyle__SC_ScrollBarSubLine QStyle__SubControl = 2
const QStyle__SC_ScrollBarAddPage QStyle__SubControl = 4
const QStyle__SC_ScrollBarSubPage QStyle__SubControl = 8
const QStyle__SC_ScrollBarFirst QStyle__SubControl = 16
const QStyle__SC_ScrollBarLast QStyle__SubControl = 32
const QStyle__SC_ScrollBarSlider QStyle__SubControl = 64
const QStyle__SC_ScrollBarGroove QStyle__SubControl = 128
const QStyle__SC_SpinBoxUp QStyle__SubControl = 1
const QStyle__SC_SpinBoxDown QStyle__SubControl = 2
const QStyle__SC_SpinBoxFrame QStyle__SubControl = 4
const QStyle__SC_SpinBoxEditField QStyle__SubControl = 8
const QStyle__SC_ComboBoxFrame QStyle__SubControl = 1
const QStyle__SC_ComboBoxEditField QStyle__SubControl = 2
const QStyle__SC_ComboBoxArrow QStyle__SubControl = 4
const QStyle__SC_ComboBoxListBoxPopup QStyle__SubControl = 8
const QStyle__SC_SliderGroove QStyle__SubControl = 1
const QStyle__SC_SliderHandle QStyle__SubControl = 2
const QStyle__SC_SliderTickmarks QStyle__SubControl = 4
const QStyle__SC_ToolButton QStyle__SubControl = 1
const QStyle__SC_ToolButtonMenu QStyle__SubControl = 2
const QStyle__SC_TitleBarSysMenu QStyle__SubControl = 1
const QStyle__SC_TitleBarMinButton QStyle__SubControl = 2
const QStyle__SC_TitleBarMaxButton QStyle__SubControl = 4
const QStyle__SC_TitleBarCloseButton QStyle__SubControl = 8
const QStyle__SC_TitleBarNormalButton QStyle__SubControl = 16
const QStyle__SC_TitleBarShadeButton QStyle__SubControl = 32
const QStyle__SC_TitleBarUnshadeButton QStyle__SubControl = 64
const QStyle__SC_TitleBarContextHelpButton QStyle__SubControl = 128
const QStyle__SC_TitleBarLabel QStyle__SubControl = 256
const QStyle__SC_DialGroove QStyle__SubControl = 1
const QStyle__SC_DialHandle QStyle__SubControl = 2
const QStyle__SC_DialTickmarks QStyle__SubControl = 4
const QStyle__SC_GroupBoxCheckBox QStyle__SubControl = 1
const QStyle__SC_GroupBoxLabel QStyle__SubControl = 2
const QStyle__SC_GroupBoxContents QStyle__SubControl = 4
const QStyle__SC_GroupBoxFrame QStyle__SubControl = 8
const QStyle__SC_MdiMinButton QStyle__SubControl = 1
const QStyle__SC_MdiNormalButton QStyle__SubControl = 2
const QStyle__SC_MdiCloseButton QStyle__SubControl = 4
const QStyle__SC_CustomBase QStyle__SubControl = 4026531840
const QStyle__SC_All QStyle__SubControl = 4294967295

type QStyle__PixelMetric = int

const QStyle__PM_ButtonMargin QStyle__PixelMetric = 0
const QStyle__PM_ButtonDefaultIndicator QStyle__PixelMetric = 1
const QStyle__PM_MenuButtonIndicator QStyle__PixelMetric = 2
const QStyle__PM_ButtonShiftHorizontal QStyle__PixelMetric = 3
const QStyle__PM_ButtonShiftVertical QStyle__PixelMetric = 4
const QStyle__PM_DefaultFrameWidth QStyle__PixelMetric = 5
const QStyle__PM_SpinBoxFrameWidth QStyle__PixelMetric = 6
const QStyle__PM_ComboBoxFrameWidth QStyle__PixelMetric = 7
const QStyle__PM_MaximumDragDistance QStyle__PixelMetric = 8
const QStyle__PM_ScrollBarExtent QStyle__PixelMetric = 9
const QStyle__PM_ScrollBarSliderMin QStyle__PixelMetric = 10
const QStyle__PM_SliderThickness QStyle__PixelMetric = 11
const QStyle__PM_SliderControlThickness QStyle__PixelMetric = 12
const QStyle__PM_SliderLength QStyle__PixelMetric = 13
const QStyle__PM_SliderTickmarkOffset QStyle__PixelMetric = 14
const QStyle__PM_SliderSpaceAvailable QStyle__PixelMetric = 15
const QStyle__PM_DockWidgetSeparatorExtent QStyle__PixelMetric = 16
const QStyle__PM_DockWidgetHandleExtent QStyle__PixelMetric = 17
const QStyle__PM_DockWidgetFrameWidth QStyle__PixelMetric = 18
const QStyle__PM_TabBarTabOverlap QStyle__PixelMetric = 19
const QStyle__PM_TabBarTabHSpace QStyle__PixelMetric = 20
const QStyle__PM_TabBarTabVSpace QStyle__PixelMetric = 21
const QStyle__PM_TabBarBaseHeight QStyle__PixelMetric = 22
const QStyle__PM_TabBarBaseOverlap QStyle__PixelMetric = 23
const QStyle__PM_ProgressBarChunkWidth QStyle__PixelMetric = 24
const QStyle__PM_SplitterWidth QStyle__PixelMetric = 25
const QStyle__PM_TitleBarHeight QStyle__PixelMetric = 26
const QStyle__PM_MenuScrollerHeight QStyle__PixelMetric = 27
const QStyle__PM_MenuHMargin QStyle__PixelMetric = 28
const QStyle__PM_MenuVMargin QStyle__PixelMetric = 29
const QStyle__PM_MenuPanelWidth QStyle__PixelMetric = 30
const QStyle__PM_MenuTearoffHeight QStyle__PixelMetric = 31
const QStyle__PM_MenuDesktopFrameWidth QStyle__PixelMetric = 32
const QStyle__PM_MenuBarPanelWidth QStyle__PixelMetric = 33
const QStyle__PM_MenuBarItemSpacing QStyle__PixelMetric = 34
const QStyle__PM_MenuBarVMargin QStyle__PixelMetric = 35
const QStyle__PM_MenuBarHMargin QStyle__PixelMetric = 36
const QStyle__PM_IndicatorWidth QStyle__PixelMetric = 37
const QStyle__PM_IndicatorHeight QStyle__PixelMetric = 38
const QStyle__PM_ExclusiveIndicatorWidth QStyle__PixelMetric = 39
const QStyle__PM_ExclusiveIndicatorHeight QStyle__PixelMetric = 40
const QStyle__PM_DialogButtonsSeparator QStyle__PixelMetric = 41
const QStyle__PM_DialogButtonsButtonWidth QStyle__PixelMetric = 42
const QStyle__PM_DialogButtonsButtonHeight QStyle__PixelMetric = 43
const QStyle__PM_MdiSubWindowFrameWidth QStyle__PixelMetric = 44
const QStyle__PM_MDIFrameWidth QStyle__PixelMetric = 44
const QStyle__PM_MdiSubWindowMinimizedWidth QStyle__PixelMetric = 45
const QStyle__PM_MDIMinimizedWidth QStyle__PixelMetric = 45
const QStyle__PM_HeaderMargin QStyle__PixelMetric = 46
const QStyle__PM_HeaderMarkSize QStyle__PixelMetric = 47
const QStyle__PM_HeaderGripMargin QStyle__PixelMetric = 48
const QStyle__PM_TabBarTabShiftHorizontal QStyle__PixelMetric = 49
const QStyle__PM_TabBarTabShiftVertical QStyle__PixelMetric = 50
const QStyle__PM_TabBarScrollButtonWidth QStyle__PixelMetric = 51
const QStyle__PM_ToolBarFrameWidth QStyle__PixelMetric = 52
const QStyle__PM_ToolBarHandleExtent QStyle__PixelMetric = 53
const QStyle__PM_ToolBarItemSpacing QStyle__PixelMetric = 54
const QStyle__PM_ToolBarItemMargin QStyle__PixelMetric = 55
const QStyle__PM_ToolBarSeparatorExtent QStyle__PixelMetric = 56
const QStyle__PM_ToolBarExtensionExtent QStyle__PixelMetric = 57
const QStyle__PM_SpinBoxSliderHeight QStyle__PixelMetric = 58
const QStyle__PM_DefaultTopLevelMargin QStyle__PixelMetric = 59
const QStyle__PM_DefaultChildMargin QStyle__PixelMetric = 60
const QStyle__PM_DefaultLayoutSpacing QStyle__PixelMetric = 61
const QStyle__PM_ToolBarIconSize QStyle__PixelMetric = 62
const QStyle__PM_ListViewIconSize QStyle__PixelMetric = 63
const QStyle__PM_IconViewIconSize QStyle__PixelMetric = 64
const QStyle__PM_SmallIconSize QStyle__PixelMetric = 65
const QStyle__PM_LargeIconSize QStyle__PixelMetric = 66
const QStyle__PM_FocusFrameVMargin QStyle__PixelMetric = 67
const QStyle__PM_FocusFrameHMargin QStyle__PixelMetric = 68
const QStyle__PM_ToolTipLabelFrameWidth QStyle__PixelMetric = 69
const QStyle__PM_CheckBoxLabelSpacing QStyle__PixelMetric = 70
const QStyle__PM_TabBarIconSize QStyle__PixelMetric = 71
const QStyle__PM_SizeGripSize QStyle__PixelMetric = 72
const QStyle__PM_DockWidgetTitleMargin QStyle__PixelMetric = 73
const QStyle__PM_MessageBoxIconSize QStyle__PixelMetric = 74
const QStyle__PM_ButtonIconSize QStyle__PixelMetric = 75
const QStyle__PM_DockWidgetTitleBarButtonMargin QStyle__PixelMetric = 76
const QStyle__PM_RadioButtonLabelSpacing QStyle__PixelMetric = 77
const QStyle__PM_LayoutLeftMargin QStyle__PixelMetric = 78
const QStyle__PM_LayoutTopMargin QStyle__PixelMetric = 79
const QStyle__PM_LayoutRightMargin QStyle__PixelMetric = 80
const QStyle__PM_LayoutBottomMargin QStyle__PixelMetric = 81
const QStyle__PM_LayoutHorizontalSpacing QStyle__PixelMetric = 82
const QStyle__PM_LayoutVerticalSpacing QStyle__PixelMetric = 83
const QStyle__PM_TabBar_ScrollButtonOverlap QStyle__PixelMetric = 84
const QStyle__PM_TextCursorWidth QStyle__PixelMetric = 85
const QStyle__PM_TabCloseIndicatorWidth QStyle__PixelMetric = 86
const QStyle__PM_TabCloseIndicatorHeight QStyle__PixelMetric = 87
const QStyle__PM_ScrollView_ScrollBarSpacing QStyle__PixelMetric = 88
const QStyle__PM_ScrollView_ScrollBarOverlap QStyle__PixelMetric = 89
const QStyle__PM_SubMenuOverlap QStyle__PixelMetric = 90
const QStyle__PM_TreeViewIndentation QStyle__PixelMetric = 91
const QStyle__PM_HeaderDefaultSectionSizeHorizontal QStyle__PixelMetric = 92
const QStyle__PM_HeaderDefaultSectionSizeVertical QStyle__PixelMetric = 93
const QStyle__PM_TitleBarButtonIconSize QStyle__PixelMetric = 94
const QStyle__PM_TitleBarButtonSize QStyle__PixelMetric = 95
const QStyle__PM_CustomBase QStyle__PixelMetric = 4026531840

type QStyle__ContentsType = int

const QStyle__CT_PushButton QStyle__ContentsType = 0
const QStyle__CT_CheckBox QStyle__ContentsType = 1
const QStyle__CT_RadioButton QStyle__ContentsType = 2
const QStyle__CT_ToolButton QStyle__ContentsType = 3
const QStyle__CT_ComboBox QStyle__ContentsType = 4
const QStyle__CT_Splitter QStyle__ContentsType = 5
const QStyle__CT_ProgressBar QStyle__ContentsType = 6
const QStyle__CT_MenuItem QStyle__ContentsType = 7
const QStyle__CT_MenuBarItem QStyle__ContentsType = 8
const QStyle__CT_MenuBar QStyle__ContentsType = 9
const QStyle__CT_Menu QStyle__ContentsType = 10
const QStyle__CT_TabBarTab QStyle__ContentsType = 11
const QStyle__CT_Slider QStyle__ContentsType = 12
const QStyle__CT_ScrollBar QStyle__ContentsType = 13
const QStyle__CT_LineEdit QStyle__ContentsType = 14
const QStyle__CT_SpinBox QStyle__ContentsType = 15
const QStyle__CT_SizeGrip QStyle__ContentsType = 16
const QStyle__CT_TabWidget QStyle__ContentsType = 17
const QStyle__CT_DialogButtons QStyle__ContentsType = 18
const QStyle__CT_HeaderSection QStyle__ContentsType = 19
const QStyle__CT_GroupBox QStyle__ContentsType = 20
const QStyle__CT_MdiControls QStyle__ContentsType = 21
const QStyle__CT_ItemViewItem QStyle__ContentsType = 22
const QStyle__CT_CustomBase QStyle__ContentsType = 4026531840

type QStyle__RequestSoftwareInputPanel = int

const QStyle__RSIP_OnMouseClickAndAlreadyFocused QStyle__RequestSoftwareInputPanel = 0
const QStyle__RSIP_OnMouseClick QStyle__RequestSoftwareInputPanel = 1

type QStyle__StyleHint = int

const QStyle__SH_EtchDisabledText QStyle__StyleHint = 0
const QStyle__SH_DitherDisabledText QStyle__StyleHint = 1
const QStyle__SH_ScrollBar_MiddleClickAbsolutePosition QStyle__StyleHint = 2
const QStyle__SH_ScrollBar_ScrollWhenPointerLeavesControl QStyle__StyleHint = 3
const QStyle__SH_TabBar_SelectMouseType QStyle__StyleHint = 4
const QStyle__SH_TabBar_Alignment QStyle__StyleHint = 5
const QStyle__SH_Header_ArrowAlignment QStyle__StyleHint = 6
const QStyle__SH_Slider_SnapToValue QStyle__StyleHint = 7
const QStyle__SH_Slider_SloppyKeyEvents QStyle__StyleHint = 8
const QStyle__SH_ProgressDialog_CenterCancelButton QStyle__StyleHint = 9
const QStyle__SH_ProgressDialog_TextLabelAlignment QStyle__StyleHint = 10
const QStyle__SH_PrintDialog_RightAlignButtons QStyle__StyleHint = 11
const QStyle__SH_MainWindow_SpaceBelowMenuBar QStyle__StyleHint = 12
const QStyle__SH_FontDialog_SelectAssociatedText QStyle__StyleHint = 13
const QStyle__SH_Menu_AllowActiveAndDisabled QStyle__StyleHint = 14
const QStyle__SH_Menu_SpaceActivatesItem QStyle__StyleHint = 15
const QStyle__SH_Menu_SubMenuPopupDelay QStyle__StyleHint = 16
const QStyle__SH_ScrollView_FrameOnlyAroundContents QStyle__StyleHint = 17
const QStyle__SH_MenuBar_AltKeyNavigation QStyle__StyleHint = 18
const QStyle__SH_ComboBox_ListMouseTracking QStyle__StyleHint = 19
const QStyle__SH_Menu_MouseTracking QStyle__StyleHint = 20
const QStyle__SH_MenuBar_MouseTracking QStyle__StyleHint = 21
const QStyle__SH_ItemView_ChangeHighlightOnFocus QStyle__StyleHint = 22
const QStyle__SH_Widget_ShareActivation QStyle__StyleHint = 23
const QStyle__SH_Workspace_FillSpaceOnMaximize QStyle__StyleHint = 24
const QStyle__SH_ComboBox_Popup QStyle__StyleHint = 25
const QStyle__SH_TitleBar_NoBorder QStyle__StyleHint = 26
const QStyle__SH_Slider_StopMouseOverSlider QStyle__StyleHint = 27
const QStyle__SH_ScrollBar_StopMouseOverSlider QStyle__StyleHint = 27
const QStyle__SH_BlinkCursorWhenTextSelected QStyle__StyleHint = 28
const QStyle__SH_RichText_FullWidthSelection QStyle__StyleHint = 29
const QStyle__SH_Menu_Scrollable QStyle__StyleHint = 30
const QStyle__SH_GroupBox_TextLabelVerticalAlignment QStyle__StyleHint = 31
const QStyle__SH_GroupBox_TextLabelColor QStyle__StyleHint = 32
const QStyle__SH_Menu_SloppySubMenus QStyle__StyleHint = 33
const QStyle__SH_Table_GridLineColor QStyle__StyleHint = 34
const QStyle__SH_LineEdit_PasswordCharacter QStyle__StyleHint = 35
const QStyle__SH_DialogButtons_DefaultButton QStyle__StyleHint = 36
const QStyle__SH_ToolBox_SelectedPageTitleBold QStyle__StyleHint = 37
const QStyle__SH_TabBar_PreferNoArrows QStyle__StyleHint = 38
const QStyle__SH_ScrollBar_LeftClickAbsolutePosition QStyle__StyleHint = 39
const QStyle__SH_ListViewExpand_SelectMouseType QStyle__StyleHint = 40
const QStyle__SH_UnderlineShortcut QStyle__StyleHint = 41
const QStyle__SH_SpinBox_AnimateButton QStyle__StyleHint = 42
const QStyle__SH_SpinBox_KeyPressAutoRepeatRate QStyle__StyleHint = 43
const QStyle__SH_SpinBox_ClickAutoRepeatRate QStyle__StyleHint = 44
const QStyle__SH_Menu_FillScreenWithScroll QStyle__StyleHint = 45
const QStyle__SH_ToolTipLabel_Opacity QStyle__StyleHint = 46
const QStyle__SH_DrawMenuBarSeparator QStyle__StyleHint = 47
const QStyle__SH_TitleBar_ModifyNotification QStyle__StyleHint = 48
const QStyle__SH_Button_FocusPolicy QStyle__StyleHint = 49
const QStyle__SH_MessageBox_UseBorderForButtonSpacing QStyle__StyleHint = 50
const QStyle__SH_TitleBar_AutoRaise QStyle__StyleHint = 51
const QStyle__SH_ToolButton_PopupDelay QStyle__StyleHint = 52
const QStyle__SH_FocusFrame_Mask QStyle__StyleHint = 53
const QStyle__SH_RubberBand_Mask QStyle__StyleHint = 54
const QStyle__SH_WindowFrame_Mask QStyle__StyleHint = 55
const QStyle__SH_SpinControls_DisableOnBounds QStyle__StyleHint = 56
const QStyle__SH_Dial_BackgroundRole QStyle__StyleHint = 57
const QStyle__SH_ComboBox_LayoutDirection QStyle__StyleHint = 58
const QStyle__SH_ItemView_EllipsisLocation QStyle__StyleHint = 59
const QStyle__SH_ItemView_ShowDecorationSelected QStyle__StyleHint = 60
const QStyle__SH_ItemView_ActivateItemOnSingleClick QStyle__StyleHint = 61
const QStyle__SH_ScrollBar_ContextMenu QStyle__StyleHint = 62
const QStyle__SH_ScrollBar_RollBetweenButtons QStyle__StyleHint = 63
const QStyle__SH_Slider_AbsoluteSetButtons QStyle__StyleHint = 64
const QStyle__SH_Slider_PageSetButtons QStyle__StyleHint = 65
const QStyle__SH_Menu_KeyboardSearch QStyle__StyleHint = 66
const QStyle__SH_TabBar_ElideMode QStyle__StyleHint = 67
const QStyle__SH_DialogButtonLayout QStyle__StyleHint = 68
const QStyle__SH_ComboBox_PopupFrameStyle QStyle__StyleHint = 69
const QStyle__SH_MessageBox_TextInteractionFlags QStyle__StyleHint = 70
const QStyle__SH_DialogButtonBox_ButtonsHaveIcons QStyle__StyleHint = 71
const QStyle__SH_SpellCheckUnderlineStyle QStyle__StyleHint = 72
const QStyle__SH_MessageBox_CenterButtons QStyle__StyleHint = 73
const QStyle__SH_Menu_SelectionWrap QStyle__StyleHint = 74
const QStyle__SH_ItemView_MovementWithoutUpdatingSelection QStyle__StyleHint = 75
const QStyle__SH_ToolTip_Mask QStyle__StyleHint = 76
const QStyle__SH_FocusFrame_AboveWidget QStyle__StyleHint = 77
const QStyle__SH_TextControl_FocusIndicatorTextCharFormat QStyle__StyleHint = 78
const QStyle__SH_WizardStyle QStyle__StyleHint = 79
const QStyle__SH_ItemView_ArrowKeysNavigateIntoChildren QStyle__StyleHint = 80
const QStyle__SH_Menu_Mask QStyle__StyleHint = 81
const QStyle__SH_Menu_FlashTriggeredItem QStyle__StyleHint = 82
const QStyle__SH_Menu_FadeOutOnHide QStyle__StyleHint = 83
const QStyle__SH_SpinBox_ClickAutoRepeatThreshold QStyle__StyleHint = 84
const QStyle__SH_ItemView_PaintAlternatingRowColorsForEmptyArea QStyle__StyleHint = 85
const QStyle__SH_FormLayoutWrapPolicy QStyle__StyleHint = 86
const QStyle__SH_TabWidget_DefaultTabPosition QStyle__StyleHint = 87
const QStyle__SH_ToolBar_Movable QStyle__StyleHint = 88
const QStyle__SH_FormLayoutFieldGrowthPolicy QStyle__StyleHint = 89
const QStyle__SH_FormLayoutFormAlignment QStyle__StyleHint = 90
const QStyle__SH_FormLayoutLabelAlignment QStyle__StyleHint = 91
const QStyle__SH_ItemView_DrawDelegateFrame QStyle__StyleHint = 92
const QStyle__SH_TabBar_CloseButtonPosition QStyle__StyleHint = 93
const QStyle__SH_DockWidget_ButtonsHaveFrame QStyle__StyleHint = 94
const QStyle__SH_ToolButtonStyle QStyle__StyleHint = 95
const QStyle__SH_RequestSoftwareInputPanel QStyle__StyleHint = 96
const QStyle__SH_ScrollBar_Transient QStyle__StyleHint = 97
const QStyle__SH_Menu_SupportsSections QStyle__StyleHint = 98
const QStyle__SH_ToolTip_WakeUpDelay QStyle__StyleHint = 99
const QStyle__SH_ToolTip_FallAsleepDelay QStyle__StyleHint = 100
const QStyle__SH_Widget_Animate QStyle__StyleHint = 101
const QStyle__SH_Splitter_OpaqueResize QStyle__StyleHint = 102
const QStyle__SH_ComboBox_UseNativePopup QStyle__StyleHint = 103
const QStyle__SH_LineEdit_PasswordMaskDelay QStyle__StyleHint = 104
const QStyle__SH_TabBar_ChangeCurrentDelay QStyle__StyleHint = 105
const QStyle__SH_Menu_SubMenuUniDirection QStyle__StyleHint = 106
const QStyle__SH_Menu_SubMenuUniDirectionFailCount QStyle__StyleHint = 107
const QStyle__SH_Menu_SubMenuSloppySelectOtherActions QStyle__StyleHint = 108
const QStyle__SH_Menu_SubMenuSloppyCloseTimeout QStyle__StyleHint = 109
const QStyle__SH_Menu_SubMenuResetWhenReenteringParent QStyle__StyleHint = 110
const QStyle__SH_Menu_SubMenuDontStartSloppyOnLeave QStyle__StyleHint = 111
const QStyle__SH_ItemView_ScrollMode QStyle__StyleHint = 112
const QStyle__SH_TitleBar_ShowToolTipsOnButtons QStyle__StyleHint = 113
const QStyle__SH_Widget_Animation_Duration QStyle__StyleHint = 114
const QStyle__SH_CustomBase QStyle__StyleHint = 4026531840

type QStyle__StandardPixmap = int

const QStyle__SP_TitleBarMenuButton QStyle__StandardPixmap = 0
const QStyle__SP_TitleBarMinButton QStyle__StandardPixmap = 1
const QStyle__SP_TitleBarMaxButton QStyle__StandardPixmap = 2
const QStyle__SP_TitleBarCloseButton QStyle__StandardPixmap = 3
const QStyle__SP_TitleBarNormalButton QStyle__StandardPixmap = 4
const QStyle__SP_TitleBarShadeButton QStyle__StandardPixmap = 5
const QStyle__SP_TitleBarUnshadeButton QStyle__StandardPixmap = 6
const QStyle__SP_TitleBarContextHelpButton QStyle__StandardPixmap = 7
const QStyle__SP_DockWidgetCloseButton QStyle__StandardPixmap = 8
const QStyle__SP_MessageBoxInformation QStyle__StandardPixmap = 9
const QStyle__SP_MessageBoxWarning QStyle__StandardPixmap = 10
const QStyle__SP_MessageBoxCritical QStyle__StandardPixmap = 11
const QStyle__SP_MessageBoxQuestion QStyle__StandardPixmap = 12
const QStyle__SP_DesktopIcon QStyle__StandardPixmap = 13
const QStyle__SP_TrashIcon QStyle__StandardPixmap = 14
const QStyle__SP_ComputerIcon QStyle__StandardPixmap = 15
const QStyle__SP_DriveFDIcon QStyle__StandardPixmap = 16
const QStyle__SP_DriveHDIcon QStyle__StandardPixmap = 17
const QStyle__SP_DriveCDIcon QStyle__StandardPixmap = 18
const QStyle__SP_DriveDVDIcon QStyle__StandardPixmap = 19
const QStyle__SP_DriveNetIcon QStyle__StandardPixmap = 20
const QStyle__SP_DirOpenIcon QStyle__StandardPixmap = 21
const QStyle__SP_DirClosedIcon QStyle__StandardPixmap = 22
const QStyle__SP_DirLinkIcon QStyle__StandardPixmap = 23
const QStyle__SP_DirLinkOpenIcon QStyle__StandardPixmap = 24
const QStyle__SP_FileIcon QStyle__StandardPixmap = 25
const QStyle__SP_FileLinkIcon QStyle__StandardPixmap = 26
const QStyle__SP_ToolBarHorizontalExtensionButton QStyle__StandardPixmap = 27
const QStyle__SP_ToolBarVerticalExtensionButton QStyle__StandardPixmap = 28
const QStyle__SP_FileDialogStart QStyle__StandardPixmap = 29
const QStyle__SP_FileDialogEnd QStyle__StandardPixmap = 30
const QStyle__SP_FileDialogToParent QStyle__StandardPixmap = 31
const QStyle__SP_FileDialogNewFolder QStyle__StandardPixmap = 32
const QStyle__SP_FileDialogDetailedView QStyle__StandardPixmap = 33
const QStyle__SP_FileDialogInfoView QStyle__StandardPixmap = 34
const QStyle__SP_FileDialogContentsView QStyle__StandardPixmap = 35
const QStyle__SP_FileDialogListView QStyle__StandardPixmap = 36
const QStyle__SP_FileDialogBack QStyle__StandardPixmap = 37
const QStyle__SP_DirIcon QStyle__StandardPixmap = 38
const QStyle__SP_DialogOkButton QStyle__StandardPixmap = 39
const QStyle__SP_DialogCancelButton QStyle__StandardPixmap = 40
const QStyle__SP_DialogHelpButton QStyle__StandardPixmap = 41
const QStyle__SP_DialogOpenButton QStyle__StandardPixmap = 42
const QStyle__SP_DialogSaveButton QStyle__StandardPixmap = 43
const QStyle__SP_DialogCloseButton QStyle__StandardPixmap = 44
const QStyle__SP_DialogApplyButton QStyle__StandardPixmap = 45
const QStyle__SP_DialogResetButton QStyle__StandardPixmap = 46
const QStyle__SP_DialogDiscardButton QStyle__StandardPixmap = 47
const QStyle__SP_DialogYesButton QStyle__StandardPixmap = 48
const QStyle__SP_DialogNoButton QStyle__StandardPixmap = 49
const QStyle__SP_ArrowUp QStyle__StandardPixmap = 50
const QStyle__SP_ArrowDown QStyle__StandardPixmap = 51
const QStyle__SP_ArrowLeft QStyle__StandardPixmap = 52
const QStyle__SP_ArrowRight QStyle__StandardPixmap = 53
const QStyle__SP_ArrowBack QStyle__StandardPixmap = 54
const QStyle__SP_ArrowForward QStyle__StandardPixmap = 55
const QStyle__SP_DirHomeIcon QStyle__StandardPixmap = 56
const QStyle__SP_CommandLink QStyle__StandardPixmap = 57
const QStyle__SP_VistaShield QStyle__StandardPixmap = 58
const QStyle__SP_BrowserReload QStyle__StandardPixmap = 59
const QStyle__SP_BrowserStop QStyle__StandardPixmap = 60
const QStyle__SP_MediaPlay QStyle__StandardPixmap = 61
const QStyle__SP_MediaStop QStyle__StandardPixmap = 62
const QStyle__SP_MediaPause QStyle__StandardPixmap = 63
const QStyle__SP_MediaSkipForward QStyle__StandardPixmap = 64
const QStyle__SP_MediaSkipBackward QStyle__StandardPixmap = 65
const QStyle__SP_MediaSeekForward QStyle__StandardPixmap = 66
const QStyle__SP_MediaSeekBackward QStyle__StandardPixmap = 67
const QStyle__SP_MediaVolume QStyle__StandardPixmap = 68
const QStyle__SP_MediaVolumeMuted QStyle__StandardPixmap = 69
const QStyle__SP_LineEditClearButton QStyle__StandardPixmap = 70
const QStyle__SP_CustomBase QStyle__StandardPixmap = 4026531840

//  body block end
