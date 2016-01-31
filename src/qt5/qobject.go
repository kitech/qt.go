package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qobject.h
// dst-file: /src/core/qobject.go
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
  // proto:  void QSignalBlocker::~QSignalBlocker();
extern void C_ZN14QSignalBlockerD2Ev(void* qthis); // 2
  // proto:  void QSignalBlocker::unblock();
extern void C_ZN14QSignalBlocker7unblockEv(void* qthis); // 2
  // proto:  void QSignalBlocker::reblock();
extern void C_ZN14QSignalBlocker7reblockEv(void* qthis); // 2
  // proto:  void QSignalBlocker::QSignalBlocker(QObject & o);
extern void* C_ZN14QSignalBlockerC2ER7QObject(void* arg0); // 1
  // proto:  void QSignalBlocker::QSignalBlocker(QObject * o);
extern void* C_ZN14QSignalBlockerC2EP7QObject(void* arg0); // 1
  // proto:  QMetaObject * QObjectData::dynamicMetaObject();
extern void C_ZNK11QObjectData17dynamicMetaObjectEv(void* qthis); // 4
  // proto:  void QObjectUserData::~QObjectUserData();
extern void C_ZN15QObjectUserDataD2Ev(void* qthis); // 4
  // proto:  bool QObject::inherits(const char * classname);
extern void C_ZNK7QObject8inheritsEPKc(void* qthis, unsigned char* arg0); // 2
  // proto:  QList<QByteArray> QObject::dynamicPropertyNames();
extern void C_ZNK7QObject20dynamicPropertyNamesEv(void* qthis); // 4
  // proto:  void QObject::installEventFilter(QObject * );
extern void C_ZN7QObject18installEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  void QObject::dumpObjectTree();
extern void C_ZN7QObject14dumpObjectTreeEv(void* qthis); // 4
  // proto:  void QObject::deleteLater();
extern void C_ZN7QObject11deleteLaterEv(void* qthis); // 4
  // proto:  void QObject::~QObject();
extern void C_ZN7QObjectD2Ev(void* qthis); // 4
  // proto:  QObject * QObject::parent();
extern void C_ZNK7QObject6parentEv(void* qthis); // 2
  // proto:  bool QObject::event(QEvent * );
extern void C_ZN7QObject5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QObject::moveToThread(QThread * thread);
extern void C_ZN7QObject12moveToThreadEP7QThread(void* qthis, void* arg0); // 4
  // proto: static bool QObject::disconnect(const QObject * sender, const char * signal, const QObject * receiver, const char * member);
extern void C_ZN7QObject10disconnectEPKS_PKcS1_S3_(void* arg0, unsigned char* arg1, void* arg2, unsigned char* arg3); // 4
  // proto:  bool QObject::disconnect(const char * signal, const QObject * receiver, const char * member);
extern void C_ZNK7QObject10disconnectEPKcPKS_S1_(void* qthis, unsigned char* arg0, void* arg1, unsigned char* arg2); // 2
  // proto:  bool QObject::disconnect(const QObject * receiver, const char * member);
extern void C_ZNK7QObject10disconnectEPKS_PKc(void* qthis, void* arg0, unsigned char* arg1); // 2
  // proto: static bool QObject::disconnect(const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & member);
extern void C_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_(void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  bool QObject::eventFilter(QObject * , QEvent * );
extern void C_ZN7QObject11eventFilterEPS_P6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QObject::setParent(QObject * );
extern void C_ZN7QObject9setParentEPS_(void* qthis, void* arg0); // 4
  // proto:  void QObject::QObject(QObject * parent);
extern void* C_ZN7QObjectC2EPS_(void* arg0); // 3
  // proto:  bool QObject::isWidgetType();
extern void C_ZNK7QObject12isWidgetTypeEv(void* qthis); // 2
  // proto:  QObjectUserData * QObject::userData(uint id);
extern void C_ZNK7QObject8userDataEj(void* qthis, int32_t arg0); // 4
  // proto:  bool QObject::blockSignals(bool b);
extern void C_ZN7QObject12blockSignalsEb(void* qthis, bool arg0); // 4
  // proto:  QString QObject::objectName();
extern void C_ZNK7QObject10objectNameEv(void* qthis); // 4
  // proto:  bool QObject::isWindowType();
extern void C_ZNK7QObject12isWindowTypeEv(void* qthis); // 2
  // proto:  bool QObject::signalsBlocked();
extern void C_ZNK7QObject14signalsBlockedEv(void* qthis); // 2
  // proto: static uint QObject::registerUserData();
extern void C_ZN7QObject16registerUserDataEv(); // 4
  // proto:  bool QObject::setProperty(const char * name, const QVariant & value);
extern void C_ZN7QObject11setPropertyEPKcRK8QVariant(void* qthis, unsigned char* arg0, void* arg1); // 4
  // proto:  void QObject::setObjectName(const QString & name);
extern void C_ZN7QObject13setObjectNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QObject::setUserData(uint id, QObjectUserData * data);
extern void C_ZN7QObject11setUserDataEjP15QObjectUserData(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QObject::metaObject();
extern void C_ZNK7QObject10metaObjectEv(void* qthis); // 4
  // proto:  QThread * QObject::thread();
extern void C_ZNK7QObject6threadEv(void* qthis); // 4
  // proto:  void QObject::removeEventFilter(QObject * );
extern void C_ZN7QObject17removeEventFilterEPS_(void* qthis, void* arg0); // 4
  // proto:  const QObjectList & QObject::children();
extern void C_ZNK7QObject8childrenEv(void* qthis); // 2
  // proto:  void QObject::killTimer(int id);
extern void C_ZN7QObject9killTimerEi(void* qthis, int32_t arg0); // 4
  // proto:  void QObject::dumpObjectInfo();
extern void C_ZN7QObject14dumpObjectInfoEv(void* qthis); // 4
  // proto:  QVariant QObject::property(const char * name);
extern void C_ZNK7QObject8propertyEPKc(void* qthis, unsigned char* arg0); // 4
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

// class sizeof(QSignalBlocker)=16
type QSignalBlocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QObjectData)=1
type QObjectData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QObjectUserData)=8
type QObjectUserData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QObject)=1
type QObject struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
//  _destroyed QObject_destroyed_signal;
//  _objectNameChanged QObject_objectNameChanged_signal;
}

// ~QSignalBlocker()
func (this *QSignalBlocker) FreeQSignalBlocker(args ...interface{}) () {
  // ~QSignalBlocker()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSignalBlockerD0Ev
    // invoke: void ~QSignalBlocker()
    C.C_ZN14QSignalBlockerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalBlocker", "~QSignalBlocker", args)
  }

}

// unblock()
func (this *QSignalBlocker) unblock(args ...interface{}) () {
  // unblock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSignalBlocker7unblockEv
    // invoke: void unblock()
    C.C_ZN14QSignalBlocker7unblockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalBlocker", "unblock", args)
  }

}

// reblock()
func (this *QSignalBlocker) reblock(args ...interface{}) () {
  // reblock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSignalBlocker7reblockEv
    // invoke: void reblock()
    C.C_ZN14QSignalBlocker7reblockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalBlocker", "reblock", args)
  }

}

// QSignalBlocker(class QObject &)
func NewQSignalBlocker(args ...interface{}) *QSignalBlocker {
  // QSignalBlocker(class QObject &)
  // QSignalBlocker(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSignalBlockerC1ER7QObject
    // invoke: void QSignalBlocker(class QObject &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QSignalBlockerC2ER7QObject(arg0)
    return &QSignalBlocker{qclsinst:qthis}
  case 1:
    // invoke: _ZN14QSignalBlockerC1EP7QObject
    // invoke: void QSignalBlocker(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QSignalBlockerC2EP7QObject(arg0)
    return &QSignalBlocker{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSignalBlocker", "QSignalBlocker", args)
  }

  return nil // QSignalBlocker{qclsinst:qthis}
}

// dynamicMetaObject()
func (this *QObjectData) dynamicMetaObject(args ...interface{}) () {
  // dynamicMetaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QObjectData17dynamicMetaObjectEv
    // invoke: QMetaObject * dynamicMetaObject()
    C.C_ZNK11QObjectData17dynamicMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectData", "dynamicMetaObject", args)
  }

}

// ~QObjectUserData()
func (this *QObjectUserData) FreeQObjectUserData(args ...interface{}) () {
  // ~QObjectUserData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QObjectUserDataD0Ev
    // invoke: void ~QObjectUserData()
    C.C_ZN15QObjectUserDataD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectUserData", "~QObjectUserData", args)
  }

}

// inherits(const char *)
func (this *QObject) inherits(args ...interface{}) () {
  // inherits(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject8inheritsEPKc
    // invoke: bool inherits(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QObject8inheritsEPKc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "inherits", args)
  }

}

// dynamicPropertyNames()
func (this *QObject) dynamicPropertyNames(args ...interface{}) () {
  // dynamicPropertyNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject20dynamicPropertyNamesEv
    // invoke: QList<QByteArray> dynamicPropertyNames()
    C.C_ZNK7QObject20dynamicPropertyNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dynamicPropertyNames", args)
  }

}

// installEventFilter(class QObject *)
func (this *QObject) installEventFilter(args ...interface{}) () {
  // installEventFilter(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject18installEventFilterEPS_
    // invoke: void installEventFilter(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject18installEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "installEventFilter", args)
  }

}

// dumpObjectTree()
func (this *QObject) dumpObjectTree(args ...interface{}) () {
  // dumpObjectTree()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject14dumpObjectTreeEv
    // invoke: void dumpObjectTree()
    C.C_ZN7QObject14dumpObjectTreeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectTree", args)
  }

}

// deleteLater()
func (this *QObject) deleteLater(args ...interface{}) () {
  // deleteLater()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject11deleteLaterEv
    // invoke: void deleteLater()
    C.C_ZN7QObject11deleteLaterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "deleteLater", args)
  }

}

// ~QObject()
func (this *QObject) FreeQObject(args ...interface{}) () {
  // ~QObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObjectD0Ev
    // invoke: void ~QObject()
    C.C_ZN7QObjectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "~QObject", args)
  }

}

// parent()
func (this *QObject) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject6parentEv
    // invoke: QObject * parent()
    var ret = C.C_ZNK7QObject6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "parent", args)
  }

}

// event(class QEvent *)
func (this *QObject) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QObject5eventEP6QEvent(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "event", args)
  }

}

// moveToThread(class QThread *)
func (this *QObject) moveToThread(args ...interface{}) () {
  // moveToThread(class QThread *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QThread{}) // "QThread *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject12moveToThreadEP7QThread
    // invoke: void moveToThread(class QThread *)
    var arg0 = args[0].(QThread).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject12moveToThreadEP7QThread(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "moveToThread", args)
  }

}

// disconnect(const class QObject *, const char *, const class QObject *, const char *)
func (this *QObject) disconnect_s(args ...interface{}) () {
  // disconnect(const class QObject *, const char *, const class QObject *, const char *)
  // disconnect(const char *, const class QObject *, const char *)
  // disconnect(const class QObject *, const char *)
  // disconnect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][1] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"
  vtys[3][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][3] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject10disconnectEPKS_PKcS1_S3_
    // invoke: bool disconnect(const class QObject *, const char *, const class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[3].([]byte)).Pointer()))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN7QObject10disconnectEPKS_PKcS1_S3_(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK7QObject10disconnectEPKcPKS_S1_
    // invoke: bool disconnect(const char *, const class QObject *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZNK7QObject10disconnectEPKcPKS_S1_(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK7QObject10disconnectEPKS_PKc
    // invoke: bool disconnect(const class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK7QObject10disconnectEPKS_PKc(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_
    // invoke: bool disconnect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMetaMethod).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QMetaMethod).qclsinst
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "disconnect", args)
  }

}

// eventFilter(class QObject *, class QEvent *)
func (this *QObject) eventFilter(args ...interface{}) () {
  // eventFilter(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject11eventFilterEPS_P6QEvent
    // invoke: bool eventFilter(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN7QObject11eventFilterEPS_P6QEvent(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "eventFilter", args)
  }

}

// setParent(class QObject *)
func (this *QObject) setParent(args ...interface{}) () {
  // setParent(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject9setParentEPS_
    // invoke: void setParent(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "setParent", args)
  }

}

// QObject(class QObject *)
func NewQObject(args ...interface{}) *QObject {
  // QObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObjectC1EPS_
    // invoke: void QObject(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QObjectC2EPS_(arg0)
    return &QObject{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QObject", "QObject", args)
  }

  return nil // QObject{qclsinst:qthis}
}

// isWidgetType()
func (this *QObject) isWidgetType(args ...interface{}) () {
  // isWidgetType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject12isWidgetTypeEv
    // invoke: bool isWidgetType()
    var ret = C.C_ZNK7QObject12isWidgetTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "isWidgetType", args)
  }

}

// userData(uint)
func (this *QObject) userData(args ...interface{}) () {
  // userData(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject8userDataEj
    // invoke: QObjectUserData * userData(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QObject8userDataEj(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "userData", args)
  }

}

// blockSignals(_Bool)
func (this *QObject) blockSignals(args ...interface{}) () {
  // blockSignals(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject12blockSignalsEb
    // invoke: bool blockSignals(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QObject12blockSignalsEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "blockSignals", args)
  }

}

// objectName()
func (this *QObject) objectName(args ...interface{}) () {
  // objectName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject10objectNameEv
    // invoke: QString objectName()
    var ret = C.C_ZNK7QObject10objectNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "objectName", args)
  }

}

// isWindowType()
func (this *QObject) isWindowType(args ...interface{}) () {
  // isWindowType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject12isWindowTypeEv
    // invoke: bool isWindowType()
    var ret = C.C_ZNK7QObject12isWindowTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "isWindowType", args)
  }

}

// signalsBlocked()
func (this *QObject) signalsBlocked(args ...interface{}) () {
  // signalsBlocked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject14signalsBlockedEv
    // invoke: bool signalsBlocked()
    var ret = C.C_ZNK7QObject14signalsBlockedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "signalsBlocked", args)
  }

}

// registerUserData()
func (this *QObject) registerUserData_s(args ...interface{}) () {
  // registerUserData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject16registerUserDataEv
    // invoke: uint registerUserData()
    var ret = C.C_ZN7QObject16registerUserDataEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "registerUserData", args)
  }

}

// setProperty(const char *, const class QVariant &)
func (this *QObject) setProperty(args ...interface{}) () {
  // setProperty(const char *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject11setPropertyEPKcRK8QVariant
    // invoke: bool setProperty(const char *, const class QVariant &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN7QObject11setPropertyEPKcRK8QVariant(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "setProperty", args)
  }

}

// setObjectName(const class QString &)
func (this *QObject) setObjectName(args ...interface{}) () {
  // setObjectName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject13setObjectNameERK7QString
    // invoke: void setObjectName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject13setObjectNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "setObjectName", args)
  }

}

// setUserData(uint, class QObjectUserData *)
func (this *QObject) setUserData(args ...interface{}) () {
  // setUserData(uint, class QObjectUserData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = reflect.TypeOf(QObjectUserData{}) // "QObjectUserData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject11setUserDataEjP15QObjectUserData
    // invoke: void setUserData(uint, class QObjectUserData *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObjectUserData).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QObject11setUserDataEjP15QObjectUserData(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QObject", "setUserData", args)
  }

}

// metaObject()
func (this *QObject) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "metaObject", args)
  }

}

// thread()
func (this *QObject) thread(args ...interface{}) () {
  // thread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject6threadEv
    // invoke: QThread * thread()
    var ret = C.C_ZNK7QObject6threadEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "thread", args)
  }

}

// removeEventFilter(class QObject *)
func (this *QObject) removeEventFilter(args ...interface{}) () {
  // removeEventFilter(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject17removeEventFilterEPS_
    // invoke: void removeEventFilter(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject17removeEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "removeEventFilter", args)
  }

}

// children()
func (this *QObject) children(args ...interface{}) () {
  // children()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject8childrenEv
    // invoke: const QObjectList & children()
    C.C_ZNK7QObject8childrenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "children", args)
  }

}

// killTimer(int)
func (this *QObject) killTimer(args ...interface{}) () {
  // killTimer(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject9killTimerEi
    // invoke: void killTimer(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QObject9killTimerEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "killTimer", args)
  }

}

// dumpObjectInfo()
func (this *QObject) dumpObjectInfo(args ...interface{}) () {
  // dumpObjectInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QObject14dumpObjectInfoEv
    // invoke: void dumpObjectInfo()
    C.C_ZN7QObject14dumpObjectInfoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectInfo", args)
  }

}

// property(const char *)
func (this *QObject) property(args ...interface{}) () {
  // property(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QObject8propertyEPKc
    // invoke: QVariant property(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QObject8propertyEPKc(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QObject", "property", args)
  }

}

// <= body block end

