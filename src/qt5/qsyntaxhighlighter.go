package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qsyntaxhighlighter.h
// dst-file: /src/gui/qsyntaxhighlighter.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QSyntaxHighlighter)=1
type QSyntaxHighlighter struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


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


func NewQSyntaxHighlighter(args ...interface{}) QSyntaxHighlighter {
  return QSyntaxHighlighter{}
}


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
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "rehighlightBlock", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSyntaxHighlighter", "setDocument", args)
  }

}


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

