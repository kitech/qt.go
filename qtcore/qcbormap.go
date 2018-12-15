package qtcore

// /usr/include/qt/QtCore/qcbormap.h
// #include <qcbormap.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QCborMap struct {
	*qtrt.CObject
}
type QCborMap_ITF interface {
	QCborMap_PTR() *QCborMap
}

func (ptr *QCborMap) QCborMap_PTR() *QCborMap { return ptr }

func (this *QCborMap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborMap) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborMapFromPointer(cthis unsafe.Pointer) *QCborMap {
	return &QCborMap{&qtrt.CObject{cthis}}
}
func (*QCborMap) NewFromPointer(cthis unsafe.Pointer) *QCborMap {
	return NewQCborMapFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcbormap.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCborMap()

/*
Constructs an empty CBOR Map object.

See also isEmpty().
*/
func (*QCborMap) NewForInherit() *QCborMap {
	return NewQCborMap()
}
func NewQCborMap() *QCborMap {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborMapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborMap)
	return gothis
}

// /usr/include/qt/QtCore/qcbormap.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCborMap()

/*

 */
func DeleteQCborMap(this *QCborMap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcbormap.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCborMap &)

/*
Swaps the contents of this map and other.
*/
func (this *QCborMap) Swap(other QCborMap_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborMap_PTR() != nil {
		convArg0 = other.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:183
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue toCborValue() const

/*
Explicitly constructs a QCborValue object that represents this map. This function is usually not necessary since QCborValue has a constructor for QCborMap, so the conversion is implicit.

Converting QCborMap to QCborValue allows it to be used in any context where QCborValues can be used, including as keys and mapped types in QCborMap, as well as QCborValue::toCbor().

See also QCborValue::QCborValue(const QCborMap &).
*/
func (this *QCborMap) ToCborValue() *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap11toCborValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] qsizetype size() const

/*
Returns the number of elements in this map.

See also isEmpty().
*/
func (this *QCborMap) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qcbormap.h:186
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this map is empty (that is, size() is 0).

See also size() and clear().
*/
func (this *QCborMap) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Empties this map.

See also isEmpty().
*/
func (this *QCborMap) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:190
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue value(qint64) const

/*
Returns the QCborValue element in this map that corresponds to key key, if there is one. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map does not contain key key, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the key was not present from the situation where the key was mapped to an undefined value.

If the map contains more than one key equal to key, it is undefined which one the return from function will reference. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also operator[](qint64), find(qint64), constFind(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) Value(key int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5valueEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:192
// index:1
// Public inline Visibility=Default Availability=Available
// [24] QCborValue value(QLatin1String) const

/*
Returns the QCborValue element in this map that corresponds to key key, if there is one. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map does not contain key key, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the key was not present from the situation where the key was mapped to an undefined value.

If the map contains more than one key equal to key, it is undefined which one the return from function will reference. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also operator[](qint64), find(qint64), constFind(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) Value1(key QLatin1String_ITF /*123*/) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5valueE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:194
// index:2
// Public inline Visibility=Default Availability=Available
// [24] QCborValue value(const QString &) const

/*
Returns the QCborValue element in this map that corresponds to key key, if there is one. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map does not contain key key, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the key was not present from the situation where the key was mapped to an undefined value.

If the map contains more than one key equal to key, it is undefined which one the return from function will reference. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also operator[](qint64), find(qint64), constFind(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) Value2(key string) *QCborValue /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:196
// index:3
// Public inline Visibility=Default Availability=Available
// [24] QCborValue value(const QCborValue &) const

/*
Returns the QCborValue element in this map that corresponds to key key, if there is one. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map does not contain key key, this function returns a QCborValue containing an undefined value. For that reason, it is not possible with this function to tell apart the situation where the key was not present from the situation where the key was mapped to an undefined value.

If the map contains more than one key equal to key, it is undefined which one the return from function will reference. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also operator[](qint64), find(qint64), constFind(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) Value3(key QCborValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5valueERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [24] const QCborValue operator[](qint64) const

/*

 */
func (this *QCborMap) Operator_get_index(key int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:204
// index:1
// Public inline Visibility=Default Availability=Available
// [24] const QCborValue operator[](QLatin1String) const

/*

 */
func (this *QCborMap) Operator_get_index1(key QLatin1String_ITF /*123*/) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:206
// index:2
// Public inline Visibility=Default Availability=Available
// [24] const QCborValue operator[](const QString &) const

/*

 */
func (this *QCborMap) Operator_get_index2(key string) *QCborValue /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:208
// index:3
// Public inline Visibility=Default Availability=Available
// [24] const QCborValue operator[](const QCborValue &) const

/*

 */
func (this *QCborMap) Operator_get_index3(key QCborValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapixERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:214
// index:4
// Public Visibility=Default Availability=Available
// [16] QCborValueRef operator[](qint64)

/*

 */
func (this *QCborMap) Operator_get_index4(key int64) *QCborValueRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:215
// index:5
// Public Visibility=Default Availability=Available
// [16] QCborValueRef operator[](QLatin1String)

/*

 */
func (this *QCborMap) Operator_get_index5(key QLatin1String_ITF /*123*/) *QCborValueRef /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapixE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:216
// index:6
// Public Visibility=Default Availability=Available
// [16] QCborValueRef operator[](const QString &)

/*

 */
func (this *QCborMap) Operator_get_index6(key string) *QCborValueRef /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:217
// index:7
// Public Visibility=Default Availability=Available
// [16] QCborValueRef operator[](const QCborValue &)

/*

 */
func (this *QCborMap) Operator_get_index7(key QCborValue_ITF) *QCborValueRef /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMapixERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValueRef)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:219
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QCborValue take(qint64)

/*
Removes the key key and the corresponding value from the map and returns the value, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

See also value(qint64), operator[](qint64), find(qint64), contains(qint64), take(QLatin1String), take(const QString &), take(const QCborValue &), and insert().
*/
func (this *QCborMap) Take(key int64) *QCborValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4takeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:221
// index:1
// Public inline Visibility=Default Availability=Available
// [24] QCborValue take(QLatin1String)

/*
Removes the key key and the corresponding value from the map and returns the value, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

See also value(qint64), operator[](qint64), find(qint64), contains(qint64), take(QLatin1String), take(const QString &), take(const QCborValue &), and insert().
*/
func (this *QCborMap) Take1(key QLatin1String_ITF /*123*/) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4takeE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:223
// index:2
// Public inline Visibility=Default Availability=Available
// [24] QCborValue take(const QString &)

/*
Removes the key key and the corresponding value from the map and returns the value, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

See also value(qint64), operator[](qint64), find(qint64), contains(qint64), take(QLatin1String), take(const QString &), take(const QCborValue &), and insert().
*/
func (this *QCborMap) Take2(key string) *QCborValue /*123*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4takeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:225
// index:3
// Public inline Visibility=Default Availability=Available
// [24] QCborValue take(const QCborValue &)

/*
Removes the key key and the corresponding value from the map and returns the value, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

See also value(qint64), operator[](qint64), find(qint64), contains(qint64), take(QLatin1String), take(const QString &), take(const QCborValue &), and insert().
*/
func (this *QCborMap) Take3(key QCborValue_ITF) *QCborValue /*123*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4takeERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborValue)
	return rv2
}

// /usr/include/qt/QtCore/qcbormap.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void remove(qint64)

/*
Removes the key key and the corresponding value from the map, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

remove(QLatin1String), remove(const QString &), remove(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), and contains(qint64).
*/
func (this *QCborMap) Remove(key int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap6removeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:229
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void remove(QLatin1String)

/*
Removes the key key and the corresponding value from the map, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

remove(QLatin1String), remove(const QString &), remove(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), and contains(qint64).
*/
func (this *QCborMap) Remove1(key QLatin1String_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap6removeE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:231
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void remove(const QString &)

/*
Removes the key key and the corresponding value from the map, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

remove(QLatin1String), remove(const QString &), remove(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), and contains(qint64).
*/
func (this *QCborMap) Remove2(key string) {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:233
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void remove(const QCborValue &)

/*
Removes the key key and the corresponding value from the map, if it is found. If the map contains no such key, this function does nothing.

If the map contains more than one key equal to key, it is undefined which one this function will remove. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

remove(QLatin1String), remove(const QString &), remove(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), and contains(qint64).
*/
func (this *QCborMap) Remove3(key QCborValue_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap6removeERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcbormap.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(qint64) const

/*
Returns true if this map contains a key-value pair identified by key key. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), contains(QLatin1String), remove(const QString &), and remove(const QCborValue &).
*/
func (this *QCborMap) Contains(key int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap8containsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:237
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String) const

/*
Returns true if this map contains a key-value pair identified by key key. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), contains(QLatin1String), remove(const QString &), and remove(const QCborValue &).
*/
func (this *QCborMap) Contains1(key QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap8containsE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:239
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &) const

/*
Returns true if this map contains a key-value pair identified by key key. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), contains(QLatin1String), remove(const QString &), and remove(const QCborValue &).
*/
func (this *QCborMap) Contains2(key string) bool {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:241
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QCborValue &) const

/*
Returns true if this map contains a key-value pair identified by key key. CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), contains(QLatin1String), remove(const QString &), and remove(const QCborValue &).
*/
func (this *QCborMap) Contains3(key QCborValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap8containsERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QCborMap &) const

/*
Compares this map and other, comparing each element in sequence, and returns an integer that indicates whether this map should be sorted prior to (if the result is negative) or after other (if the result is positive). If this function returns 0, the two maps are equal and contain the same elements.

Note that CBOR maps are unordered, which means that two maps containing the very same pairs but in different order will still compare differently. To avoid this, it is recommended to insert elements into the map in a predictable order, such as by ascending key value. In fact, maps with keys in sorted order are required for Canonical CBOR representation.

For more information on CBOR sorting order, see QCborValue::compare().

See also QCborValue::compare(), QCborArray::compare(), and operator==().
*/
func (this *QCborMap) Compare(other QCborMap_ITF) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborMap_PTR() != nil {
		convArg0 = other.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qcbormap.h:254
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QCborMap &) const

/*

 */
func (this *QCborMap) Operator_equal_equal(other QCborMap_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborMap_PTR() != nil {
		convArg0 = other.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCborMap &) const

/*

 */
func (this *QCborMap) Operator_not_equal(other QCborMap_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborMap_PTR() != nil {
		convArg0 = other.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:258
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QCborMap &) const

/*

 */
func (this *QCborMap) Operator_less_than(other QCborMap_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCborMap_PTR() != nil {
		convArg0 = other.QCborMap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMapltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:264
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::iterator begin()

/*
Returns a map iterator pointing to the first key-value pair of this map. If this map is empty, the returned iterator will be the same as end().

See also constBegin() and end().
*/
func (this *QCborMap) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:266
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator begin() const

/*
Returns a map iterator pointing to the first key-value pair of this map. If this map is empty, the returned iterator will be the same as end().

See also constBegin() and end().
*/
func (this *QCborMap) Begin1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:265
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constBegin() const

/*
Returns a map iterator pointing to the first key-value pair of this map. If this map is empty, the returned iterator will be the same as constEnd().

See also begin() and constEnd().
*/
func (this *QCborMap) ConstBegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap10constBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:267
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator cbegin() const

/*
Returns a map iterator pointing to the first key-value pair of this map. If this map is empty, the returned iterator will be the same as constEnd().

See also begin() and constEnd().
*/
func (this *QCborMap) Cbegin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::iterator end()

/*
Returns a map iterator representing an element just past the last element in the map.

See also begin(), constBegin(), find(), and constFind().
*/
func (this *QCborMap) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:270
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator end() const

/*
Returns a map iterator representing an element just past the last element in the map.

See also begin(), constBegin(), find(), and constFind().
*/
func (this *QCborMap) End1() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:269
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constEnd() const

/*
Returns a map iterator representing an element just past the last element in the map.

See also begin(), constBegin(), find(), and constFind().
*/
func (this *QCborMap) ConstEnd() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap8constEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:271
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator cend() const

/*
Returns a map iterator representing an element just past the last element in the map.

See also begin(), constBegin(), find(), and constFind().
*/
func (this *QCborMap) Cend() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:276
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty() const

/*
Synonym for isEmpty(). This function is provided for compatibility with generic code that uses the Standard Library API.

Returns true if this map is empty (size() == 0).

See also isEmpty() and size().
*/
func (this *QCborMap) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcbormap.h:278
// index:0
// Public Visibility=Default Availability=Available
// [16] QCborMap::iterator find(qint64)

/*

 */
func (this *QCborMap) Find(key int64) unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4findEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:279
// index:1
// Public Visibility=Default Availability=Available
// [16] QCborMap::iterator find(QLatin1String)

/*

 */
func (this *QCborMap) Find1(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4findE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:280
// index:2
// Public Visibility=Default Availability=Available
// [16] QCborMap::iterator find(const QString &)

/*

 */
func (this *QCborMap) Find2(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4findERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:281
// index:3
// Public Visibility=Default Availability=Available
// [16] QCborMap::iterator find(const QCborValue &)

/*

 */
func (this *QCborMap) Find3(key QCborValue_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCborMap4findERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:286
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator find(qint64) const

/*

 */
func (this *QCborMap) Find4(key int64) unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4findEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:287
// index:5
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator find(QLatin1String) const

/*

 */
func (this *QCborMap) Find5(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4findE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:288
// index:6
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator find(const QString &) const

/*

 */
func (this *QCborMap) Find6(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4findERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:289
// index:7
// Public inline Visibility=Default Availability=Available
// [16] QCborMap::const_iterator find(const QCborValue &) const

/*

 */
func (this *QCborMap) Find7(key QCborValue_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap4findERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:282
// index:0
// Public Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constFind(qint64) const

/*
Returns a map iterator to the key-value pair whose key is key, if the map contains such a pair. If it doesn't, this function returns constEnd().

CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map contains more than one key equal to key, it is undefined which one this function will find. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) ConstFind(key int64) unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap9constFindEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:283
// index:1
// Public Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constFind(QLatin1String) const

/*
Returns a map iterator to the key-value pair whose key is key, if the map contains such a pair. If it doesn't, this function returns constEnd().

CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map contains more than one key equal to key, it is undefined which one this function will find. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) ConstFind1(key QLatin1String_ITF /*123*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap9constFindE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:284
// index:2
// Public Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constFind(const QString &) const

/*
Returns a map iterator to the key-value pair whose key is key, if the map contains such a pair. If it doesn't, this function returns constEnd().

CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map contains more than one key equal to key, it is undefined which one this function will find. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) ConstFind2(key string) unsafe.Pointer /*444*/ {
	var tmpArg0 = NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap9constFindERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:285
// index:3
// Public Visibility=Default Availability=Available
// [16] QCborMap::const_iterator constFind(const QCborValue &) const

/*
Returns a map iterator to the key-value pair whose key is key, if the map contains such a pair. If it doesn't, this function returns constEnd().

CBOR recommends using integer keys, since they occupy less space and are simpler to encode and decode.

If the map contains more than one key equal to key, it is undefined which one this function will find. QCborMap does not allow inserting duplicate keys, but it is possible to create such a map by decoding a CBOR stream with them. They are usually not permitted and having duplicate keys is usually an indication of a problem in the sender.

value(QLatin1String), value(const QString &), value(const QCborValue &)

See also value(qint64), operator[](qint64), find(qint64), remove(qint64), and contains(qint64).
*/
func (this *QCborMap) ConstFind3(key QCborValue_ITF) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if key != nil && key.QCborValue_PTR() != nil {
		convArg0 = key.QCborValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap9constFindERK10QCborValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qcbormap.h:322
// index:0
// Public Visibility=Default Availability=Available
// [16] QJsonObject toJsonObject() const

/*
Recursively converts every QCborValue value in this array to JSON using QCborValue::toJsonValue() and creates a string key for all keys that aren't strings, then returns the corresponding QJsonObject composed of those associations.

Please note that CBOR contains a richer and wider type set than JSON, so some information may be lost in this conversion. For more details on what conversions are applied, see QCborValue::toJsonValue().
*/
func (this *QCborMap) ToJsonObject() *QJsonObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCborMap12toJsonObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJsonObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJsonObject)
	return rv2
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
