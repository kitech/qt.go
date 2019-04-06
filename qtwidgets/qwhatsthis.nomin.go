//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qwhatsthis.h
// #include <qwhatsthis.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

// /usr/include/qt/QtWidgets/qwhatsthis.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAction * createAction(QObject *)

/*
Returns a ready-made QAction, used to invoke "What's This?" context help, with the given parent.

The returned QAction provides a convenient way to let users enter "What's This?" mode.
*/
func (this *QWhatsThis) CreateAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWhatsThis_CreateAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var nilthis *QWhatsThis
	rv := nilthis.CreateAction(parent)
	return rv
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAction * createAction(QObject *)

/*
Returns a ready-made QAction, used to invoke "What's This?" context help, with the given parent.

The returned QAction provides a convenient way to let users enter "What's This?" mode.
*/
func (this *QWhatsThis) CreateActionp() *QAction /*777 QAction **/ {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11372() {
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
