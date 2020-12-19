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
// extern C begin: 4
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
type QSpinBox struct {
	*QAbstractSpinBox
}
type QSpinBox_ITF interface {
	QAbstractSpinBox_ITF
	QSpinBox_PTR() *QSpinBox
}

func (ptr *QSpinBox) QSpinBox_PTR() *QSpinBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QSpinBoxFromptr(cthis unsafe.Pointer) *QSpinBox {
	bcthis0 := QAbstractSpinBoxFromptr(cthis)
	return &QSpinBox{bcthis0}
}
func (*QSpinBox) Fromptr(cthis unsafe.Pointer) *QSpinBox {
	return QSpinBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSpinBox(QWidget *)

/*
 */
func (*QSpinBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QSpinBox {
	return NewQSpinBox(parent)
}
func NewQSpinBox(parent QWidget_ITF /*777 QWidget **/) *QSpinBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3685691689, "_ZN8QSpinBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSpinBox(QWidget *)

/*
 */
func (*QSpinBox) NewForInheritp() *QSpinBox {
	return NewQSpinBoxp()
}
func NewQSpinBoxp() *QSpinBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3685691689, "_ZN8QSpinBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QSpinBox")
	return gothis
}

func DeleteQSpinBox(this *QSpinBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN8QSpinBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10145() {
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
