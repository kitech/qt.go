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

// /usr/include/qt/QtWidgets/qstyleoption.h:491
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand()

/*

 */
func (*QStyleOptionRubberBand) NewForInherit() *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBand()
}
func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:495
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionRubberBand(int)

/*

 */
func (*QStyleOptionRubberBand) NewForInherit1(version int) *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBand1(version)
}
func NewQStyleOptionRubberBand1(version int) *QStyleOptionRubberBand {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionRubberBandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionRubberBand)
	return gothis
}

func DeleteQStyleOptionRubberBand(this *QStyleOptionRubberBand) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionRubberBandD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11032() {
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
