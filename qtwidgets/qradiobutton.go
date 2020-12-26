// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qradiobutton.h
// #include <qradiobutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QRadioButton struct {
	*QAbstractButton
}
type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QRadioButtonFromptr(cthis Voidptr) *QRadioButton {
	bcthis0 := QAbstractButtonFromptr(cthis)
	return &QRadioButton{bcthis0}
}
func (*QRadioButton) Fromptr(cthis Voidptr) *QRadioButton {
	return QRadioButtonFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(QWidget *)

/*
 */
func (*QRadioButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	return NewQRadioButton(parent)
}
func NewQRadioButton(parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3516896294, "_ZN12QRadioButtonC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QRadioButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(QWidget *)

/*
 */
func (*QRadioButton) NewForInheritp() *QRadioButton {
	return NewQRadioButtonp()
}
func NewQRadioButtonp() *QRadioButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3516896294, "_ZN12QRadioButtonC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QRadioButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(const QString &, QWidget *)

/*
 */
func (*QRadioButton) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	return NewQRadioButton1(text, parent)
}
func NewQRadioButton1(text string, parent QWidget_ITF /*777 QWidget **/) *QRadioButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3203710175, "_ZN12QRadioButtonC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	gothis := QRadioButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRadioButton(const QString &, QWidget *)

/*
 */
func (*QRadioButton) NewForInherit1p(text string) *QRadioButton {
	return NewQRadioButton1p(text)
}
func NewQRadioButton1p(text string) *QRadioButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3203710175, "_ZN12QRadioButtonC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	gothis := QRadioButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QRadioButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qradiobutton.h:63
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QRadioButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(1977029849, "_ZNK12QRadioButton8sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qradiobutton.h:64
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
 */
func (this *QRadioButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(2537386903, "_ZNK12QRadioButton15minimumSizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

func DeleteQRadioButton(this *QRadioButton) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QRadioButtonD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10233() {
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
