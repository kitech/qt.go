package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qobject.h
// dst-file: /src/core/qobject.go
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
// class sizeof(QSignalBlocker)=16
type QSignalBlocker struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QObjectData)=1
type QObjectData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QObjectUserData)=8
type QObjectUserData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QObject)=1
type QObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
//  _destroyed QObject_destroyed_signal;
//  _objectNameChanged QObject_objectNameChanged_signal;
}


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
  default:
    qtrt.ErrorResolve("QSignalBlocker", "unblock", args)
 }

}


func NewQSignalBlocker(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QSignalBlocker", "reblock", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObjectData", "dynamicMetaObject", args)
 }

}


func (this *QObjectData) FreeQObjectData(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QObjectData", "~QObjectData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "inherits", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "moveToThread", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "removeEventFilter", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectTree", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "eventFilter", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "setUserData", args)
 }

}


func NewQObject(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QObject", "event", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "dynamicPropertyNames", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "isWidgetType", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "property", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "thread", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "setParent", args)
 }

}


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
  case 1:
    // invoke: _ZN7QObject10disconnectERKN11QMetaObject10ConnectionE
  case 2:
    // invoke: _ZNK7QObject10disconnectEPKS_PKc
  case 3:
    // invoke: _ZNK7QObject10disconnectEPKcPKS_S1_
  case 4:
    // invoke: _ZN7QObject10disconnectEPKS_PKcS1_S3_
  default:
    qtrt.ErrorResolve("QObject", "disconnect", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "children", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "isWindowType", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "deleteLater", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "objectName", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "setProperty", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "signalsBlocked", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "userData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "parent", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "installEventFilter", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "blockSignals", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "setObjectName", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "dumpObjectInfo", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QObject", "killTimer", args)
 }

}

// <= body block end

