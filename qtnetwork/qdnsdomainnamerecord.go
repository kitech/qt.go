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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDnsDomainNameRecordFromPointer(cthis unsafe.Pointer) *QDnsDomainNameRecord {
	return &QDnsDomainNameRecord{&qtrt.CObject{cthis}}
}
func (*QDnsDomainNameRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsDomainNameRecord {
	return NewQDnsDomainNameRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:63
// index:0
// Public
// void QDnsDomainNameRecord()
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDnsDomainNameRecordC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsDomainNameRecordFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:69
// index:0
// Public
// void ~QDnsDomainNameRecord()
func DeleteQDnsDomainNameRecord(*QDnsDomainNameRecord) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDnsDomainNameRecordD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:71
// index:0
// Public inline
// void swap(class QDnsDomainNameRecord &)
func (this *QDnsDomainNameRecord) Swap(other *QDnsDomainNameRecord) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDnsDomainNameRecord4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:73
// index:0
// Public
// QString name()
func (this *QDnsDomainNameRecord) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:74
// index:0
// Public
// quint32 timeToLive()
func (this *QDnsDomainNameRecord) TimeToLive() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord10timeToLiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:75
// index:0
// Public
// QString value()
func (this *QDnsDomainNameRecord) Value() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QDnsDomainNameRecord5valueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
