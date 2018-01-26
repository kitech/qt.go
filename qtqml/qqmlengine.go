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
import "mkuse/cffiqt"
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
		ffiqt.KeepMe()
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:100
// index:0
// Public
// void QQmlEngine(class QObject *)
func NewQQmlEngine(p *qtcore.QObject /*777 QObject **/) *QQmlEngine {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngineC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlengine.h:101
// index:0
// Public virtual
// void ~QQmlEngine()
func DeleteQQmlEngine(*QQmlEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:103
// index:0
// Public
// QQmlContext * rootContext()
func (this *QQmlEngine) RootContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine11rootContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:105
// index:0
// Public
// void clearComponentCache()
func (this *QQmlEngine) ClearComponentCache() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine19clearComponentCacheEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:106
// index:0
// Public
// void trimComponentCache()
func (this *QQmlEngine) TrimComponentCache() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine18trimComponentCacheEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:109
// index:0
// Public
// void setImportPathList(const class QStringList &)
func (this *QQmlEngine) SetImportPathList(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine17setImportPathListERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:110
// index:0
// Public
// void addImportPath(const class QString &)
func (this *QQmlEngine) AddImportPath(dir *qtcore.QString) {
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine13addImportPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:113
// index:0
// Public
// void setPluginPathList(const class QStringList &)
func (this *QQmlEngine) SetPluginPathList(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine17setPluginPathListERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:114
// index:0
// Public
// void addPluginPath(const class QString &)
func (this *QQmlEngine) AddPluginPath(dir *qtcore.QString) {
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine13addPluginPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:116
// index:0
// Public
// bool addNamedBundle(const class QString &, const class QString &)
func (this *QQmlEngine) AddNamedBundle(name *qtcore.QString, fileName *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	var convArg1 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine14addNamedBundleERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:121
// index:0
// Public
// void setNetworkAccessManagerFactory(class QQmlNetworkAccessManagerFactory *)
func (this *QQmlEngine) SetNetworkAccessManagerFactory(arg0 *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine30setNetworkAccessManagerFactoryEP31QQmlNetworkAccessManagerFactory", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:122
// index:0
// Public
// QQmlNetworkAccessManagerFactory * networkAccessManagerFactory()
func (this *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine27networkAccessManagerFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlNetworkAccessManagerFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:124
// index:0
// Public
// QNetworkAccessManager * networkAccessManager()
func (this *QQmlEngine) NetworkAccessManager() *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine20networkAccessManagerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:127
// index:0
// Public
// void setUrlInterceptor(class QQmlAbstractUrlInterceptor *)
func (this *QQmlEngine) SetUrlInterceptor(urlInterceptor *QQmlAbstractUrlInterceptor /*777 QQmlAbstractUrlInterceptor **/) {
	var convArg0 = urlInterceptor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine17setUrlInterceptorEP26QQmlAbstractUrlInterceptor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:128
// index:0
// Public
// QQmlAbstractUrlInterceptor * urlInterceptor()
func (this *QQmlEngine) UrlInterceptor() *QQmlAbstractUrlInterceptor /*777 QQmlAbstractUrlInterceptor **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine14urlInterceptorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlAbstractUrlInterceptorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:130
// index:0
// Public
// void addImageProvider(const class QString &, class QQmlImageProviderBase *)
func (this *QQmlEngine) AddImageProvider(id *qtcore.QString, arg1 unsafe.Pointer /*666*/) {
	var convArg0 = id.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine16addImageProviderERK7QStringP21QQmlImageProviderBase", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:131
// index:0
// Public
// QQmlImageProviderBase * imageProvider(const class QString &)
func (this *QQmlEngine) ImageProvider(id *qtcore.QString) unsafe.Pointer /*666*/ {
	var convArg0 = id.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine13imageProviderERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQml/qqmlengine.h:132
// index:0
// Public
// void removeImageProvider(const class QString &)
func (this *QQmlEngine) RemoveImageProvider(id *qtcore.QString) {
	var convArg0 = id.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine19removeImageProviderERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:134
// index:0
// Public
// void setIncubationController(class QQmlIncubationController *)
func (this *QQmlEngine) SetIncubationController(arg0 *QQmlIncubationController /*777 QQmlIncubationController **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine23setIncubationControllerEP24QQmlIncubationController", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:135
// index:0
// Public
// QQmlIncubationController * incubationController()
func (this *QQmlEngine) IncubationController() *QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine20incubationControllerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:137
// index:0
// Public
// void setOfflineStoragePath(const class QString &)
func (this *QQmlEngine) SetOfflineStoragePath(dir *qtcore.QString) {
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine21setOfflineStoragePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:138
// index:0
// Public
// QString offlineStoragePath()
func (this *QQmlEngine) OfflineStoragePath() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine18offlineStoragePathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:139
// index:0
// Public
// QString offlineStorageDatabaseFilePath(const class QString &)
func (this *QQmlEngine) OfflineStorageDatabaseFilePath(databaseName *qtcore.QString) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = databaseName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine30offlineStorageDatabaseFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:141
// index:0
// Public
// QUrl baseUrl()
func (this *QQmlEngine) BaseUrl() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine7baseUrlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:142
// index:0
// Public
// void setBaseUrl(const class QUrl &)
func (this *QQmlEngine) SetBaseUrl(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine10setBaseUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:144
// index:0
// Public
// bool outputWarningsToStandardError()
func (this *QQmlEngine) OutputWarningsToStandardError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQmlEngine29outputWarningsToStandardErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:145
// index:0
// Public
// void setOutputWarningsToStandardError(_Bool)
func (this *QQmlEngine) SetOutputWarningsToStandardError(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine32setOutputWarningsToStandardErrorEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:148
// index:0
// Public
// void retranslate()
func (this *QQmlEngine) Retranslate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine11retranslateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:151
// index:0
// Public static
// QQmlContext * contextForObject(const class QObject *)
func (this *QQmlEngine) ContextForObject(arg0 *qtcore.QObject /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine16contextForObjectEPK7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QQmlEngine_ContextForObject(arg0 *qtcore.QObject /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var nilthis *QQmlEngine
	rv := nilthis.ContextForObject(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:152
// index:0
// Public static
// void setContextForObject(class QObject *, class QQmlContext *)
func (this *QQmlEngine) SetContextForObject(arg0 *qtcore.QObject /*777 QObject **/, arg1 *QQmlContext /*777 QQmlContext **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine19setContextForObjectEP7QObjectP11QQmlContext", ffiqt.FFI_TYPE_POINTER, arg0, arg1)
	gopp.ErrPrint(err, rv)
}
func QQmlEngine_SetContextForObject(arg0 *qtcore.QObject /*777 QObject **/, arg1 *QQmlContext /*777 QQmlContext **/) {
	var nilthis *QQmlEngine
	nilthis.SetContextForObject(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:155
// index:0
// Public static
// void setObjectOwnership(class QObject *, enum QQmlEngine::ObjectOwnership)
func (this *QQmlEngine) SetObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine18setObjectOwnershipEP7QObjectNS_15ObjectOwnershipE", ffiqt.FFI_TYPE_POINTER, arg0, arg1)
	gopp.ErrPrint(err, rv)
}
func QQmlEngine_SetObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/, arg1 int) {
	var nilthis *QQmlEngine
	nilthis.SetObjectOwnership(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:156
// index:0
// Public static
// QQmlEngine::ObjectOwnership objectOwnership(class QObject *)
func (this *QQmlEngine) ObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine15objectOwnershipEP7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QQmlEngine_ObjectOwnership(arg0 *qtcore.QObject /*777 QObject **/) int {
	var nilthis *QQmlEngine
	rv := nilthis.ObjectOwnership(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:159
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QQmlEngine) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:162
// index:0
// Public
// void quit()
func (this *QQmlEngine) Quit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine4quitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:163
// index:0
// Public
// void exit(int)
func (this *QQmlEngine) Exit(retCode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQmlEngine4exitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), retCode)
	gopp.ErrPrint(err, rv)
}

type QQmlEngine__ObjectOwnership = int

const QQmlEngine__CppOwnership QQmlEngine__ObjectOwnership = 0
const QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = 1

//  body block end
