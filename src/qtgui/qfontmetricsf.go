//  header block begin
// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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
type QFontMetricsF struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qfontmetrics.h:134
// index:0
// void QFontMetricsF(const class QFont &)
func NewQFontMetricsF(arg0 unsafe.Pointer) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFont", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	return &QFontMetricsF{cthis}
}

// /usr/include/qt/QtGui/qfontmetrics.h:135
// index:1
// void QFontMetricsF(const class QFont &, class QPaintDevice *)
func NewQFontMetricsF_1(arg0 unsafe.Pointer, pd unsafe.Pointer) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, arg0, pd)
	gopp.ErrPrint(err, rv)
	return &QFontMetricsF{cthis}
}

// /usr/include/qt/QtGui/qfontmetrics.h:136
// index:2
// void QFontMetricsF(const class QFontMetrics &)
func NewQFontMetricsF_2(arg0 unsafe.Pointer) *QFontMetricsF {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK12QFontMetrics", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	return &QFontMetricsF{cthis}
}

// /usr/include/qt/QtGui/qfontmetrics.h:138
// index:0
// void ~QFontMetricsF()
func DeleteQFontMetricsF(*QFontMetricsF) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsFD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:147
// index:0
// inline
// void swap(class QFontMetricsF &)
func (this *QFontMetricsF) Swap(other unsafe.Pointer) {
	// 0: (, QFontMetricsF & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontMetricsF4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:149
// index:0
// qreal ascent()
func (this *QFontMetricsF) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6ascentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:150
// index:0
// qreal capHeight()
func (this *QFontMetricsF) CapHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF9capHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:151
// index:0
// qreal descent()
func (this *QFontMetricsF) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7descentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:152
// index:0
// qreal height()
func (this *QFontMetricsF) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:153
// index:0
// qreal leading()
func (this *QFontMetricsF) Leading() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7leadingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:154
// index:0
// qreal lineSpacing()
func (this *QFontMetricsF) LineSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11lineSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:155
// index:0
// qreal minLeftBearing()
func (this *QFontMetricsF) MinLeftBearing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF14minLeftBearingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:156
// index:0
// qreal minRightBearing()
func (this *QFontMetricsF) MinRightBearing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF15minRightBearingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:157
// index:0
// qreal maxWidth()
func (this *QFontMetricsF) MaxWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF8maxWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:159
// index:0
// qreal xHeight()
func (this *QFontMetricsF) XHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF7xHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:160
// index:0
// qreal averageCharWidth()
func (this *QFontMetricsF) AverageCharWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF16averageCharWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:162
// index:0
// bool inFont(class QChar)
func (this *QFontMetricsF) InFont(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF6inFontE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:163
// index:0
// bool inFontUcs4(uint)
func (this *QFontMetricsF) InFontUcs4(ucs4 uint) {
	// 0: (, uint ucs4), (&ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF10inFontUcs4Ej", ffiqt.FFI_TYPE_VOID, this.cthis, &ucs4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:165
// index:0
// qreal leftBearing(class QChar)
func (this *QFontMetricsF) LeftBearing(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11leftBearingE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:166
// index:0
// qreal rightBearing(class QChar)
func (this *QFontMetricsF) RightBearing(arg0 unsafe.Pointer) {
	// 0: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12rightBearingE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:167
// index:0
// qreal width(const class QString &)
func (this *QFontMetricsF) Width(string unsafe.Pointer) {
	// 0: (, const QString & string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, string)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:169
// index:1
// qreal width(class QChar)
func (this *QFontMetricsF) Width_1(arg0 unsafe.Pointer) {
	// 1: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:171
// index:0
// QRectF boundingRect(const class QString &)
func (this *QFontMetricsF) BoundingRect(string unsafe.Pointer) {
	// 0: (, const QString & string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, string)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:172
// index:1
// QRectF boundingRect(class QChar)
func (this *QFontMetricsF) BoundingRect_1(arg0 unsafe.Pointer) {
	// 1: (, QChar arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:173
// index:2
// QRectF boundingRect(const class QRectF &, int, const class QString &, int, int *)
func (this *QFontMetricsF) BoundingRect_2(r unsafe.Pointer, flags int, string unsafe.Pointer, tabstops int, tabarray unsafe.Pointer) {
	// 2: (, const QRectF & r, int flags, const QString & string, int tabstops, int * tabarray), (r, &flags, string, &tabstops, tabarray)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", ffiqt.FFI_TYPE_VOID, this.cthis, r, &flags, string, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:174
// index:0
// QSizeF size(int, const class QString &, int, int *)
func (this *QFontMetricsF) Size(flags int, str unsafe.Pointer, tabstops int, tabarray unsafe.Pointer) {
	// 0: (, int flags, const QString & str, int tabstops, int * tabarray), (&flags, str, &tabstops, tabarray)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", ffiqt.FFI_TYPE_VOID, this.cthis, &flags, str, &tabstops, tabarray)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:176
// index:0
// QRectF tightBoundingRect(const class QString &)
func (this *QFontMetricsF) TightBoundingRect(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF17tightBoundingRectERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:178
// index:0
// QString elidedText(const class QString &, Qt::TextElideMode, qreal, int)
func (this *QFontMetricsF) ElidedText(text unsafe.Pointer, mode int, width float64, flags int) {
	// 0: (, const QString & text, Qt::TextElideMode mode, qreal width, int flags), (text, &mode, &width, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", ffiqt.FFI_TYPE_VOID, this.cthis, text, &mode, &width, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:180
// index:0
// qreal underlinePos()
func (this *QFontMetricsF) UnderlinePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12underlinePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:181
// index:0
// qreal overlinePos()
func (this *QFontMetricsF) OverlinePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF11overlinePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:182
// index:0
// qreal strikeOutPos()
func (this *QFontMetricsF) StrikeOutPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF12strikeOutPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:183
// index:0
// qreal lineWidth()
func (this *QFontMetricsF) LineWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontMetricsF9lineWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
