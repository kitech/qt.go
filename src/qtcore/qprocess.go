package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void C_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment(void* qthis, void* arg0); // 4
  // proto:  qint64 QProcess::processId();
extern int64_t C_ZNK8QProcess9processIdEv(void* qthis); // 4
  // proto:  qint64 QProcess::bytesAvailable();
extern int64_t C_ZNK8QProcess14bytesAvailableEv(void* qthis); // 4
  // proto:  QByteArray QProcess::readAllStandardError();
extern void* C_ZN8QProcess20readAllStandardErrorEv(void* qthis); // 4
  // proto:  qint64 QProcess::bytesToWrite();
extern int64_t C_ZNK8QProcess12bytesToWriteEv(void* qthis); // 4
  // proto:  Q_PID QProcess::pid();
extern int64_t C_ZNK8QProcess3pidEv(void* qthis); // 4
  // proto:  void QProcess::setArguments(const QStringList & arguments);
extern void C_ZN8QProcess12setArgumentsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QProcess::metaObject();
extern void C_ZNK8QProcess10metaObjectEv(void* qthis); // 4
  // proto:  void QProcess::kill();
extern void C_ZN8QProcess4killEv(void* qthis); // 4
  // proto:  void QProcess::setEnvironment(const QStringList & environment);
extern void C_ZN8QProcess14setEnvironmentERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QProcess::close();
extern void C_ZN8QProcess5closeEv(void* qthis); // 4
  // proto:  bool QProcess::isSequential();
extern bool C_ZNK8QProcess12isSequentialEv(void* qthis); // 4
  // proto:  void QProcess::closeWriteChannel();
extern void C_ZN8QProcess17closeWriteChannelEv(void* qthis); // 4
  // proto:  bool QProcess::canReadLine();
extern bool C_ZNK8QProcess11canReadLineEv(void* qthis); // 4
  // proto:  bool QProcess::waitForBytesWritten(int msecs);
extern bool C_ZN8QProcess19waitForBytesWrittenEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QProcess::workingDirectory();
extern void* C_ZNK8QProcess16workingDirectoryEv(void* qthis); // 4
  // proto:  QByteArray QProcess::readAllStandardOutput();
extern void* C_ZN8QProcess21readAllStandardOutputEv(void* qthis); // 4
  // proto:  bool QProcess::waitForFinished(int msecs);
extern bool C_ZN8QProcess15waitForFinishedEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProcess::setProgram(const QString & program);
extern void C_ZN8QProcess10setProgramERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QProcess::environment();
extern void C_ZNK8QProcess11environmentEv(void* qthis); // 4
  // proto:  QProcess::ProcessState QProcess::state();
extern void C_ZNK8QProcess5stateEv(void* qthis); // 4
  // proto:  void QProcess::~QProcess();
extern void C_ZN8QProcessD2Ev(void* qthis); // 4
  // proto:  QStringList QProcess::arguments();
extern void C_ZNK8QProcess9argumentsEv(void* qthis); // 4
  // proto:  QString QProcess::program();
extern void* C_ZNK8QProcess7programEv(void* qthis); // 4
  // proto: static QStringList QProcess::systemEnvironment();
extern void C_ZN8QProcess17systemEnvironmentEv(); // 4
  // proto:  void QProcess::setWorkingDirectory(const QString & dir);
extern void C_ZN8QProcess19setWorkingDirectoryERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProcess::terminate();
extern void C_ZN8QProcess9terminateEv(void* qthis); // 4
  // proto:  void QProcess::setStandardInputFile(const QString & fileName);
extern void C_ZN8QProcess20setStandardInputFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QProcess::waitForStarted(int msecs);
extern bool C_ZN8QProcess14waitForStartedEi(void* qthis, int32_t arg0); // 4
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments, const QString & workingDirectory, qint64 * pid);
extern bool C_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px(void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto: static bool QProcess::startDetached(const QString & command);
extern bool C_ZN8QProcess13startDetachedERK7QString(void* arg0); // 4
  // proto: static bool QProcess::startDetached(const QString & program, const QStringList & arguments);
extern bool C_ZN8QProcess13startDetachedERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  QProcess::ProcessChannel QProcess::readChannel();
extern void C_ZNK8QProcess11readChannelEv(void* qthis); // 4
  // proto:  QProcess::ProcessChannelMode QProcess::processChannelMode();
extern void C_ZNK8QProcess18processChannelModeEv(void* qthis); // 4
  // proto:  int QProcess::exitCode();
extern int32_t C_ZNK8QProcess8exitCodeEv(void* qthis); // 4
  // proto:  void QProcess::QProcess(QObject * parent);
extern void* C_ZN8QProcessC2EP7QObject(void* arg0); // 3
  // proto: static int QProcess::execute(const QString & command);
extern int32_t C_ZN8QProcess7executeERK7QString(void* arg0); // 4
  // proto: static int QProcess::execute(const QString & program, const QStringList & arguments);
extern int32_t C_ZN8QProcess7executeERK7QStringRK11QStringList(void* arg0, void* arg1); // 4
  // proto:  QProcessEnvironment QProcess::processEnvironment();
extern void* C_ZNK8QProcess18processEnvironmentEv(void* qthis); // 4
  // proto: static QString QProcess::nullDevice();
extern void* C_ZN8QProcess10nullDeviceEv(); // 4
  // proto:  bool QProcess::atEnd();
extern bool C_ZNK8QProcess5atEndEv(void* qthis); // 4
  // proto:  bool QProcess::waitForReadyRead(int msecs);
extern bool C_ZN8QProcess16waitForReadyReadEi(void* qthis, int32_t arg0); // 4
  // proto:  QProcess::ProcessError QProcess::error();
extern void C_ZNK8QProcess5errorEv(void* qthis); // 4
  // proto:  QProcess::InputChannelMode QProcess::inputChannelMode();
extern void C_ZNK8QProcess16inputChannelModeEv(void* qthis); // 4
  // proto:  void QProcess::setStandardOutputProcess(QProcess * destination);
extern void C_ZN8QProcess24setStandardOutputProcessEPS_(void* qthis, void* arg0); // 4
  // proto:  QProcess::ExitStatus QProcess::exitStatus();
extern void C_ZNK8QProcess10exitStatusEv(void* qthis); // 4
  // proto:  QProcess::ProcessChannelMode QProcess::readChannelMode();
extern void C_ZNK8QProcess15readChannelModeEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::QProcessEnvironment(const QProcessEnvironment & other);
extern void* C_ZN19QProcessEnvironmentC2ERKS_(void* arg0); // 3
  // proto:  void QProcessEnvironment::QProcessEnvironment();
extern void* C_ZN19QProcessEnvironmentC2Ev(); // 3
  // proto: static QProcessEnvironment QProcessEnvironment::systemEnvironment();
extern void* C_ZN19QProcessEnvironment17systemEnvironmentEv(); // 4
  // proto:  void QProcessEnvironment::insert(const QProcessEnvironment & e);
extern void C_ZN19QProcessEnvironment6insertERKS_(void* qthis, void* arg0); // 4
  // proto:  void QProcessEnvironment::insert(const QString & name, const QString & value);
extern void C_ZN19QProcessEnvironment6insertERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QProcessEnvironment::keys();
extern void C_ZNK19QProcessEnvironment4keysEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::clear();
extern void C_ZN19QProcessEnvironment5clearEv(void* qthis); // 4
  // proto:  bool QProcessEnvironment::contains(const QString & name);
extern bool C_ZNK19QProcessEnvironment8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QProcessEnvironment::value(const QString & name, const QString & defaultValue);
extern void* C_ZNK19QProcessEnvironment5valueERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QProcessEnvironment::toStringList();
extern void C_ZNK19QProcessEnvironment12toStringListEv(void* qthis); // 4
  // proto:  bool QProcessEnvironment::isEmpty();
extern bool C_ZNK19QProcessEnvironment7isEmptyEv(void* qthis); // 4
  // proto:  void QProcessEnvironment::remove(const QString & name);
extern void C_ZN19QProcessEnvironment6removeERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProcessEnvironment::~QProcessEnvironment();
extern void C_ZN19QProcessEnvironmentD2Ev(void* qthis); // 4
  // proto:  void QProcessEnvironment::swap(QProcessEnvironment & other);
extern void C_ZN19QProcessEnvironment4swapERS_(void* qthis, void* arg0); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setProcessEnvironment(const class QProcessEnvironment &)
func (this *QProcess) Setprocessenvironment(args ...interface{}) () {
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
    var arg0 = args[0].(*QProcessEnvironment).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setProcessEnvironment", args)
  }

  return
}

// processId()
func (this *QProcess) Processid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess9processIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "processId", args)
  }

  return
}

// bytesAvailable()
func (this *QProcess) Bytesavailable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess14bytesAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "bytesAvailable", args)
  }

  return
}

// readAllStandardError()
func (this *QProcess) Readallstandarderror(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QProcess20readAllStandardErrorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardError", args)
  }

  return
}

// bytesToWrite()
func (this *QProcess) Bytestowrite(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess12bytesToWriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "bytesToWrite", args)
  }

  return
}

// pid()
func (this *QProcess) Pid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess3pidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "Q_PID"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "pid", args)
  }

  return
}

// setArguments(const class QStringList &)
func (this *QProcess) Setarguments(args ...interface{}) () {
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
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess12setArgumentsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setArguments", args)
  }

  return
}

// metaObject()
func (this *QProcess) Metaobject(args ...interface{}) () {
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
    C.C_ZNK8QProcess10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "metaObject", args)
  }

  return
}

// kill()
func (this *QProcess) Kill(args ...interface{}) () {
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
    C.C_ZN8QProcess4killEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "kill", args)
  }

  return
}

// setEnvironment(const class QStringList &)
func (this *QProcess) Setenvironment(args ...interface{}) () {
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
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess14setEnvironmentERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setEnvironment", args)
  }

  return
}

// close()
func (this *QProcess) Close(args ...interface{}) () {
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
    C.C_ZN8QProcess5closeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "close", args)
  }

  return
}

// isSequential()
func (this *QProcess) Issequential(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess12isSequentialEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "isSequential", args)
  }

  return
}

// closeWriteChannel()
func (this *QProcess) Closewritechannel(args ...interface{}) () {
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
    C.C_ZN8QProcess17closeWriteChannelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "closeWriteChannel", args)
  }

  return
}

// canReadLine()
func (this *QProcess) Canreadline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess11canReadLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "canReadLine", args)
  }

  return
}

// waitForBytesWritten(int)
func (this *QProcess) Waitforbyteswritten(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess19waitForBytesWrittenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "waitForBytesWritten", args)
  }

  return
}

// workingDirectory()
func (this *QProcess) Workingdirectory(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess16workingDirectoryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "workingDirectory", args)
  }

  return
}

// readAllStandardOutput()
func (this *QProcess) Readallstandardoutput(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QProcess21readAllStandardOutputEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardOutput", args)
  }

  return
}

// waitForFinished(int)
func (this *QProcess) Waitforfinished(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess15waitForFinishedEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "waitForFinished", args)
  }

  return
}

// setProgram(const class QString &)
func (this *QProcess) Setprogram(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess10setProgramERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setProgram", args)
  }

  return
}

// environment()
func (this *QProcess) Environment(args ...interface{}) () {
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
    C.C_ZNK8QProcess11environmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "environment", args)
  }

  return
}

// state()
func (this *QProcess) State(args ...interface{}) () {
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
    C.C_ZNK8QProcess5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "state", args)
  }

  return
}

// ~QProcess()
func (this *QProcess) Freeqprocess(args ...interface{}) () {
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
    C.C_ZN8QProcessD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "~QProcess", args)
  }

  return
}

// arguments()
func (this *QProcess) Arguments(args ...interface{}) () {
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
    C.C_ZNK8QProcess9argumentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "arguments", args)
  }

  return
}

// program()
func (this *QProcess) Program(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess7programEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "program", args)
  }

  return
}

// systemEnvironment()
func (this *QProcess) Systemenvironment_S(args ...interface{}) () {
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
    C.C_ZN8QProcess17systemEnvironmentEv()
  default:
    qtrt.ErrorResolve("QProcess", "systemEnvironment", args)
  }

  return
}

// setWorkingDirectory(const class QString &)
func (this *QProcess) Setworkingdirectory(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess19setWorkingDirectoryERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setWorkingDirectory", args)
  }

  return
}

// terminate()
func (this *QProcess) Terminate(args ...interface{}) () {
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
    C.C_ZN8QProcess9terminateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "terminate", args)
  }

  return
}

// setStandardInputFile(const class QString &)
func (this *QProcess) Setstandardinputfile(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess20setStandardInputFileERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setStandardInputFile", args)
  }

  return
}

// waitForStarted(int)
func (this *QProcess) Waitforstarted(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess14waitForStartedEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "waitForStarted", args)
  }

  return
}

// startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
func (this *QProcess) Startdetached_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QProcess13startDetachedERK7QString
    // invoke: bool startDetached(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess13startDetachedERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN8QProcess13startDetachedERK7QStringRK11QStringList
    // invoke: bool startDetached(const class QString &, const class QStringList &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QProcess13startDetachedERK7QStringRK11QStringList(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "startDetached", args)
  }

  return
}

// readChannel()
func (this *QProcess) Readchannel(args ...interface{}) () {
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
    C.C_ZNK8QProcess11readChannelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readChannel", args)
  }

  return
}

// processChannelMode()
func (this *QProcess) Processchannelmode(args ...interface{}) () {
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
    C.C_ZNK8QProcess18processChannelModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "processChannelMode", args)
  }

  return
}

// exitCode()
func (this *QProcess) Exitcode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess8exitCodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "exitCode", args)
  }

  return
}

// QProcess(class QObject *)
func NewQProcess(args ...interface{}) *QProcess {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QProcessC2EP7QObject(arg0)
    return &QProcess{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QProcess", "QProcess", args)
  }

  return nil // QProcess{Qclsinst:qthis}
}

// execute(const class QString &)
func (this *QProcess) Execute_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess7executeERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QProcess7executeERK7QStringRK11QStringList
    // invoke: int execute(const class QString &, const class QStringList &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QProcess7executeERK7QStringRK11QStringList(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "execute", args)
  }

  return
}

// processEnvironment()
func (this *QProcess) Processenvironment(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess18processEnvironmentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QProcessEnvironment{}) // "QProcessEnvironment"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "processEnvironment", args)
  }

  return
}

// nullDevice()
func (this *QProcess) Nulldevice_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QProcess10nullDeviceEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "nullDevice", args)
  }

  return
}

// atEnd()
func (this *QProcess) Atend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QProcess5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "atEnd", args)
  }

  return
}

// waitForReadyRead(int)
func (this *QProcess) Waitforreadyread(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QProcess16waitForReadyReadEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcess", "waitForReadyRead", args)
  }

  return
}

// error()
func (this *QProcess) Error(args ...interface{}) () {
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
    C.C_ZNK8QProcess5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "error", args)
  }

  return
}

// inputChannelMode()
func (this *QProcess) Inputchannelmode(args ...interface{}) () {
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
    C.C_ZNK8QProcess16inputChannelModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "inputChannelMode", args)
  }

  return
}

// setStandardOutputProcess(class QProcess *)
func (this *QProcess) Setstandardoutputprocess(args ...interface{}) () {
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
    var arg0 = args[0].(*QProcess).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QProcess24setStandardOutputProcessEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcess", "setStandardOutputProcess", args)
  }

  return
}

// exitStatus()
func (this *QProcess) Exitstatus(args ...interface{}) () {
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
    C.C_ZNK8QProcess10exitStatusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "exitStatus", args)
  }

  return
}

// readChannelMode()
func (this *QProcess) Readchannelmode(args ...interface{}) () {
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
    C.C_ZNK8QProcess15readChannelModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcess", "readChannelMode", args)
  }

  return
}

// QProcessEnvironment(const class QProcessEnvironment &)
func NewQProcessEnvironment(args ...interface{}) *QProcessEnvironment {
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
    var arg0 = args[0].(*QProcessEnvironment).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QProcessEnvironmentC2ERKS_(arg0)
    return &QProcessEnvironment{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QProcessEnvironmentC1Ev
    // invoke: void QProcessEnvironment()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QProcessEnvironmentC2Ev()
    return &QProcessEnvironment{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "QProcessEnvironment", args)
  }

  return nil // QProcessEnvironment{Qclsinst:qthis}
}

// systemEnvironment()
func (this *QProcessEnvironment) Systemenvironment_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN19QProcessEnvironment17systemEnvironmentEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QProcessEnvironment{}) // "QProcessEnvironment"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "systemEnvironment", args)
  }

  return
}

// insert(const class QProcessEnvironment &)
func (this *QProcessEnvironment) Insert(args ...interface{}) () {
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
    var arg0 = args[0].(*QProcessEnvironment).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QProcessEnvironment6insertERKS_(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN19QProcessEnvironment6insertERK7QStringS2_
    // invoke: void insert(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN19QProcessEnvironment6insertERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "insert", args)
  }

  return
}

// keys()
func (this *QProcessEnvironment) Keys(args ...interface{}) () {
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
    C.C_ZNK19QProcessEnvironment4keysEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "keys", args)
  }

  return
}

// clear()
func (this *QProcessEnvironment) Clear(args ...interface{}) () {
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
    C.C_ZN19QProcessEnvironment5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "clear", args)
  }

  return
}

// contains(const class QString &)
func (this *QProcessEnvironment) Contains(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QProcessEnvironment8containsERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "contains", args)
  }

  return
}

// value(const class QString &, const class QString &)
func (this *QProcessEnvironment) Value(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK19QProcessEnvironment5valueERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "value", args)
  }

  return
}

// toStringList()
func (this *QProcessEnvironment) Tostringlist(args ...interface{}) () {
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
    C.C_ZNK19QProcessEnvironment12toStringListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "toStringList", args)
  }

  return
}

// isEmpty()
func (this *QProcessEnvironment) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QProcessEnvironment7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "isEmpty", args)
  }

  return
}

// remove(const class QString &)
func (this *QProcessEnvironment) Remove(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QProcessEnvironment6removeERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "remove", args)
  }

  return
}

// ~QProcessEnvironment()
func (this *QProcessEnvironment) Freeqprocessenvironment(args ...interface{}) () {
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
    C.C_ZN19QProcessEnvironmentD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "~QProcessEnvironment", args)
  }

  return
}

// swap(class QProcessEnvironment &)
func (this *QProcessEnvironment) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QProcessEnvironment).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QProcessEnvironment4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "swap", args)
  }

  return
}

// <= body block end

