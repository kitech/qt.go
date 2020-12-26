package qtwidgets

// /usr/include/qt/QtWidgets/qframe.h
// #include <qframe.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QFrame struct {
	*QWidget
}
type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func (ptr *QFrame) QFrame_PTR() *QFrame { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QFrameFromptr(cthis Voidptr) *QFrame {
	bcthis0 := QWidgetFromptr(cthis)
	return &QFrame{bcthis0}
}
func (*QFrame) Fromptr(cthis Voidptr) *QFrame {
	return QFrameFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
 */
func (*QFrame) NewForInherit(parent QWidget_ITF /*777 QWidget **/, f int) *QFrame {
	return NewQFrame(parent, f)
}
func NewQFrame(parent QWidget_ITF /*777 QWidget **/, f int) *QFrame {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3925201311, "_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QFrameFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
 */
func (*QFrame) NewForInheritp() *QFrame {
	return NewQFramep()
}
func NewQFramep() *QFrame {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3925201311, "_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QFrameFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFrame(QWidget *, Qt::WindowFlags)

/*
 */
func (*QFrame) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QFrame {
	return NewQFramep1(parent)
}
func NewQFramep1(parent QWidget_ITF /*777 QWidget **/) *QFrame {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(3925201311, "_ZN6QFrameC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QFrameFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qframe.h:67
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int frameStyle() const

/*
 */
func (this *QFrame) FrameStyle() int {
	rv, err := qtrt.Qtcc3(226051958, "_ZNK6QFrame10frameStyleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:68
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFrameStyle(int)

/*
 */
func (this *QFrame) SetFrameStyle(arg0 int) {
	rv, err := qtrt.Qtcc3(1788913065, "_ZN6QFrame13setFrameStyleEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qframe.h:70
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int frameWidth() const

/*
 */
func (this *QFrame) FrameWidth() int {
	rv, err := qtrt.Qtcc3(1132186928, "_ZNK6QFrame10frameWidthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qframe.h:72
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
 */
func (this *QFrame) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.Qtcc3(2037293193, "_ZNK6QFrame8sizeHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv2 := qtcore.QSizeFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

func DeleteQFrame(this *QFrame) {
	rv, err := qtrt.Qtcc1(0, "_ZN6QFrameD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QFrame__Shape = int

//
const QFrame__NoFrame QFrame__Shape = 0

//
const QFrame__Box QFrame__Shape = 1

//
const QFrame__Panel QFrame__Shape = 2

//
const QFrame__WinPanel QFrame__Shape = 3

//
const QFrame__HLine QFrame__Shape = 4

//
const QFrame__VLine QFrame__Shape = 5

//
const QFrame__StyledPanel QFrame__Shape = 6

func (this *QFrame) ShapeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFrame_ShapeItemName(val int) string {
	var nilthis *QFrame
	return nilthis.ShapeItemName(val)
}

/*


 */
type QFrame__Shadow = int

//
const QFrame__Plain QFrame__Shadow = 16

//
const QFrame__Raised QFrame__Shadow = 32

//
const QFrame__Sunken QFrame__Shadow = 48

func (this *QFrame) ShadowItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFrame_ShadowItemName(val int) string {
	var nilthis *QFrame
	return nilthis.ShadowItemName(val)
}

/*


 */
type QFrame__StyleMask = int

//
const QFrame__Shadow_Mask QFrame__StyleMask = 240

//
const QFrame__Shape_Mask QFrame__StyleMask = 15

func (this *QFrame) StyleMaskItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFrame_StyleMaskItemName(val int) string {
	var nilthis *QFrame
	return nilthis.StyleMaskItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10179() {
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
