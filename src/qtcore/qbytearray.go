//  header block begin
// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 89
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
func NewQByteArrayFromPointer(cthis unsafe.Pointer) *QByteArray {
	return &QByteArray{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbytearray.h:170
// index:0
// Public inline
// void QByteArray()
func NewQByteArray() *QByteArray {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:171
// index:1
// Public
// void QByteArray(const char *, int)
func NewQByteArray_1(arg0 string, size int) *QByteArray {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2EPKci", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &size)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:172
// index:2
// Public
// void QByteArray(int, char)
func NewQByteArray_2(size int, c byte) *QByteArray {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2Eic", ffiqt.FFI_TYPE_VOID, cthis, &size, &c)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:173
// index:3
// Public
// void QByteArray(int, Qt::Initialization)
func NewQByteArray_3(size int, arg1 int) *QByteArray {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayC2EiN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &size, &arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQByteArrayFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:175
// index:0
// Public inline
// void ~QByteArray()
func DeleteQByteArray(*QByteArray) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArrayD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:185
// index:0
// Public inline
// void swap(class QByteArray &)
func (this *QByteArray) Swap(other *QByteArray) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:188
// index:0
// Public inline
// int size()
func (this *QByteArray) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:189
// index:0
// Public inline
// bool isEmpty()
func (this *QByteArray) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:190
// index:0
// Public
// void resize(int)
func (this *QByteArray) Resize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6resizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:192
// index:0
// Public
// QByteArray & fill(char, int)
func (this *QByteArray) Fill(c byte, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4fillEci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, &size)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:194
// index:0
// Public inline
// int capacity()
func (this *QByteArray) Capacity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8capacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:195
// index:0
// Public inline
// void reserve(int)
func (this *QByteArray) Reserve(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7reserveEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:196
// index:0
// Public inline
// void squeeze()
func (this *QByteArray) Squeeze() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7squeezeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:202
// index:0
// Public inline
// char * data()
func (this *QByteArray) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:203
// index:1
// Public inline
// const char * data()
func (this *QByteArray) Data_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:204
// index:0
// Public inline
// const char * constData()
func (this *QByteArray) ConstData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray9constDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:205
// index:0
// Public inline
// void detach()
func (this *QByteArray) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:206
// index:0
// Public inline
// bool isDetached()
func (this *QByteArray) IsDetached() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:207
// index:0
// Public inline
// bool isSharedWith(const class QByteArray &)
func (this *QByteArray) IsSharedWith(other *QByteArray) interface{} {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray12isSharedWithERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:208
// index:0
// Public
// void clear()
func (this *QByteArray) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:210
// index:0
// Public inline
// char at(int)
func (this *QByteArray) At(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:215
// index:0
// Public inline
// char front()
func (this *QByteArray) Front() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:216
// index:1
// Public inline
// QByteRef front()
func (this *QByteArray) Front_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:217
// index:0
// Public inline
// char back()
func (this *QByteArray) Back() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:218
// index:1
// Public inline
// QByteRef back()
func (this *QByteArray) Back_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:220
// index:0
// Public
// int indexOf(char, int)
func (this *QByteArray) IndexOf(c byte, from int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:221
// index:1
// Public
// int indexOf(const char *, int)
func (this *QByteArray) IndexOf_1(c string, from int) interface{} {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:222
// index:2
// Public
// int indexOf(const class QByteArray &, int)
func (this *QByteArray) IndexOf_2(a *QByteArray, from int) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERKS_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:331
// index:3
// Public
// int indexOf(const class QString &, int)
func (this *QByteArray) IndexOf_3(s *QString, from int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7indexOfERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:223
// index:0
// Public
// int lastIndexOf(char, int)
func (this *QByteArray) LastIndexOf(c byte, from int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:224
// index:1
// Public
// int lastIndexOf(const char *, int)
func (this *QByteArray) LastIndexOf_1(c string, from int) interface{} {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:225
// index:2
// Public
// int lastIndexOf(const class QByteArray &, int)
func (this *QByteArray) LastIndexOf_2(a *QByteArray, from int) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERKS_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:332
// index:3
// Public
// int lastIndexOf(const class QString &, int)
func (this *QByteArray) LastIndexOf_3(s *QString, from int) interface{} {
	var convArg0 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11lastIndexOfERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &from)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:227
// index:0
// Public inline
// bool contains(char)
func (this *QByteArray) Contains(c byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:228
// index:1
// Public inline
// bool contains(const char *)
func (this *QByteArray) Contains_1(a string) interface{} {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:229
// index:2
// Public inline
// bool contains(const class QByteArray &)
func (this *QByteArray) Contains_2(a *QByteArray) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8containsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:230
// index:0
// Public
// int count(char)
func (this *QByteArray) Count(c byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:231
// index:1
// Public
// int count(const char *)
func (this *QByteArray) Count_1(a string) interface{} {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:232
// index:2
// Public
// int count(const class QByteArray &)
func (this *QByteArray) Count_2(a *QByteArray) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:433
// index:3
// Public inline
// int count()
func (this *QByteArray) Count_3() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:234
// index:0
// Public
// QByteArray left(int)
func (this *QByteArray) Left(len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4leftEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:235
// index:0
// Public
// QByteArray right(int)
func (this *QByteArray) Right(len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5rightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:236
// index:0
// Public
// QByteArray mid(int, int)
func (this *QByteArray) Mid(index int, len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray3midEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:237
// index:0
// Public inline
// QByteArray chopped(int)
func (this *QByteArray) Chopped(len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7choppedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:240
// index:0
// Public
// bool startsWith(const class QByteArray &)
func (this *QByteArray) StartsWith(a *QByteArray) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:241
// index:1
// Public
// bool startsWith(char)
func (this *QByteArray) StartsWith_1(c byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:242
// index:2
// Public
// bool startsWith(const char *)
func (this *QByteArray) StartsWith_2(c string) interface{} {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10startsWithEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:244
// index:0
// Public
// bool endsWith(const class QByteArray &)
func (this *QByteArray) EndsWith(a *QByteArray) interface{} {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:245
// index:1
// Public
// bool endsWith(char)
func (this *QByteArray) EndsWith_1(c byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:246
// index:2
// Public
// bool endsWith(const char *)
func (this *QByteArray) EndsWith_2(c string) interface{} {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8endsWithEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:248
// index:0
// Public
// void truncate(int)
func (this *QByteArray) Truncate(pos int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray8truncateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:249
// index:0
// Public
// void chop(int)
func (this *QByteArray) Chop(n int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray4chopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:259
// index:0
// Public inline
// QByteArray toLower()
func (this *QByteArray) ToLower() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:261
// index:1
// Public inline
// QByteArray toLower()
func (this *QByteArray) ToLower_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:263
// index:0
// Public inline
// QByteArray toUpper()
func (this *QByteArray) ToUpper() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:265
// index:1
// Public inline
// QByteArray toUpper()
func (this *QByteArray) ToUpper_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:267
// index:0
// Public inline
// QByteArray trimmed()
func (this *QByteArray) Trimmed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray7trimmedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:269
// index:1
// Public inline
// QByteArray trimmed()
func (this *QByteArray) Trimmed_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray7trimmedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:271
// index:0
// Public inline
// QByteArray simplified()
func (this *QByteArray) Simplified() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNKR10QByteArray10simplifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:273
// index:1
// Public inline
// QByteArray simplified()
func (this *QByteArray) Simplified_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNO10QByteArray10simplifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:285
// index:0
// Public
// QByteArray leftJustified(int, char, _Bool)
func (this *QByteArray) LeftJustified(width int, fill byte, truncate bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray13leftJustifiedEicb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width, &fill, &truncate)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:286
// index:0
// Public
// QByteArray rightJustified(int, char, _Bool)
func (this *QByteArray) RightJustified(width int, fill byte, truncate bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray14rightJustifiedEicb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width, &fill, &truncate)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:303
// index:0
// Public
// QByteArray & remove(int, int)
func (this *QByteArray) Remove(index int, len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6removeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:304
// index:0
// Public
// QByteArray & replace(int, int, const char *)
func (this *QByteArray) Replace(index int, len int, s string) interface{} {
	var convArg2 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &len, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:305
// index:1
// Public
// QByteArray & replace(int, int, const char *, int)
func (this *QByteArray) Replace_1(index int, len int, s string, alen int) interface{} {
	var convArg2 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &len, convArg2, &alen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:306
// index:2
// Public
// QByteArray & replace(int, int, const class QByteArray &)
func (this *QByteArray) Replace_2(index int, len int, s *QByteArray) interface{} {
	var convArg2 = s.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEiiRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &len, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:307
// index:3
// Public inline
// QByteArray & replace(char, const char *)
func (this *QByteArray) Replace_3(before byte, after string) interface{} {
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &before, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:308
// index:4
// Public
// QByteArray & replace(char, const class QByteArray &)
func (this *QByteArray) Replace_4(before byte, after *QByteArray) interface{} {
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &before, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:309
// index:5
// Public inline
// QByteArray & replace(const char *, const char *)
func (this *QByteArray) Replace_5(before string, after string) interface{} {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:310
// index:6
// Public
// QByteArray & replace(const char *, int, const char *, int)
func (this *QByteArray) Replace_6(before string, bsize int, after string, asize int) interface{} {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKciS1_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &bsize, convArg2, &asize)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:311
// index:7
// Public
// QByteArray & replace(const class QByteArray &, const class QByteArray &)
func (this *QByteArray) Replace_7(before *QByteArray, after *QByteArray) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:312
// index:8
// Public inline
// QByteArray & replace(const class QByteArray &, const char *)
func (this *QByteArray) Replace_8(before *QByteArray, after string) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERKS_PKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:313
// index:9
// Public
// QByteArray & replace(const char *, const class QByteArray &)
func (this *QByteArray) Replace_9(before string, after *QByteArray) interface{} {
	var convArg0 = qtrt.CString(before)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEPKcRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:314
// index:10
// Public
// QByteArray & replace(char, char)
func (this *QByteArray) Replace_10(before byte, after byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &before, &after)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:326
// index:11
// Public
// QByteArray & replace(const class QString &, const char *)
func (this *QByteArray) Replace_11(before *QString, after string) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = qtrt.CString(after)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:327
// index:12
// Public
// QByteArray & replace(char, const class QString &)
func (this *QByteArray) Replace_12(c byte, after *QString) interface{} {
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceEcRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:328
// index:13
// Public
// QByteArray & replace(const class QString &, const class QByteArray &)
func (this *QByteArray) Replace_13(before *QString, after *QByteArray) interface{} {
	var convArg0 = before.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7replaceERK7QStringRKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:319
// index:0
// Public
// QList<QByteArray> split(char)
func (this *QByteArray) Split(sep byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5splitEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sep)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:321
// index:0
// Public
// QByteArray repeated(int)
func (this *QByteArray) Repeated(times int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8repeatedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &times)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:343
// index:0
// Public
// short toShort(_Bool *, int)
func (this *QByteArray) ToShort(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:344
// index:0
// Public
// ushort toUShort(_Bool *, int)
func (this *QByteArray) ToUShort(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toUShortEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:345
// index:0
// Public
// int toInt(_Bool *, int)
func (this *QByteArray) ToInt(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:346
// index:0
// Public
// uint toUInt(_Bool *, int)
func (this *QByteArray) ToUInt(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6toUIntEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:347
// index:0
// Public
// long toLong(_Bool *, int)
func (this *QByteArray) ToLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6toLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:348
// index:0
// Public
// ulong toULong(_Bool *, int)
func (this *QByteArray) ToULong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toULongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:349
// index:0
// Public
// qlonglong toLongLong(_Bool *, int)
func (this *QByteArray) ToLongLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10toLongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:350
// index:0
// Public
// qulonglong toULongLong(_Bool *, int)
func (this *QByteArray) ToULongLong(ok unsafe.Pointer, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11toULongLongEPbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:351
// index:0
// Public
// float toFloat(_Bool *)
func (this *QByteArray) ToFloat(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray7toFloatEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:352
// index:0
// Public
// double toDouble(_Bool *)
func (this *QByteArray) ToDouble(ok unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toDoubleEPb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:354
// index:0
// Public
// QByteArray toBase64()
func (this *QByteArray) ToBase64() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8toBase64Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:355
// index:0
// Public
// QByteArray toHex()
func (this *QByteArray) ToHex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toHexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:356
// index:1
// Public
// QByteArray toHex(char)
func (this *QByteArray) ToHex_1(separator byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5toHexEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &separator)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:357
// index:0
// Public
// QByteArray toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
func (this *QByteArray) ToPercentEncoding(exclude *QByteArray, include *QByteArray, percent byte) interface{} {
	var convArg0 = exclude.GetCthis()
	var convArg1 = include.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray17toPercentEncodingERKS_S1_c", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &percent)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:361
// index:0
// Public inline
// QByteArray & setNum(short, int)
func (this *QByteArray) SetNum(arg0 int16, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEsi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:362
// index:1
// Public inline
// QByteArray & setNum(ushort, int)
func (this *QByteArray) SetNum_1(arg0 uint16, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:363
// index:2
// Public inline
// QByteArray & setNum(int, int)
func (this *QByteArray) SetNum_2(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:364
// index:3
// Public inline
// QByteArray & setNum(uint, int)
func (this *QByteArray) SetNum_3(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEji", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:365
// index:4
// Public
// QByteArray & setNum(qlonglong, int)
func (this *QByteArray) SetNum_4(arg0 int64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumExi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:366
// index:5
// Public
// QByteArray & setNum(qulonglong, int)
func (this *QByteArray) SetNum_5(arg0 uint64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEyi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &base)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:367
// index:6
// Public inline
// QByteArray & setNum(float, char, int)
func (this *QByteArray) SetNum_6(arg0 float32, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEfci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:368
// index:7
// Public
// QByteArray & setNum(double, char, int)
func (this *QByteArray) SetNum_7(arg0 float64, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6setNumEdci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &f, &prec)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:369
// index:0
// Public
// QByteArray & setRawData(const char *, uint)
func (this *QByteArray) SetRawData(a string, n uint) interface{} {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10setRawDataEPKcj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:371
// index:0
// Public static
// QByteArray number(int, int)
func (this *QByteArray) Number(arg0 int, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEii", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_Number(arg0 int, base int) {
	var nilthis *QByteArray
	nilthis.Number(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:372
// index:1
// Public static
// QByteArray number(uint, int)
func (this *QByteArray) Number_1(arg0 uint, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEji", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_Number_1(arg0 uint, base int) {
	var nilthis *QByteArray
	nilthis.Number_1(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:373
// index:2
// Public static
// QByteArray number(qlonglong, int)
func (this *QByteArray) Number_2(arg0 int64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberExi", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_Number_2(arg0 int64, base int) {
	var nilthis *QByteArray
	nilthis.Number_2(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:374
// index:3
// Public static
// QByteArray number(qulonglong, int)
func (this *QByteArray) Number_3(arg0 uint64, base int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEyi", ffiqt.FFI_TYPE_POINTER, arg0, base)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_Number_3(arg0 uint64, base int) {
	var nilthis *QByteArray
	nilthis.Number_3(arg0, base)
}

// /usr/include/qt/QtCore/qbytearray.h:375
// index:4
// Public static
// QByteArray number(double, char, int)
func (this *QByteArray) Number_4(arg0 float64, f byte, prec int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray6numberEdci", ffiqt.FFI_TYPE_POINTER, arg0, f, prec)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_Number_4(arg0 float64, f byte, prec int) {
	var nilthis *QByteArray
	nilthis.Number_4(arg0, f, prec)
}

// /usr/include/qt/QtCore/qbytearray.h:376
// index:0
// Public static
// QByteArray fromRawData(const char *, int)
func (this *QByteArray) FromRawData(arg0 string, size int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray11fromRawDataEPKci", ffiqt.FFI_TYPE_POINTER, arg0, size)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_FromRawData(arg0 string, size int) {
	var nilthis *QByteArray
	nilthis.FromRawData(arg0, size)
}

// /usr/include/qt/QtCore/qbytearray.h:378
// index:0
// Public static
// QByteArray fromBase64(const class QByteArray &)
func (this *QByteArray) FromBase64(base64 *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10fromBase64ERKS_", ffiqt.FFI_TYPE_POINTER, base64)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_FromBase64(base64 *QByteArray) {
	var nilthis *QByteArray
	nilthis.FromBase64(base64)
}

// /usr/include/qt/QtCore/qbytearray.h:379
// index:0
// Public static
// QByteArray fromHex(const class QByteArray &)
func (this *QByteArray) FromHex(hexEncoded *QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray7fromHexERKS_", ffiqt.FFI_TYPE_POINTER, hexEncoded)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_FromHex(hexEncoded *QByteArray) {
	var nilthis *QByteArray
	nilthis.FromHex(hexEncoded)
}

// /usr/include/qt/QtCore/qbytearray.h:380
// index:0
// Public static
// QByteArray fromPercentEncoding(const class QByteArray &, char)
func (this *QByteArray) FromPercentEncoding(pctEncoded *QByteArray, percent byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray19fromPercentEncodingERKS_c", ffiqt.FFI_TYPE_POINTER, pctEncoded, percent)
	gopp.ErrPrint(err, rv)
	return rv
}
func QByteArray_FromPercentEncoding(pctEncoded *QByteArray, percent byte) {
	var nilthis *QByteArray
	nilthis.FromPercentEncoding(pctEncoded, percent)
}

// /usr/include/qt/QtCore/qbytearray.h:399
// index:0
// Public inline
// QByteArray::iterator begin()
func (this *QByteArray) Begin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:400
// index:1
// Public inline
// QByteArray::const_iterator begin()
func (this *QByteArray) Begin_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:401
// index:0
// Public inline
// QByteArray::const_iterator cbegin()
func (this *QByteArray) Cbegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6cbeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:402
// index:0
// Public inline
// QByteArray::const_iterator constBegin()
func (this *QByteArray) ConstBegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:403
// index:0
// Public inline
// QByteArray::iterator end()
func (this *QByteArray) End() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:404
// index:1
// Public inline
// QByteArray::const_iterator end()
func (this *QByteArray) End_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:405
// index:0
// Public inline
// QByteArray::const_iterator cend()
func (this *QByteArray) Cend() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray4cendEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:406
// index:0
// Public inline
// QByteArray::const_iterator constEnd()
func (this *QByteArray) ConstEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:422
// index:0
// Public inline
// void push_back(char)
func (this *QByteArray) Push_back(c byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:423
// index:1
// Public inline
// void push_back(const char *)
func (this *QByteArray) Push_back_1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:424
// index:2
// Public inline
// void push_back(const class QByteArray &)
func (this *QByteArray) Push_back_2(a *QByteArray) {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray9push_backERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:425
// index:0
// Public inline
// void push_front(char)
func (this *QByteArray) Push_front(c byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:426
// index:1
// Public inline
// void push_front(const char *)
func (this *QByteArray) Push_front_1(c string) {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:427
// index:2
// Public inline
// void push_front(const class QByteArray &)
func (this *QByteArray) Push_front_2(a *QByteArray) {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray10push_frontERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:428
// index:0
// Public inline
// void shrink_to_fit()
func (this *QByteArray) Shrink_to_fit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QByteArray13shrink_to_fitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:431
// index:0
// Public inline
// std::string toStdString()
func (this *QByteArray) ToStdString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray11toStdStringB5cxx11Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:434
// index:0
// Public inline
// int length()
func (this *QByteArray) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qbytearray.h:435
// index:0
// Public
// bool isNull()
func (this *QByteArray) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QByteArray6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
