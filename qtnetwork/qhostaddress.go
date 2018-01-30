package qtnetwork

// /usr/include/qt/QtNetwork/qhostaddress.h
// #include <qhostaddress.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QHostAddress struct {
	*qtrt.CObject
}

func (this *QHostAddress) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHostAddress) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQHostAddressFromPointer(cthis unsafe.Pointer) *QHostAddress {
	return &QHostAddress{&qtrt.CObject{cthis}}
}
func (*QHostAddress) NewFromPointer(cthis unsafe.Pointer) *QHostAddress {
	return NewQHostAddressFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress()
func NewQHostAddress() *QHostAddress {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(quint32)
func NewQHostAddress_1(ip4Addr uint) *QHostAddress {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2Ej", ffiqt.FFI_TYPE_POINTER, ip4Addr)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(quint8 *)
func NewQHostAddress_2(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2EPh", ffiqt.FFI_TYPE_POINTER, &ip6Addr)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(const quint8 *)
func NewQHostAddress_3(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2EPKh", ffiqt.FFI_TYPE_POINTER, &ip6Addr)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:100
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(const QString &)
func NewQHostAddress_4(address *qtcore.QString) *QHostAddress {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:102
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(enum QHostAddress::SpecialAddress)
func NewQHostAddress_5(address int) *QHostAddress {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressC2ENS_14SpecialAddressE", ffiqt.FFI_TYPE_POINTER, address)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHostAddress()
func DeleteQHostAddress(*QHostAddress) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddressD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHostAddress &)
func (this *QHostAddress) Swap(other *QHostAddress) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAddress(quint32)
func (this *QHostAddress) SetAddress(ip4Addr uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ip4Addr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAddress(quint8 *)
func (this *QHostAddress) SetAddress_1(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ip6Addr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:121
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setAddress(const quint8 *)
func (this *QHostAddress) SetAddress_2(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPKh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ip6Addr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:124
// index:3
// Public Visibility=Default Availability=Available
// [1] bool setAddress(const QString &)
func (this *QHostAddress) SetAddress_3(address *qtcore.QString) bool {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setAddressERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:125
// index:4
// Public Visibility=Default Availability=Available
// [-2] void setAddress(enum QHostAddress::SpecialAddress)
func (this *QHostAddress) SetAddress_4(address int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setAddressENS_14SpecialAddressE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), address)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::NetworkLayerProtocol protocol()
func (this *QHostAddress) Protocol() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress8protocolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address()
func (this *QHostAddress) ToIPv4Address() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:129
// index:1
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address(_Bool *)
func (this *QHostAddress) ToIPv4Address_1(ok unsafe.Pointer /*666*/) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] Q_IPV6ADDR toIPv6Address()
func (this *QHostAddress) ToIPv6Address() *QIPv6Address /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv6AddressEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIPv6AddressFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QHostAddress) ToString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scopeId()
func (this *QHostAddress) ScopeId() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress7scopeIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScopeId(const QString &)
func (this *QHostAddress) SetScopeId(id *qtcore.QString) {
	var convArg0 = id.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress10setScopeIdERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QHostAddress) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QHostAddress) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QHostAddress5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInSubnet(const QHostAddress &, int)
func (this *QHostAddress) IsInSubnet(subnet *QHostAddress, netmask int) bool {
	var convArg0 = subnet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress10isInSubnetERKS_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, netmask)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoopback()
func (this *QHostAddress) IsLoopback() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress10isLoopbackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMulticast()
func (this *QHostAddress) IsMulticast() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QHostAddress11isMulticastEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QHostAddress__SpecialAddress = int

const QHostAddress__Null QHostAddress__SpecialAddress = 0
const QHostAddress__Broadcast QHostAddress__SpecialAddress = 1
const QHostAddress__LocalHost QHostAddress__SpecialAddress = 2
const QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = 3
const QHostAddress__Any QHostAddress__SpecialAddress = 4
const QHostAddress__AnyIPv6 QHostAddress__SpecialAddress = 5
const QHostAddress__AnyIPv4 QHostAddress__SpecialAddress = 6

type QHostAddress__ConversionModeFlag = int

const QHostAddress__ConvertV4MappedToIPv4 QHostAddress__ConversionModeFlag = 1
const QHostAddress__ConvertV4CompatToIPv4 QHostAddress__ConversionModeFlag = 2
const QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = 4
const QHostAddress__ConvertLocalHost QHostAddress__ConversionModeFlag = 8
const QHostAddress__TolerantConversion QHostAddress__ConversionModeFlag = 255
const QHostAddress__StrictConversion QHostAddress__ConversionModeFlag = 0

//  body block end
