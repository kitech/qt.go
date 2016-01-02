package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qlineedit.h
// dst-file: /src/widgets/qlineedit.go
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
  // proto:  void QLineEdit::cursorBackward(bool mark, int steps);
extern void _ZN9QLineEdit14cursorBackwardEbi(void* qthis, bool arg0, int arg1);
  // proto:  void QLineEdit::home(bool mark);
extern void _ZN9QLineEdit4homeEb(void* qthis, bool arg0);
  // proto:  int QLineEdit::selectionStart();
extern void _ZNK9QLineEdit14selectionStartEv(void* qthis);
  // proto:  void QLineEdit::setCursorPosition(int );
extern void _ZN9QLineEdit17setCursorPositionEi(void* qthis, int arg0);
  // proto:  bool QLineEdit::isRedoAvailable();
extern void _ZNK9QLineEdit15isRedoAvailableEv(void* qthis);
  // proto:  void QLineEdit::setModified(bool );
extern void _ZN9QLineEdit11setModifiedEb(void* qthis, bool arg0);
  // proto:  void QLineEdit::QLineEdit(const QString & , QWidget * parent);
extern void* dector_ZN9QLineEditC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN9QLineEditC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  bool QLineEdit::event(QEvent * );
extern void _ZN9QLineEdit5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  int QLineEdit::maxLength();
extern void _ZNK9QLineEdit9maxLengthEv(void* qthis);
  // proto:  QMenu * QLineEdit::createStandardContextMenu();
extern void _ZN9QLineEdit25createStandardContextMenuEv(void* qthis);
  // proto:  void QLineEdit::setTextMargins(const QMargins & margins);
extern void _ZN9QLineEdit14setTextMarginsERK8QMargins(void* qthis, void* arg0);
  // proto:  int QLineEdit::cursorPositionAt(const QPoint & pos);
extern void _ZN9QLineEdit16cursorPositionAtERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QLineEdit::hasSelectedText();
extern void _ZNK9QLineEdit15hasSelectedTextEv(void* qthis);
  // proto:  void QLineEdit::setPlaceholderText(const QString & );
extern void _ZN9QLineEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0);
  // proto:  QSize QLineEdit::minimumSizeHint();
extern void _ZNK9QLineEdit15minimumSizeHintEv(void* qthis);
  // proto:  void QLineEdit::cursorForward(bool mark, int steps);
extern void _ZN9QLineEdit13cursorForwardEbi(void* qthis, bool arg0, int arg1);
  // proto:  void QLineEdit::insert(const QString & );
extern void _ZN9QLineEdit6insertERK7QString(void* qthis, void* arg0);
  // proto:  void QLineEdit::setText(const QString & );
extern void _ZN9QLineEdit7setTextERK7QString(void* qthis, void* arg0);
  // proto:  const QValidator * QLineEdit::validator();
extern void _ZNK9QLineEdit9validatorEv(void* qthis);
  // proto:  void QLineEdit::deselect();
extern void _ZN9QLineEdit8deselectEv(void* qthis);
  // proto:  QString QLineEdit::inputMask();
extern void _ZNK9QLineEdit9inputMaskEv(void* qthis);
  // proto:  QString QLineEdit::placeholderText();
extern void _ZNK9QLineEdit15placeholderTextEv(void* qthis);
  // proto:  void QLineEdit::cut();
extern void _ZN9QLineEdit3cutEv(void* qthis);
  // proto:  QString QLineEdit::text();
extern void _ZNK9QLineEdit4textEv(void* qthis);
  // proto:  const QMetaObject * QLineEdit::metaObject();
extern void _ZNK9QLineEdit10metaObjectEv(void* qthis);
  // proto:  void QLineEdit::del();
extern void _ZN9QLineEdit3delEv(void* qthis);
  // proto:  bool QLineEdit::isModified();
extern void _ZNK9QLineEdit10isModifiedEv(void* qthis);
  // proto:  void QLineEdit::cursorWordForward(bool mark);
extern void _ZN9QLineEdit17cursorWordForwardEb(void* qthis, bool arg0);
  // proto:  void QLineEdit::selectAll();
extern void _ZN9QLineEdit9selectAllEv(void* qthis);
  // proto:  void QLineEdit::setSelection(int , int );
extern void _ZN9QLineEdit12setSelectionEii(void* qthis, int arg0, int arg1);
  // proto:  void QLineEdit::setCompleter(QCompleter * completer);
extern void _ZN9QLineEdit12setCompleterEP10QCompleter(void* qthis, void* arg0);
  // proto:  void QLineEdit::setMaxLength(int );
extern void _ZN9QLineEdit12setMaxLengthEi(void* qthis, int arg0);
  // proto:  void QLineEdit::~QLineEdit();
extern void _ZN9QLineEditD0Ev(void* qthis);
  // proto:  void QLineEdit::setReadOnly(bool );
extern void _ZN9QLineEdit11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  QString QLineEdit::displayText();
extern void _ZNK9QLineEdit11displayTextEv(void* qthis);
  // proto:  void QLineEdit::setFrame(bool );
extern void _ZN9QLineEdit8setFrameEb(void* qthis, bool arg0);
  // proto:  bool QLineEdit::hasAcceptableInput();
extern void _ZNK9QLineEdit18hasAcceptableInputEv(void* qthis);
  // proto:  bool QLineEdit::hasFrame();
extern void _ZNK9QLineEdit8hasFrameEv(void* qthis);
  // proto:  int QLineEdit::cursorPosition();
extern void _ZNK9QLineEdit14cursorPositionEv(void* qthis);
  // proto:  void QLineEdit::cursorWordBackward(bool mark);
extern void _ZN9QLineEdit18cursorWordBackwardEb(void* qthis, bool arg0);
  // proto:  bool QLineEdit::dragEnabled();
extern void _ZNK9QLineEdit11dragEnabledEv(void* qthis);
  // proto:  QSize QLineEdit::sizeHint();
extern void _ZNK9QLineEdit8sizeHintEv(void* qthis);
  // proto:  void QLineEdit::paste();
extern void _ZN9QLineEdit5pasteEv(void* qthis);
  // proto:  void QLineEdit::setValidator(const QValidator * );
extern void _ZN9QLineEdit12setValidatorEPK10QValidator(void* qthis, void* arg0);
  // proto:  void QLineEdit::QLineEdit(QWidget * parent);
extern void* dector_ZN9QLineEditC1EP7QWidget(void* arg0);
extern void _ZN9QLineEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QCompleter * QLineEdit::completer();
extern void _ZNK9QLineEdit9completerEv(void* qthis);
  // proto:  QMargins QLineEdit::textMargins();
extern void _ZNK9QLineEdit11textMarginsEv(void* qthis);
  // proto:  void QLineEdit::setClearButtonEnabled(bool enable);
extern void _ZN9QLineEdit21setClearButtonEnabledEb(void* qthis, bool arg0);
  // proto:  QString QLineEdit::selectedText();
extern void _ZNK9QLineEdit12selectedTextEv(void* qthis);
  // proto:  void QLineEdit::clear();
extern void _ZN9QLineEdit5clearEv(void* qthis);
  // proto:  void QLineEdit::copy();
extern void _ZNK9QLineEdit4copyEv(void* qthis);
  // proto:  bool QLineEdit::isUndoAvailable();
extern void _ZNK9QLineEdit15isUndoAvailableEv(void* qthis);
  // proto:  void QLineEdit::undo();
extern void _ZN9QLineEdit4undoEv(void* qthis);
  // proto:  bool QLineEdit::isClearButtonEnabled();
extern void _ZNK9QLineEdit20isClearButtonEnabledEv(void* qthis);
  // proto:  void QLineEdit::QLineEdit(const QLineEdit & );
extern void* dector_ZN9QLineEditC1ERKS_(void* arg0);
extern void _ZN9QLineEditC1ERKS_(void* qthis, void* arg0);
  // proto:  void QLineEdit::end(bool mark);
extern void _ZN9QLineEdit3endEb(void* qthis, bool arg0);
  // proto:  void QLineEdit::setDragEnabled(bool b);
extern void _ZN9QLineEdit14setDragEnabledEb(void* qthis, bool arg0);
  // proto:  void QLineEdit::backspace();
extern void _ZN9QLineEdit9backspaceEv(void* qthis);
  // proto:  void QLineEdit::redo();
extern void _ZN9QLineEdit4redoEv(void* qthis);
  // proto:  void QLineEdit::setTextMargins(int left, int top, int right, int bottom);
extern void _ZN9QLineEdit14setTextMarginsEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QLineEdit::setInputMask(const QString & inputMask);
extern void _ZN9QLineEdit12setInputMaskERK7QString(void* qthis, void* arg0);
  // proto:  void QLineEdit::getTextMargins(int * left, int * top, int * right, int * bottom);
extern void _ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  bool QLineEdit::isReadOnly();
extern void _ZNK9QLineEdit10isReadOnlyEv(void* qthis);
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

// class sizeof(QLineEdit)=1
type QLineEdit struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _textEdited QLineEdit_textEdited_signal;
//  _returnPressed QLineEdit_returnPressed_signal;
//  _selectionChanged QLineEdit_selectionChanged_signal;
//  _cursorPositionChanged QLineEdit_cursorPositionChanged_signal;
//  _editingFinished QLineEdit_editingFinished_signal;
//  _textChanged QLineEdit_textChanged_signal;
}

  // proto:  void QLineEdit::cursorBackward(bool mark, int steps);
func (this *QLineEdit) cursorBackward(args ...interface{}) () {
  // cursorBackward(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit14cursorBackwardEbi
    // invoke: void cursorBackward(_Bool, int)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QLineEdit14cursorBackwardEbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorBackward", args)
  }

}

  // proto:  void QLineEdit::home(bool mark);
func (this *QLineEdit) home(args ...interface{}) () {
  // home(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit4homeEb
    // invoke: void home(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit4homeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "home", args)
  }

}

  // proto:  int QLineEdit::selectionStart();
func (this *QLineEdit) selectionStart(args ...interface{}) () {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit14selectionStartEv
    // invoke: int selectionStart()
    C._ZNK9QLineEdit14selectionStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "selectionStart", args)
  }

}

  // proto:  void QLineEdit::setCursorPosition(int );
func (this *QLineEdit) setCursorPosition(args ...interface{}) () {
  // setCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit17setCursorPositionEi
    // invoke: void setCursorPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit17setCursorPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setCursorPosition", args)
  }

}

  // proto:  bool QLineEdit::isRedoAvailable();
func (this *QLineEdit) isRedoAvailable(args ...interface{}) () {
  // isRedoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15isRedoAvailableEv
    // invoke: bool isRedoAvailable()
    C._ZNK9QLineEdit15isRedoAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "isRedoAvailable", args)
  }

}

  // proto:  void QLineEdit::setModified(bool );
func (this *QLineEdit) setModified(args ...interface{}) () {
  // setModified(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit11setModifiedEb
    // invoke: void setModified(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit11setModifiedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setModified", args)
  }

}

  // proto:  void QLineEdit::QLineEdit(const QString & , QWidget * parent);
func NewQLineEdit(args ...interface{}) QLineEdit {
  return QLineEdit{}
}

  // proto:  bool QLineEdit::event(QEvent * );
func (this *QLineEdit) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "event", args)
  }

}

  // proto:  int QLineEdit::maxLength();
func (this *QLineEdit) maxLength(args ...interface{}) () {
  // maxLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit9maxLengthEv
    // invoke: int maxLength()
    C._ZNK9QLineEdit9maxLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "maxLength", args)
  }

}

  // proto:  QMenu * QLineEdit::createStandardContextMenu();
func (this *QLineEdit) createStandardContextMenu(args ...interface{}) () {
  // createStandardContextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    C._ZN9QLineEdit25createStandardContextMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "createStandardContextMenu", args)
  }

}

  // proto:  void QLineEdit::setTextMargins(const QMargins & margins);
func (this *QLineEdit) setTextMargins(args ...interface{}) () {
  // setTextMargins(const class QMargins &)
  // setTextMargins(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit14setTextMarginsERK8QMargins
    // invoke: void setTextMargins(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit14setTextMarginsERK8QMargins(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QLineEdit14setTextMarginsEiiii
    // invoke: void setTextMargins(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN9QLineEdit14setTextMarginsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLineEdit", "setTextMargins", args)
  }

}

  // proto:  int QLineEdit::cursorPositionAt(const QPoint & pos);
func (this *QLineEdit) cursorPositionAt(args ...interface{}) () {
  // cursorPositionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit16cursorPositionAtERK6QPoint
    // invoke: int cursorPositionAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit16cursorPositionAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorPositionAt", args)
  }

}

  // proto:  bool QLineEdit::hasSelectedText();
func (this *QLineEdit) hasSelectedText(args ...interface{}) () {
  // hasSelectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15hasSelectedTextEv
    // invoke: bool hasSelectedText()
    C._ZNK9QLineEdit15hasSelectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "hasSelectedText", args)
  }

}

  // proto:  void QLineEdit::setPlaceholderText(const QString & );
func (this *QLineEdit) setPlaceholderText(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit18setPlaceholderTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setPlaceholderText", args)
  }

}

  // proto:  QSize QLineEdit::minimumSizeHint();
func (this *QLineEdit) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK9QLineEdit15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "minimumSizeHint", args)
  }

}

  // proto:  void QLineEdit::cursorForward(bool mark, int steps);
func (this *QLineEdit) cursorForward(args ...interface{}) () {
  // cursorForward(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit13cursorForwardEbi
    // invoke: void cursorForward(_Bool, int)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QLineEdit13cursorForwardEbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorForward", args)
  }

}

  // proto:  void QLineEdit::insert(const QString & );
func (this *QLineEdit) insert(args ...interface{}) () {
  // insert(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit6insertERK7QString
    // invoke: void insert(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit6insertERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "insert", args)
  }

}

  // proto:  void QLineEdit::setText(const QString & );
func (this *QLineEdit) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setText", args)
  }

}

  // proto:  const QValidator * QLineEdit::validator();
func (this *QLineEdit) validator(args ...interface{}) () {
  // validator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit9validatorEv
    // invoke: const QValidator * validator()
    C._ZNK9QLineEdit9validatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "validator", args)
  }

}

  // proto:  void QLineEdit::deselect();
func (this *QLineEdit) deselect(args ...interface{}) () {
  // deselect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit8deselectEv
    // invoke: void deselect()
    C._ZN9QLineEdit8deselectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "deselect", args)
  }

}

  // proto:  QString QLineEdit::inputMask();
func (this *QLineEdit) inputMask(args ...interface{}) () {
  // inputMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit9inputMaskEv
    // invoke: QString inputMask()
    C._ZNK9QLineEdit9inputMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "inputMask", args)
  }

}

  // proto:  QString QLineEdit::placeholderText();
func (this *QLineEdit) placeholderText(args ...interface{}) () {
  // placeholderText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15placeholderTextEv
    // invoke: QString placeholderText()
    C._ZNK9QLineEdit15placeholderTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "placeholderText", args)
  }

}

  // proto:  void QLineEdit::cut();
func (this *QLineEdit) cut(args ...interface{}) () {
  // cut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit3cutEv
    // invoke: void cut()
    C._ZN9QLineEdit3cutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "cut", args)
  }

}

  // proto:  QString QLineEdit::text();
func (this *QLineEdit) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit4textEv
    // invoke: QString text()
    C._ZNK9QLineEdit4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "text", args)
  }

}

  // proto:  const QMetaObject * QLineEdit::metaObject();
func (this *QLineEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QLineEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "metaObject", args)
  }

}

  // proto:  void QLineEdit::del();
func (this *QLineEdit) del(args ...interface{}) () {
  // del()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit3delEv
    // invoke: void del()
    C._ZN9QLineEdit3delEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "del", args)
  }

}

  // proto:  bool QLineEdit::isModified();
func (this *QLineEdit) isModified(args ...interface{}) () {
  // isModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit10isModifiedEv
    // invoke: bool isModified()
    C._ZNK9QLineEdit10isModifiedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "isModified", args)
  }

}

  // proto:  void QLineEdit::cursorWordForward(bool mark);
func (this *QLineEdit) cursorWordForward(args ...interface{}) () {
  // cursorWordForward(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit17cursorWordForwardEb
    // invoke: void cursorWordForward(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit17cursorWordForwardEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordForward", args)
  }

}

  // proto:  void QLineEdit::selectAll();
func (this *QLineEdit) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit9selectAllEv
    // invoke: void selectAll()
    C._ZN9QLineEdit9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "selectAll", args)
  }

}

  // proto:  void QLineEdit::setSelection(int , int );
func (this *QLineEdit) setSelection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setSelectionEii
    // invoke: void setSelection(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QLineEdit12setSelectionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "setSelection", args)
  }

}

  // proto:  void QLineEdit::setCompleter(QCompleter * completer);
func (this *QLineEdit) setCompleter(args ...interface{}) () {
  // setCompleter(class QCompleter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCompleter{}) // "QCompleter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setCompleterEP10QCompleter
    // invoke: void setCompleter(class QCompleter *)
    var arg0 = args[0].(QCompleter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit12setCompleterEP10QCompleter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setCompleter", args)
  }

}

  // proto:  void QLineEdit::setMaxLength(int );
func (this *QLineEdit) setMaxLength(args ...interface{}) () {
  // setMaxLength(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setMaxLengthEi
    // invoke: void setMaxLength(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit12setMaxLengthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setMaxLength", args)
  }

}

  // proto:  void QLineEdit::~QLineEdit();
func (this *QLineEdit) FreeQLineEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLineEdit", "~QLineEdit", args)
  }

}

  // proto:  void QLineEdit::setReadOnly(bool );
func (this *QLineEdit) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setReadOnly", args)
  }

}

  // proto:  QString QLineEdit::displayText();
func (this *QLineEdit) displayText(args ...interface{}) () {
  // displayText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit11displayTextEv
    // invoke: QString displayText()
    C._ZNK9QLineEdit11displayTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "displayText", args)
  }

}

  // proto:  void QLineEdit::setFrame(bool );
func (this *QLineEdit) setFrame(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit8setFrameEb
    // invoke: void setFrame(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit8setFrameEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setFrame", args)
  }

}

  // proto:  bool QLineEdit::hasAcceptableInput();
func (this *QLineEdit) hasAcceptableInput(args ...interface{}) () {
  // hasAcceptableInput()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit18hasAcceptableInputEv
    // invoke: bool hasAcceptableInput()
    C._ZNK9QLineEdit18hasAcceptableInputEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "hasAcceptableInput", args)
  }

}

  // proto:  bool QLineEdit::hasFrame();
func (this *QLineEdit) hasFrame(args ...interface{}) () {
  // hasFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit8hasFrameEv
    // invoke: bool hasFrame()
    C._ZNK9QLineEdit8hasFrameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "hasFrame", args)
  }

}

  // proto:  int QLineEdit::cursorPosition();
func (this *QLineEdit) cursorPosition(args ...interface{}) () {
  // cursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit14cursorPositionEv
    // invoke: int cursorPosition()
    C._ZNK9QLineEdit14cursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorPosition", args)
  }

}

  // proto:  void QLineEdit::cursorWordBackward(bool mark);
func (this *QLineEdit) cursorWordBackward(args ...interface{}) () {
  // cursorWordBackward(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit18cursorWordBackwardEb
    // invoke: void cursorWordBackward(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit18cursorWordBackwardEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordBackward", args)
  }

}

  // proto:  bool QLineEdit::dragEnabled();
func (this *QLineEdit) dragEnabled(args ...interface{}) () {
  // dragEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit11dragEnabledEv
    // invoke: bool dragEnabled()
    C._ZNK9QLineEdit11dragEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "dragEnabled", args)
  }

}

  // proto:  QSize QLineEdit::sizeHint();
func (this *QLineEdit) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK9QLineEdit8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "sizeHint", args)
  }

}

  // proto:  void QLineEdit::paste();
func (this *QLineEdit) paste(args ...interface{}) () {
  // paste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit5pasteEv
    // invoke: void paste()
    C._ZN9QLineEdit5pasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "paste", args)
  }

}

  // proto:  void QLineEdit::setValidator(const QValidator * );
func (this *QLineEdit) setValidator(args ...interface{}) () {
  // setValidator(const class QValidator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QValidator{}) // "const QValidator *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setValidatorEPK10QValidator
    // invoke: void setValidator(const class QValidator *)
    var arg0 = args[0].(QValidator).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit12setValidatorEPK10QValidator(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setValidator", args)
  }

}

  // proto:  QCompleter * QLineEdit::completer();
func (this *QLineEdit) completer(args ...interface{}) () {
  // completer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit9completerEv
    // invoke: QCompleter * completer()
    C._ZNK9QLineEdit9completerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "completer", args)
  }

}

  // proto:  QMargins QLineEdit::textMargins();
func (this *QLineEdit) textMargins(args ...interface{}) () {
  // textMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit11textMarginsEv
    // invoke: QMargins textMargins()
    C._ZNK9QLineEdit11textMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "textMargins", args)
  }

}

  // proto:  void QLineEdit::setClearButtonEnabled(bool enable);
func (this *QLineEdit) setClearButtonEnabled(args ...interface{}) () {
  // setClearButtonEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit21setClearButtonEnabledEb
    // invoke: void setClearButtonEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit21setClearButtonEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setClearButtonEnabled", args)
  }

}

  // proto:  QString QLineEdit::selectedText();
func (this *QLineEdit) selectedText(args ...interface{}) () {
  // selectedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit12selectedTextEv
    // invoke: QString selectedText()
    C._ZNK9QLineEdit12selectedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "selectedText", args)
  }

}

  // proto:  void QLineEdit::clear();
func (this *QLineEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit5clearEv
    // invoke: void clear()
    C._ZN9QLineEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "clear", args)
  }

}

  // proto:  void QLineEdit::copy();
func (this *QLineEdit) copy(args ...interface{}) () {
  // copy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit4copyEv
    // invoke: void copy()
    C._ZNK9QLineEdit4copyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "copy", args)
  }

}

  // proto:  bool QLineEdit::isUndoAvailable();
func (this *QLineEdit) isUndoAvailable(args ...interface{}) () {
  // isUndoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15isUndoAvailableEv
    // invoke: bool isUndoAvailable()
    C._ZNK9QLineEdit15isUndoAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "isUndoAvailable", args)
  }

}

  // proto:  void QLineEdit::undo();
func (this *QLineEdit) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit4undoEv
    // invoke: void undo()
    C._ZN9QLineEdit4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "undo", args)
  }

}

  // proto:  bool QLineEdit::isClearButtonEnabled();
func (this *QLineEdit) isClearButtonEnabled(args ...interface{}) () {
  // isClearButtonEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit20isClearButtonEnabledEv
    // invoke: bool isClearButtonEnabled()
    C._ZNK9QLineEdit20isClearButtonEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "isClearButtonEnabled", args)
  }

}

  // proto:  void QLineEdit::end(bool mark);
func (this *QLineEdit) end(args ...interface{}) () {
  // end(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit3endEb
    // invoke: void end(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit3endEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "end", args)
  }

}

  // proto:  void QLineEdit::setDragEnabled(bool b);
func (this *QLineEdit) setDragEnabled(args ...interface{}) () {
  // setDragEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit14setDragEnabledEb
    // invoke: void setDragEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit14setDragEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setDragEnabled", args)
  }

}

  // proto:  void QLineEdit::backspace();
func (this *QLineEdit) backspace(args ...interface{}) () {
  // backspace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit9backspaceEv
    // invoke: void backspace()
    C._ZN9QLineEdit9backspaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "backspace", args)
  }

}

  // proto:  void QLineEdit::redo();
func (this *QLineEdit) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit4redoEv
    // invoke: void redo()
    C._ZN9QLineEdit4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "redo", args)
  }

}

  // proto:  void QLineEdit::setInputMask(const QString & inputMask);
func (this *QLineEdit) setInputMask(args ...interface{}) () {
  // setInputMask(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setInputMaskERK7QString
    // invoke: void setInputMask(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit12setInputMaskERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setInputMask", args)
  }

}

  // proto:  void QLineEdit::getTextMargins(int * left, int * top, int * right, int * bottom);
func (this *QLineEdit) getTextMargins(args ...interface{}) () {
  // getTextMargins(int *, int *, int *, int *)
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
    // invoke: _ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_
    // invoke: void getTextMargins(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLineEdit", "getTextMargins", args)
  }

}

  // proto:  bool QLineEdit::isReadOnly();
func (this *QLineEdit) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit10isReadOnlyEv
    // invoke: bool isReadOnly()
    C._ZNK9QLineEdit10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "isReadOnly", args)
  }

}

// <= body block end

