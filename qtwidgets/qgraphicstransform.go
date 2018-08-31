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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void update()
func (this *QGraphicsTransform) InheritUpdate(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "update", f)
}

/*

 */
type QGraphicsTransform struct {
	*qtcore.QObject
}
type QGraphicsTransform_ITF interface {
	qtcore.QObject_ITF
	QGraphicsTransform_PTR() *QGraphicsTransform
}

func (ptr *QGraphicsTransform) QGraphicsTransform_PTR() *QGraphicsTransform { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsTransform) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsTransform10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTransform(QObject *)

/*
Constructs a new QGraphicsTransform with the given parent.
*/
func (*QGraphicsTransform) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsTransform {
	return NewQGraphicsTransform(parent)
}
func NewQGraphicsTransform(parent qtcore.QObject_ITF /*777 QObject **/) *QGraphicsTransform {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransformC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTransform")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTransform(QObject *)

/*
Constructs a new QGraphicsTransform with the given parent.
*/
func (*QGraphicsTransform) NewForInherit__() *QGraphicsTransform {
	return NewQGraphicsTransform__()
}
func NewQGraphicsTransform__() *QGraphicsTransform {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransformC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsTransform")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsTransform()

/*

 */
func DeleteQGraphicsTransform(this *QGraphicsTransform) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsTransformD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void applyTo(QMatrix4x4 *) const

/*
This pure virtual method has to be reimplemented in derived classes.

It applies this transformation to matrix.

See also QGraphicsItem::transform() and QMatrix4x4::toTransform().
*/
func (this *QGraphicsTransform) ApplyTo(matrix qtgui.QMatrix4x4_ITF /*777 QMatrix4x4 **/) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix4x4_PTR() != nil {
		convArg0 = matrix.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsTransform7applyToEP10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:66
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void update()

/*
Notifies that this transform operation has changed its parameters in such a way that applyTo() will return a different result than before.

When implementing you own custom graphics transform, you must call this function every time you change a parameter, to let QGraphicsItem know that its transformation needs to be updated.

See also applyTo().
*/
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
