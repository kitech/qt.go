package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qopenglpixeltransferoptions.h
// dst-file: /src/gui/qopenglpixeltransferoptions.go
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
  // proto:  void QOpenGLPixelTransferOptions::setLeastSignificantByteFirst(bool lsbFirst);
extern void C_ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLPixelTransferOptions::setSkipImages(int skipImages);
extern void C_ZN27QOpenGLPixelTransferOptions13setSkipImagesEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLPixelTransferOptions::setAlignment(int alignment);
extern void C_ZN27QOpenGLPixelTransferOptions12setAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  int QOpenGLPixelTransferOptions::skipImages();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions10skipImagesEv(void* qthis); // 4
  // proto:  void QOpenGLPixelTransferOptions::~QOpenGLPixelTransferOptions();
extern void C_ZN27QOpenGLPixelTransferOptionsD2Ev(void* qthis); // 4
  // proto:  int QOpenGLPixelTransferOptions::rowLength();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions9rowLengthEv(void* qthis); // 4
  // proto:  bool QOpenGLPixelTransferOptions::isLeastSignificantBitFirst();
extern bool C_ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv(void* qthis); // 4
  // proto:  void QOpenGLPixelTransferOptions::setSkipRows(int skipRows);
extern void C_ZN27QOpenGLPixelTransferOptions11setSkipRowsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QOpenGLPixelTransferOptions::skipRows();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions8skipRowsEv(void* qthis); // 4
  // proto:  void QOpenGLPixelTransferOptions::setImageHeight(int imageHeight);
extern void C_ZN27QOpenGLPixelTransferOptions14setImageHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLPixelTransferOptions::QOpenGLPixelTransferOptions();
extern void* C_ZN27QOpenGLPixelTransferOptionsC2Ev(); // 3
  // proto:  void QOpenGLPixelTransferOptions::QOpenGLPixelTransferOptions(const QOpenGLPixelTransferOptions & );
extern void* C_ZN27QOpenGLPixelTransferOptionsC2ERKS_(void* arg0); // 3
  // proto:  void QOpenGLPixelTransferOptions::setSkipPixels(int skipPixels);
extern void C_ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLPixelTransferOptions::setSwapBytesEnabled(bool swapBytes);
extern void C_ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLPixelTransferOptions::swap(QOpenGLPixelTransferOptions & other);
extern void C_ZN27QOpenGLPixelTransferOptions4swapERS_(void* qthis, void* arg0); // 2
  // proto:  int QOpenGLPixelTransferOptions::skipPixels();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions10skipPixelsEv(void* qthis); // 4
  // proto:  void QOpenGLPixelTransferOptions::setRowLength(int rowLength);
extern void C_ZN27QOpenGLPixelTransferOptions12setRowLengthEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QOpenGLPixelTransferOptions::isSwapBytesEnabled();
extern bool C_ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv(void* qthis); // 4
  // proto:  int QOpenGLPixelTransferOptions::alignment();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions9alignmentEv(void* qthis); // 4
  // proto:  int QOpenGLPixelTransferOptions::imageHeight();
extern int32_t C_ZNK27QOpenGLPixelTransferOptions11imageHeightEv(void* qthis); // 4
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

// class sizeof(QOpenGLPixelTransferOptions)=1
type QOpenGLPixelTransferOptions struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setLeastSignificantByteFirst(_Bool)
func (this *QOpenGLPixelTransferOptions) Setleastsignificantbytefirst(args ...interface{}) () {
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
    // invoke: void setLeastSignificantByteFirst(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setLeastSignificantByteFirst", args)
  }

  return
}

// setSkipImages(int)
func (this *QOpenGLPixelTransferOptions) Setskipimages(args ...interface{}) () {
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
    // invoke: void setSkipImages(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions13setSkipImagesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipImages", args)
  }

  return
}

// setAlignment(int)
func (this *QOpenGLPixelTransferOptions) Setalignment(args ...interface{}) () {
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
    // invoke: void setAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions12setAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setAlignment", args)
  }

  return
}

// skipImages()
func (this *QOpenGLPixelTransferOptions) Skipimages(args ...interface{}) (ret interface{}) {
  // skipImages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipImagesEv
    // invoke: int skipImages()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions10skipImagesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipImages", args)
  }

  return
}

// ~QOpenGLPixelTransferOptions()
func (this *QOpenGLPixelTransferOptions) Freeqopenglpixeltransferoptions(args ...interface{}) () {
  // ~QOpenGLPixelTransferOptions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptionsD0Ev
    // invoke: void ~QOpenGLPixelTransferOptions()
    C.C_ZN27QOpenGLPixelTransferOptionsD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "~QOpenGLPixelTransferOptions", args)
  }

  return
}

// rowLength()
func (this *QOpenGLPixelTransferOptions) Rowlength(args ...interface{}) (ret interface{}) {
  // rowLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9rowLengthEv
    // invoke: int rowLength()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions9rowLengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "rowLength", args)
  }

  return
}

// isLeastSignificantBitFirst()
func (this *QOpenGLPixelTransferOptions) Isleastsignificantbitfirst(args ...interface{}) (ret interface{}) {
  // isLeastSignificantBitFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv
    // invoke: bool isLeastSignificantBitFirst()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isLeastSignificantBitFirst", args)
  }

  return
}

// setSkipRows(int)
func (this *QOpenGLPixelTransferOptions) Setskiprows(args ...interface{}) () {
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
    // invoke: void setSkipRows(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions11setSkipRowsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipRows", args)
  }

  return
}

// skipRows()
func (this *QOpenGLPixelTransferOptions) Skiprows(args ...interface{}) (ret interface{}) {
  // skipRows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions8skipRowsEv
    // invoke: int skipRows()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions8skipRowsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipRows", args)
  }

  return
}

// setImageHeight(int)
func (this *QOpenGLPixelTransferOptions) Setimageheight(args ...interface{}) () {
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
    // invoke: void setImageHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions14setImageHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setImageHeight", args)
  }

  return
}

// QOpenGLPixelTransferOptions()
func NewQOpenGLPixelTransferOptions(args ...interface{}) *QOpenGLPixelTransferOptions {
  // QOpenGLPixelTransferOptions()
  // QOpenGLPixelTransferOptions(const class QOpenGLPixelTransferOptions &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "const QOpenGLPixelTransferOptions &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptionsC1Ev
    // invoke: void QOpenGLPixelTransferOptions()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QOpenGLPixelTransferOptionsC2Ev()
    return &QOpenGLPixelTransferOptions{qclsinst:qthis}
  case 1:
    // invoke: _ZN27QOpenGLPixelTransferOptionsC1ERKS_
    // invoke: void QOpenGLPixelTransferOptions(const class QOpenGLPixelTransferOptions &)
    var arg0 = args[0].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QOpenGLPixelTransferOptionsC2ERKS_(arg0)
    return &QOpenGLPixelTransferOptions{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "QOpenGLPixelTransferOptions", args)
  }

  return nil // QOpenGLPixelTransferOptions{qclsinst:qthis}
}

// setSkipPixels(int)
func (this *QOpenGLPixelTransferOptions) Setskippixels(args ...interface{}) () {
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
    // invoke: void setSkipPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipPixels", args)
  }

  return
}

// setSwapBytesEnabled(_Bool)
func (this *QOpenGLPixelTransferOptions) Setswapbytesenabled(args ...interface{}) () {
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
    // invoke: void setSwapBytesEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSwapBytesEnabled", args)
  }

  return
}

// swap(class QOpenGLPixelTransferOptions &)
func (this *QOpenGLPixelTransferOptions) Swap(args ...interface{}) () {
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
    // invoke: void swap(class QOpenGLPixelTransferOptions &)
    var arg0 = args[0].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "swap", args)
  }

  return
}

// skipPixels()
func (this *QOpenGLPixelTransferOptions) Skippixels(args ...interface{}) (ret interface{}) {
  // skipPixels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipPixelsEv
    // invoke: int skipPixels()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions10skipPixelsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipPixels", args)
  }

  return
}

// setRowLength(int)
func (this *QOpenGLPixelTransferOptions) Setrowlength(args ...interface{}) () {
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
    // invoke: void setRowLength(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN27QOpenGLPixelTransferOptions12setRowLengthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setRowLength", args)
  }

  return
}

// isSwapBytesEnabled()
func (this *QOpenGLPixelTransferOptions) Isswapbytesenabled(args ...interface{}) (ret interface{}) {
  // isSwapBytesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv
    // invoke: bool isSwapBytesEnabled()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isSwapBytesEnabled", args)
  }

  return
}

// alignment()
func (this *QOpenGLPixelTransferOptions) Alignment(args ...interface{}) (ret interface{}) {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9alignmentEv
    // invoke: int alignment()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions9alignmentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "alignment", args)
  }

  return
}

// imageHeight()
func (this *QOpenGLPixelTransferOptions) Imageheight(args ...interface{}) (ret interface{}) {
  // imageHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions11imageHeightEv
    // invoke: int imageHeight()
    var ret0 = C.C_ZNK27QOpenGLPixelTransferOptions11imageHeightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "imageHeight", args)
  }

  return
}

// <= body block end

