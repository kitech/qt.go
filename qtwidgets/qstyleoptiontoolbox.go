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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QStyleOptionToolBox struct {
	*QStyleOption
}

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
// Public
// void QStyleOptionToolBox()
func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:476
// index:1
// Protected
// void QStyleOptionToolBox(int)
func NewQStyleOptionToolBox_1(version int) *QStyleOptionToolBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(cthis)
	return gothis
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
