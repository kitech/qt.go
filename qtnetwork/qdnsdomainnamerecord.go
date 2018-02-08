package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

type QDnsDomainNameRecord struct {
	*qtrt.CObject
}

func (this *QDnsDomainNameRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsDomainNameRecord) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDnsDomainNameRecordFromPointer(cthis unsafe.Pointer) *QDnsDomainNameRecord {
	return &QDnsDomainNameRecord{&qtrt.CObject{cthis}}
}
func (*QDnsDomainNameRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsDomainNameRecord {
	return NewQDnsDomainNameRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsDomainNameRecord()
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QDnsDomainNameRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsDomainNameRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsDomainNameRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsDomainNameRecord()
func DeleteQDnsDomainNameRecord(this *QDnsDomainNameRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QDnsDomainNameRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsDomainNameRecord &)
func (this *QDnsDomainNameRecord) Swap(other *QDnsDomainNameRecord) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QDnsDomainNameRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QDnsDomainNameRecord) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive()
func (this *QDnsDomainNameRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value()
func (this *QDnsDomainNameRecord) Value() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end
