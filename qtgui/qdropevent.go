package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

type QDropEvent struct {
	*qtcore.QEvent
}

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

// /usr/include/qt/QtGui/qevent.h:608
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDropEvent(const QPointF &, Qt::DropActions, const QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, enum QEvent::Type)
func NewQDropEvent(pos *qtcore.QPointF, actions int, data *qtcore.QMimeData /*777 const QMimeData **/, buttons int, modifiers int, type_ int) *QDropEvent {
	var convArg0 = pos.GetCthis()
	var convArg2 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEventC2ERK7QPointF6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", qtrt.FFI_TYPE_POINTER, convArg0, actions, convArg2, buttons, modifiers, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDropEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDropEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:610
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDropEvent()
func DeleteQDropEvent(this *QDropEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 72)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:612
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QDropEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:613
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF()
func (this *QDropEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:614
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons mouseButtons()
func (this *QDropEvent) MouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent12mouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:615
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers keyboardModifiers()
func (this *QDropEvent) KeyboardModifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent17keyboardModifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:617
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropActions possibleActions()
func (this *QDropEvent) PossibleActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent15possibleActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:618
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropAction proposedAction()
func (this *QDropEvent) ProposedAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent14proposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:619
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void acceptProposedAction()
func (this *QDropEvent) AcceptProposedAction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEvent20acceptProposedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:621
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::DropAction dropAction()
func (this *QDropEvent) DropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent10dropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:622
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropAction(Qt::DropAction)
func (this *QDropEvent) SetDropAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QDropEvent13setDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:624
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * source()
func (this *QDropEvent) Source() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:625
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMimeData * mimeData()
func (this *QDropEvent) MimeData() *qtcore.QMimeData /*777 const QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QDropEvent8mimeDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end
