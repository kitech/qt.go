//  header block begin
// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QTextLayout struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextlayout.h:108
// index:0
// void QTextLayout()
func NewQTextLayout() *QTextLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextLayout{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:109
// index:1
// void QTextLayout(const class QString &)
func NewQTextLayout_1(text unsafe.Pointer) *QTextLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, text)
	gopp.ErrPrint(err, rv)
	return &QTextLayout{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// void QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
func NewQTextLayout_2(text unsafe.Pointer, font unsafe.Pointer, paintdevice unsafe.Pointer) *QTextLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, text, font, paintdevice)
	gopp.ErrPrint(err, rv)
	return &QTextLayout{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:111
// index:3
// void QTextLayout(const class QTextBlock &)
func NewQTextLayout_3(b unsafe.Pointer) *QTextLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK10QTextBlock", ffiqt.FFI_TYPE_VOID, cthis, b)
	gopp.ErrPrint(err, rv)
	return &QTextLayout{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:112
// index:0
// void ~QTextLayout()
func DeleteQTextLayout(*QTextLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:114
// index:0
// void setFont(const class QFont &)
func (this *QTextLayout) SetFont(f unsafe.Pointer) {
	// 0: (, const QFont & f), (f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:115
// index:0
// QFont font()
func (this *QTextLayout) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:118
// index:0
// void setRawFont(const class QRawFont &)
func (this *QTextLayout) SetRawFont(rawFont unsafe.Pointer) {
	// 0: (, const QRawFont & rawFont), (rawFont)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10setRawFontERK8QRawFont", ffiqt.FFI_TYPE_VOID, this.cthis, rawFont)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:121
// index:0
// void setText(const class QString &)
func (this *QTextLayout) SetText(string unsafe.Pointer) {
	// 0: (, const QString & string), (string)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, string)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:122
// index:0
// QString text()
func (this *QTextLayout) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:124
// index:0
// void setTextOption(const class QTextOption &)
func (this *QTextLayout) SetTextOption(option unsafe.Pointer) {
	// 0: (, const QTextOption & option), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout13setTextOptionERK11QTextOption", ffiqt.FFI_TYPE_VOID, this.cthis, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:125
// index:0
// const QTextOption & textOption()
func (this *QTextLayout) TextOption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10textOptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:127
// index:0
// void setPreeditArea(int, const class QString &)
func (this *QTextLayout) SetPreeditArea(position int, text unsafe.Pointer) {
	// 0: (, int position, const QString & text), (&position, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout14setPreeditAreaEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &position, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:128
// index:0
// int preeditAreaPosition()
func (this *QTextLayout) PreeditAreaPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19preeditAreaPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:129
// index:0
// QString preeditAreaText()
func (this *QTextLayout) PreeditAreaText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15preeditAreaTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:143
// index:0
// QList<QTextLayout::FormatRange> additionalFormats()
func (this *QTextLayout) AdditionalFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout17additionalFormatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:144
// index:0
// void clearAdditionalFormats()
func (this *QTextLayout) ClearAdditionalFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout22clearAdditionalFormatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:147
// index:0
// QVector<QTextLayout::FormatRange> formats()
func (this *QTextLayout) Formats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout7formatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:148
// index:0
// void clearFormats()
func (this *QTextLayout) ClearFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout12clearFormatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:150
// index:0
// void setCacheEnabled(_Bool)
func (this *QTextLayout) SetCacheEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout15setCacheEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:151
// index:0
// bool cacheEnabled()
func (this *QTextLayout) CacheEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12cacheEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:153
// index:0
// void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextLayout) SetCursorMoveStyle(style int) {
	// 0: (, Qt::CursorMoveStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout18setCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:154
// index:0
// Qt::CursorMoveStyle cursorMoveStyle()
func (this *QTextLayout) CursorMoveStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15cursorMoveStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:156
// index:0
// void beginLayout()
func (this *QTextLayout) BeginLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11beginLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:157
// index:0
// void endLayout()
func (this *QTextLayout) EndLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout9endLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:158
// index:0
// void clearLayout()
func (this *QTextLayout) ClearLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11clearLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:160
// index:0
// QTextLine createLine()
func (this *QTextLayout) CreateLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10createLineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:162
// index:0
// int lineCount()
func (this *QTextLayout) LineCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout9lineCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:163
// index:0
// QTextLine lineAt(int)
func (this *QTextLayout) LineAt(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout6lineAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:164
// index:0
// QTextLine lineForTextPosition(int)
func (this *QTextLayout) LineForTextPosition(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19lineForTextPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:170
// index:0
// bool isValidCursorPosition(int)
func (this *QTextLayout) IsValidCursorPosition(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout21isValidCursorPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// int nextCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) NextCursorPosition(oldPos int, mode int) {
	// 0: (, int oldPos, QTextLayout::CursorMode mode), (&oldPos, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &oldPos, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// int previousCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) PreviousCursorPosition(oldPos int, mode int) {
	// 0: (, int oldPos, QTextLayout::CursorMode mode), (&oldPos, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &oldPos, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:173
// index:0
// int leftCursorPosition(int)
func (this *QTextLayout) LeftCursorPosition(oldPos int) {
	// 0: (, int oldPos), (&oldPos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18leftCursorPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &oldPos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:174
// index:0
// int rightCursorPosition(int)
func (this *QTextLayout) RightCursorPosition(oldPos int) {
	// 0: (, int oldPos), (&oldPos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19rightCursorPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &oldPos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:178
// index:0
// void drawCursor(class QPainter *, const class QPointF &, int)
func (this *QTextLayout) DrawCursor(p unsafe.Pointer, pos unsafe.Pointer, cursorPosition int) {
	// 0: (, QPainter * p, const QPointF & pos, int cursorPosition), (p, pos, &cursorPosition)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi", ffiqt.FFI_TYPE_VOID, this.cthis, p, pos, &cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:179
// index:1
// void drawCursor(class QPainter *, const class QPointF &, int, int)
func (this *QTextLayout) DrawCursor_1(p unsafe.Pointer, pos unsafe.Pointer, cursorPosition int, width int) {
	// 1: (, QPainter * p, const QPointF & pos, int cursorPosition, int width), (p, pos, &cursorPosition, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii", ffiqt.FFI_TYPE_VOID, this.cthis, p, pos, &cursorPosition, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:181
// index:0
// QPointF position()
func (this *QTextLayout) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:182
// index:0
// void setPosition(const class QPointF &)
func (this *QTextLayout) SetPosition(p unsafe.Pointer) {
	// 0: (, const QPointF & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11setPositionERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:184
// index:0
// QRectF boundingRect()
func (this *QTextLayout) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:186
// index:0
// qreal minimumWidth()
func (this *QTextLayout) MinimumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12minimumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:187
// index:0
// qreal maximumWidth()
func (this *QTextLayout) MaximumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12maximumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:190
// index:0
// QList<QGlyphRun> glyphRuns(int, int)
func (this *QTextLayout) GlyphRuns(from int, length int) {
	// 0: (, int from, int length), (&from, &length)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout9glyphRunsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:193
// index:0
// inline
// QTextEngine * engine()
func (this *QTextLayout) Engine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout6engineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:194
// index:0
// void setFlags(int)
func (this *QTextLayout) SetFlags(flags int) {
	// 0: (, int flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout8setFlagsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &flags)
	gopp.ErrPrint(err, rv)
}

//  body block end
