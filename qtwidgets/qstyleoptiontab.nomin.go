//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qstyleoption.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTab()

/*

 */
func (*QStyleOptionTab) NewForInherit() *QStyleOptionTab {
	return NewQStyleOptionTab()
}
func NewQStyleOptionTab() *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:289
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTab(int)

/*

 */
func (*QStyleOptionTab) NewForInherit1(version int) *QStyleOptionTab {
	return NewQStyleOptionTab1(version)
}
func NewQStyleOptionTab1(version int) *QStyleOptionTab {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTabFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTab)
	return gothis
}

func DeleteQStyleOptionTab(this *QStyleOptionTab) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStyleOptionTabD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11018() {
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
