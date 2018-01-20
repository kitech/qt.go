//  header block begin
// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QProcessEnvironment struct {
	*qtrt.CObject
}

func (this *QProcessEnvironment) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQProcessEnvironmentFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return &QProcessEnvironment{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qprocess.h:70
// index:0
// Public
// void QProcessEnvironment()
func NewQProcessEnvironment() *QProcessEnvironment {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQProcessEnvironmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:72
// index:0
// Public
// void ~QProcessEnvironment()
func DeleteQProcessEnvironment(*QProcessEnvironment) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironmentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:78
// index:0
// Public inline
// void swap(class QProcessEnvironment &)
func (this *QProcessEnvironment) Swap(other *QProcessEnvironment) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:84
// index:0
// Public
// bool isEmpty()
func (this *QProcessEnvironment) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:85
// index:0
// Public
// void clear()
func (this *QProcessEnvironment) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:87
// index:0
// Public
// bool contains(const class QString &)
func (this *QProcessEnvironment) Contains(name *QString) interface{} {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:89
// index:0
// Public
// void remove(const class QString &)
func (this *QProcessEnvironment) Remove(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:90
// index:0
// Public
// QString value(const class QString &, const class QString &)
func (this *QProcessEnvironment) Value(name *QString, defaultValue *QString) interface{} {
	var convArg0 = name.GetCthis()
	var convArg1 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:92
// index:0
// Public
// QStringList toStringList()
func (this *QProcessEnvironment) ToStringList() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment12toStringListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:94
// index:0
// Public
// QStringList keys()
func (this *QProcessEnvironment) Keys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment4keysEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:98
// index:0
// Public static
// QProcessEnvironment systemEnvironment()
func (this *QProcessEnvironment) SystemEnvironment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment17systemEnvironmentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcessEnvironment_SystemEnvironment() {
	var nilthis *QProcessEnvironment
	nilthis.SystemEnvironment()
}

//  body block end
