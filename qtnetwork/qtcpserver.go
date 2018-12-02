package qtnetwork

// /usr/include/qt/QtNetwork/qtcpserver.h
// #include <qtcpserver.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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

// void incomingConnection(qintptr)
func (this *QTcpServer) InheritIncomingConnection(f func(handle int64) /*void*/) {
	qtrt.SetAllInheritCallback(this, "incomingConnection", f)
}

// void addPendingConnection(QTcpSocket *)
func (this *QTcpServer) InheritAddPendingConnection(f func(socket *QTcpSocket /*777 QTcpSocket **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addPendingConnection", f)
}

/*

 */
type QTcpServer struct {
	*qtcore.QObject
}
type QTcpServer_ITF interface {
	qtcore.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer { return ptr }

func (this *QTcpServer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTcpServer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQTcpServerFromPointer(cthis unsafe.Pointer) *QTcpServer {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTcpServer{bcthis0}
}
func (*QTcpServer) NewFromPointer(cthis unsafe.Pointer) *QTcpServer {
	return NewQTcpServerFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTcpServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qtcpserver.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTcpServer(QObject *)

/*
Constructs a QTcpServer object.

parent is passed to the QObject constructor.

See also listen() and setSocketDescriptor().
*/
func (*QTcpServer) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QTcpServer {
	return NewQTcpServer(parent)
}
func NewQTcpServer(parent qtcore.QObject_ITF /*777 QObject **/) *QTcpServer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTcpServerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTcpServer")
	return gothis
}

// /usr/include/qt/QtNetwork/qtcpserver.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTcpServer(QObject *)

/*
Constructs a QTcpServer object.

parent is passed to the QObject constructor.

See also listen() and setSocketDescriptor().
*/
func (*QTcpServer) NewForInherit__() *QTcpServer {
	return NewQTcpServer__()
}
func NewQTcpServer__() *QTcpServer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTcpServerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTcpServer")
	return gothis
}

// /usr/include/qt/QtNetwork/qtcpserver.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTcpServer()

/*

 */
func DeleteQTcpServer(this *QTcpServer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QHostAddress &, quint16)

/*
Tells the server to listen for incoming connections on address address and port port. If port is 0, a port is chosen automatically. If address is QHostAddress::Any, the server will listen on all network interfaces.

Returns true on success; otherwise returns false.

See also isListening().
*/
func (this *QTcpServer) Listen(address QHostAddress_ITF, port uint16) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QHostAddress &, quint16)

/*
Tells the server to listen for incoming connections on address address and port port. If port is 0, a port is chosen automatically. If address is QHostAddress::Any, the server will listen on all network interfaces.

Returns true on success; otherwise returns false.

See also isListening().
*/
func (this *QTcpServer) Listen__() bool {
	// arg: 0, const QHostAddress &=LValueReference, QHostAddress=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QHostAddress &, quint16)

/*
Tells the server to listen for incoming connections on address address and port port. If port is 0, a port is chosen automatically. If address is QHostAddress::Any, the server will listen on all network interfaces.

Returns true on success; otherwise returns false.

See also isListening().
*/
func (this *QTcpServer) Listen__1(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()

/*
Closes the server. The server will no longer listen for incoming connections.

See also listen().
*/
func (this *QTcpServer) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isListening() const

/*
Returns true if the server is currently listening for incoming connections; otherwise returns false.

See also listen().
*/
func (this *QTcpServer) IsListening() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11isListeningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxPendingConnections(int)

/*
Sets the maximum number of pending accepted connections to numConnections. QTcpServer will accept no more than numConnections incoming connections before nextPendingConnection() is called. By default, the limit is 30 pending connections.

Clients may still able to connect after the server has reached its maximum number of pending connections (i.e., QTcpSocket can still emit the connected() signal). QTcpServer will stop accepting the new connections, but the operating system may still keep them in queue.

See also maxPendingConnections() and hasPendingConnections().
*/
func (this *QTcpServer) SetMaxPendingConnections(numConnections int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer24setMaxPendingConnectionsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxPendingConnections() const

/*
Returns the maximum number of pending accepted connections. The default is 30.

See also setMaxPendingConnections() and hasPendingConnections().
*/
func (this *QTcpServer) MaxPendingConnections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer21maxPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qtcpserver.h:72
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 serverPort() const

/*
Returns the server's port if the server is listening for connections; otherwise returns 0.

See also serverAddress() and listen().
*/
func (this *QTcpServer) ServerPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer10serverPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress serverAddress() const

/*
Returns the server's address if the server is listening for connections; otherwise returns QHostAddress::Null.

See also serverPort() and listen().
*/
func (this *QTcpServer) ServerAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer13serverAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qintptr socketDescriptor() const

/*
Returns the native socket descriptor the server uses to listen for incoming instructions, or -1 if the server is not listening.

If the server is using QNetworkProxy, the returned descriptor may not be usable with native socket functions.

See also setSocketDescriptor() and isListening().
*/
func (this *QTcpServer) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr)

/*
Sets the socket descriptor this server should use when listening for incoming connections to socketDescriptor. Returns true if the socket is set successfully; otherwise returns false.

The socket is assumed to be in listening state.

See also socketDescriptor() and isListening().
*/
func (this *QTcpServer) SetSocketDescriptor(socketDescriptor int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer19setSocketDescriptorEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is disadvised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QTcpServer) WaitForNewConnection(msec int, timedOut *bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is disadvised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QTcpServer) WaitForNewConnection__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msec := int(0)
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is disadvised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QTcpServer) WaitForNewConnection__1(msec int) bool {
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasPendingConnections() const

/*
Returns true if the server has a pending connection; otherwise returns false.

See also nextPendingConnection() and setMaxPendingConnections().
*/
func (this *QTcpServer) HasPendingConnections() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer21hasPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QTcpSocket * nextPendingConnection()

/*
Returns the next pending connection as a connected QTcpSocket object.

The socket is created as a child of the server, which means that it is automatically deleted when the QTcpServer object is destroyed. It is still a good idea to delete the object explicitly when you are done with it, to avoid wasting memory.

0 is returned if this function is called when there are no pending connections.

Note: The returned QTcpSocket object cannot be used from another thread. If you want to use an incoming connection from another thread, you need to override incomingConnection().

See also hasPendingConnections().
*/
func (this *QTcpServer) NextPendingConnection() *QTcpSocket /*777 QTcpSocket **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer21nextPendingConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTcpSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qtcpserver.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError serverError() const

/*
Returns an error code for the last error that occurred.

See also errorString().
*/
func (this *QTcpServer) ServerError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11serverErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a human readable description of the last error that occurred.

See also serverError().
*/
func (this *QTcpServer) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qtcpserver.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pauseAccepting()

/*
Pauses accepting new connections. Queued connections will remain in queue.

This function was introduced in  Qt 5.0.

See also resumeAccepting().
*/
func (this *QTcpServer) PauseAccepting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer14pauseAcceptingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resumeAccepting()

/*
Resumes accepting new connections.

This function was introduced in  Qt 5.0.

See also pauseAccepting().
*/
func (this *QTcpServer) ResumeAccepting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer15resumeAcceptingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxy(const QNetworkProxy &)

/*
Sets the explicit network proxy for this socket to networkProxy.

To disable the use of a proxy for this socket, use the QNetworkProxy::NoProxy proxy type:


  server->setProxy(QNetworkProxy::NoProxy);



This function was introduced in  Qt 4.1.

See also proxy() and QNetworkProxy.
*/
func (this *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if networkProxy != nil && networkProxy.QNetworkProxy_PTR() != nil {
		convArg0 = networkProxy.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer8setProxyERK13QNetworkProxy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy proxy() const

/*
Returns the network proxy for this socket. By default QNetworkProxy::DefaultProxy is used.

This function was introduced in  Qt 4.1.

See also setProxy() and QNetworkProxy.
*/
func (this *QTcpServer) Proxy() *QNetworkProxy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer5proxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void incomingConnection(qintptr)

/*
This virtual function is called by QTcpServer when a new connection is available. The socketDescriptor argument is the native socket descriptor for the accepted connection.

The base implementation creates a QTcpSocket, sets the socket descriptor and then stores the QTcpSocket in an internal list of pending connections. Finally newConnection() is emitted.

Reimplement this function to alter the server's behavior when a connection is available.

If this server is using QNetworkProxy then the socketDescriptor may not be usable with native socket functions, and should only be used with QTcpSocket::setSocketDescriptor().

Note: If another socket is created in the reimplementation of this method, it needs to be added to the Pending Connections mechanism by calling addPendingConnection().

Note: If you want to handle an incoming connection as a new QTcpSocket object in another thread you have to pass the socketDescriptor to the other thread and create the QTcpSocket object there and use its setSocketDescriptor() method.

See also newConnection(), nextPendingConnection(), and addPendingConnection().
*/
func (this *QTcpServer) IncomingConnection(handle int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer18incomingConnectionEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addPendingConnection(QTcpSocket *)

/*
This function is called by QTcpServer::incomingConnection() to add the socket to the list of pending incoming connections.

Note: Don't forget to call this member from reimplemented incomingConnection() if you do not want to break the Pending Connections mechanism.

This function was introduced in  Qt 4.7.

See also incomingConnection().
*/
func (this *QTcpServer) AddPendingConnection(socket QTcpSocket_ITF /*777 QTcpSocket **/) {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QTcpSocket_PTR() != nil {
		convArg0 = socket.QTcpSocket_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20addPendingConnectionEP10QTcpSocket", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void newConnection()

/*
This signal is emitted every time a new connection is available.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QTcpServer) NewConnection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer13newConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acceptError(QAbstractSocket::SocketError)

/*
This signal is emitted when accepting a new connection results in an error. The socketError parameter describes the type of error that occurred.

This function was introduced in  Qt 5.0.

See also pauseAccepting() and resumeAccepting().
*/
func (this *QTcpServer) AcceptError(socketError int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer11acceptErrorEN15QAbstractSocket11SocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
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
		qtcore.KeepMe()
	}
}

//  keep block end
