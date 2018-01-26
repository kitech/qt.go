package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QStyleHintReturn struct {
	*qtrt.CObject
}

func (this *QStyleHintReturn) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleHintReturn) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStyleHintReturnFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return &QStyleHintReturn{&qtrt.CObject{cthis}}
}
func (*QStyleHintReturn) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return NewQStyleHintReturnFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// Public
// void QStyleHintReturn(int, int)
func NewQStyleHintReturn(version int, type_ int) *QStyleHintReturn {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", ffiqt.FFI_TYPE_VOID, cthis, version, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:711
// index:0
// Public
// void ~QStyleHintReturn()
func DeleteQStyleHintReturn(*QStyleHintReturn) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QStyleHintReturnD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

type QStyleHintReturn__HintReturnType = int

const QStyleHintReturn__SH_Default QStyleHintReturn__HintReturnType = 61440
const QStyleHintReturn__SH_Mask QStyleHintReturn__HintReturnType = 61441
const QStyleHintReturn__SH_Variant QStyleHintReturn__HintReturnType = 61442

type QStyleHintReturn__StyleOptionType = int

const QStyleHintReturn__Type QStyleHintReturn__StyleOptionType = 61440

type QStyleHintReturn__StyleOptionVersion = int

const QStyleHintReturn__Version QStyleHintReturn__StyleOptionVersion = 1

//  body block end
