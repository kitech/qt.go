package qtcore

// /usr/include/qt/QtCore/qsystemsemaphore.h
// #include <qsystemsemaphore.h>
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
import "gopp"
import "qt.go/cffiqt"
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

type QSystemSemaphore struct {
	*qtrt.CObject
}

func (this *QSystemSemaphore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSystemSemaphore) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSystemSemaphoreFromPointer(cthis unsafe.Pointer) *QSystemSemaphore {
	return &QSystemSemaphore{&qtrt.CObject{cthis}}
}
func (*QSystemSemaphore) NewFromPointer(cthis unsafe.Pointer) *QSystemSemaphore {
	return NewQSystemSemaphoreFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, enum QSystemSemaphore::AccessMode)
func NewQSystemSemaphore(key *QString, initialValue int, mode int) *QSystemSemaphore {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSystemSemaphore()
func DeleteQSystemSemaphore(this *QSystemSemaphore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey(key *QString, initialValue int, mode int) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key()
func (this *QSystemSemaphore) Key() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acquire()
func (this *QSystemSemaphore) Acquire() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7acquireEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool release(int)
func (this *QSystemSemaphore) Release(n int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), n)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSystemSemaphore::SystemSemaphoreError error()
func (this *QSystemSemaphore) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QSystemSemaphore) ErrorString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

type QSystemSemaphore__AccessMode = int

const QSystemSemaphore__Open QSystemSemaphore__AccessMode = 0
const QSystemSemaphore__Create QSystemSemaphore__AccessMode = 1

type QSystemSemaphore__SystemSemaphoreError = int

const QSystemSemaphore__NoError QSystemSemaphore__SystemSemaphoreError = 0
const QSystemSemaphore__PermissionDenied QSystemSemaphore__SystemSemaphoreError = 1
const QSystemSemaphore__KeyError QSystemSemaphore__SystemSemaphoreError = 2
const QSystemSemaphore__AlreadyExists QSystemSemaphore__SystemSemaphoreError = 3
const QSystemSemaphore__NotFound QSystemSemaphore__SystemSemaphoreError = 4
const QSystemSemaphore__OutOfResources QSystemSemaphore__SystemSemaphoreError = 5
const QSystemSemaphore__UnknownError QSystemSemaphore__SystemSemaphoreError = 6

//  body block end
