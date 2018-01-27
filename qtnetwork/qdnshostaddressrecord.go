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
import "gopp"
import "qt.go/cffiqt"
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
type QDnsHostAddressRecord struct {
	*qtrt.CObject
}

func (this *QDnsHostAddressRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsHostAddressRecord) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDnsHostAddressRecordFromPointer(cthis unsafe.Pointer) *QDnsHostAddressRecord {
	return &QDnsHostAddressRecord{&qtrt.CObject{cthis}}
}
func (*QDnsHostAddressRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsHostAddressRecord {
	return NewQDnsHostAddressRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:87
// index:0
// Public
// void QDnsHostAddressRecord()
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QDnsHostAddressRecordC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsHostAddressRecordFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:93
// index:0
// Public
// void ~QDnsHostAddressRecord()
func DeleteQDnsHostAddressRecord(*QDnsHostAddressRecord) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QDnsHostAddressRecordD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:95
// index:0
// Public inline
// void swap(QDnsHostAddressRecord &)
func (this *QDnsHostAddressRecord) Swap(other *QDnsHostAddressRecord) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QDnsHostAddressRecord4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:97
// index:0
// Public
// QString name()
func (this *QDnsHostAddressRecord) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:98
// index:0
// Public
// quint32 timeToLive()
func (this *QDnsHostAddressRecord) TimeToLive() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord10timeToLiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:99
// index:0
// Public
// QHostAddress value()
func (this *QDnsHostAddressRecord) Value() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QDnsHostAddressRecord5valueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
