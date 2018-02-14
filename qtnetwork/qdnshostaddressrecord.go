package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QDnsHostAddressRecord struct {
	*qtrt.CObject
}
type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord { return ptr }

func (this *QDnsHostAddressRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsHostAddressRecord) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDnsHostAddressRecordFromPointer(cthis unsafe.Pointer) *QDnsHostAddressRecord {
	return &QDnsHostAddressRecord{&qtrt.CObject{cthis}}
}
func (*QDnsHostAddressRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsHostAddressRecord {
	return NewQDnsHostAddressRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsHostAddressRecord()
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QDnsHostAddressRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsHostAddressRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsHostAddressRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsHostAddressRecord()
func DeleteQDnsHostAddressRecord(this *QDnsHostAddressRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QDnsHostAddressRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsHostAddressRecord &)
func (this *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {
	var convArg0 = other.QDnsHostAddressRecord_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QDnsHostAddressRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QDnsHostAddressRecord) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive()
func (this *QDnsHostAddressRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress value()
func (this *QDnsHostAddressRecord) Value() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
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
