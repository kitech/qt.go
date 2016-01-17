package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtWidgets/qtextedit.h
// dst-file: /src/widgets/qtextedit.go
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
  // proto:  AutoFormatting QTextEdit::autoFormatting();
extern void _ZNK9QTextEdit14autoFormattingEv(void* qthis); // 4
  // proto:  void QTextEdit::setTabStopWidth(int width);
extern void _ZN9QTextEdit15setTabStopWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QColor QTextEdit::textBackgroundColor();
extern void _ZNK9QTextEdit19textBackgroundColorEv(void* qthis); // 4
  // proto:  void QTextEdit::ensureCursorVisible();
extern void _ZN9QTextEdit19ensureCursorVisibleEv(void* qthis); // 4
  // proto:  void QTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void _ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::tabChangesFocus();
extern void _ZNK9QTextEdit15tabChangesFocusEv(void* qthis); // 4
  // proto:  Qt::Alignment QTextEdit::alignment();
extern void _ZNK9QTextEdit9alignmentEv(void* qthis); // 4
  // proto:  bool QTextEdit::isUndoRedoEnabled();
extern void _ZNK9QTextEdit17isUndoRedoEnabledEv(void* qthis); // 2
  // proto:  void QTextEdit::cut();
extern void _ZN9QTextEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QTextEdit::createStandardContextMenu(const QPoint & position);
extern void _ZN9QTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QMenu * QTextEdit::createStandardContextMenu();
extern void _ZN9QTextEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontPointSize(qreal s);
extern void _ZN9QTextEdit16setFontPointSizeEd(void* qthis, double arg0); // 4
  // proto:  void QTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void _ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  QString QTextEdit::documentTitle();
extern void _ZNK9QTextEdit13documentTitleEv(void* qthis); // 2
  // proto:  QString QTextEdit::toHtml();
extern void _ZNK9QTextEdit6toHtmlEv(void* qthis); // 4
  // proto:  void QTextEdit::insertPlainText(const QString & text);
extern void _ZN9QTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  qreal QTextEdit::fontPointSize();
extern void _ZNK9QTextEdit13fontPointSizeEv(void* qthis); // 4
  // proto:  QColor QTextEdit::textColor();
extern void _ZNK9QTextEdit9textColorEv(void* qthis); // 4
  // proto:  const QMetaObject * QTextEdit::metaObject();
extern void _ZNK9QTextEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextBackgroundColor(const QColor & c);
extern void _ZN9QTextEdit22setTextBackgroundColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setTabChangesFocus(bool b);
extern void _ZN9QTextEdit18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextEdit::toPlainText();
extern void _ZNK9QTextEdit11toPlainTextEv(void* qthis); // 4
  // proto:  void QTextEdit::zoomIn(int range);
extern void _ZN9QTextEdit6zoomInEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::zoomOut(int range);
extern void _ZN9QTextEdit7zoomOutEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QTextEdit::cursorRect(const QTextCursor & cursor);
extern void _ZNK9QTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QTextEdit::cursorRect();
extern void _ZNK9QTextEdit10cursorRectEv(void* qthis); // 4
  // proto:  void QTextEdit::setUndoRedoEnabled(bool enable);
extern void _ZN9QTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0); // 2
  // proto:  QTextCursor QTextEdit::textCursor();
extern void _ZNK9QTextEdit10textCursorEv(void* qthis); // 4
  // proto:  void QTextEdit::setAcceptRichText(bool accept);
extern void _ZN9QTextEdit17setAcceptRichTextEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextEdit::anchorAt(const QPoint & pos);
extern void _ZNK9QTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setHtml(const QString & text);
extern void _ZN9QTextEdit7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setDocument(QTextDocument * document);
extern void _ZN9QTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::fontItalic();
extern void _ZNK9QTextEdit10fontItalicEv(void* qthis); // 4
  // proto:  void QTextEdit::append(const QString & text);
extern void _ZN9QTextEdit6appendERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setFontFamily(const QString & fontFamily);
extern void _ZN9QTextEdit13setFontFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::print(QPagedPaintDevice * printer);
extern void _ZNK9QTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QVariant QTextEdit::loadResource(int type, const QUrl & name);
extern void _ZN9QTextEdit12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTextEdit::undo();
extern void _ZN9QTextEdit4undoEv(void* qthis); // 4
  // proto:  int QTextEdit::lineWrapColumnOrWidth();
extern void _ZNK9QTextEdit21lineWrapColumnOrWidthEv(void* qthis); // 4
  // proto:  void QTextEdit::paste();
extern void _ZN9QTextEdit5pasteEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontWeight(int w);
extern void _ZN9QTextEdit13setFontWeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::scrollToAnchor(const QString & name);
extern void _ZN9QTextEdit14scrollToAnchorERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::selectAll();
extern void _ZN9QTextEdit9selectAllEv(void* qthis); // 4
  // proto:  QList<QTextEdit::ExtraSelection> QTextEdit::extraSelections();
extern void _ZNK9QTextEdit15extraSelectionsEv(void* qthis); // 4
  // proto:  void QTextEdit::~QTextEdit();
extern void _ZN9QTextEditD2Ev(void* qthis); // 4
  // proto:  void QTextEdit::redo();
extern void _ZN9QTextEdit4redoEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontUnderline(bool b);
extern void _ZN9QTextEdit16setFontUnderlineEb(void* qthis, bool arg0); // 4
  // proto:  bool QTextEdit::isReadOnly();
extern void _ZNK9QTextEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextCursor(const QTextCursor & cursor);
extern void _ZN9QTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::overwriteMode();
extern void _ZNK9QTextEdit13overwriteModeEv(void* qthis); // 4
  // proto:  bool QTextEdit::acceptRichText();
extern void _ZNK9QTextEdit14acceptRichTextEv(void* qthis); // 4
  // proto:  void QTextEdit::setCursorWidth(int width);
extern void _ZN9QTextEdit14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::setLineWrapColumnOrWidth(int w);
extern void _ZN9QTextEdit24setLineWrapColumnOrWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::TextInteractionFlags QTextEdit::textInteractionFlags();
extern void _ZNK9QTextEdit20textInteractionFlagsEv(void* qthis); // 4
  // proto:  QFont QTextEdit::currentFont();
extern void _ZNK9QTextEdit11currentFontEv(void* qthis); // 4
  // proto:  void QTextEdit::setOverwriteMode(bool overwrite);
extern void _ZN9QTextEdit16setOverwriteModeEb(void* qthis, bool arg0); // 4
  // proto:  void QTextEdit::copy();
extern void _ZN9QTextEdit4copyEv(void* qthis); // 4
  // proto:  bool QTextEdit::fontUnderline();
extern void _ZNK9QTextEdit13fontUnderlineEv(void* qthis); // 4
  // proto:  int QTextEdit::cursorWidth();
extern void _ZNK9QTextEdit11cursorWidthEv(void* qthis); // 4
  // proto:  void QTextEdit::setPlaceholderText(const QString & placeholderText);
extern void _ZN9QTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::clear();
extern void _ZN9QTextEdit5clearEv(void* qthis); // 4
  // proto:  void QTextEdit::insertHtml(const QString & text);
extern void _ZN9QTextEdit10insertHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setPlainText(const QString & text);
extern void _ZN9QTextEdit12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::canPaste();
extern void _ZNK9QTextEdit8canPasteEv(void* qthis); // 4
  // proto:  QString QTextEdit::fontFamily();
extern void _ZNK9QTextEdit10fontFamilyEv(void* qthis); // 4
  // proto:  void QTextEdit::QTextEdit(QWidget * parent);
extern void _ZN9QTextEditC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QTextEdit::QTextEdit(const QString & text, QWidget * parent);
extern void _ZN9QTextEditC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto:  int QTextEdit::tabStopWidth();
extern void _ZNK9QTextEdit12tabStopWidthEv(void* qthis); // 4
  // proto:  QTextOption::WrapMode QTextEdit::wordWrapMode();
extern void _ZNK9QTextEdit12wordWrapModeEv(void* qthis); // 4
  // proto:  QTextCursor QTextEdit::cursorForPosition(const QPoint & pos);
extern void _ZNK9QTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setDocumentTitle(const QString & title);
extern void _ZN9QTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextEdit::setFontItalic(bool b);
extern void _ZN9QTextEdit13setFontItalicEb(void* qthis, bool arg0); // 4
  // proto:  QTextDocument * QTextEdit::document();
extern void _ZNK9QTextEdit8documentEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextEdit::currentCharFormat();
extern void _ZNK9QTextEdit17currentCharFormatEv(void* qthis); // 4
  // proto:  int QTextEdit::fontWeight();
extern void _ZNK9QTextEdit10fontWeightEv(void* qthis); // 4
  // proto:  QString QTextEdit::placeholderText();
extern void _ZNK9QTextEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextColor(const QColor & c);
extern void _ZN9QTextEdit12setTextColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QTextEdit::LineWrapMode QTextEdit::lineWrapMode();
extern void _ZNK9QTextEdit12lineWrapModeEv(void* qthis); // 4
  // proto:  void QTextEdit::setText(const QString & text);
extern void _ZN9QTextEdit7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setCurrentFont(const QFont & f);
extern void _ZN9QTextEdit14setCurrentFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setReadOnly(bool ro);
extern void _ZN9QTextEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
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

// class sizeof(QTextEdit)=1
type QTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst unsafe.Pointer /* *C.void */;
//  _cursorPositionChanged QTextEdit_cursorPositionChanged_signal;
//  _redoAvailable QTextEdit_redoAvailable_signal;
//  _selectionChanged QTextEdit_selectionChanged_signal;
//  _currentCharFormatChanged QTextEdit_currentCharFormatChanged_signal;
//  _undoAvailable QTextEdit_undoAvailable_signal;
//  _textChanged QTextEdit_textChanged_signal;
//  _copyAvailable QTextEdit_copyAvailable_signal;
}

// autoFormatting()
func (this *QTextEdit) autoFormatting(args ...interface{}) () {
  // autoFormatting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit14autoFormattingEv
    // invoke: AutoFormatting autoFormatting()
    C._ZNK9QTextEdit14autoFormattingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "autoFormatting", args)
  }

}

// setTabStopWidth(int)
func (this *QTextEdit) setTabStopWidth(args ...interface{}) () {
  // setTabStopWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit15setTabStopWidthEi
    // invoke: void setTabStopWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit15setTabStopWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabStopWidth", args)
  }

}

// textBackgroundColor()
func (this *QTextEdit) textBackgroundColor(args ...interface{}) () {
  // textBackgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit19textBackgroundColorEv
    // invoke: QColor textBackgroundColor()
    C._ZNK9QTextEdit19textBackgroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "textBackgroundColor", args)
  }

}

// ensureCursorVisible()
func (this *QTextEdit) ensureCursorVisible(args ...interface{}) () {
  // ensureCursorVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit19ensureCursorVisibleEv
    // invoke: void ensureCursorVisible()
    C._ZN9QTextEdit19ensureCursorVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "ensureCursorVisible", args)
  }

}

// setCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) setCurrentCharFormat(args ...interface{}) () {
  // setCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat
    // invoke: void setCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentCharFormat", args)
  }

}

// tabChangesFocus()
func (this *QTextEdit) tabChangesFocus(args ...interface{}) () {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15tabChangesFocusEv
    // invoke: bool tabChangesFocus()
    C._ZNK9QTextEdit15tabChangesFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "tabChangesFocus", args)
  }

}

// alignment()
func (this *QTextEdit) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit9alignmentEv
    // invoke: Qt::Alignment alignment()
    C._ZNK9QTextEdit9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "alignment", args)
  }

}

// isUndoRedoEnabled()
func (this *QTextEdit) isUndoRedoEnabled(args ...interface{}) () {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17isUndoRedoEnabledEv
    // invoke: bool isUndoRedoEnabled()
    C._ZNK9QTextEdit17isUndoRedoEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "isUndoRedoEnabled", args)
  }

}

// cut()
func (this *QTextEdit) cut(args ...interface{}) () {
  // cut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit3cutEv
    // invoke: void cut()
    C._ZN9QTextEdit3cutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "cut", args)
  }

}

// createStandardContextMenu(const class QPoint &)
func (this *QTextEdit) createStandardContextMenu(args ...interface{}) () {
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
    // invoke: _ZN9QTextEdit25createStandardContextMenuERK6QPoint
    // invoke: QMenu * createStandardContextMenu(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit25createStandardContextMenuERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QTextEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    C._ZN9QTextEdit25createStandardContextMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "createStandardContextMenu", args)
  }

}

// setFontPointSize(qreal)
func (this *QTextEdit) setFontPointSize(args ...interface{}) () {
  // setFontPointSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setFontPointSizeEd
    // invoke: void setFontPointSize(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit16setFontPointSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontPointSize", args)
  }

}

// mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) mergeCurrentCharFormat(args ...interface{}) () {
  // mergeCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat
    // invoke: void mergeCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "mergeCurrentCharFormat", args)
  }

}

// documentTitle()
func (this *QTextEdit) documentTitle(args ...interface{}) () {
  // documentTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13documentTitleEv
    // invoke: QString documentTitle()
    C._ZNK9QTextEdit13documentTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "documentTitle", args)
  }

}

// toHtml()
func (this *QTextEdit) toHtml(args ...interface{}) () {
  // toHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit6toHtmlEv
    // invoke: QString toHtml()
    C._ZNK9QTextEdit6toHtmlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "toHtml", args)
  }

}

// insertPlainText(const class QString &)
func (this *QTextEdit) insertPlainText(args ...interface{}) () {
  // insertPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit15insertPlainTextERK7QString
    // invoke: void insertPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit15insertPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "insertPlainText", args)
  }

}

// fontPointSize()
func (this *QTextEdit) fontPointSize(args ...interface{}) () {
  // fontPointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13fontPointSizeEv
    // invoke: qreal fontPointSize()
    C._ZNK9QTextEdit13fontPointSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "fontPointSize", args)
  }

}

// textColor()
func (this *QTextEdit) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit9textColorEv
    // invoke: QColor textColor()
    C._ZNK9QTextEdit9textColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "textColor", args)
  }

}

// metaObject()
func (this *QTextEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QTextEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "metaObject", args)
  }

}

// setTextBackgroundColor(const class QColor &)
func (this *QTextEdit) setTextBackgroundColor(args ...interface{}) () {
  // setTextBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit22setTextBackgroundColorERK6QColor
    // invoke: void setTextBackgroundColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit22setTextBackgroundColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextBackgroundColor", args)
  }

}

// setTabChangesFocus(_Bool)
func (this *QTextEdit) setTabChangesFocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setTabChangesFocusEb
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit18setTabChangesFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabChangesFocus", args)
  }

}

// toPlainText()
func (this *QTextEdit) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11toPlainTextEv
    // invoke: QString toPlainText()
    C._ZNK9QTextEdit11toPlainTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "toPlainText", args)
  }

}

// zoomIn(int)
func (this *QTextEdit) zoomIn(args ...interface{}) () {
  // zoomIn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit6zoomInEi
    // invoke: void zoomIn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit6zoomInEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomIn", args)
  }

}

// zoomOut(int)
func (this *QTextEdit) zoomOut(args ...interface{}) () {
  // zoomOut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7zoomOutEi
    // invoke: void zoomOut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit7zoomOutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomOut", args)
  }

}

// cursorRect(const class QTextCursor &)
func (this *QTextEdit) cursorRect(args ...interface{}) () {
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
    // invoke: _ZNK9QTextEdit10cursorRectERK11QTextCursor
    // invoke: QRect cursorRect(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextEdit10cursorRectERK11QTextCursor(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK9QTextEdit10cursorRectEv
    // invoke: QRect cursorRect()
    C._ZNK9QTextEdit10cursorRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorRect", args)
  }

}

// setUndoRedoEnabled(_Bool)
func (this *QTextEdit) setUndoRedoEnabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setUndoRedoEnabledEb
    // invoke: void setUndoRedoEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit18setUndoRedoEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setUndoRedoEnabled", args)
  }

}

// textCursor()
func (this *QTextEdit) textCursor(args ...interface{}) () {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10textCursorEv
    // invoke: QTextCursor textCursor()
    C._ZNK9QTextEdit10textCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "textCursor", args)
  }

}

// setAcceptRichText(_Bool)
func (this *QTextEdit) setAcceptRichText(args ...interface{}) () {
  // setAcceptRichText(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit17setAcceptRichTextEb
    // invoke: void setAcceptRichText(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit17setAcceptRichTextEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setAcceptRichText", args)
  }

}

// anchorAt(const class QPoint &)
func (this *QTextEdit) anchorAt(args ...interface{}) () {
  // anchorAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8anchorAtERK6QPoint
    // invoke: QString anchorAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextEdit8anchorAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "anchorAt", args)
  }

}

// setHtml(const class QString &)
func (this *QTextEdit) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7setHtmlERK7QString
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit7setHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setHtml", args)
  }

}

// setDocument(class QTextDocument *)
func (this *QTextEdit) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit11setDocumentEP13QTextDocument(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocument", args)
  }

}

// fontItalic()
func (this *QTextEdit) fontItalic(args ...interface{}) () {
  // fontItalic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontItalicEv
    // invoke: bool fontItalic()
    C._ZNK9QTextEdit10fontItalicEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "fontItalic", args)
  }

}

// append(const class QString &)
func (this *QTextEdit) append(args ...interface{}) () {
  // append(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit6appendERK7QString
    // invoke: void append(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit6appendERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "append", args)
  }

}

// setFontFamily(const class QString &)
func (this *QTextEdit) setFontFamily(args ...interface{}) () {
  // setFontFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontFamilyERK7QString
    // invoke: void setFontFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit13setFontFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontFamily", args)
  }

}

// print(class QPagedPaintDevice *)
func (this *QTextEdit) print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit5printEP17QPagedPaintDevice
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextEdit5printEP17QPagedPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "print", args)
  }

}

// loadResource(int, const class QUrl &)
func (this *QTextEdit) loadResource(args ...interface{}) () {
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
    // invoke: _ZN9QTextEdit12loadResourceEiRK4QUrl
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QTextEdit12loadResourceEiRK4QUrl(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextEdit", "loadResource", args)
  }

}

// undo()
func (this *QTextEdit) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4undoEv
    // invoke: void undo()
    C._ZN9QTextEdit4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "undo", args)
  }

}

// lineWrapColumnOrWidth()
func (this *QTextEdit) lineWrapColumnOrWidth(args ...interface{}) () {
  // lineWrapColumnOrWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit21lineWrapColumnOrWidthEv
    // invoke: int lineWrapColumnOrWidth()
    C._ZNK9QTextEdit21lineWrapColumnOrWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "lineWrapColumnOrWidth", args)
  }

}

// paste()
func (this *QTextEdit) paste(args ...interface{}) () {
  // paste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit5pasteEv
    // invoke: void paste()
    C._ZN9QTextEdit5pasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "paste", args)
  }

}

// setFontWeight(int)
func (this *QTextEdit) setFontWeight(args ...interface{}) () {
  // setFontWeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontWeightEi
    // invoke: void setFontWeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit13setFontWeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontWeight", args)
  }

}

// scrollToAnchor(const class QString &)
func (this *QTextEdit) scrollToAnchor(args ...interface{}) () {
  // scrollToAnchor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14scrollToAnchorERK7QString
    // invoke: void scrollToAnchor(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit14scrollToAnchorERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "scrollToAnchor", args)
  }

}

// selectAll()
func (this *QTextEdit) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit9selectAllEv
    // invoke: void selectAll()
    C._ZN9QTextEdit9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "selectAll", args)
  }

}

// extraSelections()
func (this *QTextEdit) extraSelections(args ...interface{}) () {
  // extraSelections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15extraSelectionsEv
    // invoke: QList<QTextEdit::ExtraSelection> extraSelections()
    C._ZNK9QTextEdit15extraSelectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "extraSelections", args)
  }

}

// ~QTextEdit()
func (this *QTextEdit) FreeQTextEdit(args ...interface{}) () {
  // ~QTextEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEditD0Ev
    // invoke: void ~QTextEdit()
    C._ZN9QTextEditD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "~QTextEdit", args)
  }

}

// redo()
func (this *QTextEdit) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4redoEv
    // invoke: void redo()
    C._ZN9QTextEdit4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "redo", args)
  }

}

// setFontUnderline(_Bool)
func (this *QTextEdit) setFontUnderline(args ...interface{}) () {
  // setFontUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setFontUnderlineEb
    // invoke: void setFontUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit16setFontUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontUnderline", args)
  }

}

// isReadOnly()
func (this *QTextEdit) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10isReadOnlyEv
    // invoke: bool isReadOnly()
    C._ZNK9QTextEdit10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "isReadOnly", args)
  }

}

// setTextCursor(const class QTextCursor &)
func (this *QTextEdit) setTextCursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setTextCursorERK11QTextCursor
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit13setTextCursorERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextCursor", args)
  }

}

// overwriteMode()
func (this *QTextEdit) overwriteMode(args ...interface{}) () {
  // overwriteMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13overwriteModeEv
    // invoke: bool overwriteMode()
    C._ZNK9QTextEdit13overwriteModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "overwriteMode", args)
  }

}

// acceptRichText()
func (this *QTextEdit) acceptRichText(args ...interface{}) () {
  // acceptRichText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit14acceptRichTextEv
    // invoke: bool acceptRichText()
    C._ZNK9QTextEdit14acceptRichTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "acceptRichText", args)
  }

}

// setCursorWidth(int)
func (this *QTextEdit) setCursorWidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14setCursorWidthEi
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit14setCursorWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCursorWidth", args)
  }

}

// setLineWrapColumnOrWidth(int)
func (this *QTextEdit) setLineWrapColumnOrWidth(args ...interface{}) () {
  // setLineWrapColumnOrWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit24setLineWrapColumnOrWidthEi
    // invoke: void setLineWrapColumnOrWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit24setLineWrapColumnOrWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setLineWrapColumnOrWidth", args)
  }

}

// textInteractionFlags()
func (this *QTextEdit) textInteractionFlags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C._ZNK9QTextEdit20textInteractionFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "textInteractionFlags", args)
  }

}

// currentFont()
func (this *QTextEdit) currentFont(args ...interface{}) () {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11currentFontEv
    // invoke: QFont currentFont()
    C._ZNK9QTextEdit11currentFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "currentFont", args)
  }

}

// setOverwriteMode(_Bool)
func (this *QTextEdit) setOverwriteMode(args ...interface{}) () {
  // setOverwriteMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setOverwriteModeEb
    // invoke: void setOverwriteMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit16setOverwriteModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setOverwriteMode", args)
  }

}

// copy()
func (this *QTextEdit) copy(args ...interface{}) () {
  // copy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4copyEv
    // invoke: void copy()
    C._ZN9QTextEdit4copyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "copy", args)
  }

}

// fontUnderline()
func (this *QTextEdit) fontUnderline(args ...interface{}) () {
  // fontUnderline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13fontUnderlineEv
    // invoke: bool fontUnderline()
    C._ZNK9QTextEdit13fontUnderlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "fontUnderline", args)
  }

}

// cursorWidth()
func (this *QTextEdit) cursorWidth(args ...interface{}) () {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11cursorWidthEv
    // invoke: int cursorWidth()
    C._ZNK9QTextEdit11cursorWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorWidth", args)
  }

}

// setPlaceholderText(const class QString &)
func (this *QTextEdit) setPlaceholderText(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit18setPlaceholderTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlaceholderText", args)
  }

}

// clear()
func (this *QTextEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit5clearEv
    // invoke: void clear()
    C._ZN9QTextEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "clear", args)
  }

}

// insertHtml(const class QString &)
func (this *QTextEdit) insertHtml(args ...interface{}) () {
  // insertHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit10insertHtmlERK7QString
    // invoke: void insertHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit10insertHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "insertHtml", args)
  }

}

// setPlainText(const class QString &)
func (this *QTextEdit) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit12setPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlainText", args)
  }

}

// canPaste()
func (this *QTextEdit) canPaste(args ...interface{}) () {
  // canPaste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8canPasteEv
    // invoke: bool canPaste()
    C._ZNK9QTextEdit8canPasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "canPaste", args)
  }

}

// fontFamily()
func (this *QTextEdit) fontFamily(args ...interface{}) () {
  // fontFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontFamilyEv
    // invoke: QString fontFamily()
    C._ZNK9QTextEdit10fontFamilyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "fontFamily", args)
  }

}

// QTextEdit(class QWidget *)
func NewQTextEdit(args ...interface{}) QTextEdit {
  // QTextEdit(class QWidget *)
  // QTextEdit(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEditC1EP7QWidget
    // invoke: void QTextEdit(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QTextEditC2EP7QWidget(qthis, arg0)
  case 1:
    // invoke: _ZN9QTextEditC1ERK7QStringP7QWidget
    // invoke: void QTextEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QTextEditC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextEdit", "QTextEdit", args)
  }

  return QTextEdit{}
}

// tabStopWidth()
func (this *QTextEdit) tabStopWidth(args ...interface{}) () {
  // tabStopWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12tabStopWidthEv
    // invoke: int tabStopWidth()
    C._ZNK9QTextEdit12tabStopWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "tabStopWidth", args)
  }

}

// wordWrapMode()
func (this *QTextEdit) wordWrapMode(args ...interface{}) () {
  // wordWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12wordWrapModeEv
    // invoke: QTextOption::WrapMode wordWrapMode()
    C._ZNK9QTextEdit12wordWrapModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "wordWrapMode", args)
  }

}

// cursorForPosition(const class QPoint &)
func (this *QTextEdit) cursorForPosition(args ...interface{}) () {
  // cursorForPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17cursorForPositionERK6QPoint
    // invoke: QTextCursor cursorForPosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextEdit17cursorForPositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorForPosition", args)
  }

}

// setDocumentTitle(const class QString &)
func (this *QTextEdit) setDocumentTitle(args ...interface{}) () {
  // setDocumentTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setDocumentTitleERK7QString
    // invoke: void setDocumentTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit16setDocumentTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocumentTitle", args)
  }

}

// setFontItalic(_Bool)
func (this *QTextEdit) setFontItalic(args ...interface{}) () {
  // setFontItalic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontItalicEb
    // invoke: void setFontItalic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit13setFontItalicEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontItalic", args)
  }

}

// document()
func (this *QTextEdit) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8documentEv
    // invoke: QTextDocument * document()
    C._ZNK9QTextEdit8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "document", args)
  }

}

// currentCharFormat()
func (this *QTextEdit) currentCharFormat(args ...interface{}) () {
  // currentCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17currentCharFormatEv
    // invoke: QTextCharFormat currentCharFormat()
    C._ZNK9QTextEdit17currentCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "currentCharFormat", args)
  }

}

// fontWeight()
func (this *QTextEdit) fontWeight(args ...interface{}) () {
  // fontWeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontWeightEv
    // invoke: int fontWeight()
    C._ZNK9QTextEdit10fontWeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "fontWeight", args)
  }

}

// placeholderText()
func (this *QTextEdit) placeholderText(args ...interface{}) () {
  // placeholderText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15placeholderTextEv
    // invoke: QString placeholderText()
    C._ZNK9QTextEdit15placeholderTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "placeholderText", args)
  }

}

// setTextColor(const class QColor &)
func (this *QTextEdit) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit12setTextColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextColor", args)
  }

}

// lineWrapMode()
func (this *QTextEdit) lineWrapMode(args ...interface{}) () {
  // lineWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12lineWrapModeEv
    // invoke: QTextEdit::LineWrapMode lineWrapMode()
    C._ZNK9QTextEdit12lineWrapModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "lineWrapMode", args)
  }

}

// setText(const class QString &)
func (this *QTextEdit) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setText", args)
  }

}

// setCurrentFont(const class QFont &)
func (this *QTextEdit) setCurrentFont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit14setCurrentFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentFont", args)
  }

}

// setReadOnly(_Bool)
func (this *QTextEdit) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QTextEdit11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setReadOnly", args)
  }

}

// <= body block end

