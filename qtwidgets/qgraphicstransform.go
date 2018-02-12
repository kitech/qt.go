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
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void update()
func (this *QGraphicsTransform) InheritUpdate(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "update", f)
}

type QGraphicsTransform struct {
	*qtcore.QObject
}

func (this *QGraphicsTransform) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsTransform) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsTransformFromPointer(cthis unsafe.Pointer) *QGraphicsTransform {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsTransform{bcthis0}
}
func (*QGraphicsTransform) NewFromPointer(cthis unsafe.Pointer) *QGraphicsTransform {
	return NewQGraphicsTransformFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsTransform) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsTransform10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTransform(QObject *)
func NewQGraphicsTransform(parent *qtcore.QObject /*777 QObject **/) *QGraphicsTransform {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransformC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsTransform()
func DeleteQGraphicsTransform(this *QGraphicsTransform) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransformD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void applyTo(QMatrix4x4 *)
func (this *QGraphicsTransform) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsTransform7applyToEP10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:66
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void update()
func (this *QGraphicsTransform) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransform6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
