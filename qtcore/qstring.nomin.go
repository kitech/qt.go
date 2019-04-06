//  header block begin

// +build !minimal

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
// extern C begin: 52
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qstring.h:363
// index:5
// Public Visibility=Default Availability=Available
// [4] int count(const QRegularExpression &) const

/*
Returns the number of (potentially overlapping) occurrences of the string str in this string.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

See also contains() and indexOf().
*/
func (this *QString) Count5(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:357
// index:6
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const

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
func (this *QString) IndexOf6(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:357
// index:6
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const

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
func (this *QString) IndexOf6p(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:358
// index:7
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int, QRegularExpressionMatch *) const

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
func (this *QString) IndexOf7(re QRegularExpression_ITF, from int, rmatch QRegularExpressionMatch_ITF /*777 QRegularExpressionMatch **/) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rmatch != nil && rmatch.QRegularExpressionMatch_PTR() != nil {
		convArg2 = rmatch.QRegularExpressionMatch_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:359
// index:6
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const

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
func (this *QString) LastIndexOf6(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:359
// index:6
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const

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
func (this *QString) LastIndexOf6p(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:360
// index:7
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int, QRegularExpressionMatch *) const

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
func (this *QString) LastIndexOf7(re QRegularExpression_ITF, from int, rmatch QRegularExpressionMatch_ITF /*777 QRegularExpressionMatch **/) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rmatch != nil && rmatch.QRegularExpressionMatch_PTR() != nil {
		convArg2 = rmatch.QRegularExpressionMatch_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:361
// index:6
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRegularExpression &) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QString) Contains6(re QRegularExpression_ITF) bool {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:362
// index:7
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRegularExpression &, QRegularExpressionMatch *) const

/*
Returns true if this string contains an occurrence of the string str; otherwise returns false.

If cs is Qt::CaseSensitive (default), the search is case sensitive; otherwise the search is case insensitive.

Example:


  QString str = "Peter Pan";
  str.contains("peter", Qt::CaseInsensitive);    // returns true



See also indexOf() and count().
*/
func (this *QString) Contains7(re QRegularExpression_ITF, match_ QRegularExpressionMatch_ITF /*777 QRegularExpressionMatch **/) bool {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if match_ != nil && match_.QRegularExpressionMatch_PTR() != nil {
		convArg1 = match_.QRegularExpressionMatch_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:381
// index:3
// Public Visibility=Default Availability=Available
// [8] QString section(const QRegularExpression &, int, int, QString::SectionFlags) const

/*
This function returns a section of the string.

This string is treated as a sequence of fields separated by the character, sep. The returned string consists of the fields from position start to position end inclusive. If end is not specified, all fields from position start to the end of the string are included. Fields are numbered 0, 1, 2, etc., counting from the left, and -1, -2, etc., counting from right to left.

The flags argument can be used to affect some aspects of the function's behavior, e.g. whether to be case sensitive, whether to skip empty fields and how to deal with leading and trailing separators; see SectionFlags.


  QString str;
  QString csv = "forename,middlename,surname,phone";
  QString path = "/usr/local/bin/myapp"; // First field is empty
  QString::SectionFlag flag = QString::SectionSkipEmpty;

  str = csv.section(',', 2, 2);   // str == "surname"
  str = path.section('/', 3, 4);  // str == "bin/myapp"
  str = path.section('/', 3, 3, flag); // str == "myapp"



If start or end is negative, we count fields from the right of the string, the right-most field being -1, the one from right-most field being -2, and so on.


  str = csv.section(',', -3, -2);  // str == "middlename,surname"
  str = path.section('/', -1); // str == "myapp"



See also split().
*/
func (this *QString) Section3(re QRegularExpression_ITF, start int, end_ int, flags int) string {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERK18QRegularExpressionii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:381
// index:3
// Public Visibility=Default Availability=Available
// [8] QString section(const QRegularExpression &, int, int, QString::SectionFlags) const

/*
This function returns a section of the string.

This string is treated as a sequence of fields separated by the character, sep. The returned string consists of the fields from position start to position end inclusive. If end is not specified, all fields from position start to the end of the string are included. Fields are numbered 0, 1, 2, etc., counting from the left, and -1, -2, etc., counting from right to left.

The flags argument can be used to affect some aspects of the function's behavior, e.g. whether to be case sensitive, whether to skip empty fields and how to deal with leading and trailing separators; see SectionFlags.


  QString str;
  QString csv = "forename,middlename,surname,phone";
  QString path = "/usr/local/bin/myapp"; // First field is empty
  QString::SectionFlag flag = QString::SectionSkipEmpty;

  str = csv.section(',', 2, 2);   // str == "surname"
  str = path.section('/', 3, 4);  // str == "bin/myapp"
  str = path.section('/', 3, 3, flag); // str == "myapp"



If start or end is negative, we count fields from the right of the string, the right-most field being -1, the one from right-most field being -2, and so on.


  str = csv.section(',', -3, -2);  // str == "middlename,surname"
  str = path.section('/', -1); // str == "myapp"



See also split().
*/
func (this *QString) Section3p(re QRegularExpression_ITF, start int) string {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	end_ := int(-1)
	// arg: 3, QString::SectionFlags=Typedef, QString::SectionFlags=Typedef, QFlags<QString::SectionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERK18QRegularExpressionii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:381
// index:3
// Public Visibility=Default Availability=Available
// [8] QString section(const QRegularExpression &, int, int, QString::SectionFlags) const

/*
This function returns a section of the string.

This string is treated as a sequence of fields separated by the character, sep. The returned string consists of the fields from position start to position end inclusive. If end is not specified, all fields from position start to the end of the string are included. Fields are numbered 0, 1, 2, etc., counting from the left, and -1, -2, etc., counting from right to left.

The flags argument can be used to affect some aspects of the function's behavior, e.g. whether to be case sensitive, whether to skip empty fields and how to deal with leading and trailing separators; see SectionFlags.


  QString str;
  QString csv = "forename,middlename,surname,phone";
  QString path = "/usr/local/bin/myapp"; // First field is empty
  QString::SectionFlag flag = QString::SectionSkipEmpty;

  str = csv.section(',', 2, 2);   // str == "surname"
  str = path.section('/', 3, 4);  // str == "bin/myapp"
  str = path.section('/', 3, 3, flag); // str == "myapp"



If start or end is negative, we count fields from the right of the string, the right-most field being -1, the one from right-most field being -2, and so on.


  str = csv.section(',', -3, -2);  // str == "middlename,surname"
  str = path.section('/', -1); // str == "myapp"



See also split().
*/
func (this *QString) Section3p1(re QRegularExpression_ITF, start int, end_ int) string {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 3, QString::SectionFlags=Typedef, QString::SectionFlags=Typedef, QFlags<QString::SectionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERK18QRegularExpressionii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:510
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QString & remove(const QRegularExpression &)

/*
Removes n characters from the string, starting at the given position index, and returns a reference to the string.

If the specified position index is within the string, but position + n is beyond the end of the string, the string is truncated at the specified position.


  QString s = "Montreal";
  s.remove(1, 4);
  // s == "Meal"



See also insert() and replace().
*/
func (this *QString) Remove5(re QRegularExpression_ITF) string {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:509
// index:12
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QRegularExpression &, const QString &)

/*
Replaces n characters beginning at index position with the string after and returns a reference to this string.

Note: If the specified position index is within the string, but position + n goes outside the strings range, then n will be adjusted to stop at the end of the string.

Example:


  QString x = "Say yes!";
  QString y = "no";
  x.replace(4, 3, y);
  // x == "Say no!"



See also insert() and remove().
*/
func (this *QString) Replace12(re QRegularExpression_ITF, after string) string {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceERK18QRegularExpressionRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:529
// index:3
// Public Visibility=Default Availability=Available
// [8] QStringList split(const QRegularExpression &, QString::SplitBehavior) const

/*
Splits the string into substrings wherever sep occurs, and returns the list of those strings. If sep does not match anywhere in the string, split() returns a single-element list containing this string.

cs specifies whether sep should be matched case sensitively or case insensitively.

If behavior is QString::SkipEmptyParts, empty entries don't appear in the result. By default, empty entries are kept.

Example:


  QString str = "a,,b,c";

  QStringList list1 = str.split(',');
  // list1: [ "a", "", "b", "c" ]

  QStringList list2 = str.split(',', QString::SkipEmptyParts);
  // list2: [ "a", "b", "c" ]



If sep is empty, split() returns an empty string, followed by each of the string's characters, followed by another empty string:


  QString str = "abc";
  auto parts = str.split("");
  // parts: {"", "a", "b", "c", ""}



To understand this behavior, recall that the empty string matches everywhere, so the above is qualitatively the same as:


  QString str = "/a/b/c/";
  auto parts = str.split('/');
  // parts: {"", "a", "b", "c", ""}



See also QStringList::join() and section().
*/
func (this *QString) Split3(sep QRegularExpression_ITF, behavior int) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if sep != nil && sep.QRegularExpression_PTR() != nil {
		convArg0 = sep.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5splitERK18QRegularExpressionNS_13SplitBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, behavior)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:529
// index:3
// Public Visibility=Default Availability=Available
// [8] QStringList split(const QRegularExpression &, QString::SplitBehavior) const

/*
Splits the string into substrings wherever sep occurs, and returns the list of those strings. If sep does not match anywhere in the string, split() returns a single-element list containing this string.

cs specifies whether sep should be matched case sensitively or case insensitively.

If behavior is QString::SkipEmptyParts, empty entries don't appear in the result. By default, empty entries are kept.

Example:


  QString str = "a,,b,c";

  QStringList list1 = str.split(',');
  // list1: [ "a", "", "b", "c" ]

  QStringList list2 = str.split(',', QString::SkipEmptyParts);
  // list2: [ "a", "b", "c" ]



If sep is empty, split() returns an empty string, followed by each of the string's characters, followed by another empty string:


  QString str = "abc";
  auto parts = str.split("");
  // parts: {"", "a", "b", "c", ""}



To understand this behavior, recall that the empty string matches everywhere, so the above is qualitatively the same as:


  QString str = "/a/b/c/";
  auto parts = str.split('/');
  // parts: {"", "a", "b", "c", ""}



See also QStringList::join() and section().
*/
func (this *QString) Split3p(sep QRegularExpression_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if sep != nil && sep.QRegularExpression_PTR() != nil {
		convArg0 = sep.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, QString::SplitBehavior=Enum, QString::SplitBehavior=Enum, , Invalid
	behavior := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5splitERK18QRegularExpressionNS_13SplitBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, behavior)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10190() {
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
