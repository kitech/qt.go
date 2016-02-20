package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QOpenGLVertexArrayObject::isCreated();
extern bool C_ZNK24QOpenGLVertexArrayObject9isCreatedEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLVertexArrayObject::metaObject();
extern void C_ZNK24QOpenGLVertexArrayObject10metaObjectEv(void* qthis); // 4
  // proto:  GLuint QOpenGLVertexArrayObject::objectId();
extern int32_t C_ZNK24QOpenGLVertexArrayObject8objectIdEv(void* qthis); // 4
  // proto:  void QOpenGLVertexArrayObject::bind();
extern void C_ZN24QOpenGLVertexArrayObject4bindEv(void* qthis); // 4
  // proto:  void QOpenGLVertexArrayObject::~QOpenGLVertexArrayObject();
extern void C_ZN24QOpenGLVertexArrayObjectD2Ev(void* qthis); // 4
  // proto:  void QOpenGLVertexArrayObject::release();
extern void C_ZN24QOpenGLVertexArrayObject7releaseEv(void* qthis); // 4
  // proto:  void QOpenGLVertexArrayObject::destroy();
extern void C_ZN24QOpenGLVertexArrayObject7destroyEv(void* qthis); // 4
  // proto:  void QOpenGLVertexArrayObject::QOpenGLVertexArrayObject(QObject * parent);
extern void* C_ZN24QOpenGLVertexArrayObjectC2EP7QObject(void* arg0); // 3
  // proto:  bool QOpenGLVertexArrayObject::create();
extern bool C_ZN24QOpenGLVertexArrayObject6createEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QOpenGLVertexArrayObject)=1
type QOpenGLVertexArrayObject struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isCreated()
func (this *QOpenGLVertexArrayObject) Iscreated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLVertexArrayObject9isCreatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "isCreated", args)
  }

  return
}

// metaObject()
func (this *QOpenGLVertexArrayObject) Metaobject(args ...interface{}) () {
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
    C.C_ZNK24QOpenGLVertexArrayObject10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "metaObject", args)
  }

  return
}

// objectId()
func (this *QOpenGLVertexArrayObject) Objectid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLVertexArrayObject8objectIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "objectId", args)
  }

  return
}

// bind()
func (this *QOpenGLVertexArrayObject) Bind(args ...interface{}) () {
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
    C.C_ZN24QOpenGLVertexArrayObject4bindEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "bind", args)
  }

  return
}

// ~QOpenGLVertexArrayObject()
func (this *QOpenGLVertexArrayObject) Freeqopenglvertexarrayobject(args ...interface{}) () {
  // ~QOpenGLVertexArrayObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObjectD0Ev
    // invoke: void ~QOpenGLVertexArrayObject()
    C.C_ZN24QOpenGLVertexArrayObjectD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "~QOpenGLVertexArrayObject", args)
  }

  return
}

// release()
func (this *QOpenGLVertexArrayObject) Release(args ...interface{}) () {
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
    C.C_ZN24QOpenGLVertexArrayObject7releaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "release", args)
  }

  return
}

// destroy()
func (this *QOpenGLVertexArrayObject) Destroy(args ...interface{}) () {
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
    C.C_ZN24QOpenGLVertexArrayObject7destroyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "destroy", args)
  }

  return
}

// QOpenGLVertexArrayObject(class QObject *)
func NewQOpenGLVertexArrayObject(args ...interface{}) *QOpenGLVertexArrayObject {
  // QOpenGLVertexArrayObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLVertexArrayObjectC1EP7QObject
    // invoke: void QOpenGLVertexArrayObject(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLVertexArrayObjectC2EP7QObject(arg0)
    return &QOpenGLVertexArrayObject{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "QOpenGLVertexArrayObject", args)
  }

  return nil // QOpenGLVertexArrayObject{Qclsinst:qthis}
}

// create()
func (this *QOpenGLVertexArrayObject) Create(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLVertexArrayObject6createEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "create", args)
  }

  return
}

// <= body block end

