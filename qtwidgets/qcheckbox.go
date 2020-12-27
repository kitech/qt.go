// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qcheckbox.h
// #include <qcheckbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QCheckBox struct {
	*QAbstractButton
}
type QCheckBox_ITF interface {
	QAbstractButton_ITF
	QCheckBox_PTR() *QCheckBox
}

func (ptr *QCheckBox) QCheckBox_PTR() *QCheckBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QCheckBoxFromptr(cthis Voidptr) *QCheckBox {
	bcthis0 := QAbstractButtonFromptr(cthis)
	return &QCheckBox{bcthis0}
}
func (*QCheckBox) Fromptr(cthis Voidptr) *QCheckBox {
	return QCheckBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(QWidget *)

/*
 */
func (*QCheckBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QCheckBox {
	return NewQCheckBox(parent)
}
func NewQCheckBox(parent QWidget_ITF /*777 QWidget **/) *QCheckBox {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(421980486, "_ZN9QCheckBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QCheckBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QCheckBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(QWidget *)

/*
 */
func (*QCheckBox) NewForInheritp() *QCheckBox {
	return NewQCheckBoxp()
}
func NewQCheckBoxp() *QCheckBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(421980486, "_ZN9QCheckBoxC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QCheckBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QCheckBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(const QString &, QWidget *)

/*
 */
func (*QCheckBox) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QCheckBox {
	return NewQCheckBox1(text, parent)
}
func NewQCheckBox1(text string, parent QWidget_ITF /*777 QWidget **/) *QCheckBox {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(1488224333, "_ZN9QCheckBoxC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QCheckBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QCheckBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCheckBox(const QString &, QWidget *)

/*
 */
func (*QCheckBox) NewForInherit1p(text string) *QCheckBox {
	return NewQCheckBox1p(text)
}
func NewQCheckBox1p(text string) *QCheckBox {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(1488224333, "_ZN9QCheckBoxC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QCheckBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QCheckBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcheckbox.h:65
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QCheckBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(539622553, "_ZNK9QCheckBox8sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QSizeFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:66
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
 */
func (this *QCheckBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(1367630296, "_ZNK9QCheckBox15minimumSizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QSizeFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTristate(bool)

/*
 */
func (this *QCheckBox) SetTristate(y bool) {
	rv, err := qtrt.Qtcc3(4240076077, "_ZN9QCheckBox11setTristateEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&y))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:68
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTristate(bool)

/*
 */
func (this *QCheckBox) SetTristatep() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	y := true
	rv, err := qtrt.Qtcc3(4240076077, "_ZN9QCheckBox11setTristateEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&y))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:69
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTristate() const

/*
 */
func (this *QCheckBox) IsTristate() bool {
	rv, err := qtrt.Qtcc3(4236192185, "_ZNK9QCheckBox10isTristateEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtWidgets/qcheckbox.h:71
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] Qt::CheckState checkState() const

/*
 */
func (this *QCheckBox) CheckState() int {
	rv, err := qtrt.Qtcc3(713884546, "_ZNK9QCheckBox10checkStateEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int()
}

// /usr/include/qt/QtWidgets/qcheckbox.h:72
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)

/*
 */
func (this *QCheckBox) SetCheckState(state int) {
	rv, err := qtrt.Qtcc3(847293468, "_ZN9QCheckBox13setCheckStateEN2Qt10CheckStateE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&state))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qcheckbox.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void stateChanged(int)

/*
 */
func (this *QCheckBox) StateChanged(arg0 int) {
	rv, err := qtrt.Qtcc3(977411577, "_ZN9QCheckBox12stateChangedEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint3(err, rv)
}

func DeleteQCheckBox(this *QCheckBox) {
	rv, err := qtrt.Qtcc3(430589288, "_ZN9QCheckBoxD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10217() {
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
