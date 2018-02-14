package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleHintReturn struct {
	*qtrt.CObject
}
type QStyleHintReturn_ITF interface {
	QStyleHintReturn_PTR() *QStyleHintReturn
}

func (ptr *QStyleHintReturn) QStyleHintReturn_PTR() *QStyleHintReturn { return ptr }

func (this *QStyleHintReturn) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleHintReturn) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStyleHintReturnFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return &QStyleHintReturn{&qtrt.CObject{cthis}}
}
func (*QStyleHintReturn) NewFromPointer(cthis unsafe.Pointer) *QStyleHintReturn {
	return NewQStyleHintReturnFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:710
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturn(int, int)
func NewQStyleHintReturn(version int, type_ int) *QStyleHintReturn {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturn)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:711
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleHintReturn()
func DeleteQStyleHintReturn(this *QStyleHintReturn) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStyleHintReturnD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
