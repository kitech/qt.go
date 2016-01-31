package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAction * QUndoStack::createUndoAction(QObject * parent, const QString & prefix);
extern void C_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QUndoStack::setActive(bool active);
extern void C_ZN10QUndoStack9setActiveEb(void* qthis, bool arg0); // 4
  // proto:  void QUndoStack::setIndex(int idx);
extern void C_ZN10QUndoStack8setIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoStack::redo();
extern void C_ZN10QUndoStack4redoEv(void* qthis); // 4
  // proto:  bool QUndoStack::canUndo();
extern void C_ZNK10QUndoStack7canUndoEv(void* qthis); // 4
  // proto:  bool QUndoStack::isClean();
extern void C_ZNK10QUndoStack7isCleanEv(void* qthis); // 4
  // proto:  int QUndoStack::cleanIndex();
extern void C_ZNK10QUndoStack10cleanIndexEv(void* qthis); // 4
  // proto:  void QUndoStack::QUndoStack(QObject * parent);
extern void* C_ZN10QUndoStackC2EP7QObject(void* arg0); // 3
  // proto:  void QUndoStack::~QUndoStack();
extern void C_ZN10QUndoStackD2Ev(void* qthis); // 4
  // proto:  QString QUndoStack::text(int idx);
extern void C_ZNK10QUndoStack4textEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QUndoStack::redoText();
extern void C_ZNK10QUndoStack8redoTextEv(void* qthis); // 4
  // proto:  int QUndoStack::count();
extern void C_ZNK10QUndoStack5countEv(void* qthis); // 4
  // proto:  int QUndoStack::index();
extern void C_ZNK10QUndoStack5indexEv(void* qthis); // 4
  // proto:  void QUndoStack::beginMacro(const QString & text);
extern void C_ZN10QUndoStack10beginMacroERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoStack::undo();
extern void C_ZN10QUndoStack4undoEv(void* qthis); // 4
  // proto:  int QUndoStack::undoLimit();
extern void C_ZNK10QUndoStack9undoLimitEv(void* qthis); // 4
  // proto:  bool QUndoStack::canRedo();
extern void C_ZNK10QUndoStack7canRedoEv(void* qthis); // 4
  // proto:  void QUndoStack::setClean();
extern void C_ZN10QUndoStack8setCleanEv(void* qthis); // 4
  // proto:  bool QUndoStack::isActive();
extern void C_ZNK10QUndoStack8isActiveEv(void* qthis); // 4
  // proto:  QAction * QUndoStack::createRedoAction(QObject * parent, const QString & prefix);
extern void C_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QUndoStack::metaObject();
extern void C_ZNK10QUndoStack10metaObjectEv(void* qthis); // 4
  // proto:  QString QUndoStack::undoText();
extern void C_ZNK10QUndoStack8undoTextEv(void* qthis); // 4
  // proto:  void QUndoStack::clear();
extern void C_ZN10QUndoStack5clearEv(void* qthis); // 4
  // proto:  void QUndoStack::setUndoLimit(int limit);
extern void C_ZN10QUndoStack12setUndoLimitEi(void* qthis, int32_t arg0); // 4
  // proto:  const QUndoCommand * QUndoStack::command(int index);
extern void C_ZNK10QUndoStack7commandEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoStack::endMacro();
extern void C_ZN10QUndoStack8endMacroEv(void* qthis); // 4
  // proto:  void QUndoStack::push(QUndoCommand * cmd);
extern void C_ZN10QUndoStack4pushEP12QUndoCommand(void* qthis, void* arg0); // 4
  // proto:  QString QUndoCommand::actionText();
extern void C_ZNK12QUndoCommand10actionTextEv(void* qthis); // 4
  // proto:  bool QUndoCommand::mergeWith(const QUndoCommand * other);
extern void C_ZN12QUndoCommand9mergeWithEPKS_(void* qthis, void* arg0); // 4
  // proto:  QString QUndoCommand::text();
extern void C_ZNK12QUndoCommand4textEv(void* qthis); // 4
  // proto:  void QUndoCommand::setText(const QString & text);
extern void C_ZN12QUndoCommand7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoCommand::undo();
extern void C_ZN12QUndoCommand4undoEv(void* qthis); // 4
  // proto:  int QUndoCommand::childCount();
extern void C_ZNK12QUndoCommand10childCountEv(void* qthis); // 4
  // proto:  void QUndoCommand::QUndoCommand(const QString & text, QUndoCommand * parent);
extern void* C_ZN12QUndoCommandC2ERK7QStringPS_(void* arg0, void* arg1); // 3
  // proto:  void QUndoCommand::QUndoCommand(QUndoCommand * parent);
extern void* C_ZN12QUndoCommandC2EPS_(void* arg0); // 3
  // proto:  void QUndoCommand::~QUndoCommand();
extern void C_ZN12QUndoCommandD2Ev(void* qthis); // 4
  // proto:  const QUndoCommand * QUndoCommand::child(int index);
extern void C_ZNK12QUndoCommand5childEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoCommand::redo();
extern void C_ZN12QUndoCommand4redoEv(void* qthis); // 4
  // proto:  int QUndoCommand::id();
extern void C_ZNK12QUndoCommand2idEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// createUndoAction(class QObject *, const class QString &)
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
    // invoke: QAction * createUndoAction(class QObject *, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "createUndoAction", args)
  }

}

// setActive(_Bool)
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
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack9setActiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setActive", args)
  }

}

// setIndex(int)
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
    // invoke: void setIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack8setIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setIndex", args)
  }

}

// redo()
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
    // invoke: void redo()
    C.C_ZN10QUndoStack4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "redo", args)
  }

}

// canUndo()
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
    // invoke: bool canUndo()
    var ret = C.C_ZNK10QUndoStack7canUndoEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "canUndo", args)
  }

}

// isClean()
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
    // invoke: bool isClean()
    var ret = C.C_ZNK10QUndoStack7isCleanEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "isClean", args)
  }

}

// cleanIndex()
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
    // invoke: int cleanIndex()
    var ret = C.C_ZNK10QUndoStack10cleanIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "cleanIndex", args)
  }

}

// QUndoStack(class QObject *)
func NewQUndoStack(args ...interface{}) *QUndoStack {
  // QUndoStack(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStackC1EP7QObject
    // invoke: void QUndoStack(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QUndoStackC2EP7QObject(arg0)
    return &QUndoStack{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoStack", "QUndoStack", args)
  }

  return nil // QUndoStack{qclsinst:qthis}
}

// ~QUndoStack()
func (this *QUndoStack) FreeQUndoStack(args ...interface{}) () {
  // ~QUndoStack()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStackD0Ev
    // invoke: void ~QUndoStack()
    C.C_ZN10QUndoStackD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "~QUndoStack", args)
  }

}

// text(int)
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
    // invoke: QString text(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QUndoStack4textEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "text", args)
  }

}

// redoText()
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
    // invoke: QString redoText()
    var ret = C.C_ZNK10QUndoStack8redoTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "redoText", args)
  }

}

// count()
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
    // invoke: int count()
    var ret = C.C_ZNK10QUndoStack5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "count", args)
  }

}

// index()
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
    // invoke: int index()
    var ret = C.C_ZNK10QUndoStack5indexEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "index", args)
  }

}

// beginMacro(const class QString &)
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
    // invoke: void beginMacro(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack10beginMacroERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "beginMacro", args)
  }

}

// undo()
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
    // invoke: void undo()
    C.C_ZN10QUndoStack4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "undo", args)
  }

}

// undoLimit()
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
    // invoke: int undoLimit()
    var ret = C.C_ZNK10QUndoStack9undoLimitEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "undoLimit", args)
  }

}

// canRedo()
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
    // invoke: bool canRedo()
    var ret = C.C_ZNK10QUndoStack7canRedoEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "canRedo", args)
  }

}

// setClean()
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
    // invoke: void setClean()
    C.C_ZN10QUndoStack8setCleanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "setClean", args)
  }

}

// isActive()
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
    // invoke: bool isActive()
    var ret = C.C_ZNK10QUndoStack8isActiveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "isActive", args)
  }

}

// createRedoAction(class QObject *, const class QString &)
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
    // invoke: QAction * createRedoAction(class QObject *, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "createRedoAction", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QUndoStack10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "metaObject", args)
  }

}

// undoText()
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
    // invoke: QString undoText()
    var ret = C.C_ZNK10QUndoStack8undoTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "undoText", args)
  }

}

// clear()
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
    // invoke: void clear()
    C.C_ZN10QUndoStack5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "clear", args)
  }

}

// setUndoLimit(int)
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
    // invoke: void setUndoLimit(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack12setUndoLimitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setUndoLimit", args)
  }

}

// command(int)
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
    // invoke: const QUndoCommand * command(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QUndoStack7commandEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoStack", "command", args)
  }

}

// endMacro()
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
    // invoke: void endMacro()
    C.C_ZN10QUndoStack8endMacroEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "endMacro", args)
  }

}

// push(class QUndoCommand *)
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
    // invoke: void push(class QUndoCommand *)
    var arg0 = args[0].(QUndoCommand).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack4pushEP12QUndoCommand(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "push", args)
  }

}

// actionText()
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
    // invoke: QString actionText()
    var ret = C.C_ZNK12QUndoCommand10actionTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "actionText", args)
  }

}

// mergeWith(const class QUndoCommand *)
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
    // invoke: bool mergeWith(const class QUndoCommand *)
    var arg0 = args[0].(QUndoCommand).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN12QUndoCommand9mergeWithEPKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "mergeWith", args)
  }

}

// text()
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
    // invoke: QString text()
    var ret = C.C_ZNK12QUndoCommand4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "text", args)
  }

}

// setText(const class QString &)
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
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QUndoCommand7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoCommand", "setText", args)
  }

}

// undo()
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
    // invoke: void undo()
    C.C_ZN12QUndoCommand4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "undo", args)
  }

}

// childCount()
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
    // invoke: int childCount()
    var ret = C.C_ZNK12QUndoCommand10childCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "childCount", args)
  }

}

// QUndoCommand(const class QString &, class QUndoCommand *)
func NewQUndoCommand(args ...interface{}) *QUndoCommand {
  // QUndoCommand(const class QString &, class QUndoCommand *)
  // QUndoCommand(class QUndoCommand *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QUndoCommand{}) // "QUndoCommand *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUndoCommand{}) // "QUndoCommand *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommandC1ERK7QStringPS_
    // invoke: void QUndoCommand(const class QString &, class QUndoCommand *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUndoCommand).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QUndoCommandC2ERK7QStringPS_(arg0, arg1)
    return &QUndoCommand{qclsinst:qthis}
  case 1:
    // invoke: _ZN12QUndoCommandC1EPS_
    // invoke: void QUndoCommand(class QUndoCommand *)
    var arg0 = args[0].(QUndoCommand).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QUndoCommandC2EPS_(arg0)
    return &QUndoCommand{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoCommand", "QUndoCommand", args)
  }

  return nil // QUndoCommand{qclsinst:qthis}
}

// ~QUndoCommand()
func (this *QUndoCommand) FreeQUndoCommand(args ...interface{}) () {
  // ~QUndoCommand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommandD0Ev
    // invoke: void ~QUndoCommand()
    C.C_ZN12QUndoCommandD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "~QUndoCommand", args)
  }

}

// child(int)
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
    // invoke: const QUndoCommand * child(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QUndoCommand5childEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "child", args)
  }

}

// redo()
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
    // invoke: void redo()
    C.C_ZN12QUndoCommand4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "redo", args)
  }

}

// id()
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
    // invoke: int id()
    var ret = C.C_ZNK12QUndoCommand2idEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoCommand", "id", args)
  }

}

// <= body block end

