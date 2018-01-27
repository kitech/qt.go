package qtnetwork

// /usr/include/qt/QtNetwork/qdnslookup.h
// #include <qdnslookup.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QDnsTextRecord struct {
	*qtrt.CObject
}

func (this *QDnsTextRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsTextRecord) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDnsTextRecordFromPointer(cthis unsafe.Pointer) *QDnsTextRecord {
	return &QDnsTextRecord{&qtrt.CObject{cthis}}
}
func (*QDnsTextRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsTextRecord {
	return NewQDnsTextRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:163
// index:0
// Public
// void QDnsTextRecord()
func NewQDnsTextRecord() *QDnsTextRecord {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDnsTextRecordC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsTextRecordFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:169
// index:0
// Public
// void ~QDnsTextRecord()
func DeleteQDnsTextRecord(*QDnsTextRecord) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDnsTextRecordD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:171
// index:0
// Public inline
// void swap(QDnsTextRecord &)
func (this *QDnsTextRecord) Swap(other *QDnsTextRecord) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDnsTextRecord4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:173
// index:0
// Public
// QString name()
func (this *QDnsTextRecord) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDnsTextRecord4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:174
// index:0
// Public
// quint32 timeToLive()
func (this *QDnsTextRecord) TimeToLive() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDnsTextRecord10timeToLiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

//  body block end
