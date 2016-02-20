package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
  // proto:  QAction * QUndoStack::createUndoAction(QObject * parent, const QString & prefix);
extern void* C_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QUndoStack::setActive(bool active);
extern void C_ZN10QUndoStack9setActiveEb(void* qthis, bool arg0); // 4
  // proto:  void QUndoStack::setIndex(int idx);
extern void C_ZN10QUndoStack8setIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoStack::redo();
extern void C_ZN10QUndoStack4redoEv(void* qthis); // 4
  // proto:  bool QUndoStack::canUndo();
extern bool C_ZNK10QUndoStack7canUndoEv(void* qthis); // 4
  // proto:  bool QUndoStack::isClean();
extern bool C_ZNK10QUndoStack7isCleanEv(void* qthis); // 4
  // proto:  int QUndoStack::cleanIndex();
extern int32_t C_ZNK10QUndoStack10cleanIndexEv(void* qthis); // 4
  // proto:  void QUndoStack::QUndoStack(QObject * parent);
extern void* C_ZN10QUndoStackC2EP7QObject(void* arg0); // 3
  // proto:  void QUndoStack::~QUndoStack();
extern void C_ZN10QUndoStackD2Ev(void* qthis); // 4
  // proto:  QString QUndoStack::text(int idx);
extern void* C_ZNK10QUndoStack4textEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QUndoStack::redoText();
extern void* C_ZNK10QUndoStack8redoTextEv(void* qthis); // 4
  // proto:  int QUndoStack::count();
extern int32_t C_ZNK10QUndoStack5countEv(void* qthis); // 4
  // proto:  int QUndoStack::index();
extern int32_t C_ZNK10QUndoStack5indexEv(void* qthis); // 4
  // proto:  void QUndoStack::beginMacro(const QString & text);
extern void C_ZN10QUndoStack10beginMacroERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoStack::undo();
extern void C_ZN10QUndoStack4undoEv(void* qthis); // 4
  // proto:  int QUndoStack::undoLimit();
extern int32_t C_ZNK10QUndoStack9undoLimitEv(void* qthis); // 4
  // proto:  bool QUndoStack::canRedo();
extern bool C_ZNK10QUndoStack7canRedoEv(void* qthis); // 4
  // proto:  void QUndoStack::setClean();
extern void C_ZN10QUndoStack8setCleanEv(void* qthis); // 4
  // proto:  bool QUndoStack::isActive();
extern bool C_ZNK10QUndoStack8isActiveEv(void* qthis); // 4
  // proto:  QAction * QUndoStack::createRedoAction(QObject * parent, const QString & prefix);
extern void* C_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QUndoStack::metaObject();
extern void C_ZNK10QUndoStack10metaObjectEv(void* qthis); // 4
  // proto:  QString QUndoStack::undoText();
extern void* C_ZNK10QUndoStack8undoTextEv(void* qthis); // 4
  // proto:  void QUndoStack::clear();
extern void C_ZN10QUndoStack5clearEv(void* qthis); // 4
  // proto:  void QUndoStack::setUndoLimit(int limit);
extern void C_ZN10QUndoStack12setUndoLimitEi(void* qthis, int32_t arg0); // 4
  // proto:  const QUndoCommand * QUndoStack::command(int index);
extern void* C_ZNK10QUndoStack7commandEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoStack::endMacro();
extern void C_ZN10QUndoStack8endMacroEv(void* qthis); // 4
  // proto:  void QUndoStack::push(QUndoCommand * cmd);
extern void C_ZN10QUndoStack4pushEP12QUndoCommand(void* qthis, void* arg0); // 4
  // proto:  QString QUndoCommand::actionText();
extern void* C_ZNK12QUndoCommand10actionTextEv(void* qthis); // 4
  // proto:  bool QUndoCommand::mergeWith(const QUndoCommand * other);
extern bool C_ZN12QUndoCommand9mergeWithEPKS_(void* qthis, void* arg0); // 4
  // proto:  QString QUndoCommand::text();
extern void* C_ZNK12QUndoCommand4textEv(void* qthis); // 4
  // proto:  void QUndoCommand::setText(const QString & text);
extern void C_ZN12QUndoCommand7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoCommand::undo();
extern void C_ZN12QUndoCommand4undoEv(void* qthis); // 4
  // proto:  int QUndoCommand::childCount();
extern int32_t C_ZNK12QUndoCommand10childCountEv(void* qthis); // 4
  // proto:  void QUndoCommand::QUndoCommand(const QString & text, QUndoCommand * parent);
extern void* C_ZN12QUndoCommandC2ERK7QStringPS_(void* arg0, void* arg1); // 3
  // proto:  void QUndoCommand::QUndoCommand(QUndoCommand * parent);
extern void* C_ZN12QUndoCommandC2EPS_(void* arg0); // 3
  // proto:  void QUndoCommand::~QUndoCommand();
extern void C_ZN12QUndoCommandD2Ev(void* qthis); // 4
  // proto:  const QUndoCommand * QUndoCommand::child(int index);
extern void* C_ZNK12QUndoCommand5childEi(void* qthis, int32_t arg0); // 4
  // proto:  void QUndoCommand::redo();
extern void C_ZN12QUndoCommand4redoEv(void* qthis); // 4
  // proto:  int QUndoCommand::id();
extern int32_t C_ZNK12QUndoCommand2idEv(void* qthis); // 4
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

// class sizeof(QUndoStack)=1
type QUndoStack struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// createUndoAction(class QObject *, const class QString &)
func (this *QUndoStack) Createundoaction(args ...interface{}) (ret interface{}) {
  // createUndoAction(class QObject *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString
    // invoke: QAction * createUndoAction(class QObject *, const class QString &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "createUndoAction", args)
  }

  return
}

// setActive(_Bool)
func (this *QUndoStack) Setactive(args ...interface{}) () {
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
    C.C_ZN10QUndoStack9setActiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setActive", args)
  }

  return
}

// setIndex(int)
func (this *QUndoStack) Setindex(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack8setIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setIndex", args)
  }

  return
}

// redo()
func (this *QUndoStack) Redo(args ...interface{}) () {
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
    C.C_ZN10QUndoStack4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "redo", args)
  }

  return
}

// canUndo()
func (this *QUndoStack) Canundo(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack7canUndoEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "canUndo", args)
  }

  return
}

// isClean()
func (this *QUndoStack) Isclean(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack7isCleanEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "isClean", args)
  }

  return
}

// cleanIndex()
func (this *QUndoStack) Cleanindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack10cleanIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "cleanIndex", args)
  }

  return
}

// QUndoStack(class QObject *)
func NewQUndoStack(args ...interface{}) *QUndoStack {
  // QUndoStack(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStackC1EP7QObject
    // invoke: void QUndoStack(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QUndoStackC2EP7QObject(arg0)
    return &QUndoStack{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoStack", "QUndoStack", args)
  }

  return nil // QUndoStack{Qclsinst:qthis}
}

// ~QUndoStack()
func (this *QUndoStack) Freequndostack(args ...interface{}) () {
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
    C.C_ZN10QUndoStackD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "~QUndoStack", args)
  }

  return
}

// text(int)
func (this *QUndoStack) Text(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QUndoStack4textEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "text", args)
  }

  return
}

// redoText()
func (this *QUndoStack) Redotext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack8redoTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "redoText", args)
  }

  return
}

// count()
func (this *QUndoStack) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "count", args)
  }

  return
}

// index()
func (this *QUndoStack) Index(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack5indexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "index", args)
  }

  return
}

// beginMacro(const class QString &)
func (this *QUndoStack) Beginmacro(args ...interface{}) () {
  // beginMacro(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoStack10beginMacroERK7QString
    // invoke: void beginMacro(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack10beginMacroERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "beginMacro", args)
  }

  return
}

// undo()
func (this *QUndoStack) Undo(args ...interface{}) () {
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
    C.C_ZN10QUndoStack4undoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "undo", args)
  }

  return
}

// undoLimit()
func (this *QUndoStack) Undolimit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack9undoLimitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "undoLimit", args)
  }

  return
}

// canRedo()
func (this *QUndoStack) Canredo(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack7canRedoEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "canRedo", args)
  }

  return
}

// setClean()
func (this *QUndoStack) Setclean(args ...interface{}) () {
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
    C.C_ZN10QUndoStack8setCleanEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "setClean", args)
  }

  return
}

// isActive()
func (this *QUndoStack) Isactive(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack8isActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "isActive", args)
  }

  return
}

// createRedoAction(class QObject *, const class QString &)
func (this *QUndoStack) Createredoaction(args ...interface{}) (ret interface{}) {
  // createRedoAction(class QObject *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString
    // invoke: QAction * createRedoAction(class QObject *, const class QString &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "createRedoAction", args)
  }

  return
}

// metaObject()
func (this *QUndoStack) Metaobject(args ...interface{}) () {
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
    C.C_ZNK10QUndoStack10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "metaObject", args)
  }

  return
}

// undoText()
func (this *QUndoStack) Undotext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QUndoStack8undoTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "undoText", args)
  }

  return
}

// clear()
func (this *QUndoStack) Clear(args ...interface{}) () {
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
    C.C_ZN10QUndoStack5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "clear", args)
  }

  return
}

// setUndoLimit(int)
func (this *QUndoStack) Setundolimit(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack12setUndoLimitEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "setUndoLimit", args)
  }

  return
}

// command(int)
func (this *QUndoStack) Command(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QUndoStack7commandEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUndoCommand{}) // "const QUndoCommand *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoStack", "command", args)
  }

  return
}

// endMacro()
func (this *QUndoStack) Endmacro(args ...interface{}) () {
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
    C.C_ZN10QUndoStack8endMacroEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoStack", "endMacro", args)
  }

  return
}

// push(class QUndoCommand *)
func (this *QUndoStack) Push(args ...interface{}) () {
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
    var arg0 = args[0].(*QUndoCommand).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoStack4pushEP12QUndoCommand(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoStack", "push", args)
  }

  return
}

// actionText()
func (this *QUndoCommand) Actiontext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QUndoCommand10actionTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "actionText", args)
  }

  return
}

// mergeWith(const class QUndoCommand *)
func (this *QUndoCommand) Mergewith(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QUndoCommand).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QUndoCommand9mergeWithEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "mergeWith", args)
  }

  return
}

// text()
func (this *QUndoCommand) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QUndoCommand4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "text", args)
  }

  return
}

// setText(const class QString &)
func (this *QUndoCommand) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommand7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QUndoCommand7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoCommand", "setText", args)
  }

  return
}

// undo()
func (this *QUndoCommand) Undo(args ...interface{}) () {
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
    C.C_ZN12QUndoCommand4undoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "undo", args)
  }

  return
}

// childCount()
func (this *QUndoCommand) Childcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QUndoCommand10childCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "childCount", args)
  }

  return
}

// QUndoCommand(const class QString &, class QUndoCommand *)
func NewQUndoCommand(args ...interface{}) *QUndoCommand {
  // QUndoCommand(const class QString &, class QUndoCommand *)
  // QUndoCommand(class QUndoCommand *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QUndoCommand{}) // "QUndoCommand *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUndoCommand{}) // "QUndoCommand *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QUndoCommandC1ERK7QStringPS_
    // invoke: void QUndoCommand(const class QString &, class QUndoCommand *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QUndoCommand).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QUndoCommandC2ERK7QStringPS_(arg0, arg1)
    return &QUndoCommand{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QUndoCommandC1EPS_
    // invoke: void QUndoCommand(class QUndoCommand *)
    var arg0 = args[0].(*QUndoCommand).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QUndoCommandC2EPS_(arg0)
    return &QUndoCommand{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoCommand", "QUndoCommand", args)
  }

  return nil // QUndoCommand{Qclsinst:qthis}
}

// ~QUndoCommand()
func (this *QUndoCommand) Freequndocommand(args ...interface{}) () {
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
    C.C_ZN12QUndoCommandD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "~QUndoCommand", args)
  }

  return
}

// child(int)
func (this *QUndoCommand) Child(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QUndoCommand5childEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUndoCommand{}) // "const QUndoCommand *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "child", args)
  }

  return
}

// redo()
func (this *QUndoCommand) Redo(args ...interface{}) () {
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
    C.C_ZN12QUndoCommand4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoCommand", "redo", args)
  }

  return
}

// id()
func (this *QUndoCommand) Id(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QUndoCommand2idEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoCommand", "id", args)
  }

  return
}

// <= body block end

