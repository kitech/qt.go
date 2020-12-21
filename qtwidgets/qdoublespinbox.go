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
// extern C begin: 10
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
// ignore GetCthis for 1 base
func QDoubleSpinBoxFromptr(cthis Voidptr) *QDoubleSpinBox {
	bcthis0 := QAbstractSpinBoxFromptr(cthis)
	return &QDoubleSpinBox{bcthis0}
}
func (*QDoubleSpinBox) Fromptr(cthis Voidptr) *QDoubleSpinBox {
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
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(698183486, "_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(698183486, "_ZN14QDoubleSpinBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QDoubleSpinBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QDoubleSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:138
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] double value() const

/*
 */
func (this *QDoubleSpinBox) Value() float64 {
	rv, err := qtrt.Qtcc3(4042644504, "_ZNK14QDoubleSpinBox5valueEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:146
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString cleanText() const

/*
 */
func (this *QDoubleSpinBox) CleanText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(3151881166, "_ZNK14QDoubleSpinBox9cleanTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qspinbox.h:149
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setSingleStep(double)

/*
 */
func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	rv, err := qtrt.Qtcc3(1279985875, "_ZN14QDoubleSpinBox13setSingleStepEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&val))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:152
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimum(double)

/*
 */
func (this *QDoubleSpinBox) SetMinimum(min float64) {
	rv, err := qtrt.Qtcc3(216875734, "_ZN14QDoubleSpinBox10setMinimumEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&min))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:155
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximum(double)

/*
 */
func (this *QDoubleSpinBox) SetMaximum(max float64) {
	rv, err := qtrt.Qtcc3(3674665111, "_ZN14QDoubleSpinBox10setMaximumEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&max))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:157
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setRange(double, double)

/*
 */
func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	rv, err := qtrt.Qtcc3(3493293011, "_ZN14QDoubleSpinBox8setRangeEdd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&min), Voidptr(&max))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:171
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setValue(double)

/*
 */
func (this *QDoubleSpinBox) SetValue(val float64) {
	rv, err := qtrt.Qtcc3(747158720, "_ZN14QDoubleSpinBox8setValueEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&val))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:174
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void valueChanged(double)

/*
 */
func (this *QDoubleSpinBox) ValueChanged(arg0 float64) {
	rv, err := qtrt.Qtcc3(1552051280, "_ZN14QDoubleSpinBox12valueChangedEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:175
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textChanged(const QString &)

/*
 */
func (this *QDoubleSpinBox) TextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(135358936, "_ZN14QDoubleSpinBox11textChangedERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQDoubleSpinBox(this *QDoubleSpinBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN14QDoubleSpinBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10151() {
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
