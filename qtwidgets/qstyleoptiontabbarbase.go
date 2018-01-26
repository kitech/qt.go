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
import "mkuse/cffiqt"
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
type QStyleOptionTabBarBase struct {
	*QStyleOption
}

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
// Public
// void QStyleOptionTabBarBase()
func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {
	cthis := qtrt.Calloc(1, 256) // 104
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabBarBaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:199
// index:1
// Protected
// void QStyleOptionTabBarBase(int)
func NewQStyleOptionTabBarBase_1(version int) *QStyleOptionTabBarBase {
	cthis := qtrt.Calloc(1, 256) // 104
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabBarBaseFromPointer(cthis)
	return gothis
}

type QStyleOptionTabBarBase__StyleOptionType = int

const QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = 12

type QStyleOptionTabBarBase__StyleOptionVersion = int

const QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = 2

//  body block end
