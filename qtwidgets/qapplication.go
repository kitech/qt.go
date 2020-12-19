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
func QApplicationFromptr(cthis unsafe.Pointer) *QApplication {
	bcthis0 := qtgui.QGuiApplicationFromptr(cthis)
	return &QApplication{bcthis0}
}
func (*QApplication) Fromptr(cthis unsafe.Pointer) *QApplication {
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

// /usr/include/qt/QtWidgets/qapplication.h:183
// index:0
// Public static Direct Visibility=Default Availability=Available
// [4] int exec()

/*
 */
func (this *QApplication) Exec() int {
	rv, err := qtrt.Qtcc1(1512411141, "_ZN12QApplication4execEv", qtrt.FFITY_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_Exec() int {
	var nilthis *QApplication
	rv := nilthis.Exec()
	return rv
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

func init_unused_10097() {
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
