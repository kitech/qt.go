package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern double C_ZNK9QTextLine17horizontalAdvanceEv(void* qthis); // 4
  // proto:  qreal QTextLine::height();
extern double C_ZNK9QTextLine6heightEv(void* qthis); // 4
  // proto:  int QTextLine::textStart();
extern int32_t C_ZNK9QTextLine9textStartEv(void* qthis); // 4
  // proto:  void QTextLine::QTextLine();
extern void* C_ZN9QTextLineC2Ev(); // 1
  // proto:  void QTextLine::setLeadingIncluded(bool included);
extern void C_ZN9QTextLine18setLeadingIncludedEb(void* qthis, bool arg0); // 4
  // proto:  qreal QTextLine::ascent();
extern double C_ZNK9QTextLine6ascentEv(void* qthis); // 4
  // proto:  qreal QTextLine::x();
extern void C_ZNK9QTextLine1xEv(void* qthis); // 4
  // proto:  qreal QTextLine::naturalTextWidth();
extern double C_ZNK9QTextLine16naturalTextWidthEv(void* qthis); // 4
  // proto:  void QTextLine::setNumColumns(int columns);
extern void C_ZN9QTextLine13setNumColumnsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLine::setNumColumns(int columns, qreal alignmentWidth);
extern void C_ZN9QTextLine13setNumColumnsEid(void* qthis, int32_t arg0, double arg1); // 4
  // proto:  qreal QTextLine::descent();
extern double C_ZNK9QTextLine7descentEv(void* qthis); // 4
  // proto:  qreal QTextLine::leading();
extern double C_ZNK9QTextLine7leadingEv(void* qthis); // 4
  // proto:  QRectF QTextLine::naturalTextRect();
extern void* C_ZNK9QTextLine15naturalTextRectEv(void* qthis); // 4
  // proto:  qreal QTextLine::width();
extern double C_ZNK9QTextLine5widthEv(void* qthis); // 4
  // proto:  void QTextLine::setLineWidth(qreal width);
extern void C_ZN9QTextLine12setLineWidthEd(void* qthis, double arg0); // 4
  // proto:  bool QTextLine::isValid();
extern bool C_ZNK9QTextLine7isValidEv(void* qthis); // 2
  // proto:  int QTextLine::lineNumber();
extern int32_t C_ZNK9QTextLine10lineNumberEv(void* qthis); // 2
  // proto:  void QTextLine::setPosition(const QPointF & pos);
extern void C_ZN9QTextLine11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QList<QGlyphRun> QTextLine::glyphRuns(int from, int length);
extern void C_ZNK9QTextLine9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QRectF QTextLine::rect();
extern void* C_ZNK9QTextLine4rectEv(void* qthis); // 4
  // proto:  bool QTextLine::leadingIncluded();
extern bool C_ZNK9QTextLine15leadingIncludedEv(void* qthis); // 4
  // proto:  int QTextLine::textLength();
extern int32_t C_ZNK9QTextLine10textLengthEv(void* qthis); // 4
  // proto:  qreal QTextLine::y();
extern void C_ZNK9QTextLine1yEv(void* qthis); // 4
  // proto:  QPointF QTextLine::position();
extern void* C_ZNK9QTextLine8positionEv(void* qthis); // 4
  // proto:  void QTextLayout::setCacheEnabled(bool enable);
extern void C_ZN11QTextLayout15setCacheEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QRectF QTextLayout::boundingRect();
extern void* C_ZNK11QTextLayout12boundingRectEv(void* qthis); // 4
  // proto:  QString QTextLayout::text();
extern void* C_ZNK11QTextLayout4textEv(void* qthis); // 4
  // proto:  void QTextLayout::clearAdditionalFormats();
extern void C_ZN11QTextLayout22clearAdditionalFormatsEv(void* qthis); // 4
  // proto:  int QTextLayout::lineCount();
extern int32_t C_ZNK11QTextLayout9lineCountEv(void* qthis); // 4
  // proto:  void QTextLayout::clearLayout();
extern void C_ZN11QTextLayout11clearLayoutEv(void* qthis); // 4
  // proto:  QTextLine QTextLayout::createLine();
extern void* C_ZN11QTextLayout10createLineEv(void* qthis); // 4
  // proto:  QFont QTextLayout::font();
extern void* C_ZNK11QTextLayout4fontEv(void* qthis); // 4
  // proto:  const QTextOption & QTextLayout::textOption();
extern void* C_ZNK11QTextLayout10textOptionEv(void* qthis); // 4
  // proto:  QTextLine QTextLayout::lineAt(int i);
extern void* C_ZNK11QTextLayout6lineAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setRawFont(const QRawFont & rawFont);
extern void C_ZN11QTextLayout10setRawFontERK8QRawFont(void* qthis, void* arg0); // 4
  // proto:  void QTextLayout::beginLayout();
extern void C_ZN11QTextLayout11beginLayoutEv(void* qthis); // 4
  // proto:  QList<QGlyphRun> QTextLayout::glyphRuns(int from, int length);
extern void C_ZNK11QTextLayout9glyphRunsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTextLayout::rightCursorPosition(int oldPos);
extern int32_t C_ZNK11QTextLayout19rightCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::~QTextLayout();
extern void C_ZN11QTextLayoutD2Ev(void* qthis); // 4
  // proto:  void QTextLayout::setTextOption(const QTextOption & option);
extern void C_ZN11QTextLayout13setTextOptionERK11QTextOption(void* qthis, void* arg0); // 4
  // proto:  bool QTextLayout::isValidCursorPosition(int pos);
extern bool C_ZNK11QTextLayout21isValidCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setPreeditArea(int position, const QString & text);
extern void C_ZN11QTextLayout14setPreeditAreaEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTextEngine * QTextLayout::engine();
extern void C_ZNK11QTextLayout6engineEv(void* qthis); // 2
  // proto:  qreal QTextLayout::minimumWidth();
extern double C_ZNK11QTextLayout12minimumWidthEv(void* qthis); // 4
  // proto:  void QTextLayout::QTextLayout(const QString & text, const QFont & font, QPaintDevice * paintdevice);
extern void* C_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QTextLayout::QTextLayout(const QTextBlock & b);
extern void* C_ZN11QTextLayoutC2ERK10QTextBlock(void* arg0); // 3
  // proto:  void QTextLayout::QTextLayout();
extern void* C_ZN11QTextLayoutC2Ev(); // 3
  // proto:  void QTextLayout::QTextLayout(const QString & text);
extern void* C_ZN11QTextLayoutC2ERK7QString(void* arg0); // 3
  // proto:  int QTextLayout::leftCursorPosition(int oldPos);
extern int32_t C_ZNK11QTextLayout18leftCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::endLayout();
extern void C_ZN11QTextLayout9endLayoutEv(void* qthis); // 4
  // proto:  qreal QTextLayout::maximumWidth();
extern double C_ZNK11QTextLayout12maximumWidthEv(void* qthis); // 4
  // proto:  void QTextLayout::setFont(const QFont & f);
extern void C_ZN11QTextLayout7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  QTextLine QTextLayout::lineForTextPosition(int pos);
extern void* C_ZNK11QTextLayout19lineForTextPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextLayout::preeditAreaPosition();
extern int32_t C_ZNK11QTextLayout19preeditAreaPositionEv(void* qthis); // 4
  // proto:  void QTextLayout::setFlags(int flags);
extern void C_ZN11QTextLayout8setFlagsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextLayout::setPosition(const QPointF & p);
extern void C_ZN11QTextLayout11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition);
extern void C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QTextLayout::drawCursor(QPainter * p, const QPointF & pos, int cursorPosition, int width);
extern void C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  bool QTextLayout::cacheEnabled();
extern bool C_ZNK11QTextLayout12cacheEnabledEv(void* qthis); // 4
  // proto:  void QTextLayout::setText(const QString & string);
extern void C_ZN11QTextLayout7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::CursorMoveStyle QTextLayout::cursorMoveStyle();
extern void C_ZNK11QTextLayout15cursorMoveStyleEv(void* qthis); // 4
  // proto:  QString QTextLayout::preeditAreaText();
extern void* C_ZNK11QTextLayout15preeditAreaTextEv(void* qthis); // 4
  // proto:  QList<QTextLayout::FormatRange> QTextLayout::additionalFormats();
extern void C_ZNK11QTextLayout17additionalFormatsEv(void* qthis); // 4
  // proto:  QPointF QTextLayout::position();
extern void* C_ZNK11QTextLayout8positionEv(void* qthis); // 4
  // proto:  int QTextInlineObject::textPosition();
extern int32_t C_ZNK17QTextInlineObject12textPositionEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::descent();
extern double C_ZNK17QTextInlineObject7descentEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setWidth(qreal w);
extern void C_ZN17QTextInlineObject8setWidthEd(void* qthis, double arg0); // 4
  // proto:  QTextFormat QTextInlineObject::format();
extern void* C_ZNK17QTextInlineObject6formatEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setDescent(qreal d);
extern void C_ZN17QTextInlineObject10setDescentEd(void* qthis, double arg0); // 4
  // proto:  bool QTextInlineObject::isValid();
extern bool C_ZNK17QTextInlineObject7isValidEv(void* qthis); // 2
  // proto:  Qt::LayoutDirection QTextInlineObject::textDirection();
extern void C_ZNK17QTextInlineObject13textDirectionEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::height();
extern double C_ZNK17QTextInlineObject6heightEv(void* qthis); // 4
  // proto:  qreal QTextInlineObject::width();
extern double C_ZNK17QTextInlineObject5widthEv(void* qthis); // 4
  // proto:  void QTextInlineObject::setAscent(qreal a);
extern void C_ZN17QTextInlineObject9setAscentEd(void* qthis, double arg0); // 4
  // proto:  int QTextInlineObject::formatIndex();
extern int32_t C_ZNK17QTextInlineObject11formatIndexEv(void* qthis); // 4
  // proto:  void QTextInlineObject::QTextInlineObject();
extern void* C_ZN17QTextInlineObjectC2Ev(); // 1
  // proto:  qreal QTextInlineObject::ascent();
extern double C_ZNK17QTextInlineObject6ascentEv(void* qthis); // 4
  // proto:  QRectF QTextInlineObject::rect();
extern void* C_ZNK17QTextInlineObject4rectEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextLayout)=8
type QTextLayout struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextInlineObject)=16
type QTextInlineObject struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// horizontalAdvance()
func (this *QTextLine) Horizontaladvance(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine17horizontalAdvanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "horizontalAdvance", args)
  }

  return
}

// height()
func (this *QTextLine) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "height", args)
  }

  return
}

// textStart()
func (this *QTextLine) Textstart(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine9textStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "textStart", args)
  }

  return
}

// QTextLine()
func NewQTextLine(args ...interface{}) *QTextLine {
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
    qthis = C.C_ZN9QTextLineC2Ev()
    return &QTextLine{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextLine", "QTextLine", args)
  }

  return nil // QTextLine{Qclsinst:qthis}
}

// setLeadingIncluded(_Bool)
func (this *QTextLine) Setleadingincluded(args ...interface{}) () {
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
    C.C_ZN9QTextLine18setLeadingIncludedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setLeadingIncluded", args)
  }

  return
}

// ascent()
func (this *QTextLine) Ascent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "ascent", args)
  }

  return
}

// x()
func (this *QTextLine) X(args ...interface{}) () {
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
    C.C_ZNK9QTextLine1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "x", args)
  }

  return
}

// naturalTextWidth()
func (this *QTextLine) Naturaltextwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine16naturalTextWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextWidth", args)
  }

  return
}

// setNumColumns(int)
func (this *QTextLine) Setnumcolumns(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine13setNumColumnsEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN9QTextLine13setNumColumnsEid
    // invoke: void setNumColumns(int, qreal)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QTextLine13setNumColumnsEid(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLine", "setNumColumns", args)
  }

  return
}

// descent()
func (this *QTextLine) Descent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "descent", args)
  }

  return
}

// leading()
func (this *QTextLine) Leading(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine7leadingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "leading", args)
  }

  return
}

// naturalTextRect()
func (this *QTextLine) Naturaltextrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine15naturalTextRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "naturalTextRect", args)
  }

  return
}

// width()
func (this *QTextLine) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "width", args)
  }

  return
}

// setLineWidth(qreal)
func (this *QTextLine) Setlinewidth(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine12setLineWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setLineWidth", args)
  }

  return
}

// isValid()
func (this *QTextLine) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "isValid", args)
  }

  return
}

// lineNumber()
func (this *QTextLine) Linenumber(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine10lineNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "lineNumber", args)
  }

  return
}

// setPosition(const class QPointF &)
func (this *QTextLine) Setposition(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextLine11setPositionERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLine", "setPosition", args)
  }

  return
}

// glyphRuns(int, int)
func (this *QTextLine) Glyphruns(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK9QTextLine9glyphRunsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLine", "glyphRuns", args)
  }

  return
}

// rect()
func (this *QTextLine) Rect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "rect", args)
  }

  return
}

// leadingIncluded()
func (this *QTextLine) Leadingincluded(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine15leadingIncludedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "leadingIncluded", args)
  }

  return
}

// textLength()
func (this *QTextLine) Textlength(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine10textLengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "textLength", args)
  }

  return
}

// y()
func (this *QTextLine) Y(args ...interface{}) () {
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
    C.C_ZNK9QTextLine1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLine", "y", args)
  }

  return
}

// position()
func (this *QTextLine) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTextLine8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLine", "position", args)
  }

  return
}

// setCacheEnabled(_Bool)
func (this *QTextLayout) Setcacheenabled(args ...interface{}) () {
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
    C.C_ZN11QTextLayout15setCacheEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setCacheEnabled", args)
  }

  return
}

// boundingRect()
func (this *QTextLayout) Boundingrect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "boundingRect", args)
  }

  return
}

// text()
func (this *QTextLayout) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "text", args)
  }

  return
}

// clearAdditionalFormats()
func (this *QTextLayout) Clearadditionalformats(args ...interface{}) () {
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
    C.C_ZN11QTextLayout22clearAdditionalFormatsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "clearAdditionalFormats", args)
  }

  return
}

// lineCount()
func (this *QTextLayout) Linecount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout9lineCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "lineCount", args)
  }

  return
}

// clearLayout()
func (this *QTextLayout) Clearlayout(args ...interface{}) () {
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
    C.C_ZN11QTextLayout11clearLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "clearLayout", args)
  }

  return
}

// createLine()
func (this *QTextLayout) Createline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextLayout10createLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLine{}) // "QTextLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "createLine", args)
  }

  return
}

// font()
func (this *QTextLayout) Font(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "font", args)
  }

  return
}

// textOption()
func (this *QTextLayout) Textoption(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout10textOptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "textOption", args)
  }

  return
}

// lineAt(int)
func (this *QTextLayout) Lineat(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLayout6lineAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLine{}) // "QTextLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "lineAt", args)
  }

  return
}

// setRawFont(const class QRawFont &)
func (this *QTextLayout) Setrawfont(args ...interface{}) () {
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
    var arg0 = args[0].(*QRawFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout10setRawFontERK8QRawFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setRawFont", args)
  }

  return
}

// beginLayout()
func (this *QTextLayout) Beginlayout(args ...interface{}) () {
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
    C.C_ZN11QTextLayout11beginLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "beginLayout", args)
  }

  return
}

// glyphRuns(int, int)
func (this *QTextLayout) Glyphruns(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK11QTextLayout9glyphRunsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLayout", "glyphRuns", args)
  }

  return
}

// rightCursorPosition(int)
func (this *QTextLayout) Rightcursorposition(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLayout19rightCursorPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "rightCursorPosition", args)
  }

  return
}

// ~QTextLayout()
func (this *QTextLayout) Freeqtextlayout(args ...interface{}) () {
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
    C.C_ZN11QTextLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "~QTextLayout", args)
  }

  return
}

// setTextOption(const class QTextOption &)
func (this *QTextLayout) Settextoption(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextOption).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout13setTextOptionERK11QTextOption(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setTextOption", args)
  }

  return
}

// isValidCursorPosition(int)
func (this *QTextLayout) Isvalidcursorposition(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLayout21isValidCursorPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "isValidCursorPosition", args)
  }

  return
}

// setPreeditArea(int, const class QString &)
func (this *QTextLayout) Setpreeditarea(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextLayout14setPreeditAreaEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextLayout", "setPreeditArea", args)
  }

  return
}

// engine()
func (this *QTextLayout) Engine(args ...interface{}) () {
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
    C.C_ZNK11QTextLayout6engineEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "engine", args)
  }

  return
}

// minimumWidth()
func (this *QTextLayout) Minimumwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout12minimumWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "minimumWidth", args)
  }

  return
}

// QTextLayout(const class QString &, const class QFont &, class QPaintDevice *)
func NewQTextLayout(args ...interface{}) *QTextLayout {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QFont).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice(arg0, arg1, arg2)
    return &QTextLayout{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QTextLayoutC1ERK10QTextBlock
    // invoke: void QTextLayout(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextLayoutC2ERK10QTextBlock(arg0)
    return &QTextLayout{Qclsinst:qthis}
  case 2:
    // invoke: _ZN11QTextLayoutC1Ev
    // invoke: void QTextLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextLayoutC2Ev()
    return &QTextLayout{Qclsinst:qthis}
  case 3:
    // invoke: _ZN11QTextLayoutC1ERK7QString
    // invoke: void QTextLayout(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextLayoutC2ERK7QString(arg0)
    return &QTextLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextLayout", "QTextLayout", args)
  }

  return nil // QTextLayout{Qclsinst:qthis}
}

// leftCursorPosition(int)
func (this *QTextLayout) Leftcursorposition(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLayout18leftCursorPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "leftCursorPosition", args)
  }

  return
}

// endLayout()
func (this *QTextLayout) Endlayout(args ...interface{}) () {
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
    C.C_ZN11QTextLayout9endLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "endLayout", args)
  }

  return
}

// maximumWidth()
func (this *QTextLayout) Maximumwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout12maximumWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "maximumWidth", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QTextLayout) Setfont(args ...interface{}) () {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setFont", args)
  }

  return
}

// lineForTextPosition(int)
func (this *QTextLayout) Linefortextposition(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextLayout19lineForTextPositionEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextLine{}) // "QTextLine"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "lineForTextPosition", args)
  }

  return
}

// preeditAreaPosition()
func (this *QTextLayout) Preeditareaposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout19preeditAreaPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaPosition", args)
  }

  return
}

// setFlags(int)
func (this *QTextLayout) Setflags(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout8setFlagsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setFlags", args)
  }

  return
}

// setPosition(const class QPointF &)
func (this *QTextLayout) Setposition(args ...interface{}) () {
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
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout11setPositionERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setPosition", args)
  }

  return
}

// drawCursor(class QPainter *, const class QPointF &, int)
func (this *QTextLayout) Drawcursor(args ...interface{}) () {
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
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii
    // invoke: void drawCursor(class QPainter *, const class QPointF &, int, int)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextLayout", "drawCursor", args)
  }

  return
}

// cacheEnabled()
func (this *QTextLayout) Cacheenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout12cacheEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "cacheEnabled", args)
  }

  return
}

// setText(const class QString &)
func (this *QTextLayout) Settext(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextLayout7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextLayout", "setText", args)
  }

  return
}

// cursorMoveStyle()
func (this *QTextLayout) Cursormovestyle(args ...interface{}) () {
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
    C.C_ZNK11QTextLayout15cursorMoveStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "cursorMoveStyle", args)
  }

  return
}

// preeditAreaText()
func (this *QTextLayout) Preeditareatext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout15preeditAreaTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "preeditAreaText", args)
  }

  return
}

// additionalFormats()
func (this *QTextLayout) Additionalformats(args ...interface{}) () {
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
    C.C_ZNK11QTextLayout17additionalFormatsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextLayout", "additionalFormats", args)
  }

  return
}

// position()
func (this *QTextLayout) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextLayout8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextLayout", "position", args)
  }

  return
}

// textPosition()
func (this *QTextInlineObject) Textposition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject12textPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "textPosition", args)
  }

  return
}

// descent()
func (this *QTextInlineObject) Descent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "descent", args)
  }

  return
}

// setWidth(qreal)
func (this *QTextInlineObject) Setwidth(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject8setWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setWidth", args)
  }

  return
}

// format()
func (this *QTextInlineObject) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFormat{}) // "QTextFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "format", args)
  }

  return
}

// setDescent(qreal)
func (this *QTextInlineObject) Setdescent(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject10setDescentEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setDescent", args)
  }

  return
}

// isValid()
func (this *QTextInlineObject) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "isValid", args)
  }

  return
}

// textDirection()
func (this *QTextInlineObject) Textdirection(args ...interface{}) () {
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
    C.C_ZNK17QTextInlineObject13textDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "textDirection", args)
  }

  return
}

// height()
func (this *QTextInlineObject) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "height", args)
  }

  return
}

// width()
func (this *QTextInlineObject) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "width", args)
  }

  return
}

// setAscent(qreal)
func (this *QTextInlineObject) Setascent(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN17QTextInlineObject9setAscentEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextInlineObject", "setAscent", args)
  }

  return
}

// formatIndex()
func (this *QTextInlineObject) Formatindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject11formatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "formatIndex", args)
  }

  return
}

// QTextInlineObject()
func NewQTextInlineObject(args ...interface{}) *QTextInlineObject {
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
    qthis = C.C_ZN17QTextInlineObjectC2Ev()
    return &QTextInlineObject{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextInlineObject", "QTextInlineObject", args)
  }

  return nil // QTextInlineObject{Qclsinst:qthis}
}

// ascent()
func (this *QTextInlineObject) Ascent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "ascent", args)
  }

  return
}

// rect()
func (this *QTextInlineObject) Rect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK17QTextInlineObject4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextInlineObject", "rect", args)
  }

  return
}

// <= body block end

