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
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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

// /usr/include/qt/QtCore/qjsonvalue.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRefPtr(QJsonArray *, int)
func NewQJsonValueRefPtr(array *QJsonArray /*777 QJsonArray **/, idx int) *QJsonValueRefPtr {
	var convArg0 = array.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP10QJsonArrayi", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefPtrFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRefPtr)
	return gothis
}

// /usr/include/qt/QtCore/qjsonvalue.h:239
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QJsonValueRefPtr(QJsonObject *, int)
func NewQJsonValueRefPtr_1(object *QJsonObject /*777 QJsonObject **/, idx int) *QJsonValueRefPtr {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJsonValueRefPtrC2EP11QJsonObjecti", qtrt.FFI_TYPE_POINTER, convArg0, idx)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonValueRefPtrFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonValueRefPtr)
	return gothis
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
		qtrt.KeepMe()
	}
}

//  keep block end
