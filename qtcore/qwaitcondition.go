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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWaitConditionFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return &QWaitCondition{&qtrt.CObject{cthis}}
}
func (*QWaitCondition) NewFromPointer(cthis unsafe.Pointer) *QWaitCondition {
	return NewQWaitConditionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qwaitcondition.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWaitCondition()
func NewQWaitCondition() *QWaitCondition {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitConditionC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQWaitConditionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWaitCondition)
	return gothis
}

// /usr/include/qt/QtCore/qwaitcondition.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWaitCondition()
func DeleteQWaitCondition(this *QWaitCondition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitConditionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qwaitcondition.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(QMutex *, unsigned long)
func (this *QWaitCondition) Wait(lockedMutex *QMutex /*777 QMutex **/, time uint) bool {
	var convArg0 = lockedMutex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP6QMutexm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:63
// index:1
// Public Visibility=Default Availability=Available
// [1] bool wait(QReadWriteLock *, unsigned long)
func (this *QWaitCondition) Wait_1(lockedReadWriteLock *QReadWriteLock /*777 QReadWriteLock **/, time uint) bool {
	var convArg0 = lockedReadWriteLock.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP14QReadWriteLockm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeOne()
func (this *QWaitCondition) WakeOne() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition7wakeOneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeAll()
func (this *QWaitCondition) WakeAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition7wakeAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void notify_one()
func (this *QWaitCondition) Notify_one() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition10notify_oneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void notify_all()
func (this *QWaitCondition) Notify_all() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition10notify_allEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
