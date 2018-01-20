//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QInputEvent struct {
	*qtcore.QEvent
}

func (this *QInputEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:71
// index:0
// void QInputEvent(enum QEvent::Type, Qt::KeyboardModifiers)
func NewQInputEvent(type_ int, modifiers int) *QInputEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEventC2EN6QEvent4TypeE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, cthis, &type_, &modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputEventFromPointer(cthis)
	return gothis
}
func NewQInputEventFromPointer(cthis unsafe.Pointer) *QInputEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:72
// index:0
// virtual
// void ~QInputEvent()
func DeleteQInputEvent(*QInputEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:73
// index:0
// inline
// Qt::KeyboardModifiers modifiers()
func (this *QInputEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:74
// index:0
// inline
// void setModifiers(Qt::KeyboardModifiers)
func (this *QInputEvent) SetModifiers(amodifiers int) {
	// 0: (, QFlags<Qt::KeyboardModifier> amodifiers), (&amodifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &amodifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:75
// index:0
// inline
// ulong timestamp()
func (this *QInputEvent) Timestamp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9timestampEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:76
// index:0
// inline
// void setTimestamp(ulong)
func (this *QInputEvent) SetTimestamp(atimestamp uint) {
	// 0: (, atimestamp ulong), (&atimestamp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEvent12setTimestampEm", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &atimestamp)
	gopp.ErrPrint(err, rv)
}

//  body block end
