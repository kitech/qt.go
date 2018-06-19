package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h
// #include <qwebenginehttprequest.h>
// #include <QtWebEngineCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

/*

 */
type QWebEngineHttpRequest struct {
	*qtrt.CObject
}
type QWebEngineHttpRequest_ITF interface {
	QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest
}

func (ptr *QWebEngineHttpRequest) QWebEngineHttpRequest_PTR() *QWebEngineHttpRequest { return ptr }

func (this *QWebEngineHttpRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineHttpRequest) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineHttpRequestFromPointer(cthis unsafe.Pointer) *QWebEngineHttpRequest {
	return &QWebEngineHttpRequest{&qtrt.CObject{cthis}}
}
func (*QWebEngineHttpRequest) NewFromPointer(cthis unsafe.Pointer) *QWebEngineHttpRequest {
	return NewQWebEngineHttpRequestFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineHttpRequest(const QUrl &, const QWebEngineHttpRequest::Method &)

/*

 */
func NewQWebEngineHttpRequest(url qtcore.QUrl_ITF, method int) *QWebEngineHttpRequest {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestC2ERK4QUrlRKNS_6MethodE", qtrt.FFI_TYPE_POINTER, convArg0, &method)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineHttpRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineHttpRequest)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineHttpRequest(const QUrl &, const QWebEngineHttpRequest::Method &)

/*

 */
func NewQWebEngineHttpRequest__() *QWebEngineHttpRequest {
	// arg: 0, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg0 = qtcore.NewQUrl()
	// arg: 1, const QWebEngineHttpRequest::Method &=LValueReference, QWebEngineHttpRequest::Method=Enum, , Invalid
	method := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestC2ERK4QUrlRKNS_6MethodE", qtrt.FFI_TYPE_POINTER, convArg0, &method)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineHttpRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineHttpRequest)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineHttpRequest(const QUrl &, const QWebEngineHttpRequest::Method &)

/*

 */
func NewQWebEngineHttpRequest__1(url qtcore.QUrl_ITF) *QWebEngineHttpRequest {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	// arg: 1, const QWebEngineHttpRequest::Method &=LValueReference, QWebEngineHttpRequest::Method=Enum, , Invalid
	method := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestC2ERK4QUrlRKNS_6MethodE", qtrt.FFI_TYPE_POINTER, convArg0, &method)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineHttpRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineHttpRequest)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineHttpRequest()

/*

 */
func DeleteQWebEngineHttpRequest(this *QWebEngineHttpRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWebEngineHttpRequest & operator=(QWebEngineHttpRequest &&)

/*

 */
func (this *QWebEngineHttpRequest) Operator_equal(other unsafe.Pointer /*333*/) *QWebEngineHttpRequest {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHttpRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHttpRequest)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:71
// index:1
// Public Visibility=Default Availability=Available
// [8] QWebEngineHttpRequest & operator=(const QWebEngineHttpRequest &)

/*

 */
func (this *QWebEngineHttpRequest) Operator_equal_1(other QWebEngineHttpRequest_ITF) *QWebEngineHttpRequest {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = other.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequestaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHttpRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHttpRequest)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineHttpRequest &)

/*

 */
func (this *QWebEngineHttpRequest) Swap(other QWebEngineHttpRequest_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = other.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QWebEngineHttpRequest &) const

/*

 */
func (this *QWebEngineHttpRequest) Operator_equal_equal(other QWebEngineHttpRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = other.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequesteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QWebEngineHttpRequest &) const

/*

 */
func (this *QWebEngineHttpRequest) Operator_not_equal(other QWebEngineHttpRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = other.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequestneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineHttpRequest::Method method() const

/*

 */
func (this *QWebEngineHttpRequest) Method() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequest6methodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMethod(QWebEngineHttpRequest::Method)

/*

 */
func (this *QWebEngineHttpRequest) SetMethod(method int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest9setMethodENS_6MethodE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), method)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEngineHttpRequest) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequest3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)

/*

 */
func (this *QWebEngineHttpRequest) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray postData() const

/*

 */
func (this *QWebEngineHttpRequest) PostData() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequest8postDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPostData(const QByteArray &)

/*

 */
func (this *QWebEngineHttpRequest) SetPostData(postData qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if postData != nil && postData.QByteArray_PTR() != nil {
		convArg0 = postData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest11setPostDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasHeader(const QByteArray &) const

/*

 */
func (this *QWebEngineHttpRequest) HasHeader(headerName qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequest9hasHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray header(const QByteArray &) const

/*

 */
func (this *QWebEngineHttpRequest) Header(headerName qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHttpRequest6headerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(const QByteArray &, const QByteArray &)

/*

 */
func (this *QWebEngineHttpRequest) SetHeader(headerName qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest9setHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetHeader(const QByteArray &)

/*

 */
func (this *QWebEngineHttpRequest) UnsetHeader(headerName qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHttpRequest11unsetHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QWebEngineHttpRequest__Method = int

//
const QWebEngineHttpRequest__Get QWebEngineHttpRequest__Method = 0

//
const QWebEngineHttpRequest__Post QWebEngineHttpRequest__Method = 1

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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  keep block end
