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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPluginLoader) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpluginloader.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(QObject *)
func NewQPluginLoader(parent *QObject /*777 QObject **/) *QPluginLoader {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPluginLoader(const QString &, QObject *)
func NewQPluginLoader_1(fileName string, parent *QObject /*777 QObject **/) *QPluginLoader {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPluginLoaderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qpluginloader.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPluginLoader()
func DeleteQPluginLoader(this *QPluginLoader) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoaderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qpluginloader.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * instance()
func (this *QPluginLoader) Instance() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader8instanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpluginloader.h:67
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject metaData()
func (this *QPluginLoader) MetaData() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8metaDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qpluginloader.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load()
func (this *QPluginLoader) Load() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader4loadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unload()
func (this *QPluginLoader) Unload() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader6unloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoaded()
func (this *QPluginLoader) IsLoaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8isLoadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpluginloader.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QPluginLoader) SetFileName(fileName string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QPluginLoader) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qpluginloader.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QPluginLoader) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qpluginloader.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoadHints(QLibrary::LoadHints)
func (this *QPluginLoader) SetLoadHints(loadHints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QPluginLoader12setLoadHintsE6QFlagsIN8QLibrary8LoadHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), loadHints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpluginloader.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QLibrary::LoadHints loadHints()
func (this *QPluginLoader) LoadHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QPluginLoader9loadHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

//  body block end
