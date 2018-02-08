package qtwidgets

// /usr/include/qt/QtWidgets/qprogressbar.h
// #include <qprogressbar.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QProgressBar) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QProgressBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void initStyleOption(class QStyleOptionProgressBar *)
func (this *QProgressBar) InheritInitStyleOption(f func(option *QStyleOptionProgressBar /*777 QStyleOptionProgressBar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QProgressBar struct {
	*QWidget
}

func (this *QProgressBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QProgressBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQProgressBarFromPointer(cthis unsafe.Pointer) *QProgressBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QProgressBar{bcthis0}
}
func (*QProgressBar) NewFromPointer(cthis unsafe.Pointer) *QProgressBar {
	return NewQProgressBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QProgressBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressBar(QWidget *)
func NewQProgressBar(parent *QWidget /*777 QWidget **/) *QProgressBar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBarC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressbar.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProgressBar()
func DeleteQProgressBar(this *QProgressBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum()
func (this *QProgressBar) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum()
func (this *QProgressBar) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QProgressBar) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressbar.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString text()
func (this *QProgressBar) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qprogressbar.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextVisible(_Bool)
func (this *QProgressBar) SetTextVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar14setTextVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTextVisible()
func (this *QProgressBar) IsTextVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar13isTextVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressbar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QProgressBar) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QProgressBar) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QProgressBar) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressbar.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QProgressBar) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressbar.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QProgressBar) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInvertedAppearance(_Bool)
func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar21setInvertedAppearanceEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), invert)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invertedAppearance()
func (this *QProgressBar) InvertedAppearance() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar18invertedAppearanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressbar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextDirection(QProgressBar::Direction)
func (this *QProgressBar) SetTextDirection(textDirection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar16setTextDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textDirection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QProgressBar::Direction textDirection()
func (this *QProgressBar) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QString &)
func (this *QProgressBar) SetFormat(format string) {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar9setFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetFormat()
func (this *QProgressBar) ResetFormat() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar11resetFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString format()
func (this *QProgressBar) Format() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qprogressbar.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QProgressBar) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QProgressBar) SetRange(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)
func (this *QProgressBar) SetMinimum(minimum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)
func (this *QProgressBar) SetMaximum(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)
func (this *QProgressBar) SetValue(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QProgressBar) SetOrientation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)
func (this *QProgressBar) ValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar12valueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QProgressBar) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressbar.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QProgressBar) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QProgressBar10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:114
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionProgressBar *)
func (this *QProgressBar) InitStyleOption(option *QStyleOptionProgressBar /*777 QStyleOptionProgressBar **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QProgressBar__Direction = int

const QProgressBar__TopToBottom QProgressBar__Direction = 0
const QProgressBar__BottomToTop QProgressBar__Direction = 1

//  body block end
