package qtnetwork

// /usr/include/qt/QtNetwork/qhttpmultipart.h
// #include <qhttpmultipart.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QHttpMultiPart struct {
	*qtcore.QObject
}

func (this *QHttpMultiPart) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QHttpMultiPart) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQHttpMultiPartFromPointer(cthis unsafe.Pointer) *QHttpMultiPart {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QHttpMultiPart{bcthis0}
}
func (*QHttpMultiPart) NewFromPointer(cthis unsafe.Pointer) *QHttpMultiPart {
	return NewQHttpMultiPartFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QHttpMultiPart) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QHttpMultiPart10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(QObject *)
func NewQHttpMultiPart(parent *qtcore.QObject /*777 QObject **/) *QHttpMultiPart {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QHttpMultiPartC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:102
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(enum QHttpMultiPart::ContentType, QObject *)
func NewQHttpMultiPart_1(contentType int, parent *qtcore.QObject /*777 QObject **/) *QHttpMultiPart {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QHttpMultiPartC2ENS_11ContentTypeEP7QObject", ffiqt.FFI_TYPE_POINTER, contentType, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHttpMultiPart()
func DeleteQHttpMultiPart(*QHttpMultiPart) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QHttpMultiPartD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentType(enum QHttpMultiPart::ContentType)
func (this *QHttpMultiPart) SetContentType(contentType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QHttpMultiPart14setContentTypeENS_11ContentTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), contentType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray boundary()
func (this *QHttpMultiPart) Boundary() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QHttpMultiPart8boundaryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundary(const QByteArray &)
func (this *QHttpMultiPart) SetBoundary(boundary *qtcore.QByteArray) {
	var convArg0 = boundary.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QHttpMultiPart11setBoundaryERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QHttpMultiPart__ContentType = int

const QHttpMultiPart__MixedType QHttpMultiPart__ContentType = 0
const QHttpMultiPart__RelatedType QHttpMultiPart__ContentType = 1
const QHttpMultiPart__FormDataType QHttpMultiPart__ContentType = 2
const QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = 3

//  body block end
