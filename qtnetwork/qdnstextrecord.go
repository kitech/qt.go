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
// extern C begin: 11
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

/*

 */
type QDnsTextRecord struct {
	*qtrt.CObject
}
type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (ptr *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord { return ptr }

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

/*

 */
func NewQDnsTextRecord() *QDnsTextRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsTextRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsTextRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDnsTextRecord & operator=(QDnsTextRecord &&)

/*

 */
func (this *QDnsTextRecord) Operator_equal(other unsafe.Pointer /*333*/) *QDnsTextRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDnsTextRecordFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDnsTextRecord)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:168
// index:1
// Public Visibility=Default Availability=Available
// [8] QDnsTextRecord & operator=(const QDnsTextRecord &)

/*

 */
func (this *QDnsTextRecord) Operator_equal_1(other QDnsTextRecord_ITF) *QDnsTextRecord {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDnsTextRecord_PTR() != nil {
		convArg0 = other.QDnsTextRecord_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDnsTextRecordFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDnsTextRecord)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsTextRecord()

/*

 */
func DeleteQDnsTextRecord(this *QDnsTextRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsTextRecord &)

/*

 */
func (this *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDnsTextRecord_PTR() != nil {
		convArg0 = other.QDnsTextRecord_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDnsTextRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:173
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QDnsTextRecord) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDnsTextRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive() const

/*

 */
func (this *QDnsTextRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDnsTextRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
