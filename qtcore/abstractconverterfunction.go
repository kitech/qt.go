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
type AbstractConverterFunction struct {
	*qtrt.CObject
}
type AbstractConverterFunction_ITF interface {
	AbstractConverterFunction_PTR() *AbstractConverterFunction
}

func (ptr *AbstractConverterFunction) AbstractConverterFunction_PTR() *AbstractConverterFunction {
	return ptr
}

func (this *AbstractConverterFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractConverterFunction) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewAbstractConverterFunctionFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return &AbstractConverterFunction{&qtrt.CObject{cthis}}
}
func (*AbstractConverterFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return NewAbstractConverterFunctionFromPointer(cthis)
}

func DeleteAbstractConverterFunction(this *AbstractConverterFunction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25AbstractConverterFunctionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10239() {
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
