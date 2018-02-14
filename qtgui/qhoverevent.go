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
// extern C begin: 19
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

type QHoverEvent struct {
	*QInputEvent
}
type QHoverEvent_ITF interface {
	QInputEvent_ITF
	QHoverEvent_PTR() *QHoverEvent
}

func (ptr *QHoverEvent) QHoverEvent_PTR() *QHoverEvent { return ptr }

func (this *QHoverEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QHoverEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQHoverEventFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QHoverEvent{bcthis0}
}
func (*QHoverEvent) NewFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	return NewQHoverEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHoverEvent(enum QEvent::Type, const QPointF &, const QPointF &, Qt::KeyboardModifiers)
func NewQHoverEvent(type_ int, pos qtcore.QPointF_ITF, oldPos qtcore.QPointF_ITF, modifiers int) *QHoverEvent {
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg1 = pos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if oldPos != nil && oldPos.QPointF_PTR() != nil {
		convArg2 = oldPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHoverEventC2EN6QEvent4TypeERK7QPointFS4_6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHoverEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHoverEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHoverEvent()
func DeleteQHoverEvent(this *QHoverEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHoverEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QHoverEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHoverEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint oldPos()
func (this *QHoverEvent) OldPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHoverEvent6oldPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:165
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF()
func (this *QHoverEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHoverEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & oldPosF()
func (this *QHoverEvent) OldPosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHoverEvent7oldPosFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
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
