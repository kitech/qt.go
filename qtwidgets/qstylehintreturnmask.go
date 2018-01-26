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
	if this == nil {
		return nil
	} else {
		return this.QStyleHintReturn.GetCthis()
	}
}
func (this *QStyleHintReturnMask) SetCthis(cthis unsafe.Pointer) {
	this.QStyleHintReturn = NewQStyleHintReturnFromPointer(cthis)
}
func NewQStyleHintReturnMaskFromPointer(cthis unsafe.Pointer) *QStyleHintReturnMask {
	bcthis0 := NewQStyleHintReturnFromPointer(cthis)
	return &QStyleHintReturnMask{bcthis0}
}
func (*QStyleHintReturnMask) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturnMask {
	return NewQStyleHintReturnMaskFromPointer(cthis)
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

type QStyleHintReturnMask__StyleOptionType = int

const QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = 61441

type QStyleHintReturnMask__StyleOptionVersion = int

const QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = 1

//  body block end
