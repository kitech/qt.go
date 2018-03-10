package qtcore

// /usr/include/qt/QtCore/qmetatype.h
// #include <qmetatype.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type AbstractDebugStreamFunction struct {
	*qtrt.CObject
}
type AbstractDebugStreamFunction_ITF interface {
	AbstractDebugStreamFunction_PTR() *AbstractDebugStreamFunction
}

func (ptr *AbstractDebugStreamFunction) AbstractDebugStreamFunction_PTR() *AbstractDebugStreamFunction {
	return ptr
}

func (this *AbstractDebugStreamFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractDebugStreamFunction) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewAbstractDebugStreamFunctionFromPointer(cthis unsafe.Pointer) *AbstractDebugStreamFunction {
	return &AbstractDebugStreamFunction{&qtrt.CObject{cthis}}
}
func (*AbstractDebugStreamFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractDebugStreamFunction {
	return NewAbstractDebugStreamFunctionFromPointer(cthis)
}

func DeleteAbstractDebugStreamFunction(this *AbstractDebugStreamFunction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27AbstractDebugStreamFunctionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

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
}

//  keep block end
