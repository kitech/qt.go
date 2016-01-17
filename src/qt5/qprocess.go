package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QProcess::setProcessEnvironment(const QProcessEnvironment & environment);
extern void _ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment(void* qthis, void* arg0); // 4
  // proto:  qint64 QProcess::processId();
extern void _ZNK8QProcess9processIdEv(void* qthis); // 4
  // proto:  qint64 QProcess::bytesAvailable();
extern void _ZNK8QProcess14bytesAvailableEv(void* qthis); // 4
  // proto:  QByteArray QProcess::readAllStandardError();
extern void _ZN8QProcess20readAllStandardErrorEv(void* qthis); // 4
  // proto:  qint64 QProcess::bytesToWrite();
extern void _ZNK8QProcess12bytesToWriteEv(void* qthis); // 4
  // proto:  Q_PID QProcess::pid();
extern void _ZNK8QProcess3pidEv(void* qthis); // 4
  // proto:  void QProcess::setArguments(const QStringList & arguments);
extern void _ZN8QProcess12setArgumentsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QProcess::metaObject();
extern void _ZNK8QProcess10metaObjectEv(void* qthis); // 4
  // proto:  void QProcess::kill();
extern void _ZN8QProcess4killEv(void* qthis); // 4
  // proto:  void QProcess::setEnvironment(const QStringList & environment);
extern void _ZN8QProcess14setEnvironmentERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QProcess::close();
extern void _ZN8QProcess5closeEv(void* qthis); // 4
  // proto:  bool QProcess::isSequential();
extern void _ZNK8QProcess12isSequentialEv(void* qthis); // 4
  // proto:  void QProcess::closeWriteChannel();
extern void _ZN8QProcess17closeWriteChannelEv(void* qthis); // 4
  // proto:  bool QProcess::canReadLine();
extern void _ZNK8QProcess11canReadLineEv(void* qthis); // 4
  // proto:  bool QProcess::waitForBytesWritten(int msecs);
extern void _ZN8QProcess19waitForBytesWrittenEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QProcess::workingDirectory();
extern void _ZNK8QProcess16workingDirectoryEv(void* qthis); // 4
  // proto:  QByteArray QProcess::readAllStandardOutput();
extern void _ZN8QProcess21readAllStandardOutputEv(void* qthis); // 4
  // proto:  bool QProcess::waitForFinished(int msecs);
extern void _ZN8QProcess15waitForFinishedEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProcess::setProgram(const QString & program);
extern void _ZN8QProcess10setProgramERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QProcess::environment();
extern void _ZNK8QProcess11environmentEv(void* qthis); // 4
  // proto:  QProcess::ProcessState QProcess::state();
extern void _ZNK8QProcess5stateEv(void* qthis); // 4
  // proto:  void QProcess::~QProcess();
extern void _ZN8QProcessD2Ev(void* qthis); // 4
  // proto:  QStringList QProcess::arguments();
extern void _ZNK8QProcess9argumentsEv(void* qthis); // 4
  // proto:  QString QProcess::program();
extern void _ZNK8QProcess7programEv(void* qthis); // 4
  // proto: static QStringList QProcess::systemEnvironment();
extern void _ZN8QProcess17systemEnvironmentEv(); // 4
  // proto:  void QProcess::setWorkingDirectory(const QString & dir);
extern void _ZN8QProcess19setWorkingDirectoryERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProcess::terminate();
extern void _ZN8QProcess9terminateEv(void* qthis); // 4
  // proto:  void QProcess::setStandardInputFile(const QString & fileName);
extern void _ZN8QProcess20setStandardInputFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QProcess::waitForStarted(int msecs);
extern void _ZN8QProcess14waitForStartedEi(void* qthis, int32_t arg0); // 4
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments, const QString & workingDirectory, qint64 * pid);
extern void _ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px(void* arg0, void* arg1, void* arg2, int64_t* arg3); // 4
  // proto: static bool QProcess::startDetached(const QString & command);
extern void _ZN8QProcess13startDetachedERK7QString(void* arg0); // 4
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments);
extern void _ZN8QProcess13startDetachedERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  QProcess::ProcessChannel QProcess::readChannel();
extern void _ZNK8QProcess11readChannelEv(void* qthis); // 4
  // proto:  QProcess::ProcessChannelMode QProcess::processChannelMode();
extern void _ZNK8QProcess18processChannelModeEv(void* qthis); // 4
  // proto:  int QProcess::exitCode();
extern void _ZNK8QProcess8exitCodeEv(void* qthis); // 4
  // proto:  void QProcess::QProcess(QObject * parent);
extern void _ZN8QProcessC2EP7QObject(void* qthis, void* arg0); // 3
  // proto: static int QProcess::execute(const QString & command);
extern void _ZN8QProcess7executeERK7QString(void* arg0); // 4
  // proto: static int QProcess::execute(const QString & program, const QStringList & arguments);
extern void _ZN8QProcess7executeERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  QProcessEnvironment QProcess::processEnvironment();
extern void _ZNK8QProcess18processEnvironmentEv(void* qthis); // 4
  // proto: static QString QProcess::nullDevice();
extern void _ZN8QProcess10nullDeviceEv(); // 4
  // proto:  bool QProcess::atEnd();
extern void _ZNK8QProcess5atEndEv(void* qthis); // 4
  // proto:  bool QProcess::waitForReadyRead(int msecs);
extern void _ZN8QProcess16waitForReadyReadEi(void* qthis, int32_t arg0); // 4
  // proto:  QProcess::ProcessError QProcess::error();
extern void _ZNK8QProcess5errorEv(void* qthis); // 4
  // proto:  QProcess::InputChannelMode QProcess::inputChannelMode();
extern void _ZNK8QProcess16inputChannelModeEv(void* qthis); // 4
  // proto:  void QProcess::setStandardOutputProcess(QProcess * destination);
extern void _ZN8QProcess24setStandardOutputProcessEPS_(void* qthis, void* arg0); // 4
  // proto:  QProcess::ExitStatus QProcess::exitStatus();
extern void _ZNK8QProcess10exitStatusEv(void* qthis); // 4
  // proto:  QProcess::ProcessChannelMode QProcess::readChannelMode();
extern void _ZNK8QProcess15readChannelModeEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::QProcessEnvironment(const QProcessEnvironment & other);
extern void _ZN19QProcessEnvironmentC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QProcessEnvironment::QProcessEnvironment();
extern void _ZN19QProcessEnvironmentC2Ev(void* qthis); // 3
  // proto: static QProcessEnvironment QProcessEnvironment::systemEnvironment();
extern void _ZN19QProcessEnvironment17systemEnvironmentEv(); // 4
  // proto:  void QProcessEnvironment::insert(const QProcessEnvironment & e);
extern void _ZN19QProcessEnvironment6insertERKS_(void* qthis, void* arg0); // 4
  // proto:  void QProcessEnvironment::insert(const QString & name, const QString & value);
extern void _ZN19QProcessEnvironment6insertERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QProcessEnvironment::keys();
extern void _ZNK19QProcessEnvironment4keysEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::clear();
extern void _ZN19QProcessEnvironment5clearEv(void* qthis); // 4
  // proto:  bool QProcessEnvironment::contains(const QString & name);
extern void _ZNK19QProcessEnvironment8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QProcessEnvironment::value(const QString & name, const QString & defaultValue);
extern void _ZNK19QProcessEnvironment5valueERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QProcessEnvironment::toStringList();
extern void _ZNK19QProcessEnvironment12toStringListEv(void* qthis); // 4
  // proto:  bool QProcessEnvironment::isEmpty();
extern void _ZNK19QProcessEnvironment7isEmptyEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::remove(const QString & name);
extern void _ZN19QProcessEnvironment6removeERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProcessEnvironment::~QProcessEnvironment();
extern void _ZN19QProcessEnvironmentD2Ev(void* qthis); // 4
  // proto:  void QProcessEnvironment::swap(QProcessEnvironment & other);
extern void _ZN19QProcessEnvironment4swapERS_(void* qthis, void* arg0); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// setProcessEnvironment(const class QProcessEnvironment &)
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

// processId()
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

// bytesAvailable()
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

// readAllStandardError()
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

// bytesToWrite()
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

// pid()
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

// setArguments(const class QStringList &)
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

// metaObject()
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

// kill()
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

// setEnvironment(const class QStringList &)
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

// close()
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

// isSequential()
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

// closeWriteChannel()
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

// canReadLine()
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

// waitForBytesWritten(int)
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

// workingDirectory()
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

// readAllStandardOutput()
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

// waitForFinished(int)
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

// setProgram(const class QString &)
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

// environment()
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

// state()
func (this *QProcess) state(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess5stateEv
    // invoke: QProcess::ProcessState state()
    C._ZNK8QProcess5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "state", args)
  }

}

// ~QProcess()
func (this *QProcess) FreeQProcess(args ...interface{}) () {
  // ~QProcess()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcessD0Ev
    // invoke: void ~QProcess()
    C._ZN8QProcessD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "~QProcess", args)
  }

}

// arguments()
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

// program()
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

// systemEnvironment()
func (this *QProcess) systemEnvironment_s(args ...interface{}) () {
  // systemEnvironment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess17systemEnvironmentEv
    // invoke: QStringList systemEnvironment()
    C._ZN8QProcess17systemEnvironmentEv()
  default:
    qtrt.ErrorResolve("QProcess", "systemEnvironment", args)
  }

}

// setWorkingDirectory(const class QString &)
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

// terminate()
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

// setStandardInputFile(const class QString &)
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

// waitForStarted(int)
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

// startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
func (this *QProcess) startDetached_s(args ...interface{}) () {
  // startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
  // startDetached(const class QString &)
  // startDetached(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int64Ty(true) // "qint64 *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px
    // invoke: bool startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int64_t)(args[3].(*int64))
    if false {fmt.Println(arg3)}
    C._ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px(arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QProcess13startDetachedERK7QString
    // invoke: bool startDetached(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess13startDetachedERK7QString(arg0)
  case 2:
    // invoke: _ZN8QProcess13startDetachedERK7QStringRK11QStringList
    // invoke: bool startDetached(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QProcess13startDetachedERK7QStringRK11QStringList(arg0, arg1)
  default:
    qtrt.ErrorResolve("QProcess", "startDetached", args)
  }

}

// readChannel()
func (this *QProcess) readChannel(args ...interface{}) () {
  // readChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess11readChannelEv
    // invoke: QProcess::ProcessChannel readChannel()
    C._ZNK8QProcess11readChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readChannel", args)
  }

}

// processChannelMode()
func (this *QProcess) processChannelMode(args ...interface{}) () {
  // processChannelMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess18processChannelModeEv
    // invoke: QProcess::ProcessChannelMode processChannelMode()
    C._ZNK8QProcess18processChannelModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "processChannelMode", args)
  }

}

// exitCode()
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

// QProcess(class QObject *)
func NewQProcess(args ...interface{}) QProcess {
  // QProcess(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcessC1EP7QObject
    // invoke: void QProcess(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QProcessC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "QProcess", args)
  }

  return QProcess{}
}

// execute(const class QString &)
func (this *QProcess) execute_s(args ...interface{}) () {
  // execute(const class QString &)
  // execute(const class QString &, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess7executeERK7QString
    // invoke: int execute(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QProcess7executeERK7QString(arg0)
  case 1:
    // invoke: _ZN8QProcess7executeERK7QStringRK11QStringList
    // invoke: int execute(const class QString &, const class QStringList &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QProcess7executeERK7QStringRK11QStringList(arg0, arg1)
  default:
    qtrt.ErrorResolve("QProcess", "execute", args)
  }

}

// processEnvironment()
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

// nullDevice()
func (this *QProcess) nullDevice_s(args ...interface{}) () {
  // nullDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QProcess10nullDeviceEv
    // invoke: QString nullDevice()
    C._ZN8QProcess10nullDeviceEv()
  default:
    qtrt.ErrorResolve("QProcess", "nullDevice", args)
  }

}

// atEnd()
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

// waitForReadyRead(int)
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

// error()
func (this *QProcess) error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess5errorEv
    // invoke: QProcess::ProcessError error()
    C._ZNK8QProcess5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "error", args)
  }

}

// inputChannelMode()
func (this *QProcess) inputChannelMode(args ...interface{}) () {
  // inputChannelMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess16inputChannelModeEv
    // invoke: QProcess::InputChannelMode inputChannelMode()
    C._ZNK8QProcess16inputChannelModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "inputChannelMode", args)
  }

}

// setStandardOutputProcess(class QProcess *)
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

// exitStatus()
func (this *QProcess) exitStatus(args ...interface{}) () {
  // exitStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess10exitStatusEv
    // invoke: QProcess::ExitStatus exitStatus()
    C._ZNK8QProcess10exitStatusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "exitStatus", args)
  }

}

// readChannelMode()
func (this *QProcess) readChannelMode(args ...interface{}) () {
  // readChannelMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QProcess15readChannelModeEv
    // invoke: QProcess::ProcessChannelMode readChannelMode()
    C._ZNK8QProcess15readChannelModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readChannelMode", args)
  }

}

// QProcessEnvironment(const class QProcessEnvironment &)
func NewQProcessEnvironment(args ...interface{}) QProcessEnvironment {
  // QProcessEnvironment(const class QProcessEnvironment &)
  // QProcessEnvironment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProcessEnvironment{}) // "const QProcessEnvironment &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironmentC1ERKS_
    // invoke: void QProcessEnvironment(const class QProcessEnvironment &)
    var arg0 = args[0].(QProcessEnvironment).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QProcessEnvironmentC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN19QProcessEnvironmentC1Ev
    // invoke: void QProcessEnvironment()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QProcessEnvironmentC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "QProcessEnvironment", args)
  }

  return QProcessEnvironment{}
}

// systemEnvironment()
func (this *QProcessEnvironment) systemEnvironment_s(args ...interface{}) () {
  // systemEnvironment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment17systemEnvironmentEv
    // invoke: QProcessEnvironment systemEnvironment()
    C._ZN19QProcessEnvironment17systemEnvironmentEv()
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "systemEnvironment", args)
  }

}

// insert(const class QProcessEnvironment &)
func (this *QProcessEnvironment) insert(args ...interface{}) () {
  // insert(const class QProcessEnvironment &)
  // insert(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProcessEnvironment{}) // "const QProcessEnvironment &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironment6insertERKS_
    // invoke: void insert(const class QProcessEnvironment &)
    var arg0 = args[0].(QProcessEnvironment).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QProcessEnvironment6insertERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QProcessEnvironment6insertERK7QStringS2_
    // invoke: void insert(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN19QProcessEnvironment6insertERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "insert", args)
  }

}

// keys()
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

// clear()
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

// contains(const class QString &)
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

// value(const class QString &, const class QString &)
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

// toStringList()
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

// isEmpty()
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

// remove(const class QString &)
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

// ~QProcessEnvironment()
func (this *QProcessEnvironment) FreeQProcessEnvironment(args ...interface{}) () {
  // ~QProcessEnvironment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QProcessEnvironmentD0Ev
    // invoke: void ~QProcessEnvironment()
    C._ZN19QProcessEnvironmentD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "~QProcessEnvironment", args)
  }

}

// swap(class QProcessEnvironment &)
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
    C._ZN19QProcessEnvironment4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "swap", args)
  }

}

// <= body block end

