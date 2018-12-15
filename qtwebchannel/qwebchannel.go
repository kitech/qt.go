package qtwebchannel

// /usr/include/qt/QtWebChannel/qwebchannel.h
// #include <qwebchannel.h>
// #include <QtWebChannel>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QWebChannel struct {
	*qtcore.QObject
}
type QWebChannel_ITF interface {
	qtcore.QObject_ITF
	QWebChannel_PTR() *QWebChannel
}

func (ptr *QWebChannel) QWebChannel_PTR() *QWebChannel { return ptr }

func (this *QWebChannel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebChannel) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebChannelFromPointer(cthis unsafe.Pointer) *QWebChannel {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebChannel{bcthis0}
}
func (*QWebChannel) NewFromPointer(cthis unsafe.Pointer) *QWebChannel {
	return NewQWebChannelFromPointer(cthis)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebChannel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWebChannel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebChannel(QObject *)

/*
Constructs the QWebChannel object with the given parent.

Note that a QWebChannel is only fully operational once you connect it to a QWebChannelAbstractTransport. The HTML clients also need to be setup appropriately using qwebchannel.js.
*/
func (*QWebChannel) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QWebChannel {
	return NewQWebChannel(parent)
}
func NewQWebChannel(parent qtcore.QObject_ITF /*777 QObject **/) *QWebChannel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebChannelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebChannel")
	return gothis
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebChannel(QObject *)

/*
Constructs the QWebChannel object with the given parent.

Note that a QWebChannel is only fully operational once you connect it to a QWebChannelAbstractTransport. The HTML clients also need to be setup appropriately using qwebchannel.js.
*/
func (*QWebChannel) NewForInheritp() *QWebChannel {
	return NewQWebChannelp()
}
func NewQWebChannelp() *QWebChannel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebChannelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebChannel")
	return gothis
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebChannel()

/*

 */
func DeleteQWebChannel(this *QWebChannel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void registerObject(const QString &, QObject *)

/*
Registers a single object to the QWebChannel.

The properties, signals and public methods of the object are published to the remote clients. There, an object with the identifier id is then constructed.

Note: A current limitation is that objects must be registered before any client is initialized.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also QWebChannel::registerObjects(), QWebChannel::deregisterObject(), and QWebChannel::registeredObjects().
*/
func (this *QWebChannel) RegisterObject(id string, object qtcore.QObject_ITF /*777 QObject **/) {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg1 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel14registerObjectERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deregisterObject(QObject *)

/*
Deregisters the given object from the QWebChannel.

Remote clients will receive a destroyed signal for the given object.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also QWebChannel::registerObjects(), QWebChannel::registerObject(), and QWebChannel::registeredObjects().
*/
func (this *QWebChannel) DeregisterObject(object qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel16deregisterObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool blockUpdates() const

/*

 */
func (this *QWebChannel) BlockUpdates() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWebChannel12blockUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlockUpdates(bool)

/*

 */
func (this *QWebChannel) SetBlockUpdates(block bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel15setBlockUpdatesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), block)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blockUpdatesChanged(bool)

/*

 */
func (this *QWebChannel) BlockUpdatesChanged(block bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel19blockUpdatesChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), block)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectTo(QWebChannelAbstractTransport *)

/*
Connects the QWebChannel to the given transport object.

The transport object then handles the communication between the C++ application and a remote HTML client.

See also QWebChannelAbstractTransport and QWebChannel::disconnectFrom().
*/
func (this *QWebChannel) ConnectTo(transport QWebChannelAbstractTransport_ITF /*777 QWebChannelAbstractTransport **/) {
	var convArg0 unsafe.Pointer
	if transport != nil && transport.QWebChannelAbstractTransport_PTR() != nil {
		convArg0 = transport.QWebChannelAbstractTransport_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel9connectToEP28QWebChannelAbstractTransport", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannel.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnectFrom(QWebChannelAbstractTransport *)

/*
Disconnects the QWebChannel from the transport object.

See also QWebChannel::connectTo().
*/
func (this *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransport_ITF /*777 QWebChannelAbstractTransport **/) {
	var convArg0 unsafe.Pointer
	if transport != nil && transport.QWebChannelAbstractTransport_PTR() != nil {
		convArg0 = transport.QWebChannelAbstractTransport_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWebChannel14disconnectFromEP28QWebChannelAbstractTransport", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
		qtnetwork.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
