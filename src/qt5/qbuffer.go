package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
extern void _ZNK7QBuffer3posEv(void* qthis); // 4
  // proto:  void QBuffer::close();
extern void _ZN7QBuffer5closeEv(void* qthis); // 4
  // proto:  qint64 QBuffer::size();
extern void _ZNK7QBuffer4sizeEv(void* qthis); // 4
  // proto:  bool QBuffer::canReadLine();
extern void _ZNK7QBuffer11canReadLineEv(void* qthis); // 4
  // proto:  void QBuffer::setBuffer(QByteArray * a);
extern void _ZN7QBuffer9setBufferEP10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QBuffer::QBuffer(QByteArray * buf, QObject * parent);
extern void _ZN7QBufferC2EP10QByteArrayP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QBuffer::QBuffer(QObject * parent);
extern void _ZN7QBufferC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QBuffer::setData(const char * data, int len);
extern void _ZN7QBuffer7setDataEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 2
  // proto:  void QBuffer::setData(const QByteArray & data);
extern void _ZN7QBuffer7setDataERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QBuffer::~QBuffer();
extern void _ZN7QBufferD2Ev(void* qthis); // 4
  // proto:  QByteArray & QBuffer::buffer();
extern void _ZN7QBuffer6bufferEv(void* qthis); // 4
  // proto:  const QByteArray & QBuffer::data();
extern void _ZNK7QBuffer4dataEv(void* qthis); // 4
  // proto:  const QMetaObject * QBuffer::metaObject();
extern void _ZNK7QBuffer10metaObjectEv(void* qthis); // 4
  // proto:  bool QBuffer::seek(qint64 off);
extern void _ZN7QBuffer4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QBuffer::atEnd();
extern void _ZNK7QBuffer5atEndEv(void* qthis); // 4
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

// class sizeof(QBuffer)=1
type QBuffer struct {
  /*qbase*/ QIODevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

// pos()
func (this *QBuffer) pos(args ...interface{}) () {
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
    C._ZNK7QBuffer3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "pos", args)
  }

}

// close()
func (this *QBuffer) close(args ...interface{}) () {
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
    C._ZN7QBuffer5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "close", args)
  }

}

// size()
func (this *QBuffer) size(args ...interface{}) () {
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
    C._ZNK7QBuffer4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "size", args)
  }

}

// canReadLine()
func (this *QBuffer) canReadLine(args ...interface{}) () {
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
    C._ZNK7QBuffer11canReadLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "canReadLine", args)
  }

}

// setBuffer(class QByteArray *)
func (this *QBuffer) setBuffer(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QBuffer9setBufferEP10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "setBuffer", args)
  }

}

// QBuffer(class QByteArray *, class QObject *)
func NewQBuffer(args ...interface{}) QBuffer {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QBufferC2EP10QByteArrayP7QObject(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN7QBufferC1EP7QObject
    // invoke: void QBuffer(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QBufferC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "QBuffer", args)
  }

  return QBuffer{}
}

// setData(const char *, int)
func (this *QBuffer) setData(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QBuffer7setDataEPKci(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QBuffer7setDataERK10QByteArray
    // invoke: void setData(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QBuffer7setDataERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "setData", args)
  }

}

// ~QBuffer()
func (this *QBuffer) FreeQBuffer(args ...interface{}) () {
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
    C._ZN7QBufferD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "~QBuffer", args)
  }

}

// buffer()
func (this *QBuffer) buffer(args ...interface{}) () {
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
    C._ZN7QBuffer6bufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "buffer", args)
  }

}

// data()
func (this *QBuffer) data(args ...interface{}) () {
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
    C._ZNK7QBuffer4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "data", args)
  }

}

// metaObject()
func (this *QBuffer) metaObject(args ...interface{}) () {
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
    C._ZNK7QBuffer10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "metaObject", args)
  }

}

// seek(qint64)
func (this *QBuffer) seek(args ...interface{}) () {
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN7QBuffer4seekEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBuffer", "seek", args)
  }

}

// atEnd()
func (this *QBuffer) atEnd(args ...interface{}) () {
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
    C._ZNK7QBuffer5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBuffer", "atEnd", args)
  }

}

// <= body block end

