package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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

type QMouseEvent struct {
	*QInputEvent
}

func (this *QMouseEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QMouseEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQMouseEventFromPointer(cthis unsafe.Pointer) *QMouseEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QMouseEvent{bcthis0}
}
func (*QMouseEvent) NewFromPointer(cthis unsafe.Pointer) *QMouseEvent {
	return NewQMouseEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(enum QEvent::Type, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)
func NewQMouseEvent(type_ int, localPos *qtcore.QPointF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 = localPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFN2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, type_, convArg1, button, buttons, modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:109
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(enum QEvent::Type, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)
func NewQMouseEvent_1(type_ int, localPos *qtcore.QPointF, screenPos *qtcore.QPointF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 = localPos.GetCthis()
	var convArg2 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2, button, buttons, modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:112
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(enum QEvent::Type, const QPointF &, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)
func NewQMouseEvent_2(type_ int, localPos *qtcore.QPointF, windowPos *qtcore.QPointF, screenPos *qtcore.QPointF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 = localPos.GetCthis()
	var convArg2 = windowPos.GetCthis()
	var convArg3 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_S4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, button, buttons, modifiers)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:115
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(enum QEvent::Type, const QPointF &, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::MouseEventSource)
func NewQMouseEvent_3(type_ int, localPos *qtcore.QPointF, windowPos *qtcore.QPointF, screenPos *qtcore.QPointF, button int, buttons int, modifiers int, source int) *QMouseEvent {
	var convArg1 = localPos.GetCthis()
	var convArg2 = windowPos.GetCthis()
	var convArg3 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_S4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEENS5_16MouseEventSourceE", ffiqt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, button, buttons, modifiers, source)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMouseEvent()
func DeleteQMouseEvent(this *QMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 104)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QMouseEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint globalPos()
func (this *QMouseEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QMouseEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QMouseEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX()
func (this *QMouseEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY()
func (this *QMouseEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & localPos()
func (this *QMouseEvent) LocalPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent8localPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & windowPos()
func (this *QMouseEvent) WindowPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9windowPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & screenPos()
func (this *QMouseEvent) ScreenPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButton button()
func (this *QMouseEvent) Button() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons()
func (this *QMouseEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLocalPos(const QPointF &)
func (this *QMouseEvent) SetLocalPos(localPosition *qtcore.QPointF) {
	var convArg0 = localPosition.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEvent11setLocalPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source()
func (this *QMouseEvent) Source() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventFlags flags()
func (this *QMouseEvent) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
