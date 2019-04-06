package qtnetwork

// /usr/include/qt/QtNetwork/qudpsocket.h
// #include <qudpsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

/*

 */
type QUdpSocket struct {
	*QAbstractSocket
}
type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QUdpSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qudpsocket.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUdpSocket(QObject *)

/*
Creates a QUdpSocket object.

parent is passed to the QObject constructor.

See also socketType().
*/
func (*QUdpSocket) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QUdpSocket {
	return NewQUdpSocket(parent)
}
func NewQUdpSocket(parent qtcore.QObject_ITF /*777 QObject **/) *QUdpSocket {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUdpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUdpSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qudpsocket.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUdpSocket(QObject *)

/*
Creates a QUdpSocket object.

parent is passed to the QObject constructor.

See also socketType().
*/
func (*QUdpSocket) NewForInheritp() *QUdpSocket {
	return NewQUdpSocketp()
}
func NewQUdpSocketp() *QUdpSocket {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUdpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUdpSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qudpsocket.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUdpSocket()

/*

 */
func DeleteQUdpSocket(this *QUdpSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &)

/*
Joins the multicast group specified by groupAddress on the default interface chosen by the operating system. The socket must be in BoundState, otherwise an error occurs.

Note that if you are attempting to join an IPv4 group, your socket must not be bound using IPv6 (or in dual mode, using QHostAddress::Any). You must use QHostAddress::AnyIPv4 instead.

This function returns true if successful; otherwise it returns false and sets the socket error accordingly.

Note: Joining IPv6 multicast groups without an interface selection is not supported in all operating systems. Consider using the overload where the interface is specified.

This function was introduced in  Qt 4.8.

See also leaveMulticastGroup().
*/
func (this *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:65
// index:1
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &, const QNetworkInterface &)

/*
Joins the multicast group specified by groupAddress on the default interface chosen by the operating system. The socket must be in BoundState, otherwise an error occurs.

Note that if you are attempting to join an IPv4 group, your socket must not be bound using IPv6 (or in dual mode, using QHostAddress::Any). You must use QHostAddress::AnyIPv4 instead.

This function returns true if successful; otherwise it returns false and sets the socket error accordingly.

Note: Joining IPv6 multicast groups without an interface selection is not supported in all operating systems. Consider using the overload where the interface is specified.

This function was introduced in  Qt 4.8.

See also leaveMulticastGroup().
*/
func (this *QUdpSocket) JoinMulticastGroup1(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg1 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddressRK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &)

/*
Leaves the multicast group specified by groupAddress on the default interface chosen by the operating system. The socket must be in BoundState, otherwise an error occurs.

This function returns true if successful; otherwise it returns false and sets the socket error accordingly.

Note: This function should be called with the same arguments as were passed to joinMulticastGroup().

This function was introduced in  Qt 4.8.

See also joinMulticastGroup().
*/
func (this *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:68
// index:1
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &, const QNetworkInterface &)

/*
Leaves the multicast group specified by groupAddress on the default interface chosen by the operating system. The socket must be in BoundState, otherwise an error occurs.

This function returns true if successful; otherwise it returns false and sets the socket error accordingly.

Note: This function should be called with the same arguments as were passed to joinMulticastGroup().

This function was introduced in  Qt 4.8.

See also joinMulticastGroup().
*/
func (this *QUdpSocket) LeaveMulticastGroup1(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg1 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddressRK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkInterface multicastInterface() const

/*
Returns the interface for the outgoing interface for multicast datagrams. This corresponds to the IP_MULTICAST_IF socket option for IPv4 sockets and the IPV6_MULTICAST_IF socket option for IPv6 sockets. If no interface has been previously set, this function returns an invalid QNetworkInterface. The socket must be in BoundState, otherwise an invalid QNetworkInterface is returned.

This function was introduced in  Qt 4.8.

See also setMulticastInterface().
*/
func (this *QUdpSocket) MulticastInterface() *QNetworkInterface /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket18multicastInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMulticastInterface(const QNetworkInterface &)

/*
Sets the outgoing interface for multicast datagrams to the interface iface. This corresponds to the IP_MULTICAST_IF socket option for IPv4 sockets and the IPV6_MULTICAST_IF socket option for IPv6 sockets. The socket must be in BoundState, otherwise this function does nothing.

This function was introduced in  Qt 4.8.

See also multicastInterface(), joinMulticastGroup(), and leaveMulticastGroup().
*/
func (this *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg0 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket21setMulticastInterfaceERK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasPendingDatagrams() const

/*
Returns true if at least one datagram is waiting to be read; otherwise returns false.

See also pendingDatagramSize() and readDatagram().
*/
func (this *QUdpSocket) HasPendingDatagrams() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket19hasPendingDatagramsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pendingDatagramSize() const

/*
Returns the size of the first pending UDP datagram. If there is no datagram available, this function returns -1.

See also hasPendingDatagrams() and readDatagram().
*/
func (this *QUdpSocket) PendingDatagramSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket19pendingDatagramSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram receiveDatagram(qint64)

/*
Receives a datagram no larger than maxSize bytes and returns it in the QNetworkDatagram object, along with the sender's host address and port. If possible, this function will also try to determine the datagram's destination address, port, and the number of hop counts at reception time.

On failure, returns a QNetworkDatagram that reports not valid.

If maxSize is too small, the rest of the datagram will be lost. If maxSize is 0, the datagram will be discarded. If maxSize is -1 (the default), this function will attempt to read the entire datagram.

See also writeDatagram(), hasPendingDatagrams(), and pendingDatagramSize().
*/
func (this *QUdpSocket) ReceiveDatagram(maxSize int64) *QNetworkDatagram /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket15receiveDatagramEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram receiveDatagram(qint64)

/*
Receives a datagram no larger than maxSize bytes and returns it in the QNetworkDatagram object, along with the sender's host address and port. If possible, this function will also try to determine the datagram's destination address, port, and the number of hop counts at reception time.

On failure, returns a QNetworkDatagram that reports not valid.

If maxSize is too small, the rest of the datagram will be lost. If maxSize is 0, the datagram will be discarded. If maxSize is -1 (the default), this function will attempt to read the entire datagram.

See also writeDatagram(), hasPendingDatagrams(), and pendingDatagramSize().
*/
func (this *QUdpSocket) ReceiveDatagramp() *QNetworkDatagram /*123*/ {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long, LongLong
	maxSize := int64(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket15receiveDatagramEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)

/*
Receives a datagram no larger than maxSize bytes and stores it in data. The sender's host address and port is stored in *address and *port (unless the pointers are 0).

Returns the size of the datagram on success; otherwise returns -1.

If maxSize is too small, the rest of the datagram will be lost. To avoid loss of data, call pendingDatagramSize() to determine the size of the pending datagram before attempting to read it. If maxSize is 0, the datagram will be discarded.

See also writeDatagram(), hasPendingDatagrams(), and pendingDatagramSize().
*/
func (this *QUdpSocket) ReadDatagram(data string, maxlen int64, host QHostAddress_ITF /*777 QHostAddress **/, port unsafe.Pointer /*666*/) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)

/*
Receives a datagram no larger than maxSize bytes and stores it in data. The sender's host address and port is stored in *address and *port (unless the pointers are 0).

Returns the size of the datagram on success; otherwise returns -1.

If maxSize is too small, the rest of the datagram will be lost. To avoid loss of data, call pendingDatagramSize() to determine the size of the pending datagram before attempting to read it. If maxSize is 0, the datagram will be discarded.

See also writeDatagram(), hasPendingDatagrams(), and pendingDatagramSize().
*/
func (this *QUdpSocket) ReadDatagramp(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	// arg: 2, QHostAddress *=Pointer, QHostAddress=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, quint16 *=Pointer, quint16=Typedef, unsigned short, UShort
	port := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)

/*
Receives a datagram no larger than maxSize bytes and stores it in data. The sender's host address and port is stored in *address and *port (unless the pointers are 0).

Returns the size of the datagram on success; otherwise returns -1.

If maxSize is too small, the rest of the datagram will be lost. To avoid loss of data, call pendingDatagramSize() to determine the size of the pending datagram before attempting to read it. If maxSize is 0, the datagram will be discarded.

See also writeDatagram(), hasPendingDatagrams(), and pendingDatagramSize().
*/
func (this *QUdpSocket) ReadDatagramp1(data string, maxlen int64, host QHostAddress_ITF /*777 QHostAddress **/) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	// arg: 3, quint16 *=Pointer, quint16=Typedef, unsigned short, UShort
	port := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QNetworkDatagram &)

/*
Sends the datagram at data of size size to the host address address at port port. Returns the number of bytes sent on success; otherwise returns -1.

Datagrams are always written as one block. The maximum size of a datagram is highly platform-dependent, but can be as low as 8192 bytes. If the datagram is too large, this function will return -1 and error() will return DatagramTooLargeError.

Sending datagrams larger than 512 bytes is in general disadvised, as even if they are sent successfully, they are likely to be fragmented by the IP layer before arriving at their final destination.

Warning: Calling this function on a connected UDP socket may result in an error and no packet being sent. If you are using a connected socket, use write() to send datagrams.

See also readDatagram() and write().
*/
func (this *QUdpSocket) WriteDatagram(datagram QNetworkDatagram_ITF) int64 {
	var convArg0 unsafe.Pointer
	if datagram != nil && datagram.QNetworkDatagram_PTR() != nil {
		convArg0 = datagram.QNetworkDatagram_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK16QNetworkDatagram", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:81
// index:1
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const char *, qint64, const QHostAddress &, quint16)

/*
Sends the datagram at data of size size to the host address address at port port. Returns the number of bytes sent on success; otherwise returns -1.

Datagrams are always written as one block. The maximum size of a datagram is highly platform-dependent, but can be as low as 8192 bytes. If the datagram is too large, this function will return -1 and error() will return DatagramTooLargeError.

Sending datagrams larger than 512 bytes is in general disadvised, as even if they are sent successfully, they are likely to be fragmented by the IP layer before arriving at their final destination.

Warning: Calling this function on a connected UDP socket may result in an error and no packet being sent. If you are using a connected socket, use write() to send datagrams.

See also readDatagram() and write().
*/
func (this *QUdpSocket) WriteDatagram1(data string, len_ int64, host QHostAddress_ITF, port uint16) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramEPKcxRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:82
// index:2
// Public inline Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QByteArray &, const QHostAddress &, quint16)

/*
Sends the datagram at data of size size to the host address address at port port. Returns the number of bytes sent on success; otherwise returns -1.

Datagrams are always written as one block. The maximum size of a datagram is highly platform-dependent, but can be as low as 8192 bytes. If the datagram is too large, this function will return -1 and error() will return DatagramTooLargeError.

Sending datagrams larger than 512 bytes is in general disadvised, as even if they are sent successfully, they are likely to be fragmented by the IP layer before arriving at their final destination.

Warning: Calling this function on a connected UDP socket may result in an error and no packet being sent. If you are using a connected socket, use write() to send datagrams.

See also readDatagram() and write().
*/
func (this *QUdpSocket) WriteDatagram2(datagram qtcore.QByteArray_ITF, host QHostAddress_ITF, port uint16) int64 {
	var convArg0 unsafe.Pointer
	if datagram != nil && datagram.QByteArray_PTR() != nil {
		convArg0 = datagram.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg1 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

//  body block end

//  keep block begin

func init_unused_11473() {
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
