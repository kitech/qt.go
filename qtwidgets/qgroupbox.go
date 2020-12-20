// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
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
type QGroupBox struct {
	*QWidget
}
type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QGroupBoxFromptr(cthis Voidptr) *QGroupBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QGroupBox{bcthis0}
}
func (*QGroupBox) Fromptr(cthis Voidptr) *QGroupBox {
	return QGroupBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
 */
func (*QGroupBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	return NewQGroupBox(parent)
}
func NewQGroupBox(parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(662756536, "_ZN9QGroupBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
 */
func (*QGroupBox) NewForInheritp() *QGroupBox {
	return NewQGroupBoxp()
}
func NewQGroupBoxp() *QGroupBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(662756536, "_ZN9QGroupBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
 */
func (*QGroupBox) NewForInherit1(title string, parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	return NewQGroupBox1(title, parent)
}
func NewQGroupBox1(title string, parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2800323070, "_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
 */
func (*QGroupBox) NewForInherit1p(title string) *QGroupBox {
	return NewQGroupBox1p(title)
}
func NewQGroupBox1p(title string) *QGroupBox {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2800323070, "_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QGroupBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:66
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString title() const

/*
 */
func (this *QGroupBox) Title() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(3632320991, "_ZNK9QGroupBox5titleEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgroupbox.h:67
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*
 */
func (this *QGroupBox) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(2779866592, "_ZN9QGroupBox8setTitleERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:74
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isFlat() const

/*
 */
func (this *QGroupBox) IsFlat() bool {
	rv, err := qtrt.Qtcc1(2002744521, "_ZNK9QGroupBox6isFlatEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFlat(bool)

/*
 */
func (this *QGroupBox) SetFlat(flat bool) {
	rv, err := qtrt.Qtcc1(2689482562, "_ZN9QGroupBox7setFlatEb", qtrt.FFITY_POINTER, this.GetCthis(), flat)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:76
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*
 */
func (this *QGroupBox) IsCheckable() bool {
	rv, err := qtrt.Qtcc1(2216875125, "_ZNK9QGroupBox11isCheckableEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:77
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*
 */
func (this *QGroupBox) SetCheckable(checkable bool) {
	rv, err := qtrt.Qtcc1(2495419631, "_ZN9QGroupBox12setCheckableEb", qtrt.FFITY_POINTER, this.GetCthis(), checkable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:78
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isChecked() const

/*
 */
func (this *QGroupBox) IsChecked() bool {
	rv, err := qtrt.Qtcc1(3230016360, "_ZNK9QGroupBox9isCheckedEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:81
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*
 */
func (this *QGroupBox) SetChecked(checked bool) {
	rv, err := qtrt.Qtcc1(2906849573, "_ZN9QGroupBox10setCheckedEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
 */
func (this *QGroupBox) Clicked(checked bool) {
	rv, err := qtrt.Qtcc1(1674438564, "_ZN9QGroupBox7clickedEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
 */
func (this *QGroupBox) Clickedp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.Qtcc1(1674438564, "_ZN9QGroupBox7clickedEb", qtrt.FFITY_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:85
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
 */
func (this *QGroupBox) Toggled(arg0 bool) {
	rv, err := qtrt.Qtcc1(1550658212, "_ZN9QGroupBox7toggledEb", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQGroupBox(this *QGroupBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QGroupBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10131() {
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
