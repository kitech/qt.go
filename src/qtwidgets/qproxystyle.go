//  header block begin
// /usr/include/qt/QtWidgets/qproxystyle.h
// #include <qproxystyle.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QProxyStyle struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qproxystyle.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QProxyStyle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:57
// index:0
// void QProxyStyle(class QStyle *)
func NewQProxyStyle(style unsafe.Pointer) *QProxyStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleC2EP6QStyle", ffiqt.FFI_TYPE_VOID, cthis, style)
	gopp.ErrPrint(err, rv)
	return &QProxyStyle{cthis}
}

// /usr/include/qt/QtWidgets/qproxystyle.h:58
// index:1
// void QProxyStyle(const class QString &)
func NewQProxyStyle_1(key unsafe.Pointer) *QProxyStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, key)
	gopp.ErrPrint(err, rv)
	return &QProxyStyle{cthis}
}

// /usr/include/qt/QtWidgets/qproxystyle.h:59
// index:0
// virtual
// void ~QProxyStyle()
func DeleteQProxyStyle(*QProxyStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:61
// index:0
// QStyle * baseStyle()
func (this *QProxyStyle) BaseStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle9baseStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:62
// index:0
// void setBaseStyle(class QStyle *)
func (this *QProxyStyle) SetBaseStyle(style unsafe.Pointer) {
	// 0: (, QStyle * style), (style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle12setBaseStyleEP6QStyle", ffiqt.FFI_TYPE_VOID, this.cthis, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:64
// index:0
// virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawPrimitive(element int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::PrimitiveElement element, const QStyleOption * option, QPainter * painter, const QWidget * widget), (&element, option, painter, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &element, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:65
// index:0
// virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawControl(element int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::ControlElement element, const QStyleOption * option, QPainter * painter, const QWidget * widget), (&element, option, painter, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &element, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:66
// index:0
// virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QProxyStyle) DrawComplexControl(control int, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl control, const QStyleOptionComplex * option, QPainter * painter, const QWidget * widget), (&control, option, painter, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &control, option, painter, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:67
// index:0
// virtual
// void drawItemText(class QPainter *, const class QRect &, int, const class QPalette &, _Bool, const class QString &, class QPalette::ColorRole)
func (this *QProxyStyle) DrawItemText(painter unsafe.Pointer, rect unsafe.Pointer, flags int, pal unsafe.Pointer, enabled bool, text unsafe.Pointer, textRole int) {
	// 0: (, QPainter * painter, const QRect & rect, int flags, const QPalette & pal, bool enabled, const QString & text, QPalette::ColorRole textRole), (painter, rect, &flags, pal, &enabled, text, &textRole)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12drawItemTextEP8QPainterRK5QRectiRK8QPalettebRK7QStringNS5_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, painter, rect, &flags, pal, &enabled, text, &textRole)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:69
// index:0
// virtual
// void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) DrawItemPixmap(painter unsafe.Pointer, rect unsafe.Pointer, alignment int, pixmap unsafe.Pointer) {
	// 0: (, QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap), (painter, rect, &alignment, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, painter, rect, &alignment, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:71
// index:0
// virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QProxyStyle) SizeFromContents(type_ int, option unsafe.Pointer, size unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::ContentsType type, const QStyleOption * option, const QSize & size, const QWidget * widget), (&type_, option, size, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &type_, option, size, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:73
// index:0
// virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) SubElementRect(element int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::SubElement element, const QStyleOption * option, const QWidget * widget), (&element, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &element, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:74
// index:0
// virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QProxyStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, widget unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl cc, const QStyleOptionComplex * opt, QStyle::SubControl sc, const QWidget * widget), (&cc, opt, &sc, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &cc, opt, &sc, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:75
// index:0
// virtual
// QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QProxyStyle) ItemTextRect(fm unsafe.Pointer, r unsafe.Pointer, flags int, enabled bool, text unsafe.Pointer) {
	// 0: (, const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text), (fm, r, &flags, &enabled, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fm, r, &flags, &enabled, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:76
// index:0
// virtual
// QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) ItemPixmapRect(r unsafe.Pointer, flags int, pixmap unsafe.Pointer) {
	// 0: (, const QRect & r, int flags, const QPixmap & pixmap), (r, &flags, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, r, &flags, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:78
// index:0
// virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QProxyStyle) HitTestComplexControl(control int, option unsafe.Pointer, pos unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl control, const QStyleOptionComplex * option, const QPoint & pos, const QWidget * widget), (&control, option, pos, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &control, option, pos, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:79
// index:0
// virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QProxyStyle) StyleHint(hint int, option unsafe.Pointer, widget unsafe.Pointer, returnData unsafe.Pointer) {
	// 0: (, QStyle::StyleHint hint, const QStyleOption * option, const QWidget * widget, QStyleHintReturn * returnData), (&hint, option, widget, returnData)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_VOID, this.cthis, &hint, option, widget, returnData)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:80
// index:0
// virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) PixelMetric(metric int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::PixelMetric metric, const QStyleOption * option, const QWidget * widget), (&metric, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &metric, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:81
// index:0
// virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption * option, const QWidget * widget), (&control1, &control2, &orientation, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:84
// index:0
// virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) StandardIcon(standardIcon int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::StandardPixmap standardIcon, const QStyleOption * option, const QWidget * widget), (&standardIcon, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &standardIcon, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:85
// index:0
// virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QProxyStyle) StandardPixmap(standardPixmap int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::StandardPixmap standardPixmap, const QStyleOption * opt, const QWidget * widget), (&standardPixmap, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &standardPixmap, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:86
// index:0
// virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QProxyStyle) GeneratedIconPixmap(iconMode int, pixmap unsafe.Pointer, opt unsafe.Pointer) {
	// 0: (, QIcon::Mode iconMode, const QPixmap & pixmap, const QStyleOption * opt), (&iconMode, pixmap, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.cthis, &iconMode, pixmap, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:87
// index:0
// virtual
// QPalette standardPalette()
func (this *QProxyStyle) StandardPalette() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QProxyStyle15standardPaletteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:89
// index:0
// virtual
// void polish(class QWidget *)
func (this *QProxyStyle) Polish(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:90
// index:1
// virtual
// void polish(class QPalette &)
func (this *QProxyStyle) Polish_1(pal unsafe.Pointer) {
	// 1: (, QPalette & pal), (pal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishER8QPalette", ffiqt.FFI_TYPE_VOID, this.cthis, pal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:91
// index:2
// virtual
// void polish(class QApplication *)
func (this *QProxyStyle) Polish_2(app unsafe.Pointer) {
	// 2: (, QApplication * app), (app)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle6polishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.cthis, app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:93
// index:0
// virtual
// void unpolish(class QWidget *)
func (this *QProxyStyle) Unpolish(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qproxystyle.h:94
// index:1
// virtual
// void unpolish(class QApplication *)
func (this *QProxyStyle) Unpolish_1(app unsafe.Pointer) {
	// 1: (, QApplication * app), (app)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QProxyStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.cthis, app)
	gopp.ErrPrint(err, rv)
}

//  body block end
