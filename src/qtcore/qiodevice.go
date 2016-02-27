package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qiodevice.h
// dst-file: /src/core/qiodevice.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qint64 QIODevice::bytesAvailable();
extern int64_t C_ZNK9QIODevice14bytesAvailableEv(void* qthis); // 4
  // proto:  qint64 QIODevice::pos();
extern int64_t C_ZNK9QIODevice3posEv(void* qthis); // 4
  // proto:  qint64 QIODevice::bytesToWrite();
extern int64_t C_ZNK9QIODevice12bytesToWriteEv(void* qthis); // 4
  // proto:  void QIODevice::close();
extern void C_ZN9QIODevice5closeEv(void* qthis); // 4
  // proto:  bool QIODevice::seek(qint64 pos);
extern bool C_ZN9QIODevice4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QIODevice::isSequential();
extern bool C_ZNK9QIODevice12isSequentialEv(void* qthis); // 4
  // proto:  void QIODevice::ungetChar(char c);
extern void C_ZN9QIODevice9ungetCharEc(void* qthis, unsigned char arg0); // 4
  // proto:  qint64 QIODevice::size();
extern int64_t C_ZNK9QIODevice4sizeEv(void* qthis); // 4
  // proto:  qint64 QIODevice::write(const char * data, qint64 len);
extern int64_t C_ZN9QIODevice5writeEPKcx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  qint64 QIODevice::write(const char * data);
extern int64_t C_ZN9QIODevice5writeEPKc(void* qthis, void* arg0); // 4
  // proto:  qint64 QIODevice::write(const QByteArray & data);
extern int64_t C_ZN9QIODevice5writeERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  bool QIODevice::canReadLine();
extern bool C_ZNK9QIODevice11canReadLineEv(void* qthis); // 4
  // proto:  bool QIODevice::waitForBytesWritten(int msecs);
extern bool C_ZN9QIODevice19waitForBytesWrittenEi(void* qthis, int32_t arg0); // 4
  // proto:  qint64 QIODevice::read(char * data, qint64 maxlen);
extern int64_t C_ZN9QIODevice4readEPcx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  QByteArray QIODevice::read(qint64 maxlen);
extern void* C_ZN9QIODevice4readEx(void* qthis, int64_t arg0); // 4
  // proto:  qint64 QIODevice::peek(char * data, qint64 maxlen);
extern int64_t C_ZN9QIODevice4peekEPcx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  QByteArray QIODevice::peek(qint64 maxlen);
extern void* C_ZN9QIODevice4peekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QIODevice::putChar(char c);
extern bool C_ZN9QIODevice7putCharEc(void* qthis, unsigned char arg0); // 4
  // proto:  QString QIODevice::errorString();
extern void* C_ZNK9QIODevice11errorStringEv(void* qthis); // 4
  // proto:  void QIODevice::~QIODevice();
extern void C_ZN9QIODeviceD2Ev(void* qthis); // 4
  // proto:  QByteArray QIODevice::readAll();
extern void* C_ZN9QIODevice7readAllEv(void* qthis); // 4
  // proto:  void QIODevice::setTextModeEnabled(bool enabled);
extern void C_ZN9QIODevice18setTextModeEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QIODevice::QIODevice(QObject * parent);
extern void* C_ZN9QIODeviceC2EP7QObject(void* arg0); // 3
  // proto:  void QIODevice::QIODevice();
extern void* C_ZN9QIODeviceC2Ev(); // 3
  // proto:  bool QIODevice::reset();
extern bool C_ZN9QIODevice5resetEv(void* qthis); // 4
  // proto:  OpenMode QIODevice::openMode();
extern void C_ZNK9QIODevice8openModeEv(void* qthis); // 4
  // proto:  bool QIODevice::isTextModeEnabled();
extern bool C_ZNK9QIODevice17isTextModeEnabledEv(void* qthis); // 4
  // proto:  const QMetaObject * QIODevice::metaObject();
extern void C_ZNK9QIODevice10metaObjectEv(void* qthis); // 4
  // proto:  bool QIODevice::isOpen();
extern bool C_ZNK9QIODevice6isOpenEv(void* qthis); // 4
  // proto:  bool QIODevice::atEnd();
extern bool C_ZNK9QIODevice5atEndEv(void* qthis); // 4
  // proto:  bool QIODevice::waitForReadyRead(int msecs);
extern bool C_ZN9QIODevice16waitForReadyReadEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QIODevice::isReadable();
extern bool C_ZNK9QIODevice10isReadableEv(void* qthis); // 4
  // proto:  QByteArray QIODevice::readLine(qint64 maxlen);
extern void* C_ZN9QIODevice8readLineEx(void* qthis, int64_t arg0); // 4
  // proto:  qint64 QIODevice::readLine(char * data, qint64 maxlen);
extern int64_t C_ZN9QIODevice8readLineEPcx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  bool QIODevice::isWritable();
extern bool C_ZNK9QIODevice10isWritableEv(void* qthis); // 4
  // proto:  bool QIODevice::getChar(char * c);
extern bool C_ZN9QIODevice7getCharEPc(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QIODevice)=1
type QIODevice struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _readyRead QIODevice_readyRead_signal;
//  _readChannelFinished QIODevice_readChannelFinished_signal;
//  _aboutToClose QIODevice_aboutToClose_signal;
//  _bytesWritten QIODevice_bytesWritten_signal;
}

// bytesAvailable()
func (this *QIODevice) BytesAvailable(args ...interface{}) (ret interface{}) {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice14bytesAvailableEv
    // invoke: qint64 bytesAvailable()
    var ret0 = C.C_ZNK9QIODevice14bytesAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "bytesAvailable", args)
  }

  return
}

// pos()
func (this *QIODevice) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice3posEv
    // invoke: qint64 pos()
    var ret0 = C.C_ZNK9QIODevice3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "pos", args)
  }

  return
}

// bytesToWrite()
func (this *QIODevice) BytesToWrite(args ...interface{}) (ret interface{}) {
  // bytesToWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12bytesToWriteEv
    // invoke: qint64 bytesToWrite()
    var ret0 = C.C_ZNK9QIODevice12bytesToWriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "bytesToWrite", args)
  }

  return
}

// close()
func (this *QIODevice) Close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5closeEv
    // invoke: void close()
    C.C_ZN9QIODevice5closeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "close", args)
  }

  return
}

// seek(qint64)
func (this *QIODevice) Seek(args ...interface{}) (ret interface{}) {
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
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice4seekEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "seek", args)
  }

  return
}

// isSequential()
func (this *QIODevice) IsSequential(args ...interface{}) (ret interface{}) {
  // isSequential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice12isSequentialEv
    // invoke: bool isSequential()
    var ret0 = C.C_ZNK9QIODevice12isSequentialEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "isSequential", args)
  }

  return
}

// ungetChar(char)
func (this *QIODevice) UngetChar(args ...interface{}) () {
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
    // invoke: void ungetChar(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN9QIODevice9ungetCharEc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIODevice", "ungetChar", args)
  }

  return
}

// size()
func (this *QIODevice) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice4sizeEv
    // invoke: qint64 size()
    var ret0 = C.C_ZNK9QIODevice4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "size", args)
  }

  return
}

// write(const char *, qint64)
func (this *QIODevice) Write(args ...interface{}) (ret interface{}) {
  // write(const char *, qint64)
  // write(const char *)
  // write(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5writeEPKcx
    // invoke: qint64 write(const char *, qint64)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int64_t(qtrt.PrimConv(args[1], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QIODevice5writeEPKcx(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QIODevice5writeEPKc
    // invoke: qint64 write(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN9QIODevice5writeEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN9QIODevice5writeERK10QByteArray
    // invoke: qint64 write(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice5writeERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "write", args)
  }

  return
}

// canReadLine()
func (this *QIODevice) CanReadLine(args ...interface{}) (ret interface{}) {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11canReadLineEv
    // invoke: bool canReadLine()
    var ret0 = C.C_ZNK9QIODevice11canReadLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "canReadLine", args)
  }

  return
}

// waitForBytesWritten(int)
func (this *QIODevice) WaitForBytesWritten(args ...interface{}) (ret interface{}) {
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
    // invoke: bool waitForBytesWritten(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice19waitForBytesWrittenEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "waitForBytesWritten", args)
  }

  return
}

// read(char *, qint64)
func (this *QIODevice) Read(args ...interface{}) (ret interface{}) {
  // read(char *, qint64)
  // read(qint64)
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
    // invoke: _ZN9QIODevice4readEPcx
    // invoke: qint64 read(char *, qint64)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int64_t(qtrt.PrimConv(args[1], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QIODevice4readEPcx(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QIODevice4readEx
    // invoke: QByteArray read(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice4readEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "read", args)
  }

  return
}

// peek(char *, qint64)
func (this *QIODevice) Peek(args ...interface{}) (ret interface{}) {
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
    // invoke: qint64 peek(char *, qint64)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int64_t(qtrt.PrimConv(args[1], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QIODevice4peekEPcx(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QIODevice4peekEx
    // invoke: QByteArray peek(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice4peekEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "peek", args)
  }

  return
}

// putChar(char)
func (this *QIODevice) PutChar(args ...interface{}) (ret interface{}) {
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
    // invoke: bool putChar(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice7putCharEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "putChar", args)
  }

  return
}

// errorString()
func (this *QIODevice) ErrorString(args ...interface{}) (ret interface{}) {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice11errorStringEv
    // invoke: QString errorString()
    var ret0 = C.C_ZNK9QIODevice11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "errorString", args)
  }

  return
}

// ~QIODevice()
func (this *QIODevice) Free(args ...interface{}) () {
  // ~QIODevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODeviceD0Ev
    // invoke: void ~QIODevice()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QIODeviceD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QIODevice", "~QIODevice", args)
  }

  return
}

// readAll()
func (this *QIODevice) ReadAll(args ...interface{}) (ret interface{}) {
  // readAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice7readAllEv
    // invoke: QByteArray readAll()
    var ret0 = C.C_ZN9QIODevice7readAllEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "readAll", args)
  }

  return
}

// setTextModeEnabled(_Bool)
func (this *QIODevice) SetTextModeEnabled(args ...interface{}) () {
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
    // invoke: void setTextModeEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QIODevice18setTextModeEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIODevice", "setTextModeEnabled", args)
  }

  return
}

// QIODevice(class QObject *)
func GcfreeQIODevice(this *QIODevice) {
  qtrt.UniverseFree(this)
}
func NewQIODevice(args ...interface{}) *QIODevice {
  // QIODevice(class QObject *)
  // QIODevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODeviceC1EP7QObject
    // invoke: void QIODevice(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QIODeviceC2EP7QObject(arg0)
    this := &QIODevice{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQIODevice)
    return this
  case 1:
    // invoke: _ZN9QIODeviceC1Ev
    // invoke: void QIODevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QIODeviceC2Ev()
    this := &QIODevice{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQIODevice)
    return this
  default:
    qtrt.ErrorResolve("QIODevice", "QIODevice", args)
  }

  return nil // QIODevice{Qclsinst:qthis}
}

// reset()
func (this *QIODevice) Reset(args ...interface{}) (ret interface{}) {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QIODevice5resetEv
    // invoke: bool reset()
    var ret0 = C.C_ZN9QIODevice5resetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "reset", args)
  }

  return
}

// openMode()
func (this *QIODevice) OpenMode(args ...interface{}) () {
  // openMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice8openModeEv
    // invoke: OpenMode openMode()
    C.C_ZNK9QIODevice8openModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "openMode", args)
  }

  return
}

// isTextModeEnabled()
func (this *QIODevice) IsTextModeEnabled(args ...interface{}) (ret interface{}) {
  // isTextModeEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice17isTextModeEnabledEv
    // invoke: bool isTextModeEnabled()
    var ret0 = C.C_ZNK9QIODevice17isTextModeEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "isTextModeEnabled", args)
  }

  return
}

// metaObject()
func (this *QIODevice) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QIODevice10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIODevice", "metaObject", args)
  }

  return
}

// isOpen()
func (this *QIODevice) IsOpen(args ...interface{}) (ret interface{}) {
  // isOpen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice6isOpenEv
    // invoke: bool isOpen()
    var ret0 = C.C_ZNK9QIODevice6isOpenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "isOpen", args)
  }

  return
}

// atEnd()
func (this *QIODevice) AtEnd(args ...interface{}) (ret interface{}) {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice5atEndEv
    // invoke: bool atEnd()
    var ret0 = C.C_ZNK9QIODevice5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "atEnd", args)
  }

  return
}

// waitForReadyRead(int)
func (this *QIODevice) WaitForReadyRead(args ...interface{}) (ret interface{}) {
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
    // invoke: bool waitForReadyRead(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice16waitForReadyReadEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "waitForReadyRead", args)
  }

  return
}

// isReadable()
func (this *QIODevice) IsReadable(args ...interface{}) (ret interface{}) {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isReadableEv
    // invoke: bool isReadable()
    var ret0 = C.C_ZNK9QIODevice10isReadableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "isReadable", args)
  }

  return
}

// readLine(qint64)
func (this *QIODevice) ReadLine(args ...interface{}) (ret interface{}) {
  // readLine(qint64)
  // readLine(char *, qint64)
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
    // invoke: _ZN9QIODevice8readLineEx
    // invoke: QByteArray readLine(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QIODevice8readLineEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QIODevice8readLineEPcx
    // invoke: qint64 readLine(char *, qint64)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int64_t(qtrt.PrimConv(args[1], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QIODevice8readLineEPcx(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "readLine", args)
  }

  return
}

// isWritable()
func (this *QIODevice) IsWritable(args ...interface{}) (ret interface{}) {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QIODevice10isWritableEv
    // invoke: bool isWritable()
    var ret0 = C.C_ZNK9QIODevice10isWritableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "isWritable", args)
  }

  return
}

// getChar(char *)
func (this *QIODevice) GetChar(args ...interface{}) (ret interface{}) {
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
    // invoke: bool getChar(char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN9QIODevice7getCharEPc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIODevice", "getChar", args)
  }

  return
}

// <= body block end

