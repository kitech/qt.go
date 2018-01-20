//  header block begin
// /usr/include/qt/QtGui/qgenericplugin.h
// #include <qgenericplugin.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QGenericPlugin struct {
	*qtcore.QObject
}

func (this *QGenericPlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qgenericplugin.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGenericPlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QGenericPlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qgenericplugin.h:55
// index:0
// void QGenericPlugin(class QObject *)
func NewQGenericPlugin(parent unsafe.Pointer) *QGenericPlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGenericPluginFromPointer(cthis)
	return gothis
}
func NewQGenericPluginFromPointer(cthis unsafe.Pointer) *QGenericPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGenericPlugin{bcthis0}
}

// /usr/include/qt/QtGui/qgenericplugin.h:56
// index:0
// virtual
// void ~QGenericPlugin()
func DeleteQGenericPlugin(*QGenericPlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qgenericplugin.h:58
// index:0
// pure virtual
// QObject * create(const class QString &, const class QString &)
func (this *QGenericPlugin) Create(name unsafe.Pointer, spec unsafe.Pointer) {
	// 0: (, name const QString &, spec const QString &), (name, spec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QGenericPlugin6createERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name, spec)
	gopp.ErrPrint(err, rv)
}

//  body block end
