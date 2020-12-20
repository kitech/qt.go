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
func QSpinBoxFromptr(cthis Voidptr) *QSpinBox {
	bcthis0 := QAbstractSpinBoxFromptr(cthis)
	return &QSpinBox{bcthis0}
}
func (*QSpinBox) Fromptr(cthis Voidptr) *QSpinBox {
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
	var convArg0 Voidptr
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
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3685691689, "_ZN8QSpinBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:69
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int value() const

/*
 */
func (this *QSpinBox) Value() int {
	rv, err := qtrt.Qtcc1(882179790, "_ZNK8QSpinBox5valueEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:77
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString cleanText() const

/*
 */
func (this *QSpinBox) CleanText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(1673841482, "_ZNK8QSpinBox9cleanTextEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:80
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setSingleStep(int)

/*
 */
func (this *QSpinBox) SetSingleStep(val int) {
	rv, err := qtrt.Qtcc1(3398301100, "_ZN8QSpinBox13setSingleStepEi", qtrt.FFITY_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:83
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimum(int)

/*
 */
func (this *QSpinBox) SetMinimum(min int) {
	rv, err := qtrt.Qtcc1(4211881020, "_ZN8QSpinBox10setMinimumEi", qtrt.FFITY_POINTER, this.GetCthis(), min)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:86
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximum(int)

/*
 */
func (this *QSpinBox) SetMaximum(max int) {
	rv, err := qtrt.Qtcc1(753370749, "_ZN8QSpinBox10setMaximumEi", qtrt.FFITY_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:88
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*
 */
func (this *QSpinBox) SetRange(min int, max int) {
	rv, err := qtrt.Qtcc1(3910809396, "_ZN8QSpinBox8setRangeEii", qtrt.FFITY_POINTER, this.GetCthis(), min, max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setValue(int)

/*
 */
func (this *QSpinBox) SetValue(val int) {
	rv, err := qtrt.Qtcc1(40343854, "_ZN8QSpinBox8setValueEi", qtrt.FFITY_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:108
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void valueChanged(int)

/*
 */
func (this *QSpinBox) ValueChanged(arg0 int) {
	rv, err := qtrt.Qtcc1(924629808, "_ZN8QSpinBox12valueChangedEi", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textChanged(const QString &)

/*
 */
func (this *QSpinBox) TextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(2388599440, "_ZN8QSpinBox11textChangedERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQSpinBox(this *QSpinBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN8QSpinBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10149() {
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
