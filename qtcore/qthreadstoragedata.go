package qtcore

// /usr/include/qt/QtCore/qthreadstorage.h
// #include <qthreadstorage.h>
// #include <QtCore>

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

type QThreadStorageData struct {
	*qtrt.CObject
}

func (this *QThreadStorageData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QThreadStorageData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQThreadStorageDataFromPointer(cthis unsafe.Pointer) *QThreadStorageData {
	return &QThreadStorageData{&qtrt.CObject{cthis}}
}
func (*QThreadStorageData) NewFromPointer(cthis unsafe.Pointer) *QThreadStorageData {
	return NewQThreadStorageDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qthreadstorage.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QThreadStorageData()
func DeleteQThreadStorageData(this *QThreadStorageData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageDataD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthreadstorage.h:56
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** get()
func (this *QThreadStorageData) Get() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QThreadStorageData3getEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qthreadstorage.h:57
// index:0
// Public Visibility=Default Availability=Available
// [8] void ** set(void *)
func (this *QThreadStorageData) Set(p unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData3setEPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qthreadstorage.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void finish(void **)
func (this *QThreadStorageData) Finish(arg0 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QThreadStorageData6finishEPPv", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThreadStorageData_Finish(arg0 unsafe.Pointer /*666*/) {
	var nilthis *QThreadStorageData
	nilthis.Finish(arg0)
}

//  body block end
