package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QStyleOptionTabWidgetFrame struct {
	*QStyleOption
}

func (this *QStyleOptionTabWidgetFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTabWidgetFrame) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabWidgetFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabWidgetFrame{bcthis0}
}
func (*QStyleOptionTabWidgetFrame) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	return NewQStyleOptionTabWidgetFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTabWidgetFrame()
func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabWidgetFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTabWidgetFrame)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:176
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTabWidgetFrame(int)
func NewQStyleOptionTabWidgetFrame_1(version int) *QStyleOptionTabWidgetFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameC2Ei", qtrt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabWidgetFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTabWidgetFrame)
	return gothis
}

func DeleteQStyleOptionTabWidgetFrame(this *QStyleOptionTabWidgetFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionTabWidgetFrame__StyleOptionType = int

const QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = 11

type QStyleOptionTabWidgetFrame__StyleOptionVersion = int

const QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = 2

//  body block end
