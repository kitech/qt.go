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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QImageIOPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimageiohandler.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageIOPlugin(QObject *)
func NewQImageIOPlugin(parent *qtcore.QObject /*777 QObject **/) *QImageIOPlugin {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QImageIOPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImageIOPlugin()
func DeleteQImageIOPlugin(this *QImageIOPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QImageIOPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimageiohandler.h:153
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QImageIOPlugin::Capabilities capabilities(QIODevice *, const QByteArray &)
func (this *QImageIOPlugin) Capabilities(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray) int {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin12capabilitiesEP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:154
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QImageIOHandler * create(QIODevice *, const QByteArray &)
func (this *QImageIOPlugin) Create(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray) *QImageIOHandler /*777 QImageIOHandler **/ {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

type QImageIOPlugin__Capability = int

const QImageIOPlugin__CanRead QImageIOPlugin__Capability = 1
const QImageIOPlugin__CanWrite QImageIOPlugin__Capability = 2
const QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = 4

//  body block end
