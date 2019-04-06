package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QTimerEvent struct {
	*QEvent
}
type QTimerEvent_ITF interface {
	QEvent_ITF
	QTimerEvent_PTR() *QTimerEvent
}

func (ptr *QTimerEvent) QTimerEvent_PTR() *QTimerEvent { return ptr }

func (this *QTimerEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QTimerEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = NewQEventFromPointer(cthis)
}
func NewQTimerEventFromPointer(cthis unsafe.Pointer) *QTimerEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QTimerEvent{bcthis0}
}
func (*QTimerEvent) NewFromPointer(cthis unsafe.Pointer) *QTimerEvent {
	return NewQTimerEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimerEvent(int)

/*

 */
func (*QTimerEvent) NewForInherit(timerId int) *QTimerEvent {
	return NewQTimerEvent(timerId)
}
func NewQTimerEvent(timerId int) *QTimerEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTimerEventC2Ei", qtrt.FFI_TYPE_POINTER, timerId)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimerEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimerEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:341
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTimerEvent()

/*

 */
func DeleteQTimerEvent(this *QTimerEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTimerEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:342
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int timerId() const

/*

 */
func (this *QTimerEvent) TimerId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTimerEvent7timerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_10371() {
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
