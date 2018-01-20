//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QShortcutEvent struct {
	*qtcore.QEvent
}

func (this *QShortcutEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQShortcutEventFromPointer(cthis unsafe.Pointer) *QShortcutEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QShortcutEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:767
// index:0
// Public
// void QShortcutEvent(const class QKeySequence &, int, _Bool)
func NewQShortcutEvent(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventC2ERK12QKeySequenceib", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &id, &ambiguous)
	gopp.ErrPrint(err, rv)
	gothis := NewQShortcutEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:768
// index:0
// Public virtual
// void ~QShortcutEvent()
func DeleteQShortcutEvent(*QShortcutEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:770
// index:0
// Public inline
// const QKeySequence & key()
func (this *QShortcutEvent) Key() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:771
// index:0
// Public inline
// int shortcutId()
func (this *QShortcutEvent) ShortcutId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent10shortcutIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:772
// index:0
// Public inline
// bool isAmbiguous()
func (this *QShortcutEvent) IsAmbiguous() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent11isAmbiguousEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
