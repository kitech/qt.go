//  header block begin
// /usr/include/qt/QtCore/qjsonobject.h
// #include <qjsonobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QJsonObject struct {
	*qtrt.CObject
}

func (this *QJsonObject) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQJsonObjectFromPointer(cthis unsafe.Pointer) *QJsonObject {
	return &QJsonObject{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qjsonobject.h:61
// index:0
// Public
// void QJsonObject()
func NewQJsonObject() *QJsonObject {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonObjectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qjsonobject.h:72
// index:0
// Public
// void ~QJsonObject()
func DeleteQJsonObject(*QJsonObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:90
// index:0
// Public inline
// void swap(class QJsonObject &)
func (this *QJsonObject) Swap(other *QJsonObject) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:97
// index:0
// Public
// QVariantMap toVariantMap()
func (this *QJsonObject) ToVariantMap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject12toVariantMapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:99
// index:0
// Public
// QVariantHash toVariantHash()
func (this *QJsonObject) ToVariantHash() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject13toVariantHashEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:101
// index:0
// Public
// QStringList keys()
func (this *QJsonObject) Keys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4keysEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:102
// index:0
// Public
// int size()
func (this *QJsonObject) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:103
// index:0
// Public inline
// int count()
func (this *QJsonObject) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:104
// index:0
// Public inline
// int length()
func (this *QJsonObject) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:105
// index:0
// Public
// bool isEmpty()
func (this *QJsonObject) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:107
// index:0
// Public
// QJsonValue value(const class QString &)
func (this *QJsonObject) Value(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:108
// index:1
// Public
// QJsonValue value(class QLatin1String)
func (this *QJsonObject) Value_1(key *QLatin1String) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:114
// index:0
// Public
// void remove(const class QString &)
func (this *QJsonObject) Remove(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:115
// index:0
// Public
// QJsonValue take(const class QString &)
func (this *QJsonObject) Take(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4takeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:116
// index:0
// Public
// bool contains(const class QString &)
func (this *QJsonObject) Contains(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:117
// index:1
// Public
// bool contains(class QLatin1String)
func (this *QJsonObject) Contains_1(key *QLatin1String) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:214
// index:0
// Public inline
// QJsonObject::iterator begin()
func (this *QJsonObject) Begin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:215
// index:1
// Public inline
// QJsonObject::const_iterator begin()
func (this *QJsonObject) Begin_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:216
// index:0
// Public inline
// QJsonObject::const_iterator constBegin()
func (this *QJsonObject) ConstBegin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject10constBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:217
// index:0
// Public inline
// QJsonObject::iterator end()
func (this *QJsonObject) End() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:218
// index:1
// Public inline
// QJsonObject::const_iterator end()
func (this *QJsonObject) End_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:219
// index:0
// Public inline
// QJsonObject::const_iterator constEnd()
func (this *QJsonObject) ConstEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8constEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:225
// index:0
// Public
// QJsonObject::iterator find(const class QString &)
func (this *QJsonObject) Find(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:226
// index:1
// Public
// QJsonObject::iterator find(class QLatin1String)
func (this *QJsonObject) Find_1(key *QLatin1String) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:227
// index:2
// Public inline
// QJsonObject::const_iterator find(const class QString &)
func (this *QJsonObject) Find_2(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:228
// index:3
// Public inline
// QJsonObject::const_iterator find(class QLatin1String)
func (this *QJsonObject) Find_3(key *QLatin1String) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:229
// index:0
// Public
// QJsonObject::const_iterator constFind(const class QString &)
func (this *QJsonObject) ConstFind(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:230
// index:1
// Public
// QJsonObject::const_iterator constFind(class QLatin1String)
func (this *QJsonObject) ConstFind_1(key *QLatin1String) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qjsonobject.h:238
// index:0
// Public inline
// bool empty()
func (this *QJsonObject) Empty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5emptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
