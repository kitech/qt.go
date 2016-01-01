package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qopengldebug.h
// dst-file: /src/gui/qopengldebug.go
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
// class sizeof(QOpenGLDebugMessage)=1
type QOpenGLDebugMessage struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLDebugLogger)=1
type QOpenGLDebugLogger struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _messageLogged QOpenGLDebugLogger_messageLogged_signal;
}


func NewQOpenGLDebugMessage(args ...interface{}) QOpenGLDebugMessage {
  return QOpenGLDebugMessage{}
}


func (this *QOpenGLDebugMessage) FreeQOpenGLDebugMessage(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "~QOpenGLDebugMessage", args)
  }

}


func (this *QOpenGLDebugMessage) id(args ...interface{}) () {
  // id()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLDebugMessage2idEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "id", args)
  }

}


func (this *QOpenGLDebugMessage) message(args ...interface{}) () {
  // message()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLDebugMessage7messageEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "message", args)
  }

}


func (this *QOpenGLDebugMessage) swap(args ...interface{}) () {
  // swap(class QOpenGLDebugMessage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLDebugMessage{}) // "QOpenGLDebugMessage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QOpenGLDebugMessage4swapERS_
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "swap", args)
  }

}


func (this *QOpenGLDebugLogger) FreeQOpenGLDebugLogger(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "~QOpenGLDebugLogger", args)
  }

}


func (this *QOpenGLDebugLogger) maximumMessageLength(args ...interface{}) () {
  // maximumMessageLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLDebugLogger20maximumMessageLengthEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "maximumMessageLength", args)
  }

}


func (this *QOpenGLDebugLogger) isLogging(args ...interface{}) () {
  // isLogging()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLDebugLogger9isLoggingEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "isLogging", args)
  }

}


func (this *QOpenGLDebugLogger) stopLogging(args ...interface{}) () {
  // stopLogging()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLogger11stopLoggingEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "stopLogging", args)
  }

}


func (this *QOpenGLDebugLogger) logMessage(args ...interface{}) () {
  // logMessage(const class QOpenGLDebugMessage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLDebugMessage{}) // "const QOpenGLDebugMessage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLogger10logMessageERK19QOpenGLDebugMessage
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "logMessage", args)
  }

}


func NewQOpenGLDebugLogger(args ...interface{}) QOpenGLDebugLogger {
  return QOpenGLDebugLogger{}
}


func (this *QOpenGLDebugLogger) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLDebugLogger10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "metaObject", args)
  }

}


func (this *QOpenGLDebugLogger) loggedMessages(args ...interface{}) () {
  // loggedMessages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLDebugLogger14loggedMessagesEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "loggedMessages", args)
  }

}


func (this *QOpenGLDebugLogger) popGroup(args ...interface{}) () {
  // popGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLogger8popGroupEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "popGroup", args)
  }

}


func (this *QOpenGLDebugLogger) initialize(args ...interface{}) () {
  // initialize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLogger10initializeEv
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "initialize", args)
  }

}

// <= body block end

