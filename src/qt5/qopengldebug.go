package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QOpenGLDebugMessage::swap(QOpenGLDebugMessage & debugMessage);
extern void _ZN19QOpenGLDebugMessage4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QOpenGLDebugMessage::QOpenGLDebugMessage(const QOpenGLDebugMessage & debugMessage);
extern void _ZN19QOpenGLDebugMessageC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QOpenGLDebugMessage::QOpenGLDebugMessage();
extern void _ZN19QOpenGLDebugMessageC2Ev(void* qthis); // 3
  // proto:  void QOpenGLDebugMessage::~QOpenGLDebugMessage();
extern void _ZN19QOpenGLDebugMessageD2Ev(void* qthis); // 4
  // proto:  QOpenGLDebugMessage::Source QOpenGLDebugMessage::source();
extern void _ZNK19QOpenGLDebugMessage6sourceEv(void* qthis); // 4
  // proto:  QString QOpenGLDebugMessage::message();
extern void _ZNK19QOpenGLDebugMessage7messageEv(void* qthis); // 4
  // proto:  QOpenGLDebugMessage::Type QOpenGLDebugMessage::type();
extern void _ZNK19QOpenGLDebugMessage4typeEv(void* qthis); // 4
  // proto:  GLuint QOpenGLDebugMessage::id();
extern void _ZNK19QOpenGLDebugMessage2idEv(void* qthis); // 4
  // proto:  QOpenGLDebugMessage::Severity QOpenGLDebugMessage::severity();
extern void _ZNK19QOpenGLDebugMessage8severityEv(void* qthis); // 4
  // proto:  qint64 QOpenGLDebugLogger::maximumMessageLength();
extern void _ZNK18QOpenGLDebugLogger20maximumMessageLengthEv(void* qthis); // 4
  // proto:  void QOpenGLDebugLogger::~QOpenGLDebugLogger();
extern void _ZN18QOpenGLDebugLoggerD2Ev(void* qthis); // 4
  // proto:  QOpenGLDebugLogger::LoggingMode QOpenGLDebugLogger::loggingMode();
extern void _ZNK18QOpenGLDebugLogger11loggingModeEv(void* qthis); // 4
  // proto:  void QOpenGLDebugLogger::stopLogging();
extern void _ZN18QOpenGLDebugLogger11stopLoggingEv(void* qthis); // 4
  // proto:  QList<QOpenGLDebugMessage> QOpenGLDebugLogger::loggedMessages();
extern void _ZNK18QOpenGLDebugLogger14loggedMessagesEv(void* qthis); // 4
  // proto:  void QOpenGLDebugLogger::popGroup();
extern void _ZN18QOpenGLDebugLogger8popGroupEv(void* qthis); // 4
  // proto:  void QOpenGLDebugLogger::QOpenGLDebugLogger(QObject * parent);
extern void _ZN18QOpenGLDebugLoggerC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QOpenGLDebugLogger::initialize();
extern void _ZN18QOpenGLDebugLogger10initializeEv(void* qthis); // 4
  // proto:  void QOpenGLDebugLogger::logMessage(const QOpenGLDebugMessage & debugMessage);
extern void _ZN18QOpenGLDebugLogger10logMessageERK19QOpenGLDebugMessage(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QOpenGLDebugLogger::metaObject();
extern void _ZNK18QOpenGLDebugLogger10metaObjectEv(void* qthis); // 4
  // proto:  bool QOpenGLDebugLogger::isLogging();
extern void _ZNK18QOpenGLDebugLogger9isLoggingEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLDebugLogger)=1
type QOpenGLDebugLogger struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _messageLogged QOpenGLDebugLogger_messageLogged_signal;
}

// swap(class QOpenGLDebugMessage &)
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
    // invoke: void swap(class QOpenGLDebugMessage &)
    var arg0 = args[0].(QOpenGLDebugMessage).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QOpenGLDebugMessage4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "swap", args)
  }

}

// QOpenGLDebugMessage(const class QOpenGLDebugMessage &)
func NewQOpenGLDebugMessage(args ...interface{}) QOpenGLDebugMessage {
  // QOpenGLDebugMessage(const class QOpenGLDebugMessage &)
  // QOpenGLDebugMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLDebugMessage{}) // "const QOpenGLDebugMessage &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QOpenGLDebugMessageC1ERKS_
    // invoke: void QOpenGLDebugMessage(const class QOpenGLDebugMessage &)
    var arg0 = args[0].(QOpenGLDebugMessage).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QOpenGLDebugMessageC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN19QOpenGLDebugMessageC1Ev
    // invoke: void QOpenGLDebugMessage()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QOpenGLDebugMessageC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "QOpenGLDebugMessage", args)
  }

  return QOpenGLDebugMessage{}
}

// ~QOpenGLDebugMessage()
func (this *QOpenGLDebugMessage) FreeQOpenGLDebugMessage(args ...interface{}) () {
  // ~QOpenGLDebugMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QOpenGLDebugMessageD0Ev
    // invoke: void ~QOpenGLDebugMessage()
    C._ZN19QOpenGLDebugMessageD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "~QOpenGLDebugMessage", args)
  }

}

// source()
func (this *QOpenGLDebugMessage) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLDebugMessage6sourceEv
    // invoke: QOpenGLDebugMessage::Source source()
    C._ZNK19QOpenGLDebugMessage6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "source", args)
  }

}

// message()
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
    // invoke: QString message()
    C._ZNK19QOpenGLDebugMessage7messageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "message", args)
  }

}

// type()
func (this *QOpenGLDebugMessage) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLDebugMessage4typeEv
    // invoke: QOpenGLDebugMessage::Type type()
    C._ZNK19QOpenGLDebugMessage4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "type", args)
  }

}

// id()
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
    // invoke: GLuint id()
    C._ZNK19QOpenGLDebugMessage2idEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "id", args)
  }

}

// severity()
func (this *QOpenGLDebugMessage) severity(args ...interface{}) () {
  // severity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLDebugMessage8severityEv
    // invoke: QOpenGLDebugMessage::Severity severity()
    C._ZNK19QOpenGLDebugMessage8severityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugMessage", "severity", args)
  }

}

// maximumMessageLength()
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
    // invoke: qint64 maximumMessageLength()
    C._ZNK18QOpenGLDebugLogger20maximumMessageLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "maximumMessageLength", args)
  }

}

// ~QOpenGLDebugLogger()
func (this *QOpenGLDebugLogger) FreeQOpenGLDebugLogger(args ...interface{}) () {
  // ~QOpenGLDebugLogger()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLoggerD0Ev
    // invoke: void ~QOpenGLDebugLogger()
    C._ZN18QOpenGLDebugLoggerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "~QOpenGLDebugLogger", args)
  }

}

// loggingMode()
func (this *QOpenGLDebugLogger) loggingMode(args ...interface{}) () {
  // loggingMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLDebugLogger11loggingModeEv
    // invoke: QOpenGLDebugLogger::LoggingMode loggingMode()
    C._ZNK18QOpenGLDebugLogger11loggingModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "loggingMode", args)
  }

}

// stopLogging()
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
    // invoke: void stopLogging()
    C._ZN18QOpenGLDebugLogger11stopLoggingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "stopLogging", args)
  }

}

// loggedMessages()
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
    // invoke: QList<QOpenGLDebugMessage> loggedMessages()
    C._ZNK18QOpenGLDebugLogger14loggedMessagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "loggedMessages", args)
  }

}

// popGroup()
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
    // invoke: void popGroup()
    C._ZN18QOpenGLDebugLogger8popGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "popGroup", args)
  }

}

// QOpenGLDebugLogger(class QObject *)
func NewQOpenGLDebugLogger(args ...interface{}) QOpenGLDebugLogger {
  // QOpenGLDebugLogger(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLDebugLoggerC1EP7QObject
    // invoke: void QOpenGLDebugLogger(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN18QOpenGLDebugLoggerC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "QOpenGLDebugLogger", args)
  }

  return QOpenGLDebugLogger{}
}

// initialize()
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
    // invoke: bool initialize()
    C._ZN18QOpenGLDebugLogger10initializeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "initialize", args)
  }

}

// logMessage(const class QOpenGLDebugMessage &)
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
    // invoke: void logMessage(const class QOpenGLDebugMessage &)
    var arg0 = args[0].(QOpenGLDebugMessage).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QOpenGLDebugLogger10logMessageERK19QOpenGLDebugMessage(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "logMessage", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK18QOpenGLDebugLogger10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "metaObject", args)
  }

}

// isLogging()
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
    // invoke: bool isLogging()
    C._ZNK18QOpenGLDebugLogger9isLoggingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLDebugLogger", "isLogging", args)
  }

}

// <= body block end

