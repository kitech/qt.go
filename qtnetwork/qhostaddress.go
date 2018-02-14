package qtnetwork

// /usr/include/qt/QtNetwork/qhostaddress.h
// #include <qhostaddress.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QHostAddress struct {
	*qtrt.CObject
}
type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (ptr *QHostAddress) QHostAddress_PTR() *QHostAddress { return ptr }

func (this *QHostAddress) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHostAddress) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(quint32)
func NewQHostAddress_1(ip4Addr uint) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2Ej", qtrt.FFI_TYPE_POINTER, ip4Addr)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(quint8 *)
func NewQHostAddress_2(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2EPh", qtrt.FFI_TYPE_POINTER, ip6Addr)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:97
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(const quint8 *)
func NewQHostAddress_3(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2EPKh", qtrt.FFI_TYPE_POINTER, ip6Addr)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:100
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(const QString &)
func NewQHostAddress_4(address string) *QHostAddress {
	var tmpArg0 = qtcore.NewQString_5(address)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:102
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QHostAddress(enum QHostAddress::SpecialAddress)
func NewQHostAddress_5(address int) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressC2ENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, address)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostAddress)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostaddress.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHostAddress()
func DeleteQHostAddress(this *QHostAddress) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QHostAddress & operator=(QHostAddress &&)
func (this *QHostAddress) Operator_equal(other unsafe.Pointer /*333*/) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:110
// index:1
// Public Visibility=Default Availability=Available
// [8] QHostAddress & operator=(const QHostAddress &)
func (this *QHostAddress) Operator_equal_1(other QHostAddress_ITF) *QHostAddress {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHostAddress_PTR() != nil {
		convArg0 = other.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:113
// index:2
// Public Visibility=Default Availability=Available
// [8] QHostAddress & operator=(const QString &)
func (this *QHostAddress) Operator_equal_2(address string) *QHostAddress {
	var tmpArg0 = qtcore.NewQString_5(address)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressaSERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:115
// index:3
// Public Visibility=Default Availability=Available
// [8] QHostAddress & operator=(enum QHostAddress::SpecialAddress)
func (this *QHostAddress) Operator_equal_3(address int) *QHostAddress {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddressaSENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHostAddress &)
func (this *QHostAddress) Swap(other QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHostAddress_PTR() != nil {
		convArg0 = other.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAddress(quint32)
func (this *QHostAddress) SetAddress(ip4Addr uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip4Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAddress(quint8 *)
func (this *QHostAddress) SetAddress_1(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip6Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:121
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setAddress(const quint8 *)
func (this *QHostAddress) SetAddress_2(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPKh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip6Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:124
// index:3
// Public Visibility=Default Availability=Available
// [1] bool setAddress(const QString &)
func (this *QHostAddress) SetAddress_3(address string) bool {
	var tmpArg0 = qtcore.NewQString_5(address)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:125
// index:4
// Public Visibility=Default Availability=Available
// [-2] void setAddress(enum QHostAddress::SpecialAddress)
func (this *QHostAddress) SetAddress_4(address int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::NetworkLayerProtocol protocol()
func (this *QHostAddress) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address()
func (this *QHostAddress) ToIPv4Address() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:129
// index:1
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address(_Bool *)
func (this *QHostAddress) ToIPv4Address_1(ok *bool) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] Q_IPV6ADDR toIPv6Address()
func (this *QHostAddress) ToIPv6Address() *QIPv6Address /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv6AddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIPv6AddressFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QHostAddress) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qhostaddress.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scopeId()
func (this *QHostAddress) ScopeId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress7scopeIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qhostaddress.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScopeId(const QString &)
func (this *QHostAddress) SetScopeId(id string) {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setScopeIdERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEqual(const QHostAddress &, QHostAddress::ConversionMode)
func (this *QHostAddress) IsEqual(address QHostAddress_ITF, mode int) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress7isEqualERKS_6QFlagsINS_18ConversionModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:138
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QHostAddress &)
func (this *QHostAddress) Operator_equal_equal(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddresseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:139
// index:1
// Public Visibility=Default Availability=Available
// [1] bool operator==(enum QHostAddress::SpecialAddress)
func (this *QHostAddress) Operator_equal_equal_1(address int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddresseqENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QHostAddress &)
func (this *QHostAddress) Operator_not_equal(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddressneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:142
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(enum QHostAddress::SpecialAddress)
func (this *QHostAddress) Operator_not_equal_1(address int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddressneENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QHostAddress) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QHostAddress) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInSubnet(const QHostAddress &, int)
func (this *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	var convArg0 unsafe.Pointer
	if subnet != nil && subnet.QHostAddress_PTR() != nil {
		convArg0 = subnet.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress10isInSubnetERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, netmask)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoopback()
func (this *QHostAddress) IsLoopback() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress10isLoopbackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMulticast()
func (this *QHostAddress) IsMulticast() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress11isMulticastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
