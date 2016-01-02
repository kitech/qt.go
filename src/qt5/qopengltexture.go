package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qopengltexture.h
// dst-file: /src/gui/qopengltexture.go
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
  // proto:  void QOpenGLTexture::bind();
extern void _ZN14QOpenGLTexture4bindEv(void* qthis);
  // proto:  void QOpenGLTexture::setFixedSamplePositions(bool fixed);
extern void _ZN14QOpenGLTexture23setFixedSamplePositionsEb(void* qthis, bool arg0);
  // proto:  int QOpenGLTexture::height();
extern void _ZNK14QOpenGLTexture6heightEv(void* qthis);
  // proto:  bool QOpenGLTexture::isAutoMipMapGenerationEnabled();
extern void _ZNK14QOpenGLTexture29isAutoMipMapGenerationEnabledEv(void* qthis);
  // proto:  void QOpenGLTexture::setCompressedData(int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  void QOpenGLTexture::setMaximumLevelOfDetail(float value);
extern void _ZN14QOpenGLTexture23setMaximumLevelOfDetailEf(void* qthis, float arg0);
  // proto:  void QOpenGLTexture::setAutoMipMapGenerationEnabled(bool enabled);
extern void _ZN14QOpenGLTexture30setAutoMipMapGenerationEnabledEb(void* qthis, bool arg0);
  // proto:  int QOpenGLTexture::depth();
extern void _ZNK14QOpenGLTexture5depthEv(void* qthis);
  // proto:  void QOpenGLTexture::generateMipMaps(int baseLevel, bool resetBaseLevel);
extern void _ZN14QOpenGLTexture15generateMipMapsEib(void* qthis, int arg0, bool arg1);
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, int arg1, void* arg2, void* arg3);
  // proto:  void QOpenGLTexture::setMipBaseLevel(int baseLevel);
extern void _ZN14QOpenGLTexture15setMipBaseLevelEi(void* qthis, int arg0);
  // proto:  QPair<float, float> QOpenGLTexture::levelOfDetailRange();
extern void _ZNK14QOpenGLTexture18levelOfDetailRangeEv(void* qthis);
  // proto:  bool QOpenGLTexture::create();
extern void _ZN14QOpenGLTexture6createEv(void* qthis);
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int layer, int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, int arg1, int arg2, void* arg3, void* arg4);
  // proto:  void QOpenGLTexture::setLevelOfDetailRange(float min, float max);
extern void _ZN14QOpenGLTexture21setLevelOfDetailRangeEff(void* qthis, float arg0, float arg1);
  // proto:  void QOpenGLTexture::borderColor(unsigned int * border);
extern void _ZNK14QOpenGLTexture11borderColorEPj(void* qthis, unsigned int* arg0);
  // proto:  bool QOpenGLTexture::isStorageAllocated();
extern void _ZNK14QOpenGLTexture18isStorageAllocatedEv(void* qthis);
  // proto:  void QOpenGLTexture::borderColor(int * border);
extern void _ZNK14QOpenGLTexture11borderColorEPi(void* qthis, int* arg0);
  // proto:  bool QOpenGLTexture::isTextureView();
extern void _ZNK14QOpenGLTexture13isTextureViewEv(void* qthis);
  // proto:  bool QOpenGLTexture::isFixedSamplePositions();
extern void _ZNK14QOpenGLTexture22isFixedSamplePositionsEv(void* qthis);
  // proto:  int QOpenGLTexture::faces();
extern void _ZNK14QOpenGLTexture5facesEv(void* qthis);
  // proto:  void QOpenGLTexture::setLayers(int layers);
extern void _ZN14QOpenGLTexture9setLayersEi(void* qthis, int arg0);
  // proto:  void QOpenGLTexture::setCompressedData(int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  int QOpenGLTexture::width();
extern void _ZNK14QOpenGLTexture5widthEv(void* qthis);
  // proto:  int QOpenGLTexture::layers();
extern void _ZNK14QOpenGLTexture6layersEv(void* qthis);
  // proto:  void QOpenGLTexture::borderColor(float * border);
extern void _ZNK14QOpenGLTexture11borderColorEPf(void* qthis, float* arg0);
  // proto:  float QOpenGLTexture::minimumLevelOfDetail();
extern void _ZNK14QOpenGLTexture20minimumLevelOfDetailEv(void* qthis);
  // proto:  void QOpenGLTexture::setBorderColor(int r, int g, int b, int a);
extern void _ZN14QOpenGLTexture14setBorderColorEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QOpenGLTexture::setMinimumLevelOfDetail(float value);
extern void _ZN14QOpenGLTexture23setMinimumLevelOfDetailEf(void* qthis, float arg0);
  // proto:  void QOpenGLTexture::setMipLevels(int levels);
extern void _ZN14QOpenGLTexture12setMipLevelsEi(void* qthis, int arg0);
  // proto:  QPair<int, int> QOpenGLTexture::mipLevelRange();
extern void _ZNK14QOpenGLTexture13mipLevelRangeEv(void* qthis);
  // proto:  void QOpenGLTexture::setMipMaxLevel(int maxLevel);
extern void _ZN14QOpenGLTexture14setMipMaxLevelEi(void* qthis, int arg0);
  // proto:  float QOpenGLTexture::levelofDetailBias();
extern void _ZNK14QOpenGLTexture17levelofDetailBiasEv(void* qthis);
  // proto:  int QOpenGLTexture::maximumMipLevels();
extern void _ZNK14QOpenGLTexture16maximumMipLevelsEv(void* qthis);
  // proto:  bool QOpenGLTexture::isBound(uint unit);
extern void _ZN14QOpenGLTexture7isBoundEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLTexture::setBorderColor(uint r, uint g, uint b, uint a);
extern void _ZN14QOpenGLTexture14setBorderColorEjjjj(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, unsigned int arg3);
  // proto:  void QOpenGLTexture::setMaximumAnisotropy(float anisotropy);
extern void _ZN14QOpenGLTexture20setMaximumAnisotropyEf(void* qthis, float arg0);
  // proto:  void QOpenGLTexture::setSamples(int samples);
extern void _ZN14QOpenGLTexture10setSamplesEi(void* qthis, int arg0);
  // proto:  int QOpenGLTexture::mipLevels();
extern void _ZNK14QOpenGLTexture9mipLevelsEv(void* qthis);
  // proto:  void QOpenGLTexture::setLevelofDetailBias(float bias);
extern void _ZN14QOpenGLTexture20setLevelofDetailBiasEf(void* qthis, float arg0);
  // proto:  GLuint QOpenGLTexture::textureId();
extern void _ZNK14QOpenGLTexture9textureIdEv(void* qthis);
  // proto:  void QOpenGLTexture::setMipLevelRange(int baseLevel, int maxLevel);
extern void _ZN14QOpenGLTexture16setMipLevelRangeEii(void* qthis, int arg0, int arg1);
  // proto:  void QOpenGLTexture::allocateStorage();
extern void _ZN14QOpenGLTexture15allocateStorageEv(void* qthis);
  // proto:  void QOpenGLTexture::~QOpenGLTexture();
extern void _ZN14QOpenGLTextureD0Ev(void* qthis);
  // proto:  int QOpenGLTexture::mipMaxLevel();
extern void _ZNK14QOpenGLTexture11mipMaxLevelEv(void* qthis);
  // proto:  void QOpenGLTexture::setBorderColor(QColor color);
extern void _ZN14QOpenGLTexture14setBorderColorE6QColor(void* qthis, void* arg0);
  // proto:  void QOpenGLTexture::destroy();
extern void _ZN14QOpenGLTexture7destroyEv(void* qthis);
  // proto:  void QOpenGLTexture::generateMipMaps();
extern void _ZN14QOpenGLTexture15generateMipMapsEv(void* qthis);
  // proto:  void QOpenGLTexture::release();
extern void _ZN14QOpenGLTexture7releaseEv(void* qthis);
  // proto:  float QOpenGLTexture::maximumAnisotropy();
extern void _ZNK14QOpenGLTexture17maximumAnisotropyEv(void* qthis);
  // proto:  float QOpenGLTexture::maximumLevelOfDetail();
extern void _ZNK14QOpenGLTexture20maximumLevelOfDetailEv(void* qthis);
  // proto:  void QOpenGLTexture::setSize(int width, int height, int depth);
extern void _ZN14QOpenGLTexture7setSizeEiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  bool QOpenGLTexture::isCreated();
extern void _ZNK14QOpenGLTexture9isCreatedEv(void* qthis);
  // proto:  bool QOpenGLTexture::isBound();
extern void _ZNK14QOpenGLTexture7isBoundEv(void* qthis);
  // proto:  void QOpenGLTexture::setBorderColor(float r, float g, float b, float a);
extern void _ZN14QOpenGLTexture14setBorderColorEffff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  int QOpenGLTexture::samples();
extern void _ZNK14QOpenGLTexture7samplesEv(void* qthis);
  // proto:  int QOpenGLTexture::mipBaseLevel();
extern void _ZNK14QOpenGLTexture12mipBaseLevelEv(void* qthis);
  // proto:  QColor QOpenGLTexture::borderColor();
extern void _ZNK14QOpenGLTexture11borderColorEv(void* qthis);
  // proto:  void QOpenGLTexture::QOpenGLTexture(const QOpenGLTexture & );
extern void* dector_ZN14QOpenGLTextureC1ERKS_(void* arg0);
extern void _ZN14QOpenGLTextureC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int layer, int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, int arg1, int arg2, void* arg3, void* arg4);
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void _ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions(void* qthis, int arg0, int arg1, void* arg2, void* arg3);
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

// class sizeof(QOpenGLTexture)=1
type QOpenGLTexture struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QOpenGLTexture::bind();
func (this *QOpenGLTexture) bind(args ...interface{}) () {
  // bind()
  // bind(uint, enum QOpenGLTexture::TextureUnitReset)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1][1] = qtrt.Int32Ty(false) // "QOpenGLTexture::TextureUnitReset"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture4bindEv
    // invoke: void bind()
    C._ZN14QOpenGLTexture4bindEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "bind", args)
  }

}

  // proto:  void QOpenGLTexture::setFixedSamplePositions(bool fixed);
func (this *QOpenGLTexture) setFixedSamplePositions(args ...interface{}) () {
  // setFixedSamplePositions(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture23setFixedSamplePositionsEb
    // invoke: void setFixedSamplePositions(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture23setFixedSamplePositionsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setFixedSamplePositions", args)
  }

}

  // proto:  int QOpenGLTexture::height();
func (this *QOpenGLTexture) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture6heightEv
    // invoke: int height()
    C._ZNK14QOpenGLTexture6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "height", args)
  }

}

  // proto:  bool QOpenGLTexture::isAutoMipMapGenerationEnabled();
func (this *QOpenGLTexture) isAutoMipMapGenerationEnabled(args ...interface{}) () {
  // isAutoMipMapGenerationEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture29isAutoMipMapGenerationEnabledEv
    // invoke: bool isAutoMipMapGenerationEnabled()
    C._ZNK14QOpenGLTexture29isAutoMipMapGenerationEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isAutoMipMapGenerationEnabled", args)
  }

}

  // proto:  void QOpenGLTexture::setCompressedData(int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
func (this *QOpenGLTexture) setCompressedData(args ...interface{}) () {
  // setCompressedData(int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, enum QOpenGLTexture::CubeMapFace, int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, enum QOpenGLTexture::CubeMapFace, int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, void *, const class QOpenGLPixelTransferOptions *const)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "const void *"
  vtys[0][2] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "QOpenGLTexture::CubeMapFace"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.VoidpTy() // "void *"
  vtys[1][5] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.VoidpTy() // "const void *"
  vtys[2][3] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.VoidpTy() // "const void *"
  vtys[3][4] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.VoidpTy() // "void *"
  vtys[4][2] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "QOpenGLTexture::CubeMapFace"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = qtrt.VoidpTy() // "const void *"
  vtys[5][5] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[6][3] = qtrt.VoidpTy() // "void *"
  vtys[6][4] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.VoidpTy() // "void *"
  vtys[7][3] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg4)}
    C._ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 3:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg4)}
    C._ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 5:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setCompressedData", args)
  }

}

  // proto:  void QOpenGLTexture::setMaximumLevelOfDetail(float value);
func (this *QOpenGLTexture) setMaximumLevelOfDetail(args ...interface{}) () {
  // setMaximumLevelOfDetail(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture23setMaximumLevelOfDetailEf
    // invoke: void setMaximumLevelOfDetail(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture23setMaximumLevelOfDetailEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumLevelOfDetail", args)
  }

}

  // proto:  void QOpenGLTexture::setAutoMipMapGenerationEnabled(bool enabled);
func (this *QOpenGLTexture) setAutoMipMapGenerationEnabled(args ...interface{}) () {
  // setAutoMipMapGenerationEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture30setAutoMipMapGenerationEnabledEb
    // invoke: void setAutoMipMapGenerationEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture30setAutoMipMapGenerationEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setAutoMipMapGenerationEnabled", args)
  }

}

  // proto:  int QOpenGLTexture::depth();
func (this *QOpenGLTexture) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture5depthEv
    // invoke: int depth()
    C._ZNK14QOpenGLTexture5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "depth", args)
  }

}

  // proto:  void QOpenGLTexture::generateMipMaps(int baseLevel, bool resetBaseLevel);
func (this *QOpenGLTexture) generateMipMaps(args ...interface{}) () {
  // generateMipMaps(int, _Bool)
  // generateMipMaps()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture15generateMipMapsEib
    // invoke: void generateMipMaps(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
    C._ZN14QOpenGLTexture15generateMipMapsEib(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN14QOpenGLTexture15generateMipMapsEv
    // invoke: void generateMipMaps()
    C._ZN14QOpenGLTexture15generateMipMapsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "generateMipMaps", args)
  }

}

  // proto:  void QOpenGLTexture::setMipBaseLevel(int baseLevel);
func (this *QOpenGLTexture) setMipBaseLevel(args ...interface{}) () {
  // setMipBaseLevel(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture15setMipBaseLevelEi
    // invoke: void setMipBaseLevel(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture15setMipBaseLevelEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipBaseLevel", args)
  }

}

  // proto:  QPair<float, float> QOpenGLTexture::levelOfDetailRange();
func (this *QOpenGLTexture) levelOfDetailRange(args ...interface{}) () {
  // levelOfDetailRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture18levelOfDetailRangeEv
    // invoke: QPair<float, float> levelOfDetailRange()
    C._ZNK14QOpenGLTexture18levelOfDetailRangeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelOfDetailRange", args)
  }

}

  // proto:  bool QOpenGLTexture::create();
func (this *QOpenGLTexture) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture6createEv
    // invoke: bool create()
    C._ZN14QOpenGLTexture6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "create", args)
  }

}

  // proto:  void QOpenGLTexture::setLevelOfDetailRange(float min, float max);
func (this *QOpenGLTexture) setLevelOfDetailRange(args ...interface{}) () {
  // setLevelOfDetailRange(float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture21setLevelOfDetailRangeEff
    // invoke: void setLevelOfDetailRange(float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    C._ZN14QOpenGLTexture21setLevelOfDetailRangeEff(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelOfDetailRange", args)
  }

}

  // proto:  void QOpenGLTexture::borderColor(unsigned int * border);
func (this *QOpenGLTexture) borderColor(args ...interface{}) () {
  // borderColor(unsigned int *)
  // borderColor(int *)
  // borderColor(float *)
  // borderColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "unsigned int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(true) // "int *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(true) // "float *"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPj
    // invoke: void borderColor(unsigned int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLTexture11borderColorEPj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPi
    // invoke: void borderColor(int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLTexture11borderColorEPi(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPf
    // invoke: void borderColor(float *)
    var arg0 = (*C.float)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLTexture11borderColorEPf(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK14QOpenGLTexture11borderColorEv
    // invoke: QColor borderColor()
    C._ZNK14QOpenGLTexture11borderColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "borderColor", args)
  }

}

  // proto:  bool QOpenGLTexture::isStorageAllocated();
func (this *QOpenGLTexture) isStorageAllocated(args ...interface{}) () {
  // isStorageAllocated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture18isStorageAllocatedEv
    // invoke: bool isStorageAllocated()
    C._ZNK14QOpenGLTexture18isStorageAllocatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isStorageAllocated", args)
  }

}

  // proto:  bool QOpenGLTexture::isTextureView();
func (this *QOpenGLTexture) isTextureView(args ...interface{}) () {
  // isTextureView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture13isTextureViewEv
    // invoke: bool isTextureView()
    C._ZNK14QOpenGLTexture13isTextureViewEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isTextureView", args)
  }

}

  // proto:  bool QOpenGLTexture::isFixedSamplePositions();
func (this *QOpenGLTexture) isFixedSamplePositions(args ...interface{}) () {
  // isFixedSamplePositions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture22isFixedSamplePositionsEv
    // invoke: bool isFixedSamplePositions()
    C._ZNK14QOpenGLTexture22isFixedSamplePositionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isFixedSamplePositions", args)
  }

}

  // proto:  int QOpenGLTexture::faces();
func (this *QOpenGLTexture) faces(args ...interface{}) () {
  // faces()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture5facesEv
    // invoke: int faces()
    C._ZNK14QOpenGLTexture5facesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "faces", args)
  }

}

  // proto:  void QOpenGLTexture::setLayers(int layers);
func (this *QOpenGLTexture) setLayers(args ...interface{}) () {
  // setLayers(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture9setLayersEi
    // invoke: void setLayers(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture9setLayersEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLayers", args)
  }

}

  // proto:  int QOpenGLTexture::width();
func (this *QOpenGLTexture) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture5widthEv
    // invoke: int width()
    C._ZNK14QOpenGLTexture5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "width", args)
  }

}

  // proto:  int QOpenGLTexture::layers();
func (this *QOpenGLTexture) layers(args ...interface{}) () {
  // layers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture6layersEv
    // invoke: int layers()
    C._ZNK14QOpenGLTexture6layersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "layers", args)
  }

}

  // proto:  float QOpenGLTexture::minimumLevelOfDetail();
func (this *QOpenGLTexture) minimumLevelOfDetail(args ...interface{}) () {
  // minimumLevelOfDetail()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture20minimumLevelOfDetailEv
    // invoke: float minimumLevelOfDetail()
    C._ZNK14QOpenGLTexture20minimumLevelOfDetailEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "minimumLevelOfDetail", args)
  }

}

  // proto:  void QOpenGLTexture::setBorderColor(int r, int g, int b, int a);
func (this *QOpenGLTexture) setBorderColor(args ...interface{}) () {
  // setBorderColor(int, int, int, int)
  // setBorderColor(uint, uint, uint, uint)
  // setBorderColor(class QColor)
  // setBorderColor(float, float, float, float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1][1] = qtrt.Int32Ty(false) // "uint"
  vtys[1][2] = qtrt.Int32Ty(false) // "uint"
  vtys[1][3] = qtrt.Int32Ty(false) // "uint"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QColor{}) // "QColor"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.FloatTy(false) // "float"
  vtys[3][1] = qtrt.FloatTy(false) // "float"
  vtys[3][2] = qtrt.FloatTy(false) // "float"
  vtys[3][3] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEiiii
    // invoke: void setBorderColor(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN14QOpenGLTexture14setBorderColorEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEjjjj
    // invoke: void setBorderColor(uint, uint, uint, uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN14QOpenGLTexture14setBorderColorEjjjj(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN14QOpenGLTexture14setBorderColorE6QColor
    // invoke: void setBorderColor(class QColor)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture14setBorderColorE6QColor(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEffff
    // invoke: void setBorderColor(float, float, float, float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    C._ZN14QOpenGLTexture14setBorderColorEffff(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setBorderColor", args)
  }

}

  // proto:  void QOpenGLTexture::setMinimumLevelOfDetail(float value);
func (this *QOpenGLTexture) setMinimumLevelOfDetail(args ...interface{}) () {
  // setMinimumLevelOfDetail(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture23setMinimumLevelOfDetailEf
    // invoke: void setMinimumLevelOfDetail(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture23setMinimumLevelOfDetailEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMinimumLevelOfDetail", args)
  }

}

  // proto:  void QOpenGLTexture::setMipLevels(int levels);
func (this *QOpenGLTexture) setMipLevels(args ...interface{}) () {
  // setMipLevels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture12setMipLevelsEi
    // invoke: void setMipLevels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture12setMipLevelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevels", args)
  }

}

  // proto:  QPair<int, int> QOpenGLTexture::mipLevelRange();
func (this *QOpenGLTexture) mipLevelRange(args ...interface{}) () {
  // mipLevelRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture13mipLevelRangeEv
    // invoke: QPair<int, int> mipLevelRange()
    C._ZNK14QOpenGLTexture13mipLevelRangeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevelRange", args)
  }

}

  // proto:  void QOpenGLTexture::setMipMaxLevel(int maxLevel);
func (this *QOpenGLTexture) setMipMaxLevel(args ...interface{}) () {
  // setMipMaxLevel(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture14setMipMaxLevelEi
    // invoke: void setMipMaxLevel(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture14setMipMaxLevelEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipMaxLevel", args)
  }

}

  // proto:  float QOpenGLTexture::levelofDetailBias();
func (this *QOpenGLTexture) levelofDetailBias(args ...interface{}) () {
  // levelofDetailBias()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture17levelofDetailBiasEv
    // invoke: float levelofDetailBias()
    C._ZNK14QOpenGLTexture17levelofDetailBiasEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelofDetailBias", args)
  }

}

  // proto:  int QOpenGLTexture::maximumMipLevels();
func (this *QOpenGLTexture) maximumMipLevels(args ...interface{}) () {
  // maximumMipLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture16maximumMipLevelsEv
    // invoke: int maximumMipLevels()
    C._ZNK14QOpenGLTexture16maximumMipLevelsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumMipLevels", args)
  }

}

  // proto:  bool QOpenGLTexture::isBound(uint unit);
func (this *QOpenGLTexture) isBound(args ...interface{}) () {
  // isBound(uint)
  // isBound()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture7isBoundEj
    // invoke: bool isBound(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture7isBoundEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QOpenGLTexture7isBoundEv
    // invoke: bool isBound()
    C._ZNK14QOpenGLTexture7isBoundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isBound", args)
  }

}

  // proto:  void QOpenGLTexture::setMaximumAnisotropy(float anisotropy);
func (this *QOpenGLTexture) setMaximumAnisotropy(args ...interface{}) () {
  // setMaximumAnisotropy(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture20setMaximumAnisotropyEf
    // invoke: void setMaximumAnisotropy(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture20setMaximumAnisotropyEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumAnisotropy", args)
  }

}

  // proto:  void QOpenGLTexture::setSamples(int samples);
func (this *QOpenGLTexture) setSamples(args ...interface{}) () {
  // setSamples(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture10setSamplesEi
    // invoke: void setSamples(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture10setSamplesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSamples", args)
  }

}

  // proto:  int QOpenGLTexture::mipLevels();
func (this *QOpenGLTexture) mipLevels(args ...interface{}) () {
  // mipLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture9mipLevelsEv
    // invoke: int mipLevels()
    C._ZNK14QOpenGLTexture9mipLevelsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevels", args)
  }

}

  // proto:  void QOpenGLTexture::setLevelofDetailBias(float bias);
func (this *QOpenGLTexture) setLevelofDetailBias(args ...interface{}) () {
  // setLevelofDetailBias(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture20setLevelofDetailBiasEf
    // invoke: void setLevelofDetailBias(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLTexture20setLevelofDetailBiasEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelofDetailBias", args)
  }

}

  // proto:  GLuint QOpenGLTexture::textureId();
func (this *QOpenGLTexture) textureId(args ...interface{}) () {
  // textureId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture9textureIdEv
    // invoke: GLuint textureId()
    C._ZNK14QOpenGLTexture9textureIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "textureId", args)
  }

}

  // proto:  void QOpenGLTexture::setMipLevelRange(int baseLevel, int maxLevel);
func (this *QOpenGLTexture) setMipLevelRange(args ...interface{}) () {
  // setMipLevelRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture16setMipLevelRangeEii
    // invoke: void setMipLevelRange(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN14QOpenGLTexture16setMipLevelRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevelRange", args)
  }

}

  // proto:  void QOpenGLTexture::allocateStorage();
func (this *QOpenGLTexture) allocateStorage(args ...interface{}) () {
  // allocateStorage(enum QOpenGLTexture::PixelFormat, enum QOpenGLTexture::PixelType)
  // allocateStorage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QOpenGLTexture::PixelFormat"
  vtys[0][1] = qtrt.Int32Ty(false) // "QOpenGLTexture::PixelType"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture15allocateStorageEv
    // invoke: void allocateStorage()
    C._ZN14QOpenGLTexture15allocateStorageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "allocateStorage", args)
  }

}

  // proto:  void QOpenGLTexture::~QOpenGLTexture();
func (this *QOpenGLTexture) FreeQOpenGLTexture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "~QOpenGLTexture", args)
  }

}

  // proto:  int QOpenGLTexture::mipMaxLevel();
func (this *QOpenGLTexture) mipMaxLevel(args ...interface{}) () {
  // mipMaxLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture11mipMaxLevelEv
    // invoke: int mipMaxLevel()
    C._ZNK14QOpenGLTexture11mipMaxLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipMaxLevel", args)
  }

}

  // proto:  void QOpenGLTexture::destroy();
func (this *QOpenGLTexture) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture7destroyEv
    // invoke: void destroy()
    C._ZN14QOpenGLTexture7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "destroy", args)
  }

}

  // proto:  void QOpenGLTexture::release();
func (this *QOpenGLTexture) release(args ...interface{}) () {
  // release(uint, enum QOpenGLTexture::TextureUnitReset)
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = qtrt.Int32Ty(false) // "QOpenGLTexture::TextureUnitReset"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture7releaseEv
    // invoke: void release()
    C._ZN14QOpenGLTexture7releaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "release", args)
  }

}

  // proto:  float QOpenGLTexture::maximumAnisotropy();
func (this *QOpenGLTexture) maximumAnisotropy(args ...interface{}) () {
  // maximumAnisotropy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture17maximumAnisotropyEv
    // invoke: float maximumAnisotropy()
    C._ZNK14QOpenGLTexture17maximumAnisotropyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumAnisotropy", args)
  }

}

  // proto:  float QOpenGLTexture::maximumLevelOfDetail();
func (this *QOpenGLTexture) maximumLevelOfDetail(args ...interface{}) () {
  // maximumLevelOfDetail()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture20maximumLevelOfDetailEv
    // invoke: float maximumLevelOfDetail()
    C._ZNK14QOpenGLTexture20maximumLevelOfDetailEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumLevelOfDetail", args)
  }

}

  // proto:  void QOpenGLTexture::setSize(int width, int height, int depth);
func (this *QOpenGLTexture) setSize(args ...interface{}) () {
  // setSize(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture7setSizeEiii
    // invoke: void setSize(int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN14QOpenGLTexture7setSizeEiii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSize", args)
  }

}

  // proto:  bool QOpenGLTexture::isCreated();
func (this *QOpenGLTexture) isCreated(args ...interface{}) () {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture9isCreatedEv
    // invoke: bool isCreated()
    C._ZNK14QOpenGLTexture9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isCreated", args)
  }

}

  // proto:  int QOpenGLTexture::samples();
func (this *QOpenGLTexture) samples(args ...interface{}) () {
  // samples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture7samplesEv
    // invoke: int samples()
    C._ZNK14QOpenGLTexture7samplesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "samples", args)
  }

}

  // proto:  int QOpenGLTexture::mipBaseLevel();
func (this *QOpenGLTexture) mipBaseLevel(args ...interface{}) () {
  // mipBaseLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture12mipBaseLevelEv
    // invoke: int mipBaseLevel()
    C._ZNK14QOpenGLTexture12mipBaseLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipBaseLevel", args)
  }

}

  // proto:  void QOpenGLTexture::QOpenGLTexture(const QOpenGLTexture & );
func NewQOpenGLTexture(args ...interface{}) QOpenGLTexture {
  return QOpenGLTexture{}
}

// <= body block end

