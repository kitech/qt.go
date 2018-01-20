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
type QStyleOptionToolBox struct {
	*QStyleOption
}

func (this *QStyleOptionToolBox) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}
func NewQStyleOptionToolBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionToolBox {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionToolBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:472
// index:0
// Public
// void QStyleOptionToolBox()
func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:476
// index:1
// Protected
// void QStyleOptionToolBox(int)
func NewQStyleOptionToolBox_1(version int) *QStyleOptionToolBox {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionToolBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionToolBoxFromPointer(cthis)
	return gothis
}

//  body block end
