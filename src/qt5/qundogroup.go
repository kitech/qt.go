package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qundogroup.h
// dst-file: /src/widgets/qundogroup.go
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
  // proto:  QAction * QUndoGroup::createUndoAction(QObject * parent, const QString & prefix);
extern void* C_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QUndoGroup::isClean();
extern bool C_ZNK10QUndoGroup7isCleanEv(void* qthis); // 4
  // proto:  void QUndoGroup::redo();
extern void C_ZN10QUndoGroup4redoEv(void* qthis); // 4
  // proto:  bool QUndoGroup::canUndo();
extern bool C_ZNK10QUndoGroup7canUndoEv(void* qthis); // 4
  // proto:  void QUndoGroup::~QUndoGroup();
extern void C_ZN10QUndoGroupD2Ev(void* qthis); // 4
  // proto:  void QUndoGroup::QUndoGroup(QObject * parent);
extern void* C_ZN10QUndoGroupC2EP7QObject(void* arg0); // 3
  // proto:  QString QUndoGroup::redoText();
extern void* C_ZNK10QUndoGroup8redoTextEv(void* qthis); // 4
  // proto:  void QUndoGroup::undo();
extern void C_ZN10QUndoGroup4undoEv(void* qthis); // 4
  // proto:  bool QUndoGroup::canRedo();
extern bool C_ZNK10QUndoGroup7canRedoEv(void* qthis); // 4
  // proto:  void QUndoGroup::addStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup8addStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  QAction * QUndoGroup::createRedoAction(QObject * parent, const QString & prefix);
extern void* C_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QUndoGroup::metaObject();
extern void C_ZNK10QUndoGroup10metaObjectEv(void* qthis); // 4
  // proto:  QString QUndoGroup::undoText();
extern void* C_ZNK10QUndoGroup8undoTextEv(void* qthis); // 4
  // proto:  void QUndoGroup::removeStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup11removeStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  void QUndoGroup::setActiveStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup14setActiveStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  QList<QUndoStack *> QUndoGroup::stacks();
extern void C_ZNK10QUndoGroup6stacksEv(void* qthis); // 4
  // proto:  QUndoStack * QUndoGroup::activeStack();
extern void* C_ZNK10QUndoGroup11activeStackEv(void* qthis); // 4
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

// class sizeof(QUndoGroup)=1
type QUndoGroup struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _redoTextChanged QUndoGroup_redoTextChanged_signal;
//  _cleanChanged QUndoGroup_cleanChanged_signal;
//  _canUndoChanged QUndoGroup_canUndoChanged_signal;
//  _indexChanged QUndoGroup_indexChanged_signal;
//  _activeStackChanged QUndoGroup_activeStackChanged_signal;
//  _canRedoChanged QUndoGroup_canRedoChanged_signal;
//  _undoTextChanged QUndoGroup_undoTextChanged_signal;
}

// createUndoAction(class QObject *, const class QString &)
func (this *QUndoGroup) Createundoaction(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString
    // invoke: QAction * createUndoAction(class QObject *, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "createUndoAction", args)
  }

  return
}

// isClean()
func (this *QUndoGroup) Isclean(args ...interface{}) (ret interface{}) {
  // isClean()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup7isCleanEv
    // invoke: bool isClean()
    var ret0 = C.C_ZNK10QUndoGroup7isCleanEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "isClean", args)
  }

  return
}

// redo()
func (this *QUndoGroup) Redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroup4redoEv
    // invoke: void redo()
    C.C_ZN10QUndoGroup4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "redo", args)
  }

  return
}

// canUndo()
func (this *QUndoGroup) Canundo(args ...interface{}) (ret interface{}) {
  // canUndo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup7canUndoEv
    // invoke: bool canUndo()
    var ret0 = C.C_ZNK10QUndoGroup7canUndoEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "canUndo", args)
  }

  return
}

// ~QUndoGroup()
func (this *QUndoGroup) Freequndogroup(args ...interface{}) () {
  // ~QUndoGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroupD0Ev
    // invoke: void ~QUndoGroup()
    C.C_ZN10QUndoGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "~QUndoGroup", args)
  }

  return
}

// QUndoGroup(class QObject *)
func NewQUndoGroup(args ...interface{}) *QUndoGroup {
  // QUndoGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroupC1EP7QObject
    // invoke: void QUndoGroup(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QUndoGroupC2EP7QObject(arg0)
    return &QUndoGroup{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoGroup", "QUndoGroup", args)
  }

  return nil // QUndoGroup{qclsinst:qthis}
}

// redoText()
func (this *QUndoGroup) Redotext(args ...interface{}) (ret interface{}) {
  // redoText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup8redoTextEv
    // invoke: QString redoText()
    var ret0 = C.C_ZNK10QUndoGroup8redoTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "redoText", args)
  }

  return
}

// undo()
func (this *QUndoGroup) Undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroup4undoEv
    // invoke: void undo()
    C.C_ZN10QUndoGroup4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "undo", args)
  }

  return
}

// canRedo()
func (this *QUndoGroup) Canredo(args ...interface{}) (ret interface{}) {
  // canRedo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup7canRedoEv
    // invoke: bool canRedo()
    var ret0 = C.C_ZNK10QUndoGroup7canRedoEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "canRedo", args)
  }

  return
}

// addStack(class QUndoStack *)
func (this *QUndoGroup) Addstack(args ...interface{}) () {
  // addStack(class QUndoStack *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroup8addStackEP10QUndoStack
    // invoke: void addStack(class QUndoStack *)
    var arg0 = args[0].(QUndoStack).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoGroup8addStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "addStack", args)
  }

  return
}

// createRedoAction(class QObject *, const class QString &)
func (this *QUndoGroup) Createredoaction(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString
    // invoke: QAction * createRedoAction(class QObject *, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "createRedoAction", args)
  }

  return
}

// metaObject()
func (this *QUndoGroup) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QUndoGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "metaObject", args)
  }

  return
}

// undoText()
func (this *QUndoGroup) Undotext(args ...interface{}) (ret interface{}) {
  // undoText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup8undoTextEv
    // invoke: QString undoText()
    var ret0 = C.C_ZNK10QUndoGroup8undoTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "undoText", args)
  }

  return
}

// removeStack(class QUndoStack *)
func (this *QUndoGroup) Removestack(args ...interface{}) () {
  // removeStack(class QUndoStack *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroup11removeStackEP10QUndoStack
    // invoke: void removeStack(class QUndoStack *)
    var arg0 = args[0].(QUndoStack).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoGroup11removeStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "removeStack", args)
  }

  return
}

// setActiveStack(class QUndoStack *)
func (this *QUndoGroup) Setactivestack(args ...interface{}) () {
  // setActiveStack(class QUndoStack *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QUndoGroup14setActiveStackEP10QUndoStack
    // invoke: void setActiveStack(class QUndoStack *)
    var arg0 = args[0].(QUndoStack).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QUndoGroup14setActiveStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "setActiveStack", args)
  }

  return
}

// stacks()
func (this *QUndoGroup) Stacks(args ...interface{}) () {
  // stacks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup6stacksEv
    // invoke: QList<QUndoStack *> stacks()
    C.C_ZNK10QUndoGroup6stacksEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "stacks", args)
  }

  return
}

// activeStack()
func (this *QUndoGroup) Activestack(args ...interface{}) (ret interface{}) {
  // activeStack()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QUndoGroup11activeStackEv
    // invoke: QUndoStack * activeStack()
    var ret0 = C.C_ZNK10QUndoGroup11activeStackEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUndoGroup", "activeStack", args)
  }

  return
}

// <= body block end

