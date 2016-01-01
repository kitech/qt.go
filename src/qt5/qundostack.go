package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qundostack.h
// dst-file: /src/widgets/qundostack.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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
  default:
    qtrt.ErrorResolve("QUndoStack", "command", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "createUndoAction", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "createRedoAction", args)
  }

}


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


func NewQUndoStack(args ...interface{}) QUndoStack {
  return QUndoStack{}
}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "beginMacro", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "setActive", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "setIndex", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "setUndoLimit", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "text", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoStack", "push", args)
  }

}


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


func NewQUndoCommand(args ...interface{}) QUndoCommand {
  return QUndoCommand{}
}


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
  default:
    qtrt.ErrorResolve("QUndoCommand", "mergeWith", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoCommand", "child", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QUndoCommand", "setText", args)
  }

}

// <= body block end

