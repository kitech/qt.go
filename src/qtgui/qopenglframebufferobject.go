package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
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
extern int32_t C_ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::setInternalTextureFormat(GLenum internalTextureFormat);
extern void C_ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj(void* qthis, int32_t arg0); // 4
  // proto:  bool QOpenGLFramebufferObjectFormat::mipmap();
extern bool C_ZNK30QOpenGLFramebufferObjectFormat6mipmapEv(void* qthis); // 4
  // proto:  QOpenGLFramebufferObject::Attachment QOpenGLFramebufferObjectFormat::attachment();
extern void C_ZNK30QOpenGLFramebufferObjectFormat10attachmentEv(void* qthis); // 4
  // proto:  int QOpenGLFramebufferObjectFormat::samples();
extern int32_t C_ZNK30QOpenGLFramebufferObjectFormat7samplesEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::~QOpenGLFramebufferObjectFormat();
extern void C_ZN30QOpenGLFramebufferObjectFormatD2Ev(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::setTextureTarget(GLenum target);
extern void C_ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat(const QOpenGLFramebufferObjectFormat & other);
extern void* C_ZN30QOpenGLFramebufferObjectFormatC2ERKS_(void* arg0); // 3
  // proto:  void QOpenGLFramebufferObjectFormat::QOpenGLFramebufferObjectFormat();
extern void* C_ZN30QOpenGLFramebufferObjectFormatC2Ev(); // 3
  // proto:  GLenum QOpenGLFramebufferObjectFormat::internalTextureFormat();
extern int32_t C_ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv(void* qthis); // 4
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferObjects();
extern bool C_ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv(); // 4
  // proto:  int QOpenGLFramebufferObject::height();
extern int32_t C_ZNK24QOpenGLFramebufferObject6heightEv(void* qthis); // 2
  // proto:  QSize QOpenGLFramebufferObject::size();
extern void* C_ZNK24QOpenGLFramebufferObject4sizeEv(void* qthis); // 4
  // proto: static bool QOpenGLFramebufferObject::hasOpenGLFramebufferBlit();
extern bool C_ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv(); // 4
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, const QRect & targetRect, QOpenGLFramebufferObject * source, const QRect & sourceRect, GLbitfield buffers, GLenum filter);
extern void C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj(void* arg0, void* arg1, void* arg2, void* arg3, int32_t arg4, int32_t arg5); // 4
  // proto: static void QOpenGLFramebufferObject::blitFramebuffer(QOpenGLFramebufferObject * target, QOpenGLFramebufferObject * source, GLbitfield buffers, GLenum filter);
extern void C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj(void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QOpenGLFramebufferObjectFormat QOpenGLFramebufferObject::format();
extern void C_ZNK24QOpenGLFramebufferObject6formatEv(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::takeTexture();
extern int32_t C_ZN24QOpenGLFramebufferObject11takeTextureEv(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::texture();
extern int32_t C_ZNK24QOpenGLFramebufferObject7textureEv(void* qthis); // 4
  // proto:  int QOpenGLFramebufferObject::width();
extern int32_t C_ZNK24QOpenGLFramebufferObject5widthEv(void* qthis); // 2
  // proto:  QOpenGLFramebufferObject::Attachment QOpenGLFramebufferObject::attachment();
extern void C_ZNK24QOpenGLFramebufferObject10attachmentEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObject::~QOpenGLFramebufferObject();
extern void C_ZN24QOpenGLFramebufferObjectD2Ev(void* qthis); // 4
  // proto:  GLuint QOpenGLFramebufferObject::handle();
extern int32_t C_ZNK24QOpenGLFramebufferObject6handleEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::isValid();
extern bool C_ZNK24QOpenGLFramebufferObject7isValidEv(void* qthis); // 4
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, GLenum target);
extern void* C_ZN24QOpenGLFramebufferObjectC2Eiij(int32_t arg0, int32_t arg1, int32_t arg2); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(int width, int height, const QOpenGLFramebufferObjectFormat & format);
extern void* C_ZN24QOpenGLFramebufferObjectC2EiiRK30QOpenGLFramebufferObjectFormat(int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, const QOpenGLFramebufferObjectFormat & format);
extern void* C_ZN24QOpenGLFramebufferObjectC2ERK5QSizeRK30QOpenGLFramebufferObjectFormat(void* arg0, void* arg1); // 3
  // proto:  void QOpenGLFramebufferObject::QOpenGLFramebufferObject(const QSize & size, GLenum target);
extern void* C_ZN24QOpenGLFramebufferObjectC2ERK5QSizej(void* arg0, int32_t arg1); // 3
  // proto: static bool QOpenGLFramebufferObject::bindDefault();
extern bool C_ZN24QOpenGLFramebufferObject11bindDefaultEv(); // 4
  // proto:  bool QOpenGLFramebufferObject::isBound();
extern bool C_ZNK24QOpenGLFramebufferObject7isBoundEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::bind();
extern bool C_ZN24QOpenGLFramebufferObject4bindEv(void* qthis); // 4
  // proto:  QImage QOpenGLFramebufferObject::toImage(bool flipped);
extern void* C_ZNK24QOpenGLFramebufferObject7toImageEb(void* qthis, bool arg0); // 4
  // proto:  QImage QOpenGLFramebufferObject::toImage();
extern void* C_ZNK24QOpenGLFramebufferObject7toImageEv(void* qthis); // 4
  // proto:  bool QOpenGLFramebufferObject::release();
extern bool C_ZN24QOpenGLFramebufferObject7releaseEv(void* qthis); // 4
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

// class sizeof(QOpenGLFramebufferObjectFormat)=8
type QOpenGLFramebufferObjectFormat struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLFramebufferObject)=1
type QOpenGLFramebufferObject struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setMipmap(_Bool)
func (this *QOpenGLFramebufferObjectFormat) Setmipmap(args ...interface{}) () {
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
    C.C_ZN30QOpenGLFramebufferObjectFormat9setMipmapEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setMipmap", args)
  }

  return
}

// setSamples(int)
func (this *QOpenGLFramebufferObjectFormat) Setsamples(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat10setSamplesEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setSamples", args)
  }

  return
}

// textureTarget()
func (this *QOpenGLFramebufferObjectFormat) Texturetarget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK30QOpenGLFramebufferObjectFormat13textureTargetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLenum"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "textureTarget", args)
  }

  return
}

// setInternalTextureFormat(GLenum)
func (this *QOpenGLFramebufferObjectFormat) Setinternaltextureformat(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat24setInternalTextureFormatEj(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setInternalTextureFormat", args)
  }

  return
}

// mipmap()
func (this *QOpenGLFramebufferObjectFormat) Mipmap(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK30QOpenGLFramebufferObjectFormat6mipmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "mipmap", args)
  }

  return
}

// attachment()
func (this *QOpenGLFramebufferObjectFormat) Attachment(args ...interface{}) () {
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
    C.C_ZNK30QOpenGLFramebufferObjectFormat10attachmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "attachment", args)
  }

  return
}

// samples()
func (this *QOpenGLFramebufferObjectFormat) Samples(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK30QOpenGLFramebufferObjectFormat7samplesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "samples", args)
  }

  return
}

// ~QOpenGLFramebufferObjectFormat()
func (this *QOpenGLFramebufferObjectFormat) Freeqopenglframebufferobjectformat(args ...interface{}) () {
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
    C.C_ZN30QOpenGLFramebufferObjectFormatD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "~QOpenGLFramebufferObjectFormat", args)
  }

  return
}

// setTextureTarget(GLenum)
func (this *QOpenGLFramebufferObjectFormat) Settexturetarget(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN30QOpenGLFramebufferObjectFormat16setTextureTargetEj(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "setTextureTarget", args)
  }

  return
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
    var arg0 = args[0].(*QOpenGLFramebufferObjectFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QOpenGLFramebufferObjectFormatC2ERKS_(arg0)
    return &QOpenGLFramebufferObjectFormat{Qclsinst:qthis}
  case 1:
    // invoke: _ZN30QOpenGLFramebufferObjectFormatC1Ev
    // invoke: void QOpenGLFramebufferObjectFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN30QOpenGLFramebufferObjectFormatC2Ev()
    return &QOpenGLFramebufferObjectFormat{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "QOpenGLFramebufferObjectFormat", args)
  }

  return nil // QOpenGLFramebufferObjectFormat{Qclsinst:qthis}
}

// internalTextureFormat()
func (this *QOpenGLFramebufferObjectFormat) Internaltextureformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK30QOpenGLFramebufferObjectFormat21internalTextureFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLenum"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObjectFormat", "internalTextureFormat", args)
  }

  return
}

// hasOpenGLFramebufferObjects()
func (this *QOpenGLFramebufferObject) Hasopenglframebufferobjects_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject27hasOpenGLFramebufferObjectsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferObjects", args)
  }

  return
}

// height()
func (this *QOpenGLFramebufferObject) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "height", args)
  }

  return
}

// size()
func (this *QOpenGLFramebufferObject) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "size", args)
  }

  return
}

// hasOpenGLFramebufferBlit()
func (this *QOpenGLFramebufferObject) Hasopenglframebufferblit_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject24hasOpenGLFramebufferBlitEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "hasOpenGLFramebufferBlit", args)
  }

  return
}

// blitFramebuffer(class QOpenGLFramebufferObject *, const class QRect &, class QOpenGLFramebufferObject *, const class QRect &, GLbitfield, GLenum)
func (this *QOpenGLFramebufferObject) Blitframebuffer_S(args ...interface{}) () {
  // blitFramebuffer(class QOpenGLFramebufferObject *, const class QRect &, class QOpenGLFramebufferObject *, const class QRect &, GLbitfield, GLenum)
  // blitFramebuffer(class QOpenGLFramebufferObject *, class QOpenGLFramebufferObject *, GLbitfield, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][2] = reflect.TypeOf(QOpenGLFramebufferObject{}) // "QOpenGLFramebufferObject *"
  vtys[0][3] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
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
    var arg0 = args[0].(*QOpenGLFramebufferObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QOpenGLFramebufferObject).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    C.C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_RK5QRectS0_S3_jj(arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj
    // invoke: void blitFramebuffer(class QOpenGLFramebufferObject *, class QOpenGLFramebufferObject *, GLbitfield, GLenum)
    var arg0 = args[0].(*QOpenGLFramebufferObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QOpenGLFramebufferObject).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN24QOpenGLFramebufferObject15blitFramebufferEPS_S0_jj(arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "blitFramebuffer", args)
  }

  return
}

// format()
func (this *QOpenGLFramebufferObject) Format(args ...interface{}) () {
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
    C.C_ZNK24QOpenGLFramebufferObject6formatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "format", args)
  }

  return
}

// takeTexture()
func (this *QOpenGLFramebufferObject) Taketexture(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject11takeTextureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "takeTexture", args)
  }

  return
}

// texture()
func (this *QOpenGLFramebufferObject) Texture(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject7textureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "texture", args)
  }

  return
}

// width()
func (this *QOpenGLFramebufferObject) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "width", args)
  }

  return
}

// attachment()
func (this *QOpenGLFramebufferObject) Attachment(args ...interface{}) () {
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
    C.C_ZNK24QOpenGLFramebufferObject10attachmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "attachment", args)
  }

  return
}

// ~QOpenGLFramebufferObject()
func (this *QOpenGLFramebufferObject) Freeqopenglframebufferobject(args ...interface{}) () {
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
    C.C_ZN24QOpenGLFramebufferObjectD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "~QOpenGLFramebufferObject", args)
  }

  return
}

// handle()
func (this *QOpenGLFramebufferObject) Handle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject6handleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "handle", args)
  }

  return
}

// isValid()
func (this *QOpenGLFramebufferObject) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isValid", args)
  }

  return
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
  vtys[2][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  vtys[2][1] = reflect.TypeOf(QOpenGLFramebufferObjectFormat{}) // "const QOpenGLFramebufferObjectFormat &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  vtys[3][1] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QOpenGLFramebufferObjectC1Eiij
    // invoke: void QOpenGLFramebufferObject(int, int, GLenum)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2Eiij(arg0, arg1, arg2)
    return &QOpenGLFramebufferObject{Qclsinst:qthis}
  case 1:
    // invoke: _ZN24QOpenGLFramebufferObjectC1EiiRK30QOpenGLFramebufferObjectFormat
    // invoke: void QOpenGLFramebufferObject(int, int, const class QOpenGLFramebufferObjectFormat &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QOpenGLFramebufferObjectFormat).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2EiiRK30QOpenGLFramebufferObjectFormat(arg0, arg1, arg2)
    return &QOpenGLFramebufferObject{Qclsinst:qthis}
  case 2:
    // invoke: _ZN24QOpenGLFramebufferObjectC1ERK5QSizeRK30QOpenGLFramebufferObjectFormat
    // invoke: void QOpenGLFramebufferObject(const class QSize &, const class QOpenGLFramebufferObjectFormat &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QOpenGLFramebufferObjectFormat).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2ERK5QSizeRK30QOpenGLFramebufferObjectFormat(arg0, arg1)
    return &QOpenGLFramebufferObject{Qclsinst:qthis}
  case 3:
    // invoke: _ZN24QOpenGLFramebufferObjectC1ERK5QSizej
    // invoke: void QOpenGLFramebufferObject(const class QSize &, GLenum)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QOpenGLFramebufferObjectC2ERK5QSizej(arg0, arg1)
    return &QOpenGLFramebufferObject{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "QOpenGLFramebufferObject", args)
  }

  return nil // QOpenGLFramebufferObject{Qclsinst:qthis}
}

// bindDefault()
func (this *QOpenGLFramebufferObject) Binddefault_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject11bindDefaultEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bindDefault", args)
  }

  return
}

// isBound()
func (this *QOpenGLFramebufferObject) Isbound(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject7isBoundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "isBound", args)
  }

  return
}

// bind()
func (this *QOpenGLFramebufferObject) Bind(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject4bindEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "bind", args)
  }

  return
}

// toImage(_Bool)
func (this *QOpenGLFramebufferObject) Toimage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject7toImageEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK24QOpenGLFramebufferObject7toImageEv
    // invoke: QImage toImage()
    var ret0 = C.C_ZNK24QOpenGLFramebufferObject7toImageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "toImage", args)
  }

  return
}

// release()
func (this *QOpenGLFramebufferObject) Release(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN24QOpenGLFramebufferObject7releaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLFramebufferObject", "release", args)
  }

  return
}

// <= body block end

