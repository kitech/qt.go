package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qundostack.h
// dst-file: /src/widgets/qundostack.go
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
  // proto:  int QUndoStack::undoLimit();
extern void _ZNK10QUndoStack9undoLimitEv(void* qthis);
  // proto:  const QUndoCommand * QUndoStack::command(int index);
extern void _ZNK10QUndoStack7commandEi(void* qthis, int arg0);
  // proto:  bool QUndoStack::canRedo();
extern void _ZNK10QUndoStack7canRedoEv(void* qthis);
  // proto:  const QMetaObject * QUndoStack::metaObject();
extern void _ZNK10QUndoStack10metaObjectEv(void* qthis);
  // proto:  QString QUndoStack::redoText();
extern void _ZNK10QUndoStack8redoTextEv(void* qthis);
  // proto:  QAction * QUndoStack::createUndoAction(QObject * parent, const QString & prefix);
extern void _ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  int QUndoStack::count();
extern void _ZNK10QUndoStack5countEv(void* qthis);
  // proto:  QAction * QUndoStack::createRedoAction(QObject * parent, const QString & prefix);
extern void _ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  int QUndoStack::index();
extern void _ZNK10QUndoStack5indexEv(void* qthis);
  // proto:  void QUndoStack::clear();
extern void _ZN10QUndoStack5clearEv(void* qthis);
  // proto:  void QUndoStack::undo();
extern void _ZN10QUndoStack4undoEv(void* qthis);
  // proto:  bool QUndoStack::canUndo();
extern void _ZNK10QUndoStack7canUndoEv(void* qthis);
  // proto:  bool QUndoStack::isActive();
extern void _ZNK10QUndoStack8isActiveEv(void* qthis);
  // proto:  void QUndoStack::~QUndoStack();
extern void _ZN10QUndoStackD0Ev(void* qthis);
  // proto:  void QUndoStack::QUndoStack(QObject * parent);
extern void* dector_ZN10QUndoStackC1EP7QObject(void* arg0);
extern void _ZN10QUndoStackC1EP7QObject(void* qthis, void* arg0);
  // proto:  bool QUndoStack::isClean();
extern void _ZNK10QUndoStack7isCleanEv(void* qthis);
  // proto:  void QUndoStack::redo();
extern void _ZN10QUndoStack4redoEv(void* qthis);
  // proto:  void QUndoStack::beginMacro(const QString & text);
extern void _ZN10QUndoStack10beginMacroERK7QString(void* qthis, void* arg0);
  // proto:  void QUndoStack::setActive(bool active);
extern void _ZN10QUndoStack9setActiveEb(void* qthis, bool arg0);
  // proto:  QString QUndoStack::undoText();
extern void _ZNK10QUndoStack8undoTextEv(void* qthis);
  // proto:  int QUndoStack::cleanIndex();
extern void _ZNK10QUndoStack10cleanIndexEv(void* qthis);
  // proto:  void QUndoStack::setIndex(int idx);
extern void _ZN10QUndoStack8setIndexEi(void* qthis, int arg0);
  // proto:  void QUndoStack::endMacro();
extern void _ZN10QUndoStack8endMacroEv(void* qthis);
  // proto:  void QUndoStack::setUndoLimit(int limit);
extern void _ZN10QUndoStack12setUndoLimitEi(void* qthis, int arg0);
  // proto:  void QUndoStack::setClean();
extern void _ZN10QUndoStack8setCleanEv(void* qthis);
  // proto:  void QUndoStack::QUndoStack(const QUndoStack & );
extern void* dector_ZN10QUndoStackC1ERKS_(void* arg0);
extern void _ZN10QUndoStackC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QUndoStack::text(int idx);
extern void _ZNK10QUndoStack4textEi(void* qthis, int arg0);
  // proto:  void QUndoStack::push(QUndoCommand * cmd);
extern void _ZN10QUndoStack4pushEP12QUndoCommand(void* qthis, void* arg0);
  // proto:  int QUndoCommand::id();
extern void _ZNK12QUndoCommand2idEv(void* qthis);
  // proto:  void QUndoCommand::redo();
extern void _ZN12QUndoCommand4redoEv(void* qthis);
  // proto:  void QUndoCommand::QUndoCommand(const QUndoCommand & );
extern void* dector_ZN12QUndoCommandC1ERKS_(void* arg0);
extern void _ZN12QUndoCommandC1ERKS_(void* qthis, void* arg0);
  // proto:  void QUndoCommand::QUndoCommand(QUndoCommand * parent);
extern void* dector_ZN12QUndoCommandC1EPS_(void* arg0);
extern void _ZN12QUndoCommandC1EPS_(void* qthis, void* arg0);
  // proto:  void QUndoCommand::undo();
extern void _ZN12QUndoCommand4undoEv(void* qthis);
  // proto:  void QUndoCommand::QUndoCommand(const QString & text, QUndoCommand * parent);
extern void* dector_ZN12QUndoCommandC1ERK7QStringPS_(void* arg0, void* arg1);
extern void _ZN12QUndoCommandC1ERK7QStringPS_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QUndoCommand::mergeWith(const QUndoCommand * other);
extern void _ZN12QUndoCommand9mergeWithEPKS_(void* qthis, void* arg0);
  // proto:  QString QUndoCommand::text();
extern void _ZNK12QUndoCommand4textEv(void* qthis);
  // proto:  int QUndoCommand::childCount();
extern void _ZNK12QUndoCommand10childCountEv(void* qthis);
  // proto:  QString QUndoCommand::actionText();
extern void _ZNK12QUndoCommand10actionTextEv(void* qthis);
  // proto:  void QUndoCommand::~QUndoCommand();
extern void _ZN12QUndoCommandD0Ev(void* qthis);
  // proto:  const QUndoCommand * QUndoCommand::child(int index);
extern void _ZNK12QUndoCommand5childEi(void* qthis, int arg0);
  // proto:  void QUndoCommand::setText(const QString & text);
extern void _ZN12QUndoCommand7setTextERK7QString(void* qthis, void* arg0);
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

// class sizeof(QUndoStack)=1
type QUndoStack struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _redoTextChanged QUndoStack_redoTextChanged_signal;
//  _cleanChanged QUndoStack_cleanChanged_signal;
//  _canUndoChanged QUndoStack_canUndoChanged_signal;
//  _indexChanged QUndoStack_indexChanged_signal;
//  _canRedoChanged QUndoStack_canRedoChanged_signal;
//  _undoTextChanged QUndoStack_undoTextChanged_signal;
}

// class sizeof(QUndoCommand)=16
type QUndoCommand struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QUndoStack::undoLimit();
func (this *QUndoStack) undoLimit(args ...interface{}) () {
  // undoLimit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack9undoLimitEv
  default:
    qtrt.ErrorResolve("QUndoStack", "undoLimit", args)
  }

}

  // proto:  const QUndoCommand * QUndoStack::command(int index);
func (this *QUndoStack) command(args ...interface{}) () {
  // command(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack7commandEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "command", args)
  }

}

  // proto:  bool QUndoStack::canRedo();
func (this *QUndoStack) canRedo(args ...interface{}) () {
  // canRedo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack7canRedoEv
  default:
    qtrt.ErrorResolve("QUndoStack", "canRedo", args)
  }

}

  // proto:  const QMetaObject * QUndoStack::metaObject();
func (this *QUndoStack) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack10metaObjectEv
  default:
    qtrt.ErrorResolve("QUndoStack", "metaObject", args)
  }

}

  // proto:  QString QUndoStack::redoText();
func (this *QUndoStack) redoText(args ...interface{}) () {
  // redoText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack8redoTextEv
  default:
    qtrt.ErrorResolve("QUndoStack", "redoText", args)
  }

}

  // proto:  QAction * QUndoStack::createUndoAction(QObject * parent, const QString & prefix);
func (this *QUndoStack) createUndoAction(args ...interface{}) () {
  // createUndoAction(class QObject *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QUndoStack", "createUndoAction", args)
  }

}

  // proto:  int QUndoStack::count();
func (this *QUndoStack) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack5countEv
  default:
    qtrt.ErrorResolve("QUndoStack", "count", args)
  }

}

  // proto:  QAction * QUndoStack::createRedoAction(QObject * parent, const QString & prefix);
func (this *QUndoStack) createRedoAction(args ...interface{}) () {
  // createRedoAction(class QObject *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QUndoStack", "createRedoAction", args)
  }

}

  // proto:  int QUndoStack::index();
func (this *QUndoStack) index(args ...interface{}) () {
  // index()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack5indexEv
  default:
    qtrt.ErrorResolve("QUndoStack", "index", args)
  }

}

  // proto:  void QUndoStack::clear();
func (this *QUndoStack) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack5clearEv
  default:
    qtrt.ErrorResolve("QUndoStack", "clear", args)
  }

}

  // proto:  void QUndoStack::undo();
func (this *QUndoStack) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack4undoEv
  default:
    qtrt.ErrorResolve("QUndoStack", "undo", args)
  }

}

  // proto:  bool QUndoStack::canUndo();
func (this *QUndoStack) canUndo(args ...interface{}) () {
  // canUndo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack7canUndoEv
  default:
    qtrt.ErrorResolve("QUndoStack", "canUndo", args)
  }

}

  // proto:  bool QUndoStack::isActive();
func (this *QUndoStack) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack8isActiveEv
  default:
    qtrt.ErrorResolve("QUndoStack", "isActive", args)
  }

}

  // proto:  void QUndoStack::~QUndoStack();
func (this *QUndoStack) FreeQUndoStack(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUndoStack", "~QUndoStack", args)
  }

}

  // proto:  void QUndoStack::QUndoStack(QObject * parent);
func NewQUndoStack(args ...interface{}) QUndoStack {
  return QUndoStack{}
}

  // proto:  bool QUndoStack::isClean();
func (this *QUndoStack) isClean(args ...interface{}) () {
  // isClean()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack7isCleanEv
  default:
    qtrt.ErrorResolve("QUndoStack", "isClean", args)
  }

}

  // proto:  void QUndoStack::redo();
func (this *QUndoStack) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack4redoEv
  default:
    qtrt.ErrorResolve("QUndoStack", "redo", args)
  }

}

  // proto:  void QUndoStack::beginMacro(const QString & text);
func (this *QUndoStack) beginMacro(args ...interface{}) () {
  // beginMacro(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack10beginMacroERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "beginMacro", args)
  }

}

  // proto:  void QUndoStack::setActive(bool active);
func (this *QUndoStack) setActive(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack9setActiveEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "setActive", args)
  }

}

  // proto:  QString QUndoStack::undoText();
func (this *QUndoStack) undoText(args ...interface{}) () {
  // undoText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack8undoTextEv
  default:
    qtrt.ErrorResolve("QUndoStack", "undoText", args)
  }

}

  // proto:  int QUndoStack::cleanIndex();
func (this *QUndoStack) cleanIndex(args ...interface{}) () {
  // cleanIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack10cleanIndexEv
  default:
    qtrt.ErrorResolve("QUndoStack", "cleanIndex", args)
  }

}

  // proto:  void QUndoStack::setIndex(int idx);
func (this *QUndoStack) setIndex(args ...interface{}) () {
  // setIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack8setIndexEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "setIndex", args)
  }

}

  // proto:  void QUndoStack::endMacro();
func (this *QUndoStack) endMacro(args ...interface{}) () {
  // endMacro()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack8endMacroEv
  default:
    qtrt.ErrorResolve("QUndoStack", "endMacro", args)
  }

}

  // proto:  void QUndoStack::setUndoLimit(int limit);
func (this *QUndoStack) setUndoLimit(args ...interface{}) () {
  // setUndoLimit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack12setUndoLimitEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "setUndoLimit", args)
  }

}

  // proto:  void QUndoStack::setClean();
func (this *QUndoStack) setClean(args ...interface{}) () {
  // setClean()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack8setCleanEv
  default:
    qtrt.ErrorResolve("QUndoStack", "setClean", args)
  }

}

  // proto:  QString QUndoStack::text(int idx);
func (this *QUndoStack) text(args ...interface{}) () {
  // text(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack4textEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "text", args)
  }

}

  // proto:  void QUndoStack::push(QUndoCommand * cmd);
func (this *QUndoStack) push(args ...interface{}) () {
  // push(class QUndoCommand *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoCommand{}) // "QUndoCommand *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack4pushEP12QUndoCommand
    var arg0 = args[0].(QUndoCommand).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoStack", "push", args)
  }

}

  // proto:  int QUndoCommand::id();
func (this *QUndoCommand) id(args ...interface{}) () {
  // id()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QUndoCommand2idEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "id", args)
  }

}

  // proto:  void QUndoCommand::redo();
func (this *QUndoCommand) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommand4redoEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "redo", args)
  }

}

  // proto:  void QUndoCommand::QUndoCommand(const QUndoCommand & );
func NewQUndoCommand(args ...interface{}) QUndoCommand {
  return QUndoCommand{}
}

  // proto:  void QUndoCommand::undo();
func (this *QUndoCommand) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommand4undoEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "undo", args)
  }

}

  // proto:  bool QUndoCommand::mergeWith(const QUndoCommand * other);
func (this *QUndoCommand) mergeWith(args ...interface{}) () {
  // mergeWith(const class QUndoCommand *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoCommand{}) // "const QUndoCommand *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommand9mergeWithEPKS_
    var arg0 = args[0].(QUndoCommand).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "mergeWith", args)
  }

}

  // proto:  QString QUndoCommand::text();
func (this *QUndoCommand) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QUndoCommand4textEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "text", args)
  }

}

  // proto:  int QUndoCommand::childCount();
func (this *QUndoCommand) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QUndoCommand10childCountEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "childCount", args)
  }

}

  // proto:  QString QUndoCommand::actionText();
func (this *QUndoCommand) actionText(args ...interface{}) () {
  // actionText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QUndoCommand10actionTextEv
  default:
    qtrt.ErrorResolve("QUndoCommand", "actionText", args)
  }

}

  // proto:  void QUndoCommand::~QUndoCommand();
func (this *QUndoCommand) FreeQUndoCommand(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUndoCommand", "~QUndoCommand", args)
  }

}

  // proto:  const QUndoCommand * QUndoCommand::child(int index);
func (this *QUndoCommand) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QUndoCommand5childEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "child", args)
  }

}

  // proto:  void QUndoCommand::setText(const QString & text);
func (this *QUndoCommand) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommand7setTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "setText", args)
  }

}

// <= body block end

