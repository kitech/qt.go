package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkinterface.h
// #include <qnetworkinterface.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QNetworkInterface struct {
	*qtrt.CObject
}
type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (ptr *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface { return ptr }

func (this *QNetworkInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkInterfaceFromPointer(cthis unsafe.Pointer) *QNetworkInterface {
	return &QNetworkInterface{&qtrt.CObject{cthis}}
}
func (*QNetworkInterface) NewFromPointer(cthis unsafe.Pointer) *QNetworkInterface {
	return NewQNetworkInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkInterface()

/*
Constructs an empty network interface object.
*/
func (*QNetworkInterface) NewForInherit() *QNetworkInterface {
	return NewQNetworkInterface()
}
func NewQNetworkInterface() *QNetworkInterface {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkInterface)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkInterface & operator=(QNetworkInterface &&)

/*

 */
func (this *QNetworkInterface) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkInterface {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:109
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkInterface & operator=(const QNetworkInterface &)

/*

 */
func (this *QNetworkInterface) Operator_equal1(other QNetworkInterface_ITF) *QNetworkInterface {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkInterface_PTR() != nil {
		convArg0 = other.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkInterface()

/*

 */
func DeleteQNetworkInterface(this *QNetworkInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkInterface &)

/*
Swaps this network interface instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkInterface_PTR() != nil {
		convArg0 = other.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this QNetworkInterface object contains valid information about a network interface.
*/
func (this *QNetworkInterface) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int index() const

/*
Returns the interface system index, if known. This is an integer assigned by the operating system to identify this interface and it generally doesn't change. It matches the scope ID field in IPv6 addresses.

If the index isn't known, this function returns 0.

This function was introduced in  Qt 4.5.
*/
func (this *QNetworkInterface) Index() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface5indexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name of this network interface. On Unix systems, this is a string containing the type of the interface and optionally a sequence number, such as "eth0", "lo" or "pcn0". On Windows, it's an internal ID that cannot be changed by the user.
*/
func (this *QNetworkInterface) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString humanReadableName() const

/*
Returns the human-readable name of this network interface on Windows, such as "Local Area Connection", if the name could be determined. If it couldn't, this function returns the same as name(). The human-readable name is a name that the user can modify in the Windows Control Panel, so it may change during the execution of the program.

On Unix, this function currently always returns the same as name(), since Unix systems don't store a configuration for human-readable names.

This function was introduced in  Qt 4.5.
*/
func (this *QNetworkInterface) HumanReadableName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface17humanReadableNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkInterface::InterfaceFlags flags() const

/*
Returns the flags associated with this network interface.
*/
func (this *QNetworkInterface) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hardwareAddress() const

/*
Returns the low-level hardware address for this interface. On Ethernet interfaces, this will be a MAC address in string representation, separated by colons.

Other interface types may have other types of hardware addresses. Implementations should not depend on this function returning a valid MAC address.
*/
func (this *QNetworkInterface) HardwareAddress() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface15hardwareAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:123
// index:0
// Public static Visibility=Default Availability=Available
// [4] int interfaceIndexFromName(const QString &)

/*
Returns the index of the interface whose name is name or 0 if there is no interface with that name. This function should produce the same result as the following code, but will probably execute faster.


  QNetworkInterface::interfaceFromName(name).index()



This function was introduced in  Qt 5.7.

See also interfaceFromName(), interfaceNameFromIndex(), and QNetworkDatagram::interfaceIndex().
*/
func (this *QNetworkInterface) InterfaceIndexFromName(name string) int {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface22interfaceIndexFromNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QNetworkInterface_InterfaceIndexFromName(name string) int {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceIndexFromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [8] QNetworkInterface interfaceFromName(const QString &)

/*
Returns a QNetworkInterface object for the interface named name. If no such interface exists, this function returns an invalid QNetworkInterface object.

The string name may be either an actual interface name (such as "eth0" or "en1") or an interface index in string form ("1", "2", etc.).

See also name() and isValid().
*/
func (this *QNetworkInterface) InterfaceFromName(name string) *QNetworkInterface /*123*/ {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface17interfaceFromNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}
func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface /*123*/ {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceFromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [8] QNetworkInterface interfaceFromIndex(int)

/*
Returns a QNetworkInterface object for the interface whose internal ID is index. Network interfaces have a unique identifier called the "interface index" to distinguish it from other interfaces on the system. Often, this value is assigned progressively and interfaces being removed and then added again get a different value every time.

This index is also found in the IPv6 address' scope ID field.
*/
func (this *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface18interfaceFromIndexEi", qtrt.FFI_TYPE_POINTER, index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}
func QNetworkInterface_InterfaceFromIndex(index int) *QNetworkInterface /*123*/ {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceFromIndex(index)
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString interfaceNameFromIndex(int)

/*
Returns the name of the interface whose index is index or an empty string if there is no interface with that index. This function should produce the same result as the following code, but will probably execute faster.


  QNetworkInterface::interfaceFromIndex(index).name()



This function was introduced in  Qt 5.7.

See also interfaceFromIndex(), interfaceIndexFromName(), and QNetworkDatagram::interfaceIndex().
*/
func (this *QNetworkInterface) InterfaceNameFromIndex(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface22interfaceNameFromIndexEi", qtrt.FFI_TYPE_POINTER, index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QNetworkInterface_InterfaceNameFromIndex(index int) string {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceNameFromIndex(index)
	return rv
}

/*


 */
type QNetworkInterface__InterfaceFlag = int

//
const QNetworkInterface__IsUp QNetworkInterface__InterfaceFlag = 1

//
const QNetworkInterface__IsRunning QNetworkInterface__InterfaceFlag = 2

//
const QNetworkInterface__CanBroadcast QNetworkInterface__InterfaceFlag = 4

//
const QNetworkInterface__IsLoopBack QNetworkInterface__InterfaceFlag = 8

//
const QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = 16

//
const QNetworkInterface__CanMulticast QNetworkInterface__InterfaceFlag = 32

func (this *QNetworkInterface) InterfaceFlagItemName(val int) string {
	switch val {
	case QNetworkInterface__IsUp: // 1
		return "IsUp"
	case QNetworkInterface__IsRunning: // 2
		return "IsRunning"
	case QNetworkInterface__CanBroadcast: // 4
		return "CanBroadcast"
	case QNetworkInterface__IsLoopBack: // 8
		return "IsLoopBack"
	case QNetworkInterface__IsPointToPoint: // 16
		return "IsPointToPoint"
	case QNetworkInterface__CanMulticast: // 32
		return "CanMulticast"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkInterface_InterfaceFlagItemName(val int) string {
	var nilthis *QNetworkInterface
	return nilthis.InterfaceFlagItemName(val)
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
