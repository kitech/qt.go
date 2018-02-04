package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

type QTimerEvent struct {
	*QEvent
}

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
func NewQTimerEvent(timerId int) *QTimerEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTimerEventC2Ei", qtrt.FFI_TYPE_POINTER, timerId)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimerEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTimerEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:341
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTimerEvent()
func DeleteQTimerEvent(this *QTimerEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTimerEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:342
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int timerId()
func (this *QTimerEvent) TimerId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTimerEvent7timerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
