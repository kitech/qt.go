//  header block begin
// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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
type QStringRef struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qstring.h:1420
// index:0
// inline
// void QStringRef()
func NewQStringRef() *QStringRef {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStringRef{cthis}
}

// /usr/include/qt/QtCore/qstring.h:1421
// index:1
// inline
// void QStringRef(const class QString *, int, int)
func NewQStringRef_1(string unsafe.Pointer, position int, size int) *QStringRef {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QStringii", ffiqt.FFI_TYPE_VOID, cthis, string, &position, &size)
	gopp.ErrPrint(err, rv)
	return &QStringRef{cthis}
}

// /usr/include/qt/QtCore/qstring.h:1422
// index:2
// inline
// void QStringRef(const class QString *)
func NewQStringRef_2(string unsafe.Pointer) *QStringRef {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QString", ffiqt.FFI_TYPE_VOID, cthis, string)
	gopp.ErrPrint(err, rv)
	return &QStringRef{cthis}
}

// /usr/include/qt/QtCore/qstring.h:1438
// index:0
// inline
// void ~QStringRef()
func DeleteQStringRef(*QStringRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1441
// index:0
// inline
// const QString * string()
func (this *QStringRef) String() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6stringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1442
// index:0
// inline
// int position()
func (this *QStringRef) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1443
// index:0
// inline
// int size()
func (this *QStringRef) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1444
// index:0
// inline
// int count()
func (this *QStringRef) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// int count(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Count_1(s unsafe.Pointer, cs int) {
	// 1: (, const QString & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// int count(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) Count_2(c unsafe.Pointer, cs int) {
	// 2: (, QChar c, Qt::CaseSensitivity cs), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// int count(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Count_3(s unsafe.Pointer, cs int) {
	// 3: (, const QStringRef & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1445
// index:0
// inline
// int length()
func (this *QStringRef) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6lengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// int indexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf(str unsafe.Pointer, from int, cs int) {
	// 0: (, const QString & str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// int indexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_1(ch unsafe.Pointer, from int, cs int) {
	// 1: (, QChar ch, int from, Qt::CaseSensitivity cs), (ch, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, ch, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// int indexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_2(str unsafe.Pointer, from int, cs int) {
	// 2: (, QLatin1String str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// int indexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_3(str unsafe.Pointer, from int, cs int) {
	// 3: (, const QStringRef & str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// int lastIndexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf(str unsafe.Pointer, from int, cs int) {
	// 0: (, const QString & str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// int lastIndexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_1(ch unsafe.Pointer, from int, cs int) {
	// 1: (, QChar ch, int from, Qt::CaseSensitivity cs), (ch, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, ch, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// int lastIndexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_2(str unsafe.Pointer, from int, cs int) {
	// 2: (, QLatin1String str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// int lastIndexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_3(str unsafe.Pointer, from int, cs int) {
	// 3: (, const QStringRef & str, int from, Qt::CaseSensitivity cs), (str, &from, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &from, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// inline
// bool contains(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Contains(str unsafe.Pointer, cs int) {
	// 0: (, const QString & str, Qt::CaseSensitivity cs), (str, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// inline
// bool contains(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) Contains_1(ch unsafe.Pointer, cs int) {
	// 1: (, QChar ch, Qt::CaseSensitivity cs), (ch, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, ch, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// inline
// bool contains(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Contains_2(str unsafe.Pointer, cs int) {
	// 2: (, QLatin1String str, Qt::CaseSensitivity cs), (str, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// inline
// bool contains(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Contains_3(str unsafe.Pointer, cs int) {
	// 3: (, const QStringRef & str, Qt::CaseSensitivity cs), (str, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, str, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1465
// index:0
// QVector<QStringRef> split(const class QString &, class QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QStringRef) Split(sep unsafe.Pointer, behavior int, cs int) {
	// 0: (, const QString & sep, QString::SplitBehavior behavior, Qt::CaseSensitivity cs), (sep, &behavior, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5splitERK7QStringNS0_13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, sep, &behavior, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1467
// index:1
// QVector<QStringRef> split(class QChar, class QString::SplitBehavior, Qt::CaseSensitivity)
func (this *QStringRef) Split_1(sep unsafe.Pointer, behavior int, cs int) {
	// 1: (, QChar sep, QString::SplitBehavior behavior, Qt::CaseSensitivity cs), (sep, &behavior, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5splitE5QCharN7QString13SplitBehaviorEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, sep, &behavior, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1470
// index:0
// QStringRef left(int)
func (this *QStringRef) Left(n int) {
	// 0: (, int n), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4leftEi", ffiqt.FFI_TYPE_VOID, this.cthis, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1471
// index:0
// QStringRef right(int)
func (this *QStringRef) Right(n int) {
	// 0: (, int n), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5rightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// QStringRef mid(int, int)
func (this *QStringRef) Mid(pos int, n int) {
	// 0: (, int pos, int n), (&pos, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef3midEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1473
// index:0
// inline
// QStringRef chopped(int)
func (this *QStringRef) Chopped(n int) {
	// 0: (, int n), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7choppedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1476
// index:0
// inline
// void truncate(int)
func (this *QStringRef) Truncate(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef8truncateEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1477
// index:0
// inline
// void chop(int)
func (this *QStringRef) Chop(n int) {
	// 0: (, int n), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef4chopEi", ffiqt.FFI_TYPE_VOID, this.cthis, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1485
// index:0
// bool isRightToLeft()
func (this *QStringRef) IsRightToLeft() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef13isRightToLeftEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith(s unsafe.Pointer, cs int) {
	// 0: (, QStringView s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, QLatin1String s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_2(c unsafe.Pointer, cs int) {
	// 2: (, QChar c, Qt::CaseSensitivity cs), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// bool startsWith(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_3(s unsafe.Pointer, cs int) {
	// 3: (, const QString & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// bool startsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_4(c unsafe.Pointer, cs int) {
	// 4: (, const QStringRef & c, Qt::CaseSensitivity cs), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith(s unsafe.Pointer, cs int) {
	// 0: (, QStringView s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_1(s unsafe.Pointer, cs int) {
	// 1: (, QLatin1String s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_2(c unsafe.Pointer, cs int) {
	// 2: (, QChar c, Qt::CaseSensitivity cs), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// bool endsWith(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_3(s unsafe.Pointer, cs int) {
	// 3: (, const QString & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// bool endsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_4(c unsafe.Pointer, cs int) {
	// 4: (, const QStringRef & c, Qt::CaseSensitivity cs), (c, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, c, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1507
// index:0
// inline
// const QChar * unicode()
func (this *QStringRef) Unicode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7unicodeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1513
// index:0
// inline
// const QChar * data()
func (this *QStringRef) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1514
// index:0
// inline
// const QChar * constData()
func (this *QStringRef) ConstData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef9constDataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1516
// index:0
// inline
// QStringRef::const_iterator begin()
func (this *QStringRef) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5beginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1517
// index:0
// inline
// QStringRef::const_iterator cbegin()
func (this *QStringRef) Cbegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6cbeginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1518
// index:0
// inline
// QStringRef::const_iterator constBegin()
func (this *QStringRef) ConstBegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10constBeginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1519
// index:0
// inline
// QStringRef::const_iterator end()
func (this *QStringRef) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef3endEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1520
// index:0
// inline
// QStringRef::const_iterator cend()
func (this *QStringRef) Cend() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4cendEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1521
// index:0
// inline
// QStringRef::const_iterator constEnd()
func (this *QStringRef) ConstEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8constEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1531
// index:0
// QByteArray toLatin1()
func (this *QStringRef) ToLatin1() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toLatin1Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1532
// index:0
// QByteArray toUtf8()
func (this *QStringRef) ToUtf8() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toUtf8Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1533
// index:0
// QByteArray toLocal8Bit()
func (this *QStringRef) ToLocal8Bit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11toLocal8BitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1534
// index:0
// QVector<uint> toUcs4()
func (this *QStringRef) ToUcs4() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toUcs4Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1536
// index:0
// inline
// void clear()
func (this *QStringRef) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1537
// index:0
// QString toString()
func (this *QStringRef) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1538
// index:0
// inline
// bool isEmpty()
func (this *QStringRef) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1539
// index:0
// inline
// bool isNull()
func (this *QStringRef) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1541
// index:0
// QStringRef appendTo(class QString *)
func (this *QStringRef) AppendTo(string unsafe.Pointer) {
	// 0: (, QString * string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8appendToEP7QString", ffiqt.FFI_TYPE_VOID, this.cthis, string)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1543
// index:0
// inline
// const QChar at(int)
func (this *QStringRef) At(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef2atEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1546
// index:0
// inline
// QChar front()
func (this *QStringRef) Front() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5frontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1547
// index:0
// inline
// QChar back()
func (this *QStringRef) Back() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4backEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// int compare(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare(s unsafe.Pointer, cs int) {
	// 0: (, const QString & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// int compare(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_1(s unsafe.Pointer, cs int) {
	// 1: (, const QStringRef & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// int compare(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_2(s unsafe.Pointer, cs int) {
	// 2: (, QLatin1String s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// inline
// int compare(const class QByteArray &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_3(s unsafe.Pointer, cs int) {
	// 3: (, const QByteArray & s, Qt::CaseSensitivity cs), (s, &cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, s, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// static
// int compare(const class QStringRef &, const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_4(s1 unsafe.Pointer, s2 unsafe.Pointer, arg2 int) {
	// 4: (const QStringRef & s1, const QString & s2, Qt::CaseSensitivity arg2), (s1, s2, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStringRef_Compare_4(s1 unsafe.Pointer, s2 unsafe.Pointer, arg2 int) {
	// 4: (const QStringRef & s1, const QString & s2, Qt::CaseSensitivity arg2), (s1, s2, arg2)
	var nilthis *QStringRef
	nilthis.Compare_4(s1, s2, arg2)
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// static
// int compare(const class QStringRef &, const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_5(s1 unsafe.Pointer, s2 unsafe.Pointer, arg2 int) {
	// 5: (const QStringRef & s1, const QStringRef & s2, Qt::CaseSensitivity arg2), (s1, s2, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStringRef_Compare_5(s1 unsafe.Pointer, s2 unsafe.Pointer, arg2 int) {
	// 5: (const QStringRef & s1, const QStringRef & s2, Qt::CaseSensitivity arg2), (s1, s2, arg2)
	var nilthis *QStringRef
	nilthis.Compare_5(s1, s2, arg2)
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// static
// int compare(const class QStringRef &, class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_6(s1 unsafe.Pointer, s2 unsafe.Pointer, cs int) {
	// 6: (const QStringRef & s1, QLatin1String s2, Qt::CaseSensitivity cs), (s1, s2, cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStringRef_Compare_6(s1 unsafe.Pointer, s2 unsafe.Pointer, cs int) {
	// 6: (const QStringRef & s1, QLatin1String s2, Qt::CaseSensitivity cs), (s1, s2, cs)
	var nilthis *QStringRef
	nilthis.Compare_6(s1, s2, cs)
}

// /usr/include/qt/QtCore/qstring.h:1573
// index:0
// int localeAwareCompare(const class QString &)
func (this *QStringRef) LocaleAwareCompare(s unsafe.Pointer) {
	// 0: (, const QString & s), (s)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1574
// index:1
// int localeAwareCompare(const class QStringRef &)
func (this *QStringRef) LocaleAwareCompare_1(s unsafe.Pointer) {
	// 1: (, const QStringRef & s), (s)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1575
// index:2
// static
// int localeAwareCompare(const class QStringRef &, const class QString &)
func (this *QStringRef) LocaleAwareCompare_2(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 2: (const QStringRef & s1, const QString & s2), (s1, s2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_RK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStringRef_LocaleAwareCompare_2(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 2: (const QStringRef & s1, const QString & s2), (s1, s2)
	var nilthis *QStringRef
	nilthis.LocaleAwareCompare_2(s1, s2)
}

// /usr/include/qt/QtCore/qstring.h:1576
// index:3
// static
// int localeAwareCompare(const class QStringRef &, const class QStringRef &)
func (this *QStringRef) LocaleAwareCompare_3(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 3: (const QStringRef & s1, const QStringRef & s2), (s1, s2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_S1_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStringRef_LocaleAwareCompare_3(s1 unsafe.Pointer, s2 unsafe.Pointer) {
	// 3: (const QStringRef & s1, const QStringRef & s2), (s1, s2)
	var nilthis *QStringRef
	nilthis.LocaleAwareCompare_3(s1, s2)
}

// /usr/include/qt/QtCore/qstring.h:1578
// index:0
// QStringRef trimmed()
func (this *QStringRef) Trimmed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7trimmedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// short toShort(_Bool *, int)
func (this *QStringRef) ToShort(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// ushort toUShort(_Bool *, int)
func (this *QStringRef) ToUShort(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// int toInt(_Bool *, int)
func (this *QStringRef) ToInt(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// uint toUInt(_Bool *, int)
func (this *QStringRef) ToUInt(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// long toLong(_Bool *, int)
func (this *QStringRef) ToLong(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// ulong toULong(_Bool *, int)
func (this *QStringRef) ToULong(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// qlonglong toLongLong(_Bool *, int)
func (this *QStringRef) ToLongLong(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// qulonglong toULongLong(_Bool *, int)
func (this *QStringRef) ToULongLong(ok unsafe.Pointer, base int) {
	// 0: (, bool * ok, int base), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", ffiqt.FFI_TYPE_VOID, this.cthis, ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// float toFloat(_Bool *)
func (this *QStringRef) ToFloat(ok unsafe.Pointer) {
	// 0: (, bool * ok), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", ffiqt.FFI_TYPE_VOID, this.cthis, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// double toDouble(_Bool *)
func (this *QStringRef) ToDouble(ok unsafe.Pointer) {
	// 0: (, bool * ok), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", ffiqt.FFI_TYPE_VOID, this.cthis, ok)
	gopp.ErrPrint(err, rv)
}

//  body block end
