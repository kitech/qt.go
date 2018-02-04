package qtcore

// /usr/include/qt/QtCore/qjsonobject.h
// #include <qjsonobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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

type QJsonObject struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJsonObject)
	return gothis
}

// /usr/include/qt/QtCore/qjsonobject.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJsonObject()
func DeleteQJsonObject(this *QJsonObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qjsonobject.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QJsonObject &)
func (this *QJsonObject) Swap(other *QJsonObject) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QJsonObject) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QJsonObject) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length()
func (this *QJsonObject) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qjsonobject.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QJsonObject) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:107
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue value(const QString &)
func (this *QJsonObject) Value(key *QString) *QJsonValue /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:108
// index:1
// Public Visibility=Default Availability=Available
// [24] QJsonValue value(QLatin1String)
func (this *QJsonObject) Value_1(key *QLatin1String /*123*/) *QJsonValue /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)
func (this *QJsonObject) Remove(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:115
// index:0
// Public Visibility=Default Availability=Available
// [24] QJsonValue take(const QString &)
func (this *QJsonObject) Take(key *QString) *QJsonValue /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4takeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJsonValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonValue)
	return rv2
}

// /usr/include/qt/QtCore/qjsonobject.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &)
func (this *QJsonObject) Contains(key *QString) bool {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:117
// index:1
// Public Visibility=Default Availability=Available
// [1] bool contains(QLatin1String)
func (this *QJsonObject) Contains_1(key *QLatin1String /*123*/) bool {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qjsonobject.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::iterator begin()
func (this *QJsonObject) Begin() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:215
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator begin()
func (this *QJsonObject) Begin_1() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constBegin()
func (this *QJsonObject) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::iterator end()
func (this *QJsonObject) End() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:218
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator end()
func (this *QJsonObject) End_1() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constEnd()
func (this *QJsonObject) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:225
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject::iterator find(const QString &)
func (this *QJsonObject) Find(key *QString) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:226
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonObject::iterator find(QLatin1String)
func (this *QJsonObject) Find_1(key *QLatin1String /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:227
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator find(const QString &)
func (this *QJsonObject) Find_2(key *QString) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:228
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator find(QLatin1String)
func (this *QJsonObject) Find_3(key *QLatin1String /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:229
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constFind(const QString &)
func (this *QJsonObject) ConstFind(key *QString) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:230
// index:1
// Public Visibility=Default Availability=Available
// [16] QJsonObject::const_iterator constFind(QLatin1String)
func (this *QJsonObject) ConstFind_1(key *QLatin1String /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qjsonobject.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty()
func (this *QJsonObject) Empty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5emptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
