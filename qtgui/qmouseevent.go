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
// extern C begin: 11
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
type QMouseEvent struct {
	*QInputEvent
}
type QMouseEvent_ITF interface {
	QInputEvent_ITF
	QMouseEvent_PTR() *QMouseEvent
}

func (ptr *QMouseEvent) QMouseEvent_PTR() *QMouseEvent { return ptr }

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
// [-2] void QMouseEvent(QEvent::Type, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)

/*

 */
func (*QMouseEvent) NewForInherit(type_ int, localPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	return NewQMouseEvent(type_, localPos, button, buttons, modifiers)
}
func NewQMouseEvent(type_ int, localPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg1 = localPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFN2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, button, buttons, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:109
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(QEvent::Type, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)

/*

 */
func (*QMouseEvent) NewForInherit_1(type_ int, localPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	return NewQMouseEvent_1(type_, localPos, screenPos, button, buttons, modifiers)
}
func NewQMouseEvent_1(type_ int, localPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg1 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg2 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, button, buttons, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:112
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(QEvent::Type, const QPointF &, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers)

/*

 */
func (*QMouseEvent) NewForInherit_2(type_ int, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	return NewQMouseEvent_2(type_, localPos, windowPos, screenPos, button, buttons, modifiers)
}
func NewQMouseEvent_2(type_ int, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int) *QMouseEvent {
	var convArg1 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg1 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if windowPos != nil && windowPos.QPointF_PTR() != nil {
		convArg2 = windowPos.QPointF_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg3 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_S4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, button, buttons, modifiers)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:115
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QMouseEvent(QEvent::Type, const QPointF &, const QPointF &, const QPointF &, Qt::MouseButton, Qt::MouseButtons, Qt::KeyboardModifiers, Qt::MouseEventSource)

/*

 */
func (*QMouseEvent) NewForInherit_3(type_ int, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int, source int) *QMouseEvent {
	return NewQMouseEvent_3(type_, localPos, windowPos, screenPos, button, buttons, modifiers, source)
}
func NewQMouseEvent_3(type_ int, localPos qtcore.QPointF_ITF, windowPos qtcore.QPointF_ITF, screenPos qtcore.QPointF_ITF, button int, buttons int, modifiers int, source int) *QMouseEvent {
	var convArg1 unsafe.Pointer
	if localPos != nil && localPos.QPointF_PTR() != nil {
		convArg1 = localPos.QPointF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if windowPos != nil && windowPos.QPointF_PTR() != nil {
		convArg2 = windowPos.QPointF_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if screenPos != nil && screenPos.QPointF_PTR() != nil {
		convArg3 = screenPos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEventC2EN6QEvent4TypeERK7QPointFS4_S4_N2Qt11MouseButtonE6QFlagsIS6_ES7_INS5_16KeyboardModifierEENS5_16MouseEventSourceE", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2, convArg3, button, buttons, modifiers, source)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMouseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMouseEvent()

/*

 */
func DeleteQMouseEvent(this *QMouseEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 104)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint pos() const

/*

 */
func (this *QMouseEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint globalPos() const

/*

 */
func (this *QMouseEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*

 */
func (this *QMouseEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QMouseEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX() const

/*

 */
func (this *QMouseEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY() const

/*

 */
func (this *QMouseEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & localPos() const

/*

 */
func (this *QMouseEvent) LocalPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent8localPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & windowPos() const

/*

 */
func (this *QMouseEvent) WindowPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent9windowPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & screenPos() const

/*

 */
func (this *QMouseEvent) ScreenPos() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent9screenPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButton button() const

/*

 */
func (this *QMouseEvent) Button() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent6buttonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons() const

/*

 */
func (this *QMouseEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLocalPos(const QPointF &)

/*

 */
func (this *QMouseEvent) SetLocalPos(localPosition qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if localPosition != nil && localPosition.QPointF_PTR() != nil {
		convArg0 = localPosition.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMouseEvent11setLocalPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source() const

/*

 */
func (this *QMouseEvent) Source() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseEventFlags flags() const

/*

 */
func (this *QMouseEvent) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMouseEvent5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
