package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:123
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionFocusRect(int)
func NewQStyleOptionFocusRect_1(version int) *QStyleOptionFocusRect {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QStyleOptionFocusRectC2Ei", ffiqt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFocusRectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

type QStyleOptionFocusRect__StyleOptionType = int

const QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = 1

type QStyleOptionFocusRect__StyleOptionVersion = int

const QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = 1

//  body block end
