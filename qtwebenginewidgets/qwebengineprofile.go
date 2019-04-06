package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h
// #include <qwebengineprofile.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 84
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
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineProfile struct {
	*qtcore.QObject
}
type QWebEngineProfile_ITF interface {
	qtcore.QObject_ITF
	QWebEngineProfile_PTR() *QWebEngineProfile
}

func (ptr *QWebEngineProfile) QWebEngineProfile_PTR() *QWebEngineProfile { return ptr }

func (this *QWebEngineProfile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEngineProfile) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEngineProfileFromPointer(cthis unsafe.Pointer) *QWebEngineProfile {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEngineProfile{bcthis0}
}
func (*QWebEngineProfile) NewFromPointer(cthis unsafe.Pointer) *QWebEngineProfile {
	return NewQWebEngineProfileFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineProfile) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineProfile(QObject *)

/*

 */
func (*QWebEngineProfile) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineProfile {
	return NewQWebEngineProfile(parent)
}
func NewQWebEngineProfile(parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineProfile {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineProfile(QObject *)

/*

 */
func (*QWebEngineProfile) NewForInheritp() *QWebEngineProfile {
	return NewQWebEngineProfilep()
}
func NewQWebEngineProfilep() *QWebEngineProfile {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineProfile(const QString &, QObject *)

/*

 */
func (*QWebEngineProfile) NewForInherit1(name string, parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineProfile {
	return NewQWebEngineProfile1(name, parent)
}
func NewQWebEngineProfile1(name string, parent qtcore.QObject_ITF /*777 QObject **/) *QWebEngineProfile {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineProfile(const QString &, QObject *)

/*

 */
func (*QWebEngineProfile) NewForInherit1p(name string) *QWebEngineProfile {
	return NewQWebEngineProfile1p(name)
}
func NewQWebEngineProfile1p(name string) *QWebEngineProfile {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineProfile()

/*

 */
func DeleteQWebEngineProfile(this *QWebEngineProfile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString storageName() const

/*

 */
func (this *QWebEngineProfile) StorageName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile11storageNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOffTheRecord() const

/*

 */
func (this *QWebEngineProfile) IsOffTheRecord() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile14isOffTheRecordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString persistentStoragePath() const

/*

 */
func (this *QWebEngineProfile) PersistentStoragePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile21persistentStoragePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentStoragePath(const QString &)

/*

 */
func (this *QWebEngineProfile) SetPersistentStoragePath(path string) {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile24setPersistentStoragePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cachePath() const

/*

 */
func (this *QWebEngineProfile) CachePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile9cachePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCachePath(const QString &)

/*

 */
func (this *QWebEngineProfile) SetCachePath(path string) {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile12setCachePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString httpUserAgent() const

/*

 */
func (this *QWebEngineProfile) HttpUserAgent() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile13httpUserAgentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpUserAgent(const QString &)

/*

 */
func (this *QWebEngineProfile) SetHttpUserAgent(userAgent string) {
	var tmpArg0 = qtcore.NewQString5(userAgent)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile16setHttpUserAgentERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineProfile::HttpCacheType httpCacheType() const

/*

 */
func (this *QWebEngineProfile) HttpCacheType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile13httpCacheTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpCacheType(QWebEngineProfile::HttpCacheType)

/*

 */
func (this *QWebEngineProfile) SetHttpCacheType(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile16setHttpCacheTypeENS_13HttpCacheTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpAcceptLanguage(const QString &)

/*

 */
func (this *QWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	var tmpArg0 = qtcore.NewQString5(httpAcceptLanguage)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile21setHttpAcceptLanguageERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QString httpAcceptLanguage() const

/*

 */
func (this *QWebEngineProfile) HttpAcceptLanguage() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile18httpAcceptLanguageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineProfile::PersistentCookiesPolicy persistentCookiesPolicy() const

/*

 */
func (this *QWebEngineProfile) PersistentCookiesPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile23persistentCookiesPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentCookiesPolicy(QWebEngineProfile::PersistentCookiesPolicy)

/*

 */
func (this *QWebEngineProfile) SetPersistentCookiesPolicy(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile26setPersistentCookiesPolicyENS_23PersistentCookiesPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int httpCacheMaximumSize() const

/*

 */
func (this *QWebEngineProfile) HttpCacheMaximumSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile20httpCacheMaximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpCacheMaximumSize(int)

/*

 */
func (this *QWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile23setHttpCacheMaximumSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineCookieStore * cookieStore()

/*

 */
func (this *QWebEngineProfile) CookieStore() *qtwebenginecore.QWebEngineCookieStore /*777 QWebEngineCookieStore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile11cookieStoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtwebenginecore.NewQWebEngineCookieStoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRequestInterceptor(QWebEngineUrlRequestInterceptor *)

/*

 */
func (this *QWebEngineProfile) SetRequestInterceptor(interceptor qtwebenginecore.QWebEngineUrlRequestInterceptor_ITF /*777 QWebEngineUrlRequestInterceptor **/) {
	var convArg0 unsafe.Pointer
	if interceptor != nil && interceptor.QWebEngineUrlRequestInterceptor_PTR() != nil {
		convArg0 = interceptor.QWebEngineUrlRequestInterceptor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile21setRequestInterceptorEP31QWebEngineUrlRequestInterceptor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearAllVisitedLinks()

/*

 */
func (this *QWebEngineProfile) ClearAllVisitedLinks() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile20clearAllVisitedLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool visitedLinksContainsUrl(const QUrl &) const

/*

 */
func (this *QWebEngineProfile) VisitedLinksContainsUrl(url qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile23visitedLinksContainsUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineSettings * settings() const

/*

 */
func (this *QWebEngineProfile) Settings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile8settingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineScriptCollection * scripts() const

/*

 */
func (this *QWebEngineProfile) Scripts() *QWebEngineScriptCollection /*777 QWebEngineScriptCollection **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile7scriptsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineScriptCollectionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] const QWebEngineUrlSchemeHandler * urlSchemeHandler(const QByteArray &) const

/*

 */
func (this *QWebEngineProfile) UrlSchemeHandler(arg0 qtcore.QByteArray_ITF) *qtwebenginecore.QWebEngineUrlSchemeHandler /*777 const QWebEngineUrlSchemeHandler **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile16urlSchemeHandlerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtwebenginecore.NewQWebEngineUrlSchemeHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installUrlSchemeHandler(const QByteArray &, QWebEngineUrlSchemeHandler *)

/*

 */
func (this *QWebEngineProfile) InstallUrlSchemeHandler(scheme qtcore.QByteArray_ITF, arg1 qtwebenginecore.QWebEngineUrlSchemeHandler_ITF /*777 QWebEngineUrlSchemeHandler **/) {
	var convArg0 unsafe.Pointer
	if scheme != nil && scheme.QByteArray_PTR() != nil {
		convArg0 = scheme.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QWebEngineUrlSchemeHandler_PTR() != nil {
		convArg1 = arg1.QWebEngineUrlSchemeHandler_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile23installUrlSchemeHandlerERK10QByteArrayP26QWebEngineUrlSchemeHandler", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeUrlScheme(const QByteArray &)

/*

 */
func (this *QWebEngineProfile) RemoveUrlScheme(scheme qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if scheme != nil && scheme.QByteArray_PTR() != nil {
		convArg0 = scheme.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile15removeUrlSchemeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeUrlSchemeHandler(QWebEngineUrlSchemeHandler *)

/*

 */
func (this *QWebEngineProfile) RemoveUrlSchemeHandler(arg0 qtwebenginecore.QWebEngineUrlSchemeHandler_ITF /*777 QWebEngineUrlSchemeHandler **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineUrlSchemeHandler_PTR() != nil {
		convArg0 = arg0.QWebEngineUrlSchemeHandler_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile22removeUrlSchemeHandlerEP26QWebEngineUrlSchemeHandler", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAllUrlSchemeHandlers()

/*

 */
func (this *QWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile26removeAllUrlSchemeHandlersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearHttpCache()

/*

 */
func (this *QWebEngineProfile) ClearHttpCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile14clearHttpCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpellCheckLanguages(const QStringList &)

/*

 */
func (this *QWebEngineProfile) SetSpellCheckLanguages(languages qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if languages != nil && languages.QStringList_PTR() != nil {
		convArg0 = languages.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile22setSpellCheckLanguagesERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList spellCheckLanguages() const

/*

 */
func (this *QWebEngineProfile) SpellCheckLanguages() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile19spellCheckLanguagesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpellCheckEnabled(bool)

/*

 */
func (this *QWebEngineProfile) SetSpellCheckEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile20setSpellCheckEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSpellCheckEnabled() const

/*

 */
func (this *QWebEngineProfile) IsSpellCheckEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineProfile19isSpellCheckEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWebEngineProfile * defaultProfile()

/*

 */
func (this *QWebEngineProfile) DefaultProfile() *QWebEngineProfile /*777 QWebEngineProfile **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile14defaultProfileEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWebEngineProfile_DefaultProfile() *QWebEngineProfile /*777 QWebEngineProfile **/ {
	var nilthis *QWebEngineProfile
	rv := nilthis.DefaultProfile()
	return rv
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineprofile.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void downloadRequested(QWebEngineDownloadItem *)

/*

 */
func (this *QWebEngineProfile) DownloadRequested(download QWebEngineDownloadItem_ITF /*777 QWebEngineDownloadItem **/) {
	var convArg0 unsafe.Pointer
	if download != nil && download.QWebEngineDownloadItem_PTR() != nil {
		convArg0 = download.QWebEngineDownloadItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineProfile17downloadRequestedEP22QWebEngineDownloadItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QWebEngineProfile__HttpCacheType = int

//
const QWebEngineProfile__MemoryHttpCache QWebEngineProfile__HttpCacheType = 0

//
const QWebEngineProfile__DiskHttpCache QWebEngineProfile__HttpCacheType = 1

//
const QWebEngineProfile__NoCache QWebEngineProfile__HttpCacheType = 2

func (this *QWebEngineProfile) HttpCacheTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineProfile_HttpCacheTypeItemName(val int) string {
	var nilthis *QWebEngineProfile
	return nilthis.HttpCacheTypeItemName(val)
}

/*


 */
type QWebEngineProfile__PersistentCookiesPolicy = int

//
const QWebEngineProfile__NoPersistentCookies QWebEngineProfile__PersistentCookiesPolicy = 0

//
const QWebEngineProfile__AllowPersistentCookies QWebEngineProfile__PersistentCookiesPolicy = 1

//
const QWebEngineProfile__ForcePersistentCookies QWebEngineProfile__PersistentCookiesPolicy = 2

func (this *QWebEngineProfile) PersistentCookiesPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWebEngineProfile_PersistentCookiesPolicyItemName(val int) string {
	var nilthis *QWebEngineProfile
	return nilthis.PersistentCookiesPolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11723() {
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
		qtquickwidgets.KeepMe()
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
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
