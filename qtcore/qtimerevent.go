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
// Public
// void QTimerEvent(int)
func NewQTimerEvent(timerId int) *QTimerEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTimerEventC2Ei", ffiqt.FFI_TYPE_VOID, cthis, timerId)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimerEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:341
// index:0
// Public virtual
// void ~QTimerEvent()
func DeleteQTimerEvent(*QTimerEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTimerEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:342
// index:0
// Public inline
// int timerId()
func (this *QTimerEvent) TimerId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTimerEvent7timerIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
