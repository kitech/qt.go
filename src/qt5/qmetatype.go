package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qmetatype.h
// dst-file: /src/core/qmetatype.go
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
  // proto: static bool QMetaType::load(QDataStream & stream, int type, void * data);
extern bool C_ZN9QMetaType4loadER11QDataStreamiPv(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static int QMetaType::registerNormalizedTypedef(const ::QByteArray & normalizedTypeName, int aliasId);
extern int32_t C_ZN9QMetaType25registerNormalizedTypedefERK10QByteArrayi(void* arg0, int32_t arg1); // 4
  // proto: static bool QMetaType::compare(const void * lhs, const void * rhs, int typeId, int * result);
extern bool C_ZN9QMetaType7compareEPKvS1_iPi(void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto: static const char * QMetaType::typeName(int type);
extern void* C_ZN9QMetaType8typeNameEi(int32_t arg0); // 4
  // proto:  void QMetaType::QMetaType(const int type);
extern void* C_ZN9QMetaTypeC2Ei(int32_t arg0); // 3
  // proto: static bool QMetaType::isRegistered(int type);
extern bool C_ZN9QMetaType12isRegisteredEi(int32_t arg0); // 4
  // proto:  bool QMetaType::isRegistered();
extern bool C_ZNK9QMetaType12isRegisteredEv(void* qthis); // 2
  // proto: static bool QMetaType::unregisterType(int type);
extern bool C_ZN9QMetaType14unregisterTypeEi(int32_t arg0); // 4
  // proto: static TypeFlags QMetaType::typeFlags(int type);
extern void C_ZN9QMetaType9typeFlagsEi(int32_t arg0); // 4
  // proto: static void * QMetaType::create(int type, const void * copy);
extern void* C_ZN9QMetaType6createEiPKv(int32_t arg0, void* arg1); // 4
  // proto:  void * QMetaType::create(const void * copy);
extern void* C_ZNK9QMetaType6createEPKv(void* qthis, void* arg0); // 2
  // proto: static void * QMetaType::construct(int type, void * where, const void * copy);
extern void* C_ZN9QMetaType9constructEiPvPKv(int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void * QMetaType::construct(void * where, const void * copy);
extern void* C_ZNK9QMetaType9constructEPvPKv(void* qthis, void* arg0, void* arg1); // 2
  // proto: static bool QMetaType::hasRegisteredDebugStreamOperator(int typeId);
extern bool C_ZN9QMetaType32hasRegisteredDebugStreamOperatorEi(int32_t arg0); // 4
  // proto: static void QMetaType::destroy(int type, void * data);
extern void C_ZN9QMetaType7destroyEiPv(int32_t arg0, void* arg1); // 4
  // proto:  void QMetaType::destroy(void * data);
extern void C_ZNK9QMetaType7destroyEPv(void* qthis, void* arg0); // 2
  // proto: static bool QMetaType::save(QDataStream & stream, int type, const void * data);
extern bool C_ZN9QMetaType4saveER11QDataStreamiPKv(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static int QMetaType::type(const char * typeName);
extern int32_t C_ZN9QMetaType4typeEPKc(void* arg0); // 4
  // proto: static int QMetaType::type(const ::QByteArray & typeName);
extern int32_t C_ZN9QMetaType4typeERK10QByteArray(void* arg0); // 4
  // proto:  void QMetaType::~QMetaType();
extern void C_ZN9QMetaTypeD2Ev(void* qthis); // 2
  // proto: static bool QMetaType::hasRegisteredComparators(int typeId);
extern bool C_ZN9QMetaType24hasRegisteredComparatorsEi(int32_t arg0); // 4
  // proto:  bool QMetaType::isValid();
extern bool C_ZNK9QMetaType7isValidEv(void* qthis); // 2
  // proto: static bool QMetaType::equals(const void * lhs, const void * rhs, int typeId, int * result);
extern bool C_ZN9QMetaType6equalsEPKvS1_iPi(void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto: static bool QMetaType::convert(const void * from, int fromTypeId, void * to, int toTypeId);
extern bool C_ZN9QMetaType7convertEPKviPvi(void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  const QMetaObject * QMetaType::metaObject();
extern void C_ZNK9QMetaType10metaObjectEv(void* qthis); // 2
  // proto: static bool QMetaType::debugStream(QDebug & dbg, const void * rhs, int typeId);
extern bool C_ZN9QMetaType11debugStreamER6QDebugPKvi(void* arg0, void* arg1, int32_t arg2); // 4
  // proto: static const QMetaObject * QMetaType::metaObjectForType(int type);
extern void C_ZN9QMetaType17metaObjectForTypeEi(int32_t arg0); // 4
  // proto:  void QMetaType::destruct(void * data);
extern void C_ZNK9QMetaType8destructEPv(void* qthis, void* arg0); // 2
  // proto: static void QMetaType::destruct(int type, void * where);
extern void C_ZN9QMetaType8destructEiPv(int32_t arg0, void* arg1); // 4
  // proto: static bool QMetaType::hasRegisteredConverterFunction(int fromTypeId, int toTypeId);
extern bool C_ZN9QMetaType30hasRegisteredConverterFunctionEii(int32_t arg0, int32_t arg1); // 4
  // proto:  TypeFlags QMetaType::flags();
extern void C_ZNK9QMetaType5flagsEv(void* qthis); // 2
  // proto: static int QMetaType::registerTypedef(const char * typeName, int aliasId);
extern int32_t C_ZN9QMetaType15registerTypedefEPKci(void* arg0, int32_t arg1); // 4
  // proto:  int QMetaType::sizeOf();
extern int32_t C_ZNK9QMetaType6sizeOfEv(void* qthis); // 2
  // proto: static int QMetaType::sizeOf(int type);
extern int32_t C_ZN9QMetaType6sizeOfEi(int32_t arg0); // 4
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

// class sizeof(QMetaType)=80
type QMetaType struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// load(class QDataStream &, int, void *)
func (this *QMetaType) Load_S(args ...interface{}) (ret interface{}) {
  // load(class QDataStream &, int, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType4loadER11QDataStreamiPv
    // invoke: bool load(class QDataStream &, int, void *)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QMetaType4loadER11QDataStreamiPv(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "load", args)
  }

  return
}

// registerNormalizedTypedef(const ::QByteArray &, int)
func (this *QMetaType) Registernormalizedtypedef_S(args ...interface{}) (ret interface{}) {
  // registerNormalizedTypedef(const ::QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const ::QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType25registerNormalizedTypedefERK10QByteArrayi
    // invoke: int registerNormalizedTypedef(const ::QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QMetaType25registerNormalizedTypedefERK10QByteArrayi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "registerNormalizedTypedef", args)
  }

  return
}

// compare(const void *, const void *, int, int *)
func (this *QMetaType) Compare_S(args ...interface{}) (ret interface{}) {
  // compare(const void *, const void *, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType7compareEPKvS1_iPi
    // invoke: bool compare(const void *, const void *, int, int *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN9QMetaType7compareEPKvS1_iPi(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "compare", args)
  }

  return
}

// typeName(int)
func (this *QMetaType) Typename_S(args ...interface{}) (ret interface{}) {
  // typeName(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType8typeNameEi
    // invoke: const char * typeName(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType8typeNameEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "typeName", args)
  }

  return
}

// QMetaType(const int)
func NewQMetaType(args ...interface{}) *QMetaType {
  // QMetaType(const int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "const int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaTypeC1Ei
    // invoke: void QMetaType(const int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMetaTypeC2Ei(arg0)
    return &QMetaType{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaType", "QMetaType", args)
  }

  return nil // QMetaType{qclsinst:qthis}
}

// isRegistered(int)
func (this *QMetaType) Isregistered_S(args ...interface{}) (ret interface{}) {
  // isRegistered(int)
  // isRegistered()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType12isRegisteredEi
    // invoke: bool isRegistered(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType12isRegisteredEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK9QMetaType12isRegisteredEv
    // invoke: bool isRegistered()
    var ret0 = C.C_ZNK9QMetaType12isRegisteredEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "isRegistered", args)
  }

  return
}

// unregisterType(int)
func (this *QMetaType) Unregistertype_S(args ...interface{}) (ret interface{}) {
  // unregisterType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType14unregisterTypeEi
    // invoke: bool unregisterType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType14unregisterTypeEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "unregisterType", args)
  }

  return
}

// typeFlags(int)
func (this *QMetaType) Typeflags_S(args ...interface{}) () {
  // typeFlags(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType9typeFlagsEi
    // invoke: TypeFlags typeFlags(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMetaType9typeFlagsEi(arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "typeFlags", args)
  }

  return
}

// create(int, const void *)
func (this *QMetaType) Create_S(args ...interface{}) (ret interface{}) {
  // create(int, const void *)
  // create(const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType6createEiPKv
    // invoke: void * create(int, const void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QMetaType6createEiPKv(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK9QMetaType6createEPKv
    // invoke: void * create(const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMetaType6createEPKv(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "create", args)
  }

  return
}

// construct(int, void *, const void *)
func (this *QMetaType) Construct_S(args ...interface{}) (ret interface{}) {
  // construct(int, void *, const void *)
  // construct(void *, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.VoidpTy() // "const void *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "void *"
  vtys[1][1] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType9constructEiPvPKv
    // invoke: void * construct(int, void *, const void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QMetaType9constructEiPvPKv(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK9QMetaType9constructEPvPKv
    // invoke: void * construct(void *, const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QMetaType9constructEPvPKv(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "construct", args)
  }

  return
}

// hasRegisteredDebugStreamOperator(int)
func (this *QMetaType) Hasregistereddebugstreamoperator_S(args ...interface{}) (ret interface{}) {
  // hasRegisteredDebugStreamOperator(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType32hasRegisteredDebugStreamOperatorEi
    // invoke: bool hasRegisteredDebugStreamOperator(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType32hasRegisteredDebugStreamOperatorEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredDebugStreamOperator", args)
  }

  return
}

// destroy(int, void *)
func (this *QMetaType) Destroy_S(args ...interface{}) () {
  // destroy(int, void *)
  // destroy(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType7destroyEiPv
    // invoke: void destroy(int, void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C.C_ZN9QMetaType7destroyEiPv(arg0, arg1)
  case 1:
    // invoke: _ZNK9QMetaType7destroyEPv
    // invoke: void destroy(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZNK9QMetaType7destroyEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "destroy", args)
  }

  return
}

// save(class QDataStream &, int, const void *)
func (this *QMetaType) Save_S(args ...interface{}) (ret interface{}) {
  // save(class QDataStream &, int, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType4saveER11QDataStreamiPKv
    // invoke: bool save(class QDataStream &, int, const void *)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QMetaType4saveER11QDataStreamiPKv(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "save", args)
  }

  return
}

// type(const char *)
func (this *QMetaType) Type_S(args ...interface{}) (ret interface{}) {
  // type(const char *)
  // type(const ::QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const ::QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType4typeEPKc
    // invoke: int type(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN9QMetaType4typeEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN9QMetaType4typeERK10QByteArray
    // invoke: int type(const ::QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType4typeERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "type", args)
  }

  return
}

// ~QMetaType()
func (this *QMetaType) Freeqmetatype(args ...interface{}) () {
  // ~QMetaType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaTypeD0Ev
    // invoke: void ~QMetaType()
    C.C_ZN9QMetaTypeD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "~QMetaType", args)
  }

  return
}

// hasRegisteredComparators(int)
func (this *QMetaType) Hasregisteredcomparators_S(args ...interface{}) (ret interface{}) {
  // hasRegisteredComparators(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType24hasRegisteredComparatorsEi
    // invoke: bool hasRegisteredComparators(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType24hasRegisteredComparatorsEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredComparators", args)
  }

  return
}

// isValid()
func (this *QMetaType) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaType7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK9QMetaType7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "isValid", args)
  }

  return
}

// equals(const void *, const void *, int, int *)
func (this *QMetaType) Equals_S(args ...interface{}) (ret interface{}) {
  // equals(const void *, const void *, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType6equalsEPKvS1_iPi
    // invoke: bool equals(const void *, const void *, int, int *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN9QMetaType6equalsEPKvS1_iPi(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "equals", args)
  }

  return
}

// convert(const void *, int, void *, int)
func (this *QMetaType) Convert_S(args ...interface{}) (ret interface{}) {
  // convert(const void *, int, void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.VoidpTy() // "void *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType7convertEPKviPvi
    // invoke: bool convert(const void *, int, void *, int)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN9QMetaType7convertEPKviPvi(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "convert", args)
  }

  return
}

// metaObject()
func (this *QMetaType) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaType10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QMetaType10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "metaObject", args)
  }

  return
}

// debugStream(class QDebug &, const void *, int)
func (this *QMetaType) Debugstream_S(args ...interface{}) (ret interface{}) {
  // debugStream(class QDebug &, const void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDebug{}) // "QDebug &"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType11debugStreamER6QDebugPKvi
    // invoke: bool debugStream(class QDebug &, const void *, int)
    var arg0 = args[0].(QDebug).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QMetaType11debugStreamER6QDebugPKvi(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "debugStream", args)
  }

  return
}

// metaObjectForType(int)
func (this *QMetaType) Metaobjectfortype_S(args ...interface{}) () {
  // metaObjectForType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType17metaObjectForTypeEi
    // invoke: const QMetaObject * metaObjectForType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMetaType17metaObjectForTypeEi(arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "metaObjectForType", args)
  }

  return
}

// destruct(void *)
func (this *QMetaType) Destruct(args ...interface{}) () {
  // destruct(void *)
  // destruct(int, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaType8destructEPv
    // invoke: void destruct(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.C_ZNK9QMetaType8destructEPv(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QMetaType8destructEiPv
    // invoke: void destruct(int, void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C.C_ZN9QMetaType8destructEiPv(arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaType", "destruct", args)
  }

  return
}

// hasRegisteredConverterFunction(int, int)
func (this *QMetaType) Hasregisteredconverterfunction_S(args ...interface{}) (ret interface{}) {
  // hasRegisteredConverterFunction(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType30hasRegisteredConverterFunctionEii
    // invoke: bool hasRegisteredConverterFunction(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QMetaType30hasRegisteredConverterFunctionEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredConverterFunction", args)
  }

  return
}

// flags()
func (this *QMetaType) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaType5flagsEv
    // invoke: TypeFlags flags()
    C.C_ZNK9QMetaType5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "flags", args)
  }

  return
}

// registerTypedef(const char *, int)
func (this *QMetaType) Registertypedef_S(args ...interface{}) (ret interface{}) {
  // registerTypedef(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType15registerTypedefEPKci
    // invoke: int registerTypedef(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QMetaType15registerTypedefEPKci(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "registerTypedef", args)
  }

  return
}

// sizeOf()
func (this *QMetaType) Sizeof(args ...interface{}) (ret interface{}) {
  // sizeOf()
  // sizeOf(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaType6sizeOfEv
    // invoke: int sizeOf()
    var ret0 = C.C_ZNK9QMetaType6sizeOfEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN9QMetaType6sizeOfEi
    // invoke: int sizeOf(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QMetaType6sizeOfEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaType", "sizeOf", args)
  }

  return
}

// <= body block end

