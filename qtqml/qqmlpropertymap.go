package qtqml

// /usr/include/qt/QtQml/qqmlpropertymap.h
// #include <qqmlpropertymap.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
// QVariant updateValue(const class QString &, const class QVariant &)
func (this *QQmlPropertyMap) InheritUpdateValue(f func(key string, input *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "updateValue", f)
}

type QQmlPropertyMap struct {
	*qtcore.QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QQmlPropertyMap) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlPropertyMap(QObject *)
func NewQQmlPropertyMap(parent *qtcore.QObject /*777 QObject **/) *QQmlPropertyMap {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlPropertyMapFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlPropertyMap()
func DeleteQQmlPropertyMap(this *QQmlPropertyMap) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMapD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:61
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QString &)
func (this *QQmlPropertyMap) Value(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear(const QString &)
func (this *QQmlPropertyMap) Clear(key string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap5clearERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList keys()
func (this *QQmlPropertyMap) Keys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap4keysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QQmlPropertyMap) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int size()
func (this *QQmlPropertyMap) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QQmlPropertyMap) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &)
func (this *QQmlPropertyMap) Contains(key string) bool {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QQmlPropertyMap8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QString &, const QVariant &)
func (this *QQmlPropertyMap) ValueChanged(key string, value *qtcore.QVariant) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap12valueChangedERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant updateValue(const QString &, const QVariant &)
func (this *QQmlPropertyMap) UpdateValue(key string, input *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = input.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QQmlPropertyMap11updateValueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

//  body block end
