package qtcore

// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 109
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QCollatorSortKey struct {
	*qtrt.CObject
}
type QCollatorSortKey_ITF interface {
	QCollatorSortKey_PTR() *QCollatorSortKey
}

func (ptr *QCollatorSortKey) QCollatorSortKey_PTR() *QCollatorSortKey { return ptr }

func (this *QCollatorSortKey) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCollatorSortKey) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCollatorSortKeyFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return &QCollatorSortKey{&qtrt.CObject{cthis}}
}
func (*QCollatorSortKey) NewFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return NewQCollatorSortKeyFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcollator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCollatorSortKey()
func DeleteQCollatorSortKey(this *QCollatorSortKey) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKeyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcollator.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QCollatorSortKey & operator=(const QCollatorSortKey &)
func (this *QCollatorSortKey) Operator_equal(other QCollatorSortKey_ITF) *QCollatorSortKey {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCollatorSortKey_PTR() != nil {
		convArg0 = other.QCollatorSortKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKeyaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCollatorSortKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollatorSortKey)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QCollatorSortKey & operator=(QCollatorSortKey &&)
func (this *QCollatorSortKey) Operator_equal_1(other unsafe.Pointer /*333*/) *QCollatorSortKey {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKeyaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCollatorSortKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCollatorSortKey)
	return rv2
}

// /usr/include/qt/QtCore/qcollator.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(other QCollatorSortKey_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCollatorSortKey_PTR() != nil {
		convArg0 = other.QCollatorSortKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKey4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(key QCollatorSortKey_ITF) int {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCollatorSortKey_PTR() != nil {
		convArg0 = key.QCollatorSortKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCollatorSortKey7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
}

//  keep block end
