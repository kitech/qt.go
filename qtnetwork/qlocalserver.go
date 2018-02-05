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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
// void incomingConnection(quintptr)
func (this *QLocalServer) InheritIncomingConnection(f func(socketDescriptor uint64) /*void*/) {
	qtrt.SetAllInheritCallback(this, "incomingConnection", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QLocalServer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void newConnection()
func (this *QLocalServer) NewConnection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer13newConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLocalServer(QObject *)
func NewQLocalServer(parent *qtcore.QObject /*777 QObject **/) *QLocalServer {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLocalServerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qlocalserver.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLocalServer()
func DeleteQLocalServer(this *QLocalServer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()
func (this *QLocalServer) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QLocalServer) ErrorString() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasPendingConnections()
func (this *QLocalServer) HasPendingConnections() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer21hasPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isListening()
func (this *QLocalServer) IsListening() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11isListeningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool listen(const QString &)
func (this *QLocalServer) Listen(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer6listenERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:82
// index:1
// Public Visibility=Default Availability=Available
// [1] bool listen(qintptr)
func (this *QLocalServer) Listen_1(socketDescriptor int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer6listenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxPendingConnections()
func (this *QLocalServer) MaxPendingConnections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer21maxPendingConnectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qlocalserver.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLocalSocket * nextPendingConnection()
func (this *QLocalServer) NextPendingConnection() *QLocalSocket /*777 QLocalSocket **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer21nextPendingConnectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLocalSocketFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString serverName()
func (this *QLocalServer) ServerName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer10serverNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fullServerName()
func (this *QLocalServer) FullServerName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer14fullServerNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qlocalserver.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool removeServer(const QString &)
func (this *QLocalServer) RemoveServer(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer12removeServerERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
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
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::SocketError serverError()
func (this *QLocalServer) ServerError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer11serverErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxPendingConnections(int)
func (this *QLocalServer) SetMaxPendingConnections(numConnections int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer24setMaxPendingConnectionsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numConnections)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNewConnection(int, _Bool *)
func (this *QLocalServer) WaitForNewConnection(msec int, timedOut unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer20waitForNewConnectionEiPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, &timedOut)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qlocalserver.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSocketOptions(QLocalServer::SocketOptions)
func (this *QLocalServer) SetSocketOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer16setSocketOptionsE6QFlagsINS_12SocketOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QLocalServer::SocketOptions socketOptions()
func (this *QLocalServer) SocketOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer13socketOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qlocalserver.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] qintptr socketDescriptor()
func (this *QLocalServer) SocketDescriptor() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QLocalServer16socketDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qlocalserver.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void incomingConnection(quintptr)
func (this *QLocalServer) IncomingConnection(socketDescriptor uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QLocalServer18incomingConnectionEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor)
	gopp.ErrPrint(err, rv)
}

type QLocalServer__SocketOption = int

const QLocalServer__NoOptions QLocalServer__SocketOption = 0
const QLocalServer__UserAccessOption QLocalServer__SocketOption = 1
const QLocalServer__GroupAccessOption QLocalServer__SocketOption = 2
const QLocalServer__OtherAccessOption QLocalServer__SocketOption = 4
const QLocalServer__WorldAccessOption QLocalServer__SocketOption = 7

//  body block end
