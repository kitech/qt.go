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
import "gopp"
import "qt.go/qtrt"

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
func NewQStringRef() *QStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1421
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *, int, int)
func NewQStringRef_1(string string, position int, size int) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QStringii", qtrt.FFI_TYPE_POINTER, convArg0, position, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1422
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QStringRef(const QString *)
func NewQStringRef_2(string string) *QStringRef {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefC2EPK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:1438
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QStringRef()
func DeleteQStringRef(this *QStringRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstring.h:1441
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString * string()
func (this *QStringRef) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1442
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int position()
func (this *QStringRef) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1443
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size()
func (this *QStringRef) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1444
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QStringRef) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1461
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Count_1(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1462
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity)
func (this *QStringRef) Count_2(c *QChar /*123*/, cs int) int {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1463
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Count_3(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1445
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length()
func (this *QStringRef) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1447
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1448
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_1(ch *QChar /*123*/, from int, cs int) int {
	var convArg0 = ch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1449
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_2(str *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1450
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) IndexOf_3(str *QStringRef, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1451
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf(str string, from int, cs int) int {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERK7QStringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1452
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_1(ch *QChar /*123*/, from int, cs int) int {
	var convArg0 = ch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1453
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_2(str *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1454
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity)
func (this *QStringRef) LastIndexOf_3(str *QStringRef, from int, cs int) int {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1456
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1457
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity)
func (this *QStringRef) Contains_1(ch *QChar /*123*/, cs int) bool {
	var convArg0 = ch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1458
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Contains_2(str *QLatin1String /*123*/, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1459
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Contains_3(str *QStringRef, cs int) bool {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1470
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef left(int)
func (this *QStringRef) Left(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1471
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef right(int)
func (this *QStringRef) Right(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1472
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef mid(int, int)
func (this *QStringRef) Mid(pos int, n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1473
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef chopped(int)
func (this *QStringRef) Chopped(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1476
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(int)
func (this *QStringRef) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1477
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(int)
func (this *QStringRef) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1485
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft()
func (this *QStringRef) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1487
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1489
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_1(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1490
// index:2
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_2(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1492
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1493
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) StartsWith_4(c *QStringRef, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1496
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1498
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_1(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1499
// index:2
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_2(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1501
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_3(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1502
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) EndsWith_4(c *QStringRef, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1507
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * unicode()
func (this *QStringRef) Unicode() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1513
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * data()
func (this *QStringRef) Data() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1514
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * constData()
func (this *QStringRef) ConstData() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:1516
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator begin()
func (this *QStringRef) Begin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1517
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cbegin()
func (this *QStringRef) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1518
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constBegin()
func (this *QStringRef) ConstBegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1519
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator end()
func (this *QStringRef) End() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1520
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator cend()
func (this *QStringRef) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1521
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringRef::const_iterator constEnd()
func (this *QStringRef) ConstEnd() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1531
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLatin1()
func (this *QStringRef) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1532
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toUtf8()
func (this *QStringRef) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1533
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit()
func (this *QStringRef) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()
func (this *QStringRef) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1537
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QStringRef) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1538
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QStringRef) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1539
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QStringRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1541
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef appendTo(QString *)
func (this *QStringRef) AppendTo(string string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8appendToEP7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1543
// index:0
// Public inline Visibility=Default Availability=Available
// [2] const QChar at(int)
func (this *QStringRef) At(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1546
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front()
func (this *QStringRef) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1547
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back()
func (this *QStringRef) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1559
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1560
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_1(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1561
// index:2
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_2(s *QLatin1String /*123*/, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1563
// index:3
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QByteArray &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_3(s *QByteArray, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7compareERK10QByteArrayN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1566
// index:4
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QString &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_4(s1 *QStringRef, s2 string, arg2 int) int {
	var convArg0 = s1.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_RK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_4(s1 *QStringRef, s2 string, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_4(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1568
// index:5
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, const QStringRef &, Qt::CaseSensitivity)
func (this *QStringRef) Compare_5(s1 *QStringRef, s2 *QStringRef, arg2 int) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_5(s1 *QStringRef, s2 *QStringRef, arg2 int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_5(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1570
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, QLatin1String, Qt::CaseSensitivity)
func (this *QStringRef) Compare_6(s1 *QStringRef, s2 *QLatin1String /*123*/, cs int) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_Compare_6(s1 *QStringRef, s2 *QLatin1String /*123*/, cs int) int {
	var nilthis *QStringRef
	rv := nilthis.Compare_6(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1573
// index:0
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &)
func (this *QStringRef) LocaleAwareCompare(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1574
// index:1
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &)
func (this *QStringRef) LocaleAwareCompare_1(s *QStringRef) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef18localeAwareCompareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1575
// index:2
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QString &)
func (this *QStringRef) LocaleAwareCompare_2(s1 *QStringRef, s2 string) int {
	var convArg0 = s1.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_RK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_2(s1 *QStringRef, s2 string) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_2(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1576
// index:3
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &, const QStringRef &)
func (this *QStringRef) LocaleAwareCompare_3(s1 *QStringRef, s2 *QStringRef) int {
	var convArg0 = s1.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QStringRef18localeAwareCompareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QStringRef_LocaleAwareCompare_3(s1 *QStringRef, s2 *QStringRef) int {
	var nilthis *QStringRef
	rv := nilthis.LocaleAwareCompare_3(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:1578
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef trimmed()
func (this *QStringRef) Trimmed() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1579
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(_Bool *, int)
func (this *QStringRef) ToShort(ok unsafe.Pointer /*666*/, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1580
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(_Bool *, int)
func (this *QStringRef) ToUShort(ok unsafe.Pointer /*666*/, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1581
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *, int)
func (this *QStringRef) ToInt(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1582
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *, int)
func (this *QStringRef) ToUInt(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1583
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(_Bool *, int)
func (this *QStringRef) ToLong(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1584
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(_Bool *, int)
func (this *QStringRef) ToULong(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1585
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *, int)
func (this *QStringRef) ToLongLong(ok unsafe.Pointer /*666*/, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1586
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *, int)
func (this *QStringRef) ToULongLong(ok unsafe.Pointer /*666*/, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1587
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(_Bool *)
func (this *QStringRef) ToFloat(ok unsafe.Pointer /*666*/) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1588
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(_Bool *)
func (this *QStringRef) ToDouble(ok unsafe.Pointer /*666*/) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QStringRef8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

//  body block end
