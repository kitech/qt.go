package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkreply.h
// #include <qnetworkreply.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QNetworkReply struct {
	*qtcore.QIODevice
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QNetworkReply) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:114
// index:0
// Public virtual
// void ~QNetworkReply()
func DeleteQNetworkReply(*QNetworkReply) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReplyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:117
// index:0
// Public virtual
// void close()
func (this *QNetworkReply) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:118
// index:0
// Public virtual
// bool isSequential()
func (this *QNetworkReply) IsSequential() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply12isSequentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:121
// index:0
// Public
// qint64 readBufferSize()
func (this *QNetworkReply) ReadBufferSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply14readBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:122
// index:0
// Public virtual
// void setReadBufferSize(qint64)
func (this *QNetworkReply) SetReadBufferSize(size int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply17setReadBufferSizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:124
// index:0
// Public
// QNetworkAccessManager * manager()
func (this *QNetworkReply) Manager() *QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply7managerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:125
// index:0
// Public
// QNetworkAccessManager::Operation operation()
func (this *QNetworkReply) Operation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply9operationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:126
// index:0
// Public
// QNetworkRequest request()
func (this *QNetworkReply) Request() *QNetworkRequest /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply7requestEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:127
// index:0
// Public
// QNetworkReply::NetworkError error()
func (this *QNetworkReply) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:159
// index:1
// Public
// void error(QNetworkReply::NetworkError)
func (this *QNetworkReply) Error_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply5errorENS_12NetworkErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:128
// index:0
// Public
// bool isFinished()
func (this *QNetworkReply) IsFinished() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:129
// index:0
// Public
// bool isRunning()
func (this *QNetworkReply) IsRunning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:130
// index:0
// Public
// QUrl url()
func (this *QNetworkReply) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:133
// index:0
// Public
// QVariant header(QNetworkRequest::KnownHeaders)
func (this *QNetworkReply) Header(header int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply6headerEN15QNetworkRequest12KnownHeadersE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:136
// index:0
// Public
// bool hasRawHeader(const QByteArray &)
func (this *QNetworkReply) HasRawHeader(headerName *qtcore.QByteArray) bool {
	var convArg0 = headerName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply12hasRawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:138
// index:0
// Public
// QByteArray rawHeader(const QByteArray &)
func (this *QNetworkReply) RawHeader(headerName *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = headerName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply9rawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:144
// index:0
// Public
// QVariant attribute(QNetworkRequest::Attribute)
func (this *QNetworkReply) Attribute(code int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply9attributeEN15QNetworkRequest9AttributeE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), code)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:147
// index:0
// Public
// QSslConfiguration sslConfiguration()
func (this *QNetworkReply) SslConfiguration() *QSslConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply16sslConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:148
// index:0
// Public
// void setSslConfiguration(const QSslConfiguration &)
func (this *QNetworkReply) SetSslConfiguration(configuration *QSslConfiguration) {
	var convArg0 = configuration.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply19setSslConfigurationERK17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:153
// index:0
// Public pure virtual
// void abort()
func (this *QNetworkReply) Abort() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply5abortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:154
// index:0
// Public virtual
// void ignoreSslErrors()
func (this *QNetworkReply) IgnoreSslErrors() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply15ignoreSslErrorsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:157
// index:0
// Public
// void metaDataChanged()
func (this *QNetworkReply) MetaDataChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply15metaDataChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:158
// index:0
// Public
// void finished()
func (this *QNetworkReply) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:161
// index:0
// Public
// void encrypted()
func (this *QNetworkReply) Encrypted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply9encryptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:163
// index:0
// Public
// void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)
func (this *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:165
// index:0
// Public
// void redirected(const QUrl &)
func (this *QNetworkReply) Redirected(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply10redirectedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:166
// index:0
// Public
// void redirectAllowed()
func (this *QNetworkReply) RedirectAllowed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply15redirectAllowedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:168
// index:0
// Public
// void uploadProgress(qint64, qint64)
func (this *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply14uploadProgressExx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bytesSent, bytesTotal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:169
// index:0
// Public
// void downloadProgress(qint64, qint64)
func (this *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply16downloadProgressExx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bytesReceived, bytesTotal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:172
// index:0
// Protected
// void QNetworkReply(QObject *)
func NewQNetworkReply(parent *qtcore.QObject /*777 QObject **/) *QNetworkReply {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReplyC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkReplyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:174
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QNetworkReply) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:176
// index:0
// Protected
// void setOperation(QNetworkAccessManager::Operation)
func (this *QNetworkReply) SetOperation(operation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply12setOperationEN21QNetworkAccessManager9OperationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), operation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:177
// index:0
// Protected
// void setRequest(const QNetworkRequest &)
func (this *QNetworkReply) SetRequest(request *QNetworkRequest) {
	var convArg0 = request.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply10setRequestERK15QNetworkRequest", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:178
// index:0
// Protected
// void setError(QNetworkReply::NetworkError, const QString &)
func (this *QNetworkReply) SetError(errorCode int, errorString *qtcore.QString) {
	var convArg1 = errorString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply8setErrorENS_12NetworkErrorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), errorCode, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:179
// index:0
// Protected
// void setFinished(bool)
func (this *QNetworkReply) SetFinished(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply11setFinishedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:180
// index:0
// Protected
// void setUrl(const QUrl &)
func (this *QNetworkReply) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply6setUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:181
// index:0
// Protected
// void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QNetworkReply) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:182
// index:0
// Protected
// void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QNetworkReply) SetRawHeader(headerName *qtcore.QByteArray, value *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply12setRawHeaderERK10QByteArrayS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:183
// index:0
// Protected
// void setAttribute(QNetworkRequest::Attribute, const QVariant &)
func (this *QNetworkReply) SetAttribute(code int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply12setAttributeEN15QNetworkRequest9AttributeERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:185
// index:0
// Protected virtual
// void sslConfigurationImplementation(QSslConfiguration &)
func (this *QNetworkReply) SslConfigurationImplementation(arg0 *QSslConfiguration) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkReply30sslConfigurationImplementationER17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:186
// index:0
// Protected virtual
// void setSslConfigurationImplementation(const QSslConfiguration &)
func (this *QNetworkReply) SetSslConfigurationImplementation(arg0 *QSslConfiguration) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkReply33setSslConfigurationImplementationERK17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QNetworkReply__NetworkError = int

const QNetworkReply__NoError QNetworkReply__NetworkError = 0
const QNetworkReply__ConnectionRefusedError QNetworkReply__NetworkError = 1
const QNetworkReply__RemoteHostClosedError QNetworkReply__NetworkError = 2
const QNetworkReply__HostNotFoundError QNetworkReply__NetworkError = 3
const QNetworkReply__TimeoutError QNetworkReply__NetworkError = 4
const QNetworkReply__OperationCanceledError QNetworkReply__NetworkError = 5
const QNetworkReply__SslHandshakeFailedError QNetworkReply__NetworkError = 6
const QNetworkReply__TemporaryNetworkFailureError QNetworkReply__NetworkError = 7
const QNetworkReply__NetworkSessionFailedError QNetworkReply__NetworkError = 8
const QNetworkReply__BackgroundRequestNotAllowedError QNetworkReply__NetworkError = 9
const QNetworkReply__TooManyRedirectsError QNetworkReply__NetworkError = 10
const QNetworkReply__InsecureRedirectError QNetworkReply__NetworkError = 11
const QNetworkReply__UnknownNetworkError QNetworkReply__NetworkError = 99
const QNetworkReply__ProxyConnectionRefusedError QNetworkReply__NetworkError = 101
const QNetworkReply__ProxyConnectionClosedError QNetworkReply__NetworkError = 102
const QNetworkReply__ProxyNotFoundError QNetworkReply__NetworkError = 103
const QNetworkReply__ProxyTimeoutError QNetworkReply__NetworkError = 104
const QNetworkReply__ProxyAuthenticationRequiredError QNetworkReply__NetworkError = 105
const QNetworkReply__UnknownProxyError QNetworkReply__NetworkError = 199
const QNetworkReply__ContentAccessDenied QNetworkReply__NetworkError = 201
const QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = 202
const QNetworkReply__ContentNotFoundError QNetworkReply__NetworkError = 203
const QNetworkReply__AuthenticationRequiredError QNetworkReply__NetworkError = 204
const QNetworkReply__ContentReSendError QNetworkReply__NetworkError = 205
const QNetworkReply__ContentConflictError QNetworkReply__NetworkError = 206
const QNetworkReply__ContentGoneError QNetworkReply__NetworkError = 207
const QNetworkReply__UnknownContentError QNetworkReply__NetworkError = 299
const QNetworkReply__ProtocolUnknownError QNetworkReply__NetworkError = 301
const QNetworkReply__ProtocolInvalidOperationError QNetworkReply__NetworkError = 302
const QNetworkReply__ProtocolFailure QNetworkReply__NetworkError = 399
const QNetworkReply__InternalServerError QNetworkReply__NetworkError = 401
const QNetworkReply__OperationNotImplementedError QNetworkReply__NetworkError = 402
const QNetworkReply__ServiceUnavailableError QNetworkReply__NetworkError = 403
const QNetworkReply__UnknownServerError QNetworkReply__NetworkError = 499

//  body block end
