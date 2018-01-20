//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

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
type QStyleOptionTabBarBase struct {
	*QStyleOption
}

func (this *QStyleOptionTabBarBase) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}
func NewQStyleOptionTabBarBaseFromPointer(cthis unsafe.Pointer) *QStyleOptionTabBarBase {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabBarBase{bcthis0}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionTabBarBaseC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabBarBaseFromPointer(cthis)
	return gothis
}

//  body block end
