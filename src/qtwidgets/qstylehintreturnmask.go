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
type QStyleHintReturnMask struct {
	*QStyleHintReturn
}

func (this *QStyleHintReturnMask) GetCthis() unsafe.Pointer {
	return this.QStyleHintReturn.GetCthis()
}
func NewQStyleHintReturnMaskFromPointer(cthis unsafe.Pointer) *QStyleHintReturnMask {
	bcthis0 := NewQStyleHintReturnFromPointer(cthis)
	return &QStyleHintReturnMask{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:722
// index:0
// Public
// void QStyleHintReturnMask()
func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnMaskFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:723
// index:0
// Public
// void ~QStyleHintReturnMask()
func DeleteQStyleHintReturnMask(*QStyleHintReturnMask) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
