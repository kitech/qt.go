//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QDragMoveEvent struct {
	*QDropEvent
}

func (this *QDragMoveEvent) GetCthis() unsafe.Pointer {
	return this.QDropEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:642
// index:0
// void QDragMoveEvent(const class QPoint &, Qt::DropActions, const class QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, enum QEvent::Type)
func NewQDragMoveEvent(pos unsafe.Pointer, actions int, data unsafe.Pointer, buttons int, modifiers int, type_ int) *QDragMoveEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEventC2ERK6QPoint6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, pos, &actions, data, &buttons, &modifiers, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDragMoveEventFromPointer(cthis)
	return gothis
}
func NewQDragMoveEventFromPointer(cthis unsafe.Pointer) *QDragMoveEvent {
	bcthis0 := NewQDropEventFromPointer(cthis)
	return &QDragMoveEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:644
// index:0
// virtual
// void ~QDragMoveEvent()
func DeleteQDragMoveEvent(*QDragMoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:646
// index:0
// inline
// QRect answerRect()
func (this *QDragMoveEvent) AnswerRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDragMoveEvent10answerRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:648
// index:0
// inline
// void accept()
func (this *QDragMoveEvent) Accept() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:651
// index:1
// inline
// void accept(const class QRect &)
func (this *QDragMoveEvent) Accept_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:649
// index:0
// inline
// void ignore()
func (this *QDragMoveEvent) Ignore() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:652
// index:1
// inline
// void ignore(const class QRect &)
func (this *QDragMoveEvent) Ignore_1(r unsafe.Pointer) {
	// 1: (, r const QRect &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

//  body block end
