package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickTransform struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickTransform) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQuickTransform10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:62
// index:0
// Public
// void QQuickTransform(class QObject *)
func NewQQuickTransform(parent *qtcore.QObject /*777 QObject **/) *QQuickTransform {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQuickTransformC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:63
// index:0
// Public virtual
// void ~QQuickTransform()
func DeleteQQuickTransform(*QQuickTransform) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQuickTransformD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:65
// index:0
// Public
// void appendToItem(class QQuickItem *)
func (this *QQuickTransform) AppendToItem(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQuickTransform12appendToItemEP10QQuickItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:66
// index:0
// Public
// void prependToItem(class QQuickItem *)
func (this *QQuickTransform) PrependToItem(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQuickTransform13prependToItemEP10QQuickItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:68
// index:0
// Public pure virtual
// void applyTo(class QMatrix4x4 *)
func (this *QQuickTransform) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQuickTransform7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:71
// index:0
// Protected
// void update()
func (this *QQuickTransform) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQuickTransform6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
