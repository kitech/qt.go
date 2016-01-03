package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto: static void QMetaType::destroy(int type, void * data);
extern void _ZN9QMetaType7destroyEiPv(int32_t arg0, void* arg1);
  // proto: static bool QMetaType::hasRegisteredConverterFunction(int fromTypeId, int toTypeId);
extern void _ZN9QMetaType30hasRegisteredConverterFunctionEii(int32_t arg0, int32_t arg1);
  // proto:  const QMetaObject * QMetaType::metaObject();
extern void demth_ZNK9QMetaType10metaObjectEv(void* qthis);
  // proto: static bool QMetaType::hasRegisteredDebugStreamOperator(int typeId);
extern void _ZN9QMetaType32hasRegisteredDebugStreamOperatorEi(int32_t arg0);
  // proto: static void * QMetaType::create(int type, const void * copy);
extern void _ZN9QMetaType6createEiPKv(int32_t arg0, void* arg1);
  // proto:  void QMetaType::destroy(void * data);
extern void demth_ZNK9QMetaType7destroyEPv(void* qthis, void* arg0);
  // proto: static int QMetaType::registerTypedef(const char * typeName, int aliasId);
extern void _ZN9QMetaType15registerTypedefEPKci(unsigned char* arg0, int32_t arg1);
  // proto: static void QMetaType::destruct(int type, void * where);
extern void _ZN9QMetaType8destructEiPv(int32_t arg0, void* arg1);
  // proto:  bool QMetaType::isValid();
extern void demth_ZNK9QMetaType7isValidEv(void* qthis);
  // proto: static void * QMetaType::construct(int type, void * where, const void * copy);
extern void _ZN9QMetaType9constructEiPvPKv(int32_t arg0, void* arg1, void* arg2);
  // proto: static bool QMetaType::equals(const void * lhs, const void * rhs, int typeId, int * result);
extern void _ZN9QMetaType6equalsEPKvS1_iPi(void* arg0, void* arg1, int32_t arg2, int32_t* arg3);
  // proto:  void * QMetaType::construct(void * where, const void * copy);
extern void demth_ZNK9QMetaType9constructEPvPKv(void* qthis, void* arg0, void* arg1);
  // proto: static bool QMetaType::isRegistered(int type);
extern void _ZN9QMetaType12isRegisteredEi(int32_t arg0);
  // proto: static bool QMetaType::unregisterType(int type);
extern void _ZN9QMetaType14unregisterTypeEi(int32_t arg0);
  // proto:  void QMetaType::QMetaType(const QMetaType & other);
extern void* dector_ZN9QMetaTypeC1ERKS_(void* arg0);
extern void _ZN9QMetaTypeC1ERKS_(void* qthis, void* arg0);
  // proto: static const QMetaObject * QMetaType::metaObjectForType(int type);
extern void _ZN9QMetaType17metaObjectForTypeEi(int32_t arg0);
  // proto: static bool QMetaType::load(QDataStream & stream, int type, void * data);
extern void _ZN9QMetaType4loadER11QDataStreamiPv(void* arg0, int32_t arg1, void* arg2);
  // proto:  void * QMetaType::create(const void * copy);
extern void demth_ZNK9QMetaType6createEPKv(void* qthis, void* arg0);
  // proto: static int QMetaType::sizeOf(int type);
extern void _ZN9QMetaType6sizeOfEi(int32_t arg0);
  // proto: static bool QMetaType::hasRegisteredComparators(int typeId);
extern void _ZN9QMetaType24hasRegisteredComparatorsEi(int32_t arg0);
  // proto: static bool QMetaType::save(QDataStream & stream, int type, const void * data);
extern void _ZN9QMetaType4saveER11QDataStreamiPKv(void* arg0, int32_t arg1, void* arg2);
  // proto:  void QMetaType::destruct(void * data);
extern void demth_ZNK9QMetaType8destructEPv(void* qthis, void* arg0);
  // proto:  void QMetaType::~QMetaType();
extern void demth_ZN9QMetaTypeD0Ev(void* qthis);
  // proto: static int QMetaType::type(const char * typeName);
extern void _ZN9QMetaType4typeEPKc(unsigned char* arg0);
  // proto: static int QMetaType::type(const ::QByteArray & typeName);
extern void _ZN9QMetaType4typeERK10QByteArray(void* arg0);
  // proto: static bool QMetaType::debugStream(QDebug & dbg, const void * rhs, int typeId);
extern void _ZN9QMetaType11debugStreamER6QDebugPKvi(void* arg0, void* arg1, int32_t arg2);
  // proto:  int QMetaType::sizeOf();
extern void demth_ZNK9QMetaType6sizeOfEv(void* qthis);
  // proto: static const char * QMetaType::typeName(int type);
extern void _ZN9QMetaType8typeNameEi(int32_t arg0);
  // proto: static bool QMetaType::convert(const void * from, int fromTypeId, void * to, int toTypeId);
extern void _ZN9QMetaType7convertEPKviPvi(void* arg0, int32_t arg1, void* arg2, int32_t arg3);
  // proto:  void QMetaType::QMetaType(const int type);
extern void* dector_ZN9QMetaTypeC1Ei(int32_t arg0);
extern void _ZN9QMetaTypeC1Ei(void* qthis, int32_t arg0);
  // proto: static int QMetaType::registerNormalizedTypedef(const ::QByteArray & normalizedTypeName, int aliasId);
extern void _ZN9QMetaType25registerNormalizedTypedefERK10QByteArrayi(void* arg0, int32_t arg1);
  // proto: static bool QMetaType::compare(const void * lhs, const void * rhs, int typeId, int * result);
extern void _ZN9QMetaType7compareEPKvS1_iPi(void* arg0, void* arg1, int32_t arg2, int32_t* arg3);
  // proto:  bool QMetaType::isRegistered();
extern void demth_ZNK9QMetaType12isRegisteredEv(void* qthis);
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

  // proto: static void QMetaType::destroy(int type, void * data);
func (this *QMetaType) destroy_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "destroy", args)
  }

}

  // proto: static bool QMetaType::hasRegisteredConverterFunction(int fromTypeId, int toTypeId);
func (this *QMetaType) hasRegisteredConverterFunction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredConverterFunction", args)
  }

}

  // proto:  const QMetaObject * QMetaType::metaObject();
func (this *QMetaType) metaObject(args ...interface{}) () {
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
    C.demth_ZNK9QMetaType10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "metaObject", args)
  }

}

  // proto: static bool QMetaType::hasRegisteredDebugStreamOperator(int typeId);
func (this *QMetaType) hasRegisteredDebugStreamOperator_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredDebugStreamOperator", args)
  }

}

  // proto: static void * QMetaType::create(int type, const void * copy);
func (this *QMetaType) create_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "create", args)
  }

}

  // proto:  void QMetaType::destroy(void * data);
func (this *QMetaType) destroy(args ...interface{}) () {
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
    C._ZN9QMetaType7destroyEiPv(arg0, arg1)
  case 1:
    // invoke: _ZNK9QMetaType7destroyEPv
    // invoke: void destroy(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.demth_ZNK9QMetaType7destroyEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "destroy", args)
  }

}

  // proto: static int QMetaType::registerTypedef(const char * typeName, int aliasId);
func (this *QMetaType) registerTypedef_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "registerTypedef", args)
  }

}

  // proto: static void QMetaType::destruct(int type, void * where);
func (this *QMetaType) destruct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "destruct", args)
  }

}

  // proto:  bool QMetaType::isValid();
func (this *QMetaType) isValid(args ...interface{}) () {
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
    C.demth_ZNK9QMetaType7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "isValid", args)
  }

}

  // proto: static void * QMetaType::construct(int type, void * where, const void * copy);
func (this *QMetaType) construct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "construct", args)
  }

}

  // proto: static bool QMetaType::equals(const void * lhs, const void * rhs, int typeId, int * result);
func (this *QMetaType) equals_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "equals", args)
  }

}

  // proto:  void * QMetaType::construct(void * where, const void * copy);
func (this *QMetaType) construct(args ...interface{}) () {
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
    C._ZN9QMetaType9constructEiPvPKv(arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK9QMetaType9constructEPvPKv
    // invoke: void * construct(void *, const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C.demth_ZNK9QMetaType9constructEPvPKv(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaType", "construct", args)
  }

}

  // proto: static bool QMetaType::isRegistered(int type);
func (this *QMetaType) isRegistered_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "isRegistered", args)
  }

}

  // proto: static bool QMetaType::unregisterType(int type);
func (this *QMetaType) unregisterType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "unregisterType", args)
  }

}

  // proto:  void QMetaType::QMetaType(const QMetaType & other);
func NewQMetaType(args ...interface{}) QMetaType {
  return QMetaType{}
}

  // proto: static const QMetaObject * QMetaType::metaObjectForType(int type);
func (this *QMetaType) metaObjectForType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "metaObjectForType", args)
  }

}

  // proto: static bool QMetaType::load(QDataStream & stream, int type, void * data);
func (this *QMetaType) load_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "load", args)
  }

}

  // proto:  void * QMetaType::create(const void * copy);
func (this *QMetaType) create(args ...interface{}) () {
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
    C._ZN9QMetaType6createEiPKv(arg0, arg1)
  case 1:
    // invoke: _ZNK9QMetaType6createEPKv
    // invoke: void * create(const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.demth_ZNK9QMetaType6createEPKv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "create", args)
  }

}

  // proto: static int QMetaType::sizeOf(int type);
func (this *QMetaType) sizeOf_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "sizeOf", args)
  }

}

  // proto: static bool QMetaType::hasRegisteredComparators(int typeId);
func (this *QMetaType) hasRegisteredComparators_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "hasRegisteredComparators", args)
  }

}

  // proto: static bool QMetaType::save(QDataStream & stream, int type, const void * data);
func (this *QMetaType) save_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "save", args)
  }

}

  // proto:  void QMetaType::destruct(void * data);
func (this *QMetaType) destruct(args ...interface{}) () {
  // destruct(int, void *)
  // destruct(void *)
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
    // invoke: _ZN9QMetaType8destructEiPv
    // invoke: void destruct(int, void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C._ZN9QMetaType8destructEiPv(arg0, arg1)
  case 1:
    // invoke: _ZNK9QMetaType8destructEPv
    // invoke: void destruct(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    C.demth_ZNK9QMetaType8destructEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaType", "destruct", args)
  }

}

  // proto:  void QMetaType::~QMetaType();
func (this *QMetaType) FreeQMetaType(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "~QMetaType", args)
  }

}

  // proto: static int QMetaType::type(const char * typeName);
func (this *QMetaType) type_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "type", args)
  }

}

  // proto: static bool QMetaType::debugStream(QDebug & dbg, const void * rhs, int typeId);
func (this *QMetaType) debugStream_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "debugStream", args)
  }

}

  // proto:  int QMetaType::sizeOf();
func (this *QMetaType) sizeOf(args ...interface{}) () {
  // sizeOf(int)
  // sizeOf()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaType6sizeOfEi
    // invoke: int sizeOf(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QMetaType6sizeOfEi(arg0)
  case 1:
    // invoke: _ZNK9QMetaType6sizeOfEv
    // invoke: int sizeOf()
    C.demth_ZNK9QMetaType6sizeOfEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "sizeOf", args)
  }

}

  // proto: static const char * QMetaType::typeName(int type);
func (this *QMetaType) typeName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "typeName", args)
  }

}

  // proto: static bool QMetaType::convert(const void * from, int fromTypeId, void * to, int toTypeId);
func (this *QMetaType) convert_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "convert", args)
  }

}

  // proto: static int QMetaType::registerNormalizedTypedef(const ::QByteArray & normalizedTypeName, int aliasId);
func (this *QMetaType) registerNormalizedTypedef_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "registerNormalizedTypedef", args)
  }

}

  // proto: static bool QMetaType::compare(const void * lhs, const void * rhs, int typeId, int * result);
func (this *QMetaType) compare_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMetaType", "compare", args)
  }

}

  // proto:  bool QMetaType::isRegistered();
func (this *QMetaType) isRegistered(args ...interface{}) () {
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
    C._ZN9QMetaType12isRegisteredEi(arg0)
  case 1:
    // invoke: _ZNK9QMetaType12isRegisteredEv
    // invoke: bool isRegistered()
    C.demth_ZNK9QMetaType12isRegisteredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaType", "isRegistered", args)
  }

}

// <= body block end

