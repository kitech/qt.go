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
func (this *QCommonStyle) SetCthis(cthis unsafe.Pointer) {
	this.QStyle = NewQStyleFromPointer(cthis)
}
func NewQCommonStyleFromPointer(cthis unsafe.Pointer) *QCommonStyle {
	bcthis0 := NewQStyleFromPointer(cthis)
	return &QCommonStyle{bcthis0}
}
func (*QCommonStyle) NewFromPointer(cthis unsafe.Pointer) *QCommonStyle {
	return NewQCommonStyleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QCommonStyle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommonStyle()
func NewQCommonStyle() *QCommonStyle {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommonStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCommonStyle()
func DeleteQCommonStyle(this *QCommonStyle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyleD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(enum QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QCommonStyle) DrawPrimitive(pe int, opt *QStyleOption /*777 const QStyleOption **/, p *qtgui.QPainter /*777 QPainter **/, w *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(enum QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *)
func (this *QCommonStyle) DrawControl(element int, opt *QStyleOption /*777 const QStyleOption **/, p *qtgui.QPainter /*777 QPainter **/, w *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subElementRect(enum QStyle::SubElement, const QStyleOption *, const QWidget *)
func (this *QCommonStyle) SubElementRect(r int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *)
func (this *QCommonStyle) DrawComplexControl(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, p *qtgui.QPainter /*777 QPainter **/, w *QWidget /*777 const QWidget **/) {
	var convArg1 = opt.GetCthis()
	var convArg2 = p.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(enum QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *)
func (this *QCommonStyle) HitTestComplexControl(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, pt *qtcore.QPoint, w *QWidget /*777 const QWidget **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = pt.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subControlRect(enum QStyle::ComplexControl, const QStyleOptionComplex *, enum QStyle::SubControl, const QWidget *)
func (this *QCommonStyle) SubControlRect(cc int, opt *QStyleOptionComplex /*777 const QStyleOptionComplex **/, sc int, w *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg3 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, sc, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeFromContents(enum QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *)
func (this *QCommonStyle) SizeFromContents(ct int, opt *QStyleOption /*777 const QStyleOption **/, contentsSize *qtcore.QSize, widget *QWidget /*777 const QWidget **/) *qtcore.QSize /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = contentsSize.GetCthis()
	var convArg3 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ct, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(enum QStyle::PixelMetric, const QStyleOption *, const QWidget *)
func (this *QCommonStyle) PixelMetric(m int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), m, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(enum QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *)
func (this *QCommonStyle) StyleHint(sh int, opt *QStyleOption /*777 const QStyleOption **/, w *QWidget /*777 const QWidget **/, shret *QStyleHintReturn /*777 QStyleHintReturn **/) int {
	var convArg1 = opt.GetCthis()
	var convArg2 = w.GetCthis()
	var convArg3 = shret.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sh, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QCommonStyle) StandardIcon(standardIcon int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QIcon /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(enum QStyle::StandardPixmap, const QStyleOption *, const QWidget *)
func (this *QCommonStyle) StandardPixmap(sp int, opt *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = opt.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sp, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap generatedIconPixmap(QIcon::Mode, const QPixmap &, const QStyleOption *)
func (this *QCommonStyle) GeneratedIconPixmap(iconMode int, pixmap *qtgui.QPixmap, opt *QStyleOption /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 = pixmap.GetCthis()
	var convArg2 = opt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), iconMode, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *)
func (this *QCommonStyle) LayoutSpacing(control1 int, control2 int, orientation int, option *QStyleOption /*777 const QStyleOption **/, widget *QWidget /*777 const QWidget **/) int {
	var convArg3 = option.GetCthis()
	var convArg4 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QPalette &)
func (this *QCommonStyle) Polish(arg0 *qtgui.QPalette) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishER8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:89
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QApplication *)
func (this *QCommonStyle) Polish_1(app *QApplication /*777 QApplication **/) {
	var convArg0 = app.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:90
// index:2
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QWidget *)
func (this *QCommonStyle) Polish_2(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QWidget *)
func (this *QCommonStyle) Unpolish(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:92
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QApplication *)
func (this *QCommonStyle) Unpolish_1(application *QApplication /*777 QApplication **/) {
	var convArg0 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP12QApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
