package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkcookie.h
// #include <qnetworkcookie.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QNetworkCookie struct {
	*qtrt.CObject
}

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
func NewQNetworkCookie(name *qtcore.QByteArray, value *qtcore.QByteArray) *QNetworkCookie {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookieC2ERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkCookie &)
func (this *QNetworkCookie) Swap(other *QNetworkCookie) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSecure()
func (this *QNetworkCookie) IsSecure() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie8isSecureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSecure(_Bool)
func (this *QNetworkCookie) SetSecure(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9setSecureEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHttpOnly()
func (this *QNetworkCookie) IsHttpOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie10isHttpOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpOnly(_Bool)
func (this *QNetworkCookie) SetHttpOnly(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie11setHttpOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSessionCookie()
func (this *QNetworkCookie) IsSessionCookie() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie15isSessionCookieEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expirationDate()
func (this *QNetworkCookie) ExpirationDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie14expirationDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpirationDate(const QDateTime &)
func (this *QNetworkCookie) SetExpirationDate(date *qtcore.QDateTime) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie17setExpirationDateERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString domain()
func (this *QNetworkCookie) Domain() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie6domainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDomain(const QString &)
func (this *QNetworkCookie) SetDomain(domain *qtcore.QString) {
	var convArg0 = domain.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9setDomainERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QNetworkCookie) Path() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &)
func (this *QNetworkCookie) SetPath(path *qtcore.QString) {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie7setPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray name()
func (this *QNetworkCookie) Name() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QByteArray &)
func (this *QNetworkCookie) SetName(cookieName *qtcore.QByteArray) {
	var convArg0 = cookieName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie7setNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray value()
func (this *QNetworkCookie) Value() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(const QByteArray &)
func (this *QNetworkCookie) SetValue(value *qtcore.QByteArray) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie8setValueERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toRawForm(enum QNetworkCookie::RawForm)
func (this *QNetworkCookie) ToRawForm(form int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie9toRawFormENS_7RawFormE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), form)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSameIdentifier(const QNetworkCookie &)
func (this *QNetworkCookie) HasSameIdentifier(other *QNetworkCookie) bool {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QNetworkCookie17hasSameIdentifierERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void normalize(const QUrl &)
func (this *QNetworkCookie) Normalize(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QNetworkCookie9normalizeERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QNetworkCookie__RawForm = int

const QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = 0
const QNetworkCookie__Full QNetworkCookie__RawForm = 1

//  body block end
