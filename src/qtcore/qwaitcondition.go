//  header block begin
// /usr/include/qt/QtCore/qwaitcondition.h
// #include <qwaitcondition.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QWaitCondition struct {
	*qtrt.CObject
}

func (this *QWaitCondition) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qwaitcondition.h:59
// index:0
// void QWaitCondition()
func NewQWaitCondition() *QWaitCondition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitConditionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQWaitConditionFromPointer(cthis)
	return gothis
}
func NewQWaitConditionFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return &QWaitCondition{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qwaitcondition.h:60
// index:0
// void ~QWaitCondition()
func DeleteQWaitCondition(*QWaitCondition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitConditionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:62
// index:0
// bool wait(class QMutex *, unsigned long)
func (this *QWaitCondition) Wait(lockedMutex unsafe.Pointer, time uint) {
	// 0: (, lockedMutex QMutex *, time unsigned long), (lockedMutex, &time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP6QMutexm", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lockedMutex, &time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:63
// index:1
// bool wait(class QReadWriteLock *, unsigned long)
func (this *QWaitCondition) Wait_1(lockedReadWriteLock unsafe.Pointer, time uint) {
	// 1: (, lockedReadWriteLock QReadWriteLock *, time unsigned long), (lockedReadWriteLock, &time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP14QReadWriteLockm", ffiqt.FFI_TYPE_VOID, this.GetCthis(), lockedReadWriteLock, &time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:65
// index:0
// void wakeOne()
func (this *QWaitCondition) WakeOne() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition7wakeOneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:66
// index:0
// void wakeAll()
func (this *QWaitCondition) WakeAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition7wakeAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:68
// index:0
// inline
// void notify_one()
func (this *QWaitCondition) Notify_one() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition10notify_oneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:69
// index:0
// inline
// void notify_all()
func (this *QWaitCondition) Notify_all() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition10notify_allEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
