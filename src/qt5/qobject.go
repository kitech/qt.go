package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QSignalBlocker::unblock();
extern void demth_ZN14QSignalBlocker7unblockEv(void* qthis);
  // proto:  void QSignalBlocker::QSignalBlocker(QObject & o);
extern void* dector_ZN14QSignalBlockerC1ER7QObject(void* arg0);
extern void demth_ZN14QSignalBlockerC1ER7QObject(void* qthis, void* arg0);
  // proto:  void QSignalBlocker::QSignalBlocker(QObject * o);
extern void* dector_ZN14QSignalBlockerC1EP7QObject(void* arg0);
extern void demth_ZN14QSignalBlockerC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QSignalBlocker::QSignalBlocker(const QSignalBlocker & );
extern void* dector_ZN14QSignalBlockerC1ERKS_(void* arg0);
extern void _ZN14QSignalBlockerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSignalBlocker::reblock();
extern void demth_ZN14QSignalBlocker7reblockEv(void* qthis);
  // proto:  void QSignalBlocker::~QSignalBlocker();
extern void demth_ZN14QSignalBlockerD0Ev(void* qthis);
  // proto:  QMetaObject * QObjectData::dynamicMetaObject();
extern void _ZNK11QObjectData17dynamicMetaObjectEv(void* qthis);
  // proto:  void QObjectUserData::~QObjectUserData();
extern void _ZN15QObjectUserDataD0Ev(void* qthis);
  // proto:  bool QObject::inherits(const char * classname);
extern void demth_ZNK7QObject8inheritsEPKc(void* qthis, unsigned char* arg0);
  // proto:  void QObject::moveToThread(QThread * thread);
extern void _ZN7QObject12moveToThreadEP7QThread(void* qthis, void* arg0);
  // proto:  void QObject::removeEventFilter(QObject * );
extern void _ZN7QObject17removeEventFilterEPS_(void* qthis, void* arg0);
  // proto:  void QObject::dumpObjectTree();
extern void _ZN7QObject14dumpObjectTreeEv(void* qthis);
  // proto:  bool QObject::eventFilter(QObject * , QEvent * );
extern void _ZN7QObject11eventFilterEPS_P6QEvent(void* qthis, void* arg0, void* arg1);
  // proto:  void QObject::setUserData(uint id, QObjectUserData * data);
extern void _ZN7QObject11setUserDataEjP15QObjectUserData(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QObject::QObject(const QObject & );
extern void* dector_ZN7QObjectC1ERKS_(void* arg0);
extern void _ZN7QObjectC1ERKS_(void* qthis, void* arg0);
  // proto: static bool QObject::disconnect(const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & member);
extern void _ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_(void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  bool QObject::event(QEvent * );
extern void _ZN7QObject5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  QList<QByteArray> QObject::dynamicPropertyNames();
extern void _ZNK7QObject20dynamicPropertyNamesEv(void* qthis);
  // proto:  bool QObject::isWidgetType();
extern void demth_ZNK7QObject12isWidgetTypeEv(void* qthis);
  // proto:  QVariant QObject::property(const char * name);
extern void _ZNK7QObject8propertyEPKc(void* qthis, unsigned char* arg0);
  // proto:  QThread * QObject::thread();
extern void _ZNK7QObject6threadEv(void* qthis);
  // proto:  const QMetaObject * QObject::metaObject();
extern void _ZNK7QObject10metaObjectEv(void* qthis);
  // proto:  void QObject::setParent(QObject * );
extern void _ZN7QObject9setParentEPS_(void* qthis, void* arg0);
  // proto:  bool QObject::disconnect(const QObject * receiver, const char * member);
extern void demth_ZNK7QObject10disconnectEPKS_PKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  const QObjectList & QObject::children();
extern void demth_ZNK7QObject8childrenEv(void* qthis);
  // proto:  bool QObject::isWindowType();
extern void demth_ZNK7QObject12isWindowTypeEv(void* qthis);
  // proto:  bool QObject::disconnect(const char * signal, const QObject * receiver, const char * member);
extern void demth_ZNK7QObject10disconnectEPKcPKS_S1_(void* qthis, unsigned char* arg0, void* arg1, unsigned char* arg2);
  // proto:  void QObject::deleteLater();
extern void _ZN7QObject11deleteLaterEv(void* qthis);
  // proto:  void QObject::~QObject();
extern void _ZN7QObjectD0Ev(void* qthis);
  // proto:  void QObject::QObject(QObject * parent);
extern void* dector_ZN7QObjectC1EPS_(void* arg0);
extern void _ZN7QObjectC1EPS_(void* qthis, void* arg0);
  // proto:  QString QObject::objectName();
extern void _ZNK7QObject10objectNameEv(void* qthis);
  // proto:  bool QObject::setProperty(const char * name, const QVariant & value);
extern void _ZN7QObject11setPropertyEPKcRK8QVariant(void* qthis, unsigned char* arg0, void* arg1);
  // proto: static bool QObject::disconnect(const QObject * sender, const char * signal, const QObject * receiver, const char * member);
extern void _ZN7QObject10disconnectEPKS_PKcS1_S3_(void* arg0, unsigned char* arg1, void* arg2, unsigned char* arg3);
  // proto:  bool QObject::signalsBlocked();
extern void demth_ZNK7QObject14signalsBlockedEv(void* qthis);
  // proto: static uint QObject::registerUserData();
extern void _ZN7QObject16registerUserDataEv();
  // proto:  QObjectUserData * QObject::userData(uint id);
extern void _ZNK7QObject8userDataEj(void* qthis, int32_t arg0);
  // proto:  QObject * QObject::parent();
extern void demth_ZNK7QObject6parentEv(void* qthis);
  // proto:  void QObject::installEventFilter(QObject * );
extern void _ZN7QObject18installEventFilterEPS_(void* qthis, void* arg0);
  // proto:  bool QObject::blockSignals(bool b);
extern void _ZN7QObject12blockSignalsEb(void* qthis, bool arg0);
  // proto:  void QObject::setObjectName(const QString & name);
extern void _ZN7QObject13setObjectNameERK7QString(void* qthis, void* arg0);
  // proto:  void QObject::dumpObjectInfo();
extern void _ZN7QObject14dumpObjectInfoEv(void* qthis);
  // proto:  void QObject::killTimer(int id);
extern void _ZN7QObject9killTimerEi(void* qthis, int32_t arg0);
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

  // proto:  void QSignalBlocker::unblock();
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
    C.demth_ZN14QSignalBlocker7unblockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalBlocker", "unblock", args)
  }

}

  // proto:  void QSignalBlocker::QSignalBlocker(QObject & o);
func NewQSignalBlocker(args ...interface{}) QSignalBlocker {
  return QSignalBlocker{}
}

  // proto:  void QSignalBlocker::reblock();
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
    C.demth_ZN14QSignalBlocker7reblockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalBlocker", "reblock", args)
  }

}

  // proto:  void QSignalBlocker::~QSignalBlocker();
func (this *QSignalBlocker) FreeQSignalBlocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalBlocker", "~QSignalBlocker", args)
  }

}

  // proto:  QMetaObject * QObjectData::dynamicMetaObject();
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
    C._ZNK11QObjectData17dynamicMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObjectData", "dynamicMetaObject", args)
  }

}

  // proto:  void QObjectUserData::~QObjectUserData();
func (this *QObjectUserData) FreeQObjectUserData(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QObjectUserData", "~QObjectUserData", args)
  }

}

  // proto:  bool QObject::inherits(const char * classname);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QObject8inheritsEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "inherits", args)
  }

}

  // proto:  void QObject::moveToThread(QThread * thread);
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
    C._ZN7QObject12moveToThreadEP7QThread(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "moveToThread", args)
  }

}

  // proto:  void QObject::removeEventFilter(QObject * );
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
    C._ZN7QObject17removeEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "removeEventFilter", args)
  }

}

  // proto:  void QObject::dumpObjectTree();
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
    C._ZN7QObject14dumpObjectTreeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectTree", args)
  }

}

  // proto:  bool QObject::eventFilter(QObject * , QEvent * );
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
    C._ZN7QObject11eventFilterEPS_P6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QObject", "eventFilter", args)
  }

}

  // proto:  void QObject::setUserData(uint id, QObjectUserData * data);
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
    C._ZN7QObject11setUserDataEjP15QObjectUserData(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QObject", "setUserData", args)
  }

}

  // proto:  void QObject::QObject(const QObject & );
func NewQObject(args ...interface{}) QObject {
  return QObject{}
}

  // proto: static bool QObject::disconnect(const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & member);
func (this *QObject) disconnect_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QObject", "disconnect", args)
  }

}

  // proto:  bool QObject::event(QEvent * );
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
    C._ZN7QObject5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "event", args)
  }

}

  // proto:  QList<QByteArray> QObject::dynamicPropertyNames();
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
    C._ZNK7QObject20dynamicPropertyNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dynamicPropertyNames", args)
  }

}

  // proto:  bool QObject::isWidgetType();
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
    C.demth_ZNK7QObject12isWidgetTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "isWidgetType", args)
  }

}

  // proto:  QVariant QObject::property(const char * name);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK7QObject8propertyEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "property", args)
  }

}

  // proto:  QThread * QObject::thread();
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
    C._ZNK7QObject6threadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "thread", args)
  }

}

  // proto:  const QMetaObject * QObject::metaObject();
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
    C._ZNK7QObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "metaObject", args)
  }

}

  // proto:  void QObject::setParent(QObject * );
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
    C._ZN7QObject9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "setParent", args)
  }

}

  // proto:  bool QObject::disconnect(const QObject * receiver, const char * member);
func (this *QObject) disconnect(args ...interface{}) () {
  // disconnect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &)
  // disconnect(const struct QMetaObject::Connection &)
  // disconnect(const class QObject *, const char *)
  // disconnect(const char *, const class QObject *, const char *)
  // disconnect(const class QObject *, const char *, const class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QMetaObject::Connection{}) // "const QMetaObject::Connection &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][2] = qtrt.ByteTy(true) // "const char *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[4][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[4][3] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
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
    C._ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_(arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZNK7QObject10disconnectEPKS_PKc
    // invoke: bool disconnect(const class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QObject10disconnectEPKS_PKc(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QObject10disconnectEPKcPKS_S1_
    // invoke: bool disconnect(const char *, const class QObject *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    C.demth_ZNK7QObject10disconnectEPKcPKS_S1_(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN7QObject10disconnectEPKS_PKcS1_S3_
    // invoke: bool disconnect(const class QObject *, const char *, const class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[3].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg3)}
    C._ZN7QObject10disconnectEPKS_PKcS1_S3_(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QObject", "disconnect", args)
  }

}

  // proto:  const QObjectList & QObject::children();
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
    C.demth_ZNK7QObject8childrenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "children", args)
  }

}

  // proto:  bool QObject::isWindowType();
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
    C.demth_ZNK7QObject12isWindowTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "isWindowType", args)
  }

}

  // proto:  void QObject::deleteLater();
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
    C._ZN7QObject11deleteLaterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "deleteLater", args)
  }

}

  // proto:  void QObject::~QObject();
func (this *QObject) FreeQObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QObject", "~QObject", args)
  }

}

  // proto:  QString QObject::objectName();
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
    C._ZNK7QObject10objectNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "objectName", args)
  }

}

  // proto:  bool QObject::setProperty(const char * name, const QVariant & value);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QObject11setPropertyEPKcRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QObject", "setProperty", args)
  }

}

  // proto:  bool QObject::signalsBlocked();
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
    C.demth_ZNK7QObject14signalsBlockedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "signalsBlocked", args)
  }

}

  // proto: static uint QObject::registerUserData();
func (this *QObject) registerUserData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QObject", "registerUserData", args)
  }

}

  // proto:  QObjectUserData * QObject::userData(uint id);
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
    C._ZNK7QObject8userDataEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "userData", args)
  }

}

  // proto:  QObject * QObject::parent();
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
    C.demth_ZNK7QObject6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "parent", args)
  }

}

  // proto:  void QObject::installEventFilter(QObject * );
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
    C._ZN7QObject18installEventFilterEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "installEventFilter", args)
  }

}

  // proto:  bool QObject::blockSignals(bool b);
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
    C._ZN7QObject12blockSignalsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "blockSignals", args)
  }

}

  // proto:  void QObject::setObjectName(const QString & name);
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
    C._ZN7QObject13setObjectNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "setObjectName", args)
  }

}

  // proto:  void QObject::dumpObjectInfo();
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
    C._ZN7QObject14dumpObjectInfoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectInfo", args)
  }

}

  // proto:  void QObject::killTimer(int id);
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
    C._ZN7QObject9killTimerEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QObject", "killTimer", args)
  }

}

// <= body block end

