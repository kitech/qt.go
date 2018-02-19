package qtqml

// /usr/include/qt/QtQml/qqmlextensionplugin.h
// #include <qqmlextensionplugin.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QQmlExtensionPlugin struct {
	*qtcore.QObject
	*QQmlExtensionInterface
}
type QQmlExtensionPlugin_ITF interface {
	qtcore.QObject_ITF
	QQmlExtensionInterface_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QQmlExtensionPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQmlExtensionPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlExtensionPlugin(QObject *)
func NewQQmlExtensionPlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlExtensionPlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQmlExtensionPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExtensionPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExtensionPlugin")
	return gothis
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlExtensionPlugin(QObject *)
func NewQQmlExtensionPlugin__() *QQmlExtensionPlugin {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQmlExtensionPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlExtensionPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlExtensionPlugin")
	return gothis
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlExtensionPlugin()
func DeleteQQmlExtensionPlugin(this *QQmlExtensionPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQmlExtensionPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl() const
func (this *QQmlExtensionPlugin) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQmlExtensionPlugin7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerTypes(const char *)
func (this *QQmlExtensionPlugin) RegisterTypes(uri string) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQmlExtensionPlugin13registerTypesEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlextensionplugin.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void initializeEngine(QQmlEngine *, const char *)
func (this *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF /*777 QQmlEngine **/, uri string) {
	var convArg0 unsafe.Pointer
	if engine != nil && engine.QQmlEngine_PTR() != nil {
		convArg0 = engine.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQmlExtensionPlugin16initializeEngineEP10QQmlEnginePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
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
