package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qopenglframebufferobject.h
// dst-file: /src/gui/qopenglframebufferobject.go
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
  // proto:  void QOpenGLFramebufferObjectFormat::~QOpenGLFramebufferObjectFormat();
extern void _ZN30QOpenGLFramebufferObjectFormatD0Ev(void* qthis);
  // proto:  void QOpenGLFramebufferObjectFormat::setTextureTarget(GLenum target);
extern void _ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFramebufferObjectFormat::setInternalTextureFormat(GLenum internalTextureFormat);
extern void _ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat(const QOpenGLFramebufferObjectFormat & other);
extern void* dector_ZN30QOpenGLFramebufferObjectFormatC1ERKS_(void* arg0);
extern void _ZN30QOpenGLFramebufferObjectFormatC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QOpenGLFramebufferObjectFormat::mipmap();
extern void _ZNK30QOpenGLFramebufferObjectFormat6mipmapEv(void* qthis);
  // proto:  GLenum QOpenGLFramebufferObjectFormat::textureTarget();
extern void _ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv(void* qthis);
  // proto:  void QOpenGLFramebufferObjectFormat::setSamples(int samples);
extern void _ZN30QOpenGLFramebufferObjectFormat10setSamplesEi(void* qthis, int arg0);
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat();
extern void* dector_ZN30QOpenGLFramebufferObjectFormatC1Ev();
extern void _ZN30QOpenGLFramebufferObjectFormatC1Ev(void* qthis);
  // proto:  void QOpenGLFramebufferObjectFormat::setMipmap(bool enabled);
extern void _ZN30QOpenGLFramebufferObjectFormat9setMipmapEb(void* qthis, bool arg0);
  // proto:  GLenum QOpenGLFramebufferObjectFormat::internalTextureFormat();
extern void _ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv(void* qthis);
  // proto:  int QOpenGLFramebufferObjectFormat::samples();
extern void _ZNK30QOpenGLFramebufferObjectFormat7samplesEv(void* qthis);
  // proto:  bool QOpenGLFramebufferObject::isValid();
extern void _ZNK24QOpenGLFramebufferObject7isValidEv(void* qthis);
  // proto:  GLuint QOpenGLFramebufferObject::takeTexture();
extern void _ZN24QOpenGLFramebufferObject11takeTextureEv(void* qthis);
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, const QOpenGLFramebufferObjectFormat & format);
extern void* dector_ZN24QOpenGLFramebufferObjectC1ERK5QSizeRK30QOpenGLFramebufferObjectFormat(void* arg0, void* arg1);
extern void _ZN24QOpenGLFramebufferObjectC1ERK5QSizeRK30QOpenGLFramebufferObjectFormat(void* qthis, void* arg0, void* arg1);
  // proto: static bool QOpenGLFramebufferObject::bindDefault();
extern void _ZN24QOpenGLFramebufferObject11bindDefaultEv();
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferBlit();
extern void _ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv();
  // proto:  GLuint QOpenGLFramebufferObject::texture();
extern void _ZNK24QOpenGLFramebufferObject7textureEv(void* qthis);
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, GLenum target);
extern void* dector_ZN24QOpenGLFramebufferObjectC1ERK5QSizej(void* arg0, unsigned int arg1);
extern void _ZN24QOpenGLFramebufferObjectC1ERK5QSizej(void* qthis, void* arg0, unsigned int arg1);
  // proto:  bool QOpenGLFramebufferObject::release();
extern void _ZN24QOpenGLFramebufferObject7releaseEv(void* qthis);
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferObjects();
extern void _ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv();
  // proto:  QImage QOpenGLFramebufferObject::toImage(bool flipped);
extern void _ZNK24QOpenGLFramebufferObject7toImageEb(void* qthis, bool arg0);
  // proto:  GLuint QOpenGLFramebufferObject::handle();
extern void _ZNK24QOpenGLFramebufferObject6handleEv(void* qthis);
  // proto:  int QOpenGLFramebufferObject::height();
extern void _ZNK24QOpenGLFramebufferObject6heightEv(void* qthis);
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, QOpenGLFramebufferObject * source, GLbitfield buffers, GLenum filter);
extern void _ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj(void* arg0, void* arg1, unsigned int arg2, unsigned int arg3);
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, const QOpenGLFramebufferObjectFormat & format);
extern void* dector_ZN24QOpenGLFramebufferObjectC1EiiRK30QOpenGLFramebufferObjectFormat(int arg0, int arg1, void* arg2);
extern void _ZN24QOpenGLFramebufferObjectC1EiiRK30QOpenGLFramebufferObjectFormat(void* qthis, int arg0, int arg1, void* arg2);
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, const QRect & targetRect, QOpenGLFramebufferObject * source, const QRect & sourceRect, GLbitfield buffers, GLenum filter);
extern void _ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj(void* arg0, void* arg1, void* arg2, void* arg3, unsigned int arg4, unsigned int arg5);
  // proto:  QImage QOpenGLFramebufferObject::toImage();
extern void _ZNK24QOpenGLFramebufferObject7toImageEv(void* qthis);
  // proto:  QSize QOpenGLFramebufferObject::size();
extern void _ZNK24QOpenGLFramebufferObject4sizeEv(void* qthis);
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QOpenGLFramebufferObject & );
extern void* dector_ZN24QOpenGLFramebufferObjectC1ERKS_(void* arg0);
extern void _ZN24QOpenGLFramebufferObjectC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOpenGLFramebufferObject::~QOpenGLFramebufferObject();
extern void _ZN24QOpenGLFramebufferObjectD0Ev(void* qthis);
  // proto:  QOpenGLFramebufferObjectFormat QOpenGLFramebufferObject::format();
extern void _ZNK24QOpenGLFramebufferObject6formatEv(void* qthis);
  // proto:  bool QOpenGLFramebufferObject::bind();
extern void _ZN24QOpenGLFramebufferObject4bindEv(void* qthis);
  // proto:  bool QOpenGLFramebufferObject::isBound();
extern void _ZNK24QOpenGLFramebufferObject7isBoundEv(void* qthis);
  // proto:  int QOpenGLFramebufferObject::width();
extern void _ZNK24QOpenGLFramebufferObject5widthEv(void* qthis);
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, GLenum target);
extern void* dector_ZN24QOpenGLFramebufferObjectC1Eiij(int arg0, int arg1, unsigned int arg2);
extern void _ZN24QOpenGLFramebufferObjectC1Eiij(void* qthis, int arg0, int arg1, unsigned int arg2);
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

  // proto:  void QOpenGLFramebufferObjectFormat::~QOpenGLFramebufferObjectFormat();
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

  // proto:  void QOpenGLFramebufferObjectFormat::setTextureTarget(GLenum target);
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

  // proto:  void QOpenGLFramebufferObjectFormat::setInternalTextureFormat(GLenum internalTextureFormat);
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

  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat(const QOpenGLFramebufferObjectFormat & other);
func NewQOpenGLFramebufferObjectFormat(args ...interface{}) QOpenGLFramebufferObjectFormat {
  return QOpenGLFramebufferObjectFormat{}
}

  // proto:  bool QOpenGLFramebufferObjectFormat::mipmap();
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

  // proto:  GLenum QOpenGLFramebufferObjectFormat::textureTarget();
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

  // proto:  void QOpenGLFramebufferObjectFormat::setSamples(int samples);
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

  // proto:  void QOpenGLFramebufferObjectFormat::setMipmap(bool enabled);
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

  // proto:  GLenum QOpenGLFramebufferObjectFormat::internalTextureFormat();
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

  // proto:  int QOpenGLFramebufferObjectFormat::samples();
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

  // proto:  bool QOpenGLFramebufferObject::isValid();
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

  // proto:  GLuint QOpenGLFramebufferObject::takeTexture();
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

  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, const QOpenGLFramebufferObjectFormat & format);
func NewQOpenGLFramebufferObject(args ...interface{}) QOpenGLFramebufferObject {
  return QOpenGLFramebufferObject{}
}

  // proto: static bool QOpenGLFramebufferObject::bindDefault();
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

  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferBlit();
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

  // proto:  GLuint QOpenGLFramebufferObject::texture();
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

  // proto:  bool QOpenGLFramebufferObject::release();
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

  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferObjects();
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

  // proto:  QImage QOpenGLFramebufferObject::toImage(bool flipped);
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

  // proto:  GLuint QOpenGLFramebufferObject::handle();
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

  // proto:  int QOpenGLFramebufferObject::height();
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

  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, QOpenGLFramebufferObject * source, GLbitfield buffers, GLenum filter);
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

  // proto:  QSize QOpenGLFramebufferObject::size();
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

  // proto:  void QOpenGLFramebufferObject::~QOpenGLFramebufferObject();
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

  // proto:  QOpenGLFramebufferObjectFormat QOpenGLFramebufferObject::format();
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

  // proto:  bool QOpenGLFramebufferObject::bind();
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

  // proto:  bool QOpenGLFramebufferObject::isBound();
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

  // proto:  int QOpenGLFramebufferObject::width();
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

