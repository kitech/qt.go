package qtnetwork

// /usr/include/qt/QtNetwork/qabstractsocket.h
// #include <qabstractsocket.h>
// #include <QtNetwork>

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
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// long long readData(char *, qint64)
func (this *QAbstractSocket) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long readLineData(char *, qint64)
func (this *QAbstractSocket) InheritReadLineData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readLineData", f)
}

// long long writeData(const char *, qint64)
func (this *QAbstractSocket) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

// void setSocketState(enum QAbstractSocket::SocketState)
func (this *QAbstractSocket) InheritSetSocketState(f func(state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSocketState", f)
}

// void setSocketError(enum QAbstractSocket::SocketError)
func (this *QAbstractSocket) InheritSetSocketError(f func(socketError int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSocketError", f)
}

// void setLocalPort(quint16)
func (this *QAbstractSocket) InheritSetLocalPort(f func(port uint16) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLocalPort", f)
}

// void setLocalAddress(const class QHostAddress &)
func (this *QAbstractSocket) InheritSetLocalAddress(f func(address *QHostAddress) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLocalAddress", f)
}

// void setPeerPort(quint16)
func (this *QAbstractSocket) InheritSetPeerPort(f func(port uint16) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerPort", f)
}

// void setPeerAddress(const class QHostAddress &)
func (this *QAbstractSocket) InheritSetPeerAddress(f func(address *QHostAddress) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerAddress", f)
}

// void setPeerName(const class QString &)
func (this *QAbstractSocket) InheritSetPeerName(f func(name string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerName", f)
}

type QAbstractSocket struct {
	*qtcore.QIODevice
}
type QAbstractSocket_ITF interface {
	qtcore.QIODevice_ITF
	QAbstractSocket_PTR() *QAbstractSocket
}

func (ptr *QAbstractSocket) QAbstractSocket_PTR() *QAbstractSocket { return ptr }

func (this *QAbstractSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QAbstractSocket) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = qtcore.NewQIODeviceFromPointer(cthis)
}
func NewQAbstractSocketFromPointer(cthis unsafe.Pointer) *QAbstractSocket {
	bcthis0 := qtcore.NewQIODeviceFromPointer(cthis)
	return &QAbstractSocket{bcthis0}
}
func (*QAbstractSocket) NewFromPointer(cthis unsafe.Pointer) *QAbstractSocket {
	return NewQAbstractSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QAbstractSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSocket(enum QAbstractSocket::SocketType, QObject *)
func NewQAbstractSocket(socketType int, parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractSocket {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocketC2ENS_10SocketTypeEP7QObject", qtrt.FFI_TYPE_POINTER, socketType, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSocket()
func DeleteQAbstractSocket(this *QAbstractSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:142
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void resume()
func (this *QAbstractSocket) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::PauseModes pauseMode() const
func (this *QAbstractSocket) PauseMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket9pauseModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPauseMode(QAbstractSocket::PauseModes)
func (this *QAbstractSocket) SetPauseMode(pauseMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12setPauseModeE6QFlagsINS_9PauseModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pauseMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode int) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind__(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	// arg: 2, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind__1(address QHostAddress_ITF, port uint16) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind_1(port uint16, mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind_1_() bool {
	// arg: 0, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	// arg: 1, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)
func (this *QAbstractSocket) Bind_1_1(port uint16) bool {
	// arg: 1, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QAbstractSocket) ConnectToHost(hostName string, port uint16, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QAbstractSocket) ConnectToHost__(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QAbstractSocket) ConnectToHost__1(hostName string, port uint16, mode int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:152
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QHostAddress &, quint16, QIODevice::OpenMode)
func (this *QAbstractSocket) ConnectToHost_1(address QHostAddress_ITF, port uint16, mode int) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK12QHostAddresst6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:152
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QHostAddress &, quint16, QIODevice::OpenMode)
func (this *QAbstractSocket) ConnectToHost_1_(address QHostAddress_ITF, port uint16) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK12QHostAddresst6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void disconnectFromHost()
func (this *QAbstractSocket) DisconnectFromHost() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket18disconnectFromHostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QAbstractSocket) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const
func (this *QAbstractSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const
func (this *QAbstractSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:160
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const
func (this *QAbstractSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:162
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 localPort() const
func (this *QAbstractSocket) LocalPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket9localPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:163
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress localAddress() const
func (this *QAbstractSocket) LocalAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12localAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 peerPort() const
func (this *QAbstractSocket) PeerPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress peerAddress() const
func (this *QAbstractSocket) PeerAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket11peerAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerName() const
func (this *QAbstractSocket) PeerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize() const
func (this *QAbstractSocket) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:169
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)
func (this *QAbstractSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()
func (this *QAbstractSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:173
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qintptr socketDescriptor() const
func (this *QAbstractSocket) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:174
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QAbstractSocket::SocketState, QIODevice::OpenMode)
func (this *QAbstractSocket) SetSocketDescriptor(socketDescriptor int64, state int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:174
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QAbstractSocket::SocketState, QIODevice::OpenMode)
func (this *QAbstractSocket) SetSocketDescriptor__(socketDescriptor int64) bool {
	// arg: 1, QAbstractSocket::SocketState=Enum, QAbstractSocket::SocketState=Enum,
	state := 0
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:174
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QAbstractSocket::SocketState, QIODevice::OpenMode)
func (this *QAbstractSocket) SetSocketDescriptor__1(socketDescriptor int64, state int) bool {
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:177
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)
func (this *QAbstractSocket) SetSocketOption(option int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket15setSocketOptionENS_12SocketOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:178
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant socketOption(QAbstractSocket::SocketOption)
func (this *QAbstractSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12socketOptionENS_12SocketOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketType socketType() const
func (this *QAbstractSocket) SocketType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket10socketTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketState state() const
func (this *QAbstractSocket) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError error() const
func (this *QAbstractSocket) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:206
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QAbstractSocket::SocketError)
func (this *QAbstractSocket) Error_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5errorENS_11SocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QAbstractSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const
func (this *QAbstractSocket) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:187
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const
func (this *QAbstractSocket) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:188
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()
func (this *QAbstractSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:191
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QAbstractSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:191
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QAbstractSocket) WaitForConnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QAbstractSocket) WaitForReadyRead__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:193
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:193
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QAbstractSocket) WaitForBytesWritten__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QAbstractSocket) WaitForDisconnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxy(const QNetworkProxy &)
func (this *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if networkProxy != nil && networkProxy.QNetworkProxy_PTR() != nil {
		convArg0 = networkProxy.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket8setProxyERK13QNetworkProxy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy proxy() const
func (this *QAbstractSocket) Proxy() *QNetworkProxy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5proxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hostFound()
func (this *QAbstractSocket) HostFound() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9hostFoundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connected()
func (this *QAbstractSocket) Connected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9connectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnected()
func (this *QAbstractSocket) Disconnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12disconnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAbstractSocket::SocketState)
func (this *QAbstractSocket) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12stateChangedENS_11SocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void proxyAuthenticationRequired(const QNetworkProxy &, QAuthenticator *)
func (this *QAbstractSocket) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if proxy != nil && proxy.QNetworkProxy_PTR() != nil {
		convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QAuthenticator_PTR() != nil {
		convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket27proxyAuthenticationRequiredERK13QNetworkProxyP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QAbstractSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readLineData(char *, qint64)
func (this *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12readLineDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QAbstractSocket) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:216
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setSocketState(enum QAbstractSocket::SocketState)
func (this *QAbstractSocket) SetSocketState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketStateENS_11SocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:217
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setSocketError(enum QAbstractSocket::SocketError)
func (this *QAbstractSocket) SetSocketError(socketError int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketErrorENS_11SocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:218
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLocalPort(quint16)
func (this *QAbstractSocket) SetLocalPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12setLocalPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:219
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLocalAddress(const QHostAddress &)
func (this *QAbstractSocket) SetLocalAddress(address QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket15setLocalAddressERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:220
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerPort(quint16)
func (this *QAbstractSocket) SetPeerPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:221
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerAddress(const QHostAddress &)
func (this *QAbstractSocket) SetPeerAddress(address QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setPeerAddressERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:222
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerName(const QString &)
func (this *QAbstractSocket) SetPeerName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QAbstractSocket__SocketType = int

const QAbstractSocket__TcpSocket QAbstractSocket__SocketType = 0
const QAbstractSocket__UdpSocket QAbstractSocket__SocketType = 1
const QAbstractSocket__SctpSocket QAbstractSocket__SocketType = 2
const QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = -1

type QAbstractSocket__NetworkLayerProtocol = int

const QAbstractSocket__IPv4Protocol QAbstractSocket__NetworkLayerProtocol = 0
const QAbstractSocket__IPv6Protocol QAbstractSocket__NetworkLayerProtocol = 1
const QAbstractSocket__AnyIPProtocol QAbstractSocket__NetworkLayerProtocol = 2
const QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = -1

type QAbstractSocket__SocketError = int

const QAbstractSocket__ConnectionRefusedError QAbstractSocket__SocketError = 0
const QAbstractSocket__RemoteHostClosedError QAbstractSocket__SocketError = 1
const QAbstractSocket__HostNotFoundError QAbstractSocket__SocketError = 2
const QAbstractSocket__SocketAccessError QAbstractSocket__SocketError = 3
const QAbstractSocket__SocketResourceError QAbstractSocket__SocketError = 4
const QAbstractSocket__SocketTimeoutError QAbstractSocket__SocketError = 5
const QAbstractSocket__DatagramTooLargeError QAbstractSocket__SocketError = 6
const QAbstractSocket__NetworkError QAbstractSocket__SocketError = 7
const QAbstractSocket__AddressInUseError QAbstractSocket__SocketError = 8
const QAbstractSocket__SocketAddressNotAvailableError QAbstractSocket__SocketError = 9
const QAbstractSocket__UnsupportedSocketOperationError QAbstractSocket__SocketError = 10
const QAbstractSocket__UnfinishedSocketOperationError QAbstractSocket__SocketError = 11
const QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = 12
const QAbstractSocket__SslHandshakeFailedError QAbstractSocket__SocketError = 13
const QAbstractSocket__ProxyConnectionRefusedError QAbstractSocket__SocketError = 14
const QAbstractSocket__ProxyConnectionClosedError QAbstractSocket__SocketError = 15
const QAbstractSocket__ProxyConnectionTimeoutError QAbstractSocket__SocketError = 16
const QAbstractSocket__ProxyNotFoundError QAbstractSocket__SocketError = 17
const QAbstractSocket__ProxyProtocolError QAbstractSocket__SocketError = 18
const QAbstractSocket__OperationError QAbstractSocket__SocketError = 19
const QAbstractSocket__SslInternalError QAbstractSocket__SocketError = 20
const QAbstractSocket__SslInvalidUserDataError QAbstractSocket__SocketError = 21
const QAbstractSocket__TemporaryError QAbstractSocket__SocketError = 22
const QAbstractSocket__UnknownSocketError QAbstractSocket__SocketError = -1

type QAbstractSocket__SocketState = int

const QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = 0
const QAbstractSocket__HostLookupState QAbstractSocket__SocketState = 1
const QAbstractSocket__ConnectingState QAbstractSocket__SocketState = 2
const QAbstractSocket__ConnectedState QAbstractSocket__SocketState = 3
const QAbstractSocket__BoundState QAbstractSocket__SocketState = 4
const QAbstractSocket__ListeningState QAbstractSocket__SocketState = 5
const QAbstractSocket__ClosingState QAbstractSocket__SocketState = 6

type QAbstractSocket__SocketOption = int

const QAbstractSocket__LowDelayOption QAbstractSocket__SocketOption = 0
const QAbstractSocket__KeepAliveOption QAbstractSocket__SocketOption = 1
const QAbstractSocket__MulticastTtlOption QAbstractSocket__SocketOption = 2
const QAbstractSocket__MulticastLoopbackOption QAbstractSocket__SocketOption = 3
const QAbstractSocket__TypeOfServiceOption QAbstractSocket__SocketOption = 4
const QAbstractSocket__SendBufferSizeSocketOption QAbstractSocket__SocketOption = 5
const QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = 6

type QAbstractSocket__BindFlag = int

const QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = 0
const QAbstractSocket__ShareAddress QAbstractSocket__BindFlag = 1
const QAbstractSocket__DontShareAddress QAbstractSocket__BindFlag = 2
const QAbstractSocket__ReuseAddressHint QAbstractSocket__BindFlag = 4

type QAbstractSocket__PauseMode = int

const QAbstractSocket__PauseNever QAbstractSocket__PauseMode = 0
const QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = 1

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
