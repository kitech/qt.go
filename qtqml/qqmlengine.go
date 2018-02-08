package qtqml

// /usr/include/qt/QtQml/qqmlengine.h
// #include <qqmlengine.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QQmlEngine) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QQmlEngine struct {
	*QJSEngine
}

func (this *QQmlEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QJSEngine.GetCthis()
	}
}
func (this *QQmlEngine) SetCthis(cthis unsafe.Pointer) {
	this.QJSEngine = NewQJSEngineFromPointer(cthis)
}
func NewQQmlEngineFromPointer(cthis unsafe.Pointer) *QQmlEngine {
	bcthis0 := NewQJSEngineFromPointer(cthis)
	return &QQmlEngine{bcthis0}
}
func (*QQmlEngine) NewFromPointer(cthis unsafe.Pointer) *QQmlEngine {
	return NewQQmlEngineFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlengine.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQmlEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlEngine(QObject *)
func NewQQmlEngine(p *qtcore.QObject /*777 QObject **/) *QQmlEngine {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlengine.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlEngine()
func DeleteQQmlEngine(this *QQmlEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlengine.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * rootContext()
func (this *QQmlEngine) RootContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine11rootContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearComponentCache()
func (this *QQmlEngine) ClearComponentCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19clearComponentCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void trimComponentCache()
func (this *QQmlEngine) TrimComponentCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine18trimComponentCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImportPathList(const QStringList &)
func (this *QQmlEngine) SetImportPathList(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setImportPathListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addImportPath(const QString &)
func (this *QQmlEngine) AddImportPath(dir string) {
	var tmpArg0 = qtcore.NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine13addImportPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPluginPathList(const QStringList &)
func (this *QQmlEngine) SetPluginPathList(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setPluginPathListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPluginPath(const QString &)
func (this *QQmlEngine) AddPluginPath(dir string) {
	var tmpArg0 = qtcore.NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine13addPluginPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addNamedBundle(const QString &, const QString &)
func (this *QQmlEngine) AddNamedBundle(name string, fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine14addNamedBundleERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNetworkAccessManagerFactory(QQmlNetworkAccessManagerFactory *)
func (this *QQmlEngine) SetNetworkAccessManagerFactory(arg0 *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine30setNetworkAccessManagerFactoryEP31QQmlNetworkAccessManagerFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlNetworkAccessManagerFactory * networkAccessManagerFactory()
func (this *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine27networkAccessManagerFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQQmlNetworkAccessManagerFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkAccessManager * networkAccessManager()
func (this *QQmlEngine) NetworkAccessManager() *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine20networkAccessManagerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrlInterceptor(QQmlAbstractUrlInterceptor *)
func (this *QQmlEngine) SetUrlInterceptor(urlInterceptor *QQmlAbstractUrlInterceptor /*777 QQmlAbstractUrlInterceptor **/) {
	var convArg0 = urlInterceptor.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setUrlInterceptorEP26QQmlAbstractUrlInterceptor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlAbstractUrlInterceptor * urlInterceptor()
func (this *QQmlEngine) UrlInterceptor() *QQmlAbstractUrlInterceptor /*777 QQmlAbstractUrlInterceptor **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine14urlInterceptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQQmlAbstractUrlInterceptorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addImageProvider(const QString &, QQmlImageProviderBase *)
func (this *QQmlEngine) AddImageProvider(id string, arg1 unsafe.Pointer /*666*/) {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine16addImageProviderERK7QStringP21QQmlImageProviderBase", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlImageProviderBase * imageProvider(const QString &)
func (this *QQmlEngine) ImageProvider(id string) unsafe.Pointer /*666*/ {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine13imageProviderERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQml/qqmlengine.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeImageProvider(const QString &)
func (this *QQmlEngine) RemoveImageProvider(id string) {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19removeImageProviderERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIncubationController(QQmlIncubationController *)
func (this *QQmlEngine) SetIncubationController(arg0 *QQmlIncubationController /*777 QQmlIncubationController **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine23setIncubationControllerEP24QQmlIncubationController", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlIncubationController * incubationController()
func (this *QQmlEngine) IncubationController() *QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine20incubationControllerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOfflineStoragePath(const QString &)
func (this *QQmlEngine) SetOfflineStoragePath(dir string) {
	var tmpArg0 = qtcore.NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine21setOfflineStoragePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QString offlineStoragePath()
func (this *QQmlEngine) OfflineStoragePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine18offlineStoragePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlengine.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString offlineStorageDatabaseFilePath(const QString &)
func (this *QQmlEngine) OfflineStorageDatabaseFilePath(databaseName string) string {
	var tmpArg0 = qtcore.NewQString_5(databaseName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine30offlineStorageDatabaseFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlengine.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl()
func (this *QQmlEngine) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)
func (this *QQmlEngine) SetBaseUrl(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool outputWarningsToStandardError()
func (this *QQmlEngine) OutputWarningsToStandardError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine29outputWarningsToStandardErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOutputWarningsToStandardError(_Bool)
func (this *QQmlEngine) SetOutputWarningsToStandardError(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine32setOutputWarningsToStandardErrorEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void retranslate()
func (this *QQmlEngine) Retranslate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine11retranslateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQmlContext * contextForObject(const QObject *)
func (this *QQmlEngine) ContextForObject(arg0 *qtcore.QObject /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine16contextForObjectEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQmlEngine_ContextForObject(arg0 *qtcore.QObject /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var nilthis *QQmlEngine
	rv := nilthis.ContextForObject(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setContextForObject(QObject *, QQmlContext *)
func (this *QQmlEngine) SetContextForObject(arg0 *qtcore.QObject /*777 QObject **/, arg1 *QQmlContext /*777 QQmlContext **/) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19setContextForObjectEP7QObjectP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QQmlEngine_SetContextForObject(arg0 *qtcore.QObject /*777 QObject **/, arg1 *QQmlContext /*777 QQmlContext **/) {
	var nilthis *QQmlEngine
	nilthis.SetContextForObject(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:155
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setObjectOwnership(QObject *, enum QQmlEngine::ObjectOwnership)
func (this *QQmlEngine) SetObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine18setObjectOwnershipEP7QObjectNS_15ObjectOwnershipE", qtrt.FFI_TYPE_POINTER, convArg0, arg1)
	gopp.ErrPrint(err, rv)
}
func QQmlEngine_SetObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) {
	var nilthis *QQmlEngine
	nilthis.SetObjectOwnership(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [4] QQmlEngine::ObjectOwnership objectOwnership(QObject *)
func (this *QQmlEngine) ObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine15objectOwnershipEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QQmlEngine_ObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/) int {
	var nilthis *QQmlEngine
	rv := nilthis.ObjectOwnership(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQmlEngine) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()
func (this *QQmlEngine) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QQmlEngine) Exit(retCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retCode)
	gopp.ErrPrint(err, rv)
}

type QQmlEngine__ObjectOwnership = int

const QQmlEngine__CppOwnership QQmlEngine__ObjectOwnership = 0
const QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = 1

//  body block end
