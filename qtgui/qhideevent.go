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
// extern C begin: 2
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
type QHideEvent struct {
	*qtcore.QEvent
}
type QHideEvent_ITF interface {
	qtcore.QEvent_ITF
	QHideEvent_PTR() *QHideEvent
}

func (ptr *QHideEvent) QHideEvent_PTR() *QHideEvent { return ptr }

func (this *QHideEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QHideEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQHideEventFromPointer(cthis unsafe.Pointer) *QHideEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QHideEvent{bcthis0}
}
func (*QHideEvent) NewFromPointer(cthis unsafe.Pointer) *QHideEvent {
	return NewQHideEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:501
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHideEvent()

/*

 */
func NewQHideEvent() *QHideEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QHideEventC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHideEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHideEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:502
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHideEvent()

/*

 */
func DeleteQHideEvent(this *QHideEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QHideEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
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
