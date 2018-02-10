package qtandroidextras

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjnienvironment.h
// #include <qandroidjnienvironment.h>
// #include <Qtjni>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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

type QAndroidJniEnvironment struct {
	*qtrt.CObject
}

func (this *QAndroidJniEnvironment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidJniEnvironment) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidJniEnvironmentFromPointer(cthis unsafe.Pointer) *QAndroidJniEnvironment {
	return &QAndroidJniEnvironment{&qtrt.CObject{cthis}}
}
func (*QAndroidJniEnvironment) NewFromPointer(cthis unsafe.Pointer) *QAndroidJniEnvironment {
	return NewQAndroidJniEnvironmentFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjnienvironment.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniEnvironment()
func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironmentC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniEnvironmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniEnvironment)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjnienvironment.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniEnvironment()
func DeleteQAndroidJniEnvironment(this *QAndroidJniEnvironment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/jni/qandroidjnienvironment.h:57
// index:0
// Public static Visibility=Default Availability=Available
// [8] int * javaVM()
func (this *QAndroidJniEnvironment) JavaVM() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironment6javaVMEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QAndroidJniEnvironment_JavaVM() unsafe.Pointer /*666*/ {
	var nilthis *QAndroidJniEnvironment
	rv := nilthis.JavaVM()
	return rv
}

//  body block end
