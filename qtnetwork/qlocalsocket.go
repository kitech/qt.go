package qtnetwork

// /usr/include/qt/QtNetwork/qlocalsocket.h
// #include <qlocalsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
func (this *QLocalSocket) InheritReadData(f func(arg0 string, arg1 int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QLocalSocket) InheritWriteData(f func(arg0 string, arg1 int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

/*

 */
type QLocalSocket struct {
	*qtcore.QIODevice
}
type QLocalSocket_ITF interface {
	qtcore.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func (ptr *QLocalSocket) QLocalSocket_PTR() *QLocalSocket { return ptr }

func (this *QLocalSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QLocalSocket) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = qtcore.NewQIODeviceFromPointer(cthis)
}
func NewQLocalSocketFromPointer(cthis unsafe.Pointer) *QLocalSocket {
	bcthis0 := qtcore.NewQIODeviceFromPointer(cthis)
	return &QLocalSocket{bcthis0}
}
func (*QLocalSocket) NewFromPointer(cthis unsafe.Pointer) *QLocalSocket {
	return NewQLocalSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLocalSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalSocket(QObject *)

/*
Creates a new local socket. The parent argument is passed to QObject's constructor.
*/
func NewQLocalSocket(parent qtcore.QObject_ITF /*777 QObject **/) *QLocalSocket {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocalSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLocalSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalSocket(QObject *)

/*
Creates a new local socket. The parent argument is passed to QObject's constructor.
*/
func NewQLocalSocket__() *QLocalSocket {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLocalSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLocalSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLocalSocket()

/*

 */
func DeleteQLocalSocket(this *QLocalSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(QIODevice::OpenMode)

/*
Attempts to make a connection to serverName(). setServerName() must be called before you open the connection. Alternatively you can use connectToServer(const QString &name, OpenMode openMode);

The socket is opened in the given openMode and first enters ConnectingState. If a connection is established, QLocalSocket enters ConnectedState and emits connected().

After calling this function, the socket can emit error() to signal that an error occurred.

This function was introduced in  Qt 5.1.

See also state(), serverName(), and waitForConnected().
*/
func (this *QLocalSocket) ConnectToServer(openMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(QIODevice::OpenMode)

/*
Attempts to make a connection to serverName(). setServerName() must be called before you open the connection. Alternatively you can use connectToServer(const QString &name, OpenMode openMode);

The socket is opened in the given openMode and first enters ConnectingState. If a connection is established, QLocalSocket enters ConnectedState and emits connected().

After calling this function, the socket can emit error() to signal that an error occurred.

This function was introduced in  Qt 5.1.

See also state(), serverName(), and waitForConnected().
*/
func (this *QLocalSocket) ConnectToServer__() {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(const QString &, QIODevice::OpenMode)

/*
Attempts to make a connection to serverName(). setServerName() must be called before you open the connection. Alternatively you can use connectToServer(const QString &name, OpenMode openMode);

The socket is opened in the given openMode and first enters ConnectingState. If a connection is established, QLocalSocket enters ConnectedState and emits connected().

After calling this function, the socket can emit error() to signal that an error occurred.

This function was introduced in  Qt 5.1.

See also state(), serverName(), and waitForConnected().
*/
func (this *QLocalSocket) ConnectToServer_1(name string, openMode int) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(const QString &, QIODevice::OpenMode)

/*
Attempts to make a connection to serverName(). setServerName() must be called before you open the connection. Alternatively you can use connectToServer(const QString &name, OpenMode openMode);

The socket is opened in the given openMode and first enters ConnectingState. If a connection is established, QLocalSocket enters ConnectedState and emits connected().

After calling this function, the socket can emit error() to signal that an error occurred.

This function was introduced in  Qt 5.1.

See also state(), serverName(), and waitForConnected().
*/
func (this *QLocalSocket) ConnectToServer_1_(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnectFromServer()

/*
Attempts to close the socket. If there is pending data waiting to be written, QLocalSocket will enter ClosingState and wait until all data has been written. Eventually, it will enter UnconnectedState and emit the disconnectedFromServer() signal.

See also connectToServer().
*/
func (this *QLocalSocket) DisconnectFromServer() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket20disconnectFromServerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setServerName(const QString &)

/*
Set the name of the peer to connect to. On Windows name is the name of a named pipe; on Unix name is the name of a local domain socket.

This function must be called when the socket is not connected.

This function was introduced in  Qt 5.1.

See also serverName().
*/
func (this *QLocalSocket) SetServerName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket13setServerNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString serverName() const

/*
Returns the name of the peer as specified by setServerName(), or an empty QString if setServerName() has not been called or connectToServer() failed.

See also setServerName(), connectToServer(), and fullServerName().
*/
func (this *QLocalSocket) ServerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket10serverNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fullServerName() const

/*
Returns the server path that the socket is connected to.

Note: The return value of this function is platform specific.

See also connectToServer() and serverName().
*/
func (this *QLocalSocket) FullServerName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket14fullServerNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()

/*
Aborts the current connection and resets the socket. Unlike disconnectFromServer(), this function immediately closes the socket, clearing any pending data in the write buffer.

See also disconnectFromServer() and close().
*/
func (this *QLocalSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const

/*
Reimplemented from QIODevice::isSequential().
*/
func (this *QLocalSocket) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const

/*
Reimplemented from QIODevice::bytesAvailable().
*/
func (this *QLocalSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const

/*
Reimplemented from QIODevice::bytesToWrite().
*/
func (this *QLocalSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const

/*
Reimplemented from QIODevice::canReadLine().
*/
func (this *QLocalSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Equivalent to connectToServer(OpenMode mode). The socket is opened in the given openMode to the server defined by setServerName().

Note that unlike in most other QIODevice subclasses, open() may not open the device directly. The function return false if the socket was already connected or if the server to connect to was not defined and true in any other case. The connected() or error() signals will be emitted once the device is actualy open (or the connection failed).

See connectToServer() for more details.
*/
func (this *QLocalSocket) Open(openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Equivalent to connectToServer(OpenMode mode). The socket is opened in the given openMode to the server defined by setServerName().

Note that unlike in most other QIODevice subclasses, open() may not open the device directly. The function return false if the socket was already connected or if the server to connect to was not defined and true in any other case. The connected() or error() signals will be emitted once the device is actualy open (or the connection failed).

See connectToServer() for more details.
*/
func (this *QLocalSocket) Open__() bool {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().
*/
func (this *QLocalSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalSocket::LocalSocketError error() const

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QLocalSocket) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QLocalSocket::LocalSocketError)

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QLocalSocket) Error_1(socketError int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5errorENS_16LocalSocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()

/*
This function writes as much as possible from the internal write buffer to the socket, without blocking. If any data was written, this function returns true; otherwise false is returned.

Call this function if you need QLocalSocket to start sending buffered data immediately. The number of bytes successfully written depends on the operating system. In most cases, you do not need to call this function, because QLocalSocket will start sending data automatically once control goes back to the event loop. In the absence of an event loop, call waitForBytesWritten() instead.

See also write() and waitForBytesWritten().
*/
func (this *QLocalSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the socket is valid and ready for use; otherwise returns false.

Note: The socket's state must be ConnectedState before reading and writing can occur.

See also state() and connectToServer().
*/
func (this *QLocalSocket) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize() const

/*
Returns the size of the internal read buffer. This limits the amount of data that the client can receive before you call read() or readAll(). A read buffer size of 0 (the default) means that the buffer has no size limit, ensuring that no data is lost.

See also setReadBufferSize() and read().
*/
func (this *QLocalSocket) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)

/*
Sets the size of QLocalSocket's internal read buffer to be size bytes.

If the buffer size is limited to a certain size, QLocalSocket won't buffer more than this size of data. Exceptionally, a buffer size of 0 means that the read buffer is unlimited and all incoming data is buffered. This is the default.

This option is useful if you only read the data at certain points in time (e.g., in a real-time streaming application) or if you want to protect your socket against receiving too much data, which may eventually cause your application to run out of memory.

See also readBufferSize() and read().
*/
func (this *QLocalSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QLocalSocket::LocalSocketState, QIODevice::OpenMode)

/*
Initializes QLocalSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState.

Note: It is not possible to initialize two local sockets with the same native socket descriptor.

See also socketDescriptor(), state(), and openMode().
*/
func (this *QLocalSocket) SetSocketDescriptor(socketDescriptor int64, socketState int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QLocalSocket::LocalSocketState, QIODevice::OpenMode)

/*
Initializes QLocalSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState.

Note: It is not possible to initialize two local sockets with the same native socket descriptor.

See also socketDescriptor(), state(), and openMode().
*/
func (this *QLocalSocket) SetSocketDescriptor__(socketDescriptor int64) bool {
	// arg: 1, QLocalSocket::LocalSocketState=Enum, QLocalSocket::LocalSocketState=Enum, , Invalid
	socketState := 0
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, QLocalSocket::LocalSocketState, QIODevice::OpenMode)

/*
Initializes QLocalSocket with the native socket descriptor socketDescriptor. Returns true if socketDescriptor is accepted as a valid socket descriptor; otherwise returns false. The socket is opened in the mode specified by openMode, and enters the socket state specified by socketState.

Note: It is not possible to initialize two local sockets with the same native socket descriptor.

See also socketDescriptor(), state(), and openMode().
*/
func (this *QLocalSocket) SetSocketDescriptor__1(socketDescriptor int64, socketState int) bool {
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>, Unexposed
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qintptr socketDescriptor() const

/*
Returns the native socket descriptor of the QLocalSocket object if this is available; otherwise returns -1.

The socket descriptor is not available when QLocalSocket is in UnconnectedState. The type of the descriptor depends on the platform:


On Windows, the returned value is a Winsock 2 Socket Handle.
With WinRT and on INTEGRITY, the returned value is the QTcpSocket socket descriptor and the type is defined by socketDescriptor.
On all other UNIX-like operating systems, the type is a file descriptor representing a socket.


See also setSocketDescriptor().
*/
func (this *QLocalSocket) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalSocket::LocalSocketState state() const

/*
Returns the state of the socket.

See also error().
*/
func (this *QLocalSocket) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QLocalSocket) WaitForBytesWritten__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Waits until the socket is connected, up to msecs milliseconds. If the connection has been established, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be established:


  socket->connectToServer("market");
  if (socket->waitForConnected(1000))
      qDebug("Connected!");



If msecs is -1, this function will not time out.

See also connectToServer() and connected().
*/
func (this *QLocalSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForConnected(int)

/*
Waits until the socket is connected, up to msecs milliseconds. If the connection has been established, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be established:


  socket->connectToServer("market");
  if (socket->waitForConnected(1000))
      qDebug("Connected!");



If msecs is -1, this function will not time out.

See also connectToServer() and connected().
*/
func (this *QLocalSocket) WaitForConnected__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Waits until the socket has disconnected, up to msecs milliseconds. If the connection has been disconnected, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be closed:


  socket->disconnectFromServer();
  if (socket->waitForDisconnected(1000))
      qDebug("Disconnected!");



If msecs is -1, this function will not time out.

See also disconnectFromServer() and close().
*/
func (this *QLocalSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)

/*
Waits until the socket has disconnected, up to msecs milliseconds. If the connection has been disconnected, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for a connection to be closed:


  socket->disconnectFromServer();
  if (socket->waitForDisconnected(1000))
      qDebug("Disconnected!");



If msecs is -1, this function will not time out.

See also disconnectFromServer() and close().
*/
func (this *QLocalSocket) WaitForDisconnected__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().

This function blocks until data is available for reading and the readyRead() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if data is available for reading; otherwise it returns false (if an error occurred or the operation timed out).

See also waitForBytesWritten().
*/
func (this *QLocalSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().

This function blocks until data is available for reading and the readyRead() signal has been emitted. The function will timeout after msecs milliseconds; the default timeout is 30000 milliseconds.

The function returns true if data is available for reading; otherwise it returns false (if an error occurred or the operation timed out).

See also waitForBytesWritten().
*/
func (this *QLocalSocket) WaitForReadyRead__() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connected()

/*
This signal is emitted after connectToServer() has been called and a connection has been successfully established.

See also connectToServer() and disconnected().
*/
func (this *QLocalSocket) Connected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket9connectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnected()

/*
This signal is emitted when the socket has been disconnected.

See also connectToServer(), disconnectFromServer(), abort(), and connected().
*/
func (this *QLocalSocket) Disconnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket12disconnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QLocalSocket::LocalSocketState)

/*
This signal is emitted whenever QLocalSocket's state changes. The socketState parameter is the new state.

QLocalSocket::SocketState is not a registered metatype, so for queued connections, you will have to register it with Q_DECLARE_METATYPE() and qRegisterMetaType().

See also state() and Creating Custom Qt Types.
*/
func (this *QLocalSocket) StateChanged(socketState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket12stateChangedENS_16LocalSocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QLocalSocket) ReadData(arg0 string, arg1 int64) int64 {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QLocalSocket) WriteData(arg0 string, arg1 int64) int64 {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

/*
The LocalServerError enumeration represents the errors that can occur. The most recent error can be retrieved through a call to QLocalSocket::error().

QLocalSocket::ConnectionRefusedErrorQAbstractSocket::ConnectionRefusedErrorThe connection was refused by the peer (or timed out).
QLocalSocket::PeerClosedErrorQAbstractSocket::RemoteHostClosedErrorThe remote socket closed the connection. Note that the client socket (i.e., this socket) will be closed after the remote close notification has been sent.
QLocalSocket::ServerNotFoundErrorQAbstractSocket::HostNotFoundErrorThe local socket name was not found.
QLocalSocket::SocketAccessErrorQAbstractSocket::SocketAccessErrorThe socket operation failed because the application lacked the required privileges.
QLocalSocket::SocketResourceErrorQAbstractSocket::SocketResourceErrorThe local system ran out of resources (e.g., too many sockets).
QLocalSocket::SocketTimeoutErrorQAbstractSocket::SocketTimeoutErrorThe socket operation timed out.
QLocalSocket::ConnectionErrorQAbstractSocket::NetworkErrorAn error occurred with the connection.
QLocalSocket::UnsupportedSocketOperationErrorQAbstractSocket::UnsupportedSocketOperationErrorThe requested socket operation is not supported by the local operating system.
QLocalSocket::OperationErrorQAbstractSocket::OperationErrorAn operation was attempted while the socket was in a state that did not permit it.
QLocalSocket::UnknownSocketErrorQAbstractSocket::UnknownSocketErrorAn unidentified error occurred.

*/
type QLocalSocket__LocalSocketError = int

//
const QLocalSocket__ConnectionRefusedError QLocalSocket__LocalSocketError = 0

//
const QLocalSocket__PeerClosedError QLocalSocket__LocalSocketError = 1

//
const QLocalSocket__ServerNotFoundError QLocalSocket__LocalSocketError = 2

//
const QLocalSocket__SocketAccessError QLocalSocket__LocalSocketError = 3

//
const QLocalSocket__SocketResourceError QLocalSocket__LocalSocketError = 4

//
const QLocalSocket__SocketTimeoutError QLocalSocket__LocalSocketError = 5

//
const QLocalSocket__DatagramTooLargeError QLocalSocket__LocalSocketError = 6

//
const QLocalSocket__ConnectionError QLocalSocket__LocalSocketError = 7

//
const QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = 10

//
const QLocalSocket__UnknownSocketError QLocalSocket__LocalSocketError = -1

//
const QLocalSocket__OperationError QLocalSocket__LocalSocketError = 19

func (this *QLocalSocket) LocalSocketErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLocalSocket_LocalSocketErrorItemName(val int) string {
	var nilthis *QLocalSocket
	return nilthis.LocalSocketErrorItemName(val)
}

/*
This enum describes the different states in which a socket can be.

QLocalSocket::UnconnectedStateQAbstractSocket::UnconnectedStateThe socket is not connected.
QLocalSocket::ConnectingStateQAbstractSocket::ConnectingStateThe socket has started establishing a connection.
QLocalSocket::ConnectedStateQAbstractSocket::ConnectedStateA connection is established.
QLocalSocket::ClosingStateQAbstractSocket::ClosingStateThe socket is about to close (data may still be waiting to be written).


See also QLocalSocket::state().

*/
type QLocalSocket__LocalSocketState = int

//
const QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = 0

//
const QLocalSocket__ConnectingState QLocalSocket__LocalSocketState = 2

//
const QLocalSocket__ConnectedState QLocalSocket__LocalSocketState = 3

//
const QLocalSocket__ClosingState QLocalSocket__LocalSocketState = 6

func (this *QLocalSocket) LocalSocketStateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLocalSocket_LocalSocketStateItemName(val int) string {
	var nilthis *QLocalSocket
	return nilthis.LocalSocketStateItemName(val)
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
