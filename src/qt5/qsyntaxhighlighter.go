package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK18QSyntaxHighlighter8documentEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setDocument(class QTextDocument *)
func (this *QSyntaxHighlighter) Setdocument(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "setDocument", args)
  }

  return
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QSyntaxHighlighterC2EP7QObject(arg0)
    return &QSyntaxHighlighter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN18QSyntaxHighlighterC1EP13QTextDocument
    // invoke: void QSyntaxHighlighter(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QSyntaxHighlighterC2EP13QTextDocument(arg0)
    return &QSyntaxHighlighter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "QSyntaxHighlighter", args)
  }

  return nil // QSyntaxHighlighter{Qclsinst:qthis}
}

// document()
func (this *QSyntaxHighlighter) Document(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QSyntaxHighlighter8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "document", args)
  }

  return
}

// ~QSyntaxHighlighter()
func (this *QSyntaxHighlighter) Freeqsyntaxhighlighter(args ...interface{}) () {
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
    C.C_ZN18QSyntaxHighlighterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "~QSyntaxHighlighter", args)
  }

  return
}

// rehighlight()
func (this *QSyntaxHighlighter) Rehighlight(args ...interface{}) () {
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
    C.C_ZN18QSyntaxHighlighter11rehighlightEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlight", args)
  }

  return
}

// metaObject()
func (this *QSyntaxHighlighter) Metaobject(args ...interface{}) () {
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
    C.C_ZNK18QSyntaxHighlighter10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "metaObject", args)
  }

  return
}

// rehighlightBlock(const class QTextBlock &)
func (this *QSyntaxHighlighter) Rehighlightblock(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlightBlock", args)
  }

  return
}

// <= body block end

