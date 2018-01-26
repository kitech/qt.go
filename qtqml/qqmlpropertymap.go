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
import "mkuse/cffiqt"
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
		ffiqt.KeepMe()
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQmlPropertyMap) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:58
// index:0
// Public
// void QQmlPropertyMap(class QObject *)
func NewQQmlPropertyMap(parent *qtcore.QObject /*777 QObject **/) *QQmlPropertyMap {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQmlPropertyMapC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlPropertyMapFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:59
// index:0
// Public virtual
// void ~QQmlPropertyMap()
func DeleteQQmlPropertyMap(*QQmlPropertyMap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQmlPropertyMapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:61
// index:0
// Public
// QVariant value(const class QString &)
func (this *QQmlPropertyMap) Value(key *qtcore.QString) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5valueERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:63
// index:0
// Public
// void clear(const class QString &)
func (this *QQmlPropertyMap) Clear(key *qtcore.QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQmlPropertyMap5clearERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:67
// index:0
// Public
// int count()
func (this *QQmlPropertyMap) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:68
// index:0
// Public
// int size()
func (this *QQmlPropertyMap) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:69
// index:0
// Public
// bool isEmpty()
func (this *QQmlPropertyMap) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:70
// index:0
// Public
// bool contains(const class QString &)
func (this *QQmlPropertyMap) Contains(key *qtcore.QString) bool {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QQmlPropertyMap8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:76
// index:0
// Public
// void valueChanged(const class QString &, const class QVariant &)
func (this *QQmlPropertyMap) ValueChanged(key *qtcore.QString, value *qtcore.QVariant) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQmlPropertyMap12valueChangedERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlpropertymap.h:79
// index:0
// Protected virtual
// QVariant updateValue(const class QString &, const class QVariant &)
func (this *QQmlPropertyMap) UpdateValue(key *qtcore.QString, input *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = key.GetCthis()
	var convArg1 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QQmlPropertyMap11updateValueERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
