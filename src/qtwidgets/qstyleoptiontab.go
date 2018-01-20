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
type QStyleOptionTab struct {
	*QStyleOption
}

func (this *QStyleOptionTab) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyleoption.h:285
// index:0
// void QStyleOptionTab()
func NewQStyleOptionTab() *QStyleOptionTab {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(cthis)
	return gothis
}
func NewQStyleOptionTabFromPointer(cthis unsafe.Pointer) *QStyleOptionTab {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTab{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:289
// index:1
// void QStyleOptionTab(int)
func NewQStyleOptionTab_1(version int) *QStyleOptionTab {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(cthis)
	return gothis
}

//  body block end
