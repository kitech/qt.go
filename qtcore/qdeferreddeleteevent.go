package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QDeferredDeleteEvent struct {
	*QEvent
}

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
// Public
// void QDeferredDeleteEvent()
func NewQDeferredDeleteEvent() *QDeferredDeleteEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDeferredDeleteEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDeferredDeleteEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:378
// index:0
// Public virtual
// void ~QDeferredDeleteEvent()
func DeleteQDeferredDeleteEvent(*QDeferredDeleteEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDeferredDeleteEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:379
// index:0
// Public inline
// int loopLevel()
func (this *QDeferredDeleteEvent) LoopLevel() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QDeferredDeleteEvent9loopLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
