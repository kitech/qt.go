package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qtextcursor.h
// dst-file: /src/gui/qtextcursor.go
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
// class sizeof(QTextCursor)=1
type QTextCursor struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextCursor) columnNumber(args ...interface{}) () {
  // columnNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12columnNumberEv
  default:
    qtrt.ErrorResolve("QTextCursor", "columnNumber", args)
  }

}


func (this *QTextCursor) swap(args ...interface{}) () {
  // swap(class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor4swapERS_
  default:
    qtrt.ErrorResolve("QTextCursor", "swap", args)
  }

}


func (this *QTextCursor) mergeCharFormat(args ...interface{}) () {
  // mergeCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeCharFormat", args)
  }

}


func (this *QTextCursor) selection(args ...interface{}) () {
  // selection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor9selectionEv
  default:
    qtrt.ErrorResolve("QTextCursor", "selection", args)
  }

}


func (this *QTextCursor) hasComplexSelection(args ...interface{}) () {
  // hasComplexSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor19hasComplexSelectionEv
  default:
    qtrt.ErrorResolve("QTextCursor", "hasComplexSelection", args)
  }

}


func (this *QTextCursor) block(args ...interface{}) () {
  // block()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor5blockEv
  default:
    qtrt.ErrorResolve("QTextCursor", "block", args)
  }

}


func (this *QTextCursor) insertFragment(args ...interface{}) () {
  // insertFragment(const class QTextDocumentFragment &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocumentFragment{}) // "const QTextDocumentFragment &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFragment", args)
  }

}


func (this *QTextCursor) insertList(args ...interface{}) () {
  // insertList(const class QTextListFormat &)
  // insertList(class QTextListFormat::Style)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QTextListFormat::Style"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertListERK15QTextListFormat
  case 1:
    // invoke: _ZN11QTextCursor10insertListEN15QTextListFormat5StyleE
  default:
    qtrt.ErrorResolve("QTextCursor", "insertList", args)
  }

}


func (this *QTextCursor) insertImage(args ...interface{}) () {
  // insertImage(const class QTextImageFormat &)
  // insertImage(const class QString &)
  // insertImage(const class QImage &, const class QString &)
  // insertImage(const class QTextImageFormat &, class QTextFrameFormat::Position)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextImageFormat{}) // "const QTextImageFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QTextImageFormat{}) // "const QTextImageFormat &"
  vtys[3][1] = qtrt.Int32Ty(false) // "QTextFrameFormat::Position"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertImageERK16QTextImageFormat
  case 1:
    // invoke: _ZN11QTextCursor11insertImageERK7QString
  case 2:
    // invoke: _ZN11QTextCursor11insertImageERK6QImageRK7QString
  case 3:
    // invoke: _ZN11QTextCursor11insertImageERK16QTextImageFormatN16QTextFrameFormat8PositionE
  default:
    qtrt.ErrorResolve("QTextCursor", "insertImage", args)
  }

}


func (this *QTextCursor) keepPositionOnInsert(args ...interface{}) () {
  // keepPositionOnInsert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor20keepPositionOnInsertEv
  default:
    qtrt.ErrorResolve("QTextCursor", "keepPositionOnInsert", args)
  }

}


func (this *QTextCursor) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor8positionEv
  default:
    qtrt.ErrorResolve("QTextCursor", "position", args)
  }

}


func (this *QTextCursor) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor6isNullEv
  default:
    qtrt.ErrorResolve("QTextCursor", "isNull", args)
  }

}


func (this *QTextCursor) removeSelectedText(args ...interface{}) () {
  // removeSelectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor18removeSelectedTextEv
  default:
    qtrt.ErrorResolve("QTextCursor", "removeSelectedText", args)
  }

}


func (this *QTextCursor) insertHtml(args ...interface{}) () {
  // insertHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertHtmlERK7QString
  default:
    qtrt.ErrorResolve("QTextCursor", "insertHtml", args)
  }

}


func (this *QTextCursor) isCopyOf(args ...interface{}) () {
  // isCopyOf(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor8isCopyOfERKS_
  default:
    qtrt.ErrorResolve("QTextCursor", "isCopyOf", args)
  }

}


func (this *QTextCursor) insertFrame(args ...interface{}) () {
  // insertFrame(const class QTextFrameFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrameFormat{}) // "const QTextFrameFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertFrameERK16QTextFrameFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFrame", args)
  }

}


func NewQTextCursor(args ...interface{}) QTextCursor {
  return QTextCursor{}
}


func (this *QTextCursor) deleteChar(args ...interface{}) () {
  // deleteChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10deleteCharEv
  default:
    qtrt.ErrorResolve("QTextCursor", "deleteChar", args)
  }

}


func (this *QTextCursor) currentFrame(args ...interface{}) () {
  // currentFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12currentFrameEv
  default:
    qtrt.ErrorResolve("QTextCursor", "currentFrame", args)
  }

}


func (this *QTextCursor) insertBlock(args ...interface{}) () {
  // insertBlock()
  // insertBlock(const class QTextBlockFormat &)
  // insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"
  vtys[2][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertBlockEv
  case 1:
    // invoke: _ZN11QTextCursor11insertBlockERK16QTextBlockFormat
  case 2:
    // invoke: _ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "insertBlock", args)
  }

}


func (this *QTextCursor) insertTable(args ...interface{}) () {
  // insertTable(int, int)
  // insertTable(int, int, const class QTextTableFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QTextTableFormat{}) // "const QTextTableFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertTableEii
  case 1:
    // invoke: _ZN11QTextCursor11insertTableEiiRK16QTextTableFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "insertTable", args)
  }

}


func (this *QTextCursor) atStart(args ...interface{}) () {
  // atStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor7atStartEv
  default:
    qtrt.ErrorResolve("QTextCursor", "atStart", args)
  }

}


func (this *QTextCursor) selectionStart(args ...interface{}) () {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor14selectionStartEv
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionStart", args)
  }

}


func (this *QTextCursor) selectedTableCells(args ...interface{}) () {
  // selectedTableCells(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedTableCells", args)
  }

}


func (this *QTextCursor) endEditBlock(args ...interface{}) () {
  // endEditBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor12endEditBlockEv
  default:
    qtrt.ErrorResolve("QTextCursor", "endEditBlock", args)
  }

}


func (this *QTextCursor) selectedText(args ...interface{}) () {
  // selectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12selectedTextEv
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedText", args)
  }

}


func (this *QTextCursor) positionInBlock(args ...interface{}) () {
  // positionInBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor15positionInBlockEv
  default:
    qtrt.ErrorResolve("QTextCursor", "positionInBlock", args)
  }

}


func (this *QTextCursor) hasSelection(args ...interface{}) () {
  // hasSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12hasSelectionEv
  default:
    qtrt.ErrorResolve("QTextCursor", "hasSelection", args)
  }

}


func (this *QTextCursor) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor5atEndEv
  default:
    qtrt.ErrorResolve("QTextCursor", "atEnd", args)
  }

}


func (this *QTextCursor) atBlockStart(args ...interface{}) () {
  // atBlockStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12atBlockStartEv
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockStart", args)
  }

}


func (this *QTextCursor) insertText(args ...interface{}) () {
  // insertText(const class QString &)
  // insertText(const class QString &, const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertTextERK7QString
  case 1:
    // invoke: _ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "insertText", args)
  }

}


func (this *QTextCursor) visualNavigation(args ...interface{}) () {
  // visualNavigation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor16visualNavigationEv
  default:
    qtrt.ErrorResolve("QTextCursor", "visualNavigation", args)
  }

}


func (this *QTextCursor) atBlockEnd(args ...interface{}) () {
  // atBlockEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor10atBlockEndEv
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockEnd", args)
  }

}


func (this *QTextCursor) currentList(args ...interface{}) () {
  // currentList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor11currentListEv
  default:
    qtrt.ErrorResolve("QTextCursor", "currentList", args)
  }

}


func (this *QTextCursor) mergeBlockCharFormat(args ...interface{}) () {
  // mergeBlockCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockCharFormat", args)
  }

}


func (this *QTextCursor) setCharFormat(args ...interface{}) () {
  // setCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor13setCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "setCharFormat", args)
  }

}


func (this *QTextCursor) verticalMovementX(args ...interface{}) () {
  // verticalMovementX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor17verticalMovementXEv
  default:
    qtrt.ErrorResolve("QTextCursor", "verticalMovementX", args)
  }

}


func (this *QTextCursor) blockNumber(args ...interface{}) () {
  // blockNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor11blockNumberEv
  default:
    qtrt.ErrorResolve("QTextCursor", "blockNumber", args)
  }

}


func (this *QTextCursor) joinPreviousEditBlock(args ...interface{}) () {
  // joinPreviousEditBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor21joinPreviousEditBlockEv
  default:
    qtrt.ErrorResolve("QTextCursor", "joinPreviousEditBlock", args)
  }

}


func (this *QTextCursor) mergeBlockFormat(args ...interface{}) () {
  // mergeBlockFormat(const class QTextBlockFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockFormat", args)
  }

}


func (this *QTextCursor) blockFormat(args ...interface{}) () {
  // blockFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor11blockFormatEv
  default:
    qtrt.ErrorResolve("QTextCursor", "blockFormat", args)
  }

}


func (this *QTextCursor) beginEditBlock(args ...interface{}) () {
  // beginEditBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor14beginEditBlockEv
  default:
    qtrt.ErrorResolve("QTextCursor", "beginEditBlock", args)
  }

}


func (this *QTextCursor) anchor(args ...interface{}) () {
  // anchor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor6anchorEv
  default:
    qtrt.ErrorResolve("QTextCursor", "anchor", args)
  }

}


func (this *QTextCursor) charFormat(args ...interface{}) () {
  // charFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor10charFormatEv
  default:
    qtrt.ErrorResolve("QTextCursor", "charFormat", args)
  }

}


func (this *QTextCursor) deletePreviousChar(args ...interface{}) () {
  // deletePreviousChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor18deletePreviousCharEv
  default:
    qtrt.ErrorResolve("QTextCursor", "deletePreviousChar", args)
  }

}


func (this *QTextCursor) FreeQTextCursor(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCursor", "~QTextCursor", args)
  }

}


func (this *QTextCursor) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor14clearSelectionEv
  default:
    qtrt.ErrorResolve("QTextCursor", "clearSelection", args)
  }

}


func (this *QTextCursor) setVisualNavigation(args ...interface{}) () {
  // setVisualNavigation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor19setVisualNavigationEb
  default:
    qtrt.ErrorResolve("QTextCursor", "setVisualNavigation", args)
  }

}


func (this *QTextCursor) setBlockCharFormat(args ...interface{}) () {
  // setBlockCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockCharFormat", args)
  }

}


func (this *QTextCursor) currentTable(args ...interface{}) () {
  // currentTable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12currentTableEv
  default:
    qtrt.ErrorResolve("QTextCursor", "currentTable", args)
  }

}


func (this *QTextCursor) setKeepPositionOnInsert(args ...interface{}) () {
  // setKeepPositionOnInsert(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor23setKeepPositionOnInsertEb
  default:
    qtrt.ErrorResolve("QTextCursor", "setKeepPositionOnInsert", args)
  }

}


func (this *QTextCursor) setVerticalMovementX(args ...interface{}) () {
  // setVerticalMovementX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor20setVerticalMovementXEi
  default:
    qtrt.ErrorResolve("QTextCursor", "setVerticalMovementX", args)
  }

}


func (this *QTextCursor) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor8documentEv
  default:
    qtrt.ErrorResolve("QTextCursor", "document", args)
  }

}


func (this *QTextCursor) selectionEnd(args ...interface{}) () {
  // selectionEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor12selectionEndEv
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionEnd", args)
  }

}


func (this *QTextCursor) setBlockFormat(args ...interface{}) () {
  // setBlockFormat(const class QTextBlockFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockFormat", args)
  }

}


func (this *QTextCursor) createList(args ...interface{}) () {
  // createList(class QTextListFormat::Style)
  // createList(const class QTextListFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QTextListFormat::Style"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10createListEN15QTextListFormat5StyleE
  case 1:
    // invoke: _ZN11QTextCursor10createListERK15QTextListFormat
  default:
    qtrt.ErrorResolve("QTextCursor", "createList", args)
  }

}


func (this *QTextCursor) blockCharFormat(args ...interface{}) () {
  // blockCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextCursor15blockCharFormatEv
  default:
    qtrt.ErrorResolve("QTextCursor", "blockCharFormat", args)
  }

}

// <= body block end

