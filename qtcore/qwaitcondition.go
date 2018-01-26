package qtcore

// /usr/include/qt/QtCore/qwaitcondition.h
// #include <qwaitcondition.h>
// #include <QtCore>

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
type QWaitCondition struct {
	*qtrt.CObject
}

func (this *QWaitCondition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWaitCondition) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQWaitConditionFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return &QWaitCondition{&qtrt.CObject{cthis}}
}
func (*QWaitCondition) NewFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return NewQWaitConditionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qwaitcondition.h:59
// index:0
// Public
// void QWaitCondition()
func NewQWaitCondition() *QWaitCondition {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitConditionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQWaitConditionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qwaitcondition.h:60
// index:0
// Public
// void ~QWaitCondition()
func DeleteQWaitCondition(*QWaitCondition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitConditionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:62
// index:0
// Public
// bool wait(class QMutex *, unsigned long)
func (this *QWaitCondition) Wait(lockedMutex *QMutex /*777 QMutex **/, time uint) bool {
	var convArg0 = lockedMutex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP6QMutexm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:63
// index:1
// Public
// bool wait(class QReadWriteLock *, unsigned long)
func (this *QWaitCondition) Wait_1(lockedReadWriteLock *QReadWriteLock /*777 QReadWriteLock **/, time uint) bool {
	var convArg0 = lockedReadWriteLock.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP14QReadWriteLockm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:65
// index:0
// Public
// void wakeOne()
func (this *QWaitCondition) WakeOne() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition7wakeOneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:66
// index:0
// Public
// void wakeAll()
func (this *QWaitCondition) WakeAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition7wakeAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:68
// index:0
// Public inline
// void notify_one()
func (this *QWaitCondition) Notify_one() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition10notify_oneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:69
// index:0
// Public inline
// void notify_all()
func (this *QWaitCondition) Notify_all() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QWaitCondition10notify_allEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
