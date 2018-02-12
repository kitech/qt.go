package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h
// #include <qnetworkcookiejar.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// bool validateCookie(const class QNetworkCookie &, const class QUrl &)
func (this *QNetworkCookieJar) InheritValidateCookie(f func(cookie *QNetworkCookie, url *qtcore.QUrl) bool) {
	qtrt.SetAllInheritCallback(this, "validateCookie", f)
}

type QNetworkCookieJar struct {
	*qtcore.QObject
}

func (this *QNetworkCookieJar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QNetworkCookieJar) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQNetworkCookieJarFromPointer(cthis unsafe.Pointer) *QNetworkCookieJar {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QNetworkCookieJar{bcthis0}
}
func (*QNetworkCookieJar) NewFromPointer(cthis unsafe.Pointer) *QNetworkCookieJar {
	return NewQNetworkCookieJarFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QNetworkCookieJar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkCookieJar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookieJar(QObject *)
func NewQNetworkCookieJar(parent *qtcore.QObject /*777 QObject **/) *QNetworkCookieJar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkCookieJar()
func DeleteQNetworkCookieJar(this *QNetworkCookieJar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertCookie(const QNetworkCookie &)
func (this *QNetworkCookieJar) InsertCookie(cookie *QNetworkCookie) bool {
	var convArg0 = cookie.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12insertCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool updateCookie(const QNetworkCookie &)
func (this *QNetworkCookieJar) UpdateCookie(cookie *QNetworkCookie) bool {
	var convArg0 = cookie.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12updateCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool deleteCookie(const QNetworkCookie &)
func (this *QNetworkCookieJar) DeleteCookie(cookie *QNetworkCookie) bool {
	var convArg0 = cookie.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12deleteCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool validateCookie(const QNetworkCookie &, const QUrl &)
func (this *QNetworkCookieJar) ValidateCookie(cookie *QNetworkCookie, url *qtcore.QUrl) bool {
	var convArg0 = cookie.GetCthis()
	var convArg1 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkCookieJar14validateCookieERK14QNetworkCookieRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
