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
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsScale) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:88
// index:0
// Public
// void QGraphicsScale(QObject *)
func NewQGraphicsScale(parent *qtcore.QObject /*777 QObject **/) *QGraphicsScale {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScaleC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsScaleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:89
// index:0
// Public virtual
// void ~QGraphicsScale()
func DeleteQGraphicsScale(*QGraphicsScale) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScaleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:91
// index:0
// Public
// QVector3D origin()
func (this *QGraphicsScale) Origin() *qtgui.QVector3D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6originEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:92
// index:0
// Public
// void setOrigin(const QVector3D &)
func (this *QGraphicsScale) SetOrigin(point *qtgui.QVector3D) {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setOriginERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:94
// index:0
// Public
// qreal xScale()
func (this *QGraphicsScale) XScale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6xScaleEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:95
// index:0
// Public
// void setXScale(qreal)
func (this *QGraphicsScale) SetXScale(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setXScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:97
// index:0
// Public
// qreal yScale()
func (this *QGraphicsScale) YScale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6yScaleEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:98
// index:0
// Public
// void setYScale(qreal)
func (this *QGraphicsScale) SetYScale(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setYScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:100
// index:0
// Public
// qreal zScale()
func (this *QGraphicsScale) ZScale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale6zScaleEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:101
// index:0
// Public
// void setZScale(qreal)
func (this *QGraphicsScale) SetZScale(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale9setZScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:103
// index:0
// Public virtual
// void applyTo(QMatrix4x4 *)
func (this *QGraphicsScale) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGraphicsScale7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:106
// index:0
// Public
// void originChanged()
func (this *QGraphicsScale) OriginChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13originChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:107
// index:0
// Public
// void xScaleChanged()
func (this *QGraphicsScale) XScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13xScaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:108
// index:0
// Public
// void yScaleChanged()
func (this *QGraphicsScale) YScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13yScaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:109
// index:0
// Public
// void zScaleChanged()
func (this *QGraphicsScale) ZScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale13zScaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:110
// index:0
// Public
// void scaleChanged()
func (this *QGraphicsScale) ScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGraphicsScale12scaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
