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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkCookieFromPointer(cthis unsafe.Pointer) *QNetworkCookie {
	return &QNetworkCookie{&qtrt.CObject{cthis}}
}
func (*QNetworkCookie) NewFromPointer(cthis unsafe.Pointer) *QNetworkCookie {
	return NewQNetworkCookieFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:66
// index:0
// Public
// void QNetworkCookie(const QByteArray &, const QByteArray &)
func NewQNetworkCookie(name *qtcore.QByteArray, value *qtcore.QByteArray) *QNetworkCookie {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookieC2ERK10QByteArrayS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkCookieFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:68
// index:0
// Public
// void ~QNetworkCookie()
func DeleteQNetworkCookie(*QNetworkCookie) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookieD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:74
// index:0
// Public inline
// void swap(QNetworkCookie &)
func (this *QNetworkCookie) Swap(other *QNetworkCookie) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:80
// index:0
// Public
// bool isSecure()
func (this *QNetworkCookie) IsSecure() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie8isSecureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:81
// index:0
// Public
// void setSecure(bool)
func (this *QNetworkCookie) SetSecure(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie9setSecureEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:82
// index:0
// Public
// bool isHttpOnly()
func (this *QNetworkCookie) IsHttpOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie10isHttpOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:83
// index:0
// Public
// void setHttpOnly(bool)
func (this *QNetworkCookie) SetHttpOnly(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie11setHttpOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:85
// index:0
// Public
// bool isSessionCookie()
func (this *QNetworkCookie) IsSessionCookie() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie15isSessionCookieEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:86
// index:0
// Public
// QDateTime expirationDate()
func (this *QNetworkCookie) ExpirationDate() *qtcore.QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie14expirationDateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:87
// index:0
// Public
// void setExpirationDate(const QDateTime &)
func (this *QNetworkCookie) SetExpirationDate(date *qtcore.QDateTime) {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie17setExpirationDateERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:89
// index:0
// Public
// QString domain()
func (this *QNetworkCookie) Domain() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie6domainEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:90
// index:0
// Public
// void setDomain(const QString &)
func (this *QNetworkCookie) SetDomain(domain *qtcore.QString) {
	var convArg0 = domain.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie9setDomainERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:92
// index:0
// Public
// QString path()
func (this *QNetworkCookie) Path() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie4pathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:93
// index:0
// Public
// void setPath(const QString &)
func (this *QNetworkCookie) SetPath(path *qtcore.QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie7setPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:95
// index:0
// Public
// QByteArray name()
func (this *QNetworkCookie) Name() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:96
// index:0
// Public
// void setName(const QByteArray &)
func (this *QNetworkCookie) SetName(cookieName *qtcore.QByteArray) {
	var convArg0 = cookieName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie7setNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:98
// index:0
// Public
// QByteArray value()
func (this *QNetworkCookie) Value() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie5valueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:99
// index:0
// Public
// void setValue(const QByteArray &)
func (this *QNetworkCookie) SetValue(value *qtcore.QByteArray) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie8setValueERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:101
// index:0
// Public
// QByteArray toRawForm(QNetworkCookie::RawForm)
func (this *QNetworkCookie) ToRawForm(form int) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie9toRawFormENS_7RawFormE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), form)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:103
// index:0
// Public
// bool hasSameIdentifier(const QNetworkCookie &)
func (this *QNetworkCookie) HasSameIdentifier(other *QNetworkCookie) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QNetworkCookie17hasSameIdentifierERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:104
// index:0
// Public
// void normalize(const QUrl &)
func (this *QNetworkCookie) Normalize(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QNetworkCookie9normalizeERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QNetworkCookie__RawForm = int

const QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = 0
const QNetworkCookie__Full QNetworkCookie__RawForm = 1

//  body block end
