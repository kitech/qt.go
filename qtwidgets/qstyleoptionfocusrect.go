package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionFocusRect struct {
	*QStyleOption
}

func (this *QStyleOptionFocusRect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionFocusRect) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionFocusRectFromPointer(cthis unsafe.Pointer) *QStyleOptionFocusRect {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFocusRect{bcthis0}
}
func (*QStyleOptionFocusRect) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionFocusRect {
	return NewQStyleOptionFocusRectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionFocusRect()
func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFocusRect)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:123
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionFocusRect(int)
func NewQStyleOptionFocusRect_1(version int) *QStyleOptionFocusRect {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFocusRect)
	return gothis
}

func DeleteQStyleOptionFocusRect(this *QStyleOptionFocusRect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionFocusRect__StyleOptionType = int

const QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = 1

type QStyleOptionFocusRect__StyleOptionVersion = int

const QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = 1

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
