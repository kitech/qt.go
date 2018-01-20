//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h
// #include <qgraphicsitemanimation.h>
// #include <QtWidgets>
package qtwidgets

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
type QGraphicsItemAnimation struct {
	*qtcore.QObject
}

func (this *QGraphicsItemAnimation) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQGraphicsItemAnimationFromPointer(cthis unsafe.Pointer) *QGraphicsItemAnimation {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsItemAnimation{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsItemAnimation) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:61
// index:0
// Public
// void QGraphicsItemAnimation(class QObject *)
func NewQGraphicsItemAnimation(parent unsafe.Pointer) *QGraphicsItemAnimation {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
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
func (this *QGraphicsItemAnimation) Item() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation4itemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:65
// index:0
// Public
// void setItem(class QGraphicsItem *)
func (this *QGraphicsItemAnimation) SetItem(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:67
// index:0
// Public
// QTimeLine * timeLine()
func (this *QGraphicsItemAnimation) TimeLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation8timeLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:68
// index:0
// Public
// void setTimeLine(class QTimeLine *)
func (this *QGraphicsItemAnimation) SetTimeLine(timeLine unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeLine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:70
// index:0
// Public
// QPointF posAt(qreal)
func (this *QGraphicsItemAnimation) PosAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation5posAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:71
// index:0
// Public
// QList<QPair<qreal, QPointF> > posList()
func (this *QGraphicsItemAnimation) PosList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation7posListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:72
// index:0
// Public
// void setPosAt(qreal, const class QPointF &)
func (this *QGraphicsItemAnimation) SetPosAt(step float64, pos *qtcore.QPointF) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:74
// index:0
// Public
// QMatrix matrixAt(qreal)
func (this *QGraphicsItemAnimation) MatrixAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation8matrixAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:76
// index:0
// Public
// qreal rotationAt(qreal)
func (this *QGraphicsItemAnimation) RotationAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation10rotationAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:77
// index:0
// Public
// QList<QPair<qreal, qreal> > rotationList()
func (this *QGraphicsItemAnimation) RotationList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation12rotationListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:78
// index:0
// Public
// void setRotationAt(qreal, qreal)
func (this *QGraphicsItemAnimation) SetRotationAt(step float64, angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation13setRotationAtEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step, &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:80
// index:0
// Public
// qreal xTranslationAt(qreal)
func (this *QGraphicsItemAnimation) XTranslationAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation14xTranslationAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:81
// index:0
// Public
// qreal yTranslationAt(qreal)
func (this *QGraphicsItemAnimation) YTranslationAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation14yTranslationAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:82
// index:0
// Public
// QList<QPair<qreal, QPointF> > translationList()
func (this *QGraphicsItemAnimation) TranslationList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation15translationListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:83
// index:0
// Public
// void setTranslationAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetTranslationAt(step float64, dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation16setTranslationAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:85
// index:0
// Public
// qreal verticalScaleAt(qreal)
func (this *QGraphicsItemAnimation) VerticalScaleAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation15verticalScaleAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:86
// index:0
// Public
// qreal horizontalScaleAt(qreal)
func (this *QGraphicsItemAnimation) HorizontalScaleAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation17horizontalScaleAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:87
// index:0
// Public
// QList<QPair<qreal, QPointF> > scaleList()
func (this *QGraphicsItemAnimation) ScaleList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation9scaleListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:88
// index:0
// Public
// void setScaleAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetScaleAt(step float64, sx float64, sy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation10setScaleAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step, &sx, &sy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:90
// index:0
// Public
// qreal verticalShearAt(qreal)
func (this *QGraphicsItemAnimation) VerticalShearAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation15verticalShearAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:91
// index:0
// Public
// qreal horizontalShearAt(qreal)
func (this *QGraphicsItemAnimation) HorizontalShearAt(step float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation17horizontalShearAtEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:92
// index:0
// Public
// QList<QPair<qreal, QPointF> > shearList()
func (this *QGraphicsItemAnimation) ShearList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QGraphicsItemAnimation9shearListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:93
// index:0
// Public
// void setShearAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) SetShearAt(step float64, sh float64, sv float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation10setShearAtEddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step, &sh, &sv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation7setStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation19beforeAnimationStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitemanimation.h:103
// index:0
// Protected virtual
// void afterAnimationStep(qreal)
func (this *QGraphicsItemAnimation) AfterAnimationStep(step float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGraphicsItemAnimation18afterAnimationStepEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

//  body block end
