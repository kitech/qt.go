package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QUndoGroup::addStack(QUndoStack * stack);
extern void _ZN10QUndoGroup8addStackEP10QUndoStack(void* qthis, void* arg0);
  // proto:  void QUndoGroup::undo();
extern void _ZN10QUndoGroup4undoEv(void* qthis);
  // proto:  QList<QUndoStack *> QUndoGroup::stacks();
extern void _ZNK10QUndoGroup6stacksEv(void* qthis);
  // proto:  void QUndoGroup::redo();
extern void _ZN10QUndoGroup4redoEv(void* qthis);
  // proto:  void QUndoGroup::QUndoGroup(QObject * parent);
extern void* dector_ZN10QUndoGroupC1EP7QObject(void* arg0);
extern void _ZN10QUndoGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QUndoGroup::setActiveStack(QUndoStack * stack);
extern void _ZN10QUndoGroup14setActiveStackEP10QUndoStack(void* qthis, void* arg0);
  // proto:  QAction * QUndoGroup::createRedoAction(QObject * parent, const QString & prefix);
extern void _ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QUndoGroup::QUndoGroup(const QUndoGroup & );
extern void* dector_ZN10QUndoGroupC1ERKS_(void* arg0);
extern void _ZN10QUndoGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QUndoGroup::metaObject();
extern void _ZNK10QUndoGroup10metaObjectEv(void* qthis);
  // proto:  bool QUndoGroup::canRedo();
extern void _ZNK10QUndoGroup7canRedoEv(void* qthis);
  // proto:  QString QUndoGroup::redoText();
extern void _ZNK10QUndoGroup8redoTextEv(void* qthis);
  // proto:  QUndoStack * QUndoGroup::activeStack();
extern void _ZNK10QUndoGroup11activeStackEv(void* qthis);
  // proto:  QString QUndoGroup::undoText();
extern void _ZNK10QUndoGroup8undoTextEv(void* qthis);
  // proto:  bool QUndoGroup::canUndo();
extern void _ZNK10QUndoGroup7canUndoEv(void* qthis);
  // proto:  void QUndoGroup::~QUndoGroup();
extern void _ZN10QUndoGroupD0Ev(void* qthis);
  // proto:  bool QUndoGroup::isClean();
extern void _ZNK10QUndoGroup7isCleanEv(void* qthis);
  // proto:  QAction * QUndoGroup::createUndoAction(QObject * parent, const QString & prefix);
extern void _ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QUndoGroup::removeStack(QUndoStack * stack);
extern void _ZN10QUndoGroup11removeStackEP10QUndoStack(void* qthis, void* arg0);
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

  // proto:  void QUndoGroup::addStack(QUndoStack * stack);
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
    C._ZN10QUndoGroup8addStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "addStack", args)
  }

}

  // proto:  void QUndoGroup::undo();
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
    C._ZN10QUndoGroup4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "undo", args)
  }

}

  // proto:  QList<QUndoStack *> QUndoGroup::stacks();
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
    C._ZNK10QUndoGroup6stacksEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "stacks", args)
  }

}

  // proto:  void QUndoGroup::redo();
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
    C._ZN10QUndoGroup4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "redo", args)
  }

}

  // proto:  void QUndoGroup::QUndoGroup(QObject * parent);
func NewQUndoGroup(args ...interface{}) QUndoGroup {
  return QUndoGroup{}
}

  // proto:  void QUndoGroup::setActiveStack(QUndoStack * stack);
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
    C._ZN10QUndoGroup14setActiveStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "setActiveStack", args)
  }

}

  // proto:  QAction * QUndoGroup::createRedoAction(QObject * parent, const QString & prefix);
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
    C._ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUndoGroup", "createRedoAction", args)
  }

}

  // proto:  const QMetaObject * QUndoGroup::metaObject();
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
    C._ZNK10QUndoGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "metaObject", args)
  }

}

  // proto:  bool QUndoGroup::canRedo();
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
    C._ZNK10QUndoGroup7canRedoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "canRedo", args)
  }

}

  // proto:  QString QUndoGroup::redoText();
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
    C._ZNK10QUndoGroup8redoTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "redoText", args)
  }

}

  // proto:  QUndoStack * QUndoGroup::activeStack();
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
    C._ZNK10QUndoGroup11activeStackEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "activeStack", args)
  }

}

  // proto:  QString QUndoGroup::undoText();
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
    C._ZNK10QUndoGroup8undoTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "undoText", args)
  }

}

  // proto:  bool QUndoGroup::canUndo();
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
    C._ZNK10QUndoGroup7canUndoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "canUndo", args)
  }

}

  // proto:  void QUndoGroup::~QUndoGroup();
func (this *QUndoGroup) FreeQUndoGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUndoGroup", "~QUndoGroup", args)
  }

}

  // proto:  bool QUndoGroup::isClean();
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
    C._ZNK10QUndoGroup7isCleanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoGroup", "isClean", args)
  }

}

  // proto:  QAction * QUndoGroup::createUndoAction(QObject * parent, const QString & prefix);
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
    C._ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUndoGroup", "createUndoAction", args)
  }

}

  // proto:  void QUndoGroup::removeStack(QUndoStack * stack);
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
    C._ZN10QUndoGroup11removeStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoGroup", "removeStack", args)
  }

}

// <= body block end

