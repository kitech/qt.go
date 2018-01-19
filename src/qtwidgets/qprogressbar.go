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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qprogressbar.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QProgressBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:71
// index:0
// void QProgressBar(class QWidget *)
func NewQProgressBar(parent unsafe.Pointer) *QProgressBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QProgressBar{cthis}
}

// /usr/include/qt/QtWidgets/qprogressbar.h:72
// index:0
// virtual
// void ~QProgressBar()
func DeleteQProgressBar(*QProgressBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:74
// index:0
// int minimum()
func (this *QProgressBar) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar7minimumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:75
// index:0
// int maximum()
func (this *QProgressBar) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar7maximumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:77
// index:0
// int value()
func (this *QProgressBar) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:79
// index:0
// virtual
// QString text()
func (this *QProgressBar) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:80
// index:0
// void setTextVisible(_Bool)
func (this *QProgressBar) SetTextVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar14setTextVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:81
// index:0
// bool isTextVisible()
func (this *QProgressBar) IsTextVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar13isTextVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:83
// index:0
// Qt::Alignment alignment()
func (this *QProgressBar) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:86
// index:0
// virtual
// QSize sizeHint()
func (this *QProgressBar) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:87
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QProgressBar) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:89
// index:0
// Qt::Orientation orientation()
func (this *QProgressBar) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:91
// index:0
// void setInvertedAppearance(_Bool)
func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	// 0: (, bool invert), (&invert)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar21setInvertedAppearanceEb", ffiqt.FFI_TYPE_VOID, this.cthis, &invert)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:92
// index:0
// bool invertedAppearance()
func (this *QProgressBar) InvertedAppearance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar18invertedAppearanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:93
// index:0
// void setTextDirection(class QProgressBar::Direction)
func (this *QProgressBar) SetTextDirection(textDirection int) {
	// 0: (, QProgressBar::Direction textDirection), (&textDirection)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar16setTextDirectionENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &textDirection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:94
// index:0
// QProgressBar::Direction textDirection()
func (this *QProgressBar) TextDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar13textDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:96
// index:0
// void setFormat(const class QString &)
func (this *QProgressBar) SetFormat(format unsafe.Pointer) {
	// 0: (, const QString & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar9setFormatERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:97
// index:0
// void resetFormat()
func (this *QProgressBar) ResetFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar11resetFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:98
// index:0
// QString format()
func (this *QProgressBar) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QProgressBar6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:101
// index:0
// void reset()
func (this *QProgressBar) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:102
// index:0
// void setRange(int, int)
func (this *QProgressBar) SetRange(minimum int, maximum int) {
	// 0: (, int minimum, int maximum), (&minimum, &maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar8setRangeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:103
// index:0
// void setMinimum(int)
func (this *QProgressBar) SetMinimum(minimum int) {
	// 0: (, int minimum), (&minimum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar10setMinimumEi", ffiqt.FFI_TYPE_VOID, this.cthis, &minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:104
// index:0
// void setMaximum(int)
func (this *QProgressBar) SetMaximum(maximum int) {
	// 0: (, int maximum), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar10setMaximumEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:105
// index:0
// void setValue(int)
func (this *QProgressBar) SetValue(value int) {
	// 0: (, int value), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar8setValueEi", ffiqt.FFI_TYPE_VOID, this.cthis, &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:106
// index:0
// void setOrientation(Qt::Orientation)
func (this *QProgressBar) SetOrientation(arg0 int) {
	// 0: (, Qt::Orientation arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressbar.h:109
// index:0
// void valueChanged(int)
func (this *QProgressBar) ValueChanged(value int) {
	// 0: (, int value), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QProgressBar12valueChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &value)
	gopp.ErrPrint(err, rv)
}

//  body block end
