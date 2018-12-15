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
// extern C begin: 8
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
type QLatin1String struct {
	*qtrt.CObject
}
type QLatin1String_ITF interface {
	QLatin1String_PTR() *QLatin1String
}

func (ptr *QLatin1String) QLatin1String_PTR() *QLatin1String { return ptr }

func (this *QLatin1String) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLatin1String) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLatin1StringFromPointer(cthis unsafe.Pointer) *QLatin1String {
	return &QLatin1String{&qtrt.CObject{cthis}}
}
func (*QLatin1String) NewFromPointer(cthis unsafe.Pointer) *QLatin1String {
	return NewQLatin1StringFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String()

/*

 */
func (*QLatin1String) NewForInherit() *QLatin1String {
	return NewQLatin1String()
}
func NewQLatin1String() *QLatin1String {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:92
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const char *)

/*

 */
func (*QLatin1String) NewForInherit1(s string) *QLatin1String {
	return NewQLatin1String1(s)
}
func NewQLatin1String1(s string) *QLatin1String {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:93
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const char *, int)

/*

 */
func (*QLatin1String) NewForInherit2(s string, sz int) *QLatin1String {
	return NewQLatin1String2(s, sz)
}
func NewQLatin1String2(s string, sz int) *QLatin1String {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, sz)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:94
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const QByteArray &)

/*

 */
func (*QLatin1String) NewForInherit3(s QByteArray_ITF) *QLatin1String {
	return NewQLatin1String3(s)
}
func NewQLatin1String3(s QByteArray_ITF) *QLatin1String {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * latin1() const

/*

 */
func (this *QLatin1String) Latin1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String6latin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qstring.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of characters in this string.

The last character in the string is at position size() - 1. In addition, QString ensures that the character at position size() is always '\0', so that you can use the return value of data() and constData() as arguments to functions that expect '\0'-terminated strings.

Example:


  QString str = "World";
  int n = str.size();         // n == 5
  str.data()[0];              // returns 'W'
  str.data()[4];              // returns 'd'
  str.data()[5];              // returns '\0'



See also isEmpty() and resize().
*/
func (this *QLatin1String) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * data() const

/*
Returns a pointer to the data stored in the QString. The pointer can be used to access and modify the characters that compose the string. For convenience, the data is '\0'-terminated.

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
func (this *QLatin1String) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qstring.h:100
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char at(int) const

/*
Returns the character at the given index position in the string.

The position must be a valid index position in the string (i.e., 0 <= position < size()).

See also operator[]().
*/
func (this *QLatin1String) At(i int) *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char operator[](int) const

/*

 */
func (this *QLatin1String) Operator_get_index(i int) *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String mid(int) const

/*
Returns a string that contains n characters of this string, starting at the specified position index.

Returns a null string if the position index exceeds the length of the string. If there are less than n characters available in the string starting at the given position, or if n is -1 (default), the function returns all characters that are available from the specified position.

Example:


  QString x = "Nine pineapples";
  QString y = x.mid(5, 4);            // y == "pine"
  QString z = x.mid(5);               // z == "pineapples"



See also left() and right().
*/
func (this *QLatin1String) Mid(pos int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String3midEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:105
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String mid(int, int) const

/*
Returns a string that contains n characters of this string, starting at the specified position index.

Returns a null string if the position index exceeds the length of the string. If there are less than n characters available in the string starting at the given position, or if n is -1 (default), the function returns all characters that are available from the specified position.

Example:


  QString x = "Nine pineapples";
  QString y = x.mid(5, 4);            // y == "pine"
  QString z = x.mid(5);               // z == "pineapples"



See also left() and right().
*/
func (this *QLatin1String) Mid1(pos int, n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String left(int) const

/*
Returns a substring that contains the n leftmost characters of the string.

The entire string is returned if n is greater than or equal to size(), or less than zero.


  QString x = "Pineapple";
  QString y = x.left(4);      // y == "Pine"



See also right(), mid(), and startsWith().
*/
func (this *QLatin1String) Left(n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String right(int) const

/*
Returns a substring that contains the n rightmost characters of the string.

The entire string is returned if n is greater than or equal to size(), or less than zero.


  QString x = "Pineapple";
  QString y = x.right(5);      // y == "apple"



See also left(), mid(), and endsWith().
*/
func (this *QLatin1String) Right(n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QString &) const

/*

 */
func (this *QLatin1String) Operator_equal_equal(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:120
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const char *) const

/*

 */
func (this *QLatin1String) Operator_equal_equal1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:127
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_equal_equal2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QString &) const

/*

 */
func (this *QLatin1String) Operator_not_equal(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:121
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const char *) const

/*

 */
func (this *QLatin1String) Operator_not_equal1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:128
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_not_equal2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QString &) const

/*

 */
func (this *QLatin1String) Operator_greater_than(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const char *) const

/*

 */
func (this *QLatin1String) Operator_greater_than1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:130
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_greater_than2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QString &) const

/*

 */
func (this *QLatin1String) Operator_less_than(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:122
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const char *) const

/*

 */
func (this *QLatin1String) Operator_less_than1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:129
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_less_than2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QString &) const

/*

 */
func (this *QLatin1String) Operator_greater_than_equal(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:125
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const char *) const

/*

 */
func (this *QLatin1String) Operator_greater_than_equal1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:132
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_greater_than_equal2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QString &) const

/*

 */
func (this *QLatin1String) Operator_less_than_equal(s string) bool {
	var tmpArg0 = NewQString5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:124
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const char *) const

/*

 */
func (this *QLatin1String) Operator_less_than_equal1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringleEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:131
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QByteArray &) const

/*

 */
func (this *QLatin1String) Operator_less_than_equal2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringleERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQLatin1String(this *QLatin1String) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
