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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qthreadstorage.h:54
// index:0
// void ~QThreadStorageData()
func DeleteQThreadStorageData(*QThreadStorageData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadstorage.h:56
// index:0
// void ** get()
func (this *QThreadStorageData) Get() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QThreadStorageData3getEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadstorage.h:57
// index:0
// void ** set(void *)
func (this *QThreadStorageData) Set(p unsafe.Pointer) {
	// 0: (, void * p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData3setEPv", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadstorage.h:59
// index:0
// static
// void finish(void **)
func (this *QThreadStorageData) Finish(arg0 unsafe.Pointer) {
	// 0: (void ** arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData6finishEPPv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThreadStorageData_Finish(arg0 unsafe.Pointer) {
	// 0: (void ** arg0), (arg0)
	var nilthis *QThreadStorageData
	nilthis.Finish(arg0)
}

//  body block end
