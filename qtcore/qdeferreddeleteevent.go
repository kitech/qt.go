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
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDeferredDeleteEvent struct {
	*QEvent
}
type QDeferredDeleteEvent_ITF interface {
	QEvent_ITF
	QDeferredDeleteEvent_PTR() *QDeferredDeleteEvent
}

func (ptr *QDeferredDeleteEvent) QDeferredDeleteEvent_PTR() *QDeferredDeleteEvent { return ptr }

func (this *QDeferredDeleteEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDeferredDeleteEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = NewQEventFromPointer(cthis)
}
func NewQDeferredDeleteEventFromPointer(cthis unsafe.Pointer) *QDeferredDeleteEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QDeferredDeleteEvent{bcthis0}
}
func (*QDeferredDeleteEvent) NewFromPointer(cthis unsafe.Pointer) *QDeferredDeleteEvent {
	return NewQDeferredDeleteEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:377
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDeferredDeleteEvent()
func NewQDeferredDeleteEvent() *QDeferredDeleteEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QDeferredDeleteEventC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeferredDeleteEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeferredDeleteEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:378
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDeferredDeleteEvent()
func DeleteQDeferredDeleteEvent(this *QDeferredDeleteEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QDeferredDeleteEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:379
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int loopLevel() const
func (this *QDeferredDeleteEvent) LoopLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QDeferredDeleteEvent9loopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
