// +build !minimal

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
// extern C begin: 9
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
type QDnsServiceRecord struct {
	*qtrt.CObject
}
type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (ptr *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord { return ptr }

func (this *QDnsServiceRecord) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDnsServiceRecord) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDnsServiceRecordFromPointer(cthis unsafe.Pointer) *QDnsServiceRecord {
	return &QDnsServiceRecord{&qtrt.CObject{cthis}}
}
func (*QDnsServiceRecord) NewFromPointer(cthis unsafe.Pointer) *QDnsServiceRecord {
	return NewQDnsServiceRecordFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDnsServiceRecord()

/*

 */
func (*QDnsServiceRecord) NewForInherit() *QDnsServiceRecord {
	return NewQDnsServiceRecord()
}
func NewQDnsServiceRecord() *QDnsServiceRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDnsServiceRecordC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDnsServiceRecordFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDnsServiceRecord)
	return gothis
}

// /usr/include/qt/QtNetwork/qdnslookup.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QDnsServiceRecord & operator=(QDnsServiceRecord &&)

/*

 */
func (this *QDnsServiceRecord) Operator_equal(other unsafe.Pointer /*333*/) *QDnsServiceRecord {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDnsServiceRecordaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDnsServiceRecordFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDnsServiceRecord)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:143
// index:1
// Public Visibility=Default Availability=Available
// [8] QDnsServiceRecord & operator=(const QDnsServiceRecord &)

/*

 */
func (this *QDnsServiceRecord) Operator_equal1(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDnsServiceRecord_PTR() != nil {
		convArg0 = other.QDnsServiceRecord_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDnsServiceRecordaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDnsServiceRecordFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDnsServiceRecord)
	return rv2
}

// /usr/include/qt/QtNetwork/qdnslookup.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDnsServiceRecord()

/*

 */
func DeleteQDnsServiceRecord(this *QDnsServiceRecord) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDnsServiceRecordD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDnsServiceRecord &)

/*

 */
func (this *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDnsServiceRecord_PTR() != nil {
		convArg0 = other.QDnsServiceRecord_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDnsServiceRecord4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdnslookup.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QDnsServiceRecord) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:149
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 port() const

/*

 */
func (this *QDnsServiceRecord) Port() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord4portEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:150
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 priority() const

/*

 */
func (this *QDnsServiceRecord) Priority() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QString target() const

/*

 */
func (this *QDnsServiceRecord) Target() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord6targetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdnslookup.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] quint32 timeToLive() const

/*

 */
func (this *QDnsServiceRecord) TimeToLive() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord10timeToLiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qdnslookup.h:153
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 weight() const

/*

 */
func (this *QDnsServiceRecord) Weight() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDnsServiceRecord6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
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
