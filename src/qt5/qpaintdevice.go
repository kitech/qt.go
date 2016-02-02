package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qpaintdevice.h
// dst-file: /src/gui/qpaintdevice.go
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
  // proto:  int QPaintDevice::physicalDpiX();
extern int32_t C_ZNK12QPaintDevice12physicalDpiXEv(void* qthis); // 2
  // proto:  int QPaintDevice::physicalDpiY();
extern int32_t C_ZNK12QPaintDevice12physicalDpiYEv(void* qthis); // 2
  // proto:  int QPaintDevice::heightMM();
extern int32_t C_ZNK12QPaintDevice8heightMMEv(void* qthis); // 2
  // proto:  int QPaintDevice::devicePixelRatio();
extern int32_t C_ZNK12QPaintDevice16devicePixelRatioEv(void* qthis); // 2
  // proto:  int QPaintDevice::logicalDpiX();
extern int32_t C_ZNK12QPaintDevice11logicalDpiXEv(void* qthis); // 2
  // proto:  int QPaintDevice::depth();
extern int32_t C_ZNK12QPaintDevice5depthEv(void* qthis); // 2
  // proto:  int QPaintDevice::height();
extern int32_t C_ZNK12QPaintDevice6heightEv(void* qthis); // 2
  // proto:  int QPaintDevice::colorCount();
extern int32_t C_ZNK12QPaintDevice10colorCountEv(void* qthis); // 2
  // proto:  int QPaintDevice::width();
extern int32_t C_ZNK12QPaintDevice5widthEv(void* qthis); // 2
  // proto:  int QPaintDevice::devType();
extern int32_t C_ZNK12QPaintDevice7devTypeEv(void* qthis); // 2
  // proto:  int QPaintDevice::widthMM();
extern int32_t C_ZNK12QPaintDevice7widthMMEv(void* qthis); // 2
  // proto:  int QPaintDevice::logicalDpiY();
extern int32_t C_ZNK12QPaintDevice11logicalDpiYEv(void* qthis); // 2
  // proto:  void QPaintDevice::~QPaintDevice();
extern void C_ZN12QPaintDeviceD2Ev(void* qthis); // 4
  // proto:  bool QPaintDevice::paintingActive();
extern bool C_ZNK12QPaintDevice14paintingActiveEv(void* qthis); // 2
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

// class sizeof(QPaintDevice)=24
type QPaintDevice struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// physicalDpiX()
func (this *QPaintDevice) Physicaldpix(args ...interface{}) (ret interface{}) {
  // physicalDpiX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice12physicalDpiXEv
    // invoke: int physicalDpiX()
    var ret0 = C.C_ZNK12QPaintDevice12physicalDpiXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiX", args)
  }

  return
}

// physicalDpiY()
func (this *QPaintDevice) Physicaldpiy(args ...interface{}) (ret interface{}) {
  // physicalDpiY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice12physicalDpiYEv
    // invoke: int physicalDpiY()
    var ret0 = C.C_ZNK12QPaintDevice12physicalDpiYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiY", args)
  }

  return
}

// heightMM()
func (this *QPaintDevice) Heightmm(args ...interface{}) (ret interface{}) {
  // heightMM()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice8heightMMEv
    // invoke: int heightMM()
    var ret0 = C.C_ZNK12QPaintDevice8heightMMEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "heightMM", args)
  }

  return
}

// devicePixelRatio()
func (this *QPaintDevice) Devicepixelratio(args ...interface{}) (ret interface{}) {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice16devicePixelRatioEv
    // invoke: int devicePixelRatio()
    var ret0 = C.C_ZNK12QPaintDevice16devicePixelRatioEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "devicePixelRatio", args)
  }

  return
}

// logicalDpiX()
func (this *QPaintDevice) Logicaldpix(args ...interface{}) (ret interface{}) {
  // logicalDpiX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11logicalDpiXEv
    // invoke: int logicalDpiX()
    var ret0 = C.C_ZNK12QPaintDevice11logicalDpiXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiX", args)
  }

  return
}

// depth()
func (this *QPaintDevice) Depth(args ...interface{}) (ret interface{}) {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice5depthEv
    // invoke: int depth()
    var ret0 = C.C_ZNK12QPaintDevice5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "depth", args)
  }

  return
}

// height()
func (this *QPaintDevice) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK12QPaintDevice6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "height", args)
  }

  return
}

// colorCount()
func (this *QPaintDevice) Colorcount(args ...interface{}) (ret interface{}) {
  // colorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice10colorCountEv
    // invoke: int colorCount()
    var ret0 = C.C_ZNK12QPaintDevice10colorCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "colorCount", args)
  }

  return
}

// width()
func (this *QPaintDevice) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK12QPaintDevice5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "width", args)
  }

  return
}

// devType()
func (this *QPaintDevice) Devtype(args ...interface{}) (ret interface{}) {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice7devTypeEv
    // invoke: int devType()
    var ret0 = C.C_ZNK12QPaintDevice7devTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "devType", args)
  }

  return
}

// widthMM()
func (this *QPaintDevice) Widthmm(args ...interface{}) (ret interface{}) {
  // widthMM()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice7widthMMEv
    // invoke: int widthMM()
    var ret0 = C.C_ZNK12QPaintDevice7widthMMEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "widthMM", args)
  }

  return
}

// logicalDpiY()
func (this *QPaintDevice) Logicaldpiy(args ...interface{}) (ret interface{}) {
  // logicalDpiY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11logicalDpiYEv
    // invoke: int logicalDpiY()
    var ret0 = C.C_ZNK12QPaintDevice11logicalDpiYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiY", args)
  }

  return
}

// ~QPaintDevice()
func (this *QPaintDevice) Freeqpaintdevice(args ...interface{}) () {
  // ~QPaintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPaintDeviceD0Ev
    // invoke: void ~QPaintDevice()
    C.C_ZN12QPaintDeviceD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintDevice", "~QPaintDevice", args)
  }

  return
}

// paintingActive()
func (this *QPaintDevice) Paintingactive(args ...interface{}) (ret interface{}) {
  // paintingActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice14paintingActiveEv
    // invoke: bool paintingActive()
    var ret0 = C.C_ZNK12QPaintDevice14paintingActiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintingActive", args)
  }

  return
}

// <= body block end

