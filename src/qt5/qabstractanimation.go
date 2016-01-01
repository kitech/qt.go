package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qabstractanimation.h
// dst-file: /src/core/qabstractanimation.go
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
// class sizeof(QAbstractAnimation)=1
type QAbstractAnimation struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _currentLoopChanged QAbstractAnimation_currentLoopChanged_signal;
//  _finished QAbstractAnimation_finished_signal;
//  _stateChanged QAbstractAnimation_stateChanged_signal;
//  _directionChanged QAbstractAnimation_directionChanged_signal;
}

// class sizeof(QAnimationDriver)=1
type QAnimationDriver struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _started QAnimationDriver_started_signal;
//  _stopped QAnimationDriver_stopped_signal;
}


func (this *QAbstractAnimation) resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation6resumeEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "resume", args)
  }

}


func NewQAbstractAnimation(args ...interface{}) QAbstractAnimation {
  return QAbstractAnimation{}
}


func (this *QAbstractAnimation) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation4stopEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "stop", args)
  }

}


func (this *QAbstractAnimation) pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation5pauseEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "pause", args)
  }

}


func (this *QAbstractAnimation) setLoopCount(args ...interface{}) () {
  // setLoopCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation12setLoopCountEi
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setLoopCount", args)
  }

}


func (this *QAbstractAnimation) currentLoop(args ...interface{}) () {
  // currentLoop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation11currentLoopEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoop", args)
  }

}


func (this *QAbstractAnimation) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation5groupEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "group", args)
  }

}


func (this *QAbstractAnimation) setPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation9setPausedEb
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setPaused", args)
  }

}


func (this *QAbstractAnimation) totalDuration(args ...interface{}) () {
  // totalDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation13totalDurationEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "totalDuration", args)
  }

}


func (this *QAbstractAnimation) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation8durationEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "duration", args)
  }

}


func (this *QAbstractAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "metaObject", args)
  }

}


func (this *QAbstractAnimation) currentLoopTime(args ...interface{}) () {
  // currentLoopTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation15currentLoopTimeEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentLoopTime", args)
  }

}


func (this *QAbstractAnimation) currentTime(args ...interface{}) () {
  // currentTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation11currentTimeEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "currentTime", args)
  }

}


func (this *QAbstractAnimation) setCurrentTime(args ...interface{}) () {
  // setCurrentTime(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractAnimation14setCurrentTimeEi
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "setCurrentTime", args)
  }

}


func (this *QAbstractAnimation) FreeQAbstractAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "~QAbstractAnimation", args)
  }

}


func (this *QAbstractAnimation) loopCount(args ...interface{}) () {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractAnimation9loopCountEv
  default:
    qtrt.ErrorResolve("QAbstractAnimation", "loopCount", args)
  }

}


func (this *QAnimationDriver) advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver7advanceEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "advance", args)
  }

}


func (this *QAnimationDriver) FreeQAnimationDriver(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAnimationDriver", "~QAnimationDriver", args)
  }

}


func NewQAnimationDriver(args ...interface{}) QAnimationDriver {
  return QAnimationDriver{}
}


func (this *QAnimationDriver) elapsed(args ...interface{}) () {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver7elapsedEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "elapsed", args)
  }

}


func (this *QAnimationDriver) install(args ...interface{}) () {
  // install()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver7installEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "install", args)
  }

}


func (this *QAnimationDriver) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver10metaObjectEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "metaObject", args)
  }

}


func (this *QAnimationDriver) uninstall(args ...interface{}) () {
  // uninstall()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver9uninstallEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "uninstall", args)
  }

}


func (this *QAnimationDriver) isRunning(args ...interface{}) () {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver9isRunningEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "isRunning", args)
  }

}


func (this *QAnimationDriver) startTime(args ...interface{}) () {
  // startTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAnimationDriver9startTimeEv
  default:
    qtrt.ErrorResolve("QAnimationDriver", "startTime", args)
  }

}


func (this *QAnimationDriver) setStartTime(args ...interface{}) () {
  // setStartTime(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAnimationDriver12setStartTimeEx
  default:
    qtrt.ErrorResolve("QAnimationDriver", "setStartTime", args)
  }

}

// <= body block end

