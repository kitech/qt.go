package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  QString QSessionManager::sessionId();
extern void _ZNK15QSessionManager9sessionIdEv(void* qthis);
  // proto:  QString QSessionManager::sessionKey();
extern void _ZNK15QSessionManager10sessionKeyEv(void* qthis);
  // proto:  void QSessionManager::setRestartCommand(const QStringList & );
extern void _ZN15QSessionManager17setRestartCommandERK11QStringList(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSessionManager::metaObject();
extern void _ZNK15QSessionManager10metaObjectEv(void* qthis);
  // proto:  bool QSessionManager::allowsErrorInteraction();
extern void _ZN15QSessionManager22allowsErrorInteractionEv(void* qthis);
  // proto:  void QSessionManager::~QSessionManager();
extern void _ZN15QSessionManagerD0Ev(void* qthis);
  // proto:  QStringList QSessionManager::restartCommand();
extern void _ZNK15QSessionManager14restartCommandEv(void* qthis);
  // proto:  void QSessionManager::requestPhase2();
extern void _ZN15QSessionManager13requestPhase2Ev(void* qthis);
  // proto:  bool QSessionManager::isPhase2();
extern void _ZNK15QSessionManager8isPhase2Ev(void* qthis);
  // proto:  void QSessionManager::release();
extern void _ZN15QSessionManager7releaseEv(void* qthis);
  // proto:  void QSessionManager::setManagerProperty(const QString & name, const QString & value);
extern void _ZN15QSessionManager18setManagerPropertyERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QSessionManager::setManagerProperty(const QString & name, const QStringList & value);
extern void _ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList(void* qthis, void* arg0, void* arg1);
  // proto:  QStringList QSessionManager::discardCommand();
extern void _ZNK15QSessionManager14discardCommandEv(void* qthis);
  // proto:  void QSessionManager::QSessionManager(QGuiApplication * app, QString & id, QString & key);
extern void* dector_ZN15QSessionManagerC1EP15QGuiApplicationR7QStringS3_(void* arg0, void* arg1, void* arg2);
extern void _ZN15QSessionManagerC1EP15QGuiApplicationR7QStringS3_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QSessionManager::cancel();
extern void _ZN15QSessionManager6cancelEv(void* qthis);
  // proto:  void QSessionManager::setDiscardCommand(const QStringList & );
extern void _ZN15QSessionManager17setDiscardCommandERK11QStringList(void* qthis, void* arg0);
  // proto:  bool QSessionManager::allowsInteraction();
extern void _ZN15QSessionManager17allowsInteractionEv(void* qthis);
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

  // proto:  QString QSessionManager::sessionId();
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
    C._ZNK15QSessionManager9sessionIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionId", args)
  }

}

  // proto:  QString QSessionManager::sessionKey();
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
    C._ZNK15QSessionManager10sessionKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "sessionKey", args)
  }

}

  // proto:  void QSessionManager::setRestartCommand(const QStringList & );
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
    C._ZN15QSessionManager17setRestartCommandERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setRestartCommand", args)
  }

}

  // proto:  const QMetaObject * QSessionManager::metaObject();
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
    C._ZNK15QSessionManager10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "metaObject", args)
  }

}

  // proto:  bool QSessionManager::allowsErrorInteraction();
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
    C._ZN15QSessionManager22allowsErrorInteractionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsErrorInteraction", args)
  }

}

  // proto:  void QSessionManager::~QSessionManager();
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

  // proto:  QStringList QSessionManager::restartCommand();
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
    C._ZNK15QSessionManager14restartCommandEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "restartCommand", args)
  }

}

  // proto:  void QSessionManager::requestPhase2();
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
    C._ZN15QSessionManager13requestPhase2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "requestPhase2", args)
  }

}

  // proto:  bool QSessionManager::isPhase2();
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
    C._ZNK15QSessionManager8isPhase2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "isPhase2", args)
  }

}

  // proto:  void QSessionManager::release();
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
    C._ZN15QSessionManager7releaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "release", args)
  }

}

  // proto:  void QSessionManager::setManagerProperty(const QString & name, const QString & value);
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
    // invoke: void setManagerProperty(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN15QSessionManager18setManagerPropertyERK7QStringS2_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList
    // invoke: void setManagerProperty(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSessionManager", "setManagerProperty", args)
  }

}

  // proto:  QStringList QSessionManager::discardCommand();
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
    C._ZNK15QSessionManager14discardCommandEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "discardCommand", args)
  }

}

  // proto:  void QSessionManager::QSessionManager(QGuiApplication * app, QString & id, QString & key);
func NewQSessionManager(args ...interface{}) QSessionManager {
  return QSessionManager{}
}

  // proto:  void QSessionManager::cancel();
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
    C._ZN15QSessionManager6cancelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "cancel", args)
  }

}

  // proto:  void QSessionManager::setDiscardCommand(const QStringList & );
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
    C._ZN15QSessionManager17setDiscardCommandERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSessionManager", "setDiscardCommand", args)
  }

}

  // proto:  bool QSessionManager::allowsInteraction();
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
    C._ZN15QSessionManager17allowsInteractionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSessionManager", "allowsInteraction", args)
  }

}

// <= body block end

