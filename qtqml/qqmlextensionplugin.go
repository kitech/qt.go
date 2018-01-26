package qtqml

// /usr/include/qt/QtQml/qqmlextensionplugin.h
// #include <qqmlextensionplugin.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QQmlExtensionPlugin struct {
	*qtcore.QObject
	*QQmlExtensionInterface
}

func (this *QQmlExtensionPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlExtensionPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QQmlExtensionInterface = NewQQmlExtensionInterfaceFromPointer(cthis)
}
func NewQQmlExtensionPluginFromPointer(cthis unsafe.Pointer) *QQmlExtensionPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQQmlExtensionInterfaceFromPointer(cthis)
	return &QQmlExtensionPlugin{bcthis0, bcthis1}
}
func (*QQmlExtensionPlugin) NewFromPointer(cthis unsafe.Pointer) *QQmlExtensionPlugin {
	return NewQQmlExtensionPluginFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlExtensionPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQmlExtensionPlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:61
// index:0
// Public
// void QQmlExtensionPlugin(class QObject *)
func NewQQmlExtensionPlugin(parent *qtcore.QObject /*777 QObject **/) *QQmlExtensionPlugin {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQmlExtensionPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlExtensionPluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:62
// index:0
// Public virtual
// void ~QQmlExtensionPlugin()
func DeleteQQmlExtensionPlugin(*QQmlExtensionPlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQmlExtensionPluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:64
// index:0
// Public
// QUrl baseUrl()
func (this *QQmlExtensionPlugin) BaseUrl() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQmlExtensionPlugin7baseUrlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:66
// index:0
// Public pure virtual
// void registerTypes(const char *)
func (this *QQmlExtensionPlugin) RegisterTypes(uri string) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQmlExtensionPlugin13registerTypesEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:67
// index:0
// Public virtual
// void initializeEngine(class QQmlEngine *, const char *)
func (this *QQmlExtensionPlugin) InitializeEngine(engine *QQmlEngine /*777 QQmlEngine **/, uri string) {
	var convArg0 = engine.GetCthis()
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQmlExtensionPlugin16initializeEngineEP10QQmlEnginePKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
