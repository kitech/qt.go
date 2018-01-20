//  header block begin
// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
type QLatin1String struct {
	*qtrt.CObject
}

func (this *QLatin1String) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qstring.h:94
// index:0
// inline
// void QLatin1String()
func NewQLatin1String() *QLatin1String {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1StringC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(cthis)
	return gothis
}
func NewQLatin1StringFromPointer(cthis unsafe.Pointer) *QLatin1String {
	return &QLatin1String{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstring.h:95
// index:1
// inline
// void QLatin1String(const char *)
func NewQLatin1String_1(s unsafe.Pointer) *QLatin1String {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:96
// index:2
// inline
// void QLatin1String(const char *, const char *)
func NewQLatin1String_2(f unsafe.Pointer, l unsafe.Pointer) *QLatin1String {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKcS1_", ffiqt.FFI_TYPE_VOID, cthis, f, l)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:98
// index:3
// inline
// void QLatin1String(const char *, int)
func NewQLatin1String_3(s unsafe.Pointer, sz int) *QLatin1String {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1StringC2EPKci", ffiqt.FFI_TYPE_VOID, cthis, s, &sz)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:99
// index:4
// inline
// void QLatin1String(const class QByteArray &)
func NewQLatin1String_4(s unsafe.Pointer) *QLatin1String {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1StringC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, s)
	gopp.ErrPrint(err, rv)
	gothis := NewQLatin1StringFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:101
// index:0
// inline
// const char * latin1()
func (this *QLatin1String) Latin1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String6latin1Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:102
// index:0
// inline
// int size()
func (this *QLatin1String) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:103
// index:0
// inline
// const char * data()
func (this *QLatin1String) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:105
// index:0
// inline
// bool isNull()
func (this *QLatin1String) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:106
// index:0
// inline
// bool isEmpty()
func (this *QLatin1String) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:108
// index:0
// inline
// QLatin1Char at(int)
func (this *QLatin1String) At(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String2atEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:112
// index:0
// inline
// QLatin1Char front()
func (this *QLatin1String) Front() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String5frontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:113
// index:0
// inline
// QLatin1Char back()
func (this *QLatin1String) Back() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String4backEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:115
// index:0
// inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QLatin1String) StartsWith(s unsafe.Pointer, cs int) {
	// 0: (, s QStringView, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:117
// index:1
// inline
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QLatin1String) StartsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, s QLatin1String, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:119
// index:2
// inline
// bool startsWith(class QChar)
func (this *QLatin1String) StartsWith_2(c unsafe.Pointer) {
	// 2: (, c QChar), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:121
// index:3
// inline
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QLatin1String) StartsWith_3(c unsafe.Pointer, cs int) {
	// 3: (, c QChar, cs Qt::CaseSensitivity), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:124
// index:0
// inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QLatin1String) EndsWith(s unsafe.Pointer, cs int) {
	// 0: (, s QStringView, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:126
// index:1
// inline
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QLatin1String) EndsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, s QLatin1String, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:128
// index:2
// inline
// bool endsWith(class QChar)
func (this *QLatin1String) EndsWith_2(c unsafe.Pointer) {
	// 2: (, c QChar), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:130
// index:3
// inline
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QLatin1String) EndsWith_3(c unsafe.Pointer, cs int) {
	// 3: (, c QChar, cs Qt::CaseSensitivity), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:141
// index:0
// inline
// QLatin1String::const_iterator begin()
func (this *QLatin1String) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:142
// index:0
// inline
// QLatin1String::const_iterator cbegin()
func (this *QLatin1String) Cbegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String6cbeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:143
// index:0
// inline
// QLatin1String::const_iterator end()
func (this *QLatin1String) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:144
// index:0
// inline
// QLatin1String::const_iterator cend()
func (this *QLatin1String) Cend() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String4cendEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:154
// index:0
// inline
// QLatin1String mid(int)
func (this *QLatin1String) Mid(pos int) {
	// 0: (, pos int), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String3midEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:156
// index:1
// inline
// QLatin1String mid(int, int)
func (this *QLatin1String) Mid_1(pos int, n int) {
	// 1: (, pos int, n int), (&pos, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String3midEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:158
// index:0
// inline
// QLatin1String left(int)
func (this *QLatin1String) Left(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String4leftEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:160
// index:0
// inline
// QLatin1String right(int)
func (this *QLatin1String) Right(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String5rightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:162
// index:0
// inline
// QLatin1String chopped(int)
func (this *QLatin1String) Chopped(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String7choppedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:165
// index:0
// inline
// void chop(int)
func (this *QLatin1String) Chop(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1String4chopEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:167
// index:0
// inline
// void truncate(int)
func (this *QLatin1String) Truncate(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QLatin1String8truncateEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:170
// index:0
// inline
// QLatin1String trimmed()
func (this *QLatin1String) Trimmed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QLatin1String7trimmedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
