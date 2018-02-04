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

type QStyleOptionComboBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionComboBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionComboBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionComboBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionComboBox{bcthis0}
}
func (*QStyleOptionComboBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionComboBox {
	return NewQStyleOptionComboBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:601
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionComboBox()
func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComboBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:605
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionComboBox(int)
func NewQStyleOptionComboBox_1(version int) *QStyleOptionComboBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComboBox)
	return gothis
}

func DeleteQStyleOptionComboBox(this *QStyleOptionComboBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionComboBox__StyleOptionType = int

const QStyleOptionComboBox__Type QStyleOptionComboBox__StyleOptionType = 983044

type QStyleOptionComboBox__StyleOptionVersion = int

const QStyleOptionComboBox__Version QStyleOptionComboBox__StyleOptionVersion = 1

//  body block end
