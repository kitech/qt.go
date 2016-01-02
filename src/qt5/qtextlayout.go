package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QTextLine::ascent();
extern void _ZNK9QTextLine6ascentEv(void* qthis);
  // proto:  qreal QTextLine::leading();
extern void _ZNK9QTextLine7leadingEv(void* qthis);
  // proto:  int QTextLine::textStart();
extern void _ZNK9QTextLine9textStartEv(void* qthis);
  // proto:  bool QTextLine::leadingIncluded();
extern void _ZNK9QTextLine15leadingIncludedEv(void* qthis);
  // proto:  qreal QTextLine::x();
extern void _ZNK9QTextLine1xEv(void* qthis);
  // proto:  qreal QTextLine::height();
extern void _ZNK9QTextLine6heightEv(void* qthis);
  // proto:  qreal QTextLine::y();
extern void _ZNK9QTextLine1yEv(void* qthis);
  // proto:  qreal QTextLine::horizontalAdvance();
extern void _ZNK9QTextLine17horizontalAdvanceEv(void* qthis);
  // proto:  QRectF QTextLine::naturalTextRect();
extern void _ZNK9QTextLine15naturalTextRectEv(void* qthis);
  // proto:  void QTextLine::setNumColumns(int columns, qreal alignmentWidth);
extern void _ZN9QTextLine13setNumColumnsEid(void* qthis, int arg0, double arg1);
  // proto:  qreal QTextLine::width();
extern void _ZNK9QTextLine5widthEv(void* qthis);
  // proto:  void QTextLine::setLeadingIncluded(bool included);
extern void _ZN9QTextLine18setLeadingIncludedEb(void* qthis, bool arg0);
  // proto:  void QTextLine::setPosition(const QPointF & pos);
extern void _ZN9QTextLine11setPositionERK7QPointF(void* qthis, void* arg0);
  // proto:  int QTextLine::lineNumber();
extern void _ZNK9QTextLine10lineNumberEv(void* qthis);
  // proto:  QRectF QTextLine::rect();
extern void _ZNK9QTextLine4rectEv(void* qthis);
  // proto:  void QTextLine::QTextLine();
extern void* dector_ZN9QTextLineC1Ev();
extern void demth_ZN9QTextLineC1Ev(void* qthis);
  // proto:  void QTextLine::setNumColumns(int columns);
extern void _ZN9QTextLine13setNumColumnsEi(void* qthis, int arg0);
  // proto:  int QTextLine::textLength();
extern void _ZNK9QTextLine10textLengthEv(void* qthis);
  // proto:  QPointF QTextLine::position();
extern void _ZNK9QTextLine8positionEv(void* qthis);
  // proto:  QList<QGlyphRun> QTextLine::glyphRuns(int from, int length);
extern void _ZNK9QTextLine9glyphRunsEii(void* qthis, int arg0, int arg1);
  // proto:  qreal QTextLine::descent();
extern void _ZNK9QTextLine7descentEv(void* qthis);
  // proto:  qreal QTextLine::naturalTextWidth();
extern void _ZNK9QTextLine16naturalTextWidthEv(void* qthis);
  // proto:  void QTextLine::setLineWidth(qreal width);
extern void _ZN9QTextLine12setLineWidthEd(void* qthis, double arg0);
  // proto:  bool QTextLine::isValid();
extern void demth_ZNK9QTextLine7isValidEv(void* qthis);
  // proto:  void QTextLayout::setFont(const QFont & f);
extern void _ZN11QTextLayout7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QTextLayout::setText(const QString & string);
extern void _ZN11QTextLayout7setTextERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextLayout::isValidCursorPosition(int pos);
extern void _ZNK11QTextLayout21isValidCursorPositionEi(void* qthis, int arg0);
  // proto:  QRectF QTextLayout::boundingRect();
extern void _ZNK11QTextLayout12boundingRectEv(void* qthis);
  // proto:  void QTextLayout::setRawFont(const QRawFont & rawFont);
extern void _ZN11QTextLayout10setRawFontERK8QRawFont(void* qthis, void* arg0);
  // proto:  void QTextLayout::setTextOption(const QTextOption & option);
extern void _ZN11QTextLayout13setTextOptionERK11QTextOption(void* qthis, void* arg0);
  // proto:  void QTextLayout::QTextLayout(const QString & text, const QFont & font, QPaintDevice * paintdevice);
extern void* dector_ZN11QTextLayoutC1ERK7QStringRK5QFontP12QPaintDevice(void* arg0, void* arg1, void* arg2);
extern void _ZN11QTextLayoutC1ERK7QStringRK5QFontP12QPaintDevice(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QTextLayout::setPosition(const QPointF & p);
extern void _ZN11QTextLayout11setPositionERK7QPointF(void* qthis, void* arg0);
  // proto:  QTextLine QTextLayout::lineForTextPosition(int pos);
extern void _ZNK11QTextLayout19lineForTextPositionEi(void* qthis, int arg0);
  // proto:  const QTextOption & QTextLayout::textOption();
extern void _ZNK11QTextLayout10textOptionEv(void* qthis);
  // proto:  QTextEngine * QTextLayout::engine();
extern void _ZNK11QTextLayout6engineEv(void* qthis);
  // proto:  int QTextLayout::preeditAreaPosition();
extern void _ZNK11QTextLayout19preeditAreaPositionEv(void* qthis);
  // proto:  void QTextLayout::clearAdditionalFormats();
extern void _ZN11QTextLayout22clearAdditionalFormatsEv(void* qthis);
  // proto:  int QTextLayout::leftCursorPosition(int oldPos);
extern void _ZNK11QTextLayout18leftCursorPositionEi(void* qthis, int arg0);
  // proto:  int QTextLayout::lineCount();
extern void _ZNK11QTextLayout9lineCountEv(void* qthis);
  // proto:  void QTextLayout::~QTextLayout();
extern void _ZN11QTextLayoutD0Ev(void* qthis);
  // proto:  void QTextLayout::setCacheEnabled(bool enable);
extern void _ZN11QTextLayout15setCacheEnabledEb(void* qthis, bool arg0);
  // proto:  QTextLine QTextLayout::lineAt(int i);
extern void _ZNK11QTextLayout6lineAtEi(void* qthis, int arg0);
  // proto:  int QTextLayout::rightCursorPosition(int oldPos);
extern void _ZNK11QTextLayout19rightCursorPositionEi(void* qthis, int arg0);
  // proto:  void QTextLayout::QTextLayout(const QTextBlock & b);
extern void* dector_ZN11QTextLayoutC1ERK10QTextBlock(void* arg0);
extern void _ZN11QTextLayoutC1ERK10QTextBlock(void* qthis, void* arg0);
  // proto:  qreal QTextLayout::minimumWidth();
extern void _ZNK11QTextLayout12minimumWidthEv(void* qthis);
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition);
extern void _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  QFont QTextLayout::font();
extern void _ZNK11QTextLayout4fontEv(void* qthis);
  // proto:  void QTextLayout::setPreeditArea(int position, const QString & text);
extern void _ZN11QTextLayout14setPreeditAreaEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTextLayout::beginLayout();
extern void _ZN11QTextLayout11beginLayoutEv(void* qthis);
  // proto:  void QTextLayout::QTextLayout(const QString & text);
extern void* dector_ZN11QTextLayoutC1ERK7QString(void* arg0);
extern void _ZN11QTextLayoutC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QTextLayout::setFlags(int flags);
extern void _ZN11QTextLayout8setFlagsEi(void* qthis, int arg0);
  // proto:  QPointF QTextLayout::position();
extern void _ZNK11QTextLayout8positionEv(void* qthis);
  // proto:  void QTextLayout::clearLayout();
extern void _ZN11QTextLayout11clearLayoutEv(void* qthis);
  // proto:  bool QTextLayout::cacheEnabled();
extern void _ZNK11QTextLayout12cacheEnabledEv(void* qthis);
  // proto:  qreal QTextLayout::maximumWidth();
extern void _ZNK11QTextLayout12maximumWidthEv(void* qthis);
  // proto:  QString QTextLayout::text();
extern void _ZNK11QTextLayout4textEv(void* qthis);
  // proto:  void QTextLayout::QTextLayout(const QTextLayout & );
extern void* dector_ZN11QTextLayoutC1ERKS_(void* arg0);
extern void _ZN11QTextLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  QTextLine QTextLayout::createLine();
extern void _ZN11QTextLayout10createLineEv(void* qthis);
  // proto:  QString QTextLayout::preeditAreaText();
extern void _ZNK11QTextLayout15preeditAreaTextEv(void* qthis);
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition, int width);
extern void _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii(void* qthis, void* arg0, void* arg1, int arg2, int arg3);
  // proto:  void QTextLayout::endLayout();
extern void _ZN11QTextLayout9endLayoutEv(void* qthis);
  // proto:  void QTextLayout::QTextLayout();
extern void* dector_ZN11QTextLayoutC1Ev();
extern void _ZN11QTextLayoutC1Ev(void* qthis);
  // proto:  QList<QGlyphRun> QTextLayout::glyphRuns(int from, int length);
extern void _ZNK11QTextLayout9glyphRunsEii(void* qthis, int arg0, int arg1);
  // proto:  void QTextInlineObject::setAscent(qreal a);
extern void _ZN17QTextInlineObject9setAscentEd(void* qthis, double arg0);
  // proto:  qreal QTextInlineObject::width();
extern void _ZNK17QTextInlineObject5widthEv(void* qthis);
  // proto:  int QTextInlineObject::formatIndex();
extern void _ZNK17QTextInlineObject11formatIndexEv(void* qthis);
  // proto:  QRectF QTextInlineObject::rect();
extern void _ZNK17QTextInlineObject4rectEv(void* qthis);
  // proto:  int QTextInlineObject::textPosition();
extern void _ZNK17QTextInlineObject12textPositionEv(void* qthis);
  // proto:  void QTextInlineObject::setDescent(qreal d);
extern void _ZN17QTextInlineObject10setDescentEd(void* qthis, double arg0);
  // proto:  qreal QTextInlineObject::height();
extern void _ZNK17QTextInlineObject6heightEv(void* qthis);
  // proto:  bool QTextInlineObject::isValid();
extern void demth_ZNK17QTextInlineObject7isValidEv(void* qthis);
  // proto:  void QTextInlineObject::QTextInlineObject();
extern void* dector_ZN17QTextInlineObjectC1Ev();
extern void demth_ZN17QTextInlineObjectC1Ev(void* qthis);
  // proto:  QTextFormat QTextInlineObject::format();
extern void _ZNK17QTextInlineObject6formatEv(void* qthis);
  // proto:  qreal QTextInlineObject::descent();
extern void _ZNK17QTextInlineObject7descentEv(void* qthis);
  // proto:  qreal QTextInlineObject::ascent();
extern void _ZNK17QTextInlineObject6ascentEv(void* qthis);
  // proto:  void QTextInlineObject::setWidth(qreal w);
extern void _ZN17QTextInlineObject8setWidthEd(void* qthis, double arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextLayout)=8
type QTextLayout struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextInlineObject)=16
type QTextInlineObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qreal QTextLine::ascent();
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
  default:
    qtrt.ErrorResolve("QTextLine", "ascent", args)
  }

}

  // proto:  qreal QTextLine::leading();
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
  default:
    qtrt.ErrorResolve("QTextLine", "leading", args)
  }

}

  // proto:  int QTextLine::textStart();
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
  default:
    qtrt.ErrorResolve("QTextLine", "textStart", args)
  }

}

  // proto:  bool QTextLine::leadingIncluded();
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
  default:
    qtrt.ErrorResolve("QTextLine", "leadingIncluded", args)
  }

}

  // proto:  qreal QTextLine::x();
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
  default:
    qtrt.ErrorResolve("QTextLine", "x", args)
  }

}

  // proto:  qreal QTextLine::height();
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
  default:
    qtrt.ErrorResolve("QTextLine", "height", args)
  }

}

  // proto:  qreal QTextLine::y();
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
  default:
    qtrt.ErrorResolve("QTextLine", "y", args)
  }

}

  // proto:  qreal QTextLine::horizontalAdvance();
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
  default:
    qtrt.ErrorResolve("QTextLine", "horizontalAdvance", args)
  }

}

  // proto:  QRectF QTextLine::naturalTextRect();
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
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextRect", args)
  }

}

  // proto:  void QTextLine::setNumColumns(int columns, qreal alignmentWidth);
func (this *QTextLine) setNumColumns(args ...interface{}) () {
  // setNumColumns(int, qreal)
  // setNumColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextLine13setNumColumnsEid
  case 1:
    // invoke: _ZN9QTextLine13setNumColumnsEi
  default:
    qtrt.ErrorResolve("QTextLine", "setNumColumns", args)
  }

}

  // proto:  qreal QTextLine::width();
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
  default:
    qtrt.ErrorResolve("QTextLine", "width", args)
  }

}

  // proto:  void QTextLine::setLeadingIncluded(bool included);
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
  default:
    qtrt.ErrorResolve("QTextLine", "setLeadingIncluded", args)
  }

}

  // proto:  void QTextLine::setPosition(const QPointF & pos);
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
  default:
    qtrt.ErrorResolve("QTextLine", "setPosition", args)
  }

}

  // proto:  int QTextLine::lineNumber();
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
  default:
    qtrt.ErrorResolve("QTextLine", "lineNumber", args)
  }

}

  // proto:  QRectF QTextLine::rect();
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
  default:
    qtrt.ErrorResolve("QTextLine", "rect", args)
  }

}

  // proto:  void QTextLine::QTextLine();
func NewQTextLine(args ...interface{}) QTextLine {
  return QTextLine{}
}

  // proto:  int QTextLine::textLength();
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
  default:
    qtrt.ErrorResolve("QTextLine", "textLength", args)
  }

}

  // proto:  QPointF QTextLine::position();
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
  default:
    qtrt.ErrorResolve("QTextLine", "position", args)
  }

}

  // proto:  QList<QGlyphRun> QTextLine::glyphRuns(int from, int length);
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
  default:
    qtrt.ErrorResolve("QTextLine", "glyphRuns", args)
  }

}

  // proto:  qreal QTextLine::descent();
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
  default:
    qtrt.ErrorResolve("QTextLine", "descent", args)
  }

}

  // proto:  qreal QTextLine::naturalTextWidth();
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
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextWidth", args)
  }

}

  // proto:  void QTextLine::setLineWidth(qreal width);
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
  default:
    qtrt.ErrorResolve("QTextLine", "setLineWidth", args)
  }

}

  // proto:  bool QTextLine::isValid();
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
  default:
    qtrt.ErrorResolve("QTextLine", "isValid", args)
  }

}

  // proto:  void QTextLayout::setFont(const QFont & f);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setFont", args)
  }

}

  // proto:  void QTextLayout::setText(const QString & string);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setText", args)
  }

}

  // proto:  bool QTextLayout::isValidCursorPosition(int pos);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "isValidCursorPosition", args)
  }

}

  // proto:  QRectF QTextLayout::boundingRect();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "boundingRect", args)
  }

}

  // proto:  void QTextLayout::setRawFont(const QRawFont & rawFont);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setRawFont", args)
  }

}

  // proto:  void QTextLayout::setTextOption(const QTextOption & option);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setTextOption", args)
  }

}

  // proto:  void QTextLayout::QTextLayout(const QString & text, const QFont & font, QPaintDevice * paintdevice);
func NewQTextLayout(args ...interface{}) QTextLayout {
  return QTextLayout{}
}

  // proto:  void QTextLayout::setPosition(const QPointF & p);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setPosition", args)
  }

}

  // proto:  QTextLine QTextLayout::lineForTextPosition(int pos);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "lineForTextPosition", args)
  }

}

  // proto:  const QTextOption & QTextLayout::textOption();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "textOption", args)
  }

}

  // proto:  QTextEngine * QTextLayout::engine();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "engine", args)
  }

}

  // proto:  int QTextLayout::preeditAreaPosition();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaPosition", args)
  }

}

  // proto:  void QTextLayout::clearAdditionalFormats();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "clearAdditionalFormats", args)
  }

}

  // proto:  int QTextLayout::leftCursorPosition(int oldPos);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "leftCursorPosition", args)
  }

}

  // proto:  int QTextLayout::lineCount();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "lineCount", args)
  }

}

  // proto:  void QTextLayout::~QTextLayout();
func (this *QTextLayout) FreeQTextLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextLayout", "~QTextLayout", args)
  }

}

  // proto:  void QTextLayout::setCacheEnabled(bool enable);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setCacheEnabled", args)
  }

}

  // proto:  QTextLine QTextLayout::lineAt(int i);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "lineAt", args)
  }

}

  // proto:  int QTextLayout::rightCursorPosition(int oldPos);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "rightCursorPosition", args)
  }

}

  // proto:  qreal QTextLayout::minimumWidth();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "minimumWidth", args)
  }

}

  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition);
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
  case 1:
    // invoke: _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii
  default:
    qtrt.ErrorResolve("QTextLayout", "drawCursor", args)
  }

}

  // proto:  QFont QTextLayout::font();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "font", args)
  }

}

  // proto:  void QTextLayout::setPreeditArea(int position, const QString & text);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setPreeditArea", args)
  }

}

  // proto:  void QTextLayout::beginLayout();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "beginLayout", args)
  }

}

  // proto:  void QTextLayout::setFlags(int flags);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "setFlags", args)
  }

}

  // proto:  QPointF QTextLayout::position();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "position", args)
  }

}

  // proto:  void QTextLayout::clearLayout();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "clearLayout", args)
  }

}

  // proto:  bool QTextLayout::cacheEnabled();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "cacheEnabled", args)
  }

}

  // proto:  qreal QTextLayout::maximumWidth();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "maximumWidth", args)
  }

}

  // proto:  QString QTextLayout::text();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "text", args)
  }

}

  // proto:  QTextLine QTextLayout::createLine();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "createLine", args)
  }

}

  // proto:  QString QTextLayout::preeditAreaText();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaText", args)
  }

}

  // proto:  void QTextLayout::endLayout();
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
  default:
    qtrt.ErrorResolve("QTextLayout", "endLayout", args)
  }

}

  // proto:  QList<QGlyphRun> QTextLayout::glyphRuns(int from, int length);
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
  default:
    qtrt.ErrorResolve("QTextLayout", "glyphRuns", args)
  }

}

  // proto:  void QTextInlineObject::setAscent(qreal a);
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setAscent", args)
  }

}

  // proto:  qreal QTextInlineObject::width();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "width", args)
  }

}

  // proto:  int QTextInlineObject::formatIndex();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "formatIndex", args)
  }

}

  // proto:  QRectF QTextInlineObject::rect();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "rect", args)
  }

}

  // proto:  int QTextInlineObject::textPosition();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "textPosition", args)
  }

}

  // proto:  void QTextInlineObject::setDescent(qreal d);
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setDescent", args)
  }

}

  // proto:  qreal QTextInlineObject::height();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "height", args)
  }

}

  // proto:  bool QTextInlineObject::isValid();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "isValid", args)
  }

}

  // proto:  void QTextInlineObject::QTextInlineObject();
func NewQTextInlineObject(args ...interface{}) QTextInlineObject {
  return QTextInlineObject{}
}

  // proto:  QTextFormat QTextInlineObject::format();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "format", args)
  }

}

  // proto:  qreal QTextInlineObject::descent();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "descent", args)
  }

}

  // proto:  qreal QTextInlineObject::ascent();
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "ascent", args)
  }

}

  // proto:  void QTextInlineObject::setWidth(qreal w);
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
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setWidth", args)
  }

}

// <= body block end

