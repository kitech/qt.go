//  header block begin
// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:67
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAccessibleBridgePlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QAccessibleBridgePlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:69
// index:0
// void QAccessibleBridgePlugin(class QObject *)
func NewQAccessibleBridgePlugin(parent unsafe.Pointer) *QAccessibleBridgePlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleBridgePluginFromPointer(cthis)
	return gothis
}
func NewQAccessibleBridgePluginFromPointer(cthis unsafe.Pointer) *QAccessibleBridgePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAccessibleBridgePlugin{bcthis0}
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:70
// index:0
// virtual
// void ~QAccessibleBridgePlugin()
func DeleteQAccessibleBridgePlugin(*QAccessibleBridgePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:72
// index:0
// pure virtual
// QAccessibleBridge * create(const class QString &)
func (this *QAccessibleBridgePlugin) Create(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QAccessibleBridgePlugin6createERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

//  body block end
