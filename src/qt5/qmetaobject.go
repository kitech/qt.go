package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qmetaobject.h
// dst-file: /src/core/qmetaobject.go
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
  // proto:  const char * QMetaEnum::name();
extern void* C_ZNK9QMetaEnum4nameEv(void* qthis); // 4
  // proto:  bool QMetaEnum::isFlag();
extern bool C_ZNK9QMetaEnum6isFlagEv(void* qthis); // 4
  // proto:  bool QMetaEnum::isValid();
extern bool C_ZNK9QMetaEnum7isValidEv(void* qthis); // 2
  // proto:  int QMetaEnum::keyToValue(const char * key, bool * ok);
extern int32_t C_ZNK9QMetaEnum10keyToValueEPKcPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMetaEnum::QMetaEnum();
extern void* C_ZN9QMetaEnumC2Ev(); // 1
  // proto:  int QMetaEnum::value(int index);
extern int32_t C_ZNK9QMetaEnum5valueEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::valueToKey(int value);
extern void* C_ZNK9QMetaEnum10valueToKeyEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::key(int index);
extern void* C_ZNK9QMetaEnum3keyEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QMetaEnum::valueToKeys(int value);
extern void* C_ZNK9QMetaEnum11valueToKeysEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::scope();
extern void* C_ZNK9QMetaEnum5scopeEv(void* qthis); // 4
  // proto:  int QMetaEnum::keyCount();
extern int32_t C_ZNK9QMetaEnum8keyCountEv(void* qthis); // 4
  // proto:  int QMetaEnum::keysToValue(const char * keys, bool * ok);
extern int32_t C_ZNK9QMetaEnum11keysToValueEPKcPb(void* qthis, void* arg0, void* arg1); // 4
  // proto:  const QMetaObject * QMetaEnum::enclosingMetaObject();
extern void C_ZNK9QMetaEnum19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  void QMetaClassInfo::QMetaClassInfo();
extern void* C_ZN14QMetaClassInfoC2Ev(); // 1
  // proto:  const char * QMetaClassInfo::name();
extern void* C_ZNK14QMetaClassInfo4nameEv(void* qthis); // 4
  // proto:  const char * QMetaClassInfo::value();
extern void* C_ZNK14QMetaClassInfo5valueEv(void* qthis); // 4
  // proto:  const QMetaObject * QMetaClassInfo::enclosingMetaObject();
extern void C_ZNK14QMetaClassInfo19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  int QMetaMethod::parameterType(int index);
extern int32_t C_ZNK11QMetaMethod13parameterTypeEi(void* qthis, int32_t arg0); // 4
  // proto:  QList<QByteArray> QMetaMethod::parameterNames();
extern void C_ZNK11QMetaMethod14parameterNamesEv(void* qthis); // 4
  // proto:  int QMetaMethod::methodIndex();
extern int32_t C_ZNK11QMetaMethod11methodIndexEv(void* qthis); // 4
  // proto:  bool QMetaMethod::isValid();
extern bool C_ZNK11QMetaMethod7isValidEv(void* qthis); // 2
  // proto:  QByteArray QMetaMethod::name();
extern void* C_ZNK11QMetaMethod4nameEv(void* qthis); // 4
  // proto:  void QMetaMethod::getParameterTypes(int * types);
extern void C_ZNK11QMetaMethod17getParameterTypesEPi(void* qthis, void* arg0); // 4
  // proto:  int QMetaMethod::parameterCount();
extern int32_t C_ZNK11QMetaMethod14parameterCountEv(void* qthis); // 4
  // proto:  QMetaMethod::Access QMetaMethod::access();
extern void C_ZNK11QMetaMethod6accessEv(void* qthis); // 4
  // proto:  const char * QMetaMethod::typeName();
extern void* C_ZNK11QMetaMethod8typeNameEv(void* qthis); // 4
  // proto:  const char * QMetaMethod::tag();
extern void* C_ZNK11QMetaMethod3tagEv(void* qthis); // 4
  // proto:  QMetaMethod::MethodType QMetaMethod::methodType();
extern void C_ZNK11QMetaMethod10methodTypeEv(void* qthis); // 4
  // proto:  int QMetaMethod::returnType();
extern int32_t C_ZNK11QMetaMethod10returnTypeEv(void* qthis); // 4
  // proto:  int QMetaMethod::attributes();
extern int32_t C_ZNK11QMetaMethod10attributesEv(void* qthis); // 4
  // proto:  QList<QByteArray> QMetaMethod::parameterTypes();
extern void C_ZNK11QMetaMethod14parameterTypesEv(void* qthis); // 4
  // proto:  int QMetaMethod::revision();
extern int32_t C_ZNK11QMetaMethod8revisionEv(void* qthis); // 4
  // proto:  QByteArray QMetaMethod::methodSignature();
extern void* C_ZNK11QMetaMethod15methodSignatureEv(void* qthis); // 4
  // proto:  void QMetaMethod::QMetaMethod();
extern void* C_ZN11QMetaMethodC2Ev(); // 1
  // proto:  const QMetaObject * QMetaMethod::enclosingMetaObject();
extern void C_ZNK11QMetaMethod19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  bool QMetaProperty::isStored(const QObject * obj);
extern bool C_ZNK13QMetaProperty8isStoredEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isEditable(const QObject * obj);
extern bool C_ZNK13QMetaProperty10isEditableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QMetaMethod QMetaProperty::notifySignal();
extern void* C_ZNK13QMetaProperty12notifySignalEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isConstant();
extern bool C_ZNK13QMetaProperty10isConstantEv(void* qthis); // 4
  // proto:  const char * QMetaProperty::typeName();
extern void* C_ZNK13QMetaProperty8typeNameEv(void* qthis); // 4
  // proto:  void QMetaProperty::QMetaProperty();
extern void* C_ZN13QMetaPropertyC2Ev(); // 3
  // proto:  int QMetaProperty::notifySignalIndex();
extern int32_t C_ZNK13QMetaProperty17notifySignalIndexEv(void* qthis); // 4
  // proto:  QVariant QMetaProperty::readOnGadget(const void * gadget);
extern void* C_ZNK13QMetaProperty12readOnGadgetEPKv(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::hasStdCppSet();
extern bool C_ZNK13QMetaProperty12hasStdCppSetEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isFinal();
extern bool C_ZNK13QMetaProperty7isFinalEv(void* qthis); // 4
  // proto:  int QMetaProperty::propertyIndex();
extern int32_t C_ZNK13QMetaProperty13propertyIndexEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isWritable();
extern bool C_ZNK13QMetaProperty10isWritableEv(void* qthis); // 4
  // proto:  bool QMetaProperty::write(QObject * obj, const QVariant & value);
extern bool C_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QVariant::Type QMetaProperty::type();
extern void C_ZNK13QMetaProperty4typeEv(void* qthis); // 4
  // proto:  const QMetaObject * QMetaProperty::enclosingMetaObject();
extern void C_ZNK13QMetaProperty19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  bool QMetaProperty::reset(QObject * obj);
extern bool C_ZNK13QMetaProperty5resetEP7QObject(void* qthis, void* arg0); // 4
  // proto:  QMetaEnum QMetaProperty::enumerator();
extern void* C_ZNK13QMetaProperty10enumeratorEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isEnumType();
extern bool C_ZNK13QMetaProperty10isEnumTypeEv(void* qthis); // 4
  // proto:  QVariant QMetaProperty::read(const QObject * obj);
extern void* C_ZNK13QMetaProperty4readEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isValid();
extern bool C_ZNK13QMetaProperty7isValidEv(void* qthis); // 2
  // proto:  bool QMetaProperty::isDesignable(const QObject * obj);
extern bool C_ZNK13QMetaProperty12isDesignableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::writeOnGadget(void * gadget, const QVariant & value);
extern bool C_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  int QMetaProperty::userType();
extern int32_t C_ZNK13QMetaProperty8userTypeEv(void* qthis); // 4
  // proto:  bool QMetaProperty::hasNotifySignal();
extern bool C_ZNK13QMetaProperty15hasNotifySignalEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isScriptable(const QObject * obj);
extern bool C_ZNK13QMetaProperty12isScriptableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isFlagType();
extern bool C_ZNK13QMetaProperty10isFlagTypeEv(void* qthis); // 4
  // proto:  const char * QMetaProperty::name();
extern void* C_ZNK13QMetaProperty4nameEv(void* qthis); // 4
  // proto:  bool QMetaProperty::resetOnGadget(void * gadget);
extern bool C_ZNK13QMetaProperty13resetOnGadgetEPv(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isResettable();
extern bool C_ZNK13QMetaProperty12isResettableEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isReadable();
extern bool C_ZNK13QMetaProperty10isReadableEv(void* qthis); // 4
  // proto:  int QMetaProperty::revision();
extern int32_t C_ZNK13QMetaProperty8revisionEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isUser(const QObject * obj);
extern bool C_ZNK13QMetaProperty6isUserEPK7QObject(void* qthis, void* arg0); // 4
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

// class sizeof(QMetaEnum)=16
type QMetaEnum struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMetaClassInfo)=16
type QMetaClassInfo struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMetaMethod)=16
type QMetaMethod struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMetaProperty)=32
type QMetaProperty struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// name()
func (this *QMetaEnum) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum4nameEv
    // invoke: const char * name()
    var ret0 = C.C_ZNK9QMetaEnum4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "name", args)
  }

  return
}

// isFlag()
func (this *QMetaEnum) Isflag(args ...interface{}) (ret interface{}) {
  // isFlag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum6isFlagEv
    // invoke: bool isFlag()
    var ret0 = C.C_ZNK9QMetaEnum6isFlagEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "isFlag", args)
  }

  return
}

// isValid()
func (this *QMetaEnum) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK9QMetaEnum7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "isValid", args)
  }

  return
}

// keyToValue(const char *, _Bool *)
func (this *QMetaEnum) Keytovalue(args ...interface{}) (ret interface{}) {
  // keyToValue(const char *, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum10keyToValueEPKcPb
    // invoke: int keyToValue(const char *, _Bool *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QMetaEnum10keyToValueEPKcPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyToValue", args)
  }

  return
}

// QMetaEnum()
func NewQMetaEnum(args ...interface{}) *QMetaEnum {
  // QMetaEnum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMetaEnumC1Ev
    // invoke: void QMetaEnum()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMetaEnumC2Ev()
    return &QMetaEnum{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaEnum", "QMetaEnum", args)
  }

  return nil // QMetaEnum{qclsinst:qthis}
}

// value(int)
func (this *QMetaEnum) Value(args ...interface{}) (ret interface{}) {
  // value(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum5valueEi
    // invoke: int value(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMetaEnum5valueEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "value", args)
  }

  return
}

// valueToKey(int)
func (this *QMetaEnum) Valuetokey(args ...interface{}) (ret interface{}) {
  // valueToKey(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum10valueToKeyEi
    // invoke: const char * valueToKey(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMetaEnum10valueToKeyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKey", args)
  }

  return
}

// key(int)
func (this *QMetaEnum) Key(args ...interface{}) (ret interface{}) {
  // key(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum3keyEi
    // invoke: const char * key(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMetaEnum3keyEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "key", args)
  }

  return
}

// valueToKeys(int)
func (this *QMetaEnum) Valuetokeys(args ...interface{}) (ret interface{}) {
  // valueToKeys(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum11valueToKeysEi
    // invoke: QByteArray valueToKeys(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMetaEnum11valueToKeysEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKeys", args)
  }

  return
}

// scope()
func (this *QMetaEnum) Scope(args ...interface{}) (ret interface{}) {
  // scope()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum5scopeEv
    // invoke: const char * scope()
    var ret0 = C.C_ZNK9QMetaEnum5scopeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "scope", args)
  }

  return
}

// keyCount()
func (this *QMetaEnum) Keycount(args ...interface{}) (ret interface{}) {
  // keyCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum8keyCountEv
    // invoke: int keyCount()
    var ret0 = C.C_ZNK9QMetaEnum8keyCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyCount", args)
  }

  return
}

// keysToValue(const char *, _Bool *)
func (this *QMetaEnum) Keystovalue(args ...interface{}) (ret interface{}) {
  // keysToValue(const char *, _Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum11keysToValueEPKcPb
    // invoke: int keysToValue(const char *, _Bool *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QMetaEnum11keysToValueEPKcPb(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaEnum", "keysToValue", args)
  }

  return
}

// enclosingMetaObject()
func (this *QMetaEnum) Enclosingmetaobject(args ...interface{}) () {
  // enclosingMetaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMetaEnum19enclosingMetaObjectEv
    // invoke: const QMetaObject * enclosingMetaObject()
    C.C_ZNK9QMetaEnum19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "enclosingMetaObject", args)
  }

  return
}

// QMetaClassInfo()
func NewQMetaClassInfo(args ...interface{}) *QMetaClassInfo {
  // QMetaClassInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QMetaClassInfoC1Ev
    // invoke: void QMetaClassInfo()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QMetaClassInfoC2Ev()
    return &QMetaClassInfo{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "QMetaClassInfo", args)
  }

  return nil // QMetaClassInfo{qclsinst:qthis}
}

// name()
func (this *QMetaClassInfo) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMetaClassInfo4nameEv
    // invoke: const char * name()
    var ret0 = C.C_ZNK14QMetaClassInfo4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "name", args)
  }

  return
}

// value()
func (this *QMetaClassInfo) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMetaClassInfo5valueEv
    // invoke: const char * value()
    var ret0 = C.C_ZNK14QMetaClassInfo5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "value", args)
  }

  return
}

// enclosingMetaObject()
func (this *QMetaClassInfo) Enclosingmetaobject(args ...interface{}) () {
  // enclosingMetaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMetaClassInfo19enclosingMetaObjectEv
    // invoke: const QMetaObject * enclosingMetaObject()
    C.C_ZNK14QMetaClassInfo19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "enclosingMetaObject", args)
  }

  return
}

// parameterType(int)
func (this *QMetaMethod) Parametertype(args ...interface{}) (ret interface{}) {
  // parameterType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod13parameterTypeEi
    // invoke: int parameterType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QMetaMethod13parameterTypeEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterType", args)
  }

  return
}

// parameterNames()
func (this *QMetaMethod) Parameternames(args ...interface{}) () {
  // parameterNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod14parameterNamesEv
    // invoke: QList<QByteArray> parameterNames()
    C.C_ZNK11QMetaMethod14parameterNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterNames", args)
  }

  return
}

// methodIndex()
func (this *QMetaMethod) Methodindex(args ...interface{}) (ret interface{}) {
  // methodIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod11methodIndexEv
    // invoke: int methodIndex()
    var ret0 = C.C_ZNK11QMetaMethod11methodIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodIndex", args)
  }

  return
}

// isValid()
func (this *QMetaMethod) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK11QMetaMethod7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "isValid", args)
  }

  return
}

// name()
func (this *QMetaMethod) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod4nameEv
    // invoke: QByteArray name()
    var ret0 = C.C_ZNK11QMetaMethod4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "name", args)
  }

  return
}

// getParameterTypes(int *)
func (this *QMetaMethod) Getparametertypes(args ...interface{}) () {
  // getParameterTypes(int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod17getParameterTypesEPi
    // invoke: void getParameterTypes(int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMetaMethod17getParameterTypesEPi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaMethod", "getParameterTypes", args)
  }

  return
}

// parameterCount()
func (this *QMetaMethod) Parametercount(args ...interface{}) (ret interface{}) {
  // parameterCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod14parameterCountEv
    // invoke: int parameterCount()
    var ret0 = C.C_ZNK11QMetaMethod14parameterCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterCount", args)
  }

  return
}

// access()
func (this *QMetaMethod) Access(args ...interface{}) () {
  // access()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod6accessEv
    // invoke: QMetaMethod::Access access()
    C.C_ZNK11QMetaMethod6accessEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "access", args)
  }

  return
}

// typeName()
func (this *QMetaMethod) Typename(args ...interface{}) (ret interface{}) {
  // typeName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod8typeNameEv
    // invoke: const char * typeName()
    var ret0 = C.C_ZNK11QMetaMethod8typeNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "typeName", args)
  }

  return
}

// tag()
func (this *QMetaMethod) Tag(args ...interface{}) (ret interface{}) {
  // tag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod3tagEv
    // invoke: const char * tag()
    var ret0 = C.C_ZNK11QMetaMethod3tagEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "tag", args)
  }

  return
}

// methodType()
func (this *QMetaMethod) Methodtype(args ...interface{}) () {
  // methodType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod10methodTypeEv
    // invoke: QMetaMethod::MethodType methodType()
    C.C_ZNK11QMetaMethod10methodTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodType", args)
  }

  return
}

// returnType()
func (this *QMetaMethod) Returntype(args ...interface{}) (ret interface{}) {
  // returnType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod10returnTypeEv
    // invoke: int returnType()
    var ret0 = C.C_ZNK11QMetaMethod10returnTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "returnType", args)
  }

  return
}

// attributes()
func (this *QMetaMethod) Attributes(args ...interface{}) (ret interface{}) {
  // attributes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod10attributesEv
    // invoke: int attributes()
    var ret0 = C.C_ZNK11QMetaMethod10attributesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "attributes", args)
  }

  return
}

// parameterTypes()
func (this *QMetaMethod) Parametertypes(args ...interface{}) () {
  // parameterTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod14parameterTypesEv
    // invoke: QList<QByteArray> parameterTypes()
    C.C_ZNK11QMetaMethod14parameterTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterTypes", args)
  }

  return
}

// revision()
func (this *QMetaMethod) Revision(args ...interface{}) (ret interface{}) {
  // revision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod8revisionEv
    // invoke: int revision()
    var ret0 = C.C_ZNK11QMetaMethod8revisionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "revision", args)
  }

  return
}

// methodSignature()
func (this *QMetaMethod) Methodsignature(args ...interface{}) (ret interface{}) {
  // methodSignature()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod15methodSignatureEv
    // invoke: QByteArray methodSignature()
    var ret0 = C.C_ZNK11QMetaMethod15methodSignatureEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodSignature", args)
  }

  return
}

// QMetaMethod()
func NewQMetaMethod(args ...interface{}) *QMetaMethod {
  // QMetaMethod()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMetaMethodC1Ev
    // invoke: void QMetaMethod()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QMetaMethodC2Ev()
    return &QMetaMethod{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaMethod", "QMetaMethod", args)
  }

  return nil // QMetaMethod{qclsinst:qthis}
}

// enclosingMetaObject()
func (this *QMetaMethod) Enclosingmetaobject(args ...interface{}) () {
  // enclosingMetaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMetaMethod19enclosingMetaObjectEv
    // invoke: const QMetaObject * enclosingMetaObject()
    C.C_ZNK11QMetaMethod19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "enclosingMetaObject", args)
  }

  return
}

// isStored(const class QObject *)
func (this *QMetaProperty) Isstored(args ...interface{}) (ret interface{}) {
  // isStored(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty8isStoredEPK7QObject
    // invoke: bool isStored(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty8isStoredEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isStored", args)
  }

  return
}

// isEditable(const class QObject *)
func (this *QMetaProperty) Iseditable(args ...interface{}) (ret interface{}) {
  // isEditable(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isEditableEPK7QObject
    // invoke: bool isEditable(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty10isEditableEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEditable", args)
  }

  return
}

// notifySignal()
func (this *QMetaProperty) Notifysignal(args ...interface{}) (ret interface{}) {
  // notifySignal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12notifySignalEv
    // invoke: QMetaMethod notifySignal()
    var ret0 = C.C_ZNK13QMetaProperty12notifySignalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaMethod{}) // "QMetaMethod"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignal", args)
  }

  return
}

// isConstant()
func (this *QMetaProperty) Isconstant(args ...interface{}) (ret interface{}) {
  // isConstant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isConstantEv
    // invoke: bool isConstant()
    var ret0 = C.C_ZNK13QMetaProperty10isConstantEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isConstant", args)
  }

  return
}

// typeName()
func (this *QMetaProperty) Typename(args ...interface{}) (ret interface{}) {
  // typeName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty8typeNameEv
    // invoke: const char * typeName()
    var ret0 = C.C_ZNK13QMetaProperty8typeNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "typeName", args)
  }

  return
}

// QMetaProperty()
func NewQMetaProperty(args ...interface{}) *QMetaProperty {
  // QMetaProperty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMetaPropertyC1Ev
    // invoke: void QMetaProperty()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QMetaPropertyC2Ev()
    return &QMetaProperty{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMetaProperty", "QMetaProperty", args)
  }

  return nil // QMetaProperty{qclsinst:qthis}
}

// notifySignalIndex()
func (this *QMetaProperty) Notifysignalindex(args ...interface{}) (ret interface{}) {
  // notifySignalIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty17notifySignalIndexEv
    // invoke: int notifySignalIndex()
    var ret0 = C.C_ZNK13QMetaProperty17notifySignalIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignalIndex", args)
  }

  return
}

// readOnGadget(const void *)
func (this *QMetaProperty) Readongadget(args ...interface{}) (ret interface{}) {
  // readOnGadget(const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12readOnGadgetEPKv
    // invoke: QVariant readOnGadget(const void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty12readOnGadgetEPKv(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "readOnGadget", args)
  }

  return
}

// hasStdCppSet()
func (this *QMetaProperty) Hasstdcppset(args ...interface{}) (ret interface{}) {
  // hasStdCppSet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12hasStdCppSetEv
    // invoke: bool hasStdCppSet()
    var ret0 = C.C_ZNK13QMetaProperty12hasStdCppSetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasStdCppSet", args)
  }

  return
}

// isFinal()
func (this *QMetaProperty) Isfinal(args ...interface{}) (ret interface{}) {
  // isFinal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty7isFinalEv
    // invoke: bool isFinal()
    var ret0 = C.C_ZNK13QMetaProperty7isFinalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFinal", args)
  }

  return
}

// propertyIndex()
func (this *QMetaProperty) Propertyindex(args ...interface{}) (ret interface{}) {
  // propertyIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty13propertyIndexEv
    // invoke: int propertyIndex()
    var ret0 = C.C_ZNK13QMetaProperty13propertyIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "propertyIndex", args)
  }

  return
}

// isWritable()
func (this *QMetaProperty) Iswritable(args ...interface{}) (ret interface{}) {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isWritableEv
    // invoke: bool isWritable()
    var ret0 = C.C_ZNK13QMetaProperty10isWritableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isWritable", args)
  }

  return
}

// write(class QObject *, const class QVariant &)
func (this *QMetaProperty) Write(args ...interface{}) (ret interface{}) {
  // write(class QObject *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty5writeEP7QObjectRK8QVariant
    // invoke: bool write(class QObject *, const class QVariant &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "write", args)
  }

  return
}

// type()
func (this *QMetaProperty) Type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty4typeEv
    // invoke: QVariant::Type type()
    C.C_ZNK13QMetaProperty4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "type", args)
  }

  return
}

// enclosingMetaObject()
func (this *QMetaProperty) Enclosingmetaobject(args ...interface{}) () {
  // enclosingMetaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty19enclosingMetaObjectEv
    // invoke: const QMetaObject * enclosingMetaObject()
    C.C_ZNK13QMetaProperty19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "enclosingMetaObject", args)
  }

  return
}

// reset(class QObject *)
func (this *QMetaProperty) Reset(args ...interface{}) (ret interface{}) {
  // reset(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty5resetEP7QObject
    // invoke: bool reset(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty5resetEP7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "reset", args)
  }

  return
}

// enumerator()
func (this *QMetaProperty) Enumerator(args ...interface{}) (ret interface{}) {
  // enumerator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10enumeratorEv
    // invoke: QMetaEnum enumerator()
    var ret0 = C.C_ZNK13QMetaProperty10enumeratorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMetaEnum{}) // "QMetaEnum"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "enumerator", args)
  }

  return
}

// isEnumType()
func (this *QMetaProperty) Isenumtype(args ...interface{}) (ret interface{}) {
  // isEnumType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isEnumTypeEv
    // invoke: bool isEnumType()
    var ret0 = C.C_ZNK13QMetaProperty10isEnumTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEnumType", args)
  }

  return
}

// read(const class QObject *)
func (this *QMetaProperty) Read(args ...interface{}) (ret interface{}) {
  // read(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty4readEPK7QObject
    // invoke: QVariant read(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty4readEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "read", args)
  }

  return
}

// isValid()
func (this *QMetaProperty) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK13QMetaProperty7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isValid", args)
  }

  return
}

// isDesignable(const class QObject *)
func (this *QMetaProperty) Isdesignable(args ...interface{}) (ret interface{}) {
  // isDesignable(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12isDesignableEPK7QObject
    // invoke: bool isDesignable(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty12isDesignableEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isDesignable", args)
  }

  return
}

// writeOnGadget(void *, const class QVariant &)
func (this *QMetaProperty) Writeongadget(args ...interface{}) (ret interface{}) {
  // writeOnGadget(void *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant
    // invoke: bool writeOnGadget(void *, const class QVariant &)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "writeOnGadget", args)
  }

  return
}

// userType()
func (this *QMetaProperty) Usertype(args ...interface{}) (ret interface{}) {
  // userType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty8userTypeEv
    // invoke: int userType()
    var ret0 = C.C_ZNK13QMetaProperty8userTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "userType", args)
  }

  return
}

// hasNotifySignal()
func (this *QMetaProperty) Hasnotifysignal(args ...interface{}) (ret interface{}) {
  // hasNotifySignal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty15hasNotifySignalEv
    // invoke: bool hasNotifySignal()
    var ret0 = C.C_ZNK13QMetaProperty15hasNotifySignalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasNotifySignal", args)
  }

  return
}

// isScriptable(const class QObject *)
func (this *QMetaProperty) Isscriptable(args ...interface{}) (ret interface{}) {
  // isScriptable(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12isScriptableEPK7QObject
    // invoke: bool isScriptable(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty12isScriptableEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isScriptable", args)
  }

  return
}

// isFlagType()
func (this *QMetaProperty) Isflagtype(args ...interface{}) (ret interface{}) {
  // isFlagType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isFlagTypeEv
    // invoke: bool isFlagType()
    var ret0 = C.C_ZNK13QMetaProperty10isFlagTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFlagType", args)
  }

  return
}

// name()
func (this *QMetaProperty) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty4nameEv
    // invoke: const char * name()
    var ret0 = C.C_ZNK13QMetaProperty4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "name", args)
  }

  return
}

// resetOnGadget(void *)
func (this *QMetaProperty) Resetongadget(args ...interface{}) (ret interface{}) {
  // resetOnGadget(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty13resetOnGadgetEPv
    // invoke: bool resetOnGadget(void *)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty13resetOnGadgetEPv(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "resetOnGadget", args)
  }

  return
}

// isResettable()
func (this *QMetaProperty) Isresettable(args ...interface{}) (ret interface{}) {
  // isResettable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty12isResettableEv
    // invoke: bool isResettable()
    var ret0 = C.C_ZNK13QMetaProperty12isResettableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isResettable", args)
  }

  return
}

// isReadable()
func (this *QMetaProperty) Isreadable(args ...interface{}) (ret interface{}) {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty10isReadableEv
    // invoke: bool isReadable()
    var ret0 = C.C_ZNK13QMetaProperty10isReadableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isReadable", args)
  }

  return
}

// revision()
func (this *QMetaProperty) Revision(args ...interface{}) (ret interface{}) {
  // revision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty8revisionEv
    // invoke: int revision()
    var ret0 = C.C_ZNK13QMetaProperty8revisionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "revision", args)
  }

  return
}

// isUser(const class QObject *)
func (this *QMetaProperty) Isuser(args ...interface{}) (ret interface{}) {
  // isUser(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMetaProperty6isUserEPK7QObject
    // invoke: bool isUser(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QMetaProperty6isUserEPK7QObject(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMetaProperty", "isUser", args)
  }

  return
}

// <= body block end

