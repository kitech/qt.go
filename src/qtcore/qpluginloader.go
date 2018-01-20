//  header block begin
// /usr/include/qt/QtCore/qpluginloader.h
// #include <qpluginloader.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QPluginLoader struct {
	*QObject
}

func (this *QPluginLoader) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQPluginLoaderFromPointer(cthis unsafe.Pointer) *QPluginLoader {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QPluginLoader{bcthis0}
}

// /usr/include/qt/QtCore/qpluginloader.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPluginLoader) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// Public
// void QPluginLoader(class QObject *)
func NewQPluginLoader(parent unsafe.Pointer) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// Public
// void QPluginLoader(const class QString &, class QObject *)
func NewQPluginLoader_1(fileName *QString, parent unsafe.Pointer) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:64
// index:0
// Public virtual
// void ~QPluginLoader()
func DeleteQPluginLoader(*QPluginLoader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:66
// index:0
// Public
// QObject * instance()
func (this *QPluginLoader) Instance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader8instanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:67
// index:0
// Public
// QJsonObject metaData()
func (this *QPluginLoader) MetaData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8metaDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:69
// index:0
// Public static
// QObjectList staticInstances()
func (this *QPluginLoader) StaticInstances() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader15staticInstancesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPluginLoader_StaticInstances() {
	var nilthis *QPluginLoader
	nilthis.StaticInstances()
}

// /usr/include/qt/QtCore/qpluginloader.h:70
// index:0
// Public static
// QVector<QStaticPlugin> staticPlugins()
func (this *QPluginLoader) StaticPlugins() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader13staticPluginsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPluginLoader_StaticPlugins() {
	var nilthis *QPluginLoader
	nilthis.StaticPlugins()
}

// /usr/include/qt/QtCore/qpluginloader.h:72
// index:0
// Public
// bool load()
func (this *QPluginLoader) Load() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader4loadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:73
// index:0
// Public
// bool unload()
func (this *QPluginLoader) Unload() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader6unloadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:74
// index:0
// Public
// bool isLoaded()
func (this *QPluginLoader) IsLoaded() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8isLoadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:76
// index:0
// Public
// void setFileName(const class QString &)
func (this *QPluginLoader) SetFileName(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:77
// index:0
// Public
// QString fileName()
func (this *QPluginLoader) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:79
// index:0
// Public
// QString errorString()
func (this *QPluginLoader) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qpluginloader.h:82
// index:0
// Public
// QLibrary::LoadHints loadHints()
func (this *QPluginLoader) LoadHints() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader9loadHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
