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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qbitarray.h:57
// index:0
// inline
// void QBitArray()
func NewQBitArray() *QBitArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArrayC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QBitArray{cthis}
}

// /usr/include/qt/QtCore/qbitarray.h:58
// index:1
// void QBitArray(int, _Bool)
func NewQBitArray_1(size int, val bool) *QBitArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArrayC2Eib", ffiqt.FFI_TYPE_VOID, cthis, &size, &val)
	gopp.ErrPrint(err, rv)
	return &QBitArray{cthis}
}

// /usr/include/qt/QtCore/qbitarray.h:67
// index:0
// inline
// void swap(class QBitArray &)
func (this *QBitArray) Swap(other unsafe.Pointer) {
	// 0: (, QBitArray & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:69
// index:0
// inline
// int size()
func (this *QBitArray) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:70
// index:0
// inline
// int count()
func (this *QBitArray) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:71
// index:1
// int count(_Bool)
func (this *QBitArray) Count_1(on bool) {
	// 1: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray5countEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:73
// index:0
// inline
// bool isEmpty()
func (this *QBitArray) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:74
// index:0
// inline
// bool isNull()
func (this *QBitArray) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:76
// index:0
// void resize(int)
func (this *QBitArray) Resize(size int) {
	// 0: (, int size), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6resizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:78
// index:0
// inline
// void detach()
func (this *QBitArray) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6detachEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:79
// index:0
// inline
// bool isDetached()
func (this *QBitArray) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:80
// index:0
// inline
// void clear()
func (this *QBitArray) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:82
// index:0
// bool testBit(int)
func (this *QBitArray) TestBit(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray7testBitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:83
// index:0
// void setBit(int)
func (this *QBitArray) SetBit(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6setBitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:84
// index:1
// void setBit(int, _Bool)
func (this *QBitArray) SetBit_1(i int, val bool) {
	// 1: (, int i, bool val), (&i, &val)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray6setBitEib", ffiqt.FFI_TYPE_VOID, this.cthis, &i, &val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:85
// index:0
// void clearBit(int)
func (this *QBitArray) ClearBit(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray8clearBitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:86
// index:0
// bool toggleBit(int)
func (this *QBitArray) ToggleBit(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray9toggleBitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:88
// index:0
// bool at(int)
func (this *QBitArray) At(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QBitArray2atEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:102
// index:0
// inline
// bool fill(_Bool, int)
func (this *QBitArray) Fill(val bool, size int) {
	// 0: (, bool val, int size), (&val, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4fillEbi", ffiqt.FFI_TYPE_VOID, this.cthis, &val, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:103
// index:1
// void fill(_Bool, int, int)
func (this *QBitArray) Fill_1(val bool, first int, last int) {
	// 1: (, bool val, int first, int last), (&val, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray4fillEbii", ffiqt.FFI_TYPE_VOID, this.cthis, &val, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbitarray.h:105
// index:0
// inline
// void truncate(int)
func (this *QBitArray) Truncate(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QBitArray8truncateEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
