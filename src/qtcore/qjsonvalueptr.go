//  header block begin
// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QJsonValuePtr struct {
	*qtrt.CObject
}

func (this *QJsonValuePtr) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qjsonvalue.h:226
// index:0
// inline
// void QJsonValuePtr(const class QJsonValue &)
func NewQJsonValuePtr(val unsafe.Pointer) *QJsonValuePtr {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QJsonValuePtrC2ERK10QJsonValue", ffiqt.FFI_TYPE_VOID, cthis, val)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonValuePtrFromPointer(cthis)
	return gothis
}
func NewQJsonValuePtrFromPointer(cthis unsafe.Pointer) *QJsonValuePtr {
	return &QJsonValuePtr{&qtrt.CObject{cthis}}
}

//  body block end
