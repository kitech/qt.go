//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoCompletion() const

/*

 */
func (this *QComboBox) AutoCompletion() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletion(bool)

/*

 */
func (this *QComboBox) SetAutoCompletion(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity autoCompletionCaseSensitivity() const

/*

 */
func (this *QComboBox) AutoCompletionCaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)

/*

 */
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sensitivity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)

/*
Sets the completer to use instead of the current completer. If completer is 0, auto completion is disabled.

By default, for an editable combo box, a QCompleter that performs case insensitive inline completion is automatically created.

Note: The completer is removed when the editable property becomes false.

This function was introduced in  Qt 4.2.

See also completer().
*/
func (this *QComboBox) SetCompleter(c QCompleter_ITF /*777 QCompleter **/) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QCompleter_PTR() != nil {
		convArg0 = c.QCompleter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer() const

/*
Returns the completer that is used to auto complete text input for the combobox.

This function was introduced in  Qt 4.2.

See also setCompleter() and editable.
*/
func (this *QComboBox) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:239
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QComboBox) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QComboBox10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11104() {
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
