//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
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
type QAccessibleTextSelectionEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qaccessible.h:773
// index:0
// inline
// void QAccessibleTextSelectionEvent(class QObject *, int, int)
func NewQAccessibleTextSelectionEvent(obj unsafe.Pointer, start int, end int) *QAccessibleTextSelectionEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP7QObjectii", ffiqt.FFI_TYPE_VOID, cthis, obj, &start, &end)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTextSelectionEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:779
// index:1
// inline
// void QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
func NewQAccessibleTextSelectionEvent_1(iface unsafe.Pointer, start int, end int) *QAccessibleTextSelectionEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii", ffiqt.FFI_TYPE_VOID, cthis, iface, &start, &end)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTextSelectionEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:786
// index:0
// virtual
// void ~QAccessibleTextSelectionEvent()
func DeleteQAccessibleTextSelectionEvent(*QAccessibleTextSelectionEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:788
// index:0
// inline
// void setSelection(int, int)
func (this *QAccessibleTextSelectionEvent) SetSelection(start int, end int) {
	// 0: (, int start, int end), (&start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTextSelectionEvent12setSelectionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:793
// index:0
// inline
// int selectionStart()
func (this *QAccessibleTextSelectionEvent) SelectionStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent14selectionStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:794
// index:0
// inline
// int selectionEnd()
func (this *QAccessibleTextSelectionEvent) SelectionEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTextSelectionEvent12selectionEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
