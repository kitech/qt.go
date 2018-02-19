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
// extern C begin: 25
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

// bool validateCookie(const class QNetworkCookie &, const class QUrl &)
func (this *QNetworkCookieJar) InheritValidateCookie(f func(cookie *QNetworkCookie, url *qtcore.QUrl) bool) {
	qtrt.SetAllInheritCallback(this, "validateCookie", f)
}

type QNetworkCookieJar struct {
	*qtcore.QObject
}
type QNetworkCookieJar_ITF interface {
	qtcore.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar { return ptr }

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
// [8] const QMetaObject * metaObject() const
func (this *QNetworkCookieJar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkCookieJar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookieJar(QObject *)
func NewQNetworkCookieJar(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkCookieJar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkCookieJar")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCookieJar(QObject *)
func NewQNetworkCookieJar__() *QNetworkCookieJar {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkCookieJar")
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
func (this *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12insertCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool updateCookie(const QNetworkCookie &)
func (this *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12updateCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool deleteCookie(const QNetworkCookie &)
func (this *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkCookieJar12deleteCookieERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkcookiejar.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool validateCookie(const QNetworkCookie &, const QUrl &) const
func (this *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
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
