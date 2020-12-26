// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QPushButton struct {
	*QAbstractButton
}
type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QPushButtonFromptr(cthis Voidptr) *QPushButton {
	bcthis0 := QAbstractButtonFromptr(cthis)
	return &QPushButton{bcthis0}
}
func (*QPushButton) Fromptr(cthis Voidptr) *QPushButton {
	return QPushButtonFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(QWidget *)

/*
 */
func (*QPushButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	return NewQPushButton(parent)
}
func NewQPushButton(parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(638151916, "_ZN11QPushButtonC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QPushButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(QWidget *)

/*
 */
func (*QPushButton) NewForInheritp() *QPushButton {
	return NewQPushButtonp()
}
func NewQPushButtonp() *QPushButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(638151916, "_ZN11QPushButtonC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QPushButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QString &, QWidget *)

/*
 */
func (*QPushButton) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	return NewQPushButton1(text, parent)
}
func NewQPushButton1(text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3027368468, "_ZN11QPushButtonC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	gothis := QPushButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QString &, QWidget *)

/*
 */
func (*QPushButton) NewForInherit1p(text string) *QPushButton {
	return NewQPushButton1p(text)
}
func NewQPushButton1p(text string) *QPushButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3027368468, "_ZN11QPushButtonC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	gothis := QPushButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:69
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QPushButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(1743812697, "_ZNK11QPushButton8sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
 */
func (this *QPushButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(2291737008, "_ZNK11QPushButton15minimumSizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool autoDefault() const

/*
 */
func (this *QPushButton) AutoDefault() bool {
	rv, err := qtrt.Qtcc3(627767004, "_ZNK11QPushButton11autoDefaultEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoDefault(bool)

/*
 */
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	rv, err := qtrt.Qtcc3(2921668158, "_ZN11QPushButton14setAutoDefaultEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isDefault() const

/*
 */
func (this *QPushButton) IsDefault() bool {
	rv, err := qtrt.Qtcc3(1488522037, "_ZNK11QPushButton9isDefaultEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDefault(bool)

/*
 */
func (this *QPushButton) SetDefault(arg0 bool) {
	rv, err := qtrt.Qtcc3(2236558823, "_ZN11QPushButton10setDefaultEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFlat(bool)

/*
 */
func (this *QPushButton) SetFlat(arg0 bool) {
	rv, err := qtrt.Qtcc3(167659761, "_ZN11QPushButton7setFlatEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQPushButton(this *QPushButton) {
	rv, err := qtrt.Qtcc3(29964820, "_ZN11QPushButtonD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10211() {
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
