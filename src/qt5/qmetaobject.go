package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK9QMetaEnum4nameEv(void* qthis); // 4
  // proto:  bool QMetaEnum::isFlag();
extern void C_ZNK9QMetaEnum6isFlagEv(void* qthis); // 4
  // proto:  bool QMetaEnum::isValid();
extern void C_ZNK9QMetaEnum7isValidEv(void* qthis); // 2
  // proto:  int QMetaEnum::keyToValue(const char * key, bool * ok);
extern void C_ZNK9QMetaEnum10keyToValueEPKcPb(void* qthis, unsigned char* arg0, bool* arg1); // 4
  // proto:  void QMetaEnum::QMetaEnum();
extern void C_ZN9QMetaEnumC2Ev(void* qthis); // 1
  // proto:  int QMetaEnum::value(int index);
extern void C_ZNK9QMetaEnum5valueEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::valueToKey(int value);
extern void C_ZNK9QMetaEnum10valueToKeyEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::key(int index);
extern void C_ZNK9QMetaEnum3keyEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QMetaEnum::valueToKeys(int value);
extern void C_ZNK9QMetaEnum11valueToKeysEi(void* qthis, int32_t arg0); // 4
  // proto:  const char * QMetaEnum::scope();
extern void C_ZNK9QMetaEnum5scopeEv(void* qthis); // 4
  // proto:  int QMetaEnum::keyCount();
extern void C_ZNK9QMetaEnum8keyCountEv(void* qthis); // 4
  // proto:  int QMetaEnum::keysToValue(const char * keys, bool * ok);
extern void C_ZNK9QMetaEnum11keysToValueEPKcPb(void* qthis, unsigned char* arg0, bool* arg1); // 4
  // proto:  const QMetaObject * QMetaEnum::enclosingMetaObject();
extern void C_ZNK9QMetaEnum19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  void QMetaClassInfo::QMetaClassInfo();
extern void C_ZN14QMetaClassInfoC2Ev(void* qthis); // 1
  // proto:  const char * QMetaClassInfo::name();
extern void C_ZNK14QMetaClassInfo4nameEv(void* qthis); // 4
  // proto:  const char * QMetaClassInfo::value();
extern void C_ZNK14QMetaClassInfo5valueEv(void* qthis); // 4
  // proto:  const QMetaObject * QMetaClassInfo::enclosingMetaObject();
extern void C_ZNK14QMetaClassInfo19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  int QMetaMethod::parameterType(int index);
extern void C_ZNK11QMetaMethod13parameterTypeEi(void* qthis, int32_t arg0); // 4
  // proto:  QList<QByteArray> QMetaMethod::parameterNames();
extern void C_ZNK11QMetaMethod14parameterNamesEv(void* qthis); // 4
  // proto:  int QMetaMethod::methodIndex();
extern void C_ZNK11QMetaMethod11methodIndexEv(void* qthis); // 4
  // proto:  bool QMetaMethod::isValid();
extern void C_ZNK11QMetaMethod7isValidEv(void* qthis); // 2
  // proto:  QByteArray QMetaMethod::name();
extern void C_ZNK11QMetaMethod4nameEv(void* qthis); // 4
  // proto:  void QMetaMethod::getParameterTypes(int * types);
extern void C_ZNK11QMetaMethod17getParameterTypesEPi(void* qthis, int32_t* arg0); // 4
  // proto:  int QMetaMethod::parameterCount();
extern void C_ZNK11QMetaMethod14parameterCountEv(void* qthis); // 4
  // proto:  QMetaMethod::Access QMetaMethod::access();
extern void C_ZNK11QMetaMethod6accessEv(void* qthis); // 4
  // proto:  const char * QMetaMethod::typeName();
extern void C_ZNK11QMetaMethod8typeNameEv(void* qthis); // 4
  // proto:  const char * QMetaMethod::tag();
extern void C_ZNK11QMetaMethod3tagEv(void* qthis); // 4
  // proto:  QMetaMethod::MethodType QMetaMethod::methodType();
extern void C_ZNK11QMetaMethod10methodTypeEv(void* qthis); // 4
  // proto:  int QMetaMethod::returnType();
extern void C_ZNK11QMetaMethod10returnTypeEv(void* qthis); // 4
  // proto:  int QMetaMethod::attributes();
extern void C_ZNK11QMetaMethod10attributesEv(void* qthis); // 4
  // proto:  QList<QByteArray> QMetaMethod::parameterTypes();
extern void C_ZNK11QMetaMethod14parameterTypesEv(void* qthis); // 4
  // proto:  int QMetaMethod::revision();
extern void C_ZNK11QMetaMethod8revisionEv(void* qthis); // 4
  // proto:  QByteArray QMetaMethod::methodSignature();
extern void C_ZNK11QMetaMethod15methodSignatureEv(void* qthis); // 4
  // proto:  void QMetaMethod::QMetaMethod();
extern void C_ZN11QMetaMethodC2Ev(void* qthis); // 1
  // proto:  const QMetaObject * QMetaMethod::enclosingMetaObject();
extern void C_ZNK11QMetaMethod19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  bool QMetaProperty::isStored(const QObject * obj);
extern void C_ZNK13QMetaProperty8isStoredEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isEditable(const QObject * obj);
extern void C_ZNK13QMetaProperty10isEditableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QMetaMethod QMetaProperty::notifySignal();
extern void C_ZNK13QMetaProperty12notifySignalEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isConstant();
extern void C_ZNK13QMetaProperty10isConstantEv(void* qthis); // 4
  // proto:  const char * QMetaProperty::typeName();
extern void C_ZNK13QMetaProperty8typeNameEv(void* qthis); // 4
  // proto:  void QMetaProperty::QMetaProperty();
extern void C_ZN13QMetaPropertyC2Ev(void* qthis); // 3
  // proto:  int QMetaProperty::notifySignalIndex();
extern void C_ZNK13QMetaProperty17notifySignalIndexEv(void* qthis); // 4
  // proto:  QVariant QMetaProperty::readOnGadget(const void * gadget);
extern void C_ZNK13QMetaProperty12readOnGadgetEPKv(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::hasStdCppSet();
extern void C_ZNK13QMetaProperty12hasStdCppSetEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isFinal();
extern void C_ZNK13QMetaProperty7isFinalEv(void* qthis); // 4
  // proto:  int QMetaProperty::propertyIndex();
extern void C_ZNK13QMetaProperty13propertyIndexEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isWritable();
extern void C_ZNK13QMetaProperty10isWritableEv(void* qthis); // 4
  // proto:  bool QMetaProperty::write(QObject * obj, const QVariant & value);
extern void C_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QVariant::Type QMetaProperty::type();
extern void C_ZNK13QMetaProperty4typeEv(void* qthis); // 4
  // proto:  const QMetaObject * QMetaProperty::enclosingMetaObject();
extern void C_ZNK13QMetaProperty19enclosingMetaObjectEv(void* qthis); // 2
  // proto:  bool QMetaProperty::reset(QObject * obj);
extern void C_ZNK13QMetaProperty5resetEP7QObject(void* qthis, void* arg0); // 4
  // proto:  QMetaEnum QMetaProperty::enumerator();
extern void C_ZNK13QMetaProperty10enumeratorEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isEnumType();
extern void C_ZNK13QMetaProperty10isEnumTypeEv(void* qthis); // 4
  // proto:  QVariant QMetaProperty::read(const QObject * obj);
extern void C_ZNK13QMetaProperty4readEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isValid();
extern void C_ZNK13QMetaProperty7isValidEv(void* qthis); // 2
  // proto:  bool QMetaProperty::isDesignable(const QObject * obj);
extern void C_ZNK13QMetaProperty12isDesignableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::writeOnGadget(void * gadget, const QVariant & value);
extern void C_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  int QMetaProperty::userType();
extern void C_ZNK13QMetaProperty8userTypeEv(void* qthis); // 4
  // proto:  bool QMetaProperty::hasNotifySignal();
extern void C_ZNK13QMetaProperty15hasNotifySignalEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isScriptable(const QObject * obj);
extern void C_ZNK13QMetaProperty12isScriptableEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isFlagType();
extern void C_ZNK13QMetaProperty10isFlagTypeEv(void* qthis); // 4
  // proto:  const char * QMetaProperty::name();
extern void C_ZNK13QMetaProperty4nameEv(void* qthis); // 4
  // proto:  bool QMetaProperty::resetOnGadget(void * gadget);
extern void C_ZNK13QMetaProperty13resetOnGadgetEPv(void* qthis, void* arg0); // 4
  // proto:  bool QMetaProperty::isResettable();
extern void C_ZNK13QMetaProperty12isResettableEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isReadable();
extern void C_ZNK13QMetaProperty10isReadableEv(void* qthis); // 4
  // proto:  int QMetaProperty::revision();
extern void C_ZNK13QMetaProperty8revisionEv(void* qthis); // 4
  // proto:  bool QMetaProperty::isUser(const QObject * obj);
extern void C_ZNK13QMetaProperty6isUserEPK7QObject(void* qthis, void* arg0); // 4
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
func (this *QMetaEnum) name(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "name", args)
  }

}

// isFlag()
func (this *QMetaEnum) isFlag(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum6isFlagEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "isFlag", args)
  }

}

// isValid()
func (this *QMetaEnum) isValid(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "isValid", args)
  }

}

// keyToValue(const char *, _Bool *)
func (this *QMetaEnum) keyToValue(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C.C_ZNK9QMetaEnum10keyToValueEPKcPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyToValue", args)
  }

}

// QMetaEnum()
func NewQMetaEnum(args ...interface{}) QMetaEnum {
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
    C.C_ZN9QMetaEnumC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMetaEnum", "QMetaEnum", args)
  }

  return QMetaEnum{}
}

// value(int)
func (this *QMetaEnum) value(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum5valueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "value", args)
  }

}

// valueToKey(int)
func (this *QMetaEnum) valueToKey(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum10valueToKeyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKey", args)
  }

}

// key(int)
func (this *QMetaEnum) key(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum3keyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "key", args)
  }

}

// valueToKeys(int)
func (this *QMetaEnum) valueToKeys(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum11valueToKeysEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKeys", args)
  }

}

// scope()
func (this *QMetaEnum) scope(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum5scopeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "scope", args)
  }

}

// keyCount()
func (this *QMetaEnum) keyCount(args ...interface{}) () {
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
    C.C_ZNK9QMetaEnum8keyCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyCount", args)
  }

}

// keysToValue(const char *, _Bool *)
func (this *QMetaEnum) keysToValue(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C.C_ZNK9QMetaEnum11keysToValueEPKcPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keysToValue", args)
  }

}

// enclosingMetaObject()
func (this *QMetaEnum) enclosingMetaObject(args ...interface{}) () {
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

}

// QMetaClassInfo()
func NewQMetaClassInfo(args ...interface{}) QMetaClassInfo {
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
    C.C_ZN14QMetaClassInfoC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "QMetaClassInfo", args)
  }

  return QMetaClassInfo{}
}

// name()
func (this *QMetaClassInfo) name(args ...interface{}) () {
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
    C.C_ZNK14QMetaClassInfo4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "name", args)
  }

}

// value()
func (this *QMetaClassInfo) value(args ...interface{}) () {
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
    C.C_ZNK14QMetaClassInfo5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "value", args)
  }

}

// enclosingMetaObject()
func (this *QMetaClassInfo) enclosingMetaObject(args ...interface{}) () {
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

}

// parameterType(int)
func (this *QMetaMethod) parameterType(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod13parameterTypeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterType", args)
  }

}

// parameterNames()
func (this *QMetaMethod) parameterNames(args ...interface{}) () {
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

}

// methodIndex()
func (this *QMetaMethod) methodIndex(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod11methodIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodIndex", args)
  }

}

// isValid()
func (this *QMetaMethod) isValid(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "isValid", args)
  }

}

// name()
func (this *QMetaMethod) name(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "name", args)
  }

}

// getParameterTypes(int *)
func (this *QMetaMethod) getParameterTypes(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QMetaMethod17getParameterTypesEPi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaMethod", "getParameterTypes", args)
  }

}

// parameterCount()
func (this *QMetaMethod) parameterCount(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod14parameterCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterCount", args)
  }

}

// access()
func (this *QMetaMethod) access(args ...interface{}) () {
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

}

// typeName()
func (this *QMetaMethod) typeName(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod8typeNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "typeName", args)
  }

}

// tag()
func (this *QMetaMethod) tag(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod3tagEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "tag", args)
  }

}

// methodType()
func (this *QMetaMethod) methodType(args ...interface{}) () {
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

}

// returnType()
func (this *QMetaMethod) returnType(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod10returnTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "returnType", args)
  }

}

// attributes()
func (this *QMetaMethod) attributes(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod10attributesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "attributes", args)
  }

}

// parameterTypes()
func (this *QMetaMethod) parameterTypes(args ...interface{}) () {
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

}

// revision()
func (this *QMetaMethod) revision(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "revision", args)
  }

}

// methodSignature()
func (this *QMetaMethod) methodSignature(args ...interface{}) () {
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
    C.C_ZNK11QMetaMethod15methodSignatureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodSignature", args)
  }

}

// QMetaMethod()
func NewQMetaMethod(args ...interface{}) QMetaMethod {
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
    C.C_ZN11QMetaMethodC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMetaMethod", "QMetaMethod", args)
  }

  return QMetaMethod{}
}

// enclosingMetaObject()
func (this *QMetaMethod) enclosingMetaObject(args ...interface{}) () {
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

}

// isStored(const class QObject *)
func (this *QMetaProperty) isStored(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty8isStoredEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isStored", args)
  }

}

// isEditable(const class QObject *)
func (this *QMetaProperty) isEditable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isEditableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEditable", args)
  }

}

// notifySignal()
func (this *QMetaProperty) notifySignal(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12notifySignalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignal", args)
  }

}

// isConstant()
func (this *QMetaProperty) isConstant(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isConstantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isConstant", args)
  }

}

// typeName()
func (this *QMetaProperty) typeName(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty8typeNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "typeName", args)
  }

}

// QMetaProperty()
func NewQMetaProperty(args ...interface{}) QMetaProperty {
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
    C.C_ZN13QMetaPropertyC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMetaProperty", "QMetaProperty", args)
  }

  return QMetaProperty{}
}

// notifySignalIndex()
func (this *QMetaProperty) notifySignalIndex(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty17notifySignalIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignalIndex", args)
  }

}

// readOnGadget(const void *)
func (this *QMetaProperty) readOnGadget(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12readOnGadgetEPKv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "readOnGadget", args)
  }

}

// hasStdCppSet()
func (this *QMetaProperty) hasStdCppSet(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12hasStdCppSetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasStdCppSet", args)
  }

}

// isFinal()
func (this *QMetaProperty) isFinal(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty7isFinalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFinal", args)
  }

}

// propertyIndex()
func (this *QMetaProperty) propertyIndex(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty13propertyIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "propertyIndex", args)
  }

}

// isWritable()
func (this *QMetaProperty) isWritable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isWritableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isWritable", args)
  }

}

// write(class QObject *, const class QVariant &)
func (this *QMetaProperty) write(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaProperty", "write", args)
  }

}

// type()
func (this *QMetaProperty) type_(args ...interface{}) () {
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

}

// enclosingMetaObject()
func (this *QMetaProperty) enclosingMetaObject(args ...interface{}) () {
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

}

// reset(class QObject *)
func (this *QMetaProperty) reset(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty5resetEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "reset", args)
  }

}

// enumerator()
func (this *QMetaProperty) enumerator(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10enumeratorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "enumerator", args)
  }

}

// isEnumType()
func (this *QMetaProperty) isEnumType(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isEnumTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEnumType", args)
  }

}

// read(const class QObject *)
func (this *QMetaProperty) read(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty4readEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "read", args)
  }

}

// isValid()
func (this *QMetaProperty) isValid(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isValid", args)
  }

}

// isDesignable(const class QObject *)
func (this *QMetaProperty) isDesignable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12isDesignableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isDesignable", args)
  }

}

// writeOnGadget(void *, const class QVariant &)
func (this *QMetaProperty) writeOnGadget(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaProperty", "writeOnGadget", args)
  }

}

// userType()
func (this *QMetaProperty) userType(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty8userTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "userType", args)
  }

}

// hasNotifySignal()
func (this *QMetaProperty) hasNotifySignal(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty15hasNotifySignalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasNotifySignal", args)
  }

}

// isScriptable(const class QObject *)
func (this *QMetaProperty) isScriptable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12isScriptableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isScriptable", args)
  }

}

// isFlagType()
func (this *QMetaProperty) isFlagType(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isFlagTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFlagType", args)
  }

}

// name()
func (this *QMetaProperty) name(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "name", args)
  }

}

// resetOnGadget(void *)
func (this *QMetaProperty) resetOnGadget(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty13resetOnGadgetEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "resetOnGadget", args)
  }

}

// isResettable()
func (this *QMetaProperty) isResettable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty12isResettableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isResettable", args)
  }

}

// isReadable()
func (this *QMetaProperty) isReadable(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty10isReadableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isReadable", args)
  }

}

// revision()
func (this *QMetaProperty) revision(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "revision", args)
  }

}

// isUser(const class QObject *)
func (this *QMetaProperty) isUser(args ...interface{}) () {
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
    C.C_ZNK13QMetaProperty6isUserEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isUser", args)
  }

}

// <= body block end

