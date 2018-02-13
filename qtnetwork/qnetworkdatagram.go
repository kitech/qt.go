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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

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
func NewQNetworkDatagram_1(data *qtcore.QByteArray, destinationAddress *QHostAddress, port uint16) *QNetworkDatagram {
	var convArg0 = data.GetCthis()
	var convArg1 = destinationAddress.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagramC2ERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkDatagram)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QNetworkDatagram()
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
func (this *QNetworkDatagram) Swap(other *QNetworkDatagram) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QNetworkDatagram) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QNetworkDatagram) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QNetworkDatagram) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] uint interfaceIndex()
func (this *QNetworkDatagram) InterfaceIndex() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram14interfaceIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInterfaceIndex(uint)
func (this *QNetworkDatagram) SetInterfaceIndex(index uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram17setInterfaceIndexEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress senderAddress()
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
// [8] QHostAddress destinationAddress()
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
// [4] int senderPort()
func (this *QNetworkDatagram) SenderPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram10senderPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int destinationPort()
func (this *QNetworkDatagram) DestinationPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram15destinationPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSender(const QHostAddress &, quint16)
func (this *QNetworkDatagram) SetSender(address *QHostAddress, port uint16) {
	var convArg0 = address.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram9setSenderERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDestination(const QHostAddress &, quint16)
func (this *QNetworkDatagram) SetDestination(address *QHostAddress, port uint16) {
	var convArg0 = address.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram14setDestinationERK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int hopLimit()
func (this *QNetworkDatagram) HopLimit() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QNetworkDatagram8hopLimitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHopLimit(int)
func (this *QNetworkDatagram) SetHopLimit(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram11setHopLimitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray data()
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
func (this *QNetworkDatagram) SetData(data *qtcore.QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QNetworkDatagram7setDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkDatagram makeReply(const QByteArray &)
func (this *QNetworkDatagram) MakeReply(payload *qtcore.QByteArray) *QNetworkDatagram /*123*/ {
	var convArg0 = payload.GetCthis()
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
func (this *QNetworkDatagram) MakeReply_1(payload *qtcore.QByteArray) *QNetworkDatagram /*123*/ {
	var convArg0 = payload.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNO16QNetworkDatagram9makeReplyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
