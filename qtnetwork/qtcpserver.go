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

// void addPendingConnection(class QTcpSocket *)
func (this *QTcpServer) InheritAddPendingConnection(f func(socket *QTcpSocket /*777 QTcpSocket **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addPendingConnection", f)
}

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
func (this *QTcpServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qtcpserver.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTcpServer(QObject *)
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
func NewQTcpServer__() *QTcpServer {
	// arg: 0, QObject *=Pointer, QObject=Record,
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
func (this *QTcpServer) Listen__() bool {
	// arg: 0, const QHostAddress &=LValueReference, QHostAddress=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QHostAddress &, quint16)
func (this *QTcpServer) Listen__1(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()
func (this *QTcpServer) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isListening() const
func (this *QTcpServer) IsListening() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11isListeningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxPendingConnections(int)
func (this *QTcpServer) SetMaxPendingConnections(numConnections int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer24setMaxPendingConnectionsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxPendingConnections() const
func (this *QTcpServer) MaxPendingConnections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer21maxPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qtcpserver.h:72
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 serverPort() const
func (this *QTcpServer) ServerPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer10serverPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress serverAddress() const
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
func (this *QTcpServer) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr)
func (this *QTcpServer) SetSocketDescriptor(socketDescriptor int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer19setSocketDescriptorEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, _Bool *)
func (this *QTcpServer) WaitForNewConnection(msec int, timedOut *bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, _Bool *)
func (this *QTcpServer) WaitForNewConnection__() bool {
	// arg: 0, int=Int, =Invalid,
	msec := int(0)
	// arg: 1, bool *=Pointer, =Invalid,
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, _Bool *)
func (this *QTcpServer) WaitForNewConnection__1(msec int) bool {
	// arg: 1, bool *=Pointer, =Invalid,
	var timedOut unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timedOut)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasPendingConnections() const
func (this *QTcpServer) HasPendingConnections() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer21hasPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QTcpSocket * nextPendingConnection()
func (this *QTcpServer) NextPendingConnection() *QTcpSocket /*777 QTcpSocket **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer21nextPendingConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTcpSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qtcpserver.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError serverError() const
func (this *QTcpServer) ServerError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11serverErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QTcpServer) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpServer11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qtcpserver.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pauseAccepting()
func (this *QTcpServer) PauseAccepting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer14pauseAcceptingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resumeAccepting()
func (this *QTcpServer) ResumeAccepting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer15resumeAcceptingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxy(const QNetworkProxy &)
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
func (this *QTcpServer) IncomingConnection(handle int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer18incomingConnectionEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addPendingConnection(QTcpSocket *)
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
func (this *QTcpServer) NewConnection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpServer13newConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void acceptError(QAbstractSocket::SocketError)
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
