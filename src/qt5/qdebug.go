package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QNoDebug & QNoDebug::maybeSpace();
extern void C_ZN8QNoDebug10maybeSpaceEv(void* qthis); // 2
  // proto:  QNoDebug & QNoDebug::space();
extern void C_ZN8QNoDebug5spaceEv(void* qthis); // 2
  // proto:  QNoDebug & QNoDebug::quote();
extern void C_ZN8QNoDebug5quoteEv(void* qthis); // 2
  // proto:  QNoDebug & QNoDebug::nospace();
extern void C_ZN8QNoDebug7nospaceEv(void* qthis); // 2
  // proto:  QNoDebug & QNoDebug::noquote();
extern void C_ZN8QNoDebug7noquoteEv(void* qthis); // 2
  // proto:  QNoDebug & QNoDebug::maybeQuote(const char );
extern void C_ZN8QNoDebug10maybeQuoteEc(void* qthis, unsigned char arg0); // 2
  // proto:  void QDebugStateSaver::QDebugStateSaver(QDebug & dbg);
extern void* C_ZN16QDebugStateSaverC2ER6QDebug(void* arg0); // 3
  // proto:  void QDebugStateSaver::~QDebugStateSaver();
extern void C_ZN16QDebugStateSaverD2Ev(void* qthis); // 4
  // proto:  void QDebug::~QDebug();
extern void C_ZN6QDebugD2Ev(void* qthis); // 4
  // proto:  QDebug & QDebug::space();
extern void C_ZN6QDebug5spaceEv(void* qthis); // 2
  // proto:  void QDebug::setAutoInsertSpaces(bool b);
extern void C_ZN6QDebug19setAutoInsertSpacesEb(void* qthis, bool arg0); // 2
  // proto:  bool QDebug::autoInsertSpaces();
extern void C_ZNK6QDebug16autoInsertSpacesEv(void* qthis); // 2
  // proto:  QDebug & QDebug::nospace();
extern void C_ZN6QDebug7nospaceEv(void* qthis); // 2
  // proto:  QDebug & QDebug::resetFormat();
extern void C_ZN6QDebug11resetFormatEv(void* qthis); // 4
  // proto:  QDebug & QDebug::noquote();
extern void C_ZN6QDebug7noquoteEv(void* qthis); // 2
  // proto:  void QDebug::swap(QDebug & other);
extern void C_ZN6QDebug4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QDebug & QDebug::quote();
extern void C_ZN6QDebug5quoteEv(void* qthis); // 2
  // proto:  void QDebug::QDebug(const QDebug & o);
extern void* C_ZN6QDebugC2ERKS_(void* arg0); // 1
  // proto:  void QDebug::QDebug(QtMsgType t);
extern void* C_ZN6QDebugC2E9QtMsgType(int32_t arg0); // 1
  // proto:  void QDebug::QDebug(QString * string);
extern void* C_ZN6QDebugC2EP7QString(void* arg0); // 1
  // proto:  void QDebug::QDebug(QIODevice * device);
extern void* C_ZN6QDebugC2EP9QIODevice(void* arg0); // 1
  // proto:  QDebug & QDebug::maybeSpace();
extern void C_ZN6QDebug10maybeSpaceEv(void* qthis); // 2
  // proto:  QDebug & QDebug::maybeQuote(char c);
extern void C_ZN6QDebug10maybeQuoteEc(void* qthis, unsigned char arg0); // 2
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

// maybeSpace()
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
    var ret = C.C_ZN8QNoDebug10maybeSpaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeSpace", args)
  }

}

// space()
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
    var ret = C.C_ZN8QNoDebug5spaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "space", args)
  }

}

// quote()
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
    var ret = C.C_ZN8QNoDebug5quoteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "quote", args)
  }

}

// nospace()
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
    var ret = C.C_ZN8QNoDebug7nospaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "nospace", args)
  }

}

// noquote()
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
    var ret = C.C_ZN8QNoDebug7noquoteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "noquote", args)
  }

}

// maybeQuote(const char)
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
    var ret = C.C_ZN8QNoDebug10maybeQuoteEc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeQuote", args)
  }

}

// QDebugStateSaver(class QDebug &)
func NewQDebugStateSaver(args ...interface{}) *QDebugStateSaver {
  // QDebugStateSaver(class QDebug &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDebug{}) // "QDebug &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDebugStateSaverC1ER6QDebug
    // invoke: void QDebugStateSaver(class QDebug &)
    var arg0 = args[0].(QDebug).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QDebugStateSaverC2ER6QDebug(arg0)
    return &QDebugStateSaver{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDebugStateSaver", "QDebugStateSaver", args)
  }

  return nil // QDebugStateSaver{qclsinst:qthis}
}

// ~QDebugStateSaver()
func (this *QDebugStateSaver) FreeQDebugStateSaver(args ...interface{}) () {
  // ~QDebugStateSaver()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDebugStateSaverD0Ev
    // invoke: void ~QDebugStateSaver()
    C.C_ZN16QDebugStateSaverD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebugStateSaver", "~QDebugStateSaver", args)
  }

}

// ~QDebug()
func (this *QDebug) FreeQDebug(args ...interface{}) () {
  // ~QDebug()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebugD0Ev
    // invoke: void ~QDebug()
    C.C_ZN6QDebugD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDebug", "~QDebug", args)
  }

}

// space()
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
    var ret = C.C_ZN6QDebug5spaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "space", args)
  }

}

// setAutoInsertSpaces(_Bool)
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
    C.C_ZN6QDebug19setAutoInsertSpacesEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDebug", "setAutoInsertSpaces", args)
  }

}

// autoInsertSpaces()
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
    var ret = C.C_ZNK6QDebug16autoInsertSpacesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "autoInsertSpaces", args)
  }

}

// nospace()
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
    var ret = C.C_ZN6QDebug7nospaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "nospace", args)
  }

}

// resetFormat()
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
    var ret = C.C_ZN6QDebug11resetFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "resetFormat", args)
  }

}

// noquote()
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
    var ret = C.C_ZN6QDebug7noquoteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "noquote", args)
  }

}

// swap(class QDebug &)
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
    C.C_ZN6QDebug4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDebug", "swap", args)
  }

}

// quote()
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
    var ret = C.C_ZN6QDebug5quoteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "quote", args)
  }

}

// QDebug(const class QDebug &)
func NewQDebug(args ...interface{}) *QDebug {
  // QDebug(const class QDebug &)
  // QDebug(enum QtMsgType)
  // QDebug(class QString *)
  // QDebug(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDebug{}) // "const QDebug &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QtMsgType"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebugC1ERKS_
    // invoke: void QDebug(const class QDebug &)
    var arg0 = args[0].(QDebug).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QDebugC2ERKS_(arg0)
    return &QDebug{qclsinst:qthis}
  case 1:
    // invoke: _ZN6QDebugC1E9QtMsgType
    // invoke: void QDebug(enum QtMsgType)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QDebugC2E9QtMsgType(arg0)
    return &QDebug{qclsinst:qthis}
  case 2:
    // invoke: _ZN6QDebugC1EP7QString
    // invoke: void QDebug(class QString *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QDebugC2EP7QString(arg0)
    return &QDebug{qclsinst:qthis}
  case 3:
    // invoke: _ZN6QDebugC1EP9QIODevice
    // invoke: void QDebug(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QDebugC2EP9QIODevice(arg0)
    return &QDebug{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDebug", "QDebug", args)
  }

  return nil // QDebug{qclsinst:qthis}
}

// maybeSpace()
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
    var ret = C.C_ZN6QDebug10maybeSpaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "maybeSpace", args)
  }

}

// maybeQuote(char)
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
    var ret = C.C_ZN6QDebug10maybeQuoteEc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDebug", "maybeQuote", args)
  }

}

// <= body block end

