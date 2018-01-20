//  header block begin
// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QImageIOPlugin struct {
	*qtcore.QObject
}

func (this *QImageIOPlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQImageIOPluginFromPointer(cthis unsafe.Pointer) *QImageIOPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QImageIOPlugin{bcthis0}
}

// /usr/include/qt/QtGui/qimageiohandler.h:141
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QImageIOPlugin) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:143
// index:0
// Public
// void QImageIOPlugin(class QObject *)
func NewQImageIOPlugin(parent unsafe.Pointer) *QImageIOPlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QImageIOPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOPluginFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:144
// index:0
// Public virtual
// void ~QImageIOPlugin()
func DeleteQImageIOPlugin(*QImageIOPlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QImageIOPluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:153
// index:0
// Public pure virtual
// QImageIOPlugin::Capabilities capabilities(class QIODevice *, const class QByteArray &)
func (this *QImageIOPlugin) Capabilities(device unsafe.Pointer, format *qtcore.QByteArray) interface{} {
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin12capabilitiesEP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:154
// index:0
// Public pure virtual
// QImageIOHandler * create(class QIODevice *, const class QByteArray &)
func (this *QImageIOPlugin) Create(device unsafe.Pointer, format *qtcore.QByteArray) interface{} {
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
