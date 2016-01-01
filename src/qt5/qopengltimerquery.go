package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qopengltimerquery.h
// dst-file: /src/gui/qopengltimerquery.go
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
// class sizeof(QOpenGLTimerQuery)=1
type QOpenGLTimerQuery struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLTimeMonitor)=1
type QOpenGLTimeMonitor struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQOpenGLTimerQuery(args ...interface{}) QOpenGLTimerQuery {
  return QOpenGLTimerQuery{}
}


func (this *QOpenGLTimerQuery) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery6createEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "create", args)
  }

}


func (this *QOpenGLTimerQuery) isCreated(args ...interface{}) () {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery9isCreatedEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isCreated", args)
  }

}


func (this *QOpenGLTimerQuery) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery3endEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "end", args)
  }

}


func (this *QOpenGLTimerQuery) FreeQOpenGLTimerQuery(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "~QOpenGLTimerQuery", args)
  }

}


func (this *QOpenGLTimerQuery) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery5beginEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "begin", args)
  }

}


func (this *QOpenGLTimerQuery) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery7destroyEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "destroy", args)
  }

}


func (this *QOpenGLTimerQuery) waitForResult(args ...interface{}) () {
  // waitForResult()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery13waitForResultEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForResult", args)
  }

}


func (this *QOpenGLTimerQuery) objectId(args ...interface{}) () {
  // objectId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery8objectIdEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "objectId", args)
  }

}


func (this *QOpenGLTimerQuery) waitForTimestamp(args ...interface{}) () {
  // waitForTimestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery16waitForTimestampEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForTimestamp", args)
  }

}


func (this *QOpenGLTimerQuery) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "metaObject", args)
  }

}


func (this *QOpenGLTimerQuery) recordTimestamp(args ...interface{}) () {
  // recordTimestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery15recordTimestampEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "recordTimestamp", args)
  }

}


func (this *QOpenGLTimerQuery) isResultAvailable(args ...interface{}) () {
  // isResultAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery17isResultAvailableEv
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isResultAvailable", args)
  }

}


func (this *QOpenGLTimeMonitor) setSampleCount(args ...interface{}) () {
  // setSampleCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor14setSampleCountEi
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "setSampleCount", args)
  }

}


func (this *QOpenGLTimeMonitor) sampleCount(args ...interface{}) () {
  // sampleCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor11sampleCountEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "sampleCount", args)
  }

}


func (this *QOpenGLTimeMonitor) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor7destroyEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "destroy", args)
  }

}


func NewQOpenGLTimeMonitor(args ...interface{}) QOpenGLTimeMonitor {
  return QOpenGLTimeMonitor{}
}


func (this *QOpenGLTimeMonitor) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor6createEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "create", args)
  }

}


func (this *QOpenGLTimeMonitor) FreeQOpenGLTimeMonitor(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "~QOpenGLTimeMonitor", args)
  }

}


func (this *QOpenGLTimeMonitor) isResultAvailable(args ...interface{}) () {
  // isResultAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor17isResultAvailableEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isResultAvailable", args)
  }

}


func (this *QOpenGLTimeMonitor) waitForIntervals(args ...interface{}) () {
  // waitForIntervals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor16waitForIntervalsEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForIntervals", args)
  }

}


func (this *QOpenGLTimeMonitor) objectIds(args ...interface{}) () {
  // objectIds()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor9objectIdsEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "objectIds", args)
  }

}


func (this *QOpenGLTimeMonitor) recordSample(args ...interface{}) () {
  // recordSample()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor12recordSampleEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "recordSample", args)
  }

}


func (this *QOpenGLTimeMonitor) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor5resetEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "reset", args)
  }

}


func (this *QOpenGLTimeMonitor) waitForSamples(args ...interface{}) () {
  // waitForSamples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor14waitForSamplesEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForSamples", args)
  }

}


func (this *QOpenGLTimeMonitor) isCreated(args ...interface{}) () {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor9isCreatedEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isCreated", args)
  }

}


func (this *QOpenGLTimeMonitor) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "metaObject", args)
  }

}

// <= body block end

