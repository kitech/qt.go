//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlineedit.h
// #include <qlineedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 111
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qlineedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)

/*
Sets this line edit to provide auto completions from the completer, c. The completion mode is set using QCompleter::setCompletionMode().

To use a QCompleter with a QValidator or QLineEdit::inputMask, you need to ensure that the model provided to QCompleter contains valid entries. You can use the QSortFilterProxyModel to ensure that the QCompleter's model contains only valid entries.

If c == 0, setCompleter() removes the current completer, effectively disabling auto completion.

This function was introduced in  Qt 4.2.

See also completer() and QCompleter.
*/
func (this *QLineEdit) SetCompleter(completer QCompleter_ITF /*777 QCompleter **/) {
	var convArg0 unsafe.Pointer
	if completer != nil && completer.QCompleter_PTR() != nil {
		convArg0 = completer.QCompleter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setCompleterEP10QCompleter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer() const

/*
Returns the current QCompleter that provides completions.

This function was introduced in  Qt 4.2.

See also setCompleter().
*/
func (this *QLineEdit) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9completerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *, QLineEdit::ActionPosition)

/*
Adds the action to the list of actions at the position.

This function was introduced in  Qt 5.2.
*/
func (this *QLineEdit) AddAction(action QAction_ITF /*777 QAction **/, position int) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9addActionEP7QActionNS_14ActionPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:181
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, QLineEdit::ActionPosition)

/*
Adds the action to the list of actions at the position.

This function was introduced in  Qt 5.2.
*/
func (this *QLineEdit) AddAction1(icon qtgui.QIcon_ITF, position int) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9addActionERK5QIconNS_14ActionPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QWidget::dragEnterEvent().
*/
func (this *QLineEdit) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QLineEdit) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:224
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().
*/
func (this *QLineEdit) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:225
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QLineEdit) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11250() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
