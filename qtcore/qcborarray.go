package qtcore

// /usr/include/qt/QtCore/qcborarray.h
// #include <qcborarray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QCborArray struct {
	*qtrt.CObject
}
type QCborArray_ITF interface {
	QCborArray_PTR() *QCborArray
}

func (ptr *QCborArray) QCborArray_PTR() *QCborArray { return ptr }

func (this *QCborArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborArrayFromPointer(cthis unsafe.Pointer) *QCborArray {
	return &QCborArray{&qtrt.CObject{cthis}}
}
func (*QCborArray) NewFromPointer(cthis unsafe.Pointer) *QCborArray {
	return NewQCborArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborarray.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCborArray()

/*
Constructs an empty QCborArray.
*/
func (*QCborArray) NewForInherit() *QCborArray {
	return NewQCborArray()
}
func NewQCborArray() *QCborArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArrayC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborArray)
	return gothis
}

// /usr/include/qt/QtCore/qcborarray.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QCborArray & operator=(const QCborArray &)

/*

 */
func (this *QCborArray) Operator_equal(other QCborArray_ITF) *QCborArray {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArrayaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCborArray()

/*

 */
func DeleteQCborArray(this *QCborArray) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArrayD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcborarray.h:174
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCborArray &)

/*
Swaps the contents of this object and other.
*/
func (this *QCborArray) Swap(other QCborArray_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:179
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue toCborValue() const

/*
Explicitly construcuts a QCborValue object that represents this array. This function is usually not necessary since QCborValue has a constructor for QCborArray, so the conversion is implicit.

Converting QCborArray to QCborValue allows it to be used in any context where QCborValues can be used, including as items in QCborArrays and as keys and mapped types in QCborMap. Converting an array to QCborValue allows access to QCborValue::toCbor().

See also QCborValue::QCborValue(const QCborArray &).
*/
func (this *QCborArray) ToCborValue() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray11toCborValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] qsizetype size() const

/*
Returns the size of this array.

See also isEmpty().
*/
func (this *QCborArray) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcborarray.h:182
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this QCborArray is empty (that is if size() is 0).

See also size() and clear().
*/
func (this *QCborArray) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Empties this array.

See also isEmpty().
*/
func (this *QCborArray) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:185
// index:0
// Public Visibility=Default Availability=Available
// [24] QCborValue at(qsizetype) const

/*
Returns the QCborValue element at position i in the array.

If the array is smaller than i elements, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the array is not large enough from the case where the array starts with an undefined value.

See also operator[](), first(), last(), insert(), prepend(), append(), removeAt(), and takeAt().
*/
func (this *QCborArray) At(i int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray2atEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue first() const

/*
Returns the first QCborValue of this array.

If the array is empty, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the array is not large enough from the case where the array ends with an undefined value.

See also operator[](), at(), last(), insert(), prepend(), append(), removeAt(), and takeAt().
*/
func (this *QCborArray) First() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:189
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef first()

/*
Returns the first QCborValue of this array.

If the array is empty, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the array is not large enough from the case where the array ends with an undefined value.

See also operator[](), at(), last(), insert(), prepend(), append(), removeAt(), and takeAt().
*/
func (this *QCborArray) First1() *QCborValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue last() const

/*
Returns the last QCborValue of this array.

If the array is empty, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the array is not large enough from the case where the array ends with an undefined value.

See also operator[](), at(), first(), insert(), prepend(), append(), removeAt(), and takeAt().
*/
func (this *QCborArray) Last() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:190
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef last()

/*
Returns the last QCborValue of this array.

If the array is empty, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the array is not large enough from the case where the array ends with an undefined value.

See also operator[](), at(), first(), insert(), prepend(), append(), removeAt(), and takeAt().
*/
func (this *QCborArray) Last1() *QCborValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [24] const QCborValue operator[](qsizetype) const

/*

 */
func (this *QCborArray) Operator_get_index(i int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArrayixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:191
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborValueRef operator[](qsizetype)

/*

 */
func (this *QCborArray) Operator_get_index1(i int64) *QCborValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArrayixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAt(qsizetype)

/*
Removes the item at position i from the array. The array must have more than i elements before the removal.

See also takeAt(), removeFirst(), removeLast(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) RemoveAt(i int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray8removeAtEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue takeAt(qsizetype)

/*
Removes the item at position i from the array and returns it. The array must have more than i elements before the removal.

See also removeAt(), removeFirst(), removeLast(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) TakeAt(i int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray6takeAtEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeFirst()

/*
Removes the first item in the array, making the second element become the first. The array must not be empty before this call.

See also removeAt(), takeFirst(), removeLast(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) RemoveFirst() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray11removeFirstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeLast()

/*
Removes the last item in the array. The array must not be empty before this call.

See also removeAt(), takeLast(), removeFirst(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) RemoveLast() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray10removeLastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue takeFirst()

/*
Removes the first item in the array and returns it, making the second element become the first. The array must not be empty before this call.

See also takeAt(), removeFirst(), removeLast(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) TakeFirst() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray9takeFirstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue takeLast()

/*
Removes the last item in the array and returns it. The array must not be empty before this call.

See also takeAt(), removeLast(), removeFirst(), at(), operator[](), insert(), prepend(), and append().
*/
func (this *QCborArray) TakeLast() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray8takeLastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:213
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QCborValue &) const

/*
Returns true if this array contains an element that is equal to value.
*/
func (this *QCborArray) Contains(value QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if value != nil && value.QCborValue_PTR() != nil {
		convArg0 = value.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray8containsERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:215
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QCborArray &) const

/*
Compares this array and other, comparing each element in sequence, and returns an integer that indicates whether this array should be sorted before (if the result is negative) or after other (if the result is positive). If this function returns 0, the two arrays are equal and contain the same elements.

For more information on CBOR sorting order, see QCborValue::compare().

See also QCborValue::compare(), QCborMap::compare(), and operator==().
*/
func (this *QCborArray) Compare(other QCborArray_ITF) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcborarray.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QCborArray &) const

/*

 */
func (this *QCborArray) Operator_equal_equal(other QCborArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArrayeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCborArray &) const

/*

 */
func (this *QCborArray) Operator_not_equal(other QCborArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArrayneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QCborArray &) const

/*

 */
func (this *QCborArray) Operator_less_than(other QCborArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborArray_PTR() != nil {
		convArg0 = other.QCborArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArrayltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::iterator begin()

/*
Returns an array iterator pointing to the first item in this array. If the array is empty, then this function returns the same as end().

See also constBegin() and end().
*/
func (this *QCborArray) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:237
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator begin() const

/*
Returns an array iterator pointing to the first item in this array. If the array is empty, then this function returns the same as end().

See also constBegin() and end().
*/
func (this *QCborArray) Begin1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator constBegin() const

/*
Returns an array iterator pointing to the first item in this array. If the array is empty, then this function returns the same as end().

See also begin() and constEnd().
*/
func (this *QCborArray) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator cbegin() const

/*
Returns an array iterator pointing to the first item in this array. If the array is empty, then this function returns the same as end().

See also constBegin() and constEnd().
*/
func (this *QCborArray) Cbegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::iterator end()

/*
Returns an array iterator pointing to just after the last element in this array.

See also begin() and constEnd().
*/
func (this *QCborArray) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:241
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator end() const

/*
Returns an array iterator pointing to just after the last element in this array.

See also begin() and constEnd().
*/
func (this *QCborArray) End1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator constEnd() const

/*
Returns an array iterator pointing to just after the last element in this array.

See also constBegin() and end().
*/
func (this *QCborArray) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:242
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborArray::const_iterator cend() const

/*
Returns an array iterator pointing to just after the last element in this array.

See also constBegin() and constEnd().
*/
func (this *QCborArray) Cend() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcborarray.h:250
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QCborValue &)

/*
Synonym for append(). This function is provided for compatibility with generic code that uses the Standard Library API.

Appends the element t to this array.

See also append(), push_front(), pop_back(), prepend(), and insert().
*/
func (this *QCborArray) Push_back(t QCborValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QCborValue_PTR() != nil {
		convArg0 = t.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray9push_backERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:251
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QCborValue &)

/*
Synonym for prepend(). This function is provided for compatibility with generic code that uses the Standard Library API.

Prepends the element t to this array.

See also prepend(), push_back(), pop_front(), append(), and insert().
*/
func (this *QCborArray) Push_front(t QCborValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QCborValue_PTR() != nil {
		convArg0 = t.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray10push_frontERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:252
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_front()

/*
Synonym for removeFirst(). This function is provided for compatibility with generic code that uses the Standard Library API.

Removes the first element of this array. The array must not be empty before the removal

See also removeFirst(), takeFirst(), pop_back(), push_front(), prepend(), and insert().
*/
func (this *QCborArray) Pop_front() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray9pop_frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:253
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_back()

/*
Synonym for removeLast(). This function is provided for compatibility with generic code that uses the Standard Library API.

Removes the last element of this array. The array must not be empty before the removal

See also removeLast(), takeLast(), pop_front(), push_back(), append(), and insert().
*/
func (this *QCborArray) Pop_back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray8pop_backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborarray.h:254
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty() const

/*
Synonym for isEmpty(). This function is provided for compatibility with generic code that uses the Standard Library API.

Returns true if this array is empty (size() == 0).

See also isEmpty() and size().
*/
func (this *QCborArray) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborarray.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborArray operator+(const QCborValue &) const

/*

 */
func (this *QCborArray) Operator_add(v QCborValue_ITF) *QCborArray /*123*/ {
	var convArg0 unsafe.Pointer
	if v != nil && v.QCborValue_PTR() != nil {
		convArg0 = v.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArrayplERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:259
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborArray & operator+=(const QCborValue &)

/*

 */
func (this *QCborArray) Operator_add_equal(v QCborValue_ITF) *QCborArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QCborValue_PTR() != nil {
		convArg0 = v.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArraypLERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:261
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QCborArray & operator<<(const QCborValue &)

/*

 */
func (this *QCborArray) Operator_left_shift(v QCborValue_ITF) *QCborArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QCborValue_PTR() != nil {
		convArg0 = v.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArraylsERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QCborArray fromStringList(const QStringList &)

/*
Returns a QCborArray containing all the strings found in the list list.

See also fromVariantList() and fromJsonArray().
*/
func (this *QCborArray) FromStringList(list QStringList_ITF) *QCborArray /*123*/ {
	var convArg0 unsafe.Pointer
	if list != nil && list.QStringList_PTR() != nil {
		convArg0 = list.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray14fromStringListERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}
func QCborArray_FromStringList(list QStringList_ITF) *QCborArray /*123*/ {
	var nilthis *QCborArray
	rv := nilthis.FromStringList(list)
	return rv
}

// /usr/include/qt/QtCore/qcborarray.h:266
// index:0
// Public static Visibility=Default Availability=Available
// [8] QCborArray fromJsonArray(const QJsonArray &)

/*
Converts all JSON items found in the array array to CBOR using QCborValue::fromJson(), and returns the CBOR array composed of those elements.

This conversion is lossless, as the CBOR type system is a superset of JSON's. Moreover, the array returned by this function can be converted back to the original array by using toJsonArray().

See also toJsonArray(), toVariantList(), QCborValue::fromJsonValue(), and QCborMap::fromJsonObject().
*/
func (this *QCborArray) FromJsonArray(array QJsonArray_ITF) *QCborArray /*123*/ {
	var convArg0 unsafe.Pointer
	if array != nil && array.QJsonArray_PTR() != nil {
		convArg0 = array.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCborArray13fromJsonArrayERK10QJsonArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborArray)
	return rv2
}
func QCborArray_FromJsonArray(array QJsonArray_ITF) *QCborArray /*123*/ {
	var nilthis *QCborArray
	rv := nilthis.FromJsonArray(array)
	return rv
}

// /usr/include/qt/QtCore/qcborarray.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QVariantList toVariantList() const

/*
Recursively converts each QCborValue in this array using QCborValue::toVariant() and returns the QVariantList composed of the converted items.

Conversion to QVariant is not completely lossless. Please see the documentation in QCborValue::toVariant() for more information.

See also fromVariantList(), fromStringList(), toJsonArray(), QCborValue::toVariant(), and QCborMap::toVariantMap().
*/
func (this *QCborArray) ToVariantList() *QVariantList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray13toVariantListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qcborarray.h:268
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray toJsonArray() const

/*
Recursively converts every QCborValue element in this array to JSON using QCborValue::toJsonValue() and returns the corresponding QJsonArray composed of those elements.

Please note that CBOR contains a richer and wider type set than JSON, so some information may be lost in this conversion. For more details on what conversions are applied, see QCborValue::toJsonValue().

See also fromJsonArray(), QCborValue::toJsonValue(), QCborMap::toJsonObject(), and toVariantList().
*/
func (this *QCborArray) ToJsonArray() *QJsonArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCborArray11toJsonArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10355() {
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
