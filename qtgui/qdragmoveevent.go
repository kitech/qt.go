package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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
import "gopp"
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
	if this == nil {
		return nil
	} else {
		return this.QDropEvent.GetCthis()
	}
}
func (this *QDragMoveEvent) SetCthis(cthis unsafe.Pointer) {
	this.QDropEvent = NewQDropEventFromPointer(cthis)
}
func NewQDragMoveEventFromPointer(cthis unsafe.Pointer) *QDragMoveEvent {
	bcthis0 := NewQDropEventFromPointer(cthis)
	return &QDragMoveEvent{bcthis0}
}
func (*QDragMoveEvent) NewFromPointer(cthis unsafe.Pointer) *QDragMoveEvent {
	return NewQDragMoveEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:642
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDragMoveEvent(const QPoint &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, enum QEvent::Type)
func NewQDragMoveEvent(pos *qtcore.QPoint, actions int, data *qtcore.QMimeData /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDragMoveEvent {
	var convArg0 = pos.GetCthis()
	var convArg2 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEventC2ERK6QPoint6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDragMoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDragMoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:644
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDragMoveEvent()
func DeleteQDragMoveEvent(this *QDragMoveEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:646
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect answerRect()
func (this *QDragMoveEvent) AnswerRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDragMoveEvent10answerRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:648
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void accept()
func (this *QDragMoveEvent) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:651
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void accept(const QRect &)
func (this *QDragMoveEvent) Accept_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:649
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ignore()
func (this *QDragMoveEvent) Ignore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:652
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ignore(const QRect &)
func (this *QDragMoveEvent) Ignore_1(r *qtcore.QRect) {
	var convArg0 = r.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
