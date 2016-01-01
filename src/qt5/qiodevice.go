package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qiodevice.h
// dst-file: /src/core/qiodevice.go
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
// class sizeof(QIODevice)=1
type QIODevice struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _readyRead QIODevice_readyRead_signal;
//  _readChannelFinished QIODevice_readChannelFinished_signal;
//  _aboutToClose QIODevice_aboutToClose_signal;
//  _bytesWritten QIODevice_bytesWritten_signal;
}


func (this *QIODevice) ungetChar(args ...interface{}) () {
  // ungetChar(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice9ungetCharEc
  default:
    qtrt.ErrorResolve("QIODevice", "ungetChar", args)
  }

}


func (this *QIODevice) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11errorStringEv
  default:
    qtrt.ErrorResolve("QIODevice", "errorString", args)
  }

}


func (this *QIODevice) write(args ...interface{}) () {
  // write(const class QByteArray &)
  // write(const char *)
  // write(const char *, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5writeERK10QByteArray
  case 1:
    // invoke: _ZN9QIODevice5writeEPKc
  case 2:
    // invoke: _ZN9QIODevice5writeEPKcx
  default:
    qtrt.ErrorResolve("QIODevice", "write", args)
  }

}


func (this *QIODevice) isReadable(args ...interface{}) () {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isReadableEv
  default:
    qtrt.ErrorResolve("QIODevice", "isReadable", args)
  }

}


func (this *QIODevice) readLine(args ...interface{}) () {
  // readLine(char *, qint64)
  // readLine(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice8readLineEPcx
  case 1:
    // invoke: _ZN9QIODevice8readLineEx
  default:
    qtrt.ErrorResolve("QIODevice", "readLine", args)
  }

}


func (this *QIODevice) waitForReadyRead(args ...interface{}) () {
  // waitForReadyRead(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice16waitForReadyReadEi
  default:
    qtrt.ErrorResolve("QIODevice", "waitForReadyRead", args)
  }

}


func (this *QIODevice) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice4sizeEv
  default:
    qtrt.ErrorResolve("QIODevice", "size", args)
  }

}


func (this *QIODevice) getChar(args ...interface{}) () {
  // getChar(char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7getCharEPc
  default:
    qtrt.ErrorResolve("QIODevice", "getChar", args)
  }

}


func (this *QIODevice) putChar(args ...interface{}) () {
  // putChar(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7putCharEc
  default:
    qtrt.ErrorResolve("QIODevice", "putChar", args)
  }

}


func (this *QIODevice) isTextModeEnabled(args ...interface{}) () {
  // isTextModeEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice17isTextModeEnabledEv
  default:
    qtrt.ErrorResolve("QIODevice", "isTextModeEnabled", args)
  }

}


func (this *QIODevice) isSequential(args ...interface{}) () {
  // isSequential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12isSequentialEv
  default:
    qtrt.ErrorResolve("QIODevice", "isSequential", args)
  }

}


func (this *QIODevice) bytesAvailable(args ...interface{}) () {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice14bytesAvailableEv
  default:
    qtrt.ErrorResolve("QIODevice", "bytesAvailable", args)
  }

}


func (this *QIODevice) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5closeEv
  default:
    qtrt.ErrorResolve("QIODevice", "close", args)
  }

}


func (this *QIODevice) readAll(args ...interface{}) () {
  // readAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7readAllEv
  default:
    qtrt.ErrorResolve("QIODevice", "readAll", args)
  }

}


func (this *QIODevice) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice5atEndEv
  default:
    qtrt.ErrorResolve("QIODevice", "atEnd", args)
  }

}


func (this *QIODevice) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4seekEx
  default:
    qtrt.ErrorResolve("QIODevice", "seek", args)
  }

}


func NewQIODevice(args ...interface{}) QIODevice {
  return QIODevice{}
}


func (this *QIODevice) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice3posEv
  default:
    qtrt.ErrorResolve("QIODevice", "pos", args)
  }

}


func (this *QIODevice) read(args ...interface{}) () {
  // read(qint64)
  // read(char *, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "char *"
  vtys[1][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4readEx
  case 1:
    // invoke: _ZN9QIODevice4readEPcx
  default:
    qtrt.ErrorResolve("QIODevice", "read", args)
  }

}


func (this *QIODevice) peek(args ...interface{}) () {
  // peek(char *, qint64)
  // peek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice4peekEPcx
  case 1:
    // invoke: _ZN9QIODevice4peekEx
  default:
    qtrt.ErrorResolve("QIODevice", "peek", args)
  }

}


func (this *QIODevice) waitForBytesWritten(args ...interface{}) () {
  // waitForBytesWritten(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice19waitForBytesWrittenEi
  default:
    qtrt.ErrorResolve("QIODevice", "waitForBytesWritten", args)
  }

}


func (this *QIODevice) bytesToWrite(args ...interface{}) () {
  // bytesToWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12bytesToWriteEv
  default:
    qtrt.ErrorResolve("QIODevice", "bytesToWrite", args)
  }

}


func (this *QIODevice) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5resetEv
  default:
    qtrt.ErrorResolve("QIODevice", "reset", args)
  }

}


func (this *QIODevice) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isWritableEv
  default:
    qtrt.ErrorResolve("QIODevice", "isWritable", args)
  }

}


func (this *QIODevice) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10metaObjectEv
  default:
    qtrt.ErrorResolve("QIODevice", "metaObject", args)
  }

}


func (this *QIODevice) setTextModeEnabled(args ...interface{}) () {
  // setTextModeEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice18setTextModeEnabledEb
  default:
    qtrt.ErrorResolve("QIODevice", "setTextModeEnabled", args)
  }

}


func (this *QIODevice) isOpen(args ...interface{}) () {
  // isOpen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice6isOpenEv
  default:
    qtrt.ErrorResolve("QIODevice", "isOpen", args)
  }

}


func (this *QIODevice) canReadLine(args ...interface{}) () {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11canReadLineEv
  default:
    qtrt.ErrorResolve("QIODevice", "canReadLine", args)
  }

}


func (this *QIODevice) FreeQIODevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIODevice", "~QIODevice", args)
  }

}

// <= body block end

