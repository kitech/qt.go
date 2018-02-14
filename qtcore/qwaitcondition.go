package qtcore

// /usr/include/qt/QtCore/qwaitcondition.h
// #include <qwaitcondition.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QWaitCondition struct {
	*qtrt.CObject
}
type QWaitCondition_ITF interface {
	QWaitCondition_PTR() *QWaitCondition
}

func (ptr *QWaitCondition) QWaitCondition_PTR() *QWaitCondition { return ptr }

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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qwaitcondition.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(QMutex *, unsigned long)
func (this *QWaitCondition) Wait(lockedMutex QMutex_ITF /*777 QMutex **/, time uint) bool {
	var convArg0 unsafe.Pointer
	if lockedMutex != nil && lockedMutex.QMutex_PTR() != nil {
		convArg0 = lockedMutex.QMutex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP6QMutexm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:63
// index:1
// Public Visibility=Default Availability=Available
// [1] bool wait(QReadWriteLock *, unsigned long)
func (this *QWaitCondition) Wait_1(lockedReadWriteLock QReadWriteLock_ITF /*777 QReadWriteLock **/, time uint) bool {
	var convArg0 unsafe.Pointer
	if lockedReadWriteLock != nil && lockedReadWriteLock.QReadWriteLock_PTR() != nil {
		convArg0 = lockedReadWriteLock.QReadWriteLock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition4waitEP14QReadWriteLockm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qwaitcondition.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeOne()
func (this *QWaitCondition) WakeOne() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition7wakeOneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeAll()
func (this *QWaitCondition) WakeAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition7wakeAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void notify_one()
func (this *QWaitCondition) Notify_one() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition10notify_oneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qwaitcondition.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void notify_all()
func (this *QWaitCondition) Notify_all() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWaitCondition10notify_allEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

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
