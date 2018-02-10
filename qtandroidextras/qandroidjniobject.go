package qtandroidextras

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h
// #include <qandroidjniobject.h>
// #include <Qtjni>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QAndroidJniObject struct {
	*qtrt.CObject
}

func (this *QAndroidJniObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidJniObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidJniObjectFromPointer(cthis unsafe.Pointer) *QAndroidJniObject {
	return &QAndroidJniObject{&qtrt.CObject{cthis}}
}
func (*QAndroidJniObject) NewFromPointer(cthis unsafe.Pointer) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject()
func NewQAndroidJniObject() *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject(const char *)
func NewQAndroidJniObject_1(className string) *QAndroidJniObject {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:58
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject(int)
func NewQAndroidJniObject_2(clazz int) *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2Ei", qtrt.FFI_TYPE_POINTER, clazz)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:60
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject(int)
func NewQAndroidJniObject_3(obj int) *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2Ei", qtrt.FFI_TYPE_POINTER, obj)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniObject()
func DeleteQAndroidJniObject(this *QAndroidJniObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int object()
func (this *QAndroidJniObject) Object() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidJniObject getObjectField(const char *, const char *)
func (this *QAndroidJniObject) GetObjectField(fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var convArg0 = qtrt.CString(fieldName)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sig)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject14getObjectFieldEPKcS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [1] QAndroidJniObject getStaticObjectField(const char *, const char *, const char *)
func (this *QAndroidJniObject) GetStaticObjectField(className string, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(fieldName)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(sig)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject20getStaticObjectFieldEPKcS1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_GetStaticObjectField(className string, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.GetStaticObjectField(className, fieldName, sig)
	return rv
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:140
// index:1
// Public static Visibility=Default Availability=Available
// [1] QAndroidJniObject getStaticObjectField(int, const char *, const char *)
func (this *QAndroidJniObject) GetStaticObjectField_1(clazz int, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var convArg1 = qtrt.CString(fieldName)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(sig)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject20getStaticObjectFieldEiPKcS1_", qtrt.FFI_TYPE_POINTER, clazz, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_GetStaticObjectField_1(clazz int, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.GetStaticObjectField_1(clazz, fieldName, sig)
	return rv
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [1] QAndroidJniObject fromString(const QString &)
func (this *QAndroidJniObject) FromString(string string) *QAndroidJniObject /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject10fromStringERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_FromString(string string) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.FromString(string)
	return rv
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QAndroidJniObject) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:161
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isClassAvailable(const char *)
func (this *QAndroidJniObject) IsClassAvailable(className string) bool {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject16isClassAvailableEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QAndroidJniObject_IsClassAvailable(className string) bool {
	var nilthis *QAndroidJniObject
	rv := nilthis.IsClassAvailable(className)
	return rv
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QAndroidJniObject) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjniobject.h:164
// index:0
// Public static Visibility=Default Availability=Available
// [1] QAndroidJniObject fromLocalRef(int)
func (this *QAndroidJniObject) FromLocalRef(obj int) *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject12fromLocalRefEi", qtrt.FFI_TYPE_POINTER, obj)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_FromLocalRef(obj int) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.FromLocalRef(obj)
	return rv
}

//  body block end
