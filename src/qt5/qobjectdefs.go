package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qobjectdefs.h
// dst-file: /src/core/qobjectdefs.go
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
// class sizeof(Connection)=8
type Connection struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGenericReturnArgument)=16
type QGenericReturnArgument struct {
  /*qbase*/ QGenericArgument;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMetaObject)=48
type QMetaObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGenericArgument)=16
type QGenericArgument struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewConnection(args ...interface{})() {
}


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


func NewQGenericReturnArgument(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSlot", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfConstructor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "enumerator", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfMethod", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "constructor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorOffset", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "property", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "userProperty", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfProperty", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfClassInfo", args)
 }

}


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
  case 1:
    // invoke: _ZNK11QMetaObject4castEP7QObject
  default:
    qtrt.ErrorResolve("QMetaObject", "cast", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "method", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "superClass", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfo", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "className", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfSignal", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "constructorCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "propertyOffset", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "enumeratorCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoOffset", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "methodOffset", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "indexOfEnumerator", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "methodCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMetaObject", "classInfoCount", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGenericArgument", "name", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGenericArgument", "data", args)
 }

}


func NewQGenericArgument(args ...interface{})() {
}

// <= body block end

