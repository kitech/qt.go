package qtgui

// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleBridgePlugin struct {
	*qtcore.QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QAccessibleBridgePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QAccessibleBridgePlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleBridgePlugin(QObject *)
func NewQAccessibleBridgePlugin(parent *qtcore.QObject /*777 QObject **/) *QAccessibleBridgePlugin {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleBridgePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleBridgePlugin()
func DeleteQAccessibleBridgePlugin(*QAccessibleBridgePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleBridge * create(const QString &)
func (this *QAccessibleBridgePlugin) Create(key *qtcore.QString) *QAccessibleBridge /*777 QAccessibleBridge **/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePlugin6createERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleBridgeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
