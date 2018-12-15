package qtnetwork

// /usr/include/qt/QtNetwork/qlocalserver.h
// #include <qlocalserver.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// void incomingConnection(quintptr)
func (this *QLocalServer) InheritIncomingConnection(f func(socketDescriptor uint64) /*void*/) {
	qtrt.SetAllInheritCallback(this, "incomingConnection", f)
}

/*

 */
type QLocalServer struct {
	*qtcore.QObject
}
type QLocalServer_ITF interface {
	qtcore.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func (ptr *QLocalServer) QLocalServer_PTR() *QLocalServer { return ptr }

func (this *QLocalServer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLocalServer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQLocalServerFromPointer(cthis unsafe.Pointer) *QLocalServer {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QLocalServer{bcthis0}
}
func (*QLocalServer) NewFromPointer(cthis unsafe.Pointer) *QLocalServer {
	return NewQLocalServerFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLocalServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qlocalserver.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void newConnection()

/*
This signal is emitted every time a new connection is available.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QLocalServer) NewConnection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer13newConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalServer(QObject *)

/*
Create a new local socket server with the given parent.

See also listen().
*/
func (*QLocalServer) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QLocalServer {
	return NewQLocalServer(parent)
}
func NewQLocalServer(parent qtcore.QObject_ITF /*777 QObject **/) *QLocalServer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocalServerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLocalServer")
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalserver.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalServer(QObject *)

/*
Create a new local socket server with the given parent.

See also listen().
*/
func (*QLocalServer) NewForInheritp() *QLocalServer {
	return NewQLocalServerp()
}
func NewQLocalServerp() *QLocalServer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocalServerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLocalServer")
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalserver.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLocalServer()

/*

 */
func DeleteQLocalServer(this *QLocalServer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()

/*
Stop listening for incoming connections. Existing connections are not affected, but any new connections will be refused.

See also isListening() and listen().
*/
func (this *QLocalServer) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns the human-readable message appropriate to the current error reported by serverError(). If no suitable string is available, an empty string is returned.

See also serverError().
*/
func (this *QLocalServer) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qlocalserver.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasPendingConnections() const

/*
Returns true if the server has a pending connection; otherwise returns false.

See also nextPendingConnection() and setMaxPendingConnections().
*/
func (this *QLocalServer) HasPendingConnections() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer21hasPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isListening() const

/*
Returns true if the server is listening for incoming connections otherwise false.

See also listen() and close().
*/
func (this *QLocalServer) IsListening() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11isListeningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QString &)

/*
Tells the server to listen for incoming connections on name. If the server is currently listening then it will return false. Return true on success otherwise false.

name can be a single name and QLocalServer will determine the correct platform specific path. serverName() will return the name that is passed into listen.

Usually you would just pass in a name like "foo", but on Unix this could also be a path such as "/tmp/foo" and on Windows this could be a pipe path such as "\\.\pipe\foo"

Note: On Unix if the server crashes without closing listen will fail with AddressInUseError. To create a new server the file should be removed. On Windows two local servers can listen to the same pipe at the same time, but any connections will go to one of the server.

See also serverName(), isListening(), and close().
*/
func (this *QLocalServer) Listen(name string) bool {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer6listenERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:83
// index:1
// Public Visibility=Default Availability=Available
// [1] bool listen(qintptr)

/*
Tells the server to listen for incoming connections on name. If the server is currently listening then it will return false. Return true on success otherwise false.

name can be a single name and QLocalServer will determine the correct platform specific path. serverName() will return the name that is passed into listen.

Usually you would just pass in a name like "foo", but on Unix this could also be a path such as "/tmp/foo" and on Windows this could be a pipe path such as "\\.\pipe\foo"

Note: On Unix if the server crashes without closing listen will fail with AddressInUseError. To create a new server the file should be removed. On Windows two local servers can listen to the same pipe at the same time, but any connections will go to one of the server.

See also serverName(), isListening(), and close().
*/
func (this *QLocalServer) Listen1(socketDescriptor int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer6listenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxPendingConnections() const

/*
Returns the maximum number of pending accepted connections. The default is 30.

See also setMaxPendingConnections() and hasPendingConnections().
*/
func (this *QLocalServer) MaxPendingConnections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer21maxPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qlocalserver.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLocalSocket * nextPendingConnection()

/*
Returns the next pending connection as a connected QLocalSocket object.

The socket is created as a child of the server, which means that it is automatically deleted when the QLocalServer object is destroyed. It is still a good idea to delete the object explicitly when you are done with it, to avoid wasting memory.

0 is returned if this function is called when there are no pending connections.

See also hasPendingConnections(), newConnection(), and incomingConnection().
*/
func (this *QLocalServer) NextPendingConnection() *QLocalSocket /*777 QLocalSocket **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer21nextPendingConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLocalSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qlocalserver.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString serverName() const

/*
Returns the server name if the server is listening for connections; otherwise returns QString()

See also listen() and fullServerName().
*/
func (this *QLocalServer) ServerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer10serverNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qlocalserver.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fullServerName() const

/*
Returns the full path that the server is listening on.

Note: This is platform specific

See also listen() and serverName().
*/
func (this *QLocalServer) FullServerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer14fullServerNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qlocalserver.h:88
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool removeServer(const QString &)

/*
Removes any server instance that might cause a call to listen() to fail and returns true if successful; otherwise returns false. This function is meant to recover from a crash, when the previous server instance has not been cleaned up.

On Windows, this function does nothing; on Unix, it removes the socket file given by name.

Warning: Be careful to avoid removing sockets of running instances.

This function was introduced in  Qt 4.5.
*/
func (this *QLocalServer) RemoveServer(name string) bool {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer12removeServerERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QLocalServer_RemoveServer(name string) bool {
	var nilthis *QLocalServer
	rv := nilthis.RemoveServer(name)
	return rv
}

// /usr/include/qt/QtNetwork/qlocalserver.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError serverError() const

/*
Returns the type of error that occurred last or NoError.

See also errorString().
*/
func (this *QLocalServer) ServerError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11serverErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxPendingConnections(int)

/*
Sets the maximum number of pending accepted connections to numConnections. QLocalServer will accept no more than numConnections incoming connections before nextPendingConnection() is called.

Note: Even though QLocalServer will stop accepting new connections after it has reached its maximum number of pending connections, the operating system may still keep them in queue which will result in clients signaling that it is connected.

See also maxPendingConnections() and hasPendingConnections().
*/
func (this *QLocalServer) SetMaxPendingConnections(numConnections int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer24setMaxPendingConnectionsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is ill-advised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QLocalServer) WaitForNewConnection(msec int, timedOut *bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is ill-advised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QLocalServer) WaitForNewConnectionp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msec := int(0)
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, bool *)

/*
Waits for at most msec milliseconds or until an incoming connection is available. Returns true if a connection is available; otherwise returns false. If the operation timed out and timedOut is not 0, *timedOut will be set to true.

This is a blocking function call. Its use is ill-advised in a single-threaded GUI application, since the whole application will stop responding until the function returns. waitForNewConnection() is mostly useful when there is no event loop available.

The non-blocking alternative is to connect to the newConnection() signal.

If msec is -1, this function will not time out.

See also hasPendingConnections() and nextPendingConnection().
*/
func (this *QLocalServer) WaitForNewConnectionp1(msec int) bool {
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSocketOptions(QLocalServer::SocketOptions)

/*

 */
func (this *QLocalServer) SetSocketOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer16setSocketOptionsE6QFlagsINS_12SocketOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalServer::SocketOptions socketOptions() const

/*
Returns the socket options set on the socket.

This function was introduced in  Qt 5.0.

Note: Getter function for property socketOptions.

See also setSocketOptions().
*/
func (this *QLocalServer) SocketOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer13socketOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void incomingConnection(quintptr)

/*
This virtual function is called by QLocalServer when a new connection is available. socketDescriptor is the native socket descriptor for the accepted connection.

The base implementation creates a QLocalSocket, sets the socket descriptor and then stores the QLocalSocket in an internal list of pending connections. Finally newConnection() is emitted.

Reimplement this function to alter the server's behavior when a connection is available.

See also newConnection(), nextPendingConnection(), and QLocalSocket::setSocketDescriptor().
*/
func (this *QLocalServer) IncomingConnection(socketDescriptor uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer18incomingConnectionEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QLocalServer__SocketOption = int

//
const QLocalServer__NoOptions QLocalServer__SocketOption = 0

//
const QLocalServer__UserAccessOption QLocalServer__SocketOption = 1

//
const QLocalServer__GroupAccessOption QLocalServer__SocketOption = 2

//
const QLocalServer__OtherAccessOption QLocalServer__SocketOption = 4

//
const QLocalServer__WorldAccessOption QLocalServer__SocketOption = 7

func (this *QLocalServer) SocketOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLocalServer_SocketOptionItemName(val int) string {
	var nilthis *QLocalServer
	return nilthis.SocketOptionItemName(val)
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
