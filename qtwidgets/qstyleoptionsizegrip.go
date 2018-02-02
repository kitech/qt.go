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

type QStyleOptionSizeGrip struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSizeGrip) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSizeGrip) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSizeGripFromPointer(cthis unsafe.Pointer) *QStyleOptionSizeGrip {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSizeGrip{bcthis0}
}
func (*QStyleOptionSizeGrip) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGripFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:653
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionSizeGrip()
func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSizeGrip)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:656
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionSizeGrip(int)
func NewQStyleOptionSizeGrip_1(version int) *QStyleOptionSizeGrip {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripC2Ei", ffiqt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSizeGripFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSizeGrip)
	return gothis
}

func DeleteQStyleOptionSizeGrip(this *QStyleOptionSizeGrip) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionSizeGripD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionSizeGrip__StyleOptionType = int

const QStyleOptionSizeGrip__Type QStyleOptionSizeGrip__StyleOptionType = 983047

type QStyleOptionSizeGrip__StyleOptionVersion = int

const QStyleOptionSizeGrip__Version QStyleOptionSizeGrip__StyleOptionVersion = 1

//  body block end
