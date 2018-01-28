package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h
// #include <qgraphicsitemanimation.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QGraphicsItemAnimation struct {
	*qtcore.QObject
}

func (this *QGraphicsItemAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsItemAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsItemAnimationFromPointer(cthis unsafe.Pointer) *QGraphicsItemAnimation {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsItemAnimation{bcthis0}
}
func (*QGraphicsItemAnimation) NewFromPointer(cthis unsafe.Pointer) *QGraphicsItemAnimation {
	return NewQGraphicsItemAnimationFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsItemAnimation) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:61
// index:0
// Public
// void QGraphicsItemAnimation(QObject *)
func NewQGraphicsItemAnimation(parent *qtcore.QObject /*777 QObject **/) *QGraphicsItemAnimation {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:62
// index:0
// Public virtual
// void ~QGraphicsItemAnimation()
func DeleteQGraphicsItemAnimation(*QGraphicsItemAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:64
// index:0
// Public
// QGraphicsItem * item()
func (this *QGraphicsItemAnimation) Item() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation4itemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:65
// index:0
// Public
// void setItem(QGraphicsItem *)
func (this *QGraphicsItemAnimation) SetItem(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:67
// index:0
// Public
// QTimeLine * timeLine()
func (this *QGraphicsItemAnimation) TimeLine() *qtcore.QTimeLine /*777 QTimeLine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation8timeLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQTimeLineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:68
// index:0
// Public
// void setTimeLine(QTimeLine *)
func (this *QGraphicsItemAnimation) SetTimeLine(timeLine *qtcore.QTimeLine /*777 QTimeLine **/) {
	var convArg0 = timeLine.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:70
// index:0
// Public
// QPointF posAt(qreal)
func (this *QGraphicsItemAnimation) PosAt(step float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation5posAtEd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:72
// index:0
// Public
// void setPosAt(qreal, const QPointF &)
func (this *QGraphicsItemAnimation) SetPosAt(step float64, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:74
// index:0
// Public
// QMatrix matrixAt(qreal)
func (this *QGraphicsItemAnimation) MatrixAt(step float64) *qtgui.QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation8matrixAtEd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:76
// index:0
// Public
// qreal rotationAt(qreal)
func (this *QGraphicsItemAnimation) RotationAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation10rotationAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:78
// index:0
// Public
// void setRotationAt(qreal, qreal)
func (this *QGraphicsItemAnimation) SetRotationAt(step float64, angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation13setRotationAtEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:80
// index:0
// Public
// qreal xTranslationAt(qreal)
func (this *QGraphicsItemAnimation) XTranslationAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation14xTranslationAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:81
// index:0
// Public
// qreal yTranslationAt(qreal)
func (this *QGraphicsItemAnimation) YTranslationAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation14yTranslationAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:83
// index:0
// Public
// void setTranslationAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetTranslationAt(step float64, dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation16setTranslationAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:85
// index:0
// Public
// qreal verticalScaleAt(qreal)
func (this *QGraphicsItemAnimation) VerticalScaleAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation15verticalScaleAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:86
// index:0
// Public
// qreal horizontalScaleAt(qreal)
func (this *QGraphicsItemAnimation) HorizontalScaleAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation17horizontalScaleAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:88
// index:0
// Public
// void setScaleAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetScaleAt(step float64, sx float64, sy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation10setScaleAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, sx, sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:90
// index:0
// Public
// qreal verticalShearAt(qreal)
func (this *QGraphicsItemAnimation) VerticalShearAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation15verticalShearAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:91
// index:0
// Public
// qreal horizontalShearAt(qreal)
func (this *QGraphicsItemAnimation) HorizontalShearAt(step float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation17horizontalShearAtEd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:93
// index:0
// Public
// void setShearAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetShearAt(step float64, sh float64, sv float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation10setShearAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step, sh, sv)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:95
// index:0
// Public
// void clear()
func (this *QGraphicsItemAnimation) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:98
// index:0
// Public
// void setStep(qreal)
func (this *QGraphicsItemAnimation) SetStep(x float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation7setStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:99
// index:0
// Public
// void reset()
func (this *QGraphicsItemAnimation) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:102
// index:0
// Protected virtual
// void beforeAnimationStep(qreal)
func (this *QGraphicsItemAnimation) BeforeAnimationStep(step float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation19beforeAnimationStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:103
// index:0
// Protected virtual
// void afterAnimationStep(qreal)
func (this *QGraphicsItemAnimation) AfterAnimationStep(step float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation18afterAnimationStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), step)
	gopp.ErrPrint(err, rv)
}

//  body block end
