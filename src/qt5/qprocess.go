package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qprocess.h
// dst-file: /src/core/qprocess.go
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
  // proto:  void QProcess::close();
extern void _ZN8QProcess5closeEv(void* qthis);
  // proto:  void QProcess::setEnvironment(const QStringList & environment);
extern void _ZN8QProcess14setEnvironmentERK11QStringList(void* qthis, void* arg0);
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments);
extern void _ZN8QProcess13startDetachedERK7QStringRK11QStringList(void* arg0, void* arg1);
  // proto: static bool QProcess::startDetached(const QString & command);
extern void _ZN8QProcess13startDetachedERK7QString(void* arg0);
  // proto:  bool QProcess::atEnd();
extern void _ZNK8QProcess5atEndEv(void* qthis);
  // proto: static QStringList QProcess::systemEnvironment();
extern void _ZN8QProcess17systemEnvironmentEv();
  // proto:  void QProcess::setProcessEnvironment(const QProcessEnvironment & environment);
extern void _ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment(void* qthis, void* arg0);
  // proto:  void QProcess::QProcess(const QProcess & );
extern void* dector_ZN8QProcessC1ERKS_(void* arg0);
extern void _ZN8QProcessC1ERKS_(void* qthis, void* arg0);
  // proto:  Q_PID QProcess::pid();
extern void _ZNK8QProcess3pidEv(void* qthis);
  // proto:  void QProcess::setArguments(const QStringList & arguments);
extern void _ZN8QProcess12setArgumentsERK11QStringList(void* qthis, void* arg0);
  // proto:  void QProcess::~QProcess();
extern void _ZN8QProcessD0Ev(void* qthis);
  // proto: static int QProcess::execute(const QString & command);
extern void _ZN8QProcess7executeERK7QString(void* arg0);
  // proto:  void QProcess::closeWriteChannel();
extern void _ZN8QProcess17closeWriteChannelEv(void* qthis);
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments, const QString & workingDirectory, qint64 * pid);
extern void _ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px(void* arg0, void* arg1, void* arg2, long long* arg3);
  // proto:  QProcessEnvironment QProcess::processEnvironment();
extern void _ZNK8QProcess18processEnvironmentEv(void* qthis);
  // proto:  QByteArray QProcess::readAllStandardOutput();
extern void _ZN8QProcess21readAllStandardOutputEv(void* qthis);
  // proto: static QString QProcess::nullDevice();
extern void _ZN8QProcess10nullDeviceEv();
  // proto: static int QProcess::execute(const QString & program, const QStringList & arguments);
extern void _ZN8QProcess7executeERK7QStringRK11QStringList(void* arg0, void* arg1);
  // proto:  bool QProcess::waitForBytesWritten(int msecs);
extern void _ZN8QProcess19waitForBytesWrittenEi(void* qthis, int arg0);
  // proto:  void QProcess::QProcess(QObject * parent);
extern void* dector_ZN8QProcessC1EP7QObject(void* arg0);
extern void _ZN8QProcessC1EP7QObject(void* qthis, void* arg0);
  // proto:  QString QProcess::program();
extern void _ZNK8QProcess7programEv(void* qthis);
  // proto:  qint64 QProcess::processId();
extern void _ZNK8QProcess9processIdEv(void* qthis);
  // proto:  QStringList QProcess::arguments();
extern void _ZNK8QProcess9argumentsEv(void* qthis);
  // proto:  bool QProcess::isSequential();
extern void _ZNK8QProcess12isSequentialEv(void* qthis);
  // proto:  bool QProcess::waitForReadyRead(int msecs);
extern void _ZN8QProcess16waitForReadyReadEi(void* qthis, int arg0);
  // proto:  void QProcess::setWorkingDirectory(const QString & dir);
extern void _ZN8QProcess19setWorkingDirectoryERK7QString(void* qthis, void* arg0);
  // proto:  void QProcess::terminate();
extern void _ZN8QProcess9terminateEv(void* qthis);
  // proto:  void QProcess::kill();
extern void _ZN8QProcess4killEv(void* qthis);
  // proto:  qint64 QProcess::bytesAvailable();
extern void _ZNK8QProcess14bytesAvailableEv(void* qthis);
  // proto:  const QMetaObject * QProcess::metaObject();
extern void _ZNK8QProcess10metaObjectEv(void* qthis);
  // proto:  bool QProcess::waitForStarted(int msecs);
extern void _ZN8QProcess14waitForStartedEi(void* qthis, int arg0);
  // proto:  QByteArray QProcess::readAllStandardError();
extern void _ZN8QProcess20readAllStandardErrorEv(void* qthis);
  // proto:  int QProcess::exitCode();
extern void _ZNK8QProcess8exitCodeEv(void* qthis);
  // proto:  QStringList QProcess::environment();
extern void _ZNK8QProcess11environmentEv(void* qthis);
  // proto:  bool QProcess::canReadLine();
extern void _ZNK8QProcess11canReadLineEv(void* qthis);
  // proto:  void QProcess::setStandardOutputProcess(QProcess * destination);
extern void _ZN8QProcess24setStandardOutputProcessEPS_(void* qthis, void* arg0);
  // proto:  bool QProcess::waitForFinished(int msecs);
extern void _ZN8QProcess15waitForFinishedEi(void* qthis, int arg0);
  // proto:  qint64 QProcess::bytesToWrite();
extern void _ZNK8QProcess12bytesToWriteEv(void* qthis);
  // proto:  QString QProcess::workingDirectory();
extern void _ZNK8QProcess16workingDirectoryEv(void* qthis);
  // proto:  void QProcess::setProgram(const QString & program);
extern void _ZN8QProcess10setProgramERK7QString(void* qthis, void* arg0);
  // proto:  void QProcess::setStandardInputFile(const QString & fileName);
extern void _ZN8QProcess20setStandardInputFileERK7QString(void* qthis, void* arg0);
  // proto:  bool QProcessEnvironment::contains(const QString & name);
extern void _ZNK19QProcessEnvironment8containsERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QProcessEnvironment::keys();
extern void _ZNK19QProcessEnvironment4keysEv(void* qthis);
  // proto:  void QProcessEnvironment::remove(const QString & name);
extern void _ZN19QProcessEnvironment6removeERK7QString(void* qthis, void* arg0);
  // proto:  void QProcessEnvironment::clear();
extern void _ZN19QProcessEnvironment5clearEv(void* qthis);
  // proto:  QString QProcessEnvironment::value(const QString & name, const QString & defaultValue);
extern void _ZNK19QProcessEnvironment5valueERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QProcessEnvironment::isEmpty();
extern void _ZNK19QProcessEnvironment7isEmptyEv(void* qthis);
  // proto:  void QProcessEnvironment::~QProcessEnvironment();
extern void _ZN19QProcessEnvironmentD0Ev(void* qthis);
  // proto:  void QProcessEnvironment::swap(QProcessEnvironment & other);
extern void demth_ZN19QProcessEnvironment4swapERS_(void* qthis, void* arg0);
  // proto:  void QProcessEnvironment::QProcessEnvironment(const QProcessEnvironment & other);
extern void* dector_ZN19QProcessEnvironmentC1ERKS_(void* arg0);
extern void _ZN19QProcessEnvironmentC1ERKS_(void* qthis, void* arg0);
  // proto: static QProcessEnvironment QProcessEnvironment::systemEnvironment();
extern void _ZN19QProcessEnvironment17systemEnvironmentEv();
  // proto:  void QProcessEnvironment::insert(const QString & name, const QString & value);
extern void _ZN19QProcessEnvironment6insertERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QStringList QProcessEnvironment::toStringList();
extern void _ZNK19QProcessEnvironment12toStringListEv(void* qthis);
  // proto:  void QProcessEnvironment::QProcessEnvironment();
extern void* dector_ZN19QProcessEnvironmentC1Ev();
extern void _ZN19QProcessEnvironmentC1Ev(void* qthis);
  // proto:  void QProcessEnvironment::insert(const QProcessEnvironment & e);
extern void _ZN19QProcessEnvironment6insertERKS_(void* qthis, void* arg0);
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

// class sizeof(QProcess)=1
type QProcess struct {
  /*qbase*/ QIODevice;
  qclsinst uint64 /* *mut c_void*/;
//  _stateChanged QProcess_stateChanged_signal;
//  _started QProcess_started_signal;
//  _finished QProcess_finished_signal;
//  _readyReadStandardError QProcess_readyReadStandardError_signal;
//  _error QProcess_error_signal;
//  _readyReadStandardOutput QProcess_readyReadStandardOutput_signal;
}

// class sizeof(QProcessEnvironment)=1
type QProcessEnvironment struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QProcess::close();
func (this *QProcess) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess5closeEv
    // invoke: void close()
    C._ZN8QProcess5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "close", args)
  }

}

  // proto:  void QProcess::setEnvironment(const QStringList & environment);
func (this *QProcess) setEnvironment(args ...interface{}) () {
  // setEnvironment(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess14setEnvironmentERK11QStringList
    // invoke: void setEnvironment(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess14setEnvironmentERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setEnvironment", args)
  }

}

  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments);
func (this *QProcess) startDetached_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcess", "startDetached", args)
  }

}

  // proto:  bool QProcess::atEnd();
func (this *QProcess) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess5atEndEv
    // invoke: bool atEnd()
    C._ZNK8QProcess5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "atEnd", args)
  }

}

  // proto: static QStringList QProcess::systemEnvironment();
func (this *QProcess) systemEnvironment_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcess", "systemEnvironment", args)
  }

}

  // proto:  void QProcess::setProcessEnvironment(const QProcessEnvironment & environment);
func (this *QProcess) setProcessEnvironment(args ...interface{}) () {
  // setProcessEnvironment(const class QProcessEnvironment &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProcessEnvironment{}) // "const QProcessEnvironment &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment
    // invoke: void setProcessEnvironment(const class QProcessEnvironment &)
    var arg0 = args[0].(QProcessEnvironment).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setProcessEnvironment", args)
  }

}

  // proto:  void QProcess::QProcess(const QProcess & );
func NewQProcess(args ...interface{}) QProcess {
  return QProcess{}
}

  // proto:  Q_PID QProcess::pid();
func (this *QProcess) pid(args ...interface{}) () {
  // pid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess3pidEv
    // invoke: Q_PID pid()
    C._ZNK8QProcess3pidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "pid", args)
  }

}

  // proto:  void QProcess::setArguments(const QStringList & arguments);
func (this *QProcess) setArguments(args ...interface{}) () {
  // setArguments(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess12setArgumentsERK11QStringList
    // invoke: void setArguments(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess12setArgumentsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setArguments", args)
  }

}

  // proto:  void QProcess::~QProcess();
func (this *QProcess) FreeQProcess(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcess", "~QProcess", args)
  }

}

  // proto: static int QProcess::execute(const QString & command);
func (this *QProcess) execute_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcess", "execute", args)
  }

}

  // proto:  void QProcess::closeWriteChannel();
func (this *QProcess) closeWriteChannel(args ...interface{}) () {
  // closeWriteChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess17closeWriteChannelEv
    // invoke: void closeWriteChannel()
    C._ZN8QProcess17closeWriteChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "closeWriteChannel", args)
  }

}

  // proto:  QProcessEnvironment QProcess::processEnvironment();
func (this *QProcess) processEnvironment(args ...interface{}) () {
  // processEnvironment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess18processEnvironmentEv
    // invoke: QProcessEnvironment processEnvironment()
    C._ZNK8QProcess18processEnvironmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "processEnvironment", args)
  }

}

  // proto:  QByteArray QProcess::readAllStandardOutput();
func (this *QProcess) readAllStandardOutput(args ...interface{}) () {
  // readAllStandardOutput()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess21readAllStandardOutputEv
    // invoke: QByteArray readAllStandardOutput()
    C._ZN8QProcess21readAllStandardOutputEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardOutput", args)
  }

}

  // proto: static QString QProcess::nullDevice();
func (this *QProcess) nullDevice_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcess", "nullDevice", args)
  }

}

  // proto:  bool QProcess::waitForBytesWritten(int msecs);
func (this *QProcess) waitForBytesWritten(args ...interface{}) () {
  // waitForBytesWritten(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess19waitForBytesWrittenEi
    // invoke: bool waitForBytesWritten(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QProcess19waitForBytesWrittenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "waitForBytesWritten", args)
  }

}

  // proto:  QString QProcess::program();
func (this *QProcess) program(args ...interface{}) () {
  // program()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess7programEv
    // invoke: QString program()
    C._ZNK8QProcess7programEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "program", args)
  }

}

  // proto:  qint64 QProcess::processId();
func (this *QProcess) processId(args ...interface{}) () {
  // processId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess9processIdEv
    // invoke: qint64 processId()
    C._ZNK8QProcess9processIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "processId", args)
  }

}

  // proto:  QStringList QProcess::arguments();
func (this *QProcess) arguments(args ...interface{}) () {
  // arguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess9argumentsEv
    // invoke: QStringList arguments()
    C._ZNK8QProcess9argumentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "arguments", args)
  }

}

  // proto:  bool QProcess::isSequential();
func (this *QProcess) isSequential(args ...interface{}) () {
  // isSequential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess12isSequentialEv
    // invoke: bool isSequential()
    C._ZNK8QProcess12isSequentialEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "isSequential", args)
  }

}

  // proto:  bool QProcess::waitForReadyRead(int msecs);
func (this *QProcess) waitForReadyRead(args ...interface{}) () {
  // waitForReadyRead(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess16waitForReadyReadEi
    // invoke: bool waitForReadyRead(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QProcess16waitForReadyReadEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "waitForReadyRead", args)
  }

}

  // proto:  void QProcess::setWorkingDirectory(const QString & dir);
func (this *QProcess) setWorkingDirectory(args ...interface{}) () {
  // setWorkingDirectory(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess19setWorkingDirectoryERK7QString
    // invoke: void setWorkingDirectory(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess19setWorkingDirectoryERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setWorkingDirectory", args)
  }

}

  // proto:  void QProcess::terminate();
func (this *QProcess) terminate(args ...interface{}) () {
  // terminate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess9terminateEv
    // invoke: void terminate()
    C._ZN8QProcess9terminateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "terminate", args)
  }

}

  // proto:  void QProcess::kill();
func (this *QProcess) kill(args ...interface{}) () {
  // kill()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess4killEv
    // invoke: void kill()
    C._ZN8QProcess4killEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "kill", args)
  }

}

  // proto:  qint64 QProcess::bytesAvailable();
func (this *QProcess) bytesAvailable(args ...interface{}) () {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess14bytesAvailableEv
    // invoke: qint64 bytesAvailable()
    C._ZNK8QProcess14bytesAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "bytesAvailable", args)
  }

}

  // proto:  const QMetaObject * QProcess::metaObject();
func (this *QProcess) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK8QProcess10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "metaObject", args)
  }

}

  // proto:  bool QProcess::waitForStarted(int msecs);
func (this *QProcess) waitForStarted(args ...interface{}) () {
  // waitForStarted(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess14waitForStartedEi
    // invoke: bool waitForStarted(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QProcess14waitForStartedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "waitForStarted", args)
  }

}

  // proto:  QByteArray QProcess::readAllStandardError();
func (this *QProcess) readAllStandardError(args ...interface{}) () {
  // readAllStandardError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess20readAllStandardErrorEv
    // invoke: QByteArray readAllStandardError()
    C._ZN8QProcess20readAllStandardErrorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardError", args)
  }

}

  // proto:  int QProcess::exitCode();
func (this *QProcess) exitCode(args ...interface{}) () {
  // exitCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess8exitCodeEv
    // invoke: int exitCode()
    C._ZNK8QProcess8exitCodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "exitCode", args)
  }

}

  // proto:  QStringList QProcess::environment();
func (this *QProcess) environment(args ...interface{}) () {
  // environment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess11environmentEv
    // invoke: QStringList environment()
    C._ZNK8QProcess11environmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "environment", args)
  }

}

  // proto:  bool QProcess::canReadLine();
func (this *QProcess) canReadLine(args ...interface{}) () {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess11canReadLineEv
    // invoke: bool canReadLine()
    C._ZNK8QProcess11canReadLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "canReadLine", args)
  }

}

  // proto:  void QProcess::setStandardOutputProcess(QProcess * destination);
func (this *QProcess) setStandardOutputProcess(args ...interface{}) () {
  // setStandardOutputProcess(class QProcess *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProcess{}) // "QProcess *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess24setStandardOutputProcessEPS_
    // invoke: void setStandardOutputProcess(class QProcess *)
    var arg0 = args[0].(QProcess).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess24setStandardOutputProcessEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setStandardOutputProcess", args)
  }

}

  // proto:  bool QProcess::waitForFinished(int msecs);
func (this *QProcess) waitForFinished(args ...interface{}) () {
  // waitForFinished(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess15waitForFinishedEi
    // invoke: bool waitForFinished(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QProcess15waitForFinishedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "waitForFinished", args)
  }

}

  // proto:  qint64 QProcess::bytesToWrite();
func (this *QProcess) bytesToWrite(args ...interface{}) () {
  // bytesToWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess12bytesToWriteEv
    // invoke: qint64 bytesToWrite()
    C._ZNK8QProcess12bytesToWriteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "bytesToWrite", args)
  }

}

  // proto:  QString QProcess::workingDirectory();
func (this *QProcess) workingDirectory(args ...interface{}) () {
  // workingDirectory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess16workingDirectoryEv
    // invoke: QString workingDirectory()
    C._ZNK8QProcess16workingDirectoryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "workingDirectory", args)
  }

}

  // proto:  void QProcess::setProgram(const QString & program);
func (this *QProcess) setProgram(args ...interface{}) () {
  // setProgram(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess10setProgramERK7QString
    // invoke: void setProgram(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess10setProgramERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setProgram", args)
  }

}

  // proto:  void QProcess::setStandardInputFile(const QString & fileName);
func (this *QProcess) setStandardInputFile(args ...interface{}) () {
  // setStandardInputFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess20setStandardInputFileERK7QString
    // invoke: void setStandardInputFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess20setStandardInputFileERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setStandardInputFile", args)
  }

}

  // proto:  bool QProcessEnvironment::contains(const QString & name);
func (this *QProcessEnvironment) contains(args ...interface{}) () {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QProcessEnvironment8containsERK7QString
    // invoke: bool contains(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QProcessEnvironment8containsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "contains", args)
  }

}

  // proto:  QStringList QProcessEnvironment::keys();
func (this *QProcessEnvironment) keys(args ...interface{}) () {
  // keys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QProcessEnvironment4keysEv
    // invoke: QStringList keys()
    C._ZNK19QProcessEnvironment4keysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "keys", args)
  }

}

  // proto:  void QProcessEnvironment::remove(const QString & name);
func (this *QProcessEnvironment) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment6removeERK7QString
    // invoke: void remove(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QProcessEnvironment6removeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "remove", args)
  }

}

  // proto:  void QProcessEnvironment::clear();
func (this *QProcessEnvironment) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment5clearEv
    // invoke: void clear()
    C._ZN19QProcessEnvironment5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "clear", args)
  }

}

  // proto:  QString QProcessEnvironment::value(const QString & name, const QString & defaultValue);
func (this *QProcessEnvironment) value(args ...interface{}) () {
  // value(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QProcessEnvironment5valueERK7QStringS2_
    // invoke: QString value(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QProcessEnvironment5valueERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "value", args)
  }

}

  // proto:  bool QProcessEnvironment::isEmpty();
func (this *QProcessEnvironment) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QProcessEnvironment7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK19QProcessEnvironment7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "isEmpty", args)
  }

}

  // proto:  void QProcessEnvironment::~QProcessEnvironment();
func (this *QProcessEnvironment) FreeQProcessEnvironment(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "~QProcessEnvironment", args)
  }

}

  // proto:  void QProcessEnvironment::swap(QProcessEnvironment & other);
func (this *QProcessEnvironment) swap(args ...interface{}) () {
  // swap(class QProcessEnvironment &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProcessEnvironment{}) // "QProcessEnvironment &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment4swapERS_
    // invoke: void swap(class QProcessEnvironment &)
    var arg0 = args[0].(QProcessEnvironment).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN19QProcessEnvironment4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "swap", args)
  }

}

  // proto:  void QProcessEnvironment::QProcessEnvironment(const QProcessEnvironment & other);
func NewQProcessEnvironment(args ...interface{}) QProcessEnvironment {
  return QProcessEnvironment{}
}

  // proto: static QProcessEnvironment QProcessEnvironment::systemEnvironment();
func (this *QProcessEnvironment) systemEnvironment_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "systemEnvironment", args)
  }

}

  // proto:  void QProcessEnvironment::insert(const QString & name, const QString & value);
func (this *QProcessEnvironment) insert(args ...interface{}) () {
  // insert(const class QString &, const class QString &)
  // insert(const class QProcessEnvironment &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QProcessEnvironment{}) // "const QProcessEnvironment &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment6insertERK7QStringS2_
    // invoke: void insert(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN19QProcessEnvironment6insertERK7QStringS2_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN19QProcessEnvironment6insertERKS_
    // invoke: void insert(const class QProcessEnvironment &)
    var arg0 = args[0].(QProcessEnvironment).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QProcessEnvironment6insertERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "insert", args)
  }

}

  // proto:  QStringList QProcessEnvironment::toStringList();
func (this *QProcessEnvironment) toStringList(args ...interface{}) () {
  // toStringList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QProcessEnvironment12toStringListEv
    // invoke: QStringList toStringList()
    C._ZNK19QProcessEnvironment12toStringListEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "toStringList", args)
  }

}

// <= body block end

