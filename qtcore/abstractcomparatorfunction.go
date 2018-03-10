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
type AbstractComparatorFunction struct {
	*qtrt.CObject
}
type AbstractComparatorFunction_ITF interface {
	AbstractComparatorFunction_PTR() *AbstractComparatorFunction
}

func (ptr *AbstractComparatorFunction) AbstractComparatorFunction_PTR() *AbstractComparatorFunction {
	return ptr
}

func (this *AbstractComparatorFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractComparatorFunction) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewAbstractComparatorFunctionFromPointer(cthis unsafe.Pointer) *AbstractComparatorFunction {
	return &AbstractComparatorFunction{&qtrt.CObject{cthis}}
}
func (*AbstractComparatorFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractComparatorFunction {
	return NewAbstractComparatorFunctionFromPointer(cthis)
}

func DeleteAbstractComparatorFunction(this *AbstractComparatorFunction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26AbstractComparatorFunctionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
