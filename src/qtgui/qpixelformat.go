package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  QPixelFormat::YUVLayout QPixelFormat::yuvLayout();
extern void C_ZNK12QPixelFormat9yuvLayoutEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::yellowSize();
extern unsigned char C_ZNK12QPixelFormat10yellowSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaPosition QPixelFormat::alphaPosition();
extern void C_ZNK12QPixelFormat13alphaPositionEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::cyanSize();
extern unsigned char C_ZNK12QPixelFormat8cyanSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::lightnessSize();
extern unsigned char C_ZNK12QPixelFormat13lightnessSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::brightnessSize();
extern unsigned char C_ZNK12QPixelFormat14brightnessSizeEv(void* qthis); // 2
  // proto:  void QPixelFormat::QPixelFormat();
extern void* C_ZN12QPixelFormatC2Ev(); // 1
  // proto:  uchar QPixelFormat::bitsPerPixel();
extern unsigned char C_ZNK12QPixelFormat12bitsPerPixelEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::blackSize();
extern unsigned char C_ZNK12QPixelFormat9blackSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::redSize();
extern unsigned char C_ZNK12QPixelFormat7redSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::ByteOrder QPixelFormat::byteOrder();
extern void C_ZNK12QPixelFormat9byteOrderEv(void* qthis); // 2
  // proto:  QPixelFormat::TypeInterpretation QPixelFormat::typeInterpretation();
extern void C_ZNK12QPixelFormat18typeInterpretationEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::channelCount();
extern unsigned char C_ZNK12QPixelFormat12channelCountEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::alphaSize();
extern unsigned char C_ZNK12QPixelFormat9alphaSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::hueSize();
extern unsigned char C_ZNK12QPixelFormat7hueSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::saturationSize();
extern unsigned char C_ZNK12QPixelFormat14saturationSizeEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::magentaSize();
extern unsigned char C_ZNK12QPixelFormat11magentaSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaPremultiplied QPixelFormat::premultiplied();
extern void C_ZNK12QPixelFormat13premultipliedEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::subEnum();
extern unsigned char C_ZNK12QPixelFormat7subEnumEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::greenSize();
extern unsigned char C_ZNK12QPixelFormat9greenSizeEv(void* qthis); // 2
  // proto:  QPixelFormat::ColorModel QPixelFormat::colorModel();
extern void C_ZNK12QPixelFormat10colorModelEv(void* qthis); // 2
  // proto:  QPixelFormat::AlphaUsage QPixelFormat::alphaUsage();
extern void C_ZNK12QPixelFormat10alphaUsageEv(void* qthis); // 2
  // proto:  uchar QPixelFormat::blueSize();
extern unsigned char C_ZNK12QPixelFormat8blueSizeEv(void* qthis); // 2
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

// class sizeof(QPixelFormat)=8
type QPixelFormat struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// yuvLayout()
func (this *QPixelFormat) YuvLayout(args ...interface{}) () {
  // yuvLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9yuvLayoutEv
    // invoke: QPixelFormat::YUVLayout yuvLayout()
    C.C_ZNK12QPixelFormat9yuvLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "yuvLayout", args)
  }

  return
}

// yellowSize()
func (this *QPixelFormat) YellowSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat10yellowSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "yellowSize", args)
  }

  return
}

// alphaPosition()
func (this *QPixelFormat) AlphaPosition(args ...interface{}) () {
  // alphaPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13alphaPositionEv
    // invoke: QPixelFormat::AlphaPosition alphaPosition()
    C.C_ZNK12QPixelFormat13alphaPositionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaPosition", args)
  }

  return
}

// cyanSize()
func (this *QPixelFormat) CyanSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat8cyanSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "cyanSize", args)
  }

  return
}

// lightnessSize()
func (this *QPixelFormat) LightnessSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat13lightnessSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "lightnessSize", args)
  }

  return
}

// brightnessSize()
func (this *QPixelFormat) BrightnessSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat14brightnessSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "brightnessSize", args)
  }

  return
}

// QPixelFormat()
func GcfreeQPixelFormat(this *QPixelFormat) {
  qtrt.UniverseFree(this)
}
func NewQPixelFormat(args ...interface{}) *QPixelFormat {
  // QPixelFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixelFormatC1Ev
    // invoke: void QPixelFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QPixelFormatC2Ev()
    this := &QPixelFormat{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQPixelFormat)
    return this
  default:
    qtrt.ErrorResolve("QPixelFormat", "QPixelFormat", args)
  }

  return nil // QPixelFormat{Qclsinst:qthis}
}

// bitsPerPixel()
func (this *QPixelFormat) BitsPerPixel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat12bitsPerPixelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "bitsPerPixel", args)
  }

  return
}

// blackSize()
func (this *QPixelFormat) BlackSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat9blackSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "blackSize", args)
  }

  return
}

// redSize()
func (this *QPixelFormat) RedSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat7redSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "redSize", args)
  }

  return
}

// byteOrder()
func (this *QPixelFormat) ByteOrder(args ...interface{}) () {
  // byteOrder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat9byteOrderEv
    // invoke: QPixelFormat::ByteOrder byteOrder()
    C.C_ZNK12QPixelFormat9byteOrderEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "byteOrder", args)
  }

  return
}

// typeInterpretation()
func (this *QPixelFormat) TypeInterpretation(args ...interface{}) () {
  // typeInterpretation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat18typeInterpretationEv
    // invoke: QPixelFormat::TypeInterpretation typeInterpretation()
    C.C_ZNK12QPixelFormat18typeInterpretationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "typeInterpretation", args)
  }

  return
}

// channelCount()
func (this *QPixelFormat) ChannelCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat12channelCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "channelCount", args)
  }

  return
}

// alphaSize()
func (this *QPixelFormat) AlphaSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat9alphaSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaSize", args)
  }

  return
}

// hueSize()
func (this *QPixelFormat) HueSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat7hueSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "hueSize", args)
  }

  return
}

// saturationSize()
func (this *QPixelFormat) SaturationSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat14saturationSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "saturationSize", args)
  }

  return
}

// magentaSize()
func (this *QPixelFormat) MagentaSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat11magentaSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "magentaSize", args)
  }

  return
}

// premultiplied()
func (this *QPixelFormat) Premultiplied(args ...interface{}) () {
  // premultiplied()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat13premultipliedEv
    // invoke: QPixelFormat::AlphaPremultiplied premultiplied()
    C.C_ZNK12QPixelFormat13premultipliedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "premultiplied", args)
  }

  return
}

// subEnum()
func (this *QPixelFormat) SubEnum(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat7subEnumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "subEnum", args)
  }

  return
}

// greenSize()
func (this *QPixelFormat) GreenSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat9greenSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "greenSize", args)
  }

  return
}

// colorModel()
func (this *QPixelFormat) ColorModel(args ...interface{}) () {
  // colorModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10colorModelEv
    // invoke: QPixelFormat::ColorModel colorModel()
    C.C_ZNK12QPixelFormat10colorModelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "colorModel", args)
  }

  return
}

// alphaUsage()
func (this *QPixelFormat) AlphaUsage(args ...interface{}) () {
  // alphaUsage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPixelFormat10alphaUsageEv
    // invoke: QPixelFormat::AlphaUsage alphaUsage()
    C.C_ZNK12QPixelFormat10alphaUsageEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixelFormat", "alphaUsage", args)
  }

  return
}

// blueSize()
func (this *QPixelFormat) BlueSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QPixelFormat8blueSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixelFormat", "blueSize", args)
  }

  return
}

// <= body block end

