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
// extern C begin: 6
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

type QEnterEvent struct {
	*qtcore.QEvent
}
type QEnterEvent_ITF interface {
	qtcore.QEvent_ITF
	QEnterEvent_PTR() *QEnterEvent
}

func (ptr *QEnterEvent) QEnterEvent_PTR() *QEnterEvent { return ptr }

func (this *QEnterEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QEnterEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQEnterEventFromPointer(cthis unsafe.Pointer) *QEnterEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QEnterEvent{bcthis0}
}
func (*QEnterEvent) NewFromPointer(cthis unsafe.Pointer) *QEnterEvent {
	return NewQEnterEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEnterEvent(const QPointF &, const QPointF &, const QPointF &)
func NewQEnterEvent(localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF) *QEnterEvent {
	var convArg0 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg0 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if windowPos != nil && windowPos.QPointF_PTR() != nil {
		convArg1 = windowPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg2 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QEnterEventC2ERK7QPointFS2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEnterEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEnterEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEnterEvent()
func DeleteQEnterEvent(this *QEnterEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QEnterEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 72)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QEnterEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint globalPos()
func (this *QEnterEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QEnterEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QEnterEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX()
func (this *QEnterEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY()
func (this *QEnterEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & localPos()
func (this *QEnterEvent) LocalPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent8localPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & windowPos()
func (this *QEnterEvent) WindowPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent9windowPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & screenPos()
func (this *QEnterEvent) ScreenPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QEnterEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
