package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qopenglbuffer.h
// dst-file: /src/gui/qopenglbuffer.go
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
// class sizeof(QOpenGLBuffer)=8
type QOpenGLBuffer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "read", args)
  }

}


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
  case 1:
    // invoke: _ZN13QOpenGLBuffer8allocateEPKvi
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "allocate", args)
  }

}


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


func NewQOpenGLBuffer(args ...interface{}) QOpenGLBuffer {
  return QOpenGLBuffer{}
}


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
  case 1:
    // invoke: _ZN13QOpenGLBuffer7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "release", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLBuffer", "write", args)
  }

}

// <= body block end

