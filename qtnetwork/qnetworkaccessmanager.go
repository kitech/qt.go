package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h
// #include <qnetworkaccessmanager.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
type QNetworkAccessManager struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QNetworkAccessManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:103
// index:0
// Public
// void QNetworkAccessManager(class QObject *)
func NewQNetworkAccessManager(parent *qtcore.QObject /*777 QObject **/) *QNetworkAccessManager {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManagerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkAccessManagerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:104
// index:0
// Public virtual
// void ~QNetworkAccessManager()
func DeleteQNetworkAccessManager(*QNetworkAccessManager) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManagerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:109
// index:0
// Public
// void clearAccessCache()
func (this *QNetworkAccessManager) ClearAccessCache() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager16clearAccessCacheEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:111
// index:0
// Public
// void clearConnectionCache()
func (this *QNetworkAccessManager) ClearConnectionCache() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager20clearConnectionCacheEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:114
// index:0
// Public
// QNetworkProxy proxy()
func (this *QNetworkAccessManager) Proxy() *QNetworkProxy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager5proxyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:115
// index:0
// Public
// void setProxy(const class QNetworkProxy &)
func (this *QNetworkAccessManager) SetProxy(proxy *QNetworkProxy) {
	var convArg0 = proxy.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setProxyERK13QNetworkProxy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:116
// index:0
// Public
// QNetworkProxyFactory * proxyFactory()
func (this *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory /*777 QNetworkProxyFactory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager12proxyFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkProxyFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:117
// index:0
// Public
// void setProxyFactory(class QNetworkProxyFactory *)
func (this *QNetworkAccessManager) SetProxyFactory(factory *QNetworkProxyFactory /*777 QNetworkProxyFactory **/) {
	var convArg0 = factory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager15setProxyFactoryEP20QNetworkProxyFactory", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:120
// index:0
// Public
// QAbstractNetworkCache * cache()
func (this *QNetworkAccessManager) Cache() *QAbstractNetworkCache /*777 QAbstractNetworkCache **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager5cacheEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:121
// index:0
// Public
// void setCache(class QAbstractNetworkCache *)
func (this *QNetworkAccessManager) SetCache(cache *QAbstractNetworkCache /*777 QAbstractNetworkCache **/) {
	var convArg0 = cache.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setCacheEP21QAbstractNetworkCache", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:123
// index:0
// Public
// QNetworkCookieJar * cookieJar()
func (this *QNetworkAccessManager) CookieJar() *QNetworkCookieJar /*777 QNetworkCookieJar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager9cookieJarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:124
// index:0
// Public
// void setCookieJar(class QNetworkCookieJar *)
func (this *QNetworkAccessManager) SetCookieJar(cookieJar *QNetworkCookieJar /*777 QNetworkCookieJar **/) {
	var convArg0 = cookieJar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager12setCookieJarEP17QNetworkCookieJar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:126
// index:0
// Public
// void setStrictTransportSecurityEnabled(_Bool)
func (this *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager33setStrictTransportSecurityEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:127
// index:0
// Public
// bool isStrictTransportSecurityEnabled()
func (this *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager32isStrictTransportSecurityEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:128
// index:0
// Public
// void enableStrictTransportSecurityStore(_Bool, const class QString &)
func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool, storeDir *qtcore.QString) {
	var convArg1 = storeDir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager34enableStrictTransportSecurityStoreEbRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:129
// index:0
// Public
// bool isStrictTransportSecurityStoreEnabled()
func (this *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager37isStrictTransportSecurityStoreEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:133
// index:0
// Public
// QNetworkReply * head(const class QNetworkRequest &)
func (this *QNetworkAccessManager) Head(request *QNetworkRequest) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager4headERK15QNetworkRequest", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:134
// index:0
// Public
// QNetworkReply * get(const class QNetworkRequest &)
func (this *QNetworkAccessManager) Get(request *QNetworkRequest) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager3getERK15QNetworkRequest", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:135
// index:0
// Public
// QNetworkReply * post(const class QNetworkRequest &, class QIODevice *)
func (this *QNetworkAccessManager) Post(request *QNetworkRequest, data *qtcore.QIODevice /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:136
// index:1
// Public
// QNetworkReply * post(const class QNetworkRequest &, const class QByteArray &)
func (this *QNetworkAccessManager) Post_1(request *QNetworkRequest, data *qtcore.QByteArray) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:137
// index:2
// Public
// QNetworkReply * post(const class QNetworkRequest &, class QHttpMultiPart *)
func (this *QNetworkAccessManager) Post_2(request *QNetworkRequest, multiPart *QHttpMultiPart /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = multiPart.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP14QHttpMultiPart", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:138
// index:0
// Public
// QNetworkReply * put(const class QNetworkRequest &, class QIODevice *)
func (this *QNetworkAccessManager) Put(request *QNetworkRequest, data *qtcore.QIODevice /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:139
// index:1
// Public
// QNetworkReply * put(const class QNetworkRequest &, const class QByteArray &)
func (this *QNetworkAccessManager) Put_1(request *QNetworkRequest, data *qtcore.QByteArray) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:140
// index:2
// Public
// QNetworkReply * put(const class QNetworkRequest &, class QHttpMultiPart *)
func (this *QNetworkAccessManager) Put_2(request *QNetworkRequest, multiPart *QHttpMultiPart /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = multiPart.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP14QHttpMultiPart", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:141
// index:0
// Public
// QNetworkReply * deleteResource(const class QNetworkRequest &)
func (this *QNetworkAccessManager) DeleteResource(request *QNetworkRequest) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager14deleteResourceERK15QNetworkRequest", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:142
// index:0
// Public
// QNetworkReply * sendCustomRequest(const class QNetworkRequest &, const class QByteArray &, class QIODevice *)
func (this *QNetworkAccessManager) SendCustomRequest(request *QNetworkRequest, verb *qtcore.QByteArray, data *qtcore.QIODevice /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = verb.GetCthis()
	var convArg2 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:143
// index:1
// Public
// QNetworkReply * sendCustomRequest(const class QNetworkRequest &, const class QByteArray &, const class QByteArray &)
func (this *QNetworkAccessManager) SendCustomRequest_1(request *QNetworkRequest, verb *qtcore.QByteArray, data *qtcore.QByteArray) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = verb.GetCthis()
	var convArg2 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayS5_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:144
// index:2
// Public
// QNetworkReply * sendCustomRequest(const class QNetworkRequest &, const class QByteArray &, class QHttpMultiPart *)
func (this *QNetworkAccessManager) SendCustomRequest_2(request *QNetworkRequest, verb *qtcore.QByteArray, multiPart *QHttpMultiPart /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.GetCthis()
	var convArg1 = verb.GetCthis()
	var convArg2 = multiPart.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP14QHttpMultiPart", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:147
// index:0
// Public
// void setConfiguration(const class QNetworkConfiguration &)
func (this *QNetworkAccessManager) SetConfiguration(config *QNetworkConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager16setConfigurationERK21QNetworkConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:148
// index:0
// Public
// QNetworkConfiguration configuration()
func (this *QNetworkAccessManager) Configuration() *QNetworkConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager13configurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:149
// index:0
// Public
// QNetworkConfiguration activeConfiguration()
func (this *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager19activeConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:151
// index:0
// Public
// void setNetworkAccessible(enum QNetworkAccessManager::NetworkAccessibility)
func (this *QNetworkAccessManager) SetNetworkAccessible(accessible int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager20setNetworkAccessibleENS_20NetworkAccessibilityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:152
// index:0
// Public
// QNetworkAccessManager::NetworkAccessibility networkAccessible()
func (this *QNetworkAccessManager) NetworkAccessible() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager17networkAccessibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:156
// index:0
// Public
// void connectToHostEncrypted(const class QString &, quint16, const class QSslConfiguration &)
func (this *QNetworkAccessManager) ConnectToHostEncrypted(hostName *qtcore.QString, port uint16, sslConfiguration *QSslConfiguration) {
	var convArg0 = hostName.GetCthis()
	var convArg2 = sslConfiguration.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager22connectToHostEncryptedERK7QStringtRK17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:159
// index:0
// Public
// void connectToHost(const class QString &, quint16)
func (this *QNetworkAccessManager) ConnectToHost(hostName *qtcore.QString, port uint16) {
	var convArg0 = hostName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager13connectToHostERK7QStringt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:161
// index:0
// Public
// void setRedirectPolicy(class QNetworkRequest::RedirectPolicy)
func (this *QNetworkAccessManager) SetRedirectPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager17setRedirectPolicyEN15QNetworkRequest14RedirectPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:162
// index:0
// Public
// QNetworkRequest::RedirectPolicy redirectPolicy()
func (this *QNetworkAccessManager) RedirectPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkAccessManager14redirectPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:166
// index:0
// Public
// void proxyAuthenticationRequired(const class QNetworkProxy &, class QAuthenticator *)
func (this *QNetworkAccessManager) ProxyAuthenticationRequired(proxy *QNetworkProxy, authenticator *QAuthenticator /*777 QAuthenticator **/) {
	var convArg0 = proxy.GetCthis()
	var convArg1 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager27proxyAuthenticationRequiredERK13QNetworkProxyP14QAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:168
// index:0
// Public
// void authenticationRequired(class QNetworkReply *, class QAuthenticator *)
func (this *QNetworkAccessManager) AuthenticationRequired(reply *QNetworkReply /*777 QNetworkReply **/, authenticator *QAuthenticator /*777 QAuthenticator **/) {
	var convArg0 = reply.GetCthis()
	var convArg1 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager22authenticationRequiredEP13QNetworkReplyP14QAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:169
// index:0
// Public
// void finished(class QNetworkReply *)
func (this *QNetworkAccessManager) Finished(reply *QNetworkReply /*777 QNetworkReply **/) {
	var convArg0 = reply.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager8finishedEP13QNetworkReply", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:171
// index:0
// Public
// void encrypted(class QNetworkReply *)
func (this *QNetworkAccessManager) Encrypted(reply *QNetworkReply /*777 QNetworkReply **/) {
	var convArg0 = reply.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager9encryptedEP13QNetworkReply", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:173
// index:0
// Public
// void preSharedKeyAuthenticationRequired(class QNetworkReply *, class QSslPreSharedKeyAuthenticator *)
func (this *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply *QNetworkReply /*777 QNetworkReply **/, authenticator *QSslPreSharedKeyAuthenticator /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = reply.GetCthis()
	var convArg1 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager34preSharedKeyAuthenticationRequiredEP13QNetworkReplyP29QSslPreSharedKeyAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:177
// index:0
// Public
// void networkSessionConnected()
func (this *QNetworkAccessManager) NetworkSessionConnected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager23networkSessionConnectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:179
// index:0
// Public
// void networkAccessibleChanged(class QNetworkAccessManager::NetworkAccessibility)
func (this *QNetworkAccessManager) NetworkAccessibleChanged(accessible int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager24networkAccessibleChangedENS_20NetworkAccessibilityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:183
// index:0
// Protected virtual
// QNetworkReply * createRequest(enum QNetworkAccessManager::Operation, const class QNetworkRequest &, class QIODevice *)
func (this *QNetworkAccessManager) CreateRequest(op int, request *QNetworkRequest, outgoingData *qtcore.QIODevice /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg1 = request.GetCthis()
	var convArg2 = outgoingData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkAccessManager13createRequestENS_9OperationERK15QNetworkRequestP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), op, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QNetworkAccessManager__Operation = int

const QNetworkAccessManager__HeadOperation QNetworkAccessManager__Operation = 1
const QNetworkAccessManager__GetOperation QNetworkAccessManager__Operation = 2
const QNetworkAccessManager__PutOperation QNetworkAccessManager__Operation = 3
const QNetworkAccessManager__PostOperation QNetworkAccessManager__Operation = 4
const QNetworkAccessManager__DeleteOperation QNetworkAccessManager__Operation = 5
const QNetworkAccessManager__CustomOperation QNetworkAccessManager__Operation = 6
const QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = 0

type QNetworkAccessManager__NetworkAccessibility = int

const QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = 4294967295
const QNetworkAccessManager__NotAccessible QNetworkAccessManager__NetworkAccessibility = 0
const QNetworkAccessManager__Accessible QNetworkAccessManager__NetworkAccessibility = 1

//  body block end
