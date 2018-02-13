package qtcore

// /usr/include/qt/QtCore/qshareddata.h
// #include <qshareddata.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSharedData struct {
	*qtrt.CObject
}
type QSharedData_ITF interface {
	QSharedData_PTR() *QSharedData
}

func (ptr *QSharedData) QSharedData_PTR() *QSharedData { return ptr }

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
	qtrt.ErrPrint(err, rv)
	gothis := NewQSharedDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSharedData)
	return gothis
}

func DeleteQSharedData(this *QSharedData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSharedDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
}

//  keep block end
