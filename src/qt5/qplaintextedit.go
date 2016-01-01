package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qplaintextedit.h
// dst-file: /src/widgets/qplaintextedit.go
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
// class sizeof(QPlainTextDocumentLayout)=1
type QPlainTextDocumentLayout struct {
  /*qbase*/ QAbstractTextDocumentLayout;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPlainTextEdit)=1
type QPlainTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorPositionChanged QPlainTextEdit_cursorPositionChanged_signal;
//  _modificationChanged QPlainTextEdit_modificationChanged_signal;
//  _redoAvailable QPlainTextEdit_redoAvailable_signal;
//  _selectionChanged QPlainTextEdit_selectionChanged_signal;
//  _updateRequest QPlainTextEdit_updateRequest_signal;
//  _blockCountChanged QPlainTextEdit_blockCountChanged_signal;
//  _undoAvailable QPlainTextEdit_undoAvailable_signal;
//  _textChanged QPlainTextEdit_textChanged_signal;
//  _copyAvailable QPlainTextEdit_copyAvailable_signal;
}


func (this *QPlainTextDocumentLayout) requestUpdate(args ...interface{}) () {
  // requestUpdate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayout13requestUpdateEv
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "requestUpdate", args)
 }

}


func (this *QPlainTextDocumentLayout) setCursorWidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayout14setCursorWidthEi
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "setCursorWidth", args)
 }

}


func (this *QPlainTextDocumentLayout) frameBoundingRect(args ...interface{}) () {
  // frameBoundingRect(class QTextFrame *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "frameBoundingRect", args)
 }

}


func (this *QPlainTextDocumentLayout) pageCount(args ...interface{}) () {
  // pageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout9pageCountEv
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "pageCount", args)
 }

}


func (this *QPlainTextDocumentLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout10metaObjectEv
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "metaObject", args)
 }

}


func (this *QPlainTextDocumentLayout) ensureBlockLayout(args ...interface{}) () {
  // ensureBlockLayout(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "ensureBlockLayout", args)
 }

}


func (this *QPlainTextDocumentLayout) FreeQPlainTextDocumentLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "~QPlainTextDocumentLayout", args)
 }

}


func (this *QPlainTextDocumentLayout) blockBoundingRect(args ...interface{}) () {
  // blockBoundingRect(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "blockBoundingRect", args)
 }

}


func (this *QPlainTextDocumentLayout) cursorWidth(args ...interface{}) () {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout11cursorWidthEv
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "cursorWidth", args)
 }

}


func NewQPlainTextDocumentLayout(args ...interface{})() {
}


func (this *QPlainTextDocumentLayout) documentSize(args ...interface{}) () {
  // documentSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout12documentSizeEv
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "documentSize", args)
 }

}


func (this *QPlainTextEdit) createStandardContextMenu(args ...interface{}) () {
  // createStandardContextMenu(const class QPoint &)
  // createStandardContextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint
  case 1:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "createStandardContextMenu", args)
 }

}


func (this *QPlainTextEdit) ensureCursorVisible(args ...interface{}) () {
  // ensureCursorVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit19ensureCursorVisibleEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "ensureCursorVisible", args)
 }

}


func (this *QPlainTextEdit) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8documentEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "document", args)
 }

}


func (this *QPlainTextEdit) cursorRect(args ...interface{}) () {
  // cursorRect()
  // cursorRect(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10cursorRectEv
  case 1:
    // invoke: _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorRect", args)
 }

}


func (this *QPlainTextEdit) setTabChangesFocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setTabChangesFocusEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabChangesFocus", args)
 }

}


func (this *QPlainTextEdit) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit11toPlainTextEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "toPlainText", args)
 }

}


func (this *QPlainTextEdit) loadResource(args ...interface{}) () {
  // loadResource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12loadResourceEiRK4QUrl
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "loadResource", args)
 }

}


func (this *QPlainTextEdit) tabStopWidth(args ...interface{}) () {
  // tabStopWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit12tabStopWidthEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabStopWidth", args)
 }

}


func (this *QPlainTextEdit) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isReadOnly", args)
 }

}


func (this *QPlainTextEdit) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit11setReadOnlyEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setReadOnly", args)
 }

}


func (this *QPlainTextEdit) textCursor(args ...interface{}) () {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10textCursorEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textCursor", args)
 }

}


func (this *QPlainTextEdit) setCenterOnScroll(args ...interface{}) () {
  // setCenterOnScroll(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit17setCenterOnScrollEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCenterOnScroll", args)
 }

}


func (this *QPlainTextEdit) placeholderText(args ...interface{}) () {
  // placeholderText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit15placeholderTextEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "placeholderText", args)
 }

}


func (this *QPlainTextEdit) blockCount(args ...interface{}) () {
  // blockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10blockCountEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "blockCount", args)
 }

}


func (this *QPlainTextEdit) setCurrentCharFormat(args ...interface{}) () {
  // setCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCurrentCharFormat", args)
 }

}


func (this *QPlainTextEdit) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit11setDocumentEP13QTextDocument
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocument", args)
 }

}


func (this *QPlainTextEdit) print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit5printEP17QPagedPaintDevice
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "print", args)
 }

}


func (this *QPlainTextEdit) setTabStopWidth(args ...interface{}) () {
  // setTabStopWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15setTabStopWidthEi
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabStopWidth", args)
 }

}


func (this *QPlainTextEdit) backgroundVisible(args ...interface{}) () {
  // backgroundVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17backgroundVisibleEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "backgroundVisible", args)
 }

}


func (this *QPlainTextEdit) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4redoEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "redo", args)
 }

}


func NewQPlainTextEdit(args ...interface{})() {
}


func (this *QPlainTextEdit) setOverwriteMode(args ...interface{}) () {
  // setOverwriteMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit16setOverwriteModeEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setOverwriteMode", args)
 }

}


func (this *QPlainTextEdit) tabChangesFocus(args ...interface{}) () {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit15tabChangesFocusEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabChangesFocus", args)
 }

}


func (this *QPlainTextEdit) copy(args ...interface{}) () {
  // copy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4copyEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "copy", args)
 }

}


func (this *QPlainTextEdit) mergeCurrentCharFormat(args ...interface{}) () {
  // mergeCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "mergeCurrentCharFormat", args)
 }

}


func (this *QPlainTextEdit) maximumBlockCount(args ...interface{}) () {
  // maximumBlockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17maximumBlockCountEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "maximumBlockCount", args)
 }

}


func (this *QPlainTextEdit) insertPlainText(args ...interface{}) () {
  // insertPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15insertPlainTextERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "insertPlainText", args)
 }

}


func (this *QPlainTextEdit) setTextCursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit13setTextCursorERK11QTextCursor
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTextCursor", args)
 }

}


func (this *QPlainTextEdit) paste(args ...interface{}) () {
  // paste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit5pasteEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "paste", args)
 }

}


func (this *QPlainTextEdit) zoomIn(args ...interface{}) () {
  // zoomIn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit6zoomInEi
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomIn", args)
 }

}


func (this *QPlainTextEdit) setMaximumBlockCount(args ...interface{}) () {
  // setMaximumBlockCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setMaximumBlockCountEi
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setMaximumBlockCount", args)
 }

}


func (this *QPlainTextEdit) currentCharFormat(args ...interface{}) () {
  // currentCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17currentCharFormatEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "currentCharFormat", args)
 }

}


func (this *QPlainTextEdit) setCursorWidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit14setCursorWidthEi
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCursorWidth", args)
 }

}


func (this *QPlainTextEdit) documentTitle(args ...interface{}) () {
  // documentTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit13documentTitleEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "documentTitle", args)
 }

}


func (this *QPlainTextEdit) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit9selectAllEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "selectAll", args)
 }

}


func (this *QPlainTextEdit) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12setPlainTextERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlainText", args)
 }

}


func (this *QPlainTextEdit) setBackgroundVisible(args ...interface{}) () {
  // setBackgroundVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setBackgroundVisibleEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setBackgroundVisible", args)
 }

}


func (this *QPlainTextEdit) setUndoRedoEnabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setUndoRedoEnabledEb
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setUndoRedoEnabled", args)
 }

}


func (this *QPlainTextEdit) overwriteMode(args ...interface{}) () {
  // overwriteMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit13overwriteModeEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "overwriteMode", args)
 }

}


func (this *QPlainTextEdit) centerCursor(args ...interface{}) () {
  // centerCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12centerCursorEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerCursor", args)
 }

}


func (this *QPlainTextEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10metaObjectEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "metaObject", args)
 }

}


func (this *QPlainTextEdit) setDocumentTitle(args ...interface{}) () {
  // setDocumentTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit16setDocumentTitleERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocumentTitle", args)
 }

}


func (this *QPlainTextEdit) FreeQPlainTextEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "~QPlainTextEdit", args)
 }

}


func (this *QPlainTextEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit5clearEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "clear", args)
 }

}


func (this *QPlainTextEdit) anchorAt(args ...interface{}) () {
  // anchorAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8anchorAtERK6QPoint
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "anchorAt", args)
 }

}


func (this *QPlainTextEdit) canPaste(args ...interface{}) () {
  // canPaste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8canPasteEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "canPaste", args)
 }

}


func (this *QPlainTextEdit) cut(args ...interface{}) () {
  // cut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit3cutEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cut", args)
 }

}


func (this *QPlainTextEdit) appendHtml(args ...interface{}) () {
  // appendHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit10appendHtmlERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendHtml", args)
 }

}


func (this *QPlainTextEdit) isUndoRedoEnabled(args ...interface{}) () {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17isUndoRedoEnabledEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isUndoRedoEnabled", args)
 }

}


func (this *QPlainTextEdit) zoomOut(args ...interface{}) () {
  // zoomOut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit7zoomOutEi
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomOut", args)
 }

}


func (this *QPlainTextEdit) setPlaceholderText(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setPlaceholderTextERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlaceholderText", args)
 }

}


func (this *QPlainTextEdit) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4undoEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "undo", args)
 }

}


func (this *QPlainTextEdit) cursorForPosition(args ...interface{}) () {
  // cursorForPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17cursorForPositionERK6QPoint
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorForPosition", args)
 }

}


func (this *QPlainTextEdit) centerOnScroll(args ...interface{}) () {
  // centerOnScroll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit14centerOnScrollEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerOnScroll", args)
 }

}


func (this *QPlainTextEdit) appendPlainText(args ...interface{}) () {
  // appendPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15appendPlainTextERK7QString
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendPlainText", args)
 }

}


func (this *QPlainTextEdit) cursorWidth(args ...interface{}) () {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit11cursorWidthEv
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorWidth", args)
 }

}

// <= body block end

