package qtcore

// /usr/include/qt/QtCore/qsystemsemaphore.h
// #include <qsystemsemaphore.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSystemSemaphore struct {
	*qtrt.CObject
}
type QSystemSemaphore_ITF interface {
	QSystemSemaphore_PTR() *QSystemSemaphore
}

func (ptr *QSystemSemaphore) QSystemSemaphore_PTR() *QSystemSemaphore { return ptr }

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
func NewQSystemSemaphore(key string, initialValue int, mode int) *QSystemSemaphore {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, enum QSystemSemaphore::AccessMode)
func NewQSystemSemaphore__(key string) *QSystemSemaphore {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	initialValue := int(0)
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemSemaphore(const QString &, int, enum QSystemSemaphore::AccessMode)
func NewQSystemSemaphore__1(key string, initialValue int) *QSystemSemaphore {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSystemSemaphore)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSystemSemaphore()
func DeleteQSystemSemaphore(this *QSystemSemaphore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey(key string, initialValue int, mode int) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey__(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	initialValue := int(0)
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey__1(key string, initialValue int) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QSystemSemaphore::AccessMode=Enum, QSystemSemaphore::AccessMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, initialValue, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key() const
func (this *QSystemSemaphore) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acquire()
func (this *QSystemSemaphore) Acquire() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7acquireEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool release(int)
func (this *QSystemSemaphore) Release(n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool release(int)
func (this *QSystemSemaphore) Release__() bool {
	// arg: 0, int=Int, =Invalid,
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSystemSemaphore::SystemSemaphoreError error() const
func (this *QSystemSemaphore) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QSystemSemaphore) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSystemSemaphore11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
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
