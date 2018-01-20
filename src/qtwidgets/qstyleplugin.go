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

// /usr/include/qt/QtWidgets/qstyleplugin.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStylePlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QStylePlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:58
// index:0
// void QStylePlugin(class QObject *)
func NewQStylePlugin(parent unsafe.Pointer) *QStylePlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStylePluginFromPointer(cthis)
	return gothis
}
func NewQStylePluginFromPointer(cthis unsafe.Pointer) *QStylePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStylePlugin{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:59
// index:0
// virtual
// void ~QStylePlugin()
func DeleteQStylePlugin(*QStylePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleplugin.h:61
// index:0
// pure virtual
// QStyle * create(const class QString &)
func (this *QStylePlugin) Create(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStylePlugin6createERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

//  body block end
