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

type QStyleOptionTabBarBase struct {
	*QStyleOption
}
type QStyleOptionTabBarBase_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase
}

func (ptr *QStyleOptionTabBarBase) QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase { return ptr }

func (this *QStyleOptionTabBarBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTabBarBase) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabBarBaseFromPointer(cthis unsafe.Pointer) *QStyleOptionTabBarBase {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabBarBase{bcthis0}
}
func (*QStyleOptionTabBarBase) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTabBarBase {
	return NewQStyleOptionTabBarBaseFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTabBarBase()
func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabBarBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTabBarBase)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:199
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTabBarBase(int)
func NewQStyleOptionTabBarBase_1(version int) *QStyleOptionTabBarBase {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabBarBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTabBarBase)
	return gothis
}

func DeleteQStyleOptionTabBarBase(this *QStyleOptionTabBarBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionTabBarBase__StyleOptionType = int

const QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = 12

type QStyleOptionTabBarBase__StyleOptionVersion = int

const QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = 2

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
