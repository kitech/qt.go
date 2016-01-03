package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QStaticText::QStaticText(const QString & text);
extern void* dector_ZN11QStaticTextC1ERK7QString(void* arg0);
extern void _ZN11QStaticTextC1ERK7QString(void* qthis, void* arg0);
  // proto:  QSizeF QStaticText::size();
extern void _ZNK11QStaticText4sizeEv(void* qthis);
  // proto:  QString QStaticText::text();
extern void _ZNK11QStaticText4textEv(void* qthis);
  // proto:  void QStaticText::~QStaticText();
extern void _ZN11QStaticTextD0Ev(void* qthis);
  // proto:  void QStaticText::setText(const QString & text);
extern void _ZN11QStaticText7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QStaticText::QStaticText();
extern void* dector_ZN11QStaticTextC1Ev();
extern void _ZN11QStaticTextC1Ev(void* qthis);
  // proto:  void QStaticText::prepare(const QTransform & matrix, const QFont & font);
extern void _ZN11QStaticText7prepareERK10QTransformRK5QFont(void* qthis, void* arg0, void* arg1);
  // proto:  void QStaticText::setTextOption(const QTextOption & textOption);
extern void _ZN11QStaticText13setTextOptionERK11QTextOption(void* qthis, void* arg0);
  // proto:  void QStaticText::setTextWidth(qreal textWidth);
extern void _ZN11QStaticText12setTextWidthEd(void* qthis, double arg0);
  // proto:  qreal QStaticText::textWidth();
extern void _ZNK11QStaticText9textWidthEv(void* qthis);
  // proto:  void QStaticText::swap(QStaticText & other);
extern void demth_ZN11QStaticText4swapERS_(void* qthis, void* arg0);
  // proto:  QTextOption QStaticText::textOption();
extern void _ZNK11QStaticText10textOptionEv(void* qthis);
  // proto:  void QStaticText::QStaticText(const QStaticText & other);
extern void* dector_ZN11QStaticTextC1ERKS_(void* arg0);
extern void _ZN11QStaticTextC1ERKS_(void* qthis, void* arg0);
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

  // proto:  void QStaticText::QStaticText(const QString & text);
func NewQStaticText(args ...interface{}) QStaticText {
  return QStaticText{}
}

  // proto:  QSizeF QStaticText::size();
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
    C._ZNK11QStaticText4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "size", args)
  }

}

  // proto:  QString QStaticText::text();
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
    C._ZNK11QStaticText4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "text", args)
  }

}

  // proto:  void QStaticText::~QStaticText();
func (this *QStaticText) FreeQStaticText(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStaticText", "~QStaticText", args)
  }

}

  // proto:  void QStaticText::setText(const QString & text);
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
    C._ZN11QStaticText7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setText", args)
  }

}

  // proto:  void QStaticText::prepare(const QTransform & matrix, const QFont & font);
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
    C._ZN11QStaticText7prepareERK10QTransformRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStaticText", "prepare", args)
  }

}

  // proto:  void QStaticText::setTextOption(const QTextOption & textOption);
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
    C._ZN11QStaticText13setTextOptionERK11QTextOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextOption", args)
  }

}

  // proto:  void QStaticText::setTextWidth(qreal textWidth);
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
    C._ZN11QStaticText12setTextWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextWidth", args)
  }

}

  // proto:  qreal QStaticText::textWidth();
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
    C._ZNK11QStaticText9textWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "textWidth", args)
  }

}

  // proto:  void QStaticText::swap(QStaticText & other);
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
    C.demth_ZN11QStaticText4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "swap", args)
  }

}

  // proto:  QTextOption QStaticText::textOption();
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
    C._ZNK11QStaticText10textOptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "textOption", args)
  }

}

// <= body block end

