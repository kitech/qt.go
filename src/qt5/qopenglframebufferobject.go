package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QOpenGLFramebufferObjectFormat::setMipmap(bool enabled);
extern void C_ZN30QOpenGLFramebufferObjectFormat9setMipmapEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::setSamples(int samples);
extern void C_ZN30QOpenGLFramebufferObjectFormat10setSamplesEi(void* qthis, int32_t arg0); // 4
  // proto:  GLenum QOpenGLFramebufferObjectFormat::textureTarget();
extern void C_ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::setInternalTextureFormat(GLenum internalTextureFormat);
extern void C_ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj(void* qthis, int32_t arg0); // 4
  // proto:  bool QOpenGLFramebufferObjectFormat::mipmap();
extern void C_ZNK30QOpenGLFramebufferObjectFormat6mipmapEv(void* qthis); // 4
  // proto:  QOpenGLFramebufferObject::Attachment QOpenGLFramebufferObjectFormat::attachment();
extern void C_ZNK30QOpenGLFramebufferObjectFormat10attachmentEv(void* qthis); // 4
  // proto:  int QOpenGLFramebufferObjectFormat::samples();
extern void C_ZNK30QOpenGLFramebufferObjectFormat7samplesEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::~QOpenGLFramebufferObjectFormat();
extern void C_ZN30QOpenGLFramebufferObjectFormatD2Ev(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::setTextureTarget(GLenum target);
extern void C_ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat(const QOpenGLFramebufferObjectFormat & other);
extern void* C_ZN30QOpenGLFramebufferObjectFormatC2ERKS_(void* arg0); // 3
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat();
extern void* C_ZN30QOpenGLFramebufferObjectFormatC2Ev(); // 3
  // proto:  GLenum QOpenGLFramebufferObjectFormat::internalTextureFormat();
extern void C_ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv(void* qthis); // 4
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferObjects();
extern void C_ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv(); // 4
  // proto:  int QOpenGLFramebufferObject::height();
extern void C_ZNK24QOpenGLFramebufferObject6heightEv(void* qthis); // 2
  // proto:  QSize QOpenGLFramebufferObject::size();
extern void C_ZNK24QOpenGLFramebufferObject4sizeEv(void* qthis); // 4
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferBlit();
extern void C_ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv(); // 4
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, const QRect & targetRect, QOpenGLFramebufferObject * source, const QRect & sourceRect, GLbitfield buffers, GLenum filter);
extern void C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj(void* arg0, void* arg1, void* arg2, void* arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, QOpenGLFramebufferObject * source, GLbitfield buffers, GLenum filter);
extern void C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj(void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QOpenGLFramebufferObjectFormat QOpenGLFramebufferObject::format();
extern void C_ZNK24QOpenGLFramebufferObject6formatEv(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::takeTexture();
extern void C_ZN24QOpenGLFramebufferObject11takeTextureEv(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::texture();
extern void C_ZNK24QOpenGLFramebufferObject7textureEv(void* qthis); // 4
  // proto:  int QOpenGLFramebufferObject::width();
extern void C_ZNK24QOpenGLFramebufferObject5widthEv(void* qthis); // 2
  // proto:  QOpenGLFramebufferObject::Attachment QOpenGLFramebufferObject::attachment();
extern void C_ZNK24QOpenGLFramebufferObject10attachmentEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObject::~QOpenGLFramebufferObject();
extern void C_ZN24QOpenGLFramebufferObjectD2Ev(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::handle();
extern void C_ZNK24QOpenGLFramebufferObject6handleEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::isValid();
extern void C_ZNK24QOpenGLFramebufferObject7isValidEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, GLenum target);
extern void* C_ZN24QOpenGLFramebufferObjectC2Eiij(int32_t arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, const QOpenGLFramebufferObjectFormat & format);
extern void* C_ZN24QOpenGLFramebufferObjectC2EiiRK30QOpenGLFramebufferObjectFormat(int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, const QOpenGLFramebufferObjectFormat & format);
extern void* C_ZN24QOpenGLFramebufferObjectC2ERK5QSizeRK30QOpenGLFramebufferObjectFormat(void* arg0, void* arg1); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, GLenum target);
extern void* C_ZN24QOpenGLFramebufferObjectC2ERK5QSizej(void* arg0, int32_t arg1); // 3
  // proto: static bool QOpenGLFramebufferObject::bindDefault();
extern void C_ZN24QOpenGLFramebufferObject11bindDefaultEv(); // 4
  // proto:  bool QOpenGLFramebufferObject::isBound();
extern void C_ZNK24QOpenGLFramebufferObject7isBoundEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::bind();
extern void C_ZN24QOpenGLFramebufferObject4bindEv(void* qthis); // 4
  // proto:  QImage QOpenGLFramebufferObject::toImage(bool flipped);
extern void C_ZNK24QOpenGLFramebufferObject7toImageEb(void* qthis, bool arg0); // 4
  // proto:  QImage QOpenGLFramebufferObject::toImage();
extern void C_ZNK24QOpenGLFramebufferObject7toImageEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::release();
extern void C_ZN24QOpenGLFramebufferObject7releaseEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLFramebufferObject)=1
type QOpenGLFramebufferObject struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setMipmap(_Bool)
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
    // invoke: void setMipmap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat9setMipmapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setMipmap", args)
  }

}

// setSamples(int)
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
    // invoke: void setSamples(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat10setSamplesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setSamples", args)
  }

}

// textureTarget()
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
    // invoke: GLenum textureTarget()
    var ret = C.C_ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "textureTarget", args)
  }

}

// setInternalTextureFormat(GLenum)
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
    // invoke: void setInternalTextureFormat(GLenum)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setInternalTextureFormat", args)
  }

}

// mipmap()
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
    // invoke: bool mipmap()
    var ret = C.C_ZNK30QOpenGLFramebufferObjectFormat6mipmapEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "mipmap", args)
  }

}

// attachment()
func (this *QOpenGLFramebufferObjectFormat) attachment(args ...interface{}) () {
  // attachment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK30QOpenGLFramebufferObjectFormat10attachmentEv
    // invoke: QOpenGLFramebufferObject::Attachment attachment()
    C.C_ZNK30QOpenGLFramebufferObjectFormat10attachmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "attachment", args)
  }

}

// samples()
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
    // invoke: int samples()
    var ret = C.C_ZNK30QOpenGLFramebufferObjectFormat7samplesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "samples", args)
  }

}

// ~QOpenGLFramebufferObjectFormat()
func (this *QOpenGLFramebufferObjectFormat) FreeQOpenGLFramebufferObjectFormat(args ...interface{}) () {
  // ~QOpenGLFramebufferObjectFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormatD0Ev
    // invoke: void ~QOpenGLFramebufferObjectFormat()
    C.C_ZN30QOpenGLFramebufferObjectFormatD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "~QOpenGLFramebufferObjectFormat", args)
  }

}

// setTextureTarget(GLenum)
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
    // invoke: void setTextureTarget(GLenum)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setTextureTarget", args)
  }

}

// QOpenGLFramebufferObjectFormat(const class QOpenGLFramebufferObjectFormat &)
func NewQOpenGLFramebufferObjectFormat(args ...interface{}) *QOpenGLFramebufferObjectFormat {
  // QOpenGLFramebufferObjectFormat(const class QOpenGLFramebufferObjectFormat &)
  // QOpenGLFramebufferObjectFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLFramebufferObjectFormat{}) // "const QOpenGLFramebufferObjectFormat &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN30QOpenGLFramebufferObjectFormatC1ERKS_
    // invoke: void QOpenGLFramebufferObjectFormat(const class QOpenGLFramebufferObjectFormat &)
    var arg0 = args[0].(QOpenGLFramebufferObjectFormat).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QOpenGLFramebufferObjectFormatC2ERKS_(arg0)
    return &QOpenGLFramebufferObjectFormat{qclsinst:qthis}
  case 1:
    // invoke: _ZN30QOpenGLFramebufferObjectFormatC1Ev
    // invoke: void QOpenGLFramebufferObjectFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QOpenGLFramebufferObjectFormatC2Ev()
    return &QOpenGLFramebufferObjectFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "QOpenGLFramebufferObjectFormat", args)
  }

  return nil // QOpenGLFramebufferObjectFormat{qclsinst:qthis}
}

// internalTextureFormat()
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
    // invoke: GLenum internalTextureFormat()
    var ret = C.C_ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "internalTextureFormat", args)
  }

}

// hasOpenGLFramebufferObjects()
func (this *QOpenGLFramebufferObject) hasOpenGLFramebufferObjects_s(args ...interface{}) () {
  // hasOpenGLFramebufferObjects()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv
    // invoke: bool hasOpenGLFramebufferObjects()
    var ret = C.C_ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferObjects", args)
  }

}

// height()
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
    // invoke: int height()
    var ret = C.C_ZNK24QOpenGLFramebufferObject6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "height", args)
  }

}

// size()
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
    // invoke: QSize size()
    var ret = C.C_ZNK24QOpenGLFramebufferObject4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "size", args)
  }

}

// hasOpenGLFramebufferBlit()
func (this *QOpenGLFramebufferObject) hasOpenGLFramebufferBlit_s(args ...interface{}) () {
  // hasOpenGLFramebufferBlit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv
    // invoke: bool hasOpenGLFramebufferBlit()
    var ret = C.C_ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferBlit", args)
  }

}

// blitFramebuffer(class QOpenGLFramebufferObject *, const class QRect &, class QOpenGLFramebufferObject *, const class QRect &, GLbitfield, GLenum)
func (this *QOpenGLFramebufferObject) blitFramebuffer_s(args ...interface{}) () {
  // blitFramebuffer(class QOpenGLFramebufferObject *, const class QRect &, class QOpenGLFramebufferObject *, const class QRect &, GLbitfield, GLenum)
  // blitFramebuffer(class QOpenGLFramebufferObject *, class QOpenGLFramebufferObject *, GLbitfield, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[0][1] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][2] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[0][3] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLbitfield"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[1][1] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[1][2] = qtrt.Int32Ty(false) // "GLbitfield"
  vtys[1][3] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj
    // invoke: void blitFramebuffer(class QOpenGLFramebufferObject *, const class QRect &, class QOpenGLFramebufferObject *, const class QRect &, GLbitfield, GLenum)
    var arg0 = args[0].(QOpenGLFramebufferObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QOpenGLFramebufferObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRect).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj(arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj
    // invoke: void blitFramebuffer(class QOpenGLFramebufferObject *, class QOpenGLFramebufferObject *, GLbitfield, GLenum)
    var arg0 = args[0].(QOpenGLFramebufferObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QOpenGLFramebufferObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "blitFramebuffer", args)
  }

}

// format()
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
    // invoke: QOpenGLFramebufferObjectFormat format()
    C.C_ZNK24QOpenGLFramebufferObject6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "format", args)
  }

}

// takeTexture()
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
    // invoke: GLuint takeTexture()
    var ret = C.C_ZN24QOpenGLFramebufferObject11takeTextureEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "takeTexture", args)
  }

}

// texture()
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
    // invoke: GLuint texture()
    var ret = C.C_ZNK24QOpenGLFramebufferObject7textureEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "texture", args)
  }

}

// width()
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
    // invoke: int width()
    var ret = C.C_ZNK24QOpenGLFramebufferObject5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "width", args)
  }

}

// attachment()
func (this *QOpenGLFramebufferObject) attachment(args ...interface{}) () {
  // attachment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QOpenGLFramebufferObject10attachmentEv
    // invoke: QOpenGLFramebufferObject::Attachment attachment()
    C.C_ZNK24QOpenGLFramebufferObject10attachmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "attachment", args)
  }

}

// ~QOpenGLFramebufferObject()
func (this *QOpenGLFramebufferObject) FreeQOpenGLFramebufferObject(args ...interface{}) () {
  // ~QOpenGLFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObjectD0Ev
    // invoke: void ~QOpenGLFramebufferObject()
    C.C_ZN24QOpenGLFramebufferObjectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "~QOpenGLFramebufferObject", args)
  }

}

// handle()
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
    // invoke: GLuint handle()
    var ret = C.C_ZNK24QOpenGLFramebufferObject6handleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "handle", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    var ret = C.C_ZNK24QOpenGLFramebufferObject7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isValid", args)
  }

}

// QOpenGLFramebufferObject(int, int, GLenum)
func NewQOpenGLFramebufferObject(args ...interface{}) *QOpenGLFramebufferObject {
  // QOpenGLFramebufferObject(int, int, GLenum)
  // QOpenGLFramebufferObject(int, int, const class QOpenGLFramebufferObjectFormat &)
  // QOpenGLFramebufferObject(const class QSize &, const class QOpenGLFramebufferObjectFormat &)
  // QOpenGLFramebufferObject(const class QSize &, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QOpenGLFramebufferObjectFormat{}) // "const QOpenGLFramebufferObjectFormat &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[2][1] = reflect.TypeOf(QOpenGLFramebufferObjectFormat{}) // "const QOpenGLFramebufferObjectFormat &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[3][1] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObjectC1Eiij
    // invoke: void QOpenGLFramebufferObject(int, int, GLenum)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2Eiij(arg0, arg1, arg2)
    return &QOpenGLFramebufferObject{qclsinst:qthis}
  case 1:
    // invoke: _ZN24QOpenGLFramebufferObjectC1EiiRK30QOpenGLFramebufferObjectFormat
    // invoke: void QOpenGLFramebufferObject(int, int, const class QOpenGLFramebufferObjectFormat &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QOpenGLFramebufferObjectFormat).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2EiiRK30QOpenGLFramebufferObjectFormat(arg0, arg1, arg2)
    return &QOpenGLFramebufferObject{qclsinst:qthis}
  case 2:
    // invoke: _ZN24QOpenGLFramebufferObjectC1ERK5QSizeRK30QOpenGLFramebufferObjectFormat
    // invoke: void QOpenGLFramebufferObject(const class QSize &, const class QOpenGLFramebufferObjectFormat &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QOpenGLFramebufferObjectFormat).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2ERK5QSizeRK30QOpenGLFramebufferObjectFormat(arg0, arg1)
    return &QOpenGLFramebufferObject{qclsinst:qthis}
  case 3:
    // invoke: _ZN24QOpenGLFramebufferObjectC1ERK5QSizej
    // invoke: void QOpenGLFramebufferObject(const class QSize &, GLenum)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2ERK5QSizej(arg0, arg1)
    return &QOpenGLFramebufferObject{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "QOpenGLFramebufferObject", args)
  }

  return nil // QOpenGLFramebufferObject{qclsinst:qthis}
}

// bindDefault()
func (this *QOpenGLFramebufferObject) bindDefault_s(args ...interface{}) () {
  // bindDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObject11bindDefaultEv
    // invoke: bool bindDefault()
    var ret = C.C_ZN24QOpenGLFramebufferObject11bindDefaultEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bindDefault", args)
  }

}

// isBound()
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
    // invoke: bool isBound()
    var ret = C.C_ZNK24QOpenGLFramebufferObject7isBoundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isBound", args)
  }

}

// bind()
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
    // invoke: bool bind()
    var ret = C.C_ZN24QOpenGLFramebufferObject4bindEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bind", args)
  }

}

// toImage(_Bool)
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
    // invoke: QImage toImage(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK24QOpenGLFramebufferObject7toImageEb(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK24QOpenGLFramebufferObject7toImageEv
    // invoke: QImage toImage()
    var ret = C.C_ZNK24QOpenGLFramebufferObject7toImageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "toImage", args)
  }

}

// release()
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
    // invoke: bool release()
    var ret = C.C_ZN24QOpenGLFramebufferObject7releaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "release", args)
  }

}

// <= body block end

