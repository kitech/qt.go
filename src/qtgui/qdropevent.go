package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

// /usr/include/qt/QtGui/qevent.h:610
// index:0
// Public virtual
// void ~QDropEvent()
func DeleteQDropEvent(*QDropEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:612
// index:0
// Public inline
// QPoint pos()
func (this *QDropEvent) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:613
// index:0
// Public inline
// const QPointF & posF()
func (this *QDropEvent) PosF() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:614
// index:0
// Public inline
// Qt::MouseButtons mouseButtons()
func (this *QDropEvent) MouseButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent12mouseButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:615
// index:0
// Public inline
// Qt::KeyboardModifiers keyboardModifiers()
func (this *QDropEvent) KeyboardModifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent17keyboardModifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:617
// index:0
// Public inline
// Qt::DropActions possibleActions()
func (this *QDropEvent) PossibleActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent15possibleActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:618
// index:0
// Public inline
// Qt::DropAction proposedAction()
func (this *QDropEvent) ProposedAction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent14proposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:619
// index:0
// Public inline
// void acceptProposedAction()
func (this *QDropEvent) AcceptProposedAction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEvent20acceptProposedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:621
// index:0
// Public inline
// Qt::DropAction dropAction()
func (this *QDropEvent) DropAction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent10dropActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:622
// index:0
// Public
// void setDropAction(Qt::DropAction)
func (this *QDropEvent) SetDropAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QDropEvent13setDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:624
// index:0
// Public
// QObject * source()
func (this *QDropEvent) Source() *qtcore.QObject /*444 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:625
// index:0
// Public inline
// const QMimeData * mimeData()
func (this *QDropEvent) MimeData() *qtcore.QMimeData /*444 const QMimeData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QDropEvent8mimeDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
