package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
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

type QHBoxLayout struct {
	*QBoxLayout
}

func (this *QHBoxLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBoxLayout.GetCthis()
	}
}
func (this *QHBoxLayout) SetCthis(cthis unsafe.Pointer) {
	this.QBoxLayout = NewQBoxLayoutFromPointer(cthis)
}
func NewQHBoxLayoutFromPointer(cthis unsafe.Pointer) *QHBoxLayout {
	bcthis0 := NewQBoxLayoutFromPointer(cthis)
	return &QHBoxLayout{bcthis0}
}
func (*QHBoxLayout) NewFromPointer(cthis unsafe.Pointer) *QHBoxLayout {
	return NewQHBoxLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QHBoxLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHBoxLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout()
func NewQHBoxLayout() *QHBoxLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:118
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout(QWidget *)
func NewQHBoxLayout_1(parent *QWidget /*777 QWidget **/) *QHBoxLayout {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHBoxLayout()
func DeleteQHBoxLayout(this *QHBoxLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
