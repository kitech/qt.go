package qtqml

// /usr/include/qt/QtQml/qqmlpropertymap.h
// #include <qqmlpropertymap.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

// QVariant updateValue(const class QString &, const class QVariant &)
func (this *QQmlPropertyMap) InheritUpdateValue(f func(key string, input *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "updateValue", f)
}

/*

 */
type QQmlPropertyMap struct {
	*qtcore.QObject
}
type QQmlPropertyMap_ITF interface {
	qtcore.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap { return ptr }

func (this *QQmlPropertyMap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlPropertyMap) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlPropertyMapFromPointer(cthis unsafe.Pointer) *QQmlPropertyMap {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlPropertyMap{bcthis0}
}
func (*QQmlPropertyMap) NewFromPointer(cthis unsafe.Pointer) *QQmlPropertyMap {
	return NewQQmlPropertyMapFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQmlPropertyMap) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlPropertyMap(QObject *)

/*
Constructs a bindable map with parent object parent.
*/
func NewQQmlPropertyMap(parent qtcore.QObject_ITF /*777 QObject **/) *QQmlPropertyMap {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyMapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlPropertyMap")
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlPropertyMap(QObject *)

/*
Constructs a bindable map with parent object parent.
*/
func NewQQmlPropertyMap__() *QQmlPropertyMap {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyMapFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQmlPropertyMap")
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlPropertyMap()

/*

 */
func DeleteQQmlPropertyMap(this *QQmlPropertyMap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:61
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QString &) const

/*
Returns the value associated with key.

If no value has been set for this key (or if the value has been cleared), an invalid QVariant is returned.
*/
func (this *QQmlPropertyMap) Value(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear(const QString &)

/*
Clears the value (if any) associated with key.
*/
func (this *QQmlPropertyMap) Clear(key string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap5clearERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList keys() const

/*
Returns the list of keys.

Keys that have been cleared will still appear in this list, even though their associated values are invalid QVariants.
*/
func (this *QQmlPropertyMap) Keys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap4keysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
This is an overloaded function.

Same as size().
*/
func (this *QQmlPropertyMap) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int size() const

/*
Returns the number of keys in the map.

See also isEmpty() and count().
*/
func (this *QQmlPropertyMap) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the map contains no keys; otherwise returns false.

See also size().
*/
func (this *QQmlPropertyMap) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &) const

/*
Returns true if the map contains key.

See also size().
*/
func (this *QQmlPropertyMap) Contains(key string) bool {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:72
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant & operator[](const QString &)

/*

 */
func (this *QQmlPropertyMap) Operator_get_index(key string) *qtcore.QVariant {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:73
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant operator[](const QString &) const

/*

 */
func (this *QQmlPropertyMap) Operator_get_index_1(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMapixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QString &, const QVariant &)

/*
This signal is emitted whenever one of the values in the map is changed. key is the key corresponding to the value that was changed.

Note: valueChanged() is NOT emitted when changes are made by calling insert() or clear() - it is only emitted when a value is updated from QML.
*/
func (this *QQmlPropertyMap) ValueChanged(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap12valueChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant updateValue(const QString &, const QVariant &)

/*
Returns the new value to be stored for the key key. This function is provided to intercept updates to a property from QML, where the value provided from QML is input.

Override this function to manipulate the property value as it is updated. Note that this function is only invoked when the value is updated from QML.
*/
func (this *QQmlPropertyMap) UpdateValue(key string, input qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if input != nil && input.QVariant_PTR() != nil {
		convArg1 = input.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap11updateValueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
