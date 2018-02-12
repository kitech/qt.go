package qtgui

// /usr/include/qt/QtGui/qaccessibleplugin.h
// #include <qaccessibleplugin.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessiblePlugin struct {
	*qtcore.QObject
}

func (this *QAccessiblePlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAccessiblePlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAccessiblePluginFromPointer(cthis unsafe.Pointer) *QAccessiblePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAccessiblePlugin{bcthis0}
}
func (*QAccessiblePlugin) NewFromPointer(cthis unsafe.Pointer) *QAccessiblePlugin {
	return NewQAccessiblePluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAccessiblePlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessiblePlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessiblePlugin(QObject *)
func NewQAccessiblePlugin(parent *qtcore.QObject /*777 QObject **/) *QAccessiblePlugin {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessiblePluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessiblePluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessiblePlugin()
func DeleteQAccessiblePlugin(this *QAccessiblePlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessiblePluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * create(const QString &, QObject *)
func (this *QAccessiblePlugin) Create(key string, object *qtcore.QObject /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessiblePlugin6createERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
