package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

type QVBoxLayout struct {
	*QBoxLayout
}
type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout { return ptr }

func (this *QVBoxLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBoxLayout.GetCthis()
	}
}
func (this *QVBoxLayout) SetCthis(cthis unsafe.Pointer) {
	this.QBoxLayout = NewQBoxLayoutFromPointer(cthis)
}
func NewQVBoxLayoutFromPointer(cthis unsafe.Pointer) *QVBoxLayout {
	bcthis0 := NewQBoxLayoutFromPointer(cthis)
	return &QVBoxLayout{bcthis0}
}
func (*QVBoxLayout) NewFromPointer(cthis unsafe.Pointer) *QVBoxLayout {
	return NewQVBoxLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QVBoxLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QVBoxLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QVBoxLayout()
func NewQVBoxLayout() *QVBoxLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVBoxLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:131
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QVBoxLayout(QWidget *)
func NewQVBoxLayout_1(parent QWidget_ITF /*777 QWidget **/) *QVBoxLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVBoxLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVBoxLayout()
func DeleteQVBoxLayout(this *QVBoxLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QVBoxLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
