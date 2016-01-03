package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
extern void* dector_ZN22QGenericReturnArgumentC1EPKcPv(unsigned char* arg0, void* arg1);
extern void demth_ZN22QGenericReturnArgumentC1EPKcPv(void* qthis, unsigned char* arg0, void* arg1);
  // proto: static QByteArray QMetaObject::normalizedSignature(const char * method);
extern void _ZN11QMetaObject19normalizedSignatureEPKc(unsigned char* arg0);
  // proto: static bool QMetaObject::disconnectOne(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3);
  // proto:  int QMetaObject::indexOfSlot(const char * slot);
extern void _ZNK11QMetaObject11indexOfSlotEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QMetaObject::indexOfConstructor(const char * constructor);
extern void _ZNK11QMetaObject18indexOfConstructorEPKc(void* qthis, unsigned char* arg0);
  // proto:  QMetaEnum QMetaObject::enumerator(int index);
extern void _ZNK11QMetaObject10enumeratorEi(void* qthis, int32_t arg0);
  // proto:  int QMetaObject::indexOfMethod(const char * method);
extern void _ZNK11QMetaObject13indexOfMethodEPKc(void* qthis, unsigned char* arg0);
  // proto:  QMetaMethod QMetaObject::constructor(int index);
extern void _ZNK11QMetaObject11constructorEi(void* qthis, int32_t arg0);
  // proto: static bool QMetaObject::checkConnectArgs(const char * signal, const char * method);
extern void _ZN11QMetaObject16checkConnectArgsEPKcS1_(unsigned char* arg0, unsigned char* arg1);
  // proto:  int QMetaObject::enumeratorOffset();
extern void _ZNK11QMetaObject16enumeratorOffsetEv(void* qthis);
  // proto:  QMetaProperty QMetaObject::property(int index);
extern void _ZNK11QMetaObject8propertyEi(void* qthis, int32_t arg0);
  // proto: static void QMetaObject::connectSlotsByName(QObject * o);
extern void _ZN11QMetaObject18connectSlotsByNameEP7QObject(void* arg0);
  // proto:  QMetaProperty QMetaObject::userProperty();
extern void _ZNK11QMetaObject12userPropertyEv(void* qthis);
  // proto:  int QMetaObject::indexOfProperty(const char * name);
extern void _ZNK11QMetaObject15indexOfPropertyEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QMetaObject::indexOfClassInfo(const char * name);
extern void _ZNK11QMetaObject16indexOfClassInfoEPKc(void* qthis, unsigned char* arg0);
  // proto: static void QMetaObject::activate(QObject * sender, const QMetaObject * , int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectPKS_iPPv(void* arg0, void* arg1, int32_t arg2, void* arg3);
  // proto:  const QObject * QMetaObject::cast(const QObject * obj);
extern void _ZNK11QMetaObject4castEPK7QObject(void* qthis, void* arg0);
  // proto:  QMetaMethod QMetaObject::method(int index);
extern void _ZNK11QMetaObject6methodEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QMetaObject::superClass();
extern void demth_ZNK11QMetaObject10superClassEv(void* qthis);
  // proto:  QObject * QMetaObject::cast(QObject * obj);
extern void _ZNK11QMetaObject4castEP7QObject(void* qthis, void* arg0);
  // proto: static void QMetaObject::activate(QObject * sender, int signal_offset, int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiiPPv(void* arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  int QMetaObject::propertyCount();
extern void _ZNK11QMetaObject13propertyCountEv(void* qthis);
  // proto:  QMetaClassInfo QMetaObject::classInfo(int index);
extern void _ZNK11QMetaObject9classInfoEi(void* qthis, int32_t arg0);
  // proto: static bool QMetaObject::checkConnectArgs(const QMetaMethod & signal, const QMetaMethod & method);
extern void _ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(void* arg0, void* arg1);
  // proto:  const char * QMetaObject::className();
extern void _ZNK11QMetaObject9classNameEv(void* qthis);
  // proto:  int QMetaObject::indexOfSignal(const char * signal);
extern void _ZNK11QMetaObject13indexOfSignalEPKc(void* qthis, unsigned char* arg0);
  // proto: static QByteArray QMetaObject::normalizedType(const char * type);
extern void _ZN11QMetaObject14normalizedTypeEPKc(unsigned char* arg0);
  // proto:  int QMetaObject::constructorCount();
extern void _ZNK11QMetaObject16constructorCountEv(void* qthis);
  // proto:  int QMetaObject::propertyOffset();
extern void _ZNK11QMetaObject14propertyOffsetEv(void* qthis);
  // proto: static bool QMetaObject::disconnect(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject10disconnectEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3);
  // proto: static void QMetaObject::activate(QObject * sender, int signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiPPv(void* arg0, int32_t arg1, void* arg2);
  // proto:  int QMetaObject::enumeratorCount();
extern void _ZNK11QMetaObject15enumeratorCountEv(void* qthis);
  // proto:  int QMetaObject::classInfoOffset();
extern void _ZNK11QMetaObject15classInfoOffsetEv(void* qthis);
  // proto:  int QMetaObject::methodOffset();
extern void _ZNK11QMetaObject12methodOffsetEv(void* qthis);
  // proto:  int QMetaObject::indexOfEnumerator(const char * name);
extern void _ZNK11QMetaObject17indexOfEnumeratorEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QMetaObject::methodCount();
extern void _ZNK11QMetaObject11methodCountEv(void* qthis);
  // proto:  int QMetaObject::classInfoCount();
extern void _ZNK11QMetaObject14classInfoCountEv(void* qthis);
  // proto:  const char * QGenericArgument::name();
extern void demth_ZNK16QGenericArgument4nameEv(void* qthis);
  // proto:  void * QGenericArgument::data();
extern void demth_ZNK16QGenericArgument4dataEv(void* qthis);
  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
extern void* dector_ZN16QGenericArgumentC1EPKcPKv(unsigned char* arg0, void* arg1);
extern void demth_ZN16QGenericArgumentC1EPKcPKv(void* qthis, unsigned char* arg0, void* arg1);
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGenericReturnArgument)=16
type QGenericReturnArgument struct {
  /*qbase*/ QGenericArgument;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMetaObject)=48
type QMetaObject struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGenericArgument)=16
type QGenericArgument struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
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
    // invoke: int indexOfSlot(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject11indexOfSlotEPKc(this.qclsinst, arg0)
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
    // invoke: int indexOfConstructor(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject18indexOfConstructorEPKc(this.qclsinst, arg0)
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
    // invoke: QMetaEnum enumerator(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject10enumeratorEi(this.qclsinst, arg0)
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
    // invoke: int indexOfMethod(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject13indexOfMethodEPKc(this.qclsinst, arg0)
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
    // invoke: QMetaMethod constructor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject11constructorEi(this.qclsinst, arg0)
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
    // invoke: int enumeratorOffset()
    C._ZNK11QMetaObject16enumeratorOffsetEv(this.qclsinst)
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
    // invoke: QMetaProperty property(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject8propertyEi(this.qclsinst, arg0)
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
    // invoke: QMetaProperty userProperty()
    C._ZNK11QMetaObject12userPropertyEv(this.qclsinst)
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
    // invoke: int indexOfProperty(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject15indexOfPropertyEPKc(this.qclsinst, arg0)
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
    // invoke: int indexOfClassInfo(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject16indexOfClassInfoEPKc(this.qclsinst, arg0)
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
    // invoke: const QObject * cast(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject4castEPK7QObject(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK11QMetaObject4castEP7QObject
    // invoke: QObject * cast(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject4castEP7QObject(this.qclsinst, arg0)
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
    // invoke: QMetaMethod method(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject6methodEi(this.qclsinst, arg0)
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
    // invoke: const QMetaObject * superClass()
    C.demth_ZNK11QMetaObject10superClassEv(this.qclsinst)
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
    // invoke: int propertyCount()
    C._ZNK11QMetaObject13propertyCountEv(this.qclsinst)
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
    // invoke: QMetaClassInfo classInfo(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject9classInfoEi(this.qclsinst, arg0)
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
    // invoke: const char * className()
    C._ZNK11QMetaObject9classNameEv(this.qclsinst)
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
    // invoke: int indexOfSignal(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject13indexOfSignalEPKc(this.qclsinst, arg0)
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
    // invoke: int constructorCount()
    C._ZNK11QMetaObject16constructorCountEv(this.qclsinst)
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
    // invoke: int propertyOffset()
    C._ZNK11QMetaObject14propertyOffsetEv(this.qclsinst)
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
    // invoke: int enumeratorCount()
    C._ZNK11QMetaObject15enumeratorCountEv(this.qclsinst)
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
    // invoke: int classInfoOffset()
    C._ZNK11QMetaObject15classInfoOffsetEv(this.qclsinst)
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
    // invoke: int methodOffset()
    C._ZNK11QMetaObject12methodOffsetEv(this.qclsinst)
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
    // invoke: int indexOfEnumerator(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject17indexOfEnumeratorEPKc(this.qclsinst, arg0)
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
    // invoke: int methodCount()
    C._ZNK11QMetaObject11methodCountEv(this.qclsinst)
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
    // invoke: int classInfoCount()
    C._ZNK11QMetaObject14classInfoCountEv(this.qclsinst)
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
    // invoke: const char * name()
    C.demth_ZNK16QGenericArgument4nameEv(this.qclsinst)
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
    // invoke: void * data()
    C.demth_ZNK16QGenericArgument4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGenericArgument", "data", args)
  }

}

  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
func NewQGenericArgument(args ...interface{}) QGenericArgument {
  return QGenericArgument{}
}

// <= body block end

