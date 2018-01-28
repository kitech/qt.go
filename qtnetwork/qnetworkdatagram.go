package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkdatagram.h
// #include <qnetworkdatagram.h>
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
type QNetworkDatagram struct {
	*qtrt.CObject
}

func (this *QNetworkDatagram) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkDatagram) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkDatagramFromPointer(cthis unsafe.Pointer) *QNetworkDatagram {
	return &QNetworkDatagram{&qtrt.CObject{cthis}}
}
func (*QNetworkDatagram) NewFromPointer(cthis unsafe.Pointer) *QNetworkDatagram {
	return NewQNetworkDatagramFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:56
// index:0
// Public
// void QNetworkDatagram()
func NewQNetworkDatagram() *QNetworkDatagram {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagramC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:57
// index:1
// Public
// void QNetworkDatagram(const QByteArray &, const QHostAddress &, quint16)
func NewQNetworkDatagram_1(data *qtcore.QByteArray, destinationAddress *QHostAddress, port uint16) *QNetworkDatagram {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = data.GetCthis()
	var convArg1 = destinationAddress.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagramC2ERK10QByteArrayRK12QHostAddresst", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, port)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkDatagramFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:61
// index:0
// Public inline
// void ~QNetworkDatagram()
func DeleteQNetworkDatagram(*QNetworkDatagram) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagramD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:70
// index:0
// Public inline
// void swap(QNetworkDatagram &)
func (this *QNetworkDatagram) Swap(other *QNetworkDatagram) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:73
// index:0
// Public
// void clear()
func (this *QNetworkDatagram) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:74
// index:0
// Public
// bool isValid()
func (this *QNetworkDatagram) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:75
// index:0
// Public inline
// bool isNull()
func (this *QNetworkDatagram) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:78
// index:0
// Public
// uint interfaceIndex()
func (this *QNetworkDatagram) InterfaceIndex() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram14interfaceIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:79
// index:0
// Public
// void setInterfaceIndex(uint)
func (this *QNetworkDatagram) SetInterfaceIndex(index uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram17setInterfaceIndexEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:81
// index:0
// Public
// QHostAddress senderAddress()
func (this *QNetworkDatagram) SenderAddress() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram13senderAddressEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:82
// index:0
// Public
// QHostAddress destinationAddress()
func (this *QNetworkDatagram) DestinationAddress() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram18destinationAddressEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:83
// index:0
// Public
// int senderPort()
func (this *QNetworkDatagram) SenderPort() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram10senderPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:84
// index:0
// Public
// int destinationPort()
func (this *QNetworkDatagram) DestinationPort() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram15destinationPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:85
// index:0
// Public
// void setSender(const QHostAddress &, quint16)
func (this *QNetworkDatagram) SetSender(address *QHostAddress, port uint16) {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram9setSenderERK12QHostAddresst", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:86
// index:0
// Public
// void setDestination(const QHostAddress &, quint16)
func (this *QNetworkDatagram) SetDestination(address *QHostAddress, port uint16) {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram14setDestinationERK12QHostAddresst", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:88
// index:0
// Public
// int hopLimit()
func (this *QNetworkDatagram) HopLimit() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram8hopLimitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:89
// index:0
// Public
// void setHopLimit(int)
func (this *QNetworkDatagram) SetHopLimit(count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram11setHopLimitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:91
// index:0
// Public
// QByteArray data()
func (this *QNetworkDatagram) Data() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QNetworkDatagram4dataEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:92
// index:0
// Public
// void setData(const QByteArray &)
func (this *QNetworkDatagram) SetData(data *qtcore.QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QNetworkDatagram7setDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:95
// index:0
// Public inline
// QNetworkDatagram makeReply(const QByteArray &)
func (this *QNetworkDatagram) MakeReply(payload *qtcore.QByteArray) *QNetworkDatagram /*123*/ {
	var convArg0 = payload.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR16QNetworkDatagram9makeReplyERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:97
// index:1
// Public inline
// QNetworkDatagram makeReply(const QByteArray &)
func (this *QNetworkDatagram) MakeReply_1(payload *qtcore.QByteArray) *QNetworkDatagram /*123*/ {
	var convArg0 = payload.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNO16QNetworkDatagram9makeReplyERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
