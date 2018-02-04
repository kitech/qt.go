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

type QStyleOptionToolButton struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionToolButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionToolButton) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionToolButtonFromPointer(cthis unsafe.Pointer) *QStyleOptionToolButton {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionToolButton{bcthis0}
}
func (*QStyleOptionToolButton) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionToolButton {
	return NewQStyleOptionToolButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:579
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionToolButton()
func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolButton)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:583
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionToolButton(int)
func NewQStyleOptionToolButton_1(version int) *QStyleOptionToolButton {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonC2Ei", qtrt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionToolButton)
	return gothis
}

func DeleteQStyleOptionToolButton(this *QStyleOptionToolButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionToolButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionToolButton__StyleOptionType = int

const QStyleOptionToolButton__Type QStyleOptionToolButton__StyleOptionType = 983043

type QStyleOptionToolButton__StyleOptionVersion = int

const QStyleOptionToolButton__Version QStyleOptionToolButton__StyleOptionVersion = 1

type QStyleOptionToolButton__ToolButtonFeature = int

const QStyleOptionToolButton__None QStyleOptionToolButton__ToolButtonFeature = 0
const QStyleOptionToolButton__Arrow QStyleOptionToolButton__ToolButtonFeature = 1
const QStyleOptionToolButton__Menu QStyleOptionToolButton__ToolButtonFeature = 4
const QStyleOptionToolButton__MenuButtonPopup QStyleOptionToolButton__ToolButtonFeature = 4
const QStyleOptionToolButton__PopupDelay QStyleOptionToolButton__ToolButtonFeature = 8
const QStyleOptionToolButton__HasMenu QStyleOptionToolButton__ToolButtonFeature = 16

//  body block end
