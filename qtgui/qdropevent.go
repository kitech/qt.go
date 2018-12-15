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
// extern C begin: 5
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
type QDropEvent struct {
	*qtcore.QEvent
}
type QDropEvent_ITF interface {
	qtcore.QEvent_ITF
	QDropEvent_PTR() *QDropEvent
}

func (ptr *QDropEvent) QDropEvent_PTR() *QDropEvent { return ptr }

func (this *QDropEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDropEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQDropEventFromPointer(cthis unsafe.Pointer) *QDropEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QDropEvent{bcthis0}
}
func (*QDropEvent) NewFromPointer(cthis unsafe.Pointer) *QDropEvent {
	return NewQDropEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:601
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDropEvent(const QPointF &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, QEvent::Type)

/*

 */
func (*QDropEvent) NewForInherit(pos qtcore.QPointF_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDropEvent {
	return NewQDropEvent(pos, actions, data, buttons, modifiers, type_)
}
func NewQDropEvent(pos qtcore.QPointF_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDropEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEventC2ERK7QPointF6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDropEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:601
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDropEvent(const QPointF &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, QEvent::Type)

/*

 */
func (*QDropEvent) NewForInheritp(pos qtcore.QPointF_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDropEvent {
	return NewQDropEventp(pos, actions, data, buttons, modifiers)
}
func NewQDropEventp(pos qtcore.QPointF_ITF, actions int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, buttons int, modifiers int) *QDropEvent {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	// arg: 5, QEvent::Type=Enum, QEvent::Type=Enum, , Invalid
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEventC2ERK7QPointF6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDropEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:603
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDropEvent()

/*

 */
func DeleteQDropEvent(this *QDropEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 72)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:605
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos() const

/*

 */
func (this *QDropEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:606
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF() const

/*

 */
func (this *QDropEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:607
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons mouseButtons() const

/*

 */
func (this *QDropEvent) MouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent12mouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:608
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers keyboardModifiers() const

/*

 */
func (this *QDropEvent) KeyboardModifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent17keyboardModifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:610
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropActions possibleActions() const

/*

 */
func (this *QDropEvent) PossibleActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent15possibleActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:611
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropAction proposedAction() const

/*

 */
func (this *QDropEvent) ProposedAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent14proposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:612
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void acceptProposedAction()

/*

 */
func (this *QDropEvent) AcceptProposedAction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEvent20acceptProposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:614
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropAction dropAction() const

/*

 */
func (this *QDropEvent) DropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent10dropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:615
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropAction(Qt::DropAction)

/*

 */
func (this *QDropEvent) SetDropAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEvent13setDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:617
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * source() const

/*

 */
func (this *QDropEvent) Source() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:618
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMimeData * mimeData() const

/*

 */
func (this *QDropEvent) MimeData() *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent8mimeDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
