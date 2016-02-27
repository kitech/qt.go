package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
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
extern void* C_ZNK11QStaticText4textEv(void* qthis); // 4
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
extern double C_ZNK11QStaticText9textWidthEv(void* qthis); // 4
  // proto:  Qt::TextFormat QStaticText::textFormat();
extern void C_ZNK11QStaticText10textFormatEv(void* qthis); // 4
  // proto:  QTextOption QStaticText::textOption();
extern void* C_ZNK11QStaticText10textOptionEv(void* qthis); // 4
  // proto:  QStaticText::PerformanceHint QStaticText::performanceHint();
extern void C_ZNK11QStaticText15performanceHintEv(void* qthis); // 4
  // proto:  QSizeF QStaticText::size();
extern void* C_ZNK11QStaticText4sizeEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QStaticText)=1
type QStaticText struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// swap(class QStaticText &)
func (this *QStaticText) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QStaticText).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "swap", args)
  }

  return
}

// prepare(const class QTransform &, const class QFont &)
func (this *QStaticText) Prepare(args ...interface{}) () {
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
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QFont).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QStaticText7prepareERK10QTransformRK5QFont(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStaticText", "prepare", args)
  }

  return
}

// text()
func (this *QStaticText) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStaticText4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStaticText", "text", args)
  }

  return
}

// setText(const class QString &)
func (this *QStaticText) SetText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStaticText7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setText", args)
  }

  return
}

// ~QStaticText()
func (this *QStaticText) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QStaticTextD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QStaticText", "~QStaticText", args)
  }

  return
}

// setTextWidth(qreal)
func (this *QStaticText) SetTextWidth(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText12setTextWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextWidth", args)
  }

  return
}

// setTextOption(const class QTextOption &)
func (this *QStaticText) SetTextOption(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextOption).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QStaticText13setTextOptionERK11QTextOption(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStaticText", "setTextOption", args)
  }

  return
}

// QStaticText()
func GcfreeQStaticText(this *QStaticText) {
  qtrt.UniverseFree(this)
}
func NewQStaticText(args ...interface{}) *QStaticText {
  // QStaticText()
  // QStaticText(const class QString &)
  // QStaticText(const class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
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
    this := &QStaticText{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStaticText)
    return this
  case 1:
    // invoke: _ZN11QStaticTextC1ERK7QString
    // invoke: void QStaticText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStaticTextC2ERK7QString(arg0)
    this := &QStaticText{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStaticText)
    return this
  case 2:
    // invoke: _ZN11QStaticTextC1ERKS_
    // invoke: void QStaticText(const class QStaticText &)
    var arg0 = args[0].(*QStaticText).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStaticTextC2ERKS_(arg0)
    this := &QStaticText{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStaticText)
    return this
  default:
    qtrt.ErrorResolve("QStaticText", "QStaticText", args)
  }

  return nil // QStaticText{Qclsinst:qthis}
}

// textWidth()
func (this *QStaticText) TextWidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStaticText9textWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStaticText", "textWidth", args)
  }

  return
}

// textFormat()
func (this *QStaticText) TextFormat(args ...interface{}) () {
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
    C.C_ZNK11QStaticText10textFormatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "textFormat", args)
  }

  return
}

// textOption()
func (this *QStaticText) TextOption(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStaticText10textOptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextOption{}) // "QTextOption"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStaticText", "textOption", args)
  }

  return
}

// performanceHint()
func (this *QStaticText) PerformanceHint(args ...interface{}) () {
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
    C.C_ZNK11QStaticText15performanceHintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStaticText", "performanceHint", args)
  }

  return
}

// size()
func (this *QStaticText) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QStaticText4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStaticText", "size", args)
  }

  return
}

// <= body block end

