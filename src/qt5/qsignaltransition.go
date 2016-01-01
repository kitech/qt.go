package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qsignaltransition.h
// dst-file: /src/core/qsignaltransition.go
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
// class sizeof(QSignalTransition)=1
type QSignalTransition struct {
  /*qbase*/ QAbstractTransition;
  qclsinst uint64 /* *mut c_void*/;
//  _senderObjectChanged QSignalTransition_senderObjectChanged_signal;
//  _signalChanged QSignalTransition_signalChanged_signal;
}


func (this *QSignalTransition) setSenderObject(args ...interface{}) () {
  // setSenderObject(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition15setSenderObjectEPK7QObject
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSenderObject", args)
  }

}


func (this *QSignalTransition) signal(args ...interface{}) () {
  // signal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition6signalEv
  default:
    qtrt.ErrorResolve("QSignalTransition", "signal", args)
  }

}


func (this *QSignalTransition) FreeQSignalTransition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalTransition", "~QSignalTransition", args)
  }

}


func NewQSignalTransition(args ...interface{}) QSignalTransition {
  return QSignalTransition{}
}


func (this *QSignalTransition) senderObject(args ...interface{}) () {
  // senderObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition12senderObjectEv
  default:
    qtrt.ErrorResolve("QSignalTransition", "senderObject", args)
  }

}


func (this *QSignalTransition) setSignal(args ...interface{}) () {
  // setSignal(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition9setSignalERK10QByteArray
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSignal", args)
  }

}


func (this *QSignalTransition) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition10metaObjectEv
  default:
    qtrt.ErrorResolve("QSignalTransition", "metaObject", args)
  }

}

// <= body block end

