// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qprogressbar.h
// #include <qprogressbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QProgressBar struct {
	*QWidget
}
type QProgressBar_ITF interface {
	QWidget_ITF
	QProgressBar_PTR() *QProgressBar
}

func (ptr *QProgressBar) QProgressBar_PTR() *QProgressBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QProgressBarFromptr(cthis Voidptr) *QProgressBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QProgressBar{bcthis0}
}
func (*QProgressBar) Fromptr(cthis Voidptr) *QProgressBar {
	return QProgressBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressBar(QWidget *)

/*
 */
func (*QProgressBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QProgressBar {
	return NewQProgressBar(parent)
}
func NewQProgressBar(parent QWidget_ITF /*777 QWidget **/) *QProgressBar {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(572852642, "_ZN12QProgressBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QProgressBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QProgressBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressBar(QWidget *)

/*
 */
func (*QProgressBar) NewForInheritp() *QProgressBar {
	return NewQProgressBarp()
}
func NewQProgressBarp() *QProgressBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(572852642, "_ZN12QProgressBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QProgressBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QProgressBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressbar.h:74
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int minimum() const

/*
 */
func (this *QProgressBar) Minimum() int {
	rv, err := qtrt.Qtcc1(4052484470, "_ZNK12QProgressBar7minimumEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:75
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maximum() const

/*
 */
func (this *QProgressBar) Maximum() int {
	rv, err := qtrt.Qtcc1(644330295, "_ZNK12QProgressBar7maximumEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:77
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int value() const

/*
 */
func (this *QProgressBar) Value() int {
	rv, err := qtrt.Qtcc1(504073427, "_ZNK12QProgressBar5valueEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:79
// index:0
// Public virtual Indirect Visibility=Default Availability=Available
// [8] QString text() const

/*
 */
func (this *QProgressBar) Text() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(4022741238, "_ZNK12QProgressBar4textEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qprogressbar.h:80
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTextVisible(bool)

/*
 */
func (this *QProgressBar) SetTextVisible(visible bool) {
	rv, err := qtrt.Qtcc1(2485447105, "_ZN12QProgressBar14setTextVisibleEb", qtrt.FFITY_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:81
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTextVisible() const

/*
 */
func (this *QProgressBar) IsTextVisible() bool {
	rv, err := qtrt.Qtcc1(1900994530, "_ZNK12QProgressBar13isTextVisibleEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressbar.h:101
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void reset()

/*
 */
func (this *QProgressBar) Reset() {
	rv, err := qtrt.Qtcc1(2829967089, "_ZN12QProgressBar5resetEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:102
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*
 */
func (this *QProgressBar) SetRange(minimum int, maximum int) {
	rv, err := qtrt.Qtcc1(277876671, "_ZN12QProgressBar8setRangeEii", qtrt.FFITY_POINTER, this.GetCthis(), minimum, maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:103
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimum(int)

/*
 */
func (this *QProgressBar) SetMinimum(minimum int) {
	rv, err := qtrt.Qtcc1(192492712, "_ZN12QProgressBar10setMinimumEi", qtrt.FFITY_POINTER, this.GetCthis(), minimum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:104
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximum(int)

/*
 */
func (this *QProgressBar) SetMaximum(maximum int) {
	rv, err := qtrt.Qtcc1(3700590313, "_ZN12QProgressBar10setMaximumEi", qtrt.FFITY_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setValue(int)

/*
 */
func (this *QProgressBar) SetValue(value int) {
	rv, err := qtrt.Qtcc1(830504370, "_ZN12QProgressBar8setValueEi", qtrt.FFITY_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:106
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*
 */
func (this *QProgressBar) SetOrientation(arg0 int) {
	rv, err := qtrt.Qtcc1(2370329577, "_ZN12QProgressBar14setOrientationEN2Qt11OrientationE", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void valueChanged(int)

/*
 */
func (this *QProgressBar) ValueChanged(value int) {
	rv, err := qtrt.Qtcc1(573637723, "_ZN12QProgressBar12valueChangedEi", qtrt.FFITY_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

func DeleteQProgressBar(this *QProgressBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QProgressBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QProgressBar__Direction = int

//
const QProgressBar__TopToBottom QProgressBar__Direction = 0

//
const QProgressBar__BottomToTop QProgressBar__Direction = 1

func (this *QProgressBar) DirectionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProgressBar_DirectionItemName(val int) string {
	var nilthis *QProgressBar
	return nilthis.DirectionItemName(val)
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
