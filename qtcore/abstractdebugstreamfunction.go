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
type AbstractDebugStreamFunction struct {
	*qtrt.CObject
}

func (this *AbstractDebugStreamFunction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *AbstractDebugStreamFunction) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewAbstractDebugStreamFunctionFromPointer(cthis unsafe.Pointer) *AbstractDebugStreamFunction {
	return &AbstractDebugStreamFunction{&qtrt.CObject{cthis}}
}
func (*AbstractDebugStreamFunction) NewFromPointer(cthis unsafe.Pointer) *AbstractDebugStreamFunction {
	return NewAbstractDebugStreamFunctionFromPointer(cthis)
}

//  body block end
