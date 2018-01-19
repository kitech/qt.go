//  header block begin
// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QMetaEnum struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmetaobject.h:209
// index:0
// inline
// void QMetaEnum()
func NewQMetaEnum() *QMetaEnum {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMetaEnumC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMetaEnum{cthis}
}

// /usr/include/qt/QtCore/qmetaobject.h:211
// index:0
// const char * name()
func (this *QMetaEnum) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:212
// index:0
// bool isFlag()
func (this *QMetaEnum) IsFlag() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum6isFlagEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:213
// index:0
// bool isScoped()
func (this *QMetaEnum) IsScoped() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum8isScopedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:215
// index:0
// int keyCount()
func (this *QMetaEnum) KeyCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum8keyCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:216
// index:0
// const char * key(int)
func (this *QMetaEnum) Key(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum3keyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:217
// index:0
// int value(int)
func (this *QMetaEnum) Value(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum5valueEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:219
// index:0
// const char * scope()
func (this *QMetaEnum) Scope() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum5scopeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:221
// index:0
// int keyToValue(const char *, _Bool *)
func (this *QMetaEnum) KeyToValue(key unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, const char * key, bool * ok), (key, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum10keyToValueEPKcPb", ffiqt.FFI_TYPE_VOID, this.cthis, key, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:222
// index:0
// const char * valueToKey(int)
func (this *QMetaEnum) ValueToKey(value int) {
	// 0: (, int value), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum10valueToKeyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:223
// index:0
// int keysToValue(const char *, _Bool *)
func (this *QMetaEnum) KeysToValue(keys unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, const char * keys, bool * ok), (keys, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum11keysToValueEPKcPb", ffiqt.FFI_TYPE_VOID, this.cthis, keys, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:224
// index:0
// QByteArray valueToKeys(int)
func (this *QMetaEnum) ValueToKeys(value int) {
	// 0: (, int value), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum11valueToKeysEi", ffiqt.FFI_TYPE_VOID, this.cthis, &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:226
// index:0
// inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaEnum) EnclosingMetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum19enclosingMetaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:228
// index:0
// inline
// bool isValid()
func (this *QMetaEnum) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMetaEnum7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
