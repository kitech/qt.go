package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h
// #include <qwinevent.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QWinEvent struct {
	*qtcore.QEvent
}
type QWinEvent_ITF interface {
	qtcore.QEvent_ITF
	QWinEvent_PTR() *QWinEvent
}

func (ptr *QWinEvent) QWinEvent_PTR() *QWinEvent { return ptr }

func (this *QWinEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QWinEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQWinEventFromPointer(cthis unsafe.Pointer) *QWinEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QWinEvent{bcthis0}
}
func (*QWinEvent) NewFromPointer(cthis unsafe.Pointer) *QWinEvent {
	return NewQWinEventFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinEvent(int)
func NewQWinEvent(type_ int) *QWinEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QWinEventC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinEvent)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinEvent()
func DeleteQWinEvent(this *QWinEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QWinEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
