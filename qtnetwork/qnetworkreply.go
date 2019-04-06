package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkreply.h
// #include <qnetworkreply.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

// long long writeData(const char *, qint64)
func (this *QNetworkReply) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

// void setOperation(QNetworkAccessManager::Operation)
func (this *QNetworkReply) InheritSetOperation(f func(operation int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setOperation", f)
}

// void setRequest(const QNetworkRequest &)
func (this *QNetworkReply) InheritSetRequest(f func(request *QNetworkRequest) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRequest", f)
}

// void setError(QNetworkReply::NetworkError, const QString &)
func (this *QNetworkReply) InheritSetError(f func(errorCode int, errorString string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setError", f)
}

// void setFinished(bool)
func (this *QNetworkReply) InheritSetFinished(f func(arg0 bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setFinished", f)
}

// void setUrl(const QUrl &)
func (this *QNetworkReply) InheritSetUrl(f func(url *qtcore.QUrl) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setUrl", f)
}

// void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QNetworkReply) InheritSetHeader(f func(header int, value *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setHeader", f)
}

// void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QNetworkReply) InheritSetRawHeader(f func(headerName *qtcore.QByteArray, value *qtcore.QByteArray) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRawHeader", f)
}

// void setAttribute(QNetworkRequest::Attribute, const QVariant &)
func (this *QNetworkReply) InheritSetAttribute(f func(code int, value *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setAttribute", f)
}

// void sslConfigurationImplementation(QSslConfiguration &)
func (this *QNetworkReply) InheritSslConfigurationImplementation(f func(arg0 *QSslConfiguration) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sslConfigurationImplementation", f)
}

// void setSslConfigurationImplementation(const QSslConfiguration &)
func (this *QNetworkReply) InheritSetSslConfigurationImplementation(f func(arg0 *QSslConfiguration) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSslConfigurationImplementation", f)
}

/*

 */
type QNetworkReply struct {
	*qtcore.QIODevice
}
type QNetworkReply_ITF interface {
	qtcore.QIODevice_ITF
	QNetworkReply_PTR() *QNetworkReply
}

func (ptr *QNetworkReply) QNetworkReply_PTR() *QNetworkReply { return ptr }

func (this *QNetworkReply) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QNetworkReply) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = qtcore.NewQIODeviceFromPointer(cthis)
}
func NewQNetworkReplyFromPointer(cthis unsafe.Pointer) *QNetworkReply {
	bcthis0 := qtcore.NewQIODeviceFromPointer(cthis)
	return &QNetworkReply{bcthis0}
}
func (*QNetworkReply) NewFromPointer(cthis unsafe.Pointer) *QNetworkReply {
	return NewQNetworkReplyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QNetworkReply) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkReply()

/*

 */
func DeleteQNetworkReply(this *QNetworkReply) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReplyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().

Closes this device for reading. Unread data is discarded, but the network resources are not discarded until they are finished. In particular, if any upload is in progress, it will continue until it is done.

The finished() signal is emitted when all operations are over and the network resources are freed.

See also abort() and finished().
*/
func (this *QNetworkReply) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const

/*

 */
func (this *QNetworkReply) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize() const

/*
Returns the size of the read buffer, in bytes.

See also setReadBufferSize().
*/
func (this *QNetworkReply) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)

/*
Sets the size of the read buffer to be size bytes. The read buffer is the buffer that holds data that is being downloaded off the network, before it is read with QIODevice::read(). Setting the buffer size to 0 will make the buffer unlimited in size.

QNetworkReply will try to stop reading from the network once this buffer is full (i.e., bytesAvailable() returns size or more), thus causing the download to throttle down as well. If the buffer is not limited in size, QNetworkReply will try to download as fast as possible from the network.

Unlike QAbstractSocket::setReadBufferSize(), QNetworkReply cannot guarantee precision in the read buffer size. That is, bytesAvailable() can return more than size.

See also readBufferSize().
*/
func (this *QNetworkReply) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkAccessManager * manager() const

/*
Returns the QNetworkAccessManager that was used to create this QNetworkReply object. Initially, it is also the parent object.
*/
func (this *QNetworkReply) Manager() *QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply7managerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkAccessManager::Operation operation() const

/*
Returns the operation that was posted for this reply.

See also setOperation().
*/
func (this *QNetworkReply) Operation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9operationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkRequest request() const

/*
Returns the request that was posted for this reply. In special, note that the URL for the request may be different than that of the reply.

See also QNetworkRequest::url(), url(), and setRequest().
*/
func (this *QNetworkReply) Request() *QNetworkRequest /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply7requestEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkRequest)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkReply::NetworkError error() const

/*
Returns the error that was found during the processing of this request. If no error was found, returns NoError.

See also setError().
*/
func (this *QNetworkReply) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:159
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QNetworkReply::NetworkError)

/*
Returns the error that was found during the processing of this request. If no error was found, returns NoError.

See also setError().
*/
func (this *QNetworkReply) Error1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5errorENS_12NetworkErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const

/*
Returns true when the reply has finished or was aborted.

This function was introduced in  Qt 4.6.

See also isRunning().
*/
func (this *QNetworkReply) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const

/*
Returns true when the request is still processing and the reply has not finished or was aborted yet.

This function was introduced in  Qt 4.6.

See also isFinished().
*/
func (this *QNetworkReply) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*
Returns the URL of the content downloaded or uploaded. Note that the URL may be different from that of the original request. If the QNetworkRequest::FollowRedirectsAttribute was set in the request, then this function returns the current url that the network API is accessing, i.e the url emitted in the QNetworkReply::redirected signal.

See also request(), setUrl(), QNetworkRequest::url(), and redirected().
*/
func (this *QNetworkReply) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:133
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant header(QNetworkRequest::KnownHeaders) const

/*
Returns the value of the known header header, if that header was sent by the remote server. If the header was not sent, returns an invalid QVariant.

See also rawHeader(), setHeader(), and QNetworkRequest::header().
*/
func (this *QNetworkReply) Header(header int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply6headerEN15QNetworkRequest12KnownHeadersE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasRawHeader(const QByteArray &) const

/*
Returns true if the raw header of name headerName was sent by the remote server

See also rawHeader().
*/
func (this *QNetworkReply) HasRawHeader(headerName qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply12hasRawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rawHeader(const QByteArray &) const

/*
Returns the raw contents of the header headerName as sent by the remote server. If there is no such header, returns an empty byte array, which may be indistinguishable from an empty header. Use hasRawHeader() to verify if the server sent such header field.

See also setRawHeader(), hasRawHeader(), and header().
*/
func (this *QNetworkReply) RawHeader(headerName qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9rawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:144
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant attribute(QNetworkRequest::Attribute) const

/*
Returns the attribute associated with the code code. If the attribute has not been set, it returns an invalid QVariant (type QMetaType::UnknownType).

You can expect the default values listed in QNetworkRequest::Attribute to be applied to the values returned by this function.

See also setAttribute() and QNetworkRequest::Attribute.
*/
func (this *QNetworkReply) Attribute(code int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9attributeEN15QNetworkRequest9AttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration() const

/*
Returns the SSL configuration and state associated with this reply, if SSL was used. It will contain the remote server's certificate, its certificate chain leading to the Certificate Authority as well as the encryption ciphers in use.

The peer's certificate and its certificate chain will be known by the time sslErrors() is emitted, if it's emitted.

See also setSslConfiguration().
*/
func (this *QNetworkReply) SslConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply16sslConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslConfiguration(const QSslConfiguration &)

/*
Sets the SSL configuration for the network connection associated with this request, if possible, to be that of config.

See also sslConfiguration().
*/
func (this *QNetworkReply) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:153
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void abort()

/*
Aborts the operation immediately and close down any network connections still open. Uploads still in progress are also aborted.

The finished() signal will also be emitted.

See also close() and finished().
*/
func (this *QNetworkReply) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ignoreSslErrors()

/*
If this function is called, SSL errors related to network connection will be ignored, including certificate validation errors.

Warning: Be sure to always let the user inspect the errors reported by the sslErrors() signal, and only call this method upon confirmation from the user that proceeding is ok. If there are unexpected errors, the reply should be aborted. Calling this method without inspecting the actual errors will most likely pose a security risk for your application. Use it with great care!

This function can be called from the slot connected to the sslErrors() signal, which indicates which errors were found.

Note: If HTTP Strict Transport Security is enabled for QNetworkAccessManager, this function has no effect.

See also sslConfiguration(), sslErrors(), and QSslSocket::ignoreSslErrors().
*/
func (this *QNetworkReply) IgnoreSslErrors() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15ignoreSslErrorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()

/*
This signal is emitted whenever the metadata in this reply changes. metadata is any information that is not the content (data) itself, including the network headers. In the majority of cases, the metadata will be known fully by the time the first byte of data is received. However, it is possible to receive updates of headers or other metadata during the processing of the data.

See also header(), rawHeaderList(), rawHeader(), and hasRawHeader().
*/
func (this *QNetworkReply) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
This signal is emitted when the reply has finished processing. After this signal is emitted, there will be no more updates to the reply's data or metadata.

Unless close() or abort() have been called, the reply will be still be opened for reading, so the data can be retrieved by calls to read() or readAll(). In particular, if no calls to read() were made as a result of readyRead(), a call to readAll() will retrieve the full contents in a QByteArray.

This signal is emitted in tandem with QNetworkAccessManager::finished() where that signal's reply parameter is this object.

Note: Do not delete the object in the slot connected to this signal. Use deleteLater().

You can also use isFinished() to check if a QNetworkReply has finished even before you receive the finished() signal.

See also setFinished(), QNetworkAccessManager::finished(), and isFinished().
*/
func (this *QNetworkReply) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted()

/*
This signal is emitted when an SSL/TLS session has successfully completed the initial handshake. At this point, no user data has been transmitted. The signal can be used to perform additional checks on the certificate chain, for example to notify users when the certificate for a website has changed. If the reply does not match the expected criteria then it should be aborted by calling QNetworkReply::abort() by a slot connected to this signal. The SSL configuration in use can be inspected using the QNetworkReply::sslConfiguration() method.

Internally, QNetworkAccessManager may open multiple connections to a server, in order to allow it process requests in parallel. These connections may be reused, which means that the encrypted() signal would not be emitted. This means that you are only guaranteed to receive this signal for the first connection to a site in the lifespan of the QNetworkAccessManager.

This function was introduced in  Qt 5.1.

See also QSslSocket::encrypted() and QNetworkAccessManager::encrypted().
*/
func (this *QNetworkReply) Encrypted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9encryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)

/*
This signal is emitted if the SSL/TLS handshake negotiates a PSK ciphersuite, and therefore a PSK authentication is then required.

When using PSK, the client must send to the server a valid identity and a valid pre shared key, in order for the SSL handshake to continue. Applications can provide this information in a slot connected to this signal, by filling in the passed authenticator object according to their needs.

Note: Ignoring this signal, or failing to provide the required credentials, will cause the handshake to fail, and therefore the connection to be aborted.

Note: The authenticator object is owned by the reply and must not be deleted by the application.

This function was introduced in  Qt 5.5.

See also QSslPreSharedKeyAuthenticator.
*/
func (this *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirected(const QUrl &)

/*
This signal is emitted if the QNetworkRequest::FollowRedirectsAttribute was set in the request and the server responded with a 3xx status (specifically 301, 302, 303, 305, 307 or 308 status code) with a valid url in the location header, indicating a HTTP redirect. The url parameter contains the new redirect url as returned by the server in the location header.

This function was introduced in  Qt 5.6.

See also QNetworkRequest::FollowRedirectsAttribute.
*/
func (this *QNetworkReply) Redirected(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply10redirectedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirectAllowed()

/*
When client code handling the redirected() signal has verified the new URL, it emits this signal to allow the redirect to go ahead. This protocol applies to network requests whose redirects policy is set to QNetworkRequest::UserVerifiedRedirectPolicy

This function was introduced in  Qt 5.9.

See also QNetworkRequest::UserVerifiedRedirectPolicy, QNetworkAccessManager::setRedirectPolicy(), and QNetworkRequest::RedirectPolicyAttribute.
*/
func (this *QNetworkReply) RedirectAllowed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15redirectAllowedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void uploadProgress(qint64, qint64)

/*
This signal is emitted to indicate the progress of the upload part of this network request, if there's any. If there's no upload associated with this request, this signal will not be emitted.

The bytesSent parameter indicates the number of bytes uploaded, while bytesTotal indicates the total number of bytes to be uploaded. If the number of bytes to be uploaded could not be determined, bytesTotal will be -1.

The upload is finished when bytesSent is equal to bytesTotal. At that time, bytesTotal will not be -1.

See also downloadProgress().
*/
func (this *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply14uploadProgressExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytesSent, bytesTotal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void downloadProgress(qint64, qint64)

/*
This signal is emitted to indicate the progress of the download part of this network request, if there's any. If there's no download associated with this request, this signal will be emitted once with 0 as the value of both bytesReceived and bytesTotal.

The bytesReceived parameter indicates the number of bytes received, while bytesTotal indicates the total number of bytes expected to be downloaded. If the number of bytes to be downloaded is not known, bytesTotal will be -1.

The download is finished when bytesReceived is equal to bytesTotal. At that time, bytesTotal will not be -1.

Note that the values of both bytesReceived and bytesTotal may be different from size(), the total number of bytes obtained through read() or readAll(), or the value of the header(ContentLengthHeader). The reason for that is that there may be protocol overhead or the data may be compressed during the download.

See also uploadProgress() and bytesAvailable().
*/
func (this *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply16downloadProgressExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytesReceived, bytesTotal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:172
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QNetworkReply(QObject *)

/*
Creates a QNetworkReply object with parent parent.

You cannot directly instantiate QNetworkReply objects. Use QNetworkAccessManager functions to do that.
*/
func (*QNetworkReply) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkReply {
	return NewQNetworkReply(parent)
}
func NewQNetworkReply(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkReply {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReplyC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkReply")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:172
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QNetworkReply(QObject *)

/*
Creates a QNetworkReply object with parent parent.

You cannot directly instantiate QNetworkReply objects. Use QNetworkAccessManager functions to do that.
*/
func (*QNetworkReply) NewForInheritp() *QNetworkReply {
	return NewQNetworkReplyp()
}
func NewQNetworkReplyp() *QNetworkReply {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReplyC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkReply")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:174
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*

 */
func (this *QNetworkReply) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:176
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOperation(QNetworkAccessManager::Operation)

/*
Sets the associated operation for this object to be operation. This value will be returned by operation().

Note: The operation should be set when this object is created and not changed again.

See also operation() and setRequest().
*/
func (this *QNetworkReply) SetOperation(operation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setOperationEN21QNetworkAccessManager9OperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:177
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRequest(const QNetworkRequest &)

/*
Sets the associated request for this object to be request. This value will be returned by request().

Note: The request should be set when this object is created and not changed again.

See also request() and setOperation().
*/
func (this *QNetworkReply) SetRequest(request QNetworkRequest_ITF) {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply10setRequestERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:178
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setError(QNetworkReply::NetworkError, const QString &)

/*
Sets the error condition to be errorCode. The human-readable message is set with errorString.

Calling setError() does not emit the error(QNetworkReply::NetworkError) signal.

See also error() and errorString().
*/
func (this *QNetworkReply) SetError(errorCode int, errorString string) {
	var tmpArg1 = qtcore.NewQString5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply8setErrorENS_12NetworkErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), errorCode, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:179
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFinished(bool)

/*
Sets the reply as finished.

After having this set the replies data must not change.

This function was introduced in  Qt 4.8.

See also finished() and isFinished().
*/
func (this *QNetworkReply) SetFinished(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply11setFinishedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:180
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)

/*
Sets the URL being processed to be url. Normally, the URL matches that of the request that was posted, but for a variety of reasons it can be different (for example, a file path being made absolute or canonical).

See also url(), request(), and QNetworkRequest::url().
*/
func (this *QNetworkReply) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:181
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)

/*
Sets the known header header to be of value value. The corresponding raw form of the header will be set as well.

See also header(), setRawHeader(), and QNetworkRequest::setHeader().
*/
func (this *QNetworkReply) SetHeader(header int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:182
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)

/*
Sets the raw header headerName to be of value value. If headerName was previously set, it is overridden. Multiple HTTP headers of the same name are functionally equivalent to one single header with the values concatenated, separated by commas.

If headerName matches a known header, the value value will be parsed and the corresponding parsed form will also be set.

See also rawHeader(), header(), setHeader(), and QNetworkRequest::setRawHeader().
*/
func (this *QNetworkReply) SetRawHeader(headerName qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:183
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setAttribute(QNetworkRequest::Attribute, const QVariant &)

/*
Sets the attribute code to have value value. If code was previously set, it will be overridden. If value is an invalid QVariant, the attribute will be unset.

See also attribute() and QNetworkRequest::setAttribute().
*/
func (this *QNetworkReply) SetAttribute(code int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setAttributeEN15QNetworkRequest9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sslConfigurationImplementation(QSslConfiguration &) const

/*
This virtual method is provided to enable overriding the behavior of sslConfiguration(). sslConfiguration() is a public wrapper for this method. The configuration will be returned in configuration.

This function was introduced in  Qt 5.0.

See also setSslConfigurationImplementation() and sslConfiguration().
*/
func (this *QNetworkReply) SslConfigurationImplementation(arg0 QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSslConfiguration_PTR() != nil {
		convArg0 = arg0.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply30sslConfigurationImplementationER17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSslConfigurationImplementation(const QSslConfiguration &)

/*
This virtual method is provided to enable overriding the behavior of setSslConfiguration(). setSslConfiguration() is a public wrapper for this method. If you override this method use configuration to set the SSL configuration.

This function was introduced in  Qt 5.0.

See also sslConfigurationImplementation() and setSslConfiguration().
*/
func (this *QNetworkReply) SetSslConfigurationImplementation(arg0 QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSslConfiguration_PTR() != nil {
		convArg0 = arg0.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply33setSslConfigurationImplementationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Indicates all possible error conditions found during the processing of the request.



Note: When the HTTP protocol returns a redirect no error will be reported. You can check if there is a redirect with the QNetworkRequest::RedirectionTargetAttribute attribute.



See also error().

*/
type QNetworkReply__NetworkError = int

// no error condition.
const QNetworkReply__NoError QNetworkReply__NetworkError = 0

// the remote server refused the connection (the server is not accepting requests)
const QNetworkReply__ConnectionRefusedError QNetworkReply__NetworkError = 1

// the remote server closed the connection prematurely, before the entire reply was received and processed
const QNetworkReply__RemoteHostClosedError QNetworkReply__NetworkError = 2

// the remote host name was not found (invalid hostname)
const QNetworkReply__HostNotFoundError QNetworkReply__NetworkError = 3

// the connection to the remote server timed out
const QNetworkReply__TimeoutError QNetworkReply__NetworkError = 4

// the operation was canceled via calls to abort() or close() before it was finished.
const QNetworkReply__OperationCanceledError QNetworkReply__NetworkError = 5

// the SSL/TLS handshake failed and the encrypted channel could not be established. The sslErrors() signal should have been emitted.
const QNetworkReply__SslHandshakeFailedError QNetworkReply__NetworkError = 6

// the connection was broken due to disconnection from the network, however the system has initiated roaming to another access point. The request should be resubmitted and will be processed as soon as the connection is re-established.
const QNetworkReply__TemporaryNetworkFailureError QNetworkReply__NetworkError = 7

// the connection was broken due to disconnection from the network or failure to start the network.
const QNetworkReply__NetworkSessionFailedError QNetworkReply__NetworkError = 8

// the background request is not currently allowed due to platform policy.
const QNetworkReply__BackgroundRequestNotAllowedError QNetworkReply__NetworkError = 9

//
const QNetworkReply__TooManyRedirectsError QNetworkReply__NetworkError = 10

//
const QNetworkReply__InsecureRedirectError QNetworkReply__NetworkError = 11

//
const QNetworkReply__UnknownNetworkError QNetworkReply__NetworkError = 99

//
const QNetworkReply__ProxyConnectionRefusedError QNetworkReply__NetworkError = 101

//
const QNetworkReply__ProxyConnectionClosedError QNetworkReply__NetworkError = 102

//
const QNetworkReply__ProxyNotFoundError QNetworkReply__NetworkError = 103

//
const QNetworkReply__ProxyTimeoutError QNetworkReply__NetworkError = 104

//
const QNetworkReply__ProxyAuthenticationRequiredError QNetworkReply__NetworkError = 105

//
const QNetworkReply__UnknownProxyError QNetworkReply__NetworkError = 199

//
const QNetworkReply__ContentAccessDenied QNetworkReply__NetworkError = 201

//
const QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = 202

//
const QNetworkReply__ContentNotFoundError QNetworkReply__NetworkError = 203

//
const QNetworkReply__AuthenticationRequiredError QNetworkReply__NetworkError = 204

//
const QNetworkReply__ContentReSendError QNetworkReply__NetworkError = 205

//
const QNetworkReply__ContentConflictError QNetworkReply__NetworkError = 206

//
const QNetworkReply__ContentGoneError QNetworkReply__NetworkError = 207

//
const QNetworkReply__UnknownContentError QNetworkReply__NetworkError = 299

//
const QNetworkReply__ProtocolUnknownError QNetworkReply__NetworkError = 301

//
const QNetworkReply__ProtocolInvalidOperationError QNetworkReply__NetworkError = 302

//
const QNetworkReply__ProtocolFailure QNetworkReply__NetworkError = 399

//
const QNetworkReply__InternalServerError QNetworkReply__NetworkError = 401

//
const QNetworkReply__OperationNotImplementedError QNetworkReply__NetworkError = 402

//
const QNetworkReply__ServiceUnavailableError QNetworkReply__NetworkError = 403

//
const QNetworkReply__UnknownServerError QNetworkReply__NetworkError = 499

func (this *QNetworkReply) NetworkErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkReply_NetworkErrorItemName(val int) string {
	var nilthis *QNetworkReply
	return nilthis.NetworkErrorItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11457() {
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
