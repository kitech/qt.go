package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkdatagram.h
// #include <qnetworkdatagram.h>
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

/*

 */
type QNetworkDatagram struct {
	*qtrt.CObject
}
type QNetworkDatagram_ITF interface {
	QNetworkDatagram_PTR() *QNetworkDatagram
}

func (ptr *QNetworkDatagram) QNetworkDatagram_PTR() *QNetworkDatagram { return ptr }

func (this *QNetworkDatagram) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkDatagram) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkDatagramFromPointer(cthis unsafe.Pointer) *QNetworkDatagram {
	return &QNetworkDatagram{&qtrt.CObject{cthis}}
}
func (*QNetworkDatagram) NewFromPointer(cthis unsafe.Pointer) *QNetworkDatagram {
	return NewQNetworkDatagramFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDatagram()

/*
Creates a QNetworkDatagram object with no payload data and undefined destination address.

The payload can be modified by using setData() and the destination address can be set with setDestination().

If the destination address is left undefined, QUdpSocket::writeDatagram() will attempt to send the datagram to the address last associated with, by using QUdpSocket::connectToHost().
*/
func (*QNetworkDatagram) NewForInherit() *QNetworkDatagram {
	return NewQNetworkDatagram()
}
func NewQNetworkDatagram() *QNetworkDatagram {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkDatagram)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDatagram(const QByteArray &, const QHostAddress &, quint16)

/*
Creates a QNetworkDatagram object with no payload data and undefined destination address.

The payload can be modified by using setData() and the destination address can be set with setDestination().

If the destination address is left undefined, QUdpSocket::writeDatagram() will attempt to send the datagram to the address last associated with, by using QUdpSocket::connectToHost().
*/
func (*QNetworkDatagram) NewForInherit1(data qtcore.QByteArray_ITF, destinationAddress QHostAddress_ITF, port uint16) *QNetworkDatagram {
	return NewQNetworkDatagram1(data, destinationAddress, port)
}
func NewQNetworkDatagram1(data qtcore.QByteArray_ITF, destinationAddress QHostAddress_ITF, port uint16) *QNetworkDatagram {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if destinationAddress != nil && destinationAddress.QHostAddress_PTR() != nil {
		convArg1 = destinationAddress.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramC2ERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkDatagram)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDatagram(const QByteArray &, const QHostAddress &, quint16)

/*
Creates a QNetworkDatagram object with no payload data and undefined destination address.

The payload can be modified by using setData() and the destination address can be set with setDestination().

If the destination address is left undefined, QUdpSocket::writeDatagram() will attempt to send the datagram to the address last associated with, by using QUdpSocket::connectToHost().
*/
func (*QNetworkDatagram) NewForInherit1p(data qtcore.QByteArray_ITF) *QNetworkDatagram {
	return NewQNetworkDatagram1p(data)
}
func NewQNetworkDatagram1p(data qtcore.QByteArray_ITF) *QNetworkDatagram {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QHostAddress &=LValueReference, QHostAddress=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramC2ERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkDatagram)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDatagram(const QByteArray &, const QHostAddress &, quint16)

/*
Creates a QNetworkDatagram object with no payload data and undefined destination address.

The payload can be modified by using setData() and the destination address can be set with setDestination().

If the destination address is left undefined, QUdpSocket::writeDatagram() will attempt to send the datagram to the address last associated with, by using QUdpSocket::connectToHost().
*/
func (*QNetworkDatagram) NewForInherit1p1(data qtcore.QByteArray_ITF, destinationAddress QHostAddress_ITF) *QNetworkDatagram {
	return NewQNetworkDatagram1p1(data, destinationAddress)
}
func NewQNetworkDatagram1p1(data qtcore.QByteArray_ITF, destinationAddress QHostAddress_ITF) *QNetworkDatagram {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if destinationAddress != nil && destinationAddress.QHostAddress_PTR() != nil {
		convArg1 = destinationAddress.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramC2ERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkDatagram)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram & operator=(const QNetworkDatagram &)

/*

 */
func (this *QNetworkDatagram) Operator_equal(other QNetworkDatagram_ITF) *QNetworkDatagram {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkDatagram_PTR() != nil {
		convArg0 = other.QNetworkDatagram_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:67
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QNetworkDatagram & operator=(QNetworkDatagram &&)

/*

 */
func (this *QNetworkDatagram) Operator_equal1(other unsafe.Pointer /*333*/) *QNetworkDatagram {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QNetworkDatagram()

/*

 */
func DeleteQNetworkDatagram(this *QNetworkDatagram) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkDatagram &)

/*
Swaps this instance with other.
*/
func (this *QNetworkDatagram) Swap(other QNetworkDatagram_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkDatagram_PTR() != nil {
		convArg0 = other.QNetworkDatagram_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the payload data and metadata in this QNetworkDatagram object, resetting them to their default values.
*/
func (this *QNetworkDatagram) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this QNetworkDatagram object is valid. A valid QNetworkDatagram object contains at least one sender or receiver address. Valid datagrams can contain empty payloads.
*/
func (this *QNetworkDatagram) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this QNetworkDatagram object is null. This function is the opposite of isValid().
*/
func (this *QNetworkDatagram) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] uint interfaceIndex() const

/*
Returns the interface index this datagram is associated with. The interface index is a positive number that uniquely identifies the network interface in the operating system. This number matches the value returned by QNetworkInterface::index() for the interface.

If this datagram was received from the network, this is the index of the interface that the packet was received from. If this is an outgoing datagram, this is the index of the interface that the datagram should be sent on.

A value of 0 indicates that the interface index is unknown.

See also setInterfaceIndex().
*/
func (this *QNetworkDatagram) InterfaceIndex() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram14interfaceIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInterfaceIndex(uint)

/*
Sets the interface index this datagram is associated with to index. The interface index is a positive number that uniquely identifies the network interface in the operating system. This number matches the value returned by QNetworkInterface::index() for the interface.

It is usually not necessary to call this function on datagrams received from the network.

If this is an outgoing packet, this is the index of the interface the datagram should be sent on. A value of 0 indicates that the operating system should choose the interface based on other factors.

Note that the interface index can also be set with QHostAddress::setScopeId() for IPv6 destination addresses and then with setDestination(). If the scope ID set in the destination address and index are different and neither is zero, it is undefined which interface the operating system will send the datagram on.

See also interfaceIndex() and setInterfaceIndex().
*/
func (this *QNetworkDatagram) SetInterfaceIndex(index uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram17setInterfaceIndexEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress senderAddress() const

/*
Returns the sender address associated with this datagram. For a datagram received from the network, it is the address of the peer node that sent the datagram. For an outgoing datagrams, it is the local address to be used when sending.

If no sender address was set on this datagram, the returned object will report true to QHostAddress::isNull().

See also destinationAddress(), senderPort(), and setSender().
*/
func (this *QNetworkDatagram) SenderAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram13senderAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress destinationAddress() const

/*
Returns the destination address associated with this datagram. For a datagram received from the network, it is the address the peer node sent the datagram to, which can either be a local address of this machine or a multicast or broadcast address. For an outgoing datagrams, it is the address the datagram should be sent to.

If no destination address was set on this datagram, the returned object will report true to QHostAddress::isNull().

See also senderAddress(), destinationPort(), and setDestination().
*/
func (this *QNetworkDatagram) DestinationAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram18destinationAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int senderPort() const

/*
Returns the port number of the sender associated with this datagram. For a datagram received from the network, it is the port number that the peer node sent the datagram from. For an outgoing datagram, it is the local port the datagram should be sent from.

If no sender address was associated with this datagram, this function returns -1.

See also senderAddress(), destinationPort(), and setSender().
*/
func (this *QNetworkDatagram) SenderPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram10senderPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int destinationPort() const

/*
Returns the port number of the destination associated with this datagram. For a datagram received from the network, it is the local port number that the peer node sent the datagram to. For an outgoing datagram, it is the peer port the datagram should be sent to.

If no destination address was associated with this datagram, this function returns -1.

See also destinationAddress(), senderPort(), and setDestination().
*/
func (this *QNetworkDatagram) DestinationPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram15destinationPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSender(const QHostAddress &, quint16)

/*
Sets the sender address associated with this datagram to be the address address and port number port. The sender address and port numbers are usually set by QUdpSocket upon reception, so there's no need to call this function on a received datagram.

For outgoing datagrams, this function can be used to set the address the datagram should carry. The address address must usually be one of the local addresses assigned to this machine, which can be obtained using QNetworkInterface. If left unset, the operating system will choose the most appropriate address to use given the destination in question.

The port number port must be the port number associated with the socket, if there is one. The value of 0 can be used to indicate that the operating system should choose the port number.

See also QUdpSocket::writeDatagram(), senderAddress(), senderPort(), and setDestination().
*/
func (this *QNetworkDatagram) SetSender(address QHostAddress_ITF, port uint16) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram9setSenderERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSender(const QHostAddress &, quint16)

/*
Sets the sender address associated with this datagram to be the address address and port number port. The sender address and port numbers are usually set by QUdpSocket upon reception, so there's no need to call this function on a received datagram.

For outgoing datagrams, this function can be used to set the address the datagram should carry. The address address must usually be one of the local addresses assigned to this machine, which can be obtained using QNetworkInterface. If left unset, the operating system will choose the most appropriate address to use given the destination in question.

The port number port must be the port number associated with the socket, if there is one. The value of 0 can be used to indicate that the operating system should choose the port number.

See also QUdpSocket::writeDatagram(), senderAddress(), senderPort(), and setDestination().
*/
func (this *QNetworkDatagram) SetSenderp(address QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram9setSenderERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDestination(const QHostAddress &, quint16)

/*
Sets the destination address associated with this datagram to be the address address and port number port. The destination address and port numbers are usually set by QUdpSocket upon reception, so there's no need to call this function on a received datagram.

For outgoing datagrams, this function can be used to set the address the datagram should be sent to. It can be the unicast address used to communicate with the peer or a broadcast or multicast address to send to a group of devices.

See also QUdpSocket::writeDatagram(), destinationAddress(), destinationPort(), and setSender().
*/
func (this *QNetworkDatagram) SetDestination(address QHostAddress_ITF, port uint16) {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram14setDestinationERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int hopLimit() const

/*
Returns the hop count limit associated with this datagram. The hop count limit is the number of nodes that are allowed to forward the IP packet before it expires and an error is sent back to the sender of the datagram. In IPv4, this value is usually known as "time to live" (TTL).

If this datagram was received from the network, this is the remaining hop count of the datagram after reception and was decremented by 1 by each node that forwarded the packet. A value of -1 indicates that the hop limit count not be obtained.

If this is an outgoing datagram, this is the value to be set in the IP header upon sending. A value of -1 indicates the operating system should choose the value.

See also setHopLimit().
*/
func (this *QNetworkDatagram) HopLimit() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram8hopLimitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHopLimit(int)

/*
Sets the hop count limit associated with this datagram to count. The hop count limit is the number of nodes that are allowed to forward the IP packet before it expires and an error is sent back to the sender of the datagram. In IPv4, this value is usually known as "time to live" (TTL).

It is usually not necessary to call this function on datagrams received from the network.

If this is an outgoing packet, this is the value to be set in the IP header upon sending. The valid range for the value is 1 to 255. This function also accepts a value of -1 to indicate that the operating system should choose the value.

See also hopLimit().
*/
func (this *QNetworkDatagram) SetHopLimit(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram11setHopLimitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray data() const

/*
Returns the data payload of this datagram. For a datagram received from the network, it contains the payload of the datagram. For an outgoing datagram, it is the datagram to be sent.

Note that datagrams can be transmitted with no data, so the returned QByteArray may be empty.

See also setData().
*/
func (this *QNetworkDatagram) Data() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QByteArray &)

/*
Sets the data payload of this datagram to data. It is usually not necessary to call this function on received datagrams. For outgoing datagrams, this function sets the data to be sent on the network.

Since datagrams can empty, an empty QByteArray is a valid value for data.

See also data().
*/
func (this *QNetworkDatagram) SetData(data qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram7setDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkDatagram makeReply(const QByteArray &) const

/*

 */
func (this *QNetworkDatagram) MakeReply(payload qtcore.QByteArray_ITF) *QNetworkDatagram /*123*/ {
	var convArg0 unsafe.Pointer
	if payload != nil && payload.QByteArray_PTR() != nil {
		convArg0 = payload.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNKR16QNetworkDatagram9makeReplyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:97
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QNetworkDatagram makeReply(const QByteArray &)

/*

 */
func (this *QNetworkDatagram) MakeReply1(payload qtcore.QByteArray_ITF) *QNetworkDatagram /*123*/ {
	var convArg0 unsafe.Pointer
	if payload != nil && payload.QByteArray_PTR() != nil {
		convArg0 = payload.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNO16QNetworkDatagram9makeReplyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_11443() {
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
