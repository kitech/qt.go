package qtnetwork

// /usr/include/qt/QtNetwork/qlocalserver.h
// #include <qlocalserver.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QLocalServer struct {
	*qtcore.QObject
}

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

// /usr/include/qt/QtNetwork/qlocalserver.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QLocalServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:60
// index:0
// Public
// void newConnection()
func (this *QLocalServer) NewConnection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer13newConnectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:74
// index:0
// Public
// void QLocalServer(class QObject *)
func NewQLocalServer(parent *qtcore.QObject /*777 QObject **/) *QLocalServer {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocalServerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalserver.h:75
// index:0
// Public virtual
// void ~QLocalServer()
func DeleteQLocalServer(*QLocalServer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:77
// index:0
// Public
// void close()
func (this *QLocalServer) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:78
// index:0
// Public
// QString errorString()
func (this *QLocalServer) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:79
// index:0
// Public virtual
// bool hasPendingConnections()
func (this *QLocalServer) HasPendingConnections() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer21hasPendingConnectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:80
// index:0
// Public
// bool isListening()
func (this *QLocalServer) IsListening() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer11isListeningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:81
// index:0
// Public
// bool listen(const class QString &)
func (this *QLocalServer) Listen(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer6listenERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:82
// index:1
// Public
// bool listen(qintptr)
func (this *QLocalServer) Listen_1(socketDescriptor int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer6listenEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:83
// index:0
// Public
// int maxPendingConnections()
func (this *QLocalServer) MaxPendingConnections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer21maxPendingConnectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qlocalserver.h:84
// index:0
// Public virtual
// QLocalSocket * nextPendingConnection()
func (this *QLocalServer) NextPendingConnection() *QLocalSocket /*777 QLocalSocket **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer21nextPendingConnectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLocalSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:85
// index:0
// Public
// QString serverName()
func (this *QLocalServer) ServerName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer10serverNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:86
// index:0
// Public
// QString fullServerName()
func (this *QLocalServer) FullServerName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer14fullServerNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:87
// index:0
// Public static
// bool removeServer(const class QString &)
func (this *QLocalServer) RemoveServer(name *qtcore.QString) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer12removeServerERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QLocalServer_RemoveServer(name *qtcore.QString) bool {
	var nilthis *QLocalServer
	rv := nilthis.RemoveServer(name)
	return rv
}

// /usr/include/qt/QtNetwork/qlocalserver.h:88
// index:0
// Public
// QAbstractSocket::SocketError serverError()
func (this *QLocalServer) ServerError() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer11serverErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:89
// index:0
// Public
// void setMaxPendingConnections(int)
func (this *QLocalServer) SetMaxPendingConnections(numConnections int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer24setMaxPendingConnectionsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:90
// index:0
// Public
// bool waitForNewConnection(int, _Bool *)
func (this *QLocalServer) WaitForNewConnection(msec int, timedOut unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer20waitForNewConnectionEiPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec, &timedOut)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:93
// index:0
// Public
// QLocalServer::SocketOptions socketOptions()
func (this *QLocalServer) SocketOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer13socketOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:95
// index:0
// Public
// qintptr socketDescriptor()
func (this *QLocalServer) SocketDescriptor() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QLocalServer16socketDescriptorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalserver.h:98
// index:0
// Protected virtual
// void incomingConnection(quintptr)
func (this *QLocalServer) IncomingConnection(socketDescriptor uint64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QLocalServer18incomingConnectionEy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	gopp.ErrPrint(err, rv)
}

type QLocalServer__SocketOption = int

const QLocalServer__NoOptions QLocalServer__SocketOption = 0
const QLocalServer__UserAccessOption QLocalServer__SocketOption = 1
const QLocalServer__GroupAccessOption QLocalServer__SocketOption = 2
const QLocalServer__OtherAccessOption QLocalServer__SocketOption = 4
const QLocalServer__WorldAccessOption QLocalServer__SocketOption = 7

//  body block end
