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

// void setSocketState(QAbstractSocket::SocketState)
func (this *QAbstractSocket) InheritSetSocketState(f func(state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSocketState", f)
}

// void setSocketError(QAbstractSocket::SocketError)
func (this *QAbstractSocket) InheritSetSocketError(f func(socketError int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSocketError", f)
}

// void setLocalPort(quint16)
func (this *QAbstractSocket) InheritSetLocalPort(f func(port uint16) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLocalPort", f)
}

// void setLocalAddress(const QHostAddress &)
func (this *QAbstractSocket) InheritSetLocalAddress(f func(address *QHostAddress) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLocalAddress", f)
}

// void setPeerPort(quint16)
func (this *QAbstractSocket) InheritSetPeerPort(f func(port uint16) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerPort", f)
}

// void setPeerAddress(const QHostAddress &)
func (this *QAbstractSocket) InheritSetPeerAddress(f func(address *QHostAddress) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerAddress", f)
}

// void setPeerName(const QString &)
func (this *QAbstractSocket) InheritSetPeerName(f func(name string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPeerName", f)
}

/*

 */
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

/*

 */
func (this *QAbstractSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSocket(QAbstractSocket::SocketType, QObject *)

/*
Creates a new abstract socket of type socketType. The parent argument is passed to QObject's constructor.

See also socketType(), QTcpSocket, and QUdpSocket.
*/
func (*QAbstractSocket) NewForInherit(socketType int, parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractSocket {
	return NewQAbstractSocket(socketType, parent)
}
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

// /usr/include/qt/QtNetwork/qabstractsocket.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSocket()

/*

 */
func DeleteQAbstractSocket(this *QAbstractSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:143
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void resume()

/*
Continues data transfer on the socket. This method should only be used after the socket has been set to pause upon notifications and a notification has been received. The only notification currently supported is QSslSocket::sslErrors(). Calling this method if the socket is not paused results in undefined behavior.

This function was introduced in  Qt 5.0.

See also pauseMode() and setPauseMode().
*/
func (this *QAbstractSocket) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::PauseModes pauseMode() const

/*
Returns the pause mode of this socket.

This function was introduced in  Qt 5.0.

See also setPauseMode() and resume().
*/
func (this *QAbstractSocket) PauseMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket9pauseModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPauseMode(QAbstractSocket::PauseModes)

/*
Controls whether to pause upon receiving a notification. The pauseMode parameter specifies the conditions in which the socket should be paused. The only notification currently supported is QSslSocket::sslErrors(). If set to PauseOnSslErrors, data transfer on the socket will be paused and needs to be enabled explicitly again by calling resume(). By default this option is set to PauseNever. This option must be called before connecting to the server, otherwise it will result in undefined behavior.

This function was introduced in  Qt 5.0.

See also pauseMode() and resume().
*/
func (this *QAbstractSocket) SetPauseMode(pauseMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12setPauseModeE6QFlagsINS_9PauseModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pauseMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode int) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bindp(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	// arg: 2, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bind(const QHostAddress &, quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bindp1(address QHostAddress_ITF, port uint16) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindERK12QHostAddresst6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:149
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bind1(port uint16, mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:149
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bind1p() bool {
	// arg: 0, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	// arg: 1, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:149
// index:1
// Public Visibility=Default Availability=Available
// [1] bool bind(quint16, QAbstractSocket::BindMode)

/*
Binds to address on port port, using the BindMode mode.

Binds this socket to the address address and the port port.

For UDP sockets, after binding, the signal QUdpSocket::readyRead() is emitted whenever a UDP datagram arrives on the specified address and port. Thus, This function is useful to write UDP servers.

For TCP sockets, this function may be used to specify which interface to use for an outgoing connection, which is useful in case of multiple network interfaces.

By default, the socket is bound using the DefaultForPlatform BindMode. If a port is not specified, a random port is chosen.

On success, the functions returns true and the socket enters BoundState; otherwise it returns false.

This function was introduced in  Qt 5.0.
*/
func (this *QAbstractSocket) Bind1p1(port uint16) bool {
	// arg: 1, QAbstractSocket::BindMode=Typedef, QAbstractSocket::BindMode=Typedef, QFlags<QAbstractSocket::BindFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket4bindEt6QFlagsINS_8BindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Attempts to make a connection to hostName on the given port. The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

The socket is opened in the given openMode and first enters HostLookupState, then performs a host name lookup of hostName. If the lookup succeeds, hostFound() is emitted and QAbstractSocket enters ConnectingState. It then attempts to connect to the address or addresses returned by the lookup. Finally, if a connection is established, QAbstractSocket enters ConnectedState and emits connected().

At any point, the socket can emit error() to signal that an error occurred.

hostName may be an IP address in string form (e.g., "43.195.83.32"), or it may be a host name (e.g., "example.com"). QAbstractSocket will do a lookup only if required. port is in native byte order.

See also state(), peerName(), peerAddress(), peerPort(), and waitForConnected().
*/
func (this *QAbstractSocket) ConnectToHost(hostName string, port uint16, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Attempts to make a connection to hostName on the given port. The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

The socket is opened in the given openMode and first enters HostLookupState, then performs a host name lookup of hostName. If the lookup succeeds, hostFound() is emitted and QAbstractSocket enters ConnectingState. It then attempts to connect to the address or addresses returned by the lookup. Finally, if a connection is established, QAbstractSocket enters ConnectedState and emits connected().

At any point, the socket can emit error() to signal that an error occurred.

hostName may be an IP address in string form (e.g., "43.195.83.32"), or it may be a host name (e.g., "example.com"). QAbstractSocket will do a lookup only if required. port is in native byte order.

See also state(), peerName(), peerAddress(), peerPort(), and waitForConnected().
*/
func (this *QAbstractSocket) ConnectToHostp(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	mode := 0
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum, , Invalid
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Attempts to make a connection to hostName on the given port. The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

The socket is opened in the given openMode and first enters HostLookupState, then performs a host name lookup of hostName. If the lookup succeeds, hostFound() is emitted and QAbstractSocket enters ConnectingState. It then attempts to connect to the address or addresses returned by the lookup. Finally, if a connection is established, QAbstractSocket enters ConnectedState and emits connected().

At any point, the socket can emit error() to signal that an error occurred.

hostName may be an IP address in string form (e.g., "43.195.83.32"), or it may be a host name (e.g., "example.com"). QAbstractSocket will do a lookup only if required. port is in native byte order.

See also state(), peerName(), peerAddress(), peerPort(), and waitForConnected().
*/
func (this *QAbstractSocket) ConnectToHostp1(hostName string, port uint16, mode int) {
	var tmpArg0 = qtcore.NewQString5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum, , Invalid
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEENS_20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:153
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QHostAddress &, quint16, QIODevice::OpenMode)

/*
Attempts to make a connection to hostName on the given port. The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

The socket is opened in the given openMode and first enters HostLookupState, then performs a host name lookup of hostName. If the lookup succeeds, hostFound() is emitted and QAbstractSocket enters ConnectingState. It then attempts to connect to the address or addresses returned by the lookup. Finally, if a connection is established, QAbstractSocket enters ConnectedState and emits connected().

At any point, the socket can emit error() to signal that an error occurred.

hostName may be an IP address in string form (e.g., "43.195.83.32"), or it may be a host name (e.g., "example.com"). QAbstractSocket will do a lookup only if required. port is in native byte order.

See also state(), peerName(), peerAddress(), peerPort(), and waitForConnected().
*/
func (this *QAbstractSocket) ConnectToHost1(address QHostAddress_ITF, port uint16, mode int) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK12QHostAddresst6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:153
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QHostAddress &, quint16, QIODevice::OpenMode)

/*
Attempts to make a connection to hostName on the given port. The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

The socket is opened in the given openMode and first enters HostLookupState, then performs a host name lookup of hostName. If the lookup succeeds, hostFound() is emitted and QAbstractSocket enters ConnectingState. It then attempts to connect to the address or addresses returned by the lookup. Finally, if a connection is established, QAbstractSocket enters ConnectedState and emits connected().

At any point, the socket can emit error() to signal that an error occurred.

hostName may be an IP address in string form (e.g., "43.195.83.32"), or it may be a host name (e.g., "example.com"). QAbstractSocket will do a lookup only if required. port is in native byte order.

See also state(), peerName(), peerAddress(), peerPort(), and waitForConnected().
*/
func (this *QAbstractSocket) ConnectToHost1p(address QHostAddress_ITF, port uint16) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket13connectToHostERK12QHostAddresst6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void disconnectFromHost()

/*
Attempts to close the socket. If there is pending data waiting to be written, QAbstractSocket will enter ClosingState and wait until all data has been written. Eventually, it will enter UnconnectedState and emit the disconnected() signal.

See also connectToHost().
*/
func (this *QAbstractSocket) DisconnectFromHost() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket18disconnectFromHostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:156
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the socket is valid and ready for use; otherwise returns false.

Note: The socket's state must be ConnectedState before reading and writing can occur.

See also state().
*/
func (this *QAbstractSocket) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const

/*
Reimplemented from QIODevice::bytesAvailable().

Returns the number of incoming bytes that are waiting to be read.

See also bytesToWrite() and read().
*/
func (this *QAbstractSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:159
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const

/*
Reimplemented from QIODevice::bytesToWrite().

Returns the number of bytes that are waiting to be written. The bytes are written when control goes back to the event loop or when flush() is called.

See also bytesAvailable() and flush().
*/
func (this *QAbstractSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const

/*
Reimplemented from QIODevice::canReadLine().

Returns true if a line of data can be read from the socket; otherwise returns false.

See also readLine().
*/
func (this *QAbstractSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:163
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 localPort() const

/*
Returns the host port number (in native byte order) of the local socket if available; otherwise returns 0.

See also localAddress(), peerPort(), and setLocalPort().
*/
func (this *QAbstractSocket) LocalPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket9localPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress localAddress() const

/*
Returns the host address of the local socket if available; otherwise returns QHostAddress::Null.

This is normally the main IP address of the host, but can be QHostAddress::LocalHost (127.0.0.1) for connections to the local host.

See also localPort(), peerAddress(), and setLocalAddress().
*/
func (this *QAbstractSocket) LocalAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12localAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:165
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 peerPort() const

/*
Returns the port of the connected peer if the socket is in ConnectedState; otherwise returns 0.

See also peerAddress(), localPort(), and setPeerPort().
*/
func (this *QAbstractSocket) PeerPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress peerAddress() const

/*
Returns the address of the connected peer if the socket is in ConnectedState; otherwise returns QHostAddress::Null.

See also peerName(), peerPort(), localAddress(), and setPeerAddress().
*/
func (this *QAbstractSocket) PeerAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket11peerAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerName() const

/*
Returns the name of the peer as specified by connectToHost(), or an empty QString if connectToHost() has not been called.

See also peerAddress(), peerPort(), and setPeerName().
*/
func (this *QAbstractSocket) PeerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize() const

/*
Returns the size of the internal read buffer. This limits the amount of data that the client can receive before you call read() or readAll().

A read buffer size of 0 (the default) means that the buffer has no size limit, ensuring that no data is lost.

See also setReadBufferSize() and read().
*/
func (this *QAbstractSocket) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:170
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)

/*
Sets the size of QAbstractSocket's internal read buffer to be size bytes.

If the buffer size is limited to a certain size, QAbstractSocket won't buffer more than this size of data. Exceptionally, a buffer size of 0 means that the read buffer is unlimited and all incoming data is buffered. This is the default.

This option is useful if you only read the data at certain points in time (e.g., in a real-time streaming application) or if you want to protect your socket against receiving too much data, which may eventually cause your application to run out of memory.

Only QTcpSocket uses QAbstractSocket's internal buffer; QUdpSocket does not use any buffering at all, but rather relies on the implicit buffering provided by the operating system. Because of this, calling this function on QUdpSocket has no effect.

See also readBufferSize() and read().
*/
func (this *QAbstractSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()

/*
Aborts the current connection and resets the socket. Unlike disconnectFromHost(), this function immediately closes the socket, discarding any pending data in the write buffer.

See also disconnectFromHost() and close().
*/
func (this *QAbstractSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:174
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qintptr socketDescriptor() const

/*
Returns the native socket descriptor of the QAbstractSocket object if this is available; otherwise returns -1.

If the socket is using QNetworkProxy, the returned descriptor may not be usable with native socket functions.

The socket descriptor is not available when QAbstractSocket is in UnconnectedState.

See also setSocketDescriptor().
*/
func (this *QAbstractSocket) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Initializes QAbstractSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState. Read and write buffers are cleared, discarding any pending data.

Note: It is not possible to initialize two abstract sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QAbstractSocket) SetSocketDescriptor(socketDescriptor int64, state int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Initializes QAbstractSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState. Read and write buffers are cleared, discarding any pending data.

Note: It is not possible to initialize two abstract sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QAbstractSocket) SetSocketDescriptorp(socketDescriptor int64) bool {
	// arg: 1, QAbstractSocket::SocketState=Enum, QAbstractSocket::SocketState=Enum, , Invalid
	state := 0
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Initializes QAbstractSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState. Read and write buffers are cleared, discarding any pending data.

Note: It is not possible to initialize two abstract sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QAbstractSocket) SetSocketDescriptorp1(socketDescriptor int64, state int) bool {
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19setSocketDescriptorExNS_11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:178
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)

/*
Sets the given option to the value described by value.

Note: On Windows Runtime, QAbstractSocket::KeepAliveOption must be set before the socket is connected.

This function was introduced in  Qt 4.6.

See also socketOption().
*/
func (this *QAbstractSocket) SetSocketOption(option int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket15setSocketOptionENS_12SocketOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:179
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant socketOption(QAbstractSocket::SocketOption)

/*
Returns the value of the option option.

This function was introduced in  Qt 4.6.

See also setSocketOption().
*/
func (this *QAbstractSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12socketOptionENS_12SocketOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketType socketType() const

/*
Returns the socket type (TCP, UDP, or other).

See also QTcpSocket and QUdpSocket.
*/
func (this *QAbstractSocket) SocketType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket10socketTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketState state() const

/*
Returns the state of the socket.

See also error().
*/
func (this *QAbstractSocket) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:183
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError error() const

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QAbstractSocket) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:207
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QAbstractSocket::SocketError)

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QAbstractSocket) Error1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5errorENS_11SocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().

Closes the I/O device for the socket and calls disconnectFromHost() to close the socket's connection.

See QIODevice::close() for a description of the actions that occur when an I/O device is closed.

See also abort().
*/
func (this *QAbstractSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:187
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const

/*
Reimplemented from QIODevice::isSequential().
*/
func (this *QAbstractSocket) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:188
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Reimplemented from QIODevice::atEnd().

Returns true if no more data is currently available for reading; otherwise returns false.

This function is most commonly used when reading data from the socket in a loop. For example:


   // This slot is connected to QAbstractSocket::readyRead()
   void SocketClass::readyReadSlot()
   {
       while (!socket.atEnd()) {
           QByteArray data = socket.read(100);
           ....
       }
   }



See also bytesAvailable() and readyRead().
*/
func (this *QAbstractSocket) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:189
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()

/*
This function writes as much as possible from the internal write buffer to the underlying network socket, without blocking. If any data was written, this function returns true; otherwise false is returned.

Call this function if you need QAbstractSocket to start sending buffered data immediately. The number of bytes successfully written depends on the operating system. In most cases, you do not need to call this function, because QAbstractSocket will start sending data automatically once control goes back to the event loop. In the absence of an event loop, call waitForBytesWritten() instead.

See also write() and waitForBytesWritten().
*/
func (this *QAbstractSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Waits until the socket is connected, up to msecs milliseconds. If the connection has been established, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be established:


  socket->connectToHost("imap", 143);
  if (socket->waitForConnected(1000))
      qDebug("Connected!");



If msecs is -1, this function will not time out.

Note: This function may wait slightly longer than msecs, depending on the time it takes to complete the host lookup.

Note: Multiple calls to this functions do not accumulate the time. If the function times out, the connecting process will be aborted.

Note: This function may fail randomly on Windows. Consider using the event loop and the connected() signal if your software will run on Windows.

See also connectToHost() and connected().
*/
func (this *QAbstractSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Waits until the socket is connected, up to msecs milliseconds. If the connection has been established, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be established:


  socket->connectToHost("imap", 143);
  if (socket->waitForConnected(1000))
      qDebug("Connected!");



If msecs is -1, this function will not time out.

Note: This function may wait slightly longer than msecs, depending on the time it takes to complete the host lookup.

Note: Multiple calls to this functions do not accumulate the time. If the function times out, the connecting process will be aborted.

Note: This function may fail randomly on Windows. Consider using the event loop and the connected() signal if your software will run on Windows.

See also connectToHost() and connected().
*/
func (this *QAbstractSocket) WaitForConnectedp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:193
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().

This function blocks until new data is available for reading and the readyRead() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if the readyRead() signal is emitted and there is new data available for reading; otherwise it returns false (if an error occurred or the operation timed out).

Note: This function may fail randomly on Windows. Consider using the event loop and the readyRead() signal if your software will run on Windows.

See also waitForBytesWritten().
*/
func (this *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:193
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().

This function blocks until new data is available for reading and the readyRead() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if the readyRead() signal is emitted and there is new data available for reading; otherwise it returns false (if an error occurred or the operation timed out).

Note: This function may fail randomly on Windows. Consider using the event loop and the readyRead() signal if your software will run on Windows.

See also waitForBytesWritten().
*/
func (this *QAbstractSocket) WaitForReadyReadp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().

This function blocks until at least one byte has been written on the socket and the bytesWritten() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if the bytesWritten() signal is emitted; otherwise it returns false (if an error occurred or the operation timed out).

Note: This function may fail randomly on Windows. Consider using the event loop and the bytesWritten() signal if your software will run on Windows.

See also waitForReadyRead().
*/
func (this *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().

This function blocks until at least one byte has been written on the socket and the bytesWritten() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if the bytesWritten() signal is emitted; otherwise it returns false (if an error occurred or the operation timed out).

Note: This function may fail randomly on Windows. Consider using the event loop and the bytesWritten() signal if your software will run on Windows.

See also waitForReadyRead().
*/
func (this *QAbstractSocket) WaitForBytesWrittenp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:195
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Waits until the socket has disconnected, up to msecs milliseconds. If the connection has been disconnected, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be closed:


  socket->disconnectFromHost();
      if (socket->state() == QAbstractSocket::UnconnectedState ||
          socket->waitForDisconnected(1000))
          qDebug("Disconnected!");



If msecs is -1, this function will not time out.

Note: This function may fail randomly on Windows. Consider using the event loop and the disconnected() signal if your software will run on Windows.

See also disconnectFromHost() and close().
*/
func (this *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:195
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Waits until the socket has disconnected, up to msecs milliseconds. If the connection has been disconnected, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be closed:


  socket->disconnectFromHost();
      if (socket->state() == QAbstractSocket::UnconnectedState ||
          socket->waitForDisconnected(1000))
          qDebug("Disconnected!");



If msecs is -1, this function will not time out.

Note: This function may fail randomly on Windows. Consider using the event loop and the disconnected() signal if your software will run on Windows.

See also disconnectFromHost() and close().
*/
func (this *QAbstractSocket) WaitForDisconnectedp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxy(const QNetworkProxy &)

/*
Sets the explicit network proxy for this socket to networkProxy.

To disable the use of a proxy for this socket, use the QNetworkProxy::NoProxy proxy type:


  socket->setProxy(QNetworkProxy::NoProxy);



The default value for the proxy is QNetworkProxy::DefaultProxy, which means the socket will use the application settings: if a proxy is set with QNetworkProxy::setApplicationProxy, it will use that; otherwise, if a factory is set with QNetworkProxyFactory::setApplicationProxyFactory, it will query that factory with type QNetworkProxyQuery::TcpSocket.

This function was introduced in  Qt 4.1.

See also proxy(), QNetworkProxy, and QNetworkProxyFactory::queryProxy().
*/
func (this *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if networkProxy != nil && networkProxy.QNetworkProxy_PTR() != nil {
		convArg0 = networkProxy.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket8setProxyERK13QNetworkProxy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy proxy() const

/*
Returns the network proxy for this socket. By default QNetworkProxy::DefaultProxy is used, which means this socket will query the default proxy settings for the application.

This function was introduced in  Qt 4.1.

See also setProxy(), QNetworkProxy, and QNetworkProxyFactory.
*/
func (this *QAbstractSocket) Proxy() *QNetworkProxy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractSocket5proxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hostFound()

/*
This signal is emitted after connectToHost() has been called and the host lookup has succeeded.

Note: Since Qt 4.6.3 QAbstractSocket may emit hostFound() directly from the connectToHost() call since a DNS result could have been cached.

See also connected().
*/
func (this *QAbstractSocket) HostFound() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9hostFoundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connected()

/*
This signal is emitted after connectToHost() has been called and a connection has been successfully established.

Note: On some operating systems the connected() signal may be directly emitted from the connectToHost() call for connections to the localhost.

See also connectToHost() and disconnected().
*/
func (this *QAbstractSocket) Connected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9connectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnected()

/*
This signal is emitted when the socket has been disconnected.

Warning: If you need to delete the sender() of this signal in a slot connected to it, use the deleteLater() function.

See also connectToHost(), disconnectFromHost(), and abort().
*/
func (this *QAbstractSocket) Disconnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12disconnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAbstractSocket::SocketState)

/*
This signal is emitted whenever QAbstractSocket's state changes. The socketState parameter is the new state.

QAbstractSocket::SocketState is not a registered metatype, so for queued connections, you will have to register it with Q_DECLARE_METATYPE() and qRegisterMetaType().

See also state() and Creating Custom Qt Types.
*/
func (this *QAbstractSocket) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12stateChangedENS_11SocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void proxyAuthenticationRequired(const QNetworkProxy &, QAuthenticator *)

/*
This signal can be emitted when a proxy that requires authentication is used. The authenticator object can then be filled in with the required details to allow authentication and continue the connection.

Note: It is not possible to use a QueuedConnection to connect to this signal, as the connection will fail if the authenticator has not been filled in with new information when the signal returns.

This function was introduced in  Qt 4.3.

See also QAuthenticator and QNetworkProxy.
*/
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

// /usr/include/qt/QtNetwork/qabstractsocket.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QAbstractSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readLineData(char *, qint64)

/*
Reimplemented from QIODevice::readLineData().
*/
func (this *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12readLineDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QAbstractSocket) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:217
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setSocketState(QAbstractSocket::SocketState)

/*
Sets the state of the socket to state.

See also state().
*/
func (this *QAbstractSocket) SetSocketState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketStateENS_11SocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:218
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setSocketError(QAbstractSocket::SocketError)

/*
Sets the type of error that last occurred to socketError.

See also setSocketState() and setErrorString().
*/
func (this *QAbstractSocket) SetSocketError(socketError int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketErrorENS_11SocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:219
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLocalPort(quint16)

/*
Sets the port on the local side of a connection to port.

You can call this function in a subclass of QAbstractSocket to change the return value of the localPort() function after a connection has been established. This feature is commonly used by proxy connections for virtual connection settings.

Note that this function does not bind the local port of the socket prior to a connection (e.g., QAbstractSocket::bind()).

This function was introduced in  Qt 4.1.

See also localPort(), localAddress(), setLocalAddress(), and setPeerPort().
*/
func (this *QAbstractSocket) SetLocalPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket12setLocalPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:220
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLocalAddress(const QHostAddress &)

/*
Sets the address on the local side of a connection to address.

You can call this function in a subclass of QAbstractSocket to change the return value of the localAddress() function after a connection has been established. This feature is commonly used by proxy connections for virtual connection settings.

Note that this function does not bind the local address of the socket prior to a connection (e.g., QAbstractSocket::bind()).

This function was introduced in  Qt 4.1.

See also localAddress(), setLocalPort(), and setPeerAddress().
*/
func (this *QAbstractSocket) SetLocalAddress(address QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket15setLocalAddressERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:221
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerPort(quint16)

/*
Sets the port of the remote side of the connection to port.

You can call this function in a subclass of QAbstractSocket to change the return value of the peerPort() function after a connection has been established. This feature is commonly used by proxy connections for virtual connection settings.

This function was introduced in  Qt 4.1.

See also peerPort(), setPeerAddress(), and setLocalPort().
*/
func (this *QAbstractSocket) SetPeerPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:222
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerAddress(const QHostAddress &)

/*
Sets the address of the remote side of the connection to address.

You can call this function in a subclass of QAbstractSocket to change the return value of the peerAddress() function after a connection has been established. This feature is commonly used by proxy connections for virtual connection settings.

This function was introduced in  Qt 4.1.

See also peerAddress(), setPeerPort(), and setLocalAddress().
*/
func (this *QAbstractSocket) SetPeerAddress(address QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket14setPeerAddressERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:223
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPeerName(const QString &)

/*
Sets the host name of the remote peer to name.

You can call this function in a subclass of QAbstractSocket to change the return value of the peerName() function after a connection has been established. This feature is commonly used by proxy connections for virtual connection settings.

This function was introduced in  Qt 4.1.

See also peerName().
*/
func (this *QAbstractSocket) SetPeerName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the transport layer protocol.



See also QAbstractSocket::socketType().

*/
type QAbstractSocket__SocketType = int

// TCP
const QAbstractSocket__TcpSocket QAbstractSocket__SocketType = 0

// UDP
const QAbstractSocket__UdpSocket QAbstractSocket__SocketType = 1

// SCTP
const QAbstractSocket__SctpSocket QAbstractSocket__SocketType = 2

//
const QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = -1

func (this *QAbstractSocket) SocketTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_SocketTypeItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.SocketTypeItemName(val)
}

/*
This enum describes the network layer protocol values used in Qt.



See also QHostAddress::protocol().

*/
type QAbstractSocket__NetworkLayerProtocol = int

// IPv4
const QAbstractSocket__IPv4Protocol QAbstractSocket__NetworkLayerProtocol = 0

// IPv6
const QAbstractSocket__IPv6Protocol QAbstractSocket__NetworkLayerProtocol = 1

//
const QAbstractSocket__AnyIPProtocol QAbstractSocket__NetworkLayerProtocol = 2

//
const QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = -1

func (this *QAbstractSocket) NetworkLayerProtocolItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_NetworkLayerProtocolItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.NetworkLayerProtocolItemName(val)
}

/*
This enum describes the socket errors that can occur.



See also QAbstractSocket::error().

*/
type QAbstractSocket__SocketError = int

// The connection was refused by the peer (or timed out).
const QAbstractSocket__ConnectionRefusedError QAbstractSocket__SocketError = 0

// The remote host closed the connection. Note that the client socket (i.e., this socket) will be closed after the remote close notification has been sent.
const QAbstractSocket__RemoteHostClosedError QAbstractSocket__SocketError = 1

// The host address was not found.
const QAbstractSocket__HostNotFoundError QAbstractSocket__SocketError = 2

// The socket operation failed because the application lacked the required privileges.
const QAbstractSocket__SocketAccessError QAbstractSocket__SocketError = 3

// The local system ran out of resources (e.g., too many sockets).
const QAbstractSocket__SocketResourceError QAbstractSocket__SocketError = 4

// The socket operation timed out.
const QAbstractSocket__SocketTimeoutError QAbstractSocket__SocketError = 5

//
const QAbstractSocket__DatagramTooLargeError QAbstractSocket__SocketError = 6

// An error occurred with the network (e.g., the network cable was accidentally plugged out).
const QAbstractSocket__NetworkError QAbstractSocket__SocketError = 7

// The address specified to QAbstractSocket::bind() is already in use and was set to be exclusive.
const QAbstractSocket__AddressInUseError QAbstractSocket__SocketError = 8

// The address specified to QAbstractSocket::bind() does not belong to the host.
const QAbstractSocket__SocketAddressNotAvailableError QAbstractSocket__SocketError = 9

//
const QAbstractSocket__UnsupportedSocketOperationError QAbstractSocket__SocketError = 10

//
const QAbstractSocket__UnfinishedSocketOperationError QAbstractSocket__SocketError = 11

//
const QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = 12

//
const QAbstractSocket__SslHandshakeFailedError QAbstractSocket__SocketError = 13

//
const QAbstractSocket__ProxyConnectionRefusedError QAbstractSocket__SocketError = 14

//
const QAbstractSocket__ProxyConnectionClosedError QAbstractSocket__SocketError = 15

//
const QAbstractSocket__ProxyConnectionTimeoutError QAbstractSocket__SocketError = 16

//
const QAbstractSocket__ProxyNotFoundError QAbstractSocket__SocketError = 17

//
const QAbstractSocket__ProxyProtocolError QAbstractSocket__SocketError = 18

//
const QAbstractSocket__OperationError QAbstractSocket__SocketError = 19

//
const QAbstractSocket__SslInternalError QAbstractSocket__SocketError = 20

//
const QAbstractSocket__SslInvalidUserDataError QAbstractSocket__SocketError = 21

//
const QAbstractSocket__TemporaryError QAbstractSocket__SocketError = 22

//
const QAbstractSocket__UnknownSocketError QAbstractSocket__SocketError = -1

func (this *QAbstractSocket) SocketErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_SocketErrorItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.SocketErrorItemName(val)
}

/*
This enum describes the different states in which a socket can be.



See also QAbstractSocket::state().

*/
type QAbstractSocket__SocketState = int

// The socket is not connected.
const QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = 0

// The socket is performing a host name lookup.
const QAbstractSocket__HostLookupState QAbstractSocket__SocketState = 1

// The socket has started establishing a connection.
const QAbstractSocket__ConnectingState QAbstractSocket__SocketState = 2

// A connection is established.
const QAbstractSocket__ConnectedState QAbstractSocket__SocketState = 3

// The socket is bound to an address and port.
const QAbstractSocket__BoundState QAbstractSocket__SocketState = 4

// For internal use only.
const QAbstractSocket__ListeningState QAbstractSocket__SocketState = 5

// The socket is about to close (data may still be waiting to be written).
const QAbstractSocket__ClosingState QAbstractSocket__SocketState = 6

func (this *QAbstractSocket) SocketStateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_SocketStateItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.SocketStateItemName(val)
}

/*
This enum represents the options that can be set on a socket. If desired, they can be set after having received the connected() signal from the socket or after having received a new socket from a QTcpServer.



Possible values for TypeOfServiceOption are:


 ValueDescription
224Network control
192Internetwork control
160CRITIC/ECP
128Flash override
96Flash
64Immediate
32Priority
0Routine


This enum was introduced or modified in  Qt 4.6.

See also QAbstractSocket::setSocketOption() and QAbstractSocket::socketOption().

*/
type QAbstractSocket__SocketOption = int

//
const QAbstractSocket__LowDelayOption QAbstractSocket__SocketOption = 0

//
const QAbstractSocket__KeepAliveOption QAbstractSocket__SocketOption = 1

// Set this to an integer value to set IP_MULTICAST_TTL (TTL for multicast datagrams) socket option.
const QAbstractSocket__MulticastTtlOption QAbstractSocket__SocketOption = 2

//
const QAbstractSocket__MulticastLoopbackOption QAbstractSocket__SocketOption = 3

// This option is not supported on Windows. This maps to the IP_TOS socket option. For possible values, see table below.
const QAbstractSocket__TypeOfServiceOption QAbstractSocket__SocketOption = 4

//
const QAbstractSocket__SendBufferSizeSocketOption QAbstractSocket__SocketOption = 5

//
const QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = 6

//
const QAbstractSocket__PathMtuSocketOption QAbstractSocket__SocketOption = 7

func (this *QAbstractSocket) SocketOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_SocketOptionItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.SocketOptionItemName(val)
}

/*


 */
type QAbstractSocket__BindFlag = int

//
const QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = 0

//
const QAbstractSocket__ShareAddress QAbstractSocket__BindFlag = 1

//
const QAbstractSocket__DontShareAddress QAbstractSocket__BindFlag = 2

//
const QAbstractSocket__ReuseAddressHint QAbstractSocket__BindFlag = 4

func (this *QAbstractSocket) BindFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_BindFlagItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.BindFlagItemName(val)
}

/*


 */
type QAbstractSocket__PauseMode = int

//
const QAbstractSocket__PauseNever QAbstractSocket__PauseMode = 0

//
const QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = 1

func (this *QAbstractSocket) PauseModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSocket_PauseModeItemName(val int) string {
	var nilthis *QAbstractSocket
	return nilthis.PauseModeItemName(val)
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
