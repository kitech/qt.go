//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qcompleter.h
// #include <qcompleter.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(const QStringList &, QObject *)

/*
Constructs a completer object with the given parent.
*/
func (*QCompleter) NewForInherit2(completions qtcore.QStringList_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCompleter {
	return NewQCompleter2(completions, parent)
}
func NewQCompleter2(completions qtcore.QStringList_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCompleter {
	var convArg0 unsafe.Pointer
	if completions != nil && completions.QStringList_PTR() != nil {
		convArg0 = completions.QStringList_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(const QStringList &, QObject *)

/*
Constructs a completer object with the given parent.
*/
func (*QCompleter) NewForInherit2p(completions qtcore.QStringList_ITF) *QCompleter {
	return NewQCompleter2p(completions)
}
func NewQCompleter2p(completions qtcore.QStringList_ITF) *QCompleter {
	var convArg0 unsafe.Pointer
	if completions != nil && completions.QStringList_PTR() != nil {
		convArg0 = completions.QStringList_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11112() {
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
