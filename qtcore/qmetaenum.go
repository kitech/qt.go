package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QMetaEnum struct {
	*qtrt.CObject
}
type QMetaEnum_ITF interface {
	QMetaEnum_PTR() *QMetaEnum
}

func (ptr *QMetaEnum) QMetaEnum_PTR() *QMetaEnum { return ptr }

func (this *QMetaEnum) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaEnum) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaEnumFromPointer(cthis unsafe.Pointer) *QMetaEnum {
	return &QMetaEnum{&qtrt.CObject{cthis}}
}
func (*QMetaEnum) NewFromPointer(cthis unsafe.Pointer) *QMetaEnum {
	return NewQMetaEnumFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMetaEnum()

/*

 */
func (*QMetaEnum) NewForInherit() *QMetaEnum {
	return NewQMetaEnum()
}
func NewQMetaEnum() *QMetaEnum {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMetaEnumC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaEnumFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMetaEnum)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:211
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * name() const

/*

 */
func (this *QMetaEnum) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:212
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlag() const

/*

 */
func (this *QMetaEnum) IsFlag() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum6isFlagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:213
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScoped() const

/*

 */
func (this *QMetaEnum) IsScoped() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum8isScopedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:215
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyCount() const

/*

 */
func (this *QMetaEnum) KeyCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum8keyCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:216
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * key(int) const

/*

 */
func (this *QMetaEnum) Key(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum3keyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int value(int) const

/*

 */
func (this *QMetaEnum) Value(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum5valueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * scope() const

/*

 */
func (this *QMetaEnum) Scope() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum5scopeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyToValue(const char *, bool *) const

/*

 */
func (this *QMetaEnum) KeyToValue(key string, ok *bool) int {
	var convArg0 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum10keyToValueEPKcPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyToValue(const char *, bool *) const

/*

 */
func (this *QMetaEnum) KeyToValuep(key string) int {
	var convArg0 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum10keyToValueEPKcPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * valueToKey(int) const

/*

 */
func (this *QMetaEnum) ValueToKey(value int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum10valueToKeyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int keysToValue(const char *, bool *) const

/*

 */
func (this *QMetaEnum) KeysToValue(keys string, ok *bool) int {
	var convArg0 = qtrt.CString(keys)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum11keysToValueEPKcPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] int keysToValue(const char *, bool *) const

/*

 */
func (this *QMetaEnum) KeysToValuep(keys string) int {
	var convArg0 = qtrt.CString(keys)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum11keysToValueEPKcPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray valueToKeys(int) const

/*

 */
func (this *QMetaEnum) ValueToKeys(value int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum11valueToKeysEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject() const

/*

 */
func (this *QMetaEnum) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum19enclosingMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmetaobject.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QMetaEnum) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMetaEnum7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQMetaEnum(this *QMetaEnum) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMetaEnumD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
