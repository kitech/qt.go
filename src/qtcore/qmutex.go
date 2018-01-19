//  header block begin
// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>
package qtcore

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmutex.h:130
// index:0
// void QMutex(enum QMutex::RecursionMode)
func NewQMutex(mode int) *QMutex {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutexC2ENS_13RecursionModeE", ffiqt.FFI_TYPE_VOID, cthis, &mode)
	gopp.ErrPrint(err, rv)
	return &QMutex{cthis}
}

// /usr/include/qt/QtCore/qmutex.h:131
// index:0
// void ~QMutex()
func DeleteQMutex(*QMutex) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutexD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:134
// index:0
// void lock()
func (this *QMutex) Lock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex4lockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:135
// index:0
// bool tryLock(int)
func (this *QMutex) TryLock(timeout int) {
	// 0: (, int timeout), (&timeout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex7tryLockEi", ffiqt.FFI_TYPE_VOID, this.cthis, &timeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:137
// index:0
// void unlock()
func (this *QMutex) Unlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex6unlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:140
// index:0
// inline
// bool try_lock()
func (this *QMutex) Try_lock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QMutex8try_lockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:161
// index:0
// inline
// bool isRecursive()
func (this *QMutex) IsRecursive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QMutex11isRecursiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
