// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QStatusBar struct {
	*QWidget
}
type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QStatusBarFromptr(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := QWidgetFromptr(cthis)
	return &QStatusBar{bcthis0}
}
func (*QStatusBar) Fromptr(cthis unsafe.Pointer) *QStatusBar {
	return QStatusBarFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
 */
func (*QStatusBar) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QStatusBar {
	return NewQStatusBar(parent)
}
func NewQStatusBar(parent QWidget_ITF /*777 QWidget **/) *QStatusBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2740779997, "_ZN10QStatusBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QStatusBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusBar(QWidget *)

/*
 */
func (*QStatusBar) NewForInheritp() *QStatusBar {
	return NewQStatusBarp()
}
func NewQStatusBarp() *QStatusBar {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2740779997, "_ZN10QStatusBarC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QStatusBarFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QStatusBar")
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int)

/*
 */
func (this *QStatusBar) AddWidget(widget QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(2690176878, "_ZN10QStatusBar9addWidgetEP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int)

/*
 */
func (this *QStatusBar) AddWidgetp(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(2690176878, "_ZN10QStatusBar9addWidgetEP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *, int)

/*
 */
func (this *QStatusBar) InsertWidget(index int, widget QWidget_ITF /*777 QWidget **/, stretch int) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(4143092170, "_ZN10QStatusBar12insertWidgetEiP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *, int)

/*
 */
func (this *QStatusBar) InsertWidgetp(index int, widget QWidget_ITF /*777 QWidget **/) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(4143092170, "_ZN10QStatusBar12insertWidgetEiP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addPermanentWidget(QWidget *, int)

/*
 */
func (this *QStatusBar) AddPermanentWidget(widget QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(1375597379, "_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addPermanentWidget(QWidget *, int)

/*
 */
func (this *QStatusBar) AddPermanentWidgetp(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(1375597379, "_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertPermanentWidget(int, QWidget *, int)

/*
 */
func (this *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF /*777 QWidget **/, stretch int) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(2100641601, "_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int insertPermanentWidget(int, QWidget *, int)

/*
 */
func (this *QStatusBar) InsertPermanentWidgetp(index int, widget QWidget_ITF /*777 QWidget **/) int {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(2100641601, "_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstatusbar.h:66
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)

/*
 */
func (this *QStatusBar) RemoveWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(1858783838, "_ZN10QStatusBar12removeWidgetEP7QWidget", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:68
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setSizeGripEnabled(bool)

/*
 */
func (this *QStatusBar) SetSizeGripEnabled(arg0 bool) {
	rv, err := qtrt.Qtcc1(2512257864, "_ZN10QStatusBar18setSizeGripEnabledEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:69
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isSizeGripEnabled() const

/*
 */
func (this *QStatusBar) IsSizeGripEnabled() bool {
	rv, err := qtrt.Qtcc1(1342450037, "_ZNK10QStatusBar17isSizeGripEnabledEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstatusbar.h:71
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString currentMessage() const

/*
 */
func (this *QStatusBar) CurrentMessage() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(3358295982, "_ZNK10QStatusBar14currentMessageEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int)

/*
 */
func (this *QStatusBar) ShowMessage(text string, timeout int) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(3756771496, "_ZN10QStatusBar11showMessageERK7QStringi", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, int)

/*
 */
func (this *QStatusBar) ShowMessagep(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.Qtcc1(3756771496, "_ZN10QStatusBar11showMessageERK7QStringi", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clearMessage()

/*
 */
func (this *QStatusBar) ClearMessage() {
	rv, err := qtrt.Qtcc1(2515226392, "_ZN10QStatusBar12clearMessageEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:79
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void messageChanged(const QString &)

/*
 */
func (this *QStatusBar) MessageChanged(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(2199278426, "_ZN10QStatusBar14messageChangedERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQStatusBar(this *QStatusBar) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QStatusBarD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10155() {
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
