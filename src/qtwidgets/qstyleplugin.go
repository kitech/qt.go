//  header block begin
// /usr/include/qt/QtWidgets/qstyleplugin.h
// #include <qstyleplugin.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStylePlugin struct {
	*qtcore.QObject
}

func (this *QStylePlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQStylePluginFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStylePlugin{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStylePlugin) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStylePlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// Public
// void QStylePlugin(class QObject *)
func NewQStylePlugin(parent unsafe.Pointer) *QStylePlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:59
// index:0
// Public virtual
// void ~QStylePlugin()
func DeleteQStylePlugin(*QStylePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:61
// index:0
// Public pure virtual
// QStyle * create(const class QString &)
func (this *QStylePlugin) Create(key *qtcore.QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePlugin6createERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
