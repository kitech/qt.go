package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qobjectdefs.h
// dst-file: /src/core/qobjectdefs.go
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
  // proto:  void Connection::Connection();
extern void* dector_ZN11QMetaObject10ConnectionC1Ev();
extern void _ZN11QMetaObject10ConnectionC1Ev(void* qthis);
  // proto:  void Connection::~Connection();
extern void _ZN11QMetaObject10ConnectionD0Ev(void* qthis);
  // proto:  void Connection::Connection(void * data);
extern void* dector_ZN11QMetaObject10ConnectionC1EPv(void* arg0);
extern void _ZN11QMetaObject10ConnectionC1EPv(void* qthis, void* arg0);
  // proto:  void QGenericReturnArgument::QGenericReturnArgument(const char * aName, void * aData);
extern void* dector_ZN22QGenericReturnArgumentC1EPKcPv(char* arg0, void* arg1);
extern void demth_ZN22QGenericReturnArgumentC1EPKcPv(void* qthis, char* arg0, void* arg1);
  // proto: static QByteArray QMetaObject::normalizedSignature(const char * method);
extern void _ZN11QMetaObject19normalizedSignatureEPKc(char* arg0);
  // proto: static bool QMetaObject::disconnectOne(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(void* arg0, int arg1, void* arg2, int arg3);
  // proto:  int QMetaObject::indexOfSlot(const char * slot);
extern void _ZNK11QMetaObject11indexOfSlotEPKc(void* qthis, char* arg0);
  // proto:  int QMetaObject::indexOfConstructor(const char * constructor);
extern void _ZNK11QMetaObject18indexOfConstructorEPKc(void* qthis, char* arg0);
  // proto:  QMetaEnum QMetaObject::enumerator(int index);
extern void _ZNK11QMetaObject10enumeratorEi(void* qthis, int arg0);
  // proto:  int QMetaObject::indexOfMethod(const char * method);
extern void _ZNK11QMetaObject13indexOfMethodEPKc(void* qthis, char* arg0);
  // proto:  QMetaMethod QMetaObject::constructor(int index);
extern void _ZNK11QMetaObject11constructorEi(void* qthis, int arg0);
  // proto: static bool QMetaObject::checkConnectArgs(const char * signal, const char * method);
extern void _ZN11QMetaObject16checkConnectArgsEPKcS1_(char* arg0, char* arg1);
  // proto:  int QMetaObject::enumeratorOffset();
extern void _ZNK11QMetaObject16enumeratorOffsetEv(void* qthis);
  // proto:  QMetaProperty QMetaObject::property(int index);
extern void _ZNK11QMetaObject8propertyEi(void* qthis, int arg0);
  // proto: static void QMetaObject::connectSlotsByName(QObject * o);
extern void _ZN11QMetaObject18connectSlotsByNameEP7QObject(void* arg0);
  // proto:  QMetaProperty QMetaObject::userProperty();
extern void _ZNK11QMetaObject12userPropertyEv(void* qthis);
  // proto:  int QMetaObject::indexOfProperty(const char * name);
extern void _ZNK11QMetaObject15indexOfPropertyEPKc(void* qthis, char* arg0);
  // proto:  int QMetaObject::indexOfClassInfo(const char * name);
extern void _ZNK11QMetaObject16indexOfClassInfoEPKc(void* qthis, char* arg0);
  // proto: static void QMetaObject::activate(QObject * sender, const QMetaObject * , int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectPKS_iPPv(void* arg0, void* arg1, int arg2, void* arg3);
  // proto:  const QObject * QMetaObject::cast(const QObject * obj);
extern void _ZNK11QMetaObject4castEPK7QObject(void* qthis, void* arg0);
  // proto:  QMetaMethod QMetaObject::method(int index);
extern void _ZNK11QMetaObject6methodEi(void* qthis, int arg0);
  // proto:  const QMetaObject * QMetaObject::superClass();
extern void _ZNK11QMetaObject10superClassEv(void* qthis);
  // proto:  QObject * QMetaObject::cast(QObject * obj);
extern void _ZNK11QMetaObject4castEP7QObject(void* qthis, void* arg0);
  // proto: static void QMetaObject::activate(QObject * sender, int signal_offset, int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiiPPv(void* arg0, int arg1, int arg2, void* arg3);
  // proto:  int QMetaObject::propertyCount();
extern void _ZNK11QMetaObject13propertyCountEv(void* qthis);
  // proto:  QMetaClassInfo QMetaObject::classInfo(int index);
extern void _ZNK11QMetaObject9classInfoEi(void* qthis, int arg0);
  // proto: static bool QMetaObject::checkConnectArgs(const QMetaMethod & signal, const QMetaMethod & method);
extern void _ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(void* arg0, void* arg1);
  // proto:  const char * QMetaObject::className();
extern void _ZNK11QMetaObject9classNameEv(void* qthis);
  // proto:  int QMetaObject::indexOfSignal(const char * signal);
extern void _ZNK11QMetaObject13indexOfSignalEPKc(void* qthis, char* arg0);
  // proto: static QByteArray QMetaObject::normalizedType(const char * type);
extern void _ZN11QMetaObject14normalizedTypeEPKc(char* arg0);
  // proto:  int QMetaObject::constructorCount();
extern void _ZNK11QMetaObject16constructorCountEv(void* qthis);
  // proto:  int QMetaObject::propertyOffset();
extern void _ZNK11QMetaObject14propertyOffsetEv(void* qthis);
  // proto: static bool QMetaObject::disconnect(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject10disconnectEPK7QObjectiS2_i(void* arg0, int arg1, void* arg2, int arg3);
  // proto: static void QMetaObject::activate(QObject * sender, int signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiPPv(void* arg0, int arg1, void* arg2);
  // proto:  int QMetaObject::enumeratorCount();
extern void _ZNK11QMetaObject15enumeratorCountEv(void* qthis);
  // proto:  int QMetaObject::classInfoOffset();
extern void _ZNK11QMetaObject15classInfoOffsetEv(void* qthis);
  // proto:  int QMetaObject::methodOffset();
extern void _ZNK11QMetaObject12methodOffsetEv(void* qthis);
  // proto:  int QMetaObject::indexOfEnumerator(const char * name);
extern void _ZNK11QMetaObject17indexOfEnumeratorEPKc(void* qthis, char* arg0);
  // proto:  int QMetaObject::methodCount();
extern void _ZNK11QMetaObject11methodCountEv(void* qthis);
  // proto:  int QMetaObject::classInfoCount();
extern void _ZNK11QMetaObject14classInfoCountEv(void* qthis);
  // proto:  const char * QGenericArgument::name();
extern void demth_ZNK16QGenericArgument4nameEv(void* qthis);
  // proto:  void * QGenericArgument::data();
extern void demth_ZNK16QGenericArgument4dataEv(void* qthis);
  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
extern void* dector_ZN16QGenericArgumentC1EPKcPKv(char* arg0, void* arg1);
extern void demth_ZN16QGenericArgumentC1EPKcPKv(void* qthis, char* arg0, void* arg1);
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

// class sizeof(Connection)=8
type Connection struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGenericReturnArgument)=16
type QGenericReturnArgument struct {
  /*qbase*/ QGenericArgument;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMetaObject)=48
type QMetaObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGenericArgument)=16
type QGenericArgument struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void Connection::Connection();
func NewConnection(args ...interface{}) Connection {
  return Connection{}
}

  // proto:  void Connection::~Connection();
func (this *Connection) FreeConnection(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("Connection", "~Connection", args)
  }

}

  // proto:  void QGenericReturnArgument::QGenericReturnArgument(const char * aName, void * aData);
func NewQGenericReturnArgument(args ...interface{}) QGenericReturnArgument {
  return QGenericReturnArgument{}
}

  // proto: static QByteArray QMetaObject::normalizedSignature(const char * method);
func (this *QMetaObject) normalizedSignature_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedSignature", args)
  }

}

  // proto: static bool QMetaObject::disconnectOne(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
func (this *QMetaObject) disconnectOne_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnectOne", args)
  }

}

  // proto:  int QMetaObject::indexOfSlot(const char * slot);
func (this *QMetaObject) indexOfSlot(args ...interface{}) () {
  // indexOfSlot(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11indexOfSlotEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSlot", args)
  }

}

  // proto:  int QMetaObject::indexOfConstructor(const char * constructor);
func (this *QMetaObject) indexOfConstructor(args ...interface{}) () {
  // indexOfConstructor(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject18indexOfConstructorEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfConstructor", args)
  }

}

  // proto:  QMetaEnum QMetaObject::enumerator(int index);
func (this *QMetaObject) enumerator(args ...interface{}) () {
  // enumerator(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject10enumeratorEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "enumerator", args)
  }

}

  // proto:  int QMetaObject::indexOfMethod(const char * method);
func (this *QMetaObject) indexOfMethod(args ...interface{}) () {
  // indexOfMethod(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13indexOfMethodEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfMethod", args)
  }

}

  // proto:  QMetaMethod QMetaObject::constructor(int index);
func (this *QMetaObject) constructor(args ...interface{}) () {
  // constructor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11constructorEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "constructor", args)
  }

}

  // proto: static bool QMetaObject::checkConnectArgs(const char * signal, const char * method);
func (this *QMetaObject) checkConnectArgs_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "checkConnectArgs", args)
  }

}

  // proto:  int QMetaObject::enumeratorOffset();
func (this *QMetaObject) enumeratorOffset(args ...interface{}) () {
  // enumeratorOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16enumeratorOffsetEv
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorOffset", args)
  }

}

  // proto:  QMetaProperty QMetaObject::property(int index);
func (this *QMetaObject) property(args ...interface{}) () {
  // property(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject8propertyEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "property", args)
  }

}

  // proto: static void QMetaObject::connectSlotsByName(QObject * o);
func (this *QMetaObject) connectSlotsByName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "connectSlotsByName", args)
  }

}

  // proto:  QMetaProperty QMetaObject::userProperty();
func (this *QMetaObject) userProperty(args ...interface{}) () {
  // userProperty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject12userPropertyEv
  default:
    qtrt.ErrorResolve("QMetaObject", "userProperty", args)
  }

}

  // proto:  int QMetaObject::indexOfProperty(const char * name);
func (this *QMetaObject) indexOfProperty(args ...interface{}) () {
  // indexOfProperty(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15indexOfPropertyEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfProperty", args)
  }

}

  // proto:  int QMetaObject::indexOfClassInfo(const char * name);
func (this *QMetaObject) indexOfClassInfo(args ...interface{}) () {
  // indexOfClassInfo(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16indexOfClassInfoEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfClassInfo", args)
  }

}

  // proto: static void QMetaObject::activate(QObject * sender, const QMetaObject * , int local_signal_index, void ** argv);
func (this *QMetaObject) activate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "activate", args)
  }

}

  // proto:  const QObject * QMetaObject::cast(const QObject * obj);
func (this *QMetaObject) cast(args ...interface{}) () {
  // cast(const class QObject *)
  // cast(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject4castEPK7QObject
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK11QMetaObject4castEP7QObject
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "cast", args)
  }

}

  // proto:  QMetaMethod QMetaObject::method(int index);
func (this *QMetaObject) method(args ...interface{}) () {
  // method(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject6methodEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "method", args)
  }

}

  // proto:  const QMetaObject * QMetaObject::superClass();
func (this *QMetaObject) superClass(args ...interface{}) () {
  // superClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject10superClassEv
  default:
    qtrt.ErrorResolve("QMetaObject", "superClass", args)
  }

}

  // proto:  int QMetaObject::propertyCount();
func (this *QMetaObject) propertyCount(args ...interface{}) () {
  // propertyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13propertyCountEv
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyCount", args)
  }

}

  // proto:  QMetaClassInfo QMetaObject::classInfo(int index);
func (this *QMetaObject) classInfo(args ...interface{}) () {
  // classInfo(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject9classInfoEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfo", args)
  }

}

  // proto:  const char * QMetaObject::className();
func (this *QMetaObject) className(args ...interface{}) () {
  // className()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject9classNameEv
  default:
    qtrt.ErrorResolve("QMetaObject", "className", args)
  }

}

  // proto:  int QMetaObject::indexOfSignal(const char * signal);
func (this *QMetaObject) indexOfSignal(args ...interface{}) () {
  // indexOfSignal(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13indexOfSignalEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSignal", args)
  }

}

  // proto: static QByteArray QMetaObject::normalizedType(const char * type);
func (this *QMetaObject) normalizedType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedType", args)
  }

}

  // proto:  int QMetaObject::constructorCount();
func (this *QMetaObject) constructorCount(args ...interface{}) () {
  // constructorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16constructorCountEv
  default:
    qtrt.ErrorResolve("QMetaObject", "constructorCount", args)
  }

}

  // proto:  int QMetaObject::propertyOffset();
func (this *QMetaObject) propertyOffset(args ...interface{}) () {
  // propertyOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject14propertyOffsetEv
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyOffset", args)
  }

}

  // proto: static bool QMetaObject::disconnect(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
func (this *QMetaObject) disconnect_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnect", args)
  }

}

  // proto:  int QMetaObject::enumeratorCount();
func (this *QMetaObject) enumeratorCount(args ...interface{}) () {
  // enumeratorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15enumeratorCountEv
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorCount", args)
  }

}

  // proto:  int QMetaObject::classInfoOffset();
func (this *QMetaObject) classInfoOffset(args ...interface{}) () {
  // classInfoOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15classInfoOffsetEv
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoOffset", args)
  }

}

  // proto:  int QMetaObject::methodOffset();
func (this *QMetaObject) methodOffset(args ...interface{}) () {
  // methodOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject12methodOffsetEv
  default:
    qtrt.ErrorResolve("QMetaObject", "methodOffset", args)
  }

}

  // proto:  int QMetaObject::indexOfEnumerator(const char * name);
func (this *QMetaObject) indexOfEnumerator(args ...interface{}) () {
  // indexOfEnumerator(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject17indexOfEnumeratorEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfEnumerator", args)
  }

}

  // proto:  int QMetaObject::methodCount();
func (this *QMetaObject) methodCount(args ...interface{}) () {
  // methodCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11methodCountEv
  default:
    qtrt.ErrorResolve("QMetaObject", "methodCount", args)
  }

}

  // proto:  int QMetaObject::classInfoCount();
func (this *QMetaObject) classInfoCount(args ...interface{}) () {
  // classInfoCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject14classInfoCountEv
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoCount", args)
  }

}

  // proto:  const char * QGenericArgument::name();
func (this *QGenericArgument) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QGenericArgument4nameEv
  default:
    qtrt.ErrorResolve("QGenericArgument", "name", args)
  }

}

  // proto:  void * QGenericArgument::data();
func (this *QGenericArgument) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QGenericArgument4dataEv
  default:
    qtrt.ErrorResolve("QGenericArgument", "data", args)
  }

}

  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
func NewQGenericArgument(args ...interface{}) QGenericArgument {
  return QGenericArgument{}
}

// <= body block end

