package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// void update()
func (this *QQuickTransform) InheritUpdate(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "update", f)
}

type QQuickTransform struct {
	*qtcore.QObject
}
type QQuickTransform_ITF interface {
	qtcore.QObject_ITF
	QQuickTransform_PTR() *QQuickTransform
}

func (ptr *QQuickTransform) QQuickTransform_PTR() *QQuickTransform { return ptr }

func (this *QQuickTransform) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickTransform) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickTransformFromPointer(cthis unsafe.Pointer) *QQuickTransform {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickTransform{bcthis0}
}
func (*QQuickTransform) NewFromPointer(cthis unsafe.Pointer) *QQuickTransform {
	return NewQQuickTransformFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickitem.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickTransform) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQuickTransform10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickTransform(QObject *)
func NewQQuickTransform(parent *qtcore.QObject /*777 QObject **/) *QQuickTransform {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQuickTransformC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickTransformFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickTransform()
func DeleteQQuickTransform(this *QQuickTransform) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQuickTransformD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickitem.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendToItem(QQuickItem *)
func (this *QQuickTransform) AppendToItem(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQuickTransform12appendToItemEP10QQuickItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prependToItem(QQuickItem *)
func (this *QQuickTransform) PrependToItem(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQuickTransform13prependToItemEP10QQuickItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void applyTo(QMatrix4x4 *)
func (this *QQuickTransform) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQuickTransform7applyToEP10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:71
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void update()
func (this *QQuickTransform) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQuickTransform6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
