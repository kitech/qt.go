package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QToolBarChangeEvent struct {
	*qtcore.QEvent
}

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

// /usr/include/qt/QtGui/qevent.h:754
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QToolBarChangeEvent(_Bool)
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QToolBarChangeEventC2Eb", ffiqt.FFI_TYPE_POINTER, t)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQToolBarChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:755
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QToolBarChangeEvent()
func DeleteQToolBarChangeEvent(this *QToolBarChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QToolBarChangeEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:757
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool toggle()
func (this *QToolBarChangeEvent) Toggle() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QToolBarChangeEvent6toggleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
