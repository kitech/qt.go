//  header block begin
// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QObjectData struct {
	*qtrt.CObject
}

func (this *QObjectData) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQObjectDataFromPointer(cthis unsafe.Pointer) *QObjectData {
	return &QObjectData{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qobject.h:97
// index:0
// Public pure virtual
// void ~QObjectData()
func DeleteQObjectData(*QObjectData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QObjectDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:112
// index:0
// Public
// QMetaObject * dynamicMetaObject()
func (this *QObjectData) DynamicMetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QObjectData17dynamicMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
