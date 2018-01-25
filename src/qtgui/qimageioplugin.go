package qtgui

// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>

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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QImageIOPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQImageIOPluginFromPointer(cthis unsafe.Pointer) *QImageIOPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QImageIOPlugin{bcthis0}
}
func (*QImageIOPlugin) NewFromPointer(cthis unsafe.Pointer) *QImageIOPlugin {
	return NewQImageIOPluginFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimageiohandler.h:141
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QImageIOPlugin) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:143
// index:0
// Public
// void QImageIOPlugin(class QObject *)
func NewQImageIOPlugin(parent *qtcore.QObject /*444 QObject **/) *QImageIOPlugin {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QImageIOPluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QImageIOPlugin) Capabilities(device *qtcore.QIODevice /*444 QIODevice **/, format *qtcore.QByteArray) int {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin12capabilitiesEP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:154
// index:0
// Public pure virtual
// QImageIOHandler * create(class QIODevice *, const class QByteArray &)
func (this *QImageIOPlugin) Create(device *qtcore.QIODevice /*444 QIODevice **/, format *qtcore.QByteArray) *QImageIOHandler /*444 QImageIOHandler **/ {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QImageIOPlugin__Capability = int

const QImageIOPlugin__CanRead QImageIOPlugin__Capability = 1
const QImageIOPlugin__CanWrite QImageIOPlugin__Capability = 2
const QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = 4

//  body block end
