// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qlabel.h
// #include <qlabel.h>
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

/*
 */
// size 48
type QLabel struct {
	*QFrame
}
type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func (ptr *QLabel) QLabel_PTR() *QLabel { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QLabelFromptr(cthis unsafe.Pointer) *QLabel {
	bcthis0 := QFrameFromptr(cthis)
	return &QLabel{bcthis0}
}
func (*QLabel) Fromptr(cthis unsafe.Pointer) *QLabel {
	return QLabelFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInherit(parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	return NewQLabel(parent, f)
}
func NewQLabel(parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInheritp() *QLabel {
	return NewQLabelp()
}
func NewQLabelp() *QLabel {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLabel(QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QLabel {
	return NewQLabelp1(parent)
}
func NewQLabelp1(parent QWidget_ITF /*777 QWidget **/) *QLabel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	return NewQLabel1(text, parent, f)
}
func NewQLabel1(text string, parent QWidget_ITF /*777 QWidget **/, f int) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInherit1p(text string) *QLabel {
	return NewQLabel1p(text)
}
func NewQLabel1p(text string) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLabel(const QString &, QWidget *, Qt::WindowFlags)

/*
 */
func (*QLabel) NewForInherit1p1(text string, parent QWidget_ITF /*777 QWidget **/) *QLabel {
	return NewQLabel1p1(text, parent)
}
func NewQLabel1p1(text string, parent QWidget_ITF /*777 QWidget **/) *QLabel {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITY_POINTER, cthis, convArg0, convArg1, f)
	qtrt.ErrPrint(err, rv)
	gothis := QLabelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QLabel")
	return gothis
}

// /usr/include/qt/QtWidgets/qlabel.h:74
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString text() const

/*
 */
func (this *QLabel) Text() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(1414356847, "_ZNK6QLabel4textEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlabel.h:136
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
 */
func (this *QLabel) SetText(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(1605932811, "_ZN6QLabel7setTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:144
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setNum(int)

/*
 */
func (this *QLabel) SetNum(arg0 int) {
	rv, err := qtrt.Qtcc1(711141672, "_ZN6QLabel6setNumEi", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:145
// index:1
// Public Ignore Visibility=Default Availability=Available
// [-2] void setNum(double)

/*
 */
func (this *QLabel) SetNum1(arg0 float64) {
	rv, err := qtrt.Qtcc1(1423070613, "_ZN6QLabel6setNumEd", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:146
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QLabel) Clear() {
	rv, err := qtrt.Qtcc1(2401623602, "_ZN6QLabel5clearEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQLabel(this *QLabel) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QLabelD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10129() {
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
