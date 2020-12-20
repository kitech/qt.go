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
// extern C begin: 86
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
func QAbstractButtonFromptr(cthis Voidptr) *QAbstractButton {
	bcthis0 := QWidgetFromptr(cthis)
	return &QAbstractButton{bcthis0}
}
func (*QAbstractButton) Fromptr(cthis Voidptr) *QAbstractButton {
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
	var convArg0 Voidptr
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
	var convArg0 Voidptr
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
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:91
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*
 */
func (this *QAbstractButton) SetCheckable(arg0 bool) {
	rv, err := qtrt.Qtcc1(2802838592, "_ZN15QAbstractButton12setCheckableEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:92
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*
 */
func (this *QAbstractButton) IsCheckable() bool {
	rv, err := qtrt.Qtcc1(1736830991, "_ZNK15QAbstractButton11isCheckableEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:94
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isChecked() const

/*
 */
func (this *QAbstractButton) IsChecked() bool {
	rv, err := qtrt.Qtcc1(3862269578, "_ZNK15QAbstractButton9isCheckedEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:96
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDown(bool)

/*
 */
func (this *QAbstractButton) SetDown(arg0 bool) {
	rv, err := qtrt.Qtcc1(2815823766, "_ZN15QAbstractButton7setDownEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:97
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isDown() const

/*
 */
func (this *QAbstractButton) IsDown() bool {
	rv, err := qtrt.Qtcc1(1977163016, "_ZNK15QAbstractButton6isDownEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:99
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoRepeat(bool)

/*
 */
func (this *QAbstractButton) SetAutoRepeat(arg0 bool) {
	rv, err := qtrt.Qtcc1(1398880692, "_ZN15QAbstractButton13setAutoRepeatEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:100
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool autoRepeat() const

/*
 */
func (this *QAbstractButton) AutoRepeat() bool {
	rv, err := qtrt.Qtcc1(3492544815, "_ZNK15QAbstractButton10autoRepeatEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:102
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoRepeatDelay(int)

/*
 */
func (this *QAbstractButton) SetAutoRepeatDelay(arg0 int) {
	rv, err := qtrt.Qtcc1(2923678386, "_ZN15QAbstractButton18setAutoRepeatDelayEi", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:103
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int autoRepeatDelay() const

/*
 */
func (this *QAbstractButton) AutoRepeatDelay() int {
	rv, err := qtrt.Qtcc1(4206853868, "_ZNK15QAbstractButton15autoRepeatDelayEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoRepeatInterval(int)

/*
 */
func (this *QAbstractButton) SetAutoRepeatInterval(arg0 int) {
	rv, err := qtrt.Qtcc1(477848196, "_ZN15QAbstractButton21setAutoRepeatIntervalEi", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:106
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int autoRepeatInterval() const

/*
 */
func (this *QAbstractButton) AutoRepeatInterval() int {
	rv, err := qtrt.Qtcc1(94216004, "_ZNK15QAbstractButton18autoRepeatIntervalEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:108
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setAutoExclusive(bool)

/*
 */
func (this *QAbstractButton) SetAutoExclusive(arg0 bool) {
	rv, err := qtrt.Qtcc1(3712112795, "_ZN15QAbstractButton16setAutoExclusiveEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:109
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool autoExclusive() const

/*
 */
func (this *QAbstractButton) AutoExclusive() bool {
	rv, err := qtrt.Qtcc1(3776291618, "_ZNK15QAbstractButton13autoExclusiveEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:119
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void toggle()

/*
 */
func (this *QAbstractButton) Toggle() {
	rv, err := qtrt.Qtcc1(3222906074, "_ZN15QAbstractButton6toggleEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:120
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*
 */
func (this *QAbstractButton) SetChecked(arg0 bool) {
	rv, err := qtrt.Qtcc1(1138771998, "_ZN15QAbstractButton10setCheckedEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:123
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void pressed()

/*
 */
func (this *QAbstractButton) Pressed() {
	rv, err := qtrt.Qtcc1(2025725746, "_ZN15QAbstractButton7pressedEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:124
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void released()

/*
 */
func (this *QAbstractButton) Released() {
	rv, err := qtrt.Qtcc1(868254391, "_ZN15QAbstractButton8releasedEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
 */
func (this *QAbstractButton) Clicked(checked bool) {
	rv, err := qtrt.Qtcc1(184427086, "_ZN15QAbstractButton7clickedEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:125
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
 */
func (this *QAbstractButton) Clickedp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.Qtcc1(184427086, "_ZN15QAbstractButton7clickedEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractbutton.h:126
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
 */
func (this *QAbstractButton) Toggled(checked bool) {
	rv, err := qtrt.Qtcc1(895414094, "_ZN15QAbstractButton7toggledEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
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
