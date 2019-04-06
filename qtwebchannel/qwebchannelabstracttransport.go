package qtwebchannel

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h
// #include <qwebchannelabstracttransport.h>
// #include <QtWebChannel>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QWebChannelAbstractTransport struct {
	*qtcore.QObject
}
type QWebChannelAbstractTransport_ITF interface {
	qtcore.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (this *QWebChannelAbstractTransport) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebChannelAbstractTransport) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebChannelAbstractTransportFromPointer(cthis unsafe.Pointer) *QWebChannelAbstractTransport {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebChannelAbstractTransport{bcthis0}
}
func (*QWebChannelAbstractTransport) NewFromPointer(cthis unsafe.Pointer) *QWebChannelAbstractTransport {
	return NewQWebChannelAbstractTransportFromPointer(cthis)
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:51
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebChannelAbstractTransport) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QWebChannelAbstractTransport10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebChannelAbstractTransport(QObject *)

/*
Constructs a transport object with the given parent.
*/
func (*QWebChannelAbstractTransport) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QWebChannelAbstractTransport {
	return NewQWebChannelAbstractTransport(parent)
}
func NewQWebChannelAbstractTransport(parent qtcore.QObject_ITF /*777 QObject **/) *QWebChannelAbstractTransport {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QWebChannelAbstractTransportC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebChannelAbstractTransportFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebChannelAbstractTransport")
	return gothis
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebChannelAbstractTransport(QObject *)

/*
Constructs a transport object with the given parent.
*/
func (*QWebChannelAbstractTransport) NewForInheritp() *QWebChannelAbstractTransport {
	return NewQWebChannelAbstractTransportp()
}
func NewQWebChannelAbstractTransportp() *QWebChannelAbstractTransport {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN28QWebChannelAbstractTransportC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebChannelAbstractTransportFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebChannelAbstractTransport")
	return gothis
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebChannelAbstractTransport()

/*

 */
func DeleteQWebChannelAbstractTransport(this *QWebChannelAbstractTransport) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QWebChannelAbstractTransportD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void sendMessage(const QJsonObject &)

/*
Sends a JSON message to the remote client. An implementation would serialize the message and transmit it to the remote JavaScript client.
*/
func (this *QWebChannelAbstractTransport) SendMessage(message qtcore.QJsonObject_ITF) {
	var convArg0 unsafe.Pointer
	if message != nil && message.QJsonObject_PTR() != nil {
		convArg0 = message.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QWebChannelAbstractTransport11sendMessageERK11QJsonObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qwebchannelabstracttransport.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void messageReceived(const QJsonObject &, QWebChannelAbstractTransport *)

/*
This signal must be emitted when a new JSON message was received from the remote client. The transport argument should be set to this transport object.
*/
func (this *QWebChannelAbstractTransport) MessageReceived(message qtcore.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF /*777 QWebChannelAbstractTransport **/) {
	var convArg0 unsafe.Pointer
	if message != nil && message.QJsonObject_PTR() != nil {
		convArg0 = message.QJsonObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if transport != nil && transport.QWebChannelAbstractTransport_PTR() != nil {
		convArg1 = transport.QWebChannelAbstractTransport_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QWebChannelAbstractTransport15messageReceivedERK11QJsonObjectPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11667() {
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
