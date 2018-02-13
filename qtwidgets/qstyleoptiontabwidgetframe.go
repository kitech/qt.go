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
// extern C begin: 2
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

type QStyleOptionTabWidgetFrame struct {
	*QStyleOption
}
type QStyleOptionTabWidgetFrame_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame
}

func (ptr *QStyleOptionTabWidgetFrame) QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame {
	return ptr
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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabWidgetFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTabWidgetFrame)
	return gothis
}

func DeleteQStyleOptionTabWidgetFrame(this *QStyleOptionTabWidgetFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QStyleOptionTabWidgetFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionTabWidgetFrame__StyleOptionType = int

const QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = 11

type QStyleOptionTabWidgetFrame__StyleOptionVersion = int

const QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = 2

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
