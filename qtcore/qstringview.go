package qtcore

// /usr/include/qt/QtCore/qstringview.h
// #include <qstringview.h>
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
type QStringView struct {
	*qtrt.CObject
}
type QStringView_ITF interface {
	QStringView_PTR() *QStringView
}

func (ptr *QStringView) QStringView_PTR() *QStringView { return ptr }

func (this *QStringView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringView) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringViewFromPointer(cthis unsafe.Pointer) *QStringView {
	return &QStringView{&qtrt.CObject{cthis}}
}
func (*QStringView) NewFromPointer(cthis unsafe.Pointer) *QStringView {
	return NewQStringViewFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringview.h:172
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringView()

/*
Constructs a null string view.

See also isNull().
*/
func (*QStringView) NewForInherit() *QStringView {
	return NewQStringView()
}
func NewQStringView() *QStringView {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringViewC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringView)
	return gothis
}

// /usr/include/qt/QtCore/qstringview.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns a deep copy of this string view's data as a QString.

The return value will be the null QString if and only if this string view is null.

Warning: QStringView can store strings with more than 230 characters while QString cannot. Calling this function on a string view for which size() returns a value greater than INT_MAX / 2 constitutes undefined behavior.
*/
func (this *QStringView) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstringview.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qsizetype size() const

/*
Returns the size of this string view, in UTF-16 code points (that is, surrogate pairs count as two for the purposes of this function, the same as in QString and QStringRef).

See also empty(), isEmpty(), isNull(), and length().
*/
func (this *QStringView) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstringview.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_pointer data() const

/*
Returns a const pointer to the first character in the string.

Note: The character array represented by the return value is not null-terminated.

See also begin(), end(), and utf16().
*/
func (this *QStringView) Data() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QStringView::storage_type * utf16() const

/*
Returns a const pointer to the first character in the string.

storage_type is char16_t, except on MSVC 2013 (which lacks char16_t support), where it is wchar_t instead.

Note: The character array represented by the return value is not null-terminated.

See also begin(), end(), and data().
*/
func (this *QStringView) Utf16() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5utf16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qstringview.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar operator[](qsizetype) const

/*

 */
func (this *QStringView) Operator_get_index(n int64) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringViewixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLatin1() const

/*
Returns a Latin-1 representation of the string as a QByteArray.

The behavior is undefined if the string contains non-Latin1 characters.

See also toUtf8(), toLocal8Bit(), QTextCodec, and qConvertToLatin1().
*/
func (this *QStringView) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUtf8() const

/*
Returns a UTF-8 representation of the string as a QByteArray.

UTF-8 is a Unicode codec and can represent all characters in a Unicode string like QString.

See also toLatin1(), toLocal8Bit(), QTextCodec, and qConvertToUtf8().
*/
func (this *QStringView) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit() const

/*
Returns a local 8-bit representation of the string as a QByteArray.

QTextCodec::codecForLocale() is used to perform the conversion from Unicode. If the locale's encoding could not be determined, this function does the same as toLatin1().

The behavior is undefined if the string contains characters not supported by the locale's 8-bit encoding.

See also toLatin1(), toUtf8(), and QTextCodec.
*/
func (this *QStringView) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar at(qsizetype) const

/*
Returns the character at position n in this string view.

The behavior is undefined if n is negative or not less than size().

See also operator[](), front(), and back().
*/
func (this *QStringView) At(n int64) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView2atEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView mid(qsizetype) const

/*
Returns the substring starting at position start in this object, and extending to the end of the string.

Note: The behavior is undefined when start < 0 or start > size().

See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QStringView) Mid(pos int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3midEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:237
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QStringView mid(qsizetype, qsizetype) const

/*
Returns the substring starting at position start in this object, and extending to the end of the string.

Note: The behavior is undefined when start < 0 or start > size().

See also left(), right(), chopped(), chop(), and truncate().
*/
func (this *QStringView) Mid_1(pos int64, n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3midExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView left(qsizetype) const

/*
Returns the substring of length length starting at position 0 in this object.

Note: The behavior is undefined when length < 0 or length > size().

See also mid(), right(), chopped(), chop(), and truncate().
*/
func (this *QStringView) Left(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4leftEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:241
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView right(qsizetype) const

/*
Returns the substring of length length starting at position size() - length in this object.

Note: The behavior is undefined when length < 0 or length > size().

See also mid(), left(), chopped(), chop(), and truncate().
*/
func (this *QStringView) Right(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5rightEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:243
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView chopped(qsizetype) const

/*
Returns the substring of length size() - length starting at the beginning of this object.

Same as left(size() - length).

Note: The behavior is undefined when length < 0 or length > size().

See also mid(), left(), right(), chop(), and truncate().
*/
func (this *QStringView) Chopped(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7choppedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(qsizetype)

/*
Truncates this string view to length length.

Same as *this = left(length).

Note: The behavior is undefined when length < 0 or length > size().

See also mid(), left(), right(), chopped(), and chop().
*/
func (this *QStringView) Truncate(n int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringView8truncateEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:248
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(qsizetype)

/*
Truncates this string view by length characters.

Same as *this = left(size() - length).

Note: The behavior is undefined when length < 0 or length > size().

See also mid(), left(), right(), chopped(), and truncate().
*/
func (this *QStringView) Chop(n int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringView4chopEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:251
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView trimmed() const

/*
Strips leading and trailing whitespace and returns the result.

Whitespace means any character for which QChar::isSpace() returns true. This includes the ASCII characters '\t', '\n', '\v', '\f', '\r', and ' '.
*/
func (this *QStringView) Trimmed() *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:253
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:253
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:255
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:255
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:256
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:258
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if this string-view starts with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also endsWith().
*/
func (this *QStringView) StartsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:261
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:261
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:263
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:263
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum, , Invalid
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:264
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:266
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const

/*
Returns true if this string-view ends with string-view str, Latin-1 string l1, or character ch, respectively; otherwise returns false.

If cs is Qt::CaseSensitive (the default), the search is case-sensitive; otherwise the search is case-insensitive.

See also startsWith().
*/
func (this *QStringView) EndsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:272
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator begin() const

/*
Returns a const STL-style iterator pointing to the first character in the string.

This function is provided for STL compatibility.

See also end(), cbegin(), rbegin(), and data().
*/
func (this *QStringView) Begin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator end() const

/*
Returns a const STL-style iterator pointing to the imaginary character after the last character in the list.

This function is provided for STL compatibility.

See also begin(), cend(), and rend().
*/
func (this *QStringView) End() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:274
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator cbegin() const

/*
Same as begin().

This function is provided for STL compatibility.

See also cend(), begin(), crbegin(), and data().
*/
func (this *QStringView) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:275
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator cend() const

/*
Same as end().

This function is provided for STL compatibility.

See also cbegin(), end(), and crend().
*/
func (this *QStringView) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty() const

/*
Returns whether this string view is empty - that is, whether size() == 0.

This function is provided for STL compatibility.

See also isEmpty(), isNull(), size(), and length().
*/
func (this *QStringView) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:282
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front() const

/*
Returns the first character in the string. Same as first().

This function is provided for STL compatibility.

Warning: Calling this function on an empty string view constitutes undefined behavior.

See also back(), first(), and last().
*/
func (this *QStringView) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:283
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back() const

/*
Returns the last character in the string. Same as last().

This function is provided for STL compatibility.

Warning: Calling this function on an empty string view constitutes undefined behavior.

See also front(), first(), and last().
*/
func (this *QStringView) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns whether this string view is null - that is, whether data() == nullptr.

This functions is provided for compatibility with other Qt containers.

See also empty(), isEmpty(), size(), and length().
*/
func (this *QStringView) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:289
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns whether this string view is empty - that is, whether size() == 0.

This function is provided for compatibility with other Qt containers.

See also empty(), isNull(), size(), and length().
*/
func (this *QStringView) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:290
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length() const

/*
Same as size(), except returns the result as an int.

This function is provided for compatibility with other Qt containers.

Warning: QStringView can represent strings with more than 231 characters. Calling this function on a string view for which size() returns a value greater than INT_MAX constitutes undefined behavior.

See also empty(), isEmpty(), isNull(), and size().
*/
func (this *QStringView) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringview.h:292
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar first() const

/*
Returns the first character in the string. Same as front().

This function is provided for compatibility with other Qt containers.

Warning: Calling this function on an empty string view constitutes undefined behavior.

See also front(), back(), and last().
*/
func (this *QStringView) First() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:293
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar last() const

/*
Returns the last character in the string. Same as back().

This function is provided for compatibility with other Qt containers.

Warning: Calling this function on an empty string view constitutes undefined behavior.

See also back(), front(), and first().
*/
func (this *QStringView) Last() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

func DeleteQStringView(this *QStringView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
