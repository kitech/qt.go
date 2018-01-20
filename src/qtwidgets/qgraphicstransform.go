//  header block begin
// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QGraphicsTransform struct {
	*qtcore.QObject
}

func (this *QGraphicsTransform) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsTransform) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsTransform10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:60
// index:0
// void QGraphicsTransform(class QObject *)
func NewQGraphicsTransform(parent unsafe.Pointer) *QGraphicsTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransformC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(cthis)
	return gothis
}
func NewQGraphicsTransformFromPointer(cthis unsafe.Pointer) *QGraphicsTransform {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsTransform{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:69
// index:1
// void QGraphicsTransform(class QGraphicsTransformPrivate &, class QObject *)
func NewQGraphicsTransform_1(p unsafe.Pointer, parent unsafe.Pointer) *QGraphicsTransform {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransformC1ER25QGraphicsTransformPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, p, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:61
// index:0
// virtual
// void ~QGraphicsTransform()
func DeleteQGraphicsTransform(*QGraphicsTransform) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransformD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:63
// index:0
// pure virtual
// void applyTo(class QMatrix4x4 *)
func (this *QGraphicsTransform) ApplyTo(matrix unsafe.Pointer) {
	// 0: (, matrix QMatrix4x4 *), (matrix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsTransform7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:66
// index:0
// void update()
func (this *QGraphicsTransform) Update() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransform6updateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
