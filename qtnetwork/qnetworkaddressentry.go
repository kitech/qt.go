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
// extern C begin: 16
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

type QNetworkAddressEntry struct {
	*qtrt.CObject
}
type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry { return ptr }

func (this *QNetworkAddressEntry) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkAddressEntry) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkAddressEntryFromPointer(cthis unsafe.Pointer) *QNetworkAddressEntry {
	return &QNetworkAddressEntry{&qtrt.CObject{cthis}}
}
func (*QNetworkAddressEntry) NewFromPointer(cthis unsafe.Pointer) *QNetworkAddressEntry {
	return NewQNetworkAddressEntryFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkAddressEntry()
func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntryC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkAddressEntryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkAddressEntry)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkAddressEntry & operator=(QNetworkAddressEntry &&)
func (this *QNetworkAddressEntry) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkAddressEntry {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntryaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkAddressEntryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkAddressEntry)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:64
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkAddressEntry & operator=(const QNetworkAddressEntry &)
func (this *QNetworkAddressEntry) Operator_equal_1(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkAddressEntry_PTR() != nil {
		convArg0 = other.QNetworkAddressEntry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntryaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkAddressEntryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkAddressEntry)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkAddressEntry()
func DeleteQNetworkAddressEntry(this *QNetworkAddressEntry) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkAddressEntry &)
func (this *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkAddressEntry_PTR() != nil {
		convArg0 = other.QNetworkAddressEntry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntry4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkAddressEntry &)
func (this *QNetworkAddressEntry) Operator_equal_equal(other QNetworkAddressEntry_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkAddressEntry_PTR() != nil {
		convArg0 = other.QNetworkAddressEntry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntryeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkAddressEntry &)
func (this *QNetworkAddressEntry) Operator_not_equal(other QNetworkAddressEntry_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkAddressEntry_PTR() != nil {
		convArg0 = other.QNetworkAddressEntry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntryneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress ip()
func (this *QNetworkAddressEntry) Ip() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry2ipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIp(const QHostAddress &)
func (this *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if newIp != nil && newIp.QHostAddress_PTR() != nil {
		convArg0 = newIp.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntry5setIpERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress netmask()
func (this *QNetworkAddressEntry) Netmask() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry7netmaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNetmask(const QHostAddress &)
func (this *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if newNetmask != nil && newNetmask.QHostAddress_PTR() != nil {
		convArg0 = newNetmask.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntry10setNetmaskERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int prefixLength()
func (this *QNetworkAddressEntry) PrefixLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry12prefixLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrefixLength(int)
func (this *QNetworkAddressEntry) SetPrefixLength(length int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntry15setPrefixLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress broadcast()
func (this *QNetworkAddressEntry) Broadcast() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry9broadcastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBroadcast(const QHostAddress &)
func (this *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if newBroadcast != nil && newBroadcast.QHostAddress_PTR() != nil {
		convArg0 = newBroadcast.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QNetworkAddressEntry12setBroadcastERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
