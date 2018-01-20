//  header block begin
// /usr/include/qt/QtCore/qstringview.h
// #include <qstringview.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QStringView struct {
	*qtrt.CObject
}

func (this *QStringView) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qstringview.h:171
// index:0
// inline
// void QStringView()
func NewQStringView() *QStringView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringViewC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringViewFromPointer(cthis)
	return gothis
}
func NewQStringViewFromPointer(cthis unsafe.Pointer) *QStringView {
	return &QStringView{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstringview.h:173
// index:1
// inline
// void QStringView(std::nullptr_t)
func NewQStringView_1(arg0 int) *QStringView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringViewC2EDn", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringview.h:214
// index:0
// inline
// QString toString()
func (this *QStringView) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8toStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:216
// index:0
// inline
// qsizetype size()
func (this *QStringView) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:217
// index:0
// inline
// QStringView::const_pointer data()
func (this *QStringView) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:218
// index:0
// inline
// const QStringView::storage_type * utf16()
func (this *QStringView) Utf16() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5utf16Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:227
// index:0
// inline
// QByteArray toLatin1()
func (this *QStringView) ToLatin1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8toLatin1Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:228
// index:0
// inline
// QByteArray toUtf8()
func (this *QStringView) ToUtf8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6toUtf8Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:229
// index:0
// inline
// QByteArray toLocal8Bit()
func (this *QStringView) ToLocal8Bit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView11toLocal8BitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:230
// index:0
// inline
// QVector<uint> toUcs4()
func (this *QStringView) ToUcs4() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6toUcs4Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:232
// index:0
// inline
// QChar at(qsizetype)
func (this *QStringView) At(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView2atEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:234
// index:0
// inline
// QStringView mid(qsizetype)
func (this *QStringView) Mid(pos int64) {
	// 0: (, pos qsizetype), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3midEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:236
// index:1
// inline
// QStringView mid(qsizetype, qsizetype)
func (this *QStringView) Mid_1(pos int64, n int64) {
	// 1: (, pos qsizetype, n qsizetype), (&pos, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3midExx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:238
// index:0
// inline
// QStringView left(qsizetype)
func (this *QStringView) Left(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4leftEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:240
// index:0
// inline
// QStringView right(qsizetype)
func (this *QStringView) Right(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5rightEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:242
// index:0
// inline
// QStringView chopped(qsizetype)
func (this *QStringView) Chopped(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7choppedEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:245
// index:0
// inline
// void truncate(qsizetype)
func (this *QStringView) Truncate(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringView8truncateEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:247
// index:0
// inline
// void chop(qsizetype)
func (this *QStringView) Chop(n int64) {
	// 0: (, n qsizetype), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringView4chopEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:250
// index:0
// inline
// QStringView trimmed()
func (this *QStringView) Trimmed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7trimmedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:252
// index:0
// inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringView) StartsWith(s unsafe.Pointer, cs int) {
	// 0: (, s QStringView, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:254
// index:1
// inline
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, s QLatin1String, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:255
// index:2
// inline
// bool startsWith(class QChar)
func (this *QStringView) StartsWith_2(c unsafe.Pointer) {
	// 2: (, c QChar), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:257
// index:3
// inline
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_3(c unsafe.Pointer, cs int) {
	// 3: (, c QChar, cs Qt::CaseSensitivity), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:260
// index:0
// inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringView) EndsWith(s unsafe.Pointer, cs int) {
	// 0: (, s QStringView, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:262
// index:1
// inline
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, s QLatin1String, cs Qt::CaseSensitivity), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:263
// index:2
// inline
// bool endsWith(class QChar)
func (this *QStringView) EndsWith_2(c unsafe.Pointer) {
	// 2: (, c QChar), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:265
// index:3
// inline
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_3(c unsafe.Pointer, cs int) {
	// 3: (, c QChar, cs Qt::CaseSensitivity), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:271
// index:0
// inline
// QStringView::const_iterator begin()
func (this *QStringView) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:272
// index:0
// inline
// QStringView::const_iterator end()
func (this *QStringView) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:273
// index:0
// inline
// QStringView::const_iterator cbegin()
func (this *QStringView) Cbegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6cbeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:274
// index:0
// inline
// QStringView::const_iterator cend()
func (this *QStringView) Cend() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4cendEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:280
// index:0
// inline
// bool empty()
func (this *QStringView) Empty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5emptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:281
// index:0
// inline
// QChar front()
func (this *QStringView) Front() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5frontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:282
// index:0
// inline
// QChar back()
func (this *QStringView) Back() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4backEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:287
// index:0
// inline
// bool isNull()
func (this *QStringView) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:288
// index:0
// inline
// bool isEmpty()
func (this *QStringView) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:289
// index:0
// inline
// int length()
func (this *QStringView) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:291
// index:0
// inline
// QChar first()
func (this *QStringView) First() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5firstEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:292
// index:0
// inline
// QChar last()
func (this *QStringView) Last() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4lastEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
