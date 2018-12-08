package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
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
type QAccessibleTextSelectionEvent struct {
	*QAccessibleTextCursorEvent
}
type QAccessibleTextSelectionEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextSelectionEvent_PTR() *QAccessibleTextSelectionEvent
}

func (ptr *QAccessibleTextSelectionEvent) QAccessibleTextSelectionEvent_PTR() *QAccessibleTextSelectionEvent {
	return ptr
}

func (this *QAccessibleTextSelectionEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleTextCursorEvent.GetCthis()
	}
}
func (this *QAccessibleTextSelectionEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleTextCursorEvent = NewQAccessibleTextCursorEventFromPointer(cthis)
}
func NewQAccessibleTextSelectionEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextSelectionEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextSelectionEvent{bcthis0}
}
func (*QAccessibleTextSelectionEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextSelectionEvent {
	return NewQAccessibleTextSelectionEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:773
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextSelectionEvent(QObject *, int, int)

/*

 */
func (*QAccessibleTextSelectionEvent) NewForInherit(obj qtcore.QObject_ITF /*777 QObject **/, start int, end_ int) *QAccessibleTextSelectionEvent {
	return NewQAccessibleTextSelectionEvent(obj, start, end_)
}
func NewQAccessibleTextSelectionEvent(obj qtcore.QObject_ITF /*777 QObject **/, start int, end_ int) *QAccessibleTextSelectionEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP7QObjectii", qtrt.FFI_TYPE_POINTER, convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextSelectionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextSelectionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:779
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextSelectionEvent(QAccessibleInterface *, int, int)

/*

 */
func (*QAccessibleTextSelectionEvent) NewForInherit1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, start int, end_ int) *QAccessibleTextSelectionEvent {
	return NewQAccessibleTextSelectionEvent1(iface, start, end_)
}
func NewQAccessibleTextSelectionEvent1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, start int, end_ int) *QAccessibleTextSelectionEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii", qtrt.FFI_TYPE_POINTER, convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextSelectionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextSelectionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:786
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextSelectionEvent()

/*

 */
func DeleteQAccessibleTextSelectionEvent(this *QAccessibleTextSelectionEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:788
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelection(int, int)

/*

 */
func (this *QAccessibleTextSelectionEvent) SetSelection(start int, end_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEvent12setSelectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:793
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int selectionStart() const

/*

 */
func (this *QAccessibleTextSelectionEvent) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:794
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int selectionEnd() const

/*

 */
func (this *QAccessibleTextSelectionEvent) SelectionEnd() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent12selectionEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
