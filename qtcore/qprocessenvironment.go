package qtcore

// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QProcessEnvironment struct {
	*qtrt.CObject
}

func (this *QProcessEnvironment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QProcessEnvironment) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQProcessEnvironmentFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return &QProcessEnvironment{&qtrt.CObject{cthis}}
}
func (*QProcessEnvironment) NewFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return NewQProcessEnvironmentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qprocess.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProcessEnvironment()
func NewQProcessEnvironment() *QProcessEnvironment {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQProcessEnvironment)
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QProcessEnvironment()
func DeleteQProcessEnvironment(this *QProcessEnvironment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qprocess.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QProcessEnvironment &)
func (this *QProcessEnvironment) Swap(other *QProcessEnvironment) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QProcessEnvironment) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QProcessEnvironment) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &)
func (this *QProcessEnvironment) Contains(name *QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)
func (this *QProcessEnvironment) Remove(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &, const QString &)
func (this *QProcessEnvironment) Value(name *QString, defaultValue *QString) *QString /*123*/ {
	var convArg0 = name.GetCthis()
	var convArg1 = defaultValue.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QProcessEnvironment systemEnvironment()
func (this *QProcessEnvironment) SystemEnvironment() *QProcessEnvironment /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment17systemEnvironmentEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQProcessEnvironment)
	return rv2
}
func QProcessEnvironment_SystemEnvironment() *QProcessEnvironment /*123*/ {
	var nilthis *QProcessEnvironment
	rv := nilthis.SystemEnvironment()
	return rv
}

//  body block end
