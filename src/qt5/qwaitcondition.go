package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qwaitcondition.h
// dst-file: /src/core/qwaitcondition.go
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
// class sizeof(QWaitCondition)=8
type QWaitCondition struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QWaitCondition) wait(args ...interface{}) () {
  // wait(class QReadWriteLock *, unsigned long)
  // wait(class QMutex *, unsigned long)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QReadWriteLock{}) // "QReadWriteLock *"
  vtys[0][1] = qtrt.Int32Ty(false) // "unsigned long"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMutex{}) // "QMutex *"
  vtys[1][1] = qtrt.Int32Ty(false) // "unsigned long"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition4waitEP14QReadWriteLockm
  case 1:
    // invoke: _ZN14QWaitCondition4waitEP6QMutexm
  default:
    qtrt.ErrorResolve("QWaitCondition", "wait", args)
  }

}


func (this *QWaitCondition) wakeAll(args ...interface{}) () {
  // wakeAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeAllEv
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeAll", args)
  }

}


func (this *QWaitCondition) wakeOne(args ...interface{}) () {
  // wakeOne()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QWaitCondition7wakeOneEv
  default:
    qtrt.ErrorResolve("QWaitCondition", "wakeOne", args)
  }

}


func NewQWaitCondition(args ...interface{}) QWaitCondition {
  return QWaitCondition{}
}


func (this *QWaitCondition) FreeQWaitCondition(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWaitCondition", "~QWaitCondition", args)
  }

}

// <= body block end

