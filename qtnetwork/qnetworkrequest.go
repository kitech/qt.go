package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkrequest.h
// #include <qnetworkrequest.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QNetworkRequest struct {
	*qtrt.CObject
}
type QNetworkRequest_ITF interface {
	QNetworkRequest_PTR() *QNetworkRequest
}

func (ptr *QNetworkRequest) QNetworkRequest_PTR() *QNetworkRequest { return ptr }

func (this *QNetworkRequest) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkRequest) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkRequestFromPointer(cthis unsafe.Pointer) *QNetworkRequest {
	return &QNetworkRequest{&qtrt.CObject{cthis}}
}
func (*QNetworkRequest) NewFromPointer(cthis unsafe.Pointer) *QNetworkRequest {
	return NewQNetworkRequestFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkRequest(const QUrl &)
func NewQNetworkRequest(url qtcore.QUrl_ITF) *QNetworkRequest {
	var convArg0 = url.QUrl_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkRequest)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkRequest()
func DeleteQNetworkRequest(this *QNetworkRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkRequest &)
func (this *QNetworkRequest) Swap(other QNetworkRequest_ITF) {
	var convArg0 = other.QNetworkRequest_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QNetworkRequest) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)
func (this *QNetworkRequest) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 = url.QUrl_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:142
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant header(enum QNetworkRequest::KnownHeaders)
func (this *QNetworkRequest) Header(header int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest6headerENS_12KnownHeadersE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(enum QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QNetworkRequest) SetHeader(header int, value qtcore.QVariant_ITF) {
	var convArg1 = value.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest9setHeaderENS_12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasRawHeader(const QByteArray &)
func (this *QNetworkRequest) HasRawHeader(headerName qtcore.QByteArray_ITF) bool {
	var convArg0 = headerName.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest12hasRawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rawHeader(const QByteArray &)
func (this *QNetworkRequest) RawHeader(headerName qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 = headerName.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest9rawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QNetworkRequest) SetRawHeader(headerName qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 = headerName.QByteArray_PTR().GetCthis()
	var convArg1 = value.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:152
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant attribute(enum QNetworkRequest::Attribute, const QVariant &)
func (this *QNetworkRequest) Attribute(code int, defaultValue qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 = defaultValue.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest9attributeENS_9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(enum QNetworkRequest::Attribute, const QVariant &)
func (this *QNetworkRequest) SetAttribute(code int, value qtcore.QVariant_ITF) {
	var convArg1 = value.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest12setAttributeENS_9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration()
func (this *QNetworkRequest) SslConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest16sslConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslConfiguration(const QSslConfiguration &)
func (this *QNetworkRequest) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOriginatingObject(QObject *)
func (this *QNetworkRequest) SetOriginatingObject(object qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 = object.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest20setOriginatingObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * originatingObject()
func (this *QNetworkRequest) OriginatingObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest17originatingObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkRequest::Priority priority()
func (this *QNetworkRequest) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(enum QNetworkRequest::Priority)
func (this *QNetworkRequest) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumRedirectsAllowed()
func (this *QNetworkRequest) MaximumRedirectsAllowed() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest23maximumRedirectsAllowedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumRedirectsAllowed(int)
func (this *QNetworkRequest) SetMaximumRedirectsAllowed(maximumRedirectsAllowed int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest26setMaximumRedirectsAllowedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximumRedirectsAllowed)
	qtrt.ErrPrint(err, rv)
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
