//  header block begin
// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 90
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
type QByteArray struct {
	*qtrt.CObject
}

func (this *QByteArray) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qbytearray.h:170
// index:0
// inline
// void QByteArray()
func NewQByteArray() *QByteArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}
func NewQByteArrayFromPointer(cthis unsafe.Pointer) *QByteArray {
	return &QByteArray{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbytearray.h:171
// index:1
// void QByteArray(const char *, int)
func NewQByteArray_1(arg0 unsafe.Pointer, size int) *QByteArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2EPKci", ffiqt.FFI_TYPE_VOID, cthis, arg0, &size)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:172
// index:2
// void QByteArray(int, char)
func NewQByteArray_2(size int, c byte) *QByteArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2Eic", ffiqt.FFI_TYPE_VOID, cthis, &size, &c)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:173
// index:3
// void QByteArray(int, Qt::Initialization)
func NewQByteArray_3(size int, arg1 int) *QByteArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2EiN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &size, &arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:437
// index:4
// inline
// void QByteArray(struct QByteArrayDataPtr)
func NewQByteArray_4(dd unsafe.Pointer) *QByteArray {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2E17QByteArrayDataPtr", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:175
// index:0
// inline
// void ~QByteArray()
func DeleteQByteArray(*QByteArray) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:185
// index:0
// inline
// void swap(class QByteArray &)
func (this *QByteArray) Swap(other unsafe.Pointer) {
	// 0: (, other QByteArray &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:188
// index:0
// inline
// int size()
func (this *QByteArray) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:189
// index:0
// inline
// bool isEmpty()
func (this *QByteArray) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:190
// index:0
// void resize(int)
func (this *QByteArray) Resize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6resizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:192
// index:0
// QByteArray & fill(char, int)
func (this *QByteArray) Fill(c byte, size int) {
	// 0: (, c char, size int), (&c, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4fillEci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:194
// index:0
// inline
// int capacity()
func (this *QByteArray) Capacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8capacityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:195
// index:0
// inline
// void reserve(int)
func (this *QByteArray) Reserve(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7reserveEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:196
// index:0
// inline
// void squeeze()
func (this *QByteArray) Squeeze() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7squeezeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:202
// index:0
// inline
// char * data()
func (this *QByteArray) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:203
// index:1
// inline
// const char * data()
func (this *QByteArray) Data_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4dataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:204
// index:0
// inline
// const char * constData()
func (this *QByteArray) ConstData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray9constDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:205
// index:0
// inline
// void detach()
func (this *QByteArray) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:206
// index:0
// inline
// bool isDetached()
func (this *QByteArray) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:207
// index:0
// inline
// bool isSharedWith(const class QByteArray &)
func (this *QByteArray) IsSharedWith(other unsafe.Pointer) {
	// 0: (, other const QByteArray &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray12isSharedWithERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:208
// index:0
// void clear()
func (this *QByteArray) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:210
// index:0
// inline
// char at(int)
func (this *QByteArray) At(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray2atEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:215
// index:0
// inline
// char front()
func (this *QByteArray) Front() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5frontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:216
// index:1
// inline
// QByteRef front()
func (this *QByteArray) Front_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5frontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:217
// index:0
// inline
// char back()
func (this *QByteArray) Back() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4backEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:218
// index:1
// inline
// QByteRef back()
func (this *QByteArray) Back_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4backEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:220
// index:0
// int indexOf(char, int)
func (this *QByteArray) IndexOf(c byte, from int) {
	// 0: (, c char, from int), (&c, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:221
// index:1
// int indexOf(const char *, int)
func (this *QByteArray) IndexOf_1(c unsafe.Pointer, from int) {
	// 1: (, c const char *, from int), (c, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:222
// index:2
// int indexOf(const class QByteArray &, int)
func (this *QByteArray) IndexOf_2(a unsafe.Pointer, from int) {
	// 2: (, a const QByteArray &, from int), (a, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERKS_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:331
// index:3
// int indexOf(const class QString &, int)
func (this *QByteArray) IndexOf_3(s unsafe.Pointer, from int) {
	// 3: (, s const QString &, from int), (s, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:223
// index:0
// int lastIndexOf(char, int)
func (this *QByteArray) LastIndexOf(c byte, from int) {
	// 0: (, c char, from int), (&c, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:224
// index:1
// int lastIndexOf(const char *, int)
func (this *QByteArray) LastIndexOf_1(c unsafe.Pointer, from int) {
	// 1: (, c const char *, from int), (c, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:225
// index:2
// int lastIndexOf(const class QByteArray &, int)
func (this *QByteArray) LastIndexOf_2(a unsafe.Pointer, from int) {
	// 2: (, a const QByteArray &, from int), (a, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERKS_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:332
// index:3
// int lastIndexOf(const class QString &, int)
func (this *QByteArray) LastIndexOf_3(s unsafe.Pointer, from int) {
	// 3: (, s const QString &, from int), (s, &from)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), s, &from)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:227
// index:0
// inline
// bool contains(char)
func (this *QByteArray) Contains(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:228
// index:1
// inline
// bool contains(const char *)
func (this *QByteArray) Contains_1(a unsafe.Pointer) {
	// 1: (, a const char *), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:229
// index:2
// inline
// bool contains(const class QByteArray &)
func (this *QByteArray) Contains_2(a unsafe.Pointer) {
	// 2: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:230
// index:0
// int count(char)
func (this *QByteArray) Count(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:231
// index:1
// int count(const char *)
func (this *QByteArray) Count_1(a unsafe.Pointer) {
	// 1: (, a const char *), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:232
// index:2
// int count(const class QByteArray &)
func (this *QByteArray) Count_2(a unsafe.Pointer) {
	// 2: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:433
// index:3
// inline
// int count()
func (this *QByteArray) Count_3() {
	// 3: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:234
// index:0
// QByteArray left(int)
func (this *QByteArray) Left(len int) {
	// 0: (, len int), (&len)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4leftEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:235
// index:0
// QByteArray right(int)
func (this *QByteArray) Right(len int) {
	// 0: (, len int), (&len)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5rightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:0
// QByteArray mid(int, int)
func (this *QByteArray) Mid(index int, len int) {
	// 0: (, index int, len int), (&index, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray3midEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:237
// index:0
// inline
// QByteArray chopped(int)
func (this *QByteArray) Chopped(len int) {
	// 0: (, len int), (&len)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7choppedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:240
// index:0
// bool startsWith(const class QByteArray &)
func (this *QByteArray) StartsWith(a unsafe.Pointer) {
	// 0: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:241
// index:1
// bool startsWith(char)
func (this *QByteArray) StartsWith_1(c byte) {
	// 1: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:242
// index:2
// bool startsWith(const char *)
func (this *QByteArray) StartsWith_2(c unsafe.Pointer) {
	// 2: (, c const char *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:244
// index:0
// bool endsWith(const class QByteArray &)
func (this *QByteArray) EndsWith(a unsafe.Pointer) {
	// 0: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:245
// index:1
// bool endsWith(char)
func (this *QByteArray) EndsWith_1(c byte) {
	// 1: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:246
// index:2
// bool endsWith(const char *)
func (this *QByteArray) EndsWith_2(c unsafe.Pointer) {
	// 2: (, c const char *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:248
// index:0
// void truncate(int)
func (this *QByteArray) Truncate(pos int) {
	// 0: (, pos int), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray8truncateEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:249
// index:0
// void chop(int)
func (this *QByteArray) Chop(n int) {
	// 0: (, n int), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4chopEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:259
// index:0
// inline
// QByteArray toLower()
func (this *QByteArray) ToLower() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7toLowerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:261
// index:1
// inline
// QByteArray toLower()
func (this *QByteArray) ToLower_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7toLowerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:263
// index:0
// inline
// QByteArray toUpper()
func (this *QByteArray) ToUpper() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7toUpperEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:265
// index:1
// inline
// QByteArray toUpper()
func (this *QByteArray) ToUpper_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7toUpperEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:267
// index:0
// inline
// QByteArray trimmed()
func (this *QByteArray) Trimmed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7trimmedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:269
// index:1
// inline
// QByteArray trimmed()
func (this *QByteArray) Trimmed_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7trimmedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:271
// index:0
// inline
// QByteArray simplified()
func (this *QByteArray) Simplified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray10simplifiedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:273
// index:1
// inline
// QByteArray simplified()
func (this *QByteArray) Simplified_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray10simplifiedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:285
// index:0
// QByteArray leftJustified(int, char, _Bool)
func (this *QByteArray) LeftJustified(width int, fill byte, truncate bool) {
	// 0: (, width int, fill char, truncate bool), (&width, &fill, &truncate)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width, &fill, &truncate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:286
// index:0
// QByteArray rightJustified(int, char, _Bool)
func (this *QByteArray) RightJustified(width int, fill byte, truncate bool) {
	// 0: (, width int, fill char, truncate bool), (&width, &fill, &truncate)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width, &fill, &truncate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:303
// index:0
// QByteArray & remove(int, int)
func (this *QByteArray) Remove(index int, len int) {
	// 0: (, index int, len int), (&index, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6removeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:304
// index:0
// QByteArray & replace(int, int, const char *)
func (this *QByteArray) Replace(index int, len int, s unsafe.Pointer) {
	// 0: (, index int, len int, s const char *), (&index, &len, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &len, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:305
// index:1
// QByteArray & replace(int, int, const char *, int)
func (this *QByteArray) Replace_1(index int, len int, s unsafe.Pointer, alen int) {
	// 1: (, index int, len int, s const char *, alen int), (&index, &len, s, &alen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &len, s, &alen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:306
// index:2
// QByteArray & replace(int, int, const class QByteArray &)
func (this *QByteArray) Replace_2(index int, len int, s unsafe.Pointer) {
	// 2: (, index int, len int, s const QByteArray &), (&index, &len, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiRKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &len, s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:307
// index:3
// inline
// QByteArray & replace(char, const char *)
func (this *QByteArray) Replace_3(before byte, after unsafe.Pointer) {
	// 3: (, before char, after const char *), (&before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:308
// index:4
// QByteArray & replace(char, const class QByteArray &)
func (this *QByteArray) Replace_4(before byte, after unsafe.Pointer) {
	// 4: (, before char, after const QByteArray &), (&before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:309
// index:5
// inline
// QByteArray & replace(const char *, const char *)
func (this *QByteArray) Replace_5(before unsafe.Pointer, after unsafe.Pointer) {
	// 5: (, before const char *, after const char *), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:310
// index:6
// QByteArray & replace(const char *, int, const char *, int)
func (this *QByteArray) Replace_6(before unsafe.Pointer, bsize int, after unsafe.Pointer, asize int) {
	// 6: (, before const char *, bsize int, after const char *, asize int), (before, &bsize, after, &asize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKciS1_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, &bsize, after, &asize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:311
// index:7
// QByteArray & replace(const class QByteArray &, const class QByteArray &)
func (this *QByteArray) Replace_7(before unsafe.Pointer, after unsafe.Pointer) {
	// 7: (, before const QByteArray &, after const QByteArray &), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_S1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:312
// index:8
// inline
// QByteArray & replace(const class QByteArray &, const char *)
func (this *QByteArray) Replace_8(before unsafe.Pointer, after unsafe.Pointer) {
	// 8: (, before const QByteArray &, after const char *), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_PKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:313
// index:9
// QByteArray & replace(const char *, const class QByteArray &)
func (this *QByteArray) Replace_9(before unsafe.Pointer, after unsafe.Pointer) {
	// 9: (, before const char *, after const QByteArray &), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcRKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:314
// index:10
// QByteArray & replace(char, char)
func (this *QByteArray) Replace_10(before byte, after byte) {
	// 10: (, before char, after char), (&before, &after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &before, &after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:326
// index:11
// QByteArray & replace(const class QString &, const char *)
func (this *QByteArray) Replace_11(before unsafe.Pointer, after unsafe.Pointer) {
	// 11: (, before const QString &, after const char *), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:327
// index:12
// QByteArray & replace(char, const class QString &)
func (this *QByteArray) Replace_12(c byte, after unsafe.Pointer) {
	// 12: (, c char, after const QString &), (&c, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:328
// index:13
// QByteArray & replace(const class QString &, const class QByteArray &)
func (this *QByteArray) Replace_13(before unsafe.Pointer, after unsafe.Pointer) {
	// 13: (, before const QString &, after const QByteArray &), (before, after)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringRKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), before, after)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:319
// index:0
// QList<QByteArray> split(char)
func (this *QByteArray) Split(sep byte) {
	// 0: (, sep char), (&sep)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5splitEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sep)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:321
// index:0
// QByteArray repeated(int)
func (this *QByteArray) Repeated(times int) {
	// 0: (, times int), (&times)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8repeatedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &times)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// short toShort(_Bool *, int)
func (this *QByteArray) ToShort(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// ushort toUShort(_Bool *, int)
func (this *QByteArray) ToUShort(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// int toInt(_Bool *, int)
func (this *QByteArray) ToInt(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// uint toUInt(_Bool *, int)
func (this *QByteArray) ToUInt(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// long toLong(_Bool *, int)
func (this *QByteArray) ToLong(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:348
// index:0
// ulong toULong(_Bool *, int)
func (this *QByteArray) ToULong(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:349
// index:0
// qlonglong toLongLong(_Bool *, int)
func (this *QByteArray) ToLongLong(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// qulonglong toULongLong(_Bool *, int)
func (this *QByteArray) ToULongLong(ok unsafe.Pointer, base int) {
	// 0: (, ok bool *, base int), (ok, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// float toFloat(_Bool *)
func (this *QByteArray) ToFloat(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// double toDouble(_Bool *)
func (this *QByteArray) ToDouble(ok unsafe.Pointer) {
	// 0: (, ok bool *), (ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:353
// index:0
// QByteArray toBase64(QByteArray::Base64Options)
func (this *QByteArray) ToBase64(options int) {
	// 0: (, QFlags<QByteArray::Base64Option> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toBase64E6QFlagsINS_12Base64OptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:354
// index:1
// QByteArray toBase64()
func (this *QByteArray) ToBase64_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toBase64Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// QByteArray toHex()
func (this *QByteArray) ToHex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toHexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:1
// QByteArray toHex(char)
func (this *QByteArray) ToHex_1(separator byte) {
	// 1: (, separator char), (&separator)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toHexEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &separator)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// QByteArray toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
func (this *QByteArray) ToPercentEncoding(exclude unsafe.Pointer, include unsafe.Pointer, percent byte) {
	// 0: (, exclude const QByteArray &, include const QByteArray &, percent char), (exclude, include, &percent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", ffiqt.FFI_TYPE_VOID, this.GetCthis(), exclude, include, &percent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:361
// index:0
// inline
// QByteArray & setNum(short, int)
func (this *QByteArray) SetNum(arg0 int16, base int) {
	// 0: (, short arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEsi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:362
// index:1
// inline
// QByteArray & setNum(ushort, int)
func (this *QByteArray) SetNum_1(arg0 uint16, base int) {
	// 1: (, ushort arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:363
// index:2
// inline
// QByteArray & setNum(int, int)
func (this *QByteArray) SetNum_2(arg0 int, base int) {
	// 2: (, int arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:364
// index:3
// inline
// QByteArray & setNum(uint, int)
func (this *QByteArray) SetNum_3(arg0 uint, base int) {
	// 3: (, uint arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:365
// index:4
// QByteArray & setNum(qlonglong, int)
func (this *QByteArray) SetNum_4(arg0 int64, base int) {
	// 4: (, qlonglong arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:366
// index:5
// QByteArray & setNum(qulonglong, int)
func (this *QByteArray) SetNum_5(arg0 uint64, base int) {
	// 5: (, qulonglong arg0, base int), (&arg0, &base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:367
// index:6
// inline
// QByteArray & setNum(float, char, int)
func (this *QByteArray) SetNum_6(arg0 float32, f byte, prec int) {
	// 6: (, float arg0, f char, prec int), (&arg0, &f, &prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:368
// index:7
// QByteArray & setNum(double, char, int)
func (this *QByteArray) SetNum_7(arg0 float64, f byte, prec int) {
	// 7: (, double arg0, f char, prec int), (&arg0, &f, &prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:369
// index:0
// QByteArray & setRawData(const char *, uint)
func (this *QByteArray) SetRawData(a unsafe.Pointer, n uint) {
	// 0: (, a const char *, n uint), (a, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10setRawDataEPKcj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:371
// index:0
// static
// QByteArray number(int, int)
func (this *QByteArray) Number(arg0 int, base int) {
	// 0: (int arg0, base int), (arg0, base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_Number(arg0 int, base int) {
	// 0: (int arg0, base int), (arg0, base)
	var nilthis *QByteArray
	nilthis.Number(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:372
// index:1
// static
// QByteArray number(uint, int)
func (this *QByteArray) Number_1(arg0 uint, base int) {
	// 1: (uint arg0, base int), (arg0, base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEji", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_Number_1(arg0 uint, base int) {
	// 1: (uint arg0, base int), (arg0, base)
	var nilthis *QByteArray
	nilthis.Number_1(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:373
// index:2
// static
// QByteArray number(qlonglong, int)
func (this *QByteArray) Number_2(arg0 int64, base int) {
	// 2: (qlonglong arg0, base int), (arg0, base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberExi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_Number_2(arg0 int64, base int) {
	// 2: (qlonglong arg0, base int), (arg0, base)
	var nilthis *QByteArray
	nilthis.Number_2(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:374
// index:3
// static
// QByteArray number(qulonglong, int)
func (this *QByteArray) Number_3(arg0 uint64, base int) {
	// 3: (qulonglong arg0, base int), (arg0, base)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_Number_3(arg0 uint64, base int) {
	// 3: (qulonglong arg0, base int), (arg0, base)
	var nilthis *QByteArray
	nilthis.Number_3(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:375
// index:4
// static
// QByteArray number(double, char, int)
func (this *QByteArray) Number_4(arg0 float64, f byte, prec int) {
	// 4: (double arg0, f char, prec int), (arg0, f, prec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_Number_4(arg0 float64, f byte, prec int) {
	// 4: (double arg0, f char, prec int), (arg0, f, prec)
	var nilthis *QByteArray
	nilthis.Number_4(arg0, f, prec)
}

// /usr/include/qt/QtCore/qbytearray.h:376
// index:0
// static
// QByteArray fromRawData(const char *, int)
func (this *QByteArray) FromRawData(arg0 unsafe.Pointer, size int) {
	// 0: (const char * arg0, size int), (arg0, size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray11fromRawDataEPKci", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_FromRawData(arg0 unsafe.Pointer, size int) {
	// 0: (const char * arg0, size int), (arg0, size)
	var nilthis *QByteArray
	nilthis.FromRawData(arg0, size)
}

// /usr/include/qt/QtCore/qbytearray.h:377
// index:0
// static
// QByteArray fromBase64(const class QByteArray &, QByteArray::Base64Options)
func (this *QByteArray) FromBase64(base64 unsafe.Pointer, options int) {
	// 0: (base64 const QByteArray &, QFlags<QByteArray::Base64Option> options), (base64, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10fromBase64ERKS_6QFlagsINS_12Base64OptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_FromBase64(base64 unsafe.Pointer, options int) {
	// 0: (base64 const QByteArray &, QFlags<QByteArray::Base64Option> options), (base64, options)
	var nilthis *QByteArray
	nilthis.FromBase64(base64, options)
}

// /usr/include/qt/QtCore/qbytearray.h:378
// index:1
// static
// QByteArray fromBase64(const class QByteArray &)
func (this *QByteArray) FromBase64_1(base64 unsafe.Pointer) {
	// 1: (base64 const QByteArray &), (base64)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10fromBase64ERKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_FromBase64_1(base64 unsafe.Pointer) {
	// 1: (base64 const QByteArray &), (base64)
	var nilthis *QByteArray
	nilthis.FromBase64_1(base64)
}

// /usr/include/qt/QtCore/qbytearray.h:379
// index:0
// static
// QByteArray fromHex(const class QByteArray &)
func (this *QByteArray) FromHex(hexEncoded unsafe.Pointer) {
	// 0: (hexEncoded const QByteArray &), (hexEncoded)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7fromHexERKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_FromHex(hexEncoded unsafe.Pointer) {
	// 0: (hexEncoded const QByteArray &), (hexEncoded)
	var nilthis *QByteArray
	nilthis.FromHex(hexEncoded)
}

// /usr/include/qt/QtCore/qbytearray.h:380
// index:0
// static
// QByteArray fromPercentEncoding(const class QByteArray &, char)
func (this *QByteArray) FromPercentEncoding(pctEncoded unsafe.Pointer, percent byte) {
	// 0: (pctEncoded const QByteArray &, percent char), (pctEncoded, percent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray19fromPercentEncodingERKS_c", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QByteArray_FromPercentEncoding(pctEncoded unsafe.Pointer, percent byte) {
	// 0: (pctEncoded const QByteArray &, percent char), (pctEncoded, percent)
	var nilthis *QByteArray
	nilthis.FromPercentEncoding(pctEncoded, percent)
}

// /usr/include/qt/QtCore/qbytearray.h:399
// index:0
// inline
// QByteArray::iterator begin()
func (this *QByteArray) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:400
// index:1
// inline
// QByteArray::const_iterator begin()
func (this *QByteArray) Begin_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:401
// index:0
// inline
// QByteArray::const_iterator cbegin()
func (this *QByteArray) Cbegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6cbeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:402
// index:0
// inline
// QByteArray::const_iterator constBegin()
func (this *QByteArray) ConstBegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10constBeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:403
// index:0
// inline
// QByteArray::iterator end()
func (this *QByteArray) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:404
// index:1
// inline
// QByteArray::const_iterator end()
func (this *QByteArray) End_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:405
// index:0
// inline
// QByteArray::const_iterator cend()
func (this *QByteArray) Cend() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4cendEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:406
// index:0
// inline
// QByteArray::const_iterator constEnd()
func (this *QByteArray) ConstEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8constEndEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:422
// index:0
// inline
// void push_back(char)
func (this *QByteArray) Push_back(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:423
// index:1
// inline
// void push_back(const char *)
func (this *QByteArray) Push_back_1(c unsafe.Pointer) {
	// 1: (, c const char *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:424
// index:2
// inline
// void push_back(const class QByteArray &)
func (this *QByteArray) Push_back_2(a unsafe.Pointer) {
	// 2: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:425
// index:0
// inline
// void push_front(char)
func (this *QByteArray) Push_front(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:426
// index:1
// inline
// void push_front(const char *)
func (this *QByteArray) Push_front_1(c unsafe.Pointer) {
	// 1: (, c const char *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:427
// index:2
// inline
// void push_front(const class QByteArray &)
func (this *QByteArray) Push_front_2(a unsafe.Pointer) {
	// 2: (, a const QByteArray &), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:428
// index:0
// inline
// void shrink_to_fit()
func (this *QByteArray) Shrink_to_fit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray13shrink_to_fitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:431
// index:0
// inline
// std::string toStdString()
func (this *QByteArray) ToStdString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11toStdStringB5cxx11Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:434
// index:0
// inline
// int length()
func (this *QByteArray) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:435
// index:0
// bool isNull()
func (this *QByteArray) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
