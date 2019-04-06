//  header block begin

// +build !minimal

package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h
// #include <qnetworkaccessmanager.h>
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

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:144
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, QHttpMultiPart *)

/*
Sends an HTTP POST request to the destination specified by request and returns a new QNetworkReply object opened for reading that will contain the reply sent by the server. The contents of the data device will be uploaded to the server.

data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: Sending a POST request on protocols other than HTTP and HTTPS is undefined and will probably fail.

See also get(), put(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:145
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, QHttpMultiPart *)

/*
Uploads the contents of data to the destination request and returns a new QNetworkReply object that will be open for reply.

data must be opened for reading when this function is called and must remain valid until the finished() signal is emitted for this reply.

Whether anything will be available for reading from the returned object is protocol dependent. For HTTP, the server may send a small HTML page indicating the upload was successful (or not). Other protocols will probably have content in their replies.

Note: For HTTP, this request will send a PUT request, which most servers do not allow. Form upload mechanisms, including that of uploading files through HTML forms, use the POST mechanism.

See also get(), post(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:146
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QHttpMultiPart *)

/*
Sends a custom request to the server identified by the URL of request.

It is the user's responsibility to send a verb to the server that is valid according to the HTTP specification.

This method provides means to send verbs other than the common ones provided via get() or post() etc., for instance sending an HTTP OPTIONS command.

If data is not empty, the contents of the data device will be uploaded to the server; in that case, data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: This feature is currently available for HTTP(S) only.

This function was introduced in  Qt 4.7.

See also get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) SendCustomRequest2(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if verb != nil && verb.QByteArray_PTR() != nil {
		convArg1 = verb.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg2 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11434() {
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
