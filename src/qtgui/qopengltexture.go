package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
  // proto:  float QOpenGLTexture::maximumAnisotropy();
extern float C_ZNK14QOpenGLTexture17maximumAnisotropyEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::maximumMipLevels();
extern int32_t C_ZNK14QOpenGLTexture16maximumMipLevelsEv(void* qthis); // 4
  // proto:  QOpenGLTexture::TextureFormat QOpenGLTexture::format();
extern void C_ZNK14QOpenGLTexture6formatEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::isTextureView();
extern bool C_ZNK14QOpenGLTexture13isTextureViewEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::height();
extern int32_t C_ZNK14QOpenGLTexture6heightEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMinimumLevelOfDetail(float value);
extern void C_ZN14QOpenGLTexture23setMinimumLevelOfDetailEf(void* qthis, float arg0); // 4
  // proto:  QOpenGLTexture::ComparisonFunction QOpenGLTexture::comparisonFunction();
extern void C_ZNK14QOpenGLTexture18comparisonFunctionEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::allocateStorage();
extern void C_ZN14QOpenGLTexture15allocateStorageEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::isFixedSamplePositions();
extern bool C_ZNK14QOpenGLTexture22isFixedSamplePositionsEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::borderColor(float * border);
extern void C_ZNK14QOpenGLTexture11borderColorEPf(void* qthis, void* arg0); // 4
  // proto:  QColor QOpenGLTexture::borderColor();
extern void* C_ZNK14QOpenGLTexture11borderColorEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::borderColor(unsigned int * border);
extern void C_ZNK14QOpenGLTexture11borderColorEPj(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLTexture::borderColor(int * border);
extern void C_ZNK14QOpenGLTexture11borderColorEPi(void* qthis, void* arg0); // 4
  // proto:  QOpenGLTexture::Filter QOpenGLTexture::minificationFilter();
extern void C_ZNK14QOpenGLTexture18minificationFilterEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::isCreated();
extern bool C_ZNK14QOpenGLTexture9isCreatedEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setSamples(int samples);
extern void C_ZN14QOpenGLTexture10setSamplesEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLTexture::setAutoMipMapGenerationEnabled(bool enabled);
extern void C_ZN14QOpenGLTexture30setAutoMipMapGenerationEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QOpenGLTexture::Filter QOpenGLTexture::magnificationFilter();
extern void C_ZNK14QOpenGLTexture19magnificationFilterEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::create();
extern bool C_ZN14QOpenGLTexture6createEv(void* qthis); // 4
  // proto:  float QOpenGLTexture::minimumLevelOfDetail();
extern float C_ZNK14QOpenGLTexture20minimumLevelOfDetailEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setLevelofDetailBias(float bias);
extern void C_ZN14QOpenGLTexture20setLevelofDetailBiasEf(void* qthis, float arg0); // 4
  // proto:  float QOpenGLTexture::maximumLevelOfDetail();
extern float C_ZNK14QOpenGLTexture20maximumLevelOfDetailEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::width();
extern int32_t C_ZNK14QOpenGLTexture5widthEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMipLevelRange(int baseLevel, int maxLevel);
extern void C_ZN14QOpenGLTexture16setMipLevelRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QOpenGLTexture::samples();
extern int32_t C_ZNK14QOpenGLTexture7samplesEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::isStorageAllocated();
extern bool C_ZNK14QOpenGLTexture18isStorageAllocatedEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setLevelOfDetailRange(float min, float max);
extern void C_ZN14QOpenGLTexture21setLevelOfDetailRangeEff(void* qthis, float arg0, float arg1); // 4
  // proto:  GLuint QOpenGLTexture::textureId();
extern int32_t C_ZNK14QOpenGLTexture9textureIdEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::~QOpenGLTexture();
extern void C_ZN14QOpenGLTextureD2Ev(void* qthis); // 4
  // proto:  void QOpenGLTexture::destroy();
extern void C_ZN14QOpenGLTexture7destroyEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMipBaseLevel(int baseLevel);
extern void C_ZN14QOpenGLTexture15setMipBaseLevelEi(void* qthis, int32_t arg0); // 4
  // proto:  int QOpenGLTexture::layers();
extern int32_t C_ZNK14QOpenGLTexture6layersEv(void* qthis); // 4
  // proto:  QPair<float, float> QOpenGLTexture::levelOfDetailRange();
extern void C_ZNK14QOpenGLTexture18levelOfDetailRangeEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setBorderColor(QColor color);
extern void C_ZN14QOpenGLTexture14setBorderColorE6QColor(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLTexture::setBorderColor(int r, int g, int b, int a);
extern void C_ZN14QOpenGLTexture14setBorderColorEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLTexture::setBorderColor(float r, float g, float b, float a);
extern void C_ZN14QOpenGLTexture14setBorderColorEffff(void* qthis, float arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QOpenGLTexture::setBorderColor(uint r, uint g, uint b, uint a);
extern void C_ZN14QOpenGLTexture14setBorderColorEjjjj(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QOpenGLTexture::mipBaseLevel();
extern int32_t C_ZNK14QOpenGLTexture12mipBaseLevelEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setFixedSamplePositions(bool fixed);
extern void C_ZN14QOpenGLTexture23setFixedSamplePositionsEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLTexture::setMipMaxLevel(int maxLevel);
extern void C_ZN14QOpenGLTexture14setMipMaxLevelEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int layer, int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3, void* arg4); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int dataSize, void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3); // 4
  // proto:  void QOpenGLTexture::setCompressedData(int mipLevel, int layer, int dataSize, const void * data, const QOpenGLPixelTransferOptions *const options);
extern void C_ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3, void* arg4); // 4
  // proto:  bool QOpenGLTexture::isAutoMipMapGenerationEnabled();
extern bool C_ZNK14QOpenGLTexture29isAutoMipMapGenerationEnabledEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::mipLevels();
extern int32_t C_ZNK14QOpenGLTexture9mipLevelsEv(void* qthis); // 4
  // proto:  float QOpenGLTexture::levelofDetailBias();
extern float C_ZNK14QOpenGLTexture17levelofDetailBiasEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMaximumLevelOfDetail(float value);
extern void C_ZN14QOpenGLTexture23setMaximumLevelOfDetailEf(void* qthis, float arg0); // 4
  // proto:  QOpenGLTexture::ComparisonMode QOpenGLTexture::comparisonMode();
extern void C_ZNK14QOpenGLTexture14comparisonModeEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMaximumAnisotropy(float anisotropy);
extern void C_ZN14QOpenGLTexture20setMaximumAnisotropyEf(void* qthis, float arg0); // 4
  // proto:  void QOpenGLTexture::setLayers(int layers);
extern void C_ZN14QOpenGLTexture9setLayersEi(void* qthis, int32_t arg0); // 4
  // proto:  QOpenGLTexture::DepthStencilMode QOpenGLTexture::depthStencilMode();
extern void C_ZNK14QOpenGLTexture16depthStencilModeEv(void* qthis); // 4
  // proto:  QOpenGLTexture::Target QOpenGLTexture::target();
extern void C_ZNK14QOpenGLTexture6targetEv(void* qthis); // 4
  // proto:  QPair<int, int> QOpenGLTexture::mipLevelRange();
extern void C_ZNK14QOpenGLTexture13mipLevelRangeEv(void* qthis); // 4
  // proto:  bool QOpenGLTexture::isBound(uint unit);
extern bool C_ZN14QOpenGLTexture7isBoundEj(void* qthis, int32_t arg0); // 4
  // proto:  bool QOpenGLTexture::isBound();
extern bool C_ZNK14QOpenGLTexture7isBoundEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::bind();
extern void C_ZN14QOpenGLTexture4bindEv(void* qthis); // 4
  // proto:  QPair<QOpenGLTexture::Filter, QOpenGLTexture::Filter> QOpenGLTexture::minMagFilters();
extern void C_ZNK14QOpenGLTexture13minMagFiltersEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::mipMaxLevel();
extern int32_t C_ZNK14QOpenGLTexture11mipMaxLevelEv(void* qthis); // 4
  // proto:  int QOpenGLTexture::depth();
extern int32_t C_ZNK14QOpenGLTexture5depthEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setMipLevels(int levels);
extern void C_ZN14QOpenGLTexture12setMipLevelsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QOpenGLTexture::faces();
extern int32_t C_ZNK14QOpenGLTexture5facesEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::release();
extern void C_ZN14QOpenGLTexture7releaseEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::generateMipMaps(int baseLevel, bool resetBaseLevel);
extern void C_ZN14QOpenGLTexture15generateMipMapsEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QOpenGLTexture::generateMipMaps();
extern void C_ZN14QOpenGLTexture15generateMipMapsEv(void* qthis); // 4
  // proto:  void QOpenGLTexture::setSize(int width, int height, int depth);
extern void C_ZN14QOpenGLTexture7setSizeEiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QOpenGLTexture)=1
type QOpenGLTexture struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// maximumAnisotropy()
func (this *QOpenGLTexture) MaximumAnisotropy(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture17maximumAnisotropyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumAnisotropy", args)
  }

  return
}

// maximumMipLevels()
func (this *QOpenGLTexture) MaximumMipLevels(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture16maximumMipLevelsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumMipLevels", args)
  }

  return
}

// format()
func (this *QOpenGLTexture) Format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture6formatEv
    // invoke: QOpenGLTexture::TextureFormat format()
    C.C_ZNK14QOpenGLTexture6formatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "format", args)
  }

  return
}

// isTextureView()
func (this *QOpenGLTexture) IsTextureView(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture13isTextureViewEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isTextureView", args)
  }

  return
}

// height()
func (this *QOpenGLTexture) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "height", args)
  }

  return
}

// setMinimumLevelOfDetail(float)
func (this *QOpenGLTexture) SetMinimumLevelOfDetail(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture23setMinimumLevelOfDetailEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMinimumLevelOfDetail", args)
  }

  return
}

// comparisonFunction()
func (this *QOpenGLTexture) ComparisonFunction(args ...interface{}) () {
  // comparisonFunction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture18comparisonFunctionEv
    // invoke: QOpenGLTexture::ComparisonFunction comparisonFunction()
    C.C_ZNK14QOpenGLTexture18comparisonFunctionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "comparisonFunction", args)
  }

  return
}

// allocateStorage()
func (this *QOpenGLTexture) AllocateStorage(args ...interface{}) () {
  // allocateStorage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture15allocateStorageEv
    // invoke: void allocateStorage()
    C.C_ZN14QOpenGLTexture15allocateStorageEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "allocateStorage", args)
  }

  return
}

// isFixedSamplePositions()
func (this *QOpenGLTexture) IsFixedSamplePositions(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture22isFixedSamplePositionsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isFixedSamplePositions", args)
  }

  return
}

// borderColor(float *)
func (this *QOpenGLTexture) BorderColor(args ...interface{}) () {
  // borderColor(float *)
  // borderColor()
  // borderColor(unsigned int *)
  // borderColor(int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(true) // "float *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(true) // "unsigned int *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPf
    // invoke: void borderColor(float *)
    var arg0 = (unsafe.Pointer)(args[0].(*float32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QOpenGLTexture11borderColorEPf(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QOpenGLTexture11borderColorEv
    // invoke: QColor borderColor()
    var ret0 = C.C_ZNK14QOpenGLTexture11borderColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
  case 2:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPj
    // invoke: void borderColor(unsigned int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QOpenGLTexture11borderColorEPj(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPi
    // invoke: void borderColor(int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QOpenGLTexture11borderColorEPi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "borderColor", args)
  }

  return
}

// minificationFilter()
func (this *QOpenGLTexture) MinificationFilter(args ...interface{}) () {
  // minificationFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture18minificationFilterEv
    // invoke: QOpenGLTexture::Filter minificationFilter()
    C.C_ZNK14QOpenGLTexture18minificationFilterEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "minificationFilter", args)
  }

  return
}

// isCreated()
func (this *QOpenGLTexture) IsCreated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture9isCreatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isCreated", args)
  }

  return
}

// setSamples(int)
func (this *QOpenGLTexture) SetSamples(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture10setSamplesEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSamples", args)
  }

  return
}

// setAutoMipMapGenerationEnabled(_Bool)
func (this *QOpenGLTexture) SetAutoMipMapGenerationEnabled(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture30setAutoMipMapGenerationEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setAutoMipMapGenerationEnabled", args)
  }

  return
}

// magnificationFilter()
func (this *QOpenGLTexture) MagnificationFilter(args ...interface{}) () {
  // magnificationFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture19magnificationFilterEv
    // invoke: QOpenGLTexture::Filter magnificationFilter()
    C.C_ZNK14QOpenGLTexture19magnificationFilterEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "magnificationFilter", args)
  }

  return
}

// create()
func (this *QOpenGLTexture) Create(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLTexture6createEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "create", args)
  }

  return
}

// minimumLevelOfDetail()
func (this *QOpenGLTexture) MinimumLevelOfDetail(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture20minimumLevelOfDetailEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "minimumLevelOfDetail", args)
  }

  return
}

// setLevelofDetailBias(float)
func (this *QOpenGLTexture) SetLevelofDetailBias(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture20setLevelofDetailBiasEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelofDetailBias", args)
  }

  return
}

// maximumLevelOfDetail()
func (this *QOpenGLTexture) MaximumLevelOfDetail(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture20maximumLevelOfDetailEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumLevelOfDetail", args)
  }

  return
}

// width()
func (this *QOpenGLTexture) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "width", args)
  }

  return
}

// setMipLevelRange(int, int)
func (this *QOpenGLTexture) SetMipLevelRange(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN14QOpenGLTexture16setMipLevelRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevelRange", args)
  }

  return
}

// samples()
func (this *QOpenGLTexture) Samples(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture7samplesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "samples", args)
  }

  return
}

// isStorageAllocated()
func (this *QOpenGLTexture) IsStorageAllocated(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture18isStorageAllocatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isStorageAllocated", args)
  }

  return
}

// setLevelOfDetailRange(float, float)
func (this *QOpenGLTexture) SetLevelOfDetailRange(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    C.C_ZN14QOpenGLTexture21setLevelOfDetailRangeEff(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelOfDetailRange", args)
  }

  return
}

// textureId()
func (this *QOpenGLTexture) TextureId(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture9textureIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "textureId", args)
  }

  return
}

// ~QOpenGLTexture()
func (this *QOpenGLTexture) Free(args ...interface{}) () {
  // ~QOpenGLTexture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTextureD0Ev
    // invoke: void ~QOpenGLTexture()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QOpenGLTextureD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "~QOpenGLTexture", args)
  }

  return
}

// destroy()
func (this *QOpenGLTexture) Destroy(args ...interface{}) () {
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
    C.C_ZN14QOpenGLTexture7destroyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "destroy", args)
  }

  return
}

// setMipBaseLevel(int)
func (this *QOpenGLTexture) SetMipBaseLevel(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture15setMipBaseLevelEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipBaseLevel", args)
  }

  return
}

// layers()
func (this *QOpenGLTexture) Layers(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture6layersEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "layers", args)
  }

  return
}

// levelOfDetailRange()
func (this *QOpenGLTexture) LevelOfDetailRange(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLTexture18levelOfDetailRangeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelOfDetailRange", args)
  }

  return
}

// setBorderColor(class QColor)
func (this *QOpenGLTexture) SetBorderColor(args ...interface{}) () {
  // setBorderColor(class QColor)
  // setBorderColor(int, int, int, int)
  // setBorderColor(float, float, float, float)
  // setBorderColor(uint, uint, uint, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "QColor"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.FloatTy(false) // "float"
  vtys[2][2] = qtrt.FloatTy(false) // "float"
  vtys[2][3] = qtrt.FloatTy(false) // "float"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "uint"
  vtys[3][1] = qtrt.Int32Ty(false) // "uint"
  vtys[3][2] = qtrt.Int32Ty(false) // "uint"
  vtys[3][3] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture14setBorderColorE6QColor
    // invoke: void setBorderColor(class QColor)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture14setBorderColorE6QColor(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEiiii
    // invoke: void setBorderColor(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN14QOpenGLTexture14setBorderColorEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEffff
    // invoke: void setBorderColor(float, float, float, float)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    C.C_ZN14QOpenGLTexture14setBorderColorEffff(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEjjjj
    // invoke: void setBorderColor(uint, uint, uint, uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN14QOpenGLTexture14setBorderColorEjjjj(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setBorderColor", args)
  }

  return
}

// mipBaseLevel()
func (this *QOpenGLTexture) MipBaseLevel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture12mipBaseLevelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipBaseLevel", args)
  }

  return
}

// setFixedSamplePositions(_Bool)
func (this *QOpenGLTexture) SetFixedSamplePositions(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture23setFixedSamplePositionsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setFixedSamplePositions", args)
  }

  return
}

// setMipMaxLevel(int)
func (this *QOpenGLTexture) SetMipMaxLevel(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture14setMipMaxLevelEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipMaxLevel", args)
  }

  return
}

// setCompressedData(int, int, int, void *, const class QOpenGLPixelTransferOptions *const)
func (this *QOpenGLTexture) SetCompressedData(args ...interface{}) () {
  // setCompressedData(int, int, int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, const void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, void *, const class QOpenGLPixelTransferOptions *const)
  // setCompressedData(int, int, int, const void *, const class QOpenGLPixelTransferOptions *const)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.VoidpTy() // "void *"
  vtys[0][4] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.VoidpTy() // "void *"
  vtys[1][2] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.VoidpTy() // "const void *"
  vtys[2][3] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.VoidpTy() // "const void *"
  vtys[3][2] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.VoidpTy() // "void *"
  vtys[4][3] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = qtrt.VoidpTy() // "const void *"
  vtys[5][4] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions *const"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiPKvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 5:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions
    // invoke: void setCompressedData(int, int, int, const void *, const class QOpenGLPixelTransferOptions *const)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QOpenGLPixelTransferOptions).Qclsinst
    if false {fmt.Println(arg4)}
    C.C_ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setCompressedData", args)
  }

  return
}

// isAutoMipMapGenerationEnabled()
func (this *QOpenGLTexture) IsAutoMipMapGenerationEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture29isAutoMipMapGenerationEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isAutoMipMapGenerationEnabled", args)
  }

  return
}

// mipLevels()
func (this *QOpenGLTexture) MipLevels(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture9mipLevelsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevels", args)
  }

  return
}

// levelofDetailBias()
func (this *QOpenGLTexture) LevelofDetailBias(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture17levelofDetailBiasEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelofDetailBias", args)
  }

  return
}

// setMaximumLevelOfDetail(float)
func (this *QOpenGLTexture) SetMaximumLevelOfDetail(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture23setMaximumLevelOfDetailEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumLevelOfDetail", args)
  }

  return
}

// comparisonMode()
func (this *QOpenGLTexture) ComparisonMode(args ...interface{}) () {
  // comparisonMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture14comparisonModeEv
    // invoke: QOpenGLTexture::ComparisonMode comparisonMode()
    C.C_ZNK14QOpenGLTexture14comparisonModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "comparisonMode", args)
  }

  return
}

// setMaximumAnisotropy(float)
func (this *QOpenGLTexture) SetMaximumAnisotropy(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture20setMaximumAnisotropyEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumAnisotropy", args)
  }

  return
}

// setLayers(int)
func (this *QOpenGLTexture) SetLayers(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture9setLayersEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLayers", args)
  }

  return
}

// depthStencilMode()
func (this *QOpenGLTexture) DepthStencilMode(args ...interface{}) () {
  // depthStencilMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture16depthStencilModeEv
    // invoke: QOpenGLTexture::DepthStencilMode depthStencilMode()
    C.C_ZNK14QOpenGLTexture16depthStencilModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "depthStencilMode", args)
  }

  return
}

// target()
func (this *QOpenGLTexture) Target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture6targetEv
    // invoke: QOpenGLTexture::Target target()
    C.C_ZNK14QOpenGLTexture6targetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "target", args)
  }

  return
}

// mipLevelRange()
func (this *QOpenGLTexture) MipLevelRange(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLTexture13mipLevelRangeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevelRange", args)
  }

  return
}

// isBound(uint)
func (this *QOpenGLTexture) IsBound(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QOpenGLTexture7isBoundEj(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK14QOpenGLTexture7isBoundEv
    // invoke: bool isBound()
    var ret0 = C.C_ZNK14QOpenGLTexture7isBoundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isBound", args)
  }

  return
}

// bind()
func (this *QOpenGLTexture) Bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture4bindEv
    // invoke: void bind()
    C.C_ZN14QOpenGLTexture4bindEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "bind", args)
  }

  return
}

// minMagFilters()
func (this *QOpenGLTexture) MinMagFilters(args ...interface{}) () {
  // minMagFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLTexture13minMagFiltersEv
    // invoke: QPair<QOpenGLTexture::Filter, QOpenGLTexture::Filter> minMagFilters()
    C.C_ZNK14QOpenGLTexture13minMagFiltersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "minMagFilters", args)
  }

  return
}

// mipMaxLevel()
func (this *QOpenGLTexture) MipMaxLevel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture11mipMaxLevelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipMaxLevel", args)
  }

  return
}

// depth()
func (this *QOpenGLTexture) Depth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "depth", args)
  }

  return
}

// setMipLevels(int)
func (this *QOpenGLTexture) SetMipLevels(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QOpenGLTexture12setMipLevelsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevels", args)
  }

  return
}

// faces()
func (this *QOpenGLTexture) Faces(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLTexture5facesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "faces", args)
  }

  return
}

// release()
func (this *QOpenGLTexture) Release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLTexture7releaseEv
    // invoke: void release()
    C.C_ZN14QOpenGLTexture7releaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "release", args)
  }

  return
}

// generateMipMaps(int, _Bool)
func (this *QOpenGLTexture) GenerateMipMaps(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN14QOpenGLTexture15generateMipMapsEib(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN14QOpenGLTexture15generateMipMapsEv
    // invoke: void generateMipMaps()
    C.C_ZN14QOpenGLTexture15generateMipMapsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "generateMipMaps", args)
  }

  return
}

// setSize(int, int, int)
func (this *QOpenGLTexture) SetSize(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN14QOpenGLTexture7setSizeEiii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSize", args)
  }

  return
}

// <= body block end

