package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QIODevice::ungetChar(char c);
extern void _ZN9QIODevice9ungetCharEc(void* qthis, char arg0);
  // proto:  QString QIODevice::errorString();
extern void _ZNK9QIODevice11errorStringEv(void* qthis);
  // proto:  qint64 QIODevice::write(const QByteArray & data);
extern void demth_ZN9QIODevice5writeERK10QByteArray(void* qthis, void* arg0);
  // proto:  qint64 QIODevice::write(const char * data);
extern void _ZN9QIODevice5writeEPKc(void* qthis, char* arg0);
  // proto:  bool QIODevice::isReadable();
extern void _ZNK9QIODevice10isReadableEv(void* qthis);
  // proto:  qint64 QIODevice::readLine(char * data, qint64 maxlen);
extern void _ZN9QIODevice8readLineEPcx(void* qthis, char* arg0, long long arg1);
  // proto:  QByteArray QIODevice::readLine(qint64 maxlen);
extern void _ZN9QIODevice8readLineEx(void* qthis, long long arg0);
  // proto:  bool QIODevice::waitForReadyRead(int msecs);
extern void _ZN9QIODevice16waitForReadyReadEi(void* qthis, int arg0);
  // proto:  qint64 QIODevice::size();
extern void _ZNK9QIODevice4sizeEv(void* qthis);
  // proto:  bool QIODevice::getChar(char * c);
extern void _ZN9QIODevice7getCharEPc(void* qthis, char* arg0);
  // proto:  bool QIODevice::putChar(char c);
extern void _ZN9QIODevice7putCharEc(void* qthis, char arg0);
  // proto:  bool QIODevice::isTextModeEnabled();
extern void _ZNK9QIODevice17isTextModeEnabledEv(void* qthis);
  // proto:  bool QIODevice::isSequential();
extern void _ZNK9QIODevice12isSequentialEv(void* qthis);
  // proto:  qint64 QIODevice::bytesAvailable();
extern void _ZNK9QIODevice14bytesAvailableEv(void* qthis);
  // proto:  qint64 QIODevice::write(const char * data, qint64 len);
extern void _ZN9QIODevice5writeEPKcx(void* qthis, char* arg0, long long arg1);
  // proto:  void QIODevice::close();
extern void _ZN9QIODevice5closeEv(void* qthis);
  // proto:  QByteArray QIODevice::readAll();
extern void _ZN9QIODevice7readAllEv(void* qthis);
  // proto:  bool QIODevice::atEnd();
extern void _ZNK9QIODevice5atEndEv(void* qthis);
  // proto:  bool QIODevice::seek(qint64 pos);
extern void _ZN9QIODevice4seekEx(void* qthis, long long arg0);
  // proto:  void QIODevice::QIODevice(const QIODevice & );
extern void* dector_ZN9QIODeviceC1ERKS_(void* arg0);
extern void _ZN9QIODeviceC1ERKS_(void* qthis, void* arg0);
  // proto:  qint64 QIODevice::pos();
extern void _ZNK9QIODevice3posEv(void* qthis);
  // proto:  QByteArray QIODevice::read(qint64 maxlen);
extern void _ZN9QIODevice4readEx(void* qthis, long long arg0);
  // proto:  qint64 QIODevice::peek(char * data, qint64 maxlen);
extern void _ZN9QIODevice4peekEPcx(void* qthis, char* arg0, long long arg1);
  // proto:  qint64 QIODevice::read(char * data, qint64 maxlen);
extern void _ZN9QIODevice4readEPcx(void* qthis, char* arg0, long long arg1);
  // proto:  bool QIODevice::waitForBytesWritten(int msecs);
extern void _ZN9QIODevice19waitForBytesWrittenEi(void* qthis, int arg0);
  // proto:  qint64 QIODevice::bytesToWrite();
extern void _ZNK9QIODevice12bytesToWriteEv(void* qthis);
  // proto:  bool QIODevice::reset();
extern void _ZN9QIODevice5resetEv(void* qthis);
  // proto:  bool QIODevice::isWritable();
extern void _ZNK9QIODevice10isWritableEv(void* qthis);
  // proto:  QByteArray QIODevice::peek(qint64 maxlen);
extern void _ZN9QIODevice4peekEx(void* qthis, long long arg0);
  // proto:  void QIODevice::QIODevice(QObject * parent);
extern void* dector_ZN9QIODeviceC1EP7QObject(void* arg0);
extern void _ZN9QIODeviceC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QIODevice::metaObject();
extern void _ZNK9QIODevice10metaObjectEv(void* qthis);
  // proto:  void QIODevice::setTextModeEnabled(bool enabled);
extern void _ZN9QIODevice18setTextModeEnabledEb(void* qthis, bool arg0);
  // proto:  void QIODevice::QIODevice();
extern void* dector_ZN9QIODeviceC1Ev();
extern void _ZN9QIODeviceC1Ev(void* qthis);
  // proto:  bool QIODevice::isOpen();
extern void _ZNK9QIODevice6isOpenEv(void* qthis);
  // proto:  bool QIODevice::canReadLine();
extern void _ZNK9QIODevice11canReadLineEv(void* qthis);
  // proto:  void QIODevice::~QIODevice();
extern void _ZN9QIODeviceD0Ev(void* qthis);
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

// class sizeof(QIODevice)=1
type QIODevice struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _readyRead QIODevice_readyRead_signal;
//  _readChannelFinished QIODevice_readChannelFinished_signal;
//  _aboutToClose QIODevice_aboutToClose_signal;
//  _bytesWritten QIODevice_bytesWritten_signal;
}

  // proto:  void QIODevice::ungetChar(char c);
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "ungetChar", args)
  }

}

  // proto:  QString QIODevice::errorString();
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

  // proto:  qint64 QIODevice::write(const QByteArray & data);
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN9QIODevice5writeEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZN9QIODevice5writeEPKcx
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QIODevice", "write", args)
  }

}

  // proto:  bool QIODevice::isReadable();
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

  // proto:  qint64 QIODevice::readLine(char * data, qint64 maxlen);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN9QIODevice8readLineEx
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "readLine", args)
  }

}

  // proto:  bool QIODevice::waitForReadyRead(int msecs);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "waitForReadyRead", args)
  }

}

  // proto:  qint64 QIODevice::size();
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

  // proto:  bool QIODevice::getChar(char * c);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "getChar", args)
  }

}

  // proto:  bool QIODevice::putChar(char c);
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "putChar", args)
  }

}

  // proto:  bool QIODevice::isTextModeEnabled();
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

  // proto:  bool QIODevice::isSequential();
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

  // proto:  qint64 QIODevice::bytesAvailable();
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

  // proto:  void QIODevice::close();
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

  // proto:  QByteArray QIODevice::readAll();
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

  // proto:  bool QIODevice::atEnd();
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

  // proto:  bool QIODevice::seek(qint64 pos);
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "seek", args)
  }

}

  // proto:  void QIODevice::QIODevice(const QIODevice & );
func NewQIODevice(args ...interface{}) QIODevice {
  return QIODevice{}
}

  // proto:  qint64 QIODevice::pos();
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

  // proto:  QByteArray QIODevice::read(qint64 maxlen);
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN9QIODevice4readEPcx
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QIODevice", "read", args)
  }

}

  // proto:  qint64 QIODevice::peek(char * data, qint64 maxlen);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN9QIODevice4peekEx
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "peek", args)
  }

}

  // proto:  bool QIODevice::waitForBytesWritten(int msecs);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "waitForBytesWritten", args)
  }

}

  // proto:  qint64 QIODevice::bytesToWrite();
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

  // proto:  bool QIODevice::reset();
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

  // proto:  bool QIODevice::isWritable();
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

  // proto:  const QMetaObject * QIODevice::metaObject();
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

  // proto:  void QIODevice::setTextModeEnabled(bool enabled);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QIODevice", "setTextModeEnabled", args)
  }

}

  // proto:  bool QIODevice::isOpen();
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

  // proto:  bool QIODevice::canReadLine();
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

  // proto:  void QIODevice::~QIODevice();
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

