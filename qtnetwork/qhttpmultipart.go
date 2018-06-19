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
// extern C begin: 11
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

/*

 */
type QHttpMultiPart struct {
	*qtcore.QObject
}
type QHttpMultiPart_ITF interface {
	qtcore.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QHttpMultiPart) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QHttpMultiPart10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(QObject *)

/*
Constructs a QHttpMultiPart with content type MixedType and sets parent as the parent object.

See also QHttpMultiPart::ContentType.
*/
func NewQHttpMultiPart(parent qtcore.QObject_ITF /*777 QObject **/) *QHttpMultiPart {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPartC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHttpMultiPart")
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(QObject *)

/*
Constructs a QHttpMultiPart with content type MixedType and sets parent as the parent object.

See also QHttpMultiPart::ContentType.
*/
func NewQHttpMultiPart__() *QHttpMultiPart {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPartC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHttpMultiPart")
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:102
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(QHttpMultiPart::ContentType, QObject *)

/*
Constructs a QHttpMultiPart with content type MixedType and sets parent as the parent object.

See also QHttpMultiPart::ContentType.
*/
func NewQHttpMultiPart_1(contentType int, parent qtcore.QObject_ITF /*777 QObject **/) *QHttpMultiPart {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPartC2ENS_11ContentTypeEP7QObject", qtrt.FFI_TYPE_POINTER, contentType, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHttpMultiPart")
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:102
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHttpMultiPart(QHttpMultiPart::ContentType, QObject *)

/*
Constructs a QHttpMultiPart with content type MixedType and sets parent as the parent object.

See also QHttpMultiPart::ContentType.
*/
func NewQHttpMultiPart_1_(contentType int) *QHttpMultiPart {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPartC2ENS_11ContentTypeEP7QObject", qtrt.FFI_TYPE_POINTER, contentType, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHttpMultiPartFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHttpMultiPart")
	return gothis
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHttpMultiPart()

/*

 */
func DeleteQHttpMultiPart(this *QHttpMultiPart) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPartD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentType(QHttpMultiPart::ContentType)

/*
Sets the content type to contentType. The content type will be used in the HTTP header section when sending the multipart message via QNetworkAccessManager::post(). In case you want to use a multipart subtype not contained in QHttpMultiPart::ContentType, you can add the "Content-Type" header field to the QNetworkRequest by hand, and then use this request together with the multipart message for posting.

See also QHttpMultiPart::ContentType and QNetworkAccessManager::post().
*/
func (this *QHttpMultiPart) SetContentType(contentType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPart14setContentTypeENS_11ContentTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), contentType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray boundary() const

/*
returns the boundary.

See also setBoundary().
*/
func (this *QHttpMultiPart) Boundary() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QHttpMultiPart8boundaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qhttpmultipart.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundary(const QByteArray &)

/*
Sets the boundary to boundary.

Usually, you do not need to generate a boundary yourself; upon construction the boundary is initiated with the string "boundary_.oOo._" followed by random characters, and provides enough uniqueness to make sure it does not occur inside the parts itself.

See also boundary().
*/
func (this *QHttpMultiPart) SetBoundary(boundary qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if boundary != nil && boundary.QByteArray_PTR() != nil {
		convArg0 = boundary.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QHttpMultiPart11setBoundaryERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
List of known content types for a multipart subtype as described in RFC 2046 and others.



See also setContentType().

*/
type QHttpMultiPart__ContentType = int

//
const QHttpMultiPart__MixedType QHttpMultiPart__ContentType = 0

//
const QHttpMultiPart__RelatedType QHttpMultiPart__ContentType = 1

//
const QHttpMultiPart__FormDataType QHttpMultiPart__ContentType = 2

//
const QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = 3

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
