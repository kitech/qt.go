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
func NewQAccessiblePluginFromPointer(cthis unsafe.Pointer) *QAccessiblePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAccessiblePlugin{bcthis0}
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:61
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAccessiblePlugin) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessiblePlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:63
// index:0
// Public
// void QAccessiblePlugin(class QObject *)
func NewQAccessiblePlugin(parent *qtcore.QObject /*444 QObject **/) *QAccessiblePlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessiblePluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:64
// index:0
// Public virtual
// void ~QAccessiblePlugin()
func DeleteQAccessiblePlugin(*QAccessiblePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:66
// index:0
// Public pure virtual
// QAccessibleInterface * create(const class QString &, class QObject *)
func (this *QAccessiblePlugin) Create(key *qtcore.QString, object *qtcore.QObject /*444 QObject **/) *QAccessibleInterface /*444 QAccessibleInterface **/ {
	var convArg0 = key.GetCthis()
	var convArg1 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePlugin6createERK7QStringP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
