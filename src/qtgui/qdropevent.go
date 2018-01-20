//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QDropEvent struct {
	*qtcore.QEvent
}

func (this *QDropEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:608
// index:0
// void QDropEvent(const class QPointF &, Qt::DropActions, const class QMimeData *, Qt::MouseButtons, Qt::KeyboardModifiers, enum QEvent::Type)
func NewQDropEvent(pos unsafe.Pointer, actions int, data unsafe.Pointer, buttons int, modifiers int, type_ int) *QDropEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEventC2ERK7QPointF6QFlagsIN2Qt10DropActionEEPK9QMimeDataS3_INS4_11MouseButtonEES3_INS4_16KeyboardModifierEEN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, pos, &actions, data, &buttons, &modifiers, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDropEventFromPointer(cthis)
	return gothis
}
func NewQDropEventFromPointer(cthis unsafe.Pointer) *QDropEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QDropEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:610
// index:0
// virtual
// void ~QDropEvent()
func DeleteQDropEvent(*QDropEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:612
// index:0
// inline
// QPoint pos()
func (this *QDropEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:613
// index:0
// inline
// const QPointF & posF()
func (this *QDropEvent) PosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent4posFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:614
// index:0
// inline
// Qt::MouseButtons mouseButtons()
func (this *QDropEvent) MouseButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent12mouseButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:615
// index:0
// inline
// Qt::KeyboardModifiers keyboardModifiers()
func (this *QDropEvent) KeyboardModifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent17keyboardModifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:617
// index:0
// inline
// Qt::DropActions possibleActions()
func (this *QDropEvent) PossibleActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent15possibleActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:618
// index:0
// inline
// Qt::DropAction proposedAction()
func (this *QDropEvent) ProposedAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent14proposedActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:619
// index:0
// inline
// void acceptProposedAction()
func (this *QDropEvent) AcceptProposedAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEvent20acceptProposedActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:621
// index:0
// inline
// Qt::DropAction dropAction()
func (this *QDropEvent) DropAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent10dropActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:622
// index:0
// void setDropAction(Qt::DropAction)
func (this *QDropEvent) SetDropAction(action int) {
	// 0: (, action Qt::DropAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEvent13setDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:624
// index:0
// QObject * source()
func (this *QDropEvent) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:625
// index:0
// inline
// const QMimeData * mimeData()
func (this *QDropEvent) MimeData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent8mimeDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
