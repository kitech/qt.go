package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QByteArray struct {
	*qtrt.CObject
}
type QByteArray_ITF interface {
	QByteArray_PTR() *QByteArray
}

func (ptr *QByteArray) QByteArray_PTR() *QByteArray { return ptr }

func (this *QByteArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQByteArrayFromPointer(cthis unsafe.Pointer) *QByteArray {
	return &QByteArray{&qtrt.CObject{cthis}}
}
func (*QByteArray) NewFromPointer(cthis unsafe.Pointer) *QByteArray {
	return NewQByteArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearray.h:170
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QByteArray()

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit() *QByteArray {
	return NewQByteArray()
}
func NewQByteArray() *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit_1(arg0 string, size int) *QByteArray {
	return NewQByteArray_1(arg0, size)
}
func NewQByteArray_1(arg0 string, size int) *QByteArray {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit_1_(arg0 string) *QByteArray {
	return NewQByteArray_1_(arg0)
}
func NewQByteArray_1_(arg0 string) *QByteArray {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:172
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(int, char)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit_2(size int, c byte) *QByteArray {
	return NewQByteArray_2(size, c)
}
func NewQByteArray_2(size int, c byte) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2Eic", qtrt.FFI_TYPE_POINTER, size, c)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:173
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(int, Qt::Initialization)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit_3(size int, arg1 int) *QByteArray {
	return NewQByteArray_3(size, arg1)
}
func NewQByteArray_3(size int, arg1 int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2EiN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, size, arg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QByteArray()

/*

 */
func DeleteQByteArray(this *QByteArray) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qbytearray.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & operator=(const QByteArray &)

/*

 */
func (this *QByteArray) Operator_equal(arg0 QByteArray_ITF) *QByteArray {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:178
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray & operator=(const char *)

/*

 */
func (this *QByteArray) Operator_equal_1(str string) *QByteArray {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayaSEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:181
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator=(QByteArray &&)

/*

 */
func (this *QByteArray) Operator_equal_2(other unsafe.Pointer /*333*/) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:185
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QByteArray &)

/*
Swaps byte array other with this byte array. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QByteArray) Swap(other QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QByteArray_PTR() != nil {
		convArg0 = other.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:188
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of bytes in this byte array.

The last byte in the byte array is at position size() - 1. In addition, QByteArray ensures that the byte at position size() is always '\0', so that you can use the return value of data() and constData() as arguments to functions that expect '\0'-terminated strings. If the QByteArray object was created from a raw data that didn't include the trailing null-termination character then QByteArray doesn't add it automaticall unless the deep copy is created.

Example:


  QByteArray ba("Hello");
  int n = ba.size();          // n == 5
  ba.data()[0];               // returns 'H'
  ba.data()[4];               // returns 'o'
  ba.data()[5];               // returns '\0'



See also isEmpty() and resize().
*/
func (this *QByteArray) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the byte array has size 0; otherwise returns false.

Example:


  QByteArray().isEmpty();         // returns true
  QByteArray("").isEmpty();       // returns true
  QByteArray("abc").isEmpty();    // returns false



See also size().
*/
func (this *QByteArray) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int)

/*
Sets the size of the byte array to size bytes.

If size is greater than the current size, the byte array is extended to make it size bytes with the extra bytes added to the end. The new bytes are uninitialized.

If size is less than the current size, bytes are removed from the end.

See also size() and truncate().
*/
func (this *QByteArray) Resize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6resizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & fill(char, int)

/*
Sets every byte in the byte array to character ch. If size is different from -1 (the default), the byte array is resized to size size beforehand.

Example:


  QByteArray ba("Istambul");
  ba.fill('o');
  // ba == "oooooooo"

  ba.fill('X', 2);
  // ba == "XX"



See also resize().
*/
func (this *QByteArray) Fill(c byte, size int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4fillEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, size)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & fill(char, int)

/*
Sets every byte in the byte array to character ch. If size is different from -1 (the default), the byte array is resized to size size beforehand.

Example:


  QByteArray ba("Istambul");
  ba.fill('o');
  // ba == "oooooooo"

  ba.fill('X', 2);
  // ba == "XX"



See also resize().
*/
func (this *QByteArray) Fill__(c byte) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4fillEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, size)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:194
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int capacity() const

/*
Returns the maximum number of bytes that can be stored in the byte array without forcing a reallocation.

The sole purpose of this function is to provide a means of fine tuning QByteArray's memory usage. In general, you will rarely ever need to call this function. If you want to know how many bytes are in the byte array, call size().

See also reserve() and squeeze().
*/
func (this *QByteArray) Capacity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8capacityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:195
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void reserve(int)

/*
Attempts to allocate memory for at least size bytes. If you know in advance how large the byte array will be, you can call this function, and if you call resize() often you are likely to get better performance. If size is an underestimate, the worst that will happen is that the QByteArray will be a bit slower.

The sole purpose of this function is to provide a means of fine tuning QByteArray's memory usage. In general, you will rarely ever need to call this function. If you want to change the size of the byte array, call resize().

See also squeeze() and capacity().
*/
func (this *QByteArray) Reserve(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7reserveEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void squeeze()

/*
Releases any memory not required to store the array's data.

The sole purpose of this function is to provide a means of fine tuning QByteArray's memory usage. In general, you will rarely ever need to call this function.

See also reserve() and capacity().
*/
func (this *QByteArray) Squeeze() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7squeezeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [8] char * data()

/*
Returns a pointer to the data stored in the byte array. The pointer can be used to access and modify the bytes that compose the array. The data is '\0'-terminated, i.e. the number of bytes in the returned character string is size() + 1 for the '\0' terminator.

Example:


  QByteArray ba("Hello world");
  char *data = ba.data();
  while (*data) {
      cout << "[" << *data << "]" << endl;
      ++data;
  }



The pointer remains valid as long as the byte array isn't reallocated or destroyed. For read-only access, constData() is faster because it never causes a deep copy to occur.

This function is mostly useful to pass a byte array to a function that accepts a const char *.

The following example makes a copy of the char* returned by data(), but it will corrupt the heap and cause a crash because it does not allocate a byte for the '\0' at the end:


  QString tmp = "test";
  QByteArray text = tmp.toLocal8Bit();
  char *data = new char[text.size()];
  strcpy(data, text.data());
  delete [] data;



This one allocates the correct amount of space:


  QString tmp = "test";
  QByteArray text = tmp.toLocal8Bit();
  char *data = new char[text.size() + 1];
  strcpy(data, text.data());
  delete [] data;



Note: A QByteArray can store any byte values including '\0's, but most functions that take char * arguments assume that the data ends at the first '\0' they encounter.

See also constData() and operator[]().
*/
func (this *QByteArray) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:203
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const char * data() const

/*
Returns a pointer to the data stored in the byte array. The pointer can be used to access and modify the bytes that compose the array. The data is '\0'-terminated, i.e. the number of bytes in the returned character string is size() + 1 for the '\0' terminator.

Example:


  QByteArray ba("Hello world");
  char *data = ba.data();
  while (*data) {
      cout << "[" << *data << "]" << endl;
      ++data;
  }



The pointer remains valid as long as the byte array isn't reallocated or destroyed. For read-only access, constData() is faster because it never causes a deep copy to occur.

This function is mostly useful to pass a byte array to a function that accepts a const char *.

The following example makes a copy of the char* returned by data(), but it will corrupt the heap and cause a crash because it does not allocate a byte for the '\0' at the end:


  QString tmp = "test";
  QByteArray text = tmp.toLocal8Bit();
  char *data = new char[text.size()];
  strcpy(data, text.data());
  delete [] data;



This one allocates the correct amount of space:


  QString tmp = "test";
  QByteArray text = tmp.toLocal8Bit();
  char *data = new char[text.size() + 1];
  strcpy(data, text.data());
  delete [] data;



Note: A QByteArray can store any byte values including '\0's, but most functions that take char * arguments assume that the data ends at the first '\0' they encounter.

See also constData() and operator[]().
*/
func (this *QByteArray) Data_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:204
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * constData() const

/*
Returns a pointer to the data stored in the byte array. The pointer can be used to access the bytes that compose the array. The data is '\0'-terminated unless the QByteArray object was created from raw data. The pointer remains valid as long as the byte array isn't reallocated or destroyed.

This function is mostly useful to pass a byte array to a function that accepts a const char *.

Note: A QByteArray can store any byte values including '\0's, but most functions that take char * arguments assume that the data ends at the first '\0' they encounter.

See also data(), operator[](), and fromRawData().
*/
func (this *QByteArray) ConstData() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:205
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QByteArray) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QByteArray) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSharedWith(const QByteArray &) const

/*

 */
func (this *QByteArray) IsSharedWith(other QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QByteArray_PTR() != nil {
		convArg0 = other.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray12isSharedWithERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the contents of the byte array and makes it null.

See also resize() and isNull().
*/
func (this *QByteArray) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:210
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char at(int) const

/*
Returns the character at index position i in the byte array.

i must be a valid index position in the byte array (i.e., 0 <= i < size()).

See also operator[]().
*/
func (this *QByteArray) At(i int) byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:211
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char operator[](int) const

/*

 */
func (this *QByteArray) Operator_get_index(i int) byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:212
// index:1
// Public inline Visibility=Default Availability=Available
// [1] char operator[](uint) const

/*

 */
func (this *QByteArray) Operator_get_index_1(i uint) byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:213
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QByteRef operator[](int)

/*

 */
func (this *QByteArray) Operator_get_index_2(i int) *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:214
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QByteRef operator[](uint)

/*

 */
func (this *QByteArray) Operator_get_index_3(i uint) *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char front() const

/*
Returns the first character in the byte array. Same as at(0).

This function is provided for STL compatibility.

Warning: Calling this function on an empty byte array constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also back(), at(), and operator[]().
*/
func (this *QByteArray) Front() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:216
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QByteRef front()

/*
Returns the first character in the byte array. Same as at(0).

This function is provided for STL compatibility.

Warning: Calling this function on an empty byte array constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also back(), at(), and operator[]().
*/
func (this *QByteArray) Front_1() *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char back() const

/*
Returns the last character in the byte array. Same as at(size() - 1).

This function is provided for STL compatibility.

Warning: Calling this function on an empty byte array constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also front(), at(), and operator[]().
*/
func (this *QByteArray) Back() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QByteRef back()

/*
Returns the last character in the byte array. Same as at(size() - 1).

This function is provided for STL compatibility.

Warning: Calling this function on an empty byte array constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also front(), at(), and operator[]().
*/
func (this *QByteArray) Back_1() *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:220
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(char, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf(c byte, from int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:220
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(char, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf__(c byte) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:221
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(const char *, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_1(c string, from int) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:221
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(const char *, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_1_(c string) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:222
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QByteArray &, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_2(a QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:222
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QByteArray &, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_2_(a QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:331
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_3(s string, from int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:331
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int) const

/*
Returns the index position of the first occurrence of the byte array ba in this byte array, searching forward from index position from. Returns -1 if ba could not be found.

Example:


  QByteArray x("sticky question");
  QByteArray y("sti");
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



See also lastIndexOf(), contains(), and count().
*/
func (this *QByteArray) IndexOf_3_(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(char, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf(c byte, from int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(char, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf__(c byte) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:224
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const char *, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_1(c string, from int) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:224
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const char *, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_1_(c string) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:225
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QByteArray &, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_2(a QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:225
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QByteArray &, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_2_(a QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:332
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_3(s string, from int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:332
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int) const

/*
Returns the index position of the last occurrence of the byte array ba in this byte array, searching backward from index position from. If from is -1 (the default), the search starts at the last byte. Returns -1 if ba could not be found.

Example:


  QByteArray x("crazy azimuths");
  QByteArray y("az");
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QByteArray) LastIndexOf_3_(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(char) const

/*
Returns true if the byte array contains an occurrence of the byte array ba; otherwise returns false.

See also indexOf() and count().
*/
func (this *QByteArray) Contains(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8containsEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:228
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const char *) const

/*
Returns true if the byte array contains an occurrence of the byte array ba; otherwise returns false.

See also indexOf() and count().
*/
func (this *QByteArray) Contains_1(a string) bool {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8containsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:229
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QByteArray &) const

/*
Returns true if the byte array contains an occurrence of the byte array ba; otherwise returns false.

See also indexOf() and count().
*/
func (this *QByteArray) Contains_2(a QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:230
// index:0
// Public Visibility=Default Availability=Available
// [4] int count(char) const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count(c byte) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:231
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const char *) const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count_1(a string) int {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:232
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(const QByteArray &) const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count_2(a QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:433
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count_3() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray left(int) const

/*
Returns a byte array that contains the leftmost len bytes of this byte array.

The entire byte array is returned if len is greater than size().

Example:


  QByteArray x("Pineapple");
  QByteArray y = x.left(4);
  // y == "Pine"



See also startsWith(), right(), mid(), chopped(), chop(), and truncate().
*/
func (this *QByteArray) Left(len_ int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray right(int) const

/*
Returns a byte array that contains the rightmost len bytes of this byte array.

The entire byte array is returned if len is greater than size().

Example:


  QByteArray x("Pineapple");
  QByteArray y = x.right(5);
  // y == "apple"



See also endsWith(), left(), mid(), chopped(), chop(), and truncate().
*/
func (this *QByteArray) Right(len_ int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray mid(int, int) const

/*
Returns a byte array containing len bytes from this byte array, starting at position pos.

If len is -1 (the default), or pos + len >= size(), returns a byte array containing all bytes starting at position pos until the end of the byte array.

Example:


  QByteArray x("Five pineapples");
  QByteArray y = x.mid(5, 4);     // y == "pine"
  QByteArray z = x.mid(5);        // z == "pineapples"



See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QByteArray) Mid(index int, len_ int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray mid(int, int) const

/*
Returns a byte array containing len bytes from this byte array, starting at position pos.

If len is -1 (the default), or pos + len >= size(), returns a byte array containing all bytes starting at position pos until the end of the byte array.

Example:


  QByteArray x("Five pineapples");
  QByteArray y = x.mid(5, 4);     // y == "pine"
  QByteArray z = x.mid(5);        // z == "pineapples"



See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QByteArray) Mid__(index int) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	len_ := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray chopped(int) const

/*
Returns a byte array that contains the leftmost size() - len bytes of this byte array.

Note: The behavior is undefined if len is negative or greater than size().

This function was introduced in  Qt 5.10.

See also endsWith(), left(), right(), mid(), chop(), and truncate().
*/
func (this *QByteArray) Chopped(len_ int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:240
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QByteArray &) const

/*
Returns true if this byte array starts with byte array ba; otherwise returns false.

Example:


  QByteArray url("ftp://ftp.qt-project.org/");
  if (url.startsWith("ftp:"))
      ...



See also endsWith() and left().
*/
func (this *QByteArray) StartsWith(a QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10startsWithERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:241
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(char) const

/*
Returns true if this byte array starts with byte array ba; otherwise returns false.

Example:


  QByteArray url("ftp://ftp.qt-project.org/");
  if (url.startsWith("ftp:"))
      ...



See also endsWith() and left().
*/
func (this *QByteArray) StartsWith_1(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:242
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const char *) const

/*
Returns true if this byte array starts with byte array ba; otherwise returns false.

Example:


  QByteArray url("ftp://ftp.qt-project.org/");
  if (url.startsWith("ftp:"))
      ...



See also endsWith() and left().
*/
func (this *QByteArray) StartsWith_2(c string) bool {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:244
// index:0
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QByteArray &) const

/*
Returns true if this byte array ends with byte array ba; otherwise returns false.

Example:


  QByteArray url("http://qt-project.org/doc/qt-5.0/qtdoc/index.html");
  if (url.endsWith(".html"))
      ...



See also startsWith() and right().
*/
func (this *QByteArray) EndsWith(a QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8endsWithERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:245
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(char) const

/*
Returns true if this byte array ends with byte array ba; otherwise returns false.

Example:


  QByteArray url("http://qt-project.org/doc/qt-5.0/qtdoc/index.html");
  if (url.endsWith(".html"))
      ...



See also startsWith() and right().
*/
func (this *QByteArray) EndsWith_1(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:246
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const char *) const

/*
Returns true if this byte array ends with byte array ba; otherwise returns false.

Example:


  QByteArray url("http://qt-project.org/doc/qt-5.0/qtdoc/index.html");
  if (url.endsWith(".html"))
      ...



See also startsWith() and right().
*/
func (this *QByteArray) EndsWith_2(c string) bool {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void truncate(int)

/*
Truncates the byte array at index position pos.

If pos is beyond the end of the array, nothing happens.

Example:


  QByteArray ba("Stockholm");
  ba.truncate(5);             // ba == "Stock"



See also chop(), resize(), and left().
*/
func (this *QByteArray) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void chop(int)

/*
Removes n bytes from the end of the byte array.

If n is greater than size(), the result is an empty byte array.

Example:


  QByteArray ba("STARTTLS\r\n");
  ba.chop(2);                 // ba == "STARTTLS"



See also truncate(), resize(), and left().
*/
func (this *QByteArray) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:259
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLower() const

/*
Returns a lowercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toLower();
  // y == "qt by the qt company"



See also toUpper() and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToLower() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:261
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLower()

/*
Returns a lowercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toLower();
  // y == "qt by the qt company"



See also toUpper() and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToLower_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:263
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUpper() const

/*
Returns an uppercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toUpper();
  // y == "QT BY THE QT COMPANY"



See also toLower() and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToUpper() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:265
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUpper()

/*
Returns an uppercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toUpper();
  // y == "QT BY THE QT COMPANY"



See also toLower() and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToUpper_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:267
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray trimmed() const

/*
Returns a byte array that has whitespace removed from the start and the end.

Whitespace means any character for which the standard C++ isspace() function returns true in the C locale. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.

Example:


  QByteArray ba("  lots\t of\nwhitespace\r\n ");
  ba = ba.trimmed();
  // ba == "lots\t of\nwhitespace";



Unlike simplified(), trimmed() leaves internal whitespace alone.

See also simplified().
*/
func (this *QByteArray) Trimmed() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:269
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray trimmed()

/*
Returns a byte array that has whitespace removed from the start and the end.

Whitespace means any character for which the standard C++ isspace() function returns true in the C locale. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.

Example:


  QByteArray ba("  lots\t of\nwhitespace\r\n ");
  ba = ba.trimmed();
  // ba == "lots\t of\nwhitespace";



Unlike simplified(), trimmed() leaves internal whitespace alone.

See also simplified().
*/
func (this *QByteArray) Trimmed_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:271
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray simplified() const

/*
Returns a byte array that has whitespace removed from the start and the end, and which has each sequence of internal whitespace replaced with a single space.

Whitespace means any character for which the standard C++ isspace() function returns true in the C locale. This includes the ASCII isspace() function returns true in the C locale. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.

Example:


  QByteArray ba("  lots\t of\nwhitespace\r\n ");
  ba = ba.simplified();
  // ba == "lots of whitespace";



See also trimmed().
*/
func (this *QByteArray) Simplified() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:273
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray simplified()

/*
Returns a byte array that has whitespace removed from the start and the end, and which has each sequence of internal whitespace replaced with a single space.

Whitespace means any character for which the standard C++ isspace() function returns true in the C locale. This includes the ASCII isspace() function returns true in the C locale. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.

Example:


  QByteArray ba("  lots\t of\nwhitespace\r\n ");
  ba = ba.simplified();
  // ba == "lots of whitespace";



See also trimmed().
*/
func (this *QByteArray) Simplified_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray leftJustified(int, char, bool) const

/*
Returns a byte array of size width that contains this byte array padded by the fill character.

If truncate is false and the size() of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size() of the byte array is more than width, then any bytes in a copy of the byte array after position width are removed, and the copy is returned.

Example:


  QByteArray x("apple");
  QByteArray y = x.leftJustified(8, '.');   // y == "apple..."



See also rightJustified().
*/
func (this *QByteArray) LeftJustified(width int, fill byte, truncate bool) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray leftJustified(int, char, bool) const

/*
Returns a byte array of size width that contains this byte array padded by the fill character.

If truncate is false and the size() of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size() of the byte array is more than width, then any bytes in a copy of the byte array after position width are removed, and the copy is returned.

Example:


  QByteArray x("apple");
  QByteArray y = x.leftJustified(8, '.');   // y == "apple..."



See also rightJustified().
*/
func (this *QByteArray) LeftJustified__(width int) *QByteArray /*123*/ {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	fill := ' '
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray leftJustified(int, char, bool) const

/*
Returns a byte array of size width that contains this byte array padded by the fill character.

If truncate is false and the size() of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size() of the byte array is more than width, then any bytes in a copy of the byte array after position width are removed, and the copy is returned.

Example:


  QByteArray x("apple");
  QByteArray y = x.leftJustified(8, '.');   // y == "apple..."



See also rightJustified().
*/
func (this *QByteArray) LeftJustified__1(width int, fill byte) *QByteArray /*123*/ {
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rightJustified(int, char, bool) const

/*
Returns a byte array of size width that contains the fill character followed by this byte array.

If truncate is false and the size of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size of the byte array is more than width, then the resulting byte array is truncated at position width.

Example:


  QByteArray x("apple");
  QByteArray y = x.rightJustified(8, '.');    // y == "...apple"



See also leftJustified().
*/
func (this *QByteArray) RightJustified(width int, fill byte, truncate bool) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rightJustified(int, char, bool) const

/*
Returns a byte array of size width that contains the fill character followed by this byte array.

If truncate is false and the size of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size of the byte array is more than width, then the resulting byte array is truncated at position width.

Example:


  QByteArray x("apple");
  QByteArray y = x.rightJustified(8, '.');    // y == "...apple"



See also leftJustified().
*/
func (this *QByteArray) RightJustified__(width int) *QByteArray /*123*/ {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	fill := ' '
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rightJustified(int, char, bool) const

/*
Returns a byte array of size width that contains the fill character followed by this byte array.

If truncate is false and the size of the byte array is more than width, then the returned byte array is a copy of this byte array.

If truncate is true and the size of the byte array is more than width, then the resulting byte array is truncated at position width.

Example:


  QByteArray x("apple");
  QByteArray y = x.rightJustified(8, '.');    // y == "...apple"



See also leftJustified().
*/
func (this *QByteArray) RightJustified__1(width int, fill byte) *QByteArray /*123*/ {
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & remove(int, int)

/*
Removes len bytes from the array, starting at index position pos, and returns a reference to the array.

If pos is out of range, nothing happens. If pos is valid, but pos + len is larger than the size of the array, the array is truncated at position pos.

Example:


  QByteArray ba("Montreal");
  ba.remove(1, 4);
  // ba == "Meal"



See also insert() and replace().
*/
func (this *QByteArray) Remove(index int, len_ int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6removeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:304
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(int, int, const char *)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace(index int, len_ int, s string) *QByteArray {
	var convArg2 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:305
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(int, int, const char *, int)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_1(index int, len_ int, s string, alen int) *QByteArray {
	var convArg2 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_, convArg2, alen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:306
// index:2
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(int, int, const QByteArray &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_2(index int, len_ int, s QByteArray_ITF) *QByteArray {
	var convArg2 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg2 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:307
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & replace(char, const char *)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_3(before byte, after string) *QByteArray {
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), before, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:308
// index:4
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(char, const QByteArray &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_4(before byte, after QByteArray_ITF) *QByteArray {
	var convArg1 unsafe.Pointer
	if after != nil && after.QByteArray_PTR() != nil {
		convArg1 = after.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), before, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:309
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & replace(const char *, const char *)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_5(before string, after string) *QByteArray {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:310
// index:6
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(const char *, int, const char *, int)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_6(before string, bsize int, after string, asize int) *QByteArray {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKciS1_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, bsize, convArg2, asize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:311
// index:7
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(const QByteArray &, const QByteArray &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_7(before QByteArray_ITF, after QByteArray_ITF) *QByteArray {
	var convArg0 unsafe.Pointer
	if before != nil && before.QByteArray_PTR() != nil {
		convArg0 = before.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QByteArray_PTR() != nil {
		convArg1 = after.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:312
// index:8
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & replace(const QByteArray &, const char *)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_8(before QByteArray_ITF, after string) *QByteArray {
	var convArg0 unsafe.Pointer
	if before != nil && before.QByteArray_PTR() != nil {
		convArg0 = before.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_PKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:313
// index:9
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(const char *, const QByteArray &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_9(before string, after QByteArray_ITF) *QByteArray {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if after != nil && after.QByteArray_PTR() != nil {
		convArg1 = after.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:314
// index:10
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(char, char)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_10(before byte, after byte) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), before, after)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:326
// index:11
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(const QString &, const char *)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_11(before string, after string) *QByteArray {
	var tmpArg0 = NewQString_5(before)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:327
// index:12
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(char, const QString &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_12(c byte, after string) *QByteArray {
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:328
// index:13
// Public Visibility=Default Availability=Available
// [8] QByteArray & replace(const QString &, const QByteArray &)

/*
Replaces len bytes from index position pos with the byte array after, and returns a reference to this byte array.

Example:


  QByteArray x("Say yes!");
  QByteArray y("no");
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QByteArray) Replace_13(before string, after QByteArray_ITF) *QByteArray {
	var tmpArg0 = NewQString_5(before)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if after != nil && after.QByteArray_PTR() != nil {
		convArg1 = after.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator+=(char)

/*

 */
func (this *QByteArray) Operator_add_equal(c byte) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:316
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const char *)

/*

 */
func (this *QByteArray) Operator_add_equal_1(s string) *QByteArray {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:317
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const QByteArray &)

/*

 */
func (this *QByteArray) Operator_add_equal_2(a QByteArray_ITF) *QByteArray {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:330
// index:3
// Public Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const QString &)

/*

 */
func (this *QByteArray) Operator_add_equal_3(s string) *QByteArray {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:321
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray repeated(int) const

/*
Returns a copy of this byte array repeated the specified number of times.

If times is less than 1, an empty byte array is returned.

Example:


  QByteArray ba("ab");
  ba.repeated(4);             // returns "abababab"



This function was introduced in  Qt 4.5.
*/
func (this *QByteArray) Repeated(times int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8repeatedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), times)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:335
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QString &) const

/*

 */
func (this *QByteArray) Operator_equal_equal(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayeqERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:336
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QString &) const

/*

 */
func (this *QByteArray) Operator_not_equal(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayneERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:337
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QString &) const

/*

 */
func (this *QByteArray) Operator_less_than(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayltERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:338
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QString &) const

/*

 */
func (this *QByteArray) Operator_greater_than(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArraygtERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:339
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QString &) const

/*

 */
func (this *QByteArray) Operator_less_than_equal(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:340
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QString &) const

/*

 */
func (this *QByteArray) Operator_greater_than_equal(s2 string) bool {
	var tmpArg0 = NewQString_5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArraygeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShort(ok *bool, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShort__() int16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShort__1(ok *bool) int16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShort(ok *bool, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShort__() uint16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShort__1(ok *bool) uint16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  int hex = str.toInt(&ok, 16);     // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);     // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToInt(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  int hex = str.toInt(&ok, 16);     // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);     // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToInt__() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  int hex = str.toInt(&ok, 16);     // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);     // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToInt__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUInt(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUInt__() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUInt__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  long hex = str.toLong(&ok, 16);   // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);   // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToLong(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  long hex = str.toLong(&ok, 16);   // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);   // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToLong__() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray str("FF");
  bool ok;
  long hex = str.toLong(&ok, 16);   // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);   // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToLong__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:348
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULong(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:348
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULong__() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:348
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULong__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:349
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLong(ok *bool, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:349
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLong__() int64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:349
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLong__1(ok *bool) int64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLong(ok *bool, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLong__() uint64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLong__1(ok *bool) uint64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the byte array converted to a float value.

Returns 0.0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToFloat(ok *bool) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the byte array converted to a float value.

Returns 0.0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToFloat__() float32 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the byte array converted to a double value.

Returns 0.0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray string("1234.56");
  double a = string.toDouble();   // a == 1234.56



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToDouble(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the byte array converted to a double value.

Returns 0.0 if the conversion fails.

If ok is not 0: if a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QByteArray string("1234.56");
  double a = string.toDouble();   // a == 1234.56



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToDouble__() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:353
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toBase64(QByteArray::Base64Options) const

/*
Returns a copy of the byte array, encoded as Base64.


  QByteArray text("Qt is great!");
  text.toBase64();        // returns "UXQgaXMgZ3JlYXQh"



The algorithm used to encode Base64-encoded data is defined in RFC 4648.

See also fromBase64().
*/
func (this *QByteArray) ToBase64(options int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toBase64E6QFlagsINS_12Base64OptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:354
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toBase64() const

/*
Returns a copy of the byte array, encoded as Base64.


  QByteArray text("Qt is great!");
  text.toBase64();        // returns "UXQgaXMgZ3JlYXQh"



The algorithm used to encode Base64-encoded data is defined in RFC 4648.

See also fromBase64().
*/
func (this *QByteArray) ToBase64_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toBase64Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toHex() const

/*
Returns a hex encoded copy of the byte array. The hex encoding uses the numbers 0-9 and the letters a-f.

See also fromHex().
*/
func (this *QByteArray) ToHex() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toHexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toHex(char) const

/*
Returns a hex encoded copy of the byte array. The hex encoding uses the numbers 0-9 and the letters a-f.

See also fromHex().
*/
func (this *QByteArray) ToHex_1(separator byte) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toHexEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), separator)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QByteArray &, const QByteArray &, char) const

/*
Returns a URI/URL-style percent-encoded copy of this byte array. The percent parameter allows you to override the default '%' character for another.

By default, this function will encode all characters that are not one of the following:

ALPHA ("a" to "z" and "A" to "Z") / DIGIT (0 to 9) / "-" / "." / "_" / "~"

To prevent characters from being encoded pass them to exclude. To force characters to be encoded pass them to include. The percent character is always encoded.

Example:


  QByteArray text = "{a fishy string?}";
  QByteArray ba = text.toPercentEncoding("{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"



The hex encoding uses the numbers 0-9 and the uppercase letters A-F.

This function was introduced in  Qt 4.4.

See also fromPercentEncoding() and QUrl::toPercentEncoding().
*/
func (this *QByteArray) ToPercentEncoding(exclude QByteArray_ITF, include_ QByteArray_ITF, percent byte) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg0 = exclude.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if include_ != nil && include_.QByteArray_PTR() != nil {
		convArg1 = include_.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QByteArray &, const QByteArray &, char) const

/*
Returns a URI/URL-style percent-encoded copy of this byte array. The percent parameter allows you to override the default '%' character for another.

By default, this function will encode all characters that are not one of the following:

ALPHA ("a" to "z" and "A" to "Z") / DIGIT (0 to 9) / "-" / "." / "_" / "~"

To prevent characters from being encoded pass them to exclude. To force characters to be encoded pass them to include. The percent character is always encoded.

Example:


  QByteArray text = "{a fishy string?}";
  QByteArray ba = text.toPercentEncoding("{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"



The hex encoding uses the numbers 0-9 and the uppercase letters A-F.

This function was introduced in  Qt 4.4.

See also fromPercentEncoding() and QUrl::toPercentEncoding().
*/
func (this *QByteArray) ToPercentEncoding__() *QByteArray /*123*/ {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = NewQByteArray()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = NewQByteArray()
	// arg: 2, char=Char_S, =Invalid, , Invalid
	percent := '%'
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QByteArray &, const QByteArray &, char) const

/*
Returns a URI/URL-style percent-encoded copy of this byte array. The percent parameter allows you to override the default '%' character for another.

By default, this function will encode all characters that are not one of the following:

ALPHA ("a" to "z" and "A" to "Z") / DIGIT (0 to 9) / "-" / "." / "_" / "~"

To prevent characters from being encoded pass them to exclude. To force characters to be encoded pass them to include. The percent character is always encoded.

Example:


  QByteArray text = "{a fishy string?}";
  QByteArray ba = text.toPercentEncoding("{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"



The hex encoding uses the numbers 0-9 and the uppercase letters A-F.

This function was introduced in  Qt 4.4.

See also fromPercentEncoding() and QUrl::toPercentEncoding().
*/
func (this *QByteArray) ToPercentEncoding__1(exclude QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg0 = exclude.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = NewQByteArray()
	// arg: 2, char=Char_S, =Invalid, , Invalid
	percent := '%'
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QByteArray &, const QByteArray &, char) const

/*
Returns a URI/URL-style percent-encoded copy of this byte array. The percent parameter allows you to override the default '%' character for another.

By default, this function will encode all characters that are not one of the following:

ALPHA ("a" to "z" and "A" to "Z") / DIGIT (0 to 9) / "-" / "." / "_" / "~"

To prevent characters from being encoded pass them to exclude. To force characters to be encoded pass them to include. The percent character is always encoded.

Example:


  QByteArray text = "{a fishy string?}";
  QByteArray ba = text.toPercentEncoding("{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"



The hex encoding uses the numbers 0-9 and the uppercase letters A-F.

This function was introduced in  Qt 4.4.

See also fromPercentEncoding() and QUrl::toPercentEncoding().
*/
func (this *QByteArray) ToPercentEncoding__2(exclude QByteArray_ITF, include_ QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg0 = exclude.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if include_ != nil && include_.QByteArray_PTR() != nil {
		convArg1 = include_.QByteArray_PTR().GetCthis()
	}
	// arg: 2, char=Char_S, =Invalid, , Invalid
	percent := '%'
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:361
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(short, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum(arg0 int16, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEsi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:361
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(short, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum__(arg0 int16) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEsi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(ushort, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_1(arg0 uint16, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:362
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(ushort, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_1_(arg0 uint16) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:363
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(int, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_2(arg0 int, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:363
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(int, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_2_(arg0 int) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:364
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(uint, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_3(arg0 uint, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:364
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(uint, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_3_(arg0 uint) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:365
// index:4
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(qlonglong, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_4(arg0 int64, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:365
// index:4
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(qlonglong, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_4_(arg0 int64) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:366
// index:5
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(qulonglong, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_5(arg0 uint64, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:366
// index:5
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(qulonglong, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_5_(arg0 uint64) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:367
// index:6
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(float, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_6(arg0 float32, f byte, prec int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:367
// index:6
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(float, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_6_(arg0 float32) *QByteArray {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	f := 'g'
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:367
// index:6
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & setNum(float, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_6_1(arg0 float32, f byte) *QByteArray {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:368
// index:7
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(double, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_7(arg0 float64, f byte, prec int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:368
// index:7
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(double, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_7_(arg0 float64) *QByteArray {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	f := 'g'
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:368
// index:7
// Public Visibility=Default Availability=Available
// [8] QByteArray & setNum(double, char, int)

/*
Sets the byte array to the printed value of n in base base (10 by default) and returns a reference to the byte array. The base can be any value between 2 and 36. For bases other than 10, n is treated as an unsigned integer.

Example:


  QByteArray ba;
  int n = 63;
  ba.setNum(n);           // ba == "63"
  ba.setNum(n, 16);       // ba == "3f"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also number() and toInt().
*/
func (this *QByteArray) SetNum_7_1(arg0 float64, f byte) *QByteArray {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:369
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & setRawData(const char *, uint)

/*
Resets the QByteArray to use the first size bytes of the data array. The bytes are not copied. The QByteArray will contain the data pointer. The caller guarantees that data will not be deleted or modified as long as this QByteArray and any copies of it exist that have not been modified.

This function can be used instead of fromRawData() to re-use existing QByteArray objects to save memory re-allocations.

This function was introduced in  Qt 4.7.

See also fromRawData(), data(), and constData().
*/
func (this *QByteArray) SetRawData(a string, n uint) *QByteArray {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10setRawDataEPKcj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:371
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(int, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number(arg0 int, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEii", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number(arg0 int, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:371
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(int, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number__(arg0 int) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEii", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:372
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(uint, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_1(arg0 uint, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEji", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number_1(arg0 uint, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number_1(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:372
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(uint, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_1_(arg0 uint) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEji", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:373
// index:2
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(qlonglong, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_2(arg0 int64, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberExi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number_2(arg0 int64, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number_2(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:373
// index:2
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(qlonglong, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_2_(arg0 int64) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberExi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:374
// index:3
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(qulonglong, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_3(arg0 uint64, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number_3(arg0 uint64, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number_3(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:374
// index:3
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(qulonglong, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_3_(arg0 uint64) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:375
// index:4
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(double, char, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_4(arg0 float64, f byte, prec int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number_4(arg0 float64, f byte, prec int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number_4(arg0, f, prec)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:375
// index:4
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(double, char, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_4_(arg0 float64) *QByteArray /*123*/ {
	// arg: 1, char=Char_S, =Invalid, , Invalid
	f := 'g'
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:375
// index:4
// Public static Visibility=Default Availability=Available
// [8] QByteArray number(double, char, int)

/*
Returns a byte array containing the string equivalent of the number n to base base (10 by default). The base can be any value between 2 and 36.

Example:


  int n = 63;
  QByteArray::number(n);              // returns "63"
  QByteArray::number(n, 16);          // returns "3f"
  QByteArray::number(n, 16).toUpper();  // returns "3F"



Note: The format of the number is not localized; the default C locale is used irrespective of the user's locale.

See also setNum() and toInt().
*/
func (this *QByteArray) Number_4_1(arg0 float64, f byte) *QByteArray /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:376
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromRawData(const char *, int)

/*
Constructs a QByteArray that uses the first size bytes of the data array. The bytes are not copied. The QByteArray will contain the data pointer. The caller guarantees that data will not be deleted or modified as long as this QByteArray and any copies of it exist that have not been modified. In other words, because QByteArray is an implicitly shared class and the instance returned by this function contains the data pointer, the caller must not delete data or modify it directly as long as the returned QByteArray and any copies exist. However, QByteArray does not take ownership of data, so the QByteArray destructor will never delete the raw data, even when the last QByteArray referring to data is destroyed.

A subsequent attempt to modify the contents of the returned QByteArray or any copy made from it will cause it to create a deep copy of the data array before doing the modification. This ensures that the raw data array itself will never be modified by QByteArray.

Here is an example of how to read data using a QDataStream on raw data in memory without copying the raw data into a QByteArray:


   static const char mydata[] = {
      0x00, 0x00, 0x03, 0x84, 0x78, 0x9c, 0x3b, 0x76,
      0xec, 0x18, 0xc3, 0x31, 0x0a, 0xf1, 0xcc, 0x99,
      ...
      0x6d, 0x5b
  };

  QByteArray data = QByteArray::fromRawData(mydata, sizeof(mydata));
  QDataStream in(&data, QIODevice::ReadOnly);
  ...



Warning: A byte array created with fromRawData() is not null-terminated, unless the raw data contains a 0 character at position size. While that does not matter for QDataStream or functions like indexOf(), passing the byte array to a function accepting a const char * expected to be '\0'-terminated will fail.

See also setRawData(), data(), and constData().
*/
func (this *QByteArray) FromRawData(arg0 string, size int) *QByteArray /*123*/ {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray11fromRawDataEPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_FromRawData(arg0 string, size int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromRawData(arg0, size)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:377
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromBase64(const QByteArray &, QByteArray::Base64Options)

/*
Returns a decoded copy of the Base64 array base64. Input is not checked for validity; invalid characters in the input are skipped, enabling the decoding process to continue with subsequent characters.

For example:


  QByteArray text = QByteArray::fromBase64("UXQgaXMgZ3JlYXQh");
  text.data();            // returns "Qt is great!"



The algorithm used to decode Base64-encoded data is defined in RFC 4648.

See also toBase64().
*/
func (this *QByteArray) FromBase64(base64 QByteArray_ITF, options int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if base64 != nil && base64.QByteArray_PTR() != nil {
		convArg0 = base64.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10fromBase64ERKS_6QFlagsINS_12Base64OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_FromBase64(base64 QByteArray_ITF, options int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromBase64(base64, options)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:378
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromBase64(const QByteArray &)

/*
Returns a decoded copy of the Base64 array base64. Input is not checked for validity; invalid characters in the input are skipped, enabling the decoding process to continue with subsequent characters.

For example:


  QByteArray text = QByteArray::fromBase64("UXQgaXMgZ3JlYXQh");
  text.data();            // returns "Qt is great!"



The algorithm used to decode Base64-encoded data is defined in RFC 4648.

See also toBase64().
*/
func (this *QByteArray) FromBase64_1(base64 QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if base64 != nil && base64.QByteArray_PTR() != nil {
		convArg0 = base64.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10fromBase64ERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_FromBase64_1(base64 QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromBase64_1(base64)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:379
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromHex(const QByteArray &)

/*
Returns a decoded copy of the hex encoded array hexEncoded. Input is not checked for validity; invalid characters in the input are skipped, enabling the decoding process to continue with subsequent characters.

For example:


  QByteArray text = QByteArray::fromHex("517420697320677265617421");
  text.data();            // returns "Qt is great!"



See also toHex().
*/
func (this *QByteArray) FromHex(hexEncoded QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if hexEncoded != nil && hexEncoded.QByteArray_PTR() != nil {
		convArg0 = hexEncoded.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7fromHexERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_FromHex(hexEncoded QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromHex(hexEncoded)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:380
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromPercentEncoding(const QByteArray &, char)

/*
Returns a decoded copy of the URI/URL-style percent-encoded input. The percent parameter allows you to replace the '%' character for another (for instance, '_' or '=').

For example:


  QByteArray text = QByteArray::fromPercentEncoding("Qt%20is%20great%33");
  text.data();            // returns "Qt is great!"



Note: Given invalid input (such as a string containing the sequence "%G5", which is not a valid hexadecimal number) the output will be invalid as well. As an example: the sequence "%G5" could be decoded to 'W'.

This function was introduced in  Qt 4.4.

See also toPercentEncoding() and QUrl::fromPercentEncoding().
*/
func (this *QByteArray) FromPercentEncoding(pctEncoded QByteArray_ITF, percent byte) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if pctEncoded != nil && pctEncoded.QByteArray_PTR() != nil {
		convArg0 = pctEncoded.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray19fromPercentEncodingERKS_c", qtrt.FFI_TYPE_POINTER, convArg0, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_FromPercentEncoding(pctEncoded QByteArray_ITF, percent byte) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromPercentEncoding(pctEncoded, percent)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:380
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromPercentEncoding(const QByteArray &, char)

/*
Returns a decoded copy of the URI/URL-style percent-encoded input. The percent parameter allows you to replace the '%' character for another (for instance, '_' or '=').

For example:


  QByteArray text = QByteArray::fromPercentEncoding("Qt%20is%20great%33");
  text.data();            // returns "Qt is great!"



Note: Given invalid input (such as a string containing the sequence "%G5", which is not a valid hexadecimal number) the output will be invalid as well. As an example: the sequence "%G5" could be decoded to 'W'.

This function was introduced in  Qt 4.4.

See also toPercentEncoding() and QUrl::fromPercentEncoding().
*/
func (this *QByteArray) FromPercentEncoding__(pctEncoded QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if pctEncoded != nil && pctEncoded.QByteArray_PTR() != nil {
		convArg0 = pctEncoded.QByteArray_PTR().GetCthis()
	}
	// arg: 1, char=Char_S, =Invalid, , Invalid
	percent := '%'
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray19fromPercentEncodingERKS_c", qtrt.FFI_TYPE_POINTER, convArg0, percent)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:399
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::iterator begin()

/*
Returns an STL-style iterator pointing to the first character in the byte-array.

See also constBegin() and end().
*/
func (this *QByteArray) Begin() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:400
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator begin() const

/*
Returns an STL-style iterator pointing to the first character in the byte-array.

See also constBegin() and end().
*/
func (this *QByteArray) Begin_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:401
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator cbegin() const

/*
Returns a const STL-style iterator pointing to the first character in the byte-array.

This function was introduced in  Qt 5.0.

See also begin() and cend().
*/
func (this *QByteArray) Cbegin() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:402
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator constBegin() const

/*
Returns a const STL-style iterator pointing to the first character in the byte-array.

See also begin() and constEnd().
*/
func (this *QByteArray) ConstBegin() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:403
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::iterator end()

/*
Returns an STL-style iterator pointing to the imaginary character after the last character in the byte-array.

See also begin() and constEnd().
*/
func (this *QByteArray) End() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:404
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator end() const

/*
Returns an STL-style iterator pointing to the imaginary character after the last character in the byte-array.

See also begin() and constEnd().
*/
func (this *QByteArray) End_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:405
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator cend() const

/*
Returns a const STL-style iterator pointing to the imaginary character after the last character in the list.

This function was introduced in  Qt 5.0.

See also cbegin() and end().
*/
func (this *QByteArray) Cend() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:406
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator constEnd() const

/*
Returns a const STL-style iterator pointing to the imaginary character after the last character in the list.

See also constBegin() and end().
*/
func (this *QByteArray) ConstEnd() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:422
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(char)

/*
This function is provided for STL compatibility. It is equivalent to append(other).
*/
func (this *QByteArray) Push_back(c byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray9push_backEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:423
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const char *)

/*
This function is provided for STL compatibility. It is equivalent to append(other).
*/
func (this *QByteArray) Push_back_1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray9push_backEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:424
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QByteArray &)

/*
This function is provided for STL compatibility. It is equivalent to append(other).
*/
func (this *QByteArray) Push_back_2(a QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray9push_backERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:425
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(char)

/*
This function is provided for STL compatibility. It is equivalent to prepend(other).
*/
func (this *QByteArray) Push_front(c byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10push_frontEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:426
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const char *)

/*
This function is provided for STL compatibility. It is equivalent to prepend(other).
*/
func (this *QByteArray) Push_front_1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10push_frontEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:427
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QByteArray &)

/*
This function is provided for STL compatibility. It is equivalent to prepend(other).
*/
func (this *QByteArray) Push_front_2(a QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10push_frontERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:428
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void shrink_to_fit()

/*
This function is provided for STL compatibility. It is equivalent to squeeze().

This function was introduced in  Qt 5.10.
*/
func (this *QByteArray) Shrink_to_fit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray13shrink_to_fitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:431
// index:0
// Public inline Visibility=Default Availability=Available
// [32] std::string toStdString() const

/*
Returns a std::string object with the data contained in this QByteArray.

This operator is mostly useful to pass a QByteArray to a function that accepts a std::string object.

This function was introduced in  Qt 5.4.

See also fromStdString() and QString::toStdString().
*/
func (this *QByteArray) ToStdString() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toStdStringB5cxx11Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:434
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length() const

/*
Same as size().
*/
func (this *QByteArray) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:435
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this byte array is null; otherwise returns false.

Example:


  QByteArray().isNull();          // returns true
  QByteArray("").isNull();        // returns false
  QByteArray("abc").isNull();     // returns false



Qt makes a distinction between null byte arrays and empty byte arrays for historical reasons. For most applications, what matters is whether or not a byte array contains any data, and this can be determined using isEmpty().

See also isEmpty().
*/
func (this *QByteArray) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QByteArray__Base64Option = int

//
const QByteArray__Base64Encoding QByteArray__Base64Option = 0

//
const QByteArray__Base64UrlEncoding QByteArray__Base64Option = 1

//
const QByteArray__KeepTrailingEquals QByteArray__Base64Option = 0

//
const QByteArray__OmitTrailingEquals QByteArray__Base64Option = 2

func (this *QByteArray) Base64OptionItemName(val int) string {
	switch val {
	case QByteArray__Base64Encoding: // 0
		return "Base64Encoding,KeepTrailingEquals"
	case QByteArray__Base64UrlEncoding: // 1
		return "Base64UrlEncoding"
		// case QByteArray__KeepTrailingEquals: // 0
		// return ""
	case QByteArray__OmitTrailingEquals: // 2
		return "OmitTrailingEquals"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QByteArray_Base64OptionItemName(val int) string {
	var nilthis *QByteArray
	return nilthis.Base64OptionItemName(val)
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
