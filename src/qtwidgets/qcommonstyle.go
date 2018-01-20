//  header block begin
// /usr/include/qt/QtWidgets/qcommonstyle.h
// #include <qcommonstyle.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
	*QStyle
}

func (this *QCommonStyle) GetCthis() unsafe.Pointer {
	return this.QStyle.GetCthis()
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCommonStyle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:55
// index:0
// void QCommonStyle()
func NewQCommonStyle() *QCommonStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommonStyleFromPointer(cthis)
	return gothis
}
func NewQCommonStyleFromPointer(cthis unsafe.Pointer) *QCommonStyle {
	bcthis0 := NewQStyleFromPointer(cthis)
	return &QCommonStyle{bcthis0}
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:95
// index:1
// void QCommonStyle(class QCommonStylePrivate &)
func NewQCommonStyle_1(dd unsafe.Pointer) *QCommonStyle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleC2ER19QCommonStylePrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommonStyleFromPointer(cthis)
	return gothis
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
	// 0: (, pe QStyle::PrimitiveElement, opt const QStyleOption *, p QPainter *, w const QWidget *), (&pe, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pe, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawControl(element int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, element QStyle::ControlElement, opt const QStyleOption *, p QPainter *, w const QWidget *), (&element, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &element, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) SubElementRect(r int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, r QStyle::SubElement, opt const QStyleOption *, widget const QWidget *), (&r, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &r, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawComplexControl(cc int, opt unsafe.Pointer, p unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, p QPainter *, w const QWidget *), (&cc, opt, p, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, p, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QCommonStyle) HitTestComplexControl(cc int, opt unsafe.Pointer, pt unsafe.Pointer, w unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, pt const QPoint &, w const QWidget *), (&cc, opt, pt, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, pt, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QCommonStyle) SubControlRect(cc int, opt unsafe.Pointer, sc int, w unsafe.Pointer) {
	// 0: (, cc QStyle::ComplexControl, opt const QStyleOptionComplex *, sc QStyle::SubControl, w const QWidget *), (&cc, opt, &sc, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cc, opt, &sc, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QCommonStyle) SizeFromContents(ct int, opt unsafe.Pointer, contentsSize unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, ct QStyle::ContentsType, opt const QStyleOption *, contentsSize const QSize &, widget const QWidget *), (&ct, opt, contentsSize, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ct, opt, contentsSize, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) PixelMetric(m int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, m QStyle::PixelMetric, opt const QStyleOption *, widget const QWidget *), (&m, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &m, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QCommonStyle) StyleHint(sh int, opt unsafe.Pointer, w unsafe.Pointer, shret unsafe.Pointer) {
	// 0: (, sh QStyle::StyleHint, opt const QStyleOption *, w const QWidget *, shret QStyleHintReturn *), (&sh, opt, w, shret)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sh, opt, w, shret)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardIcon(standardIcon int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, standardIcon QStyle::StandardPixmap, opt const QStyleOption *, widget const QWidget *), (&standardIcon, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &standardIcon, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardPixmap(sp int, opt unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, sp QStyle::StandardPixmap, opt const QStyleOption *, widget const QWidget *), (&sp, opt, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sp, opt, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:82
// index:0
// virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QCommonStyle) GeneratedIconPixmap(iconMode int, pixmap unsafe.Pointer, opt unsafe.Pointer) {
	// 0: (, iconMode QIcon::Mode, pixmap const QPixmap &, opt const QStyleOption *), (&iconMode, pixmap, opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &iconMode, pixmap, opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) LayoutSpacing(control1 int, control2 int, orientation int, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, control1 QSizePolicy::ControlType, control2 QSizePolicy::ControlType, orientation Qt::Orientation, option const QStyleOption *, widget const QWidget *), (&control1, &control2, &orientation, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &control1, &control2, &orientation, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:88
// index:0
// virtual
// void polish(class QPalette &)
func (this *QCommonStyle) Polish(arg0 unsafe.Pointer) {
	// 0: (, QPalette & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishER8QPalette", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:89
// index:1
// virtual
// void polish(class QApplication *)
func (this *QCommonStyle) Polish_1(app unsafe.Pointer) {
	// 1: (, app QApplication *), (app)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.GetCthis(), app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:90
// index:2
// virtual
// void polish(class QWidget *)
func (this *QCommonStyle) Polish_2(widget unsafe.Pointer) {
	// 2: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:91
// index:0
// virtual
// void unpolish(class QWidget *)
func (this *QCommonStyle) Unpolish(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:92
// index:1
// virtual
// void unpolish(class QApplication *)
func (this *QCommonStyle) Unpolish_1(application unsafe.Pointer) {
	// 1: (, application QApplication *), (application)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_VOID, this.GetCthis(), application)
	gopp.ErrPrint(err, rv)
}

//  body block end
