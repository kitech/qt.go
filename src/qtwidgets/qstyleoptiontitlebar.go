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
type QStyleOptionTitleBar struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionTitleBar) GetCthis() unsafe.Pointer {
	return this.QStyleOptionComplex.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyleoption.h:619
// index:0
// void QStyleOptionTitleBar()
func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(cthis)
	return gothis
}
func NewQStyleOptionTitleBarFromPointer(cthis unsafe.Pointer) *QStyleOptionTitleBar {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionTitleBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:623
// index:1
// void QStyleOptionTitleBar(int)
func NewQStyleOptionTitleBar_1(version int) *QStyleOptionTitleBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(cthis)
	return gothis
}

//  body block end
