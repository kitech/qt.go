package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkinterface.h
// #include <qnetworkinterface.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
type QNetworkAddressEntry struct {
	*qtrt.CObject
}

func (this *QNetworkAddressEntry) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkAddressEntry) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkAddressEntryFromPointer(cthis unsafe.Pointer) *QNetworkAddressEntry {
	return &QNetworkAddressEntry{&qtrt.CObject{cthis}}
}
func (*QNetworkAddressEntry) NewFromPointer(cthis unsafe.Pointer) *QNetworkAddressEntry {
	return NewQNetworkAddressEntryFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:59
// index:0
// Public
// void QNetworkAddressEntry()
func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntryC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkAddressEntryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:65
// index:0
// Public
// void ~QNetworkAddressEntry()
func DeleteQNetworkAddressEntry(*QNetworkAddressEntry) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:67
// index:0
// Public inline
// void swap(class QNetworkAddressEntry &)
func (this *QNetworkAddressEntry) Swap(other *QNetworkAddressEntry) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntry4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:73
// index:0
// Public
// QHostAddress ip()
func (this *QNetworkAddressEntry) Ip() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry2ipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:74
// index:0
// Public
// void setIp(const class QHostAddress &)
func (this *QNetworkAddressEntry) SetIp(newIp *QHostAddress) {
	var convArg0 = newIp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntry5setIpERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:76
// index:0
// Public
// QHostAddress netmask()
func (this *QNetworkAddressEntry) Netmask() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry7netmaskEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:77
// index:0
// Public
// void setNetmask(const class QHostAddress &)
func (this *QNetworkAddressEntry) SetNetmask(newNetmask *QHostAddress) {
	var convArg0 = newNetmask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntry10setNetmaskERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:78
// index:0
// Public
// int prefixLength()
func (this *QNetworkAddressEntry) PrefixLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry12prefixLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:79
// index:0
// Public
// void setPrefixLength(int)
func (this *QNetworkAddressEntry) SetPrefixLength(length int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntry15setPrefixLengthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:81
// index:0
// Public
// QHostAddress broadcast()
func (this *QNetworkAddressEntry) Broadcast() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QNetworkAddressEntry9broadcastEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:82
// index:0
// Public
// void setBroadcast(const class QHostAddress &)
func (this *QNetworkAddressEntry) SetBroadcast(newBroadcast *QHostAddress) {
	var convArg0 = newBroadcast.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkAddressEntry12setBroadcastERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
