//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QHoverEvent struct {
	*QInputEvent
}

func (this *QHoverEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:157
// index:0
// void QHoverEvent(enum QEvent::Type, const class QPointF &, const class QPointF &, Qt::KeyboardModifiers)
func NewQHoverEvent(type_ int, pos unsafe.Pointer, oldPos unsafe.Pointer, modifiers int) *QHoverEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHoverEventC2EN6QEvent4TypeERK7QPointFS4_6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, cthis, &type_, pos, oldPos, &modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQHoverEventFromPointer(cthis)
	return gothis
}
func NewQHoverEventFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QHoverEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:158
// index:0
// virtual
// void ~QHoverEvent()
func DeleteQHoverEvent(*QHoverEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHoverEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:161
// index:0
// inline
// QPoint pos()
func (this *QHoverEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:162
// index:0
// inline
// QPoint oldPos()
func (this *QHoverEvent) OldPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent6oldPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:165
// index:0
// inline
// const QPointF & posF()
func (this *QHoverEvent) PosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent4posFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:166
// index:0
// inline
// const QPointF & oldPosF()
func (this *QHoverEvent) OldPosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent7oldPosFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
