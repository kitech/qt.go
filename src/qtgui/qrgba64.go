package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtGui/qrgba64.h
// dst-file: /src/gui/qrgba64.go
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
  // proto:  void QRgba64::setBlue(quint16 _blue);
extern void C_ZN7QRgba647setBlueEt(void* qthis, int16_t arg0); // 2
  // proto:  bool QRgba64::isOpaque();
extern bool C_ZNK7QRgba648isOpaqueEv(void* qthis); // 2
  // proto:  bool QRgba64::isTransparent();
extern bool C_ZNK7QRgba6413isTransparentEv(void* qthis); // 2
  // proto:  void QRgba64::QRgba64();
extern void* C_ZN7QRgba64C2Ev(); // 1
  // proto:  void QRgba64::setGreen(quint16 _green);
extern void C_ZN7QRgba648setGreenEt(void* qthis, int16_t arg0); // 2
  // proto:  void QRgba64::setRed(quint16 _red);
extern void C_ZN7QRgba646setRedEt(void* qthis, int16_t arg0); // 2
  // proto: static QRgba64 QRgba64::fromRgba64(quint64 c);
extern void* C_ZN7QRgba6410fromRgba64Ey(int64_t arg0); // 2
  // proto: static QRgba64 QRgba64::fromRgba64(quint16 red, quint16 green, quint16 blue, quint16 alpha);
extern void* C_ZN7QRgba6410fromRgba64Etttt(int16_t arg0, int16_t arg1, int16_t arg2, int16_t arg3); // 2
  // proto:  quint16 QRgba64::blue();
extern int16_t C_ZNK7QRgba644blueEv(void* qthis); // 2
  // proto:  quint8 QRgba64::alpha8();
extern unsigned char C_ZNK7QRgba646alpha8Ev(void* qthis); // 2
  // proto:  quint8 QRgba64::green8();
extern unsigned char C_ZNK7QRgba646green8Ev(void* qthis); // 2
  // proto:  void QRgba64::setAlpha(quint16 _alpha);
extern void C_ZN7QRgba648setAlphaEt(void* qthis, int16_t arg0); // 2
  // proto:  QRgba64 QRgba64::unpremultiplied();
extern void* C_ZNK7QRgba6415unpremultipliedEv(void* qthis); // 2
  // proto:  quint16 QRgba64::red();
extern int16_t C_ZNK7QRgba643redEv(void* qthis); // 2
  // proto: static QRgba64 QRgba64::fromArgb32(uint rgb);
extern void* C_ZN7QRgba6410fromArgb32Ej(int32_t arg0); // 2
  // proto:  quint8 QRgba64::red8();
extern unsigned char C_ZNK7QRgba644red8Ev(void* qthis); // 2
  // proto:  quint8 QRgba64::blue8();
extern unsigned char C_ZNK7QRgba645blue8Ev(void* qthis); // 2
  // proto:  ushort QRgba64::toRgb16();
extern int16_t C_ZNK7QRgba647toRgb16Ev(void* qthis); // 2
  // proto:  QRgba64 QRgba64::premultiplied();
extern void* C_ZNK7QRgba6413premultipliedEv(void* qthis); // 2
  // proto:  quint16 QRgba64::alpha();
extern int16_t C_ZNK7QRgba645alphaEv(void* qthis); // 2
  // proto:  uint QRgba64::toArgb32();
extern int32_t C_ZNK7QRgba648toArgb32Ev(void* qthis); // 2
  // proto:  quint16 QRgba64::green();
extern int16_t C_ZNK7QRgba645greenEv(void* qthis); // 2
  // proto: static QRgba64 QRgba64::fromRgba(quint8 red, quint8 green, quint8 blue, quint8 alpha);
extern void* C_ZN7QRgba648fromRgbaEhhhh(unsigned char arg0, unsigned char arg1, unsigned char arg2, unsigned char arg3); // 2
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

// class sizeof(QRgba64)=8
type QRgba64 struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setBlue(quint16)
func (this *QRgba64) Setblue(args ...interface{}) () {
  // setBlue(quint16)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "quint16"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba647setBlueEt
    // invoke: void setBlue(quint16)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    C.C_ZN7QRgba647setBlueEt(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRgba64", "setBlue", args)
  }

  return
}

// isOpaque()
func (this *QRgba64) Isopaque(args ...interface{}) (ret interface{}) {
  // isOpaque()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba648isOpaqueEv
    // invoke: bool isOpaque()
    var ret0 = C.C_ZNK7QRgba648isOpaqueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "isOpaque", args)
  }

  return
}

// isTransparent()
func (this *QRgba64) Istransparent(args ...interface{}) (ret interface{}) {
  // isTransparent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba6413isTransparentEv
    // invoke: bool isTransparent()
    var ret0 = C.C_ZNK7QRgba6413isTransparentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "isTransparent", args)
  }

  return
}

// QRgba64()
func NewQRgba64(args ...interface{}) *QRgba64 {
  // QRgba64()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba64C1Ev
    // invoke: void QRgba64()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QRgba64C2Ev()
    return &QRgba64{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QRgba64", "QRgba64", args)
  }

  return nil // QRgba64{Qclsinst:qthis}
}

// setGreen(quint16)
func (this *QRgba64) Setgreen(args ...interface{}) () {
  // setGreen(quint16)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "quint16"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba648setGreenEt
    // invoke: void setGreen(quint16)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    C.C_ZN7QRgba648setGreenEt(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRgba64", "setGreen", args)
  }

  return
}

// setRed(quint16)
func (this *QRgba64) Setred(args ...interface{}) () {
  // setRed(quint16)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "quint16"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba646setRedEt
    // invoke: void setRed(quint16)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    C.C_ZN7QRgba646setRedEt(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRgba64", "setRed", args)
  }

  return
}

// fromRgba64(quint64)
func (this *QRgba64) Fromrgba64_S(args ...interface{}) (ret interface{}) {
  // fromRgba64(quint64)
  // fromRgba64(quint16, quint16, quint16, quint16)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "quint64"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int16Ty(false) // "quint16"
  vtys[1][1] = qtrt.Int16Ty(false) // "quint16"
  vtys[1][2] = qtrt.Int16Ty(false) // "quint16"
  vtys[1][3] = qtrt.Int16Ty(false) // "quint16"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba6410fromRgba64Ey
    // invoke: QRgba64 fromRgba64(quint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QRgba6410fromRgba64Ey(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QRgba6410fromRgba64Etttt
    // invoke: QRgba64 fromRgba64(quint16, quint16, quint16, quint16)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int16_t(qtrt.PrimConv(args[1], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg1)}
    var arg2 = C.int16_t(qtrt.PrimConv(args[2], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg2)}
    var arg3 = C.int16_t(qtrt.PrimConv(args[3], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN7QRgba6410fromRgba64Etttt(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "fromRgba64", args)
  }

  return
}

// blue()
func (this *QRgba64) Blue(args ...interface{}) (ret interface{}) {
  // blue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba644blueEv
    // invoke: quint16 blue()
    var ret0 = C.C_ZNK7QRgba644blueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "quint16"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "blue", args)
  }

  return
}

// alpha8()
func (this *QRgba64) Alpha8(args ...interface{}) (ret interface{}) {
  // alpha8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba646alpha8Ev
    // invoke: quint8 alpha8()
    var ret0 = C.C_ZNK7QRgba646alpha8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "quint8"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "alpha8", args)
  }

  return
}

// green8()
func (this *QRgba64) Green8(args ...interface{}) (ret interface{}) {
  // green8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba646green8Ev
    // invoke: quint8 green8()
    var ret0 = C.C_ZNK7QRgba646green8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "quint8"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "green8", args)
  }

  return
}

// setAlpha(quint16)
func (this *QRgba64) Setalpha(args ...interface{}) () {
  // setAlpha(quint16)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "quint16"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba648setAlphaEt
    // invoke: void setAlpha(quint16)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    C.C_ZN7QRgba648setAlphaEt(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRgba64", "setAlpha", args)
  }

  return
}

// unpremultiplied()
func (this *QRgba64) Unpremultiplied(args ...interface{}) (ret interface{}) {
  // unpremultiplied()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba6415unpremultipliedEv
    // invoke: QRgba64 unpremultiplied()
    var ret0 = C.C_ZNK7QRgba6415unpremultipliedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "unpremultiplied", args)
  }

  return
}

// red()
func (this *QRgba64) Red(args ...interface{}) (ret interface{}) {
  // red()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba643redEv
    // invoke: quint16 red()
    var ret0 = C.C_ZNK7QRgba643redEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "quint16"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "red", args)
  }

  return
}

// fromArgb32(uint)
func (this *QRgba64) Fromargb32_S(args ...interface{}) (ret interface{}) {
  // fromArgb32(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba6410fromArgb32Ej
    // invoke: QRgba64 fromArgb32(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QRgba6410fromArgb32Ej(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "fromArgb32", args)
  }

  return
}

// red8()
func (this *QRgba64) Red8(args ...interface{}) (ret interface{}) {
  // red8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba644red8Ev
    // invoke: quint8 red8()
    var ret0 = C.C_ZNK7QRgba644red8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "quint8"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "red8", args)
  }

  return
}

// blue8()
func (this *QRgba64) Blue8(args ...interface{}) (ret interface{}) {
  // blue8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba645blue8Ev
    // invoke: quint8 blue8()
    var ret0 = C.C_ZNK7QRgba645blue8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "quint8"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "blue8", args)
  }

  return
}

// toRgb16()
func (this *QRgba64) Torgb16(args ...interface{}) (ret interface{}) {
  // toRgb16()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba647toRgb16Ev
    // invoke: ushort toRgb16()
    var ret0 = C.C_ZNK7QRgba647toRgb16Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "toRgb16", args)
  }

  return
}

// premultiplied()
func (this *QRgba64) Premultiplied(args ...interface{}) (ret interface{}) {
  // premultiplied()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba6413premultipliedEv
    // invoke: QRgba64 premultiplied()
    var ret0 = C.C_ZNK7QRgba6413premultipliedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "premultiplied", args)
  }

  return
}

// alpha()
func (this *QRgba64) Alpha(args ...interface{}) (ret interface{}) {
  // alpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba645alphaEv
    // invoke: quint16 alpha()
    var ret0 = C.C_ZNK7QRgba645alphaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "quint16"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "alpha", args)
  }

  return
}

// toArgb32()
func (this *QRgba64) Toargb32(args ...interface{}) (ret interface{}) {
  // toArgb32()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba648toArgb32Ev
    // invoke: uint toArgb32()
    var ret0 = C.C_ZNK7QRgba648toArgb32Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "toArgb32", args)
  }

  return
}

// green()
func (this *QRgba64) Green(args ...interface{}) (ret interface{}) {
  // green()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRgba645greenEv
    // invoke: quint16 green()
    var ret0 = C.C_ZNK7QRgba645greenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "quint16"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "green", args)
  }

  return
}

// fromRgba(quint8, quint8, quint8, quint8)
func (this *QRgba64) Fromrgba_S(args ...interface{}) (ret interface{}) {
  // fromRgba(quint8, quint8, quint8, quint8)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "quint8"
  vtys[0][1] = qtrt.ByteTy(false) // "quint8"
  vtys[0][2] = qtrt.ByteTy(false) // "quint8"
  vtys[0][3] = qtrt.ByteTy(false) // "quint8"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRgba648fromRgbaEhhhh
    // invoke: QRgba64 fromRgba(quint8, quint8, quint8, quint8)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.uchar(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = C.uchar(args[3].(byte))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN7QRgba648fromRgbaEhhhh(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRgba64{}) // "QRgba64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QRgba64", "fromRgba", args)
  }

  return
}

// <= body block end

