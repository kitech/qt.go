package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h
// #include <qaudiosystemplugin.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QAudioSystemPlugin struct {
	*qtcore.QObject
	*QAudioSystemFactoryInterface
}
type QAudioSystemPlugin_ITF interface {
	qtcore.QObject_ITF
	QAudioSystemFactoryInterface_ITF
	QAudioSystemPlugin_PTR() *QAudioSystemPlugin
}

func (ptr *QAudioSystemPlugin) QAudioSystemPlugin_PTR() *QAudioSystemPlugin { return ptr }

func (this *QAudioSystemPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAudioSystemPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QAudioSystemFactoryInterface = NewQAudioSystemFactoryInterfaceFromPointer(cthis)
}
func NewQAudioSystemPluginFromPointer(cthis unsafe.Pointer) *QAudioSystemPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQAudioSystemFactoryInterfaceFromPointer(cthis)
	return &QAudioSystemPlugin{bcthis0, bcthis1}
}
func (*QAudioSystemPlugin) NewFromPointer(cthis unsafe.Pointer) *QAudioSystemPlugin {
	return NewQAudioSystemPluginFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAudioSystemPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAudioSystemPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioSystemPlugin(QObject *)

/*
Construct a new audio plugin with parent. This is invoked automatically by the Q_PLUGIN_METADATA() macro.
*/
func (*QAudioSystemPlugin) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioSystemPlugin {
	return NewQAudioSystemPlugin(parent)
}
func NewQAudioSystemPlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QAudioSystemPlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioSystemPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioSystemPlugin")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAudioSystemPlugin(QObject *)

/*
Construct a new audio plugin with parent. This is invoked automatically by the Q_PLUGIN_METADATA() macro.
*/
func (*QAudioSystemPlugin) NewForInherit__() *QAudioSystemPlugin {
	return NewQAudioSystemPlugin__()
}
func NewQAudioSystemPlugin__() *QAudioSystemPlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAudioSystemPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAudioSystemPlugin")
	return gothis
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioSystemPlugin()

/*

 */
func DeleteQAudioSystemPlugin(this *QAudioSystemPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioInput * createInput(const QByteArray &)

/*
Returns a pointer to a QAbstractAudioInput created using device identifier
*/
func (this *QAudioSystemPlugin) CreateInput(device qtcore.QByteArray_ITF) *QAbstractAudioInput /*777 QAbstractAudioInput **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPlugin11createInputERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioInputFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioOutput * createOutput(const QByteArray &)

/*
Returns a pointer to a QAbstractAudioOutput created using device identifier
*/
func (this *QAudioSystemPlugin) CreateOutput(device qtcore.QByteArray_ITF) *QAbstractAudioOutput /*777 QAbstractAudioOutput **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPlugin12createOutputERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioOutputFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioDeviceInfo * createDeviceInfo(const QByteArray &, QAudio::Mode)

/*
Returns a pointer to a QAbstractAudioDeviceInfo created using device and mode
*/
func (this *QAudioSystemPlugin) CreateDeviceInfo(device qtcore.QByteArray_ITF, mode int) *QAbstractAudioDeviceInfo /*777 QAbstractAudioDeviceInfo **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAudioSystemPlugin16createDeviceInfoERK10QByteArrayN6QAudio4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
