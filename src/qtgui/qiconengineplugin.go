//  header block begin
// /usr/include/qt/QtGui/qiconengineplugin.h
// #include <qiconengineplugin.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QIconEnginePlugin struct {
	*qtcore.QObject
}

func (this *QIconEnginePlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qiconengineplugin.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QIconEnginePlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QIconEnginePlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:58
// index:0
// void QIconEnginePlugin(class QObject *)
func NewQIconEnginePlugin(parent unsafe.Pointer) *QIconEnginePlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconEnginePluginFromPointer(cthis)
	return gothis
}
func NewQIconEnginePluginFromPointer(cthis unsafe.Pointer) *QIconEnginePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QIconEnginePlugin{bcthis0}
}

// /usr/include/qt/QtGui/qiconengineplugin.h:59
// index:0
// virtual
// void ~QIconEnginePlugin()
func DeleteQIconEnginePlugin(*QIconEnginePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengineplugin.h:61
// index:0
// pure virtual
// QIconEngine * create(const class QString &)
func (this *QIconEnginePlugin) Create(filename unsafe.Pointer) {
	// 0: (, filename const QString &), (filename)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIconEnginePlugin6createERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filename)
	gopp.ErrPrint(err, rv)
}

//  body block end
