package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkcookie.h
// #include <qnetworkcookie.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QNetworkCookie struct {
	*qtrt.CObject
}
type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (ptr *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie { return ptr }

func (this *QNetworkCookie) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkCookie) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkCookieFromPointer(cthis unsafe.Pointer) *QNetworkCookie {
	return &QNetworkCookie{&qtrt.CObject{cthis}}
}
func (*QNetworkCookie) NewFromPointer(cthis unsafe.Pointer) *QNetworkCookie {
	return NewQNetworkCookieFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookie(const QByteArray &, const QByteArray &)

/*
Create a new QNetworkCookie object, initializing the cookie name to name and its value to value.

A cookie is only valid if it has a name. However, the value is opaque to the application and being empty may have significance to the remote server.
*/
func (*QNetworkCookie) NewForInherit(name qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) *QNetworkCookie {
	return NewQNetworkCookie(name, value)
}
func NewQNetworkCookie(name qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) *QNetworkCookie {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieC2ERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkCookie)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookie(const QByteArray &, const QByteArray &)

/*
Create a new QNetworkCookie object, initializing the cookie name to name and its value to value.

A cookie is only valid if it has a name. However, the value is opaque to the application and being empty may have significance to the remote server.
*/
func (*QNetworkCookie) NewForInherit__() *QNetworkCookie {
	return NewQNetworkCookie__()
}
func NewQNetworkCookie__() *QNetworkCookie {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieC2ERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkCookie)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookie(const QByteArray &, const QByteArray &)

/*
Create a new QNetworkCookie object, initializing the cookie name to name and its value to value.

A cookie is only valid if it has a name. However, the value is opaque to the application and being empty may have significance to the remote server.
*/
func (*QNetworkCookie) NewForInherit__1(name qtcore.QByteArray_ITF) *QNetworkCookie {
	return NewQNetworkCookie__1(name)
}
func NewQNetworkCookie__1(name qtcore.QByteArray_ITF) *QNetworkCookie {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieC2ERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkCookie)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkCookie()

/*

 */
func DeleteQNetworkCookie(this *QNetworkCookie) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkCookie & operator=(QNetworkCookie &&)

/*

 */
func (this *QNetworkCookie) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkCookie {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkCookieFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkCookie)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:72
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkCookie & operator=(const QNetworkCookie &)

/*

 */
func (this *QNetworkCookie) Operator_equal_1(other QNetworkCookie_ITF) *QNetworkCookie {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkCookie_PTR() != nil {
		convArg0 = other.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkCookieFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkCookie)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkCookie &)

/*
Swaps this cookie with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkCookie) Swap(other QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkCookie_PTR() != nil {
		convArg0 = other.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkCookie &) const

/*

 */
func (this *QNetworkCookie) Operator_equal_equal(other QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkCookie_PTR() != nil {
		convArg0 = other.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookieeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkCookie &) const

/*

 */
func (this *QNetworkCookie) Operator_not_equal(other QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkCookie_PTR() != nil {
		convArg0 = other.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookieneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSecure() const

/*
Returns true if the "secure" option was specified in the cookie string, false otherwise.

Secure cookies may contain private information and should not be resent over unencrypted connections.

See also setSecure().
*/
func (this *QNetworkCookie) IsSecure() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie8isSecureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSecure(bool)

/*
Sets the secure flag of this cookie to enable.

Secure cookies may contain private information and should not be resent over unencrypted connections.

See also isSecure().
*/
func (this *QNetworkCookie) SetSecure(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9setSecureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHttpOnly() const

/*
Returns true if the "HttpOnly" flag is enabled for this cookie.

A cookie that is "HttpOnly" is only set and retrieved by the network requests and replies; i.e., the HTTP protocol. It is not accessible from scripts running on browsers.

This function was introduced in  Qt 4.5.

See also isSecure().
*/
func (this *QNetworkCookie) IsHttpOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie10isHttpOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpOnly(bool)

/*
Sets this cookie's "HttpOnly" flag to enable.

This function was introduced in  Qt 4.5.

See also isHttpOnly().
*/
func (this *QNetworkCookie) SetHttpOnly(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie11setHttpOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSessionCookie() const

/*
Returns true if this cookie is a session cookie. A session cookie is a cookie which has no expiration date, which means it should be discarded when the application's concept of session is over (usually, when the application exits).

See also expirationDate() and setExpirationDate().
*/
func (this *QNetworkCookie) IsSessionCookie() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie15isSessionCookieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expirationDate() const

/*
Returns the expiration date for this cookie. If this cookie is a session cookie, the QDateTime returned will not be valid. If the date is in the past, this cookie has already expired and should not be sent again back to a remote server.

The expiration date corresponds to the parameters of the "expires" entry in the cookie string.

See also isSessionCookie() and setExpirationDate().
*/
func (this *QNetworkCookie) ExpirationDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie14expirationDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpirationDate(const QDateTime &)

/*
Sets the expiration date of this cookie to date. Setting an invalid expiration date to this cookie will mean it's a session cookie.

See also isSessionCookie() and expirationDate().
*/
func (this *QNetworkCookie) SetExpirationDate(date qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDateTime_PTR() != nil {
		convArg0 = date.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie17setExpirationDateERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString domain() const

/*
Returns the domain this cookie is associated with. This corresponds to the "domain" field of the cookie string.

Note that the domain here may start with a dot, which is not a valid hostname. However, it means this cookie matches all hostnames ending with that domain name.

See also setDomain().
*/
func (this *QNetworkCookie) Domain() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie6domainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDomain(const QString &)

/*
Sets the domain associated with this cookie to be domain.

See also domain().
*/
func (this *QNetworkCookie) SetDomain(domain string) {
	var tmpArg0 = qtcore.NewQString_5(domain)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9setDomainERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const

/*
Returns the path associated with this cookie. This corresponds to the "path" field of the cookie string.

See also setPath().
*/
func (this *QNetworkCookie) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)

/*
Sets the path associated with this cookie to be path.

See also path().
*/
func (this *QNetworkCookie) SetPath(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray name() const

/*
Returns the name of this cookie. The only mandatory field of a cookie is its name, without which it is not considered valid.

See also setName() and value().
*/
func (this *QNetworkCookie) Name() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QByteArray &)

/*
Sets the name of this cookie to be cookieName. Note that setting a cookie name to an empty QByteArray will make this cookie invalid.

See also name() and value().
*/
func (this *QNetworkCookie) SetName(cookieName qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if cookieName != nil && cookieName.QByteArray_PTR() != nil {
		convArg0 = cookieName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie7setNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray value() const

/*
Returns this cookies value, as specified in the cookie string. Note that a cookie is still valid if its value is empty.

Cookie name-value pairs are considered opaque to the application: that is, their values don't mean anything.

See also setValue() and name().
*/
func (this *QNetworkCookie) Value() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(const QByteArray &)

/*
Sets the value of this cookie to be value.

See also value() and name().
*/
func (this *QNetworkCookie) SetValue(value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg0 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie8setValueERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toRawForm(QNetworkCookie::RawForm) const

/*
Returns the raw form of this QNetworkCookie. The QByteArray returned by this function is suitable for an HTTP header, either in a server response (the Set-Cookie header) or the client request (the Cookie header). You can choose from one of two formats, using form.

See also parseCookies().
*/
func (this *QNetworkCookie) ToRawForm(form int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie9toRawFormENS_7RawFormE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), form)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toRawForm(QNetworkCookie::RawForm) const

/*
Returns the raw form of this QNetworkCookie. The QByteArray returned by this function is suitable for an HTTP header, either in a server response (the Set-Cookie header) or the client request (the Cookie header). You can choose from one of two formats, using form.

See also parseCookies().
*/
func (this *QNetworkCookie) ToRawForm__() *qtcore.QByteArray /*123*/ {
	// arg: 0, QNetworkCookie::RawForm=Enum, QNetworkCookie::RawForm=Enum, , Invalid
	form := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie9toRawFormENS_7RawFormE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), form)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSameIdentifier(const QNetworkCookie &) const

/*
Returns true if this cookie has the same identifier tuple as other. The identifier tuple is composed of the name, domain and path.

See also operator==().
*/
func (this *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkCookie_PTR() != nil {
		convArg0 = other.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie17hasSameIdentifierERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize(const QUrl &)

/*
This functions normalizes the path and domain of the cookie if they were previously empty. The url parameter is used to determine the correct domain and path.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkCookie) Normalize(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9normalizeERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum is used with the toRawForm() function to declare which form of a cookie shall be returned.



Note that only the Full form of the cookie can be parsed back into its original contents.

See also toRawForm() and parseCookies().

*/
type QNetworkCookie__RawForm = int

// makes toRawForm() return only the "NAME=VALUE" part of the cookie, as suitable for sending back to a server in a client request's "Cookie:" header. Multiple cookies are separated by a semi-colon in the "Cookie:" header field.
const QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = 0

// makes toRawForm() return the full cookie contents, as suitable for sending to a client in a server's "Set-Cookie:" header.
const QNetworkCookie__Full QNetworkCookie__RawForm = 1

func (this *QNetworkCookie) RawFormItemName(val int) string {
	switch val {
	case QNetworkCookie__NameAndValueOnly: // 0
		return "NameAndValueOnly"
	case QNetworkCookie__Full: // 1
		return "Full"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkCookie_RawFormItemName(val int) string {
	var nilthis *QNetworkCookie
	return nilthis.RawFormItemName(val)
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
