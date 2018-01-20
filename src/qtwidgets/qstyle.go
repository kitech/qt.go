//  header block begin
// /usr/include/qt/QtWidgets/qstyle.h
// #include <qstyle.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	return this.QObject.GetCthis()
}
func NewQStyleFromPointer(cthis unsafe.Pointer) *QStyle {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStyle{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyle.h:66
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStyle) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
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
// void polish(class QWidget *)
func (this *QStyle) Polish(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:79
// index:1
// Public virtual
// void polish(class QApplication *)
func (this *QStyle) Polish_1(application unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), application)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:82
// index:2
// Public virtual
// void polish(class QPalette &)
func (this *QStyle) Polish_2(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle6polishER8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:77
// index:0
// Public virtual
// void unpolish(class QWidget *)
func (this *QStyle) Unpolish(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:80
// index:1
// Public virtual
// void unpolish(class QApplication *)
func (this *QStyle) Unpolish_1(application unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), application)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:84
// index:0
// Public virtual
// QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QStyle) ItemTextRect(fm *qtgui.QFontMetrics, r *qtcore.QRect, flags int, enabled bool, text *qtcore.QString) interface{} {
	var convArg0 = fm.GetCthis()
	var convArg1 = r.GetCthis()
	var convArg4 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &flags, &enabled, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:88
// index:0
// Public virtual
// QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QStyle) ItemPixmapRect(r *qtcore.QRect, flags int, pixmap *qtgui.QPixmap) interface{} {
	var convArg0 = r.GetCthis()
	var convArg2 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &flags, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:90
// index:0
// Public virtual
// void drawItemText(class QPainter *, const class QRect &, int, const class QPalette &, _Bool, const class QString &, class QPalette::ColorRole)
func (this *QStyle) DrawItemText(painter unsafe.Pointer, rect *qtcore.QRect, flags int, pal *qtgui.QPalette, enabled bool, text *qtcore.QString, textRole int) {
	var convArg1 = rect.GetCthis()
	var convArg3 = pal.GetCthis()
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, &flags, convArg3, &enabled, convArg5, &textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:94
// index:0
// Public virtual
// void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QStyle) DrawItemPixmap(painter unsafe.Pointer, rect *qtcore.QRect, alignment int, pixmap *qtgui.QPixmap) {
	var convArg1 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, &alignment, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:97
// index:0
// Public virtual
// QPalette standardPalette()
func (this *QStyle) StandardPalette() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle15standardPaletteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:204
// index:0
// Public pure virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawPrimitive(pe int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13drawPrimitiveENS_16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pe, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:275
// index:0
// Public pure virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawControl(element int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11drawControlENS_14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &element, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:364
// index:0
// Public pure virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QStyle) SubElementRect(subElement int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subElementRectENS_10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &subElement, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:443
// index:0
// Public pure virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QStyle) DrawComplexControl(cc int, opt unsafe.Pointer, p unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle18drawComplexControlENS_14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, opt, p, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyle.h:445
// index:0
// Public pure virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QStyle) HitTestComplexControl(cc int, opt unsafe.Pointer, pt *qtcore.QPoint, widget unsafe.Pointer) interface{} {
	var convArg2 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle21hitTestComplexControlENS_14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, opt, convArg2, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:447
// index:0
// Public pure virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14subControlRectENS_14ComplexControlEPK19QStyleOptionComplexNS_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, opt, &sc, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:582
// index:0
// Public pure virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QStyle) PixelMetric(metric int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle11pixelMetricENS_11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &metric, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:614
// index:0
// Public pure virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QStyle) SizeFromContents(ct int, opt unsafe.Pointer, contentsSize *qtcore.QSize, w unsafe.Pointer) interface{} {
	var convArg2 = contentsSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle16sizeFromContentsENS_12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ct, opt, convArg2, w)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:748
// index:0
// Public pure virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QStyle) StyleHint(stylehint int, opt unsafe.Pointer, widget unsafe.Pointer, returnData unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle9styleHintENS_9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stylehint, opt, widget, returnData)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:828
// index:0
// Public pure virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QStyle) StandardPixmap(standardPixmap int, opt unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle14standardPixmapENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &standardPixmap, opt, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:831
// index:0
// Public pure virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QStyle) StandardIcon(standardIcon int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle12standardIconENS_14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &standardIcon, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:834
// index:0
// Public pure virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt unsafe.Pointer) interface{} {
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &iconMode, convArg1, opt)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:837
// index:0
// Public static
// QRect visualRect(Qt::LayoutDirection, const class QRect &, const class QRect &)
func (this *QStyle) VisualRect(direction int, boundingRect *qtcore.QRect, logicalRect *qtcore.QRect) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle10visualRectEN2Qt15LayoutDirectionERK5QRectS4_", ffiqt.FFI_TYPE_POINTER, direction, boundingRect, logicalRect)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyle_VisualRect(direction int, boundingRect *qtcore.QRect, logicalRect *qtcore.QRect) {
	var nilthis *QStyle
	nilthis.VisualRect(direction, boundingRect, logicalRect)
}

// /usr/include/qt/QtWidgets/qstyle.h:839
// index:0
// Public static
// QPoint visualPos(Qt::LayoutDirection, const class QRect &, const class QPoint &)
func (this *QStyle) VisualPos(direction int, boundingRect *qtcore.QRect, logicalPos *qtcore.QPoint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle9visualPosEN2Qt15LayoutDirectionERK5QRectRK6QPoint", ffiqt.FFI_TYPE_POINTER, direction, boundingRect, logicalPos)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyle_VisualPos(direction int, boundingRect *qtcore.QRect, logicalPos *qtcore.QPoint) {
	var nilthis *QStyle
	nilthis.VisualPos(direction, boundingRect, logicalPos)
}

// /usr/include/qt/QtWidgets/qstyle.h:841
// index:0
// Public static
// int sliderPositionFromValue(int, int, int, int, _Bool)
func (this *QStyle) SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderPositionFromValueEiiiib", ffiqt.FFI_TYPE_POINTER, min, max, val, space, upsideDown)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyle_SliderPositionFromValue(min int, max int, val int, space int, upsideDown bool) {
	var nilthis *QStyle
	nilthis.SliderPositionFromValue(min, max, val, space, upsideDown)
}

// /usr/include/qt/QtWidgets/qstyle.h:843
// index:0
// Public static
// int sliderValueFromPosition(int, int, int, int, _Bool)
func (this *QStyle) SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QStyle23sliderValueFromPositionEiiiib", ffiqt.FFI_TYPE_POINTER, min, max, pos, space, upsideDown)
	gopp.ErrPrint(err, rv)
	return rv
}
func QStyle_SliderValueFromPosition(min int, max int, pos int, space int, upsideDown bool) {
	var nilthis *QStyle
	nilthis.SliderValueFromPosition(min, max, pos, space, upsideDown)
}

// /usr/include/qt/QtWidgets/qstyle.h:849
// index:0
// Public pure virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyle.h:856
// index:0
// Public
// const QStyle * proxy()
func (this *QStyle) Proxy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QStyle5proxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
