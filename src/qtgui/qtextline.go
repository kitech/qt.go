//  header block begin
// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
type QTextLine struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextlayout.h:213
// index:0
// inline
// void QTextLine()
func NewQTextLine() *QTextLine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLineC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextLine{cthis}
}

// /usr/include/qt/QtGui/qtextlayout.h:214
// index:0
// inline
// bool isValid()
func (this *QTextLine) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:216
// index:0
// QRectF rect()
func (this *QTextLine) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:217
// index:0
// qreal x()
func (this *QTextLine) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:218
// index:0
// qreal y()
func (this *QTextLine) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:219
// index:0
// qreal width()
func (this *QTextLine) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:220
// index:0
// qreal ascent()
func (this *QTextLine) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine6ascentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:221
// index:0
// qreal descent()
func (this *QTextLine) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7descentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:222
// index:0
// qreal height()
func (this *QTextLine) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:223
// index:0
// qreal leading()
func (this *QTextLine) Leading() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7leadingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:225
// index:0
// void setLeadingIncluded(_Bool)
func (this *QTextLine) SetLeadingIncluded(included bool) {
	// 0: (, bool included), (&included)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine18setLeadingIncludedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &included)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:226
// index:0
// bool leadingIncluded()
func (this *QTextLine) LeadingIncluded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine15leadingIncludedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:228
// index:0
// qreal naturalTextWidth()
func (this *QTextLine) NaturalTextWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine16naturalTextWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:229
// index:0
// qreal horizontalAdvance()
func (this *QTextLine) HorizontalAdvance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine17horizontalAdvanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:230
// index:0
// QRectF naturalTextRect()
func (this *QTextLine) NaturalTextRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine15naturalTextRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:242
// index:0
// qreal cursorToX(int *, enum QTextLine::Edge)
func (this *QTextLine) CursorToX(cursorPos unsafe.Pointer, edge int) {
	// 0: (, int * cursorPos, QTextLine::Edge edge), (cursorPos, &edge)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEPiNS_4EdgeE", ffiqt.FFI_TYPE_VOID, this.cthis, cursorPos, &edge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:243
// index:1
// inline
// qreal cursorToX(int, enum QTextLine::Edge)
func (this *QTextLine) CursorToX_1(cursorPos int, edge int) {
	// 1: (, int cursorPos, QTextLine::Edge edge), (&cursorPos, &edge)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEiNS_4EdgeE", ffiqt.FFI_TYPE_VOID, this.cthis, &cursorPos, &edge)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:244
// index:0
// int xToCursor(qreal, enum QTextLine::CursorPosition)
func (this *QTextLine) XToCursor(x float64, arg1 int) {
	// 0: (, qreal x, QTextLine::CursorPosition arg1), (&x, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9xToCursorEdNS_14CursorPositionE", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:246
// index:0
// void setLineWidth(qreal)
func (this *QTextLine) SetLineWidth(width float64) {
	// 0: (, qreal width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine12setLineWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:247
// index:0
// void setNumColumns(int)
func (this *QTextLine) SetNumColumns(columns int) {
	// 0: (, int columns), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:248
// index:1
// void setNumColumns(int, qreal)
func (this *QTextLine) SetNumColumns_1(columns int, alignmentWidth float64) {
	// 1: (, int columns, qreal alignmentWidth), (&columns, &alignmentWidth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEid", ffiqt.FFI_TYPE_VOID, this.cthis, &columns, &alignmentWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:250
// index:0
// void setPosition(const class QPointF &)
func (this *QTextLine) SetPosition(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine11setPositionERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:251
// index:0
// QPointF position()
func (this *QTextLine) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:253
// index:0
// int textStart()
func (this *QTextLine) TextStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9textStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:254
// index:0
// int textLength()
func (this *QTextLine) TextLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine10textLengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:256
// index:0
// inline
// int lineNumber()
func (this *QTextLine) LineNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine10lineNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:258
// index:0
// void draw(class QPainter *, const class QPointF &, const class QTextLayout::FormatRange *)
func (this *QTextLine) Draw(p unsafe.Pointer, point unsafe.Pointer, selection unsafe.Pointer) {
	// 0: (, QPainter * p, const QPointF & point, const QTextLayout::FormatRange * selection), (p, point, selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine4drawEP8QPainterRK7QPointFPKN11QTextLayout11FormatRangeE", ffiqt.FFI_TYPE_VOID, this.cthis, p, point, selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:261
// index:0
// QList<QGlyphRun> glyphRuns(int, int)
func (this *QTextLine) GlyphRuns(from int, length int) {
	// 0: (, int from, int length), (&from, &length)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9glyphRunsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &length)
	gopp.ErrPrint(err, rv)
}

//  body block end
