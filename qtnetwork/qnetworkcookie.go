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
func NewQNetworkCookie__() *QNetworkCookie {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg0 = qtcore.NewQByteArray()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
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
func NewQNetworkCookie__1(name qtcore.QByteArray_ITF) *QNetworkCookie {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
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
func (this *QNetworkCookie) IsSecure() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie8isSecureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSecure(_Bool)
func (this *QNetworkCookie) SetSecure(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9setSecureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHttpOnly() const
func (this *QNetworkCookie) IsHttpOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie10isHttpOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpOnly(_Bool)
func (this *QNetworkCookie) SetHttpOnly(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie11setHttpOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSessionCookie() const
func (this *QNetworkCookie) IsSessionCookie() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie15isSessionCookieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expirationDate() const
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
// [8] QByteArray toRawForm(enum QNetworkCookie::RawForm) const
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
// [8] QByteArray toRawForm(enum QNetworkCookie::RawForm) const
func (this *QNetworkCookie) ToRawForm__() *qtcore.QByteArray /*123*/ {
	// arg: 0, QNetworkCookie::RawForm=Enum, QNetworkCookie::RawForm=Enum,
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
func (this *QNetworkCookie) Normalize(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9normalizeERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QNetworkCookie__RawForm = int

const QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = 0
const QNetworkCookie__Full QNetworkCookie__RawForm = 1

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
