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
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QJsonValueRefPtr struct {
	*qtrt.CObject
}
type QJsonValueRefPtr_ITF interface {
	QJsonValueRefPtr_PTR() *QJsonValueRefPtr
}

func (ptr *QJsonValueRefPtr) QJsonValueRefPtr_PTR() *QJsonValueRefPtr { return ptr }

func (this *QJsonValueRefPtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonValueRefPtr) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonValueRefPtrFromPointer(cthis unsafe.Pointer) *QJsonValueRefPtr {
	return &QJsonValueRefPtr{&qtrt.CObject{cthis}}
}
func (*QJsonValueRefPtr) NewFromPointer(cthis unsafe.Pointer) *QJsonValueRefPtr {
	return NewQJsonValueRefPtrFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonvalue.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRefPtr(QJsonArray *, int)

/*

 */
func (*QJsonValueRefPtr) NewForInherit(array QJsonArray_ITF /*777 QJsonArray **/, idx int) *QJsonValueRefPtr {
	return NewQJsonValueRefPtr(array, idx)
}
func NewQJsonValueRefPtr(array QJsonArray_ITF /*777 QJsonArray **/, idx int) *QJsonValueRefPtr {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP10QJsonArrayi", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefPtrFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRefPtr)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:240
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRefPtr(QJsonObject *, int)

/*

 */
func (*QJsonValueRefPtr) NewForInherit1(object QJsonObject_ITF /*777 QJsonObject **/, idx int) *QJsonValueRefPtr {
	return NewQJsonValueRefPtr1(object, idx)
}
func NewQJsonValueRefPtr1(object QJsonObject_ITF /*777 QJsonObject **/, idx int) *QJsonValueRefPtr {
	var convArg0 unsafe.Pointer
	if object != nil && object.QJsonObject_PTR() != nil {
		convArg0 = object.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP11QJsonObjecti", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefPtrFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRefPtr)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:243
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonValueRef & operator*()

/*

 */
func (this *QJsonValueRefPtr) Operator_mul() *QJsonValueRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrdeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonvalue.h:244
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QJsonValueRef * operator->()

/*

 */
func (this *QJsonValueRefPtr) Operator_minus_greater() *QJsonValueRef /*777 QJsonValueRef **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQJsonValueRefPtr(this *QJsonValueRefPtr) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
