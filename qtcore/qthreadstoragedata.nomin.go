//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qthreadstorage.h
// #include <qthreadstorage.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qthreadstorage.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QThreadStorageData()

/*

 */
func DeleteQThreadStorageData(this *QThreadStorageData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QThreadStorageDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 4)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthreadstorage.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** get() const

/*

 */
func (this *QThreadStorageData) Get() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QThreadStorageData3getEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qthreadstorage.h:57
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** set(void *)

/*

 */
func (this *QThreadStorageData) Set(p unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QThreadStorageData3setEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), p)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qthreadstorage.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void finish(void **)

/*

 */
func (this *QThreadStorageData) Finish(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QThreadStorageData6finishEPPv", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QThreadStorageData_Finish(arg0 unsafe.Pointer /*666*/) {
	var nilthis *QThreadStorageData
	nilthis.Finish(arg0)
}

//  body block end

//  keep block begin

func init_unused_10600() {
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
