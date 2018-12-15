package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h
// #include <qandroidjnienvironment.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAndroidJniEnvironment struct {
	*qtrt.CObject
}
type QAndroidJniEnvironment_ITF interface {
	QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment
}

func (ptr *QAndroidJniEnvironment) QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment { return ptr }

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

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniEnvironment()

/*
Constructs a new QAndroidJniEnvironment object and attach the current thread to the Java VM.


  bool exceptionCheck()
  {
      /-*
        The QAndroidJniEnvironment attaches the current thread to the JavaVM on
        creation and detach when it goes out of scope.
       *-/
      QAndroidJniEnvironment qjniEnv;
      return qjniEnv->ExceptionCheck();
  }
*/
func (*QAndroidJniEnvironment) NewForInherit() *QAndroidJniEnvironment {
	return NewQAndroidJniEnvironment()
}
func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironmentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniEnvironmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniEnvironment)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniEnvironment()

/*

 */
func DeleteQAndroidJniEnvironment(this *QAndroidJniEnvironment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:57
// index:0
// Public static Visibility=Default Availability=Available
// [8] JavaVM * javaVM()

/*
Returns the Java VM interface.
*/
func (this *QAndroidJniEnvironment) JavaVM() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironment6javaVMEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QAndroidJniEnvironment_JavaVM() unsafe.Pointer /*666*/ {
	var nilthis *QAndroidJniEnvironment
	rv := nilthis.JavaVM()
	return rv
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:58
// index:0
// Public Visibility=Default Availability=Available
// [8] JNIEnv * operator->()

/*

 */
func (this *QAndroidJniEnvironment) Operator_minus_greater() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironmentptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] jclass findClass(const char *)

/*
Searches for className using all available class loaders. Qt on Android uses a custom class loader to load all the .jar files and it must be used to find any classes that are created by that class loader because these classes are not visible in the default class loader.

Returns the class pointer or null if is not found.

This function was introduced in  Qt 5.12.
*/
func (this *QAndroidJniEnvironment) FindClass(className string) unsafe.Pointer /*666*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAndroidJniEnvironment9findClassEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
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
}

//  keep block end
