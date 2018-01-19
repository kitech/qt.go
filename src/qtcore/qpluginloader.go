//  header block begin
// /usr/include/qt/QtCore/qpluginloader.h
// #include <qpluginloader.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qpluginloader.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPluginLoader) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// void QPluginLoader(class QObject *)
func NewQPluginLoader(parent unsafe.Pointer) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QPluginLoader{cthis}
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// void QPluginLoader(const class QString &, class QObject *)
func NewQPluginLoader_1(fileName unsafe.Pointer, parent unsafe.Pointer) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, parent)
	gopp.ErrPrint(err, rv)
	return &QPluginLoader{cthis}
}

// /usr/include/qt/QtCore/qpluginloader.h:64
// index:0
// virtual
// void ~QPluginLoader()
func DeleteQPluginLoader(*QPluginLoader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:66
// index:0
// QObject * instance()
func (this *QPluginLoader) Instance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader8instanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:67
// index:0
// QJsonObject metaData()
func (this *QPluginLoader) MetaData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8metaDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:69
// index:0
// static
// QObjectList staticInstances()
func (this *QPluginLoader) StaticInstances() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader15staticInstancesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPluginLoader_StaticInstances() {
	// 0: (), ()
	var nilthis *QPluginLoader
	nilthis.StaticInstances()
}

// /usr/include/qt/QtCore/qpluginloader.h:70
// index:0
// static
// QVector<QStaticPlugin> staticPlugins()
func (this *QPluginLoader) StaticPlugins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader13staticPluginsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPluginLoader_StaticPlugins() {
	// 0: (), ()
	var nilthis *QPluginLoader
	nilthis.StaticPlugins()
}

// /usr/include/qt/QtCore/qpluginloader.h:72
// index:0
// bool load()
func (this *QPluginLoader) Load() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader4loadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:73
// index:0
// bool unload()
func (this *QPluginLoader) Unload() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader6unloadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:74
// index:0
// bool isLoaded()
func (this *QPluginLoader) IsLoaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8isLoadedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:76
// index:0
// void setFileName(const class QString &)
func (this *QPluginLoader) SetFileName(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:77
// index:0
// QString fileName()
func (this *QPluginLoader) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:79
// index:0
// QString errorString()
func (this *QPluginLoader) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:82
// index:0
// QLibrary::LoadHints loadHints()
func (this *QPluginLoader) LoadHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader9loadHintsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
