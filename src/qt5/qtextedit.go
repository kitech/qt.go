package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  int QTextEdit::lineWrapColumnOrWidth();
extern void _ZNK9QTextEdit21lineWrapColumnOrWidthEv(void* qthis);
  // proto:  void QTextEdit::setFontFamily(const QString & fontFamily);
extern void _ZN9QTextEdit13setFontFamilyERK7QString(void* qthis, void* arg0);
  // proto:  QString QTextEdit::toPlainText();
extern void _ZNK9QTextEdit11toPlainTextEv(void* qthis);
  // proto:  void QTextEdit::setCursorWidth(int width);
extern void _ZN9QTextEdit14setCursorWidthEi(void* qthis, int arg0);
  // proto:  QMenu * QTextEdit::createStandardContextMenu();
extern void _ZN9QTextEdit25createStandardContextMenuEv(void* qthis);
  // proto:  QTextDocument * QTextEdit::document();
extern void _ZNK9QTextEdit8documentEv(void* qthis);
  // proto:  QRect QTextEdit::cursorRect();
extern void _ZNK9QTextEdit10cursorRectEv(void* qthis);
  // proto:  void QTextEdit::setTextColor(const QColor & c);
extern void _ZN9QTextEdit12setTextColorERK6QColor(void* qthis, void* arg0);
  // proto:  bool QTextEdit::acceptRichText();
extern void _ZNK9QTextEdit14acceptRichTextEv(void* qthis);
  // proto:  void QTextEdit::clear();
extern void _ZN9QTextEdit5clearEv(void* qthis);
  // proto:  void QTextEdit::insertHtml(const QString & text);
extern void _ZN9QTextEdit10insertHtmlERK7QString(void* qthis, void* arg0);
  // proto:  QString QTextEdit::fontFamily();
extern void _ZNK9QTextEdit10fontFamilyEv(void* qthis);
  // proto:  void QTextEdit::setFontUnderline(bool b);
extern void _ZN9QTextEdit16setFontUnderlineEb(void* qthis, bool arg0);
  // proto:  void QTextEdit::cut();
extern void _ZN9QTextEdit3cutEv(void* qthis);
  // proto:  QString QTextEdit::anchorAt(const QPoint & pos);
extern void _ZNK9QTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0);
  // proto:  int QTextEdit::cursorWidth();
extern void _ZNK9QTextEdit11cursorWidthEv(void* qthis);
  // proto:  void QTextEdit::setTextBackgroundColor(const QColor & c);
extern void _ZN9QTextEdit22setTextBackgroundColorERK6QColor(void* qthis, void* arg0);
  // proto:  int QTextEdit::tabStopWidth();
extern void _ZNK9QTextEdit12tabStopWidthEv(void* qthis);
  // proto:  void QTextEdit::setFontWeight(int w);
extern void _ZN9QTextEdit13setFontWeightEi(void* qthis, int arg0);
  // proto:  void QTextEdit::selectAll();
extern void _ZN9QTextEdit9selectAllEv(void* qthis);
  // proto:  void QTextEdit::zoomOut(int range);
extern void _ZN9QTextEdit7zoomOutEi(void* qthis, int arg0);
  // proto:  void QTextEdit::redo();
extern void _ZN9QTextEdit4redoEv(void* qthis);
  // proto:  void QTextEdit::setFontPointSize(qreal s);
extern void _ZN9QTextEdit16setFontPointSizeEd(void* qthis, double arg0);
  // proto:  bool QTextEdit::overwriteMode();
extern void _ZNK9QTextEdit13overwriteModeEv(void* qthis);
  // proto:  QTextCursor QTextEdit::textCursor();
extern void _ZNK9QTextEdit10textCursorEv(void* qthis);
  // proto:  void QTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void _ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  void QTextEdit::setPlainText(const QString & text);
extern void _ZN9QTextEdit12setPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  QString QTextEdit::placeholderText();
extern void _ZNK9QTextEdit15placeholderTextEv(void* qthis);
  // proto:  void QTextEdit::~QTextEdit();
extern void _ZN9QTextEditD0Ev(void* qthis);
  // proto:  bool QTextEdit::fontItalic();
extern void _ZNK9QTextEdit10fontItalicEv(void* qthis);
  // proto:  void QTextEdit::copy();
extern void _ZN9QTextEdit4copyEv(void* qthis);
  // proto:  qreal QTextEdit::fontPointSize();
extern void _ZNK9QTextEdit13fontPointSizeEv(void* qthis);
  // proto:  void QTextEdit::setDocument(QTextDocument * document);
extern void _ZN9QTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QTextEdit::setOverwriteMode(bool overwrite);
extern void _ZN9QTextEdit16setOverwriteModeEb(void* qthis, bool arg0);
  // proto:  void QTextEdit::undo();
extern void _ZN9QTextEdit4undoEv(void* qthis);
  // proto:  void QTextEdit::zoomIn(int range);
extern void _ZN9QTextEdit6zoomInEi(void* qthis, int arg0);
  // proto:  void QTextEdit::setDocumentTitle(const QString & title);
extern void demth_ZN9QTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextEdit::canPaste();
extern void _ZNK9QTextEdit8canPasteEv(void* qthis);
  // proto:  QString QTextEdit::toHtml();
extern void _ZNK9QTextEdit6toHtmlEv(void* qthis);
  // proto:  QMenu * QTextEdit::createStandardContextMenu(const QPoint & position);
extern void _ZN9QTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTextEdit::setTabStopWidth(int width);
extern void _ZN9QTextEdit15setTabStopWidthEi(void* qthis, int arg0);
  // proto:  QString QTextEdit::documentTitle();
extern void demth_ZNK9QTextEdit13documentTitleEv(void* qthis);
  // proto:  bool QTextEdit::isUndoRedoEnabled();
extern void demth_ZNK9QTextEdit17isUndoRedoEnabledEv(void* qthis);
  // proto:  void QTextEdit::setText(const QString & text);
extern void _ZN9QTextEdit7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QTextEdit::ensureCursorVisible();
extern void _ZN9QTextEdit19ensureCursorVisibleEv(void* qthis);
  // proto:  void QTextEdit::setAcceptRichText(bool accept);
extern void _ZN9QTextEdit17setAcceptRichTextEb(void* qthis, bool arg0);
  // proto:  void QTextEdit::setPlaceholderText(const QString & placeholderText);
extern void _ZN9QTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0);
  // proto:  bool QTextEdit::isReadOnly();
extern void _ZNK9QTextEdit10isReadOnlyEv(void* qthis);
  // proto:  void QTextEdit::setUndoRedoEnabled(bool enable);
extern void demth_ZN9QTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0);
  // proto:  void QTextEdit::QTextEdit(const QTextEdit & );
extern void* dector_ZN9QTextEditC1ERKS_(void* arg0);
extern void _ZN9QTextEditC1ERKS_(void* qthis, void* arg0);
  // proto:  QTextCharFormat QTextEdit::currentCharFormat();
extern void _ZNK9QTextEdit17currentCharFormatEv(void* qthis);
  // proto:  QTextCursor QTextEdit::cursorForPosition(const QPoint & pos);
extern void _ZNK9QTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0);
  // proto:  void QTextEdit::scrollToAnchor(const QString & name);
extern void _ZN9QTextEdit14scrollToAnchorERK7QString(void* qthis, void* arg0);
  // proto:  QFont QTextEdit::currentFont();
extern void _ZNK9QTextEdit11currentFontEv(void* qthis);
  // proto:  void QTextEdit::paste();
extern void _ZN9QTextEdit5pasteEv(void* qthis);
  // proto:  void QTextEdit::setTextCursor(const QTextCursor & cursor);
extern void _ZN9QTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void _ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  QVariant QTextEdit::loadResource(int type, const QUrl & name);
extern void _ZN9QTextEdit12loadResourceEiRK4QUrl(void* qthis, int arg0, void* arg1);
  // proto:  void QTextEdit::setTabChangesFocus(bool b);
extern void _ZN9QTextEdit18setTabChangesFocusEb(void* qthis, bool arg0);
  // proto:  void QTextEdit::setHtml(const QString & text);
extern void _ZN9QTextEdit7setHtmlERK7QString(void* qthis, void* arg0);
  // proto:  QRect QTextEdit::cursorRect(const QTextCursor & cursor);
extern void _ZNK9QTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QTextEdit::setLineWrapColumnOrWidth(int w);
extern void _ZN9QTextEdit24setLineWrapColumnOrWidthEi(void* qthis, int arg0);
  // proto:  void QTextEdit::setFontItalic(bool b);
extern void _ZN9QTextEdit13setFontItalicEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QTextEdit::metaObject();
extern void _ZNK9QTextEdit10metaObjectEv(void* qthis);
  // proto:  void QTextEdit::setCurrentFont(const QFont & f);
extern void _ZN9QTextEdit14setCurrentFontERK5QFont(void* qthis, void* arg0);
  // proto:  bool QTextEdit::tabChangesFocus();
extern void _ZNK9QTextEdit15tabChangesFocusEv(void* qthis);
  // proto:  QColor QTextEdit::textBackgroundColor();
extern void _ZNK9QTextEdit19textBackgroundColorEv(void* qthis);
  // proto:  void QTextEdit::QTextEdit(const QString & text, QWidget * parent);
extern void* dector_ZN9QTextEditC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN9QTextEditC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QTextEdit::print(QPagedPaintDevice * printer);
extern void _ZNK9QTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0);
  // proto:  bool QTextEdit::fontUnderline();
extern void _ZNK9QTextEdit13fontUnderlineEv(void* qthis);
  // proto:  void QTextEdit::insertPlainText(const QString & text);
extern void _ZN9QTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  int QTextEdit::fontWeight();
extern void _ZNK9QTextEdit10fontWeightEv(void* qthis);
  // proto:  QColor QTextEdit::textColor();
extern void _ZNK9QTextEdit9textColorEv(void* qthis);
  // proto:  void QTextEdit::append(const QString & text);
extern void _ZN9QTextEdit6appendERK7QString(void* qthis, void* arg0);
  // proto:  void QTextEdit::QTextEdit(QWidget * parent);
extern void* dector_ZN9QTextEditC1EP7QWidget(void* arg0);
extern void _ZN9QTextEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QTextEdit::setReadOnly(bool ro);
extern void _ZN9QTextEdit11setReadOnlyEb(void* qthis, bool arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _cursorPositionChanged QTextEdit_cursorPositionChanged_signal;
//  _redoAvailable QTextEdit_redoAvailable_signal;
//  _selectionChanged QTextEdit_selectionChanged_signal;
//  _currentCharFormatChanged QTextEdit_currentCharFormatChanged_signal;
//  _undoAvailable QTextEdit_undoAvailable_signal;
//  _textChanged QTextEdit_textChanged_signal;
//  _copyAvailable QTextEdit_copyAvailable_signal;
}

  // proto:  int QTextEdit::lineWrapColumnOrWidth();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "lineWrapColumnOrWidth", args)
  }

}

  // proto:  void QTextEdit::setFontFamily(const QString & fontFamily);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontFamily", args)
  }

}

  // proto:  QString QTextEdit::toPlainText();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "toPlainText", args)
  }

}

  // proto:  void QTextEdit::setCursorWidth(int width);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setCursorWidth", args)
  }

}

  // proto:  QMenu * QTextEdit::createStandardContextMenu();
func (this *QTextEdit) createStandardContextMenu(args ...interface{}) () {
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
    // invoke: _ZN9QTextEdit25createStandardContextMenuEv
  case 1:
    // invoke: _ZN9QTextEdit25createStandardContextMenuERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "createStandardContextMenu", args)
  }

}

  // proto:  QTextDocument * QTextEdit::document();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "document", args)
  }

}

  // proto:  QRect QTextEdit::cursorRect();
func (this *QTextEdit) cursorRect(args ...interface{}) () {
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
    // invoke: _ZNK9QTextEdit10cursorRectEv
  case 1:
    // invoke: _ZNK9QTextEdit10cursorRectERK11QTextCursor
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorRect", args)
  }

}

  // proto:  void QTextEdit::setTextColor(const QColor & c);
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextColor", args)
  }

}

  // proto:  bool QTextEdit::acceptRichText();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "acceptRichText", args)
  }

}

  // proto:  void QTextEdit::clear();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "clear", args)
  }

}

  // proto:  void QTextEdit::insertHtml(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "insertHtml", args)
  }

}

  // proto:  QString QTextEdit::fontFamily();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "fontFamily", args)
  }

}

  // proto:  void QTextEdit::setFontUnderline(bool b);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontUnderline", args)
  }

}

  // proto:  void QTextEdit::cut();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "cut", args)
  }

}

  // proto:  QString QTextEdit::anchorAt(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "anchorAt", args)
  }

}

  // proto:  int QTextEdit::cursorWidth();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorWidth", args)
  }

}

  // proto:  void QTextEdit::setTextBackgroundColor(const QColor & c);
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextBackgroundColor", args)
  }

}

  // proto:  int QTextEdit::tabStopWidth();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "tabStopWidth", args)
  }

}

  // proto:  void QTextEdit::setFontWeight(int w);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontWeight", args)
  }

}

  // proto:  void QTextEdit::selectAll();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "selectAll", args)
  }

}

  // proto:  void QTextEdit::zoomOut(int range);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomOut", args)
  }

}

  // proto:  void QTextEdit::redo();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "redo", args)
  }

}

  // proto:  void QTextEdit::setFontPointSize(qreal s);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontPointSize", args)
  }

}

  // proto:  bool QTextEdit::overwriteMode();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "overwriteMode", args)
  }

}

  // proto:  QTextCursor QTextEdit::textCursor();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "textCursor", args)
  }

}

  // proto:  void QTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
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
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "mergeCurrentCharFormat", args)
  }

}

  // proto:  void QTextEdit::setPlainText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlainText", args)
  }

}

  // proto:  QString QTextEdit::placeholderText();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "placeholderText", args)
  }

}

  // proto:  void QTextEdit::~QTextEdit();
func (this *QTextEdit) FreeQTextEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextEdit", "~QTextEdit", args)
  }

}

  // proto:  bool QTextEdit::fontItalic();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "fontItalic", args)
  }

}

  // proto:  void QTextEdit::copy();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "copy", args)
  }

}

  // proto:  qreal QTextEdit::fontPointSize();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "fontPointSize", args)
  }

}

  // proto:  void QTextEdit::setDocument(QTextDocument * document);
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
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocument", args)
  }

}

  // proto:  void QTextEdit::setOverwriteMode(bool overwrite);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setOverwriteMode", args)
  }

}

  // proto:  void QTextEdit::undo();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "undo", args)
  }

}

  // proto:  void QTextEdit::zoomIn(int range);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomIn", args)
  }

}

  // proto:  void QTextEdit::setDocumentTitle(const QString & title);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocumentTitle", args)
  }

}

  // proto:  bool QTextEdit::canPaste();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "canPaste", args)
  }

}

  // proto:  QString QTextEdit::toHtml();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "toHtml", args)
  }

}

  // proto:  void QTextEdit::setTabStopWidth(int width);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabStopWidth", args)
  }

}

  // proto:  QString QTextEdit::documentTitle();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "documentTitle", args)
  }

}

  // proto:  bool QTextEdit::isUndoRedoEnabled();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "isUndoRedoEnabled", args)
  }

}

  // proto:  void QTextEdit::setText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setText", args)
  }

}

  // proto:  void QTextEdit::ensureCursorVisible();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "ensureCursorVisible", args)
  }

}

  // proto:  void QTextEdit::setAcceptRichText(bool accept);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setAcceptRichText", args)
  }

}

  // proto:  void QTextEdit::setPlaceholderText(const QString & placeholderText);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlaceholderText", args)
  }

}

  // proto:  bool QTextEdit::isReadOnly();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "isReadOnly", args)
  }

}

  // proto:  void QTextEdit::setUndoRedoEnabled(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setUndoRedoEnabled", args)
  }

}

  // proto:  void QTextEdit::QTextEdit(const QTextEdit & );
func NewQTextEdit(args ...interface{}) QTextEdit {
  return QTextEdit{}
}

  // proto:  QTextCharFormat QTextEdit::currentCharFormat();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "currentCharFormat", args)
  }

}

  // proto:  QTextCursor QTextEdit::cursorForPosition(const QPoint & pos);
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorForPosition", args)
  }

}

  // proto:  void QTextEdit::scrollToAnchor(const QString & name);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "scrollToAnchor", args)
  }

}

  // proto:  QFont QTextEdit::currentFont();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "currentFont", args)
  }

}

  // proto:  void QTextEdit::paste();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "paste", args)
  }

}

  // proto:  void QTextEdit::setTextCursor(const QTextCursor & cursor);
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
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextCursor", args)
  }

}

  // proto:  void QTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
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
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentCharFormat", args)
  }

}

  // proto:  QVariant QTextEdit::loadResource(int type, const QUrl & name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTextEdit", "loadResource", args)
  }

}

  // proto:  void QTextEdit::setTabChangesFocus(bool b);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabChangesFocus", args)
  }

}

  // proto:  void QTextEdit::setHtml(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setHtml", args)
  }

}

  // proto:  void QTextEdit::setLineWrapColumnOrWidth(int w);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setLineWrapColumnOrWidth", args)
  }

}

  // proto:  void QTextEdit::setFontItalic(bool b);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontItalic", args)
  }

}

  // proto:  const QMetaObject * QTextEdit::metaObject();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "metaObject", args)
  }

}

  // proto:  void QTextEdit::setCurrentFont(const QFont & f);
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
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentFont", args)
  }

}

  // proto:  bool QTextEdit::tabChangesFocus();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "tabChangesFocus", args)
  }

}

  // proto:  QColor QTextEdit::textBackgroundColor();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "textBackgroundColor", args)
  }

}

  // proto:  void QTextEdit::print(QPagedPaintDevice * printer);
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
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "print", args)
  }

}

  // proto:  bool QTextEdit::fontUnderline();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "fontUnderline", args)
  }

}

  // proto:  void QTextEdit::insertPlainText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "insertPlainText", args)
  }

}

  // proto:  int QTextEdit::fontWeight();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "fontWeight", args)
  }

}

  // proto:  QColor QTextEdit::textColor();
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
  default:
    qtrt.ErrorResolve("QTextEdit", "textColor", args)
  }

}

  // proto:  void QTextEdit::append(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "append", args)
  }

}

  // proto:  void QTextEdit::setReadOnly(bool ro);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextEdit", "setReadOnly", args)
  }

}

// <= body block end

