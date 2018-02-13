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
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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

// /usr/include/qt/QtNetwork/qhttpmultipart.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHttpPart &)
func (this *QHttpPart) Swap(other *QHttpPart) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QHttpPart) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QHttpPart) SetRawHeader(headerName *qtcore.QByteArray, headerValue *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = headerValue.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBody(const QByteArray &)
func (this *QHttpPart) SetBody(body *qtcore.QByteArray) {
	var convArg0 = body.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHttpPart7setBodyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBodyDevice(QIODevice *)
func (this *QHttpPart) SetBodyDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
