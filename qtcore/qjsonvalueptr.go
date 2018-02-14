package qtcore

// /usr/include/qt/QtCore/qjsonvalue.h
// #include <qjsonvalue.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QJsonValuePtr struct {
	*qtrt.CObject
}
type QJsonValuePtr_ITF interface {
	QJsonValuePtr_PTR() *QJsonValuePtr
}

func (ptr *QJsonValuePtr) QJsonValuePtr_PTR() *QJsonValuePtr { return ptr }

func (this *QJsonValuePtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValuePtr) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonValuePtrFromPointer(cthis unsafe.Pointer) *QJsonValuePtr {
	return &QJsonValuePtr{&qtrt.CObject{cthis}}
}
func (*QJsonValuePtr) NewFromPointer(cthis unsafe.Pointer) *QJsonValuePtr {
	return NewQJsonValuePtrFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValuePtr(const QJsonValue &)
func NewQJsonValuePtr(val QJsonValue_ITF) *QJsonValuePtr {
	var convArg0 unsafe.Pointer
	if val != nil && val.QJsonValue_PTR() != nil {
		convArg0 = val.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValuePtrC2ERK10QJsonValue", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValuePtrFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValuePtr)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QJsonValue & operator*()
func (this *QJsonValuePtr) Operator_mul() *QJsonValue {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValuePtrdeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QJsonValue * operator->()
func (this *QJsonValuePtr) Operator_minus_greater() *QJsonValue /*777 QJsonValue **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValuePtrptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQJsonValuePtr(this *QJsonValuePtr) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QJsonValuePtrD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
