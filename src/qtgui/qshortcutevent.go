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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:767
// index:0
// void QShortcutEvent(const class QKeySequence &, int, _Bool)
func NewQShortcutEvent(key unsafe.Pointer, id int, ambiguous bool) *QShortcutEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventC2ERK12QKeySequenceib", ffiqt.FFI_TYPE_VOID, cthis, key, &id, &ambiguous)
	gopp.ErrPrint(err, rv)
	return &QShortcutEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:768
// index:0
// virtual
// void ~QShortcutEvent()
func DeleteQShortcutEvent(*QShortcutEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:770
// index:0
// inline
// const QKeySequence & key()
func (this *QShortcutEvent) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:771
// index:0
// inline
// int shortcutId()
func (this *QShortcutEvent) ShortcutId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent10shortcutIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:772
// index:0
// inline
// bool isAmbiguous()
func (this *QShortcutEvent) IsAmbiguous() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent11isAmbiguousEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
