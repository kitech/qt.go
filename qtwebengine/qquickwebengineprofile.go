package qtwebengine

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h
// #include <qquickwebengineprofile.h>
// #include <QtWebEngine>

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
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QQuickWebEngineProfile struct {
	*qtcore.QObject
}
type QQuickWebEngineProfile_ITF interface {
	qtcore.QObject_ITF
	QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile
}

func (ptr *QQuickWebEngineProfile) QQuickWebEngineProfile_PTR() *QQuickWebEngineProfile { return ptr }

func (this *QQuickWebEngineProfile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickWebEngineProfile) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickWebEngineProfileFromPointer(cthis unsafe.Pointer) *QQuickWebEngineProfile {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickWebEngineProfile{bcthis0}
}
func (*QQuickWebEngineProfile) NewFromPointer(cthis unsafe.Pointer) *QQuickWebEngineProfile {
	return NewQQuickWebEngineProfileFromPointer(cthis)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickWebEngineProfile) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWebEngineProfile(QObject *)

/*
Constructs a new profile with the parent parent.
*/
func (*QQuickWebEngineProfile) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickWebEngineProfile {
	return NewQQuickWebEngineProfile(parent)
}
func NewQQuickWebEngineProfile(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickWebEngineProfile {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWebEngineProfile(QObject *)

/*
Constructs a new profile with the parent parent.
*/
func (*QQuickWebEngineProfile) NewForInheritp() *QQuickWebEngineProfile {
	return NewQQuickWebEngineProfilep()
}
func NewQQuickWebEngineProfilep() *QQuickWebEngineProfile {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWebEngineProfile")
	return gothis
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWebEngineProfile()

/*

 */
func DeleteQQuickWebEngineProfile(this *QQuickWebEngineProfile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QString storageName() const

/*

 */
func (this *QQuickWebEngineProfile) StorageName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile11storageNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStorageName(const QString &)

/*

 */
func (this *QQuickWebEngineProfile) SetStorageName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile14setStorageNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOffTheRecord() const

/*

 */
func (this *QQuickWebEngineProfile) IsOffTheRecord() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile14isOffTheRecordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOffTheRecord(bool)

/*

 */
func (this *QQuickWebEngineProfile) SetOffTheRecord(offTheRecord bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile15setOffTheRecordEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), offTheRecord)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QString persistentStoragePath() const

/*

 */
func (this *QQuickWebEngineProfile) PersistentStoragePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile21persistentStoragePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentStoragePath(const QString &)

/*

 */
func (this *QQuickWebEngineProfile) SetPersistentStoragePath(path string) {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile24setPersistentStoragePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cachePath() const

/*

 */
func (this *QQuickWebEngineProfile) CachePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile9cachePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCachePath(const QString &)

/*

 */
func (this *QQuickWebEngineProfile) SetCachePath(path string) {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile12setCachePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QString httpUserAgent() const

/*

 */
func (this *QQuickWebEngineProfile) HttpUserAgent() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile13httpUserAgentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpUserAgent(const QString &)

/*

 */
func (this *QQuickWebEngineProfile) SetHttpUserAgent(userAgent string) {
	var tmpArg0 = qtcore.NewQString5(userAgent)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile16setHttpUserAgentERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWebEngineProfile::HttpCacheType httpCacheType() const

/*

 */
func (this *QQuickWebEngineProfile) HttpCacheType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile13httpCacheTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpCacheType(QQuickWebEngineProfile::HttpCacheType)

/*

 */
func (this *QQuickWebEngineProfile) SetHttpCacheType(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile16setHttpCacheTypeENS_13HttpCacheTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWebEngineProfile::PersistentCookiesPolicy persistentCookiesPolicy() const

/*

 */
func (this *QQuickWebEngineProfile) PersistentCookiesPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile23persistentCookiesPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentCookiesPolicy(QQuickWebEngineProfile::PersistentCookiesPolicy)

/*

 */
func (this *QQuickWebEngineProfile) SetPersistentCookiesPolicy(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile26setPersistentCookiesPolicyENS_23PersistentCookiesPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] int httpCacheMaximumSize() const

/*

 */
func (this *QQuickWebEngineProfile) HttpCacheMaximumSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile20httpCacheMaximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpCacheMaximumSize(int)

/*

 */
func (this *QQuickWebEngineProfile) SetHttpCacheMaximumSize(maxSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile23setHttpCacheMaximumSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString httpAcceptLanguage() const

/*

 */
func (this *QQuickWebEngineProfile) HttpAcceptLanguage() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile18httpAcceptLanguageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHttpAcceptLanguage(const QString &)

/*

 */
func (this *QQuickWebEngineProfile) SetHttpAcceptLanguage(httpAcceptLanguage string) {
	var tmpArg0 = qtcore.NewQString5(httpAcceptLanguage)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile21setHttpAcceptLanguageERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineCookieStore * cookieStore() const

/*
Returns the cookie store for this profile.
*/
func (this *QQuickWebEngineProfile) CookieStore() *qtwebenginecore.QWebEngineCookieStore /*777 QWebEngineCookieStore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile11cookieStoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtwebenginecore.NewQWebEngineCookieStoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRequestInterceptor(QWebEngineUrlRequestInterceptor *)

/*
Registers a request interceptor singleton interceptor to intercept URL requests.

The profile does not take ownership of the pointer.

See also QWebEngineUrlRequestInterceptor.
*/
func (this *QQuickWebEngineProfile) SetRequestInterceptor(interceptor qtwebenginecore.QWebEngineUrlRequestInterceptor_ITF /*777 QWebEngineUrlRequestInterceptor **/) {
	var convArg0 unsafe.Pointer
	if interceptor != nil && interceptor.QWebEngineUrlRequestInterceptor_PTR() != nil {
		convArg0 = interceptor.QWebEngineUrlRequestInterceptor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile21setRequestInterceptorEP31QWebEngineUrlRequestInterceptor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] const QWebEngineUrlSchemeHandler * urlSchemeHandler(const QByteArray &) const

/*
Returns the custom URL scheme handler register for the URL scheme scheme.
*/
func (this *QQuickWebEngineProfile) UrlSchemeHandler(arg0 qtcore.QByteArray_ITF) *qtwebenginecore.QWebEngineUrlSchemeHandler /*777 const QWebEngineUrlSchemeHandler **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile16urlSchemeHandlerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtwebenginecore.NewQWebEngineUrlSchemeHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installUrlSchemeHandler(const QByteArray &, QWebEngineUrlSchemeHandler *)

/*
Registers a handler handler for custom URL scheme scheme in the profile.

It is necessary to first register the scheme with QWebEngineUrlScheme::registerScheme at application startup.
*/
func (this *QQuickWebEngineProfile) InstallUrlSchemeHandler(scheme qtcore.QByteArray_ITF, arg1 qtwebenginecore.QWebEngineUrlSchemeHandler_ITF /*777 QWebEngineUrlSchemeHandler **/) {
	var convArg0 unsafe.Pointer
	if scheme != nil && scheme.QByteArray_PTR() != nil {
		convArg0 = scheme.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QWebEngineUrlSchemeHandler_PTR() != nil {
		convArg1 = arg1.QWebEngineUrlSchemeHandler_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile23installUrlSchemeHandlerERK10QByteArrayP26QWebEngineUrlSchemeHandler", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeUrlScheme(const QByteArray &)

/*
Removes the custom URL scheme scheme from the profile.

See also removeUrlSchemeHandler().
*/
func (this *QQuickWebEngineProfile) RemoveUrlScheme(scheme qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if scheme != nil && scheme.QByteArray_PTR() != nil {
		convArg0 = scheme.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile15removeUrlSchemeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeUrlSchemeHandler(QWebEngineUrlSchemeHandler *)

/*
Removes the custom URL scheme handler handler from the profile.

See also removeUrlScheme().
*/
func (this *QQuickWebEngineProfile) RemoveUrlSchemeHandler(arg0 qtwebenginecore.QWebEngineUrlSchemeHandler_ITF /*777 QWebEngineUrlSchemeHandler **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineUrlSchemeHandler_PTR() != nil {
		convArg0 = arg0.QWebEngineUrlSchemeHandler_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile22removeUrlSchemeHandlerEP26QWebEngineUrlSchemeHandler", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAllUrlSchemeHandlers()

/*
Removes all custom URL scheme handlers installed in the profile.
*/
func (this *QQuickWebEngineProfile) RemoveAllUrlSchemeHandlers() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile26removeAllUrlSchemeHandlersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearHttpCache()

/*
Removes the profile's cache entries.

This function was introduced in  Qt 5.7.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also WebEngineProfile::clearHttpCache.
*/
func (this *QQuickWebEngineProfile) ClearHttpCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile14clearHttpCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpellCheckLanguages(const QStringList &)

/*

 */
func (this *QQuickWebEngineProfile) SetSpellCheckLanguages(languages qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if languages != nil && languages.QStringList_PTR() != nil {
		convArg0 = languages.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile22setSpellCheckLanguagesERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList spellCheckLanguages() const

/*
Returns the list of languages used by the spell checker.

This function was introduced in  Qt 5.8.

Note: Getter function for property spellCheckLanguages.

See also setSpellCheckLanguages().
*/
func (this *QQuickWebEngineProfile) SpellCheckLanguages() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile19spellCheckLanguagesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpellCheckEnabled(bool)

/*

 */
func (this *QQuickWebEngineProfile) SetSpellCheckEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile20setSpellCheckEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSpellCheckEnabled() const

/*

 */
func (this *QQuickWebEngineProfile) IsSpellCheckEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QQuickWebEngineProfile19isSpellCheckEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:140
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQuickWebEngineProfile * defaultProfile()

/*
Returns the default profile.

The default profile uses the storage name "Default".

See also storageName().
*/
func (this *QQuickWebEngineProfile) DefaultProfile() *QQuickWebEngineProfile /*777 QQuickWebEngineProfile **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile14defaultProfileEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQuickWebEngineProfile_DefaultProfile() *QQuickWebEngineProfile /*777 QQuickWebEngineProfile **/ {
	var nilthis *QQuickWebEngineProfile
	rv := nilthis.DefaultProfile()
	return rv
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void storageNameChanged()

/*

 */
func (this *QQuickWebEngineProfile) StorageNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile18storageNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void offTheRecordChanged()

/*

 */
func (this *QQuickWebEngineProfile) OffTheRecordChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile19offTheRecordChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void persistentStoragePathChanged()

/*

 */
func (this *QQuickWebEngineProfile) PersistentStoragePathChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile28persistentStoragePathChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cachePathChanged()

/*

 */
func (this *QQuickWebEngineProfile) CachePathChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile16cachePathChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void httpUserAgentChanged()

/*

 */
func (this *QQuickWebEngineProfile) HttpUserAgentChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile20httpUserAgentChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void httpCacheTypeChanged()

/*

 */
func (this *QQuickWebEngineProfile) HttpCacheTypeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile20httpCacheTypeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void persistentCookiesPolicyChanged()

/*

 */
func (this *QQuickWebEngineProfile) PersistentCookiesPolicyChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile30persistentCookiesPolicyChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void httpCacheMaximumSizeChanged()

/*

 */
func (this *QQuickWebEngineProfile) HttpCacheMaximumSizeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile27httpCacheMaximumSizeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void httpAcceptLanguageChanged()

/*

 */
func (this *QQuickWebEngineProfile) HttpAcceptLanguageChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile25httpAcceptLanguageChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void spellCheckLanguagesChanged()

/*

 */
func (this *QQuickWebEngineProfile) SpellCheckLanguagesChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile26spellCheckLanguagesChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngine/qquickwebengineprofile.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void spellCheckEnabledChanged()

/*

 */
func (this *QQuickWebEngineProfile) SpellCheckEnabledChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQuickWebEngineProfile24spellCheckEnabledChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the HTTP cache type:


*/
type QQuickWebEngineProfile__HttpCacheType = int

// Use an in-memory cache. This is the default if off-the-record is set.
const QQuickWebEngineProfile__MemoryHttpCache QQuickWebEngineProfile__HttpCacheType = 0

// Use a disk cache. This is the default if off-the-record is not set. Falls back to MemoryHttpCache if off-the-record is set.
const QQuickWebEngineProfile__DiskHttpCache QQuickWebEngineProfile__HttpCacheType = 1

//
const QQuickWebEngineProfile__NoCache QQuickWebEngineProfile__HttpCacheType = 2

func (this *QQuickWebEngineProfile) HttpCacheTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWebEngineProfile_HttpCacheTypeItemName(val int) string {
	var nilthis *QQuickWebEngineProfile
	return nilthis.HttpCacheTypeItemName(val)
}

/*
This enum describes policy for cookie persistency:


*/
type QQuickWebEngineProfile__PersistentCookiesPolicy = int

// Both session and persistent cookies are stored in memory. This is the only setting possible if off-the-record is set or no persistent data path is available.
const QQuickWebEngineProfile__NoPersistentCookies QQuickWebEngineProfile__PersistentCookiesPolicy = 0

// Cookies marked persistent are saved to and restored from disk, whereas session cookies are only stored to disk for crash recovery. This is the default setting.
const QQuickWebEngineProfile__AllowPersistentCookies QQuickWebEngineProfile__PersistentCookiesPolicy = 1

// Both session and persistent cookies are saved to and restored from disk.
const QQuickWebEngineProfile__ForcePersistentCookies QQuickWebEngineProfile__PersistentCookiesPolicy = 2

func (this *QQuickWebEngineProfile) PersistentCookiesPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWebEngineProfile_PersistentCookiesPolicyItemName(val int) string {
	var nilthis *QQuickWebEngineProfile
	return nilthis.PersistentCookiesPolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11687() {
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
		qtpositioning.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
