package qtnetwork

// /usr/include/qt/QtNetwork/qsslsocket.h
// #include <qsslsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
func (this *QSslSocket) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QSslSocket) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

/*

 */
type QSslSocket struct {
	*QTcpSocket
}
type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func (ptr *QSslSocket) QSslSocket_PTR() *QSslSocket { return ptr }

func (this *QSslSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTcpSocket.GetCthis()
	}
}
func (this *QSslSocket) SetCthis(cthis unsafe.Pointer) {
	this.QTcpSocket = NewQTcpSocketFromPointer(cthis)
}
func NewQSslSocketFromPointer(cthis unsafe.Pointer) *QSslSocket {
	bcthis0 := NewQTcpSocketFromPointer(cthis)
	return &QSslSocket{bcthis0}
}
func (*QSslSocket) NewFromPointer(cthis unsafe.Pointer) *QSslSocket {
	return NewQSslSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSslSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qsslsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslSocket(QObject *)

/*
Constructs a QSslSocket object. parent is passed to QObject's constructor. The new socket's cipher suite is set to the one returned by the static method defaultCiphers().
*/
func NewQSslSocket(parent qtcore.QObject_ITF /*777 QObject **/) *QSslSocket {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSslSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qsslsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslSocket(QObject *)

/*
Constructs a QSslSocket object. parent is passed to QObject's constructor. The new socket's cipher suite is set to the one returned by the static method defaultCiphers().
*/
func NewQSslSocket__() *QSslSocket {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSslSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qsslsocket.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSslSocket()

/*

 */
func DeleteQSslSocket(this *QSslSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void resume()

/*
Reimplemented from QAbstractSocket::resume().

Continues data transfer on the socket after it has been paused. If "setPauseMode(QAbstractSocket::PauseOnSslErrors);" has been called on this socket and a sslErrors() signal is received, calling this method is necessary for the socket to continue.

This function was introduced in  Qt 5.0.

See also QAbstractSocket::pauseMode() and QAbstractSocket::setPauseMode().
*/
func (this *QSslSocket) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted__(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted__1(hostName string, port uint16, mode int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:88
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QString &, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted_1(hostName string, port uint16, sslPeerName string, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(sslPeerName)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringtS2_6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:88
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QString &, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted_1_(hostName string, port uint16, sslPeerName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(sslPeerName)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	// arg: 4, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringtS2_6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:88
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QString &, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*
Starts an encrypted connection to the device hostName on port, using mode as the OpenMode. This is equivalent to calling connectToHost() to establish the connection, followed by a call to startClientEncryption(). The protocol parameter can be used to specify which network protocol to use (eg. IPv4 or IPv6).

QSslSocket first enters the HostLookupState. Then, after entering either the event loop or one of the waitFor...() functions, it enters the ConnectingState, emits connected(), and then initiates the SSL client handshake. At each state change, QSslSocket emits signal stateChanged().

After initiating the SSL client handshake, if the identity of the peer can't be established, signal sslErrors() is emitted. If you want to ignore the errors and continue connecting, you must call ignoreSslErrors(), either from inside a slot function connected to the sslErrors() signal, or prior to entering encrypted mode. If ignoreSslErrors() is not called, the connection is dropped, signal disconnected() is emitted, and QSslSocket returns to the UnconnectedState.

If the SSL handshake is successful, QSslSocket emits encrypted().


  QSslSocket socket;
  connect(&socket, SIGNAL(encrypted()), receiver, SLOT(socketEncrypted()));

  socket.connectToHostEncrypted("imap", 993);
  socket->write("1 CAPABILITY\r\n");



Note: The example above shows that text can be written to the socket immediately after requesting the encrypted connection, before the encrypted() signal has been emitted. In such cases, the text is queued in the object and written to the socket after the connection is established and the encrypted() signal has been emitted.

The default for mode is ReadWrite.

If you want to create a QSslSocket on the server side of a connection, you should instead call startServerEncryption() upon receiving the incoming connection through QTcpServer.

See also connectToHost(), startClientEncryption(), waitForConnected(), and waitForEncrypted().
*/
func (this *QSslSocket) ConnectToHostEncrypted_1_1(hostName string, port uint16, sslPeerName string, mode int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(sslPeerName)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringtS2_6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Reimplemented from QAbstractSocket::setSocketDescriptor().

Initializes QSslSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by state.

Note: It is not possible to initialize two sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QSslSocket) SetSocketDescriptor(socketDescriptor int64, state int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSocketDescriptorExN15QAbstractSocket11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Reimplemented from QAbstractSocket::setSocketDescriptor().

Initializes QSslSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by state.

Note: It is not possible to initialize two sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QSslSocket) SetSocketDescriptor__(socketDescriptor int64) bool {
	// arg: 1, QAbstractSocket::SocketState=Enum, QAbstractSocket::SocketState=Enum,
	state := 0
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSocketDescriptorExN15QAbstractSocket11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QAbstractSocket::SocketState, QIODevice::OpenMode)

/*
Reimplemented from QAbstractSocket::setSocketDescriptor().

Initializes QSslSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by state.

Note: It is not possible to initialize two sockets with the same native socket descriptor.

See also socketDescriptor().
*/
func (this *QSslSocket) SetSocketDescriptor__1(socketDescriptor int64, state int) bool {
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSocketDescriptorExN15QAbstractSocket11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*

 */
func (this *QSslSocket) ConnectToHost(hostName string, port uint16, openMode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, openMode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*

 */
func (this *QSslSocket) ConnectToHost__(hostName string, port uint16) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, openMode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, QAbstractSocket::NetworkLayerProtocol)

/*

 */
func (this *QSslSocket) ConnectToHost__1(hostName string, port uint16, openMode int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, QAbstractSocket::NetworkLayerProtocol=Enum, QAbstractSocket::NetworkLayerProtocol=Enum,
	protocol := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, openMode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void disconnectFromHost()

/*

 */
func (this *QSslSocket) DisconnectFromHost() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18disconnectFromHostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)

/*
Reimplemented from QAbstractSocket::setSocketOption().

Sets the given option to the value described by value.

This function was introduced in  Qt 4.6.

See also socketOption().
*/
func (this *QSslSocket) SetSocketOption(option int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15setSocketOptionEN15QAbstractSocket12SocketOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant socketOption(QAbstractSocket::SocketOption)

/*
Reimplemented from QAbstractSocket::socketOption().

Returns the value of the option option.

This function was introduced in  Qt 4.6.

See also setSocketOption().
*/
func (this *QSslSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket12socketOptionEN15QAbstractSocket12SocketOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::SslMode mode() const

/*
Returns the current mode for the socket; either UnencryptedMode, where QSslSocket behaves identially to QTcpSocket, or one of SslClientMode or SslServerMode, where the client is either negotiating or in encrypted mode.

When the mode changes, QSslSocket emits modeChanged()

See also SslMode.
*/
func (this *QSslSocket) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEncrypted() const

/*
Returns true if the socket is encrypted; otherwise, false is returned.

An encrypted socket encrypts all data that is written by calling write() or putChar() before the data is written to the network, and decrypts all incoming data as the data is received from the network, before you call read(), readLine() or getChar().

QSslSocket emits encrypted() when it enters encrypted mode.

You can call sessionCipher() to find which cryptographic cipher is used to encrypt and decrypt your data.

See also mode().
*/
func (this *QSslSocket) IsEncrypted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11isEncryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol() const

/*
Returns the socket's SSL protocol. By default, QSsl::SecureProtocols is used.

See also setProtocol().
*/
func (this *QSslSocket) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocol(QSsl::SslProtocol)

/*
Sets the socket's SSL protocol to protocol. This will affect the next initiated handshake; calling this function on an already-encrypted socket will not affect the socket's protocol.

See also protocol().
*/
func (this *QSslSocket) SetProtocol(protocol int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11setProtocolEN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::PeerVerifyMode peerVerifyMode() const

/*
Returns the socket's verify mode. This mode decides whether QSslSocket should request a certificate from the peer (i.e., the client requests a certificate from the server, or a server requesting a certificate from the client), and whether it should require that this certificate is valid.

The default mode is AutoVerifyPeer, which tells QSslSocket to use VerifyPeer for clients and QueryPeer for servers.

This function was introduced in  Qt 4.4.

See also setPeerVerifyMode(), peerVerifyDepth(), and mode().
*/
func (this *QSslSocket) PeerVerifyMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyMode(QSslSocket::PeerVerifyMode)

/*
Sets the socket's verify mode to mode. This mode decides whether QSslSocket should request a certificate from the peer (i.e., the client requests a certificate from the server, or a server requesting a certificate from the client), and whether it should require that this certificate is valid.

The default mode is AutoVerifyPeer, which tells QSslSocket to use VerifyPeer for clients and QueryPeer for servers.

Setting this mode after encryption has started has no effect on the current connection.

This function was introduced in  Qt 4.4.

See also peerVerifyMode(), setPeerVerifyDepth(), and mode().
*/
func (this *QSslSocket) SetPeerVerifyMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyModeENS_14PeerVerifyModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerVerifyDepth() const

/*
Returns the maximum number of certificates in the peer's certificate chain to be checked during the SSL handshake phase, or 0 (the default) if no maximum depth has been set, indicating that the whole certificate chain should be checked.

The certificates are checked in issuing order, starting with the peer's own certificate, then its issuer's certificate, and so on.

This function was introduced in  Qt 4.4.

See also setPeerVerifyDepth() and peerVerifyMode().
*/
func (this *QSslSocket) PeerVerifyDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15peerVerifyDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslsocket.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyDepth(int)

/*
Sets the maximum number of certificates in the peer's certificate chain to be checked during the SSL handshake phase, to depth. Setting a depth of 0 means that no maximum depth is set, indicating that the whole certificate chain should be checked.

The certificates are checked in issuing order, starting with the peer's own certificate, then its issuer's certificate, and so on.

This function was introduced in  Qt 4.4.

See also peerVerifyDepth() and setPeerVerifyMode().
*/
func (this *QSslSocket) SetPeerVerifyDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18setPeerVerifyDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerVerifyName() const

/*
Returns the different hostname for the certificate validation, as set by setPeerVerifyName or by connectToHostEncrypted.

This function was introduced in  Qt 4.8.

See also setPeerVerifyName() and connectToHostEncrypted().
*/
func (this *QSslSocket) PeerVerifyName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslsocket.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyName(const QString &)

/*
Sets a different host name, given by hostName, for the certificate validation instead of the one used for the TCP connection.

This function was introduced in  Qt 4.8.

See also peerVerifyName() and connectToHostEncrypted().
*/
func (this *QSslSocket) SetPeerVerifyName(hostName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const

/*
Reimplemented from QIODevice::bytesAvailable().

Returns the number of decrypted bytes that are immediately available for reading.
*/
func (this *QSslSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const

/*
Reimplemented from QIODevice::bytesToWrite().

Returns the number of unencrypted bytes that are waiting to be encrypted and written to the network.
*/
func (this *QSslSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const

/*
Reimplemented from QIODevice::canReadLine().

Returns true if you can read one while line (terminated by a single ASCII '\n' character) of decrypted characters; otherwise, false is returned.
*/
func (this *QSslSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().
*/
func (this *QSslSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Reimplemented from QIODevice::atEnd().
*/
func (this *QSslSocket) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()

/*
This function writes as much as possible from the internal write buffer to the underlying network socket, without blocking. If any data was written, this function returns true; otherwise false is returned.

Call this function if you need QSslSocket to start sending buffered data immediately. The number of bytes successfully written depends on the operating system. In most cases, you do not need to call this function, because QAbstractSocket will start sending data automatically once control goes back to the event loop. In the absence of an event loop, call waitForBytesWritten() instead.

See also write() and waitForBytesWritten().
*/
func (this *QSslSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()

/*
Aborts the current connection and resets the socket. Unlike disconnectFromHost(), this function immediately closes the socket, clearing any pending data in the write buffer.

See also disconnectFromHost() and close().
*/
func (this *QSslSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)

/*
Reimplemented from QAbstractSocket::setReadBufferSize().

Sets the size of QSslSocket's internal read buffer to be size bytes.

This function was introduced in  Qt 4.4.
*/
func (this *QSslSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesAvailable() const

/*
Returns the number of encrypted bytes that are awaiting decryption. Normally, this function will return 0 because QSslSocket decrypts its incoming data as soon as it can.

This function was introduced in  Qt 4.4.
*/
func (this *QSslSocket) EncryptedBytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket23encryptedBytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesToWrite() const

/*
Returns the number of encrypted bytes that are waiting to be written to the network.

This function was introduced in  Qt 4.4.
*/
func (this *QSslSocket) EncryptedBytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket21encryptedBytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration() const

/*
Returns the socket's SSL configuration state. The default SSL configuration of a socket is to use the default ciphers, default CA certificates, no local private key or certificate.

The SSL configuration also contains fields that can change with time without notice.

This function was introduced in  Qt 4.4.

See also setSslConfiguration(), localCertificate(), peerCertificate(), peerCertificateChain(), sessionCipher(), privateKey(), ciphers(), and caCertificates().
*/
func (this *QSslSocket) SslConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16sslConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslConfiguration(const QSslConfiguration &)

/*
Sets the socket's SSL configuration to be the contents of configuration. This function sets the local certificate, the ciphers, the private key and the CA certificates to those stored in configuration.

It is not possible to set the SSL-state related fields.

This function was introduced in  Qt 4.4.

See also sslConfiguration(), setLocalCertificate(), setPrivateKey(), setCaCertificates(), and setCiphers().
*/
func (this *QSslSocket) SetSslConfiguration(config QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QSslConfiguration_PTR() != nil {
		convArg0 = config.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QSslCertificate &)

/*
Sets the socket's local certificate to certificate. The local certificate is necessary if you need to confirm your identity to the peer. It is used together with the private key; if you set the local certificate, you must also set the private key.

The local certificate and private key are always necessary for server sockets, but are also rarely used by client sockets if the server requires the client to authenticate.

Note: Secure Transport SSL backend on macOS may update the default keychain (the default is probably your login keychain) by importing your local certificates and keys. This can also result in system dialogs showing up and asking for permission when your application is using these private keys. If such behavior is undesired, set the QT_SSL_USE_TEMPORARY_KEYCHAIN environment variable to a non-zero value; this will prompt QSslSocket to use its own temporary keychain.

See also localCertificate() and setPrivateKey().
*/
func (this *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:139
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QString &, QSsl::EncodingFormat)

/*
Sets the socket's local certificate to certificate. The local certificate is necessary if you need to confirm your identity to the peer. It is used together with the private key; if you set the local certificate, you must also set the private key.

The local certificate and private key are always necessary for server sockets, but are also rarely used by client sockets if the server requires the client to authenticate.

Note: Secure Transport SSL backend on macOS may update the default keychain (the default is probably your login keychain) by importing your local certificates and keys. This can also result in system dialogs showing up and asking for permission when your application is using these private keys. If such behavior is undesired, set the QT_SSL_USE_TEMPORARY_KEYCHAIN environment variable to a non-zero value; this will prompt QSslSocket to use its own temporary keychain.

See also localCertificate() and setPrivateKey().
*/
func (this *QSslSocket) SetLocalCertificate_1(fileName string, format int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK7QStringN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:139
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QString &, QSsl::EncodingFormat)

/*
Sets the socket's local certificate to certificate. The local certificate is necessary if you need to confirm your identity to the peer. It is used together with the private key; if you set the local certificate, you must also set the private key.

The local certificate and private key are always necessary for server sockets, but are also rarely used by client sockets if the server requires the client to authenticate.

Note: Secure Transport SSL backend on macOS may update the default keychain (the default is probably your login keychain) by importing your local certificates and keys. This can also result in system dialogs showing up and asking for permission when your application is using these private keys. If such behavior is undesired, set the QT_SSL_USE_TEMPORARY_KEYCHAIN environment variable to a non-zero value; this will prompt QSslSocket to use its own temporary keychain.

See also localCertificate() and setPrivateKey().
*/
func (this *QSslSocket) SetLocalCertificate_1_(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum,
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK7QStringN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate localCertificate() const

/*
Returns the socket's local certificate, or an empty certificate if no local certificate has been assigned.

See also setLocalCertificate() and privateKey().
*/
func (this *QSslSocket) LocalCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16localCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate peerCertificate() const

/*
Returns the peer's digital certificate (i.e., the immediate certificate of the host you are connected to), or a null certificate, if the peer has not assigned a certificate.

The peer certificate is checked automatically during the handshake phase, so this function is normally used to fetch the certificate for display or for connection diagnostic purposes. It contains information about the peer, including its host name, the certificate issuer, and the peer's public key.

Because the peer certificate is set during the handshake phase, it is safe to access the peer certificate from a slot connected to the sslErrors() signal or the encrypted() signal.

If a null certificate is returned, it can mean the SSL handshake failed, or it can mean the host you are connected to doesn't have a certificate, or it can mean there is no connection.

If you want to check the peer's complete chain of certificates, use peerCertificateChain() to get them all at once.

See also peerCertificateChain().
*/
func (this *QSslSocket) PeerCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15peerCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCipher sessionCipher() const

/*
Returns the socket's cryptographic cipher, or a null cipher if the connection isn't encrypted. The socket's cipher for the session is set during the handshake phase. The cipher is used to encrypt and decrypt data transmitted through the socket.

QSslSocket also provides functions for setting the ordered list of ciphers from which the handshake phase will eventually select the session cipher. This ordered list must be in place before the handshake phase begins.

See also ciphers(), setCiphers(), setDefaultCiphers(), defaultCiphers(), and supportedCiphers().
*/
func (this *QSslSocket) SessionCipher() *QSslCipher /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket13sessionCipherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol sessionProtocol() const

/*
Returns the socket's SSL/TLS protocol or UnknownProtocol if the connection isn't encrypted. The socket's protocol for the session is set during the handshake phase.

This function was introduced in  Qt 5.4.

See also protocol() and setProtocol().
*/
func (this *QSslSocket) SessionProtocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15sessionProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QSslKey &)

/*
Sets the socket's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QSslKey", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)

/*
Sets the socket's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslSocket) SetPrivateKey_1(fileName string, algorithm int, format int, passPhrase qtcore.QByteArray_ITF) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg3 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg3 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)

/*
Sets the socket's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslSocket) SetPrivateKey_1_(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QSsl::KeyAlgorithm=Elaborated, QSsl::KeyAlgorithm=Enum,
	algorithm := 0
	// arg: 2, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum,
	format := 0
	// arg: 3, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg3 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)

/*
Sets the socket's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslSocket) SetPrivateKey_1_1(fileName string, algorithm int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum,
	format := 0
	// arg: 3, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg3 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)

/*
Sets the socket's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslSocket) SetPrivateKey_1_2(fileName string, algorithm int, format int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg3 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey privateKey() const

/*
Returns this socket's private key.

See also setPrivateKey() and localCertificate().
*/
func (this *QSslSocket) PrivateKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket10privateKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCiphers(const QString &)

/*

 */
func (this *QSslSocket) SetCiphers(ciphers string) {
	var tmpArg0 = qtcore.NewQString_5(ciphers)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket10setCiphersERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates encoded in the specified format and adds them to this socket's CA certificate database. path must be a file or a pattern matching one or more files, as specified by syntax. Returns true if one or more certificates are added to the socket's CA certificate database; otherwise returns false.

The CA certificate database is used by the socket during the handshake phase to validate the peer's certificate.

For more precise control, use addCaCertificate().

See also addCaCertificate() and QSslCertificate::fromPath().
*/
func (this *QSslSocket) AddCaCertificates(path string, format int, syntax int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates encoded in the specified format and adds them to this socket's CA certificate database. path must be a file or a pattern matching one or more files, as specified by syntax. Returns true if one or more certificates are added to the socket's CA certificate database; otherwise returns false.

The CA certificate database is used by the socket during the handshake phase to validate the peer's certificate.

For more precise control, use addCaCertificate().

See also addCaCertificate() and QSslCertificate::fromPath().
*/
func (this *QSslSocket) AddCaCertificates__(path string) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum,
	format := 0
	// arg: 2, QRegExp::PatternSyntax=Elaborated, QRegExp::PatternSyntax=Enum,
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates encoded in the specified format and adds them to this socket's CA certificate database. path must be a file or a pattern matching one or more files, as specified by syntax. Returns true if one or more certificates are added to the socket's CA certificate database; otherwise returns false.

The CA certificate database is used by the socket during the handshake phase to validate the peer's certificate.

For more precise control, use addCaCertificate().

See also addCaCertificate() and QSslCertificate::fromPath().
*/
func (this *QSslSocket) AddCaCertificates__1(path string, format int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegExp::PatternSyntax=Elaborated, QRegExp::PatternSyntax=Enum,
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCaCertificate(const QSslCertificate &)

/*
Adds the certificate to this socket's CA certificate database. The CA certificate database is used by the socket during the handshake phase to validate the peer's certificate.

To add multiple certificates, use addCaCertificates().

See also caCertificates() and setCaCertificates().
*/
func (this *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16addCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool addDefaultCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates with the specified encoding and adds them to the default CA certificate database. path can be an explicit file, or it can contain wildcards in the format specified by syntax. Returns true if any CA certificates are added to the default database.

Each SSL socket's CA certificate database is initialized to the default CA certificate database.

See also defaultCaCertificates(), addCaCertificates(), and addDefaultCaCertificate().
*/
func (this *QSslSocket) AddDefaultCaCertificates(path string, format int, syntax int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QSslSocket_AddDefaultCaCertificates(path string, format int, syntax int) bool {
	var nilthis *QSslSocket
	rv := nilthis.AddDefaultCaCertificates(path, format, syntax)
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool addDefaultCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates with the specified encoding and adds them to the default CA certificate database. path can be an explicit file, or it can contain wildcards in the format specified by syntax. Returns true if any CA certificates are added to the default database.

Each SSL socket's CA certificate database is initialized to the default CA certificate database.

See also defaultCaCertificates(), addCaCertificates(), and addDefaultCaCertificate().
*/
func (this *QSslSocket) AddDefaultCaCertificates__(path string) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum,
	format := 0
	// arg: 2, QRegExp::PatternSyntax=Elaborated, QRegExp::PatternSyntax=Enum,
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool addDefaultCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)

/*
Searches all files in the path for certificates with the specified encoding and adds them to the default CA certificate database. path can be an explicit file, or it can contain wildcards in the format specified by syntax. Returns true if any CA certificates are added to the default database.

Each SSL socket's CA certificate database is initialized to the default CA certificate database.

See also defaultCaCertificates(), addCaCertificates(), and addDefaultCaCertificate().
*/
func (this *QSslSocket) AddDefaultCaCertificates__1(path string, format int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QRegExp::PatternSyntax=Elaborated, QRegExp::PatternSyntax=Enum,
	syntax := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:174
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addDefaultCaCertificate(const QSslCertificate &)

/*
Adds certificate to the default CA certificate database. Each SSL socket's CA certificate database is initialized to the default CA certificate database.

See also defaultCaCertificates() and addCaCertificates().
*/
func (this *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23addDefaultCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	var nilthis *QSslSocket
	nilthis.AddDefaultCaCertificate(certificate)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Reimplemented from QAbstractSocket::waitForConnected().

Waits until the socket is connected, or msecs milliseconds, whichever happens first. If the connection has been established, this function returns true; otherwise it returns false.

See also QAbstractSocket::waitForConnected().
*/
func (this *QSslSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Reimplemented from QAbstractSocket::waitForConnected().

Waits until the socket is connected, or msecs milliseconds, whichever happens first. If the connection has been established, this function returns true; otherwise it returns false.

See also QAbstractSocket::waitForConnected().
*/
func (this *QSslSocket) WaitForConnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:183
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForEncrypted(int)

/*
Waits until the socket has completed the SSL handshake and has emitted encrypted(), or msecs milliseconds, whichever comes first. If encrypted() has been emitted, this function returns true; otherwise (e.g., the socket is disconnected, or the SSL handshake fails), false is returned.

The following example waits up to one second for the socket to be encrypted:


  socket->connectToHostEncrypted("imap", 993);
  if (socket->waitForEncrypted(1000))
      qDebug("Encrypted!");



If msecs is -1, this function will not time out.

See also startClientEncryption(), startServerEncryption(), encrypted(), and isEncrypted().
*/
func (this *QSslSocket) WaitForEncrypted(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForEncryptedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:183
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForEncrypted(int)

/*
Waits until the socket has completed the SSL handshake and has emitted encrypted(), or msecs milliseconds, whichever comes first. If encrypted() has been emitted, this function returns true; otherwise (e.g., the socket is disconnected, or the SSL handshake fails), false is returned.

The following example waits up to one second for the socket to be encrypted:


  socket->connectToHostEncrypted("imap", 993);
  if (socket->waitForEncrypted(1000))
      qDebug("Encrypted!");



If msecs is -1, this function will not time out.

See also startClientEncryption(), startServerEncryption(), encrypted(), and isEncrypted().
*/
func (this *QSslSocket) WaitForEncrypted__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForEncryptedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().
*/
func (this *QSslSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().
*/
func (this *QSslSocket) WaitForReadyRead__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QSslSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QSslSocket) WaitForBytesWritten__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Reimplemented from QAbstractSocket::waitForDisconnected().

Waits until the socket has disconnected or msecs milliseconds, whichever comes first. If the connection has been disconnected, this function returns true; otherwise it returns false.

See also QAbstractSocket::waitForDisconnected().
*/
func (this *QSslSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Reimplemented from QAbstractSocket::waitForDisconnected().

Waits until the socket has disconnected or msecs milliseconds, whichever comes first. If the connection has been disconnected, this function returns true; otherwise it returns false.

See also QAbstractSocket::waitForDisconnected().
*/
func (this *QSslSocket) WaitForDisconnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:190
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool supportsSsl()

/*
Returns true if this platform supports SSL; otherwise, returns false. If the platform doesn't support SSL, the socket will fail in the connection phase.
*/
func (this *QSslSocket) SupportsSsl() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11supportsSslEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QSslSocket_SupportsSsl() bool {
	var nilthis *QSslSocket
	rv := nilthis.SupportsSsl()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:191
// index:0
// Public static Visibility=Default Availability=Available
// [8] long sslLibraryVersionNumber()

/*
Returns the version number of the SSL library in use. Note that this is the version of the library in use at run-time not compile time. If no SSL support is available then this will return an undefined value.

This function was introduced in  Qt 5.0.
*/
func (this *QSslSocket) SslLibraryVersionNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionNumberEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QSslSocket_SslLibraryVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:192
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sslLibraryVersionString()

/*
Returns the version string of the SSL library in use. Note that this is the version of the library in use at run-time not compile time. If no SSL support is available then this will return an empty value.

This function was introduced in  Qt 5.0.
*/
func (this *QSslSocket) SslLibraryVersionString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionStringEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QSslSocket_SslLibraryVersionString() string {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:193
// index:0
// Public static Visibility=Default Availability=Available
// [8] long sslLibraryBuildVersionNumber()

/*
Returns the version number of the SSL library in use at compile time. If no SSL support is available then this will return an undefined value.

This function was introduced in  Qt 5.4.

See also sslLibraryVersionNumber().
*/
func (this *QSslSocket) SslLibraryBuildVersionNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionNumberEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QSslSocket_SslLibraryBuildVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:194
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sslLibraryBuildVersionString()

/*
Returns the version string of the SSL library in use at compile time. If no SSL support is available then this will return an empty value.

This function was introduced in  Qt 5.4.

See also sslLibraryVersionString().
*/
func (this *QSslSocket) SslLibraryBuildVersionString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionStringEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QSslSocket_SslLibraryBuildVersionString() string {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startClientEncryption()

/*
Starts a delayed SSL handshake for a client connection. This function can be called when the socket is in the ConnectedState but still in the UnencryptedMode. If it is not yet connected, or if it is already encrypted, this function has no effect.

Clients that implement STARTTLS functionality often make use of delayed SSL handshakes. Most other clients can avoid calling this function directly by using connectToHostEncrypted() instead, which automatically performs the handshake.

See also connectToHostEncrypted() and startServerEncryption().
*/
func (this *QSslSocket) StartClientEncryption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21startClientEncryptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startServerEncryption()

/*
Starts a delayed SSL handshake for a server connection. This function can be called when the socket is in the ConnectedState but still in UnencryptedMode. If it is not connected or it is already encrypted, the function has no effect.

For server sockets, calling this function is the only way to initiate the SSL handshake. Most servers will call this function immediately upon receiving a connection, or as a result of having received a protocol-specific command to enter SSL mode (e.g, the server may respond to receiving the string "STARTTLS\r\n" by calling this function).

The most common way to implement an SSL server is to create a subclass of QTcpServer and reimplement QTcpServer::incomingConnection(). The returned socket descriptor is then passed to QSslSocket::setSocketDescriptor().

See also connectToHostEncrypted() and startClientEncryption().
*/
func (this *QSslSocket) StartServerEncryption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21startServerEncryptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignoreSslErrors()

/*
This slot tells QSslSocket to ignore errors during QSslSocket's handshake phase and continue connecting. If you want to continue with the connection even if errors occur during the handshake phase, then you must call this slot, either from a slot connected to sslErrors(), or before the handshake phase. If you don't call this slot, either in response to errors or before the handshake, the connection will be dropped after the sslErrors() signal has been emitted.

If there are no errors during the SSL handshake phase (i.e., the identity of the peer is established with no problems), QSslSocket will not emit the sslErrors() signal, and it is unnecessary to call this function.

Warning: Be sure to always let the user inspect the errors reported by the sslErrors() signal, and only call this method upon confirmation from the user that proceeding is ok. If there are unexpected errors, the connection should be aborted. Calling this method without inspecting the actual errors will most likely pose a security risk for your application. Use it with great care!

See also sslErrors().
*/
func (this *QSslSocket) IgnoreSslErrors() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15ignoreSslErrorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted()

/*
This signal is emitted when QSslSocket enters encrypted mode. After this signal has been emitted, QSslSocket::isEncrypted() will return true, and all further transmissions on the socket will be encrypted.

See also QSslSocket::connectToHostEncrypted() and QSslSocket::isEncrypted().
*/
func (this *QSslSocket) Encrypted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket9encryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void peerVerifyError(const QSslError &)

/*
QSslSocket can emit this signal several times during the SSL handshake, before encryption has been established, to indicate that an error has occurred while establishing the identity of the peer. The error is usually an indication that QSslSocket is unable to securely identify the peer.

This signal provides you with an early indication when something's wrong. By connecting to this signal, you can manually choose to tear down the connection from inside the connected slot before the handshake has completed. If no action is taken, QSslSocket will proceed to emitting QSslSocket::sslErrors().

This function was introduced in  Qt 4.4.

See also sslErrors().
*/
func (this *QSslSocket) PeerVerifyError(error QSslError_ITF) {
	var convArg0 unsafe.Pointer
	if error != nil && error.QSslError_PTR() != nil {
		convArg0 = error.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15peerVerifyErrorERK9QSslError", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modeChanged(QSslSocket::SslMode)

/*
This signal is emitted when QSslSocket changes from QSslSocket::UnencryptedMode to either QSslSocket::SslClientMode or QSslSocket::SslServerMode. mode is the new mode.

See also QSslSocket::mode().
*/
func (this *QSslSocket) ModeChanged(newMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11modeChangedENS_7SslModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encryptedBytesWritten(qint64)

/*
This signal is emitted when QSslSocket writes its encrypted data to the network. The written parameter contains the number of bytes that were successfully written.

This function was introduced in  Qt 4.4.

See also QIODevice::bytesWritten().
*/
func (this *QSslSocket) EncryptedBytesWritten(totalBytes int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21encryptedBytesWrittenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), totalBytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)

/*
QSslSocket emits this signal when it negotiates a PSK ciphersuite, and therefore a PSK authentication is then required.

When using PSK, the client must send to the server a valid identity and a valid pre shared key, in order for the SSL handshake to continue. Applications can provide this information in a slot connected to this signal, by filling in the passed authenticator object according to their needs.

Note: Ignoring this signal, or failing to provide the required credentials, will cause the handshake to fail, and therefore the connection to be aborted.

Note: The authenticator object is owned by the socket and must not be deleted by the application.

This function was introduced in  Qt 5.5.

See also QSslPreSharedKeyAuthenticator.
*/
func (this *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QSslSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QSslSocket) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

/*
Describes the connection modes available for QSslSocket.


*/
type QSslSocket__SslMode = int

// The socket is unencrypted. Its behavior is identical to QTcpSocket.
const QSslSocket__UnencryptedMode QSslSocket__SslMode = 0

// The socket is a client-side SSL socket. It is either alreayd encrypted, or it is in the SSL handshake phase (see QSslSocket::isEncrypted()).
const QSslSocket__SslClientMode QSslSocket__SslMode = 1

// The socket is a server-side SSL socket. It is either already encrypted, or it is in the SSL handshake phase (see QSslSocket::isEncrypted()).
const QSslSocket__SslServerMode QSslSocket__SslMode = 2

/*
Describes the peer verification modes for QSslSocket. The default mode is AutoVerifyPeer, which selects an appropriate mode depending on the socket's QSocket::SslMode.



This enum was introduced or modified in  Qt 4.4.

See also QSslSocket::peerVerifyMode().

*/
type QSslSocket__PeerVerifyMode = int

// QSslSocket will not request a certificate from the peer. You can set this mode if you are not interested in the identity of the other side of the connection. The connection will still be encrypted, and your socket will still send its local certificate to the peer if it's requested.
const QSslSocket__VerifyNone QSslSocket__PeerVerifyMode = 0

// QSslSocket will request a certificate from the peer, but does not require this certificate to be valid. This is useful when you want to display peer certificate details to the user without affecting the actual SSL handshake. This mode is the default for servers.
const QSslSocket__QueryPeer QSslSocket__PeerVerifyMode = 1

// QSslSocket will request a certificate from the peer during the SSL handshake phase, and requires that this certificate is valid. On failure, QSslSocket will emit the QSslSocket::sslErrors() signal. This mode is the default for clients.
const QSslSocket__VerifyPeer QSslSocket__PeerVerifyMode = 2

// QSslSocket will automatically use QueryPeer for server sockets and VerifyPeer for client sockets.
const QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = 3

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
