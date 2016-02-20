package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPlainTextDocumentLayout::ensureBlockLayout(const QTextBlock & block);
extern void C_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextDocumentLayout::pageCount();
extern int32_t C_ZNK24QPlainTextDocumentLayout9pageCountEv(void* qthis); // 4
  // proto:  QRectF QPlainTextDocumentLayout::blockBoundingRect(const QTextBlock & block);
extern void* C_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  QRectF QPlainTextDocumentLayout::frameBoundingRect(QTextFrame * );
extern void* C_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextDocumentLayout::cursorWidth();
extern int32_t C_ZNK24QPlainTextDocumentLayout11cursorWidthEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::requestUpdate();
extern void C_ZN24QPlainTextDocumentLayout13requestUpdateEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::setCursorWidth(int width);
extern void C_ZN24QPlainTextDocumentLayout14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSizeF QPlainTextDocumentLayout::documentSize();
extern void* C_ZNK24QPlainTextDocumentLayout12documentSizeEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::~QPlainTextDocumentLayout();
extern void C_ZN24QPlainTextDocumentLayoutD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QPlainTextDocumentLayout::metaObject();
extern void C_ZNK24QPlainTextDocumentLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QPlainTextDocumentLayout::QPlainTextDocumentLayout(QTextDocument * document);
extern void* C_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument(void* arg0); // 3
  // proto:  void QPlainTextEdit::appendPlainText(const QString & text);
extern void C_ZN14QPlainTextEdit15appendPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QPlainTextEdit::documentTitle();
extern void* C_ZNK14QPlainTextEdit13documentTitleEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::setTabStopWidth(int width);
extern void C_ZN14QPlainTextEdit15setTabStopWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPlainTextEdit::ensureCursorVisible();
extern void C_ZN14QPlainTextEdit19ensureCursorVisibleEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void C_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextEdit::tabStopWidth();
extern int32_t C_ZNK14QPlainTextEdit12tabStopWidthEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::isUndoRedoEnabled();
extern bool C_ZNK14QPlainTextEdit17isUndoRedoEnabledEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::cut();
extern void C_ZN14QPlainTextEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu();
extern void* C_ZN14QPlainTextEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu(const QPoint & position);
extern void* C_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QList<QTextEdit::ExtraSelection> QPlainTextEdit::extraSelections();
extern void C_ZNK14QPlainTextEdit15extraSelectionsEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void C_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QPlainTextEdit::metaObject();
extern void C_ZNK14QPlainTextEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setTabChangesFocus(bool b);
extern void C_ZN14QPlainTextEdit18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  bool QPlainTextEdit::centerOnScroll();
extern bool C_ZNK14QPlainTextEdit14centerOnScrollEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::toPlainText();
extern void* C_ZNK14QPlainTextEdit11toPlainTextEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::zoomIn(int range);
extern void C_ZN14QPlainTextEdit6zoomInEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPlainTextEdit::backgroundVisible();
extern bool C_ZNK14QPlainTextEdit17backgroundVisibleEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::zoomOut(int range);
extern void C_ZN14QPlainTextEdit7zoomOutEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QPlainTextEdit::cursorRect(const QTextCursor & cursor);
extern void* C_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QPlainTextEdit::cursorRect();
extern void* C_ZNK14QPlainTextEdit10cursorRectEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setUndoRedoEnabled(bool enable);
extern void C_ZN14QPlainTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0); // 2
  // proto:  QTextCursor QPlainTextEdit::textCursor();
extern void* C_ZNK14QPlainTextEdit10textCursorEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::anchorAt(const QPoint & pos);
extern void* C_ZNK14QPlainTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setDocument(QTextDocument * document);
extern void C_ZN14QPlainTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::print(QPagedPaintDevice * printer);
extern void C_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QVariant QPlainTextEdit::loadResource(int type, const QUrl & name);
extern void* C_ZN14QPlainTextEdit12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QPlainTextEdit::setCenterOnScroll(bool enabled);
extern void C_ZN14QPlainTextEdit17setCenterOnScrollEb(void* qthis, bool arg0); // 4
  // proto:  void QPlainTextEdit::undo();
extern void C_ZN14QPlainTextEdit4undoEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::paste();
extern void C_ZN14QPlainTextEdit5pasteEv(void* qthis); // 4
  // proto:  int QPlainTextEdit::blockCount();
extern int32_t C_ZNK14QPlainTextEdit10blockCountEv(void* qthis); // 4
  // proto:  QString QPlainTextEdit::placeholderText();
extern void* C_ZNK14QPlainTextEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::selectAll();
extern void C_ZN14QPlainTextEdit9selectAllEv(void* qthis); // 4
  // proto:  int QPlainTextEdit::maximumBlockCount();
extern int32_t C_ZNK14QPlainTextEdit17maximumBlockCountEv(void* qthis); // 2
  // proto:  void QPlainTextEdit::redo();
extern void C_ZN14QPlainTextEdit4redoEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setBackgroundVisible(bool visible);
extern void C_ZN14QPlainTextEdit20setBackgroundVisibleEb(void* qthis, bool arg0); // 4
  // proto:  bool QPlainTextEdit::isReadOnly();
extern bool C_ZNK14QPlainTextEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setTextCursor(const QTextCursor & cursor);
extern void C_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  bool QPlainTextEdit::overwriteMode();
extern bool C_ZNK14QPlainTextEdit13overwriteModeEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setCursorWidth(int width);
extern void C_ZN14QPlainTextEdit14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QPlainTextEdit::appendHtml(const QString & html);
extern void C_ZN14QPlainTextEdit10appendHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::TextInteractionFlags QPlainTextEdit::textInteractionFlags();
extern void C_ZNK14QPlainTextEdit20textInteractionFlagsEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::canPaste();
extern bool C_ZNK14QPlainTextEdit8canPasteEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setOverwriteMode(bool overwrite);
extern void C_ZN14QPlainTextEdit16setOverwriteModeEb(void* qthis, bool arg0); // 4
  // proto:  void QPlainTextEdit::copy();
extern void C_ZN14QPlainTextEdit4copyEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::~QPlainTextEdit();
extern void C_ZN14QPlainTextEditD2Ev(void* qthis); // 4
  // proto:  void QPlainTextEdit::setPlaceholderText(const QString & placeholderText);
extern void C_ZN14QPlainTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setMaximumBlockCount(int maximum);
extern void C_ZN14QPlainTextEdit20setMaximumBlockCountEi(void* qthis, int32_t arg0); // 2
  // proto:  void QPlainTextEdit::QPlainTextEdit(const QString & text, QWidget * parent);
extern void* C_ZN14QPlainTextEditC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QPlainTextEdit::QPlainTextEdit(QWidget * parent);
extern void* C_ZN14QPlainTextEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QPlainTextEdit::setPlainText(const QString & text);
extern void C_ZN14QPlainTextEdit12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::insertPlainText(const QString & text);
extern void C_ZN14QPlainTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QPlainTextEdit::cursorWidth();
extern int32_t C_ZNK14QPlainTextEdit11cursorWidthEv(void* qthis); // 4
  // proto:  bool QPlainTextEdit::tabChangesFocus();
extern bool C_ZNK14QPlainTextEdit15tabChangesFocusEv(void* qthis); // 4
  // proto:  QTextOption::WrapMode QPlainTextEdit::wordWrapMode();
extern void C_ZNK14QPlainTextEdit12wordWrapModeEv(void* qthis); // 4
  // proto:  QTextCursor QPlainTextEdit::cursorForPosition(const QPoint & pos);
extern void* C_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QPlainTextEdit::setDocumentTitle(const QString & title);
extern void C_ZN14QPlainTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0); // 2
  // proto:  QTextDocument * QPlainTextEdit::document();
extern void* C_ZNK14QPlainTextEdit8documentEv(void* qthis); // 4
  // proto:  QTextCharFormat QPlainTextEdit::currentCharFormat();
extern void* C_ZNK14QPlainTextEdit17currentCharFormatEv(void* qthis); // 4
  // proto:  QPlainTextEdit::LineWrapMode QPlainTextEdit::lineWrapMode();
extern void C_ZNK14QPlainTextEdit12lineWrapModeEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::clear();
extern void C_ZN14QPlainTextEdit5clearEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::centerCursor();
extern void C_ZN14QPlainTextEdit12centerCursorEv(void* qthis); // 4
  // proto:  void QPlainTextEdit::setReadOnly(bool ro);
extern void C_ZN14QPlainTextEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPlainTextDocumentLayout)=1
type QPlainTextDocumentLayout struct {
  /*qbase*/ qtgui.QAbstractTextDocumentLayout;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPlainTextEdit)=1
type QPlainTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
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
func (this *QPlainTextDocumentLayout) Ensureblocklayout(args ...interface{}) () {
  // ensureBlockLayout(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock
    // invoke: void ensureBlockLayout(const class QTextBlock &)
    var arg0 = args[0].(*qtgui.QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "ensureBlockLayout", args)
  }

  return
}

// pageCount()
func (this *QPlainTextDocumentLayout) Pagecount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QPlainTextDocumentLayout9pageCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "pageCount", args)
  }

  return
}

// blockBoundingRect(const class QTextBlock &)
func (this *QPlainTextDocumentLayout) Blockboundingrect(args ...interface{}) (ret interface{}) {
  // blockBoundingRect(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock
    // invoke: QRectF blockBoundingRect(const class QTextBlock &)
    var arg0 = args[0].(*qtgui.QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "blockBoundingRect", args)
  }

  return
}

// frameBoundingRect(class QTextFrame *)
func (this *QPlainTextDocumentLayout) Frameboundingrect(args ...interface{}) (ret interface{}) {
  // frameBoundingRect(class QTextFrame *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextFrame{}) // "QTextFrame *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame
    // invoke: QRectF frameBoundingRect(class QTextFrame *)
    var arg0 = args[0].(*qtgui.QTextFrame).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "frameBoundingRect", args)
  }

  return
}

// cursorWidth()
func (this *QPlainTextDocumentLayout) Cursorwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QPlainTextDocumentLayout11cursorWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "cursorWidth", args)
  }

  return
}

// requestUpdate()
func (this *QPlainTextDocumentLayout) Requestupdate(args ...interface{}) () {
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
    C.C_ZN24QPlainTextDocumentLayout13requestUpdateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "requestUpdate", args)
  }

  return
}

// setCursorWidth(int)
func (this *QPlainTextDocumentLayout) Setcursorwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN24QPlainTextDocumentLayout14setCursorWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "setCursorWidth", args)
  }

  return
}

// documentSize()
func (this *QPlainTextDocumentLayout) Documentsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QPlainTextDocumentLayout12documentSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "documentSize", args)
  }

  return
}

// ~QPlainTextDocumentLayout()
func (this *QPlainTextDocumentLayout) Freeqplaintextdocumentlayout(args ...interface{}) () {
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
    C.C_ZN24QPlainTextDocumentLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "~QPlainTextDocumentLayout", args)
  }

  return
}

// metaObject()
func (this *QPlainTextDocumentLayout) Metaobject(args ...interface{}) () {
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
    C.C_ZNK24QPlainTextDocumentLayout10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "metaObject", args)
  }

  return
}

// QPlainTextDocumentLayout(class QTextDocument *)
func NewQPlainTextDocumentLayout(args ...interface{}) *QPlainTextDocumentLayout {
  // QPlainTextDocumentLayout(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayoutC1EP13QTextDocument
    // invoke: void QPlainTextDocumentLayout(class QTextDocument *)
    var arg0 = args[0].(*qtgui.QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument(arg0)
    return &QPlainTextDocumentLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "QPlainTextDocumentLayout", args)
  }

  return nil // QPlainTextDocumentLayout{Qclsinst:qthis}
}

// appendPlainText(const class QString &)
func (this *QPlainTextEdit) Appendplaintext(args ...interface{}) () {
  // appendPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15appendPlainTextERK7QString
    // invoke: void appendPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit15appendPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendPlainText", args)
  }

  return
}

// documentTitle()
func (this *QPlainTextEdit) Documenttitle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit13documentTitleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "documentTitle", args)
  }

  return
}

// setTabStopWidth(int)
func (this *QPlainTextEdit) Settabstopwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit15setTabStopWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabStopWidth", args)
  }

  return
}

// ensureCursorVisible()
func (this *QPlainTextEdit) Ensurecursorvisible(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit19ensureCursorVisibleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "ensureCursorVisible", args)
  }

  return
}

// setCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) Setcurrentcharformat(args ...interface{}) () {
  // setCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat
    // invoke: void setCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(*qtgui.QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCurrentCharFormat", args)
  }

  return
}

// tabStopWidth()
func (this *QPlainTextEdit) Tabstopwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit12tabStopWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabStopWidth", args)
  }

  return
}

// isUndoRedoEnabled()
func (this *QPlainTextEdit) Isundoredoenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit17isUndoRedoEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isUndoRedoEnabled", args)
  }

  return
}

// cut()
func (this *QPlainTextEdit) Cut(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit3cutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cut", args)
  }

  return
}

// createStandardContextMenu()
func (this *QPlainTextEdit) Createstandardcontextmenu(args ...interface{}) (ret interface{}) {
  // createStandardContextMenu()
  // createStandardContextMenu(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    var ret0 = C.C_ZN14QPlainTextEdit25createStandardContextMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint
    // invoke: QMenu * createStandardContextMenu(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "createStandardContextMenu", args)
  }

  return
}

// extraSelections()
func (this *QPlainTextEdit) Extraselections(args ...interface{}) () {
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
    C.C_ZNK14QPlainTextEdit15extraSelectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "extraSelections", args)
  }

  return
}

// mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) Mergecurrentcharformat(args ...interface{}) () {
  // mergeCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat
    // invoke: void mergeCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(*qtgui.QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "mergeCurrentCharFormat", args)
  }

  return
}

// metaObject()
func (this *QPlainTextEdit) Metaobject(args ...interface{}) () {
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
    C.C_ZNK14QPlainTextEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "metaObject", args)
  }

  return
}

// setTabChangesFocus(_Bool)
func (this *QPlainTextEdit) Settabchangesfocus(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit18setTabChangesFocusEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabChangesFocus", args)
  }

  return
}

// centerOnScroll()
func (this *QPlainTextEdit) Centeronscroll(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit14centerOnScrollEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerOnScroll", args)
  }

  return
}

// toPlainText()
func (this *QPlainTextEdit) Toplaintext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit11toPlainTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "toPlainText", args)
  }

  return
}

// zoomIn(int)
func (this *QPlainTextEdit) Zoomin(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit6zoomInEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomIn", args)
  }

  return
}

// backgroundVisible()
func (this *QPlainTextEdit) Backgroundvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit17backgroundVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "backgroundVisible", args)
  }

  return
}

// zoomOut(int)
func (this *QPlainTextEdit) Zoomout(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit7zoomOutEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomOut", args)
  }

  return
}

// cursorRect(const class QTextCursor &)
func (this *QPlainTextEdit) Cursorrect(args ...interface{}) (ret interface{}) {
  // cursorRect(const class QTextCursor &)
  // cursorRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCursor{}) // "const QTextCursor &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor
    // invoke: QRect cursorRect(const class QTextCursor &)
    var arg0 = args[0].(*qtgui.QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK14QPlainTextEdit10cursorRectEv
    // invoke: QRect cursorRect()
    var ret0 = C.C_ZNK14QPlainTextEdit10cursorRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorRect", args)
  }

  return
}

// setUndoRedoEnabled(_Bool)
func (this *QPlainTextEdit) Setundoredoenabled(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit18setUndoRedoEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setUndoRedoEnabled", args)
  }

  return
}

// textCursor()
func (this *QPlainTextEdit) Textcursor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit10textCursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textCursor", args)
  }

  return
}

// anchorAt(const class QPoint &)
func (this *QPlainTextEdit) Anchorat(args ...interface{}) (ret interface{}) {
  // anchorAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8anchorAtERK6QPoint
    // invoke: QString anchorAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QPlainTextEdit8anchorAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "anchorAt", args)
  }

  return
}

// setDocument(class QTextDocument *)
func (this *QPlainTextEdit) Setdocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(*qtgui.QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit11setDocumentEP13QTextDocument(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocument", args)
  }

  return
}

// print(class QPagedPaintDevice *)
func (this *QPlainTextEdit) Print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit5printEP17QPagedPaintDevice
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(*qtgui.QPagedPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "print", args)
  }

  return
}

// loadResource(int, const class QUrl &)
func (this *QPlainTextEdit) Loadresource(args ...interface{}) (ret interface{}) {
  // loadResource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12loadResourceEiRK4QUrl
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QPlainTextEdit12loadResourceEiRK4QUrl(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "loadResource", args)
  }

  return
}

// setCenterOnScroll(_Bool)
func (this *QPlainTextEdit) Setcenteronscroll(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit17setCenterOnScrollEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCenterOnScroll", args)
  }

  return
}

// undo()
func (this *QPlainTextEdit) Undo(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit4undoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "undo", args)
  }

  return
}

// paste()
func (this *QPlainTextEdit) Paste(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit5pasteEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "paste", args)
  }

  return
}

// blockCount()
func (this *QPlainTextEdit) Blockcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit10blockCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "blockCount", args)
  }

  return
}

// placeholderText()
func (this *QPlainTextEdit) Placeholdertext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit15placeholderTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "placeholderText", args)
  }

  return
}

// selectAll()
func (this *QPlainTextEdit) Selectall(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "selectAll", args)
  }

  return
}

// maximumBlockCount()
func (this *QPlainTextEdit) Maximumblockcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit17maximumBlockCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "maximumBlockCount", args)
  }

  return
}

// redo()
func (this *QPlainTextEdit) Redo(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "redo", args)
  }

  return
}

// setBackgroundVisible(_Bool)
func (this *QPlainTextEdit) Setbackgroundvisible(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit20setBackgroundVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setBackgroundVisible", args)
  }

  return
}

// isReadOnly()
func (this *QPlainTextEdit) Isreadonly(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isReadOnly", args)
  }

  return
}

// setTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) Settextcursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit13setTextCursorERK11QTextCursor
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(*qtgui.QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTextCursor", args)
  }

  return
}

// overwriteMode()
func (this *QPlainTextEdit) Overwritemode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit13overwriteModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "overwriteMode", args)
  }

  return
}

// setCursorWidth(int)
func (this *QPlainTextEdit) Setcursorwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit14setCursorWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCursorWidth", args)
  }

  return
}

// appendHtml(const class QString &)
func (this *QPlainTextEdit) Appendhtml(args ...interface{}) () {
  // appendHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit10appendHtmlERK7QString
    // invoke: void appendHtml(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit10appendHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendHtml", args)
  }

  return
}

// textInteractionFlags()
func (this *QPlainTextEdit) Textinteractionflags(args ...interface{}) () {
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
    C.C_ZNK14QPlainTextEdit20textInteractionFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textInteractionFlags", args)
  }

  return
}

// canPaste()
func (this *QPlainTextEdit) Canpaste(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit8canPasteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "canPaste", args)
  }

  return
}

// setOverwriteMode(_Bool)
func (this *QPlainTextEdit) Setoverwritemode(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit16setOverwriteModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setOverwriteMode", args)
  }

  return
}

// copy()
func (this *QPlainTextEdit) Copy(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit4copyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "copy", args)
  }

  return
}

// ~QPlainTextEdit()
func (this *QPlainTextEdit) Freeqplaintextedit(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "~QPlainTextEdit", args)
  }

  return
}

// setPlaceholderText(const class QString &)
func (this *QPlainTextEdit) Setplaceholdertext(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit18setPlaceholderTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlaceholderText", args)
  }

  return
}

// setMaximumBlockCount(int)
func (this *QPlainTextEdit) Setmaximumblockcount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit20setMaximumBlockCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setMaximumBlockCount", args)
  }

  return
}

// QPlainTextEdit(const class QString &, class QWidget *)
func NewQPlainTextEdit(args ...interface{}) *QPlainTextEdit {
  // QPlainTextEdit(const class QString &, class QWidget *)
  // QPlainTextEdit(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEditC1ERK7QStringP7QWidget
    // invoke: void QPlainTextEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QPlainTextEditC2ERK7QStringP7QWidget(arg0, arg1)
    return &QPlainTextEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QPlainTextEditC1EP7QWidget
    // invoke: void QPlainTextEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QPlainTextEditC2EP7QWidget(arg0)
    return &QPlainTextEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "QPlainTextEdit", args)
  }

  return nil // QPlainTextEdit{Qclsinst:qthis}
}

// setPlainText(const class QString &)
func (this *QPlainTextEdit) Setplaintext(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit12setPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlainText", args)
  }

  return
}

// insertPlainText(const class QString &)
func (this *QPlainTextEdit) Insertplaintext(args ...interface{}) () {
  // insertPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15insertPlainTextERK7QString
    // invoke: void insertPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit15insertPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "insertPlainText", args)
  }

  return
}

// cursorWidth()
func (this *QPlainTextEdit) Cursorwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit11cursorWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorWidth", args)
  }

  return
}

// tabChangesFocus()
func (this *QPlainTextEdit) Tabchangesfocus(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit15tabChangesFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabChangesFocus", args)
  }

  return
}

// wordWrapMode()
func (this *QPlainTextEdit) Wordwrapmode(args ...interface{}) () {
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
    C.C_ZNK14QPlainTextEdit12wordWrapModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "wordWrapMode", args)
  }

  return
}

// cursorForPosition(const class QPoint &)
func (this *QPlainTextEdit) Cursorforposition(args ...interface{}) (ret interface{}) {
  // cursorForPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17cursorForPositionERK6QPoint
    // invoke: QTextCursor cursorForPosition(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorForPosition", args)
  }

  return
}

// setDocumentTitle(const class QString &)
func (this *QPlainTextEdit) Setdocumenttitle(args ...interface{}) () {
  // setDocumentTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit16setDocumentTitleERK7QString
    // invoke: void setDocumentTitle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QPlainTextEdit16setDocumentTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocumentTitle", args)
  }

  return
}

// document()
func (this *QPlainTextEdit) Document(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "document", args)
  }

  return
}

// currentCharFormat()
func (this *QPlainTextEdit) Currentcharformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QPlainTextEdit17currentCharFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "currentCharFormat", args)
  }

  return
}

// lineWrapMode()
func (this *QPlainTextEdit) Linewrapmode(args ...interface{}) () {
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
    C.C_ZNK14QPlainTextEdit12lineWrapModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "lineWrapMode", args)
  }

  return
}

// clear()
func (this *QPlainTextEdit) Clear(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "clear", args)
  }

  return
}

// centerCursor()
func (this *QPlainTextEdit) Centercursor(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit12centerCursorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerCursor", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QPlainTextEdit) Setreadonly(args ...interface{}) () {
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
    C.C_ZN14QPlainTextEdit11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setReadOnly", args)
  }

  return
}

// <= body block end

