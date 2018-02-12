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

type QStyleOptionTab struct {
	*QStyleOption
}

func (this *QStyleOptionTab) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTab) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabFromPointer(cthis unsafe.Pointer) *QStyleOptionTab {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTab{bcthis0}
}
func (*QStyleOptionTab) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTab {
	return NewQStyleOptionTabFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTab()
func NewQStyleOptionTab() *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:289
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTab(int)
func NewQStyleOptionTab_1(version int) *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

func DeleteQStyleOptionTab(this *QStyleOptionTab) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionTab__StyleOptionType = int

const QStyleOptionTab__Type QStyleOptionTab__StyleOptionType = 3

type QStyleOptionTab__StyleOptionVersion = int

const QStyleOptionTab__Version QStyleOptionTab__StyleOptionVersion = 3

type QStyleOptionTab__TabPosition = int

const QStyleOptionTab__Beginning QStyleOptionTab__TabPosition = 0
const QStyleOptionTab__Middle QStyleOptionTab__TabPosition = 1
const QStyleOptionTab__End QStyleOptionTab__TabPosition = 2
const QStyleOptionTab__OnlyOneTab QStyleOptionTab__TabPosition = 3

type QStyleOptionTab__SelectedPosition = int

const QStyleOptionTab__NotAdjacent QStyleOptionTab__SelectedPosition = 0
const QStyleOptionTab__NextIsSelected QStyleOptionTab__SelectedPosition = 1
const QStyleOptionTab__PreviousIsSelected QStyleOptionTab__SelectedPosition = 2

type QStyleOptionTab__CornerWidget = int

const QStyleOptionTab__NoCornerWidgets QStyleOptionTab__CornerWidget = 0
const QStyleOptionTab__LeftCornerWidget QStyleOptionTab__CornerWidget = 1
const QStyleOptionTab__RightCornerWidget QStyleOptionTab__CornerWidget = 2

type QStyleOptionTab__TabFeature = int

const QStyleOptionTab__None QStyleOptionTab__TabFeature = 0
const QStyleOptionTab__HasFrame QStyleOptionTab__TabFeature = 1

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
