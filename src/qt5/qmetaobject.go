package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
extern void _ZNK9QMetaEnum5valueEi(void* qthis, int arg0);
  // proto:  const char * QMetaEnum::name();
extern void _ZNK9QMetaEnum4nameEv(void* qthis);
  // proto:  bool QMetaEnum::isFlag();
extern void _ZNK9QMetaEnum6isFlagEv(void* qthis);
  // proto:  const char * QMetaEnum::scope();
extern void _ZNK9QMetaEnum5scopeEv(void* qthis);
  // proto:  int QMetaEnum::keyToValue(const char * key, bool * ok);
extern void _ZNK9QMetaEnum10keyToValueEPKcPb(void* qthis, char* arg0, bool* arg1);
  // proto:  const QMetaObject * QMetaEnum::enclosingMetaObject();
extern void demth_ZNK9QMetaEnum19enclosingMetaObjectEv(void* qthis);
  // proto:  QByteArray QMetaEnum::valueToKeys(int value);
extern void _ZNK9QMetaEnum11valueToKeysEi(void* qthis, int arg0);
  // proto:  void QMetaEnum::QMetaEnum();
extern void* dector_ZN9QMetaEnumC1Ev();
extern void _ZN9QMetaEnumC1Ev(void* qthis);
  // proto:  int QMetaEnum::keysToValue(const char * keys, bool * ok);
extern void _ZNK9QMetaEnum11keysToValueEPKcPb(void* qthis, char* arg0, bool* arg1);
  // proto:  const char * QMetaEnum::key(int index);
extern void _ZNK9QMetaEnum3keyEi(void* qthis, int arg0);
  // proto:  const char * QMetaEnum::valueToKey(int value);
extern void _ZNK9QMetaEnum10valueToKeyEi(void* qthis, int arg0);
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
extern void _ZNK11QMetaMethod17getParameterTypesEPi(void* qthis, int* arg0);
  // proto:  void QMetaMethod::QMetaMethod();
extern void* dector_ZN11QMetaMethodC1Ev();
extern void _ZN11QMetaMethodC1Ev(void* qthis);
  // proto:  int QMetaMethod::parameterType(int index);
extern void _ZNK11QMetaMethod13parameterTypeEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMetaClassInfo)=16
type QMetaClassInfo struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMetaMethod)=16
type QMetaMethod struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMetaProperty)=32
type QMetaProperty struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
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
  default:
    qtrt.ErrorResolve("QMetaProperty", "isDesignable", args)
  }

}

// <= body block end

