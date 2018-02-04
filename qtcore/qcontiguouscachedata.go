package qtcore

// /usr/include/qt/QtCore/qcontiguouscache.h
// #include <qcontiguouscache.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QContiguousCacheData struct {
	*qtrt.CObject
}

func (this *QContiguousCacheData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QContiguousCacheData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQContiguousCacheDataFromPointer(cthis unsafe.Pointer) *QContiguousCacheData {
	return &QContiguousCacheData{&qtrt.CObject{cthis}}
}
func (*QContiguousCacheData) NewFromPointer(cthis unsafe.Pointer) *QContiguousCacheData {
	return NewQContiguousCacheDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcontiguouscache.h:67
// index:0
// Public static Visibility=Default Availability=Available
// [8] QContiguousCacheData * allocateData(int, int)
func (this *QContiguousCacheData) AllocateData(size int, alignment int) *QContiguousCacheData /*777 QContiguousCacheData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QContiguousCacheData12allocateDataEii", qtrt.FFI_TYPE_POINTER, size, alignment)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQContiguousCacheDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QContiguousCacheData_AllocateData(size int, alignment int) *QContiguousCacheData /*777 QContiguousCacheData **/ {
	var nilthis *QContiguousCacheData
	rv := nilthis.AllocateData(size, alignment)
	return rv
}

func DeleteQContiguousCacheData(this *QContiguousCacheData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QContiguousCacheDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
