package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  int QMetaEnum::value(int index);
extern void _ZNK9QMetaEnum5valueEi(void* qthis, int32_t arg0);
  // proto:  const char * QMetaEnum::name();
extern void _ZNK9QMetaEnum4nameEv(void* qthis);
  // proto:  bool QMetaEnum::isFlag();
extern void _ZNK9QMetaEnum6isFlagEv(void* qthis);
  // proto:  const char * QMetaEnum::scope();
extern void _ZNK9QMetaEnum5scopeEv(void* qthis);
  // proto:  int QMetaEnum::keyToValue(const char * key, bool * ok);
extern void _ZNK9QMetaEnum10keyToValueEPKcPb(void* qthis, unsigned char* arg0, bool* arg1);
  // proto:  const QMetaObject * QMetaEnum::enclosingMetaObject();
extern void demth_ZNK9QMetaEnum19enclosingMetaObjectEv(void* qthis);
  // proto:  QByteArray QMetaEnum::valueToKeys(int value);
extern void _ZNK9QMetaEnum11valueToKeysEi(void* qthis, int32_t arg0);
  // proto:  void QMetaEnum::QMetaEnum();
extern void* dector_ZN9QMetaEnumC1Ev();
extern void _ZN9QMetaEnumC1Ev(void* qthis);
  // proto:  int QMetaEnum::keysToValue(const char * keys, bool * ok);
extern void _ZNK9QMetaEnum11keysToValueEPKcPb(void* qthis, unsigned char* arg0, bool* arg1);
  // proto:  const char * QMetaEnum::key(int index);
extern void _ZNK9QMetaEnum3keyEi(void* qthis, int32_t arg0);
  // proto:  const char * QMetaEnum::valueToKey(int value);
extern void _ZNK9QMetaEnum10valueToKeyEi(void* qthis, int32_t arg0);
  // proto:  int QMetaEnum::keyCount();
extern void _ZNK9QMetaEnum8keyCountEv(void* qthis);
  // proto:  bool QMetaEnum::isValid();
extern void demth_ZNK9QMetaEnum7isValidEv(void* qthis);
  // proto:  void QMetaClassInfo::QMetaClassInfo();
extern void* dector_ZN14QMetaClassInfoC1Ev();
extern void _ZN14QMetaClassInfoC1Ev(void* qthis);
  // proto:  const QMetaObject * QMetaClassInfo::enclosingMetaObject();
extern void demth_ZNK14QMetaClassInfo19enclosingMetaObjectEv(void* qthis);
  // proto:  const char * QMetaClassInfo::name();
extern void _ZNK14QMetaClassInfo4nameEv(void* qthis);
  // proto:  const char * QMetaClassInfo::value();
extern void _ZNK14QMetaClassInfo5valueEv(void* qthis);
  // proto:  QList<QByteArray> QMetaMethod::parameterTypes();
extern void _ZNK11QMetaMethod14parameterTypesEv(void* qthis);
  // proto:  QList<QByteArray> QMetaMethod::parameterNames();
extern void _ZNK11QMetaMethod14parameterNamesEv(void* qthis);
  // proto:  QByteArray QMetaMethod::methodSignature();
extern void _ZNK11QMetaMethod15methodSignatureEv(void* qthis);
  // proto:  const char * QMetaMethod::typeName();
extern void _ZNK11QMetaMethod8typeNameEv(void* qthis);
  // proto:  int QMetaMethod::attributes();
extern void _ZNK11QMetaMethod10attributesEv(void* qthis);
  // proto:  void QMetaMethod::getParameterTypes(int * types);
extern void _ZNK11QMetaMethod17getParameterTypesEPi(void* qthis, int32_t* arg0);
  // proto:  void QMetaMethod::QMetaMethod();
extern void* dector_ZN11QMetaMethodC1Ev();
extern void _ZN11QMetaMethodC1Ev(void* qthis);
  // proto:  int QMetaMethod::parameterType(int index);
extern void _ZNK11QMetaMethod13parameterTypeEi(void* qthis, int32_t arg0);
  // proto:  QByteArray QMetaMethod::name();
extern void _ZNK11QMetaMethod4nameEv(void* qthis);
  // proto:  int QMetaMethod::returnType();
extern void _ZNK11QMetaMethod10returnTypeEv(void* qthis);
  // proto:  int QMetaMethod::methodIndex();
extern void _ZNK11QMetaMethod11methodIndexEv(void* qthis);
  // proto:  int QMetaMethod::parameterCount();
extern void _ZNK11QMetaMethod14parameterCountEv(void* qthis);
  // proto:  const QMetaObject * QMetaMethod::enclosingMetaObject();
extern void demth_ZNK11QMetaMethod19enclosingMetaObjectEv(void* qthis);
  // proto:  int QMetaMethod::revision();
extern void _ZNK11QMetaMethod8revisionEv(void* qthis);
  // proto:  const char * QMetaMethod::tag();
extern void _ZNK11QMetaMethod3tagEv(void* qthis);
  // proto:  bool QMetaMethod::isValid();
extern void demth_ZNK11QMetaMethod7isValidEv(void* qthis);
  // proto:  bool QMetaProperty::isEnumType();
extern void _ZNK13QMetaProperty10isEnumTypeEv(void* qthis);
  // proto:  void QMetaProperty::QMetaProperty();
extern void* dector_ZN13QMetaPropertyC1Ev();
extern void _ZN13QMetaPropertyC1Ev(void* qthis);
  // proto:  bool QMetaProperty::isValid();
extern void demth_ZNK13QMetaProperty7isValidEv(void* qthis);
  // proto:  QVariant QMetaProperty::readOnGadget(const void * gadget);
extern void _ZNK13QMetaProperty12readOnGadgetEPKv(void* qthis, void* arg0);
  // proto:  const QMetaObject * QMetaProperty::enclosingMetaObject();
extern void demth_ZNK13QMetaProperty19enclosingMetaObjectEv(void* qthis);
  // proto:  bool QMetaProperty::resetOnGadget(void * gadget);
extern void _ZNK13QMetaProperty13resetOnGadgetEPv(void* qthis, void* arg0);
  // proto:  int QMetaProperty::propertyIndex();
extern void _ZNK13QMetaProperty13propertyIndexEv(void* qthis);
  // proto:  bool QMetaProperty::isStored(const QObject * obj);
extern void _ZNK13QMetaProperty8isStoredEPK7QObject(void* qthis, void* arg0);
  // proto:  QMetaEnum QMetaProperty::enumerator();
extern void _ZNK13QMetaProperty10enumeratorEv(void* qthis);
  // proto:  bool QMetaProperty::write(QObject * obj, const QVariant & value);
extern void _ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  bool QMetaProperty::isResettable();
extern void _ZNK13QMetaProperty12isResettableEv(void* qthis);
  // proto:  bool QMetaProperty::isEditable(const QObject * obj);
extern void _ZNK13QMetaProperty10isEditableEPK7QObject(void* qthis, void* arg0);
  // proto:  bool QMetaProperty::hasStdCppSet();
extern void _ZNK13QMetaProperty12hasStdCppSetEv(void* qthis);
  // proto:  bool QMetaProperty::hasNotifySignal();
extern void _ZNK13QMetaProperty15hasNotifySignalEv(void* qthis);
  // proto:  bool QMetaProperty::isConstant();
extern void _ZNK13QMetaProperty10isConstantEv(void* qthis);
  // proto:  const char * QMetaProperty::typeName();
extern void _ZNK13QMetaProperty8typeNameEv(void* qthis);
  // proto:  bool QMetaProperty::isReadable();
extern void _ZNK13QMetaProperty10isReadableEv(void* qthis);
  // proto:  int QMetaProperty::userType();
extern void _ZNK13QMetaProperty8userTypeEv(void* qthis);
  // proto:  bool QMetaProperty::isWritable();
extern void _ZNK13QMetaProperty10isWritableEv(void* qthis);
  // proto:  bool QMetaProperty::writeOnGadget(void * gadget, const QVariant & value);
extern void _ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  int QMetaProperty::notifySignalIndex();
extern void _ZNK13QMetaProperty17notifySignalIndexEv(void* qthis);
  // proto:  bool QMetaProperty::isUser(const QObject * obj);
extern void _ZNK13QMetaProperty6isUserEPK7QObject(void* qthis, void* arg0);
  // proto:  bool QMetaProperty::isFlagType();
extern void _ZNK13QMetaProperty10isFlagTypeEv(void* qthis);
  // proto:  bool QMetaProperty::isFinal();
extern void _ZNK13QMetaProperty7isFinalEv(void* qthis);
  // proto:  const char * QMetaProperty::name();
extern void _ZNK13QMetaProperty4nameEv(void* qthis);
  // proto:  bool QMetaProperty::reset(QObject * obj);
extern void _ZNK13QMetaProperty5resetEP7QObject(void* qthis, void* arg0);
  // proto:  int QMetaProperty::revision();
extern void _ZNK13QMetaProperty8revisionEv(void* qthis);
  // proto:  bool QMetaProperty::isScriptable(const QObject * obj);
extern void _ZNK13QMetaProperty12isScriptableEPK7QObject(void* qthis, void* arg0);
  // proto:  QVariant QMetaProperty::read(const QObject * obj);
extern void _ZNK13QMetaProperty4readEPK7QObject(void* qthis, void* arg0);
  // proto:  QMetaMethod QMetaProperty::notifySignal();
extern void _ZNK13QMetaProperty12notifySignalEv(void* qthis);
  // proto:  bool QMetaProperty::isDesignable(const QObject * obj);
extern void _ZNK13QMetaProperty12isDesignableEPK7QObject(void* qthis, void* arg0);
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

  // proto:  int QMetaEnum::value(int index);
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
    C._ZNK9QMetaEnum5valueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "value", args)
  }

}

  // proto:  const char * QMetaEnum::name();
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
    C._ZNK9QMetaEnum4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "name", args)
  }

}

  // proto:  bool QMetaEnum::isFlag();
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
    C._ZNK9QMetaEnum6isFlagEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "isFlag", args)
  }

}

  // proto:  const char * QMetaEnum::scope();
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
    C._ZNK9QMetaEnum5scopeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "scope", args)
  }

}

  // proto:  int QMetaEnum::keyToValue(const char * key, bool * ok);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK9QMetaEnum10keyToValueEPKcPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyToValue", args)
  }

}

  // proto:  const QMetaObject * QMetaEnum::enclosingMetaObject();
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
    C.demth_ZNK9QMetaEnum19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "enclosingMetaObject", args)
  }

}

  // proto:  QByteArray QMetaEnum::valueToKeys(int value);
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
    C._ZNK9QMetaEnum11valueToKeysEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKeys", args)
  }

}

  // proto:  void QMetaEnum::QMetaEnum();
func NewQMetaEnum(args ...interface{}) QMetaEnum {
  return QMetaEnum{}
}

  // proto:  int QMetaEnum::keysToValue(const char * keys, bool * ok);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    C._ZNK9QMetaEnum11keysToValueEPKcPb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keysToValue", args)
  }

}

  // proto:  const char * QMetaEnum::key(int index);
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
    C._ZNK9QMetaEnum3keyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "key", args)
  }

}

  // proto:  const char * QMetaEnum::valueToKey(int value);
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
    C._ZNK9QMetaEnum10valueToKeyEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaEnum", "valueToKey", args)
  }

}

  // proto:  int QMetaEnum::keyCount();
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
    C._ZNK9QMetaEnum8keyCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "keyCount", args)
  }

}

  // proto:  bool QMetaEnum::isValid();
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
    C.demth_ZNK9QMetaEnum7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaEnum", "isValid", args)
  }

}

  // proto:  void QMetaClassInfo::QMetaClassInfo();
func NewQMetaClassInfo(args ...interface{}) QMetaClassInfo {
  return QMetaClassInfo{}
}

  // proto:  const QMetaObject * QMetaClassInfo::enclosingMetaObject();
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
    C.demth_ZNK14QMetaClassInfo19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "enclosingMetaObject", args)
  }

}

  // proto:  const char * QMetaClassInfo::name();
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
    C._ZNK14QMetaClassInfo4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "name", args)
  }

}

  // proto:  const char * QMetaClassInfo::value();
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
    C._ZNK14QMetaClassInfo5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaClassInfo", "value", args)
  }

}

  // proto:  QList<QByteArray> QMetaMethod::parameterTypes();
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
    C._ZNK11QMetaMethod14parameterTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterTypes", args)
  }

}

  // proto:  QList<QByteArray> QMetaMethod::parameterNames();
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
    C._ZNK11QMetaMethod14parameterNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterNames", args)
  }

}

  // proto:  QByteArray QMetaMethod::methodSignature();
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
    C._ZNK11QMetaMethod15methodSignatureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodSignature", args)
  }

}

  // proto:  const char * QMetaMethod::typeName();
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
    C._ZNK11QMetaMethod8typeNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "typeName", args)
  }

}

  // proto:  int QMetaMethod::attributes();
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
    C._ZNK11QMetaMethod10attributesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "attributes", args)
  }

}

  // proto:  void QMetaMethod::getParameterTypes(int * types);
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
    C._ZNK11QMetaMethod17getParameterTypesEPi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaMethod", "getParameterTypes", args)
  }

}

  // proto:  void QMetaMethod::QMetaMethod();
func NewQMetaMethod(args ...interface{}) QMetaMethod {
  return QMetaMethod{}
}

  // proto:  int QMetaMethod::parameterType(int index);
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
    C._ZNK11QMetaMethod13parameterTypeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterType", args)
  }

}

  // proto:  QByteArray QMetaMethod::name();
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
    C._ZNK11QMetaMethod4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "name", args)
  }

}

  // proto:  int QMetaMethod::returnType();
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
    C._ZNK11QMetaMethod10returnTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "returnType", args)
  }

}

  // proto:  int QMetaMethod::methodIndex();
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
    C._ZNK11QMetaMethod11methodIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "methodIndex", args)
  }

}

  // proto:  int QMetaMethod::parameterCount();
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
    C._ZNK11QMetaMethod14parameterCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "parameterCount", args)
  }

}

  // proto:  const QMetaObject * QMetaMethod::enclosingMetaObject();
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
    C.demth_ZNK11QMetaMethod19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "enclosingMetaObject", args)
  }

}

  // proto:  int QMetaMethod::revision();
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
    C._ZNK11QMetaMethod8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "revision", args)
  }

}

  // proto:  const char * QMetaMethod::tag();
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
    C._ZNK11QMetaMethod3tagEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "tag", args)
  }

}

  // proto:  bool QMetaMethod::isValid();
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
    C.demth_ZNK11QMetaMethod7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaMethod", "isValid", args)
  }

}

  // proto:  bool QMetaProperty::isEnumType();
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
    C._ZNK13QMetaProperty10isEnumTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEnumType", args)
  }

}

  // proto:  void QMetaProperty::QMetaProperty();
func NewQMetaProperty(args ...interface{}) QMetaProperty {
  return QMetaProperty{}
}

  // proto:  bool QMetaProperty::isValid();
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
    C.demth_ZNK13QMetaProperty7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isValid", args)
  }

}

  // proto:  QVariant QMetaProperty::readOnGadget(const void * gadget);
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
    C._ZNK13QMetaProperty12readOnGadgetEPKv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "readOnGadget", args)
  }

}

  // proto:  const QMetaObject * QMetaProperty::enclosingMetaObject();
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
    C.demth_ZNK13QMetaProperty19enclosingMetaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "enclosingMetaObject", args)
  }

}

  // proto:  bool QMetaProperty::resetOnGadget(void * gadget);
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
    C._ZNK13QMetaProperty13resetOnGadgetEPv(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "resetOnGadget", args)
  }

}

  // proto:  int QMetaProperty::propertyIndex();
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
    C._ZNK13QMetaProperty13propertyIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "propertyIndex", args)
  }

}

  // proto:  bool QMetaProperty::isStored(const QObject * obj);
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
    C._ZNK13QMetaProperty8isStoredEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isStored", args)
  }

}

  // proto:  QMetaEnum QMetaProperty::enumerator();
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
    C._ZNK13QMetaProperty10enumeratorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "enumerator", args)
  }

}

  // proto:  bool QMetaProperty::write(QObject * obj, const QVariant & value);
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
    C._ZNK13QMetaProperty5writeEP7QObjectRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaProperty", "write", args)
  }

}

  // proto:  bool QMetaProperty::isResettable();
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
    C._ZNK13QMetaProperty12isResettableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isResettable", args)
  }

}

  // proto:  bool QMetaProperty::isEditable(const QObject * obj);
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
    C._ZNK13QMetaProperty10isEditableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isEditable", args)
  }

}

  // proto:  bool QMetaProperty::hasStdCppSet();
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
    C._ZNK13QMetaProperty12hasStdCppSetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasStdCppSet", args)
  }

}

  // proto:  bool QMetaProperty::hasNotifySignal();
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
    C._ZNK13QMetaProperty15hasNotifySignalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "hasNotifySignal", args)
  }

}

  // proto:  bool QMetaProperty::isConstant();
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
    C._ZNK13QMetaProperty10isConstantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isConstant", args)
  }

}

  // proto:  const char * QMetaProperty::typeName();
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
    C._ZNK13QMetaProperty8typeNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "typeName", args)
  }

}

  // proto:  bool QMetaProperty::isReadable();
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
    C._ZNK13QMetaProperty10isReadableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isReadable", args)
  }

}

  // proto:  int QMetaProperty::userType();
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
    C._ZNK13QMetaProperty8userTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "userType", args)
  }

}

  // proto:  bool QMetaProperty::isWritable();
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
    C._ZNK13QMetaProperty10isWritableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isWritable", args)
  }

}

  // proto:  bool QMetaProperty::writeOnGadget(void * gadget, const QVariant & value);
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
    C._ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMetaProperty", "writeOnGadget", args)
  }

}

  // proto:  int QMetaProperty::notifySignalIndex();
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
    C._ZNK13QMetaProperty17notifySignalIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignalIndex", args)
  }

}

  // proto:  bool QMetaProperty::isUser(const QObject * obj);
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
    C._ZNK13QMetaProperty6isUserEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isUser", args)
  }

}

  // proto:  bool QMetaProperty::isFlagType();
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
    C._ZNK13QMetaProperty10isFlagTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFlagType", args)
  }

}

  // proto:  bool QMetaProperty::isFinal();
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
    C._ZNK13QMetaProperty7isFinalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isFinal", args)
  }

}

  // proto:  const char * QMetaProperty::name();
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
    C._ZNK13QMetaProperty4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "name", args)
  }

}

  // proto:  bool QMetaProperty::reset(QObject * obj);
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
    C._ZNK13QMetaProperty5resetEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "reset", args)
  }

}

  // proto:  int QMetaProperty::revision();
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
    C._ZNK13QMetaProperty8revisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "revision", args)
  }

}

  // proto:  bool QMetaProperty::isScriptable(const QObject * obj);
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
    C._ZNK13QMetaProperty12isScriptableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isScriptable", args)
  }

}

  // proto:  QVariant QMetaProperty::read(const QObject * obj);
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
    C._ZNK13QMetaProperty4readEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "read", args)
  }

}

  // proto:  QMetaMethod QMetaProperty::notifySignal();
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
    C._ZNK13QMetaProperty12notifySignalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMetaProperty", "notifySignal", args)
  }

}

  // proto:  bool QMetaProperty::isDesignable(const QObject * obj);
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
    C._ZNK13QMetaProperty12isDesignableEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMetaProperty", "isDesignable", args)
  }

}

// <= body block end

