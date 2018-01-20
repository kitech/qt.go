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
	*qtrt.CObject
}

func (this *QTextLayout) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextLayoutFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return &QTextLayout{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextlayout.h:108
// index:0
// Public
// void QTextLayout()
func NewQTextLayout() *QTextLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:109
// index:1
// Public
// void QTextLayout(const class QString &)
func NewQTextLayout_1(text *qtcore.QString) *QTextLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// Public
// void QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
func NewQTextLayout_2(text *qtcore.QString, font *QFont, paintdevice unsafe.Pointer) *QTextLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = text.GetCthis()
	var convArg1 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, paintdevice)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:111
// index:3
// Public
// void QTextLayout(const class QTextBlock &)
func NewQTextLayout_3(b *QTextBlock) *QTextLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = b.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK10QTextBlock", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:112
// index:0
// Public
// void ~QTextLayout()
func DeleteQTextLayout(*QTextLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:114
// index:0
// Public
// void setFont(const class QFont &)
func (this *QTextLayout) SetFont(f *QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:115
// index:0
// Public
// QFont font()
func (this *QTextLayout) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:118
// index:0
// Public
// void setRawFont(const class QRawFont &)
func (this *QTextLayout) SetRawFont(rawFont *QRawFont) {
	var convArg0 = rawFont.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10setRawFontERK8QRawFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:121
// index:0
// Public
// void setText(const class QString &)
func (this *QTextLayout) SetText(string *qtcore.QString) {
	var convArg0 = string.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:122
// index:0
// Public
// QString text()
func (this *QTextLayout) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:124
// index:0
// Public
// void setTextOption(const class QTextOption &)
func (this *QTextLayout) SetTextOption(option *QTextOption) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout13setTextOptionERK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:125
// index:0
// Public
// const QTextOption & textOption()
func (this *QTextLayout) TextOption() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10textOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:127
// index:0
// Public
// void setPreeditArea(int, const class QString &)
func (this *QTextLayout) SetPreeditArea(position int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout14setPreeditAreaEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:128
// index:0
// Public
// int preeditAreaPosition()
func (this *QTextLayout) PreeditAreaPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19preeditAreaPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:129
// index:0
// Public
// QString preeditAreaText()
func (this *QTextLayout) PreeditAreaText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15preeditAreaTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:143
// index:0
// Public
// QList<QTextLayout::FormatRange> additionalFormats()
func (this *QTextLayout) AdditionalFormats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout17additionalFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:144
// index:0
// Public
// void clearAdditionalFormats()
func (this *QTextLayout) ClearAdditionalFormats() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout22clearAdditionalFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:147
// index:0
// Public
// QVector<QTextLayout::FormatRange> formats()
func (this *QTextLayout) Formats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout7formatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:148
// index:0
// Public
// void clearFormats()
func (this *QTextLayout) ClearFormats() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout12clearFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:150
// index:0
// Public
// void setCacheEnabled(_Bool)
func (this *QTextLayout) SetCacheEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout15setCacheEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:151
// index:0
// Public
// bool cacheEnabled()
func (this *QTextLayout) CacheEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12cacheEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:153
// index:0
// Public
// void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextLayout) SetCursorMoveStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout18setCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:154
// index:0
// Public
// Qt::CursorMoveStyle cursorMoveStyle()
func (this *QTextLayout) CursorMoveStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout15cursorMoveStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:156
// index:0
// Public
// void beginLayout()
func (this *QTextLayout) BeginLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11beginLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:157
// index:0
// Public
// void endLayout()
func (this *QTextLayout) EndLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout9endLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:158
// index:0
// Public
// void clearLayout()
func (this *QTextLayout) ClearLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11clearLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:160
// index:0
// Public
// QTextLine createLine()
func (this *QTextLayout) CreateLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout10createLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:162
// index:0
// Public
// int lineCount()
func (this *QTextLayout) LineCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout9lineCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:163
// index:0
// Public
// QTextLine lineAt(int)
func (this *QTextLayout) LineAt(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout6lineAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:164
// index:0
// Public
// QTextLine lineForTextPosition(int)
func (this *QTextLayout) LineForTextPosition(pos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19lineForTextPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:170
// index:0
// Public
// bool isValidCursorPosition(int)
func (this *QTextLayout) IsValidCursorPosition(pos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout21isValidCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// Public
// int nextCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) NextCursorPosition(oldPos int, mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldPos, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// Public
// int previousCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) PreviousCursorPosition(oldPos int, mode int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldPos, &mode)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:173
// index:0
// Public
// int leftCursorPosition(int)
func (this *QTextLayout) LeftCursorPosition(oldPos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout18leftCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldPos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:174
// index:0
// Public
// int rightCursorPosition(int)
func (this *QTextLayout) RightCursorPosition(oldPos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout19rightCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldPos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:178
// index:0
// Public
// void drawCursor(class QPainter *, const class QPointF &, int)
func (this *QTextLayout) DrawCursor(p unsafe.Pointer, pos *qtcore.QPointF, cursorPosition int) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p, convArg1, &cursorPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:179
// index:1
// Public
// void drawCursor(class QPainter *, const class QPointF &, int, int)
func (this *QTextLayout) DrawCursor_1(p unsafe.Pointer, pos *qtcore.QPointF, cursorPosition int, width int) {
	var convArg1 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p, convArg1, &cursorPosition, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:181
// index:0
// Public
// QPointF position()
func (this *QTextLayout) Position() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:182
// index:0
// Public
// void setPosition(const class QPointF &)
func (this *QTextLayout) SetPosition(p *qtcore.QPointF) {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:184
// index:0
// Public
// QRectF boundingRect()
func (this *QTextLayout) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:186
// index:0
// Public
// qreal minimumWidth()
func (this *QTextLayout) MinimumWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12minimumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:187
// index:0
// Public
// qreal maximumWidth()
func (this *QTextLayout) MaximumWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout12maximumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:190
// index:0
// Public
// QList<QGlyphRun> glyphRuns(int, int)
func (this *QTextLayout) GlyphRuns(from int, length int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout9glyphRunsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &length)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:193
// index:0
// Public inline
// QTextEngine * engine()
func (this *QTextLayout) Engine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLayout6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextlayout.h:194
// index:0
// Public
// void setFlags(int)
func (this *QTextLayout) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLayout8setFlagsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

//  body block end
