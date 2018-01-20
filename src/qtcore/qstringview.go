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
func NewQStringViewFromPointer(cthis unsafe.Pointer) *QStringView {
	return &QStringView{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qstringview.h:171
// index:0
// Public inline
// void QStringView()
func NewQStringView() *QStringView {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringViewC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstringview.h:214
// index:0
// Public inline
// QString toString()
func (this *QStringView) ToString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:216
// index:0
// Public inline
// qsizetype size()
func (this *QStringView) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:217
// index:0
// Public inline
// QStringView::const_pointer data()
func (this *QStringView) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:218
// index:0
// Public inline
// const QStringView::storage_type * utf16()
func (this *QStringView) Utf16() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5utf16Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:227
// index:0
// Public inline
// QByteArray toLatin1()
func (this *QStringView) ToLatin1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:228
// index:0
// Public inline
// QByteArray toUtf8()
func (this *QStringView) ToUtf8() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6toUtf8Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:229
// index:0
// Public inline
// QByteArray toLocal8Bit()
func (this *QStringView) ToLocal8Bit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView11toLocal8BitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:230
// index:0
// Public inline
// QVector<uint> toUcs4()
func (this *QStringView) ToUcs4() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6toUcs4Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:232
// index:0
// Public inline
// QChar at(qsizetype)
func (this *QStringView) At(n int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView2atEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:234
// index:0
// Public inline
// QStringView mid(qsizetype)
func (this *QStringView) Mid(pos int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3midEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:236
// index:1
// Public inline
// QStringView mid(qsizetype, qsizetype)
func (this *QStringView) Mid_1(pos int64, n int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3midExx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos, &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:238
// index:0
// Public inline
// QStringView left(qsizetype)
func (this *QStringView) Left(n int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4leftEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:240
// index:0
// Public inline
// QStringView right(qsizetype)
func (this *QStringView) Right(n int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5rightEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:242
// index:0
// Public inline
// QStringView chopped(qsizetype)
func (this *QStringView) Chopped(n int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7choppedEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:245
// index:0
// Public inline
// void truncate(qsizetype)
func (this *QStringView) Truncate(n int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringView8truncateEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:247
// index:0
// Public inline
// void chop(qsizetype)
func (this *QStringView) Chop(n int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStringView4chopEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:250
// index:0
// Public inline
// QStringView trimmed()
func (this *QStringView) Trimmed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7trimmedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:252
// index:0
// Public inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringView) StartsWith(s *QStringView, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:254
// index:1
// Public inline
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_1(s *QLatin1String, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:255
// index:2
// Public inline
// bool startsWith(class QChar)
func (this *QStringView) StartsWith_2(c *QChar) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:257
// index:3
// Public inline
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_3(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:260
// index:0
// Public inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringView) EndsWith(s *QStringView, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithES_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:262
// index:1
// Public inline
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_1(s *QLatin1String, cs int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:263
// index:2
// Public inline
// bool endsWith(class QChar)
func (this *QStringView) EndsWith_2(c *QChar) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:265
// index:3
// Public inline
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_3(c *QChar, cs int) interface{} {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &cs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:271
// index:0
// Public inline
// QStringView::const_iterator begin()
func (this *QStringView) Begin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:272
// index:0
// Public inline
// QStringView::const_iterator end()
func (this *QStringView) End() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:273
// index:0
// Public inline
// QStringView::const_iterator cbegin()
func (this *QStringView) Cbegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6cbeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:274
// index:0
// Public inline
// QStringView::const_iterator cend()
func (this *QStringView) Cend() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4cendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:280
// index:0
// Public inline
// bool empty()
func (this *QStringView) Empty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5emptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:281
// index:0
// Public inline
// QChar front()
func (this *QStringView) Front() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:282
// index:0
// Public inline
// QChar back()
func (this *QStringView) Back() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:287
// index:0
// Public inline
// bool isNull()
func (this *QStringView) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:288
// index:0
// Public inline
// bool isEmpty()
func (this *QStringView) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:289
// index:0
// Public inline
// int length()
func (this *QStringView) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:291
// index:0
// Public inline
// QChar first()
func (this *QStringView) First() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView5firstEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qstringview.h:292
// index:0
// Public inline
// QChar last()
func (this *QStringView) Last() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStringView4lastEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
