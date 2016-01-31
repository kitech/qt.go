package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qsyntaxhighlighter.h
// dst-file: /src/gui/qsyntaxhighlighter.go
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
  // proto:  void QSyntaxHighlighter::setDocument(QTextDocument * doc);
extern void C_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(QObject * parent);
extern void* C_ZN18QSyntaxHighlighterC2EP7QObject(void* arg0); // 3
  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(QTextDocument * parent);
extern void* C_ZN18QSyntaxHighlighterC2EP13QTextDocument(void* arg0); // 3
  // proto:  QTextDocument * QSyntaxHighlighter::document();
extern void C_ZNK18QSyntaxHighlighter8documentEv(void* qthis); // 4
  // proto:  void QSyntaxHighlighter::~QSyntaxHighlighter();
extern void C_ZN18QSyntaxHighlighterD2Ev(void* qthis); // 4
  // proto:  void QSyntaxHighlighter::rehighlight();
extern void C_ZN18QSyntaxHighlighter11rehighlightEv(void* qthis); // 4
  // proto:  const QMetaObject * QSyntaxHighlighter::metaObject();
extern void C_ZNK18QSyntaxHighlighter10metaObjectEv(void* qthis); // 4
  // proto:  void QSyntaxHighlighter::rehighlightBlock(const QTextBlock & block);
extern void C_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock(void* qthis, void* arg0); // 4
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

// class sizeof(QSyntaxHighlighter)=1
type QSyntaxHighlighter struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setDocument(class QTextDocument *)
func (this *QSyntaxHighlighter) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "setDocument", args)
  }

}

// QSyntaxHighlighter(class QObject *)
func NewQSyntaxHighlighter(args ...interface{}) *QSyntaxHighlighter {
  // QSyntaxHighlighter(class QObject *)
  // QSyntaxHighlighter(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QSyntaxHighlighterC1EP7QObject
    // invoke: void QSyntaxHighlighter(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QSyntaxHighlighterC2EP7QObject(arg0)
    return &QSyntaxHighlighter{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QSyntaxHighlighterC1EP13QTextDocument
    // invoke: void QSyntaxHighlighter(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QSyntaxHighlighterC2EP13QTextDocument(arg0)
    return &QSyntaxHighlighter{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "QSyntaxHighlighter", args)
  }

  return nil // QSyntaxHighlighter{qclsinst:qthis}
}

// document()
func (this *QSyntaxHighlighter) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QSyntaxHighlighter8documentEv
    // invoke: QTextDocument * document()
    var ret = C.C_ZNK18QSyntaxHighlighter8documentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "document", args)
  }

}

// ~QSyntaxHighlighter()
func (this *QSyntaxHighlighter) FreeQSyntaxHighlighter(args ...interface{}) () {
  // ~QSyntaxHighlighter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QSyntaxHighlighterD0Ev
    // invoke: void ~QSyntaxHighlighter()
    C.C_ZN18QSyntaxHighlighterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "~QSyntaxHighlighter", args)
  }

}

// rehighlight()
func (this *QSyntaxHighlighter) rehighlight(args ...interface{}) () {
  // rehighlight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QSyntaxHighlighter11rehighlightEv
    // invoke: void rehighlight()
    C.C_ZN18QSyntaxHighlighter11rehighlightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlight", args)
  }

}

// metaObject()
func (this *QSyntaxHighlighter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QSyntaxHighlighter10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QSyntaxHighlighter10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "metaObject", args)
  }

}

// rehighlightBlock(const class QTextBlock &)
func (this *QSyntaxHighlighter) rehighlightBlock(args ...interface{}) () {
  // rehighlightBlock(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock
    // invoke: void rehighlightBlock(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlightBlock", args)
  }

}

// <= body block end

