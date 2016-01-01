package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qtextdocument.h
// dst-file: /src/gui/qtextdocument.go
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
// class sizeof(QTextDocument)=1
type QTextDocument struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorPositionChanged QTextDocument_cursorPositionChanged_signal;
//  _documentLayoutChanged QTextDocument_documentLayoutChanged_signal;
//  _undoCommandAdded QTextDocument_undoCommandAdded_signal;
//  _modificationChanged QTextDocument_modificationChanged_signal;
//  _redoAvailable QTextDocument_redoAvailable_signal;
//  _contentsChanged QTextDocument_contentsChanged_signal;
//  _baseUrlChanged QTextDocument_baseUrlChanged_signal;
//  _blockCountChanged QTextDocument_blockCountChanged_signal;
//  _undoAvailable QTextDocument_undoAvailable_signal;
//  _contentsChange QTextDocument_contentsChange_signal;
}

// class sizeof(QAbstractUndoItem)=8
type QAbstractUndoItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextDocument) setDefaultStyleSheet(args ...interface{}) () {
  // setDefaultStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setDefaultStyleSheetERK7QString
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultStyleSheet", args)
  }

}


func (this *QTextDocument) documentLayout(args ...interface{}) () {
  // documentLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14documentLayoutEv
  default:
    qtrt.ErrorResolve("QTextDocument", "documentLayout", args)
  }

}


func (this *QTextDocument) isModified(args ...interface{}) () {
  // isModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10isModifiedEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isModified", args)
  }

}


func (this *QTextDocument) revision(args ...interface{}) () {
  // revision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8revisionEv
  default:
    qtrt.ErrorResolve("QTextDocument", "revision", args)
  }

}


func (this *QTextDocument) pageSize(args ...interface{}) () {
  // pageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8pageSizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "pageSize", args)
  }

}


func (this *QTextDocument) redo(args ...interface{}) () {
  // redo(class QTextCursor *)
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument4redoEP11QTextCursor
  case 1:
    // invoke: _ZN13QTextDocument4redoEv
  default:
    qtrt.ErrorResolve("QTextDocument", "redo", args)
  }

}


func (this *QTextDocument) lineCount(args ...interface{}) () {
  // lineCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9lineCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "lineCount", args)
  }

}


func (this *QTextDocument) print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5printEP17QPagedPaintDevice
  default:
    qtrt.ErrorResolve("QTextDocument", "print", args)
  }

}


func (this *QTextDocument) docHandle(args ...interface{}) () {
  // docHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9docHandleEv
  default:
    qtrt.ErrorResolve("QTextDocument", "docHandle", args)
  }

}


func (this *QTextDocument) toHtml(args ...interface{}) () {
  // toHtml(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument6toHtmlERK10QByteArray
  default:
    qtrt.ErrorResolve("QTextDocument", "toHtml", args)
  }

}


func (this *QTextDocument) availableUndoSteps(args ...interface{}) () {
  // availableUndoSteps()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument18availableUndoStepsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "availableUndoSteps", args)
  }

}


func (this *QTextDocument) setUndoRedoEnabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument18setUndoRedoEnabledEb
  default:
    qtrt.ErrorResolve("QTextDocument", "setUndoRedoEnabled", args)
  }

}


func (this *QTextDocument) undo(args ...interface{}) () {
  // undo(class QTextCursor *)
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument4undoEP11QTextCursor
  case 1:
    // invoke: _ZN13QTextDocument4undoEv
  default:
    qtrt.ErrorResolve("QTextDocument", "undo", args)
  }

}


func (this *QTextDocument) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11toPlainTextEv
  default:
    qtrt.ErrorResolve("QTextDocument", "toPlainText", args)
  }

}


func (this *QTextDocument) addResource(args ...interface{}) () {
  // addResource(int, const class QUrl &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant
  default:
    qtrt.ErrorResolve("QTextDocument", "addResource", args)
  }

}


func (this *QTextDocument) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument4sizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "size", args)
  }

}


func (this *QTextDocument) object(args ...interface{}) () {
  // object(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument6objectEi
  default:
    qtrt.ErrorResolve("QTextDocument", "object", args)
  }

}


func (this *QTextDocument) clone(args ...interface{}) () {
  // clone(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5cloneEP7QObject
  default:
    qtrt.ErrorResolve("QTextDocument", "clone", args)
  }

}


func (this *QTextDocument) markContentsDirty(args ...interface{}) () {
  // markContentsDirty(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17markContentsDirtyEii
  default:
    qtrt.ErrorResolve("QTextDocument", "markContentsDirty", args)
  }

}


func NewQTextDocument(args ...interface{}) QTextDocument {
  return QTextDocument{}
}


func (this *QTextDocument) characterCount(args ...interface{}) () {
  // characterCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14characterCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "characterCount", args)
  }

}


func (this *QTextDocument) rootFrame(args ...interface{}) () {
  // rootFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9rootFrameEv
  default:
    qtrt.ErrorResolve("QTextDocument", "rootFrame", args)
  }

}


func (this *QTextDocument) firstBlock(args ...interface{}) () {
  // firstBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10firstBlockEv
  default:
    qtrt.ErrorResolve("QTextDocument", "firstBlock", args)
  }

}


func (this *QTextDocument) blockCount(args ...interface{}) () {
  // blockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10blockCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "blockCount", args)
  }

}


func (this *QTextDocument) idealWidth(args ...interface{}) () {
  // idealWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10idealWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "idealWidth", args)
  }

}


func (this *QTextDocument) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument10adjustSizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "adjustSize", args)
  }

}


func (this *QTextDocument) isRedoAvailable(args ...interface{}) () {
  // isRedoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15isRedoAvailableEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isRedoAvailable", args)
  }

}


func (this *QTextDocument) allFormats(args ...interface{}) () {
  // allFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10allFormatsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "allFormats", args)
  }

}


func (this *QTextDocument) defaultStyleSheet(args ...interface{}) () {
  // defaultStyleSheet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17defaultStyleSheetEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultStyleSheet", args)
  }

}


func (this *QTextDocument) lastBlock(args ...interface{}) () {
  // lastBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9lastBlockEv
  default:
    qtrt.ErrorResolve("QTextDocument", "lastBlock", args)
  }

}


func (this *QTextDocument) useDesignMetrics(args ...interface{}) () {
  // useDesignMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument16useDesignMetricsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "useDesignMetrics", args)
  }

}


func (this *QTextDocument) pageCount(args ...interface{}) () {
  // pageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9pageCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "pageCount", args)
  }

}


func (this *QTextDocument) setTextWidth(args ...interface{}) () {
  // setTextWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12setTextWidthEd
  default:
    qtrt.ErrorResolve("QTextDocument", "setTextWidth", args)
  }

}


func (this *QTextDocument) setDocumentMargin(args ...interface{}) () {
  // setDocumentMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17setDocumentMarginEd
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentMargin", args)
  }

}


func (this *QTextDocument) isUndoAvailable(args ...interface{}) () {
  // isUndoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15isUndoAvailableEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoAvailable", args)
  }

}


func (this *QTextDocument) indentWidth(args ...interface{}) () {
  // indentWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11indentWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "indentWidth", args)
  }

}


func (this *QTextDocument) setUseDesignMetrics(args ...interface{}) () {
  // setUseDesignMetrics(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument19setUseDesignMetricsEb
  default:
    qtrt.ErrorResolve("QTextDocument", "setUseDesignMetrics", args)
  }

}


func (this *QTextDocument) setDocumentLayout(args ...interface{}) () {
  // setDocumentLayout(class QAbstractTextDocumentLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractTextDocumentLayout{}) // "QAbstractTextDocumentLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentLayout", args)
  }

}


func (this *QTextDocument) setIndentWidth(args ...interface{}) () {
  // setIndentWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14setIndentWidthEd
  default:
    qtrt.ErrorResolve("QTextDocument", "setIndentWidth", args)
  }

}


func (this *QTextDocument) baseUrl(args ...interface{}) () {
  // baseUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7baseUrlEv
  default:
    qtrt.ErrorResolve("QTextDocument", "baseUrl", args)
  }

}


func (this *QTextDocument) frameAt(args ...interface{}) () {
  // frameAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7frameAtEi
  default:
    qtrt.ErrorResolve("QTextDocument", "frameAt", args)
  }

}


func (this *QTextDocument) setDefaultTextOption(args ...interface{}) () {
  // setDefaultTextOption(const class QTextOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setDefaultTextOptionERK11QTextOption
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultTextOption", args)
  }

}


func (this *QTextDocument) defaultFont(args ...interface{}) () {
  // defaultFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11defaultFontEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultFont", args)
  }

}


func (this *QTextDocument) findBlockByNumber(args ...interface{}) () {
  // findBlockByNumber(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17findBlockByNumberEi
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByNumber", args)
  }

}


func (this *QTextDocument) defaultTextOption(args ...interface{}) () {
  // defaultTextOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17defaultTextOptionEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultTextOption", args)
  }

}


func (this *QTextDocument) findBlock(args ...interface{}) () {
  // findBlock(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9findBlockEi
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlock", args)
  }

}


func (this *QTextDocument) setBaseUrl(args ...interface{}) () {
  // setBaseUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument10setBaseUrlERK4QUrl
  default:
    qtrt.ErrorResolve("QTextDocument", "setBaseUrl", args)
  }

}


func (this *QTextDocument) appendUndoItem(args ...interface{}) () {
  // appendUndoItem(class QAbstractUndoItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractUndoItem{}) // "QAbstractUndoItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem
  default:
    qtrt.ErrorResolve("QTextDocument", "appendUndoItem", args)
  }

}


func (this *QTextDocument) drawContents(args ...interface{}) () {
  // drawContents(class QPainter *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF
  default:
    qtrt.ErrorResolve("QTextDocument", "drawContents", args)
  }

}


func (this *QTextDocument) findBlockByLineNumber(args ...interface{}) () {
  // findBlockByLineNumber(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument21findBlockByLineNumberEi
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByLineNumber", args)
  }

}


func (this *QTextDocument) textWidth(args ...interface{}) () {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9textWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "textWidth", args)
  }

}


func (this *QTextDocument) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10metaObjectEv
  default:
    qtrt.ErrorResolve("QTextDocument", "metaObject", args)
  }

}


func (this *QTextDocument) availableRedoSteps(args ...interface{}) () {
  // availableRedoSteps()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument18availableRedoStepsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "availableRedoSteps", args)
  }

}


func (this *QTextDocument) characterAt(args ...interface{}) () {
  // characterAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11characterAtEi
  default:
    qtrt.ErrorResolve("QTextDocument", "characterAt", args)
  }

}


func (this *QTextDocument) setDefaultFont(args ...interface{}) () {
  // setDefaultFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14setDefaultFontERK5QFont
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultFont", args)
  }

}


func (this *QTextDocument) objectForFormat(args ...interface{}) () {
  // objectForFormat(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15objectForFormatERK11QTextFormat
  default:
    qtrt.ErrorResolve("QTextDocument", "objectForFormat", args)
  }

}


func (this *QTextDocument) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7isEmptyEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isEmpty", args)
  }

}


func (this *QTextDocument) isUndoRedoEnabled(args ...interface{}) () {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17isUndoRedoEnabledEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoRedoEnabled", args)
  }

}


func (this *QTextDocument) FreeQTextDocument(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocument", "~QTextDocument", args)
  }

}


func (this *QTextDocument) documentMargin(args ...interface{}) () {
  // documentMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14documentMarginEv
  default:
    qtrt.ErrorResolve("QTextDocument", "documentMargin", args)
  }

}


func (this *QTextDocument) setPageSize(args ...interface{}) () {
  // setPageSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11setPageSizeERK6QSizeF
  default:
    qtrt.ErrorResolve("QTextDocument", "setPageSize", args)
  }

}


func (this *QTextDocument) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument7setHtmlERK7QString
  default:
    qtrt.ErrorResolve("QTextDocument", "setHtml", args)
  }

}


func (this *QTextDocument) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument3endEv
  default:
    qtrt.ErrorResolve("QTextDocument", "end", args)
  }

}


func (this *QTextDocument) maximumBlockCount(args ...interface{}) () {
  // maximumBlockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17maximumBlockCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "maximumBlockCount", args)
  }

}


func (this *QTextDocument) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12setPlainTextERK7QString
  default:
    qtrt.ErrorResolve("QTextDocument", "setPlainText", args)
  }

}


func (this *QTextDocument) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument5clearEv
  default:
    qtrt.ErrorResolve("QTextDocument", "clear", args)
  }

}


func (this *QTextDocument) resource(args ...interface{}) () {
  // resource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8resourceEiRK4QUrl
  default:
    qtrt.ErrorResolve("QTextDocument", "resource", args)
  }

}


func (this *QTextDocument) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5beginEv
  default:
    qtrt.ErrorResolve("QTextDocument", "begin", args)
  }

}


func (this *QTextDocument) setMaximumBlockCount(args ...interface{}) () {
  // setMaximumBlockCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setMaximumBlockCountEi
  default:
    qtrt.ErrorResolve("QTextDocument", "setMaximumBlockCount", args)
  }

}


func (this *QTextDocument) setModified(args ...interface{}) () {
  // setModified(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11setModifiedEb
  default:
    qtrt.ErrorResolve("QTextDocument", "setModified", args)
  }

}


func (this *QAbstractUndoItem) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractUndoItem4undoEv
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "undo", args)
  }

}


func (this *QAbstractUndoItem) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractUndoItem4redoEv
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "redo", args)
  }

}


func (this *QAbstractUndoItem) FreeQAbstractUndoItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "~QAbstractUndoItem", args)
  }

}

// <= body block end

