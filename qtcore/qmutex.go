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
// Public
// void QMutex(enum QMutex::RecursionMode)
func NewQMutex(mode int) *QMutex {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", ffiqt.FFI_TYPE_VOID, cthis, mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQMutexFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:131
// index:0
// Public
// void ~QMutex()
func DeleteQMutex(*QMutex) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutexD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:134
// index:0
// Public
// void lock()
func (this *QMutex) Lock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex4lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// Public
// bool tryLock(int)
func (this *QMutex) TryLock(timeout int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:137
// index:0
// Public
// void unlock()
func (this *QMutex) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:140
// index:0
// Public inline
// bool try_lock()
func (this *QMutex) Try_lock() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex8try_lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:161
// index:0
// Public inline
// bool isRecursive()
func (this *QMutex) IsRecursive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMutex11isRecursiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QMutex__RecursionMode = int

const QMutex__NonRecursive QMutex__RecursionMode = 0
const QMutex__Recursive QMutex__RecursionMode = 1

//  body block end
