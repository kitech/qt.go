package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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

type QString struct {
	*qtrt.CObject
}

func (this *QString) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QString) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringFromPointer(cthis unsafe.Pointer) *QString {
	return &QString{&qtrt.CObject{cthis}}
}
func (*QString) NewFromPointer(cthis unsafe.Pointer) *QString {
	return NewQStringFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QString()
func NewQString() *QString {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:218
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QString(const QChar *, int)
func NewQString_1(unicode *QChar /*777 const QChar **/, size int) *QString {
	var convArg0 = unicode.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2EPK5QChari", qtrt.FFI_TYPE_POINTER, convArg0, size)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:219
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QString(QChar)
func NewQString_2(c *QChar /*123*/) *QString {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2E5QChar", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:220
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QString(int, QChar)
func NewQString_3(size int, c *QChar /*123*/) *QString {
	var convArg1 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2Ei5QChar", qtrt.FFI_TYPE_POINTER, size, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:221
// index:4
// Public inline Visibility=Default Availability=Available
// [-2] void QString(QLatin1String)
func NewQString_4(latin1 *QLatin1String /*123*/) *QString {
	var convArg0 = latin1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2E13QLatin1String", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:682
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void QString(const char *)
func NewQString_5(ch string) *QString {
	var convArg0 = qtrt.CString(ch)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:685
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void QString(const QByteArray &)
func NewQString_6(a *QByteArray) *QString {
	var convArg0 = a.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:811
// index:7
// Public Visibility=Default Availability=Available
// [-2] void QString(int, Qt::Initialization)
func NewQString_7(size int, arg1 int) *QString {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringC2EiN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, size, arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

// /usr/include/qt/QtCore/qstring.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QString()
func DeleteQString(this *QString) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QStringD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qstring.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QString &)
func (this *QString) Swap(other string) {
	var tmpArg0 = NewQString_5(other)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:233
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size()
func (this *QString) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:234
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QString) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:335
// index:1
// Public Visibility=Default Availability=Available
// [4] int count(QChar, Qt::CaseSensitivity)
func (this *QString) Count_1(c *QChar /*123*/, cs int) int {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:336
// index:2
// Public Visibility=Default Availability=Available
// [4] int count(const QString &, Qt::CaseSensitivity)
func (this *QString) Count_2(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:337
// index:3
// Public Visibility=Default Availability=Available
// [4] int count(const QStringRef &, Qt::CaseSensitivity)
func (this *QString) Count_3(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countERK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:343
// index:4
// Public Visibility=Default Availability=Available
// [4] int count(const QRegExp &)
func (this *QString) Count_4(arg0 *QRegExp) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:357
// index:5
// Public Visibility=Default Availability=Available
// [4] int count(const QRegularExpression &)
func (this *QString) Count_5(re *QRegularExpression) int {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5countERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length()
func (this *QString) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QString) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int)
func (this *QString) Resize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6resizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:238
// index:1
// Public Visibility=Default Availability=Available
// [-2] void resize(int, QChar)
func (this *QString) Resize_1(size int, fillChar *QChar /*123*/) {
	var convArg1 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6resizeEi5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:240
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & fill(QChar, int)
func (this *QString) Fill(c *QChar /*123*/, size int) string {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString4fillE5QChari", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void truncate(int)
func (this *QString) Truncate(pos int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8truncateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void chop(int)
func (this *QString) Chop(n int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString4chopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int capacity()
func (this *QString) Capacity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8capacityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:245
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void reserve(int)
func (this *QString) Reserve(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7reserveEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void squeeze()
func (this *QString) Squeeze() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7squeezeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:248
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * unicode()
func (this *QString) Unicode() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:249
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QChar * data()
func (this *QString) Data() *QChar /*777 QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:250
// index:1
// Public inline Visibility=Default Availability=Available
// [8] const QChar * data()
func (this *QString) Data_1() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:251
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QChar * constData()
func (this *QString) ConstData() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString9constDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstring.h:253
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void detach()
func (this *QString) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:254
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QString) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:255
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSharedWith(const QString &)
func (this *QString) IsSharedWith(other string) bool {
	var tmpArg0 = NewQString_5(other)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString12isSharedWithERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QString) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:258
// index:0
// Public inline Visibility=Default Availability=Available
// [2] const QChar at(int)
func (this *QString) At(i int) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString2atEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:264
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front()
func (this *QString) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:265
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCharRef front()
func (this *QString) Front_1() *QCharRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back()
func (this *QString) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:267
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCharRef back()
func (this *QString) Back_1() *QCharRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:269
// index:0
// Public Visibility=Default Availability=Available
// [8] QString arg(qlonglong, int, int, QChar)
func (this *QString) Arg(a int64, fieldwidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argExii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldwidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:271
// index:1
// Public Visibility=Default Availability=Available
// [8] QString arg(qulonglong, int, int, QChar)
func (this *QString) Arg_1(a uint64, fieldwidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEyii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldwidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:273
// index:2
// Public Visibility=Default Availability=Available
// [8] QString arg(long, int, int, QChar)
func (this *QString) Arg_2(a int, fieldwidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argElii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldwidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:275
// index:3
// Public Visibility=Default Availability=Available
// [8] QString arg(ulong, int, int, QChar)
func (this *QString) Arg_3(a uint, fieldwidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEmii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldwidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:277
// index:4
// Public Visibility=Default Availability=Available
// [8] QString arg(int, int, int, QChar)
func (this *QString) Arg_4(a int, fieldWidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEiii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:279
// index:5
// Public Visibility=Default Availability=Available
// [8] QString arg(uint, int, int, QChar)
func (this *QString) Arg_5(a uint, fieldWidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEjii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:281
// index:6
// Public Visibility=Default Availability=Available
// [8] QString arg(short, int, int, QChar)
func (this *QString) Arg_6(a int16, fieldWidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEsii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:283
// index:7
// Public Visibility=Default Availability=Available
// [8] QString arg(ushort, int, int, QChar)
func (this *QString) Arg_7(a uint16, fieldWidth int, base int, fillChar *QChar /*123*/) string {
	var convArg3 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEtii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, base, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:285
// index:8
// Public Visibility=Default Availability=Available
// [8] QString arg(double, int, char, int, QChar)
func (this *QString) Arg_8(a float64, fieldWidth int, fmt byte, prec int, fillChar *QChar /*123*/) string {
	var convArg4 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEdici5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, fmt, prec, convArg4)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:287
// index:9
// Public Visibility=Default Availability=Available
// [8] QString arg(char, int, QChar)
func (this *QString) Arg_9(a byte, fieldWidth int, fillChar *QChar /*123*/) string {
	var convArg2 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argEci5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:289
// index:10
// Public Visibility=Default Availability=Available
// [8] QString arg(QChar, int, QChar)
func (this *QString) Arg_10(a *QChar /*123*/, fieldWidth int, fillChar *QChar /*123*/) string {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argE5QChariS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:292
// index:11
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, int, QChar)
func (this *QString) Arg_11(a string, fieldWidth int, fillChar *QChar /*123*/) string {
	var tmpArg0 = NewQString_5(a)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_i5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:295
// index:12
// Public Visibility=Default Availability=Available
// [8] QString arg(QStringView, int, QChar)
func (this *QString) Arg_12(a *QStringView /*123*/, fieldWidth int, fillChar *QChar /*123*/) string {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argE11QStringViewi5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:297
// index:13
// Public Visibility=Default Availability=Available
// [8] QString arg(QLatin1String, int, QChar)
func (this *QString) Arg_13(a *QLatin1String /*123*/, fieldWidth int, fillChar *QChar /*123*/) string {
	var convArg0 = a.GetCthis()
	var convArg2 = fillChar.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argE13QLatin1Stringi5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fieldWidth, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:299
// index:14
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &)
func (this *QString) Arg_14(a1 string, a2 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:300
// index:15
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &)
func (this *QString) Arg_15(a1 string, a2 string, a3 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:301
// index:16
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_16(a1 string, a2 string, a3 string, a4 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:303
// index:17
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_17(a1 string, a2 string, a3 string, a4 string, a5 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString_5(a5)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:305
// index:18
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_18(a1 string, a2 string, a3 string, a4 string, a5 string, a6 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString_5(a5)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = NewQString_5(a6)
	var convArg5 = tmpArg5.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:307
// index:19
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_19(a1 string, a2 string, a3 string, a4 string, a5 string, a6 string, a7 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString_5(a5)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = NewQString_5(a6)
	var convArg5 = tmpArg5.GetCthis()
	var tmpArg6 = NewQString_5(a7)
	var convArg6 = tmpArg6.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:310
// index:20
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_20(a1 string, a2 string, a3 string, a4 string, a5 string, a6 string, a7 string, a8 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString_5(a5)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = NewQString_5(a6)
	var convArg5 = tmpArg5.GetCthis()
	var tmpArg6 = NewQString_5(a7)
	var convArg6 = tmpArg6.GetCthis()
	var tmpArg7 = NewQString_5(a8)
	var convArg7 = tmpArg7.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:313
// index:21
// Public Visibility=Default Availability=Available
// [8] QString arg(const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &, const QString &)
func (this *QString) Arg_21(a1 string, a2 string, a3 string, a4 string, a5 string, a6 string, a7 string, a8 string, a9 string) string {
	var tmpArg0 = NewQString_5(a1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(a2)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(a3)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(a4)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString_5(a5)
	var convArg4 = tmpArg4.GetCthis()
	var tmpArg5 = NewQString_5(a6)
	var convArg5 = tmpArg5.GetCthis()
	var tmpArg6 = NewQString_5(a7)
	var convArg6 = tmpArg6.GetCthis()
	var tmpArg7 = NewQString_5(a8)
	var convArg7 = tmpArg7.GetCthis()
	var tmpArg8 = NewQString_5(a9)
	var convArg8 = tmpArg8.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:322
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QChar, int, Qt::CaseSensitivity)
func (this *QString) IndexOf(c *QChar /*123*/, from int, cs int) int {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:323
// index:1
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QString &, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_1(s string, from int, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:324
// index:2
// Public Visibility=Default Availability=Available
// [4] int indexOf(QLatin1String, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_2(s *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:325
// index:3
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QStringRef &, int, Qt::CaseSensitivity)
func (this *QString) IndexOf_3(s *QStringRef, from int, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK10QStringRefiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:340
// index:4
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int)
func (this *QString) IndexOf_4(arg0 *QRegExp, from int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:345
// index:5
// Public Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int)
func (this *QString) IndexOf_5(arg0 *QRegExp, from int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:351
// index:6
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int)
func (this *QString) IndexOf_6(re *QRegularExpression, from int) int {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:352
// index:7
// Public Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int, QRegularExpressionMatch *)
func (this *QString) IndexOf_7(re *QRegularExpression, from int, rmatch *QRegularExpressionMatch /*777 QRegularExpressionMatch **/) int {
	var convArg0 = re.GetCthis()
	var convArg2 = rmatch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, convArg2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:326
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QChar, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf(c *QChar /*123*/, from int, cs int) int {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfE5QChariN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:327
// index:1
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QString &, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_1(s string, from int, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERKS_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:328
// index:2
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QLatin1String, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_2(s *QLatin1String /*123*/, from int, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:329
// index:3
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QStringRef &, int, Qt::CaseSensitivity)
func (this *QString) LastIndexOf_3(s *QStringRef, from int, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK10QStringRefiN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:341
// index:4
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int)
func (this *QString) LastIndexOf_4(arg0 *QRegExp, from int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:346
// index:5
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int)
func (this *QString) LastIndexOf_5(arg0 *QRegExp, from int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:353
// index:6
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int)
func (this *QString) LastIndexOf_6(re *QRegularExpression, from int) int {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:354
// index:7
// Public Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int, QRegularExpressionMatch *)
func (this *QString) LastIndexOf_7(re *QRegularExpression, from int, rmatch *QRegularExpressionMatch /*777 QRegularExpressionMatch **/) int {
	var convArg0 = re.GetCthis()
	var convArg2 = rmatch.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, convArg2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:331
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QChar, Qt::CaseSensitivity)
func (this *QString) Contains(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:332
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity)
func (this *QString) Contains_1(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:333
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity)
func (this *QString) Contains_2(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:334
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QStringRef &, Qt::CaseSensitivity)
func (this *QString) Contains_3(s *QStringRef, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:342
// index:4
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QRegExp &)
func (this *QString) Contains_4(rx *QRegExp) bool {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:347
// index:5
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QRegExp &)
func (this *QString) Contains_5(rx *QRegExp) bool {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsER7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:355
// index:6
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRegularExpression &)
func (this *QString) Contains_6(re *QRegularExpression) bool {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:356
// index:7
// Public Visibility=Default Availability=Available
// [1] bool contains(const QRegularExpression &, QRegularExpressionMatch *)
func (this *QString) Contains_7(re *QRegularExpression, match_ *QRegularExpressionMatch /*777 QRegularExpressionMatch **/) bool {
	var convArg0 = re.GetCthis()
	var convArg1 = match_.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:369
// index:0
// Public Visibility=Default Availability=Available
// [8] QString section(QChar, int, int, QString::SectionFlags)
func (this *QString) Section(sep *QChar /*123*/, start int, end int, flags int) string {
	var convArg0 = sep.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionE5QCharii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end, flags)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:370
// index:1
// Public Visibility=Default Availability=Available
// [8] QString section(const QString &, int, int, QString::SectionFlags)
func (this *QString) Section_1(in_sep string, start int, end int, flags int) string {
	var tmpArg0 = NewQString_5(in_sep)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERKS_ii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end, flags)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:372
// index:2
// Public Visibility=Default Availability=Available
// [8] QString section(const QRegExp &, int, int, QString::SectionFlags)
func (this *QString) Section_2(reg *QRegExp, start int, end int, flags int) string {
	var convArg0 = reg.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERK7QRegExpii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end, flags)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:375
// index:3
// Public Visibility=Default Availability=Available
// [8] QString section(const QRegularExpression &, int, int, QString::SectionFlags)
func (this *QString) Section_3(re *QRegularExpression, start int, end int, flags int) string {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7sectionERK18QRegularExpressionii6QFlagsINS_11SectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end, flags)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:377
// index:0
// Public Visibility=Default Availability=Available
// [8] QString left(int)
func (this *QString) Left(n int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString4leftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:378
// index:0
// Public Visibility=Default Availability=Available
// [8] QString right(int)
func (this *QString) Right(n int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5rightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:379
// index:0
// Public Visibility=Default Availability=Available
// [8] QString mid(int, int)
func (this *QString) Mid(position int, n int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3midEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position, n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:380
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString chopped(int)
func (this *QString) Chopped(n int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7choppedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:384
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef leftRef(int)
func (this *QString) LeftRef(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7leftRefEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:385
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef rightRef(int)
func (this *QString) RightRef(n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8rightRefEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:386
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef midRef(int, int)
func (this *QString) MidRef(position int, n int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6midRefEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position, n)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:389
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QString &, Qt::CaseSensitivity)
func (this *QString) StartsWith(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10startsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:390
// index:1
// Public Visibility=Default Availability=Available
// [1] bool startsWith(const QStringRef &, Qt::CaseSensitivity)
func (this *QString) StartsWith_1(s *QStringRef, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10startsWithERK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:392
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity)
func (this *QString) StartsWith_2(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10startsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:394
// index:3
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QString) StartsWith_3(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:395
// index:4
// Public Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity)
func (this *QString) StartsWith_4(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:398
// index:0
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QString &, Qt::CaseSensitivity)
func (this *QString) EndsWith(s string, cs int) bool {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8endsWithERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:399
// index:1
// Public Visibility=Default Availability=Available
// [1] bool endsWith(const QStringRef &, Qt::CaseSensitivity)
func (this *QString) EndsWith_1(s *QStringRef, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8endsWithERK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:401
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity)
func (this *QString) EndsWith_2(s *QStringView /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8endsWithE11QStringViewN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:403
// index:3
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QString) EndsWith_3(s *QLatin1String /*123*/, cs int) bool {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:404
// index:4
// Public Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity)
func (this *QString) EndsWith_4(c *QChar /*123*/, cs int) bool {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:406
// index:0
// Public Visibility=Default Availability=Available
// [8] QString leftJustified(int, QChar, _Bool)
func (this *QString) LeftJustified(width int, fill *QChar /*123*/, trunc bool) string {
	var convArg1 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString13leftJustifiedEi5QCharb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, convArg1, trunc)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:407
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rightJustified(int, QChar, _Bool)
func (this *QString) RightJustified(width int, fill *QChar /*123*/, trunc bool) string {
	var convArg1 = fill.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString14rightJustifiedEi5QCharb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, convArg1, trunc)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:417
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toLower()
func (this *QString) ToLower() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:419
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString toLower()
func (this *QString) ToLower_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:421
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toUpper()
func (this *QString) ToUpper() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:423
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString toUpper()
func (this *QString) ToUpper_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:425
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toCaseFolded()
func (this *QString) ToCaseFolded() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString12toCaseFoldedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:427
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString toCaseFolded()
func (this *QString) ToCaseFolded_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString12toCaseFoldedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:429
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString trimmed()
func (this *QString) Trimmed() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:431
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString trimmed()
func (this *QString) Trimmed_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:433
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString simplified()
func (this *QString) Simplified() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:435
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString simplified()
func (this *QString) Simplified_1() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString10simplifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:447
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtmlEscaped()
func (this *QString) ToHtmlEscaped() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString13toHtmlEscapedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:478
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & remove(int, int)
func (this *QString) Remove(i int, len int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, len)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:479
// index:1
// Public Visibility=Default Availability=Available
// [8] QString & remove(QChar, Qt::CaseSensitivity)
func (this *QString) Remove_1(c *QChar /*123*/, cs int) string {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:480
// index:2
// Public Visibility=Default Availability=Available
// [8] QString & remove(const QString &, Qt::CaseSensitivity)
func (this *QString) Remove_2(s string, cs int) string {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:495
// index:3
// Public inline Visibility=Default Availability=Available
// [8] QString & remove(const QRegExp &)
func (this *QString) Remove_3(rx *QRegExp) string {
	var convArg0 = rx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:500
// index:4
// Public inline Visibility=Default Availability=Available
// [8] QString & remove(const QRegularExpression &)
func (this *QString) Remove_4(re *QRegularExpression) string {
	var convArg0 = re.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6removeERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:481
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & replace(int, int, QChar)
func (this *QString) Replace(i int, len int, after *QChar /*123*/) string {
	var convArg2 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceEii5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, len, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:482
// index:1
// Public Visibility=Default Availability=Available
// [8] QString & replace(int, int, const QChar *, int)
func (this *QString) Replace_1(i int, len int, s *QChar /*777 const QChar **/, slen int) string {
	var convArg2 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceEiiPK5QChari", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, len, convArg2, slen)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:483
// index:2
// Public Visibility=Default Availability=Available
// [8] QString & replace(int, int, const QString &)
func (this *QString) Replace_2(i int, len int, after string) string {
	var tmpArg2 = NewQString_5(after)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceEiiRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, len, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:484
// index:3
// Public Visibility=Default Availability=Available
// [8] QString & replace(QChar, QChar, Qt::CaseSensitivity)
func (this *QString) Replace_3(before *QChar /*123*/, after *QChar /*123*/, cs int) string {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceE5QCharS0_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:485
// index:4
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QChar *, int, const QChar *, int, Qt::CaseSensitivity)
func (this *QString) Replace_4(before *QChar /*777 const QChar **/, blen int, after *QChar /*777 const QChar **/, alen int, cs int) string {
	var convArg0 = before.GetCthis()
	var convArg2 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceEPK5QChariS2_iN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, blen, convArg2, alen, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:486
// index:5
// Public Visibility=Default Availability=Available
// [8] QString & replace(QLatin1String, QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_5(before *QLatin1String /*123*/, after *QLatin1String /*123*/, cs int) string {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceE13QLatin1StringS0_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:487
// index:6
// Public Visibility=Default Availability=Available
// [8] QString & replace(QLatin1String, const QString &, Qt::CaseSensitivity)
func (this *QString) Replace_6(before *QLatin1String /*123*/, after string, cs int) string {
	var convArg0 = before.GetCthis()
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceE13QLatin1StringRKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:488
// index:7
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QString &, QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_7(before string, after *QLatin1String /*123*/, cs int) string {
	var tmpArg0 = NewQString_5(before)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:489
// index:8
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QString &, const QString &, Qt::CaseSensitivity)
func (this *QString) Replace_8(before string, after string, cs int) string {
	var tmpArg0 = NewQString_5(before)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:491
// index:9
// Public Visibility=Default Availability=Available
// [8] QString & replace(QChar, const QString &, Qt::CaseSensitivity)
func (this *QString) Replace_9(c *QChar /*123*/, after string, cs int) string {
	var convArg0 = c.GetCthis()
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceE5QCharRKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:492
// index:10
// Public Visibility=Default Availability=Available
// [8] QString & replace(QChar, QLatin1String, Qt::CaseSensitivity)
func (this *QString) Replace_10(c *QChar /*123*/, after *QLatin1String /*123*/, cs int) string {
	var convArg0 = c.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceE5QChar13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:494
// index:11
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QRegExp &, const QString &)
func (this *QString) Replace_11(rx *QRegExp, after string) string {
	var convArg0 = rx.GetCthis()
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceERK7QRegExpRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:499
// index:12
// Public Visibility=Default Availability=Available
// [8] QString & replace(const QRegularExpression &, const QString &)
func (this *QString) Replace_12(re *QRegularExpression, after string) string {
	var convArg0 = re.GetCthis()
	var tmpArg1 = NewQString_5(after)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7replaceERK18QRegularExpressionRKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:528
// index:0
// Public Visibility=Default Availability=Available
// [8] QString normalized(enum QString::NormalizationForm, QChar::UnicodeVersion)
func (this *QString) Normalized(mode int, version int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10normalizedENS_17NormalizationFormEN5QChar14UnicodeVersionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, version)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:530
// index:0
// Public Visibility=Default Availability=Available
// [8] QString repeated(int)
func (this *QString) Repeated(times int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8repeatedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), times)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:532
// index:0
// Public Visibility=Default Availability=Available
// [8] const ushort * utf16()
func (this *QString) Utf16() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5utf16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qstring.h:535
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLatin1()
func (this *QString) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:537
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLatin1()
func (this *QString) ToLatin1_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:539
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUtf8()
func (this *QString) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:541
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUtf8()
func (this *QString) ToUtf8_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:543
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit()
func (this *QString) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNKR7QString11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:545
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit()
func (this *QString) ToLocal8Bit_1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNO7QString11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:555
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString fromLatin1(const char *, int)
func (this *QString) FromLatin1(str string, size int) string {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10fromLatin1EPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromLatin1(str string, size int) string {
	var nilthis *QString
	rv := nilthis.FromLatin1(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:568
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString fromLatin1(const QByteArray &)
func (this *QString) FromLatin1_1(str *QByteArray) string {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10fromLatin1ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromLatin1_1(str *QByteArray) string {
	var nilthis *QString
	rv := nilthis.FromLatin1_1(str)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:560
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString fromUtf8(const char *, int)
func (this *QString) FromUtf8(str string, size int) string {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8fromUtf8EPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUtf8(str string, size int) string {
	var nilthis *QString
	rv := nilthis.FromUtf8(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:570
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString fromUtf8(const QByteArray &)
func (this *QString) FromUtf8_1(str *QByteArray) string {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8fromUtf8ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUtf8_1(str *QByteArray) string {
	var nilthis *QString
	rv := nilthis.FromUtf8_1(str)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:564
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString fromLocal8Bit(const char *, int)
func (this *QString) FromLocal8Bit(str string, size int) string {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString13fromLocal8BitEPKci", qtrt.FFI_TYPE_POINTER, convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromLocal8Bit(str string, size int) string {
	var nilthis *QString
	rv := nilthis.FromLocal8Bit(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:572
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString fromLocal8Bit(const QByteArray &)
func (this *QString) FromLocal8Bit_1(str *QByteArray) string {
	var convArg0 = str.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString13fromLocal8BitERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromLocal8Bit_1(str *QByteArray) string {
	var nilthis *QString
	rv := nilthis.FromLocal8Bit_1(str)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:574
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromUtf16(const ushort *, int)
func (this *QString) FromUtf16(arg0 unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString9fromUtf16EPKti", qtrt.FFI_TYPE_POINTER, &arg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUtf16(arg0 unsafe.Pointer /*666*/, size int) string {
	var nilthis *QString
	rv := nilthis.FromUtf16(arg0, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:579
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString fromUtf16(const char16_t *, int)
func (this *QString) FromUtf16_1(str unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString9fromUtf16EPKDsi", qtrt.FFI_TYPE_POINTER, &str, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUtf16_1(str unsafe.Pointer /*666*/, size int) string {
	var nilthis *QString
	rv := nilthis.FromUtf16_1(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:575
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromUcs4(const uint *, int)
func (this *QString) FromUcs4(arg0 unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8fromUcs4EPKji", qtrt.FFI_TYPE_POINTER, &arg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUcs4(arg0 unsafe.Pointer /*666*/, size int) string {
	var nilthis *QString
	rv := nilthis.FromUcs4(arg0, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:581
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString fromUcs4(const char32_t *, int)
func (this *QString) FromUcs4_1(str unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8fromUcs4EPKDii", qtrt.FFI_TYPE_POINTER, str, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUcs4_1(str unsafe.Pointer /*666*/, size int) string {
	var nilthis *QString
	rv := nilthis.FromUcs4_1(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:576
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromRawData(const QChar *, int)
func (this *QString) FromRawData(arg0 *QChar /*777 const QChar **/, size int) string {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString11fromRawDataEPK5QChari", qtrt.FFI_TYPE_POINTER, convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromRawData(arg0 *QChar /*777 const QChar **/, size int) string {
	var nilthis *QString
	rv := nilthis.FromRawData(arg0, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:594
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int toWCharArray(wchar_t *)
func (this *QString) ToWCharArray(array unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString12toWCharArrayEPw", qtrt.FFI_TYPE_POINTER, this.GetCthis(), array)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:595
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString fromWCharArray(const wchar_t *, int)
func (this *QString) FromWCharArray(string unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString14fromWCharArrayEPKwi", qtrt.FFI_TYPE_POINTER, string, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromWCharArray(string unsafe.Pointer /*666*/, size int) string {
	var nilthis *QString
	rv := nilthis.FromWCharArray(string, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:597
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & setRawData(const QChar *, int)
func (this *QString) SetRawData(unicode *QChar /*777 const QChar **/, size int) string {
	var convArg0 = unicode.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10setRawDataEPK5QChari", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:598
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & setUnicode(const QChar *, int)
func (this *QString) SetUnicode(unicode *QChar /*777 const QChar **/, size int) string {
	var convArg0 = unicode.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10setUnicodeEPK5QChari", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:599
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString & setUtf16(const ushort *, int)
func (this *QString) SetUtf16(utf16 unsafe.Pointer /*666*/, size int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString8setUtf16EPKti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &utf16, size)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:601
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QString &, Qt::CaseSensitivity)
func (this *QString) Compare(s string, cs int) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7compareERKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:602
// index:1
// Public Visibility=Default Availability=Available
// [4] int compare(QLatin1String, Qt::CaseSensitivity)
func (this *QString) Compare_1(other *QLatin1String /*123*/, cs int) int {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7compareE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:604
// index:2
// Public static inline Visibility=Default Availability=Available
// [4] int compare(const QString &, const QString &, Qt::CaseSensitivity)
func (this *QString) Compare_2(s1 string, s2 string, cs int) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7compareERKS_S1_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_Compare_2(s1 string, s2 string, cs int) int {
	var nilthis *QString
	rv := nilthis.Compare_2(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:608
// index:3
// Public static inline Visibility=Default Availability=Available
// [4] int compare(const QString &, QLatin1String, Qt::CaseSensitivity)
func (this *QString) Compare_3(s1 string, s2 *QLatin1String /*123*/, cs int) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7compareERKS_13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_Compare_3(s1 string, s2 *QLatin1String /*123*/, cs int) int {
	var nilthis *QString
	rv := nilthis.Compare_3(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:611
// index:4
// Public static inline Visibility=Default Availability=Available
// [4] int compare(QLatin1String, const QString &, Qt::CaseSensitivity)
func (this *QString) Compare_4(s1 *QLatin1String /*123*/, s2 string, cs int) int {
	var convArg0 = s1.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7compareE13QLatin1StringRKS_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_Compare_4(s1 *QLatin1String /*123*/, s2 string, cs int) int {
	var nilthis *QString
	rv := nilthis.Compare_4(s1, s2, cs)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:615
// index:5
// Public inline Visibility=Default Availability=Available
// [4] int compare(const QStringRef &, Qt::CaseSensitivity)
func (this *QString) Compare_5(s *QStringRef, cs int) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7compareERK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:616
// index:6
// Public static Visibility=Default Availability=Available
// [4] int compare(const QString &, const QStringRef &, Qt::CaseSensitivity)
func (this *QString) Compare_6(s1 string, s2 *QStringRef, arg2 int) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString7compareERKS_RK10QStringRefN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_Compare_6(s1 string, s2 *QStringRef, arg2 int) int {
	var nilthis *QString
	rv := nilthis.Compare_6(s1, s2, arg2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:619
// index:0
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &)
func (this *QString) LocaleAwareCompare(s string) int {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString18localeAwareCompareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:620
// index:1
// Public static inline Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &, const QString &)
func (this *QString) LocaleAwareCompare_1(s1 string, s2 string) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(s2)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString18localeAwareCompareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_LocaleAwareCompare_1(s1 string, s2 string) int {
	var nilthis *QString
	rv := nilthis.LocaleAwareCompare_1(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:623
// index:2
// Public Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QStringRef &)
func (this *QString) LocaleAwareCompare_2(s *QStringRef) int {
	var convArg0 = s.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString18localeAwareCompareERK10QStringRef", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:624
// index:3
// Public static Visibility=Default Availability=Available
// [4] int localeAwareCompare(const QString &, const QStringRef &)
func (this *QString) LocaleAwareCompare_3(s1 string, s2 *QStringRef) int {
	var tmpArg0 = NewQString_5(s1)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = s2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString18localeAwareCompareERKS_RK10QStringRef", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QString_LocaleAwareCompare_3(s1 string, s2 *QStringRef) int {
	var nilthis *QString
	rv := nilthis.LocaleAwareCompare_3(s1, s2)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:627
// index:0
// Public Visibility=Default Availability=Available
// [2] short toShort(_Bool *, int)
func (this *QString) ToShort(ok unsafe.Pointer /*666*/, base int) int16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7toShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int16", rv).(int16) // 1111
}

// /usr/include/qt/QtCore/qstring.h:628
// index:0
// Public Visibility=Default Availability=Available
// [2] ushort toUShort(_Bool *, int)
func (this *QString) ToUShort(ok unsafe.Pointer /*666*/, base int) uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8toUShortEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:629
// index:0
// Public Visibility=Default Availability=Available
// [4] int toInt(_Bool *, int)
func (this *QString) ToInt(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5toIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:630
// index:0
// Public Visibility=Default Availability=Available
// [4] uint toUInt(_Bool *, int)
func (this *QString) ToUInt(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6toUIntEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:631
// index:0
// Public Visibility=Default Availability=Available
// [8] long toLong(_Bool *, int)
func (this *QString) ToLong(ok unsafe.Pointer /*666*/, base int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6toLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:632
// index:0
// Public Visibility=Default Availability=Available
// [8] ulong toULong(_Bool *, int)
func (this *QString) ToULong(ok unsafe.Pointer /*666*/, base int) uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7toULongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:633
// index:0
// Public Visibility=Default Availability=Available
// [8] qlonglong toLongLong(_Bool *, int)
func (this *QString) ToLongLong(ok unsafe.Pointer /*666*/, base int) int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10toLongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:634
// index:0
// Public Visibility=Default Availability=Available
// [8] qulonglong toULongLong(_Bool *, int)
func (this *QString) ToULongLong(ok unsafe.Pointer /*666*/, base int) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11toULongLongEPbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &ok, base)
	gopp.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:635
// index:0
// Public Visibility=Default Availability=Available
// [4] float toFloat(_Bool *)
func (this *QString) ToFloat(ok unsafe.Pointer /*666*/) float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString7toFloatEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qstring.h:636
// index:0
// Public Visibility=Default Availability=Available
// [8] double toDouble(_Bool *)
func (this *QString) ToDouble(ok unsafe.Pointer /*666*/) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8toDoubleEPb", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &ok)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qstring.h:638
// index:0
// Public Visibility=Default Availability=Available
// [8] QString & setNum(short, int)
func (this *QString) SetNum(arg0 int16, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEsi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:639
// index:1
// Public Visibility=Default Availability=Available
// [8] QString & setNum(ushort, int)
func (this *QString) SetNum_1(arg0 uint16, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:640
// index:2
// Public Visibility=Default Availability=Available
// [8] QString & setNum(int, int)
func (this *QString) SetNum_2(arg0 int, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:641
// index:3
// Public Visibility=Default Availability=Available
// [8] QString & setNum(uint, int)
func (this *QString) SetNum_3(arg0 uint, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEji", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:642
// index:4
// Public Visibility=Default Availability=Available
// [8] QString & setNum(long, int)
func (this *QString) SetNum_4(arg0 int, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEli", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:643
// index:5
// Public Visibility=Default Availability=Available
// [8] QString & setNum(ulong, int)
func (this *QString) SetNum_5(arg0 uint, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEmi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:644
// index:6
// Public Visibility=Default Availability=Available
// [8] QString & setNum(qlonglong, int)
func (this *QString) SetNum_6(arg0 int64, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumExi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:645
// index:7
// Public Visibility=Default Availability=Available
// [8] QString & setNum(qulonglong, int)
func (this *QString) SetNum_7(arg0 uint64, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEyi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:646
// index:8
// Public Visibility=Default Availability=Available
// [8] QString & setNum(float, char, int)
func (this *QString) SetNum_8(arg0 float32, f byte, prec int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEfci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:647
// index:9
// Public Visibility=Default Availability=Available
// [8] QString & setNum(double, char, int)
func (this *QString) SetNum_9(arg0 float64, f byte, prec int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6setNumEdci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, f, prec)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:649
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString number(int, int)
func (this *QString) Number(arg0 int, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEii", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number(arg0 int, base int) string {
	var nilthis *QString
	rv := nilthis.Number(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:650
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString number(uint, int)
func (this *QString) Number_1(arg0 uint, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEji", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_1(arg0 uint, base int) string {
	var nilthis *QString
	rv := nilthis.Number_1(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:651
// index:2
// Public static Visibility=Default Availability=Available
// [8] QString number(long, int)
func (this *QString) Number_2(arg0 int, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEli", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_2(arg0 int, base int) string {
	var nilthis *QString
	rv := nilthis.Number_2(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:652
// index:3
// Public static Visibility=Default Availability=Available
// [8] QString number(ulong, int)
func (this *QString) Number_3(arg0 uint, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEmi", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_3(arg0 uint, base int) string {
	var nilthis *QString
	rv := nilthis.Number_3(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:653
// index:4
// Public static Visibility=Default Availability=Available
// [8] QString number(qlonglong, int)
func (this *QString) Number_4(arg0 int64, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberExi", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_4(arg0 int64, base int) string {
	var nilthis *QString
	rv := nilthis.Number_4(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:654
// index:5
// Public static Visibility=Default Availability=Available
// [8] QString number(qulonglong, int)
func (this *QString) Number_5(arg0 uint64, base int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEyi", qtrt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_5(arg0 uint64, base int) string {
	var nilthis *QString
	rv := nilthis.Number_5(arg0, base)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:655
// index:6
// Public static Visibility=Default Availability=Available
// [8] QString number(double, char, int)
func (this *QString) Number_6(arg0 float64, f byte, prec int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString6numberEdci", qtrt.FFI_TYPE_POINTER, arg0, f, prec)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_Number_6(arg0 float64, f byte, prec int) string {
	var nilthis *QString
	rv := nilthis.Number_6(arg0, f, prec)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:750
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::iterator begin()
func (this *QString) Begin() *QChar /*777 QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:751
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator begin()
func (this *QString) Begin_1() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:752
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator cbegin()
func (this *QString) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:753
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator constBegin()
func (this *QString) ConstBegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:754
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::iterator end()
func (this *QString) End() *QChar /*777 QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:755
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator end()
func (this *QString) End_1() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:756
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator cend()
func (this *QString) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:757
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString::const_iterator constEnd()
func (this *QString) ConstEnd() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:773
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(QChar)
func (this *QString) Push_back(c *QChar /*123*/) {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString9push_backE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:774
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QString &)
func (this *QString) Push_back_1(s string) {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString9push_backERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:775
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(QChar)
func (this *QString) Push_front(c *QChar /*123*/) {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10push_frontE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:776
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QString &)
func (this *QString) Push_front_1(s string) {
	var tmpArg0 = NewQString_5(s)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString10push_frontERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:777
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void shrink_to_fit()
func (this *QString) Shrink_to_fit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QString13shrink_to_fitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:780
// index:0
// Public inline Visibility=Default Availability=Available
// [32] std::string toStdString()
func (this *QString) ToStdString() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString11toStdStringB5cxx11Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:782
// index:0
// Public inline Visibility=Default Availability=Available
// [32] std::wstring toStdWString()
func (this *QString) ToStdWString() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString12toStdWStringB5cxx11Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:786
// index:0
// Public inline Visibility=Default Availability=Available
// [32] std::u16string toStdU16String()
func (this *QString) ToStdU16String() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString14toStdU16StringB5cxx11Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:788
// index:0
// Public inline Visibility=Default Availability=Available
// [32] std::u32string toStdU32String()
func (this *QString) ToStdU32String() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString14toStdU32StringB5cxx11Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:805
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QString) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:808
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSimpleText()
func (this *QString) IsSimpleText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString12isSimpleTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:809
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRightToLeft()
func (this *QString) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QString13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

type QString__SectionFlag = int

const QString__SectionDefault QString__SectionFlag = 0
const QString__SectionSkipEmpty QString__SectionFlag = 1
const QString__SectionIncludeLeadingSep QString__SectionFlag = 2
const QString__SectionIncludeTrailingSep QString__SectionFlag = 4
const QString__SectionCaseInsensitiveSeps QString__SectionFlag = 8

type QString__SplitBehavior = int

const QString__KeepEmptyParts QString__SplitBehavior = 0
const QString__SkipEmptyParts QString__SplitBehavior = 1

type QString__NormalizationForm = int

const QString__NormalizationForm_D QString__NormalizationForm = 0
const QString__NormalizationForm_C QString__NormalizationForm = 1
const QString__NormalizationForm_KD QString__NormalizationForm = 2
const QString__NormalizationForm_KC QString__NormalizationForm = 3

//  body block end
