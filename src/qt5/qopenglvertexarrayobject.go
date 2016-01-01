package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qopenglvertexarrayobject.h
// dst-file: /src/gui/qopenglvertexarrayobject.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QOpenGLVertexArrayObject)=1
type QOpenGLVertexArrayObject struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQOpenGLVertexArrayObject(args ...interface{}) QOpenGLVertexArrayObject {
  return QOpenGLVertexArrayObject{}
}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "objectId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "release", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "bind", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "isCreated", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "destroy", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVertexArrayObject", "create", args)
  }

}


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

