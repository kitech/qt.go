package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qsessionmanager.h
// dst-file: /src/gui/qsessionmanager.go
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
// class sizeof(QSessionManager)=1
type QSessionManager struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSessionManager) sessionId(args ...interface{}) () {
  // sessionId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager9sessionIdEv
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionId", args)
 }

}


func (this *QSessionManager) sessionKey(args ...interface{}) () {
  // sessionKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager10sessionKeyEv
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionKey", args)
 }

}


func (this *QSessionManager) setRestartCommand(args ...interface{}) () {
  // setRestartCommand(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17setRestartCommandERK11QStringList
  default:
    qtrt.ErrorResolve("QSessionManager", "setRestartCommand", args)
 }

}


func (this *QSessionManager) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager10metaObjectEv
  default:
    qtrt.ErrorResolve("QSessionManager", "metaObject", args)
 }

}


func (this *QSessionManager) allowsErrorInteraction(args ...interface{}) () {
  // allowsErrorInteraction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager22allowsErrorInteractionEv
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsErrorInteraction", args)
 }

}


func (this *QSessionManager) FreeQSessionManager(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSessionManager", "~QSessionManager", args)
 }

}


func (this *QSessionManager) restartCommand(args ...interface{}) () {
  // restartCommand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager14restartCommandEv
  default:
    qtrt.ErrorResolve("QSessionManager", "restartCommand", args)
 }

}


func (this *QSessionManager) requestPhase2(args ...interface{}) () {
  // requestPhase2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager13requestPhase2Ev
  default:
    qtrt.ErrorResolve("QSessionManager", "requestPhase2", args)
 }

}


func (this *QSessionManager) isPhase2(args ...interface{}) () {
  // isPhase2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager8isPhase2Ev
  default:
    qtrt.ErrorResolve("QSessionManager", "isPhase2", args)
 }

}


func (this *QSessionManager) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager7releaseEv
  default:
    qtrt.ErrorResolve("QSessionManager", "release", args)
 }

}


func (this *QSessionManager) setManagerProperty(args ...interface{}) () {
  // setManagerProperty(const class QString &, const class QString &)
  // setManagerProperty(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringS2_
  case 1:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList
  default:
    qtrt.ErrorResolve("QSessionManager", "setManagerProperty", args)
 }

}


func (this *QSessionManager) discardCommand(args ...interface{}) () {
  // discardCommand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager14discardCommandEv
  default:
    qtrt.ErrorResolve("QSessionManager", "discardCommand", args)
 }

}


func NewQSessionManager(args ...interface{})() {
}


func (this *QSessionManager) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager6cancelEv
  default:
    qtrt.ErrorResolve("QSessionManager", "cancel", args)
 }

}


func (this *QSessionManager) setDiscardCommand(args ...interface{}) () {
  // setDiscardCommand(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17setDiscardCommandERK11QStringList
  default:
    qtrt.ErrorResolve("QSessionManager", "setDiscardCommand", args)
 }

}


func (this *QSessionManager) allowsInteraction(args ...interface{}) () {
  // allowsInteraction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17allowsInteractionEv
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsInteraction", args)
 }

}

// <= body block end

