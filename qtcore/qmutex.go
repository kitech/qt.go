package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QMutex struct {
	*QBasicMutex
}

func (this *QMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBasicMutex.GetCthis()
	}
}
func (this *QMutex) SetCthis(cthis unsafe.Pointer) {
	this.QBasicMutex = NewQBasicMutexFromPointer(cthis)
}
func NewQMutexFromPointer(cthis unsafe.Pointer) *QMutex {
	bcthis0 := NewQBasicMutexFromPointer(cthis)
	return &QMutex{bcthis0}
}
func (*QMutex) NewFromPointer(cthis unsafe.Pointer) *QMutex {
	return NewQMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMutex(enum QMutex::RecursionMode)
func NewQMutex(mode int) *QMutex {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMutex()
func DeleteQMutex(this *QMutex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmutex.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lock()
func (this *QMutex) Lock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)
func (this *QMutex) TryLock(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QMutex) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool try_lock()
func (this *QMutex) Try_lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QMutex8try_lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRecursive()
func (this *QMutex) IsRecursive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

type QMutex__RecursionMode = int

const QMutex__NonRecursive QMutex__RecursionMode = 0
const QMutex__Recursive QMutex__RecursionMode = 1

//  body block end
