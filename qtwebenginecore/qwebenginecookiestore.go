package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h
// #include <qwebenginecookiestore.h>
// #include <QtWebEngineCore>

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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

/*

 */
type QWebEngineCookieStore struct {
	*qtcore.QObject
}
type QWebEngineCookieStore_ITF interface {
	qtcore.QObject_ITF
	QWebEngineCookieStore_PTR() *QWebEngineCookieStore
}

func (ptr *QWebEngineCookieStore) QWebEngineCookieStore_PTR() *QWebEngineCookieStore { return ptr }

func (this *QWebEngineCookieStore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineCookieStore) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineCookieStoreFromPointer(cthis unsafe.Pointer) *QWebEngineCookieStore {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineCookieStore{bcthis0}
}
func (*QWebEngineCookieStore) NewFromPointer(cthis unsafe.Pointer) *QWebEngineCookieStore {
	return NewQWebEngineCookieStoreFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineCookieStore) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineCookieStore10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineCookieStore()

/*

 */
func DeleteQWebEngineCookieStore(this *QWebEngineCookieStore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCookie(const QNetworkCookie &, const QUrl &)

/*

 */
func (this *QWebEngineCookieStore) SetCookie(cookie qtnetwork.QNetworkCookie_ITF, origin qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if origin != nil && origin.QUrl_PTR() != nil {
		convArg1 = origin.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore9setCookieERK14QNetworkCookieRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCookie(const QNetworkCookie &, const QUrl &)

/*

 */
func (this *QWebEngineCookieStore) SetCookiep(cookie qtnetwork.QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore9setCookieERK14QNetworkCookieRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteCookie(const QNetworkCookie &, const QUrl &)

/*

 */
func (this *QWebEngineCookieStore) DeleteCookie(cookie qtnetwork.QNetworkCookie_ITF, origin qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if origin != nil && origin.QUrl_PTR() != nil {
		convArg1 = origin.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore12deleteCookieERK14QNetworkCookieRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteCookie(const QNetworkCookie &, const QUrl &)

/*

 */
func (this *QWebEngineCookieStore) DeleteCookiep(cookie qtnetwork.QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore12deleteCookieERK14QNetworkCookieRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteSessionCookies()

/*

 */
func (this *QWebEngineCookieStore) DeleteSessionCookies() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore20deleteSessionCookiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteAllCookies()

/*

 */
func (this *QWebEngineCookieStore) DeleteAllCookies() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore16deleteAllCookiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadAllCookies()

/*

 */
func (this *QWebEngineCookieStore) LoadAllCookies() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore14loadAllCookiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cookieAdded(const QNetworkCookie &)

/*

 */
func (this *QWebEngineCookieStore) CookieAdded(cookie qtnetwork.QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore11cookieAddedERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebenginecookiestore.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cookieRemoved(const QNetworkCookie &)

/*

 */
func (this *QWebEngineCookieStore) CookieRemoved(cookie qtnetwork.QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if cookie != nil && cookie.QNetworkCookie_PTR() != nil {
		convArg0 = cookie.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineCookieStore13cookieRemovedERK14QNetworkCookie", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11669() {
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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  keep block end
