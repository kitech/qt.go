// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QProcessEnvironment struct {
	*qtrt.CObject
}
type QProcessEnvironment_ITF interface {
	QProcessEnvironment_PTR() *QProcessEnvironment
}

func (ptr *QProcessEnvironment) QProcessEnvironment_PTR() *QProcessEnvironment { return ptr }

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

/*

 */
func (*QProcessEnvironment) NewForInherit() *QProcessEnvironment {
	return NewQProcessEnvironment()
}
func NewQProcessEnvironment() *QProcessEnvironment {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQProcessEnvironment)
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QProcessEnvironment()

/*

 */
func DeleteQProcessEnvironment(this *QProcessEnvironment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qprocess.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QProcessEnvironment & operator=(QProcessEnvironment &&)

/*

 */
func (this *QProcessEnvironment) Operator_equal(other unsafe.Pointer /*333*/) *QProcessEnvironment {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQProcessEnvironment)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:76
// index:1
// Public Visibility=Default Availability=Available
// [8] QProcessEnvironment & operator=(const QProcessEnvironment &)

/*

 */
func (this *QProcessEnvironment) Operator_equal1(other QProcessEnvironment_ITF) *QProcessEnvironment {
	var convArg0 unsafe.Pointer
	if other != nil && other.QProcessEnvironment_PTR() != nil {
		convArg0 = other.QProcessEnvironment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironmentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQProcessEnvironment)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QProcessEnvironment &)

/*

 */
func (this *QProcessEnvironment) Swap(other QProcessEnvironment_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QProcessEnvironment_PTR() != nil {
		convArg0 = other.QProcessEnvironment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QProcessEnvironment &) const

/*

 */
func (this *QProcessEnvironment) Operator_equal_equal(other QProcessEnvironment_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QProcessEnvironment_PTR() != nil {
		convArg0 = other.QProcessEnvironment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironmenteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QProcessEnvironment &) const

/*

 */
func (this *QProcessEnvironment) Operator_not_equal(other QProcessEnvironment_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QProcessEnvironment_PTR() != nil {
		convArg0 = other.QProcessEnvironment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironmentneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*

 */
func (this *QProcessEnvironment) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*

 */
func (this *QProcessEnvironment) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &) const

/*

 */
func (this *QProcessEnvironment) Contains(name string) bool {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)

/*

 */
func (this *QProcessEnvironment) Remove(name string) {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &, const QString &) const

/*

 */
func (this *QProcessEnvironment) Value(name string, defaultValue string) string {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(defaultValue)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qprocess.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &, const QString &) const

/*

 */
func (this *QProcessEnvironment) Valuep(name string) string {
	var tmpArg0 = NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment5valueERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qprocess.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList toStringList() const

/*

 */
func (this *QProcessEnvironment) ToStringList() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment12toStringListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList keys() const

/*

 */
func (this *QProcessEnvironment) Keys() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QProcessEnvironment4keysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QProcessEnvironment systemEnvironment()

/*
Returns the environment of the calling process as a list of key=value pairs. Example:


  QStringList environment = QProcess::systemEnvironment();
  // environment = {"PATH=/usr/bin:/usr/local/bin",
  //                "USER=greg", "HOME=/home/greg"}



This function does not cache the system environment. Therefore, it's possible to obtain an updated version of the environment if low-level C library functions like setenv or putenv have been called.

However, note that repeated calls to this function will recreate the list of environment variables, which is a non-trivial operation.

Note: For new code, it is recommended to use QProcessEnvironment::systemEnvironment()

This function was introduced in  Qt 4.1.

See also QProcessEnvironment::systemEnvironment() and setProcessEnvironment().
*/
func (this *QProcessEnvironment) SystemEnvironment() *QProcessEnvironment /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QProcessEnvironment17systemEnvironmentEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

//  keep block begin

func init_unused_10501() {
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
