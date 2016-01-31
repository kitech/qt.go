package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
import "qtrt"
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
extern void C_ZNK15QSessionManager10sessionKeyEv(void* qthis); // 4
  // proto:  bool QSessionManager::isPhase2();
extern void C_ZNK15QSessionManager8isPhase2Ev(void* qthis); // 4
  // proto:  void QSessionManager::setDiscardCommand(const QStringList & );
extern void C_ZN15QSessionManager17setDiscardCommandERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  bool QSessionManager::allowsInteraction();
extern void C_ZN15QSessionManager17allowsInteractionEv(void* qthis); // 4
  // proto:  QSessionManager::RestartHint QSessionManager::restartHint();
extern void C_ZNK15QSessionManager11restartHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QSessionManager::metaObject();
extern void C_ZNK15QSessionManager10metaObjectEv(void* qthis); // 4
  // proto:  bool QSessionManager::allowsErrorInteraction();
extern void C_ZN15QSessionManager22allowsErrorInteractionEv(void* qthis); // 4
  // proto:  QString QSessionManager::sessionId();
extern void C_ZNK15QSessionManager9sessionIdEv(void* qthis); // 4
  // proto:  void QSessionManager::release();
extern void C_ZN15QSessionManager7releaseEv(void* qthis); // 4
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

// class sizeof(QSessionManager)=1
type QSessionManager struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// cancel()
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
    // invoke: void cancel()
    C.C_ZN15QSessionManager6cancelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "cancel", args)
  }

}

// restartCommand()
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
    // invoke: QStringList restartCommand()
    C.C_ZNK15QSessionManager14restartCommandEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "restartCommand", args)
  }

}

// setManagerProperty(const class QString &, const class QStringList &)
func (this *QSessionManager) setManagerProperty(args ...interface{}) () {
  // setManagerProperty(const class QString &, const class QStringList &)
  // setManagerProperty(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList
    // invoke: void setManagerProperty(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringS2_
    // invoke: void setManagerProperty(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QSessionManager18setManagerPropertyERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSessionManager", "setManagerProperty", args)
  }

}

// requestPhase2()
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
    // invoke: void requestPhase2()
    C.C_ZN15QSessionManager13requestPhase2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "requestPhase2", args)
  }

}

// setRestartCommand(const class QStringList &)
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
    // invoke: void setRestartCommand(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSessionManager17setRestartCommandERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setRestartCommand", args)
  }

}

// discardCommand()
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
    // invoke: QStringList discardCommand()
    C.C_ZNK15QSessionManager14discardCommandEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "discardCommand", args)
  }

}

// sessionKey()
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
    // invoke: QString sessionKey()
    var ret = C.C_ZNK15QSessionManager10sessionKeyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionKey", args)
  }

}

// isPhase2()
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
    // invoke: bool isPhase2()
    var ret = C.C_ZNK15QSessionManager8isPhase2Ev(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSessionManager", "isPhase2", args)
  }

}

// setDiscardCommand(const class QStringList &)
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
    // invoke: void setDiscardCommand(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSessionManager17setDiscardCommandERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setDiscardCommand", args)
  }

}

// allowsInteraction()
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
    // invoke: bool allowsInteraction()
    var ret = C.C_ZN15QSessionManager17allowsInteractionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsInteraction", args)
  }

}

// restartHint()
func (this *QSessionManager) restartHint(args ...interface{}) () {
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
    C.C_ZNK15QSessionManager11restartHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "restartHint", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QSessionManager10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "metaObject", args)
  }

}

// allowsErrorInteraction()
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
    // invoke: bool allowsErrorInteraction()
    var ret = C.C_ZN15QSessionManager22allowsErrorInteractionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsErrorInteraction", args)
  }

}

// sessionId()
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
    // invoke: QString sessionId()
    var ret = C.C_ZNK15QSessionManager9sessionIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionId", args)
  }

}

// release()
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
    // invoke: void release()
    C.C_ZN15QSessionManager7releaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "release", args)
  }

}

// <= body block end

