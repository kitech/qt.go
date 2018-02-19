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
func (this *QLocalSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalSocket(QObject *)
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
func NewQLocalSocket__() *QLocalSocket {
	// arg: 0, QObject *=Pointer, QObject=Record,
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
func (this *QLocalSocket) ConnectToServer(openMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(QIODevice::OpenMode)
func (this *QLocalSocket) ConnectToServer__() {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToServer(const QString &, QIODevice::OpenMode)
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
func (this *QLocalSocket) ConnectToServer_1_(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket15connectToServerERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, openMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnectFromServer()
func (this *QLocalSocket) DisconnectFromServer() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket20disconnectFromServerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setServerName(const QString &)
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
func (this *QLocalSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const
func (this *QLocalSocket) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const
func (this *QLocalSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const
func (this *QLocalSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const
func (this *QLocalSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QLocalSocket) Open(openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QLocalSocket) Open__() bool {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QLocalSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalSocket::LocalSocketError error() const
func (this *QLocalSocket) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QLocalSocket::LocalSocketError)
func (this *QLocalSocket) Error_1(socketError int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5errorENS_16LocalSocketErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()
func (this *QLocalSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QLocalSocket) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readBufferSize() const
func (this *QLocalSocket) ReadBufferSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket14readBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)
func (this *QLocalSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QLocalSocket::LocalSocketState, QIODevice::OpenMode)
func (this *QLocalSocket) SetSocketDescriptor(socketDescriptor int64, socketState int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QLocalSocket::LocalSocketState, QIODevice::OpenMode)
func (this *QLocalSocket) SetSocketDescriptor__(socketDescriptor int64) bool {
	// arg: 1, QLocalSocket::LocalSocketState=Enum, QLocalSocket::LocalSocketState=Enum,
	socketState := 0
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QLocalSocket::LocalSocketState, QIODevice::OpenMode)
func (this *QLocalSocket) SetSocketDescriptor__1(socketDescriptor int64, socketState int) bool {
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	openMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19setSocketDescriptorExNS_16LocalSocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, socketState, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] qintptr socketDescriptor() const
func (this *QLocalSocket) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalSocket::LocalSocketState state() const
func (this *QLocalSocket) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalSocket5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QLocalSocket) WaitForBytesWritten__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := 30000
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QLocalSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QLocalSocket) WaitForConnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := 30000
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QLocalSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QLocalSocket) WaitForDisconnected__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := 30000
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QLocalSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QLocalSocket) WaitForReadyRead__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := 30000
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connected()
func (this *QLocalSocket) Connected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket9connectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void disconnected()
func (this *QLocalSocket) Disconnected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket12disconnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QLocalSocket::LocalSocketState)
func (this *QLocalSocket) StateChanged(socketState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket12stateChangedENS_16LocalSocketStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalsocket.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
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
func (this *QLocalSocket) WriteData(arg0 string, arg1 int64) int64 {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

type QLocalSocket__LocalSocketError = int

const QLocalSocket__ConnectionRefusedError QLocalSocket__LocalSocketError = 0
const QLocalSocket__PeerClosedError QLocalSocket__LocalSocketError = 1
const QLocalSocket__ServerNotFoundError QLocalSocket__LocalSocketError = 2
const QLocalSocket__SocketAccessError QLocalSocket__LocalSocketError = 3
const QLocalSocket__SocketResourceError QLocalSocket__LocalSocketError = 4
const QLocalSocket__SocketTimeoutError QLocalSocket__LocalSocketError = 5
const QLocalSocket__DatagramTooLargeError QLocalSocket__LocalSocketError = 6
const QLocalSocket__ConnectionError QLocalSocket__LocalSocketError = 7
const QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = 10
const QLocalSocket__UnknownSocketError QLocalSocket__LocalSocketError = -1
const QLocalSocket__OperationError QLocalSocket__LocalSocketError = 19

type QLocalSocket__LocalSocketState = int

const QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = 0
const QLocalSocket__ConnectingState QLocalSocket__LocalSocketState = 2
const QLocalSocket__ConnectedState QLocalSocket__LocalSocketState = 3
const QLocalSocket__ClosingState QLocalSocket__LocalSocketState = 6

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
