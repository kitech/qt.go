package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QApplicationStateChangeEvent struct {
	*qtcore.QEvent
}
type QApplicationStateChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QApplicationStateChangeEvent_PTR() *QApplicationStateChangeEvent
}

func (ptr *QApplicationStateChangeEvent) QApplicationStateChangeEvent_PTR() *QApplicationStateChangeEvent {
	return ptr
}

func (this *QApplicationStateChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QApplicationStateChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQApplicationStateChangeEventFromPointer(cthis unsafe.Pointer) *QApplicationStateChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QApplicationStateChangeEvent{bcthis0}
}
func (*QApplicationStateChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QApplicationStateChangeEvent {
	return NewQApplicationStateChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:1052
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplicationStateChangeEvent(Qt::ApplicationState)

/*

 */
func NewQApplicationStateChangeEvent(state int) *QApplicationStateChangeEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QApplicationStateChangeEventC2EN2Qt16ApplicationStateE", qtrt.FFI_TYPE_POINTER, state)
	qtrt.ErrPrint(err, rv)
	gothis := NewQApplicationStateChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQApplicationStateChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1053
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ApplicationState applicationState() const

/*

 */
func (this *QApplicationStateChangeEvent) ApplicationState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QApplicationStateChangeEvent16applicationStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQApplicationStateChangeEvent(this *QApplicationStateChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QApplicationStateChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
