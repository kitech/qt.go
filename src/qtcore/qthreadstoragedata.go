//  header block begin
// /usr/include/qt/QtCore/qthreadstorage.h
// #include <qthreadstorage.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QThreadStorageData struct {
	*qtrt.CObject
}

func (this *QThreadStorageData) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQThreadStorageDataFromPointer(cthis unsafe.Pointer) *QThreadStorageData {
	return &QThreadStorageData{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qthreadstorage.h:54
// index:0
// Public
// void ~QThreadStorageData()
func DeleteQThreadStorageData(*QThreadStorageData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadstorage.h:56
// index:0
// Public
// void ** get()
func (this *QThreadStorageData) Get() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QThreadStorageData3getEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadstorage.h:57
// index:0
// Public
// void ** set(void *)
func (this *QThreadStorageData) Set(p unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData3setEPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadstorage.h:59
// index:0
// Public static
// void finish(void **)
func (this *QThreadStorageData) Finish(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData6finishEPPv", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThreadStorageData_Finish(arg0 unsafe.Pointer) {
	var nilthis *QThreadStorageData
	nilthis.Finish(arg0)
}

//  body block end
