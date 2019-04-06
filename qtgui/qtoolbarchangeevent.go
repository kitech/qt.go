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
// extern C begin: 6
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
type QToolBarChangeEvent struct {
	*qtcore.QEvent
}
type QToolBarChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QToolBarChangeEvent_PTR() *QToolBarChangeEvent
}

func (ptr *QToolBarChangeEvent) QToolBarChangeEvent_PTR() *QToolBarChangeEvent { return ptr }

func (this *QToolBarChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QToolBarChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQToolBarChangeEventFromPointer(cthis unsafe.Pointer) *QToolBarChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QToolBarChangeEvent{bcthis0}
}
func (*QToolBarChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QToolBarChangeEvent {
	return NewQToolBarChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:759
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBarChangeEvent(bool)

/*

 */
func (*QToolBarChangeEvent) NewForInherit(t bool) *QToolBarChangeEvent {
	return NewQToolBarChangeEvent(t)
}
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QToolBarChangeEventC2Eb", qtrt.FFI_TYPE_POINTER, t)
	qtrt.ErrPrint(err, rv)
	gothis := NewQToolBarChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQToolBarChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:760
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolBarChangeEvent()

/*

 */
func DeleteQToolBarChangeEvent(this *QToolBarChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QToolBarChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:762
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toggle() const

/*

 */
func (this *QToolBarChangeEvent) Toggle() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QToolBarChangeEvent6toggleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10703() {
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
