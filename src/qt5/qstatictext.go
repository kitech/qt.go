package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qstatictext.h
// dst-file: /src/gui/qstatictext.go
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
  // proto:  void QStaticText::swap(QStaticText & other);
extern void C_ZN11QStaticText4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QStaticText::prepare(const QTransform & matrix, const QFont & font);
extern void C_ZN11QStaticText7prepareERK10QTransformRK5QFont(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QStaticText::text();
extern void C_ZNK11QStaticText4textEv(void* qthis); // 4
  // proto:  void QStaticText::setText(const QString & text);
extern void C_ZN11QStaticText7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QStaticText::~QStaticText();
extern void C_ZN11QStaticTextD2Ev(void* qthis); // 4
  // proto:  void QStaticText::setTextWidth(qreal textWidth);
extern void C_ZN11QStaticText12setTextWidthEd(void* qthis, double arg0); // 4
  // proto:  void QStaticText::setTextOption(const QTextOption & textOption);
extern void C_ZN11QStaticText13setTextOptionERK11QTextOption(void* qthis, void* arg0); // 4
  // proto:  void QStaticText::QStaticText();
extern void* C_ZN11QStaticTextC2Ev(); // 3
  // proto:  void QStaticText::QStaticText(const QString & text);
extern void* C_ZN11QStaticTextC2ERK7QString(void* arg0); // 3
  // proto:  void QStaticText::QStaticText(const QStaticText & other);
extern void* C_ZN11QStaticTextC2ERKS_(void* arg0); // 3
  // proto:  qreal QStaticText::textWidth();
extern void C_ZNK11QStaticText9textWidthEv(void* qthis); // 4
  // proto:  Qt::TextFormat QStaticText::textFormat();
extern void C_ZNK11QStaticText10textFormatEv(void* qthis); // 4
  // proto:  QTextOption QStaticText::textOption();
extern void C_ZNK11QStaticText10textOptionEv(void* qthis); // 4
  // proto:  QStaticText::PerformanceHint QStaticText::performanceHint();
extern void C_ZNK11QStaticText15performanceHintEv(void* qthis); // 4
  // proto:  QSizeF QStaticText::size();
extern void C_ZNK11QStaticText4sizeEv(void* qthis); // 4
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

// class sizeof(QStaticText)=1
type QStaticText struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// swap(class QStaticText &)
func (this *QStaticText) swap(args ...interface{}) () {
  // swap(class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStaticText{}) // "QStaticText &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText4swapERS_
    // invoke: void swap(class QStaticText &)
    var arg0 = args[0].(QStaticText).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "swap", args)
  }

}

// prepare(const class QTransform &, const class QFont &)
func (this *QStaticText) prepare(args ...interface{}) () {
  // prepare(const class QTransform &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText7prepareERK10QTransformRK5QFont
    // invoke: void prepare(const class QTransform &, const class QFont &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QStaticText7prepareERK10QTransformRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStaticText", "prepare", args)
  }

}

// text()
func (this *QStaticText) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText4textEv
    // invoke: QString text()
    var ret = C.C_ZNK11QStaticText4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStaticText", "text", args)
  }

}

// setText(const class QString &)
func (this *QStaticText) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setText", args)
  }

}

// ~QStaticText()
func (this *QStaticText) FreeQStaticText(args ...interface{}) () {
  // ~QStaticText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticTextD0Ev
    // invoke: void ~QStaticText()
    C.C_ZN11QStaticTextD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "~QStaticText", args)
  }

}

// setTextWidth(qreal)
func (this *QStaticText) setTextWidth(args ...interface{}) () {
  // setTextWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText12setTextWidthEd
    // invoke: void setTextWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText12setTextWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextWidth", args)
  }

}

// setTextOption(const class QTextOption &)
func (this *QStaticText) setTextOption(args ...interface{}) () {
  // setTextOption(const class QTextOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText13setTextOptionERK11QTextOption
    // invoke: void setTextOption(const class QTextOption &)
    var arg0 = args[0].(QTextOption).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText13setTextOptionERK11QTextOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextOption", args)
  }

}

// QStaticText()
func NewQStaticText(args ...interface{}) *QStaticText {
  // QStaticText()
  // QStaticText(const class QString &)
  // QStaticText(const class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticTextC1Ev
    // invoke: void QStaticText()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStaticTextC2Ev()
    return &QStaticText{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QStaticTextC1ERK7QString
    // invoke: void QStaticText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStaticTextC2ERK7QString(arg0)
    return &QStaticText{qclsinst:qthis}
  case 2:
    // invoke: _ZN11QStaticTextC1ERKS_
    // invoke: void QStaticText(const class QStaticText &)
    var arg0 = args[0].(QStaticText).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStaticTextC2ERKS_(arg0)
    return &QStaticText{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStaticText", "QStaticText", args)
  }

  return nil // QStaticText{qclsinst:qthis}
}

// textWidth()
func (this *QStaticText) textWidth(args ...interface{}) () {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText9textWidthEv
    // invoke: qreal textWidth()
    var ret = C.C_ZNK11QStaticText9textWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStaticText", "textWidth", args)
  }

}

// textFormat()
func (this *QStaticText) textFormat(args ...interface{}) () {
  // textFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText10textFormatEv
    // invoke: Qt::TextFormat textFormat()
    C.C_ZNK11QStaticText10textFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "textFormat", args)
  }

}

// textOption()
func (this *QStaticText) textOption(args ...interface{}) () {
  // textOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText10textOptionEv
    // invoke: QTextOption textOption()
    var ret = C.C_ZNK11QStaticText10textOptionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStaticText", "textOption", args)
  }

}

// performanceHint()
func (this *QStaticText) performanceHint(args ...interface{}) () {
  // performanceHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText15performanceHintEv
    // invoke: QStaticText::PerformanceHint performanceHint()
    C.C_ZNK11QStaticText15performanceHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "performanceHint", args)
  }

}

// size()
func (this *QStaticText) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStaticText4sizeEv
    // invoke: QSizeF size()
    var ret = C.C_ZNK11QStaticText4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QStaticText", "size", args)
  }

}

// <= body block end

