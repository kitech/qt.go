package qtnetwork

// /usr/include/qt/QtNetwork/qhttpmultipart.h
// #include <qhttpmultipart.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QHttpPart struct {
	*qtrt.CObject
}

func (this *QHttpPart) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHttpPart) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQHttpPartFromPointer(cthis unsafe.Pointer) *QHttpPart {
	return &QHttpPart{&qtrt.CObject{cthis}}
}
func (*QHttpPart) NewFromPointer(cthis unsafe.Pointer) *QHttpPart {
	return NewQHttpPartFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:58
// index:0
// Public
// void QHttpPart()
func NewQHttpPart() *QHttpPart {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPartC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQHttpPartFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:60
// index:0
// Public
// void ~QHttpPart()
func DeleteQHttpPart(*QHttpPart) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPartD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:66
// index:0
// Public inline
// void swap(class QHttpPart &)
func (this *QHttpPart) Swap(other *QHttpPart) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPart4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:72
// index:0
// Public
// void setHeader(class QNetworkRequest::KnownHeaders, const class QVariant &)
func (this *QHttpPart) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPart9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:73
// index:0
// Public
// void setRawHeader(const class QByteArray &, const class QByteArray &)
func (this *QHttpPart) SetRawHeader(headerName *qtcore.QByteArray, headerValue *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = headerValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPart12setRawHeaderERK10QByteArrayS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:75
// index:0
// Public
// void setBody(const class QByteArray &)
func (this *QHttpPart) SetBody(body *qtcore.QByteArray) {
	var convArg0 = body.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPart7setBodyERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:76
// index:0
// Public
// void setBodyDevice(class QIODevice *)
func (this *QHttpPart) SetBodyDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHttpPart13setBodyDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
