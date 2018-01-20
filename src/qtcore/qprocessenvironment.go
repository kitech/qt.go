//  header block begin
// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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

// /usr/include/qt/QtCore/qprocess.h:70
// index:0
// void QProcessEnvironment()
func NewQProcessEnvironment() *QProcessEnvironment {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQProcessEnvironmentFromPointer(cthis)
	return gothis
}
func NewQProcessEnvironmentFromPointer(cthis unsafe.Pointer) *QProcessEnvironment {
	return &QProcessEnvironment{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qprocess.h:72
// index:0
// void ~QProcessEnvironment()
func DeleteQProcessEnvironment(*QProcessEnvironment) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironmentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:78
// index:0
// inline
// void swap(class QProcessEnvironment &)
func (this *QProcessEnvironment) Swap(other unsafe.Pointer) {
	// 0: (, other QProcessEnvironment &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:84
// index:0
// bool isEmpty()
func (this *QProcessEnvironment) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:85
// index:0
// void clear()
func (this *QProcessEnvironment) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:87
// index:0
// bool contains(const class QString &)
func (this *QProcessEnvironment) Contains(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment8containsERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:89
// index:0
// void remove(const class QString &)
func (this *QProcessEnvironment) Remove(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment6removeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:90
// index:0
// QString value(const class QString &, const class QString &)
func (this *QProcessEnvironment) Value(name unsafe.Pointer, defaultValue unsafe.Pointer) {
	// 0: (, name const QString &, defaultValue const QString &), (name, defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name, defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:92
// index:0
// QStringList toStringList()
func (this *QProcessEnvironment) ToStringList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment12toStringListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:94
// index:0
// QStringList keys()
func (this *QProcessEnvironment) Keys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QProcessEnvironment4keysEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:98
// index:0
// static
// QProcessEnvironment systemEnvironment()
func (this *QProcessEnvironment) SystemEnvironment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QProcessEnvironment17systemEnvironmentEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcessEnvironment_SystemEnvironment() {
	// 0: (), ()
	var nilthis *QProcessEnvironment
	nilthis.SystemEnvironment()
}

//  body block end
