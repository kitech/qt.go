package qtcore

// /usr/include/qt/QtCore/qjsonarray.h
// #include <qjsonarray.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QJsonArray struct {
	*qtrt.CObject
}

func (this *QJsonArray) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonArray) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonArrayFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return &QJsonArray{&qtrt.CObject{cthis}}
}
func (*QJsonArray) NewFromPointer(cthis unsafe.Pointer) *QJsonArray {
	return NewQJsonArrayFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonarray.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonArray()
func NewQJsonArray() *QJsonArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArrayC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonArray)
	return gothis
}

// /usr/include/qt/QtCore/qjsonarray.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonArray()
func DeleteQJsonArray(this *QJsonArray) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArrayD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonarray.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [16] QJsonArray fromStringList(const QStringList &)
func (this *QJsonArray) FromStringList(list *QStringList) *QJsonArray /*123*/ {
	var convArg0 = list.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray14fromStringListERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQJsonArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonArray)
	return rv2
}
func QJsonArray_FromStringList(list *QStringList) *QJsonArray /*123*/ {
	var nilthis *QJsonArray
	rv := nilthis.FromStringList(list)
	return rv
}

// /usr/include/qt/QtCore/qjsonarray.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QJsonArray) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QJsonArray) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonarray.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QJsonArray) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:97
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue at(int)
func (this *QJsonArray) At(i int) *QJsonValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray2atEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:98
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue first()
func (this *QJsonArray) First() *QJsonValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray5firstEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:99
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue last()
func (this *QJsonArray) Last() *QJsonValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray4lastEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QJsonArray) RemoveAt(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray8removeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:104
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue takeAt(int)
func (this *QJsonArray) TakeAt(i int) *QJsonValue /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonarray.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeFirst()
func (this *QJsonArray) RemoveFirst() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray11removeFirstEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeLast()
func (this *QJsonArray) RemoveLast() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray10removeLastEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replace(int, const QJsonValue &)
func (this *QJsonArray) Replace(i int, value *QJsonValue) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray7replaceEiRK10QJsonValue", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QJsonValue &)
func (this *QJsonArray) Contains(element *QJsonValue) bool {
	var convArg0 = element.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray8containsERK10QJsonValue", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonarray.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonArray &)
func (this *QJsonArray) Swap(other *QJsonArray) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator begin()
func (this *QJsonArray) Begin() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:215
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator begin()
func (this *QJsonArray) Begin_1() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constBegin()
func (this *QJsonArray) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::iterator end()
func (this *QJsonArray) End() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator end()
func (this *QJsonArray) End_1() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonArray::const_iterator constEnd()
func (this *QJsonArray) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonarray.h:236
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_back(const QJsonValue &)
func (this *QJsonArray) Push_back(t *QJsonValue) {
	var convArg0 = t.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray9push_backERK10QJsonValue", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:237
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void push_front(const QJsonValue &)
func (this *QJsonArray) Push_front(t *QJsonValue) {
	var convArg0 = t.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray10push_frontERK10QJsonValue", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_front()
func (this *QJsonArray) Pop_front() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray9pop_frontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:239
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void pop_back()
func (this *QJsonArray) Pop_back() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QJsonArray8pop_backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonarray.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty()
func (this *QJsonArray) Empty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QJsonArray5emptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
