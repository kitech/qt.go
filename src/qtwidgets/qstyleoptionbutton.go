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
type QStyleOptionButton struct {
	*QStyleOption
}

func (this *QStyleOptionButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionButton) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionButtonFromPointer(cthis unsafe.Pointer) *QStyleOptionButton {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionButton{bcthis0}
}
func (*QStyleOptionButton) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionButton {
	return NewQStyleOptionButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:248
// index:0
// Public
// void QStyleOptionButton()
func NewQStyleOptionButton() *QStyleOptionButton {
	cthis := qtrt.Calloc(1, 256) // 96
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionButtonC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:252
// index:1
// Protected
// void QStyleOptionButton(int)
func NewQStyleOptionButton_1(version int) *QStyleOptionButton {
	cthis := qtrt.Calloc(1, 256) // 96
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStyleOptionButtonC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionButtonFromPointer(cthis)
	return gothis
}

type QStyleOptionButton__StyleOptionType = int

const QStyleOptionButton__Type QStyleOptionButton__StyleOptionType = 2

type QStyleOptionButton__StyleOptionVersion = int

const QStyleOptionButton__Version QStyleOptionButton__StyleOptionVersion = 1

type QStyleOptionButton__ButtonFeature = int

const QStyleOptionButton__None QStyleOptionButton__ButtonFeature = 0
const QStyleOptionButton__Flat QStyleOptionButton__ButtonFeature = 1
const QStyleOptionButton__HasMenu QStyleOptionButton__ButtonFeature = 2
const QStyleOptionButton__DefaultButton QStyleOptionButton__ButtonFeature = 4
const QStyleOptionButton__AutoDefaultButton QStyleOptionButton__ButtonFeature = 8
const QStyleOptionButton__CommandLinkButton QStyleOptionButton__ButtonFeature = 16

//  body block end
