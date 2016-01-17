package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qplaintextedit.h
// dst-file: /src/widgets/qplaintextedit.go
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
  // proto:  void QPlainTextDocumentLayout::ensureBlockLayout(const QTextBlock & block);
extern void _ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextDocumentLayout::pageCount();
extern void _ZNK24QPlainTextDocumentLayout9pageCountEv(void* qthis); // 4
  // proto:  QRectF QPlainTextDocumentLayout::blockBoundingRect(const QTextBlock & block);
extern void _ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  QRectF QPlainTextDocumentLayout::frameBoundingRect(QTextFrame * );
extern void _ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextDocumentLayout::cursorWidth();
extern void _ZNK24QPlainTextDocumentLayout11cursorWidthEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::requestUpdate();
extern void _ZN24QPlainTextDocumentLayout13requestUpdateEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::setCursorWidth(int width);
extern void _ZN24QPlainTextDocumentLayout14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSizeF QPlainTextDocumentLayout::documentSize();
extern void _ZNK24QPlainTextDocumentLayout12documentSizeEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::~QPlainTextDocumentLayout();
extern void _ZN24QPlainTextDocumentLayoutD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QPlainTextDocumentLayout::metaObject();
extern void _ZNK24QPlainTextDocumentLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::QPlainTextDocumentLayout(QTextDocument * document);
extern void _ZN24QPlainTextDocumentLayoutC2EP13QTextDocument(void* qthis, void* arg0); // 3
  // proto:  void QPlainTextEdit::appendPlainText(const QString & text);
extern void _ZN14QPlainTextEdit15appendPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QPlainTextEdit::documentTitle();
extern void _ZNK14QPlainTextEdit13documentTitleEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::setTabStopWidth(int width);
extern void _ZN14QPlainTextEdit15setTabStopWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPlainTextEdit::ensureCursorVisible();
extern void _ZN14QPlainTextEdit19ensureCursorVisibleEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void _ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextEdit::tabStopWidth();
extern void _ZNK14QPlainTextEdit12tabStopWidthEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::isUndoRedoEnabled();
extern void _ZNK14QPlainTextEdit17isUndoRedoEnabledEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::cut();
extern void _ZN14QPlainTextEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu();
extern void _ZN14QPlainTextEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu(const QPoint & position);
extern void _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QList<QTextEdit::ExtraSelection> QPlainTextEdit::extraSelections();
extern void _ZNK14QPlainTextEdit15extraSelectionsEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void _ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QPlainTextEdit::metaObject();
extern void _ZNK14QPlainTextEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setTabChangesFocus(bool b);
extern void _ZN14QPlainTextEdit18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  bool QPlainTextEdit::centerOnScroll();
extern void _ZNK14QPlainTextEdit14centerOnScrollEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::toPlainText();
extern void _ZNK14QPlainTextEdit11toPlainTextEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::zoomIn(int range);
extern void _ZN14QPlainTextEdit6zoomInEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPlainTextEdit::backgroundVisible();
extern void _ZNK14QPlainTextEdit17backgroundVisibleEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::zoomOut(int range);
extern void _ZN14QPlainTextEdit7zoomOutEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QPlainTextEdit::cursorRect(const QTextCursor & cursor);
extern void _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QPlainTextEdit::cursorRect();
extern void _ZNK14QPlainTextEdit10cursorRectEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setUndoRedoEnabled(bool enable);
extern void _ZN14QPlainTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0); // 2
  // proto:  QTextCursor QPlainTextEdit::textCursor();
extern void _ZNK14QPlainTextEdit10textCursorEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::anchorAt(const QPoint & pos);
extern void _ZNK14QPlainTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setDocument(QTextDocument * document);
extern void _ZN14QPlainTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::print(QPagedPaintDevice * printer);
extern void _ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QVariant QPlainTextEdit::loadResource(int type, const QUrl & name);
extern void _ZN14QPlainTextEdit12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QPlainTextEdit::setCenterOnScroll(bool enabled);
extern void _ZN14QPlainTextEdit17setCenterOnScrollEb(void* qthis, bool arg0); // 4
  // proto:  void QPlainTextEdit::undo();
extern void _ZN14QPlainTextEdit4undoEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::paste();
extern void _ZN14QPlainTextEdit5pasteEv(void* qthis); // 4
  // proto:  int QPlainTextEdit::blockCount();
extern void _ZNK14QPlainTextEdit10blockCountEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::placeholderText();
extern void _ZNK14QPlainTextEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::selectAll();
extern void _ZN14QPlainTextEdit9selectAllEv(void* qthis); // 4
  // proto:  int QPlainTextEdit::maximumBlockCount();
extern void _ZNK14QPlainTextEdit17maximumBlockCountEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::redo();
extern void _ZN14QPlainTextEdit4redoEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setBackgroundVisible(bool visible);
extern void _ZN14QPlainTextEdit20setBackgroundVisibleEb(void* qthis, bool arg0); // 4
  // proto:  bool QPlainTextEdit::isReadOnly();
extern void _ZNK14QPlainTextEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setTextCursor(const QTextCursor & cursor);
extern void _ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  bool QPlainTextEdit::overwriteMode();
extern void _ZNK14QPlainTextEdit13overwriteModeEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setCursorWidth(int width);
extern void _ZN14QPlainTextEdit14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPlainTextEdit::appendHtml(const QString & html);
extern void _ZN14QPlainTextEdit10appendHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::TextInteractionFlags QPlainTextEdit::textInteractionFlags();
extern void _ZNK14QPlainTextEdit20textInteractionFlagsEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::canPaste();
extern void _ZNK14QPlainTextEdit8canPasteEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setOverwriteMode(bool overwrite);
extern void _ZN14QPlainTextEdit16setOverwriteModeEb(void* qthis, bool arg0); // 4
  // proto:  void QPlainTextEdit::copy();
extern void _ZN14QPlainTextEdit4copyEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::~QPlainTextEdit();
extern void _ZN14QPlainTextEditD2Ev(void* qthis); // 4
  // proto:  void QPlainTextEdit::setPlaceholderText(const QString & placeholderText);
extern void _ZN14QPlainTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setMaximumBlockCount(int maximum);
extern void _ZN14QPlainTextEdit20setMaximumBlockCountEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPlainTextEdit::QPlainTextEdit(const QString & text, QWidget * parent);
extern void _ZN14QPlainTextEditC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QPlainTextEdit::QPlainTextEdit(QWidget * parent);
extern void _ZN14QPlainTextEditC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QPlainTextEdit::setPlainText(const QString & text);
extern void _ZN14QPlainTextEdit12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::insertPlainText(const QString & text);
extern void _ZN14QPlainTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextEdit::cursorWidth();
extern void _ZNK14QPlainTextEdit11cursorWidthEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::tabChangesFocus();
extern void _ZNK14QPlainTextEdit15tabChangesFocusEv(void* qthis); // 4
  // proto:  QTextOption::WrapMode QPlainTextEdit::wordWrapMode();
extern void _ZNK14QPlainTextEdit12wordWrapModeEv(void* qthis); // 4
  // proto:  QTextCursor QPlainTextEdit::cursorForPosition(const QPoint & pos);
extern void _ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setDocumentTitle(const QString & title);
extern void _ZN14QPlainTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0); // 2
  // proto:  QTextDocument * QPlainTextEdit::document();
extern void _ZNK14QPlainTextEdit8documentEv(void* qthis); // 4
  // proto:  QTextCharFormat QPlainTextEdit::currentCharFormat();
extern void _ZNK14QPlainTextEdit17currentCharFormatEv(void* qthis); // 4
  // proto:  QPlainTextEdit::LineWrapMode QPlainTextEdit::lineWrapMode();
extern void _ZNK14QPlainTextEdit12lineWrapModeEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::clear();
extern void _ZN14QPlainTextEdit5clearEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::centerCursor();
extern void _ZN14QPlainTextEdit12centerCursorEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setReadOnly(bool ro);
extern void _ZN14QPlainTextEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
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

// class sizeof(QPlainTextDocumentLayout)=1
type QPlainTextDocumentLayout struct {
  /*qbase*/ QAbstractTextDocumentLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPlainTextEdit)=1
type QPlainTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst unsafe.Pointer /* *C.void */;
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

// ensureBlockLayout(const class QTextBlock &)
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
    // invoke: void ensureBlockLayout(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "ensureBlockLayout", args)
  }

}

// pageCount()
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
    // invoke: int pageCount()
    C._ZNK24QPlainTextDocumentLayout9pageCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "pageCount", args)
  }

}

// blockBoundingRect(const class QTextBlock &)
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
    // invoke: QRectF blockBoundingRect(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "blockBoundingRect", args)
  }

}

// frameBoundingRect(class QTextFrame *)
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
    // invoke: QRectF frameBoundingRect(class QTextFrame *)
    var arg0 = args[0].(QTextFrame).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "frameBoundingRect", args)
  }

}

// cursorWidth()
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
    // invoke: int cursorWidth()
    C._ZNK24QPlainTextDocumentLayout11cursorWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "cursorWidth", args)
  }

}

// requestUpdate()
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
    // invoke: void requestUpdate()
    C._ZN24QPlainTextDocumentLayout13requestUpdateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "requestUpdate", args)
  }

}

// setCursorWidth(int)
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
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN24QPlainTextDocumentLayout14setCursorWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "setCursorWidth", args)
  }

}

// documentSize()
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
    // invoke: QSizeF documentSize()
    C._ZNK24QPlainTextDocumentLayout12documentSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "documentSize", args)
  }

}

// ~QPlainTextDocumentLayout()
func (this *QPlainTextDocumentLayout) FreeQPlainTextDocumentLayout(args ...interface{}) () {
  // ~QPlainTextDocumentLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayoutD0Ev
    // invoke: void ~QPlainTextDocumentLayout()
    C._ZN24QPlainTextDocumentLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "~QPlainTextDocumentLayout", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK24QPlainTextDocumentLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "metaObject", args)
  }

}

// QPlainTextDocumentLayout(class QTextDocument *)
func NewQPlainTextDocumentLayout(args ...interface{}) QPlainTextDocumentLayout {
  // QPlainTextDocumentLayout(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayoutC1EP13QTextDocument
    // invoke: void QPlainTextDocumentLayout(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN24QPlainTextDocumentLayoutC2EP13QTextDocument(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "QPlainTextDocumentLayout", args)
  }

  return QPlainTextDocumentLayout{}
}

// appendPlainText(const class QString &)
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
    // invoke: void appendPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15appendPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendPlainText", args)
  }

}

// documentTitle()
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
    // invoke: QString documentTitle()
    C._ZNK14QPlainTextEdit13documentTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "documentTitle", args)
  }

}

// setTabStopWidth(int)
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
    // invoke: void setTabStopWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15setTabStopWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabStopWidth", args)
  }

}

// ensureCursorVisible()
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
    // invoke: void ensureCursorVisible()
    C._ZN14QPlainTextEdit19ensureCursorVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "ensureCursorVisible", args)
  }

}

// setCurrentCharFormat(const class QTextCharFormat &)
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
    // invoke: void setCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCurrentCharFormat", args)
  }

}

// tabStopWidth()
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
    // invoke: int tabStopWidth()
    C._ZNK14QPlainTextEdit12tabStopWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabStopWidth", args)
  }

}

// isUndoRedoEnabled()
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
    // invoke: bool isUndoRedoEnabled()
    C._ZNK14QPlainTextEdit17isUndoRedoEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isUndoRedoEnabled", args)
  }

}

// cut()
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
    // invoke: void cut()
    C._ZN14QPlainTextEdit3cutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cut", args)
  }

}

// createStandardContextMenu()
func (this *QPlainTextEdit) createStandardContextMenu(args ...interface{}) () {
  // createStandardContextMenu()
  // createStandardContextMenu(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    C._ZN14QPlainTextEdit25createStandardContextMenuEv(this.qclsinst)
  case 1:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint
    // invoke: QMenu * createStandardContextMenu(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "createStandardContextMenu", args)
  }

}

// extraSelections()
func (this *QPlainTextEdit) extraSelections(args ...interface{}) () {
  // extraSelections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit15extraSelectionsEv
    // invoke: QList<QTextEdit::ExtraSelection> extraSelections()
    C._ZNK14QPlainTextEdit15extraSelectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "extraSelections", args)
  }

}

// mergeCurrentCharFormat(const class QTextCharFormat &)
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
    // invoke: void mergeCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "mergeCurrentCharFormat", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QPlainTextEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "metaObject", args)
  }

}

// setTabChangesFocus(_Bool)
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
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit18setTabChangesFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabChangesFocus", args)
  }

}

// centerOnScroll()
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
    // invoke: bool centerOnScroll()
    C._ZNK14QPlainTextEdit14centerOnScrollEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerOnScroll", args)
  }

}

// toPlainText()
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
    // invoke: QString toPlainText()
    C._ZNK14QPlainTextEdit11toPlainTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "toPlainText", args)
  }

}

// zoomIn(int)
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
    // invoke: void zoomIn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit6zoomInEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomIn", args)
  }

}

// backgroundVisible()
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
    // invoke: bool backgroundVisible()
    C._ZNK14QPlainTextEdit17backgroundVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "backgroundVisible", args)
  }

}

// zoomOut(int)
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
    // invoke: void zoomOut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit7zoomOutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomOut", args)
  }

}

// cursorRect(const class QTextCursor &)
func (this *QPlainTextEdit) cursorRect(args ...interface{}) () {
  // cursorRect(const class QTextCursor &)
  // cursorRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor
    // invoke: QRect cursorRect(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QPlainTextEdit10cursorRectEv
    // invoke: QRect cursorRect()
    C._ZNK14QPlainTextEdit10cursorRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorRect", args)
  }

}

// setUndoRedoEnabled(_Bool)
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
    // invoke: void setUndoRedoEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit18setUndoRedoEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setUndoRedoEnabled", args)
  }

}

// textCursor()
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
    // invoke: QTextCursor textCursor()
    C._ZNK14QPlainTextEdit10textCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textCursor", args)
  }

}

// anchorAt(const class QPoint &)
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
    // invoke: QString anchorAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit8anchorAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "anchorAt", args)
  }

}

// setDocument(class QTextDocument *)
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
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit11setDocumentEP13QTextDocument(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocument", args)
  }

}

// print(class QPagedPaintDevice *)
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
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "print", args)
  }

}

// loadResource(int, const class QUrl &)
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
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QPlainTextEdit12loadResourceEiRK4QUrl(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "loadResource", args)
  }

}

// setCenterOnScroll(_Bool)
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
    // invoke: void setCenterOnScroll(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit17setCenterOnScrollEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCenterOnScroll", args)
  }

}

// undo()
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
    // invoke: void undo()
    C._ZN14QPlainTextEdit4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "undo", args)
  }

}

// paste()
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
    // invoke: void paste()
    C._ZN14QPlainTextEdit5pasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "paste", args)
  }

}

// blockCount()
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
    // invoke: int blockCount()
    C._ZNK14QPlainTextEdit10blockCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "blockCount", args)
  }

}

// placeholderText()
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
    // invoke: QString placeholderText()
    C._ZNK14QPlainTextEdit15placeholderTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "placeholderText", args)
  }

}

// selectAll()
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
    // invoke: void selectAll()
    C._ZN14QPlainTextEdit9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "selectAll", args)
  }

}

// maximumBlockCount()
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
    // invoke: int maximumBlockCount()
    C._ZNK14QPlainTextEdit17maximumBlockCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "maximumBlockCount", args)
  }

}

// redo()
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
    // invoke: void redo()
    C._ZN14QPlainTextEdit4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "redo", args)
  }

}

// setBackgroundVisible(_Bool)
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
    // invoke: void setBackgroundVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit20setBackgroundVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setBackgroundVisible", args)
  }

}

// isReadOnly()
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
    // invoke: bool isReadOnly()
    C._ZNK14QPlainTextEdit10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isReadOnly", args)
  }

}

// setTextCursor(const class QTextCursor &)
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
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTextCursor", args)
  }

}

// overwriteMode()
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
    // invoke: bool overwriteMode()
    C._ZNK14QPlainTextEdit13overwriteModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "overwriteMode", args)
  }

}

// setCursorWidth(int)
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
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit14setCursorWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCursorWidth", args)
  }

}

// appendHtml(const class QString &)
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
    // invoke: void appendHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit10appendHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendHtml", args)
  }

}

// textInteractionFlags()
func (this *QPlainTextEdit) textInteractionFlags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C._ZNK14QPlainTextEdit20textInteractionFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textInteractionFlags", args)
  }

}

// canPaste()
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
    // invoke: bool canPaste()
    C._ZNK14QPlainTextEdit8canPasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "canPaste", args)
  }

}

// setOverwriteMode(_Bool)
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
    // invoke: void setOverwriteMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit16setOverwriteModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setOverwriteMode", args)
  }

}

// copy()
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
    // invoke: void copy()
    C._ZN14QPlainTextEdit4copyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "copy", args)
  }

}

// ~QPlainTextEdit()
func (this *QPlainTextEdit) FreeQPlainTextEdit(args ...interface{}) () {
  // ~QPlainTextEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEditD0Ev
    // invoke: void ~QPlainTextEdit()
    C._ZN14QPlainTextEditD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "~QPlainTextEdit", args)
  }

}

// setPlaceholderText(const class QString &)
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
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit18setPlaceholderTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlaceholderText", args)
  }

}

// setMaximumBlockCount(int)
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
    // invoke: void setMaximumBlockCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit20setMaximumBlockCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setMaximumBlockCount", args)
  }

}

// QPlainTextEdit(const class QString &, class QWidget *)
func NewQPlainTextEdit(args ...interface{}) QPlainTextEdit {
  // QPlainTextEdit(const class QString &, class QWidget *)
  // QPlainTextEdit(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEditC1ERK7QStringP7QWidget
    // invoke: void QPlainTextEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QPlainTextEditC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN14QPlainTextEditC1EP7QWidget
    // invoke: void QPlainTextEdit(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QPlainTextEditC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "QPlainTextEdit", args)
  }

  return QPlainTextEdit{}
}

// setPlainText(const class QString &)
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
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit12setPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlainText", args)
  }

}

// insertPlainText(const class QString &)
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
    // invoke: void insertPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15insertPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "insertPlainText", args)
  }

}

// cursorWidth()
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
    // invoke: int cursorWidth()
    C._ZNK14QPlainTextEdit11cursorWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorWidth", args)
  }

}

// tabChangesFocus()
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
    // invoke: bool tabChangesFocus()
    C._ZNK14QPlainTextEdit15tabChangesFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabChangesFocus", args)
  }

}

// wordWrapMode()
func (this *QPlainTextEdit) wordWrapMode(args ...interface{}) () {
  // wordWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit12wordWrapModeEv
    // invoke: QTextOption::WrapMode wordWrapMode()
    C._ZNK14QPlainTextEdit12wordWrapModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "wordWrapMode", args)
  }

}

// cursorForPosition(const class QPoint &)
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
    // invoke: QTextCursor cursorForPosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorForPosition", args)
  }

}

// setDocumentTitle(const class QString &)
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
    // invoke: void setDocumentTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit16setDocumentTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocumentTitle", args)
  }

}

// document()
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
    // invoke: QTextDocument * document()
    C._ZNK14QPlainTextEdit8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "document", args)
  }

}

// currentCharFormat()
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
    // invoke: QTextCharFormat currentCharFormat()
    C._ZNK14QPlainTextEdit17currentCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "currentCharFormat", args)
  }

}

// lineWrapMode()
func (this *QPlainTextEdit) lineWrapMode(args ...interface{}) () {
  // lineWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit12lineWrapModeEv
    // invoke: QPlainTextEdit::LineWrapMode lineWrapMode()
    C._ZNK14QPlainTextEdit12lineWrapModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "lineWrapMode", args)
  }

}

// clear()
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
    // invoke: void clear()
    C._ZN14QPlainTextEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "clear", args)
  }

}

// centerCursor()
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
    // invoke: void centerCursor()
    C._ZN14QPlainTextEdit12centerCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerCursor", args)
  }

}

// setReadOnly(_Bool)
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
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setReadOnly", args)
  }

}

// <= body block end

