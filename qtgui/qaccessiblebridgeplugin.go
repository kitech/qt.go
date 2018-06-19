package qtgui

// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>

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

//  ext block end

//  body block begin

/*

 */
type QAccessibleBridgePlugin struct {
	*qtcore.QObject
}
type QAccessibleBridgePlugin_ITF interface {
	qtcore.QObject_ITF
	QAccessibleBridgePlugin_PTR() *QAccessibleBridgePlugin
}

func (ptr *QAccessibleBridgePlugin) QAccessibleBridgePlugin_PTR() *QAccessibleBridgePlugin { return ptr }

func (this *QAccessibleBridgePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAccessibleBridgePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAccessibleBridgePluginFromPointer(cthis unsafe.Pointer) *QAccessibleBridgePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAccessibleBridgePlugin{bcthis0}
}
func (*QAccessibleBridgePlugin) NewFromPointer(cthis unsafe.Pointer) *QAccessibleBridgePlugin {
	return NewQAccessibleBridgePluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAccessibleBridgePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QAccessibleBridgePlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleBridgePlugin(QObject *)

/*

 */
func NewQAccessibleBridgePlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QAccessibleBridgePlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleBridgePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAccessibleBridgePlugin")
	return gothis
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleBridgePlugin(QObject *)

/*

 */
func NewQAccessibleBridgePlugin__() *QAccessibleBridgePlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleBridgePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAccessibleBridgePlugin")
	return gothis
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleBridgePlugin()

/*

 */
func DeleteQAccessibleBridgePlugin(this *QAccessibleBridgePlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleBridge * create(const QString &)

/*

 */
func (this *QAccessibleBridgePlugin) Create(key string) *QAccessibleBridge /*777 QAccessibleBridge **/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QAccessibleBridgePlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleBridgeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
}

//  keep block end
