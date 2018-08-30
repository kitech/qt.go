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
// extern C begin: 39
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

/*
Constructs a QNetworkRequest object with url as the URL to be requested.

See also url() and setUrl().
*/
func NewQNetworkRequest(url qtcore.QUrl_ITF) *QNetworkRequest {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestC2ERK4QUrl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkRequest)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkRequest(const QUrl &)

/*
Constructs a QNetworkRequest object with url as the URL to be requested.

See also url() and setUrl().
*/
func NewQNetworkRequest__() *QNetworkRequest {
	// arg: 0, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg0 = qtcore.NewQUrl()
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

/*

 */
func DeleteQNetworkRequest(this *QNetworkRequest) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkRequest & operator=(QNetworkRequest &&)

/*

 */
func (this *QNetworkRequest) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkRequest {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkRequest)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:130
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkRequest & operator=(const QNetworkRequest &)

/*

 */
func (this *QNetworkRequest) Operator_equal_1(other QNetworkRequest_ITF) *QNetworkRequest {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkRequest_PTR() != nil {
		convArg0 = other.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequestaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkRequestFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkRequest)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkRequest &)

/*
Swaps this network request with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkRequest) Swap(other QNetworkRequest_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkRequest_PTR() != nil {
		convArg0 = other.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:134
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkRequest &) const

/*

 */
func (this *QNetworkRequest) Operator_equal_equal(other QNetworkRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkRequest_PTR() != nil {
		convArg0 = other.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequesteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkRequest &) const

/*

 */
func (this *QNetworkRequest) Operator_not_equal(other QNetworkRequest_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkRequest_PTR() != nil {
		convArg0 = other.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequestneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*
Returns the URL this network request is referring to.

See also setUrl().
*/
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

/*
Sets the URL this network request is referring to be url.

See also url().
*/
func (this *QNetworkRequest) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:142
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant header(QNetworkRequest::KnownHeaders) const

/*
Returns the value of the known network header header if it is present in this request. If it is not present, returns QVariant() (i.e., an invalid variant).

See also KnownHeaders, rawHeader(), and setHeader().
*/
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
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)

/*
Sets the value of the known header header to be value, overriding any previously set headers. This operation also sets the equivalent raw HTTP header.

See also KnownHeaders, setRawHeader(), and header().
*/
func (this *QNetworkRequest) SetHeader(header int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest9setHeaderENS_12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasRawHeader(const QByteArray &) const

/*
Returns true if the raw header headerName is present in this network request.

See also rawHeader() and setRawHeader().
*/
func (this *QNetworkRequest) HasRawHeader(headerName qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest12hasRawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rawHeader(const QByteArray &) const

/*
Returns the raw form of header headerName. If no such header is present, an empty QByteArray is returned, which may be indistinguishable from a header that is present but has no content (use hasRawHeader() to find out if the header exists or not).

Raw headers can be set with setRawHeader() or with setHeader().

See also header() and setRawHeader().
*/
func (this *QNetworkRequest) RawHeader(headerName qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
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

/*
Sets the header headerName to be of value headerValue. If headerName corresponds to a known header (see QNetworkRequest::KnownHeaders), the raw format will be parsed and the corresponding "cooked" header will be set as well.

For example:


  request.setRawHeader(QByteArray("Last-Modified"), QByteArray("Sun, 06 Nov 1994 08:49:37 GMT"));



will also set the known header LastModifiedHeader to be the QDateTime object of the parsed date.

Note: Setting the same header twice overrides the previous setting. To accomplish the behaviour of multiple HTTP headers of the same name, you should concatenate the two values, separating them with a comma (",") and set one single raw header.

See also KnownHeaders, setHeader(), hasRawHeader(), and rawHeader().
*/
func (this *QNetworkRequest) SetRawHeader(headerName qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:152
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant attribute(QNetworkRequest::Attribute, const QVariant &) const

/*
Returns the attribute associated with the code code. If the attribute has not been set, it returns defaultValue.

Note: This function does not apply the defaults listed in QNetworkRequest::Attribute.

See also setAttribute() and QNetworkRequest::Attribute.
*/
func (this *QNetworkRequest) Attribute(code int, defaultValue qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if defaultValue != nil && defaultValue.QVariant_PTR() != nil {
		convArg1 = defaultValue.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest9attributeENS_9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:152
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant attribute(QNetworkRequest::Attribute, const QVariant &) const

/*
Returns the attribute associated with the code code. If the attribute has not been set, it returns defaultValue.

Note: This function does not apply the defaults listed in QNetworkRequest::Attribute.

See also setAttribute() and QNetworkRequest::Attribute.
*/
func (this *QNetworkRequest) Attribute__(code int) *qtcore.QVariant /*123*/ {
	// arg: 1, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg1 = qtcore.NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest9attributeENS_9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(QNetworkRequest::Attribute, const QVariant &)

/*
Sets the attribute associated with code code to be value value. If the attribute is already set, the previous value is discarded. In special, if value is an invalid QVariant, the attribute is unset.

See also attribute() and QNetworkRequest::Attribute.
*/
func (this *QNetworkRequest) SetAttribute(code int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest12setAttributeENS_9AttributeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration() const

/*
Returns this network request's SSL configuration. By default, no SSL settings are specified.

See also setSslConfiguration().
*/
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

/*
Sets this network request's SSL configuration to be config. The settings that apply are the private key, the local certificate, the SSL protocol (SSLv2, SSLv3, TLSv1.0 where applicable), the CA certificates and the ciphers that the SSL backend is allowed to use.

By default, no SSL configuration is set, which allows the backends to choose freely what configuration is best for them.

See also sslConfiguration() and QSslConfiguration::defaultConfiguration().
*/
func (this *QNetworkRequest) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOriginatingObject(QObject *)

/*
Allows setting a reference to the object initiating the request.

For example Qt WebKit sets the originating object to the QWebFrame that initiated the request.

This function was introduced in  Qt 4.6.

See also originatingObject().
*/
func (this *QNetworkRequest) SetOriginatingObject(object qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest20setOriginatingObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * originatingObject() const

/*
Returns a reference to the object that initiated this network request; returns 0 if not set or the object has been destroyed.

This function was introduced in  Qt 4.6.

See also setOriginatingObject().
*/
func (this *QNetworkRequest) OriginatingObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest17originatingObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkRequest::Priority priority() const

/*
Return the priority of this request.

This function was introduced in  Qt 4.7.

See also setPriority().
*/
func (this *QNetworkRequest) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(QNetworkRequest::Priority)

/*
Set the priority of this request to priority.

Note: The priority is only a hint to the network access manager. It can use it or not. Currently it is used for HTTP to decide which request should be sent first to a server.

This function was introduced in  Qt 4.7.

See also priority().
*/
func (this *QNetworkRequest) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumRedirectsAllowed() const

/*
Returns the maximum number of redirects allowed to be followed for this request.

This function was introduced in  Qt 5.6.

See also setMaximumRedirectsAllowed().
*/
func (this *QNetworkRequest) MaximumRedirectsAllowed() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkRequest23maximumRedirectsAllowedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumRedirectsAllowed(int)

/*
Sets the maximum number of redirects allowed to be followed for this request to maxRedirectsAllowed.

This function was introduced in  Qt 5.6.

See also maximumRedirectsAllowed().
*/
func (this *QNetworkRequest) SetMaximumRedirectsAllowed(maximumRedirectsAllowed int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkRequest26setMaximumRedirectsAllowedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximumRedirectsAllowed)
	qtrt.ErrPrint(err, rv)
}

/*
List of known header types that QNetworkRequest parses. Each known header is also represented in raw form with its full HTTP name.



See also header(), setHeader(), rawHeader(), and setRawHeader().

*/
type QNetworkRequest__KnownHeaders = int

// Corresponds to the HTTP Content-Type header and contains a string containing the media (MIME) type and any auxiliary data (for instance, charset).
const QNetworkRequest__ContentTypeHeader QNetworkRequest__KnownHeaders = 0

// Corresponds to the HTTP Content-Length header and contains the length in bytes of the data transmitted.
const QNetworkRequest__ContentLengthHeader QNetworkRequest__KnownHeaders = 1

// Corresponds to the HTTP Location header and contains a URL representing the actual location of the data, including the destination URL in case of redirections.
const QNetworkRequest__LocationHeader QNetworkRequest__KnownHeaders = 2

// Corresponds to the HTTP Last-Modified header and contains a QDateTime representing the last modification date of the contents.
const QNetworkRequest__LastModifiedHeader QNetworkRequest__KnownHeaders = 3

// Corresponds to the HTTP Cookie header and contains a QList<QNetworkCookie> representing the cookies to be sent back to the server.
const QNetworkRequest__CookieHeader QNetworkRequest__KnownHeaders = 4

// Corresponds to the HTTP Set-Cookie header and contains a QList<QNetworkCookie> representing the cookies sent by the server to be stored locally.
const QNetworkRequest__SetCookieHeader QNetworkRequest__KnownHeaders = 5

// Corresponds to the HTTP Content-Disposition header and contains a string containing the disposition type (for instance, attachment) and a parameter (for instance, filename).
const QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = 6

// The User-Agent header sent by HTTP clients.
const QNetworkRequest__UserAgentHeader QNetworkRequest__KnownHeaders = 7

// The Server header received by HTTP clients.
const QNetworkRequest__ServerHeader QNetworkRequest__KnownHeaders = 8

func (this *QNetworkRequest) KnownHeadersItemName(val int) string {
	switch val {
	case QNetworkRequest__ContentTypeHeader: // 0
		return "ContentTypeHeader"
	case QNetworkRequest__ContentLengthHeader: // 1
		return "ContentLengthHeader"
	case QNetworkRequest__LocationHeader: // 2
		return "LocationHeader"
	case QNetworkRequest__LastModifiedHeader: // 3
		return "LastModifiedHeader"
	case QNetworkRequest__CookieHeader: // 4
		return "CookieHeader"
	case QNetworkRequest__SetCookieHeader: // 5
		return "SetCookieHeader"
	case QNetworkRequest__ContentDispositionHeader: // 6
		return "ContentDispositionHeader"
	case QNetworkRequest__UserAgentHeader: // 7
		return "UserAgentHeader"
	case QNetworkRequest__ServerHeader: // 8
		return "ServerHeader"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_KnownHeadersItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.KnownHeadersItemName(val)
}

/*
Attribute codes for the QNetworkRequest and QNetworkReply.

Attributes are extra meta-data that are used to control the behavior of the request and to pass further information from the reply back to the application. Attributes are also extensible, allowing custom implementations to pass custom values.

The following table explains what the default attribute codes are, the QVariant types associated, the default value if said attribute is missing and whether it's used in requests or replies.



This enum was introduced or modified in  Qt 4.7.

*/
type QNetworkRequest__Attribute = int

//
const QNetworkRequest__HttpStatusCodeAttribute QNetworkRequest__Attribute = 0

// Replies only, type: QMetaType::QByteArray (no default) Indicates the HTTP reason phrase as received from the HTTP server (like "Ok", "Found", "Not Found", "Access Denied", etc.) This is the human-readable representation of the status code (see above). If the connection was not HTTP-based, this attribute will not be present.
const QNetworkRequest__HttpReasonPhraseAttribute QNetworkRequest__Attribute = 1

// Replies only, type: QMetaType::QUrl (no default) If present, it indicates that the server is redirecting the request to a different URL. The Network Access API does not by default follow redirections: the application can determine if the requested redirection should be allowed, according to its security policies, or it can set QNetworkRequest::FollowRedirectsAttribute to true (in which case the redirection will be followed and this attribute will not be present in the reply). The returned URL might be relative. Use QUrl::resolved() to create an absolute URL out of it.
const QNetworkRequest__RedirectionTargetAttribute QNetworkRequest__Attribute = 2

// Replies only, type: QMetaType::Bool (default: false) Indicates whether the data was obtained through an encrypted (secure) connection.
const QNetworkRequest__ConnectionEncryptedAttribute QNetworkRequest__Attribute = 3

// Requests only, type: QMetaType::Int (default: QNetworkRequest::PreferNetwork) Controls how the cache should be accessed. The possible values are those of QNetworkRequest::CacheLoadControl. Note that the default QNetworkAccessManager implementation does not support caching. However, this attribute may be used by certain backends to modify their requests (for example, for caching proxies).
const QNetworkRequest__CacheLoadControlAttribute QNetworkRequest__Attribute = 4

// Requests only, type: QMetaType::Bool (default: true) Controls if the data obtained should be saved to cache for future uses. If the value is false, the data obtained will not be automatically cached. If true, data may be cached, provided it is cacheable (what is cacheable depends on the protocol being used).
const QNetworkRequest__CacheSaveControlAttribute QNetworkRequest__Attribute = 5

// Replies only, type: QMetaType::Bool (default: false) Indicates whether the data was obtained from cache or not.
const QNetworkRequest__SourceIsFromCacheAttribute QNetworkRequest__Attribute = 6

// Requests only, type: QMetaType::Bool (default: false) Indicates whether the QNetworkAccessManager code is allowed to buffer the upload data, e.g. when doing a HTTP POST. When using this flag with sequential upload data, the ContentLengthHeader header must be set.
const QNetworkRequest__DoNotBufferUploadDataAttribute QNetworkRequest__Attribute = 7

// Requests only, type: QMetaType::Bool (default: false) Indicates whether the QNetworkAccessManager code is allowed to use HTTP pipelining with this request.
const QNetworkRequest__HttpPipeliningAllowedAttribute QNetworkRequest__Attribute = 8

// Replies only, type: QMetaType::Bool Indicates whether the HTTP pipelining was used for receiving this reply.
const QNetworkRequest__HttpPipeliningWasUsedAttribute QNetworkRequest__Attribute = 9

//
const QNetworkRequest__CustomVerbAttribute QNetworkRequest__Attribute = 10

//
const QNetworkRequest__CookieLoadControlAttribute QNetworkRequest__Attribute = 11

//
const QNetworkRequest__AuthenticationReuseAttribute QNetworkRequest__Attribute = 12

//
const QNetworkRequest__CookieSaveControlAttribute QNetworkRequest__Attribute = 13

//
const QNetworkRequest__MaximumDownloadBufferSizeAttribute QNetworkRequest__Attribute = 14

//
const QNetworkRequest__DownloadBufferAttribute QNetworkRequest__Attribute = 15

//
const QNetworkRequest__SynchronousRequestAttribute QNetworkRequest__Attribute = 16

//
const QNetworkRequest__BackgroundRequestAttribute QNetworkRequest__Attribute = 17

//
const QNetworkRequest__SpdyAllowedAttribute QNetworkRequest__Attribute = 18

//
const QNetworkRequest__SpdyWasUsedAttribute QNetworkRequest__Attribute = 19

//
const QNetworkRequest__EmitAllUploadProgressSignalsAttribute QNetworkRequest__Attribute = 20

//
const QNetworkRequest__FollowRedirectsAttribute QNetworkRequest__Attribute = 21

//
const QNetworkRequest__HTTP2AllowedAttribute QNetworkRequest__Attribute = 22

//
const QNetworkRequest__HTTP2WasUsedAttribute QNetworkRequest__Attribute = 23

//
const QNetworkRequest__OriginalContentLengthAttribute QNetworkRequest__Attribute = 24

//
const QNetworkRequest__RedirectPolicyAttribute QNetworkRequest__Attribute = 25

//
const QNetworkRequest__User QNetworkRequest__Attribute = 1000

//
const QNetworkRequest__UserMax QNetworkRequest__Attribute = 32767

func (this *QNetworkRequest) AttributeItemName(val int) string {
	switch val {
	case QNetworkRequest__HttpStatusCodeAttribute: // 0
		return "HttpStatusCodeAttribute"
	case QNetworkRequest__HttpReasonPhraseAttribute: // 1
		return "HttpReasonPhraseAttribute"
	case QNetworkRequest__RedirectionTargetAttribute: // 2
		return "RedirectionTargetAttribute"
	case QNetworkRequest__ConnectionEncryptedAttribute: // 3
		return "ConnectionEncryptedAttribute"
	case QNetworkRequest__CacheLoadControlAttribute: // 4
		return "CacheLoadControlAttribute"
	case QNetworkRequest__CacheSaveControlAttribute: // 5
		return "CacheSaveControlAttribute"
	case QNetworkRequest__SourceIsFromCacheAttribute: // 6
		return "SourceIsFromCacheAttribute"
	case QNetworkRequest__DoNotBufferUploadDataAttribute: // 7
		return "DoNotBufferUploadDataAttribute"
	case QNetworkRequest__HttpPipeliningAllowedAttribute: // 8
		return "HttpPipeliningAllowedAttribute"
	case QNetworkRequest__HttpPipeliningWasUsedAttribute: // 9
		return "HttpPipeliningWasUsedAttribute"
	case QNetworkRequest__CustomVerbAttribute: // 10
		return "CustomVerbAttribute"
	case QNetworkRequest__CookieLoadControlAttribute: // 11
		return "CookieLoadControlAttribute"
	case QNetworkRequest__AuthenticationReuseAttribute: // 12
		return "AuthenticationReuseAttribute"
	case QNetworkRequest__CookieSaveControlAttribute: // 13
		return "CookieSaveControlAttribute"
	case QNetworkRequest__MaximumDownloadBufferSizeAttribute: // 14
		return "MaximumDownloadBufferSizeAttribute"
	case QNetworkRequest__DownloadBufferAttribute: // 15
		return "DownloadBufferAttribute"
	case QNetworkRequest__SynchronousRequestAttribute: // 16
		return "SynchronousRequestAttribute"
	case QNetworkRequest__BackgroundRequestAttribute: // 17
		return "BackgroundRequestAttribute"
	case QNetworkRequest__SpdyAllowedAttribute: // 18
		return "SpdyAllowedAttribute"
	case QNetworkRequest__SpdyWasUsedAttribute: // 19
		return "SpdyWasUsedAttribute"
	case QNetworkRequest__EmitAllUploadProgressSignalsAttribute: // 20
		return "EmitAllUploadProgressSignalsAttribute"
	case QNetworkRequest__FollowRedirectsAttribute: // 21
		return "FollowRedirectsAttribute"
	case QNetworkRequest__HTTP2AllowedAttribute: // 22
		return "HTTP2AllowedAttribute"
	case QNetworkRequest__HTTP2WasUsedAttribute: // 23
		return "HTTP2WasUsedAttribute"
	case QNetworkRequest__OriginalContentLengthAttribute: // 24
		return "OriginalContentLengthAttribute"
	case QNetworkRequest__RedirectPolicyAttribute: // 25
		return "RedirectPolicyAttribute"
	case QNetworkRequest__User: // 1000
		return "User"
	case QNetworkRequest__UserMax: // 32767
		return "UserMax"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_AttributeItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.AttributeItemName(val)
}

/*
Controls the caching mechanism of QNetworkAccessManager.


*/
type QNetworkRequest__CacheLoadControl = int

// always load from network and do not check if the cache has a valid entry (similar to the "Reload" feature in browsers); in addition, force intermediate caches to re-validate.
const QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = 0

// default value; load from the network if the cached entry is older than the network entry. This will never return stale data from the cache, but revalidate resources that have become stale.
const QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = 1

// load from cache if available, otherwise load from network. Note that this can return possibly stale (but not expired) items from cache.
const QNetworkRequest__PreferCache QNetworkRequest__CacheLoadControl = 2

// only load from cache, indicating error if the item was not cached (i.e., off-line mode)
const QNetworkRequest__AlwaysCache QNetworkRequest__CacheLoadControl = 3

func (this *QNetworkRequest) CacheLoadControlItemName(val int) string {
	switch val {
	case QNetworkRequest__AlwaysNetwork: // 0
		return "AlwaysNetwork"
	case QNetworkRequest__PreferNetwork: // 1
		return "PreferNetwork"
	case QNetworkRequest__PreferCache: // 2
		return "PreferCache"
	case QNetworkRequest__AlwaysCache: // 3
		return "AlwaysCache"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_CacheLoadControlItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.CacheLoadControlItemName(val)
}

/*
Indicates if an aspect of the request's loading mechanism has been manually overridden, e.g. by Qt WebKit.



This enum was introduced or modified in  Qt 4.7.

*/
type QNetworkRequest__LoadControl = int

// default value: indicates default behaviour.
const QNetworkRequest__Automatic QNetworkRequest__LoadControl = 0

// indicates behaviour has been manually overridden.
const QNetworkRequest__Manual QNetworkRequest__LoadControl = 1

func (this *QNetworkRequest) LoadControlItemName(val int) string {
	switch val {
	case QNetworkRequest__Automatic: // 0
		return "Automatic"
	case QNetworkRequest__Manual: // 1
		return "Manual"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_LoadControlItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.LoadControlItemName(val)
}

/*
This enum lists the possible network request priorities.



This enum was introduced or modified in  Qt 4.7.

*/
type QNetworkRequest__Priority = int

// High priority
const QNetworkRequest__HighPriority QNetworkRequest__Priority = 1

// Normal priority
const QNetworkRequest__NormalPriority QNetworkRequest__Priority = 3

// Low priority
const QNetworkRequest__LowPriority QNetworkRequest__Priority = 5

func (this *QNetworkRequest) PriorityItemName(val int) string {
	switch val {
	case QNetworkRequest__HighPriority: // 1
		return "HighPriority"
	case QNetworkRequest__NormalPriority: // 3
		return "NormalPriority"
	case QNetworkRequest__LowPriority: // 5
		return "LowPriority"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_PriorityItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.PriorityItemName(val)
}

/*
Indicates whether the Network Access API should automatically follow a HTTP redirect response or not.



This enum was introduced or modified in  Qt 5.9.

*/
type QNetworkRequest__RedirectPolicy = int

// Default value: not following any redirects.
const QNetworkRequest__ManualRedirectPolicy QNetworkRequest__RedirectPolicy = 0

// Only "http"->"http", "http" -> "https" or "https" -> "https" redirects are allowed. Equivalent to setting the old FollowRedirectsAttribute to true
const QNetworkRequest__NoLessSafeRedirectPolicy QNetworkRequest__RedirectPolicy = 1

//
const QNetworkRequest__SameOriginRedirectPolicy QNetworkRequest__RedirectPolicy = 2

// Client decides whether to follow each redirect by handling the redirected() signal, emitting redirectAllowed() on the QNetworkReply object to allow the redirect or aborting/finishing it to reject the redirect. This can be used, for example, to ask the user whether to accept the redirect, or to decide based on some app-specific configuration.
const QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = 3

func (this *QNetworkRequest) RedirectPolicyItemName(val int) string {
	switch val {
	case QNetworkRequest__ManualRedirectPolicy: // 0
		return "ManualRedirectPolicy"
	case QNetworkRequest__NoLessSafeRedirectPolicy: // 1
		return "NoLessSafeRedirectPolicy"
	case QNetworkRequest__SameOriginRedirectPolicy: // 2
		return "SameOriginRedirectPolicy"
	case QNetworkRequest__UserVerifiedRedirectPolicy: // 3
		return "UserVerifiedRedirectPolicy"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkRequest_RedirectPolicyItemName(val int) string {
	var nilthis *QNetworkRequest
	return nilthis.RedirectPolicyItemName(val)
}

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
