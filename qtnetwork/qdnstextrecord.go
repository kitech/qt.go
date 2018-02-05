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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDnsTextRecordFromPointer(cthis unsafe.Pointer) *QDnsTextRecord {
	return &QDnsTextRecord{&qtrt.CObject{cthis}}
}
func (*QDnsTextRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsTextRecord {
	return NewQDnsTextRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsTextRecord()
func NewQDnsTextRecord() *QDnsTextRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQDnsTextRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsTextRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsTextRecord()
func DeleteQDnsTextRecord(this *QDnsTextRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsTextRecord &)
func (this *QDnsTextRecord) Swap(other *QDnsTextRecord) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:173
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QDnsTextRecord) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDnsTextRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive()
func (this *QDnsTextRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDnsTextRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

//  body block end
