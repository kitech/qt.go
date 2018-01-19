//  header block begin
// /usr/include/qt/QtWidgets/qcommonstyle.h
// #include <qcommonstyle.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QCommonStyle struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCommonStyle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:55
// index:0
// void QCommonStyle()
func NewQCommonStyle() *QCommonStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QCommonStyle{cthis}
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:56
// index:0
// virtual
// void ~QCommonStyle()
func DeleteQCommonStyle(*QCommonStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:58
// index:0
// virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawPrimitive(pe int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, QStyle::PrimitiveElement pe, const QStyleOption * opt, QPainter * p, const QWidget * w), (&pe, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &pe, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawControl(element int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, QStyle::ControlElement element, const QStyleOption * opt, QPainter * p, const QWidget * w), (&element, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &element, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) SubElementRect(r int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::SubElement r, const QStyleOption * opt, const QWidget * widget), (&r, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &r, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawComplexControl(cc int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl cc, const QStyleOptionComplex * opt, QPainter * p, const QWidget * w), (&cc, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &cc, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QCommonStyle) HitTestComplexControl(cc int, opt unsafe.Pointer, pt unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl cc, const QStyleOptionComplex * opt, const QPoint & pt, const QWidget * w), (&cc, opt, pt, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &cc, opt, pt, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QCommonStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, w unsafe.Pointer) {
	// 0: (, QStyle::ComplexControl cc, const QStyleOptionComplex * opt, QStyle::SubControl sc, const QWidget * w), (&cc, opt, &sc, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &cc, opt, &sc, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QCommonStyle) SizeFromContents(ct int, opt unsafe.Pointer, contentsSize unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::ContentsType ct, const QStyleOption * opt, const QSize & contentsSize, const QWidget * widget), (&ct, opt, contentsSize, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &ct, opt, contentsSize, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) PixelMetric(m int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::PixelMetric m, const QStyleOption * opt, const QWidget * widget), (&m, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &m, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QCommonStyle) StyleHint(sh int, opt unsafe.Pointer, w unsafe.Pointer, shret unsafe.Pointer) {
	// 0: (, QStyle::StyleHint sh, const QStyleOption * opt, const QWidget * w, QStyleHintReturn * shret), (&sh, opt, w, shret)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_VOID, this.cthis, &sh, opt, w, shret)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardIcon(standardIcon int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::StandardPixmap standardIcon, const QStyleOption * opt, const QWidget * widget), (&standardIcon, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &standardIcon, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardPixmap(sp int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QStyle::StandardPixmap sp, const QStyleOption * opt, const QWidget * widget), (&sp, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &sp, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:82
// index:0
// virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QCommonStyle) GeneratedIconPixmap(iconMode int, pixmap unsafe.Pointer, opt unsafe.Pointer) {
	// 0: (, QIcon::Mode iconMode, const QPixmap & pixmap, const QStyleOption * opt), (&iconMode, pixmap, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.cthis, &iconMode, pixmap, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption * option, const QWidget * widget), (&control1, &control2, &orientation, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:88
// index:0
// virtual
// void polish(class QPalette &)
func (this *QCommonStyle) Polish(arg0 unsafe.Pointer) {
	// 0: (, QPalette & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishER8QPalette", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:89
// index:1
// virtual
// void polish(class QApplication *)
func (this *QCommonStyle) Polish_1(app unsafe.Pointer) {
	// 1: (, QApplication * app), (app)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.cthis, app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:90
// index:2
// virtual
// void polish(class QWidget *)
func (this *QCommonStyle) Polish_2(widget unsafe.Pointer) {
	// 2: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:91
// index:0
// virtual
// void unpolish(class QWidget *)
func (this *QCommonStyle) Unpolish(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:92
// index:1
// virtual
// void unpolish(class QApplication *)
func (this *QCommonStyle) Unpolish_1(application unsafe.Pointer) {
	// 1: (, QApplication * application), (application)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.cthis, application)
	gopp.ErrPrint(err, rv)
}

//  body block end
