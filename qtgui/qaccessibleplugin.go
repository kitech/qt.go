package qtgui

// /usr/include/qt/QtGui/qaccessibleplugin.h
// #include <qaccessibleplugin.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

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
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessiblePlugin(QObject *)
func NewQAccessiblePlugin(parent *qtcore.QObject /*777 QObject **/) *QAccessiblePlugin {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessiblePluginC1EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * create(const QString &, QObject *)
func (this *QAccessiblePlugin) Create(key *qtcore.QString, object *qtcore.QObject /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var convArg0 = key.GetCthis()
	var convArg1 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessiblePlugin6createERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
