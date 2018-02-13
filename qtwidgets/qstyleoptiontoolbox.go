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

type QStyleOptionToolBox struct {
	*QStyleOption
}
type QStyleOptionToolBox_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBox_PTR() *QStyleOptionToolBox
}

func (ptr *QStyleOptionToolBox) QStyleOptionToolBox_PTR() *QStyleOptionToolBox { return ptr }

func (this *QStyleOptionToolBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionToolBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionToolBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBox {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionToolBox{bcthis0}
}
func (*QStyleOptionToolBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBox {
	return NewQStyleOptionToolBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:472
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBox()
func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:476
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBox(int)
func NewQStyleOptionToolBox_1(version int) *QStyleOptionToolBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBox)
	return gothis
}

func DeleteQStyleOptionToolBox(this *QStyleOptionToolBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionToolBox__StyleOptionType = int

const QStyleOptionToolBox__Type QStyleOptionToolBox__StyleOptionType = 7

type QStyleOptionToolBox__StyleOptionVersion = int

const QStyleOptionToolBox__Version QStyleOptionToolBox__StyleOptionVersion = 2

type QStyleOptionToolBox__TabPosition = int

const QStyleOptionToolBox__Beginning QStyleOptionToolBox__TabPosition = 0
const QStyleOptionToolBox__Middle QStyleOptionToolBox__TabPosition = 1
const QStyleOptionToolBox__End QStyleOptionToolBox__TabPosition = 2
const QStyleOptionToolBox__OnlyOneTab QStyleOptionToolBox__TabPosition = 3

type QStyleOptionToolBox__SelectedPosition = int

const QStyleOptionToolBox__NotAdjacent QStyleOptionToolBox__SelectedPosition = 0
const QStyleOptionToolBox__NextIsSelected QStyleOptionToolBox__SelectedPosition = 1
const QStyleOptionToolBox__PreviousIsSelected QStyleOptionToolBox__SelectedPosition = 2

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
