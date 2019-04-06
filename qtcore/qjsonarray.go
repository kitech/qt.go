package qtcore

// /usr/include/qt/QtCore/qjsonarray.h
// #include <qjsonarray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QJsonArray struct {
	*qtrt.CObject
}
type QJsonArray_ITF interface {
	QJsonArray_PTR() *QJsonArray
}

func (ptr *QJsonArray) QJsonArray_PTR() *QJsonArray { return ptr }

func (this *QJsonArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonArrayFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return &QJsonArray{&qtrt.CObject{cthis}}
}
func (*QJsonArray) NewFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return NewQJsonArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonarray.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonArray()

/*
Creates an empty array.
*/
func (*QJsonArray) NewForInherit() *QJsonArray {
	return NewQJsonArray()
}
func NewQJsonArray() *QJsonArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonArray)
	return gothis
}

// /usr/include/qt/QtCore/qjsonarray.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonArray()

/*

 */
func DeleteQJsonArray(this *QJsonArray) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonarray.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonArray & operator=(const QJsonArray &)

/*

 */
func (this *QJsonArray) Operator_equal(other QJsonArray_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:83
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator=(QJsonArray &&)

/*

 */
func (this *QJsonArray) Operator_equal1(other unsafe.Pointer /*333*/) *QJsonArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [16] QJsonArray fromStringList(const QStringList &)

/*
Converts the string list list to a QJsonArray.

The values in list will be converted to JSON values.

See also toVariantList() and QJsonValue::fromVariant().
*/
func (this *QJsonArray) FromStringList(list QStringList_ITF) *QJsonArray /*123*/ {
	var convArg0 unsafe.Pointer
	if list != nil && list.QStringList_PTR() != nil {
		convArg0 = list.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray14fromStringListERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}
func QJsonArray_FromStringList(list QStringList_ITF) *QJsonArray /*123*/ {
	var nilthis *QJsonArray
	rv := nilthis.FromStringList(list)
	return rv
}

// /usr/include/qt/QtCore/qjsonarray.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QVariantList toVariantList() const

/*
Converts this object to a QVariantList.

Returns the created map.
*/
func (this *QJsonArray) ToVariantList() *QVariantList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray13toVariantListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of values stored in the array.
*/
func (this *QJsonArray) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const

/*
Same as size().

See also size().
*/
func (this *QJsonArray) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the object is empty. This is the same as size() == 0.

See also size().
*/
func (this *QJsonArray) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:97
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue at(int) const

/*
Returns a QJsonValue representing the value for index i.

The returned QJsonValue is Undefined, if i is out of bounds.
*/
func (this *QJsonArray) At(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:98
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue first() const

/*
Returns the first value stored in the array.

Same as at(0).

See also at().
*/
func (this *QJsonArray) First() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:99
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue last() const

/*
Returns the last value stored in the array.

Same as at(size() - 1).

See also at().
*/
func (this *QJsonArray) Last() *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAt(int)

/*
Removes the value at index position i. i must be a valid index position in the array (i.e., 0 <= i < size()).

See also insert() and replace().
*/
func (this *QJsonArray) RemoveAt(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:104
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue takeAt(int)

/*
Removes the item at index position i and returns it. i must be a valid index position in the array (i.e., 0 <= i < size()).

If you don't use the return value, removeAt() is more efficient.

See also removeAt().
*/
func (this *QJsonArray) TakeAt(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeFirst()

/*
Removes the first item in the array. Calling this function is equivalent to calling removeAt(0). The array must not be empty. If the array can be empty, call isEmpty() before calling this function.

See also removeAt() and removeLast().
*/
func (this *QJsonArray) RemoveFirst() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray11removeFirstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeLast()

/*
Removes the last item in the array. Calling this function is equivalent to calling removeAt(size() - 1). The array must not be empty. If the array can be empty, call isEmpty() before calling this function.

See also removeAt() and removeFirst().
*/
func (this *QJsonArray) RemoveLast() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray10removeLastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replace(int, const QJsonValue &)

/*
Replaces the item at index position i with value. i must be a valid index position in the array (i.e., 0 <= i < size()).

See also operator[]() and removeAt().
*/
func (this *QJsonArray) Replace(i int, value QJsonValue_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QJsonValue_PTR() != nil {
		convArg1 = value.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray7replaceEiRK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QJsonValue &) const

/*
Returns true if the array contains an occurrence of value, otherwise false.

See also count().
*/
func (this *QJsonArray) Contains(element QJsonValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if element != nil && element.QJsonValue_PTR() != nil {
		convArg0 = element.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray8containsERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef operator[](int)

/*

 */
func (this *QJsonArray) Operator_get_index(i int) *QJsonValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:113
// index:1
// Public Visibility=Default Availability=Available
// [24] QJsonValue operator[](int) const

/*

 */
func (this *QJsonArray) Operator_get_index1(i int) *QJsonValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonArray &) const

/*

 */
func (this *QJsonArray) Operator_equal_equal(other QJsonArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonArray &) const

/*

 */
func (this *QJsonArray) Operator_not_equal(other QJsonArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonArray &)

/*
Swaps the array other with this. This operation is very fast and never fails.

This function was introduced in  Qt 5.10.
*/
func (this *QJsonArray) Swap(other QJsonArray_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonArray_PTR() != nil {
		convArg0 = other.QJsonArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator begin()

/*
Returns an STL-style iterator pointing to the first item in the array.

See also constBegin() and end().
*/
func (this *QJsonArray) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:215
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator begin() const

/*
Returns an STL-style iterator pointing to the first item in the array.

See also constBegin() and end().
*/
func (this *QJsonArray) Begin1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constBegin() const

/*
Returns a const STL-style iterator pointing to the first item in the array.

See also begin() and constEnd().
*/
func (this *QJsonArray) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator end()

/*
Returns an STL-style iterator pointing to the imaginary item after the last item in the array.

See also begin() and constEnd().
*/
func (this *QJsonArray) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator end() const

/*
Returns an STL-style iterator pointing to the imaginary item after the last item in the array.

See also begin() and constEnd().
*/
func (this *QJsonArray) End1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constEnd() const

/*
Returns a const STL-style iterator pointing to the imaginary item after the last item in the array.

See also constBegin() and end().
*/
func (this *QJsonArray) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray operator+(const QJsonValue &) const

/*

 */
func (this *QJsonArray) Operator_add(v QJsonValue_ITF) *QJsonArray /*123*/ {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArrayplERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator+=(const QJsonValue &)

/*

 */
func (this *QJsonArray) Operator_add_equal(v QJsonValue_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArraypLERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray & operator<<(const QJsonValue &)

/*

 */
func (this *QJsonArray) Operator_left_shift(v QJsonValue_ITF) *QJsonArray {
	var convArg0 unsafe.Pointer
	if v != nil && v.QJsonValue_PTR() != nil {
		convArg0 = v.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArraylsERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QJsonValue &)

/*
This function is provided for STL compatibility. It is equivalent to append(value) and will append value to the array.
*/
func (this *QJsonArray) Push_back(t QJsonValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QJsonValue_PTR() != nil {
		convArg0 = t.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray9push_backERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QJsonValue &)

/*
This function is provided for STL compatibility. It is equivalent to prepend(value) and will prepend value to the array.
*/
func (this *QJsonArray) Push_front(t QJsonValue_ITF) {
	var convArg0 unsafe.Pointer
	if t != nil && t.QJsonValue_PTR() != nil {
		convArg0 = t.QJsonValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray10push_frontERK10QJsonValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_front()

/*
This function is provided for STL compatibility. It is equivalent to removeFirst(). The array must not be empty. If the array can be empty, call isEmpty() before calling this function.
*/
func (this *QJsonArray) Pop_front() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray9pop_frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_back()

/*
This function is provided for STL compatibility. It is equivalent to removeLast(). The array must not be empty. If the array can be empty, call isEmpty() before calling this function.
*/
func (this *QJsonArray) Pop_back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QJsonArray8pop_backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty() const

/*
This function is provided for STL compatibility. It is equivalent to isEmpty() and returns true if the array is empty.
*/
func (this *QJsonArray) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QJsonArray5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10443() {
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
