package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QStringRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringRef) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStringRefFromPointer(cthis unsafe.Pointer) *QStringRef {
	return &QStringRef{&qtrt.CObject{cthis}}
}
func (*QStringRef) NewFromPointer(cthis unsafe.Pointer) *QStringRef {
	return NewQStringRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1420
// index:0
// Public inline
// void QStringRef()
func NewQStringRef() *QStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1421
// index:1
// Public inline
// void QStringRef(const class QString *, int, int)
func NewQStringRef_1(string *QString /*444 const QString **/, position int, size int) *QStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QStringii", ffiqt.FFI_TYPE_VOID, cthis, convArg0, position, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1422
// index:2
// Public inline
// void QStringRef(const class QString *)
func NewQStringRef_2(string *QString /*444 const QString **/) *QStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1438
// index:0
// Public inline
// void ~QStringRef()
func DeleteQStringRef(*QStringRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRefD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1441
// index:0
// Public inline
// const QString * string()
func (this *QStringRef) String() *QString /*444 const QString **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6stringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1442
// index:0
// Public inline
// int position()
func (this *QStringRef) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1443
// index:0
// Public inline
// int size()
func (this *QStringRef) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1444
// index:0
// Public inline
// int count()
func (this *QStringRef) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public
// int count(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Count_1(s *QString, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public
// int count(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) Count_2(c *QChar /*123*/, cs int) int {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public
// int count(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Count_3(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1445
// index:0
// Public inline
// int length()
func (this *QStringRef) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public
// int indexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf(str *QString, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public
// int indexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_1(ch *QChar /*123*/, from int, cs int) int {
	var convArg0 = ch.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public
// int indexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_2(str *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public
// int indexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_3(str *QStringRef, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public
// int lastIndexOf(const class QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf(str *QString, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public
// int lastIndexOf(class QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_1(ch *QChar /*123*/, from int, cs int) int {
	var convArg0 = ch.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public
// int lastIndexOf(class QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_2(str *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public
// int lastIndexOf(const class QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_3(str *QStringRef, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline
// bool contains(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Contains(str *QString, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline
// bool contains(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) Contains_1(ch *QChar /*123*/, cs int) bool {
	var convArg0 = ch.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline
// bool contains(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Contains_2(str *QLatin1String /*123*/, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline
// bool contains(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Contains_3(str *QStringRef, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1470
// index:0
// Public
// QStringRef left(int)
func (this *QStringRef) Left(n int) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4leftEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1471
// index:0
// Public
// QStringRef right(int)
func (this *QStringRef) Right(n int) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5rightEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public
// QStringRef mid(int, int)
func (this *QStringRef) Mid(pos int, n int) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef3midEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pos, n)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1473
// index:0
// Public inline
// QStringRef chopped(int)
func (this *QStringRef) Chopped(n int) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7choppedEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1476
// index:0
// Public inline
// void truncate(int)
func (this *QStringRef) Truncate(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef8truncateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1477
// index:0
// Public inline
// void chop(int)
func (this *QStringRef) Chop(n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef4chopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1485
// index:0
// Public
// bool isRightToLeft()
func (this *QStringRef) IsRightToLeft() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef13isRightToLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline
// bool startsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public
// bool startsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_1(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public
// bool startsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_2(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public
// bool startsWith(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_3(s *QString, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public
// bool startsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_4(c *QStringRef, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline
// bool endsWith(class QStringView, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public
// bool endsWith(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_1(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public
// bool endsWith(class QChar, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_2(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public
// bool endsWith(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_3(s *QString, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public
// bool endsWith(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_4(c *QStringRef, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1507
// index:0
// Public inline
// const QChar * unicode()
func (this *QStringRef) Unicode() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1513
// index:0
// Public inline
// const QChar * data()
func (this *QStringRef) Data() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1514
// index:0
// Public inline
// const QChar * constData()
func (this *QStringRef) ConstData() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef9constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1516
// index:0
// Public inline
// QStringRef::const_iterator begin()
func (this *QStringRef) Begin() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1517
// index:0
// Public inline
// QStringRef::const_iterator cbegin()
func (this *QStringRef) Cbegin() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6cbeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1518
// index:0
// Public inline
// QStringRef::const_iterator constBegin()
func (this *QStringRef) ConstBegin() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1519
// index:0
// Public inline
// QStringRef::const_iterator end()
func (this *QStringRef) End() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1520
// index:0
// Public inline
// QStringRef::const_iterator cend()
func (this *QStringRef) Cend() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4cendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1521
// index:0
// Public inline
// QStringRef::const_iterator constEnd()
func (this *QStringRef) ConstEnd() *QChar /*444 const QChar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1531
// index:0
// Public
// QByteArray toLatin1()
func (this *QStringRef) ToLatin1() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1532
// index:0
// Public
// QByteArray toUtf8()
func (this *QStringRef) ToUtf8() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toUtf8Ev", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1533
// index:0
// Public
// QByteArray toLocal8Bit()
func (this *QStringRef) ToLocal8Bit() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11toLocal8BitEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1536
// index:0
// Public inline
// void clear()
func (this *QStringRef) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1537
// index:0
// Public
// QString toString()
func (this *QStringRef) ToString() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1538
// index:0
// Public inline
// bool isEmpty()
func (this *QStringRef) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1539
// index:0
// Public inline
// bool isNull()
func (this *QStringRef) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1541
// index:0
// Public
// QStringRef appendTo(class QString *)
func (this *QStringRef) AppendTo(string *QString /*444 QString **/) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8appendToEP7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1543
// index:0
// Public inline
// const QChar at(int)
func (this *QStringRef) At(i int) *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef2atEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1546
// index:0
// Public inline
// QChar front()
func (this *QStringRef) Front() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5frontEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1547
// index:0
// Public inline
// QChar back()
func (this *QStringRef) Back() *QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef4backEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public
// int compare(const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare(s *QString, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public
// int compare(const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_1(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public
// int compare(class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_2(s *QLatin1String /*123*/, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline
// int compare(const class QByteArray &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_3(s *QByteArray, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static
// int compare(const class QStringRef &, const class QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_4(s1 *QStringRef, s2 *QString, arg2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, arg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStringRef_Compare_4(s1 *QStringRef, s2 *QString, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_4(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static
// int compare(const class QStringRef &, const class QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_5(s1 *QStringRef, s2 *QStringRef, arg2 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, arg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStringRef_Compare_5(s1 *QStringRef, s2 *QStringRef, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_5(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static
// int compare(const class QStringRef &, class QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_6(s1 *QStringRef, s2 *QLatin1String /*123*/, cs int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, s1, s2, cs)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStringRef_Compare_6(s1 *QStringRef, s2 *QLatin1String /*123*/, cs int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_6(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1573
// index:0
// Public
// int localeAwareCompare(const class QString &)
func (this *QStringRef) LocaleAwareCompare(s *QString) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1574
// index:1
// Public
// int localeAwareCompare(const class QStringRef &)
func (this *QStringRef) LocaleAwareCompare_1(s *QStringRef) int {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1575
// index:2
// Public static
// int localeAwareCompare(const class QStringRef &, const class QString &)
func (this *QStringRef) LocaleAwareCompare_2(s1 *QStringRef, s2 *QString) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_RK7QString", ffiqt.FFI_TYPE_POINTER, s1, s2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStringRef_LocaleAwareCompare_2(s1 *QStringRef, s2 *QString) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_2(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1576
// index:3
// Public static
// int localeAwareCompare(const class QStringRef &, const class QStringRef &)
func (this *QStringRef) LocaleAwareCompare_3(s1 *QStringRef, s2 *QStringRef) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_S1_", ffiqt.FFI_TYPE_POINTER, s1, s2)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QStringRef_LocaleAwareCompare_3(s1 *QStringRef, s2 *QStringRef) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_3(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1578
// index:0
// Public
// QStringRef trimmed()
func (this *QStringRef) Trimmed() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7trimmedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public
// short toShort(_Bool *, int)
func (this *QStringRef) ToShort(ok unsafe.Pointer /*666*/, base int) int16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int16(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public
// ushort toUShort(_Bool *, int)
func (this *QStringRef) ToUShort(ok unsafe.Pointer /*666*/, base int) uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public
// int toInt(_Bool *, int)
func (this *QStringRef) ToInt(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public
// uint toUInt(_Bool *, int)
func (this *QStringRef) ToUInt(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public
// long toLong(_Bool *, int)
func (this *QStringRef) ToLong(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public
// ulong toULong(_Bool *, int)
func (this *QStringRef) ToULong(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public
// qlonglong toLongLong(_Bool *, int)
func (this *QStringRef) ToLongLong(ok unsafe.Pointer /*666*/, base int) int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public
// qulonglong toULongLong(_Bool *, int)
func (this *QStringRef) ToULongLong(ok unsafe.Pointer /*666*/, base int) uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public
// float toFloat(_Bool *)
func (this *QStringRef) ToFloat(ok unsafe.Pointer /*666*/) float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public
// double toDouble(_Bool *)
func (this *QStringRef) ToDouble(ok unsafe.Pointer /*666*/) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

//  body block end
