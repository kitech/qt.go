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
type QStyleOptionMenuItem struct {
	*QStyleOption
}

func (this *QStyleOptionMenuItem) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}
func NewQStyleOptionMenuItemFromPointer(cthis unsafe.Pointer) *QStyleOptionMenuItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionMenuItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:372
// index:0
// Public
// void QStyleOptionMenuItem()
func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:376
// index:1
// Protected
// void QStyleOptionMenuItem(int)
func NewQStyleOptionMenuItem_1(version int) *QStyleOptionMenuItem {
	cthis := qtrt.Calloc(1, 256) // 136
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionMenuItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionMenuItemFromPointer(cthis)
	return gothis
}

//  body block end
