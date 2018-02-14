package qtnetwork

// /usr/include/qt/QtNetwork/qhttpmultipart.h
// #include <qhttpmultipart.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QHttpPart struct {
	*qtrt.CObject
}
type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (ptr *QHttpPart) QHttpPart_PTR() *QHttpPart { return ptr }

func (this *QHttpPart) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHttpPart) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHttpPartFromPointer(cthis unsafe.Pointer) *QHttpPart {
	return &QHttpPart{&qtrt.CObject{cthis}}
}
func (*QHttpPart) NewFromPointer(cthis unsafe.Pointer) *QHttpPart {
	return NewQHttpPartFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHttpPart()
func NewQHttpPart() *QHttpPart {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPartC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHttpPartFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHttpPart)
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHttpPart()
func DeleteQHttpPart(this *QHttpPart) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPartD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QHttpPart & operator=(QHttpPart &&)
func (this *QHttpPart) Operator_equal(other unsafe.Pointer /*333*/) *QHttpPart {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPartaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHttpPartFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHttpPart)
	return rv2
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:64
// index:1
// Public Visibility=Default Availability=Available
// [8] QHttpPart & operator=(const QHttpPart &)
func (this *QHttpPart) Operator_equal_1(other QHttpPart_ITF) *QHttpPart {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHttpPart_PTR() != nil {
		convArg0 = other.QHttpPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPartaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHttpPartFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHttpPart)
	return rv2
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHttpPart &)
func (this *QHttpPart) Swap(other QHttpPart_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHttpPart_PTR() != nil {
		convArg0 = other.QHttpPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QHttpPart &)
func (this *QHttpPart) Operator_equal_equal(other QHttpPart_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHttpPart_PTR() != nil {
		convArg0 = other.QHttpPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHttpParteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QHttpPart &)
func (this *QHttpPart) Operator_not_equal(other QHttpPart_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHttpPart_PTR() != nil {
		convArg0 = other.QHttpPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHttpPartneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QHttpPart) SetHeader(header int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QHttpPart) SetRawHeader(headerName qtcore.QByteArray_ITF, headerValue qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if headerValue != nil && headerValue.QByteArray_PTR() != nil {
		convArg1 = headerValue.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBody(const QByteArray &)
func (this *QHttpPart) SetBody(body qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if body != nil && body.QByteArray_PTR() != nil {
		convArg0 = body.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart7setBodyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBodyDevice(QIODevice *)
func (this *QHttpPart) SetBodyDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart13setBodyDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
