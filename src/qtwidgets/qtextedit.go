package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  AutoFormatting QTextEdit::autoFormatting();
extern void C_ZNK9QTextEdit14autoFormattingEv(void* qthis); // 4
  // proto:  void QTextEdit::setTabStopWidth(int width);
extern void C_ZN9QTextEdit15setTabStopWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QColor QTextEdit::textBackgroundColor();
extern void* C_ZNK9QTextEdit19textBackgroundColorEv(void* qthis); // 4
  // proto:  void QTextEdit::ensureCursorVisible();
extern void C_ZN9QTextEdit19ensureCursorVisibleEv(void* qthis); // 4
  // proto:  void QTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void C_ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::tabChangesFocus();
extern bool C_ZNK9QTextEdit15tabChangesFocusEv(void* qthis); // 4
  // proto:  Qt::Alignment QTextEdit::alignment();
extern void C_ZNK9QTextEdit9alignmentEv(void* qthis); // 4
  // proto:  bool QTextEdit::isUndoRedoEnabled();
extern bool C_ZNK9QTextEdit17isUndoRedoEnabledEv(void* qthis); // 2
  // proto:  void QTextEdit::cut();
extern void C_ZN9QTextEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QTextEdit::createStandardContextMenu(const QPoint & position);
extern void* C_ZN9QTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QMenu * QTextEdit::createStandardContextMenu();
extern void* C_ZN9QTextEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontPointSize(qreal s);
extern void C_ZN9QTextEdit16setFontPointSizeEd(void* qthis, double arg0); // 4
  // proto:  void QTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void C_ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  QString QTextEdit::documentTitle();
extern void* C_ZNK9QTextEdit13documentTitleEv(void* qthis); // 2
  // proto:  QString QTextEdit::toHtml();
extern void* C_ZNK9QTextEdit6toHtmlEv(void* qthis); // 4
  // proto:  void QTextEdit::insertPlainText(const QString & text);
extern void C_ZN9QTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  qreal QTextEdit::fontPointSize();
extern double C_ZNK9QTextEdit13fontPointSizeEv(void* qthis); // 4
  // proto:  QColor QTextEdit::textColor();
extern void* C_ZNK9QTextEdit9textColorEv(void* qthis); // 4
  // proto:  const QMetaObject * QTextEdit::metaObject();
extern void C_ZNK9QTextEdit10metaObjectEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextBackgroundColor(const QColor & c);
extern void C_ZN9QTextEdit22setTextBackgroundColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setTabChangesFocus(bool b);
extern void C_ZN9QTextEdit18setTabChangesFocusEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextEdit::toPlainText();
extern void* C_ZNK9QTextEdit11toPlainTextEv(void* qthis); // 4
  // proto:  void QTextEdit::zoomIn(int range);
extern void C_ZN9QTextEdit6zoomInEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::zoomOut(int range);
extern void C_ZN9QTextEdit7zoomOutEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QTextEdit::cursorRect(const QTextCursor & cursor);
extern void* C_ZNK9QTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QTextEdit::cursorRect();
extern void* C_ZNK9QTextEdit10cursorRectEv(void* qthis); // 4
  // proto:  void QTextEdit::setUndoRedoEnabled(bool enable);
extern void C_ZN9QTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0); // 2
  // proto:  QTextCursor QTextEdit::textCursor();
extern void* C_ZNK9QTextEdit10textCursorEv(void* qthis); // 4
  // proto:  void QTextEdit::setAcceptRichText(bool accept);
extern void C_ZN9QTextEdit17setAcceptRichTextEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextEdit::anchorAt(const QPoint & pos);
extern void* C_ZNK9QTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setHtml(const QString & text);
extern void C_ZN9QTextEdit7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setDocument(QTextDocument * document);
extern void C_ZN9QTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::fontItalic();
extern bool C_ZNK9QTextEdit10fontItalicEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontFamily(const QString & fontFamily);
extern void C_ZN9QTextEdit13setFontFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::print(QPagedPaintDevice * printer);
extern void C_ZNK9QTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QVariant QTextEdit::loadResource(int type, const QUrl & name);
extern void* C_ZN9QTextEdit12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTextEdit::undo();
extern void C_ZN9QTextEdit4undoEv(void* qthis); // 4
  // proto:  int QTextEdit::lineWrapColumnOrWidth();
extern int32_t C_ZNK9QTextEdit21lineWrapColumnOrWidthEv(void* qthis); // 4
  // proto:  void QTextEdit::paste();
extern void C_ZN9QTextEdit5pasteEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontWeight(int w);
extern void C_ZN9QTextEdit13setFontWeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::scrollToAnchor(const QString & name);
extern void C_ZN9QTextEdit14scrollToAnchorERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::selectAll();
extern void C_ZN9QTextEdit9selectAllEv(void* qthis); // 4
  // proto:  QList<QTextEdit::ExtraSelection> QTextEdit::extraSelections();
extern void C_ZNK9QTextEdit15extraSelectionsEv(void* qthis); // 4
  // proto:  void QTextEdit::~QTextEdit();
extern void C_ZN9QTextEditD2Ev(void* qthis); // 4
  // proto:  void QTextEdit::redo();
extern void C_ZN9QTextEdit4redoEv(void* qthis); // 4
  // proto:  void QTextEdit::setFontUnderline(bool b);
extern void C_ZN9QTextEdit16setFontUnderlineEb(void* qthis, bool arg0); // 4
  // proto:  bool QTextEdit::isReadOnly();
extern bool C_ZNK9QTextEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextCursor(const QTextCursor & cursor);
extern void C_ZN9QTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::overwriteMode();
extern bool C_ZNK9QTextEdit13overwriteModeEv(void* qthis); // 4
  // proto:  bool QTextEdit::acceptRichText();
extern bool C_ZNK9QTextEdit14acceptRichTextEv(void* qthis); // 4
  // proto:  void QTextEdit::setCursorWidth(int width);
extern void C_ZN9QTextEdit14setCursorWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextEdit::setLineWrapColumnOrWidth(int w);
extern void C_ZN9QTextEdit24setLineWrapColumnOrWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::TextInteractionFlags QTextEdit::textInteractionFlags();
extern void C_ZNK9QTextEdit20textInteractionFlagsEv(void* qthis); // 4
  // proto:  QFont QTextEdit::currentFont();
extern void* C_ZNK9QTextEdit11currentFontEv(void* qthis); // 4
  // proto:  void QTextEdit::setOverwriteMode(bool overwrite);
extern void C_ZN9QTextEdit16setOverwriteModeEb(void* qthis, bool arg0); // 4
  // proto:  void QTextEdit::copy();
extern void C_ZN9QTextEdit4copyEv(void* qthis); // 4
  // proto:  bool QTextEdit::fontUnderline();
extern bool C_ZNK9QTextEdit13fontUnderlineEv(void* qthis); // 4
  // proto:  int QTextEdit::cursorWidth();
extern int32_t C_ZNK9QTextEdit11cursorWidthEv(void* qthis); // 4
  // proto:  void QTextEdit::setPlaceholderText(const QString & placeholderText);
extern void C_ZN9QTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::clear();
extern void C_ZN9QTextEdit5clearEv(void* qthis); // 4
  // proto:  void QTextEdit::insertHtml(const QString & text);
extern void C_ZN9QTextEdit10insertHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setPlainText(const QString & text);
extern void C_ZN9QTextEdit12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QTextEdit::canPaste();
extern bool C_ZNK9QTextEdit8canPasteEv(void* qthis); // 4
  // proto:  QString QTextEdit::fontFamily();
extern void* C_ZNK9QTextEdit10fontFamilyEv(void* qthis); // 4
  // proto:  void QTextEdit::QTextEdit(QWidget * parent);
extern void* C_ZN9QTextEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QTextEdit::QTextEdit(const QString & text, QWidget * parent);
extern void* C_ZN9QTextEditC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  int QTextEdit::tabStopWidth();
extern int32_t C_ZNK9QTextEdit12tabStopWidthEv(void* qthis); // 4
  // proto:  QTextOption::WrapMode QTextEdit::wordWrapMode();
extern void C_ZNK9QTextEdit12wordWrapModeEv(void* qthis); // 4
  // proto:  QTextCursor QTextEdit::cursorForPosition(const QPoint & pos);
extern void* C_ZNK9QTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setDocumentTitle(const QString & title);
extern void C_ZN9QTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTextEdit::setFontItalic(bool b);
extern void C_ZN9QTextEdit13setFontItalicEb(void* qthis, bool arg0); // 4
  // proto:  QTextDocument * QTextEdit::document();
extern void* C_ZNK9QTextEdit8documentEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextEdit::currentCharFormat();
extern void* C_ZNK9QTextEdit17currentCharFormatEv(void* qthis); // 4
  // proto:  int QTextEdit::fontWeight();
extern int32_t C_ZNK9QTextEdit10fontWeightEv(void* qthis); // 4
  // proto:  QString QTextEdit::placeholderText();
extern void* C_ZNK9QTextEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QTextEdit::setTextColor(const QColor & c);
extern void C_ZN9QTextEdit12setTextColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QTextEdit::LineWrapMode QTextEdit::lineWrapMode();
extern void C_ZNK9QTextEdit12lineWrapModeEv(void* qthis); // 4
  // proto:  void QTextEdit::setText(const QString & text);
extern void C_ZN9QTextEdit7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setCurrentFont(const QFont & f);
extern void C_ZN9QTextEdit14setCurrentFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QTextEdit::setReadOnly(bool ro);
extern void C_ZN9QTextEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
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

// class sizeof(QTextEdit)=1
type QTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _cursorPositionChanged QTextEdit_cursorPositionChanged_signal;
//  _redoAvailable QTextEdit_redoAvailable_signal;
//  _selectionChanged QTextEdit_selectionChanged_signal;
//  _currentCharFormatChanged QTextEdit_currentCharFormatChanged_signal;
//  _undoAvailable QTextEdit_undoAvailable_signal;
//  _textChanged QTextEdit_textChanged_signal;
//  _copyAvailable QTextEdit_copyAvailable_signal;
}

// autoFormatting()
func (this *QTextEdit) Autoformatting(args ...interface{}) () {
  // autoFormatting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit14autoFormattingEv
    // invoke: AutoFormatting autoFormatting()
    C.C_ZNK9QTextEdit14autoFormattingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "autoFormatting", args)
  }

  return
}

// setTabStopWidth(int)
func (this *QTextEdit) Settabstopwidth(args ...interface{}) () {
  // setTabStopWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit15setTabStopWidthEi
    // invoke: void setTabStopWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit15setTabStopWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabStopWidth", args)
  }

  return
}

// textBackgroundColor()
func (this *QTextEdit) Textbackgroundcolor(args ...interface{}) (ret interface{}) {
  // textBackgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit19textBackgroundColorEv
    // invoke: QColor textBackgroundColor()
    var ret0 = C.C_ZNK9QTextEdit19textBackgroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "textBackgroundColor", args)
  }

  return
}

// ensureCursorVisible()
func (this *QTextEdit) Ensurecursorvisible(args ...interface{}) () {
  // ensureCursorVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit19ensureCursorVisibleEv
    // invoke: void ensureCursorVisible()
    C.C_ZN9QTextEdit19ensureCursorVisibleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "ensureCursorVisible", args)
  }

  return
}

// setCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) Setcurrentcharformat(args ...interface{}) () {
  // setCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCharFormat{}) // "const QTextCharFormat &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat
    // invoke: void setCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(*qtgui.QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentCharFormat", args)
  }

  return
}

// tabChangesFocus()
func (this *QTextEdit) Tabchangesfocus(args ...interface{}) (ret interface{}) {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15tabChangesFocusEv
    // invoke: bool tabChangesFocus()
    var ret0 = C.C_ZNK9QTextEdit15tabChangesFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "tabChangesFocus", args)
  }

  return
}

// alignment()
func (this *QTextEdit) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK9QTextEdit9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "alignment", args)
  }

  return
}

// isUndoRedoEnabled()
func (this *QTextEdit) Isundoredoenabled(args ...interface{}) (ret interface{}) {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17isUndoRedoEnabledEv
    // invoke: bool isUndoRedoEnabled()
    var ret0 = C.C_ZNK9QTextEdit17isUndoRedoEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "isUndoRedoEnabled", args)
  }

  return
}

// cut()
func (this *QTextEdit) Cut(args ...interface{}) () {
  // cut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit3cutEv
    // invoke: void cut()
    C.C_ZN9QTextEdit3cutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "cut", args)
  }

  return
}

// createStandardContextMenu(const class QPoint &)
func (this *QTextEdit) Createstandardcontextmenu(args ...interface{}) (ret interface{}) {
  // createStandardContextMenu(const class QPoint &)
  // createStandardContextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit25createStandardContextMenuERK6QPoint
    // invoke: QMenu * createStandardContextMenu(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QTextEdit25createStandardContextMenuERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QTextEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    var ret0 = C.C_ZN9QTextEdit25createStandardContextMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "createStandardContextMenu", args)
  }

  return
}

// setFontPointSize(qreal)
func (this *QTextEdit) Setfontpointsize(args ...interface{}) () {
  // setFontPointSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setFontPointSizeEd
    // invoke: void setFontPointSize(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit16setFontPointSizeEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontPointSize", args)
  }

  return
}

// mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) Mergecurrentcharformat(args ...interface{}) () {
  // mergeCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCharFormat{}) // "const QTextCharFormat &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat
    // invoke: void mergeCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(*qtgui.QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "mergeCurrentCharFormat", args)
  }

  return
}

// documentTitle()
func (this *QTextEdit) Documenttitle(args ...interface{}) (ret interface{}) {
  // documentTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13documentTitleEv
    // invoke: QString documentTitle()
    var ret0 = C.C_ZNK9QTextEdit13documentTitleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "documentTitle", args)
  }

  return
}

// toHtml()
func (this *QTextEdit) Tohtml(args ...interface{}) (ret interface{}) {
  // toHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit6toHtmlEv
    // invoke: QString toHtml()
    var ret0 = C.C_ZNK9QTextEdit6toHtmlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "toHtml", args)
  }

  return
}

// insertPlainText(const class QString &)
func (this *QTextEdit) Insertplaintext(args ...interface{}) () {
  // insertPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit15insertPlainTextERK7QString
    // invoke: void insertPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit15insertPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "insertPlainText", args)
  }

  return
}

// fontPointSize()
func (this *QTextEdit) Fontpointsize(args ...interface{}) (ret interface{}) {
  // fontPointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13fontPointSizeEv
    // invoke: qreal fontPointSize()
    var ret0 = C.C_ZNK9QTextEdit13fontPointSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "fontPointSize", args)
  }

  return
}

// textColor()
func (this *QTextEdit) Textcolor(args ...interface{}) (ret interface{}) {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit9textColorEv
    // invoke: QColor textColor()
    var ret0 = C.C_ZNK9QTextEdit9textColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "textColor", args)
  }

  return
}

// metaObject()
func (this *QTextEdit) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QTextEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "metaObject", args)
  }

  return
}

// setTextBackgroundColor(const class QColor &)
func (this *QTextEdit) Settextbackgroundcolor(args ...interface{}) () {
  // setTextBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit22setTextBackgroundColorERK6QColor
    // invoke: void setTextBackgroundColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit22setTextBackgroundColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextBackgroundColor", args)
  }

  return
}

// setTabChangesFocus(_Bool)
func (this *QTextEdit) Settabchangesfocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setTabChangesFocusEb
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit18setTabChangesFocusEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTabChangesFocus", args)
  }

  return
}

// toPlainText()
func (this *QTextEdit) Toplaintext(args ...interface{}) (ret interface{}) {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11toPlainTextEv
    // invoke: QString toPlainText()
    var ret0 = C.C_ZNK9QTextEdit11toPlainTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "toPlainText", args)
  }

  return
}

// zoomIn(int)
func (this *QTextEdit) Zoomin(args ...interface{}) () {
  // zoomIn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit6zoomInEi
    // invoke: void zoomIn(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit6zoomInEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomIn", args)
  }

  return
}

// zoomOut(int)
func (this *QTextEdit) Zoomout(args ...interface{}) () {
  // zoomOut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7zoomOutEi
    // invoke: void zoomOut(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit7zoomOutEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "zoomOut", args)
  }

  return
}

// cursorRect(const class QTextCursor &)
func (this *QTextEdit) Cursorrect(args ...interface{}) (ret interface{}) {
  // cursorRect(const class QTextCursor &)
  // cursorRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCursor{}) // "const QTextCursor &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10cursorRectERK11QTextCursor
    // invoke: QRect cursorRect(const class QTextCursor &)
    var arg0 = args[0].(*qtgui.QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextEdit10cursorRectERK11QTextCursor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QTextEdit10cursorRectEv
    // invoke: QRect cursorRect()
    var ret0 = C.C_ZNK9QTextEdit10cursorRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorRect", args)
  }

  return
}

// setUndoRedoEnabled(_Bool)
func (this *QTextEdit) Setundoredoenabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setUndoRedoEnabledEb
    // invoke: void setUndoRedoEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit18setUndoRedoEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setUndoRedoEnabled", args)
  }

  return
}

// textCursor()
func (this *QTextEdit) Textcursor(args ...interface{}) (ret interface{}) {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10textCursorEv
    // invoke: QTextCursor textCursor()
    var ret0 = C.C_ZNK9QTextEdit10textCursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "textCursor", args)
  }

  return
}

// setAcceptRichText(_Bool)
func (this *QTextEdit) Setacceptrichtext(args ...interface{}) () {
  // setAcceptRichText(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit17setAcceptRichTextEb
    // invoke: void setAcceptRichText(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit17setAcceptRichTextEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setAcceptRichText", args)
  }

  return
}

// anchorAt(const class QPoint &)
func (this *QTextEdit) Anchorat(args ...interface{}) (ret interface{}) {
  // anchorAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8anchorAtERK6QPoint
    // invoke: QString anchorAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextEdit8anchorAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "anchorAt", args)
  }

  return
}

// setHtml(const class QString &)
func (this *QTextEdit) Sethtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7setHtmlERK7QString
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit7setHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setHtml", args)
  }

  return
}

// setDocument(class QTextDocument *)
func (this *QTextEdit) Setdocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextDocument{}) // "QTextDocument *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(*qtgui.QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit11setDocumentEP13QTextDocument(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocument", args)
  }

  return
}

// fontItalic()
func (this *QTextEdit) Fontitalic(args ...interface{}) (ret interface{}) {
  // fontItalic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontItalicEv
    // invoke: bool fontItalic()
    var ret0 = C.C_ZNK9QTextEdit10fontItalicEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "fontItalic", args)
  }

  return
}

// setFontFamily(const class QString &)
func (this *QTextEdit) Setfontfamily(args ...interface{}) () {
  // setFontFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontFamilyERK7QString
    // invoke: void setFontFamily(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit13setFontFamilyERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontFamily", args)
  }

  return
}

// print(class QPagedPaintDevice *)
func (this *QTextEdit) Print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPagedPaintDevice{}) // "QPagedPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit5printEP17QPagedPaintDevice
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(*qtgui.QPagedPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTextEdit5printEP17QPagedPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "print", args)
  }

  return
}

// loadResource(int, const class QUrl &)
func (this *QTextEdit) Loadresource(args ...interface{}) (ret interface{}) {
  // loadResource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit12loadResourceEiRK4QUrl
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QTextEdit12loadResourceEiRK4QUrl(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "loadResource", args)
  }

  return
}

// undo()
func (this *QTextEdit) Undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4undoEv
    // invoke: void undo()
    C.C_ZN9QTextEdit4undoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "undo", args)
  }

  return
}

// lineWrapColumnOrWidth()
func (this *QTextEdit) Linewrapcolumnorwidth(args ...interface{}) (ret interface{}) {
  // lineWrapColumnOrWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit21lineWrapColumnOrWidthEv
    // invoke: int lineWrapColumnOrWidth()
    var ret0 = C.C_ZNK9QTextEdit21lineWrapColumnOrWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "lineWrapColumnOrWidth", args)
  }

  return
}

// paste()
func (this *QTextEdit) Paste(args ...interface{}) () {
  // paste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit5pasteEv
    // invoke: void paste()
    C.C_ZN9QTextEdit5pasteEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "paste", args)
  }

  return
}

// setFontWeight(int)
func (this *QTextEdit) Setfontweight(args ...interface{}) () {
  // setFontWeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontWeightEi
    // invoke: void setFontWeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit13setFontWeightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontWeight", args)
  }

  return
}

// scrollToAnchor(const class QString &)
func (this *QTextEdit) Scrolltoanchor(args ...interface{}) () {
  // scrollToAnchor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14scrollToAnchorERK7QString
    // invoke: void scrollToAnchor(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit14scrollToAnchorERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "scrollToAnchor", args)
  }

  return
}

// selectAll()
func (this *QTextEdit) Selectall(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit9selectAllEv
    // invoke: void selectAll()
    C.C_ZN9QTextEdit9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "selectAll", args)
  }

  return
}

// extraSelections()
func (this *QTextEdit) Extraselections(args ...interface{}) () {
  // extraSelections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15extraSelectionsEv
    // invoke: QList<QTextEdit::ExtraSelection> extraSelections()
    C.C_ZNK9QTextEdit15extraSelectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "extraSelections", args)
  }

  return
}

// ~QTextEdit()
func (this *QTextEdit) Freeqtextedit(args ...interface{}) () {
  // ~QTextEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEditD0Ev
    // invoke: void ~QTextEdit()
    C.C_ZN9QTextEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "~QTextEdit", args)
  }

  return
}

// redo()
func (this *QTextEdit) Redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4redoEv
    // invoke: void redo()
    C.C_ZN9QTextEdit4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "redo", args)
  }

  return
}

// setFontUnderline(_Bool)
func (this *QTextEdit) Setfontunderline(args ...interface{}) () {
  // setFontUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setFontUnderlineEb
    // invoke: void setFontUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit16setFontUnderlineEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontUnderline", args)
  }

  return
}

// isReadOnly()
func (this *QTextEdit) Isreadonly(args ...interface{}) (ret interface{}) {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK9QTextEdit10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "isReadOnly", args)
  }

  return
}

// setTextCursor(const class QTextCursor &)
func (this *QTextEdit) Settextcursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QTextCursor{}) // "const QTextCursor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setTextCursorERK11QTextCursor
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(*qtgui.QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit13setTextCursorERK11QTextCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextCursor", args)
  }

  return
}

// overwriteMode()
func (this *QTextEdit) Overwritemode(args ...interface{}) (ret interface{}) {
  // overwriteMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13overwriteModeEv
    // invoke: bool overwriteMode()
    var ret0 = C.C_ZNK9QTextEdit13overwriteModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "overwriteMode", args)
  }

  return
}

// acceptRichText()
func (this *QTextEdit) Acceptrichtext(args ...interface{}) (ret interface{}) {
  // acceptRichText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit14acceptRichTextEv
    // invoke: bool acceptRichText()
    var ret0 = C.C_ZNK9QTextEdit14acceptRichTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "acceptRichText", args)
  }

  return
}

// setCursorWidth(int)
func (this *QTextEdit) Setcursorwidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14setCursorWidthEi
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit14setCursorWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCursorWidth", args)
  }

  return
}

// setLineWrapColumnOrWidth(int)
func (this *QTextEdit) Setlinewrapcolumnorwidth(args ...interface{}) () {
  // setLineWrapColumnOrWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit24setLineWrapColumnOrWidthEi
    // invoke: void setLineWrapColumnOrWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit24setLineWrapColumnOrWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setLineWrapColumnOrWidth", args)
  }

  return
}

// textInteractionFlags()
func (this *QTextEdit) Textinteractionflags(args ...interface{}) () {
  // textInteractionFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit20textInteractionFlagsEv
    // invoke: Qt::TextInteractionFlags textInteractionFlags()
    C.C_ZNK9QTextEdit20textInteractionFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "textInteractionFlags", args)
  }

  return
}

// currentFont()
func (this *QTextEdit) Currentfont(args ...interface{}) (ret interface{}) {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11currentFontEv
    // invoke: QFont currentFont()
    var ret0 = C.C_ZNK9QTextEdit11currentFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "currentFont", args)
  }

  return
}

// setOverwriteMode(_Bool)
func (this *QTextEdit) Setoverwritemode(args ...interface{}) () {
  // setOverwriteMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setOverwriteModeEb
    // invoke: void setOverwriteMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit16setOverwriteModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setOverwriteMode", args)
  }

  return
}

// copy()
func (this *QTextEdit) Copy(args ...interface{}) () {
  // copy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit4copyEv
    // invoke: void copy()
    C.C_ZN9QTextEdit4copyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "copy", args)
  }

  return
}

// fontUnderline()
func (this *QTextEdit) Fontunderline(args ...interface{}) (ret interface{}) {
  // fontUnderline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit13fontUnderlineEv
    // invoke: bool fontUnderline()
    var ret0 = C.C_ZNK9QTextEdit13fontUnderlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "fontUnderline", args)
  }

  return
}

// cursorWidth()
func (this *QTextEdit) Cursorwidth(args ...interface{}) (ret interface{}) {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit11cursorWidthEv
    // invoke: int cursorWidth()
    var ret0 = C.C_ZNK9QTextEdit11cursorWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorWidth", args)
  }

  return
}

// setPlaceholderText(const class QString &)
func (this *QTextEdit) Setplaceholdertext(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit18setPlaceholderTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlaceholderText", args)
  }

  return
}

// clear()
func (this *QTextEdit) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit5clearEv
    // invoke: void clear()
    C.C_ZN9QTextEdit5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "clear", args)
  }

  return
}

// insertHtml(const class QString &)
func (this *QTextEdit) Inserthtml(args ...interface{}) () {
  // insertHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit10insertHtmlERK7QString
    // invoke: void insertHtml(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit10insertHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "insertHtml", args)
  }

  return
}

// setPlainText(const class QString &)
func (this *QTextEdit) Setplaintext(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit12setPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setPlainText", args)
  }

  return
}

// canPaste()
func (this *QTextEdit) Canpaste(args ...interface{}) (ret interface{}) {
  // canPaste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8canPasteEv
    // invoke: bool canPaste()
    var ret0 = C.C_ZNK9QTextEdit8canPasteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "canPaste", args)
  }

  return
}

// fontFamily()
func (this *QTextEdit) Fontfamily(args ...interface{}) (ret interface{}) {
  // fontFamily()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontFamilyEv
    // invoke: QString fontFamily()
    var ret0 = C.C_ZNK9QTextEdit10fontFamilyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "fontFamily", args)
  }

  return
}

// QTextEdit(class QWidget *)
func NewQTextEdit(args ...interface{}) *QTextEdit {
  // QTextEdit(class QWidget *)
  // QTextEdit(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEditC1EP7QWidget
    // invoke: void QTextEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTextEditC2EP7QWidget(arg0)
    return &QTextEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QTextEditC1ERK7QStringP7QWidget
    // invoke: void QTextEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTextEditC2ERK7QStringP7QWidget(arg0, arg1)
    return &QTextEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextEdit", "QTextEdit", args)
  }

  return nil // QTextEdit{Qclsinst:qthis}
}

// tabStopWidth()
func (this *QTextEdit) Tabstopwidth(args ...interface{}) (ret interface{}) {
  // tabStopWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12tabStopWidthEv
    // invoke: int tabStopWidth()
    var ret0 = C.C_ZNK9QTextEdit12tabStopWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "tabStopWidth", args)
  }

  return
}

// wordWrapMode()
func (this *QTextEdit) Wordwrapmode(args ...interface{}) () {
  // wordWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12wordWrapModeEv
    // invoke: QTextOption::WrapMode wordWrapMode()
    C.C_ZNK9QTextEdit12wordWrapModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "wordWrapMode", args)
  }

  return
}

// cursorForPosition(const class QPoint &)
func (this *QTextEdit) Cursorforposition(args ...interface{}) (ret interface{}) {
  // cursorForPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17cursorForPositionERK6QPoint
    // invoke: QTextCursor cursorForPosition(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextEdit17cursorForPositionERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "cursorForPosition", args)
  }

  return
}

// setDocumentTitle(const class QString &)
func (this *QTextEdit) Setdocumenttitle(args ...interface{}) () {
  // setDocumentTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit16setDocumentTitleERK7QString
    // invoke: void setDocumentTitle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit16setDocumentTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setDocumentTitle", args)
  }

  return
}

// setFontItalic(_Bool)
func (this *QTextEdit) Setfontitalic(args ...interface{}) () {
  // setFontItalic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit13setFontItalicEb
    // invoke: void setFontItalic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit13setFontItalicEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setFontItalic", args)
  }

  return
}

// document()
func (this *QTextEdit) Document(args ...interface{}) (ret interface{}) {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit8documentEv
    // invoke: QTextDocument * document()
    var ret0 = C.C_ZNK9QTextEdit8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "document", args)
  }

  return
}

// currentCharFormat()
func (this *QTextEdit) Currentcharformat(args ...interface{}) (ret interface{}) {
  // currentCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit17currentCharFormatEv
    // invoke: QTextCharFormat currentCharFormat()
    var ret0 = C.C_ZNK9QTextEdit17currentCharFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "currentCharFormat", args)
  }

  return
}

// fontWeight()
func (this *QTextEdit) Fontweight(args ...interface{}) (ret interface{}) {
  // fontWeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit10fontWeightEv
    // invoke: int fontWeight()
    var ret0 = C.C_ZNK9QTextEdit10fontWeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "fontWeight", args)
  }

  return
}

// placeholderText()
func (this *QTextEdit) Placeholdertext(args ...interface{}) (ret interface{}) {
  // placeholderText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit15placeholderTextEv
    // invoke: QString placeholderText()
    var ret0 = C.C_ZNK9QTextEdit15placeholderTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEdit", "placeholderText", args)
  }

  return
}

// setTextColor(const class QColor &)
func (this *QTextEdit) Settextcolor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit12setTextColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setTextColor", args)
  }

  return
}

// lineWrapMode()
func (this *QTextEdit) Linewrapmode(args ...interface{}) () {
  // lineWrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextEdit12lineWrapModeEv
    // invoke: QTextEdit::LineWrapMode lineWrapMode()
    C.C_ZNK9QTextEdit12lineWrapModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextEdit", "lineWrapMode", args)
  }

  return
}

// setText(const class QString &)
func (this *QTextEdit) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setText", args)
  }

  return
}

// setCurrentFont(const class QFont &)
func (this *QTextEdit) Setcurrentfont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit14setCurrentFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setCurrentFont", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QTextEdit) Setreadonly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextEdit11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextEdit11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEdit", "setReadOnly", args)
  }

  return
}

// <= body block end

