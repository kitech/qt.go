package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QStringRef struct {
	*qtrt.CObject
}
type QStringRef_ITF interface {
	QStringRef_PTR() *QStringRef
}

func (ptr *QStringRef) QStringRef_PTR() *QStringRef { return ptr }

func (this *QStringRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringRefFromPointer(cthis unsafe.Pointer) *QStringRef {
	return &QStringRef{&qtrt.CObject{cthis}}
}
func (*QStringRef) NewFromPointer(cthis unsafe.Pointer) *QStringRef {
	return NewQStringRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1420
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef()

/*

 */
func (*QStringRef) NewForInherit() *QStringRef {
	return NewQStringRef()
}
func NewQStringRef() *QStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1421
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *, int, int)

/*

 */
func (*QStringRef) NewForInherit_1(string string, position int, size int) *QStringRef {
	return NewQStringRef_1(string, position, size)
}
func NewQStringRef_1(string string, position int, size int) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QStringii", qtrt.FFI_TYPE_POINTER, convArg0, position, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1422
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *)

/*

 */
func (*QStringRef) NewForInherit_2(string string) *QStringRef {
	return NewQStringRef_2(string)
}
func NewQStringRef_2(string string) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1431
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(QStringRef &&)

/*

 */
func (this *QStringRef) Operator_equal(other unsafe.Pointer /*333*/) *QStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1433
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(const QStringRef &)

/*

 */
func (this *QStringRef) Operator_equal_1(other QStringRef_ITF) *QStringRef {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStringRef_PTR() != nil {
		convArg0 = other.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1505
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QStringRef & operator=(const QString *)

/*

 */
func (this *QStringRef) Operator_equal_2(string string) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefaSEPK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1438
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QStringRef()

/*

 */
func DeleteQStringRef(this *QStringRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstring.h:1441
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString * string() const

/*

 */
func (this *QStringRef) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1442
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int position() const

/*

 */
func (this *QStringRef) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1443
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of characters in this string.

The last character in the string is at position size() - 1.

Example:


  QString str = "World";
  int n = str.size();         // n == 5
  str.data()[0];              // returns 'W'
  str.data()[4];              // returns 'd'



See also isEmpty() and resize().
*/
func (this *QStringRef) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1444
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_1(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_1_(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_2(c QChar_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_2_(c QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_3(s QStringRef_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QStringRef) Count_3_(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1445
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length() const

/*
Returns the number of characters in this string. Equivalent to size().

See also resize().
*/
func (this *QStringRef) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf__1(str string, from int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_1(ch QChar_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_1_(ch QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_1_1(ch QChar_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_2(str QLatin1String_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_2_(str QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_2_1(str QLatin1String_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_3(str QStringRef_ITF, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_3_(str QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the first occurrence of the string str in this string, searching forward from index position from. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "sticky question";
  QString y = "sti";
  x.indexOf(y);               // returns 0
  x.indexOf(y, 1);            // returns 10
  x.indexOf(y, 10);           // returns 10
  x.indexOf(y, 11);           // returns -1



If from is -1, the search starts at the last character; if it is -2, at the next to last character and so on.

See also lastIndexOf(), contains(), and count().
*/
func (this *QStringRef) IndexOf_3_1(str QStringRef_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf__(str string) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf__1(str string, from int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_1(ch QChar_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_1_(ch QChar_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_1_1(ch QChar_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_2(str QLatin1String_ITF /*123*/, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_2_(str QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_2_1(str QLatin1String_ITF /*123*/, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_3(str QStringRef_ITF, from int, cs int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_3_(str QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity) const

/*
Returns the index position of the last occurrence of the string str in this string, searching backward from index position from. If from is -1 (default), the search starts at the last character; if from is -2, at the next to last character and so on. Returns -1 if str is not found.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString x = "crazy azimuths";
  QString y = "az";
  x.lastIndexOf(y);           // returns 6
  x.lastIndexOf(y, 6);        // returns 6
  x.lastIndexOf(y, 5);        // returns 2
  x.lastIndexOf(y, 1);        // returns -1



See also indexOf(), contains(), and count().
*/
func (this *QStringRef) LastIndexOf_3_1(str QStringRef_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains__(str string) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_1(ch QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_1_(ch QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_2(str QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_2_(str QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_3(str QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QStringRef) Contains_3_(str QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QStringRef_PTR() != nil {
		convArg0 = str.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1470
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef left(int) const

/*
Returns a substring that contains the n leftmost characters of the string.

The entire string is returned if n is greater than or equal to size(), or less than zero.


  QString x = "Pineapple";
  QString y = x.left(4);      // y == "Pine"



See also right(), mid(), startsWith(), chopped(), chop(), and truncate().
*/
func (this *QStringRef) Left(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1471
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef right(int) const

/*
Returns a substring that contains the n rightmost characters of the string.

The entire string is returned if n is greater than or equal to size(), or less than zero.


  QString x = "Pineapple";
  QString y = x.right(5);      // y == "apple"



See also left(), mid(), endsWith(), chopped(), chop(), and truncate().
*/
func (this *QStringRef) Right(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef mid(int, int) const

/*
Returns a string that contains n characters of this string, starting at the specified position index.

Returns a null string if the position index exceeds the length of the string. If there are less than n characters available in the string starting at the given position, or if n is -1 (default), the function returns all characters that are available from the specified position.

Example:


  QString x = "Nine pineapples";
  QString y = x.mid(5, 4);            // y == "pine"
  QString z = x.mid(5);               // z == "pineapples"



See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QStringRef) Mid(pos int, n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef mid(int, int) const

/*
Returns a string that contains n characters of this string, starting at the specified position index.

Returns a null string if the position index exceeds the length of the string. If there are less than n characters available in the string starting at the given position, or if n is -1 (default), the function returns all characters that are available from the specified position.

Example:


  QString x = "Nine pineapples";
  QString y = x.mid(5, 4);            // y == "pine"
  QString z = x.mid(5);               // z == "pineapples"



See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QStringRef) Mid__(pos int) *QStringRef /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	n := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1473
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef chopped(int) const

/*
Returns a substring that contains the size() - len leftmost characters of this string.

Note: The behavior is undefined if len is negative or greater than size().

This function was introduced in  Qt 5.10.

See also endsWith(), left(), right(), mid(), chop(), and truncate().
*/
func (this *QStringRef) Chopped(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1476
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(int)

/*
Truncates the string at the given position index.

If the specified position index is beyond the end of the string, nothing happens.

Example:


  QString str = "Vladivostok";
  str.truncate(4);
  // str == "Vlad"



If position is negative, it is equivalent to passing zero.

See also chop(), resize(), left(), and QStringRef::truncate().
*/
func (this *QStringRef) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1477
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(int)

/*
Removes n characters from the end of the string.

If n is greater than or equal to size(), the result is an empty string; if n is negative, it is equivalent to passing zero.

Example:


  QString str("LOGOUT\r\n");
  str.chop(2);
  // str == "LOGOUT"



If you want to remove characters from the beginning of the string, use remove() instead.

See also truncate(), resize(), remove(), and QStringRef::chop().
*/
func (this *QStringRef) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1485
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft() const

/*
Returns true if the string is read right to left.

See also QStringRef::isRightToLeft().
*/
func (this *QStringRef) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_2(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_2_(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_3_(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_4(c QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if the string starts with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.startsWith("Ban");     // returns true
  str.startsWith("Car");     // returns false



See also endsWith().
*/
func (this *QStringRef) StartsWith_4_(c QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_2(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_2_(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_3_(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_4(c QStringRef_ITF, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity) const

/*
Returns true if the string ends with s; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.


  QString str = "Bananas";
  str.endsWith("anas");         // returns true
  str.endsWith("pple");         // returns false



See also startsWith().
*/
func (this *QStringRef) EndsWith_4_(c QStringRef_ITF) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QStringRef_PTR() != nil {
		convArg0 = c.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1507
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * unicode() const

/*
Returns a Unicode representation of the string. The result remains valid until the string is modified.

Note: The returned string may not be '\0'-terminated. Use size() to determine the length of the array.

See also setUnicode(), utf16(), and fromRawData().
*/
func (this *QStringRef) Unicode() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1513
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * data() const

/*
Returns a pointer to the data stored in the QString. The pointer can be used to access and modify the characters that compose the string.

Unlike constData() and unicode(), the returned data is always '\0'-terminated.

Example:


  QString str = "Hello world";
  QChar *data = str.data();
  while (!data->isNull()) {
      qDebug() << data->unicode();
      ++data;
  }



Note that the pointer remains valid only as long as the string is not modified by other means. For read-only access, constData() is faster because it never causes a deep copy to occur.

See also constData() and operator[]().
*/
func (this *QStringRef) Data() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1514
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * constData() const

/*
Returns a pointer to the data stored in the QString. The pointer can be used to access the characters that compose the string.

Note that the pointer remains valid only as long as the string is not modified.

Note: The returned string may not be '\0'-terminated. Use size() to determine the length of the array.

See also data(), operator[](), and fromRawData().
*/
func (this *QStringRef) ConstData() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1516
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator begin() const

/*
Returns an STL-style iterator pointing to the first character in the string.

See also constBegin() and end().
*/
func (this *QStringRef) Begin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1517
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cbegin() const

/*
Returns a const STL-style iterator pointing to the first character in the string.

This function was introduced in  Qt 5.0.

See also begin() and cend().
*/
func (this *QStringRef) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1518
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constBegin() const

/*
Returns a const STL-style iterator pointing to the first character in the string.

See also begin() and constEnd().
*/
func (this *QStringRef) ConstBegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1519
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator end() const

/*
Returns an STL-style iterator pointing to the imaginary character after the last character in the string.

See also begin() and constEnd().
*/
func (this *QStringRef) End() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1520
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cend() const

/*
Returns a const STL-style iterator pointing to the imaginary character after the last character in the list.

This function was introduced in  Qt 5.0.

See also cbegin() and end().
*/
func (this *QStringRef) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1521
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constEnd() const

/*
Returns a const STL-style iterator pointing to the imaginary character after the last character in the list.

See also constBegin() and end().
*/
func (this *QStringRef) ConstEnd() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1531
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLatin1() const

/*
Returns a Latin-1 representation of the string as a QByteArray.

The returned byte array is undefined if the string contains non-Latin1 characters. Those characters may be suppressed or replaced with a question mark.

See also fromLatin1(), toUtf8(), toLocal8Bit(), QTextCodec, and qConvertToLatin1().
*/
func (this *QStringRef) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1532
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toUtf8() const

/*
Returns a UTF-8 representation of the string as a QByteArray.

UTF-8 is a Unicode codec and can represent all characters in a Unicode string like QString.

See also fromUtf8(), toLatin1(), toLocal8Bit(), QTextCodec, and qConvertToUtf8().
*/
func (this *QStringRef) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1533
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit() const

/*
Returns the local 8-bit representation of the string as a QByteArray. The returned byte array is undefined if the string contains characters not supported by the local 8-bit encoding.

QTextCodec::codecForLocale() is used to perform the conversion from Unicode. If the locale encoding could not be determined, this function does the same as toLatin1().

If this string contains any characters that cannot be encoded in the locale, the returned byte array is undefined. Those characters may be suppressed or replaced by another.

See also fromLocal8Bit(), toLatin1(), toUtf8(), and QTextCodec.
*/
func (this *QStringRef) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the contents of the string and makes it null.

See also resize() and isNull().
*/
func (this *QStringRef) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1537
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*

 */
func (this *QStringRef) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1538
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the string has no characters; otherwise returns false.

Example:


  QString().isEmpty();            // returns true
  QString("").isEmpty();          // returns true
  QString("x").isEmpty();         // returns false
  QString("abc").isEmpty();       // returns false



See also size().
*/
func (this *QStringRef) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1539
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this string is null; otherwise returns false.

Example:


  QString().isNull();             // returns true
  QString("").isNull();           // returns false
  QString("abc").isNull();        // returns false



Qt makes a distinction between null strings and empty strings for historical reasons. For most applications, what matters is whether or not a string contains any data, and this can be determined using the isEmpty() function.

See also isEmpty().
*/
func (this *QStringRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1541
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef appendTo(QString *) const

/*

 */
func (this *QStringRef) AppendTo(string string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8appendToEP7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1543
// index:0
// Public inline Visibility=Default Availability=Available
// [2] const QChar at(int) const

/*
Returns the character at the given index position in the string.

The position must be a valid index position in the string (i.e., 0 <= position < size()).

See also operator[]().
*/
func (this *QStringRef) At(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1545
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar operator[](int) const

/*

 */
func (this *QStringRef) Operator_get_index(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1546
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front() const

/*
Returns the first character in the string. Same as at(0).

This function is provided for STL compatibility.

Warning: Calling this function on an empty string constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also back(), at(), and operator[]().
*/
func (this *QStringRef) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1547
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back() const

/*
Returns the last character in the string. Same as at(size() - 1).

This function is provided for STL compatibility.

Warning: Calling this function on an empty string constitutes undefined behavior.

This function was introduced in  Qt 5.10.

See also front(), at(), and operator[]().
*/
func (this *QStringRef) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1551
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const char *) const

/*

 */
func (this *QStringRef) Operator_equal_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefeqEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1552
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const char *) const

/*

 */
func (this *QStringRef) Operator_not_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefneEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1553
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const char *) const

/*

 */
func (this *QStringRef) Operator_less_than(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefltEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1554
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const char *) const

/*

 */
func (this *QStringRef) Operator_less_than_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefleEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1555
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const char *) const

/*

 */
func (this *QStringRef) Operator_greater_than(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefgtEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1556
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const char *) const

/*

 */
func (this *QStringRef) Operator_greater_than_equal(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRefgeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare__(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_1(s QStringRef_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_1_(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_2(s QLatin1String_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_2_(s QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_3(s QByteArray_ITF, cs int) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_3_(s QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QString &, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_4(s1 QStringRef_ITF, s2 string, arg2 int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_4(s1 QStringRef_ITF, s2 string, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_4(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QString &, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_4_(s1 QStringRef_ITF, s2 string) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_5(s1 QStringRef_ITF, s2 QStringRef_ITF, arg2 int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_5(s1 QStringRef_ITF, s2 QStringRef_ITF, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_5(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_5_(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, QLatin1String, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_6(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/, cs int) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QLatin1String_PTR() != nil {
		convArg1 = s2.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_6(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/, cs int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_6(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, QLatin1String, Qt::CaseSensitivity)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

If cs is Qt::CaseSensitive, the comparison is case sensitive; otherwise the comparison is case insensitive.

Case sensitive comparison is based exclusively on the numeric Unicode values of the characters and is very fast, but is not what a human would expect. Consider sorting user-visible strings with localeAwareCompare().


  int x = QString::compare("aUtO", "AuTo", Qt::CaseInsensitive);  // x == 0
  int y = QString::compare("auto", "Car", Qt::CaseSensitive);     // y > 0
  int z = QString::compare("auto", "Car", Qt::CaseInsensitive);   // z < 0



This function was introduced in  Qt 4.2.

See also operator==(), operator<(), and operator>().
*/
func (this *QStringRef) Compare_6_(s1 QStringRef_ITF, s2 QLatin1String_ITF /*123*/) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QLatin1String_PTR() != nil {
		convArg1 = s2.QLatin1String_PTR().GetCthis()
	}
	// arg: 2, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1573
// index:0
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

The comparison is performed in a locale- and also platform-dependent manner. Use this function to present sorted lists of strings to the user.

On macOS and iOS this function compares according the "Order for sorted lists" setting in the International preferences panel.

See also compare() and QLocale.
*/
func (this *QStringRef) LocaleAwareCompare(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1574
// index:1
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &) const

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

The comparison is performed in a locale- and also platform-dependent manner. Use this function to present sorted lists of strings to the user.

On macOS and iOS this function compares according the "Order for sorted lists" setting in the International preferences panel.

See also compare() and QLocale.
*/
func (this *QStringRef) LocaleAwareCompare_1(s QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringRef_PTR() != nil {
		convArg0 = s.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1575
// index:2
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QString &)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

The comparison is performed in a locale- and also platform-dependent manner. Use this function to present sorted lists of strings to the user.

On macOS and iOS this function compares according the "Order for sorted lists" setting in the International preferences panel.

See also compare() and QLocale.
*/
func (this *QStringRef) LocaleAwareCompare_2(s1 QStringRef_ITF, s2 string) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_2(s1 QStringRef_ITF, s2 string) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_2(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1576
// index:3
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QStringRef &)

/*
Compares s1 with s2 and returns an integer less than, equal to, or greater than zero if s1 is less than, equal to, or greater than s2.

The comparison is performed in a locale- and also platform-dependent manner. Use this function to present sorted lists of strings to the user.

On macOS and iOS this function compares according the "Order for sorted lists" setting in the International preferences panel.

See also compare() and QLocale.
*/
func (this *QStringRef) LocaleAwareCompare_3(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var convArg0 unsafe.Pointer
	if s1 != nil && s1.QStringRef_PTR() != nil {
		convArg0 = s1.QStringRef_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s2 != nil && s2.QStringRef_PTR() != nil {
		convArg1 = s2.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_3(s1 QStringRef_ITF, s2 QStringRef_ITF) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_3(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1578
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef trimmed() const

/*
Returns a string that has whitespace removed from the start and the end.

Whitespace means any character for which QChar::isSpace() returns true. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.

Example:


  QString str = "  lots\t of\nwhitespace\r\n ";
  str = str.trimmed();
  // str == "lots\t of\nwhitespace"



Unlike simplified(), trimmed() leaves internal whitespace alone.

See also simplified().
*/
func (this *QStringRef) Trimmed() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the string converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toShort()

Example:


  QString str = "FF";
  bool ok;

  short hex = str.toShort(&ok, 16);   // hex == 255, ok == true
  short dec = str.toShort(&ok, 10);   // dec == 0, ok == false



See also number(), toUShort(), toInt(), and QLocale::toShort().
*/
func (this *QStringRef) ToShort(ok *bool, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the string converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toShort()

Example:


  QString str = "FF";
  bool ok;

  short hex = str.toShort(&ok, 16);   // hex == 255, ok == true
  short dec = str.toShort(&ok, 10);   // dec == 0, ok == false



See also number(), toUShort(), toInt(), and QLocale::toShort().
*/
func (this *QStringRef) ToShort__() int16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(bool *, int) const

/*
Returns the string converted to a short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toShort()

Example:


  QString str = "FF";
  bool ok;

  short hex = str.toShort(&ok, 16);   // hex == 255, ok == true
  short dec = str.toShort(&ok, 10);   // dec == 0, ok == false



See also number(), toUShort(), toInt(), and QLocale::toShort().
*/
func (this *QStringRef) ToShort__1(ok *bool) int16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the string converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUShort()

Example:


  QString str = "FF";
  bool ok;

  ushort hex = str.toUShort(&ok, 16);     // hex == 255, ok == true
  ushort dec = str.toUShort(&ok, 10);     // dec == 0, ok == false



See also number(), toShort(), and QLocale::toUShort().
*/
func (this *QStringRef) ToUShort(ok *bool, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the string converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUShort()

Example:


  QString str = "FF";
  bool ok;

  ushort hex = str.toUShort(&ok, 16);     // hex == 255, ok == true
  ushort dec = str.toUShort(&ok, 10);     // dec == 0, ok == false



See also number(), toShort(), and QLocale::toUShort().
*/
func (this *QStringRef) ToUShort__() uint16 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(bool *, int) const

/*
Returns the string converted to an unsigned short using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUShort()

Example:


  QString str = "FF";
  bool ok;

  ushort hex = str.toUShort(&ok, 16);     // hex == 255, ok == true
  ushort dec = str.toUShort(&ok, 10);     // dec == 0, ok == false



See also number(), toShort(), and QLocale::toUShort().
*/
func (this *QStringRef) ToUShort__1(ok *bool) uint16 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the string converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toInt()

Example:


  QString str = "FF";
  bool ok;
  int hex = str.toInt(&ok, 16);       // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);       // dec == 0, ok == false



See also number(), toUInt(), toDouble(), and QLocale::toInt().
*/
func (this *QStringRef) ToInt(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the string converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toInt()

Example:


  QString str = "FF";
  bool ok;
  int hex = str.toInt(&ok, 16);       // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);       // dec == 0, ok == false



See also number(), toUInt(), toDouble(), and QLocale::toInt().
*/
func (this *QStringRef) ToInt__() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(bool *, int) const

/*
Returns the string converted to an int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toInt()

Example:


  QString str = "FF";
  bool ok;
  int hex = str.toInt(&ok, 16);       // hex == 255, ok == true
  int dec = str.toInt(&ok, 10);       // dec == 0, ok == false



See also number(), toUInt(), toDouble(), and QLocale::toInt().
*/
func (this *QStringRef) ToInt__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the string converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUInt()

Example:


  QString str = "FF";
  bool ok;

  uint hex = str.toUInt(&ok, 16);     // hex == 255, ok == true
  uint dec = str.toUInt(&ok, 10);     // dec == 0, ok == false



See also number(), toInt(), and QLocale::toUInt().
*/
func (this *QStringRef) ToUInt(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the string converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUInt()

Example:


  QString str = "FF";
  bool ok;

  uint hex = str.toUInt(&ok, 16);     // hex == 255, ok == true
  uint dec = str.toUInt(&ok, 10);     // dec == 0, ok == false



See also number(), toInt(), and QLocale::toUInt().
*/
func (this *QStringRef) ToUInt__() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(bool *, int) const

/*
Returns the string converted to an unsigned int using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toUInt()

Example:


  QString str = "FF";
  bool ok;

  uint hex = str.toUInt(&ok, 16);     // hex == 255, ok == true
  uint dec = str.toUInt(&ok, 10);     // dec == 0, ok == false



See also number(), toInt(), and QLocale::toUInt().
*/
func (this *QStringRef) ToUInt__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the string converted to a long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  long hex = str.toLong(&ok, 16);     // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);     // dec == 0, ok == false



See also number(), toULong(), toInt(), and QLocale::toInt().
*/
func (this *QStringRef) ToLong(ok *bool, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the string converted to a long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  long hex = str.toLong(&ok, 16);     // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);     // dec == 0, ok == false



See also number(), toULong(), toInt(), and QLocale::toInt().
*/
func (this *QStringRef) ToLong__() int {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(bool *, int) const

/*
Returns the string converted to a long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  long hex = str.toLong(&ok, 16);     // hex == 255, ok == true
  long dec = str.toLong(&ok, 10);     // dec == 0, ok == false



See also number(), toULong(), toInt(), and QLocale::toInt().
*/
func (this *QStringRef) ToLong__1(ok *bool) int {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the string converted to an unsigned long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  ulong hex = str.toULong(&ok, 16);   // hex == 255, ok == true
  ulong dec = str.toULong(&ok, 10);   // dec == 0, ok == false



See also number() and QLocale::toUInt().
*/
func (this *QStringRef) ToULong(ok *bool, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the string converted to an unsigned long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  ulong hex = str.toULong(&ok, 16);   // hex == 255, ok == true
  ulong dec = str.toULong(&ok, 10);   // dec == 0, ok == false



See also number() and QLocale::toUInt().
*/
func (this *QStringRef) ToULong__() uint {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(bool *, int) const

/*
Returns the string converted to an unsigned long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  ulong hex = str.toULong(&ok, 16);   // hex == 255, ok == true
  ulong dec = str.toULong(&ok, 10);   // dec == 0, ok == false



See also number() and QLocale::toUInt().
*/
func (this *QStringRef) ToULong__1(ok *bool) uint {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the string converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  qint64 hex = str.toLongLong(&ok, 16);      // hex == 255, ok == true
  qint64 dec = str.toLongLong(&ok, 10);      // dec == 0, ok == false



See also number(), toULongLong(), toInt(), and QLocale::toLongLong().
*/
func (this *QStringRef) ToLongLong(ok *bool, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the string converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  qint64 hex = str.toLongLong(&ok, 16);      // hex == 255, ok == true
  qint64 dec = str.toLongLong(&ok, 10);      // dec == 0, ok == false



See also number(), toULongLong(), toInt(), and QLocale::toLongLong().
*/
func (this *QStringRef) ToLongLong__() int64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(bool *, int) const

/*
Returns the string converted to a long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toLongLong()

Example:


  QString str = "FF";
  bool ok;

  qint64 hex = str.toLongLong(&ok, 16);      // hex == 255, ok == true
  qint64 dec = str.toLongLong(&ok, 10);      // dec == 0, ok == false



See also number(), toULongLong(), toInt(), and QLocale::toLongLong().
*/
func (this *QStringRef) ToLongLong__1(ok *bool) int64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the string converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  quint64 hex = str.toULongLong(&ok, 16);    // hex == 255, ok == true
  quint64 dec = str.toULongLong(&ok, 10);    // dec == 0, ok == false



See also number(), toLongLong(), and QLocale::toULongLong().
*/
func (this *QStringRef) ToULongLong(ok *bool, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the string converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  quint64 hex = str.toULongLong(&ok, 16);    // hex == 255, ok == true
  quint64 dec = str.toULongLong(&ok, 10);    // dec == 0, ok == false



See also number(), toLongLong(), and QLocale::toULongLong().
*/
func (this *QStringRef) ToULongLong__() uint64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(bool *, int) const

/*
Returns the string converted to an unsigned long long using base base, which is 10 by default and must be between 2 and 36, or 0. Returns 0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.

If base is 0, the C language convention is used: If the string begins with "0x", base 16 is used; if the string begins with "0", base 8 is used; otherwise, base 10 is used.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toULongLong()

Example:


  QString str = "FF";
  bool ok;

  quint64 hex = str.toULongLong(&ok, 16);    // hex == 255, ok == true
  quint64 dec = str.toULongLong(&ok, 10);    // dec == 0, ok == false



See also number(), toLongLong(), and QLocale::toULongLong().
*/
func (this *QStringRef) ToULongLong__1(ok *bool) uint64 {
	// arg: 1, int=Int, =Invalid, , Invalid
	base := int(10)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok, base)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the string converted to a float value.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true. Returns 0.0 if the conversion fails.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toFloat()

Example:


  QString str1 = "1234.56";
  str1.toFloat();             // returns 1234.56

  bool ok;
  QString str2 = "R2D2";
  str2.toFloat(&ok);          // returns 0.0, sets ok to false



See also number(), toDouble(), toInt(), and QLocale::toFloat().
*/
func (this *QStringRef) ToFloat(ok *bool) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(bool *) const

/*
Returns the string converted to a float value.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true. Returns 0.0 if the conversion fails.

The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toFloat()

Example:


  QString str1 = "1234.56";
  str1.toFloat();             // returns 1234.56

  bool ok;
  QString str2 = "R2D2";
  str2.toFloat(&ok);          // returns 0.0, sets ok to false



See also number(), toDouble(), toInt(), and QLocale::toFloat().
*/
func (this *QStringRef) ToFloat__() float32 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the string converted to a double value.

Returns 0.0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QString str = "1234.56";
  double val = str.toDouble();   // val == 1234.56



Warning: The QString content may only contain valid numerical characters which includes the plus/minus sign, the characters g and e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.


  bool ok;
  double d;

  d = QString( "1234.56e-02" ).toDouble(&ok); // ok == true, d == 12.3456



The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toDouble()


  d = QString( "1234,56" ).toDouble(&ok); // ok == false
  d = QString( "1234.56" ).toDouble(&ok); // ok == true, d == 1234.56



For historical reasons, this function does not handle thousands group separators. If you need to convert such numbers, use QLocale::toDouble().


  d = QString( "1,234,567.89" ).toDouble(&ok); // ok == false
  d = QString( "1234567.89" ).toDouble(&ok); // ok == true



See also number(), QLocale::setDefault(), QLocale::toDouble(), and trimmed().
*/
func (this *QStringRef) ToDouble(ok *bool) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(bool *) const

/*
Returns the string converted to a double value.

Returns 0.0 if the conversion fails.

If a conversion error occurs, *ok is set to false; otherwise *ok is set to true.


  QString str = "1234.56";
  double val = str.toDouble();   // val == 1234.56



Warning: The QString content may only contain valid numerical characters which includes the plus/minus sign, the characters g and e used in scientific notation, and the decimal point. Including the unit or additional characters leads to a conversion error.


  bool ok;
  double d;

  d = QString( "1234.56e-02" ).toDouble(&ok); // ok == true, d == 12.3456



The string conversion will always happen in the 'C' locale. For locale dependent conversion use QLocale::toDouble()


  d = QString( "1234,56" ).toDouble(&ok); // ok == false
  d = QString( "1234.56" ).toDouble(&ok); // ok == true, d == 1234.56



For historical reasons, this function does not handle thousands group separators. If you need to convert such numbers, use QLocale::toDouble().


  d = QString( "1,234,567.89" ).toDouble(&ok); // ok == false
  d = QString( "1234567.89" ).toDouble(&ok); // ok == true



See also number(), QLocale::setDefault(), QLocale::toDouble(), and trimmed().
*/
func (this *QStringRef) ToDouble__() float64 {
	// arg: 0, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
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
