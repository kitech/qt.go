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
type QStyleOptionSpinBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionSpinBox) GetCthis() unsafe.Pointer {
	return this.QStyleOptionComplex.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyleoption.h:552
// index:0
// void QStyleOptionSpinBox()
func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(cthis)
	return gothis
}
func NewQStyleOptionSpinBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSpinBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:556
// index:1
// void QStyleOptionSpinBox(int)
func NewQStyleOptionSpinBox_1(version int) *QStyleOptionSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(cthis)
	return gothis
}

//  body block end
