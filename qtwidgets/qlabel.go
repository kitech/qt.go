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
// extern C begin: 12
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
// ignore GetCthis for 1 base
func QLabelFromptr(cthis Voidptr) *QLabel {
	bcthis0 := QFrameFromptr(cthis)
	return &QLabel{bcthis0}
}
func (*QLabel) Fromptr(cthis Voidptr) *QLabel {
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
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3619358992, "_ZN6QLabelC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	var convArg1 Voidptr
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(199160535, "_ZN6QLabelC2ERK7QStringP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
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
	rv, err := qtrt.Qtcc3(1414356847, "_ZNK6QLabel4textEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlabel.h:105
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWordWrap(bool)

/*
 */
func (this *QLabel) SetWordWrap(on bool) {
	rv, err := qtrt.Qtcc3(405278513, "_ZN6QLabel11setWordWrapEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&on))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setIndent(int)

/*
 */
func (this *QLabel) SetIndent(arg0 int) {
	rv, err := qtrt.Qtcc3(3589004144, "_ZN6QLabel9setIndentEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:125
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(bool)

/*
 */
func (this *QLabel) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.Qtcc3(496473961, "_ZN6QLabel20setOpenExternalLinksEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&open))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*
 */
func (this *QLabel) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.Qtcc3(844288563, "_ZN6QLabel23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&flags))
	qtrt.ErrPrint2(err, rv)
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
	rv, err := qtrt.Qtcc3(1605932811, "_ZN6QLabel7setTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:137
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPixmap(const QPixmap &)

/*
 */
func (this *QLabel) SetPixmap(arg0 qtgui.QPixmap_ITF) {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(261774113, "_ZN6QLabel9setPixmapERK7QPixmap", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:144
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setNum(int)

/*
 */
func (this *QLabel) SetNum(arg0 int) {
	rv, err := qtrt.Qtcc3(711141672, "_ZN6QLabel6setNumEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:145
// index:1
// Public Ignore Visibility=Default Availability=Available
// [-2] void setNum(double)

/*
 */
func (this *QLabel) SetNum1(arg0 float64) {
	rv, err := qtrt.Qtcc3(1423070613, "_ZN6QLabel6setNumEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_DOUBLE, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qlabel.h:146
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QLabel) Clear() {
	rv, err := qtrt.Qtcc3(2401623602, "_ZN6QLabel5clearEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

func DeleteQLabel(this *QLabel) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QLabelD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10133() {
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
