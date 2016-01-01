package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglframebufferobject.h
// dst-file: /src/gui/qopenglframebufferobject.go
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
// class sizeof(QOpenGLFramebufferObjectFormat)=8
type QOpenGLFramebufferObjectFormat struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFramebufferObject)=1
type QOpenGLFramebufferObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QOpenGLFramebufferObjectFormat) FreeQOpenGLFramebufferObjectFormat(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "~QOpenGLFramebufferObjectFormat", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) setTextureTarget(args ...interface{}) () {
  // setTextureTarget(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setTextureTarget", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) setInternalTextureFormat(args ...interface{}) () {
  // setInternalTextureFormat(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setInternalTextureFormat", args)
 }

}


func NewQOpenGLFramebufferObjectFormat(args ...interface{})() {
}


func (this *QOpenGLFramebufferObjectFormat) mipmap(args ...interface{}) () {
  // mipmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK30QOpenGLFramebufferObjectFormat6mipmapEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "mipmap", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) textureTarget(args ...interface{}) () {
  // textureTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "textureTarget", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) setSamples(args ...interface{}) () {
  // setSamples(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormat10setSamplesEi
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setSamples", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) setMipmap(args ...interface{}) () {
  // setMipmap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormat9setMipmapEb
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setMipmap", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) internalTextureFormat(args ...interface{}) () {
  // internalTextureFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "internalTextureFormat", args)
 }

}


func (this *QOpenGLFramebufferObjectFormat) samples(args ...interface{}) () {
  // samples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK30QOpenGLFramebufferObjectFormat7samplesEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "samples", args)
 }

}


func (this *QOpenGLFramebufferObject) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject7isValidEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isValid", args)
 }

}


func (this *QOpenGLFramebufferObject) takeTexture(args ...interface{}) () {
  // takeTexture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject11takeTextureEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "takeTexture", args)
 }

}


func NewQOpenGLFramebufferObject(args ...interface{})() {
}


func (this *QOpenGLFramebufferObject) bindDefault_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bindDefault", args)
 }

}


func (this *QOpenGLFramebufferObject) hasOpenGLFramebufferBlit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferBlit", args)
 }

}


func (this *QOpenGLFramebufferObject) texture(args ...interface{}) () {
  // texture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject7textureEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "texture", args)
 }

}


func (this *QOpenGLFramebufferObject) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "release", args)
 }

}


func (this *QOpenGLFramebufferObject) hasOpenGLFramebufferObjects_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferObjects", args)
 }

}


func (this *QOpenGLFramebufferObject) toImage(args ...interface{}) () {
  // toImage(_Bool)
  // toImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject7toImageEb
  case 1:
    // invoke: _ZNK24QOpenGLFramebufferObject7toImageEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "toImage", args)
 }

}


func (this *QOpenGLFramebufferObject) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject6handleEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "handle", args)
 }

}


func (this *QOpenGLFramebufferObject) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject6heightEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "height", args)
 }

}


func (this *QOpenGLFramebufferObject) blitFramebuffer_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "blitFramebuffer", args)
 }

}


func (this *QOpenGLFramebufferObject) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject4sizeEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "size", args)
 }

}


func (this *QOpenGLFramebufferObject) FreeQOpenGLFramebufferObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "~QOpenGLFramebufferObject", args)
 }

}


func (this *QOpenGLFramebufferObject) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject6formatEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "format", args)
 }

}


func (this *QOpenGLFramebufferObject) bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject4bindEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bind", args)
 }

}


func (this *QOpenGLFramebufferObject) isBound(args ...interface{}) () {
  // isBound()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject7isBoundEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isBound", args)
 }

}


func (this *QOpenGLFramebufferObject) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject5widthEv
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "width", args)
 }

}

// <= body block end

