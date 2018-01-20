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
type QStyleHintReturnVariant struct {
	*QStyleHintReturn
}

func (this *QStyleHintReturnVariant) GetCthis() unsafe.Pointer {
	return this.QStyleHintReturn.GetCthis()
}
func NewQStyleHintReturnVariantFromPointer(cthis unsafe.Pointer) *QStyleHintReturnVariant {
	bcthis0 := NewQStyleHintReturnFromPointer(cthis)
	return &QStyleHintReturnVariant{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:733
// index:0
// Public
// void QStyleHintReturnVariant()
func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QStyleHintReturnVariantC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnVariantFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:734
// index:0
// Public
// void ~QStyleHintReturnVariant()
func DeleteQStyleHintReturnVariant(*QStyleHintReturnVariant) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QStyleHintReturnVariantD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
