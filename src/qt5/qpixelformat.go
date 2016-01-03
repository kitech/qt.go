package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qpixelformat.h
// dst-file: /src/gui/qpixelformat.go
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
  // proto:  uchar QPixelFormat::blackSize();
extern void demth_ZNK12QPixelFormat9blackSizeEv(void* qthis);
  // proto:  void QPixelFormat::QPixelFormat();
extern void* dector_ZN12QPixelFormatC1Ev();
extern void _ZN12QPixelFormatC1Ev(void* qthis);
  // proto:  uchar QPixelFormat::subEnum();
extern void demth_ZNK12QPixelFormat7subEnumEv(void* qthis);
  // proto:  uchar QPixelFormat::greenSize();
extern void demth_ZNK12QPixelFormat9greenSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::lightnessSize();
extern void demth_ZNK12QPixelFormat13lightnessSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::bitsPerPixel();
extern void demth_ZNK12QPixelFormat12bitsPerPixelEv(void* qthis);
  // proto:  uchar QPixelFormat::alphaSize();
extern void demth_ZNK12QPixelFormat9alphaSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::magentaSize();
extern void demth_ZNK12QPixelFormat11magentaSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::hueSize();
extern void demth_ZNK12QPixelFormat7hueSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::saturationSize();
extern void demth_ZNK12QPixelFormat14saturationSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::brightnessSize();
extern void demth_ZNK12QPixelFormat14brightnessSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::yellowSize();
extern void demth_ZNK12QPixelFormat10yellowSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::redSize();
extern void demth_ZNK12QPixelFormat7redSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::blueSize();
extern void demth_ZNK12QPixelFormat8blueSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::cyanSize();
extern void demth_ZNK12QPixelFormat8cyanSizeEv(void* qthis);
  // proto:  uchar QPixelFormat::channelCount();
extern void demth_ZNK12QPixelFormat12channelCountEv(void* qthis);
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

// class sizeof(QPixelFormat)=8
type QPixelFormat struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  uchar QPixelFormat::blackSize();
func (this *QPixelFormat) blackSize(args ...interface{}) () {
  // blackSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9blackSizeEv
    // invoke: uchar blackSize()
    C.demth_ZNK12QPixelFormat9blackSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "blackSize", args)
  }

}

  // proto:  void QPixelFormat::QPixelFormat();
func NewQPixelFormat(args ...interface{}) QPixelFormat {
  return QPixelFormat{}
}

  // proto:  uchar QPixelFormat::subEnum();
func (this *QPixelFormat) subEnum(args ...interface{}) () {
  // subEnum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7subEnumEv
    // invoke: uchar subEnum()
    C.demth_ZNK12QPixelFormat7subEnumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "subEnum", args)
  }

}

  // proto:  uchar QPixelFormat::greenSize();
func (this *QPixelFormat) greenSize(args ...interface{}) () {
  // greenSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9greenSizeEv
    // invoke: uchar greenSize()
    C.demth_ZNK12QPixelFormat9greenSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "greenSize", args)
  }

}

  // proto:  uchar QPixelFormat::lightnessSize();
func (this *QPixelFormat) lightnessSize(args ...interface{}) () {
  // lightnessSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13lightnessSizeEv
    // invoke: uchar lightnessSize()
    C.demth_ZNK12QPixelFormat13lightnessSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "lightnessSize", args)
  }

}

  // proto:  uchar QPixelFormat::bitsPerPixel();
func (this *QPixelFormat) bitsPerPixel(args ...interface{}) () {
  // bitsPerPixel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat12bitsPerPixelEv
    // invoke: uchar bitsPerPixel()
    C.demth_ZNK12QPixelFormat12bitsPerPixelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "bitsPerPixel", args)
  }

}

  // proto:  uchar QPixelFormat::alphaSize();
func (this *QPixelFormat) alphaSize(args ...interface{}) () {
  // alphaSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9alphaSizeEv
    // invoke: uchar alphaSize()
    C.demth_ZNK12QPixelFormat9alphaSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaSize", args)
  }

}

  // proto:  uchar QPixelFormat::magentaSize();
func (this *QPixelFormat) magentaSize(args ...interface{}) () {
  // magentaSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat11magentaSizeEv
    // invoke: uchar magentaSize()
    C.demth_ZNK12QPixelFormat11magentaSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "magentaSize", args)
  }

}

  // proto:  uchar QPixelFormat::hueSize();
func (this *QPixelFormat) hueSize(args ...interface{}) () {
  // hueSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7hueSizeEv
    // invoke: uchar hueSize()
    C.demth_ZNK12QPixelFormat7hueSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "hueSize", args)
  }

}

  // proto:  uchar QPixelFormat::saturationSize();
func (this *QPixelFormat) saturationSize(args ...interface{}) () {
  // saturationSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat14saturationSizeEv
    // invoke: uchar saturationSize()
    C.demth_ZNK12QPixelFormat14saturationSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "saturationSize", args)
  }

}

  // proto:  uchar QPixelFormat::brightnessSize();
func (this *QPixelFormat) brightnessSize(args ...interface{}) () {
  // brightnessSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat14brightnessSizeEv
    // invoke: uchar brightnessSize()
    C.demth_ZNK12QPixelFormat14brightnessSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "brightnessSize", args)
  }

}

  // proto:  uchar QPixelFormat::yellowSize();
func (this *QPixelFormat) yellowSize(args ...interface{}) () {
  // yellowSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10yellowSizeEv
    // invoke: uchar yellowSize()
    C.demth_ZNK12QPixelFormat10yellowSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "yellowSize", args)
  }

}

  // proto:  uchar QPixelFormat::redSize();
func (this *QPixelFormat) redSize(args ...interface{}) () {
  // redSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat7redSizeEv
    // invoke: uchar redSize()
    C.demth_ZNK12QPixelFormat7redSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "redSize", args)
  }

}

  // proto:  uchar QPixelFormat::blueSize();
func (this *QPixelFormat) blueSize(args ...interface{}) () {
  // blueSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat8blueSizeEv
    // invoke: uchar blueSize()
    C.demth_ZNK12QPixelFormat8blueSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "blueSize", args)
  }

}

  // proto:  uchar QPixelFormat::cyanSize();
func (this *QPixelFormat) cyanSize(args ...interface{}) () {
  // cyanSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat8cyanSizeEv
    // invoke: uchar cyanSize()
    C.demth_ZNK12QPixelFormat8cyanSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "cyanSize", args)
  }

}

  // proto:  uchar QPixelFormat::channelCount();
func (this *QPixelFormat) channelCount(args ...interface{}) () {
  // channelCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat12channelCountEv
    // invoke: uchar channelCount()
    C.demth_ZNK12QPixelFormat12channelCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "channelCount", args)
  }

}

// <= body block end

