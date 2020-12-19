package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
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
// size 32
type QBoxLayout struct {
	*QLayout
}
type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QBoxLayoutFromptr(cthis unsafe.Pointer) *QBoxLayout {
	bcthis0 := QLayoutFromptr(cthis)
	return &QBoxLayout{bcthis0}
}
func (*QBoxLayout) Fromptr(cthis unsafe.Pointer) *QBoxLayout {
	return QBoxLayoutFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBoxLayout(QBoxLayout::Direction, QWidget *)

/*
 */
func (*QBoxLayout) NewForInherit(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QBoxLayout {
	return NewQBoxLayout(arg0, parent)
}
func NewQBoxLayout(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QBoxLayout {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc1(3198258420, "_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", qtrt.FFITY_POINTER, cthis, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBoxLayout(QBoxLayout::Direction, QWidget *)

/*
 */
func (*QBoxLayout) NewForInheritp(arg0 int) *QBoxLayout {
	return NewQBoxLayoutp(arg0)
}
func NewQBoxLayoutp(arg0 int) *QBoxLayout {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc1(3198258420, "_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", qtrt.FFITY_POINTER, cthis, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QBoxLayoutFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:68
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QBoxLayout::Direction direction() const

/*
 */
func (this *QBoxLayout) Direction() int {
	rv, err := qtrt.Qtcc1(2927993936, "_ZNK10QBoxLayout9directionEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:69
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDirection(QBoxLayout::Direction)

/*
 */
func (this *QBoxLayout) SetDirection(arg0 int) {
	rv, err := qtrt.Qtcc1(1715852557, "_ZN10QBoxLayout12setDirectionENS_9DirectionE", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:71
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addSpacing(int)

/*
 */
func (this *QBoxLayout) AddSpacing(size int) {
	rv, err := qtrt.Qtcc1(955949330, "_ZN10QBoxLayout10addSpacingEi", qtrt.FFITY_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
 */
func (this *QBoxLayout) AddStretch(stretch int) {
	rv, err := qtrt.Qtcc1(1570309749, "_ZN10QBoxLayout10addStretchEi", qtrt.FFITY_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
 */
func (this *QBoxLayout) AddStretchp() {
	// arg: 0, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(1570309749, "_ZN10QBoxLayout10addStretchEi", qtrt.FFITY_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
 */
func (this *QBoxLayout) AddWidget(arg0 QWidget_ITF /*777 QWidget **/, stretch int, alignment int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(821085071, "_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
 */
func (this *QBoxLayout) AddWidgetp(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	// arg: 2, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.Qtcc1(821085071, "_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
 */
func (this *QBoxLayout) AddWidgetp1(arg0 QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.Qtcc1(821085071, "_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int)

/*
 */
func (this *QBoxLayout) AddLayout(layout QLayout_ITF /*777 QLayout **/, stretch int) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(4220267810, "_ZN10QBoxLayout9addLayoutEP7QLayouti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int)

/*
 */
func (this *QBoxLayout) AddLayoutp(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.Qtcc1(4220267810, "_ZN10QBoxLayout9addLayoutEP7QLayouti", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:76
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void addStrut(int)

/*
 */
func (this *QBoxLayout) AddStrut(arg0 int) {
	rv, err := qtrt.Qtcc1(4240054649, "_ZN10QBoxLayout8addStrutEi", qtrt.FFITY_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQBoxLayout(this *QBoxLayout) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QBoxLayoutD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QBoxLayout__Direction = int

//
const QBoxLayout__LeftToRight QBoxLayout__Direction = 0

//
const QBoxLayout__RightToLeft QBoxLayout__Direction = 1

//
const QBoxLayout__TopToBottom QBoxLayout__Direction = 2

//
const QBoxLayout__BottomToTop QBoxLayout__Direction = 3

//
const QBoxLayout__Down QBoxLayout__Direction = 2

//
const QBoxLayout__Up QBoxLayout__Direction = 3

func (this *QBoxLayout) DirectionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QBoxLayout_DirectionItemName(val int) string {
	var nilthis *QBoxLayout
	return nilthis.DirectionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10109() {
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
