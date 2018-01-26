package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDnsMailExchangeRecordFromPointer(cthis unsafe.Pointer) *QDnsMailExchangeRecord {
	return &QDnsMailExchangeRecord{&qtrt.CObject{cthis}}
}
func (*QDnsMailExchangeRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsMailExchangeRecord {
	return NewQDnsMailExchangeRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:111
// index:0
// Public
// void QDnsMailExchangeRecord()
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecordC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsMailExchangeRecordFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:117
// index:0
// Public
// void ~QDnsMailExchangeRecord()
func DeleteQDnsMailExchangeRecord(*QDnsMailExchangeRecord) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecordD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:119
// index:0
// Public inline
// void swap(class QDnsMailExchangeRecord &)
func (this *QDnsMailExchangeRecord) Swap(other *QDnsMailExchangeRecord) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QDnsMailExchangeRecord4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:121
// index:0
// Public
// QString exchange()
func (this *QDnsMailExchangeRecord) Exchange() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord8exchangeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:122
// index:0
// Public
// QString name()
func (this *QDnsMailExchangeRecord) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:123
// index:0
// Public
// quint16 preference()
func (this *QDnsMailExchangeRecord) Preference() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord10preferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:124
// index:0
// Public
// quint32 timeToLive()
func (this *QDnsMailExchangeRecord) TimeToLive() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QDnsMailExchangeRecord10timeToLiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

//  body block end
