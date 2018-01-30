package qtnetwork

// /usr/include/qt/QtNetwork/qudpsocket.h
// #include <qudpsocket.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QUdpSocket struct {
	*QAbstractSocket
}

func (this *QUdpSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSocket.GetCthis()
	}
}
func (this *QUdpSocket) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSocket = NewQAbstractSocketFromPointer(cthis)
}
func NewQUdpSocketFromPointer(cthis unsafe.Pointer) *QUdpSocket {
	bcthis0 := NewQAbstractSocketFromPointer(cthis)
	return &QUdpSocket{bcthis0}
}
func (*QUdpSocket) NewFromPointer(cthis unsafe.Pointer) *QUdpSocket {
	return NewQUdpSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QUdpSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUdpSocket10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUdpSocket(QObject *)
func NewQUdpSocket(parent *qtcore.QObject /*777 QObject **/) *QUdpSocket {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocketC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUdpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qudpsocket.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUdpSocket()
func DeleteQUdpSocket(*QUdpSocket) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocketD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &)
func (this *QUdpSocket) JoinMulticastGroup(groupAddress *QHostAddress) bool {
	var convArg0 = groupAddress.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:65
// index:1
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &, const QNetworkInterface &)
func (this *QUdpSocket) JoinMulticastGroup_1(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	var convArg0 = groupAddress.GetCthis()
	var convArg1 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddressRK17QNetworkInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &)
func (this *QUdpSocket) LeaveMulticastGroup(groupAddress *QHostAddress) bool {
	var convArg0 = groupAddress.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:68
// index:1
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &, const QNetworkInterface &)
func (this *QUdpSocket) LeaveMulticastGroup_1(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	var convArg0 = groupAddress.GetCthis()
	var convArg1 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddressRK17QNetworkInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkInterface multicastInterface()
func (this *QUdpSocket) MulticastInterface() *QNetworkInterface /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUdpSocket18multicastInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMulticastInterface(const QNetworkInterface &)
func (this *QUdpSocket) SetMulticastInterface(iface *QNetworkInterface) {
	var convArg0 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket21setMulticastInterfaceERK17QNetworkInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasPendingDatagrams()
func (this *QUdpSocket) HasPendingDatagrams() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUdpSocket19hasPendingDatagramsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pendingDatagramSize()
func (this *QUdpSocket) PendingDatagramSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUdpSocket19pendingDatagramSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram receiveDatagram(qint64)
func (this *QUdpSocket) ReceiveDatagram(maxSize int64) *QNetworkDatagram /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket15receiveDatagramEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)
func (this *QUdpSocket) ReadDatagram(data string, maxlen int64, host *QHostAddress /*777 QHostAddress **/, port unsafe.Pointer /*666*/) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, &port)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QNetworkDatagram &)
func (this *QUdpSocket) WriteDatagram(datagram *QNetworkDatagram) int64 {
	var convArg0 = datagram.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK16QNetworkDatagram", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:81
// index:1
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const char *, qint64, const QHostAddress &, quint16)
func (this *QUdpSocket) WriteDatagram_1(data string, len int64, host *QHostAddress, port uint16) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramEPKcxRK12QHostAddresst", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len, convArg2, port)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:82
// index:2
// Public inline Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QByteArray &, const QHostAddress &, quint16)
func (this *QUdpSocket) WriteDatagram_2(datagram *qtcore.QByteArray, host *QHostAddress, port uint16) int64 {
	var convArg0 = datagram.GetCthis()
	var convArg1 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK10QByteArrayRK12QHostAddresst", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, port)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

//  body block end
