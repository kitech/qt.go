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

// /usr/include/qt/QtCore/qbytearray.h:171
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

// /usr/include/qt/QtCore/qbytearray.h:172
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit1(arg0 string, size int) *QByteArray {
	return NewQByteArray1(arg0, size)
}
func NewQByteArray1(arg0 string, size int) *QByteArray {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:172
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit1p(arg0 string) *QByteArray {
	return NewQByteArray1p(arg0)
}
func NewQByteArray1p(arg0 string) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:173
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(int, char)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit2(size int, c byte) *QByteArray {
	return NewQByteArray2(size, c)
}
func NewQByteArray2(size int, c byte) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2Eic", qtrt.FFI_TYPE_POINTER, size, c)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:174
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(int, Qt::Initialization)

/*
Constructs an empty byte array.

See also isEmpty().
*/
func (*QByteArray) NewForInherit3(size int, arg1 int) *QByteArray {
	return NewQByteArray3(size, arg1)
}
func NewQByteArray3(size int, arg1 int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayC2EiN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, size, arg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:176
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

// /usr/include/qt/QtCore/qbytearray.h:178
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

// /usr/include/qt/QtCore/qbytearray.h:179
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray & operator=(const char *)

/*

 */
func (this *QByteArray) Operator_equal1(str string) *QByteArray {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayaSEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:182
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator=(QByteArray &&)

/*

 */
func (this *QByteArray) Operator_equal2(other unsafe.Pointer /*333*/) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:186
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

// /usr/include/qt/QtCore/qbytearray.h:189
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

// /usr/include/qt/QtCore/qbytearray.h:190
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

// /usr/include/qt/QtCore/qbytearray.h:191
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

// /usr/include/qt/QtCore/qbytearray.h:193
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

// /usr/include/qt/QtCore/qbytearray.h:193
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
func (this *QByteArray) Fillp(c byte) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4fillEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, size)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:195
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

// /usr/include/qt/QtCore/qbytearray.h:196
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

// /usr/include/qt/QtCore/qbytearray.h:197
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

// /usr/include/qt/QtCore/qbytearray.h:203
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

// /usr/include/qt/QtCore/qbytearray.h:204
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
func (this *QByteArray) Data1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:205
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

// /usr/include/qt/QtCore/qbytearray.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QByteArray) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:207
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

// /usr/include/qt/QtCore/qbytearray.h:208
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

// /usr/include/qt/QtCore/qbytearray.h:209
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

// /usr/include/qt/QtCore/qbytearray.h:211
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

// /usr/include/qt/QtCore/qbytearray.h:212
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

// /usr/include/qt/QtCore/qbytearray.h:213
// index:1
// Public inline Visibility=Default Availability=Available
// [1] char operator[](uint) const

/*

 */
func (this *QByteArray) Operator_get_index1(i uint) byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:214
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QByteRef operator[](int)

/*

 */
func (this *QByteArray) Operator_get_index2(i int) *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:215
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QByteRef operator[](uint)

/*

 */
func (this *QByteArray) Operator_get_index3(i uint) *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArrayixEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:216
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

// /usr/include/qt/QtCore/qbytearray.h:217
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
func (this *QByteArray) Front1() *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:218
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

// /usr/include/qt/QtCore/qbytearray.h:219
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
func (this *QByteArray) Back1() *QByteRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:221
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

// /usr/include/qt/QtCore/qbytearray.h:221
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
func (this *QByteArray) IndexOfp(c byte) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:222
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
func (this *QByteArray) IndexOf1(c string, from int) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:222
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
func (this *QByteArray) IndexOf1p(c string) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:223
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
func (this *QByteArray) IndexOf2(a QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:223
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
func (this *QByteArray) IndexOf2p(a QByteArray_ITF) int {
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

// /usr/include/qt/QtCore/qbytearray.h:338
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
func (this *QByteArray) IndexOf3(s string, from int) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:338
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
func (this *QByteArray) IndexOf3p(s string) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:224
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

// /usr/include/qt/QtCore/qbytearray.h:224
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
func (this *QByteArray) LastIndexOfp(c byte) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:225
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
func (this *QByteArray) LastIndexOf1(c string, from int) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:225
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
func (this *QByteArray) LastIndexOf1p(c string) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:226
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
func (this *QByteArray) LastIndexOf2(a QByteArray_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:226
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
func (this *QByteArray) LastIndexOf2p(a QByteArray_ITF) int {
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

// /usr/include/qt/QtCore/qbytearray.h:339
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
func (this *QByteArray) LastIndexOf3(s string, from int) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:339
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
func (this *QByteArray) LastIndexOf3p(s string) int {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:228
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

// /usr/include/qt/QtCore/qbytearray.h:229
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const char *) const

/*
Returns true if the byte array contains an occurrence of the byte array ba; otherwise returns false.

See also indexOf() and count().
*/
func (this *QByteArray) Contains1(a string) bool {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8containsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:230
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QByteArray &) const

/*
Returns true if the byte array contains an occurrence of the byte array ba; otherwise returns false.

See also indexOf() and count().
*/
func (this *QByteArray) Contains2(a QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8containsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:231
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

// /usr/include/qt/QtCore/qbytearray.h:232
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const char *) const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count1(a string) int {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:233
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(const QByteArray &) const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count2(a QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:440
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of (potentially overlapping) occurrences of byte array ba in this byte array.

See also contains() and indexOf().
*/
func (this *QByteArray) Count3() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int compare(const char *, Qt::CaseSensitivity) const

/*
Returns an integer less than, equal to, or greater than zero depending on whether this QByteArray sorts before, at the same position, or after the string pointed to by c. The comparison is performed according to case sensitivity cs.

This function was introduced in  Qt 5.12.

See also operator==.
*/
func (this *QByteArray) Compare(c string, cs int) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7compareEPKcN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int compare(const char *, Qt::CaseSensitivity) const

/*
Returns an integer less than, equal to, or greater than zero depending on whether this QByteArray sorts before, at the same position, or after the string pointed to by c. The comparison is performed according to case sensitivity cs.

This function was introduced in  Qt 5.12.

See also operator==.
*/
func (this *QByteArray) Comparep(c string) int {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7compareEPKcN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const

/*
Returns an integer less than, equal to, or greater than zero depending on whether this QByteArray sorts before, at the same position, or after the string pointed to by c. The comparison is performed according to case sensitivity cs.

This function was introduced in  Qt 5.12.

See also operator==.
*/
func (this *QByteArray) Compare1(a QByteArray_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const

/*
Returns an integer less than, equal to, or greater than zero depending on whether this QByteArray sorts before, at the same position, or after the string pointed to by c. The comparison is performed according to case sensitivity cs.

This function was introduced in  Qt 5.12.

See also operator==.
*/
func (this *QByteArray) Compare1p(a QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:238
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

// /usr/include/qt/QtCore/qbytearray.h:239
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

// /usr/include/qt/QtCore/qbytearray.h:240
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

// /usr/include/qt/QtCore/qbytearray.h:240
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
func (this *QByteArray) Midp(index int) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	len_ := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:241
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

// /usr/include/qt/QtCore/qbytearray.h:244
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

// /usr/include/qt/QtCore/qbytearray.h:245
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
func (this *QByteArray) StartsWith1(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:246
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
func (this *QByteArray) StartsWith2(c string) bool {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:248
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

// /usr/include/qt/QtCore/qbytearray.h:249
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
func (this *QByteArray) EndsWith1(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:250
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
func (this *QByteArray) EndsWith2(c string) bool {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:252
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUpper() const

/*
Returns true if this byte array contains only uppercase letters, otherwise returns false. The byte array is interpreted as a Latin-1 encoded string.

This function was introduced in  Qt 5.12.

See also isLower() and toUpper().
*/
func (this *QByteArray) IsUpper() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7isUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLower() const

/*
Returns true if this byte array contains only lowercase letters, otherwise returns false. The byte array is interpreted as a Latin-1 encoded string.

This function was introduced in  Qt 5.12.

See also isUpper() and toLower().
*/
func (this *QByteArray) IsLower() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7isLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:255
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

// /usr/include/qt/QtCore/qbytearray.h:256
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

// /usr/include/qt/QtCore/qbytearray.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLower() const

/*
Returns a lowercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toLower();
  // y == "qt by the qt company"



See also isLower(), toUpper(), and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToLower() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:268
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLower()

/*
Returns a lowercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toLower();
  // y == "qt by the qt company"



See also isLower(), toUpper(), and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToLower1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:270
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUpper() const

/*
Returns an uppercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toUpper();
  // y == "QT BY THE QT COMPANY"



See also isUpper(), toLower(), and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToUpper() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR10QByteArray7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:272
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUpper()

/*
Returns an uppercase copy of the byte array. The bytearray is interpreted as a Latin-1 encoded string.

Example:


  QByteArray x("Qt by THE QT COMPANY");
  QByteArray y = x.toUpper();
  // y == "QT BY THE QT COMPANY"



See also isUpper(), toLower(), and 8-bit Character Comparisons.
*/
func (this *QByteArray) ToUpper1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:274
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

// /usr/include/qt/QtCore/qbytearray.h:276
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
func (this *QByteArray) Trimmed1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:278
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

// /usr/include/qt/QtCore/qbytearray.h:280
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
func (this *QByteArray) Simplified1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO10QByteArray10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:292
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

// /usr/include/qt/QtCore/qbytearray.h:292
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
func (this *QByteArray) LeftJustifiedp(width int) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:292
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
func (this *QByteArray) LeftJustifiedp1(width int, fill byte) *QByteArray /*123*/ {
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:293
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

// /usr/include/qt/QtCore/qbytearray.h:293
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
func (this *QByteArray) RightJustifiedp(width int) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:293
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
func (this *QByteArray) RightJustifiedp1(width int, fill byte) *QByteArray /*123*/ {
	// arg: 2, bool=Bool, =Invalid, , Invalid
	truncate := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, fill, truncate)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:310
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

// /usr/include/qt/QtCore/qbytearray.h:311
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

// /usr/include/qt/QtCore/qbytearray.h:312
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
func (this *QByteArray) Replace1(index int, len_ int, s string, alen int) *QByteArray {
	var convArg2 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, len_, convArg2, alen)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:313
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
func (this *QByteArray) Replace2(index int, len_ int, s QByteArray_ITF) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:314
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
func (this *QByteArray) Replace3(before byte, after string) *QByteArray {
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), before, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:315
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
func (this *QByteArray) Replace4(before byte, after QByteArray_ITF) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:316
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
func (this *QByteArray) Replace5(before string, after string) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:317
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
func (this *QByteArray) Replace6(before string, bsize int, after string, asize int) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:318
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
func (this *QByteArray) Replace7(before QByteArray_ITF, after QByteArray_ITF) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:319
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
func (this *QByteArray) Replace8(before QByteArray_ITF, after string) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:320
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
func (this *QByteArray) Replace9(before string, after QByteArray_ITF) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:321
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
func (this *QByteArray) Replace10(before byte, after byte) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), before, after)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:333
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
func (this *QByteArray) Replace11(before string, after string) *QByteArray {
	var tmpArg0 = NewQString5(before)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:334
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
func (this *QByteArray) Replace12(c byte, after string) *QByteArray {
	var tmpArg1 = NewQString5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:335
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
func (this *QByteArray) Replace13(before string, after QByteArray_ITF) *QByteArray {
	var tmpArg0 = NewQString5(before)
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

// /usr/include/qt/QtCore/qbytearray.h:322
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

// /usr/include/qt/QtCore/qbytearray.h:323
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const char *)

/*

 */
func (this *QByteArray) Operator_add_equal1(s string) *QByteArray {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:324
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const QByteArray &)

/*

 */
func (this *QByteArray) Operator_add_equal2(a QByteArray_ITF) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:337
// index:3
// Public Visibility=Default Availability=Available
// [8] QByteArray & operator+=(const QString &)

/*

 */
func (this *QByteArray) Operator_add_equal3(s string) *QByteArray {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArraypLERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:328
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

// /usr/include/qt/QtCore/qbytearray.h:342
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QString &) const

/*

 */
func (this *QByteArray) Operator_equal_equal(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayeqERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QString &) const

/*

 */
func (this *QByteArray) Operator_not_equal(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayneERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QString &) const

/*

 */
func (this *QByteArray) Operator_less_than(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayltERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QString &) const

/*

 */
func (this *QByteArray) Operator_greater_than(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArraygtERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QString &) const

/*

 */
func (this *QByteArray) Operator_less_than_equal(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArrayleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QString &) const

/*

 */
func (this *QByteArray) Operator_greater_than_equal(s2 string) bool {
	var tmpArg0 = NewQString5(s2)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArraygeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShort(ok *bool, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShortp() int16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the byte array converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToShortp1(ok *bool) int16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShort(ok *bool, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShortp() uint16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the byte array converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUShortp1(ok *bool) uint16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


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

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray str("FF");
  bool ok;
  int hex = str.toInt(&ok, 16);     // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);     // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToIntp() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the byte array converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray str("FF");
  bool ok;
  int hex = str.toInt(&ok, 16);     // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);     // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToIntp1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:353
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUInt(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:353
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUIntp() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:353
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the byte array converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToUIntp1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:354
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


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

// /usr/include/qt/QtCore/qbytearray.h:354
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray str("FF");
  bool ok;
  long hex = str.toLong(&ok, 16);   // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);   // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToLongp() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:354
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the byte array converted to a long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray str("FF");
  bool ok;
  long hex = str.toLong(&ok, 16);   // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);   // dec == 0, ok == false



Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToLongp1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULong(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULongp() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the byte array converted to an unsigned long int using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function was introduced in  Qt 4.1.

See also number().
*/
func (this *QByteArray) ToULongp1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLong(ok *bool, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLongp() int64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the byte array converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToLongLongp1(ok *bool) int64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLong(ok *bool, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLongp() uint64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the byte array converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0.

If base is 0, the base is determined automatically using the following rules: If the byte array begins with "0x", it is assumed to be hexadecimal; if it begins with "0", it is assumed to be octal; otherwise it is assumed to be decimal.

Returns 0 if the conversion fails.

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

See also number().
*/
func (this *QByteArray) ToULongLongp1(ok *bool) uint64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:358
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the byte array converted to a float value.

Returns an infinity if the conversion overflows or 0.0 if the conversion fails for other reasons (e.g. underflow).

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray string("1234.56");
  bool ok;
  float a = string.toFloat(&ok);    // a == 1234.56, ok == true

  string = "1234.56 Volt";
  a = str.toFloat(&ok);              // a == 0, ok == false



Warning: The QByteArray content may only contain valid numerical characters which includes the plus/minus sign, the character e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function ignores leading and trailing whitespace.

See also number().
*/
func (this *QByteArray) ToFloat(ok *bool) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:358
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the byte array converted to a float value.

Returns an infinity if the conversion overflows or 0.0 if the conversion fails for other reasons (e.g. underflow).

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray string("1234.56");
  bool ok;
  float a = string.toFloat(&ok);    // a == 1234.56, ok == true

  string = "1234.56 Volt";
  a = str.toFloat(&ok);              // a == 0, ok == false



Warning: The QByteArray content may only contain valid numerical characters which includes the plus/minus sign, the character e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function ignores leading and trailing whitespace.

See also number().
*/
func (this *QByteArray) ToFloatp() float32 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:359
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the byte array converted to a double value.

Returns an infinity if the conversion overflows or 0.0 if the conversion fails for other reasons (e.g. underflow).

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray string("1234.56");
  bool ok;
  double a = string.toDouble(&ok);   // a == 1234.56, ok == true

  string = "1234.56 Volt";
  a = str.toDouble(&ok);             // a == 0, ok == false



Warning: The QByteArray content may only contain valid numerical characters which includes the plus/minus sign, the character e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function ignores leading and trailing whitespace.

See also number().
*/
func (this *QByteArray) ToDouble(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:359
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the byte array converted to a double value.

Returns an infinity if the conversion overflows or 0.0 if the conversion fails for other reasons (e.g. underflow).

If ok is not nullptr, failure is reported by setting *ok to false, and success by setting *ok to true.


  QByteArray string("1234.56");
  bool ok;
  double a = string.toDouble(&ok);   // a == 1234.56, ok == true

  string = "1234.56 Volt";
  a = str.toDouble(&ok);             // a == 0, ok == false



Warning: The QByteArray content may only contain valid numerical characters which includes the plus/minus sign, the character e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.

Note: The conversion of the number is performed in the default C locale, irrespective of the user's locale.

This function ignores leading and trailing whitespace.

See also number().
*/
func (this *QByteArray) ToDoublep() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:360
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

// /usr/include/qt/QtCore/qbytearray.h:361
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
func (this *QByteArray) ToBase641() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray8toBase64Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:362
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

// /usr/include/qt/QtCore/qbytearray.h:363
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray toHex(char) const

/*
Returns a hex encoded copy of the byte array. The hex encoding uses the numbers 0-9 and the letters a-f.

See also fromHex().
*/
func (this *QByteArray) ToHex1(separator byte) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5toHexEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), separator)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:364
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

// /usr/include/qt/QtCore/qbytearray.h:364
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
func (this *QByteArray) ToPercentEncodingp() *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:364
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
func (this *QByteArray) ToPercentEncodingp1(exclude QByteArray_ITF) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:364
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
func (this *QByteArray) ToPercentEncodingp2(exclude QByteArray_ITF, include_ QByteArray_ITF) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:368
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

// /usr/include/qt/QtCore/qbytearray.h:368
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
func (this *QByteArray) SetNump(arg0 int16) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEsi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:369
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
func (this *QByteArray) SetNum1(arg0 uint16, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:369
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
func (this *QByteArray) SetNum1p(arg0 uint16) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:370
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
func (this *QByteArray) SetNum2(arg0 int, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:370
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
func (this *QByteArray) SetNum2p(arg0 int) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:371
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
func (this *QByteArray) SetNum3(arg0 uint, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:371
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
func (this *QByteArray) SetNum3p(arg0 uint) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:372
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
func (this *QByteArray) SetNum4(arg0 int64, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:372
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
func (this *QByteArray) SetNum4p(arg0 int64) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:373
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
func (this *QByteArray) SetNum5(arg0 uint64, base int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:373
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
func (this *QByteArray) SetNum5p(arg0 uint64) *QByteArray {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:374
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
func (this *QByteArray) SetNum6(arg0 float32, f byte, prec int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:374
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
func (this *QByteArray) SetNum6p(arg0 float32) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:374
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
func (this *QByteArray) SetNum6p1(arg0 float32, f byte) *QByteArray {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:375
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
func (this *QByteArray) SetNum7(arg0 float64, f byte, prec int) *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:375
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
func (this *QByteArray) SetNum7p(arg0 float64) *QByteArray {
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

// /usr/include/qt/QtCore/qbytearray.h:375
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
func (this *QByteArray) SetNum7p1(arg0 float64, f byte) *QByteArray {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:376
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

// /usr/include/qt/QtCore/qbytearray.h:378
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

// /usr/include/qt/QtCore/qbytearray.h:378
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
func (this *QByteArray) Numberp(arg0 int) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEii", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:379
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
func (this *QByteArray) Number1(arg0 uint, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEji", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number1(arg0 uint, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number1(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:379
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
func (this *QByteArray) Number1p(arg0 uint) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEji", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:380
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
func (this *QByteArray) Number2(arg0 int64, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberExi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number2(arg0 int64, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number2(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:380
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
func (this *QByteArray) Number2p(arg0 int64) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberExi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:381
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
func (this *QByteArray) Number3(arg0 uint64, base int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number3(arg0 uint64, base int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number3(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:381
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
func (this *QByteArray) Number3p(arg0 uint64) *QByteArray /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", qtrt.FFI_TYPE_POINTER, arg0, base)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:382
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
func (this *QByteArray) Number4(arg0 float64, f byte, prec int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QByteArray_Number4(arg0 float64, f byte, prec int) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.Number4(arg0, f, prec)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:382
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
func (this *QByteArray) Number4p(arg0 float64) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:382
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
func (this *QByteArray) Number4p1(arg0 float64, f byte) *QByteArray /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	prec := int(6)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:383
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray fromRawData(const char *, int)

/*
Constructs a QByteArray that uses the first size bytes of the data array. The bytes are not copied. The QByteArray will contain the data pointer. The caller guarantees that data will not be deleted or modified as long as this QByteArray and any copies of it exist that have not been modified. In other words, because QByteArray is an implicitly shared class and the instance returned by this function contains the data pointer, the caller must not delete data or modify it directly as long as the returned QByteArray and any copies exist. However, QByteArray does not take ownership of data, so the QByteArray destructor will never delete the raw data, even when the last QByteArray referring to data is destroyed.

A subsequent attempt to modify the contents of the returned QByteArray or any copy made from it will cause it to create a deep copy of the data array before doing the modification. This ensures that the raw data array itself will never be modified by QByteArray.

Here is an example of how to read data using a QDataStream on raw data in memory without copying the raw data into a QByteArray:


   static const char mydata[] = {
      '\x00', '\x00', '\x03', '\x84', '\x78', '\x9c', '\x3b', '\x76',
      '\xec', '\x18', '\xc3', '\x31', '\x0a', '\xf1', '\xcc', '\x99',
      ...
      '\x6d', '\x5b'
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

// /usr/include/qt/QtCore/qbytearray.h:384
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

// /usr/include/qt/QtCore/qbytearray.h:385
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
func (this *QByteArray) FromBase641(base64 QByteArray_ITF) *QByteArray /*123*/ {
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
func QByteArray_FromBase641(base64 QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QByteArray
	rv := nilthis.FromBase641(base64)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:386
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

// /usr/include/qt/QtCore/qbytearray.h:387
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

// /usr/include/qt/QtCore/qbytearray.h:387
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
func (this *QByteArray) FromPercentEncodingp(pctEncoded QByteArray_ITF) *QByteArray /*123*/ {
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

// /usr/include/qt/QtCore/qbytearray.h:406
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

// /usr/include/qt/QtCore/qbytearray.h:407
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator begin() const

/*
Returns an STL-style iterator pointing to the first character in the byte-array.

See also constBegin() and end().
*/
func (this *QByteArray) Begin1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:408
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

// /usr/include/qt/QtCore/qbytearray.h:409
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

// /usr/include/qt/QtCore/qbytearray.h:410
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

// /usr/include/qt/QtCore/qbytearray.h:411
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray::const_iterator end() const

/*
Returns an STL-style iterator pointing to the imaginary character after the last character in the byte-array.

See also begin() and constEnd().
*/
func (this *QByteArray) End1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QByteArray3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:412
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

// /usr/include/qt/QtCore/qbytearray.h:413
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

// /usr/include/qt/QtCore/qbytearray.h:429
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

// /usr/include/qt/QtCore/qbytearray.h:430
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const char *)

/*
This function is provided for STL compatibility. It is equivalent to append(other).
*/
func (this *QByteArray) Push_back1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray9push_backEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:431
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QByteArray &)

/*
This function is provided for STL compatibility. It is equivalent to append(other).
*/
func (this *QByteArray) Push_back2(a QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray9push_backERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:432
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

// /usr/include/qt/QtCore/qbytearray.h:433
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const char *)

/*
This function is provided for STL compatibility. It is equivalent to prepend(other).
*/
func (this *QByteArray) Push_front1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10push_frontEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:434
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QByteArray &)

/*
This function is provided for STL compatibility. It is equivalent to prepend(other).
*/
func (this *QByteArray) Push_front2(a QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray10push_frontERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:435
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

// /usr/include/qt/QtCore/qbytearray.h:438
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

// /usr/include/qt/QtCore/qbytearray.h:441
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

// /usr/include/qt/QtCore/qbytearray.h:442
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

func init_unused_10179() {
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
