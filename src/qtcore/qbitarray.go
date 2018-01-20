//  header block begin
// /usr/include/qt/QtCore/qbitarray.h
// #include <qbitarray.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QBitArray struct {
	*qtrt.CObject
}

func (this *QBitArray) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQBitArrayFromPointer(cthis unsafe.Pointer) *QBitArray {
	return &QBitArray{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbitarray.h:57
// index:0
// Public inline
// void QBitArray()
func NewQBitArray() *QBitArray {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArrayC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbitarray.h:58
// index:1
// Public
// void QBitArray(int, _Bool)
func NewQBitArray_1(size int, val bool) *QBitArray {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArrayC2Eib", ffiqt.FFI_TYPE_VOID, cthis, &size, &val)
	gopp.ErrPrint(err, rv)
	gothis := NewQBitArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbitarray.h:67
// index:0
// Public inline
// void swap(class QBitArray &)
func (this *QBitArray) Swap(other *QBitArray) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:69
// index:0
// Public inline
// int size()
func (this *QBitArray) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:70
// index:0
// Public inline
// int count()
func (this *QBitArray) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:71
// index:1
// Public
// int count(_Bool)
func (this *QBitArray) Count_1(on bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray5countEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:73
// index:0
// Public inline
// bool isEmpty()
func (this *QBitArray) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:74
// index:0
// Public inline
// bool isNull()
func (this *QBitArray) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:76
// index:0
// Public
// void resize(int)
func (this *QBitArray) Resize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6resizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:78
// index:0
// Public inline
// void detach()
func (this *QBitArray) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:79
// index:0
// Public inline
// bool isDetached()
func (this *QBitArray) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:80
// index:0
// Public inline
// void clear()
func (this *QBitArray) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:82
// index:0
// Public
// bool testBit(int)
func (this *QBitArray) TestBit(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray7testBitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:83
// index:0
// Public
// void setBit(int)
func (this *QBitArray) SetBit(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6setBitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:84
// index:1
// Public
// void setBit(int, _Bool)
func (this *QBitArray) SetBit_1(i int, val bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6setBitEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:85
// index:0
// Public
// void clearBit(int)
func (this *QBitArray) ClearBit(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray8clearBitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:86
// index:0
// Public
// bool toggleBit(int)
func (this *QBitArray) ToggleBit(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray9toggleBitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:88
// index:0
// Public
// bool at(int)
func (this *QBitArray) At(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:102
// index:0
// Public inline
// bool fill(_Bool, int)
func (this *QBitArray) Fill(val bool, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4fillEbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &val, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbitarray.h:103
// index:1
// Public
// void fill(_Bool, int, int)
func (this *QBitArray) Fill_1(val bool, first int, last int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4fillEbii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &val, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:105
// index:0
// Public inline
// void truncate(int)
func (this *QBitArray) Truncate(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray8truncateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
