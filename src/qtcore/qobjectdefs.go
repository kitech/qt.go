package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void C_ZN11QMetaObject10ConnectionD2Ev(void* qthis); // 4
  // proto:  void QMetaObject::Connection::Connection();
extern void* C_ZN11QMetaObject10ConnectionC2Ev(); // 3
  // proto:  void QGenericReturnArgument::QGenericReturnArgument(const char * aName, void * aData);
extern void* C_ZN22QGenericReturnArgumentC2EPKcPv(void* arg0, void* arg1); // 1
  // proto:  bool QMetaObject::inherits(const QMetaObject * metaObject);
extern bool C_ZNK11QMetaObject8inheritsEPKS_(void* qthis, void* arg0); // 4
  // proto: static void QMetaObject::activate(QObject * sender, const QMetaObject * , int local_signal_index, void ** argv);
extern void C_ZN11QMetaObject8activateEP7QObjectPKS_iPPv(void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto: static void QMetaObject::activate(QObject * sender, int signal_index, void ** argv);
extern void C_ZN11QMetaObject8activateEP7QObjectiPPv(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static void QMetaObject::activate(QObject * sender, int signal_offset, int local_signal_index, void ** argv);
extern void C_ZN11QMetaObject8activateEP7QObjectiiPPv(void* arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  int QMetaObject::indexOfMethod(const char * method);
extern int32_t C_ZNK11QMetaObject13indexOfMethodEPKc(void* qthis, void* arg0); // 4
  // proto: static bool QMetaObject::disconnectOne(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern bool C_ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QMetaObject::indexOfProperty(const char * name);
extern int32_t C_ZNK11QMetaObject15indexOfPropertyEPKc(void* qthis, void* arg0); // 4
  // proto: static QMetaObject::Connection QMetaObject::connect(const QObject * sender, int signal_index, const QObject * receiver, int method_index, int type, int * types);
extern void C_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi(void* arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4, void* arg5); // 4
  // proto:  int QMetaObject::propertyOffset();
extern int32_t C_ZNK11QMetaObject14propertyOffsetEv(void* qthis); // 4
  // proto: static void QMetaObject::connectSlotsByName(QObject * o);
extern void C_ZN11QMetaObject18connectSlotsByNameEP7QObject(void* arg0); // 4
  // proto:  int QMetaObject::indexOfConstructor(const char * constructor);
extern int32_t C_ZNK11QMetaObject18indexOfConstructorEPKc(void* qthis, void* arg0); // 4
  // proto:  int QMetaObject::methodCount();
extern int32_t C_ZNK11QMetaObject11methodCountEv(void* qthis); // 4
  // proto: static bool QMetaObject::disconnect(const QObject * sender, int signal_index, const QObject * receiver, int method_index);
extern bool C_ZN11QMetaObject10disconnectEPK7QObjectiS2_i(void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  int QMetaObject::methodOffset();
extern int32_t C_ZNK11QMetaObject12methodOffsetEv(void* qthis); // 4
  // proto: static QByteArray QMetaObject::normalizedType(const char * type);
extern void* C_ZN11QMetaObject14normalizedTypeEPKc(void* arg0); // 4
  // proto:  int QMetaObject::indexOfEnumerator(const char * name);
extern int32_t C_ZNK11QMetaObject17indexOfEnumeratorEPKc(void* qthis, void* arg0); // 4
  // proto: static bool QMetaObject::checkConnectArgs(const char * signal, const char * method);
extern bool C_ZN11QMetaObject16checkConnectArgsEPKcS1_(void* arg0, void* arg1); // 4
  // proto: static bool QMetaObject::checkConnectArgs(const QMetaMethod & signal, const QMetaMethod & method);
extern bool C_ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(void* arg0, void* arg1); // 4
  // proto:  int QMetaObject::indexOfSlot(const char * slot);
extern int32_t C_ZNK11QMetaObject11indexOfSlotEPKc(void* qthis, void* arg0); // 4
  // proto:  QMetaMethod QMetaObject::method(int index);
extern void* C_ZNK11QMetaObject6methodEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::indexOfClassInfo(const char * name);
extern int32_t C_ZNK11QMetaObject16indexOfClassInfoEPKc(void* qthis, void* arg0); // 4
  // proto:  QMetaProperty QMetaObject::userProperty();
extern void* C_ZNK11QMetaObject12userPropertyEv(void* qthis); // 4
  // proto:  QMetaEnum QMetaObject::enumerator(int index);
extern void* C_ZNK11QMetaObject10enumeratorEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::classInfoOffset();
extern int32_t C_ZNK11QMetaObject15classInfoOffsetEv(void* qthis); // 4
  // proto:  int QMetaObject::indexOfSignal(const char * signal);
extern int32_t C_ZNK11QMetaObject13indexOfSignalEPKc(void* qthis, void* arg0); // 4
  // proto:  int QMetaObject::enumeratorCount();
extern int32_t C_ZNK11QMetaObject15enumeratorCountEv(void* qthis); // 4
  // proto:  int QMetaObject::enumeratorOffset();
extern int32_t C_ZNK11QMetaObject16enumeratorOffsetEv(void* qthis); // 4
  // proto:  int QMetaObject::classInfoCount();
extern int32_t C_ZNK11QMetaObject14classInfoCountEv(void* qthis); // 4
  // proto:  int QMetaObject::constructorCount();
extern int32_t C_ZNK11QMetaObject16constructorCountEv(void* qthis); // 4
  // proto:  QMetaClassInfo QMetaObject::classInfo(int index);
extern void* C_ZNK11QMetaObject9classInfoEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaObject::className();
extern void* C_ZNK11QMetaObject9classNameEv(void* qthis); // 4
  // proto:  const QObject * QMetaObject::cast(const QObject * obj);
extern void* C_ZNK11QMetaObject4castEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QMetaObject::cast(QObject * obj);
extern void* C_ZNK11QMetaObject4castEP7QObject(void* qthis, void* arg0); // 4
  // proto: static QByteArray QMetaObject::normalizedSignature(const char * method);
extern void* C_ZN11QMetaObject19normalizedSignatureEPKc(void* arg0); // 4
  // proto:  const QMetaObject * QMetaObject::superClass();
extern void C_ZNK11QMetaObject10superClassEv(void* qthis); // 2
  // proto:  QMetaMethod QMetaObject::constructor(int index);
extern void* C_ZNK11QMetaObject11constructorEi(void* qthis, int32_t arg0); // 4
  // proto:  QMetaProperty QMetaObject::property(int index);
extern void* C_ZNK11QMetaObject8propertyEi(void* qthis, int32_t arg0); // 4
  // proto:  int QMetaObject::propertyCount();
extern int32_t C_ZNK11QMetaObject13propertyCountEv(void* qthis); // 4
  // proto:  void QGenericArgument::QGenericArgument(const char * aName, const void * aData);
extern void* C_ZN16QGenericArgumentC2EPKcPKv(void* arg0, void* arg1); // 1
  // proto:  void * QGenericArgument::data();
extern void C_ZNK16QGenericArgument4dataEv(void* qthis); // 2
  // proto:  const char * QGenericArgument::name();
extern void* C_ZNK16QGenericArgument4nameEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGenericReturnArgument)=16
type QGenericReturnArgument struct {
  /*qbase*/ QGenericArgument;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMetaObject)=48
type QMetaObject struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGenericArgument)=16
type QGenericArgument struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~Connection()
func (this *QMetaObject__Connection) Freeqmetaobject__Connection(args ...interface{}) () {
  // ~Connection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10ConnectionD0Ev
    // invoke: void ~Connection()
    C.C_ZN11QMetaObject10ConnectionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMetaObject::Connection", "~Connection", args)
  }

  return
}

// Connection()
func NewQMetaObject__Connection(args ...interface{}) *QMetaObject__Connection {
  // Connection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10ConnectionC1Ev
    // invoke: void Connection()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QMetaObject10ConnectionC2Ev()
    return &QMetaObject__Connection{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaObject::Connection", "Connection", args)
  }

  return nil // QMetaObject__Connection{Qclsinst:qthis}
}

// QGenericReturnArgument(const char *, void *)
func NewQGenericReturnArgument(args ...interface{}) *QGenericReturnArgument {
  // QGenericReturnArgument(const char *, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGenericReturnArgumentC1EPKcPv
    // invoke: void QGenericReturnArgument(const char *, void *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QGenericReturnArgumentC2EPKcPv(arg0, arg1)
    return &QGenericReturnArgument{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGenericReturnArgument", "QGenericReturnArgument", args)
  }

  return nil // QGenericReturnArgument{Qclsinst:qthis}
}

// inherits(const struct QMetaObject *)
func (this *QMetaObject) Inherits(args ...interface{}) (ret interface{}) {
  // inherits(const struct QMetaObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMetaObject{}) // "const QMetaObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject8inheritsEPKS_
    // invoke: bool inherits(const struct QMetaObject *)
    var arg0 = args[0].(*QMetaObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject8inheritsEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "inherits", args)
  }

  return
}

// activate(class QObject *, const struct QMetaObject *, int, void **)
func (this *QMetaObject) Activate_S(args ...interface{}) () {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject8activateEP7QObjectPKS_iPPv
    // invoke: void activate(class QObject *, const struct QMetaObject *, int, void **)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMetaObject).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    C.C_ZN11QMetaObject8activateEP7QObjectPKS_iPPv(arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN11QMetaObject8activateEP7QObjectiPPv
    // invoke: void activate(class QObject *, int, void **)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    C.C_ZN11QMetaObject8activateEP7QObjectiPPv(arg0, arg1, arg2)
  case 2:
    // invoke: _ZN11QMetaObject8activateEP7QObjectiiPPv
    // invoke: void activate(class QObject *, int, int, void **)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    C.C_ZN11QMetaObject8activateEP7QObjectiiPPv(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMetaObject", "activate", args)
  }

  return
}

// indexOfMethod(const char *)
func (this *QMetaObject) Indexofmethod(args ...interface{}) (ret interface{}) {
  // indexOfMethod(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13indexOfMethodEPKc
    // invoke: int indexOfMethod(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject13indexOfMethodEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfMethod", args)
  }

  return
}

// disconnectOne(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) Disconnectone_S(args ...interface{}) (ret interface{}) {
  // disconnectOne(const class QObject *, int, const class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i
    // invoke: bool disconnectOne(const class QObject *, int, const class QObject *, int)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnectOne", args)
  }

  return
}

// indexOfProperty(const char *)
func (this *QMetaObject) Indexofproperty(args ...interface{}) (ret interface{}) {
  // indexOfProperty(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15indexOfPropertyEPKc
    // invoke: int indexOfProperty(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject15indexOfPropertyEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfProperty", args)
  }

  return
}

// connect(const class QObject *, int, const class QObject *, int, int, int *)
func (this *QMetaObject) Connect_S(args ...interface{}) () {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject7connectEPK7QObjectiS2_iiPi
    // invoke: QMetaObject::Connection connect(const class QObject *, int, const class QObject *, int, int, int *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = (unsafe.Pointer)(args[5].(*int32))
    if false {fmt.Println(arg5)}
    C.C_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi(arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QMetaObject", "connect", args)
  }

  return
}

// propertyOffset()
func (this *QMetaObject) Propertyoffset(args ...interface{}) (ret interface{}) {
  // propertyOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject14propertyOffsetEv
    // invoke: int propertyOffset()
    var ret0 = C.C_ZNK11QMetaObject14propertyOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyOffset", args)
  }

  return
}

// connectSlotsByName(class QObject *)
func (this *QMetaObject) Connectslotsbyname_S(args ...interface{}) () {
  // connectSlotsByName(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject18connectSlotsByNameEP7QObject
    // invoke: void connectSlotsByName(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QMetaObject18connectSlotsByNameEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QMetaObject", "connectSlotsByName", args)
  }

  return
}

// indexOfConstructor(const char *)
func (this *QMetaObject) Indexofconstructor(args ...interface{}) (ret interface{}) {
  // indexOfConstructor(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject18indexOfConstructorEPKc
    // invoke: int indexOfConstructor(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject18indexOfConstructorEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfConstructor", args)
  }

  return
}

// methodCount()
func (this *QMetaObject) Methodcount(args ...interface{}) (ret interface{}) {
  // methodCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11methodCountEv
    // invoke: int methodCount()
    var ret0 = C.C_ZNK11QMetaObject11methodCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "methodCount", args)
  }

  return
}

// disconnect(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) Disconnect_S(args ...interface{}) (ret interface{}) {
  // disconnect(const class QObject *, int, const class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject10disconnectEPK7QObjectiS2_i
    // invoke: bool disconnect(const class QObject *, int, const class QObject *, int)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN11QMetaObject10disconnectEPK7QObjectiS2_i(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "disconnect", args)
  }

  return
}

// methodOffset()
func (this *QMetaObject) Methodoffset(args ...interface{}) (ret interface{}) {
  // methodOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject12methodOffsetEv
    // invoke: int methodOffset()
    var ret0 = C.C_ZNK11QMetaObject12methodOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "methodOffset", args)
  }

  return
}

// normalizedType(const char *)
func (this *QMetaObject) Normalizedtype_S(args ...interface{}) (ret interface{}) {
  // normalizedType(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject14normalizedTypeEPKc
    // invoke: QByteArray normalizedType(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN11QMetaObject14normalizedTypeEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedType", args)
  }

  return
}

// indexOfEnumerator(const char *)
func (this *QMetaObject) Indexofenumerator(args ...interface{}) (ret interface{}) {
  // indexOfEnumerator(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject17indexOfEnumeratorEPKc
    // invoke: int indexOfEnumerator(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject17indexOfEnumeratorEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfEnumerator", args)
  }

  return
}

// checkConnectArgs(const char *, const char *)
func (this *QMetaObject) Checkconnectargs_S(args ...interface{}) (ret interface{}) {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject16checkConnectArgsEPKcS1_
    // invoke: bool checkConnectArgs(const char *, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN11QMetaObject16checkConnectArgsEPKcS1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_
    // invoke: bool checkConnectArgs(const class QMetaMethod &, const class QMetaMethod &)
    var arg0 = args[0].(*QMetaMethod).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMetaMethod).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "checkConnectArgs", args)
  }

  return
}

// indexOfSlot(const char *)
func (this *QMetaObject) Indexofslot(args ...interface{}) (ret interface{}) {
  // indexOfSlot(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11indexOfSlotEPKc
    // invoke: int indexOfSlot(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject11indexOfSlotEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSlot", args)
  }

  return
}

// method(int)
func (this *QMetaObject) Method(args ...interface{}) (ret interface{}) {
  // method(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject6methodEi
    // invoke: QMetaMethod method(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject6methodEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaMethod{}) // "QMetaMethod"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "method", args)
  }

  return
}

// indexOfClassInfo(const char *)
func (this *QMetaObject) Indexofclassinfo(args ...interface{}) (ret interface{}) {
  // indexOfClassInfo(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16indexOfClassInfoEPKc
    // invoke: int indexOfClassInfo(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject16indexOfClassInfoEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfClassInfo", args)
  }

  return
}

// userProperty()
func (this *QMetaObject) Userproperty(args ...interface{}) (ret interface{}) {
  // userProperty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject12userPropertyEv
    // invoke: QMetaProperty userProperty()
    var ret0 = C.C_ZNK11QMetaObject12userPropertyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaProperty{}) // "QMetaProperty"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "userProperty", args)
  }

  return
}

// enumerator(int)
func (this *QMetaObject) Enumerator(args ...interface{}) (ret interface{}) {
  // enumerator(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject10enumeratorEi
    // invoke: QMetaEnum enumerator(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject10enumeratorEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaEnum{}) // "QMetaEnum"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "enumerator", args)
  }

  return
}

// classInfoOffset()
func (this *QMetaObject) Classinfooffset(args ...interface{}) (ret interface{}) {
  // classInfoOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15classInfoOffsetEv
    // invoke: int classInfoOffset()
    var ret0 = C.C_ZNK11QMetaObject15classInfoOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoOffset", args)
  }

  return
}

// indexOfSignal(const char *)
func (this *QMetaObject) Indexofsignal(args ...interface{}) (ret interface{}) {
  // indexOfSignal(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13indexOfSignalEPKc
    // invoke: int indexOfSignal(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK11QMetaObject13indexOfSignalEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSignal", args)
  }

  return
}

// enumeratorCount()
func (this *QMetaObject) Enumeratorcount(args ...interface{}) (ret interface{}) {
  // enumeratorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject15enumeratorCountEv
    // invoke: int enumeratorCount()
    var ret0 = C.C_ZNK11QMetaObject15enumeratorCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorCount", args)
  }

  return
}

// enumeratorOffset()
func (this *QMetaObject) Enumeratoroffset(args ...interface{}) (ret interface{}) {
  // enumeratorOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16enumeratorOffsetEv
    // invoke: int enumeratorOffset()
    var ret0 = C.C_ZNK11QMetaObject16enumeratorOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorOffset", args)
  }

  return
}

// classInfoCount()
func (this *QMetaObject) Classinfocount(args ...interface{}) (ret interface{}) {
  // classInfoCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject14classInfoCountEv
    // invoke: int classInfoCount()
    var ret0 = C.C_ZNK11QMetaObject14classInfoCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoCount", args)
  }

  return
}

// constructorCount()
func (this *QMetaObject) Constructorcount(args ...interface{}) (ret interface{}) {
  // constructorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject16constructorCountEv
    // invoke: int constructorCount()
    var ret0 = C.C_ZNK11QMetaObject16constructorCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "constructorCount", args)
  }

  return
}

// classInfo(int)
func (this *QMetaObject) Classinfo(args ...interface{}) (ret interface{}) {
  // classInfo(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject9classInfoEi
    // invoke: QMetaClassInfo classInfo(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject9classInfoEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaClassInfo{}) // "QMetaClassInfo"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfo", args)
  }

  return
}

// className()
func (this *QMetaObject) Classname(args ...interface{}) (ret interface{}) {
  // className()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject9classNameEv
    // invoke: const char * className()
    var ret0 = C.C_ZNK11QMetaObject9classNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "className", args)
  }

  return
}

// cast(const class QObject *)
func (this *QMetaObject) Cast(args ...interface{}) (ret interface{}) {
  // cast(const class QObject *)
  // cast(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject4castEPK7QObject
    // invoke: const QObject * cast(const class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject4castEPK7QObject(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "const QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK11QMetaObject4castEP7QObject
    // invoke: QObject * cast(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject4castEP7QObject(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "cast", args)
  }

  return
}

// normalizedSignature(const char *)
func (this *QMetaObject) Normalizedsignature_S(args ...interface{}) (ret interface{}) {
  // normalizedSignature(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaObject19normalizedSignatureEPKc
    // invoke: QByteArray normalizedSignature(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN11QMetaObject19normalizedSignatureEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "normalizedSignature", args)
  }

  return
}

// superClass()
func (this *QMetaObject) Superclass(args ...interface{}) () {
  // superClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject10superClassEv
    // invoke: const QMetaObject * superClass()
    C.C_ZNK11QMetaObject10superClassEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMetaObject", "superClass", args)
  }

  return
}

// constructor(int)
func (this *QMetaObject) Constructor(args ...interface{}) (ret interface{}) {
  // constructor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject11constructorEi
    // invoke: QMetaMethod constructor(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject11constructorEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaMethod{}) // "QMetaMethod"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "constructor", args)
  }

  return
}

// property(int)
func (this *QMetaObject) Property(args ...interface{}) (ret interface{}) {
  // property(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject8propertyEi
    // invoke: QMetaProperty property(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaObject8propertyEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaProperty{}) // "QMetaProperty"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "property", args)
  }

  return
}

// propertyCount()
func (this *QMetaObject) Propertycount(args ...interface{}) (ret interface{}) {
  // propertyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaObject13propertyCountEv
    // invoke: int propertyCount()
    var ret0 = C.C_ZNK11QMetaObject13propertyCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyCount", args)
  }

  return
}

// QGenericArgument(const char *, const void *)
func NewQGenericArgument(args ...interface{}) *QGenericArgument {
  // QGenericArgument(const char *, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QGenericArgumentC1EPKcPKv
    // invoke: void QGenericArgument(const char *, const void *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QGenericArgumentC2EPKcPKv(arg0, arg1)
    return &QGenericArgument{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGenericArgument", "QGenericArgument", args)
  }

  return nil // QGenericArgument{Qclsinst:qthis}
}

// data()
func (this *QGenericArgument) Data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QGenericArgument4dataEv
    // invoke: void * data()
    C.C_ZNK16QGenericArgument4dataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGenericArgument", "data", args)
  }

  return
}

// name()
func (this *QGenericArgument) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QGenericArgument4nameEv
    // invoke: const char * name()
    var ret0 = C.C_ZNK16QGenericArgument4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGenericArgument", "name", args)
  }

  return
}

// <= body block end

