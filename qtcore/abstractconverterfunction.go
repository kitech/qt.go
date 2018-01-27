package qtcore

// /usr/include/qt/QtCore/qmetatype.h
// #include <qmetatype.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type AbstractConverterFunction struct {
	*qtrt.CObject
}

func (this *AbstractConverterFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractConverterFunction) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewAbstractConverterFunctionFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return &AbstractConverterFunction{&qtrt.CObject{cthis}}
}
func (*AbstractConverterFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractConverterFunction {
	return NewAbstractConverterFunctionFromPointer(cthis)
}

//  body block end
