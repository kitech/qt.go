package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qsignalmapper.h
// dst-file: /src/core/qsignalmapper.go
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
// class sizeof(QSignalMapper)=1
type QSignalMapper struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _mapped QSignalMapper_mapped_signal;
}


func (this *QSignalMapper) removeMappings(args ...interface{}) () {
  // removeMappings(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper14removeMappingsEP7QObject
  default:
    qtrt.ErrorResolve("QSignalMapper", "removeMappings", args)
 }

}


func (this *QSignalMapper) map_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalMapper", "map", args)
 }

}


func NewQSignalMapper(args ...interface{})() {
}


func (this *QSignalMapper) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper10metaObjectEv
  default:
    qtrt.ErrorResolve("QSignalMapper", "metaObject", args)
 }

}


func (this *QSignalMapper) setMapping(args ...interface{}) () {
  // setMapping(class QObject *, class QObject *)
  // setMapping(class QObject *, int)
  // setMapping(class QObject *, const class QString &)
  // setMapping(class QObject *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectS1_
  case 1:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjecti
  case 2:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectRK7QString
  case 3:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectP7QWidget
  default:
    qtrt.ErrorResolve("QSignalMapper", "setMapping", args)
 }

}


func (this *QSignalMapper) mapping(args ...interface{}) () {
  // mapping(int)
  // mapping(const class QString &)
  // mapping(class QObject *)
  // mapping(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper7mappingEi
  case 1:
    // invoke: _ZNK13QSignalMapper7mappingERK7QString
  case 2:
    // invoke: _ZNK13QSignalMapper7mappingEP7QObject
  case 3:
    // invoke: _ZNK13QSignalMapper7mappingEP7QWidget
  default:
    qtrt.ErrorResolve("QSignalMapper", "mapping", args)
 }

}


func (this *QSignalMapper) FreeQSignalMapper(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalMapper", "~QSignalMapper", args)
 }

}

// <= body block end

