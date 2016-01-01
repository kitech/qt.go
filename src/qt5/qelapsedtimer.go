package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qelapsedtimer.h
// dst-file: /src/core/qelapsedtimer.go
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
// class sizeof(QElapsedTimer)=16
type QElapsedTimer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QElapsedTimer) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer5startEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "start", args)
  }

}


func (this *QElapsedTimer) nsecsElapsed(args ...interface{}) () {
  // nsecsElapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer12nsecsElapsedEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "nsecsElapsed", args)
  }

}


func (this *QElapsedTimer) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer10invalidateEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "invalidate", args)
  }

}


func (this *QElapsedTimer) isMonotonic_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isMonotonic", args)
  }

}


func NewQElapsedTimer(args ...interface{}) QElapsedTimer {
  return QElapsedTimer{}
}


func (this *QElapsedTimer) msecsTo(args ...interface{}) () {
  // msecsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7msecsToERKS_
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsTo", args)
  }

}


func (this *QElapsedTimer) msecsSinceReference(args ...interface{}) () {
  // msecsSinceReference()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer19msecsSinceReferenceEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "msecsSinceReference", args)
  }

}


func (this *QElapsedTimer) hasExpired(args ...interface{}) () {
  // hasExpired(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer10hasExpiredEx
  default:
    qtrt.ErrorResolve("QElapsedTimer", "hasExpired", args)
  }

}


func (this *QElapsedTimer) restart(args ...interface{}) () {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QElapsedTimer7restartEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "restart", args)
  }

}


func (this *QElapsedTimer) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7isValidEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "isValid", args)
  }

}


func (this *QElapsedTimer) secsTo(args ...interface{}) () {
  // secsTo(const class QElapsedTimer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QElapsedTimer{}) // "const QElapsedTimer &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer6secsToERKS_
  default:
    qtrt.ErrorResolve("QElapsedTimer", "secsTo", args)
  }

}


func (this *QElapsedTimer) elapsed(args ...interface{}) () {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QElapsedTimer7elapsedEv
  default:
    qtrt.ErrorResolve("QElapsedTimer", "elapsed", args)
  }

}

// <= body block end

