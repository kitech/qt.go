package qtwidgets

// /usr/include/qt/QtWidgets/qapplication.h
// #include <qapplication.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
// size 16
type QApplication struct {
	*qtgui.QGuiApplication
}
type QApplication_ITF interface {
	qtgui.QGuiApplication_ITF
	QApplication_PTR() *QApplication
}

func (ptr *QApplication) QApplication_PTR() *QApplication { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QApplicationFromptr(cthis Voidptr) *QApplication {
	bcthis0 := qtgui.QGuiApplicationFromptr(cthis)
	return &QApplication{bcthis0}
}
func (*QApplication) Fromptr(cthis Voidptr) *QApplication {
	return QApplicationFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qapplication.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplication(int &, char **, int)

/*
 */
func (*QApplication) NewForInherit(argc int, argv []string, arg2 int) *QApplication {
	return NewQApplication(argc, argv, arg2)
}
func NewQApplication(argc int, argv []string, arg2 int) *QApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(3158502489, "_ZN12QApplicationC2ERiPPci", qtrt.FFITY_POINTER, cthis, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := QApplicationFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QApplication")
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplication(int &, char **, int)

/*
 */
func (*QApplication) NewForInheritp(argc int, argv []string) *QApplication {
	return NewQApplicationp(argc, argv)
}
func NewQApplicationp(argc int, argv []string) *QApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	// arg: 2, int=Int, =Invalid, , Invalid
	arg2 := 0
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(3158502489, "_ZN12QApplicationC2ERiPPci", qtrt.FFITY_POINTER, cthis, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := QApplicationFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QApplication")
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:147
// index:0
// Public static Ignore Visibility=Default Availability=Available
// [-2] void beep()

/*
 */
func (this *QApplication) Beep() {
	rv, err := qtrt.Qtcc1(3241209511, "_ZN12QApplication4beepEv", qtrt.FFITY_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QApplication_Beep() {
	var nilthis *QApplication
	nilthis.Beep()
}

// /usr/include/qt/QtWidgets/qapplication.h:198
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void focusChanged(QWidget *, QWidget *)

/*
 */
func (this *QApplication) FocusChanged(old QWidget_ITF /*777 QWidget **/, now QWidget_ITF /*777 QWidget **/) {
	var convArg0 Voidptr
	if old != nil && old.QWidget_PTR() != nil {
		convArg0 = old.QWidget_PTR().GetCthis()
	}
	var convArg1 Voidptr
	if now != nil && now.QWidget_PTR() != nil {
		convArg1 = now.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(299776149, "_ZN12QApplication12focusChangedEP7QWidgetS1_", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:201
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString styleSheet() const

/*
 */
func (this *QApplication) StyleSheet() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(306547508, "_ZNK12QApplication10styleSheetEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qapplication.h:204
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)

/*
 */
func (this *QApplication) SetStyleSheet(sheet string) {
	var tmpArg0 = qtcore.NewQString5(sheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(412095265, "_ZN12QApplication13setStyleSheetERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:209
// index:0
// Public static Ignore Visibility=Default Availability=Available
// [-2] void aboutQt()

/*
 */
func (this *QApplication) AboutQt() {
	rv, err := qtrt.Qtcc1(799315374, "_ZN12QApplication7aboutQtEv", qtrt.FFITY_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QApplication_AboutQt() {
	var nilthis *QApplication
	nilthis.AboutQt()
}

func DeleteQApplication(this *QApplication) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QApplicationD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QApplication__ColorSpec = int

//
const QApplication__NormalColor QApplication__ColorSpec = 0

//
const QApplication__CustomColor QApplication__ColorSpec = 1

//
const QApplication__ManyColor QApplication__ColorSpec = 2

func (this *QApplication) ColorSpecItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QApplication_ColorSpecItemName(val int) string {
	var nilthis *QApplication
	return nilthis.ColorSpecItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10101() {
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
