package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qdatastream.h
// dst-file: /src/core/qdatastream.go
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
  // proto:  void QDataStream::~QDataStream();
extern void C_ZN11QDataStreamD2Ev(void* qthis); // 4
  // proto:  int QDataStream::skipRawData(int len);
extern void C_ZN11QDataStream11skipRawDataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDataStream::setVersion(int );
extern void C_ZN11QDataStream10setVersionEi(void* qthis, int32_t arg0); // 2
  // proto:  QDataStream & QDataStream::writeBytes(const char * , uint len);
extern void C_ZN11QDataStream10writeBytesEPKcj(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  void QDataStream::unsetDevice();
extern void C_ZN11QDataStream11unsetDeviceEv(void* qthis); // 4
  // proto:  int QDataStream::version();
extern void C_ZNK11QDataStream7versionEv(void* qthis); // 2
  // proto:  QDataStream & QDataStream::readBytes(char *& , uint & len);
extern void C_ZN11QDataStream9readBytesERPcRj(void* qthis, unsigned char* arg0, int32_t* arg1); // 4
  // proto:  void QDataStream::setDevice(QIODevice * );
extern void C_ZN11QDataStream9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QDataStream::ByteOrder QDataStream::byteOrder();
extern void C_ZNK11QDataStream9byteOrderEv(void* qthis); // 2
  // proto:  QDataStream::Status QDataStream::status();
extern void C_ZNK11QDataStream6statusEv(void* qthis); // 4
  // proto:  int QDataStream::readRawData(char * , int len);
extern void C_ZN11QDataStream11readRawDataEPci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  QDataStream::FloatingPointPrecision QDataStream::floatingPointPrecision();
extern void C_ZNK11QDataStream22floatingPointPrecisionEv(void* qthis); // 4
  // proto:  int QDataStream::writeRawData(const char * , int len);
extern void C_ZN11QDataStream12writeRawDataEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  QIODevice * QDataStream::device();
extern void C_ZNK11QDataStream6deviceEv(void* qthis); // 2
  // proto:  void QDataStream::resetStatus();
extern void C_ZN11QDataStream11resetStatusEv(void* qthis); // 4
  // proto:  bool QDataStream::atEnd();
extern void C_ZNK11QDataStream5atEndEv(void* qthis); // 4
  // proto:  void QDataStream::QDataStream(const QByteArray & );
extern void C_ZN11QDataStreamC2ERK10QByteArray(void* qthis, void* arg0); // 3
  // proto:  void QDataStream::QDataStream();
extern void C_ZN11QDataStreamC2Ev(void* qthis); // 3
  // proto:  void QDataStream::QDataStream(QIODevice * );
extern void C_ZN11QDataStreamC2EP9QIODevice(void* qthis, void* arg0); // 3
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

// class sizeof(QDataStream)=1
type QDataStream struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QDataStream()
func (this *QDataStream) FreeQDataStream(args ...interface{}) () {
  // ~QDataStream()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStreamD0Ev
    // invoke: void ~QDataStream()
    C.C_ZN11QDataStreamD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "~QDataStream", args)
  }

}

// skipRawData(int)
func (this *QDataStream) skipRawData(args ...interface{}) () {
  // skipRawData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11skipRawDataEi
    // invoke: int skipRawData(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QDataStream11skipRawDataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "skipRawData", args)
  }

}

// setVersion(int)
func (this *QDataStream) setVersion(args ...interface{}) () {
  // setVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream10setVersionEi
    // invoke: void setVersion(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QDataStream10setVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setVersion", args)
  }

}

// writeBytes(const char *, uint)
func (this *QDataStream) writeBytes(args ...interface{}) () {
  // writeBytes(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream10writeBytesEPKcj
    // invoke: QDataStream & writeBytes(const char *, uint)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QDataStream10writeBytesEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "writeBytes", args)
  }

}

// unsetDevice()
func (this *QDataStream) unsetDevice(args ...interface{}) () {
  // unsetDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11unsetDeviceEv
    // invoke: void unsetDevice()
    C.C_ZN11QDataStream11unsetDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "unsetDevice", args)
  }

}

// version()
func (this *QDataStream) version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream7versionEv
    // invoke: int version()
    C.C_ZNK11QDataStream7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "version", args)
  }

}

// readBytes(char *&, uint &)
func (this *QDataStream) readBytes(args ...interface{}) () {
  // readBytes(char *&, uint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.StringTy(false) // "char *&"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream9readBytesERPcRj
    // invoke: QDataStream & readBytes(char *&, uint &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QDataStream9readBytesERPcRj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "readBytes", args)
  }

}

// setDevice(class QIODevice *)
func (this *QDataStream) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QDataStream9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setDevice", args)
  }

}

// byteOrder()
func (this *QDataStream) byteOrder(args ...interface{}) () {
  // byteOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream9byteOrderEv
    // invoke: QDataStream::ByteOrder byteOrder()
    C.C_ZNK11QDataStream9byteOrderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "byteOrder", args)
  }

}

// status()
func (this *QDataStream) status(args ...interface{}) () {
  // status()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream6statusEv
    // invoke: QDataStream::Status status()
    C.C_ZNK11QDataStream6statusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "status", args)
  }

}

// readRawData(char *, int)
func (this *QDataStream) readRawData(args ...interface{}) () {
  // readRawData(char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11readRawDataEPci
    // invoke: int readRawData(char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QDataStream11readRawDataEPci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "readRawData", args)
  }

}

// floatingPointPrecision()
func (this *QDataStream) floatingPointPrecision(args ...interface{}) () {
  // floatingPointPrecision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream22floatingPointPrecisionEv
    // invoke: QDataStream::FloatingPointPrecision floatingPointPrecision()
    C.C_ZNK11QDataStream22floatingPointPrecisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "floatingPointPrecision", args)
  }

}

// writeRawData(const char *, int)
func (this *QDataStream) writeRawData(args ...interface{}) () {
  // writeRawData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream12writeRawDataEPKci
    // invoke: int writeRawData(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QDataStream12writeRawDataEPKci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "writeRawData", args)
  }

}

// device()
func (this *QDataStream) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream6deviceEv
    // invoke: QIODevice * device()
    C.C_ZNK11QDataStream6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "device", args)
  }

}

// resetStatus()
func (this *QDataStream) resetStatus(args ...interface{}) () {
  // resetStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11resetStatusEv
    // invoke: void resetStatus()
    C.C_ZN11QDataStream11resetStatusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "resetStatus", args)
  }

}

// atEnd()
func (this *QDataStream) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream5atEndEv
    // invoke: bool atEnd()
    C.C_ZNK11QDataStream5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "atEnd", args)
  }

}

// QDataStream(const class QByteArray &)
func NewQDataStream(args ...interface{}) QDataStream {
  // QDataStream(const class QByteArray &)
  // QDataStream()
  // QDataStream(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStreamC1ERK10QByteArray
    // invoke: void QDataStream(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QDataStreamC2ERK10QByteArray(qthis, arg0)
  case 1:
    // invoke: _ZN11QDataStreamC1Ev
    // invoke: void QDataStream()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QDataStreamC2Ev(qthis)
  case 2:
    // invoke: _ZN11QDataStreamC1EP9QIODevice
    // invoke: void QDataStream(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QDataStreamC2EP9QIODevice(qthis, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "QDataStream", args)
  }

  return QDataStream{}
}

// <= body block end

