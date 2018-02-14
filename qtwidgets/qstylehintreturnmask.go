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
// extern C begin: 2
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

type QStyleHintReturnMask struct {
	*QStyleHintReturn
}
type QStyleHintReturnMask_ITF interface {
	QStyleHintReturn_ITF
	QStyleHintReturnMask_PTR() *QStyleHintReturnMask
}

func (ptr *QStyleHintReturnMask) QStyleHintReturnMask_PTR() *QStyleHintReturnMask { return ptr }

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
// Public Visibility=Default Availability=Available
// [-2] void QStyleHintReturnMask()
func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleHintReturnMaskFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleHintReturnMask)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:723
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStyleHintReturnMask()
func DeleteQStyleHintReturnMask(this *QStyleHintReturnMask) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleHintReturnMaskD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleHintReturnMask__StyleOptionType = int

const QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = 61441

type QStyleHintReturnMask__StyleOptionVersion = int

const QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = 1

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
