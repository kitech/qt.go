// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 82
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
type QAbstractButton struct {
	*QWidget
}
type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QAbstractButtonFromptr(cthis unsafe.Pointer) *QAbstractButton {
	bcthis0 := QWidgetFromptr(cthis)
	return &QAbstractButton{bcthis0}
}
func (*QAbstractButton) Fromptr(cthis unsafe.Pointer) *QAbstractButton {
	return QAbstractButtonFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractButton(QWidget *)

/*
 */
func (*QAbstractButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractButton {
	return NewQAbstractButton(parent)
}
func NewQAbstractButton(parent QWidget_ITF /*777 QWidget **/) *QAbstractButton {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(1746901202, "_ZN15QAbstractButtonC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractButton(QWidget *)

/*
 */
func (*QAbstractButton) NewForInheritp() *QAbstractButton {
	return NewQAbstractButtonp()
}
func NewQAbstractButtonp() *QAbstractButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(1746901202, "_ZN15QAbstractButtonC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QAbstractButtonFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:78
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
 */
func (this *QAbstractButton) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(199757402, "_ZN15QAbstractButton7setTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:79
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString text() const

/*
 */
func (this *QAbstractButton) Text() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(786980147, "_ZNK15QAbstractButton4textEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQAbstractButton(this *QAbstractButton) {
	rv, err := qtrt.Qtcc1(0, "_ZN15QAbstractButtonD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10085() {
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
