package qtcore

// /usr/include/qt/QtCore/qshareddata.h
// #include <qshareddata.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QSharedData struct {
	*qtrt.CObject
}

func (this *QSharedData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSharedData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSharedDataFromPointer(cthis unsafe.Pointer) *QSharedData {
	return &QSharedData{&qtrt.CObject{cthis}}
}
func (*QSharedData) NewFromPointer(cthis unsafe.Pointer) *QSharedData {
	return NewQSharedDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qshareddata.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSharedData()
func NewQSharedData() *QSharedData {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSharedDataC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSharedDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSharedData)
	return gothis
}

func DeleteQSharedData(this *QSharedData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSharedDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
