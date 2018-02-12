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

type QDnsMailExchangeRecord struct {
	*qtrt.CObject
}

func (this *QDnsMailExchangeRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsMailExchangeRecord) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDnsMailExchangeRecordFromPointer(cthis unsafe.Pointer) *QDnsMailExchangeRecord {
	return &QDnsMailExchangeRecord{&qtrt.CObject{cthis}}
}
func (*QDnsMailExchangeRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsMailExchangeRecord {
	return NewQDnsMailExchangeRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsMailExchangeRecord()
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsMailExchangeRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsMailExchangeRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsMailExchangeRecord()
func DeleteQDnsMailExchangeRecord(this *QDnsMailExchangeRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:119
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsMailExchangeRecord &)
func (this *QDnsMailExchangeRecord) Swap(other *QDnsMailExchangeRecord) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QString exchange()
func (this *QDnsMailExchangeRecord) Exchange() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord8exchangeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QDnsMailExchangeRecord) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:123
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 preference()
func (this *QDnsMailExchangeRecord) Preference() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord10preferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive()
func (this *QDnsMailExchangeRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
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
