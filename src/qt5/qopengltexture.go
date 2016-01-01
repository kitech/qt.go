package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopengltexture.h
// dst-file: /src/gui/qopengltexture.go
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
// class sizeof(QOpenGLTexture)=1
type QOpenGLTexture struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  case 1:
    // invoke: _ZN14QOpenGLTexture4bindEjNS_16TextureUnitResetE
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "bind", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setFixedSamplePositions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "height", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isAutoMipMapGenerationEnabled", args)
 }

}


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
  case 1:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiNS_11CubeMapFaceEiPvPK27QOpenGLPixelTransferOptions
  case 2:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPKvPK27QOpenGLPixelTransferOptions
  case 3:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPKvPK27QOpenGLPixelTransferOptions
  case 4:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiPvPK27QOpenGLPixelTransferOptions
  case 5:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiNS_11CubeMapFaceEiPKvPK27QOpenGLPixelTransferOptions
  case 6:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiiPvPK27QOpenGLPixelTransferOptions
  case 7:
    // invoke: _ZN14QOpenGLTexture17setCompressedDataEiiPvPK27QOpenGLPixelTransferOptions
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setCompressedData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumLevelOfDetail", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setAutoMipMapGenerationEnabled", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "depth", args)
 }

}


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
  case 1:
    // invoke: _ZN14QOpenGLTexture15generateMipMapsEv
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "generateMipMaps", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipBaseLevel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelOfDetailRange", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "create", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelOfDetailRange", args)
 }

}


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
  case 1:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPi
  case 2:
    // invoke: _ZNK14QOpenGLTexture11borderColorEPf
  case 3:
    // invoke: _ZNK14QOpenGLTexture11borderColorEv
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "borderColor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isStorageAllocated", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isTextureView", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isFixedSamplePositions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "faces", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLayers", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "width", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "layers", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "minimumLevelOfDetail", args)
 }

}


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
  case 1:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEjjjj
  case 2:
    // invoke: _ZN14QOpenGLTexture14setBorderColorE6QColor
  case 3:
    // invoke: _ZN14QOpenGLTexture14setBorderColorEffff
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setBorderColor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMinimumLevelOfDetail", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevels", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevelRange", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipMaxLevel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "levelofDetailBias", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumMipLevels", args)
 }

}


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
  case 1:
    // invoke: _ZNK14QOpenGLTexture7isBoundEv
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isBound", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMaximumAnisotropy", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSamples", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipLevels", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setLevelofDetailBias", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "textureId", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setMipLevelRange", args)
 }

}


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
    // invoke: _ZN14QOpenGLTexture15allocateStorageENS_11PixelFormatENS_9PixelTypeE
  case 1:
    // invoke: _ZN14QOpenGLTexture15allocateStorageEv
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "allocateStorage", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipMaxLevel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "destroy", args)
 }

}


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
    // invoke: _ZN14QOpenGLTexture7releaseEjNS_16TextureUnitResetE
  case 1:
    // invoke: _ZN14QOpenGLTexture7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "release", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumAnisotropy", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "maximumLevelOfDetail", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "setSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "isCreated", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "samples", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLTexture", "mipBaseLevel", args)
 }

}


func NewQOpenGLTexture(args ...interface{})() {
}

// <= body block end

