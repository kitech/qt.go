package qtcore

// /usr/include/qt/QtCore/qbitarray.h
// #include <qbitarray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QBitArray struct {
	*qtrt.CObject
}
type QBitArray_ITF interface {
	QBitArray_PTR() *QBitArray
}

func (ptr *QBitArray) QBitArray_PTR() *QBitArray { return ptr }

func (this *QBitArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBitArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBitArrayFromPointer(cthis unsafe.Pointer) *QBitArray {
	return &QBitArray{&qtrt.CObject{cthis}}
}
func (*QBitArray) NewFromPointer(cthis unsafe.Pointer) *QBitArray {
	return NewQBitArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbitarray.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QBitArray()
func NewQBitArray() *QBitArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitArray)
	return gothis
}

// /usr/include/qt/QtCore/qbitarray.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBitArray(int, _Bool)
func NewQBitArray_1(size int, val bool) *QBitArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayC2Eib", qtrt.FFI_TYPE_POINTER, size, val)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitArray)
	return gothis
}

// /usr/include/qt/QtCore/qbitarray.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBitArray(int, _Bool)
func NewQBitArray_1_(size int) *QBitArray {
	// arg: 1, bool=Bool, =Invalid,
	val := false
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayC2Eib", qtrt.FFI_TYPE_POINTER, size, val)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBitArray)
	return gothis
}

// /usr/include/qt/QtCore/qbitarray.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBitArray & operator=(const QBitArray &)
func (this *QBitArray) Operator_equal(other QBitArray_ITF) *QBitArray {
	var convArg0 unsafe.Pointer
	if other != nil && other.QBitArray_PTR() != nil {
		convArg0 = other.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:63
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QBitArray & operator=(QBitArray &&)
func (this *QBitArray) Operator_equal_1(other unsafe.Pointer /*333*/) *QBitArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QBitArray &)
func (this *QBitArray) Swap(other QBitArray_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QBitArray_PTR() != nil {
		convArg0 = other.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const
func (this *QBitArray) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbitarray.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const
func (this *QBitArray) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbitarray.h:71
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(_Bool) const
func (this *QBitArray) Count_1(on bool) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray5countEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbitarray.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QBitArray) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QBitArray) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int)
func (this *QBitArray) Resize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray6resizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void detach()
func (this *QBitArray) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached() const
func (this *QBitArray) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()
func (this *QBitArray) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testBit(int) const
func (this *QBitArray) TestBit(i int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray7testBitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBit(int)
func (this *QBitArray) SetBit(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray6setBitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:84
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBit(int, _Bool)
func (this *QBitArray) SetBit_1(i int, val bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray6setBitEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearBit(int)
func (this *QBitArray) ClearBit(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray8clearBitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:86
// index:0
// Public Visibility=Default Availability=Available
// [1] bool toggleBit(int)
func (this *QBitArray) ToggleBit(i int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray9toggleBitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool at(int) const
func (this *QBitArray) At(i int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArray2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:89
// index:0
// Public Visibility=Default Availability=Available
// [16] QBitRef operator[](int)
func (this *QBitArray) Operator_get_index(i int) *QBitRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitRef)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:90
// index:1
// Public Visibility=Default Availability=Available
// [1] bool operator[](int) const
func (this *QBitArray) Operator_get_index_1(i int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:91
// index:2
// Public Visibility=Default Availability=Available
// [16] QBitRef operator[](uint)
func (this *QBitArray) Operator_get_index_2(i uint) *QBitRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitRef)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:92
// index:3
// Public Visibility=Default Availability=Available
// [1] bool operator[](uint) const
func (this *QBitArray) Operator_get_index_3(i uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray & operator&=(const QBitArray &)
func (this *QBitArray) Operator_and_equal(arg0 QBitArray_ITF) *QBitArray {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitArray_PTR() != nil {
		convArg0 = arg0.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayaNERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray & operator|=(const QBitArray &)
func (this *QBitArray) Operator_or_equal(arg0 QBitArray_ITF) *QBitArray {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitArray_PTR() != nil {
		convArg0 = arg0.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayoRERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray & operator^=(const QBitArray &)
func (this *QBitArray) Operator_caret_equal(arg0 QBitArray_ITF) *QBitArray {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QBitArray_PTR() != nil {
		convArg0 = arg0.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayeOERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QBitArray operator~() const
func (this *QBitArray) Operator_around() *QBitArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArraycoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitArray)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QBitArray &) const
func (this *QBitArray) Operator_equal_equal(other QBitArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QBitArray_PTR() != nil {
		convArg0 = other.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArrayeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QBitArray &) const
func (this *QBitArray) Operator_not_equal(other QBitArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QBitArray_PTR() != nil {
		convArg0 = other.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QBitArrayneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fill(_Bool, int)
func (this *QBitArray) Fill(val bool, size int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray4fillEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val, size)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fill(_Bool, int)
func (this *QBitArray) Fill__(val bool) bool {
	// arg: 1, int=Int, =Invalid,
	size := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray4fillEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val, size)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void fill(_Bool, int, int)
func (this *QBitArray) Fill_1(val bool, first int, last int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray4fillEbii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(int)
func (this *QBitArray) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArray8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

func DeleteQBitArray(this *QBitArray) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QBitArrayD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
