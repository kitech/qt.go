package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qopenglbuffer.h
// dst-file: /src/gui/qopenglbuffer.go
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
  // proto:  bool QOpenGLBuffer::isCreated();
extern void C_ZNK13QOpenGLBuffer9isCreatedEv(void* qthis); // 4
  // proto:  bool QOpenGLBuffer::unmap();
extern void C_ZN13QOpenGLBuffer5unmapEv(void* qthis); // 4
  // proto:  GLuint QOpenGLBuffer::bufferId();
extern void C_ZNK13QOpenGLBuffer8bufferIdEv(void* qthis); // 4
  // proto:  bool QOpenGLBuffer::read(int offset, void * data, int count);
extern void C_ZN13QOpenGLBuffer4readEiPvi(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QOpenGLBuffer::bind();
extern void C_ZN13QOpenGLBuffer4bindEv(void* qthis); // 4
  // proto:  bool QOpenGLBuffer::create();
extern void C_ZN13QOpenGLBuffer6createEv(void* qthis); // 4
  // proto:  void QOpenGLBuffer::write(int offset, const void * data, int count);
extern void C_ZN13QOpenGLBuffer5writeEiPKvi(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  QOpenGLBuffer::UsagePattern QOpenGLBuffer::usagePattern();
extern void C_ZNK13QOpenGLBuffer12usagePatternEv(void* qthis); // 4
  // proto:  void QOpenGLBuffer::allocate(int count);
extern void C_ZN13QOpenGLBuffer8allocateEi(void* qthis, int32_t arg0); // 2
  // proto:  void QOpenGLBuffer::allocate(const void * data, int count);
extern void C_ZN13QOpenGLBuffer8allocateEPKvi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QOpenGLBuffer::release();
extern void C_ZN13QOpenGLBuffer7releaseEv(void* qthis); // 4
  // proto:  void QOpenGLBuffer::destroy();
extern void C_ZN13QOpenGLBuffer7destroyEv(void* qthis); // 4
  // proto:  int QOpenGLBuffer::size();
extern void C_ZNK13QOpenGLBuffer4sizeEv(void* qthis); // 4
  // proto:  QOpenGLBuffer::Type QOpenGLBuffer::type();
extern void C_ZNK13QOpenGLBuffer4typeEv(void* qthis); // 4
  // proto:  void QOpenGLBuffer::~QOpenGLBuffer();
extern void C_ZN13QOpenGLBufferD2Ev(void* qthis); // 4
  // proto:  void QOpenGLBuffer::QOpenGLBuffer();
extern void* C_ZN13QOpenGLBufferC2Ev(); // 3
  // proto:  void QOpenGLBuffer::QOpenGLBuffer(const QOpenGLBuffer & other);
extern void* C_ZN13QOpenGLBufferC2ERKS_(void* arg0); // 3
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

// class sizeof(QOpenGLBuffer)=8
type QOpenGLBuffer struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isCreated()
func (this *QOpenGLBuffer) isCreated(args ...interface{}) () {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLBuffer9isCreatedEv
    // invoke: bool isCreated()
    var ret = C.C_ZNK13QOpenGLBuffer9isCreatedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "isCreated", args)
  }

}

// unmap()
func (this *QOpenGLBuffer) unmap(args ...interface{}) () {
  // unmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer5unmapEv
    // invoke: bool unmap()
    var ret = C.C_ZN13QOpenGLBuffer5unmapEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "unmap", args)
  }

}

// bufferId()
func (this *QOpenGLBuffer) bufferId(args ...interface{}) () {
  // bufferId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLBuffer8bufferIdEv
    // invoke: GLuint bufferId()
    var ret = C.C_ZNK13QOpenGLBuffer8bufferIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "bufferId", args)
  }

}

// read(int, void *, int)
func (this *QOpenGLBuffer) read(args ...interface{}) () {
  // read(int, void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer4readEiPvi
    // invoke: bool read(int, void *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN13QOpenGLBuffer4readEiPvi(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "read", args)
  }

}

// bind()
func (this *QOpenGLBuffer) bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer4bindEv
    // invoke: bool bind()
    var ret = C.C_ZN13QOpenGLBuffer4bindEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "bind", args)
  }

}

// create()
func (this *QOpenGLBuffer) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer6createEv
    // invoke: bool create()
    var ret = C.C_ZN13QOpenGLBuffer6createEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "create", args)
  }

}

// write(int, const void *, int)
func (this *QOpenGLBuffer) write(args ...interface{}) () {
  // write(int, const void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer5writeEiPKvi
    // invoke: void write(int, const void *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN13QOpenGLBuffer5writeEiPKvi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "write", args)
  }

}

// usagePattern()
func (this *QOpenGLBuffer) usagePattern(args ...interface{}) () {
  // usagePattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLBuffer12usagePatternEv
    // invoke: QOpenGLBuffer::UsagePattern usagePattern()
    C.C_ZNK13QOpenGLBuffer12usagePatternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "usagePattern", args)
  }

}

// allocate(int)
func (this *QOpenGLBuffer) allocate(args ...interface{}) () {
  // allocate(int)
  // allocate(const void *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "const void *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer8allocateEi
    // invoke: void allocate(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QOpenGLBuffer8allocateEi(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QOpenGLBuffer8allocateEPKvi
    // invoke: void allocate(const void *, int)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QOpenGLBuffer8allocateEPKvi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "allocate", args)
  }

}

// release()
func (this *QOpenGLBuffer) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer7releaseEv
    // invoke: void release()
    C.C_ZN13QOpenGLBuffer7releaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "release", args)
  }

}

// destroy()
func (this *QOpenGLBuffer) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer7destroyEv
    // invoke: void destroy()
    C.C_ZN13QOpenGLBuffer7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "destroy", args)
  }

}

// size()
func (this *QOpenGLBuffer) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLBuffer4sizeEv
    // invoke: int size()
    var ret = C.C_ZNK13QOpenGLBuffer4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "size", args)
  }

}

// type()
func (this *QOpenGLBuffer) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLBuffer4typeEv
    // invoke: QOpenGLBuffer::Type type()
    C.C_ZNK13QOpenGLBuffer4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "type", args)
  }

}

// ~QOpenGLBuffer()
func (this *QOpenGLBuffer) FreeQOpenGLBuffer(args ...interface{}) () {
  // ~QOpenGLBuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBufferD0Ev
    // invoke: void ~QOpenGLBuffer()
    C.C_ZN13QOpenGLBufferD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "~QOpenGLBuffer", args)
  }

}

// QOpenGLBuffer()
func NewQOpenGLBuffer(args ...interface{}) *QOpenGLBuffer {
  // QOpenGLBuffer()
  // QOpenGLBuffer(const class QOpenGLBuffer &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QOpenGLBuffer{}) // "const QOpenGLBuffer &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBufferC1Ev
    // invoke: void QOpenGLBuffer()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QOpenGLBufferC2Ev()
    return &QOpenGLBuffer{qclsinst:qthis}
  case 1:
    // invoke: _ZN13QOpenGLBufferC1ERKS_
    // invoke: void QOpenGLBuffer(const class QOpenGLBuffer &)
    var arg0 = args[0].(QOpenGLBuffer).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QOpenGLBufferC2ERKS_(arg0)
    return &QOpenGLBuffer{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "QOpenGLBuffer", args)
  }

  return nil // QOpenGLBuffer{qclsinst:qthis}
}

// <= body block end

