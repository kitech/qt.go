package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkinterface.h
// #include <qnetworkinterface.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

type QNetworkInterface struct {
	*qtrt.CObject
}

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
func NewQNetworkInterface() *QNetworkInterface {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkInterface)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkInterface()
func DeleteQNetworkInterface(this *QNetworkInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkInterface &)
func (this *QNetworkInterface) Swap(other *QNetworkInterface) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QNetworkInterface) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int index()
func (this *QNetworkInterface) Index() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface5indexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:117
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QNetworkInterface) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QString humanReadableName()
func (this *QNetworkInterface) HumanReadableName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface17humanReadableNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkInterface::InterfaceFlags flags()
func (this *QNetworkInterface) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hardwareAddress()
func (this *QNetworkInterface) HardwareAddress() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkInterface15hardwareAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:123
// index:0
// Public static Visibility=Default Availability=Available
// [4] int interfaceIndexFromName(const QString &)
func (this *QNetworkInterface) InterfaceIndexFromName(name *qtcore.QString) int {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface22interfaceIndexFromNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QNetworkInterface_InterfaceIndexFromName(name *qtcore.QString) int {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceIndexFromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [8] QNetworkInterface interfaceFromName(const QString &)
func (this *QNetworkInterface) InterfaceFromName(name *qtcore.QString) *QNetworkInterface /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface17interfaceFromNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}
func QNetworkInterface_InterfaceFromName(name *qtcore.QString) *QNetworkInterface /*123*/ {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceFromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [8] QNetworkInterface interfaceFromIndex(int)
func (this *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface18interfaceFromIndexEi", qtrt.FFI_TYPE_POINTER, index)
	gopp.ErrPrint(err, rv)
	// return rv
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
func (this *QNetworkInterface) InterfaceNameFromIndex(index int) *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkInterface22interfaceNameFromIndexEi", qtrt.FFI_TYPE_POINTER, index)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QNetworkInterface_InterfaceNameFromIndex(index int) *qtcore.QString /*123*/ {
	var nilthis *QNetworkInterface
	rv := nilthis.InterfaceNameFromIndex(index)
	return rv
}

type QNetworkInterface__InterfaceFlag = int

const QNetworkInterface__IsUp QNetworkInterface__InterfaceFlag = 1
const QNetworkInterface__IsRunning QNetworkInterface__InterfaceFlag = 2
const QNetworkInterface__CanBroadcast QNetworkInterface__InterfaceFlag = 4
const QNetworkInterface__IsLoopBack QNetworkInterface__InterfaceFlag = 8
const QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = 16
const QNetworkInterface__CanMulticast QNetworkInterface__InterfaceFlag = 32

//  body block end
