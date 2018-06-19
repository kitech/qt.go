package qtwidgets

// /usr/include/qt/QtWidgets/qcommonstyle.h
// #include <qcommonstyle.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QCommonStyle struct {
	*QStyle
}
type QCommonStyle_ITF interface {
	QStyle_ITF
	QCommonStyle_PTR() *QCommonStyle
}

func (ptr *QCommonStyle) QCommonStyle_PTR() *QCommonStyle { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCommonStyle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommonStyle()

/*
Constructs a QCommonStyle.
*/
func NewQCommonStyle() *QCommonStyle {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyleC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommonStyleFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCommonStyle")
	return gothis
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCommonStyle()

/*

 */
func DeleteQCommonStyle(this *QCommonStyle) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawPrimitive().
*/
func (this *QCommonStyle) DrawPrimitive(pe int, opt QStyleOption_ITF /*777 const QStyleOption **/, p qtgui.QPainter_ITF /*777 QPainter **/, w QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg3 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawPrimitive(QStyle::PrimitiveElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawPrimitive().
*/
func (this *QCommonStyle) DrawPrimitive__(pe int, opt QStyleOption_ITF /*777 const QStyleOption **/, p qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle13drawPrimitiveEN6QStyle16PrimitiveElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pe, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawControl().
*/
func (this *QCommonStyle) DrawControl(element int, opt QStyleOption_ITF /*777 const QStyleOption **/, p qtgui.QPainter_ITF /*777 QPainter **/, w QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg3 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawControl(QStyle::ControlElement, const QStyleOption *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawControl().
*/
func (this *QCommonStyle) DrawControl__(element int, opt QStyleOption_ITF /*777 const QStyleOption **/, p qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle11drawControlEN6QStyle14ControlElementEPK12QStyleOptionP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), element, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subElementRect(QStyle::SubElement, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::subElementRect().
*/
func (this *QCommonStyle) SubElementRect(r int, opt QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), r, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subElementRect(QStyle::SubElement, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::subElementRect().
*/
func (this *QCommonStyle) SubElementRect__(r int, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14subElementRectEN6QStyle10SubElementEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), r, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawComplexControl().
*/
func (this *QCommonStyle) DrawComplexControl(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, p qtgui.QPainter_ITF /*777 QPainter **/, w QWidget_ITF /*777 const QWidget **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg3 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void drawComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, QPainter *, const QWidget *) const

/*
Reimplemented from QStyle::drawComplexControl().
*/
func (this *QCommonStyle) DrawComplexControl__(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, p qtgui.QPainter_ITF /*777 QPainter **/) {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg2 = p.QPainter_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle18drawComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexP8QPainterPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *) const

/*
Reimplemented from QStyle::hitTestComplexControl().
*/
func (this *QCommonStyle) HitTestComplexControl(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, pt qtcore.QPoint_ITF, w QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg2 = pt.QPoint_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg3 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl, const QStyleOptionComplex *, const QPoint &, const QWidget *) const

/*
Reimplemented from QStyle::hitTestComplexControl().
*/
func (this *QCommonStyle) HitTestComplexControl__(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, pt qtcore.QPoint_ITF) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg2 = pt.QPoint_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle21hitTestComplexControlEN6QStyle14ComplexControlEPK19QStyleOptionComplexRK6QPointPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subControlRect(QStyle::ComplexControl, const QStyleOptionComplex *, QStyle::SubControl, const QWidget *) const

/*
Reimplemented from QStyle::subControlRect().
*/
func (this *QCommonStyle) SubControlRect(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, sc int, w QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg3 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, sc, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect subControlRect(QStyle::ComplexControl, const QStyleOptionComplex *, QStyle::SubControl, const QWidget *) const

/*
Reimplemented from QStyle::subControlRect().
*/
func (this *QCommonStyle) SubControlRect__(cc int, opt QStyleOptionComplex_ITF /*777 const QStyleOptionComplex **/, sc int) *qtcore.QRect /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOptionComplex_PTR() != nil {
		convArg1 = opt.QStyleOptionComplex_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14subControlRectEN6QStyle14ComplexControlEPK19QStyleOptionComplexNS0_10SubControlEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cc, convArg1, sc, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeFromContents(QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *) const

/*
Reimplemented from QStyle::sizeFromContents().
*/
func (this *QCommonStyle) SizeFromContents(ct int, opt QStyleOption_ITF /*777 const QStyleOption **/, contentsSize qtcore.QSize_ITF, widget QWidget_ITF /*777 const QWidget **/) *qtcore.QSize /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if contentsSize != nil && contentsSize.QSize_PTR() != nil {
		convArg2 = contentsSize.QSize_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg3 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ct, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeFromContents(QStyle::ContentsType, const QStyleOption *, const QSize &, const QWidget *) const

/*
Reimplemented from QStyle::sizeFromContents().
*/
func (this *QCommonStyle) SizeFromContents__(ct int, opt QStyleOption_ITF /*777 const QStyleOption **/, contentsSize qtcore.QSize_ITF) *qtcore.QSize /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if contentsSize != nil && contentsSize.QSize_PTR() != nil {
		convArg2 = contentsSize.QSize_PTR().GetCthis()
	}
	// arg: 3, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle16sizeFromContentsEN6QStyle12ContentsTypeEPK12QStyleOptionRK5QSizePK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ct, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QCommonStyle) PixelMetric(m int, opt QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QCommonStyle) PixelMetric__(m int) int {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int pixelMetric(QStyle::PixelMetric, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::pixelMetric().
*/
func (this *QCommonStyle) PixelMetric__1(m int, opt QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle11pixelMetricEN6QStyle11PixelMetricEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QCommonStyle) StyleHint(sh int, opt QStyleOption_ITF /*777 const QStyleOption **/, w QWidget_ITF /*777 const QWidget **/, shret QStyleHintReturn_ITF /*777 QStyleHintReturn **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg2 = w.QWidget_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if shret != nil && shret.QStyleHintReturn_PTR() != nil {
		convArg3 = shret.QStyleHintReturn_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QCommonStyle) StyleHint__(sh int) int {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QCommonStyle) StyleHint__1(sh int, opt QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int styleHint(QStyle::StyleHint, const QStyleOption *, const QWidget *, QStyleHintReturn *) const

/*
Reimplemented from QStyle::styleHint().
*/
func (this *QCommonStyle) StyleHint__2(sh int, opt QStyleOption_ITF /*777 const QStyleOption **/, w QWidget_ITF /*777 const QWidget **/) int {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg2 = w.QWidget_PTR().GetCthis()
	}
	// arg: 3, QStyleHintReturn *=Pointer, QStyleHintReturn=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle9styleHintEN6QStyle9StyleHintEPK12QStyleOptionPK7QWidgetP16QStyleHintReturn", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sh, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*

 */
func (this *QCommonStyle) StandardIcon(standardIcon int, opt QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtgui.QIcon /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*

 */
func (this *QCommonStyle) StandardIcon__(standardIcon int) *qtgui.QIcon /*123*/ {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon standardIcon(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*

 */
func (this *QCommonStyle) StandardIcon__1(standardIcon int, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QIcon /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle12standardIconEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardIcon, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardPixmap().
*/
func (this *QCommonStyle) StandardPixmap(sp int, opt QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sp, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardPixmap().
*/
func (this *QCommonStyle) StandardPixmap__(sp int) *qtgui.QPixmap /*123*/ {
	// arg: 1, const QStyleOption *=Pointer, QStyleOption=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sp, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap standardPixmap(QStyle::StandardPixmap, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::standardPixmap().
*/
func (this *QCommonStyle) StandardPixmap__1(sp int, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg1 = opt.QStyleOption_PTR().GetCthis()
	}
	// arg: 2, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle14standardPixmapEN6QStyle14StandardPixmapEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sp, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap generatedIconPixmap(QIcon::Mode, const QPixmap &, const QStyleOption *) const

/*
Reimplemented from QStyle::generatedIconPixmap().
*/
func (this *QCommonStyle) GeneratedIconPixmap(iconMode int, pixmap qtgui.QPixmap_ITF, opt QStyleOption_ITF /*777 const QStyleOption **/) *qtgui.QPixmap /*123*/ {
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if opt != nil && opt.QStyleOption_PTR() != nil {
		convArg2 = opt.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle19generatedIconPixmapEN5QIcon4ModeERK7QPixmapPK12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), iconMode, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().
*/
func (this *QCommonStyle) LayoutSpacing(control1 int, control2 int, orientation int, option QStyleOption_ITF /*777 const QStyleOption **/, widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg3 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg3 = option.QStyleOption_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg4 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().
*/
func (this *QCommonStyle) LayoutSpacing__(control1 int, control2 int, orientation int) int {
	// arg: 3, const QStyleOption *=Pointer, QStyleOption=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int layoutSpacing(QSizePolicy::ControlType, QSizePolicy::ControlType, Qt::Orientation, const QStyleOption *, const QWidget *) const

/*
Reimplemented from QStyle::layoutSpacing().
*/
func (this *QCommonStyle) LayoutSpacing__1(control1 int, control2 int, orientation int, option QStyleOption_ITF /*777 const QStyleOption **/) int {
	var convArg3 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg3 = option.QStyleOption_PTR().GetCthis()
	}
	// arg: 4, const QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QCommonStyle13layoutSpacingEN11QSizePolicy11ControlTypeES1_N2Qt11OrientationEPK12QStyleOptionPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), control1, control2, orientation, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QPalette &)

/*
Reimplemented from QStyle::polish().
*/
func (this *QCommonStyle) Polish(arg0 qtgui.QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPalette_PTR() != nil {
		convArg0 = arg0.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyle6polishER8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:89
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QApplication *)

/*
Reimplemented from QStyle::polish().
*/
func (this *QCommonStyle) Polish_1(app QApplication_ITF /*777 QApplication **/) {
	var convArg0 unsafe.Pointer
	if app != nil && app.QApplication_PTR() != nil {
		convArg0 = app.QApplication_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:90
// index:2
// Public virtual Visibility=Default Availability=Available
// [-2] void polish(QWidget *)

/*
Reimplemented from QStyle::polish().
*/
func (this *QCommonStyle) Polish_2(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyle6polishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QWidget *)

/*
Reimplemented from QStyle::unpolish().
*/
func (this *QCommonStyle) Unpolish(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommonstyle.h:92
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void unpolish(QApplication *)

/*
Reimplemented from QStyle::unpolish().
*/
func (this *QCommonStyle) Unpolish_1(application QApplication_ITF /*777 QApplication **/) {
	var convArg0 unsafe.Pointer
	if application != nil && application.QApplication_PTR() != nil {
		convArg0 = application.QApplication_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QCommonStyle8unpolishEP12QApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
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

//  keep block end
