package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h
// #include <qnetworkaccessmanager.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// QNetworkReply * createRequest(QNetworkAccessManager::Operation, const QNetworkRequest &, QIODevice *)
func (this *QNetworkAccessManager) InheritCreateRequest(f func(op int, request *QNetworkRequest, outgoingData *qtcore.QIODevice /*777 QIODevice **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createRequest", f)
}

// QStringList supportedSchemesImplementation()
func (this *QNetworkAccessManager) InheritSupportedSchemesImplementation(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "supportedSchemesImplementation", f)
}

/*

 */
type QNetworkAccessManager struct {
	*qtcore.QObject
}
type QNetworkAccessManager_ITF interface {
	qtcore.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
}

func (ptr *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager { return ptr }

func (this *QNetworkAccessManager) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QNetworkAccessManager) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQNetworkAccessManagerFromPointer(cthis unsafe.Pointer) *QNetworkAccessManager {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QNetworkAccessManager{bcthis0}
}
func (*QNetworkAccessManager) NewFromPointer(cthis unsafe.Pointer) *QNetworkAccessManager {
	return NewQNetworkAccessManagerFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QNetworkAccessManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkAccessManager(QObject *)

/*
Constructs a QNetworkAccessManager object that is the center of the Network Access API and sets parent as the parent object.
*/
func (*QNetworkAccessManager) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkAccessManager {
	return NewQNetworkAccessManager(parent)
}
func NewQNetworkAccessManager(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkAccessManager {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManagerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkAccessManager")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkAccessManager(QObject *)

/*
Constructs a QNetworkAccessManager object that is the center of the Network Access API and sets parent as the parent object.
*/
func (*QNetworkAccessManager) NewForInherit__() *QNetworkAccessManager {
	return NewQNetworkAccessManager__()
}
func NewQNetworkAccessManager__() *QNetworkAccessManager {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManagerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkAccessManager")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkAccessManager()

/*

 */
func DeleteQNetworkAccessManager(this *QNetworkAccessManager) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManagerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedSchemes() const

/*
Lists all the URL schemes supported by the access manager.

This function was introduced in  Qt 5.2.

See also supportedSchemesImplementation().
*/
func (this *QNetworkAccessManager) SupportedSchemes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager16supportedSchemesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearAccessCache()

/*
Flushes the internal cache of authentication data and network connections.

This function is useful for doing auto tests.

This function was introduced in  Qt 5.0.

See also clearConnectionCache().
*/
func (this *QNetworkAccessManager) ClearAccessCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager16clearAccessCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearConnectionCache()

/*
Flushes the internal cache of network connections. In contrast to clearAccessCache() the authentication data is preserved.

This function was introduced in  Qt 5.9.

See also clearAccessCache().
*/
func (this *QNetworkAccessManager) ClearConnectionCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager20clearConnectionCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy proxy() const

/*
Returns the QNetworkProxy that the requests sent using this QNetworkAccessManager object will use. The default value for the proxy is QNetworkProxy::DefaultProxy.

See also setProxy(), setProxyFactory(), and proxyAuthenticationRequired().
*/
func (this *QNetworkAccessManager) Proxy() *QNetworkProxy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager5proxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxy(const QNetworkProxy &)

/*
Sets the proxy to be used in future requests to be proxy. This does not affect requests that have already been sent. The proxyAuthenticationRequired() signal will be emitted if the proxy requests authentication.

A proxy set with this function will be used for all requests issued by QNetworkAccessManager. In some cases, it might be necessary to select different proxies depending on the type of request being sent or the destination host. If that's the case, you should consider using setProxyFactory().

See also proxy() and proxyAuthenticationRequired().
*/
func (this *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if proxy != nil && proxy.QNetworkProxy_PTR() != nil {
		convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setProxyERK13QNetworkProxy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxyFactory * proxyFactory() const

/*
Returns the proxy factory that this QNetworkAccessManager object is using to determine the proxies to be used for requests.

Note that the pointer returned by this function is managed by QNetworkAccessManager and could be deleted at any time.

This function was introduced in  Qt 4.5.

See also setProxyFactory() and proxy().
*/
func (this *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory /*777 QNetworkProxyFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager12proxyFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkProxyFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxyFactory(QNetworkProxyFactory *)

/*
Sets the proxy factory for this class to be factory. A proxy factory is used to determine a more specific list of proxies to be used for a given request, instead of trying to use the same proxy value for all requests.

All queries sent by QNetworkAccessManager will have type QNetworkProxyQuery::UrlRequest.

For example, a proxy factory could apply the following rules:


if the target address is in the local network (for example, if the hostname contains no dots or if it's an IP address in the organization's range), return QNetworkProxy::NoProxy
if the request is FTP, return an FTP proxy
if the request is HTTP or HTTPS, then return an HTTP proxy
otherwise, return a SOCKSv5 proxy server


The lifetime of the object factory will be managed by QNetworkAccessManager. It will delete the object when necessary.

Note: If a specific proxy is set with setProxy(), the factory will not be used.

This function was introduced in  Qt 4.5.

See also proxyFactory(), setProxy(), and QNetworkProxyQuery.
*/
func (this *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF /*777 QNetworkProxyFactory **/) {
	var convArg0 unsafe.Pointer
	if factory != nil && factory.QNetworkProxyFactory_PTR() != nil {
		convArg0 = factory.QNetworkProxyFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager15setProxyFactoryEP20QNetworkProxyFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractNetworkCache * cache() const

/*
Returns the cache that is used to store data obtained from the network.

This function was introduced in  Qt 4.5.

See also setCache().
*/
func (this *QNetworkAccessManager) Cache() *QAbstractNetworkCache /*777 QAbstractNetworkCache **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager5cacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCache(QAbstractNetworkCache *)

/*
Sets the manager's network cache to be the cache specified. The cache is used for all requests dispatched by the manager.

Use this function to set the network cache object to a class that implements additional features, like saving the cookies to permanent storage.

Note: QNetworkAccessManager takes ownership of the cache object.

QNetworkAccessManager by default does not have a set cache. Qt provides a simple disk cache, QNetworkDiskCache, which can be used.

This function was introduced in  Qt 4.5.

See also cache() and QNetworkRequest::CacheLoadControl.
*/
func (this *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF /*777 QAbstractNetworkCache **/) {
	var convArg0 unsafe.Pointer
	if cache != nil && cache.QAbstractNetworkCache_PTR() != nil {
		convArg0 = cache.QAbstractNetworkCache_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setCacheEP21QAbstractNetworkCache", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkCookieJar * cookieJar() const

/*
Returns the QNetworkCookieJar that is used to store cookies obtained from the network as well as cookies that are about to be sent.

See also setCookieJar().
*/
func (this *QNetworkAccessManager) CookieJar() *QNetworkCookieJar /*777 QNetworkCookieJar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager9cookieJarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCookieJar(QNetworkCookieJar *)

/*
Sets the manager's cookie jar to be the cookieJar specified. The cookie jar is used by all requests dispatched by the manager.

Use this function to set the cookie jar object to a class that implements additional features, like saving the cookies to permanent storage.

Note: QNetworkAccessManager takes ownership of the cookieJar object.

If cookieJar is in the same thread as this QNetworkAccessManager, it will set the parent of the cookieJar so that the cookie jar is deleted when this object is deleted as well. If you want to share cookie jars between different QNetworkAccessManager objects, you may want to set the cookie jar's parent to 0 after calling this function.

QNetworkAccessManager by default does not implement any cookie policy of its own: it accepts all cookies sent by the server, as long as they are well formed and meet the minimum security requirements (cookie domain matches the request's and cookie path matches the request's). In order to implement your own security policy, override the QNetworkCookieJar::cookiesForUrl() and QNetworkCookieJar::setCookiesFromUrl() virtual functions. Those functions are called by QNetworkAccessManager when it detects a new cookie.

See also cookieJar(), QNetworkCookieJar::cookiesForUrl(), and QNetworkCookieJar::setCookiesFromUrl().
*/
func (this *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF /*777 QNetworkCookieJar **/) {
	var convArg0 unsafe.Pointer
	if cookieJar != nil && cookieJar.QNetworkCookieJar_PTR() != nil {
		convArg0 = cookieJar.QNetworkCookieJar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager12setCookieJarEP17QNetworkCookieJar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrictTransportSecurityEnabled(bool)

/*
If enabled is true, QNetworkAccessManager follows the HTTP Strict Transport Security policy (HSTS, RFC6797). When processing a request, QNetworkAccessManager automatically replaces the "http" scheme with "https" and uses a secure transport for HSTS hosts. If it's set explicitly, port 80 is replaced by port 443.

When HSTS is enabled, for each HTTP response containing HSTS header and received over a secure transport, QNetworkAccessManager will update its HSTS cache, either remembering a host with a valid policy or removing a host with an expired or disabled HSTS policy.

This function was introduced in  Qt 5.9.

See also isStrictTransportSecurityEnabled().
*/
func (this *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager33setStrictTransportSecurityEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStrictTransportSecurityEnabled() const

/*
Returns true if HTTP Strict Transport Security (HSTS) was enabled. By default HSTS is disabled.

This function was introduced in  Qt 5.9.

See also setStrictTransportSecurityEnabled().
*/
func (this *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager32isStrictTransportSecurityEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enableStrictTransportSecurityStore(bool, const QString &)

/*
If enabled is true, the internal HSTS cache will use a persistent store to read and write HSTS policies. storeDir defines where this store will be located. The default location is defined by QStandardPaths::CacheLocation. If there is no writable QStandartPaths::CacheLocation and storeDir is an empty string, the store will be located in the program's working directory.

Note: If HSTS cache already contains HSTS policies by the time persistent store is enabled, these policies will be preserved in the store. In case both cache and store contain the same known hosts, policies from cache are considered to be more up-to-date (and thus will overwrite the previous values in the store). If this behavior is undesired, enable HSTS store before enabling Strict Tranport Security. By default, the persistent store of HSTS policies is disabled.

This function was introduced in  Qt 5.10.

See also isStrictTransportSecurityStoreEnabled(), setStrictTransportSecurityEnabled(), and QStandardPaths::standardLocations().
*/
func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool, storeDir string) {
	var tmpArg1 = qtcore.NewQString_5(storeDir)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager34enableStrictTransportSecurityStoreEbRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enableStrictTransportSecurityStore(bool, const QString &)

/*
If enabled is true, the internal HSTS cache will use a persistent store to read and write HSTS policies. storeDir defines where this store will be located. The default location is defined by QStandardPaths::CacheLocation. If there is no writable QStandartPaths::CacheLocation and storeDir is an empty string, the store will be located in the program's working directory.

Note: If HSTS cache already contains HSTS policies by the time persistent store is enabled, these policies will be preserved in the store. In case both cache and store contain the same known hosts, policies from cache are considered to be more up-to-date (and thus will overwrite the previous values in the store). If this behavior is undesired, enable HSTS store before enabling Strict Tranport Security. By default, the persistent store of HSTS policies is disabled.

This function was introduced in  Qt 5.10.

See also isStrictTransportSecurityStoreEnabled(), setStrictTransportSecurityEnabled(), and QStandardPaths::standardLocations().
*/
func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore__(enabled bool) {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager34enableStrictTransportSecurityStoreEbRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStrictTransportSecurityStoreEnabled() const

/*
Returns true if HSTS cache uses a permanent store to load and store HSTS policies.

This function was introduced in  Qt 5.10.

See also enableStrictTransportSecurityStore().
*/
func (this *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager37isStrictTransportSecurityStoreEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * head(const QNetworkRequest &)

/*
Posts a request to obtain the network headers for request and returns a new QNetworkReply object which will contain such headers.

The function is named after the HTTP request associated (HEAD).
*/
func (this *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4headERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * get(const QNetworkRequest &)

/*
Posts a request to obtain the contents of the target request and returns a new QNetworkReply object opened for reading which emits the readyRead() signal whenever new data arrives.

The contents as well as associated headers will be downloaded.

See also post(), put(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3getERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, QIODevice *)

/*
Sends an HTTP POST request to the destination specified by request and returns a new QNetworkReply object opened for reading that will contain the reply sent by the server. The contents of the data device will be uploaded to the server.

data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: Sending a POST request on protocols other than HTTP and HTTPS is undefined and will probably fail.

See also get(), put(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if data != nil && data.QIODevice_PTR() != nil {
		convArg1 = data.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, const QByteArray &)

/*
Sends an HTTP POST request to the destination specified by request and returns a new QNetworkReply object opened for reading that will contain the reply sent by the server. The contents of the data device will be uploaded to the server.

data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: Sending a POST request on protocols other than HTTP and HTTPS is undefined and will probably fail.

See also get(), put(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Post_1(request QNetworkRequest_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg1 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:137
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, QHttpMultiPart *)

/*
Sends an HTTP POST request to the destination specified by request and returns a new QNetworkReply object opened for reading that will contain the reply sent by the server. The contents of the data device will be uploaded to the server.

data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: Sending a POST request on protocols other than HTTP and HTTPS is undefined and will probably fail.

See also get(), put(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Post_2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, QIODevice *)

/*
Uploads the contents of data to the destination request and returnes a new QNetworkReply object that will be open for reply.

data must be opened for reading when this function is called and must remain valid until the finished() signal is emitted for this reply.

Whether anything will be available for reading from the returned object is protocol dependent. For HTTP, the server may send a small HTML page indicating the upload was successful (or not). Other protocols will probably have content in their replies.

Note: For HTTP, this request will send a PUT request, which most servers do not allow. Form upload mechanisms, including that of uploading files through HTML forms, use the POST mechanism.

See also get(), post(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if data != nil && data.QIODevice_PTR() != nil {
		convArg1 = data.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:139
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, const QByteArray &)

/*
Uploads the contents of data to the destination request and returnes a new QNetworkReply object that will be open for reply.

data must be opened for reading when this function is called and must remain valid until the finished() signal is emitted for this reply.

Whether anything will be available for reading from the returned object is protocol dependent. For HTTP, the server may send a small HTML page indicating the upload was successful (or not). Other protocols will probably have content in their replies.

Note: For HTTP, this request will send a PUT request, which most servers do not allow. Form upload mechanisms, including that of uploading files through HTML forms, use the POST mechanism.

See also get(), post(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Put_1(request QNetworkRequest_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg1 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:140
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, QHttpMultiPart *)

/*
Uploads the contents of data to the destination request and returnes a new QNetworkReply object that will be open for reply.

data must be opened for reading when this function is called and must remain valid until the finished() signal is emitted for this reply.

Whether anything will be available for reading from the returned object is protocol dependent. For HTTP, the server may send a small HTML page indicating the upload was successful (or not). Other protocols will probably have content in their replies.

Note: For HTTP, this request will send a PUT request, which most servers do not allow. Form upload mechanisms, including that of uploading files through HTML forms, use the POST mechanism.

See also get(), post(), deleteResource(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) Put_2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * deleteResource(const QNetworkRequest &)

/*
Sends a request to delete the resource identified by the URL of request.

Note: This feature is currently available for HTTP only, performing an HTTP DELETE request.

This function was introduced in  Qt 4.6.

See also get(), post(), put(), and sendCustomRequest().
*/
func (this *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager14deleteResourceERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QIODevice *)

/*
Sends a custom request to the server identified by the URL of request.

It is the user's responsibility to send a verb to the server that is valid according to the HTTP specification.

This method provides means to send verbs other than the common ones provided via get() or post() etc., for instance sending an HTTP OPTIONS command.

If data is not empty, the contents of the data device will be uploaded to the server; in that case, data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: This feature is currently available for HTTP(S) only.

This function was introduced in  Qt 4.7.

See also get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if verb != nil && verb.QByteArray_PTR() != nil {
		convArg1 = verb.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QIODevice_PTR() != nil {
		convArg2 = data.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QIODevice *)

/*
Sends a custom request to the server identified by the URL of request.

It is the user's responsibility to send a verb to the server that is valid according to the HTTP specification.

This method provides means to send verbs other than the common ones provided via get() or post() etc., for instance sending an HTTP OPTIONS command.

If data is not empty, the contents of the data device will be uploaded to the server; in that case, data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: This feature is currently available for HTTP(S) only.

This function was introduced in  Qt 4.7.

See also get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) SendCustomRequest__(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if verb != nil && verb.QByteArray_PTR() != nil {
		convArg1 = verb.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QIODevice *=Pointer, QIODevice=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:143
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, const QByteArray &)

/*
Sends a custom request to the server identified by the URL of request.

It is the user's responsibility to send a verb to the server that is valid according to the HTTP specification.

This method provides means to send verbs other than the common ones provided via get() or post() etc., for instance sending an HTTP OPTIONS command.

If data is not empty, the contents of the data device will be uploaded to the server; in that case, data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: This feature is currently available for HTTP(S) only.

This function was introduced in  Qt 4.7.

See also get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) SendCustomRequest_1(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if verb != nil && verb.QByteArray_PTR() != nil {
		convArg1 = verb.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg2 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:144
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QHttpMultiPart *)

/*
Sends a custom request to the server identified by the URL of request.

It is the user's responsibility to send a verb to the server that is valid according to the HTTP specification.

This method provides means to send verbs other than the common ones provided via get() or post() etc., for instance sending an HTTP OPTIONS command.

If data is not empty, the contents of the data device will be uploaded to the server; in that case, data must be open for reading and must remain valid until the finished() signal is emitted for this reply.

Note: This feature is currently available for HTTP(S) only.

This function was introduced in  Qt 4.7.

See also get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) SendCustomRequest_2(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg0 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if verb != nil && verb.QByteArray_PTR() != nil {
		convArg1 = verb.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if multiPart != nil && multiPart.QHttpMultiPart_PTR() != nil {
		convArg2 = multiPart.QHttpMultiPart_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setConfiguration(const QNetworkConfiguration &)

/*
Sets the network configuration that will be used when creating the network session to config.

The network configuration is used to create and open a network session before any request that requires network access is process. If no network configuration is explicitly set via this function the network configuration returned by QNetworkConfigurationManager::defaultConfiguration() will be used.

To restore the default network configuration set the network configuration to the value returned from QNetworkConfigurationManager::defaultConfiguration().

Setting a network configuration means that the QNetworkAccessManager instance will only be using the specified one. In particular, if the default network configuration changes (upon e.g. Wifi being available), this new configuration needs to be enabled manually if desired.


  QNetworkConfigurationManager manager;
  networkAccessManager->setConfiguration(manager.defaultConfiguration());



If an invalid network configuration is set, a network session will not be created. In this case network requests will be processed regardless, but may fail. For example:


  networkAccessManager->setConfiguration(QNetworkConfiguration());



This function was introduced in  Qt 4.7.

See also configuration() and QNetworkSession.
*/
func (this *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QNetworkConfiguration_PTR() != nil {
		convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager16setConfigurationERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration configuration() const

/*
Returns the network configuration that will be used to create the network session which will be used when processing network requests.

This function was introduced in  Qt 4.7.

See also setConfiguration() and activeConfiguration().
*/
func (this *QNetworkAccessManager) Configuration() *QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager13configurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration activeConfiguration() const

/*
Returns the current active network configuration.

If the network configuration returned by configuration() is of type QNetworkConfiguration::ServiceNetwork this function will return the current active child network configuration of that configuration. Otherwise returns the same network configuration as configuration().

Use this function to return the actual network configuration currently in use by the network session.

This function was introduced in  Qt 4.7.

See also configuration().
*/
func (this *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager19activeConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNetworkAccessible(QNetworkAccessManager::NetworkAccessibility)

/*
Overrides the reported network accessibility. If accessible is NotAccessible the reported network accessiblity will always be NotAccessible. Otherwise the reported network accessibility will reflect the actual device state.

This function was introduced in  Qt 4.7.

Note: Setter function for property networkAccessible.

See also networkAccessible().
*/
func (this *QNetworkAccessManager) SetNetworkAccessible(accessible int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager20setNetworkAccessibleENS_20NetworkAccessibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkAccessManager::NetworkAccessibility networkAccessible() const

/*
Returns the current network accessibility.

This function was introduced in  Qt 4.7.

Note: Getter function for property networkAccessible.

See also setNetworkAccessible().
*/
func (this *QNetworkAccessManager) NetworkAccessible() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager17networkAccessibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QSslConfiguration &)

/*
Initiates a connection to the host given by hostName at port port, using sslConfiguration. This function is useful to complete the TCP and SSL handshake to a host before the HTTPS request is made, resulting in a lower network latency.

Note: Preconnecting a SPDY connection can be done by calling setAllowedNextProtocols() on sslConfiguration with QSslConfiguration::NextProtocolSpdy3_0 contained in the list of allowed protocols. When using SPDY, one single connection per host is enough, i.e. calling this method multiple times per host will not result in faster network transactions.

Note: This function has no possibility to report errors.

This function was introduced in  Qt 5.2.

See also connectToHost(), get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 unsafe.Pointer
	if sslConfiguration != nil && sslConfiguration.QSslConfiguration_PTR() != nil {
		convArg2 = sslConfiguration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22connectToHostEncryptedERK7QStringtRK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QSslConfiguration &)

/*
Initiates a connection to the host given by hostName at port port, using sslConfiguration. This function is useful to complete the TCP and SSL handshake to a host before the HTTPS request is made, resulting in a lower network latency.

Note: Preconnecting a SPDY connection can be done by calling setAllowedNextProtocols() on sslConfiguration with QSslConfiguration::NextProtocolSpdy3_0 contained in the list of allowed protocols. When using SPDY, one single connection per host is enough, i.e. calling this method multiple times per host will not result in faster network transactions.

Note: This function has no possibility to report errors.

This function was introduced in  Qt 5.2.

See also connectToHost(), get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) ConnectToHostEncrypted__(hostName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(443)
	// arg: 2, const QSslConfiguration &=LValueReference, QSslConfiguration=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22connectToHostEncryptedERK7QStringtRK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QSslConfiguration &)

/*
Initiates a connection to the host given by hostName at port port, using sslConfiguration. This function is useful to complete the TCP and SSL handshake to a host before the HTTPS request is made, resulting in a lower network latency.

Note: Preconnecting a SPDY connection can be done by calling setAllowedNextProtocols() on sslConfiguration with QSslConfiguration::NextProtocolSpdy3_0 contained in the list of allowed protocols. When using SPDY, one single connection per host is enough, i.e. calling this method multiple times per host will not result in faster network transactions.

Note: This function has no possibility to report errors.

This function was introduced in  Qt 5.2.

See also connectToHost(), get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) ConnectToHostEncrypted__1(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, const QSslConfiguration &=LValueReference, QSslConfiguration=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22connectToHostEncryptedERK7QStringtRK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16)

/*
Initiates a connection to the host given by hostName at port port. This function is useful to complete the TCP handshake to a host before the HTTP request is made, resulting in a lower network latency.

Note: This function has no possibility to report errors.

This function was introduced in  Qt 5.2.

See also connectToHostEncrypted(), get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13connectToHostERK7QStringt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16)

/*
Initiates a connection to the host given by hostName at port port. This function is useful to complete the TCP handshake to a host before the HTTP request is made, resulting in a lower network latency.

Note: This function has no possibility to report errors.

This function was introduced in  Qt 5.2.

See also connectToHostEncrypted(), get(), post(), put(), and deleteResource().
*/
func (this *QNetworkAccessManager) ConnectToHost__(hostName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(80)
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13connectToHostERK7QStringt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRedirectPolicy(QNetworkRequest::RedirectPolicy)

/*
Sets the manager's redirect policy to be the policy specified. This policy will affect all subsequent requests created by the manager.

Use this function to enable or disable HTTP redirects on the manager's level.

Note: When creating a request QNetworkRequest::RedirectAttributePolicy has the highest priority, next by priority is QNetworkRequest::FollowRedirectsAttribute. Finally, the manager's policy has the lowest priority.

For backwards compatibility the default value is QNetworkRequest::ManualRedirectPolicy. This may change in the future and some type of auto-redirect policy will become the default; clients relying on manual redirect handling are encouraged to set this policy explicitly in their code.

This function was introduced in  Qt 5.9.

See also redirectPolicy(), QNetworkRequest::RedirectPolicy, and QNetworkRequest::FollowRedirectsAttribute.
*/
func (this *QNetworkAccessManager) SetRedirectPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17setRedirectPolicyEN15QNetworkRequest14RedirectPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkRequest::RedirectPolicy redirectPolicy() const

/*
Returns the redirect policy that is used when creating new requests.

This function was introduced in  Qt 5.9.

See also setRedirectPolicy() and QNetworkRequest::RedirectPolicy.
*/
func (this *QNetworkAccessManager) RedirectPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager14redirectPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void proxyAuthenticationRequired(const QNetworkProxy &, QAuthenticator *)

/*
This signal is emitted whenever a proxy requests authentication and QNetworkAccessManager cannot find a valid, cached credential. The slot connected to this signal should fill in the credentials for the proxy proxy in the authenticator object.

QNetworkAccessManager will cache the credentials internally. The next time the proxy requests authentication, QNetworkAccessManager will automatically send the same credential without emitting the proxyAuthenticationRequired signal again.

If the proxy rejects the credentials, QNetworkAccessManager will emit the signal again.

See also proxy(), setProxy(), and authenticationRequired().
*/
func (this *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if proxy != nil && proxy.QNetworkProxy_PTR() != nil {
		convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QAuthenticator_PTR() != nil {
		convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager27proxyAuthenticationRequiredERK13QNetworkProxyP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void authenticationRequired(QNetworkReply *, QAuthenticator *)

/*
This signal is emitted whenever a final server requests authentication before it delivers the requested contents. The slot connected to this signal should fill the credentials for the contents (which can be determined by inspecting the reply object) in the authenticator object.

QNetworkAccessManager will cache the credentials internally and will send the same values if the server requires authentication again, without emitting the authenticationRequired() signal. If it rejects the credentials, this signal will be emitted again.

Note: To have the request not send credentials you must not call setUser() or setPassword() on the authenticator object. This will result in the finished() signal being emitted with a QNetworkReply with error AuthenticationRequiredError.

Note: It is not possible to use a QueuedConnection to connect to this signal, as the connection will fail if the authenticator has not been filled in with new information when the signal returns.

See also proxyAuthenticationRequired(), QAuthenticator::setUser(), and QAuthenticator::setPassword().
*/
func (this *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF /*777 QNetworkReply **/, authenticator QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if reply != nil && reply.QNetworkReply_PTR() != nil {
		convArg0 = reply.QNetworkReply_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QAuthenticator_PTR() != nil {
		convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22authenticationRequiredEP13QNetworkReplyP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(QNetworkReply *)

/*
This signal is emitted whenever a pending network reply is finished. The reply parameter will contain a pointer to the reply that has just finished. This signal is emitted in tandem with the QNetworkReply::finished() signal.

See QNetworkReply::finished() for information on the status that the object will be in.

Note: Do not delete the reply object in the slot connected to this signal. Use deleteLater().

See also QNetworkReply::finished() and QNetworkReply::error().
*/
func (this *QNetworkAccessManager) Finished(reply QNetworkReply_ITF /*777 QNetworkReply **/) {
	var convArg0 unsafe.Pointer
	if reply != nil && reply.QNetworkReply_PTR() != nil {
		convArg0 = reply.QNetworkReply_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8finishedEP13QNetworkReply", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted(QNetworkReply *)

/*
This signal is emitted when an SSL/TLS session has successfully completed the initial handshake. At this point, no user data has been transmitted. The signal can be used to perform additional checks on the certificate chain, for example to notify users when the certificate for a website has changed. The reply parameter specifies which network reply is responsible. If the reply does not match the expected criteria then it should be aborted by calling QNetworkReply::abort() by a slot connected to this signal. The SSL configuration in use can be inspected using the QNetworkReply::sslConfiguration() method.

Internally, QNetworkAccessManager may open multiple connections to a server, in order to allow it process requests in parallel. These connections may be reused, which means that the encrypted() signal would not be emitted. This means that you are only guaranteed to receive this signal for the first connection to a site in the lifespan of the QNetworkAccessManager.

This function was introduced in  Qt 5.1.

See also QSslSocket::encrypted() and QNetworkReply::encrypted().
*/
func (this *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF /*777 QNetworkReply **/) {
	var convArg0 unsafe.Pointer
	if reply != nil && reply.QNetworkReply_PTR() != nil {
		convArg0 = reply.QNetworkReply_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager9encryptedEP13QNetworkReply", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QNetworkReply *, QSslPreSharedKeyAuthenticator *)

/*
This signal is emitted if the SSL/TLS handshake negotiates a PSK ciphersuite, and therefore a PSK authentication is then required. The reply object is the QNetworkReply that is negotiating such ciphersuites.

When using PSK, the client must send to the server a valid identity and a valid pre shared key, in order for the SSL handshake to continue. Applications can provide this information in a slot connected to this signal, by filling in the passed authenticator object according to their needs.

Note: Ignoring this signal, or failing to provide the required credentials, will cause the handshake to fail, and therefore the connection to be aborted.

Note: The authenticator object is owned by the reply and must not be deleted by the application.

This function was introduced in  Qt 5.5.

See also QSslPreSharedKeyAuthenticator.
*/
func (this *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF /*777 QNetworkReply **/, authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if reply != nil && reply.QNetworkReply_PTR() != nil {
		convArg0 = reply.QNetworkReply_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg1 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager34preSharedKeyAuthenticationRequiredEP13QNetworkReplyP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void networkSessionConnected()

/*

 */
func (this *QNetworkAccessManager) NetworkSessionConnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager23networkSessionConnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void networkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility)

/*
This signal is emitted when the value of the networkAccessible property changes. accessible is the new network accessibility.

Note: Notifier signal for property networkAccessible.
*/
func (this *QNetworkAccessManager) NetworkAccessibleChanged(accessible int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager24networkAccessibleChangedENS_20NetworkAccessibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QNetworkReply * createRequest(QNetworkAccessManager::Operation, const QNetworkRequest &, QIODevice *)

/*
Returns a new QNetworkReply object to handle the operation op and request originalReq. The device outgoingData is always 0 for Get and Head requests, but is the value passed to post() and put() in those operations (the QByteArray variants will pass a QBuffer object).

The default implementation calls QNetworkCookieJar::cookiesForUrl() on the cookie jar set with setCookieJar() to obtain the cookies to be sent to the remote server.

The returned object must be in an open state.
*/
func (this *QNetworkAccessManager) CreateRequest(op int, request QNetworkRequest_ITF, outgoingData qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg1 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg1 = request.QNetworkRequest_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if outgoingData != nil && outgoingData.QIODevice_PTR() != nil {
		convArg2 = outgoingData.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13createRequestENS_9OperationERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QNetworkReply * createRequest(QNetworkAccessManager::Operation, const QNetworkRequest &, QIODevice *)

/*
Returns a new QNetworkReply object to handle the operation op and request originalReq. The device outgoingData is always 0 for Get and Head requests, but is the value passed to post() and put() in those operations (the QByteArray variants will pass a QBuffer object).

The default implementation calls QNetworkCookieJar::cookiesForUrl() on the cookie jar set with setCookieJar() to obtain the cookies to be sent to the remote server.

The returned object must be in an open state.
*/
func (this *QNetworkAccessManager) CreateRequest__(op int, request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg1 unsafe.Pointer
	if request != nil && request.QNetworkRequest_PTR() != nil {
		convArg1 = request.QNetworkRequest_PTR().GetCthis()
	}
	// arg: 2, QIODevice *=Pointer, QIODevice=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13createRequestENS_9OperationERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:187
// index:0
// Protected Visibility=Default Availability=Available
// [8] QStringList supportedSchemesImplementation() const

/*
Lists all the URL schemes supported by the access manager.

You should not call this function directly; use QNetworkAccessManager::supportedSchemes() instead.

Reimplement this slot to provide your own supported schemes in a QNetworkAccessManager subclass. It is for instance necessary when your subclass provides support for new protocols.

Because of binary compatibility constraints, the supportedSchemes() method (introduced in Qt 5.2) is not virtual. Instead, supportedSchemes() will dynamically detect and call this slot.

This function was introduced in  Qt 5.2.

See also supportedSchemes().
*/
func (this *QNetworkAccessManager) SupportedSchemesImplementation() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager30supportedSchemesImplementationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

/*
Indicates the operation this reply is processing.



This enum was introduced or modified in  Qt 4.7.

See also QNetworkReply::operation().

*/
type QNetworkAccessManager__Operation = int

// retrieve headers operation (created with head())
const QNetworkAccessManager__HeadOperation QNetworkAccessManager__Operation = 1

// retrieve headers and download contents (created with get())
const QNetworkAccessManager__GetOperation QNetworkAccessManager__Operation = 2

// upload contents operation (created with put())
const QNetworkAccessManager__PutOperation QNetworkAccessManager__Operation = 3

// send the contents of an HTML form for processing via HTTP POST (created with post())
const QNetworkAccessManager__PostOperation QNetworkAccessManager__Operation = 4

// delete contents operation (created with deleteResource())
const QNetworkAccessManager__DeleteOperation QNetworkAccessManager__Operation = 5

// custom operation (created with sendCustomRequest())
const QNetworkAccessManager__CustomOperation QNetworkAccessManager__Operation = 6

//
const QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = 0

func (this *QNetworkAccessManager) OperationItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkAccessManager_OperationItemName(val int) string {
	var nilthis *QNetworkAccessManager
	return nilthis.OperationItemName(val)
}

/*
Indicates whether the network is accessible via this network access manager.



See also networkAccessible.

*/
type QNetworkAccessManager__NetworkAccessibility = int

//
const QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = -1

// The network is not currently accessible, either because there is currently no network coverage or network access has been explicitly disabled by a call to setNetworkAccessible().
const QNetworkAccessManager__NotAccessible QNetworkAccessManager__NetworkAccessibility = 0

// The network is accessible.
const QNetworkAccessManager__Accessible QNetworkAccessManager__NetworkAccessibility = 1

func (this *QNetworkAccessManager) NetworkAccessibilityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkAccessManager_NetworkAccessibilityItemName(val int) string {
	var nilthis *QNetworkAccessManager
	return nilthis.NetworkAccessibilityItemName(val)
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
