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
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// QNetworkReply * createRequest(enum QNetworkAccessManager::Operation, const class QNetworkRequest &, class QIODevice *)
func (this *QNetworkAccessManager) InheritCreateRequest(f func(op int, request *QNetworkRequest, outgoingData *qtcore.QIODevice /*777 QIODevice **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createRequest", f)
}

// QStringList supportedSchemesImplementation()
func (this *QNetworkAccessManager) InheritSupportedSchemesImplementation(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "supportedSchemesImplementation", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QNetworkAccessManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkAccessManager(QObject *)
func NewQNetworkAccessManager(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkAccessManager {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManagerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkAccessManager()
func DeleteQNetworkAccessManager(this *QNetworkAccessManager) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManagerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedSchemes()
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
func (this *QNetworkAccessManager) ClearAccessCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager16clearAccessCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearConnectionCache()
func (this *QNetworkAccessManager) ClearConnectionCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager20clearConnectionCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy proxy()
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
func (this *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	var convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setProxyERK13QNetworkProxy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxyFactory * proxyFactory()
func (this *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory /*777 QNetworkProxyFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager12proxyFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkProxyFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxyFactory(QNetworkProxyFactory *)
func (this *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF /*777 QNetworkProxyFactory **/) {
	var convArg0 = factory.QNetworkProxyFactory_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager15setProxyFactoryEP20QNetworkProxyFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractNetworkCache * cache()
func (this *QNetworkAccessManager) Cache() *QAbstractNetworkCache /*777 QAbstractNetworkCache **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager5cacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCache(QAbstractNetworkCache *)
func (this *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF /*777 QAbstractNetworkCache **/) {
	var convArg0 = cache.QAbstractNetworkCache_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8setCacheEP21QAbstractNetworkCache", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkCookieJar * cookieJar()
func (this *QNetworkAccessManager) CookieJar() *QNetworkCookieJar /*777 QNetworkCookieJar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager9cookieJarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkCookieJarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCookieJar(QNetworkCookieJar *)
func (this *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF /*777 QNetworkCookieJar **/) {
	var convArg0 = cookieJar.QNetworkCookieJar_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager12setCookieJarEP17QNetworkCookieJar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrictTransportSecurityEnabled(_Bool)
func (this *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager33setStrictTransportSecurityEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStrictTransportSecurityEnabled()
func (this *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager32isStrictTransportSecurityEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enableStrictTransportSecurityStore(_Bool, const QString &)
func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool, storeDir string) {
	var tmpArg1 = qtcore.NewQString_5(storeDir)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager34enableStrictTransportSecurityStoreEbRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStrictTransportSecurityStoreEnabled()
func (this *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager37isStrictTransportSecurityStoreEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * head(const QNetworkRequest &)
func (this *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4headERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * get(const QNetworkRequest &)
func (this *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3getERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, QIODevice *)
func (this *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = data.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:136
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, const QByteArray &)
func (this *QNetworkAccessManager) Post_1(request QNetworkRequest_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = data.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:137
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * post(const QNetworkRequest &, QHttpMultiPart *)
func (this *QNetworkAccessManager) Post_2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager4postERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, QIODevice *)
func (this *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = data.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:139
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, const QByteArray &)
func (this *QNetworkAccessManager) Put_1(request QNetworkRequest_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = data.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:140
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * put(const QNetworkRequest &, QHttpMultiPart *)
func (this *QNetworkAccessManager) Put_2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = multiPart.QHttpMultiPart_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager3putERK15QNetworkRequestP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * deleteResource(const QNetworkRequest &)
func (this *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager14deleteResourceERK15QNetworkRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QIODevice *)
func (this *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, data qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = verb.QByteArray_PTR().GetCthis()
	var convArg2 = data.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:143
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, const QByteArray &)
func (this *QNetworkAccessManager) SendCustomRequest_1(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, data qtcore.QByteArray_ITF) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = verb.QByteArray_PTR().GetCthis()
	var convArg2 = data.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:144
// index:2
// Public Visibility=Default Availability=Available
// [8] QNetworkReply * sendCustomRequest(const QNetworkRequest &, const QByteArray &, QHttpMultiPart *)
func (this *QNetworkAccessManager) SendCustomRequest_2(request QNetworkRequest_ITF, verb qtcore.QByteArray_ITF, multiPart QHttpMultiPart_ITF /*777 QHttpMultiPart **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg0 = request.QNetworkRequest_PTR().GetCthis()
	var convArg1 = verb.QByteArray_PTR().GetCthis()
	var convArg2 = multiPart.QHttpMultiPart_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17sendCustomRequestERK15QNetworkRequestRK10QByteArrayP14QHttpMultiPart", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setConfiguration(const QNetworkConfiguration &)
func (this *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	var convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager16setConfigurationERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration configuration()
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
// [8] QNetworkConfiguration activeConfiguration()
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
// [-2] void setNetworkAccessible(enum QNetworkAccessManager::NetworkAccessibility)
func (this *QNetworkAccessManager) SetNetworkAccessible(accessible int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager20setNetworkAccessibleENS_20NetworkAccessibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkAccessManager::NetworkAccessibility networkAccessible()
func (this *QNetworkAccessManager) NetworkAccessible() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager17networkAccessibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QSslConfiguration &)
func (this *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 = sslConfiguration.QSslConfiguration_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22connectToHostEncryptedERK7QStringtRK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16)
func (this *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13connectToHostERK7QStringt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRedirectPolicy(QNetworkRequest::RedirectPolicy)
func (this *QNetworkAccessManager) SetRedirectPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager17setRedirectPolicyEN15QNetworkRequest14RedirectPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkRequest::RedirectPolicy redirectPolicy()
func (this *QNetworkAccessManager) RedirectPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager14redirectPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void proxyAuthenticationRequired(const QNetworkProxy &, QAuthenticator *)
func (this *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	var convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager27proxyAuthenticationRequiredERK13QNetworkProxyP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void authenticationRequired(QNetworkReply *, QAuthenticator *)
func (this *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF /*777 QNetworkReply **/, authenticator QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 = reply.QNetworkReply_PTR().GetCthis()
	var convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager22authenticationRequiredEP13QNetworkReplyP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(QNetworkReply *)
func (this *QNetworkAccessManager) Finished(reply QNetworkReply_ITF /*777 QNetworkReply **/) {
	var convArg0 = reply.QNetworkReply_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager8finishedEP13QNetworkReply", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted(QNetworkReply *)
func (this *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF /*777 QNetworkReply **/) {
	var convArg0 = reply.QNetworkReply_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager9encryptedEP13QNetworkReply", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QNetworkReply *, QSslPreSharedKeyAuthenticator *)
func (this *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF /*777 QNetworkReply **/, authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = reply.QNetworkReply_PTR().GetCthis()
	var convArg1 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager34preSharedKeyAuthenticationRequiredEP13QNetworkReplyP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void networkSessionConnected()
func (this *QNetworkAccessManager) NetworkSessionConnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager23networkSessionConnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void networkAccessibleChanged(QNetworkAccessManager::NetworkAccessibility)
func (this *QNetworkAccessManager) NetworkAccessibleChanged(accessible int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager24networkAccessibleChangedENS_20NetworkAccessibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accessible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QNetworkReply * createRequest(enum QNetworkAccessManager::Operation, const QNetworkRequest &, QIODevice *)
func (this *QNetworkAccessManager) CreateRequest(op int, request QNetworkRequest_ITF, outgoingData qtcore.QIODevice_ITF /*777 QIODevice **/) *QNetworkReply /*777 QNetworkReply **/ {
	var convArg1 = request.QNetworkRequest_PTR().GetCthis()
	var convArg2 = outgoingData.QIODevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkAccessManager13createRequestENS_9OperationERK15QNetworkRequestP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQNetworkReplyFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkaccessmanager.h:187
// index:0
// Protected Visibility=Default Availability=Available
// [8] QStringList supportedSchemesImplementation()
func (this *QNetworkAccessManager) SupportedSchemesImplementation() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkAccessManager30supportedSchemesImplementationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
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

const QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = -1
const QNetworkAccessManager__NotAccessible QNetworkAccessManager__NetworkAccessibility = 0
const QNetworkAccessManager__Accessible QNetworkAccessManager__NetworkAccessibility = 1

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
