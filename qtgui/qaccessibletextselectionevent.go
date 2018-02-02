package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
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

type QAccessibleTextSelectionEvent struct {
	*QAccessibleTextCursorEvent
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
func NewQAccessibleTextSelectionEvent(obj *qtcore.QObject /*777 QObject **/, start int, end int) *QAccessibleTextSelectionEvent {
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP7QObjectii", ffiqt.FFI_TYPE_POINTER, convArg0, start, end)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextSelectionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextSelectionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:779
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextSelectionEvent(QAccessibleInterface *, int, int)
func NewQAccessibleTextSelectionEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, start int, end int) *QAccessibleTextSelectionEvent {
	var convArg0 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii", ffiqt.FFI_TYPE_POINTER, convArg0, start, end)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleTextSelectionEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextSelectionEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:786
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextSelectionEvent()
func DeleteQAccessibleTextSelectionEvent(this *QAccessibleTextSelectionEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:788
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelection(int, int)
func (this *QAccessibleTextSelectionEvent) SetSelection(start int, end int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEvent12setSelectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:793
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int selectionStart()
func (this *QAccessibleTextSelectionEvent) SelectionStart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent14selectionStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:794
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int selectionEnd()
func (this *QAccessibleTextSelectionEvent) SelectionEnd() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent12selectionEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
