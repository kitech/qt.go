package qtwidgets

// /usr/include/qt/QtWidgets/qcommonstyle.h
// #include <qcommonstyle.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.QStyle.GetCthis()
	}
}
func NewQCommonStyleFromPointer(cthis unsafe.Pointer) *QCommonStyle {
	bcthis0 := NewQStyleFromPointer(cthis)
	return &QCommonStyle{bcthis0}
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:52
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QCommonStyle) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:55
// index:0
// Public
// void QCommonStyle()
func NewQCommonStyle() *QCommonStyle {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommonStyleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:56
// index:0
// Public virtual
// void ~QCommonStyle()
func DeleteQCommonStyle(*QCommonStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:58
// index:0
// Public virtual
// void drawPrimitive(enum QStyle::PrimitiveElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawPrimitive(pe int, opt *QStyleOption /*444 const QStyleOption **/, p *qtgui.QPainter /*444 QPainter **/, w *QWidget /*444 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pe, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// Public virtual
// void drawControl(enum QStyle::ControlElement, const class QStyleOption *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawControl(element int, opt *QStyleOption /*444 const QStyleOption **/, p *qtgui.QPainter /*444 QPainter **/, w *QWidget /*444 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &element, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// Public virtual
// QRect subElementRect(enum QStyle::SubElement, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) SubElementRect(r int, opt *QStyleOption /*444 const QStyleOption **/, widget *QWidget /*444 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &r, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// Public virtual
// void drawComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, class QPainter *, const class QWidget *)
func (this *QCommonStyle) DrawComplexControl(cc int, opt *QStyleOptionComplex /*444 const QStyleOptionComplex **/, p *qtgui.QPainter /*444 QPainter **/, w *QWidget /*444 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// Public virtual
// QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const class QStyleOptionComplex *, const class QPoint &, const class QWidget *)
func (this *QCommonStyle) HitTestComplexControl(cc int, opt *QStyleOptionComplex /*444 const QStyleOptionComplex **/, pt *qtcore.QPoint, w *QWidget /*444 const QWidget **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = pt.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// Public virtual
// QRect subControlRect(enum QStyle::ComplexControl, const class QStyleOptionComplex *, enum QStyle::SubControl, const class QWidget *)
func (this *QCommonStyle) SubControlRect(cc int, opt *QStyleOptionComplex /*444 const QStyleOptionComplex **/, sc int, w *QWidget /*444 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cc, convArg1, &sc, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// Public virtual
// QSize sizeFromContents(enum QStyle::ContentsType, const class QStyleOption *, const class QSize &, const class QWidget *)
func (this *QCommonStyle) SizeFromContents(ct int, opt *QStyleOption /*444 const QStyleOption **/, contentsSize *qtcore.QSize, widget *QWidget /*444 const QWidget **/) *qtcore.QSize /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = contentsSize.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ct, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// Public virtual
// int pixelMetric(enum QStyle::PixelMetric, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) PixelMetric(m int, opt *QStyleOption /*444 const QStyleOption **/, widget *QWidget /*444 const QWidget **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &m, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual
// int styleHint(enum QStyle::StyleHint, const class QStyleOption *, const class QWidget *, class QStyleHintReturn *)
func (this *QCommonStyle) StyleHint(sh int, opt *QStyleOption /*444 const QStyleOption **/, w *QWidget /*444 const QWidget **/, shret *QStyleHintReturn /*444 QStyleHintReturn **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = w.GetCthis()
	var convArg3 = shret.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sh, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// Public virtual
// QIcon standardIcon(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardIcon(standardIcon int, opt *QStyleOption /*444 const QStyleOption **/, widget *QWidget /*444 const QWidget **/) *qtgui.QIcon /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &standardIcon, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// Public virtual
// QPixmap standardPixmap(enum QStyle::StandardPixmap, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) StandardPixmap(sp int, opt *QStyleOption /*444 const QStyleOption **/, widget *QWidget /*444 const QWidget **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sp, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:82
// index:0
// Public virtual
// QPixmap generatedIconPixmap(class QIcon::Mode, const class QPixmap &, const class QStyleOption *)
func (this *QCommonStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt *QStyleOption /*444 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = pixmap.GetCthis()
	var convArg2 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &iconMode, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// Public virtual
// int layoutSpacing(class QSizePolicy::ControlType, class QSizePolicy::ControlType, Qt::Orientation, const class QStyleOption *, const class QWidget *)
func (this *QCommonStyle) LayoutSpacing(control1 int, control2 int, orientation int, option *QStyleOption /*444 const QStyleOption **/, widget *QWidget /*444 const QWidget **/) int {
	var convArg3 = option.GetCthis()
	var convArg4 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &control1, &control2, &orientation, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:88
// index:0
// Public virtual
// void polish(class QPalette &)
func (this *QCommonStyle) Polish(arg0 *qtgui.QPalette) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishER8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:89
// index:1
// Public virtual
// void polish(class QApplication *)
func (this *QCommonStyle) Polish_1(app *QApplication /*444 QApplication **/) {
	var convArg0 = app.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:90
// index:2
// Public virtual
// void polish(class QWidget *)
func (this *QCommonStyle) Polish_2(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:91
// index:0
// Public virtual
// void unpolish(class QWidget *)
func (this *QCommonStyle) Unpolish(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:92
// index:1
// Public virtual
// void unpolish(class QApplication *)
func (this *QCommonStyle) Unpolish_1(application *QApplication /*444 QApplication **/) {
	var convArg0 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
