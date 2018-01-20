//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:754
// index:0
// void QToolBarChangeEvent(_Bool)
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QToolBarChangeEventC2Eb", ffiqt.FFI_TYPE_VOID, cthis, &t)
	gopp.ErrPrint(err, rv)
	gothis := NewQToolBarChangeEventFromPointer(cthis)
	return gothis
}
func NewQToolBarChangeEventFromPointer(cthis unsafe.Pointer) *QToolBarChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QToolBarChangeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:755
// index:0
// virtual
// void ~QToolBarChangeEvent()
func DeleteQToolBarChangeEvent(*QToolBarChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QToolBarChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:757
// index:0
// inline
// bool toggle()
func (this *QToolBarChangeEvent) Toggle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QToolBarChangeEvent6toggleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
