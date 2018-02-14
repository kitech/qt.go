package qtcore

// /usr/include/qt/QtCore/qjsonobject.h
// #include <qjsonobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QJsonObject struct {
	*qtrt.CObject
}
type QJsonObject_ITF interface {
	QJsonObject_PTR() *QJsonObject
}

func (ptr *QJsonObject) QJsonObject_PTR() *QJsonObject { return ptr }

func (this *QJsonObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonObjectFromPointer(cthis unsafe.Pointer) *QJsonObject {
	return &QJsonObject{&qtrt.CObject{cthis}}
}
func (*QJsonObject) NewFromPointer(cthis unsafe.Pointer) *QJsonObject {
	return NewQJsonObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsonobject.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJsonObject()
func NewQJsonObject() *QJsonObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonObject)
	return gothis
}

// /usr/include/qt/QtCore/qjsonobject.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonObject()
func DeleteQJsonObject(this *QJsonObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonobject.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject & operator=(const QJsonObject &)
func (this *QJsonObject) Operator_equal(other QJsonObject_ITF) *QJsonObject {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonObject_PTR() != nil {
		convArg0 = other.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:84
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject & operator=(QJsonObject &&)
func (this *QJsonObject) Operator_equal_1(other unsafe.Pointer /*333*/) *QJsonObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonObject &)
func (this *QJsonObject) Swap(other QJsonObject_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonObject_PTR() != nil {
		convArg0 = other.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList keys()
func (this *QJsonObject) Keys() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject4keysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QJsonObject) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QJsonObject) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length()
func (this *QJsonObject) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QJsonObject) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:107
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue value(const QString &)
func (this *QJsonObject) Value(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:108
// index:1
// Public Visibility=Default Availability=Available
// [24] QJsonValue value(QLatin1String)
func (this *QJsonObject) Value_1(key QLatin1String_ITF /*123*/) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject5valueE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:109
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue operator[](const QString &)
func (this *QJsonObject) Operator_get_index(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObjectixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:110
// index:1
// Public inline Visibility=Default Availability=Available
// [24] QJsonValue operator[](QLatin1String)
func (this *QJsonObject) Operator_get_index_1(key QLatin1String_ITF /*123*/) *QJsonValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObjectixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:111
// index:2
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef operator[](const QString &)
func (this *QJsonObject) Operator_get_index_2(key string) *QJsonValueRef /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:112
// index:3
// Public Visibility=Default Availability=Available
// [16] QJsonValueRef operator[](QLatin1String)
func (this *QJsonObject) Operator_get_index_3(key QLatin1String_ITF /*123*/) *QJsonValueRef /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObjectixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)
func (this *QJsonObject) Remove(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:115
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue take(const QString &)
func (this *QJsonObject) Take(key string) *QJsonValue /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject4takeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &)
func (this *QJsonObject) Contains(key string) bool {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:117
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(QLatin1String)
func (this *QJsonObject) Contains_1(key QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject8containsE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QJsonObject &)
func (this *QJsonObject) Operator_equal_equal(other QJsonObject_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonObject_PTR() != nil {
		convArg0 = other.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObjecteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QJsonObject &)
func (this *QJsonObject) Operator_not_equal(other QJsonObject_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QJsonObject_PTR() != nil {
		convArg0 = other.QJsonObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObjectneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::iterator begin()
func (this *QJsonObject) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:215
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator begin()
func (this *QJsonObject) Begin_1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constBegin()
func (this *QJsonObject) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::iterator end()
func (this *QJsonObject) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator end()
func (this *QJsonObject) End_1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constEnd()
func (this *QJsonObject) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:225
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject::iterator find(const QString &)
func (this *QJsonObject) Find(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject4findERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:226
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonObject::iterator find(QLatin1String)
func (this *QJsonObject) Find_1(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QJsonObject4findE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:227
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator find(const QString &)
func (this *QJsonObject) Find_2(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject4findERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:228
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator find(QLatin1String)
func (this *QJsonObject) Find_3(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject4findE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:229
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constFind(const QString &)
func (this *QJsonObject) ConstFind(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject9constFindERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:230
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constFind(QLatin1String)
func (this *QJsonObject) ConstFind_1(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject9constFindE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty()
func (this *QJsonObject) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QJsonObject5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
