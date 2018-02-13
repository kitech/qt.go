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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// qint64 writeData(const char *, qint64)
func (this *QNetworkReply) InheritWriteData(f func(data string, len int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

// void setOperation(class QNetworkAccessManager::Operation)
func (this *QNetworkReply) InheritSetOperation(f func(operation int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setOperation", f)
}

// void setRequest(const class QNetworkRequest &)
func (this *QNetworkReply) InheritSetRequest(f func(request *QNetworkRequest) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRequest", f)
}

// void setError(enum QNetworkReply::NetworkError, const class QString &)
func (this *QNetworkReply) InheritSetError(f func(errorCode int, errorString string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setError", f)
}

// void setFinished(_Bool)
func (this *QNetworkReply) InheritSetFinished(f func(arg0 bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setFinished", f)
}

// void setUrl(const class QUrl &)
func (this *QNetworkReply) InheritSetUrl(f func(url *qtcore.QUrl) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setUrl", f)
}

// void setHeader(class QNetworkRequest::KnownHeaders, const class QVariant &)
func (this *QNetworkReply) InheritSetHeader(f func(header int, value *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setHeader", f)
}

// void setRawHeader(const class QByteArray &, const class QByteArray &)
func (this *QNetworkReply) InheritSetRawHeader(f func(headerName *qtcore.QByteArray, value *qtcore.QByteArray) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRawHeader", f)
}

// void setAttribute(class QNetworkRequest::Attribute, const class QVariant &)
func (this *QNetworkReply) InheritSetAttribute(f func(code int, value *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setAttribute", f)
}

// void sslConfigurationImplementation(class QSslConfiguration &)
func (this *QNetworkReply) InheritSslConfigurationImplementation(f func(arg0 *QSslConfiguration) /*void*/) {
	qtrt.SetAllInheritCallback(this, "sslConfigurationImplementation", f)
}

// void setSslConfigurationImplementation(const class QSslConfiguration &)
func (this *QNetworkReply) InheritSetSslConfigurationImplementation(f func(arg0 *QSslConfiguration) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSslConfigurationImplementation", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QNetworkReply) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkReply()
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
func (this *QNetworkReply) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential()
func (this *QNetworkReply) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize()
func (this *QNetworkReply) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)
func (this *QNetworkReply) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkAccessManager * manager()
func (this *QNetworkReply) Manager() *QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply7managerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkAccessManager::Operation operation()
func (this *QNetworkReply) Operation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9operationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkRequest request()
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
// [4] QNetworkReply::NetworkError error()
func (this *QNetworkReply) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:159
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QNetworkReply::NetworkError)
func (this *QNetworkReply) Error_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5errorENS_12NetworkErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished()
func (this *QNetworkReply) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QNetworkReply) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
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
// [16] QVariant header(QNetworkRequest::KnownHeaders)
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
// [1] bool hasRawHeader(const QByteArray &)
func (this *QNetworkReply) HasRawHeader(headerName *qtcore.QByteArray) bool {
	var convArg0 = headerName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply12hasRawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rawHeader(const QByteArray &)
func (this *QNetworkReply) RawHeader(headerName *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	var convArg0 = headerName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply9rawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:144
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant attribute(QNetworkRequest::Attribute)
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
// [8] QSslConfiguration sslConfiguration()
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
func (this *QNetworkReply) SetSslConfiguration(configuration *QSslConfiguration) {
	var convArg0 = configuration.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:153
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void abort()
func (this *QNetworkReply) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ignoreSslErrors()
func (this *QNetworkReply) IgnoreSslErrors() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15ignoreSslErrorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void metaDataChanged()
func (this *QNetworkReply) MetaDataChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15metaDataChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()
func (this *QNetworkReply) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted()
func (this *QNetworkReply) Encrypted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9encryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)
func (this *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = authenticator.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirected(const QUrl &)
func (this *QNetworkReply) Redirected(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply10redirectedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redirectAllowed()
func (this *QNetworkReply) RedirectAllowed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply15redirectAllowedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void uploadProgress(qint64, qint64)
func (this *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply14uploadProgressExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytesSent, bytesTotal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void downloadProgress(qint64, qint64)
func (this *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply16downloadProgressExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bytesReceived, bytesTotal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:172
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QNetworkReply(QObject *)
func NewQNetworkReply(parent *qtcore.QObject /*777 QObject **/) *QNetworkReply {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReplyC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:174
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QNetworkReply) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:176
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOperation(QNetworkAccessManager::Operation)
func (this *QNetworkReply) SetOperation(operation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setOperationEN21QNetworkAccessManager9OperationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:177
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRequest(const QNetworkRequest &)
func (this *QNetworkReply) SetRequest(request *QNetworkRequest) {
	var convArg0 = request.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply10setRequestERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:178
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setError(enum QNetworkReply::NetworkError, const QString &)
func (this *QNetworkReply) SetError(errorCode int, errorString string) {
	var tmpArg1 = qtcore.NewQString_5(errorString)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply8setErrorENS_12NetworkErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), errorCode, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:179
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFinished(_Bool)
func (this *QNetworkReply) SetFinished(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply11setFinishedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:180
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)
func (this *QNetworkReply) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:181
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QNetworkReply) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:182
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QNetworkReply) SetRawHeader(headerName *qtcore.QByteArray, value *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:183
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setAttribute(QNetworkRequest::Attribute, const QVariant &)
func (this *QNetworkReply) SetAttribute(code int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply12setAttributeEN15QNetworkRequest9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void sslConfigurationImplementation(QSslConfiguration &)
func (this *QNetworkReply) SslConfigurationImplementation(arg0 *QSslConfiguration) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkReply30sslConfigurationImplementationER17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkreply.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSslConfigurationImplementation(const QSslConfiguration &)
func (this *QNetworkReply) SetSslConfigurationImplementation(arg0 *QSslConfiguration) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkReply33setSslConfigurationImplementationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
