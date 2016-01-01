package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qcoreevent.h
// dst-file: /src/core/qcoreevent.go
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
// class sizeof(QDeferredDeleteEvent)=24
type QDeferredDeleteEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDynamicPropertyChangeEvent)=32
type QDynamicPropertyChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTimerEvent)=24
type QTimerEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QChildEvent)=32
type QChildEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEvent)=24
type QEvent struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QDeferredDeleteEvent) loopLevel(args ...interface{}) () {
  // loopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QDeferredDeleteEvent9loopLevelEv
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "loopLevel", args)
  }

}


func (this *QDeferredDeleteEvent) FreeQDeferredDeleteEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "~QDeferredDeleteEvent", args)
  }

}


func NewQDeferredDeleteEvent(args ...interface{}) QDeferredDeleteEvent {
  return QDeferredDeleteEvent{}
}


func (this *QDynamicPropertyChangeEvent) FreeQDynamicPropertyChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "~QDynamicPropertyChangeEvent", args)
  }

}


func NewQDynamicPropertyChangeEvent(args ...interface{}) QDynamicPropertyChangeEvent {
  return QDynamicPropertyChangeEvent{}
}


func (this *QDynamicPropertyChangeEvent) propertyName(args ...interface{}) () {
  // propertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QDynamicPropertyChangeEvent12propertyNameEv
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "propertyName", args)
  }

}


func NewQTimerEvent(args ...interface{}) QTimerEvent {
  return QTimerEvent{}
}


func (this *QTimerEvent) FreeQTimerEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimerEvent", "~QTimerEvent", args)
  }

}


func (this *QTimerEvent) timerId(args ...interface{}) () {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTimerEvent7timerIdEv
  default:
    qtrt.ErrorResolve("QTimerEvent", "timerId", args)
  }

}


func (this *QChildEvent) added(args ...interface{}) () {
  // added()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5addedEv
  default:
    qtrt.ErrorResolve("QChildEvent", "added", args)
  }

}


func (this *QChildEvent) polished(args ...interface{}) () {
  // polished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent8polishedEv
  default:
    qtrt.ErrorResolve("QChildEvent", "polished", args)
  }

}


func (this *QChildEvent) FreeQChildEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChildEvent", "~QChildEvent", args)
  }

}


func (this *QChildEvent) removed(args ...interface{}) () {
  // removed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent7removedEv
  default:
    qtrt.ErrorResolve("QChildEvent", "removed", args)
  }

}


func (this *QChildEvent) child(args ...interface{}) () {
  // child()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5childEv
  default:
    qtrt.ErrorResolve("QChildEvent", "child", args)
  }

}


func (this *QEvent) setAccepted(args ...interface{}) () {
  // setAccepted(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent11setAcceptedEb
  default:
    qtrt.ErrorResolve("QEvent", "setAccepted", args)
  }

}


func (this *QEvent) ignore(args ...interface{}) () {
  // ignore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6ignoreEv
  default:
    qtrt.ErrorResolve("QEvent", "ignore", args)
  }

}


func (this *QEvent) isAccepted(args ...interface{}) () {
  // isAccepted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent10isAcceptedEv
  default:
    qtrt.ErrorResolve("QEvent", "isAccepted", args)
  }

}


func (this *QEvent) FreeQEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEvent", "~QEvent", args)
  }

}


func NewQEvent(args ...interface{}) QEvent {
  return QEvent{}
}


func (this *QEvent) accept(args ...interface{}) () {
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6acceptEv
  default:
    qtrt.ErrorResolve("QEvent", "accept", args)
  }

}


func (this *QEvent) registerEventType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEvent", "registerEventType", args)
  }

}


func (this *QEvent) spontaneous(args ...interface{}) () {
  // spontaneous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent11spontaneousEv
  default:
    qtrt.ErrorResolve("QEvent", "spontaneous", args)
  }

}

// <= body block end

