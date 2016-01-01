package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qprocess.h
// dst-file: /src/core/qprocess.go
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
  default:
    qtrt.ErrorResolve("QProcess", "close", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setEnvironment", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "atEnd", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setProcessEnvironment", args)
 }

}


func NewQProcess(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QProcess", "pid", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setArguments", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "closeWriteChannel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "processEnvironment", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardOutput", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "waitForBytesWritten", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "program", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "processId", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "arguments", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "isSequential", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "waitForReadyRead", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setWorkingDirectory", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "terminate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "kill", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "bytesAvailable", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "waitForStarted", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "readAllStandardError", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "exitCode", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "environment", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "canReadLine", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setStandardOutputProcess", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "waitForFinished", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "bytesToWrite", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "workingDirectory", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setProgram", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcess", "setStandardInputFile", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "contains", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "keys", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "remove", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "clear", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "value", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "isEmpty", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "swap", args)
 }

}


func NewQProcessEnvironment(args ...interface{})() {
}


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
  case 1:
    // invoke: _ZN19QProcessEnvironment6insertERKS_
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "insert", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QProcessEnvironment", "toStringList", args)
 }

}

// <= body block end

