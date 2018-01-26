package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkrequest.h
// #include <qnetworkrequest.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QNetworkRequest struct {
	*qtrt.CObject
}

func (this *QNetworkRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkRequest) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkRequestFromPointer(cthis unsafe.Pointer) *QNetworkRequest {
	return &QNetworkRequest{&qtrt.CObject{cthis}}
}
func (*QNetworkRequest) NewFromPointer(cthis unsafe.Pointer) *QNetworkRequest {
	return NewQNetworkRequestFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:124
// index:0
// Public
// void QNetworkRequest(const class QUrl &)
func NewQNetworkRequest(url *qtcore.QUrl) *QNetworkRequest {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequestC2ERK4QUrl", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkRequestFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:126
// index:0
// Public
// void ~QNetworkRequest()
func DeleteQNetworkRequest(*QNetworkRequest) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequestD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:132
// index:0
// Public inline
// void swap(class QNetworkRequest &)
func (this *QNetworkRequest) Swap(other *QNetworkRequest) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:138
// index:0
// Public
// QUrl url()
func (this *QNetworkRequest) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:139
// index:0
// Public
// void setUrl(const class QUrl &)
func (this *QNetworkRequest) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest6setUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:142
// index:0
// Public
// QVariant header(enum QNetworkRequest::KnownHeaders)
func (this *QNetworkRequest) Header(header int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest6headerENS_12KnownHeadersE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:143
// index:0
// Public
// void setHeader(enum QNetworkRequest::KnownHeaders, const class QVariant &)
func (this *QNetworkRequest) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest9setHeaderENS_12KnownHeadersERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:146
// index:0
// Public
// bool hasRawHeader(const class QByteArray &)
func (this *QNetworkRequest) HasRawHeader(headerName *qtcore.QByteArray) bool {
	var convArg0 = headerName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest12hasRawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:148
// index:0
// Public
// QByteArray rawHeader(const class QByteArray &)
func (this *QNetworkRequest) RawHeader(headerName *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = headerName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest9rawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:149
// index:0
// Public
// void setRawHeader(const class QByteArray &, const class QByteArray &)
func (this *QNetworkRequest) SetRawHeader(headerName *qtcore.QByteArray, value *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest12setRawHeaderERK10QByteArrayS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:152
// index:0
// Public
// QVariant attribute(enum QNetworkRequest::Attribute, const class QVariant &)
func (this *QNetworkRequest) Attribute(code int, defaultValue *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest9attributeENS_9AttributeERK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), code, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:153
// index:0
// Public
// void setAttribute(enum QNetworkRequest::Attribute, const class QVariant &)
func (this *QNetworkRequest) SetAttribute(code int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest12setAttributeENS_9AttributeERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:156
// index:0
// Public
// QSslConfiguration sslConfiguration()
func (this *QNetworkRequest) SslConfiguration() *QSslConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest16sslConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:157
// index:0
// Public
// void setSslConfiguration(const class QSslConfiguration &)
func (this *QNetworkRequest) SetSslConfiguration(configuration *QSslConfiguration) {
	var convArg0 = configuration.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest19setSslConfigurationERK17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:160
// index:0
// Public
// void setOriginatingObject(class QObject *)
func (this *QNetworkRequest) SetOriginatingObject(object *qtcore.QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest20setOriginatingObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:161
// index:0
// Public
// QObject * originatingObject()
func (this *QNetworkRequest) OriginatingObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest17originatingObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:163
// index:0
// Public
// QNetworkRequest::Priority priority()
func (this *QNetworkRequest) Priority() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest8priorityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:164
// index:0
// Public
// void setPriority(enum QNetworkRequest::Priority)
func (this *QNetworkRequest) SetPriority(priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:167
// index:0
// Public
// int maximumRedirectsAllowed()
func (this *QNetworkRequest) MaximumRedirectsAllowed() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkRequest23maximumRedirectsAllowedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:168
// index:0
// Public
// void setMaximumRedirectsAllowed(int)
func (this *QNetworkRequest) SetMaximumRedirectsAllowed(maximumRedirectsAllowed int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkRequest26setMaximumRedirectsAllowedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maximumRedirectsAllowed)
	gopp.ErrPrint(err, rv)
}

type QNetworkRequest__KnownHeaders = int

const QNetworkRequest__ContentTypeHeader QNetworkRequest__KnownHeaders = 0
const QNetworkRequest__ContentLengthHeader QNetworkRequest__KnownHeaders = 1
const QNetworkRequest__LocationHeader QNetworkRequest__KnownHeaders = 2
const QNetworkRequest__LastModifiedHeader QNetworkRequest__KnownHeaders = 3
const QNetworkRequest__CookieHeader QNetworkRequest__KnownHeaders = 4
const QNetworkRequest__SetCookieHeader QNetworkRequest__KnownHeaders = 5
const QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = 6
const QNetworkRequest__UserAgentHeader QNetworkRequest__KnownHeaders = 7
const QNetworkRequest__ServerHeader QNetworkRequest__KnownHeaders = 8

type QNetworkRequest__Attribute = int

const QNetworkRequest__HttpStatusCodeAttribute QNetworkRequest__Attribute = 0
const QNetworkRequest__HttpReasonPhraseAttribute QNetworkRequest__Attribute = 1
const QNetworkRequest__RedirectionTargetAttribute QNetworkRequest__Attribute = 2
const QNetworkRequest__ConnectionEncryptedAttribute QNetworkRequest__Attribute = 3
const QNetworkRequest__CacheLoadControlAttribute QNetworkRequest__Attribute = 4
const QNetworkRequest__CacheSaveControlAttribute QNetworkRequest__Attribute = 5
const QNetworkRequest__SourceIsFromCacheAttribute QNetworkRequest__Attribute = 6
const QNetworkRequest__DoNotBufferUploadDataAttribute QNetworkRequest__Attribute = 7
const QNetworkRequest__HttpPipeliningAllowedAttribute QNetworkRequest__Attribute = 8
const QNetworkRequest__HttpPipeliningWasUsedAttribute QNetworkRequest__Attribute = 9
const QNetworkRequest__CustomVerbAttribute QNetworkRequest__Attribute = 10
const QNetworkRequest__CookieLoadControlAttribute QNetworkRequest__Attribute = 11
const QNetworkRequest__AuthenticationReuseAttribute QNetworkRequest__Attribute = 12
const QNetworkRequest__CookieSaveControlAttribute QNetworkRequest__Attribute = 13
const QNetworkRequest__MaximumDownloadBufferSizeAttribute QNetworkRequest__Attribute = 14
const QNetworkRequest__DownloadBufferAttribute QNetworkRequest__Attribute = 15
const QNetworkRequest__SynchronousRequestAttribute QNetworkRequest__Attribute = 16
const QNetworkRequest__BackgroundRequestAttribute QNetworkRequest__Attribute = 17
const QNetworkRequest__SpdyAllowedAttribute QNetworkRequest__Attribute = 18
const QNetworkRequest__SpdyWasUsedAttribute QNetworkRequest__Attribute = 19
const QNetworkRequest__EmitAllUploadProgressSignalsAttribute QNetworkRequest__Attribute = 20
const QNetworkRequest__FollowRedirectsAttribute QNetworkRequest__Attribute = 21
const QNetworkRequest__HTTP2AllowedAttribute QNetworkRequest__Attribute = 22
const QNetworkRequest__HTTP2WasUsedAttribute QNetworkRequest__Attribute = 23
const QNetworkRequest__OriginalContentLengthAttribute QNetworkRequest__Attribute = 24
const QNetworkRequest__RedirectPolicyAttribute QNetworkRequest__Attribute = 25
const QNetworkRequest__User QNetworkRequest__Attribute = 1000
const QNetworkRequest__UserMax QNetworkRequest__Attribute = 32767

type QNetworkRequest__CacheLoadControl = int

const QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = 0
const QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = 1
const QNetworkRequest__PreferCache QNetworkRequest__CacheLoadControl = 2
const QNetworkRequest__AlwaysCache QNetworkRequest__CacheLoadControl = 3

type QNetworkRequest__LoadControl = int

const QNetworkRequest__Automatic QNetworkRequest__LoadControl = 0
const QNetworkRequest__Manual QNetworkRequest__LoadControl = 1

type QNetworkRequest__Priority = int

const QNetworkRequest__HighPriority QNetworkRequest__Priority = 1
const QNetworkRequest__NormalPriority QNetworkRequest__Priority = 3
const QNetworkRequest__LowPriority QNetworkRequest__Priority = 5

type QNetworkRequest__RedirectPolicy = int

const QNetworkRequest__ManualRedirectPolicy QNetworkRequest__RedirectPolicy = 0
const QNetworkRequest__NoLessSafeRedirectPolicy QNetworkRequest__RedirectPolicy = 1
const QNetworkRequest__SameOriginRedirectPolicy QNetworkRequest__RedirectPolicy = 2
const QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = 3

//  body block end
