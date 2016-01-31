package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qtextcursor.h
// dst-file: /src/gui/qtextcursor.go
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
  // proto:  bool QTextCursor::isCopyOf(const QTextCursor & other);
extern void C_ZNK11QTextCursor8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertText(const QString & text, const QTextCharFormat & format);
extern void C_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QTextCursor::insertText(const QString & text);
extern void C_ZN11QTextCursor10insertTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::atStart();
extern void C_ZNK11QTextCursor7atStartEv(void* qthis); // 4
  // proto:  QTextDocumentFragment QTextCursor::selection();
extern void C_ZNK11QTextCursor9selectionEv(void* qthis); // 4
  // proto:  QTextFrame * QTextCursor::insertFrame(const QTextFrameFormat & format);
extern void C_ZN11QTextCursor11insertFrameERK16QTextFrameFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertHtml(const QString & html);
extern void C_ZN11QTextCursor10insertHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::deleteChar();
extern void C_ZN11QTextCursor10deleteCharEv(void* qthis); // 4
  // proto:  void QTextCursor::endEditBlock();
extern void C_ZN11QTextCursor12endEditBlockEv(void* qthis); // 4
  // proto:  bool QTextCursor::atBlockEnd();
extern void C_ZNK11QTextCursor10atBlockEndEv(void* qthis); // 4
  // proto:  int QTextCursor::selectionStart();
extern void C_ZNK11QTextCursor14selectionStartEv(void* qthis); // 4
  // proto:  int QTextCursor::selectionEnd();
extern void C_ZNK11QTextCursor12selectionEndEv(void* qthis); // 4
  // proto:  void QTextCursor::QTextCursor(QTextDocument * document);
extern void C_ZN11QTextCursorC2EP13QTextDocument(void* qthis, void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor(const QTextCursor & cursor);
extern void C_ZN11QTextCursorC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor(QTextFrame * frame);
extern void C_ZN11QTextCursorC2EP10QTextFrame(void* qthis, void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor();
extern void C_ZN11QTextCursorC2Ev(void* qthis); // 3
  // proto:  void QTextCursor::QTextCursor(const QTextBlock & block);
extern void C_ZN11QTextCursorC2ERK10QTextBlock(void* qthis, void* arg0); // 3
  // proto:  QTextTable * QTextCursor::currentTable();
extern void C_ZNK11QTextCursor12currentTableEv(void* qthis); // 4
  // proto:  void QTextCursor::setKeepPositionOnInsert(bool b);
extern void C_ZN11QTextCursor23setKeepPositionOnInsertEb(void* qthis, bool arg0); // 4
  // proto:  void QTextCursor::clearSelection();
extern void C_ZN11QTextCursor14clearSelectionEv(void* qthis); // 4
  // proto:  bool QTextCursor::atBlockStart();
extern void C_ZNK11QTextCursor12atBlockStartEv(void* qthis); // 4
  // proto:  void QTextCursor::setVerticalMovementX(int x);
extern void C_ZN11QTextCursor20setVerticalMovementXEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextCursor::deletePreviousChar();
extern void C_ZN11QTextCursor18deletePreviousCharEv(void* qthis); // 4
  // proto:  void QTextCursor::mergeCharFormat(const QTextCharFormat & modifier);
extern void C_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::setVisualNavigation(bool b);
extern void C_ZN11QTextCursor19setVisualNavigationEb(void* qthis, bool arg0); // 4
  // proto:  void QTextCursor::mergeBlockCharFormat(const QTextCharFormat & modifier);
extern void C_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::beginEditBlock();
extern void C_ZN11QTextCursor14beginEditBlockEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextCursor::charFormat();
extern void C_ZNK11QTextCursor10charFormatEv(void* qthis); // 4
  // proto:  QTextList * QTextCursor::currentList();
extern void C_ZNK11QTextCursor11currentListEv(void* qthis); // 4
  // proto:  QTextBlockFormat QTextCursor::blockFormat();
extern void C_ZNK11QTextCursor11blockFormatEv(void* qthis); // 4
  // proto:  void QTextCursor::swap(QTextCursor & other);
extern void C_ZN11QTextCursor4swapERS_(void* qthis, void* arg0); // 2
  // proto:  int QTextCursor::positionInBlock();
extern void C_ZNK11QTextCursor15positionInBlockEv(void* qthis); // 4
  // proto:  QTextDocument * QTextCursor::document();
extern void C_ZNK11QTextCursor8documentEv(void* qthis); // 4
  // proto:  bool QTextCursor::hasSelection();
extern void C_ZNK11QTextCursor12hasSelectionEv(void* qthis); // 4
  // proto:  void QTextCursor::mergeBlockFormat(const QTextBlockFormat & modifier);
extern void C_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertFragment(const QTextDocumentFragment & fragment);
extern void C_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::setCharFormat(const QTextCharFormat & format);
extern void C_ZN11QTextCursor13setCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::~QTextCursor();
extern void C_ZN11QTextCursorD2Ev(void* qthis); // 4
  // proto:  QTextCharFormat QTextCursor::blockCharFormat();
extern void C_ZNK11QTextCursor15blockCharFormatEv(void* qthis); // 4
  // proto:  void QTextCursor::selectedTableCells(int * firstRow, int * numRows, int * firstColumn, int * numColumns);
extern void C_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  int QTextCursor::blockNumber();
extern void C_ZNK11QTextCursor11blockNumberEv(void* qthis); // 4
  // proto:  QString QTextCursor::selectedText();
extern void C_ZNK11QTextCursor12selectedTextEv(void* qthis); // 4
  // proto:  void QTextCursor::setBlockFormat(const QTextBlockFormat & format);
extern void C_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  QTextList * QTextCursor::insertList(const QTextListFormat & format);
extern void C_ZN11QTextCursor10insertListERK15QTextListFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::visualNavigation();
extern void C_ZNK11QTextCursor16visualNavigationEv(void* qthis); // 4
  // proto:  void QTextCursor::insertBlock(const QTextBlockFormat & format);
extern void C_ZN11QTextCursor11insertBlockERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertBlock();
extern void C_ZN11QTextCursor11insertBlockEv(void* qthis); // 4
  // proto:  void QTextCursor::insertBlock(const QTextBlockFormat & format, const QTextCharFormat & charFormat);
extern void C_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QTextFrame * QTextCursor::currentFrame();
extern void C_ZNK11QTextCursor12currentFrameEv(void* qthis); // 4
  // proto:  void QTextCursor::setBlockCharFormat(const QTextCharFormat & format);
extern void C_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  int QTextCursor::columnNumber();
extern void C_ZNK11QTextCursor12columnNumberEv(void* qthis); // 4
  // proto:  void QTextCursor::insertImage(const QTextImageFormat & format);
extern void C_ZN11QTextCursor11insertImageERK16QTextImageFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertImage(const QImage & image, const QString & name);
extern void C_ZN11QTextCursor11insertImageERK6QImageRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QTextCursor::insertImage(const QString & name);
extern void C_ZN11QTextCursor11insertImageERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::keepPositionOnInsert();
extern void C_ZNK11QTextCursor20keepPositionOnInsertEv(void* qthis); // 4
  // proto:  void QTextCursor::joinPreviousEditBlock();
extern void C_ZN11QTextCursor21joinPreviousEditBlockEv(void* qthis); // 4
  // proto:  bool QTextCursor::atEnd();
extern void C_ZNK11QTextCursor5atEndEv(void* qthis); // 4
  // proto:  QTextTable * QTextCursor::insertTable(int rows, int cols);
extern void C_ZN11QTextCursor11insertTableEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTextTable * QTextCursor::insertTable(int rows, int cols, const QTextTableFormat & format);
extern void C_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QTextCursor::anchor();
extern void C_ZNK11QTextCursor6anchorEv(void* qthis); // 4
  // proto:  bool QTextCursor::isNull();
extern void C_ZNK11QTextCursor6isNullEv(void* qthis); // 4
  // proto:  int QTextCursor::verticalMovementX();
extern void C_ZNK11QTextCursor17verticalMovementXEv(void* qthis); // 4
  // proto:  void QTextCursor::removeSelectedText();
extern void C_ZN11QTextCursor18removeSelectedTextEv(void* qthis); // 4
  // proto:  QTextList * QTextCursor::createList(const QTextListFormat & format);
extern void C_ZN11QTextCursor10createListERK15QTextListFormat(void* qthis, void* arg0); // 4
  // proto:  int QTextCursor::position();
extern void C_ZNK11QTextCursor8positionEv(void* qthis); // 4
  // proto:  bool QTextCursor::hasComplexSelection();
extern void C_ZNK11QTextCursor19hasComplexSelectionEv(void* qthis); // 4
  // proto:  QTextBlock QTextCursor::block();
extern void C_ZNK11QTextCursor5blockEv(void* qthis); // 4
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

// class sizeof(QTextCursor)=1
type QTextCursor struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isCopyOf(const class QTextCursor &)
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
    // invoke: bool isCopyOf(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QTextCursor8isCopyOfERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "isCopyOf", args)
  }

}

// insertText(const class QString &, const class QTextCharFormat &)
func (this *QTextCursor) insertText(args ...interface{}) () {
  // insertText(const class QString &, const class QTextCharFormat &)
  // insertText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat
    // invoke: void insertText(const class QString &, const class QTextCharFormat &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QTextCursor10insertTextERK7QString
    // invoke: void insertText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10insertTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertText", args)
  }

}

// atStart()
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
    // invoke: bool atStart()
    C.C_ZNK11QTextCursor7atStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "atStart", args)
  }

}

// selection()
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
    // invoke: QTextDocumentFragment selection()
    C.C_ZNK11QTextCursor9selectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "selection", args)
  }

}

// insertFrame(const class QTextFrameFormat &)
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
    // invoke: QTextFrame * insertFrame(const class QTextFrameFormat &)
    var arg0 = args[0].(QTextFrameFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertFrameERK16QTextFrameFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFrame", args)
  }

}

// insertHtml(const class QString &)
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
    // invoke: void insertHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10insertHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertHtml", args)
  }

}

// deleteChar()
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
    // invoke: void deleteChar()
    C.C_ZN11QTextCursor10deleteCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "deleteChar", args)
  }

}

// endEditBlock()
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
    // invoke: void endEditBlock()
    C.C_ZN11QTextCursor12endEditBlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "endEditBlock", args)
  }

}

// atBlockEnd()
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
    // invoke: bool atBlockEnd()
    C.C_ZNK11QTextCursor10atBlockEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockEnd", args)
  }

}

// selectionStart()
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
    // invoke: int selectionStart()
    C.C_ZNK11QTextCursor14selectionStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionStart", args)
  }

}

// selectionEnd()
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
    // invoke: int selectionEnd()
    C.C_ZNK11QTextCursor12selectionEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionEnd", args)
  }

}

// QTextCursor(class QTextDocument *)
func NewQTextCursor(args ...interface{}) QTextCursor {
  // QTextCursor(class QTextDocument *)
  // QTextCursor(const class QTextCursor &)
  // QTextCursor(class QTextFrame *)
  // QTextCursor()
  // QTextCursor(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursorC1EP13QTextDocument
    // invoke: void QTextCursor(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextCursorC2EP13QTextDocument(qthis, arg0)
  case 1:
    // invoke: _ZN11QTextCursorC1ERKS_
    // invoke: void QTextCursor(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextCursorC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN11QTextCursorC1EP10QTextFrame
    // invoke: void QTextCursor(class QTextFrame *)
    var arg0 = args[0].(QTextFrame).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextCursorC2EP10QTextFrame(qthis, arg0)
  case 3:
    // invoke: _ZN11QTextCursorC1Ev
    // invoke: void QTextCursor()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextCursorC2Ev(qthis)
  case 4:
    // invoke: _ZN11QTextCursorC1ERK10QTextBlock
    // invoke: void QTextCursor(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextCursorC2ERK10QTextBlock(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "QTextCursor", args)
  }

  return QTextCursor{}
}

// currentTable()
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
    // invoke: QTextTable * currentTable()
    C.C_ZNK11QTextCursor12currentTableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "currentTable", args)
  }

}

// setKeepPositionOnInsert(_Bool)
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
    // invoke: void setKeepPositionOnInsert(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor23setKeepPositionOnInsertEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setKeepPositionOnInsert", args)
  }

}

// clearSelection()
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
    // invoke: void clearSelection()
    C.C_ZN11QTextCursor14clearSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "clearSelection", args)
  }

}

// atBlockStart()
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
    // invoke: bool atBlockStart()
    C.C_ZNK11QTextCursor12atBlockStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockStart", args)
  }

}

// setVerticalMovementX(int)
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
    // invoke: void setVerticalMovementX(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor20setVerticalMovementXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setVerticalMovementX", args)
  }

}

// deletePreviousChar()
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
    // invoke: void deletePreviousChar()
    C.C_ZN11QTextCursor18deletePreviousCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "deletePreviousChar", args)
  }

}

// mergeCharFormat(const class QTextCharFormat &)
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
    // invoke: void mergeCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeCharFormat", args)
  }

}

// setVisualNavigation(_Bool)
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
    // invoke: void setVisualNavigation(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor19setVisualNavigationEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setVisualNavigation", args)
  }

}

// mergeBlockCharFormat(const class QTextCharFormat &)
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
    // invoke: void mergeBlockCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockCharFormat", args)
  }

}

// beginEditBlock()
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
    // invoke: void beginEditBlock()
    C.C_ZN11QTextCursor14beginEditBlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "beginEditBlock", args)
  }

}

// charFormat()
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
    // invoke: QTextCharFormat charFormat()
    C.C_ZNK11QTextCursor10charFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "charFormat", args)
  }

}

// currentList()
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
    // invoke: QTextList * currentList()
    C.C_ZNK11QTextCursor11currentListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "currentList", args)
  }

}

// blockFormat()
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
    // invoke: QTextBlockFormat blockFormat()
    C.C_ZNK11QTextCursor11blockFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "blockFormat", args)
  }

}

// swap(class QTextCursor &)
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
    // invoke: void swap(class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "swap", args)
  }

}

// positionInBlock()
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
    // invoke: int positionInBlock()
    C.C_ZNK11QTextCursor15positionInBlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "positionInBlock", args)
  }

}

// document()
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
    // invoke: QTextDocument * document()
    C.C_ZNK11QTextCursor8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "document", args)
  }

}

// hasSelection()
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
    // invoke: bool hasSelection()
    C.C_ZNK11QTextCursor12hasSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "hasSelection", args)
  }

}

// mergeBlockFormat(const class QTextBlockFormat &)
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
    // invoke: void mergeBlockFormat(const class QTextBlockFormat &)
    var arg0 = args[0].(QTextBlockFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockFormat", args)
  }

}

// insertFragment(const class QTextDocumentFragment &)
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
    // invoke: void insertFragment(const class QTextDocumentFragment &)
    var arg0 = args[0].(QTextDocumentFragment).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFragment", args)
  }

}

// setCharFormat(const class QTextCharFormat &)
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
    // invoke: void setCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor13setCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setCharFormat", args)
  }

}

// ~QTextCursor()
func (this *QTextCursor) FreeQTextCursor(args ...interface{}) () {
  // ~QTextCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursorD0Ev
    // invoke: void ~QTextCursor()
    C.C_ZN11QTextCursorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "~QTextCursor", args)
  }

}

// blockCharFormat()
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
    // invoke: QTextCharFormat blockCharFormat()
    C.C_ZNK11QTextCursor15blockCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "blockCharFormat", args)
  }

}

// selectedTableCells(int *, int *, int *, int *)
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
    // invoke: void selectedTableCells(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedTableCells", args)
  }

}

// blockNumber()
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
    // invoke: int blockNumber()
    C.C_ZNK11QTextCursor11blockNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "blockNumber", args)
  }

}

// selectedText()
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
    // invoke: QString selectedText()
    C.C_ZNK11QTextCursor12selectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedText", args)
  }

}

// setBlockFormat(const class QTextBlockFormat &)
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
    // invoke: void setBlockFormat(const class QTextBlockFormat &)
    var arg0 = args[0].(QTextBlockFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockFormat", args)
  }

}

// insertList(const class QTextListFormat &)
func (this *QTextCursor) insertList(args ...interface{}) () {
  // insertList(const class QTextListFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertListERK15QTextListFormat
    // invoke: QTextList * insertList(const class QTextListFormat &)
    var arg0 = args[0].(QTextListFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10insertListERK15QTextListFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertList", args)
  }

}

// visualNavigation()
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
    // invoke: bool visualNavigation()
    C.C_ZNK11QTextCursor16visualNavigationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "visualNavigation", args)
  }

}

// insertBlock(const class QTextBlockFormat &)
func (this *QTextCursor) insertBlock(args ...interface{}) () {
  // insertBlock(const class QTextBlockFormat &)
  // insertBlock()
  // insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextBlockFormat{}) // "const QTextBlockFormat &"
  vtys[2][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertBlockERK16QTextBlockFormat
    // invoke: void insertBlock(const class QTextBlockFormat &)
    var arg0 = args[0].(QTextBlockFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertBlockERK16QTextBlockFormat(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextCursor11insertBlockEv
    // invoke: void insertBlock()
    C.C_ZN11QTextCursor11insertBlockEv(this.qclsinst)
  case 2:
    // invoke: _ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat
    // invoke: void insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
    var arg0 = args[0].(QTextBlockFormat).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertBlock", args)
  }

}

// currentFrame()
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
    // invoke: QTextFrame * currentFrame()
    C.C_ZNK11QTextCursor12currentFrameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "currentFrame", args)
  }

}

// setBlockCharFormat(const class QTextCharFormat &)
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
    // invoke: void setBlockCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockCharFormat", args)
  }

}

// columnNumber()
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
    // invoke: int columnNumber()
    C.C_ZNK11QTextCursor12columnNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "columnNumber", args)
  }

}

// insertImage(const class QTextImageFormat &)
func (this *QTextCursor) insertImage(args ...interface{}) () {
  // insertImage(const class QTextImageFormat &)
  // insertImage(const class QImage &, const class QString &)
  // insertImage(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextImageFormat{}) // "const QTextImageFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertImageERK16QTextImageFormat
    // invoke: void insertImage(const class QTextImageFormat &)
    var arg0 = args[0].(QTextImageFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertImageERK16QTextImageFormat(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextCursor11insertImageERK6QImageRK7QString
    // invoke: void insertImage(const class QImage &, const class QString &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor11insertImageERK6QImageRK7QString(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN11QTextCursor11insertImageERK7QString
    // invoke: void insertImage(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertImageERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertImage", args)
  }

}

// keepPositionOnInsert()
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
    // invoke: bool keepPositionOnInsert()
    C.C_ZNK11QTextCursor20keepPositionOnInsertEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "keepPositionOnInsert", args)
  }

}

// joinPreviousEditBlock()
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
    // invoke: void joinPreviousEditBlock()
    C.C_ZN11QTextCursor21joinPreviousEditBlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "joinPreviousEditBlock", args)
  }

}

// atEnd()
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
    // invoke: bool atEnd()
    C.C_ZNK11QTextCursor5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "atEnd", args)
  }

}

// insertTable(int, int)
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
    // invoke: QTextTable * insertTable(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor11insertTableEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QTextCursor11insertTableEiiRK16QTextTableFormat
    // invoke: QTextTable * insertTable(int, int, const class QTextTableFormat &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTextTableFormat).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertTable", args)
  }

}

// anchor()
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
    // invoke: int anchor()
    C.C_ZNK11QTextCursor6anchorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "anchor", args)
  }

}

// isNull()
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
    // invoke: bool isNull()
    C.C_ZNK11QTextCursor6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "isNull", args)
  }

}

// verticalMovementX()
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
    // invoke: int verticalMovementX()
    C.C_ZNK11QTextCursor17verticalMovementXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "verticalMovementX", args)
  }

}

// removeSelectedText()
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
    // invoke: void removeSelectedText()
    C.C_ZN11QTextCursor18removeSelectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "removeSelectedText", args)
  }

}

// createList(const class QTextListFormat &)
func (this *QTextCursor) createList(args ...interface{}) () {
  // createList(const class QTextListFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10createListERK15QTextListFormat
    // invoke: QTextList * createList(const class QTextListFormat &)
    var arg0 = args[0].(QTextListFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10createListERK15QTextListFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "createList", args)
  }

}

// position()
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
    // invoke: int position()
    C.C_ZNK11QTextCursor8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "position", args)
  }

}

// hasComplexSelection()
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
    // invoke: bool hasComplexSelection()
    C.C_ZNK11QTextCursor19hasComplexSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "hasComplexSelection", args)
  }

}

// block()
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
    // invoke: QTextBlock block()
    C.C_ZNK11QTextCursor5blockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "block", args)
  }

}

// <= body block end

