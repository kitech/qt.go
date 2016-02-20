package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QTextCursor::isCopyOf(const QTextCursor & other);
extern bool C_ZNK11QTextCursor8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertText(const QString & text, const QTextCharFormat & format);
extern void C_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QTextCursor::insertText(const QString & text);
extern void C_ZN11QTextCursor10insertTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::atStart();
extern bool C_ZNK11QTextCursor7atStartEv(void* qthis); // 4
  // proto:  QTextDocumentFragment QTextCursor::selection();
extern void* C_ZNK11QTextCursor9selectionEv(void* qthis); // 4
  // proto:  QTextFrame * QTextCursor::insertFrame(const QTextFrameFormat & format);
extern void* C_ZN11QTextCursor11insertFrameERK16QTextFrameFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertHtml(const QString & html);
extern void C_ZN11QTextCursor10insertHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::deleteChar();
extern void C_ZN11QTextCursor10deleteCharEv(void* qthis); // 4
  // proto:  void QTextCursor::endEditBlock();
extern void C_ZN11QTextCursor12endEditBlockEv(void* qthis); // 4
  // proto:  bool QTextCursor::atBlockEnd();
extern bool C_ZNK11QTextCursor10atBlockEndEv(void* qthis); // 4
  // proto:  int QTextCursor::selectionStart();
extern int32_t C_ZNK11QTextCursor14selectionStartEv(void* qthis); // 4
  // proto:  int QTextCursor::selectionEnd();
extern int32_t C_ZNK11QTextCursor12selectionEndEv(void* qthis); // 4
  // proto:  void QTextCursor::QTextCursor(QTextDocument * document);
extern void* C_ZN11QTextCursorC2EP13QTextDocument(void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor(const QTextCursor & cursor);
extern void* C_ZN11QTextCursorC2ERKS_(void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor(QTextFrame * frame);
extern void* C_ZN11QTextCursorC2EP10QTextFrame(void* arg0); // 3
  // proto:  void QTextCursor::QTextCursor();
extern void* C_ZN11QTextCursorC2Ev(); // 3
  // proto:  void QTextCursor::QTextCursor(const QTextBlock & block);
extern void* C_ZN11QTextCursorC2ERK10QTextBlock(void* arg0); // 3
  // proto:  QTextTable * QTextCursor::currentTable();
extern void* C_ZNK11QTextCursor12currentTableEv(void* qthis); // 4
  // proto:  void QTextCursor::setKeepPositionOnInsert(bool b);
extern void C_ZN11QTextCursor23setKeepPositionOnInsertEb(void* qthis, bool arg0); // 4
  // proto:  void QTextCursor::clearSelection();
extern void C_ZN11QTextCursor14clearSelectionEv(void* qthis); // 4
  // proto:  bool QTextCursor::atBlockStart();
extern bool C_ZNK11QTextCursor12atBlockStartEv(void* qthis); // 4
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
extern void* C_ZNK11QTextCursor10charFormatEv(void* qthis); // 4
  // proto:  QTextList * QTextCursor::currentList();
extern void* C_ZNK11QTextCursor11currentListEv(void* qthis); // 4
  // proto:  QTextBlockFormat QTextCursor::blockFormat();
extern void* C_ZNK11QTextCursor11blockFormatEv(void* qthis); // 4
  // proto:  void QTextCursor::swap(QTextCursor & other);
extern void C_ZN11QTextCursor4swapERS_(void* qthis, void* arg0); // 2
  // proto:  int QTextCursor::positionInBlock();
extern int32_t C_ZNK11QTextCursor15positionInBlockEv(void* qthis); // 4
  // proto:  QTextDocument * QTextCursor::document();
extern void* C_ZNK11QTextCursor8documentEv(void* qthis); // 4
  // proto:  bool QTextCursor::hasSelection();
extern bool C_ZNK11QTextCursor12hasSelectionEv(void* qthis); // 4
  // proto:  void QTextCursor::mergeBlockFormat(const QTextBlockFormat & modifier);
extern void C_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertFragment(const QTextDocumentFragment & fragment);
extern void C_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::setCharFormat(const QTextCharFormat & format);
extern void C_ZN11QTextCursor13setCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::~QTextCursor();
extern void C_ZN11QTextCursorD2Ev(void* qthis); // 4
  // proto:  QTextCharFormat QTextCursor::blockCharFormat();
extern void* C_ZNK11QTextCursor15blockCharFormatEv(void* qthis); // 4
  // proto:  void QTextCursor::selectedTableCells(int * firstRow, int * numRows, int * firstColumn, int * numColumns);
extern void C_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  int QTextCursor::blockNumber();
extern int32_t C_ZNK11QTextCursor11blockNumberEv(void* qthis); // 4
  // proto:  QString QTextCursor::selectedText();
extern void* C_ZNK11QTextCursor12selectedTextEv(void* qthis); // 4
  // proto:  void QTextCursor::setBlockFormat(const QTextBlockFormat & format);
extern void C_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  QTextList * QTextCursor::insertList(const QTextListFormat & format);
extern void* C_ZN11QTextCursor10insertListERK15QTextListFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::visualNavigation();
extern bool C_ZNK11QTextCursor16visualNavigationEv(void* qthis); // 4
  // proto:  void QTextCursor::insertBlock(const QTextBlockFormat & format);
extern void C_ZN11QTextCursor11insertBlockERK16QTextBlockFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertBlock();
extern void C_ZN11QTextCursor11insertBlockEv(void* qthis); // 4
  // proto:  void QTextCursor::insertBlock(const QTextBlockFormat & format, const QTextCharFormat & charFormat);
extern void C_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QTextFrame * QTextCursor::currentFrame();
extern void* C_ZNK11QTextCursor12currentFrameEv(void* qthis); // 4
  // proto:  void QTextCursor::setBlockCharFormat(const QTextCharFormat & format);
extern void C_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  int QTextCursor::columnNumber();
extern int32_t C_ZNK11QTextCursor12columnNumberEv(void* qthis); // 4
  // proto:  void QTextCursor::insertImage(const QTextImageFormat & format);
extern void C_ZN11QTextCursor11insertImageERK16QTextImageFormat(void* qthis, void* arg0); // 4
  // proto:  void QTextCursor::insertImage(const QImage & image, const QString & name);
extern void C_ZN11QTextCursor11insertImageERK6QImageRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QTextCursor::insertImage(const QString & name);
extern void C_ZN11QTextCursor11insertImageERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextCursor::keepPositionOnInsert();
extern bool C_ZNK11QTextCursor20keepPositionOnInsertEv(void* qthis); // 4
  // proto:  void QTextCursor::joinPreviousEditBlock();
extern void C_ZN11QTextCursor21joinPreviousEditBlockEv(void* qthis); // 4
  // proto:  bool QTextCursor::atEnd();
extern bool C_ZNK11QTextCursor5atEndEv(void* qthis); // 4
  // proto:  QTextTable * QTextCursor::insertTable(int rows, int cols);
extern void* C_ZN11QTextCursor11insertTableEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTextTable * QTextCursor::insertTable(int rows, int cols, const QTextTableFormat & format);
extern void* C_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QTextCursor::anchor();
extern int32_t C_ZNK11QTextCursor6anchorEv(void* qthis); // 4
  // proto:  bool QTextCursor::isNull();
extern bool C_ZNK11QTextCursor6isNullEv(void* qthis); // 4
  // proto:  int QTextCursor::verticalMovementX();
extern int32_t C_ZNK11QTextCursor17verticalMovementXEv(void* qthis); // 4
  // proto:  void QTextCursor::removeSelectedText();
extern void C_ZN11QTextCursor18removeSelectedTextEv(void* qthis); // 4
  // proto:  QTextList * QTextCursor::createList(const QTextListFormat & format);
extern void* C_ZN11QTextCursor10createListERK15QTextListFormat(void* qthis, void* arg0); // 4
  // proto:  int QTextCursor::position();
extern int32_t C_ZNK11QTextCursor8positionEv(void* qthis); // 4
  // proto:  bool QTextCursor::hasComplexSelection();
extern bool C_ZNK11QTextCursor19hasComplexSelectionEv(void* qthis); // 4
  // proto:  QTextBlock QTextCursor::block();
extern void* C_ZNK11QTextCursor5blockEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextCursor)=1
type QTextCursor struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isCopyOf(const class QTextCursor &)
func (this *QTextCursor) Iscopyof(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTextCursor8isCopyOfERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "isCopyOf", args)
  }

  return
}

// insertText(const class QString &, const class QTextCharFormat &)
func (this *QTextCursor) Inserttext(args ...interface{}) () {
  // insertText(const class QString &, const class QTextCharFormat &)
  // insertText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat
    // invoke: void insertText(const class QString &, const class QTextCharFormat &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QTextCursor10insertTextERK7QString
    // invoke: void insertText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10insertTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertText", args)
  }

  return
}

// atStart()
func (this *QTextCursor) Atstart(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor7atStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "atStart", args)
  }

  return
}

// selection()
func (this *QTextCursor) Selection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor9selectionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocumentFragment{}) // "QTextDocumentFragment"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "selection", args)
  }

  return
}

// insertFrame(const class QTextFrameFormat &)
func (this *QTextCursor) Insertframe(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTextFrameFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QTextCursor11insertFrameERK16QTextFrameFormat(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFrame", args)
  }

  return
}

// insertHtml(const class QString &)
func (this *QTextCursor) Inserthtml(args ...interface{}) () {
  // insertHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor10insertHtmlERK7QString
    // invoke: void insertHtml(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor10insertHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertHtml", args)
  }

  return
}

// deleteChar()
func (this *QTextCursor) Deletechar(args ...interface{}) () {
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
    C.C_ZN11QTextCursor10deleteCharEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "deleteChar", args)
  }

  return
}

// endEditBlock()
func (this *QTextCursor) Endeditblock(args ...interface{}) () {
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
    C.C_ZN11QTextCursor12endEditBlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "endEditBlock", args)
  }

  return
}

// atBlockEnd()
func (this *QTextCursor) Atblockend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor10atBlockEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockEnd", args)
  }

  return
}

// selectionStart()
func (this *QTextCursor) Selectionstart(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor14selectionStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionStart", args)
  }

  return
}

// selectionEnd()
func (this *QTextCursor) Selectionend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12selectionEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "selectionEnd", args)
  }

  return
}

// QTextCursor(class QTextDocument *)
func NewQTextCursor(args ...interface{}) *QTextCursor {
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
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextCursorC2EP13QTextDocument(arg0)
    return &QTextCursor{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QTextCursorC1ERKS_
    // invoke: void QTextCursor(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextCursorC2ERKS_(arg0)
    return &QTextCursor{Qclsinst:qthis}
  case 2:
    // invoke: _ZN11QTextCursorC1EP10QTextFrame
    // invoke: void QTextCursor(class QTextFrame *)
    var arg0 = args[0].(*QTextFrame).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextCursorC2EP10QTextFrame(arg0)
    return &QTextCursor{Qclsinst:qthis}
  case 3:
    // invoke: _ZN11QTextCursorC1Ev
    // invoke: void QTextCursor()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextCursorC2Ev()
    return &QTextCursor{Qclsinst:qthis}
  case 4:
    // invoke: _ZN11QTextCursorC1ERK10QTextBlock
    // invoke: void QTextCursor(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextCursorC2ERK10QTextBlock(arg0)
    return &QTextCursor{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextCursor", "QTextCursor", args)
  }

  return nil // QTextCursor{Qclsinst:qthis}
}

// currentTable()
func (this *QTextCursor) Currenttable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12currentTableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTable{}) // "QTextTable *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "currentTable", args)
  }

  return
}

// setKeepPositionOnInsert(_Bool)
func (this *QTextCursor) Setkeeppositiononinsert(args ...interface{}) () {
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
    C.C_ZN11QTextCursor23setKeepPositionOnInsertEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setKeepPositionOnInsert", args)
  }

  return
}

// clearSelection()
func (this *QTextCursor) Clearselection(args ...interface{}) () {
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
    C.C_ZN11QTextCursor14clearSelectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "clearSelection", args)
  }

  return
}

// atBlockStart()
func (this *QTextCursor) Atblockstart(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12atBlockStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "atBlockStart", args)
  }

  return
}

// setVerticalMovementX(int)
func (this *QTextCursor) Setverticalmovementx(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor20setVerticalMovementXEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setVerticalMovementX", args)
  }

  return
}

// deletePreviousChar()
func (this *QTextCursor) Deletepreviouschar(args ...interface{}) () {
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
    C.C_ZN11QTextCursor18deletePreviousCharEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "deletePreviousChar", args)
  }

  return
}

// mergeCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) Mergecharformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeCharFormat", args)
  }

  return
}

// setVisualNavigation(_Bool)
func (this *QTextCursor) Setvisualnavigation(args ...interface{}) () {
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
    C.C_ZN11QTextCursor19setVisualNavigationEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setVisualNavigation", args)
  }

  return
}

// mergeBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) Mergeblockcharformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockCharFormat", args)
  }

  return
}

// beginEditBlock()
func (this *QTextCursor) Begineditblock(args ...interface{}) () {
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
    C.C_ZN11QTextCursor14beginEditBlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "beginEditBlock", args)
  }

  return
}

// charFormat()
func (this *QTextCursor) Charformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor10charFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "charFormat", args)
  }

  return
}

// currentList()
func (this *QTextCursor) Currentlist(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor11currentListEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextList{}) // "QTextList *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "currentList", args)
  }

  return
}

// blockFormat()
func (this *QTextCursor) Blockformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor11blockFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlockFormat{}) // "QTextBlockFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "blockFormat", args)
  }

  return
}

// swap(class QTextCursor &)
func (this *QTextCursor) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "swap", args)
  }

  return
}

// positionInBlock()
func (this *QTextCursor) Positioninblock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor15positionInBlockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "positionInBlock", args)
  }

  return
}

// document()
func (this *QTextCursor) Document(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "document", args)
  }

  return
}

// hasSelection()
func (this *QTextCursor) Hasselection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12hasSelectionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "hasSelection", args)
  }

  return
}

// mergeBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) Mergeblockformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextBlockFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "mergeBlockFormat", args)
  }

  return
}

// insertFragment(const class QTextDocumentFragment &)
func (this *QTextCursor) Insertfragment(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextDocumentFragment).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertFragment", args)
  }

  return
}

// setCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) Setcharformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor13setCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setCharFormat", args)
  }

  return
}

// ~QTextCursor()
func (this *QTextCursor) Freeqtextcursor(args ...interface{}) () {
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
    C.C_ZN11QTextCursorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "~QTextCursor", args)
  }

  return
}

// blockCharFormat()
func (this *QTextCursor) Blockcharformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor15blockCharFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "blockCharFormat", args)
  }

  return
}

// selectedTableCells(int *, int *, int *, int *)
func (this *QTextCursor) Selectedtablecells(args ...interface{}) () {
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
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedTableCells", args)
  }

  return
}

// blockNumber()
func (this *QTextCursor) Blocknumber(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor11blockNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "blockNumber", args)
  }

  return
}

// selectedText()
func (this *QTextCursor) Selectedtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12selectedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "selectedText", args)
  }

  return
}

// setBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) Setblockformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextBlockFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockFormat", args)
  }

  return
}

// insertList(const class QTextListFormat &)
func (this *QTextCursor) Insertlist(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTextListFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QTextCursor10insertListERK15QTextListFormat(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextList{}) // "QTextList *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "insertList", args)
  }

  return
}

// visualNavigation()
func (this *QTextCursor) Visualnavigation(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor16visualNavigationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "visualNavigation", args)
  }

  return
}

// insertBlock(const class QTextBlockFormat &)
func (this *QTextCursor) Insertblock(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextBlockFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertBlockERK16QTextBlockFormat(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextCursor11insertBlockEv
    // invoke: void insertBlock()
    C.C_ZN11QTextCursor11insertBlockEv(this.Qclsinst)
  case 2:
    // invoke: _ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat
    // invoke: void insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
    var arg0 = args[0].(*QTextBlockFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertBlock", args)
  }

  return
}

// currentFrame()
func (this *QTextCursor) Currentframe(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12currentFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "currentFrame", args)
  }

  return
}

// setBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) Setblockcharformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "setBlockCharFormat", args)
  }

  return
}

// columnNumber()
func (this *QTextCursor) Columnnumber(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor12columnNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "columnNumber", args)
  }

  return
}

// insertImage(const class QTextImageFormat &)
func (this *QTextCursor) Insertimage(args ...interface{}) () {
  // insertImage(const class QTextImageFormat &)
  // insertImage(const class QImage &, const class QString &)
  // insertImage(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextImageFormat{}) // "const QTextImageFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextCursor11insertImageERK16QTextImageFormat
    // invoke: void insertImage(const class QTextImageFormat &)
    var arg0 = args[0].(*QTextImageFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertImageERK16QTextImageFormat(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextCursor11insertImageERK6QImageRK7QString
    // invoke: void insertImage(const class QImage &, const class QString &)
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextCursor11insertImageERK6QImageRK7QString(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN11QTextCursor11insertImageERK7QString
    // invoke: void insertImage(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextCursor11insertImageERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCursor", "insertImage", args)
  }

  return
}

// keepPositionOnInsert()
func (this *QTextCursor) Keeppositiononinsert(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor20keepPositionOnInsertEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "keepPositionOnInsert", args)
  }

  return
}

// joinPreviousEditBlock()
func (this *QTextCursor) Joinpreviouseditblock(args ...interface{}) () {
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
    C.C_ZN11QTextCursor21joinPreviousEditBlockEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "joinPreviousEditBlock", args)
  }

  return
}

// atEnd()
func (this *QTextCursor) Atend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "atEnd", args)
  }

  return
}

// insertTable(int, int)
func (this *QTextCursor) Inserttable(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QTextCursor11insertTableEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTable{}) // "QTextTable *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QTextCursor11insertTableEiiRK16QTextTableFormat
    // invoke: QTextTable * insertTable(int, int, const class QTextTableFormat &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTextTableFormat).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTable{}) // "QTextTable *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "insertTable", args)
  }

  return
}

// anchor()
func (this *QTextCursor) Anchor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor6anchorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "anchor", args)
  }

  return
}

// isNull()
func (this *QTextCursor) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "isNull", args)
  }

  return
}

// verticalMovementX()
func (this *QTextCursor) Verticalmovementx(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor17verticalMovementXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "verticalMovementX", args)
  }

  return
}

// removeSelectedText()
func (this *QTextCursor) Removeselectedtext(args ...interface{}) () {
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
    C.C_ZN11QTextCursor18removeSelectedTextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCursor", "removeSelectedText", args)
  }

  return
}

// createList(const class QTextListFormat &)
func (this *QTextCursor) Createlist(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTextListFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QTextCursor10createListERK15QTextListFormat(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextList{}) // "QTextList *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "createList", args)
  }

  return
}

// position()
func (this *QTextCursor) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "position", args)
  }

  return
}

// hasComplexSelection()
func (this *QTextCursor) Hascomplexselection(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor19hasComplexSelectionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "hasComplexSelection", args)
  }

  return
}

// block()
func (this *QTextCursor) Block(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextCursor5blockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCursor", "block", args)
  }

  return
}

// <= body block end

