//  header block begin
// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>
package qtwidgets

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
type QGraphicsRotation struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:120
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsRotation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:126
// index:0
// void QGraphicsRotation(class QObject *)
func NewQGraphicsRotation(parent unsafe.Pointer) *QGraphicsRotation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsRotation{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:127
// index:0
// virtual
// void ~QGraphicsRotation()
func DeleteQGraphicsRotation(*QGraphicsRotation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:129
// index:0
// QVector3D origin()
func (this *QGraphicsRotation) Origin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation6originEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:130
// index:0
// void setOrigin(const class QVector3D &)
func (this *QGraphicsRotation) SetOrigin(point unsafe.Pointer) {
	// 0: (, const QVector3D & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation9setOriginERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:132
// index:0
// qreal angle()
func (this *QGraphicsRotation) Angle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation5angleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:133
// index:0
// void setAngle(qreal)
func (this *QGraphicsRotation) SetAngle(arg0 float64) {
	// 0: (, qreal arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation8setAngleEd", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:135
// index:0
// QVector3D axis()
func (this *QGraphicsRotation) Axis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation4axisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:136
// index:0
// void setAxis(const class QVector3D &)
func (this *QGraphicsRotation) SetAxis(axis unsafe.Pointer) {
	// 0: (, const QVector3D & axis), (axis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisERK9QVector3D", ffiqt.FFI_TYPE_VOID, this.cthis, axis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:137
// index:1
// void setAxis(Qt::Axis)
func (this *QGraphicsRotation) SetAxis_1(axis int) {
	// 1: (, Qt::Axis axis), (&axis)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation7setAxisEN2Qt4AxisE", ffiqt.FFI_TYPE_VOID, this.cthis, &axis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:139
// index:0
// virtual
// void applyTo(class QMatrix4x4 *)
func (this *QGraphicsRotation) ApplyTo(matrix unsafe.Pointer) {
	// 0: (, QMatrix4x4 * matrix), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsRotation7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_VOID, this.cthis, matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:142
// index:0
// void originChanged()
func (this *QGraphicsRotation) OriginChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation13originChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:143
// index:0
// void angleChanged()
func (this *QGraphicsRotation) AngleChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation12angleChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:144
// index:0
// void axisChanged()
func (this *QGraphicsRotation) AxisChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsRotation11axisChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
