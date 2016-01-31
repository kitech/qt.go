package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qtextlayout.h
// dst-file: /src/gui/qtextlayout.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextLine::horizontalAdvance();
extern void C_ZNK9QTextLine17horizontalAdvanceEv(void* qthis); // 4
  // proto:  qreal QTextLine::height();
extern void C_ZNK9QTextLine6heightEv(void* qthis); // 4
  // proto:  int QTextLine::textStart();
extern void C_ZNK9QTextLine9textStartEv(void* qthis); // 4
  // proto:  void QTextLine::QTextLine();
extern void C_ZN9QTextLineC2Ev(void* qthis); // 1
  // proto:  void QTextLine::setLeadingIncluded(bool included);
extern void C_ZN9QTextLine18setLeadingIncludedEb(void* qthis, bool arg0); // 4
  // proto:  qreal QTextLine::ascent();
extern void C_ZNK9QTextLine6ascentEv(void* qthis); // 4
  // proto:  qreal QTextLine::x();
extern void C_ZNK9QTextLine1xEv(void* qthis); // 4
  // proto:  qreal QTextLine::naturalTextWidth();
extern void C_ZNK9QTextLine16naturalTextWidthEv(void* qthis); // 4
  // proto:  void QTextLine::setNumColumns(int columns);
extern void C_ZN9QTextLine13setNumColumnsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLine::setNumColumns(int columns, qreal alignmentWidth);
extern void C_ZN9QTextLine13setNumColumnsEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QTextLine::descent();
extern void C_ZNK9QTextLine7descentEv(void* qthis); // 4
  // proto:  qreal QTextLine::leading();
extern void C_ZNK9QTextLine7leadingEv(void* qthis); // 4
  // proto:  QRectF QTextLine::naturalTextRect();
extern void C_ZNK9QTextLine15naturalTextRectEv(void* qthis); // 4
  // proto:  qreal QTextLine::width();
extern void C_ZNK9QTextLine5widthEv(void* qthis); // 4
  // proto:  void QTextLine::setLineWidth(qreal width);
extern void C_ZN9QTextLine12setLineWidthEd(void* qthis, double arg0); // 4
  // proto:  bool QTextLine::isValid();
extern void C_ZNK9QTextLine7isValidEv(void* qthis); // 2
  // proto:  int QTextLine::lineNumber();
extern void C_ZNK9QTextLine10lineNumberEv(void* qthis); // 2
  // proto:  void QTextLine::setPosition(const QPointF & pos);
extern void C_ZN9QTextLine11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QList<QGlyphRun> QTextLine::glyphRuns(int from, int length);
extern void C_ZNK9QTextLine9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRectF QTextLine::rect();
extern void C_ZNK9QTextLine4rectEv(void* qthis); // 4
  // proto:  bool QTextLine::leadingIncluded();
extern void C_ZNK9QTextLine15leadingIncludedEv(void* qthis); // 4
  // proto:  int QTextLine::textLength();
extern void C_ZNK9QTextLine10textLengthEv(void* qthis); // 4
  // proto:  qreal QTextLine::y();
extern void C_ZNK9QTextLine1yEv(void* qthis); // 4
  // proto:  QPointF QTextLine::position();
extern void C_ZNK9QTextLine8positionEv(void* qthis); // 4
  // proto:  void QTextLayout::setCacheEnabled(bool enable);
extern void C_ZN11QTextLayout15setCacheEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QRectF QTextLayout::boundingRect();
extern void C_ZNK11QTextLayout12boundingRectEv(void* qthis); // 4
  // proto:  QString QTextLayout::text();
extern void C_ZNK11QTextLayout4textEv(void* qthis); // 4
  // proto:  void QTextLayout::clearAdditionalFormats();
extern void C_ZN11QTextLayout22clearAdditionalFormatsEv(void* qthis); // 4
  // proto:  int QTextLayout::lineCount();
extern void C_ZNK11QTextLayout9lineCountEv(void* qthis); // 4
  // proto:  void QTextLayout::clearLayout();
extern void C_ZN11QTextLayout11clearLayoutEv(void* qthis); // 4
  // proto:  QTextLine QTextLayout::createLine();
extern void C_ZN11QTextLayout10createLineEv(void* qthis); // 4
  // proto:  QFont QTextLayout::font();
extern void C_ZNK11QTextLayout4fontEv(void* qthis); // 4
  // proto:  const QTextOption & QTextLayout::textOption();
extern void C_ZNK11QTextLayout10textOptionEv(void* qthis); // 4
  // proto:  QTextLine QTextLayout::lineAt(int i);
extern void C_ZNK11QTextLayout6lineAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setRawFont(const QRawFont & rawFont);
extern void C_ZN11QTextLayout10setRawFontERK8QRawFont(void* qthis, void* arg0); // 4
  // proto:  void QTextLayout::beginLayout();
extern void C_ZN11QTextLayout11beginLayoutEv(void* qthis); // 4
  // proto:  QList<QGlyphRun> QTextLayout::glyphRuns(int from, int length);
extern void C_ZNK11QTextLayout9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTextLayout::rightCursorPosition(int oldPos);
extern void C_ZNK11QTextLayout19rightCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::~QTextLayout();
extern void C_ZN11QTextLayoutD2Ev(void* qthis); // 4
  // proto:  void QTextLayout::setTextOption(const QTextOption & option);
extern void C_ZN11QTextLayout13setTextOptionERK11QTextOption(void* qthis, void* arg0); // 4
  // proto:  bool QTextLayout::isValidCursorPosition(int pos);
extern void C_ZNK11QTextLayout21isValidCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setPreeditArea(int position, const QString & text);
extern void C_ZN11QTextLayout14setPreeditAreaEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTextEngine * QTextLayout::engine();
extern void C_ZNK11QTextLayout6engineEv(void* qthis); // 2
  // proto:  qreal QTextLayout::minimumWidth();
extern void C_ZNK11QTextLayout12minimumWidthEv(void* qthis); // 4
  // proto:  void QTextLayout::QTextLayout(const QString & text, const QFont & font, QPaintDevice * paintdevice);
extern void C_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice(void* qthis, void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QTextLayout::QTextLayout(const QTextBlock & b);
extern void C_ZN11QTextLayoutC2ERK10QTextBlock(void* qthis, void* arg0); // 3
  // proto:  void QTextLayout::QTextLayout();
extern void C_ZN11QTextLayoutC2Ev(void* qthis); // 3
  // proto:  void QTextLayout::QTextLayout(const QString & text);
extern void C_ZN11QTextLayoutC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  int QTextLayout::leftCursorPosition(int oldPos);
extern void C_ZNK11QTextLayout18leftCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::endLayout();
extern void C_ZN11QTextLayout9endLayoutEv(void* qthis); // 4
  // proto:  qreal QTextLayout::maximumWidth();
extern void C_ZNK11QTextLayout12maximumWidthEv(void* qthis); // 4
  // proto:  void QTextLayout::setFont(const QFont & f);
extern void C_ZN11QTextLayout7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  QTextLine QTextLayout::lineForTextPosition(int pos);
extern void C_ZNK11QTextLayout19lineForTextPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextLayout::preeditAreaPosition();
extern void C_ZNK11QTextLayout19preeditAreaPositionEv(void* qthis); // 4
  // proto:  void QTextLayout::setFlags(int flags);
extern void C_ZN11QTextLayout8setFlagsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setPosition(const QPointF & p);
extern void C_ZN11QTextLayout11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition);
extern void C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition, int width);
extern void C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  bool QTextLayout::cacheEnabled();
extern void C_ZNK11QTextLayout12cacheEnabledEv(void* qthis); // 4
  // proto:  void QTextLayout::setText(const QString & string);
extern void C_ZN11QTextLayout7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::CursorMoveStyle QTextLayout::cursorMoveStyle();
extern void C_ZNK11QTextLayout15cursorMoveStyleEv(void* qthis); // 4
  // proto:  QString QTextLayout::preeditAreaText();
extern void C_ZNK11QTextLayout15preeditAreaTextEv(void* qthis); // 4
  // proto:  QList<QTextLayout::FormatRange> QTextLayout::additionalFormats();
extern void C_ZNK11QTextLayout17additionalFormatsEv(void* qthis); // 4
  // proto:  QPointF QTextLayout::position();
extern void C_ZNK11QTextLayout8positionEv(void* qthis); // 4
  // proto:  int QTextInlineObject::textPosition();
extern void C_ZNK17QTextInlineObject12textPositionEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::descent();
extern void C_ZNK17QTextInlineObject7descentEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setWidth(qreal w);
extern void C_ZN17QTextInlineObject8setWidthEd(void* qthis, double arg0); // 4
  // proto:  QTextFormat QTextInlineObject::format();
extern void C_ZNK17QTextInlineObject6formatEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setDescent(qreal d);
extern void C_ZN17QTextInlineObject10setDescentEd(void* qthis, double arg0); // 4
  // proto:  bool QTextInlineObject::isValid();
extern void C_ZNK17QTextInlineObject7isValidEv(void* qthis); // 2
  // proto:  Qt::LayoutDirection QTextInlineObject::textDirection();
extern void C_ZNK17QTextInlineObject13textDirectionEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::height();
extern void C_ZNK17QTextInlineObject6heightEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::width();
extern void C_ZNK17QTextInlineObject5widthEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setAscent(qreal a);
extern void C_ZN17QTextInlineObject9setAscentEd(void* qthis, double arg0); // 4
  // proto:  int QTextInlineObject::formatIndex();
extern void C_ZNK17QTextInlineObject11formatIndexEv(void* qthis); // 4
  // proto:  void QTextInlineObject::QTextInlineObject();
extern void C_ZN17QTextInlineObjectC2Ev(void* qthis); // 1
  // proto:  qreal QTextInlineObject::ascent();
extern void C_ZNK17QTextInlineObject6ascentEv(void* qthis); // 4
  // proto:  QRectF QTextInlineObject::rect();
extern void C_ZNK17QTextInlineObject4rectEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextLine)=16
type QTextLine struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextLayout)=8
type QTextLayout struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextInlineObject)=16
type QTextInlineObject struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// horizontalAdvance()
func (this *QTextLine) horizontalAdvance(args ...interface{}) () {
  // horizontalAdvance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine17horizontalAdvanceEv
    // invoke: qreal horizontalAdvance()
    C.C_ZNK9QTextLine17horizontalAdvanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "horizontalAdvance", args)
  }

}

// height()
func (this *QTextLine) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine6heightEv
    // invoke: qreal height()
    C.C_ZNK9QTextLine6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "height", args)
  }

}

// textStart()
func (this *QTextLine) textStart(args ...interface{}) () {
  // textStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine9textStartEv
    // invoke: int textStart()
    C.C_ZNK9QTextLine9textStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "textStart", args)
  }

}

// QTextLine()
func NewQTextLine(args ...interface{}) QTextLine {
  // QTextLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLineC1Ev
    // invoke: void QTextLine()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTextLineC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextLine", "QTextLine", args)
  }

  return QTextLine{}
}

// setLeadingIncluded(_Bool)
func (this *QTextLine) setLeadingIncluded(args ...interface{}) () {
  // setLeadingIncluded(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLine18setLeadingIncludedEb
    // invoke: void setLeadingIncluded(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine18setLeadingIncludedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setLeadingIncluded", args)
  }

}

// ascent()
func (this *QTextLine) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine6ascentEv
    // invoke: qreal ascent()
    C.C_ZNK9QTextLine6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "ascent", args)
  }

}

// x()
func (this *QTextLine) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine1xEv
    // invoke: qreal x()
    C.C_ZNK9QTextLine1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "x", args)
  }

}

// naturalTextWidth()
func (this *QTextLine) naturalTextWidth(args ...interface{}) () {
  // naturalTextWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine16naturalTextWidthEv
    // invoke: qreal naturalTextWidth()
    C.C_ZNK9QTextLine16naturalTextWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextWidth", args)
  }

}

// setNumColumns(int)
func (this *QTextLine) setNumColumns(args ...interface{}) () {
  // setNumColumns(int)
  // setNumColumns(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLine13setNumColumnsEi
    // invoke: void setNumColumns(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine13setNumColumnsEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QTextLine13setNumColumnsEid
    // invoke: void setNumColumns(int, qreal)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTextLine13setNumColumnsEid(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLine", "setNumColumns", args)
  }

}

// descent()
func (this *QTextLine) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine7descentEv
    // invoke: qreal descent()
    C.C_ZNK9QTextLine7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "descent", args)
  }

}

// leading()
func (this *QTextLine) leading(args ...interface{}) () {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine7leadingEv
    // invoke: qreal leading()
    C.C_ZNK9QTextLine7leadingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "leading", args)
  }

}

// naturalTextRect()
func (this *QTextLine) naturalTextRect(args ...interface{}) () {
  // naturalTextRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine15naturalTextRectEv
    // invoke: QRectF naturalTextRect()
    C.C_ZNK9QTextLine15naturalTextRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextRect", args)
  }

}

// width()
func (this *QTextLine) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine5widthEv
    // invoke: qreal width()
    C.C_ZNK9QTextLine5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "width", args)
  }

}

// setLineWidth(qreal)
func (this *QTextLine) setLineWidth(args ...interface{}) () {
  // setLineWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLine12setLineWidthEd
    // invoke: void setLineWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine12setLineWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setLineWidth", args)
  }

}

// isValid()
func (this *QTextLine) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine7isValidEv
    // invoke: bool isValid()
    C.C_ZNK9QTextLine7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "isValid", args)
  }

}

// lineNumber()
func (this *QTextLine) lineNumber(args ...interface{}) () {
  // lineNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine10lineNumberEv
    // invoke: int lineNumber()
    C.C_ZNK9QTextLine10lineNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "lineNumber", args)
  }

}

// setPosition(const class QPointF &)
func (this *QTextLine) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLine11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine11setPositionERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setPosition", args)
  }

}

// glyphRuns(int, int)
func (this *QTextLine) glyphRuns(args ...interface{}) () {
  // glyphRuns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine9glyphRunsEii
    // invoke: QList<QGlyphRun> glyphRuns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK9QTextLine9glyphRunsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLine", "glyphRuns", args)
  }

}

// rect()
func (this *QTextLine) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine4rectEv
    // invoke: QRectF rect()
    C.C_ZNK9QTextLine4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "rect", args)
  }

}

// leadingIncluded()
func (this *QTextLine) leadingIncluded(args ...interface{}) () {
  // leadingIncluded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine15leadingIncludedEv
    // invoke: bool leadingIncluded()
    C.C_ZNK9QTextLine15leadingIncludedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "leadingIncluded", args)
  }

}

// textLength()
func (this *QTextLine) textLength(args ...interface{}) () {
  // textLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine10textLengthEv
    // invoke: int textLength()
    C.C_ZNK9QTextLine10textLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "textLength", args)
  }

}

// y()
func (this *QTextLine) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine1yEv
    // invoke: qreal y()
    C.C_ZNK9QTextLine1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "y", args)
  }

}

// position()
func (this *QTextLine) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextLine8positionEv
    // invoke: QPointF position()
    C.C_ZNK9QTextLine8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "position", args)
  }

}

// setCacheEnabled(_Bool)
func (this *QTextLayout) setCacheEnabled(args ...interface{}) () {
  // setCacheEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout15setCacheEnabledEb
    // invoke: void setCacheEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout15setCacheEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setCacheEnabled", args)
  }

}

// boundingRect()
func (this *QTextLayout) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout12boundingRectEv
    // invoke: QRectF boundingRect()
    C.C_ZNK11QTextLayout12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "boundingRect", args)
  }

}

// text()
func (this *QTextLayout) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout4textEv
    // invoke: QString text()
    C.C_ZNK11QTextLayout4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "text", args)
  }

}

// clearAdditionalFormats()
func (this *QTextLayout) clearAdditionalFormats(args ...interface{}) () {
  // clearAdditionalFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout22clearAdditionalFormatsEv
    // invoke: void clearAdditionalFormats()
    C.C_ZN11QTextLayout22clearAdditionalFormatsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "clearAdditionalFormats", args)
  }

}

// lineCount()
func (this *QTextLayout) lineCount(args ...interface{}) () {
  // lineCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout9lineCountEv
    // invoke: int lineCount()
    C.C_ZNK11QTextLayout9lineCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "lineCount", args)
  }

}

// clearLayout()
func (this *QTextLayout) clearLayout(args ...interface{}) () {
  // clearLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout11clearLayoutEv
    // invoke: void clearLayout()
    C.C_ZN11QTextLayout11clearLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "clearLayout", args)
  }

}

// createLine()
func (this *QTextLayout) createLine(args ...interface{}) () {
  // createLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout10createLineEv
    // invoke: QTextLine createLine()
    C.C_ZN11QTextLayout10createLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "createLine", args)
  }

}

// font()
func (this *QTextLayout) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout4fontEv
    // invoke: QFont font()
    C.C_ZNK11QTextLayout4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "font", args)
  }

}

// textOption()
func (this *QTextLayout) textOption(args ...interface{}) () {
  // textOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout10textOptionEv
    // invoke: const QTextOption & textOption()
    C.C_ZNK11QTextLayout10textOptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "textOption", args)
  }

}

// lineAt(int)
func (this *QTextLayout) lineAt(args ...interface{}) () {
  // lineAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout6lineAtEi
    // invoke: QTextLine lineAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLayout6lineAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "lineAt", args)
  }

}

// setRawFont(const class QRawFont &)
func (this *QTextLayout) setRawFont(args ...interface{}) () {
  // setRawFont(const class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRawFont{}) // "const QRawFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout10setRawFontERK8QRawFont
    // invoke: void setRawFont(const class QRawFont &)
    var arg0 = args[0].(QRawFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout10setRawFontERK8QRawFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setRawFont", args)
  }

}

// beginLayout()
func (this *QTextLayout) beginLayout(args ...interface{}) () {
  // beginLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout11beginLayoutEv
    // invoke: void beginLayout()
    C.C_ZN11QTextLayout11beginLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "beginLayout", args)
  }

}

// glyphRuns(int, int)
func (this *QTextLayout) glyphRuns(args ...interface{}) () {
  // glyphRuns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout9glyphRunsEii
    // invoke: QList<QGlyphRun> glyphRuns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QTextLayout9glyphRunsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLayout", "glyphRuns", args)
  }

}

// rightCursorPosition(int)
func (this *QTextLayout) rightCursorPosition(args ...interface{}) () {
  // rightCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout19rightCursorPositionEi
    // invoke: int rightCursorPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLayout19rightCursorPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "rightCursorPosition", args)
  }

}

// ~QTextLayout()
func (this *QTextLayout) FreeQTextLayout(args ...interface{}) () {
  // ~QTextLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayoutD0Ev
    // invoke: void ~QTextLayout()
    C.C_ZN11QTextLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "~QTextLayout", args)
  }

}

// setTextOption(const class QTextOption &)
func (this *QTextLayout) setTextOption(args ...interface{}) () {
  // setTextOption(const class QTextOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout13setTextOptionERK11QTextOption
    // invoke: void setTextOption(const class QTextOption &)
    var arg0 = args[0].(QTextOption).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout13setTextOptionERK11QTextOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setTextOption", args)
  }

}

// isValidCursorPosition(int)
func (this *QTextLayout) isValidCursorPosition(args ...interface{}) () {
  // isValidCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout21isValidCursorPositionEi
    // invoke: bool isValidCursorPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLayout21isValidCursorPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "isValidCursorPosition", args)
  }

}

// setPreeditArea(int, const class QString &)
func (this *QTextLayout) setPreeditArea(args ...interface{}) () {
  // setPreeditArea(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout14setPreeditAreaEiRK7QString
    // invoke: void setPreeditArea(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextLayout14setPreeditAreaEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLayout", "setPreeditArea", args)
  }

}

// engine()
func (this *QTextLayout) engine(args ...interface{}) () {
  // engine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout6engineEv
    // invoke: QTextEngine * engine()
    C.C_ZNK11QTextLayout6engineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "engine", args)
  }

}

// minimumWidth()
func (this *QTextLayout) minimumWidth(args ...interface{}) () {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout12minimumWidthEv
    // invoke: qreal minimumWidth()
    C.C_ZNK11QTextLayout12minimumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "minimumWidth", args)
  }

}

// QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
func NewQTextLayout(args ...interface{}) QTextLayout {
  // QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
  // QTextLayout(const class QTextBlock &)
  // QTextLayout()
  // QTextLayout(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[0][2] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayoutC1ERK7QStringRK5QFontP12QPaintDevice
    // invoke: void QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPaintDevice).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice(qthis, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN11QTextLayoutC1ERK10QTextBlock
    // invoke: void QTextLayout(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextLayoutC2ERK10QTextBlock(qthis, arg0)
  case 2:
    // invoke: _ZN11QTextLayoutC1Ev
    // invoke: void QTextLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextLayoutC2Ev(qthis)
  case 3:
    // invoke: _ZN11QTextLayoutC1ERK7QString
    // invoke: void QTextLayout(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextLayoutC2ERK7QString(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "QTextLayout", args)
  }

  return QTextLayout{}
}

// leftCursorPosition(int)
func (this *QTextLayout) leftCursorPosition(args ...interface{}) () {
  // leftCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout18leftCursorPositionEi
    // invoke: int leftCursorPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLayout18leftCursorPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "leftCursorPosition", args)
  }

}

// endLayout()
func (this *QTextLayout) endLayout(args ...interface{}) () {
  // endLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout9endLayoutEv
    // invoke: void endLayout()
    C.C_ZN11QTextLayout9endLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "endLayout", args)
  }

}

// maximumWidth()
func (this *QTextLayout) maximumWidth(args ...interface{}) () {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout12maximumWidthEv
    // invoke: qreal maximumWidth()
    C.C_ZNK11QTextLayout12maximumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "maximumWidth", args)
  }

}

// setFont(const class QFont &)
func (this *QTextLayout) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setFont", args)
  }

}

// lineForTextPosition(int)
func (this *QTextLayout) lineForTextPosition(args ...interface{}) () {
  // lineForTextPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout19lineForTextPositionEi
    // invoke: QTextLine lineForTextPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextLayout19lineForTextPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "lineForTextPosition", args)
  }

}

// preeditAreaPosition()
func (this *QTextLayout) preeditAreaPosition(args ...interface{}) () {
  // preeditAreaPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout19preeditAreaPositionEv
    // invoke: int preeditAreaPosition()
    C.C_ZNK11QTextLayout19preeditAreaPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaPosition", args)
  }

}

// setFlags(int)
func (this *QTextLayout) setFlags(args ...interface{}) () {
  // setFlags(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout8setFlagsEi
    // invoke: void setFlags(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout8setFlagsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setFlags", args)
  }

}

// setPosition(const class QPointF &)
func (this *QTextLayout) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout11setPositionERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setPosition", args)
  }

}

// drawCursor(class QPainter *, const class QPointF &, int)
func (this *QTextLayout) drawCursor(args ...interface{}) () {
  // drawCursor(class QPainter *, const class QPointF &, int)
  // drawCursor(class QPainter *, const class QPointF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi
    // invoke: void drawCursor(class QPainter *, const class QPointF &, int)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii
    // invoke: void drawCursor(class QPainter *, const class QPointF &, int, int)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextLayout", "drawCursor", args)
  }

}

// cacheEnabled()
func (this *QTextLayout) cacheEnabled(args ...interface{}) () {
  // cacheEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout12cacheEnabledEv
    // invoke: bool cacheEnabled()
    C.C_ZNK11QTextLayout12cacheEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "cacheEnabled", args)
  }

}

// setText(const class QString &)
func (this *QTextLayout) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextLayout7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setText", args)
  }

}

// cursorMoveStyle()
func (this *QTextLayout) cursorMoveStyle(args ...interface{}) () {
  // cursorMoveStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout15cursorMoveStyleEv
    // invoke: Qt::CursorMoveStyle cursorMoveStyle()
    C.C_ZNK11QTextLayout15cursorMoveStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "cursorMoveStyle", args)
  }

}

// preeditAreaText()
func (this *QTextLayout) preeditAreaText(args ...interface{}) () {
  // preeditAreaText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout15preeditAreaTextEv
    // invoke: QString preeditAreaText()
    C.C_ZNK11QTextLayout15preeditAreaTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaText", args)
  }

}

// additionalFormats()
func (this *QTextLayout) additionalFormats(args ...interface{}) () {
  // additionalFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout17additionalFormatsEv
    // invoke: QList<QTextLayout::FormatRange> additionalFormats()
    C.C_ZNK11QTextLayout17additionalFormatsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "additionalFormats", args)
  }

}

// position()
func (this *QTextLayout) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextLayout8positionEv
    // invoke: QPointF position()
    C.C_ZNK11QTextLayout8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "position", args)
  }

}

// textPosition()
func (this *QTextInlineObject) textPosition(args ...interface{}) () {
  // textPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject12textPositionEv
    // invoke: int textPosition()
    C.C_ZNK17QTextInlineObject12textPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "textPosition", args)
  }

}

// descent()
func (this *QTextInlineObject) descent(args ...interface{}) () {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject7descentEv
    // invoke: qreal descent()
    C.C_ZNK17QTextInlineObject7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "descent", args)
  }

}

// setWidth(qreal)
func (this *QTextInlineObject) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QTextInlineObject8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setWidth", args)
  }

}

// format()
func (this *QTextInlineObject) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject6formatEv
    // invoke: QTextFormat format()
    C.C_ZNK17QTextInlineObject6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "format", args)
  }

}

// setDescent(qreal)
func (this *QTextInlineObject) setDescent(args ...interface{}) () {
  // setDescent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QTextInlineObject10setDescentEd
    // invoke: void setDescent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject10setDescentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setDescent", args)
  }

}

// isValid()
func (this *QTextInlineObject) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject7isValidEv
    // invoke: bool isValid()
    C.C_ZNK17QTextInlineObject7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "isValid", args)
  }

}

// textDirection()
func (this *QTextInlineObject) textDirection(args ...interface{}) () {
  // textDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject13textDirectionEv
    // invoke: Qt::LayoutDirection textDirection()
    C.C_ZNK17QTextInlineObject13textDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "textDirection", args)
  }

}

// height()
func (this *QTextInlineObject) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject6heightEv
    // invoke: qreal height()
    C.C_ZNK17QTextInlineObject6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "height", args)
  }

}

// width()
func (this *QTextInlineObject) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject5widthEv
    // invoke: qreal width()
    C.C_ZNK17QTextInlineObject5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "width", args)
  }

}

// setAscent(qreal)
func (this *QTextInlineObject) setAscent(args ...interface{}) () {
  // setAscent(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QTextInlineObject9setAscentEd
    // invoke: void setAscent(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject9setAscentEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setAscent", args)
  }

}

// formatIndex()
func (this *QTextInlineObject) formatIndex(args ...interface{}) () {
  // formatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject11formatIndexEv
    // invoke: int formatIndex()
    C.C_ZNK17QTextInlineObject11formatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "formatIndex", args)
  }

}

// QTextInlineObject()
func NewQTextInlineObject(args ...interface{}) QTextInlineObject {
  // QTextInlineObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QTextInlineObjectC1Ev
    // invoke: void QTextInlineObject()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN17QTextInlineObjectC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "QTextInlineObject", args)
  }

  return QTextInlineObject{}
}

// ascent()
func (this *QTextInlineObject) ascent(args ...interface{}) () {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject6ascentEv
    // invoke: qreal ascent()
    C.C_ZNK17QTextInlineObject6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "ascent", args)
  }

}

// rect()
func (this *QTextInlineObject) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QTextInlineObject4rectEv
    // invoke: QRectF rect()
    C.C_ZNK17QTextInlineObject4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "rect", args)
  }

}

// <= body block end

