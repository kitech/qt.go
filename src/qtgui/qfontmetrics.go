//  header block begin
// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QFontMetrics struct {
	*qtrt.CObject
}

func (this *QFontMetrics) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qfontmetrics.h:61
// index:0
// void QFontMetrics(const class QFont &)
func NewQFontMetrics(arg0 unsafe.Pointer) *QFontMetrics {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(cthis)
	return gothis
}
func NewQFontMetricsFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return &QFontMetrics{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontmetrics.h:62
// index:1
// void QFontMetrics(const class QFont &, class QPaintDevice *)
func NewQFontMetrics_1(arg0 unsafe.Pointer, pd unsafe.Pointer) *QFontMetrics {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, arg0, pd)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:64
// index:0
// void ~QFontMetrics()
func DeleteQFontMetrics(*QFontMetrics) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:72
// index:0
// inline
// void swap(class QFontMetrics &)
func (this *QFontMetrics) Swap(other unsafe.Pointer) {
	// 0: (, other QFontMetrics &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetrics4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:75
// index:0
// int ascent()
func (this *QFontMetrics) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6ascentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:76
// index:0
// int capHeight()
func (this *QFontMetrics) CapHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9capHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:77
// index:0
// int descent()
func (this *QFontMetrics) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7descentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:78
// index:0
// int height()
func (this *QFontMetrics) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:79
// index:0
// int leading()
func (this *QFontMetrics) Leading() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7leadingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:80
// index:0
// int lineSpacing()
func (this *QFontMetrics) LineSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11lineSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:81
// index:0
// int minLeftBearing()
func (this *QFontMetrics) MinLeftBearing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics14minLeftBearingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:82
// index:0
// int minRightBearing()
func (this *QFontMetrics) MinRightBearing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics15minRightBearingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:83
// index:0
// int maxWidth()
func (this *QFontMetrics) MaxWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics8maxWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:85
// index:0
// int xHeight()
func (this *QFontMetrics) XHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7xHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:86
// index:0
// int averageCharWidth()
func (this *QFontMetrics) AverageCharWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics16averageCharWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:88
// index:0
// bool inFont(class QChar)
func (this *QFontMetrics) InFont(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6inFontE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:89
// index:0
// bool inFontUcs4(uint)
func (this *QFontMetrics) InFontUcs4(ucs4 uint) {
	// 0: (, ucs4 uint), (&ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10inFontUcs4Ej", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ucs4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:91
// index:0
// int leftBearing(class QChar)
func (this *QFontMetrics) LeftBearing(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11leftBearingE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:92
// index:0
// int rightBearing(class QChar)
func (this *QFontMetrics) RightBearing(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12rightBearingE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:93
// index:0
// int width(const class QString &, int)
func (this *QFontMetrics) Width(arg0 unsafe.Pointer, len int) {
	// 0: (, const QString & arg0, len int), (arg0, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:94
// index:1
// int width(const class QString &, int, int)
func (this *QFontMetrics) Width_1(arg0 unsafe.Pointer, len int, flags int) {
	// 1: (, const QString & arg0, len int, flags int), (arg0, &len, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &len, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:96
// index:2
// int width(class QChar)
func (this *QFontMetrics) Width_2(arg0 unsafe.Pointer) {
	// 2: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:98
// index:0
// int charWidth(const class QString &, int)
func (this *QFontMetrics) CharWidth(str unsafe.Pointer, pos int) {
	// 0: (, str const QString &, pos int), (str, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9charWidthERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), str, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:101
// index:0
// QRect boundingRect(class QChar)
func (this *QFontMetrics) BoundingRect(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectE5QChar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:103
// index:1
// QRect boundingRect(const class QString &)
func (this *QFontMetrics) BoundingRect_1(text unsafe.Pointer) {
	// 1: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// QRect boundingRect(const class QRect &, int, const class QString &, int, int *)
func (this *QFontMetrics) BoundingRect_2(r unsafe.Pointer, flags int, text unsafe.Pointer, tabstops int, tabarray unsafe.Pointer) {
	// 2: (, r const QRect &, flags int, text const QString &, tabstops int, tabarray int *), (r, &flags, text, &tabstops, tabarray)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r, &flags, text, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// inline
// QRect boundingRect(int, int, int, int, int, const class QString &, int, int *)
func (this *QFontMetrics) BoundingRect_3(x int, y int, w int, h int, flags int, text unsafe.Pointer, tabstops int, tabarray unsafe.Pointer) {
	// 3: (, x int, y int, w int, h int, flags int, text const QString &, tabstops int, tabarray int *), (&x, &y, &w, &h, &flags, text, &tabstops, tabarray)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &flags, text, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// QSize size(int, const class QString &, int, int *)
func (this *QFontMetrics) Size(flags int, str unsafe.Pointer, tabstops int, tabarray unsafe.Pointer) {
	// 0: (, flags int, str const QString &, tabstops int, tabarray int *), (&flags, str, &tabstops, tabarray)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags, str, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:110
// index:0
// QRect tightBoundingRect(const class QString &)
func (this *QFontMetrics) TightBoundingRect(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics17tightBoundingRectERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:112
// index:0
// QString elidedText(const class QString &, Qt::TextElideMode, int, int)
func (this *QFontMetrics) ElidedText(text unsafe.Pointer, mode int, width int, flags int) {
	// 0: (, text const QString &, mode Qt::TextElideMode, width int, flags int), (text, &mode, &width, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10elidedTextERK7QStringN2Qt13TextElideModeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &mode, &width, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:114
// index:0
// int underlinePos()
func (this *QFontMetrics) UnderlinePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12underlinePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:115
// index:0
// int overlinePos()
func (this *QFontMetrics) OverlinePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11overlinePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:116
// index:0
// int strikeOutPos()
func (this *QFontMetrics) StrikeOutPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12strikeOutPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:117
// index:0
// int lineWidth()
func (this *QFontMetrics) LineWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9lineWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
