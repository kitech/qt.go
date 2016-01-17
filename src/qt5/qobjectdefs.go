package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMetaObject::Connection::~Connection();
extern void _ZN11QMetaObject10ConnectionD2Ev(void* qthis); // 4
  // proto:  void QMetaObject::Connection::Connection();
extern void _ZN11QMetaObject10ConnectionC2Ev(void* qthis); // 3
  // proto:  void QGenericReturnArgument::QGenericReturnArgument(const char * aName, void * aData);
extern void _ZN22QGenericReturnArgumentC2EPKcPv(void* qthis, unsigned char* arg0, void* arg1); // 1
  // proto: static void QMetaObject::activate(QObject * sender, const QMetaObject * , int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectPKS_iPPv(void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto: static void QMetaObject::activate(QObject * sender, int signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiPPv(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static void QMetaObject::activate(QObject * sender, int signal_offset, int local_signal_index, void ** argv);
extern void _ZN11QMetaObject8activateEP7QObjectiiPPv(void* arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  int QMetaObject::indexOfMethod(const char * method);
extern void _ZNK11QMetaObject13indexOfMethodEPKc(void* qthis, unsigned char* arg0); // 4
  // proto: static bool QMetaObject::disconnectOne(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QMetaObject::indexOfProperty(const char * name);
extern void _ZNK11QMetaObject15indexOfPropertyEPKc(void* qthis, unsigned char* arg0); // 4
  // proto: static QMetaObject::Connection QMetaObject::connect(const QObject * sender, int signal_index, const QObject * receiver, int method_index, int type, int * types);
extern void _ZN11QMetaObject7connectEPK7QObjectiS2_iiPi(void* arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t* arg5); // 4
  // proto:  int QMetaObject::propertyOffset();
extern void _ZNK11QMetaObject14propertyOffsetEv(void* qthis); // 4
  // proto: static void QMetaObject::connectSlotsByName(QObject * o);
extern void _ZN11QMetaObject18connectSlotsByNameEP7QObject(void* arg0); // 4
  // proto:  int QMetaObject::indexOfConstructor(const char * constructor);
extern void _ZNK11QMetaObject18indexOfConstructorEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  int QMetaObject::methodCount();
extern void _ZNK11QMetaObject11methodCountEv(void* qthis); // 4
  // proto: static bool QMetaObject::disconnect(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern void _ZN11QMetaObject10disconnectEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QMetaObject::methodOffset();
extern void _ZNK11QMetaObject12methodOffsetEv(void* qthis); // 4
  // proto: static QByteArray QMetaObject::normalizedType(const char * type);
extern void _ZN11QMetaObject14normalizedTypeEPKc(unsigned char* arg0); // 4
  // proto:  int QMetaObject::indexOfEnumerator(const char * name);
extern void _ZNK11QMetaObject17indexOfEnumeratorEPKc(void* qthis, unsigned char* arg0); // 4
  // proto: static bool QMetaObject::checkConnectArgs(const char * signal, const char * method);
extern void _ZN11QMetaObject16checkConnectArgsEPKcS1_(unsigned char* arg0, unsigned char* arg1); // 4
  // proto: static bool QMetaObject::checkConnectArgs(const QMetaMethod & signal, const QMetaMethod & method);
extern void _ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(void* arg0, void* arg1); // 4
  // proto:  int QMetaObject::indexOfSlot(const char * slot);
extern void _ZNK11QMetaObject11indexOfSlotEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  QMetaMethod QMetaObject::method(int index);
extern void _ZNK11QMetaObject6methodEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::indexOfClassInfo(const char * name);
extern void _ZNK11QMetaObject16indexOfClassInfoEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  QMetaProperty QMetaObject::userProperty();
extern void _ZNK11QMetaObject12userPropertyEv(void* qthis); // 4
  // proto:  QMetaEnum QMetaObject::enumerator(int index);
extern void _ZNK11QMetaObject10enumeratorEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::classInfoOffset();
extern void _ZNK11QMetaObject15classInfoOffsetEv(void* qthis); // 4
  // proto:  int QMetaObject::indexOfSignal(const char * signal);
extern void _ZNK11QMetaObject13indexOfSignalEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  int QMetaObject::enumeratorCount();
extern void _ZNK11QMetaObject15enumeratorCountEv(void* qthis); // 4
  // proto:  int QMetaObject::enumeratorOffset();
extern void _ZNK11QMetaObject16enumeratorOffsetEv(void* qthis); // 4
  // proto:  int QMetaObject::classInfoCount();
extern void _ZNK11QMetaObject14classInfoCountEv(void* qthis); // 4
  // proto:  int QMetaObject::constructorCount();
extern void _ZNK11QMetaObject16constructorCountEv(void* qthis); // 4
  // proto:  QMetaClassInfo QMetaObject::classInfo(int index);
extern void _ZNK11QMetaObject9classInfoEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaObject::className();
extern void _ZNK11QMetaObject9classNameEv(void* qthis); // 4
  // proto:  const QObject * QMetaObject::cast(const QObject * obj);
extern void _ZNK11QMetaObject4castEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QMetaObject::cast(QObject * obj);
extern void _ZNK11QMetaObject4castEP7QObject(void* qthis, void* arg0); // 4
  // proto: static QByteArray QMetaObject::normalizedSignature(const char * method);
extern void _ZN11QMetaObject19normalizedSignatureEPKc(unsigned char* arg0); // 4
  // proto:  const QMetaObject * QMetaObject::superClass();
extern void _ZNK11QMetaObject10superClassEv(void* qthis); // 2
  // proto:  QMetaMethod QMetaObject::constructor(int index);
extern void _ZNK11QMetaObject11constructorEi(void* qthis, int32_t arg0); // 4
  // proto:  QMetaProperty QMetaObject::property(int index);
extern void _ZNK11QMetaObject8propertyEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::propertyCount();
extern void _ZNK11QMetaObject13propertyCountEv(void* qthis); // 4
  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
extern void _ZN16QGenericArgumentC2EPKcPKv(void* qthis, unsigned char* arg0, void* arg1); // 1
  // proto:  void * QGenericArgument::data();
extern void _ZNK16QGenericArgument4dataEv(void* qthis); // 2
  // proto:  const char * QGenericArgument::name();
extern void _ZNK16QGenericArgument4nameEv(void* qthis); // 2
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

// class sizeof(QMetaObject::Connection)=8
type QMetaObject__Connection struct {
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

// ~Connection()
func (this *QMetaObject__Connection) FreeQMetaObject__Connection(args ...interface{}) () {
  // ~Connection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10ConnectionD0Ev
    // invoke: void ~Connection()
    C._ZN11QMetaObject10ConnectionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaObject::Connection", "~Connection", args)
  }

}

// Connection()
func NewQMetaObject__Connection(args ...interface{}) QMetaObject__Connection {
  // Connection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10ConnectionC1Ev
    // invoke: void Connection()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QMetaObject10ConnectionC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMetaObject::Connection", "Connection", args)
  }

  return QMetaObject__Connection{}
}

// QGenericReturnArgument(const char *, void *)
func NewQGenericReturnArgument(args ...interface{}) QGenericReturnArgument {
  // QGenericReturnArgument(const char *, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGenericReturnArgumentC1EPKcPv
    // invoke: void QGenericReturnArgument(const char *, void *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN22QGenericReturnArgumentC2EPKcPv(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGenericReturnArgument", "QGenericReturnArgument", args)
  }

  return QGenericReturnArgument{}
}

// activate(class QObject *, const struct QMetaObject *, int, void **)
func (this *QMetaObject) activate_s(args ...interface{}) () {
  // activate(class QObject *, const struct QMetaObject *, int, void **)
  // activate(class QObject *, int, void **)
  // activate(class QObject *, int, int, void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QMetaObject{}) // "const QMetaObject *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.VoidpTy() // "void **"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.VoidpTy() // "void **"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject8activateEP7QObjectPKS_iPPv
    // invoke: void activate(class QObject *, const struct QMetaObject *, int, void **)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMetaObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    C._ZN11QMetaObject8activateEP7QObjectPKS_iPPv(arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN11QMetaObject8activateEP7QObjectiPPv
    // invoke: void activate(class QObject *, int, void **)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    C._ZN11QMetaObject8activateEP7QObjectiPPv(arg0, arg1, arg2)
  case 2:
    // invoke: _ZN11QMetaObject8activateEP7QObjectiiPPv
    // invoke: void activate(class QObject *, int, int, void **)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    C._ZN11QMetaObject8activateEP7QObjectiiPPv(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMetaObject", "activate", args)
  }

}

// indexOfMethod(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject13indexOfMethodEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfMethod", args)
  }

}

// disconnectOne(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) disconnectOne_s(args ...interface{}) () {
  // disconnectOne(const class QObject *, int, const class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i
    // invoke: bool disconnectOne(const class QObject *, int, const class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnectOne", args)
  }

}

// indexOfProperty(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject15indexOfPropertyEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfProperty", args)
  }

}

// connect(const class QObject *, int, const class QObject *, int, int, int *)
func (this *QMetaObject) connect_s(args ...interface{}) () {
  // connect(const class QObject *, int, const class QObject *, int, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject7connectEPK7QObjectiS2_iiPi
    // invoke: QMetaObject::Connection connect(const class QObject *, int, const class QObject *, int, int, int *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = (*C.int32_t)(args[5].(*int32))
    if false {fmt.Println(arg5)}
    C._ZN11QMetaObject7connectEPK7QObjectiS2_iiPi(arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QMetaObject", "connect", args)
  }

}

// propertyOffset()
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

// connectSlotsByName(class QObject *)
func (this *QMetaObject) connectSlotsByName_s(args ...interface{}) () {
  // connectSlotsByName(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject18connectSlotsByNameEP7QObject
    // invoke: void connectSlotsByName(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QMetaObject18connectSlotsByNameEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "connectSlotsByName", args)
  }

}

// indexOfConstructor(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject18indexOfConstructorEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfConstructor", args)
  }

}

// methodCount()
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

// disconnect(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) disconnect_s(args ...interface{}) () {
  // disconnect(const class QObject *, int, const class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10disconnectEPK7QObjectiS2_i
    // invoke: bool disconnect(const class QObject *, int, const class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN11QMetaObject10disconnectEPK7QObjectiS2_i(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnect", args)
  }

}

// methodOffset()
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

// normalizedType(const char *)
func (this *QMetaObject) normalizedType_s(args ...interface{}) () {
  // normalizedType(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject14normalizedTypeEPKc
    // invoke: QByteArray normalizedType(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN11QMetaObject14normalizedTypeEPKc(arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedType", args)
  }

}

// indexOfEnumerator(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject17indexOfEnumeratorEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfEnumerator", args)
  }

}

// checkConnectArgs(const char *, const char *)
func (this *QMetaObject) checkConnectArgs_s(args ...interface{}) () {
  // checkConnectArgs(const char *, const char *)
  // checkConnectArgs(const class QMetaMethod &, const class QMetaMethod &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"
  vtys[1][1] = reflect.TypeOf(QMetaMethod{}) // "const QMetaMethod &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject16checkConnectArgsEPKcS1_
    // invoke: bool checkConnectArgs(const char *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN11QMetaObject16checkConnectArgsEPKcS1_(arg0, arg1)
  case 1:
    // invoke: _ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_
    // invoke: bool checkConnectArgs(const class QMetaMethod &, const class QMetaMethod &)
    var arg0 = args[0].(QMetaMethod).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMetaMethod).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaObject", "checkConnectArgs", args)
  }

}

// indexOfSlot(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject11indexOfSlotEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSlot", args)
  }

}

// method(int)
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

// indexOfClassInfo(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject16indexOfClassInfoEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfClassInfo", args)
  }

}

// userProperty()
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

// enumerator(int)
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

// classInfoOffset()
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

// indexOfSignal(const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK11QMetaObject13indexOfSignalEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSignal", args)
  }

}

// enumeratorCount()
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

// enumeratorOffset()
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

// classInfoCount()
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

// constructorCount()
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

// classInfo(int)
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

// className()
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

// cast(const class QObject *)
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

// normalizedSignature(const char *)
func (this *QMetaObject) normalizedSignature_s(args ...interface{}) () {
  // normalizedSignature(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject19normalizedSignatureEPKc
    // invoke: QByteArray normalizedSignature(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN11QMetaObject19normalizedSignatureEPKc(arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedSignature", args)
  }

}

// superClass()
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
    C._ZNK11QMetaObject10superClassEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaObject", "superClass", args)
  }

}

// constructor(int)
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

// property(int)
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

// propertyCount()
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

// QGenericArgument(const char *, const void *)
func NewQGenericArgument(args ...interface{}) QGenericArgument {
  // QGenericArgument(const char *, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QGenericArgumentC1EPKcPKv
    // invoke: void QGenericArgument(const char *, const void *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN16QGenericArgumentC2EPKcPKv(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGenericArgument", "QGenericArgument", args)
  }

  return QGenericArgument{}
}

// data()
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
    C._ZNK16QGenericArgument4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGenericArgument", "data", args)
  }

}

// name()
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
    C._ZNK16QGenericArgument4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGenericArgument", "name", args)
  }

}

// <= body block end

