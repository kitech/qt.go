package qtnetwork

// /usr/include/qt/QtNetwork/qtcpserver.h
// #include <qtcpserver.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTcpServer struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QTcpServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:61
// index:0
// Public
// void QTcpServer(QObject *)
func NewQTcpServer(parent *qtcore.QObject /*777 QObject **/) *QTcpServer {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTcpServerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qtcpserver.h:62
// index:0
// Public virtual
// void ~QTcpServer()
func DeleteQTcpServer(*QTcpServer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:64
// index:0
// Public
// bool listen(const QHostAddress &, quint16)
func (this *QTcpServer) Listen(address *QHostAddress, port uint16) bool {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer6listenERK12QHostAddresst", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:65
// index:0
// Public
// void close()
func (this *QTcpServer) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:67
// index:0
// Public
// bool isListening()
func (this *QTcpServer) IsListening() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer11isListeningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:69
// index:0
// Public
// void setMaxPendingConnections(int)
func (this *QTcpServer) SetMaxPendingConnections(numConnections int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer24setMaxPendingConnectionsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:70
// index:0
// Public
// int maxPendingConnections()
func (this *QTcpServer) MaxPendingConnections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer21maxPendingConnectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qtcpserver.h:72
// index:0
// Public
// quint16 serverPort()
func (this *QTcpServer) ServerPort() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer10serverPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:73
// index:0
// Public
// QHostAddress serverAddress()
func (this *QTcpServer) ServerAddress() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer13serverAddressEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:75
// index:0
// Public
// qintptr socketDescriptor()
func (this *QTcpServer) SocketDescriptor() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer16socketDescriptorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qtcpserver.h:76
// index:0
// Public
// bool setSocketDescriptor(qintptr)
func (this *QTcpServer) SetSocketDescriptor(socketDescriptor int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer19setSocketDescriptorEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:78
// index:0
// Public
// bool waitForNewConnection(int, _Bool *)
func (this *QTcpServer) WaitForNewConnection(msec int, timedOut unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer20waitForNewConnectionEiPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec, &timedOut)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:79
// index:0
// Public virtual
// bool hasPendingConnections()
func (this *QTcpServer) HasPendingConnections() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer21hasPendingConnectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qtcpserver.h:80
// index:0
// Public virtual
// QTcpSocket * nextPendingConnection()
func (this *QTcpServer) NextPendingConnection() *QTcpSocket /*777 QTcpSocket **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer21nextPendingConnectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTcpSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:82
// index:0
// Public
// QAbstractSocket::SocketError serverError()
func (this *QTcpServer) ServerError() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer11serverErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:83
// index:0
// Public
// QString errorString()
func (this *QTcpServer) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:85
// index:0
// Public
// void pauseAccepting()
func (this *QTcpServer) PauseAccepting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer14pauseAcceptingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:86
// index:0
// Public
// void resumeAccepting()
func (this *QTcpServer) ResumeAccepting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer15resumeAcceptingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:89
// index:0
// Public
// void setProxy(const QNetworkProxy &)
func (this *QTcpServer) SetProxy(networkProxy *QNetworkProxy) {
	var convArg0 = networkProxy.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer8setProxyERK13QNetworkProxy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:90
// index:0
// Public
// QNetworkProxy proxy()
func (this *QTcpServer) Proxy() *QNetworkProxy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpServer5proxyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpserver.h:94
// index:0
// Protected virtual
// void incomingConnection(qintptr)
func (this *QTcpServer) IncomingConnection(handle int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer18incomingConnectionEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:95
// index:0
// Protected
// void addPendingConnection(QTcpSocket *)
func (this *QTcpServer) AddPendingConnection(socket *QTcpSocket /*777 QTcpSocket **/) {
	var convArg0 = socket.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer20addPendingConnectionEP10QTcpSocket", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:101
// index:0
// Public
// void newConnection()
func (this *QTcpServer) NewConnection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer13newConnectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qtcpserver.h:102
// index:0
// Public
// void acceptError(QAbstractSocket::SocketError)
func (this *QTcpServer) AcceptError(socketError int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpServer11acceptErrorEN15QAbstractSocket11SocketErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	gopp.ErrPrint(err, rv)
}

//  body block end
