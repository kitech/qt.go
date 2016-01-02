package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qopenglvertexarrayobject.h
// dst-file: /src/gui/qopenglvertexarrayobject.go
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
  // proto:  void QOpenGLVertexArrayObject::QOpenGLVertexArrayObject(const QOpenGLVertexArrayObject & );
extern void* dector_ZN24QOpenGLVertexArrayObjectC1ERKS_(void* arg0);
extern void _ZN24QOpenGLVertexArrayObjectC1ERKS_(void* qthis, void* arg0);
  // proto:  GLuint QOpenGLVertexArrayObject::objectId();
extern void _ZNK24QOpenGLVertexArrayObject8objectIdEv(void* qthis);
  // proto:  void QOpenGLVertexArrayObject::release();
extern void _ZN24QOpenGLVertexArrayObject7releaseEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLVertexArrayObject::metaObject();
extern void _ZNK24QOpenGLVertexArrayObject10metaObjectEv(void* qthis);
  // proto:  void QOpenGLVertexArrayObject::QOpenGLVertexArrayObject(QObject * parent);
extern void* dector_ZN24QOpenGLVertexArrayObjectC1EP7QObject(void* arg0);
extern void _ZN24QOpenGLVertexArrayObjectC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QOpenGLVertexArrayObject::bind();
extern void _ZN24QOpenGLVertexArrayObject4bindEv(void* qthis);
  // proto:  bool QOpenGLVertexArrayObject::isCreated();
extern void _ZNK24QOpenGLVertexArrayObject9isCreatedEv(void* qthis);
  // proto:  void QOpenGLVertexArrayObject::destroy();
extern void _ZN24QOpenGLVertexArrayObject7destroyEv(void* qthis);
  // proto:  bool QOpenGLVertexArrayObject::create();
extern void _ZN24QOpenGLVertexArrayObject6createEv(void* qthis);
  // proto:  void QOpenGLVertexArrayObject::~QOpenGLVertexArrayObject();
extern void _ZN24QOpenGLVertexArrayObjectD0Ev(void* qthis);
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

// class sizeof(QOpenGLVertexArrayObject)=1
type QOpenGLVertexArrayObject struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QOpenGLVertexArrayObject::QOpenGLVertexArrayObject(const QOpenGLVertexArrayObject & );
func NewQOpenGLVertexArrayObject(args ...interface{}) QOpenGLVertexArrayObject {
  return QOpenGLVertexArrayObject{}
}

  // proto:  GLuint QOpenGLVertexArrayObject::objectId();
func (this *QOpenGLVertexArrayObject) objectId(args ...interface{}) () {
  // objectId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLVertexArrayObject8objectIdEv
    // invoke: GLuint objectId()
    C._ZNK24QOpenGLVertexArrayObject8objectIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "objectId", args)
  }

}

  // proto:  void QOpenGLVertexArrayObject::release();
func (this *QOpenGLVertexArrayObject) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObject7releaseEv
    // invoke: void release()
    C._ZN24QOpenGLVertexArrayObject7releaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "release", args)
  }

}

  // proto:  const QMetaObject * QOpenGLVertexArrayObject::metaObject();
func (this *QOpenGLVertexArrayObject) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLVertexArrayObject10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK24QOpenGLVertexArrayObject10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "metaObject", args)
  }

}

  // proto:  void QOpenGLVertexArrayObject::bind();
func (this *QOpenGLVertexArrayObject) bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObject4bindEv
    // invoke: void bind()
    C._ZN24QOpenGLVertexArrayObject4bindEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "bind", args)
  }

}

  // proto:  bool QOpenGLVertexArrayObject::isCreated();
func (this *QOpenGLVertexArrayObject) isCreated(args ...interface{}) () {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLVertexArrayObject9isCreatedEv
    // invoke: bool isCreated()
    C._ZNK24QOpenGLVertexArrayObject9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "isCreated", args)
  }

}

  // proto:  void QOpenGLVertexArrayObject::destroy();
func (this *QOpenGLVertexArrayObject) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObject7destroyEv
    // invoke: void destroy()
    C._ZN24QOpenGLVertexArrayObject7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "destroy", args)
  }

}

  // proto:  bool QOpenGLVertexArrayObject::create();
func (this *QOpenGLVertexArrayObject) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObject6createEv
    // invoke: bool create()
    C._ZN24QOpenGLVertexArrayObject6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "create", args)
  }

}

  // proto:  void QOpenGLVertexArrayObject::~QOpenGLVertexArrayObject();
func (this *QOpenGLVertexArrayObject) FreeQOpenGLVertexArrayObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "~QOpenGLVertexArrayObject", args)
  }

}

// <= body block end

