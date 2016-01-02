package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  bool QOpenGLBuffer::read(int offset, void * data, int count);
extern void _ZN13QOpenGLBuffer4readEiPvi(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  bool QOpenGLBuffer::bind();
extern void _ZN13QOpenGLBuffer4bindEv(void* qthis);
  // proto:  void QOpenGLBuffer::destroy();
extern void _ZN13QOpenGLBuffer7destroyEv(void* qthis);
  // proto:  void QOpenGLBuffer::allocate(int count);
extern void demth_ZN13QOpenGLBuffer8allocateEi(void* qthis, int arg0);
  // proto:  bool QOpenGLBuffer::unmap();
extern void _ZN13QOpenGLBuffer5unmapEv(void* qthis);
  // proto:  void QOpenGLBuffer::QOpenGLBuffer(const QOpenGLBuffer & other);
extern void* dector_ZN13QOpenGLBufferC1ERKS_(void* arg0);
extern void _ZN13QOpenGLBufferC1ERKS_(void* qthis, void* arg0);
  // proto:  int QOpenGLBuffer::size();
extern void _ZNK13QOpenGLBuffer4sizeEv(void* qthis);
  // proto:  void QOpenGLBuffer::allocate(const void * data, int count);
extern void _ZN13QOpenGLBuffer8allocateEPKvi(void* qthis, void* arg0, int arg1);
  // proto:  GLuint QOpenGLBuffer::bufferId();
extern void _ZNK13QOpenGLBuffer8bufferIdEv(void* qthis);
  // proto:  void QOpenGLBuffer::QOpenGLBuffer();
extern void* dector_ZN13QOpenGLBufferC1Ev();
extern void _ZN13QOpenGLBufferC1Ev(void* qthis);
  // proto:  bool QOpenGLBuffer::create();
extern void _ZN13QOpenGLBuffer6createEv(void* qthis);
  // proto:  void QOpenGLBuffer::~QOpenGLBuffer();
extern void _ZN13QOpenGLBufferD0Ev(void* qthis);
  // proto:  void QOpenGLBuffer::release();
extern void _ZN13QOpenGLBuffer7releaseEv(void* qthis);
  // proto:  bool QOpenGLBuffer::isCreated();
extern void _ZNK13QOpenGLBuffer9isCreatedEv(void* qthis);
  // proto:  void QOpenGLBuffer::write(int offset, const void * data, int count);
extern void _ZN13QOpenGLBuffer5writeEiPKvi(void* qthis, int arg0, void* arg1, int arg2);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QOpenGLBuffer::read(int offset, void * data, int count);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "read", args)
  }

}

  // proto:  bool QOpenGLBuffer::bind();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "bind", args)
  }

}

  // proto:  void QOpenGLBuffer::destroy();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "destroy", args)
  }

}

  // proto:  void QOpenGLBuffer::allocate(int count);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QOpenGLBuffer8allocateEPKvi
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "allocate", args)
  }

}

  // proto:  bool QOpenGLBuffer::unmap();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "unmap", args)
  }

}

  // proto:  void QOpenGLBuffer::QOpenGLBuffer(const QOpenGLBuffer & other);
func NewQOpenGLBuffer(args ...interface{}) QOpenGLBuffer {
  return QOpenGLBuffer{}
}

  // proto:  int QOpenGLBuffer::size();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "size", args)
  }

}

  // proto:  GLuint QOpenGLBuffer::bufferId();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "bufferId", args)
  }

}

  // proto:  bool QOpenGLBuffer::create();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "create", args)
  }

}

  // proto:  void QOpenGLBuffer::~QOpenGLBuffer();
func (this *QOpenGLBuffer) FreeQOpenGLBuffer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "~QOpenGLBuffer", args)
  }

}

  // proto:  void QOpenGLBuffer::release();
func (this *QOpenGLBuffer) release(args ...interface{}) () {
  // release(class QOpenGLBuffer::Type)
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QOpenGLBuffer::Type"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLBuffer7releaseENS_4TypeE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QOpenGLBuffer7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "release", args)
  }

}

  // proto:  bool QOpenGLBuffer::isCreated();
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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "isCreated", args)
  }

}

  // proto:  void QOpenGLBuffer::write(int offset, const void * data, int count);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "write", args)
  }

}

// <= body block end

