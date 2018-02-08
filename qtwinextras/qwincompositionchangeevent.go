package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h
// #include <qwinevent.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QWinCompositionChangeEvent struct {
	*QWinEvent
}

func (this *QWinCompositionChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWinEvent.GetCthis()
	}
}
func (this *QWinCompositionChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QWinEvent = NewQWinEventFromPointer(cthis)
}
func NewQWinCompositionChangeEventFromPointer(cthis unsafe.Pointer) *QWinCompositionChangeEvent {
	bcthis0 := NewQWinEventFromPointer(cthis)
	return &QWinCompositionChangeEvent{bcthis0}
}
func (*QWinCompositionChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QWinCompositionChangeEvent {
	return NewQWinCompositionChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinCompositionChangeEvent(_Bool)
func NewQWinCompositionChangeEvent(enabled bool) *QWinCompositionChangeEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWinCompositionChangeEventC2Eb", qtrt.FFI_TYPE_POINTER, enabled)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinCompositionChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinCompositionChangeEvent)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinCompositionChangeEvent()
func DeleteQWinCompositionChangeEvent(this *QWinCompositionChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWinCompositionChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCompositionEnabled()
func (this *QWinCompositionChangeEvent) IsCompositionEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWinCompositionChangeEvent20isCompositionEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
