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
func NewQFontMetricsFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return &QFontMetrics{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontmetrics.h:61
// index:0
// Public
// void QFontMetrics(const class QFont &)
func NewQFontMetrics(arg0 *QFont) *QFontMetrics {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:62
// index:1
// Public
// void QFontMetrics(const class QFont &, class QPaintDevice *)
func NewQFontMetrics_1(arg0 *QFont, pd unsafe.Pointer) *QFontMetrics {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0, pd)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:64
// index:0
// Public
// void ~QFontMetrics()
func DeleteQFontMetrics(*QFontMetrics) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:72
// index:0
// Public inline
// void swap(class QFontMetrics &)
func (this *QFontMetrics) Swap(other *QFontMetrics) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetrics4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:75
// index:0
// Public
// int ascent()
func (this *QFontMetrics) Ascent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:76
// index:0
// Public
// int capHeight()
func (this *QFontMetrics) CapHeight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9capHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:77
// index:0
// Public
// int descent()
func (this *QFontMetrics) Descent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:78
// index:0
// Public
// int height()
func (this *QFontMetrics) Height() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:79
// index:0
// Public
// int leading()
func (this *QFontMetrics) Leading() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7leadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:80
// index:0
// Public
// int lineSpacing()
func (this *QFontMetrics) LineSpacing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11lineSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:81
// index:0
// Public
// int minLeftBearing()
func (this *QFontMetrics) MinLeftBearing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics14minLeftBearingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:82
// index:0
// Public
// int minRightBearing()
func (this *QFontMetrics) MinRightBearing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics15minRightBearingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:83
// index:0
// Public
// int maxWidth()
func (this *QFontMetrics) MaxWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics8maxWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:85
// index:0
// Public
// int xHeight()
func (this *QFontMetrics) XHeight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7xHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:86
// index:0
// Public
// int averageCharWidth()
func (this *QFontMetrics) AverageCharWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics16averageCharWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:88
// index:0
// Public
// bool inFont(class QChar)
func (this *QFontMetrics) InFont(arg0 *qtcore.QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6inFontE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:89
// index:0
// Public
// bool inFontUcs4(uint)
func (this *QFontMetrics) InFontUcs4(ucs4 uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10inFontUcs4Ej", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ucs4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:91
// index:0
// Public
// int leftBearing(class QChar)
func (this *QFontMetrics) LeftBearing(arg0 *qtcore.QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11leftBearingE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:92
// index:0
// Public
// int rightBearing(class QChar)
func (this *QFontMetrics) RightBearing(arg0 *qtcore.QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12rightBearingE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:93
// index:0
// Public
// int width(const class QString &, int)
func (this *QFontMetrics) Width(arg0 *qtcore.QString, len int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:94
// index:1
// Public
// int width(const class QString &, int, int)
func (this *QFontMetrics) Width_1(arg0 *qtcore.QString, len int, flags int) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len, &flags)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:96
// index:2
// Public
// int width(class QChar)
func (this *QFontMetrics) Width_2(arg0 *qtcore.QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:98
// index:0
// Public
// int charWidth(const class QString &, int)
func (this *QFontMetrics) CharWidth(str *qtcore.QString, pos int) interface{} {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9charWidthERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:101
// index:0
// Public
// QRect boundingRect(class QChar)
func (this *QFontMetrics) BoundingRect(arg0 *qtcore.QChar) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:103
// index:1
// Public
// QRect boundingRect(const class QString &)
func (this *QFontMetrics) BoundingRect_1(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// Public
// QRect boundingRect(const class QRect &, int, const class QString &, int, int *)
func (this *QFontMetrics) BoundingRect_2(r *qtcore.QRect, flags int, text *qtcore.QString, tabstops int, tabarray unsafe.Pointer) interface{} {
	var convArg0 = r.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &flags, convArg2, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// Public inline
// QRect boundingRect(int, int, int, int, int, const class QString &, int, int *)
func (this *QFontMetrics) BoundingRect_3(x int, y int, w int, h int, flags int, text *qtcore.QString, tabstops int, tabarray unsafe.Pointer) interface{} {
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &w, &h, &flags, convArg5, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// Public
// QSize size(int, const class QString &, int, int *)
func (this *QFontMetrics) Size(flags int, str *qtcore.QString, tabstops int, tabarray unsafe.Pointer) interface{} {
	var convArg1 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &flags, convArg1, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:110
// index:0
// Public
// QRect tightBoundingRect(const class QString &)
func (this *QFontMetrics) TightBoundingRect(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics17tightBoundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:112
// index:0
// Public
// QString elidedText(const class QString &, Qt::TextElideMode, int, int)
func (this *QFontMetrics) ElidedText(text *qtcore.QString, mode int, width int, flags int) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10elidedTextERK7QStringN2Qt13TextElideModeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &mode, &width, &flags)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:114
// index:0
// Public
// int underlinePos()
func (this *QFontMetrics) UnderlinePos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12underlinePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:115
// index:0
// Public
// int overlinePos()
func (this *QFontMetrics) OverlinePos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11overlinePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:116
// index:0
// Public
// int strikeOutPos()
func (this *QFontMetrics) StrikeOutPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12strikeOutPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontmetrics.h:117
// index:0
// Public
// int lineWidth()
func (this *QFontMetrics) LineWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9lineWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
