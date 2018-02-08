package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QGraphicsScale struct {
	*QGraphicsTransform
}

func (this *QGraphicsScale) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsTransform.GetCthis()
	}
}
func (this *QGraphicsScale) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsTransform = NewQGraphicsTransformFromPointer(cthis)
}
func NewQGraphicsScaleFromPointer(cthis unsafe.Pointer) *QGraphicsScale {
	bcthis0 := NewQGraphicsTransformFromPointer(cthis)
	return &QGraphicsScale{bcthis0}
}
func (*QGraphicsScale) NewFromPointer(cthis unsafe.Pointer) *QGraphicsScale {
	return NewQGraphicsScaleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsScale) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsScale(QObject *)
func NewQGraphicsScale(parent *qtcore.QObject /*777 QObject **/) *QGraphicsScale {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScaleC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsScaleFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsScale()
func DeleteQGraphicsScale(this *QGraphicsScale) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScaleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:91
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D origin()
func (this *QGraphicsScale) Origin() *qtgui.QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale6originEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrigin(const QVector3D &)
func (this *QGraphicsScale) SetOrigin(point *qtgui.QVector3D) {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale9setOriginERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xScale()
func (this *QGraphicsScale) XScale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale6xScaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setXScale(qreal)
func (this *QGraphicsScale) SetXScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale9setXScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal yScale()
func (this *QGraphicsScale) YScale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale6yScaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setYScale(qreal)
func (this *QGraphicsScale) SetYScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale9setYScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal zScale()
func (this *QGraphicsScale) ZScale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale6zScaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZScale(qreal)
func (this *QGraphicsScale) SetZScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale9setZScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void applyTo(QMatrix4x4 *)
func (this *QGraphicsScale) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QGraphicsScale7applyToEP10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void originChanged()
func (this *QGraphicsScale) OriginChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale13originChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void xScaleChanged()
func (this *QGraphicsScale) XScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale13xScaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void yScaleChanged()
func (this *QGraphicsScale) YScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale13yScaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zScaleChanged()
func (this *QGraphicsScale) ZScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale13zScaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scaleChanged()
func (this *QGraphicsScale) ScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QGraphicsScale12scaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
