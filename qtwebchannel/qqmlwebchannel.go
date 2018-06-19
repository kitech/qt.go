package qtwebchannel

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h
// #include <qqmlwebchannel.h>
// #include <QtWebChannel>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QQmlWebChannel struct {
	*QWebChannel
}
type QQmlWebChannel_ITF interface {
	QWebChannel_ITF
	QQmlWebChannel_PTR() *QQmlWebChannel
}

func (ptr *QQmlWebChannel) QQmlWebChannel_PTR() *QQmlWebChannel { return ptr }

func (this *QQmlWebChannel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWebChannel.GetCthis()
	}
}
func (this *QQmlWebChannel) SetCthis(cthis unsafe.Pointer) {
	this.QWebChannel = NewQWebChannelFromPointer(cthis)
}
func NewQQmlWebChannelFromPointer(cthis unsafe.Pointer) *QQmlWebChannel {
	bcthis0 := NewQWebChannelFromPointer(cthis)
	return &QQmlWebChannel{bcthis0}
}
func (*QQmlWebChannel) NewFromPointer(cthis unsafe.Pointer) *QQmlWebChannel {
	return NewQQmlWebChannelFromPointer(cthis)
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlWebChannel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QQmlWebChannel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlWebChannel(QObject *)

/*

 */
func NewQQmlWebChannel(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlWebChannel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlWebChannelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlWebChannelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlWebChannel")
	return gothis
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlWebChannel(QObject *)

/*

 */
func NewQQmlWebChannel__() *QQmlWebChannel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlWebChannelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlWebChannelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlWebChannel")
	return gothis
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlWebChannel()

/*

 */
func DeleteQQmlWebChannel(this *QQmlWebChannel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlWebChannelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectTo(QObject *)

/*

 */
func (this *QQmlWebChannel) ConnectTo(transport qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if transport != nil && transport.QObject_PTR() != nil {
		convArg0 = transport.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlWebChannel9connectToEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebChannel/qqmlwebchannel.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnectFrom(QObject *)

/*

 */
func (this *QQmlWebChannel) DisconnectFrom(transport qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if transport != nil && transport.QObject_PTR() != nil {
		convArg0 = transport.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QQmlWebChannel14disconnectFromEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
