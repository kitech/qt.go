package qtcore

// /usr/include/qt/QtCore/qpluginloader.h
// #include <qpluginloader.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QPluginLoader) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// Public
// void QPluginLoader(QObject *)
func NewQPluginLoader(parent *QObject /*777 QObject **/) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// Public
// void QPluginLoader(const QString &, QObject *)
func NewQPluginLoader_1(fileName *QString, parent *QObject /*777 QObject **/) *QPluginLoader {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
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
func (this *QPluginLoader) Instance() *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader8instanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:67
// index:0
// Public
// QJsonObject metaData()
func (this *QPluginLoader) MetaData() *QJsonObject /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8metaDataEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:72
// index:0
// Public
// bool load()
func (this *QPluginLoader) Load() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader4loadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:73
// index:0
// Public
// bool unload()
func (this *QPluginLoader) Unload() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader6unloadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:74
// index:0
// Public
// bool isLoaded()
func (this *QPluginLoader) IsLoaded() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8isLoadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:76
// index:0
// Public
// void setFileName(const QString &)
func (this *QPluginLoader) SetFileName(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QPluginLoader11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:77
// index:0
// Public
// QString fileName()
func (this *QPluginLoader) FileName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:79
// index:0
// Public
// QString errorString()
func (this *QPluginLoader) ErrorString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:82
// index:0
// Public
// QLibrary::LoadHints loadHints()
func (this *QPluginLoader) LoadHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QPluginLoader9loadHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
