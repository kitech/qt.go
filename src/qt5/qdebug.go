package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qdebug.h
// dst-file: /src/core/qdebug.go
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
  // proto:  QNoDebug & QNoDebug::maybeQuote(const char );
extern void demth_ZN8QNoDebug10maybeQuoteEc(void* qthis, unsigned char arg0);
  // proto:  QNoDebug & QNoDebug::quote();
extern void demth_ZN8QNoDebug5quoteEv(void* qthis);
  // proto:  QNoDebug & QNoDebug::space();
extern void demth_ZN8QNoDebug5spaceEv(void* qthis);
  // proto:  QNoDebug & QNoDebug::nospace();
extern void demth_ZN8QNoDebug7nospaceEv(void* qthis);
  // proto:  QNoDebug & QNoDebug::noquote();
extern void demth_ZN8QNoDebug7noquoteEv(void* qthis);
  // proto:  QNoDebug & QNoDebug::maybeSpace();
extern void demth_ZN8QNoDebug10maybeSpaceEv(void* qthis);
  // proto:  void QDebugStateSaver::QDebugStateSaver(QDebug & dbg);
extern void* dector_ZN16QDebugStateSaverC1ER6QDebug(void* arg0);
extern void _ZN16QDebugStateSaverC1ER6QDebug(void* qthis, void* arg0);
  // proto:  void QDebugStateSaver::QDebugStateSaver(const QDebugStateSaver & );
extern void* dector_ZN16QDebugStateSaverC1ERKS_(void* arg0);
extern void _ZN16QDebugStateSaverC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDebugStateSaver::~QDebugStateSaver();
extern void _ZN16QDebugStateSaverD0Ev(void* qthis);
  // proto:  QDebug & QDebug::noquote();
extern void demth_ZN6QDebug7noquoteEv(void* qthis);
  // proto:  void QDebug::~QDebug();
extern void _ZN6QDebugD0Ev(void* qthis);
  // proto:  void QDebug::QDebug(const QDebug & o);
extern void* dector_ZN6QDebugC1ERKS_(void* arg0);
extern void demth_ZN6QDebugC1ERKS_(void* qthis, void* arg0);
  // proto:  QDebug & QDebug::space();
extern void demth_ZN6QDebug5spaceEv(void* qthis);
  // proto:  void QDebug::QDebug(QtMsgType t);
extern void* dector_ZN6QDebugC1E9QtMsgType(int32_t arg0);
extern void demth_ZN6QDebugC1E9QtMsgType(void* qthis, int32_t arg0);
  // proto:  QDebug & QDebug::maybeSpace();
extern void demth_ZN6QDebug10maybeSpaceEv(void* qthis);
  // proto:  QDebug & QDebug::resetFormat();
extern void _ZN6QDebug11resetFormatEv(void* qthis);
  // proto:  void QDebug::setAutoInsertSpaces(bool b);
extern void demth_ZN6QDebug19setAutoInsertSpacesEb(void* qthis, bool arg0);
  // proto:  void QDebug::QDebug(QString * string);
extern void* dector_ZN6QDebugC1EP7QString(void* arg0);
extern void demth_ZN6QDebugC1EP7QString(void* qthis, void* arg0);
  // proto:  void QDebug::swap(QDebug & other);
extern void demth_ZN6QDebug4swapERS_(void* qthis, void* arg0);
  // proto:  QDebug & QDebug::nospace();
extern void demth_ZN6QDebug7nospaceEv(void* qthis);
  // proto:  bool QDebug::autoInsertSpaces();
extern void demth_ZNK6QDebug16autoInsertSpacesEv(void* qthis);
  // proto:  void QDebug::QDebug(QIODevice * device);
extern void* dector_ZN6QDebugC1EP9QIODevice(void* arg0);
extern void demth_ZN6QDebugC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  QDebug & QDebug::quote();
extern void demth_ZN6QDebug5quoteEv(void* qthis);
  // proto:  QDebug & QDebug::maybeQuote(char c);
extern void demth_ZN6QDebug10maybeQuoteEc(void* qthis, unsigned char arg0);
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

// class sizeof(QNoDebug)=1
type QNoDebug struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDebugStateSaver)=1
type QDebugStateSaver struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDebug)=8
type QDebug struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QNoDebug & QNoDebug::maybeQuote(const char );
func (this *QNoDebug) maybeQuote(args ...interface{}) () {
  // maybeQuote(const char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "const char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug10maybeQuoteEc
    // invoke: QNoDebug & maybeQuote(const char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN8QNoDebug10maybeQuoteEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeQuote", args)
  }

}

  // proto:  QNoDebug & QNoDebug::quote();
func (this *QNoDebug) quote(args ...interface{}) () {
  // quote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug5quoteEv
    // invoke: QNoDebug & quote()
    C.demth_ZN8QNoDebug5quoteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNoDebug", "quote", args)
  }

}

  // proto:  QNoDebug & QNoDebug::space();
func (this *QNoDebug) space(args ...interface{}) () {
  // space()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug5spaceEv
    // invoke: QNoDebug & space()
    C.demth_ZN8QNoDebug5spaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNoDebug", "space", args)
  }

}

  // proto:  QNoDebug & QNoDebug::nospace();
func (this *QNoDebug) nospace(args ...interface{}) () {
  // nospace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug7nospaceEv
    // invoke: QNoDebug & nospace()
    C.demth_ZN8QNoDebug7nospaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNoDebug", "nospace", args)
  }

}

  // proto:  QNoDebug & QNoDebug::noquote();
func (this *QNoDebug) noquote(args ...interface{}) () {
  // noquote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug7noquoteEv
    // invoke: QNoDebug & noquote()
    C.demth_ZN8QNoDebug7noquoteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNoDebug", "noquote", args)
  }

}

  // proto:  QNoDebug & QNoDebug::maybeSpace();
func (this *QNoDebug) maybeSpace(args ...interface{}) () {
  // maybeSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug10maybeSpaceEv
    // invoke: QNoDebug & maybeSpace()
    C.demth_ZN8QNoDebug10maybeSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeSpace", args)
  }

}

  // proto:  void QDebugStateSaver::QDebugStateSaver(QDebug & dbg);
func NewQDebugStateSaver(args ...interface{}) QDebugStateSaver {
  return QDebugStateSaver{}
}

  // proto:  void QDebugStateSaver::~QDebugStateSaver();
func (this *QDebugStateSaver) FreeQDebugStateSaver(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDebugStateSaver", "~QDebugStateSaver", args)
  }

}

  // proto:  QDebug & QDebug::noquote();
func (this *QDebug) noquote(args ...interface{}) () {
  // noquote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug7noquoteEv
    // invoke: QDebug & noquote()
    C.demth_ZN6QDebug7noquoteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "noquote", args)
  }

}

  // proto:  void QDebug::~QDebug();
func (this *QDebug) FreeQDebug(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDebug", "~QDebug", args)
  }

}

  // proto:  void QDebug::QDebug(const QDebug & o);
func NewQDebug(args ...interface{}) QDebug {
  return QDebug{}
}

  // proto:  QDebug & QDebug::space();
func (this *QDebug) space(args ...interface{}) () {
  // space()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug5spaceEv
    // invoke: QDebug & space()
    C.demth_ZN6QDebug5spaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "space", args)
  }

}

  // proto:  QDebug & QDebug::maybeSpace();
func (this *QDebug) maybeSpace(args ...interface{}) () {
  // maybeSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug10maybeSpaceEv
    // invoke: QDebug & maybeSpace()
    C.demth_ZN6QDebug10maybeSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "maybeSpace", args)
  }

}

  // proto:  QDebug & QDebug::resetFormat();
func (this *QDebug) resetFormat(args ...interface{}) () {
  // resetFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug11resetFormatEv
    // invoke: QDebug & resetFormat()
    C._ZN6QDebug11resetFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "resetFormat", args)
  }

}

  // proto:  void QDebug::setAutoInsertSpaces(bool b);
func (this *QDebug) setAutoInsertSpaces(args ...interface{}) () {
  // setAutoInsertSpaces(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug19setAutoInsertSpacesEb
    // invoke: void setAutoInsertSpaces(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.demth_ZN6QDebug19setAutoInsertSpacesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDebug", "setAutoInsertSpaces", args)
  }

}

  // proto:  void QDebug::swap(QDebug & other);
func (this *QDebug) swap(args ...interface{}) () {
  // swap(class QDebug &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDebug{}) // "QDebug &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug4swapERS_
    // invoke: void swap(class QDebug &)
    var arg0 = args[0].(QDebug).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN6QDebug4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDebug", "swap", args)
  }

}

  // proto:  QDebug & QDebug::nospace();
func (this *QDebug) nospace(args ...interface{}) () {
  // nospace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug7nospaceEv
    // invoke: QDebug & nospace()
    C.demth_ZN6QDebug7nospaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "nospace", args)
  }

}

  // proto:  bool QDebug::autoInsertSpaces();
func (this *QDebug) autoInsertSpaces(args ...interface{}) () {
  // autoInsertSpaces()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QDebug16autoInsertSpacesEv
    // invoke: bool autoInsertSpaces()
    C.demth_ZNK6QDebug16autoInsertSpacesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "autoInsertSpaces", args)
  }

}

  // proto:  QDebug & QDebug::quote();
func (this *QDebug) quote(args ...interface{}) () {
  // quote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug5quoteEv
    // invoke: QDebug & quote()
    C.demth_ZN6QDebug5quoteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "quote", args)
  }

}

  // proto:  QDebug & QDebug::maybeQuote(char c);
func (this *QDebug) maybeQuote(args ...interface{}) () {
  // maybeQuote(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug10maybeQuoteEc
    // invoke: QDebug & maybeQuote(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN6QDebug10maybeQuoteEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDebug", "maybeQuote", args)
  }

}

// <= body block end

