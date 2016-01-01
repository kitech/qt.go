package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qeventtransition.h
// dst-file: /src/core/qeventtransition.go
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
// class sizeof(QEventTransition)=1
type QEventTransition struct {
  /*qbase*/ QAbstractTransition;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQEventTransition(args ...interface{})() {
}


func (this *QEventTransition) FreeQEventTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QEventTransition", "~QEventTransition", args)
 }

}


func (this *QEventTransition) setEventSource(args ...interface{}) () {
  // setEventSource(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QEventTransition14setEventSourceEP7QObject
  default:
    qtrt.ErrorResolve("QEventTransition", "setEventSource", args)
 }

}


func (this *QEventTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition10metaObjectEv
  default:
    qtrt.ErrorResolve("QEventTransition", "metaObject", args)
 }

}


func (this *QEventTransition) eventSource(args ...interface{}) () {
  // eventSource()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QEventTransition11eventSourceEv
  default:
    qtrt.ErrorResolve("QEventTransition", "eventSource", args)
 }

}

// <= body block end

