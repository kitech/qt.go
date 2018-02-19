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
// extern C begin: 38
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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

// /usr/include/qt/QtCore/qstring.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String()
func NewQLatin1String() *QLatin1String {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:95
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const char *)
func NewQLatin1String_1(s string) *QLatin1String {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:96
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const char *, const char *)
func NewQLatin1String_2(f string, l string) *QLatin1String {
	var convArg0 = qtrt.CString(f)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(l)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKcS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:98
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const char *, int)
func NewQLatin1String_3(s string, sz int) *QLatin1String {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKci", qtrt.FFI_TYPE_POINTER, convArg0, sz)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLatin1String)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:99
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QLatin1String(const QByteArray &)
func NewQLatin1String_4(s QByteArray_ITF) *QLatin1String {
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

// /usr/include/qt/QtCore/qstring.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * latin1() const
func (this *QLatin1String) Latin1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String6latin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qstring.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const
func (this *QLatin1String) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * data() const
func (this *QLatin1String) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qstring.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QLatin1String) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QLatin1String) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char at(int) const
func (this *QLatin1String) At(i int) *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char operator[](int) const
func (this *QLatin1String) Operator_get_index(i int) *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringixEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char front() const
func (this *QLatin1String) Front() *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [1] QLatin1Char back() const
func (this *QLatin1String) Back() *QLatin1Char /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1CharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1Char)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const
func (this *QLatin1String) StartsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity) const
func (this *QLatin1String) StartsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:117
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QLatin1String) StartsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:117
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QLatin1String) StartsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:119
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar) const
func (this *QLatin1String) StartsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:121
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity) const
func (this *QLatin1String) StartsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const
func (this *QLatin1String) EndsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity) const
func (this *QLatin1String) EndsWith__(s QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:126
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QLatin1String) EndsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:126
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity) const
func (this *QLatin1String) EndsWith_1_(s QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:128
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar) const
func (this *QLatin1String) EndsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:130
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity) const
func (this *QLatin1String) EndsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QLatin1String::const_iterator begin() const
func (this *QLatin1String) Begin() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return string(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QLatin1String::const_iterator cbegin() const
func (this *QLatin1String) Cbegin() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return string(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QLatin1String::const_iterator end() const
func (this *QLatin1String) End() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return string(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QLatin1String::const_iterator cend() const
func (this *QLatin1String) Cend() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return string(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String mid(int) const
func (this *QLatin1String) Mid(pos int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String3midEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:156
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String mid(int, int) const
func (this *QLatin1String) Mid_1(pos int, n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String left(int) const
func (this *QLatin1String) Left(n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:160
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String right(int) const
func (this *QLatin1String) Right(n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String chopped(int) const
func (this *QLatin1String) Chopped(n int) *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:165
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(int)
func (this *QLatin1String) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1String4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:167
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(int)
func (this *QLatin1String) Truncate(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QLatin1String8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:170
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QLatin1String trimmed() const
func (this *QLatin1String) Trimmed() *QLatin1String /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1String7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLatin1StringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLatin1String)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:172
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QString &) const
func (this *QLatin1String) Operator_equal_equal(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:180
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const char *) const
func (this *QLatin1String) Operator_equal_equal_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:187
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QByteArray &) const
func (this *QLatin1String) Operator_equal_equal_2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringeqERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QString &) const
func (this *QLatin1String) Operator_not_equal(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:181
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const char *) const
func (this *QLatin1String) Operator_not_equal_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:188
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QByteArray &) const
func (this *QLatin1String) Operator_not_equal_2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringneERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:174
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QString &) const
func (this *QLatin1String) Operator_greater_than(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:183
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const char *) const
func (this *QLatin1String) Operator_greater_than_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:190
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(const QByteArray &) const
func (this *QLatin1String) Operator_greater_than_2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgtERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QString &) const
func (this *QLatin1String) Operator_less_than(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:182
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const char *) const
func (this *QLatin1String) Operator_less_than_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:189
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QByteArray &) const
func (this *QLatin1String) Operator_less_than_2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringltERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:176
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QString &) const
func (this *QLatin1String) Operator_greater_than_equal(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:185
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const char *) const
func (this *QLatin1String) Operator_greater_than_equal_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:192
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(const QByteArray &) const
func (this *QLatin1String) Operator_greater_than_equal_2(s QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QByteArray_PTR() != nil {
		convArg0 = s.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringgeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:177
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QString &) const
func (this *QLatin1String) Operator_less_than_equal(s string) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:184
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const char *) const
func (this *QLatin1String) Operator_less_than_equal_1(s string) bool {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QLatin1StringleEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:191
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(const QByteArray &) const
func (this *QLatin1String) Operator_less_than_equal_2(s QByteArray_ITF) bool {
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
