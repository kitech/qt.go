package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qbuffer.h
// dst-file: /src/core/qbuffer.go
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
  // proto:  qint64 QBuffer::pos();
extern int64_t C_ZNK7QBuffer3posEv(void* qthis); // 4
  // proto:  void QBuffer::close();
extern void C_ZN7QBuffer5closeEv(void* qthis); // 4
  // proto:  qint64 QBuffer::size();
extern int64_t C_ZNK7QBuffer4sizeEv(void* qthis); // 4
  // proto:  bool QBuffer::canReadLine();
extern bool C_ZNK7QBuffer11canReadLineEv(void* qthis); // 4
  // proto:  void QBuffer::setBuffer(QByteArray * a);
extern void C_ZN7QBuffer9setBufferEP10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QBuffer::QBuffer(QByteArray * buf, QObject * parent);
extern void* C_ZN7QBufferC2EP10QByteArrayP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QBuffer::QBuffer(QObject * parent);
extern void* C_ZN7QBufferC2EP7QObject(void* arg0); // 3
  // proto:  void QBuffer::setData(const char * data, int len);
extern void C_ZN7QBuffer7setDataEPKci(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QBuffer::setData(const QByteArray & data);
extern void C_ZN7QBuffer7setDataERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QBuffer::~QBuffer();
extern void C_ZN7QBufferD2Ev(void* qthis); // 4
  // proto:  QByteArray & QBuffer::buffer();
extern void* C_ZN7QBuffer6bufferEv(void* qthis); // 4
  // proto:  const QByteArray & QBuffer::data();
extern void* C_ZNK7QBuffer4dataEv(void* qthis); // 4
  // proto:  const QMetaObject * QBuffer::metaObject();
extern void C_ZNK7QBuffer10metaObjectEv(void* qthis); // 4
  // proto:  bool QBuffer::seek(qint64 off);
extern bool C_ZN7QBuffer4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QBuffer::atEnd();
extern bool C_ZNK7QBuffer5atEndEv(void* qthis); // 4
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

// class sizeof(QBuffer)=1
type QBuffer struct {
  /*qbase*/ QIODevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// pos()
func (this *QBuffer) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer3posEv
    // invoke: qint64 pos()
    var ret0 = C.C_ZNK7QBuffer3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "pos", args)
  }

  return
}

// close()
func (this *QBuffer) Close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer5closeEv
    // invoke: void close()
    C.C_ZN7QBuffer5closeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "close", args)
  }

  return
}

// size()
func (this *QBuffer) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4sizeEv
    // invoke: qint64 size()
    var ret0 = C.C_ZNK7QBuffer4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "size", args)
  }

  return
}

// canReadLine()
func (this *QBuffer) CanReadLine(args ...interface{}) (ret interface{}) {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer11canReadLineEv
    // invoke: bool canReadLine()
    var ret0 = C.C_ZNK7QBuffer11canReadLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "canReadLine", args)
  }

  return
}

// setBuffer(class QByteArray *)
func (this *QBuffer) SetBuffer(args ...interface{}) () {
  // setBuffer(class QByteArray *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer9setBufferEP10QByteArray
    // invoke: void setBuffer(class QByteArray *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QBuffer9setBufferEP10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "setBuffer", args)
  }

  return
}

// QBuffer(class QByteArray *, class QObject *)
func GcfreeQBuffer(this *QBuffer) {
  qtrt.UniverseFree(this)
}
func NewQBuffer(args ...interface{}) *QBuffer {
  // QBuffer(class QByteArray *, class QObject *)
  // QBuffer(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray *"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBufferC1EP10QByteArrayP7QObject
    // invoke: void QBuffer(class QByteArray *, class QObject *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBufferC2EP10QByteArrayP7QObject(arg0, arg1)
    this := &QBuffer{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQBuffer)
    return this
  case 1:
    // invoke: _ZN7QBufferC1EP7QObject
    // invoke: void QBuffer(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBufferC2EP7QObject(arg0)
    this := &QBuffer{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQBuffer)
    return this
  default:
    qtrt.ErrorResolve("QBuffer", "QBuffer", args)
  }

  return nil // QBuffer{Qclsinst:qthis}
}

// setData(const char *, int)
func (this *QBuffer) SetData(args ...interface{}) () {
  // setData(const char *, int)
  // setData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer7setDataEPKci
    // invoke: void setData(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QBuffer7setDataEPKci(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QBuffer7setDataERK10QByteArray
    // invoke: void setData(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QBuffer7setDataERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "setData", args)
  }

  return
}

// ~QBuffer()
func (this *QBuffer) Free(args ...interface{}) () {
  // ~QBuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBufferD0Ev
    // invoke: void ~QBuffer()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN7QBufferD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QBuffer", "~QBuffer", args)
  }

  return
}

// buffer()
func (this *QBuffer) Buffer(args ...interface{}) (ret interface{}) {
  // buffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer6bufferEv
    // invoke: QByteArray & buffer()
    var ret0 = C.C_ZN7QBuffer6bufferEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "buffer", args)
  }

  return
}

// data()
func (this *QBuffer) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4dataEv
    // invoke: const QByteArray & data()
    var ret0 = C.C_ZNK7QBuffer4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "data", args)
  }

  return
}

// metaObject()
func (this *QBuffer) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QBuffer10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "metaObject", args)
  }

  return
}

// seek(qint64)
func (this *QBuffer) Seek(args ...interface{}) (ret interface{}) {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer4seekEx
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QBuffer4seekEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "seek", args)
  }

  return
}

// atEnd()
func (this *QBuffer) AtEnd(args ...interface{}) (ret interface{}) {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer5atEndEv
    // invoke: bool atEnd()
    var ret0 = C.C_ZNK7QBuffer5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QBuffer", "atEnd", args)
  }

  return
}

// <= body block end

