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

// /usr/include/qt/QtCore/qjsonobject.h:61
// index:0
// void QJsonObject()
func NewQJsonObject() *QJsonObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQJsonObjectFromPointer(cthis)
	return gothis
}
func NewQJsonObjectFromPointer(cthis unsafe.Pointer) *QJsonObject {
	return &QJsonObject{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qjsonobject.h:72
// index:0
// void ~QJsonObject()
func DeleteQJsonObject(*QJsonObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObjectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:90
// index:0
// inline
// void swap(class QJsonObject &)
func (this *QJsonObject) Swap(other unsafe.Pointer) {
	// 0: (, other QJsonObject &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:97
// index:0
// QVariantMap toVariantMap()
func (this *QJsonObject) ToVariantMap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject12toVariantMapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:99
// index:0
// QVariantHash toVariantHash()
func (this *QJsonObject) ToVariantHash() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject13toVariantHashEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:101
// index:0
// QStringList keys()
func (this *QJsonObject) Keys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4keysEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:102
// index:0
// int size()
func (this *QJsonObject) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:103
// index:0
// inline
// int count()
func (this *QJsonObject) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:104
// index:0
// inline
// int length()
func (this *QJsonObject) Length() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject6lengthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:105
// index:0
// bool isEmpty()
func (this *QJsonObject) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:107
// index:0
// QJsonValue value(const class QString &)
func (this *QJsonObject) Value(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:108
// index:1
// QJsonValue value(class QLatin1String)
func (this *QJsonObject) Value_1(key unsafe.Pointer) {
	// 1: (, key QLatin1String), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5valueE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:114
// index:0
// void remove(const class QString &)
func (this *QJsonObject) Remove(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject6removeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:115
// index:0
// QJsonValue take(const class QString &)
func (this *QJsonObject) Take(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4takeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:116
// index:0
// bool contains(const class QString &)
func (this *QJsonObject) Contains(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:117
// index:1
// bool contains(class QLatin1String)
func (this *QJsonObject) Contains_1(key unsafe.Pointer) {
	// 1: (, key QLatin1String), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8containsE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:214
// index:0
// inline
// QJsonObject::iterator begin()
func (this *QJsonObject) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:215
// index:1
// inline
// QJsonObject::const_iterator begin()
func (this *QJsonObject) Begin_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5beginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:216
// index:0
// inline
// QJsonObject::const_iterator constBegin()
func (this *QJsonObject) ConstBegin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject10constBeginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:217
// index:0
// inline
// QJsonObject::iterator end()
func (this *QJsonObject) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:218
// index:1
// inline
// QJsonObject::const_iterator end()
func (this *QJsonObject) End_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject3endEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:219
// index:0
// inline
// QJsonObject::const_iterator constEnd()
func (this *QJsonObject) ConstEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject8constEndEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:220
// index:0
// QJsonObject::iterator erase(class QJsonObject::iterator)
func (this *QJsonObject) Erase(it unsafe.Pointer) {
	// 0: (, it QJsonObject::iterator), (it)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject5eraseENS_8iteratorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), it)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:225
// index:0
// QJsonObject::iterator find(const class QString &)
func (this *QJsonObject) Find(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:226
// index:1
// QJsonObject::iterator find(class QLatin1String)
func (this *QJsonObject) Find_1(key unsafe.Pointer) {
	// 1: (, key QLatin1String), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:227
// index:2
// inline
// QJsonObject::const_iterator find(const class QString &)
func (this *QJsonObject) Find_2(key unsafe.Pointer) {
	// 2: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:228
// index:3
// inline
// QJsonObject::const_iterator find(class QLatin1String)
func (this *QJsonObject) Find_3(key unsafe.Pointer) {
	// 3: (, key QLatin1String), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject4findE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:229
// index:0
// QJsonObject::const_iterator constFind(const class QString &)
func (this *QJsonObject) ConstFind(key unsafe.Pointer) {
	// 0: (, key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:230
// index:1
// QJsonObject::const_iterator constFind(class QLatin1String)
func (this *QJsonObject) ConstFind_1(key unsafe.Pointer) {
	// 1: (, key QLatin1String), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject9constFindE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qjsonobject.h:238
// index:0
// inline
// bool empty()
func (this *QJsonObject) Empty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QJsonObject5emptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
