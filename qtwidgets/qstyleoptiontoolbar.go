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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionToolBar struct {
	*QStyleOption
}
type QStyleOptionToolBar_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBar_PTR() *QStyleOptionToolBar
}

func (ptr *QStyleOptionToolBar) QStyleOptionToolBar_PTR() *QStyleOptionToolBar { return ptr }

func (this *QStyleOptionToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionToolBarFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBar {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionToolBar{bcthis0}
}
func (*QStyleOptionToolBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBar {
	return NewQStyleOptionToolBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBar()
func NewQStyleOptionToolBar() *QStyleOptionToolBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBarC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBar)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:319
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionToolBar(int)
func NewQStyleOptionToolBar_1(version int) *QStyleOptionToolBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBarC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolBar)
	return gothis
}

func DeleteQStyleOptionToolBar(this *QStyleOptionToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionToolBar__StyleOptionType = int

const QStyleOptionToolBar__Type QStyleOptionToolBar__StyleOptionType = 14

type QStyleOptionToolBar__StyleOptionVersion = int

const QStyleOptionToolBar__Version QStyleOptionToolBar__StyleOptionVersion = 1

type QStyleOptionToolBar__ToolBarPosition = int

const QStyleOptionToolBar__Beginning QStyleOptionToolBar__ToolBarPosition = 0
const QStyleOptionToolBar__Middle QStyleOptionToolBar__ToolBarPosition = 1
const QStyleOptionToolBar__End QStyleOptionToolBar__ToolBarPosition = 2
const QStyleOptionToolBar__OnlyOne QStyleOptionToolBar__ToolBarPosition = 3

type QStyleOptionToolBar__ToolBarFeature = int

const QStyleOptionToolBar__None QStyleOptionToolBar__ToolBarFeature = 0
const QStyleOptionToolBar__Movable QStyleOptionToolBar__ToolBarFeature = 1

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
