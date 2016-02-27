package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qsessionmanager.h
// dst-file: /src/gui/qsessionmanager.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QSessionManager::cancel();
extern void C_ZN15QSessionManager6cancelEv(void* qthis); // 4
  // proto:  QStringList QSessionManager::restartCommand();
extern void C_ZNK15QSessionManager14restartCommandEv(void* qthis); // 4
  // proto:  void QSessionManager::setManagerProperty(const QString & name, const QStringList & value);
extern void C_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSessionManager::setManagerProperty(const QString & name, const QString & value);
extern void C_ZN15QSessionManager18setManagerPropertyERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSessionManager::requestPhase2();
extern void C_ZN15QSessionManager13requestPhase2Ev(void* qthis); // 4
  // proto:  void QSessionManager::setRestartCommand(const QStringList & );
extern void C_ZN15QSessionManager17setRestartCommandERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QStringList QSessionManager::discardCommand();
extern void C_ZNK15QSessionManager14discardCommandEv(void* qthis); // 4
  // proto:  QString QSessionManager::sessionKey();
extern void* C_ZNK15QSessionManager10sessionKeyEv(void* qthis); // 4
  // proto:  bool QSessionManager::isPhase2();
extern bool C_ZNK15QSessionManager8isPhase2Ev(void* qthis); // 4
  // proto:  void QSessionManager::setDiscardCommand(const QStringList & );
extern void C_ZN15QSessionManager17setDiscardCommandERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  bool QSessionManager::allowsInteraction();
extern bool C_ZN15QSessionManager17allowsInteractionEv(void* qthis); // 4
  // proto:  QSessionManager::RestartHint QSessionManager::restartHint();
extern void C_ZNK15QSessionManager11restartHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QSessionManager::metaObject();
extern void C_ZNK15QSessionManager10metaObjectEv(void* qthis); // 4
  // proto:  bool QSessionManager::allowsErrorInteraction();
extern bool C_ZN15QSessionManager22allowsErrorInteractionEv(void* qthis); // 4
  // proto:  QString QSessionManager::sessionId();
extern void* C_ZNK15QSessionManager9sessionIdEv(void* qthis); // 4
  // proto:  void QSessionManager::release();
extern void C_ZN15QSessionManager7releaseEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QSessionManager)=1
type QSessionManager struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// cancel()
func (this *QSessionManager) Cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager6cancelEv
    // invoke: void cancel()
    C.C_ZN15QSessionManager6cancelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "cancel", args)
  }

  return
}

// restartCommand()
func (this *QSessionManager) RestartCommand(args ...interface{}) () {
  // restartCommand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager14restartCommandEv
    // invoke: QStringList restartCommand()
    C.C_ZNK15QSessionManager14restartCommandEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "restartCommand", args)
  }

  return
}

// setManagerProperty(const class QString &, const class QStringList &)
func (this *QSessionManager) SetManagerProperty(args ...interface{}) () {
  // setManagerProperty(const class QString &, const class QStringList &)
  // setManagerProperty(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList
    // invoke: void setManagerProperty(const class QString &, const class QStringList &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringS2_
    // invoke: void setManagerProperty(const class QString &, const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QSessionManager18setManagerPropertyERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSessionManager", "setManagerProperty", args)
  }

  return
}

// requestPhase2()
func (this *QSessionManager) RequestPhase2(args ...interface{}) () {
  // requestPhase2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager13requestPhase2Ev
    // invoke: void requestPhase2()
    C.C_ZN15QSessionManager13requestPhase2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "requestPhase2", args)
  }

  return
}

// setRestartCommand(const class QStringList &)
func (this *QSessionManager) SetRestartCommand(args ...interface{}) () {
  // setRestartCommand(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17setRestartCommandERK11QStringList
    // invoke: void setRestartCommand(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSessionManager17setRestartCommandERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setRestartCommand", args)
  }

  return
}

// discardCommand()
func (this *QSessionManager) DiscardCommand(args ...interface{}) () {
  // discardCommand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager14discardCommandEv
    // invoke: QStringList discardCommand()
    C.C_ZNK15QSessionManager14discardCommandEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "discardCommand", args)
  }

  return
}

// sessionKey()
func (this *QSessionManager) SessionKey(args ...interface{}) (ret interface{}) {
  // sessionKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager10sessionKeyEv
    // invoke: QString sessionKey()
    var ret0 = C.C_ZNK15QSessionManager10sessionKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionKey", args)
  }

  return
}

// isPhase2()
func (this *QSessionManager) IsPhase2(args ...interface{}) (ret interface{}) {
  // isPhase2()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager8isPhase2Ev
    // invoke: bool isPhase2()
    var ret0 = C.C_ZNK15QSessionManager8isPhase2Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSessionManager", "isPhase2", args)
  }

  return
}

// setDiscardCommand(const class QStringList &)
func (this *QSessionManager) SetDiscardCommand(args ...interface{}) () {
  // setDiscardCommand(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17setDiscardCommandERK11QStringList
    // invoke: void setDiscardCommand(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSessionManager17setDiscardCommandERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setDiscardCommand", args)
  }

  return
}

// allowsInteraction()
func (this *QSessionManager) AllowsInteraction(args ...interface{}) (ret interface{}) {
  // allowsInteraction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager17allowsInteractionEv
    // invoke: bool allowsInteraction()
    var ret0 = C.C_ZN15QSessionManager17allowsInteractionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsInteraction", args)
  }

  return
}

// restartHint()
func (this *QSessionManager) RestartHint(args ...interface{}) () {
  // restartHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager11restartHintEv
    // invoke: QSessionManager::RestartHint restartHint()
    C.C_ZNK15QSessionManager11restartHintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "restartHint", args)
  }

  return
}

// metaObject()
func (this *QSessionManager) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QSessionManager10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "metaObject", args)
  }

  return
}

// allowsErrorInteraction()
func (this *QSessionManager) AllowsErrorInteraction(args ...interface{}) (ret interface{}) {
  // allowsErrorInteraction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager22allowsErrorInteractionEv
    // invoke: bool allowsErrorInteraction()
    var ret0 = C.C_ZN15QSessionManager22allowsErrorInteractionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsErrorInteraction", args)
  }

  return
}

// sessionId()
func (this *QSessionManager) SessionId(args ...interface{}) (ret interface{}) {
  // sessionId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSessionManager9sessionIdEv
    // invoke: QString sessionId()
    var ret0 = C.C_ZNK15QSessionManager9sessionIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionId", args)
  }

  return
}

// release()
func (this *QSessionManager) Release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager7releaseEv
    // invoke: void release()
    C.C_ZN15QSessionManager7releaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "release", args)
  }

  return
}

// <= body block end

