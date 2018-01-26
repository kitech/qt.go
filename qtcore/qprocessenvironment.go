package qtcore

// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QProcessEnvironment) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQProcessEnvironmentFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return &QProcessEnvironment{&qtrt.CObject{cthis}}
}
func (*QProcessEnvironment) NewFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return NewQProcessEnvironmentFromPointer(cthis)
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
func (this *QProcessEnvironment) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func (this *QProcessEnvironment) Contains(name *QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func (this *QProcessEnvironment) Value(name *QString, defaultValue *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	var convArg1 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:98
// index:0
// Public static
// QProcessEnvironment systemEnvironment()
func (this *QProcessEnvironment) SystemEnvironment() *QProcessEnvironment /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment17systemEnvironmentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QProcessEnvironment_SystemEnvironment() *QProcessEnvironment /*123*/ {
	var nilthis *QProcessEnvironment
	rv := nilthis.SystemEnvironment()
	return rv
}

//  body block end
