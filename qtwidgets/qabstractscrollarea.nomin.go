//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
 */
func (*QAbstractScrollArea) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	return NewQAbstractScrollArea(parent)
}
func NewQAbstractScrollArea(parent QWidget_ITF /*777 QWidget **/) *QAbstractScrollArea {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(538689117, "_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractScrollAreaFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)

/*
 */
func (*QAbstractScrollArea) NewForInheritp() *QAbstractScrollArea {
	return NewQAbstractScrollAreap()
}
func NewQAbstractScrollAreap() *QAbstractScrollArea {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(538689117, "_ZN19QAbstractScrollAreaC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractScrollAreaFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractScrollArea")
	return gothis
}

func DeleteQAbstractScrollArea(this *QAbstractScrollArea) {
	rv, err := qtrt.Qtcc1(0, "_ZN19QAbstractScrollAreaD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10092() {
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
