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
// extern C begin: 2
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
type QAccessibleTextCursorEvent struct {
	*QAccessibleEvent
}
type QAccessibleTextCursorEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleTextCursorEvent_PTR() *QAccessibleTextCursorEvent
}

func (ptr *QAccessibleTextCursorEvent) QAccessibleTextCursorEvent_PTR() *QAccessibleTextCursorEvent {
	return ptr
}

func (this *QAccessibleTextCursorEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func (this *QAccessibleTextCursorEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleEvent = NewQAccessibleEventFromPointer(cthis)
}
func NewQAccessibleTextCursorEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextCursorEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleTextCursorEvent{bcthis0}
}
func (*QAccessibleTextCursorEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextCursorEvent {
	return NewQAccessibleTextCursorEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:747
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextCursorEvent(QObject *, int)

/*

 */
func NewQAccessibleTextCursorEvent(obj qtcore.QObject_ITF /*777 QObject **/, cursorPos int) *QAccessibleTextCursorEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, cursorPos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextCursorEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:753
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextCursorEvent(QAccessibleInterface *, int)

/*

 */
func NewQAccessibleTextCursorEvent_1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, cursorPos int) *QAccessibleTextCursorEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei", qtrt.FFI_TYPE_POINTER, convArg0, cursorPos)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextCursorEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextCursorEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:760
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextCursorEvent()

/*

 */
func DeleteQAccessibleTextCursorEvent(this *QAccessibleTextCursorEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:762
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCursorPosition(int)

/*

 */
func (this *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextCursorEvent17setCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:763
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int cursorPosition() const

/*

 */
func (this *QAccessibleTextCursorEvent) CursorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextCursorEvent14cursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
