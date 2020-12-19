// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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

/*
 */
// size 48
type QDoubleSpinBox struct {
	*QAbstractSpinBox
}
type QDoubleSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QDoubleSpinBox_PTR() *QDoubleSpinBox
}

func (ptr *QDoubleSpinBox) QDoubleSpinBox_PTR() *QDoubleSpinBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QDoubleSpinBoxFromptr(cthis unsafe.Pointer) *QDoubleSpinBox {
	bcthis0 := QAbstractSpinBoxFromptr(cthis)
	return &QDoubleSpinBox{bcthis0}
}
func (*QDoubleSpinBox) Fromptr(cthis unsafe.Pointer) *QDoubleSpinBox {
	return QDoubleSpinBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleSpinBox(QWidget *)

/*
 */
func (*QDoubleSpinBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QDoubleSpinBox {
	return NewQDoubleSpinBox(parent)
}
func NewQDoubleSpinBox(parent QWidget_ITF /*777 QWidget **/) *QDoubleSpinBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(698183486, "_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QDoubleSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QDoubleSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDoubleSpinBox(QWidget *)

/*
 */
func (*QDoubleSpinBox) NewForInheritp() *QDoubleSpinBox {
	return NewQDoubleSpinBoxp()
}
func NewQDoubleSpinBoxp() *QDoubleSpinBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(698183486, "_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QDoubleSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QDoubleSpinBox")
	return gothis
}

func DeleteQDoubleSpinBox(this *QDoubleSpinBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN14QDoubleSpinBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10147() {
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
