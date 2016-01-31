package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QUndoGroup::isClean();
extern void C_ZNK10QUndoGroup7isCleanEv(void* qthis); // 4
  // proto:  void QUndoGroup::redo();
extern void C_ZN10QUndoGroup4redoEv(void* qthis); // 4
  // proto:  bool QUndoGroup::canUndo();
extern void C_ZNK10QUndoGroup7canUndoEv(void* qthis); // 4
  // proto:  void QUndoGroup::~QUndoGroup();
extern void C_ZN10QUndoGroupD2Ev(void* qthis); // 4
  // proto:  void QUndoGroup::QUndoGroup(QObject * parent);
extern void C_ZN10QUndoGroupC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QString QUndoGroup::redoText();
extern void C_ZNK10QUndoGroup8redoTextEv(void* qthis); // 4
  // proto:  void QUndoGroup::undo();
extern void C_ZN10QUndoGroup4undoEv(void* qthis); // 4
  // proto:  bool QUndoGroup::canRedo();
extern void C_ZNK10QUndoGroup7canRedoEv(void* qthis); // 4
  // proto:  void QUndoGroup::addStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup8addStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  QAction * QUndoGroup::createRedoAction(QObject * parent, const QString & prefix);
extern void C_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QUndoGroup::metaObject();
extern void C_ZNK10QUndoGroup10metaObjectEv(void* qthis); // 4
  // proto:  QString QUndoGroup::undoText();
extern void C_ZNK10QUndoGroup8undoTextEv(void* qthis); // 4
  // proto:  void QUndoGroup::removeStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup11removeStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  void QUndoGroup::setActiveStack(QUndoStack * stack);
extern void C_ZN10QUndoGroup14setActiveStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  QList<QUndoStack *> QUndoGroup::stacks();
extern void C_ZNK10QUndoGroup6stacksEv(void* qthis); // 4
  // proto:  QUndoStack * QUndoGroup::activeStack();
extern void C_ZNK10QUndoGroup11activeStackEv(void* qthis); // 4
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
func (this *QUndoGroup) createUndoAction(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUndoGroup", "createUndoAction", args)
  }

}

// isClean()
func (this *QUndoGroup) isClean(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup7isCleanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "isClean", args)
  }

}

// redo()
func (this *QUndoGroup) redo(args ...interface{}) () {
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

}

// canUndo()
func (this *QUndoGroup) canUndo(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup7canUndoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "canUndo", args)
  }

}

// ~QUndoGroup()
func (this *QUndoGroup) FreeQUndoGroup(args ...interface{}) () {
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

}

// QUndoGroup(class QObject *)
func NewQUndoGroup(args ...interface{}) QUndoGroup {
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
    C.C_ZN10QUndoGroupC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "QUndoGroup", args)
  }

  return QUndoGroup{}
}

// redoText()
func (this *QUndoGroup) redoText(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup8redoTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "redoText", args)
  }

}

// undo()
func (this *QUndoGroup) undo(args ...interface{}) () {
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

}

// canRedo()
func (this *QUndoGroup) canRedo(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup7canRedoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "canRedo", args)
  }

}

// addStack(class QUndoStack *)
func (this *QUndoGroup) addStack(args ...interface{}) () {
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

}

// createRedoAction(class QObject *, const class QString &)
func (this *QUndoGroup) createRedoAction(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUndoGroup", "createRedoAction", args)
  }

}

// metaObject()
func (this *QUndoGroup) metaObject(args ...interface{}) () {
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

}

// undoText()
func (this *QUndoGroup) undoText(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup8undoTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "undoText", args)
  }

}

// removeStack(class QUndoStack *)
func (this *QUndoGroup) removeStack(args ...interface{}) () {
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

}

// setActiveStack(class QUndoStack *)
func (this *QUndoGroup) setActiveStack(args ...interface{}) () {
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

}

// stacks()
func (this *QUndoGroup) stacks(args ...interface{}) () {
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

}

// activeStack()
func (this *QUndoGroup) activeStack(args ...interface{}) () {
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
    C.C_ZNK10QUndoGroup11activeStackEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "activeStack", args)
  }

}

// <= body block end

