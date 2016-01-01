package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qmetaobject.h
// dst-file: /src/core/qmetaobject.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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


func NewQMetaEnum(args ...interface{}) QMetaEnum {
  return QMetaEnum{}
}


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


func NewQMetaClassInfo(args ...interface{}) QMetaClassInfo {
  return QMetaClassInfo{}
}


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


func NewQMetaMethod(args ...interface{}) QMetaMethod {
  return QMetaMethod{}
}


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


func NewQMetaProperty(args ...interface{}) QMetaProperty {
  return QMetaProperty{}
}


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

