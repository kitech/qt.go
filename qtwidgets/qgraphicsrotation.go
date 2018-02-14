package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QGraphicsRotation struct {
	*QGraphicsTransform
}
type QGraphicsRotation_ITF interface {
	QGraphicsTransform_ITF
	QGraphicsRotation_PTR() *QGraphicsRotation
}

func (ptr *QGraphicsRotation) QGraphicsRotation_PTR() *QGraphicsRotation { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsRotation) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRotation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRotation(QObject *)
func NewQGraphicsRotation(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsRotation {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRotationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsRotation")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsRotation()
func DeleteQGraphicsRotation(this *QGraphicsRotation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:129
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D origin()
func (this *QGraphicsRotation) Origin() *qtgui.QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRotation6originEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrigin(const QVector3D &)
func (this *QGraphicsRotation) SetOrigin(point qtgui.QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QVector3D_PTR() != nil {
		convArg0 = point.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation9setOriginERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angle()
func (this *QGraphicsRotation) Angle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRotation5angleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAngle(qreal)
func (this *QGraphicsRotation) SetAngle(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation8setAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:135
// index:0
// Public Visibility=Default Availability=Available
// [12] QVector3D axis()
func (this *QGraphicsRotation) Axis() *qtgui.QVector3D /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRotation4axisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQVector3DFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQVector3D)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAxis(const QVector3D &)
func (this *QGraphicsRotation) SetAxis(axis qtgui.QVector3D_ITF) {
	var convArg0 unsafe.Pointer
	if axis != nil && axis.QVector3D_PTR() != nil {
		convArg0 = axis.QVector3D_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisERK9QVector3D", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:137
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAxis(Qt::Axis)
func (this *QGraphicsRotation) SetAxis_1(axis int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisEN2Qt4AxisE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), axis)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void applyTo(QMatrix4x4 *)
func (this *QGraphicsRotation) ApplyTo(matrix qtgui.QMatrix4x4_ITF /*777 QMatrix4x4 **/) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix4x4_PTR() != nil {
		convArg0 = matrix.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRotation7applyToEP10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void originChanged()
func (this *QGraphicsRotation) OriginChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation13originChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void angleChanged()
func (this *QGraphicsRotation) AngleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation12angleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void axisChanged()
func (this *QGraphicsRotation) AxisChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRotation11axisChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
