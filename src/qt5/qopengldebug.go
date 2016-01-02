package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qopengldebug.h
// dst-file: /src/gui/qopengldebug.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QOpenGLDebugMessage::QOpenGLDebugMessage();
extern void* dector_ZN19QOpenGLDebugMessageC1Ev();
extern void _ZN19QOpenGLDebugMessageC1Ev(void* qthis);
  // proto:  void QOpenGLDebugMessage::~QOpenGLDebugMessage();
extern void _ZN19QOpenGLDebugMessageD0Ev(void* qthis);
  // proto:  GLuint QOpenGLDebugMessage::id();
extern void _ZNK19QOpenGLDebugMessage2idEv(void* qthis);
  // proto:  void QOpenGLDebugMessage::QOpenGLDebugMessage(const QOpenGLDebugMessage & debugMessage);
extern void* dector_ZN19QOpenGLDebugMessageC1ERKS_(void* arg0);
extern void _ZN19QOpenGLDebugMessageC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QOpenGLDebugMessage::message();
extern void _ZNK19QOpenGLDebugMessage7messageEv(void* qthis);
  // proto:  void QOpenGLDebugMessage::swap(QOpenGLDebugMessage & debugMessage);
extern void demth_ZN19QOpenGLDebugMessage4swapERS_(void* qthis, void* arg0);
  // proto:  void QOpenGLDebugLogger::~QOpenGLDebugLogger();
extern void _ZN18QOpenGLDebugLoggerD0Ev(void* qthis);
  // proto:  qint64 QOpenGLDebugLogger::maximumMessageLength();
extern void _ZNK18QOpenGLDebugLogger20maximumMessageLengthEv(void* qthis);
  // proto:  bool QOpenGLDebugLogger::isLogging();
extern void _ZNK18QOpenGLDebugLogger9isLoggingEv(void* qthis);
  // proto:  void QOpenGLDebugLogger::stopLogging();
extern void _ZN18QOpenGLDebugLogger11stopLoggingEv(void* qthis);
  // proto:  void QOpenGLDebugLogger::logMessage(const QOpenGLDebugMessage & debugMessage);
extern void _ZN18QOpenGLDebugLogger10logMessageERK19QOpenGLDebugMessage(void* qthis, void* arg0);
  // proto:  void QOpenGLDebugLogger::QOpenGLDebugLogger(const QOpenGLDebugLogger & );
extern void* dector_ZN18QOpenGLDebugLoggerC1ERKS_(void* arg0);
extern void _ZN18QOpenGLDebugLoggerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOpenGLDebugLogger::QOpenGLDebugLogger(QObject * parent);
extern void* dector_ZN18QOpenGLDebugLoggerC1EP7QObject(void* arg0);
extern void _ZN18QOpenGLDebugLoggerC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QOpenGLDebugLogger::metaObject();
extern void _ZNK18QOpenGLDebugLogger10metaObjectEv(void* qthis);
  // proto:  QList<QOpenGLDebugMessage> QOpenGLDebugLogger::loggedMessages();
extern void _ZNK18QOpenGLDebugLogger14loggedMessagesEv(void* qthis);
  // proto:  void QOpenGLDebugLogger::popGroup();
extern void _ZN18QOpenGLDebugLogger8popGroupEv(void* qthis);
  // proto:  bool QOpenGLDebugLogger::initialize();
extern void _ZN18QOpenGLDebugLogger10initializeEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

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

  // proto:  void QOpenGLDebugMessage::QOpenGLDebugMessage();
func NewQOpenGLDebugMessage(args ...interface{}) QOpenGLDebugMessage {
  return QOpenGLDebugMessage{}
}

  // proto:  void QOpenGLDebugMessage::~QOpenGLDebugMessage();
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

  // proto:  GLuint QOpenGLDebugMessage::id();
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

  // proto:  QString QOpenGLDebugMessage::message();
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

  // proto:  void QOpenGLDebugMessage::swap(QOpenGLDebugMessage & debugMessage);
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
    var arg0 = args[0].(QOpenGLDebugMessage).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "swap", args)
  }

}

  // proto:  void QOpenGLDebugLogger::~QOpenGLDebugLogger();
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

  // proto:  qint64 QOpenGLDebugLogger::maximumMessageLength();
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

  // proto:  bool QOpenGLDebugLogger::isLogging();
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

  // proto:  void QOpenGLDebugLogger::stopLogging();
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

  // proto:  void QOpenGLDebugLogger::logMessage(const QOpenGLDebugMessage & debugMessage);
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
    var arg0 = args[0].(QOpenGLDebugMessage).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "logMessage", args)
  }

}

  // proto:  void QOpenGLDebugLogger::QOpenGLDebugLogger(const QOpenGLDebugLogger & );
func NewQOpenGLDebugLogger(args ...interface{}) QOpenGLDebugLogger {
  return QOpenGLDebugLogger{}
}

  // proto:  const QMetaObject * QOpenGLDebugLogger::metaObject();
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

  // proto:  QList<QOpenGLDebugMessage> QOpenGLDebugLogger::loggedMessages();
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

  // proto:  void QOpenGLDebugLogger::popGroup();
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

  // proto:  bool QOpenGLDebugLogger::initialize();
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

