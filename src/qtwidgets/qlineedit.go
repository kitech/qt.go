package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  QLineEdit::EchoMode QLineEdit::echoMode();
extern void C_ZNK9QLineEdit8echoModeEv(void* qthis); // 4
  // proto:  QString QLineEdit::text();
extern void* C_ZNK9QLineEdit4textEv(void* qthis); // 4
  // proto:  void QLineEdit::setDragEnabled(bool b);
extern void C_ZN9QLineEdit14setDragEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QLineEdit::hasFrame();
extern bool C_ZNK9QLineEdit8hasFrameEv(void* qthis); // 4
  // proto:  bool QLineEdit::hasSelectedText();
extern bool C_ZNK9QLineEdit15hasSelectedTextEv(void* qthis); // 4
  // proto:  Qt::Alignment QLineEdit::alignment();
extern void C_ZNK9QLineEdit9alignmentEv(void* qthis); // 4
  // proto:  void QLineEdit::cut();
extern void C_ZN9QLineEdit3cutEv(void* qthis); // 4
  // proto:  QMenu * QLineEdit::createStandardContextMenu();
extern void* C_ZN9QLineEdit25createStandardContextMenuEv(void* qthis); // 4
  // proto:  QString QLineEdit::displayText();
extern void* C_ZNK9QLineEdit11displayTextEv(void* qthis); // 4
  // proto:  void QLineEdit::QLineEdit(const QString & , QWidget * parent);
extern void* C_ZN9QLineEditC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QLineEdit::QLineEdit(QWidget * parent);
extern void* C_ZN9QLineEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QLineEdit::cursorBackward(bool mark, int steps);
extern void C_ZN9QLineEdit14cursorBackwardEbi(void* qthis, bool arg0, int32_t arg1); // 4
  // proto:  int QLineEdit::cursorPositionAt(const QPoint & pos);
extern int32_t C_ZN9QLineEdit16cursorPositionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::getTextMargins(int * left, int * top, int * right, int * bottom);
extern void C_ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  int QLineEdit::maxLength();
extern int32_t C_ZNK9QLineEdit9maxLengthEv(void* qthis); // 4
  // proto:  const QMetaObject * QLineEdit::metaObject();
extern void C_ZNK9QLineEdit10metaObjectEv(void* qthis); // 4
  // proto:  bool QLineEdit::hasAcceptableInput();
extern bool C_ZNK9QLineEdit18hasAcceptableInputEv(void* qthis); // 4
  // proto:  bool QLineEdit::isClearButtonEnabled();
extern bool C_ZNK9QLineEdit20isClearButtonEnabledEv(void* qthis); // 4
  // proto:  bool QLineEdit::isUndoAvailable();
extern bool C_ZNK9QLineEdit15isUndoAvailableEv(void* qthis); // 4
  // proto:  int QLineEdit::selectionStart();
extern int32_t C_ZNK9QLineEdit14selectionStartEv(void* qthis); // 4
  // proto:  void QLineEdit::cursorWordForward(bool mark);
extern void C_ZN9QLineEdit17cursorWordForwardEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::home(bool mark);
extern void C_ZN9QLineEdit4homeEb(void* qthis, bool arg0); // 4
  // proto:  bool QLineEdit::event(QEvent * );
extern bool C_ZN9QLineEdit5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::deselect();
extern void C_ZN9QLineEdit8deselectEv(void* qthis); // 4
  // proto:  QString QLineEdit::inputMask();
extern void* C_ZNK9QLineEdit9inputMaskEv(void* qthis); // 4
  // proto:  void QLineEdit::backspace();
extern void C_ZN9QLineEdit9backspaceEv(void* qthis); // 4
  // proto:  void QLineEdit::setModified(bool );
extern void C_ZN9QLineEdit11setModifiedEb(void* qthis, bool arg0); // 4
  // proto:  QString QLineEdit::selectedText();
extern void* C_ZNK9QLineEdit12selectedTextEv(void* qthis); // 4
  // proto:  void QLineEdit::undo();
extern void C_ZN9QLineEdit4undoEv(void* qthis); // 4
  // proto:  void QLineEdit::paste();
extern void C_ZN9QLineEdit5pasteEv(void* qthis); // 4
  // proto:  QSize QLineEdit::sizeHint();
extern void* C_ZNK9QLineEdit8sizeHintEv(void* qthis); // 4
  // proto:  void QLineEdit::~QLineEdit();
extern void C_ZN9QLineEditD2Ev(void* qthis); // 4
  // proto:  bool QLineEdit::isRedoAvailable();
extern bool C_ZNK9QLineEdit15isRedoAvailableEv(void* qthis); // 4
  // proto:  void QLineEdit::selectAll();
extern void C_ZN9QLineEdit9selectAllEv(void* qthis); // 4
  // proto:  void QLineEdit::setClearButtonEnabled(bool enable);
extern void C_ZN9QLineEdit21setClearButtonEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::redo();
extern void C_ZN9QLineEdit4redoEv(void* qthis); // 4
  // proto:  int QLineEdit::cursorPosition();
extern int32_t C_ZNK9QLineEdit14cursorPositionEv(void* qthis); // 4
  // proto:  bool QLineEdit::isModified();
extern bool C_ZNK9QLineEdit10isModifiedEv(void* qthis); // 4
  // proto:  bool QLineEdit::isReadOnly();
extern bool C_ZNK9QLineEdit10isReadOnlyEv(void* qthis); // 4
  // proto:  void QLineEdit::setCompleter(QCompleter * completer);
extern void C_ZN9QLineEdit12setCompleterEP10QCompleter(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setInputMask(const QString & inputMask);
extern void C_ZN9QLineEdit12setInputMaskERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setValidator(const QValidator * );
extern void C_ZN9QLineEdit12setValidatorEPK10QValidator(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::copy();
extern void C_ZNK9QLineEdit4copyEv(void* qthis); // 4
  // proto:  QCompleter * QLineEdit::completer();
extern void* C_ZNK9QLineEdit9completerEv(void* qthis); // 4
  // proto:  void QLineEdit::setCursorPosition(int );
extern void C_ZN9QLineEdit17setCursorPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QLineEdit::minimumSizeHint();
extern void* C_ZNK9QLineEdit15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QLineEdit::setPlaceholderText(const QString & );
extern void C_ZN9QLineEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QLineEdit::dragEnabled();
extern bool C_ZNK9QLineEdit11dragEnabledEv(void* qthis); // 4
  // proto:  void QLineEdit::setMaxLength(int );
extern void C_ZN9QLineEdit12setMaxLengthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLineEdit::del();
extern void C_ZN9QLineEdit3delEv(void* qthis); // 4
  // proto:  void QLineEdit::clear();
extern void C_ZN9QLineEdit5clearEv(void* qthis); // 4
  // proto:  void QLineEdit::cursorWordBackward(bool mark);
extern void C_ZN9QLineEdit18cursorWordBackwardEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::end(bool mark);
extern void C_ZN9QLineEdit3endEb(void* qthis, bool arg0); // 4
  // proto:  QMargins QLineEdit::textMargins();
extern void* C_ZNK9QLineEdit11textMarginsEv(void* qthis); // 4
  // proto:  QString QLineEdit::placeholderText();
extern void* C_ZNK9QLineEdit15placeholderTextEv(void* qthis); // 4
  // proto:  void QLineEdit::setTextMargins(const QMargins & margins);
extern void C_ZN9QLineEdit14setTextMarginsERK8QMargins(void* qthis, void* arg0); // 4
  // proto:  void QLineEdit::setTextMargins(int left, int top, int right, int bottom);
extern void C_ZN9QLineEdit14setTextMarginsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QLineEdit::setSelection(int , int );
extern void C_ZN9QLineEdit12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QLineEdit::setFrame(bool );
extern void C_ZN9QLineEdit8setFrameEb(void* qthis, bool arg0); // 4
  // proto:  void QLineEdit::cursorForward(bool mark, int steps);
extern void C_ZN9QLineEdit13cursorForwardEbi(void* qthis, bool arg0, int32_t arg1); // 4
  // proto:  void QLineEdit::setText(const QString & );
extern void C_ZN9QLineEdit7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::CursorMoveStyle QLineEdit::cursorMoveStyle();
extern void C_ZNK9QLineEdit15cursorMoveStyleEv(void* qthis); // 4
  // proto:  void QLineEdit::setReadOnly(bool );
extern void C_ZN9QLineEdit11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  const QValidator * QLineEdit::validator();
extern void* C_ZNK9QLineEdit9validatorEv(void* qthis); // 4
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

// class sizeof(QLineEdit)=1
type QLineEdit struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _textEdited QLineEdit_textEdited_signal;
//  _returnPressed QLineEdit_returnPressed_signal;
//  _selectionChanged QLineEdit_selectionChanged_signal;
//  _cursorPositionChanged QLineEdit_cursorPositionChanged_signal;
//  _editingFinished QLineEdit_editingFinished_signal;
//  _textChanged QLineEdit_textChanged_signal;
}

// echoMode()
func (this *QLineEdit) Echomode(args ...interface{}) () {
  // echoMode()
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
    // invoke: _ZNK9QLineEdit8echoModeEv
    // invoke: QLineEdit::EchoMode echoMode()
    C.C_ZNK9QLineEdit8echoModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "echoMode", args)
  }

  return
}

// text()
func (this *QLineEdit) Text(args ...interface{}) (ret interface{}) {
  // text()
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
    // invoke: _ZNK9QLineEdit4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK9QLineEdit4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "text", args)
  }

  return
}

// setDragEnabled(_Bool)
func (this *QLineEdit) Setdragenabled(args ...interface{}) () {
  // setDragEnabled(_Bool)
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
    // invoke: _ZN9QLineEdit14setDragEnabledEb
    // invoke: void setDragEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit14setDragEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setDragEnabled", args)
  }

  return
}

// hasFrame()
func (this *QLineEdit) Hasframe(args ...interface{}) (ret interface{}) {
  // hasFrame()
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
    // invoke: _ZNK9QLineEdit8hasFrameEv
    // invoke: bool hasFrame()
    var ret0 = C.C_ZNK9QLineEdit8hasFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "hasFrame", args)
  }

  return
}

// hasSelectedText()
func (this *QLineEdit) Hasselectedtext(args ...interface{}) (ret interface{}) {
  // hasSelectedText()
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
    // invoke: _ZNK9QLineEdit15hasSelectedTextEv
    // invoke: bool hasSelectedText()
    var ret0 = C.C_ZNK9QLineEdit15hasSelectedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "hasSelectedText", args)
  }

  return
}

// alignment()
func (this *QLineEdit) Alignment(args ...interface{}) () {
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
    // invoke: _ZNK9QLineEdit9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK9QLineEdit9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "alignment", args)
  }

  return
}

// cut()
func (this *QLineEdit) Cut(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit3cutEv
    // invoke: void cut()
    C.C_ZN9QLineEdit3cutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "cut", args)
  }

  return
}

// createStandardContextMenu()
func (this *QLineEdit) Createstandardcontextmenu(args ...interface{}) (ret interface{}) {
  // createStandardContextMenu()
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
    // invoke: _ZN9QLineEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    var ret0 = C.C_ZN9QLineEdit25createStandardContextMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "createStandardContextMenu", args)
  }

  return
}

// displayText()
func (this *QLineEdit) Displaytext(args ...interface{}) (ret interface{}) {
  // displayText()
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
    // invoke: _ZNK9QLineEdit11displayTextEv
    // invoke: QString displayText()
    var ret0 = C.C_ZNK9QLineEdit11displayTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "displayText", args)
  }

  return
}

// QLineEdit(const class QString &, class QWidget *)
func NewQLineEdit(args ...interface{}) *QLineEdit {
  // QLineEdit(const class QString &, class QWidget *)
  // QLineEdit(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEditC1ERK7QStringP7QWidget
    // invoke: void QLineEdit(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QLineEditC2ERK7QStringP7QWidget(arg0, arg1)
    return &QLineEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QLineEditC1EP7QWidget
    // invoke: void QLineEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QLineEditC2EP7QWidget(arg0)
    return &QLineEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLineEdit", "QLineEdit", args)
  }

  return nil // QLineEdit{Qclsinst:qthis}
}

// cursorBackward(_Bool, int)
func (this *QLineEdit) Cursorbackward(args ...interface{}) () {
  // cursorBackward(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit14cursorBackwardEbi
    // invoke: void cursorBackward(_Bool, int)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QLineEdit14cursorBackwardEbi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorBackward", args)
  }

  return
}

// cursorPositionAt(const class QPoint &)
func (this *QLineEdit) Cursorpositionat(args ...interface{}) (ret interface{}) {
  // cursorPositionAt(const class QPoint &)
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
    // invoke: _ZN9QLineEdit16cursorPositionAtERK6QPoint
    // invoke: int cursorPositionAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QLineEdit16cursorPositionAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorPositionAt", args)
  }

  return
}

// getTextMargins(int *, int *, int *, int *)
func (this *QLineEdit) Gettextmargins(args ...interface{}) () {
  // getTextMargins(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_
    // invoke: void getTextMargins(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLineEdit", "getTextMargins", args)
  }

  return
}

// maxLength()
func (this *QLineEdit) Maxlength(args ...interface{}) (ret interface{}) {
  // maxLength()
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
    // invoke: _ZNK9QLineEdit9maxLengthEv
    // invoke: int maxLength()
    var ret0 = C.C_ZNK9QLineEdit9maxLengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "maxLength", args)
  }

  return
}

// metaObject()
func (this *QLineEdit) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK9QLineEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QLineEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "metaObject", args)
  }

  return
}

// hasAcceptableInput()
func (this *QLineEdit) Hasacceptableinput(args ...interface{}) (ret interface{}) {
  // hasAcceptableInput()
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
    // invoke: _ZNK9QLineEdit18hasAcceptableInputEv
    // invoke: bool hasAcceptableInput()
    var ret0 = C.C_ZNK9QLineEdit18hasAcceptableInputEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "hasAcceptableInput", args)
  }

  return
}

// isClearButtonEnabled()
func (this *QLineEdit) Isclearbuttonenabled(args ...interface{}) (ret interface{}) {
  // isClearButtonEnabled()
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
    // invoke: _ZNK9QLineEdit20isClearButtonEnabledEv
    // invoke: bool isClearButtonEnabled()
    var ret0 = C.C_ZNK9QLineEdit20isClearButtonEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "isClearButtonEnabled", args)
  }

  return
}

// isUndoAvailable()
func (this *QLineEdit) Isundoavailable(args ...interface{}) (ret interface{}) {
  // isUndoAvailable()
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
    // invoke: _ZNK9QLineEdit15isUndoAvailableEv
    // invoke: bool isUndoAvailable()
    var ret0 = C.C_ZNK9QLineEdit15isUndoAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "isUndoAvailable", args)
  }

  return
}

// selectionStart()
func (this *QLineEdit) Selectionstart(args ...interface{}) (ret interface{}) {
  // selectionStart()
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
    // invoke: _ZNK9QLineEdit14selectionStartEv
    // invoke: int selectionStart()
    var ret0 = C.C_ZNK9QLineEdit14selectionStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "selectionStart", args)
  }

  return
}

// cursorWordForward(_Bool)
func (this *QLineEdit) Cursorwordforward(args ...interface{}) () {
  // cursorWordForward(_Bool)
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
    // invoke: _ZN9QLineEdit17cursorWordForwardEb
    // invoke: void cursorWordForward(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit17cursorWordForwardEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordForward", args)
  }

  return
}

// home(_Bool)
func (this *QLineEdit) Home(args ...interface{}) () {
  // home(_Bool)
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
    // invoke: _ZN9QLineEdit4homeEb
    // invoke: void home(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit4homeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "home", args)
  }

  return
}

// event(class QEvent *)
func (this *QLineEdit) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QLineEdit5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "event", args)
  }

  return
}

// deselect()
func (this *QLineEdit) Deselect(args ...interface{}) () {
  // deselect()
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
    // invoke: _ZN9QLineEdit8deselectEv
    // invoke: void deselect()
    C.C_ZN9QLineEdit8deselectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "deselect", args)
  }

  return
}

// inputMask()
func (this *QLineEdit) Inputmask(args ...interface{}) (ret interface{}) {
  // inputMask()
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
    // invoke: _ZNK9QLineEdit9inputMaskEv
    // invoke: QString inputMask()
    var ret0 = C.C_ZNK9QLineEdit9inputMaskEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "inputMask", args)
  }

  return
}

// backspace()
func (this *QLineEdit) Backspace(args ...interface{}) () {
  // backspace()
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
    // invoke: _ZN9QLineEdit9backspaceEv
    // invoke: void backspace()
    C.C_ZN9QLineEdit9backspaceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "backspace", args)
  }

  return
}

// setModified(_Bool)
func (this *QLineEdit) Setmodified(args ...interface{}) () {
  // setModified(_Bool)
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
    // invoke: _ZN9QLineEdit11setModifiedEb
    // invoke: void setModified(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit11setModifiedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setModified", args)
  }

  return
}

// selectedText()
func (this *QLineEdit) Selectedtext(args ...interface{}) (ret interface{}) {
  // selectedText()
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
    // invoke: _ZNK9QLineEdit12selectedTextEv
    // invoke: QString selectedText()
    var ret0 = C.C_ZNK9QLineEdit12selectedTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "selectedText", args)
  }

  return
}

// undo()
func (this *QLineEdit) Undo(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit4undoEv
    // invoke: void undo()
    C.C_ZN9QLineEdit4undoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "undo", args)
  }

  return
}

// paste()
func (this *QLineEdit) Paste(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit5pasteEv
    // invoke: void paste()
    C.C_ZN9QLineEdit5pasteEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "paste", args)
  }

  return
}

// sizeHint()
func (this *QLineEdit) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
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
    // invoke: _ZNK9QLineEdit8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK9QLineEdit8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "sizeHint", args)
  }

  return
}

// ~QLineEdit()
func (this *QLineEdit) Freeqlineedit(args ...interface{}) () {
  // ~QLineEdit()
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
    // invoke: _ZN9QLineEditD0Ev
    // invoke: void ~QLineEdit()
    C.C_ZN9QLineEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "~QLineEdit", args)
  }

  return
}

// isRedoAvailable()
func (this *QLineEdit) Isredoavailable(args ...interface{}) (ret interface{}) {
  // isRedoAvailable()
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
    // invoke: _ZNK9QLineEdit15isRedoAvailableEv
    // invoke: bool isRedoAvailable()
    var ret0 = C.C_ZNK9QLineEdit15isRedoAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "isRedoAvailable", args)
  }

  return
}

// selectAll()
func (this *QLineEdit) Selectall(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit9selectAllEv
    // invoke: void selectAll()
    C.C_ZN9QLineEdit9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "selectAll", args)
  }

  return
}

// setClearButtonEnabled(_Bool)
func (this *QLineEdit) Setclearbuttonenabled(args ...interface{}) () {
  // setClearButtonEnabled(_Bool)
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
    // invoke: _ZN9QLineEdit21setClearButtonEnabledEb
    // invoke: void setClearButtonEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit21setClearButtonEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setClearButtonEnabled", args)
  }

  return
}

// redo()
func (this *QLineEdit) Redo(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit4redoEv
    // invoke: void redo()
    C.C_ZN9QLineEdit4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "redo", args)
  }

  return
}

// cursorPosition()
func (this *QLineEdit) Cursorposition(args ...interface{}) (ret interface{}) {
  // cursorPosition()
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
    // invoke: _ZNK9QLineEdit14cursorPositionEv
    // invoke: int cursorPosition()
    var ret0 = C.C_ZNK9QLineEdit14cursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorPosition", args)
  }

  return
}

// isModified()
func (this *QLineEdit) Ismodified(args ...interface{}) (ret interface{}) {
  // isModified()
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
    // invoke: _ZNK9QLineEdit10isModifiedEv
    // invoke: bool isModified()
    var ret0 = C.C_ZNK9QLineEdit10isModifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "isModified", args)
  }

  return
}

// isReadOnly()
func (this *QLineEdit) Isreadonly(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK9QLineEdit10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK9QLineEdit10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "isReadOnly", args)
  }

  return
}

// setCompleter(class QCompleter *)
func (this *QLineEdit) Setcompleter(args ...interface{}) () {
  // setCompleter(class QCompleter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCompleter{}) // "QCompleter *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setCompleterEP10QCompleter
    // invoke: void setCompleter(class QCompleter *)
    var arg0 = args[0].(*QCompleter).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit12setCompleterEP10QCompleter(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setCompleter", args)
  }

  return
}

// setInputMask(const class QString &)
func (this *QLineEdit) Setinputmask(args ...interface{}) () {
  // setInputMask(const class QString &)
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
    // invoke: _ZN9QLineEdit12setInputMaskERK7QString
    // invoke: void setInputMask(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit12setInputMaskERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setInputMask", args)
  }

  return
}

// setValidator(const class QValidator *)
func (this *QLineEdit) Setvalidator(args ...interface{}) () {
  // setValidator(const class QValidator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QValidator{}) // "const QValidator *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setValidatorEPK10QValidator
    // invoke: void setValidator(const class QValidator *)
    var arg0 = args[0].(*qtgui.QValidator).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit12setValidatorEPK10QValidator(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setValidator", args)
  }

  return
}

// copy()
func (this *QLineEdit) Copy(args ...interface{}) () {
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
    // invoke: _ZNK9QLineEdit4copyEv
    // invoke: void copy()
    C.C_ZNK9QLineEdit4copyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "copy", args)
  }

  return
}

// completer()
func (this *QLineEdit) Completer(args ...interface{}) (ret interface{}) {
  // completer()
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
    // invoke: _ZNK9QLineEdit9completerEv
    // invoke: QCompleter * completer()
    var ret0 = C.C_ZNK9QLineEdit9completerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCompleter{}) // "QCompleter *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "completer", args)
  }

  return
}

// setCursorPosition(int)
func (this *QLineEdit) Setcursorposition(args ...interface{}) () {
  // setCursorPosition(int)
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
    // invoke: _ZN9QLineEdit17setCursorPositionEi
    // invoke: void setCursorPosition(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit17setCursorPositionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setCursorPosition", args)
  }

  return
}

// minimumSizeHint()
func (this *QLineEdit) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
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
    // invoke: _ZNK9QLineEdit15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK9QLineEdit15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "minimumSizeHint", args)
  }

  return
}

// setPlaceholderText(const class QString &)
func (this *QLineEdit) Setplaceholdertext(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit18setPlaceholderTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setPlaceholderText", args)
  }

  return
}

// dragEnabled()
func (this *QLineEdit) Dragenabled(args ...interface{}) (ret interface{}) {
  // dragEnabled()
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
    // invoke: _ZNK9QLineEdit11dragEnabledEv
    // invoke: bool dragEnabled()
    var ret0 = C.C_ZNK9QLineEdit11dragEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "dragEnabled", args)
  }

  return
}

// setMaxLength(int)
func (this *QLineEdit) Setmaxlength(args ...interface{}) () {
  // setMaxLength(int)
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
    // invoke: _ZN9QLineEdit12setMaxLengthEi
    // invoke: void setMaxLength(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit12setMaxLengthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setMaxLength", args)
  }

  return
}

// del()
func (this *QLineEdit) Del(args ...interface{}) () {
  // del()
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
    // invoke: _ZN9QLineEdit3delEv
    // invoke: void del()
    C.C_ZN9QLineEdit3delEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "del", args)
  }

  return
}

// clear()
func (this *QLineEdit) Clear(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit5clearEv
    // invoke: void clear()
    C.C_ZN9QLineEdit5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "clear", args)
  }

  return
}

// cursorWordBackward(_Bool)
func (this *QLineEdit) Cursorwordbackward(args ...interface{}) () {
  // cursorWordBackward(_Bool)
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
    // invoke: _ZN9QLineEdit18cursorWordBackwardEb
    // invoke: void cursorWordBackward(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit18cursorWordBackwardEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorWordBackward", args)
  }

  return
}

// end(_Bool)
func (this *QLineEdit) End(args ...interface{}) () {
  // end(_Bool)
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
    // invoke: _ZN9QLineEdit3endEb
    // invoke: void end(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit3endEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "end", args)
  }

  return
}

// textMargins()
func (this *QLineEdit) Textmargins(args ...interface{}) (ret interface{}) {
  // textMargins()
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
    // invoke: _ZNK9QLineEdit11textMarginsEv
    // invoke: QMargins textMargins()
    var ret0 = C.C_ZNK9QLineEdit11textMarginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMargins{}) // "QMargins"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "textMargins", args)
  }

  return
}

// placeholderText()
func (this *QLineEdit) Placeholdertext(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK9QLineEdit15placeholderTextEv
    // invoke: QString placeholderText()
    var ret0 = C.C_ZNK9QLineEdit15placeholderTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "placeholderText", args)
  }

  return
}

// setTextMargins(const class QMargins &)
func (this *QLineEdit) Settextmargins(args ...interface{}) () {
  // setTextMargins(const class QMargins &)
  // setTextMargins(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QMargins{}) // "const QMargins &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit14setTextMarginsERK8QMargins
    // invoke: void setTextMargins(const class QMargins &)
    var arg0 = args[0].(*qtcore.QMargins).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit14setTextMarginsERK8QMargins(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN9QLineEdit14setTextMarginsEiiii
    // invoke: void setTextMargins(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN9QLineEdit14setTextMarginsEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLineEdit", "setTextMargins", args)
  }

  return
}

// setSelection(int, int)
func (this *QLineEdit) Setselection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit12setSelectionEii
    // invoke: void setSelection(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QLineEdit12setSelectionEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "setSelection", args)
  }

  return
}

// setFrame(_Bool)
func (this *QLineEdit) Setframe(args ...interface{}) () {
  // setFrame(_Bool)
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
    // invoke: _ZN9QLineEdit8setFrameEb
    // invoke: void setFrame(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit8setFrameEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setFrame", args)
  }

  return
}

// cursorForward(_Bool, int)
func (this *QLineEdit) Cursorforward(args ...interface{}) () {
  // cursorForward(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QLineEdit13cursorForwardEbi
    // invoke: void cursorForward(_Bool, int)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QLineEdit13cursorForwardEbi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorForward", args)
  }

  return
}

// setText(const class QString &)
func (this *QLineEdit) Settext(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setText", args)
  }

  return
}

// cursorMoveStyle()
func (this *QLineEdit) Cursormovestyle(args ...interface{}) () {
  // cursorMoveStyle()
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
    // invoke: _ZNK9QLineEdit15cursorMoveStyleEv
    // invoke: Qt::CursorMoveStyle cursorMoveStyle()
    C.C_ZNK9QLineEdit15cursorMoveStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QLineEdit", "cursorMoveStyle", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QLineEdit) Setreadonly(args ...interface{}) () {
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
    // invoke: _ZN9QLineEdit11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QLineEdit11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLineEdit", "setReadOnly", args)
  }

  return
}

// validator()
func (this *QLineEdit) Validator(args ...interface{}) (ret interface{}) {
  // validator()
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
    // invoke: _ZNK9QLineEdit9validatorEv
    // invoke: const QValidator * validator()
    var ret0 = C.C_ZNK9QLineEdit9validatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QValidator{}) // "const QValidator *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLineEdit", "validator", args)
  }

  return
}

// <= body block end

