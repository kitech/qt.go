//  header block begin
// /usr/include/qt/QtWidgets/qprogressbar.h
// #include <qprogressbar.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
type QProgressBar struct {
	*QWidget
}

func (this *QProgressBar) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQProgressBarFromPointer(cthis unsafe.Pointer) *QProgressBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QProgressBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qprogressbar.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QProgressBar) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// Public
// void QProgressBar(class QWidget *)
func NewQProgressBar(parent unsafe.Pointer) *QProgressBar {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressbar.h:72
// index:0
// Public virtual
// void ~QProgressBar()
func DeleteQProgressBar(*QProgressBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:74
// index:0
// Public
// int minimum()
func (this *QProgressBar) Minimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:75
// index:0
// Public
// int maximum()
func (this *QProgressBar) Maximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:77
// index:0
// Public
// int value()
func (this *QProgressBar) Value() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:79
// index:0
// Public virtual
// QString text()
func (this *QProgressBar) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:80
// index:0
// Public
// void setTextVisible(_Bool)
func (this *QProgressBar) SetTextVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar14setTextVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:81
// index:0
// Public
// bool isTextVisible()
func (this *QProgressBar) IsTextVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar13isTextVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:83
// index:0
// Public
// Qt::Alignment alignment()
func (this *QProgressBar) Alignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:86
// index:0
// Public virtual
// QSize sizeHint()
func (this *QProgressBar) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:87
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QProgressBar) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:89
// index:0
// Public
// Qt::Orientation orientation()
func (this *QProgressBar) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:91
// index:0
// Public
// void setInvertedAppearance(_Bool)
func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar21setInvertedAppearanceEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &invert)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:92
// index:0
// Public
// bool invertedAppearance()
func (this *QProgressBar) InvertedAppearance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar18invertedAppearanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:93
// index:0
// Public
// void setTextDirection(class QProgressBar::Direction)
func (this *QProgressBar) SetTextDirection(textDirection int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar16setTextDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &textDirection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:94
// index:0
// Public
// QProgressBar::Direction textDirection()
func (this *QProgressBar) TextDirection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar13textDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:96
// index:0
// Public
// void setFormat(const class QString &)
func (this *QProgressBar) SetFormat(format *qtcore.QString) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar9setFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:97
// index:0
// Public
// void resetFormat()
func (this *QProgressBar) ResetFormat() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar11resetFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:98
// index:0
// Public
// QString format()
func (this *QProgressBar) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:101
// index:0
// Public
// void reset()
func (this *QProgressBar) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:102
// index:0
// Public
// void setRange(int, int)
func (this *QProgressBar) SetRange(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:103
// index:0
// Public
// void setMinimum(int)
func (this *QProgressBar) SetMinimum(minimum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar10setMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:104
// index:0
// Public
// void setMaximum(int)
func (this *QProgressBar) SetMaximum(maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar10setMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:105
// index:0
// Public
// void setValue(int)
func (this *QProgressBar) SetValue(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar8setValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:106
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QProgressBar) SetOrientation(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:109
// index:0
// Public
// void valueChanged(int)
func (this *QProgressBar) ValueChanged(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar12valueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:112
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QProgressBar) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressbar.h:113
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QProgressBar) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:114
// index:0
// Protected
// void initStyleOption(class QStyleOptionProgressBar *)
func (this *QProgressBar) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
