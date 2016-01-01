package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qobjectcleanuphandler.h
// dst-file: /src/core/qobjectcleanuphandler.go
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
// class sizeof(QObjectCleanupHandler)=1
type QObjectCleanupHandler struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QObjectCleanupHandler) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler5clearEv
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "clear", args)
 }

}


func (this *QObjectCleanupHandler) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QObjectCleanupHandler7isEmptyEv
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "isEmpty", args)
 }

}


func (this *QObjectCleanupHandler) FreeQObjectCleanupHandler(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "~QObjectCleanupHandler", args)
 }

}


func (this *QObjectCleanupHandler) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QObjectCleanupHandler10metaObjectEv
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "metaObject", args)
 }

}


func (this *QObjectCleanupHandler) remove(args ...interface{}) () {
  // remove(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler6removeEP7QObject
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "remove", args)
 }

}


func (this *QObjectCleanupHandler) add(args ...interface{}) () {
  // add(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QObjectCleanupHandler3addEP7QObject
  default:
    qtrt.ErrorResolve("QObjectCleanupHandler", "add", args)
 }

}


func NewQObjectCleanupHandler(args ...interface{})() {
}

// <= body block end

