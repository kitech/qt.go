package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QSyntaxHighlighter::~QSyntaxHighlighter();
extern void _ZN18QSyntaxHighlighterD0Ev(void* qthis);
  // proto:  void QSyntaxHighlighter::rehighlight();
extern void _ZN18QSyntaxHighlighter11rehighlightEv(void* qthis);
  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(QObject * parent);
extern void* dector_ZN18QSyntaxHighlighterC1EP7QObject(void* arg0);
extern void _ZN18QSyntaxHighlighterC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QSyntaxHighlighter::rehighlightBlock(const QTextBlock & block);
extern void _ZN18QSyntaxHighlighter16rehighlightBlockERK10QTextBlock(void* qthis, void* arg0);
  // proto:  void QSyntaxHighlighter::setDocument(QTextDocument * doc);
extern void _ZN18QSyntaxHighlighter11setDocumentEP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(QTextDocument * parent);
extern void* dector_ZN18QSyntaxHighlighterC1EP13QTextDocument(void* arg0);
extern void _ZN18QSyntaxHighlighterC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(const QSyntaxHighlighter & );
extern void* dector_ZN18QSyntaxHighlighterC1ERKS_(void* arg0);
extern void _ZN18QSyntaxHighlighterC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSyntaxHighlighter::metaObject();
extern void _ZNK18QSyntaxHighlighter10metaObjectEv(void* qthis);
  // proto:  QTextDocument * QSyntaxHighlighter::document();
extern void _ZNK18QSyntaxHighlighter8documentEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSyntaxHighlighter::~QSyntaxHighlighter();
func (this *QSyntaxHighlighter) FreeQSyntaxHighlighter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "~QSyntaxHighlighter", args)
  }

}

  // proto:  void QSyntaxHighlighter::rehighlight();
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
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlight", args)
  }

}

  // proto:  void QSyntaxHighlighter::QSyntaxHighlighter(QObject * parent);
func NewQSyntaxHighlighter(args ...interface{}) QSyntaxHighlighter {
  return QSyntaxHighlighter{}
}

  // proto:  void QSyntaxHighlighter::rehighlightBlock(const QTextBlock & block);
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
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlightBlock", args)
  }

}

  // proto:  void QSyntaxHighlighter::setDocument(QTextDocument * doc);
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
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "setDocument", args)
  }

}

  // proto:  const QMetaObject * QSyntaxHighlighter::metaObject();
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
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "metaObject", args)
  }

}

  // proto:  QTextDocument * QSyntaxHighlighter::document();
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
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "document", args)
  }

}

// <= body block end

