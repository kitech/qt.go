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
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QInputEvent struct {
	*qtcore.QEvent
}

func (this *QInputEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QInputEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQInputEventFromPointer(cthis unsafe.Pointer) *QInputEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputEvent{bcthis0}
}
func (*QInputEvent) NewFromPointer(cthis unsafe.Pointer) *QInputEvent {
	return NewQInputEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputEvent(enum QEvent::Type, Qt::KeyboardModifiers)
func NewQInputEvent(type_ int, modifiers int) *QInputEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QInputEventC2EN6QEvent4TypeE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, type_, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQInputEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQInputEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputEvent()
func DeleteQInputEvent(this *QInputEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QInputEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QInputEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QInputEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QInputEvent) SetModifiers(amodifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QInputEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), amodifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [8] ulong timestamp()
func (this *QInputEvent) Timestamp() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QInputEvent9timestampEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTimestamp(ulong)
func (this *QInputEvent) SetTimestamp(atimestamp uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QInputEvent12setTimestampEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), atimestamp)
	qtrt.ErrPrint(err, rv)
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
}

//  keep block end
