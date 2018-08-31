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

/*

 */
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

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit() *QHostAddress {
	return NewQHostAddress()
}
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

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit_1(ip4Addr uint) *QHostAddress {
	return NewQHostAddress_1(ip4Addr)
}
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

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit_2(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	return NewQHostAddress_2(ip6Addr)
}
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

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit_3(ip6Addr unsafe.Pointer /*666*/) *QHostAddress {
	return NewQHostAddress_3(ip6Addr)
}
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

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit_4(address string) *QHostAddress {
	return NewQHostAddress_4(address)
}
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
// [-2] void QHostAddress(QHostAddress::SpecialAddress)

/*
Constructs a null host address object, i.e. an address which is not valid for any host or interface.

See also clear().
*/
func (*QHostAddress) NewForInherit_5(address int) *QHostAddress {
	return NewQHostAddress_5(address)
}
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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
// [8] QHostAddress & operator=(QHostAddress::SpecialAddress)

/*

 */
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

/*
Swaps this host address with other. This operation is very fast and never fails.

This function was introduced in  Qt 5.6.
*/
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

/*
Set the IPv4 address specified by ip4Addr.
*/
func (this *QHostAddress) SetAddress(ip4Addr uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip4Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAddress(quint8 *)

/*
Set the IPv4 address specified by ip4Addr.
*/
func (this *QHostAddress) SetAddress_1(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip6Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:121
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setAddress(const quint8 *)

/*
Set the IPv4 address specified by ip4Addr.
*/
func (this *QHostAddress) SetAddress_2(ip6Addr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressEPKh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ip6Addr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:124
// index:3
// Public Visibility=Default Availability=Available
// [1] bool setAddress(const QString &)

/*
Set the IPv4 address specified by ip4Addr.
*/
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
// [-2] void setAddress(QHostAddress::SpecialAddress)

/*
Set the IPv4 address specified by ip4Addr.
*/
func (this *QHostAddress) SetAddress_4(address int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setAddressENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSocket::NetworkLayerProtocol protocol() const

/*
Returns the network layer protocol of the host address.
*/
func (this *QHostAddress) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address() const

/*
Returns the IPv4 address as a number.

For example, if the address is 127.0.0.1, the returned value is 2130706433 (i.e. 0x7f000001).

This value is valid if the protocol() is IPv4Protocol, or if the protocol is IPv6Protocol, and the IPv6 address is an IPv4 mapped address. (RFC4291)

See also toString().
*/
func (this *QHostAddress) ToIPv4Address() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:129
// index:1
// Public Visibility=Default Availability=Available
// [4] quint32 toIPv4Address(bool *) const

/*
Returns the IPv4 address as a number.

For example, if the address is 127.0.0.1, the returned value is 2130706433 (i.e. 0x7f000001).

This value is valid if the protocol() is IPv4Protocol, or if the protocol is IPv6Protocol, and the IPv6 address is an IPv4 mapped address. (RFC4291)

See also toString().
*/
func (this *QHostAddress) ToIPv4Address_1(ok *bool) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv4AddressEPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] Q_IPV6ADDR toIPv6Address() const

/*
Returns the IPv6 address as a Q_IPV6ADDR structure. The structure consists of 16 unsigned characters.


  Q_IPV6ADDR addr = hostAddr.toIPv6Address();
  // addr contains 16 unsigned characters

  for (int i = 0; i < 16; ++i) {
      // process addr[i]
  }



This value is valid if the protocol() is IPv6Protocol. If the protocol is IPv4Protocol, then the address is returned an an IPv4 mapped IPv6 address. (RFC4291)

See also toString().
*/
func (this *QHostAddress) ToIPv6Address() *QIPv6Address /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress13toIPv6AddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIPv6AddressFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns the address as a string.

For example, if the address is the IPv4 address 127.0.0.1, the returned string is "127.0.0.1". For IPv6 the string format will follow the RFC5952 recommendation. For QHostAddress::Any, its IPv4 address will be returned ("0.0.0.0")

See also toIPv4Address().
*/
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
// [8] QString scopeId() const

/*
Returns the scope ID of an IPv6 address. For IPv4 addresses, or if the address does not contain a scope ID, an empty QString is returned.

The IPv6 scope ID specifies the scope of reachability for non-global IPv6 addresses, limiting the area in which the address can be used. All IPv6 addresses are associated with such a reachability scope. The scope ID is used to disambiguate addresses that are not guaranteed to be globally unique.

IPv6 specifies the following four levels of reachability:


Node-local: Addresses that are only used for communicating with services on the same interface (e.g., the loopback interface "::1").
Link-local: Addresses that are local to the network interface (link). There is always one link-local address for each IPv6 interface on your host. Link-local addresses ("fe80...") are generated from the MAC address of the local network adaptor, and are not guaranteed to be unique.
Global: For globally routable addresses, such as public servers on the Internet.


When using a link-local or site-local address for IPv6 connections, you must specify the scope ID. The scope ID for a link-local address is usually the same as the interface name (e.g., "eth0", "en1") or number (e.g., "1", "2").

This function was introduced in  Qt 4.1.

See also setScopeId(), QNetworkInterface, and QNetworkInterface::interfaceFromName.
*/
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

/*
Sets the IPv6 scope ID of the address to id. If the address protocol is not IPv6, this function does nothing. The scope ID may be set as an interface name (such as "eth0" or "en1") or as an integer representing the interface index. If id is an interface name, QtNetwork will convert to an interface index using QNetworkInterface::interfaceIndexFromName() before calling the operating system networking functions.

This function was introduced in  Qt 4.1.

See also scopeId(), QNetworkInterface, and QNetworkInterface::interfaceFromName.
*/
func (this *QHostAddress) SetScopeId(id string) {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress10setScopeIdERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEqual(const QHostAddress &, QHostAddress::ConversionMode) const

/*
Returns true if this host address is the same as the other address given; otherwise returns false.

The parameter mode controls which conversions are preformed between addresses of differing protocols. If no mode is given, TolerantConversion is performed by default.

This function was introduced in  Qt 5.8.

See also ConversionMode and operator==().
*/
func (this *QHostAddress) IsEqual(address QHostAddress_ITF, mode int) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress7isEqualERKS_6QFlagsINS_18ConversionModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEqual(const QHostAddress &, QHostAddress::ConversionMode) const

/*
Returns true if this host address is the same as the other address given; otherwise returns false.

The parameter mode controls which conversions are preformed between addresses of differing protocols. If no mode is given, TolerantConversion is performed by default.

This function was introduced in  Qt 5.8.

See also ConversionMode and operator==().
*/
func (this *QHostAddress) IsEqual__(address QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 1, QHostAddress::ConversionMode=Typedef, QHostAddress::ConversionMode=Typedef, QFlags<QHostAddress::ConversionModeFlag>, Unexposed
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress7isEqualERKS_6QFlagsINS_18ConversionModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:138
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QHostAddress &) const

/*

 */
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
// [1] bool operator==(QHostAddress::SpecialAddress) const

/*

 */
func (this *QHostAddress) Operator_equal_equal_1(address int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddresseqENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QHostAddress &) const

/*

 */
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
// [1] bool operator!=(QHostAddress::SpecialAddress) const

/*

 */
func (this *QHostAddress) Operator_not_equal_1(address int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddressneENS_14SpecialAddressE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), address)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this host address is not valid for any host or interface.

The default constructor creates a null address.

See also QHostAddress::Null.
*/
func (this *QHostAddress) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Sets the host address to null.

See also QHostAddress::Null.
*/
func (this *QHostAddress) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHostAddress5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInSubnet(const QHostAddress &, int) const

/*
Returns true if this IP is in the subnet described by the network prefix subnet and netmask netmask.

An IP is considered to belong to a subnet if it is contained between the lowest and the highest address in that subnet. In the case of IP version 4, the lowest address is the network address, while the highest address is the broadcast address.

The subnet argument does not have to be the actual network address (the lowest address in the subnet). It can be any valid IP belonging to that subnet. In particular, if it is equal to the IP address held by this object, this function will always return true (provided the netmask is a valid value).

This function was introduced in  Qt 4.5.

See also parseSubnet().
*/
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
// [1] bool isLoopback() const

/*
returns true if the address is the IPv6 loopback address, or any of the IPv4 loopback addresses.

This function was introduced in  Qt 5.0.
*/
func (this *QHostAddress) IsLoopback() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress10isLoopbackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMulticast() const

/*
Returns true if the address is an IPv4 or IPv6 multicast address, false otherwise.

This function was introduced in  Qt 5.6.
*/
func (this *QHostAddress) IsMulticast() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QHostAddress11isMulticastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*

 */
type QHostAddress__SpecialAddress = int

// The null address object. Equivalent to QHostAddress(). See also QHostAddress::isNull().
const QHostAddress__Null QHostAddress__SpecialAddress = 0

//
const QHostAddress__Broadcast QHostAddress__SpecialAddress = 1

//
const QHostAddress__LocalHost QHostAddress__SpecialAddress = 2

//
const QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = 3

//
const QHostAddress__Any QHostAddress__SpecialAddress = 4

//
const QHostAddress__AnyIPv6 QHostAddress__SpecialAddress = 5

//
const QHostAddress__AnyIPv4 QHostAddress__SpecialAddress = 6

func (this *QHostAddress) SpecialAddressItemName(val int) string {
	switch val {
	case QHostAddress__Null: // 0
		return "Null"
	case QHostAddress__Broadcast: // 1
		return "Broadcast"
	case QHostAddress__LocalHost: // 2
		return "LocalHost"
	case QHostAddress__LocalHostIPv6: // 3
		return "LocalHostIPv6"
	case QHostAddress__Any: // 4
		return "Any"
	case QHostAddress__AnyIPv6: // 5
		return "AnyIPv6"
	case QHostAddress__AnyIPv4: // 6
		return "AnyIPv4"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QHostAddress_SpecialAddressItemName(val int) string {
	var nilthis *QHostAddress
	return nilthis.SpecialAddressItemName(val)
}

/*


 */
type QHostAddress__ConversionModeFlag = int

//
const QHostAddress__ConvertV4MappedToIPv4 QHostAddress__ConversionModeFlag = 1

//
const QHostAddress__ConvertV4CompatToIPv4 QHostAddress__ConversionModeFlag = 2

//
const QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = 4

//
const QHostAddress__ConvertLocalHost QHostAddress__ConversionModeFlag = 8

//
const QHostAddress__TolerantConversion QHostAddress__ConversionModeFlag = 255

//
const QHostAddress__StrictConversion QHostAddress__ConversionModeFlag = 0

func (this *QHostAddress) ConversionModeFlagItemName(val int) string {
	switch val {
	case QHostAddress__ConvertV4MappedToIPv4: // 1
		return "ConvertV4MappedToIPv4"
	case QHostAddress__ConvertV4CompatToIPv4: // 2
		return "ConvertV4CompatToIPv4"
	case QHostAddress__ConvertUnspecifiedAddress: // 4
		return "ConvertUnspecifiedAddress"
	case QHostAddress__ConvertLocalHost: // 8
		return "ConvertLocalHost"
	case QHostAddress__TolerantConversion: // 255
		return "TolerantConversion"
	case QHostAddress__StrictConversion: // 0
		return "StrictConversion"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QHostAddress_ConversionModeFlagItemName(val int) string {
	var nilthis *QHostAddress
	return nilthis.ConversionModeFlagItemName(val)
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
