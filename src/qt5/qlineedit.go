package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QLineEdit::EchoMode QLineEdit::echoMode();
extern void _ZNK9QLineEdit8echoModeEv(void* qthis); // 4
  // proto:  QString QLineEdit::text();
extern void _ZNK9QLineEdit4textEv(void* qthis); // 4
  // proto:  void QLineEdit::setDragEnabled(bool b);
extern void _ZN9QLineEdit14setDragEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QLineEdit::hasFrame();
extern void _ZNK9QLineEdit8hasFrameEv(void* qthis); // 4
  // proto:  bool QLineEdit::hasSelectedText();
extern void _ZNK9QLineEdit15hasSelectedTextEv(void* qthis); // 4
  // proto:  Qt::Alignment QLineEdit::alignment();
extern void _ZNK9QLineEdit9alignmentEv(void* qthis); // 4
  // proto:  void QLineEdit::cut();
extern void _ZN9QLineEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QLineEdit::createStandardContextMenu();
extern void _ZN9QLineEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  QString QLineEdit::displayText();
extern void _ZNK9QLineEdit11displayTextEv(void* qthis); // 4
  // proto:  void QLineEdit::QLineEdit(const QString & , QWidget * parent);
extern void _ZN9QLineEditC2ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QLineEdit::QLineEdit(QWidget * parent);
extern void _ZN9QLineEditC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QLineEdit::cursorBackward(bool mark, int steps);
extern void _ZN9QLineEdit14cursorBackwardEbi(void* qthis, bool arg0, int32_t arg1); // 4
  // proto:  int QLineEdit::cursorPositionAt(const QPoint & pos);
extern void _ZN9QLineEdit16cursorPositionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::getTextMargins(int * left, int * top, int * right, int * bottom);
extern void _ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  int QLineEdit::maxLength();
extern void _ZNK9QLineEdit9maxLengthEv(void* qthis); // 4
  // proto:  void QLineEdit::insert(const QString & );
extern void _ZN9QLineEdit6insertERK7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QLineEdit::metaObject();
extern void _ZNK9QLineEdit10metaObjectEv(void* qthis); // 4
  // proto:  bool QLineEdit::hasAcceptableInput();
extern void _ZNK9QLineEdit18hasAcceptableInputEv(void* qthis); // 4
  // proto:  bool QLineEdit::isClearButtonEnabled();
extern void _ZNK9QLineEdit20isClearButtonEnabledEv(void* qthis); // 4
  // proto:  bool QLineEdit::isUndoAvailable();
extern void _ZNK9QLineEdit15isUndoAvailableEv(void* qthis); // 4
  // proto:  int QLineEdit::selectionStart();
extern void _ZNK9QLineEdit14selectionStartEv(void* qthis); // 4
  // proto:  void QLineEdit::cursorWordForward(bool mark);
extern void _ZN9QLineEdit17cursorWordForwardEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::home(bool mark);
extern void _ZN9QLineEdit4homeEb(void* qthis, bool arg0); // 4
  // proto:  bool QLineEdit::event(QEvent * );
extern void _ZN9QLineEdit5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::deselect();
extern void _ZN9QLineEdit8deselectEv(void* qthis); // 4
  // proto:  QString QLineEdit::inputMask();
extern void _ZNK9QLineEdit9inputMaskEv(void* qthis); // 4
  // proto:  void QLineEdit::backspace();
extern void _ZN9QLineEdit9backspaceEv(void* qthis); // 4
  // proto:  void QLineEdit::setModified(bool );
extern void _ZN9QLineEdit11setModifiedEb(void* qthis, bool arg0); // 4
  // proto:  QString QLineEdit::selectedText();
extern void _ZNK9QLineEdit12selectedTextEv(void* qthis); // 4
  // proto:  void QLineEdit::undo();
extern void _ZN9QLineEdit4undoEv(void* qthis); // 4
  // proto:  void QLineEdit::paste();
extern void _ZN9QLineEdit5pasteEv(void* qthis); // 4
  // proto:  QSize QLineEdit::sizeHint();
extern void _ZNK9QLineEdit8sizeHintEv(void* qthis); // 4
  // proto:  void QLineEdit::~QLineEdit();
extern void _ZN9QLineEditD2Ev(void* qthis); // 4
  // proto:  bool QLineEdit::isRedoAvailable();
extern void _ZNK9QLineEdit15isRedoAvailableEv(void* qthis); // 4
  // proto:  void QLineEdit::selectAll();
extern void _ZN9QLineEdit9selectAllEv(void* qthis); // 4
  // proto:  void QLineEdit::setClearButtonEnabled(bool enable);
extern void _ZN9QLineEdit21setClearButtonEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::redo();
extern void _ZN9QLineEdit4redoEv(void* qthis); // 4
  // proto:  int QLineEdit::cursorPosition();
extern void _ZNK9QLineEdit14cursorPositionEv(void* qthis); // 4
  // proto:  bool QLineEdit::isModified();
extern void _ZNK9QLineEdit10isModifiedEv(void* qthis); // 4
  // proto:  bool QLineEdit::isReadOnly();
extern void _ZNK9QLineEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QLineEdit::setCompleter(QCompleter * completer);
extern void _ZN9QLineEdit12setCompleterEP10QCompleter(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setInputMask(const QString & inputMask);
extern void _ZN9QLineEdit12setInputMaskERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setValidator(const QValidator * );
extern void _ZN9QLineEdit12setValidatorEPK10QValidator(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::copy();
extern void _ZNK9QLineEdit4copyEv(void* qthis); // 4
  // proto:  QCompleter * QLineEdit::completer();
extern void _ZNK9QLineEdit9completerEv(void* qthis); // 4
  // proto:  void QLineEdit::setCursorPosition(int );
extern void _ZN9QLineEdit17setCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QLineEdit::minimumSizeHint();
extern void _ZNK9QLineEdit15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QLineEdit::setPlaceholderText(const QString & );
extern void _ZN9QLineEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QLineEdit::dragEnabled();
extern void _ZNK9QLineEdit11dragEnabledEv(void* qthis); // 4
  // proto:  void QLineEdit::setMaxLength(int );
extern void _ZN9QLineEdit12setMaxLengthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLineEdit::del();
extern void _ZN9QLineEdit3delEv(void* qthis); // 4
  // proto:  void QLineEdit::clear();
extern void _ZN9QLineEdit5clearEv(void* qthis); // 4
  // proto:  void QLineEdit::cursorWordBackward(bool mark);
extern void _ZN9QLineEdit18cursorWordBackwardEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::end(bool mark);
extern void _ZN9QLineEdit3endEb(void* qthis, bool arg0); // 4
  // proto:  QMargins QLineEdit::textMargins();
extern void _ZNK9QLineEdit11textMarginsEv(void* qthis); // 4
  // proto:  QString QLineEdit::placeholderText();
extern void _ZNK9QLineEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QLineEdit::setTextMargins(const QMargins & margins);
extern void _ZN9QLineEdit14setTextMarginsERK8QMargins(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setTextMargins(int left, int top, int right, int bottom);
extern void _ZN9QLineEdit14setTextMarginsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QLineEdit::setSelection(int , int );
extern void _ZN9QLineEdit12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QLineEdit::setFrame(bool );
extern void _ZN9QLineEdit8setFrameEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::cursorForward(bool mark, int steps);
extern void _ZN9QLineEdit13cursorForwardEbi(void* qthis, bool arg0, int32_t arg1); // 4
  // proto:  void QLineEdit::setText(const QString & );
extern void _ZN9QLineEdit7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::CursorMoveStyle QLineEdit::cursorMoveStyle();
extern void _ZNK9QLineEdit15cursorMoveStyleEv(void* qthis); // 4
  // proto:  void QLineEdit::setReadOnly(bool );
extern void _ZN9QLineEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  const QValidator * QLineEdit::validator();
extern void _ZNK9QLineEdit9validatorEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _textEdited QLineEdit_textEdited_signal;
//  _returnPressed QLineEdit_returnPressed_signal;
//  _selectionChanged QLineEdit_selectionChanged_signal;
//  _cursorPositionChanged QLineEdit_cursorPositionChanged_signal;
//  _editingFinished QLineEdit_editingFinished_signal;
//  _textChanged QLineEdit_textChanged_signal;
}

// echoMode()
func (this *QLineEdit) echoMode(args ...interface{}) () {
  // echoMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit8echoModeEv
    // invoke: QLineEdit::EchoMode echoMode()
    C._ZNK9QLineEdit8echoModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "echoMode", args)
  }

}

// text()
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

// setDragEnabled(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit14setDragEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setDragEnabled", args)
  }

}

// hasFrame()
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

// hasSelectedText()
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

// alignment()
func (this *QLineEdit) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit9alignmentEv
    // invoke: Qt::Alignment alignment()
    C._ZNK9QLineEdit9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "alignment", args)
  }

}

// cut()
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

// createStandardContextMenu()
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

// displayText()
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

// QLineEdit(const class QString &, class QWidget *)
func NewQLineEdit(args ...interface{}) QLineEdit {
  // QLineEdit(const class QString &, class QWidget *)
  // QLineEdit(class QWidget *)
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
    // invoke: _ZN9QLineEditC1ERK7QStringP7QWidget
    // invoke: void QLineEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QLineEditC2ERK7QStringP7QWidget(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN9QLineEditC1EP7QWidget
    // invoke: void QLineEdit(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QLineEditC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "QLineEdit", args)
  }

  return QLineEdit{}
}

// cursorBackward(_Bool, int)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QLineEdit14cursorBackwardEbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorBackward", args)
  }

}

// cursorPositionAt(const class QPoint &)
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

// getTextMargins(int *, int *, int *, int *)
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

// maxLength()
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

// insert(const class QString &)
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

// metaObject()
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

// hasAcceptableInput()
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

// isClearButtonEnabled()
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

// isUndoAvailable()
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

// selectionStart()
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

// cursorWordForward(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit17cursorWordForwardEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordForward", args)
  }

}

// home(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit4homeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "home", args)
  }

}

// event(class QEvent *)
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

// deselect()
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

// inputMask()
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

// backspace()
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

// setModified(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit11setModifiedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setModified", args)
  }

}

// selectedText()
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

// undo()
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

// paste()
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

// sizeHint()
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

// ~QLineEdit()
func (this *QLineEdit) FreeQLineEdit(args ...interface{}) () {
  // ~QLineEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEditD0Ev
    // invoke: void ~QLineEdit()
    C._ZN9QLineEditD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "~QLineEdit", args)
  }

}

// isRedoAvailable()
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

// selectAll()
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

// setClearButtonEnabled(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit21setClearButtonEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setClearButtonEnabled", args)
  }

}

// redo()
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

// cursorPosition()
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

// isModified()
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

// isReadOnly()
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

// setCompleter(class QCompleter *)
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

// setInputMask(const class QString &)
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

// setValidator(const class QValidator *)
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

// copy()
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

// completer()
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

// setCursorPosition(int)
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

// minimumSizeHint()
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

// setPlaceholderText(const class QString &)
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

// dragEnabled()
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

// setMaxLength(int)
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

// del()
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

// clear()
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

// cursorWordBackward(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit18cursorWordBackwardEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordBackward", args)
  }

}

// end(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit3endEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "end", args)
  }

}

// textMargins()
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

// placeholderText()
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

// setTextMargins(const class QMargins &)
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

// setSelection(int, int)
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

// setFrame(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit8setFrameEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setFrame", args)
  }

}

// cursorForward(_Bool, int)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QLineEdit13cursorForwardEbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorForward", args)
  }

}

// setText(const class QString &)
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

// cursorMoveStyle()
func (this *QLineEdit) cursorMoveStyle(args ...interface{}) () {
  // cursorMoveStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit15cursorMoveStyleEv
    // invoke: Qt::CursorMoveStyle cursorMoveStyle()
    C._ZNK9QLineEdit15cursorMoveStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorMoveStyle", args)
  }

}

// setReadOnly(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QLineEdit11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setReadOnly", args)
  }

}

// validator()
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

// <= body block end

