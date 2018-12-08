package qtcore

// /usr/include/qt/QtCore/qpluginloader.h
// #include <qpluginloader.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QPluginLoader struct {
	*QObject
}
type QPluginLoader_ITF interface {
	QObject_ITF
	QPluginLoader_PTR() *QPluginLoader
}

func (ptr *QPluginLoader) QPluginLoader_PTR() *QPluginLoader { return ptr }

func (this *QPluginLoader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPluginLoader) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQPluginLoaderFromPointer(cthis unsafe.Pointer) *QPluginLoader {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QPluginLoader{bcthis0}
}
func (*QPluginLoader) NewFromPointer(cthis unsafe.Pointer) *QPluginLoader {
	return NewQPluginLoaderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpluginloader.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPluginLoader) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(QObject *)

/*
Constructs a plugin loader with the given parent.
*/
func (*QPluginLoader) NewForInherit(parent QObject_ITF /*777 QObject **/) *QPluginLoader {
	return NewQPluginLoader(parent)
}
func NewQPluginLoader(parent QObject_ITF /*777 QObject **/) *QPluginLoader {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPluginLoader")
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(QObject *)

/*
Constructs a plugin loader with the given parent.
*/
func (*QPluginLoader) NewForInheritp() *QPluginLoader {
	return NewQPluginLoaderp()
}
func NewQPluginLoaderp() *QPluginLoader {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPluginLoader")
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(const QString &, QObject *)

/*
Constructs a plugin loader with the given parent.
*/
func (*QPluginLoader) NewForInherit1(fileName string, parent QObject_ITF /*777 QObject **/) *QPluginLoader {
	return NewQPluginLoader1(fileName, parent)
}
func NewQPluginLoader1(fileName string, parent QObject_ITF /*777 QObject **/) *QPluginLoader {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPluginLoader")
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(const QString &, QObject *)

/*
Constructs a plugin loader with the given parent.
*/
func (*QPluginLoader) NewForInherit1p(fileName string) *QPluginLoader {
	return NewQPluginLoader1p(fileName)
}
func NewQPluginLoader1p(fileName string) *QPluginLoader {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPluginLoader")
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPluginLoader()

/*

 */
func DeleteQPluginLoader(this *QPluginLoader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qpluginloader.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * instance()

/*
Returns the root component object of the plugin. The plugin is loaded if necessary. The function returns 0 if the plugin could not be loaded or if the root component object could not be instantiated.

If the root component object was destroyed, calling this function creates a new instance.

The root component, returned by this function, is not deleted when the QPluginLoader is destroyed. If you want to ensure that the root component is deleted, you should call unload() as soon you don't need to access the core component anymore. When the library is finally unloaded, the root component will automatically be deleted.

The component object is a QObject. Use qobject_cast() to access interfaces you are interested in.

See also load().
*/
func (this *QPluginLoader) Instance() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader8instanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpluginloader.h:67
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject metaData() const

/*
Returns the meta data for this plugin. The meta data is data specified in a json format using the Q_PLUGIN_METADATA() macro when compiling the plugin.

The meta data can be queried in a fast and inexpensive way without actually loading the plugin. This makes it possible to e.g. store capabilities of the plugin in there, and make the decision whether to load the plugin dependent on this meta data.
*/
func (this *QPluginLoader) MetaData() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8metaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [8] QObjectList staticInstances()

/*
Returns a list of static plugin instances (root components) held by the plugin loader.

See also staticPlugins().
*/
func (this *QPluginLoader) StaticInstances() *QObjectList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader15staticInstancesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQObjectListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}
func QPluginLoader_StaticInstances() *QObjectList /*667*/ {
	var nilthis *QPluginLoader
	rv := nilthis.StaticInstances()
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load()

/*
Loads the plugin and returns true if the plugin was loaded successfully; otherwise returns false. Since instance() always calls this function before resolving any symbols it is not necessary to call it explicitly. In some situations you might want the plugin loaded in advance, in which case you would use this function.

See also unload().
*/
func (this *QPluginLoader) Load() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader4loadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unload()

/*
Unloads the plugin and returns true if the plugin could be unloaded; otherwise returns false.

This happens automatically on application termination, so you shouldn't normally need to call this function.

If other instances of QPluginLoader are using the same plugin, the call will fail, and unloading will only happen when every instance has called unload().

Don't try to delete the root component. Instead rely on that unload() will automatically delete it when needed.

See also instance() and load().
*/
func (this *QPluginLoader) Unload() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader6unloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoaded() const

/*
Returns true if the plugin is loaded; otherwise returns false.

See also load().
*/
func (this *QPluginLoader) IsLoaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8isLoadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*

 */
func (this *QPluginLoader) SetFileName(fileName string) {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*

 */
func (this *QPluginLoader) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qpluginloader.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a text string with the description of the last error that occurred.

This function was introduced in  Qt 4.2.
*/
func (this *QPluginLoader) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qpluginloader.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoadHints(QLibrary::LoadHints)

/*

 */
func (this *QPluginLoader) SetLoadHints(loadHints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader12setLoadHintsE6QFlagsIN8QLibrary8LoadHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), loadHints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QLibrary::LoadHints loadHints() const

/*

 */
func (this *QPluginLoader) LoadHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader9loadHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
}

//  keep block end
