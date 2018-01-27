package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QGraphicsRotation struct {
	*QGraphicsTransform
}

func (this *QGraphicsRotation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsTransform.GetCthis()
	}
}
func (this *QGraphicsRotation) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsTransform = NewQGraphicsTransformFromPointer(cthis)
}
func NewQGraphicsRotationFromPointer(cthis unsafe.Pointer) *QGraphicsRotation {
	bcthis0 := NewQGraphicsTransformFromPointer(cthis)
	return &QGraphicsRotation{bcthis0}
}
func (*QGraphicsRotation) NewFromPointer(cthis unsafe.Pointer) *QGraphicsRotation {
	return NewQGraphicsRotationFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:120
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsRotation) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:126
// index:0
// Public
// void QGraphicsRotation(QObject *)
func NewQGraphicsRotation(parent *qtcore.QObject /*777 QObject **/) *QGraphicsRotation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsRotationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:127
// index:0
// Public virtual
// void ~QGraphicsRotation()
func DeleteQGraphicsRotation(*QGraphicsRotation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:129
// index:0
// Public
// QVector3D origin()
func (this *QGraphicsRotation) Origin() *qtgui.QVector3D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation6originEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:130
// index:0
// Public
// void setOrigin(const QVector3D &)
func (this *QGraphicsRotation) SetOrigin(point *qtgui.QVector3D) {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation9setOriginERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:132
// index:0
// Public
// qreal angle()
func (this *QGraphicsRotation) Angle() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation5angleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:133
// index:0
// Public
// void setAngle(qreal)
func (this *QGraphicsRotation) SetAngle(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation8setAngleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:135
// index:0
// Public
// QVector3D axis()
func (this *QGraphicsRotation) Axis() *qtgui.QVector3D /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation4axisEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:136
// index:0
// Public
// void setAxis(const QVector3D &)
func (this *QGraphicsRotation) SetAxis(axis *qtgui.QVector3D) {
	var convArg0 = axis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:137
// index:1
// Public
// void setAxis(Qt::Axis)
func (this *QGraphicsRotation) SetAxis_1(axis int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisEN2Qt4AxisE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), axis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:139
// index:0
// Public virtual
// void applyTo(QMatrix4x4 *)
func (this *QGraphicsRotation) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:142
// index:0
// Public
// void originChanged()
func (this *QGraphicsRotation) OriginChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation13originChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:143
// index:0
// Public
// void angleChanged()
func (this *QGraphicsRotation) AngleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation12angleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:144
// index:0
// Public
// void axisChanged()
func (this *QGraphicsRotation) AxisChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation11axisChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
