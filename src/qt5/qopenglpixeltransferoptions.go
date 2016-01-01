package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglpixeltransferoptions.h
// dst-file: /src/gui/qopenglpixeltransferoptions.go
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
// class sizeof(QOpenGLPixelTransferOptions)=1
type QOpenGLPixelTransferOptions struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQOpenGLPixelTransferOptions(args ...interface{})() {
}


func (this *QOpenGLPixelTransferOptions) FreeQOpenGLPixelTransferOptions(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "~QOpenGLPixelTransferOptions", args)
 }

}


func (this *QOpenGLPixelTransferOptions) isSwapBytesEnabled(args ...interface{}) () {
  // isSwapBytesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isSwapBytesEnabled", args)
 }

}


func (this *QOpenGLPixelTransferOptions) swap(args ...interface{}) () {
  // swap(class QOpenGLPixelTransferOptions &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "QOpenGLPixelTransferOptions &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions4swapERS_
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "swap", args)
 }

}


func (this *QOpenGLPixelTransferOptions) skipImages(args ...interface{}) () {
  // skipImages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipImagesEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipImages", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setSkipRows(args ...interface{}) () {
  // setSkipRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions11setSkipRowsEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipRows", args)
 }

}


func (this *QOpenGLPixelTransferOptions) skipPixels(args ...interface{}) () {
  // skipPixels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipPixelsEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipPixels", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setRowLength(args ...interface{}) () {
  // setRowLength(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions12setRowLengthEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setRowLength", args)
 }

}


func (this *QOpenGLPixelTransferOptions) imageHeight(args ...interface{}) () {
  // imageHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions11imageHeightEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "imageHeight", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setImageHeight(args ...interface{}) () {
  // setImageHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions14setImageHeightEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setImageHeight", args)
 }

}


func (this *QOpenGLPixelTransferOptions) skipRows(args ...interface{}) () {
  // skipRows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions8skipRowsEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipRows", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setAlignment(args ...interface{}) () {
  // setAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions12setAlignmentEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setAlignment", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setSkipImages(args ...interface{}) () {
  // setSkipImages(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions13setSkipImagesEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipImages", args)
 }

}


func (this *QOpenGLPixelTransferOptions) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9alignmentEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "alignment", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setSkipPixels(args ...interface{}) () {
  // setSkipPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipPixels", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setSwapBytesEnabled(args ...interface{}) () {
  // setSwapBytesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSwapBytesEnabled", args)
 }

}


func (this *QOpenGLPixelTransferOptions) setLeastSignificantByteFirst(args ...interface{}) () {
  // setLeastSignificantByteFirst(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setLeastSignificantByteFirst", args)
 }

}


func (this *QOpenGLPixelTransferOptions) isLeastSignificantBitFirst(args ...interface{}) () {
  // isLeastSignificantBitFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isLeastSignificantBitFirst", args)
 }

}


func (this *QOpenGLPixelTransferOptions) rowLength(args ...interface{}) () {
  // rowLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9rowLengthEv
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "rowLength", args)
 }

}

// <= body block end

