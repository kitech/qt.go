package qtgui

// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QImageIOPlugin struct {
	*qtcore.QObject
}
type QImageIOPlugin_ITF interface {
	qtcore.QObject_ITF
	QImageIOPlugin_PTR() *QImageIOPlugin
}

func (ptr *QImageIOPlugin) QImageIOPlugin_PTR() *QImageIOPlugin { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QImageIOPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimageiohandler.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageIOPlugin(QObject *)

/*

 */
func (*QImageIOPlugin) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QImageIOPlugin {
	return NewQImageIOPlugin(parent)
}
func NewQImageIOPlugin(parent qtcore.QObject_ITF /*777 QObject **/) *QImageIOPlugin {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QImageIOPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageIOPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QImageIOPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageIOPlugin(QObject *)

/*

 */
func (*QImageIOPlugin) NewForInherit__() *QImageIOPlugin {
	return NewQImageIOPlugin__()
}
func NewQImageIOPlugin__() *QImageIOPlugin {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QImageIOPluginC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageIOPluginFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QImageIOPlugin")
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImageIOPlugin()

/*

 */
func DeleteQImageIOPlugin(this *QImageIOPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QImageIOPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimageiohandler.h:153
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QImageIOPlugin::Capabilities capabilities(QIODevice *, const QByteArray &) const

/*

 */
func (this *QImageIOPlugin) Capabilities(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin12capabilitiesEP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:154
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QImageIOHandler * create(QIODevice *, const QByteArray &) const

/*

 */
func (this *QImageIOPlugin) Create(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QImageIOHandler /*777 QImageIOHandler **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimageiohandler.h:154
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QImageIOHandler * create(QIODevice *, const QByteArray &) const

/*

 */
func (this *QImageIOPlugin) Create__(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QImageIOHandler /*777 QImageIOHandler **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*


 */
type QImageIOPlugin__Capability = int

//
const QImageIOPlugin__CanRead QImageIOPlugin__Capability = 1

//
const QImageIOPlugin__CanWrite QImageIOPlugin__Capability = 2

//
const QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = 4

func (this *QImageIOPlugin) CapabilityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QImageIOPlugin_CapabilityItemName(val int) string {
	var nilthis *QImageIOPlugin
	return nilthis.CapabilityItemName(val)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
