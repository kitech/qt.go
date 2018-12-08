package qtqml

// /usr/include/qt/QtQml/qqmlengine.h
// #include <qqmlengine.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QQmlEngine) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QQmlEngine struct {
	*QJSEngine
}
type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlEngine(QObject *)

/*
Create a new QQmlEngine with the given parent.
*/
func (*QQmlEngine) NewForInherit(p qtcore.QObject_ITF /*777 QObject **/) *QQmlEngine {
	return NewQQmlEngine(p)
}
func NewQQmlEngine(p qtcore.QObject_ITF /*777 QObject **/) *QQmlEngine {
	var convArg0 unsafe.Pointer
	if p != nil && p.QObject_PTR() != nil {
		convArg0 = p.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlengine.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlEngine(QObject *)

/*
Create a new QQmlEngine with the given parent.
*/
func (*QQmlEngine) NewForInheritp() *QQmlEngine {
	return NewQQmlEnginep()
}
func NewQQmlEnginep() *QQmlEngine {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlEngine")
	return gothis
}

// /usr/include/qt/QtQml/qqmlengine.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlEngine()

/*

 */
func DeleteQQmlEngine(this *QQmlEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlengine.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * rootContext() const

/*
Returns the engine's root context.

The root context is automatically created by the QQmlEngine. Data that should be available to all QML component instances instantiated by the engine should be put in the root context.

Additional data that should only be available to a subset of component instances should be added to sub-contexts parented to the root context.
*/
func (this *QQmlEngine) RootContext() *QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine11rootContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearComponentCache()

/*
Clears the engine's internal component cache.

This function causes the property metadata of all components previously loaded by the engine to be destroyed. All previously loaded components and the property bindings for all extant objects created from those components will cease to function.

This function returns the engine to a state where it does not contain any loaded component data. This may be useful in order to reload a smaller subset of the previous component set, or to load a new version of a previously loaded component.

Once the component cache has been cleared, components must be loaded before any new objects can be created.

See also trimComponentCache().
*/
func (this *QQmlEngine) ClearComponentCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19clearComponentCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void trimComponentCache()

/*
Trims the engine's internal component cache.

This function causes the property metadata of any loaded components which are not currently in use to be destroyed.

A component is considered to be in use if there are any extant instances of the component itself, any instances of other components that use the component, or any objects instantiated by any of those components.

See also clearComponentCache().
*/
func (this *QQmlEngine) TrimComponentCache() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine18trimComponentCacheEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList importPathList() const

/*
Returns the list of directories where the engine searches for installed modules in a URL-based directory structure.

For example, if /opt/MyApp/lib/imports is in the path, then QML that imports com.mycompany.Feature will cause the QQmlEngine to look in /opt/MyApp/lib/imports/com/mycompany/Feature/ for the components provided by that module. A qmldir file is required for defining the type version mapping and possibly QML extensions plugins.

By default, the list contains the directory of the application executable, paths specified in the QML2_IMPORT_PATH environment variable, and the builtin Qml2ImportsPath from QLibraryInfo.

See also addImportPath() and setImportPathList().
*/
func (this *QQmlEngine) ImportPathList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine14importPathListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImportPathList(const QStringList &)

/*
Sets paths as the list of directories where the engine searches for installed modules in a URL-based directory structure.

By default, the list contains the directory of the application executable, paths specified in the QML2_IMPORT_PATH environment variable, and the builtin Qml2ImportsPath from QLibraryInfo.

See also importPathList() and addImportPath().
*/
func (this *QQmlEngine) SetImportPathList(paths qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setImportPathListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addImportPath(const QString &)

/*
Adds path as a directory where the engine searches for installed modules in a URL-based directory structure.

The path may be a local filesystem directory, a Qt Resource path (:/imports), a Qt Resource url (qrc:/imports) or a URL.

The path will be converted into canonical form before it is added to the import path list.

The newly added path will be first in the importPathList().

See also setImportPathList() and QML Modules.
*/
func (this *QQmlEngine) AddImportPath(dir string) {
	var tmpArg0 = qtcore.NewQString5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine13addImportPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList pluginPathList() const

/*
Returns the list of directories where the engine searches for native plugins for imported modules (referenced in the qmldir file).

By default, the list contains only ., i.e. the engine searches in the directory of the qmldir file itself.

See also addPluginPath() and setPluginPathList().
*/
func (this *QQmlEngine) PluginPathList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine14pluginPathListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPluginPathList(const QStringList &)

/*
Sets the list of directories where the engine searches for native plugins for imported modules (referenced in the qmldir file) to paths.

By default, the list contains only ., i.e. the engine searches in the directory of the qmldir file itself.

See also pluginPathList() and addPluginPath().
*/
func (this *QQmlEngine) SetPluginPathList(paths qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setPluginPathListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPluginPath(const QString &)

/*
Adds path as a directory where the engine searches for native plugins for imported modules (referenced in the qmldir file).

By default, the list contains only ., i.e. the engine searches in the directory of the qmldir file itself.

The newly added path will be first in the pluginPathList().

See also setPluginPathList().
*/
func (this *QQmlEngine) AddPluginPath(dir string) {
	var tmpArg0 = qtcore.NewQString5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine13addPluginPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addNamedBundle(const QString &, const QString &)

/*

 */
func (this *QQmlEngine) AddNamedBundle(name string, fileName string) bool {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(fileName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine14addNamedBundleERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNetworkAccessManagerFactory(QQmlNetworkAccessManagerFactory *)

/*
Sets the factory to use for creating QNetworkAccessManager(s).

QNetworkAccessManager is used for all network access by QML. By implementing a factory it is possible to create custom QNetworkAccessManager with specialized caching, proxy and cookie support.

The factory must be set before executing the engine.

See also networkAccessManagerFactory().
*/
func (this *QQmlEngine) SetNetworkAccessManagerFactory(arg0 QQmlNetworkAccessManagerFactory_ITF /*777 QQmlNetworkAccessManagerFactory **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlNetworkAccessManagerFactory_PTR() != nil {
		convArg0 = arg0.QQmlNetworkAccessManagerFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine30setNetworkAccessManagerFactoryEP31QQmlNetworkAccessManagerFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlNetworkAccessManagerFactory * networkAccessManagerFactory() const

/*
Returns the current QQmlNetworkAccessManagerFactory.

See also setNetworkAccessManagerFactory().
*/
func (this *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory /*777 QQmlNetworkAccessManagerFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine27networkAccessManagerFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlNetworkAccessManagerFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkAccessManager * networkAccessManager() const

/*
Returns a common QNetworkAccessManager which can be used by any QML type instantiated by this engine.

If a QQmlNetworkAccessManagerFactory has been set and a QNetworkAccessManager has not yet been created, the QQmlNetworkAccessManagerFactory will be used to create the QNetworkAccessManager; otherwise the returned QNetworkAccessManager will have no proxy or cache set.

See also setNetworkAccessManagerFactory().
*/
func (this *QQmlEngine) NetworkAccessManager() *qtnetwork.QNetworkAccessManager /*777 QNetworkAccessManager **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine20networkAccessManagerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtnetwork.NewQNetworkAccessManagerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrlInterceptor(QQmlAbstractUrlInterceptor *)

/*

 */
func (this *QQmlEngine) SetUrlInterceptor(urlInterceptor QQmlAbstractUrlInterceptor_ITF /*777 QQmlAbstractUrlInterceptor **/) {
	var convArg0 unsafe.Pointer
	if urlInterceptor != nil && urlInterceptor.QQmlAbstractUrlInterceptor_PTR() != nil {
		convArg0 = urlInterceptor.QQmlAbstractUrlInterceptor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine17setUrlInterceptorEP26QQmlAbstractUrlInterceptor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlAbstractUrlInterceptor * urlInterceptor() const

/*

 */
func (this *QQmlEngine) UrlInterceptor() *QQmlAbstractUrlInterceptor /*777 QQmlAbstractUrlInterceptor **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine14urlInterceptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlAbstractUrlInterceptorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addImageProvider(const QString &, QQmlImageProviderBase *)

/*
Sets the provider to use for images requested via the image: url scheme, with host providerId. The QQmlEngine takes ownership of provider.

Image providers enable support for pixmap and threaded image requests. See the QQuickImageProvider documentation for details on implementing and using image providers.

All required image providers should be added to the engine before any QML sources files are loaded.

See also removeImageProvider(), QQuickImageProvider, and QQmlImageProviderBase.
*/
func (this *QQmlEngine) AddImageProvider(id string, arg1 unsafe.Pointer /*666*/) {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine16addImageProviderERK7QStringP21QQmlImageProviderBase", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlImageProviderBase * imageProvider(const QString &) const

/*
Returns the image provider set for providerId.

Returns the provider if it was found; otherwise returns 0.

See also QQuickImageProvider.
*/
func (this *QQmlEngine) ImageProvider(id string) unsafe.Pointer /*666*/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine13imageProviderERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQml/qqmlengine.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeImageProvider(const QString &)

/*
Removes the image provider for providerId.

See also addImageProvider() and QQuickImageProvider.
*/
func (this *QQmlEngine) RemoveImageProvider(id string) {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19removeImageProviderERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIncubationController(QQmlIncubationController *)

/*
Sets the engine's incubation controller. The engine can only have one active controller and it does not take ownership of it.

See also incubationController().
*/
func (this *QQmlEngine) SetIncubationController(arg0 QQmlIncubationController_ITF /*777 QQmlIncubationController **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlIncubationController_PTR() != nil {
		convArg0 = arg0.QQmlIncubationController_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine23setIncubationControllerEP24QQmlIncubationController", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlIncubationController * incubationController() const

/*
Returns the currently set incubation controller, or 0 if no controller has been set.

See also setIncubationController().
*/
func (this *QQmlEngine) IncubationController() *QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine20incubationControllerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlengine.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOfflineStoragePath(const QString &)

/*

 */
func (this *QQmlEngine) SetOfflineStoragePath(dir string) {
	var tmpArg0 = qtcore.NewQString5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine21setOfflineStoragePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QString offlineStoragePath() const

/*

 */
func (this *QQmlEngine) OfflineStoragePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine18offlineStoragePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlengine.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString offlineStorageDatabaseFilePath(const QString &) const

/*
Returns the file path where a Local Storage database with the identifier databaseName is (or would be) located.

This function was introduced in  Qt 5.9.

See also LocalStorage.openDatabaseSync().
*/
func (this *QQmlEngine) OfflineStorageDatabaseFilePath(databaseName string) string {
	var tmpArg0 = qtcore.NewQString5(databaseName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine30offlineStorageDatabaseFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlengine.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl() const

/*
Return the base URL for this engine. The base URL is only used to resolve components when a relative URL is passed to the QQmlComponent constructor.

If a base URL has not been explicitly set, this method returns the application's current working directory.

See also setBaseUrl().
*/
func (this *QQmlEngine) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlengine.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)

/*
Set the base URL for this engine to url.

See also baseUrl().
*/
func (this *QQmlEngine) SetBaseUrl(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool outputWarningsToStandardError() const

/*
Returns true if warning messages will be output to stderr in addition to being emitted by the warnings() signal, otherwise false.

The default value is true.

See also setOutputWarningsToStandardError().
*/
func (this *QQmlEngine) OutputWarningsToStandardError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQmlEngine29outputWarningsToStandardErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOutputWarningsToStandardError(bool)

/*
Set whether warning messages will be output to stderr to enabled.

If enabled is true, any warning messages generated by QML will be output to stderr and emitted by the warnings() signal. If enabled is false, only the warnings() signal will be emitted. This allows applications to handle warning output themselves.

The default value is true.

See also outputWarningsToStandardError().
*/
func (this *QQmlEngine) SetOutputWarningsToStandardError(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine32setOutputWarningsToStandardErrorEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void retranslate()

/*
Refreshes all binding expressions that use strings marked for translation.

Call this function after you have installed a new translator with QCoreApplication::installTranslator, to ensure that your user-interface shows up-to-date translations.

Note: Due to a limitation in the implementation, this function refreshes all the engine's bindings, not only those that use strings marked for translation. This may be optimized in a future release.

This function was introduced in  Qt 5.10.
*/
func (this *QQmlEngine) Retranslate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine11retranslateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQmlContext * contextForObject(const QObject *)

/*
Returns the QQmlContext for the object, or 0 if no context has been set.

When the QQmlEngine instantiates a QObject, the context is set automatically.

See also setContextForObject(), qmlContext(), and qmlEngine().
*/
func (this *QQmlEngine) ContextForObject(arg0 qtcore.QObject_ITF /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine16contextForObjectEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQmlEngine_ContextForObject(arg0 qtcore.QObject_ITF /*777 const QObject **/) *QQmlContext /*777 QQmlContext **/ {
	var nilthis *QQmlEngine
	rv := nilthis.ContextForObject(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setContextForObject(QObject *, QQmlContext *)

/*
Sets the QQmlContext for the object to context. If the object already has a context, a warning is output, but the context is not changed.

When the QQmlEngine instantiates a QObject, the context is set automatically.

See also contextForObject().
*/
func (this *QQmlEngine) SetContextForObject(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 QQmlContext_ITF /*777 QQmlContext **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QQmlContext_PTR() != nil {
		convArg1 = arg1.QQmlContext_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine19setContextForObjectEP7QObjectP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QQmlEngine_SetContextForObject(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 QQmlContext_ITF /*777 QQmlContext **/) {
	var nilthis *QQmlEngine
	nilthis.SetContextForObject(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:155
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setObjectOwnership(QObject *, QQmlEngine::ObjectOwnership)

/*
Sets the ownership of object.

See also objectOwnership().
*/
func (this *QQmlEngine) SetObjectOwnership(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine18setObjectOwnershipEP7QObjectNS_15ObjectOwnershipE", qtrt.FFI_TYPE_POINTER, convArg0, arg1)
	qtrt.ErrPrint(err, rv)
}
func QQmlEngine_SetObjectOwnership(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 int) {
	var nilthis *QQmlEngine
	nilthis.SetObjectOwnership(arg0, arg1)
}

// /usr/include/qt/QtQml/qqmlengine.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [4] QQmlEngine::ObjectOwnership objectOwnership(QObject *)

/*
Returns the ownership of object.

See also setObjectOwnership().
*/
func (this *QQmlEngine) ObjectOwnership(arg0 qtcore.QObject_ITF /*777 QObject **/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine15objectOwnershipEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QQmlEngine_ObjectOwnership(arg0 qtcore.QObject_ITF /*777 QObject **/) int {
	var nilthis *QQmlEngine
	rv := nilthis.ObjectOwnership(arg0)
	return rv
}

// /usr/include/qt/QtQml/qqmlengine.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QQmlEngine) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlengine.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()

/*
This signal is emitted when the QML loaded by the engine would like to quit.

See also exit().
*/
func (this *QQmlEngine) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlengine.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)

/*
This signal is emitted when the QML loaded by the engine would like to exit from the event loop with the specified return code retCode.

This function was introduced in  Qt 5.8.

See also quit().
*/
func (this *QQmlEngine) Exit(retCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQmlEngine4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retCode)
	qtrt.ErrPrint(err, rv)
}

/*
ObjectOwnership controls whether or not QML automatically destroys the QObject when the corresponding JavaScript object is garbage collected by the engine. The two ownership options are:



Generally an application doesn't need to set an object's ownership explicitly. QML uses a heuristic to set the default ownership. By default, an object that is created by QML has JavaScriptOwnership. The exception to this are the root objects created by calling QQmlComponent::create() or QQmlComponent::beginCreate(), which have CppOwnership by default. The ownership of these root-level objects is considered to have been transferred to the C++ caller.

Objects not-created by QML have CppOwnership by default. The exception to this are objects returned from C++ method calls; their ownership will be set to JavaScriptOwnership. This applies only to explicit invocations of Q_INVOKABLE methods or slots, but not to property getter invocations.

Calling setObjectOwnership() overrides the default ownership heuristic used by QML.

*/
type QQmlEngine__ObjectOwnership = int

// The object is owned by C++ code and QML will never delete it. The JavaScript destroy() method cannot be used on these objects. This option is similar to QScriptEngine::QtOwnership.
const QQmlEngine__CppOwnership QQmlEngine__ObjectOwnership = 0

// The object is owned by JavaScript. When the object is returned to QML as the return value of a method call, QML will track it and delete it if there are no remaining JavaScript references to it and it has no QObject::parent(). An object tracked by one QQmlEngine will be deleted during that QQmlEngine's destructor. Thus, JavaScript references between objects with JavaScriptOwnership from two different engines will not be valid if one of these engines is deleted. This option is similar to QScriptEngine::ScriptOwnership.
const QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = 1

func (this *QQmlEngine) ObjectOwnershipItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQmlEngine_ObjectOwnershipItemName(val int) string {
	var nilthis *QQmlEngine
	return nilthis.ObjectOwnershipItemName(val)
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
