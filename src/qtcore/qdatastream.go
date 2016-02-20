package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern int32_t C_ZN11QDataStream11skipRawDataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDataStream::setVersion(int );
extern void C_ZN11QDataStream10setVersionEi(void* qthis, int32_t arg0); // 2
  // proto:  QDataStream & QDataStream::writeBytes(const char * , uint len);
extern void* C_ZN11QDataStream10writeBytesEPKcj(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QDataStream::unsetDevice();
extern void C_ZN11QDataStream11unsetDeviceEv(void* qthis); // 4
  // proto:  int QDataStream::version();
extern int32_t C_ZNK11QDataStream7versionEv(void* qthis); // 2
  // proto:  QDataStream & QDataStream::readBytes(char *& , uint & len);
extern void* C_ZN11QDataStream9readBytesERPcRj(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QDataStream::setDevice(QIODevice * );
extern void C_ZN11QDataStream9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QDataStream::ByteOrder QDataStream::byteOrder();
extern void C_ZNK11QDataStream9byteOrderEv(void* qthis); // 2
  // proto:  QDataStream::Status QDataStream::status();
extern void C_ZNK11QDataStream6statusEv(void* qthis); // 4
  // proto:  int QDataStream::readRawData(char * , int len);
extern int32_t C_ZN11QDataStream11readRawDataEPci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QDataStream::FloatingPointPrecision QDataStream::floatingPointPrecision();
extern void C_ZNK11QDataStream22floatingPointPrecisionEv(void* qthis); // 4
  // proto:  int QDataStream::writeRawData(const char * , int len);
extern int32_t C_ZN11QDataStream12writeRawDataEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QIODevice * QDataStream::device();
extern void* C_ZNK11QDataStream6deviceEv(void* qthis); // 2
  // proto:  void QDataStream::resetStatus();
extern void C_ZN11QDataStream11resetStatusEv(void* qthis); // 4
  // proto:  bool QDataStream::atEnd();
extern bool C_ZNK11QDataStream5atEndEv(void* qthis); // 4
  // proto:  void QDataStream::QDataStream(const QByteArray & );
extern void* C_ZN11QDataStreamC2ERK10QByteArray(void* arg0); // 3
  // proto:  void QDataStream::QDataStream();
extern void* C_ZN11QDataStreamC2Ev(); // 3
  // proto:  void QDataStream::QDataStream(QIODevice * );
extern void* C_ZN11QDataStreamC2EP9QIODevice(void* arg0); // 3
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QDataStream()
func (this *QDataStream) Freeqdatastream(args ...interface{}) () {
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
    C.C_ZN11QDataStreamD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "~QDataStream", args)
  }

  return
}

// skipRawData(int)
func (this *QDataStream) Skiprawdata(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QDataStream11skipRawDataEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "skipRawData", args)
  }

  return
}

// setVersion(int)
func (this *QDataStream) Setversion(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QDataStream10setVersionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setVersion", args)
  }

  return
}

// writeBytes(const char *, uint)
func (this *QDataStream) Writebytes(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QDataStream10writeBytesEPKcj(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDataStream{}) // "QDataStream &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "writeBytes", args)
  }

  return
}

// unsetDevice()
func (this *QDataStream) Unsetdevice(args ...interface{}) () {
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
    C.C_ZN11QDataStream11unsetDeviceEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "unsetDevice", args)
  }

  return
}

// version()
func (this *QDataStream) Version(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDataStream7versionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "version", args)
  }

  return
}

// readBytes(char *&, uint &)
func (this *QDataStream) Readbytes(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QDataStream9readBytesERPcRj(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDataStream{}) // "QDataStream &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "readBytes", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QDataStream) Setdevice(args ...interface{}) () {
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
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QDataStream9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setDevice", args)
  }

  return
}

// byteOrder()
func (this *QDataStream) Byteorder(args ...interface{}) () {
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
    C.C_ZNK11QDataStream9byteOrderEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "byteOrder", args)
  }

  return
}

// status()
func (this *QDataStream) Status(args ...interface{}) () {
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
    C.C_ZNK11QDataStream6statusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "status", args)
  }

  return
}

// readRawData(char *, int)
func (this *QDataStream) Readrawdata(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QDataStream11readRawDataEPci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "readRawData", args)
  }

  return
}

// floatingPointPrecision()
func (this *QDataStream) Floatingpointprecision(args ...interface{}) () {
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
    C.C_ZNK11QDataStream22floatingPointPrecisionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "floatingPointPrecision", args)
  }

  return
}

// writeRawData(const char *, int)
func (this *QDataStream) Writerawdata(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QDataStream12writeRawDataEPKci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "writeRawData", args)
  }

  return
}

// device()
func (this *QDataStream) Device(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDataStream6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "device", args)
  }

  return
}

// resetStatus()
func (this *QDataStream) Resetstatus(args ...interface{}) () {
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
    C.C_ZN11QDataStream11resetStatusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "resetStatus", args)
  }

  return
}

// atEnd()
func (this *QDataStream) Atend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QDataStream5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDataStream", "atEnd", args)
  }

  return
}

// QDataStream(const class QByteArray &)
func NewQDataStream(args ...interface{}) *QDataStream {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QDataStreamC2ERK10QByteArray(arg0)
    return &QDataStream{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QDataStreamC1Ev
    // invoke: void QDataStream()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QDataStreamC2Ev()
    return &QDataStream{Qclsinst:qthis}
  case 2:
    // invoke: _ZN11QDataStreamC1EP9QIODevice
    // invoke: void QDataStream(class QIODevice *)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QDataStreamC2EP9QIODevice(arg0)
    return &QDataStream{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDataStream", "QDataStream", args)
  }

  return nil // QDataStream{Qclsinst:qthis}
}

// <= body block end

