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
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QDragMoveEvent struct {
	*QDropEvent
}
type QDragMoveEvent_ITF interface {
	QDropEvent_ITF
	QDragMoveEvent_PTR() *QDragMoveEvent
}

func (ptr *QDragMoveEvent) QDragMoveEvent_PTR() *QDragMoveEvent { return ptr }

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
// [-2] void QDragMoveEvent(const QPoint &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, QEvent::Type)

/*

 */
func (*QDragMoveEvent) NewForInherit(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDragMoveEvent {
	return NewQDragMoveEvent(pos, actions, data, buttons, modifiers, type_)
}
func NewQDragMoveEvent(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDragMoveEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEventC2ERK6QPoint6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDragMoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDragMoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:642
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDragMoveEvent(const QPoint &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, QEvent::Type)

/*

 */
func (*QDragMoveEvent) NewForInheritp(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDragMoveEvent {
	return NewQDragMoveEventp(pos, actions, data, buttons, modifiers)
}
func NewQDragMoveEventp(pos qtcore.QPoint_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDragMoveEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	// arg: 5, QEvent::Type=Enum, QEvent::Type=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEventC2ERK6QPoint6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDragMoveEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDragMoveEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:644
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDragMoveEvent()

/*

 */
func DeleteQDragMoveEvent(this *QDragMoveEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:646
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect answerRect() const

/*

 */
func (this *QDragMoveEvent) AnswerRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDragMoveEvent10answerRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:648
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void accept()

/*

 */
func (this *QDragMoveEvent) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:651
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void accept(const QRect &)

/*

 */
func (this *QDragMoveEvent) Accept1(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6acceptERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:649
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ignore()

/*

 */
func (this *QDragMoveEvent) Ignore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:652
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ignore(const QRect &)

/*

 */
func (this *QDragMoveEvent) Ignore1(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDragMoveEvent6ignoreERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
