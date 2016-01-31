package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK12QPaintDevice12physicalDpiXEv(void* qthis); // 2
  // proto:  int QPaintDevice::physicalDpiY();
extern void C_ZNK12QPaintDevice12physicalDpiYEv(void* qthis); // 2
  // proto:  int QPaintDevice::heightMM();
extern void C_ZNK12QPaintDevice8heightMMEv(void* qthis); // 2
  // proto:  int QPaintDevice::devicePixelRatio();
extern void C_ZNK12QPaintDevice16devicePixelRatioEv(void* qthis); // 2
  // proto:  int QPaintDevice::logicalDpiX();
extern void C_ZNK12QPaintDevice11logicalDpiXEv(void* qthis); // 2
  // proto:  int QPaintDevice::depth();
extern void C_ZNK12QPaintDevice5depthEv(void* qthis); // 2
  // proto:  int QPaintDevice::height();
extern void C_ZNK12QPaintDevice6heightEv(void* qthis); // 2
  // proto:  int QPaintDevice::colorCount();
extern void C_ZNK12QPaintDevice10colorCountEv(void* qthis); // 2
  // proto:  int QPaintDevice::width();
extern void C_ZNK12QPaintDevice5widthEv(void* qthis); // 2
  // proto:  int QPaintDevice::devType();
extern void C_ZNK12QPaintDevice7devTypeEv(void* qthis); // 2
  // proto:  int QPaintDevice::widthMM();
extern void C_ZNK12QPaintDevice7widthMMEv(void* qthis); // 2
  // proto:  int QPaintDevice::logicalDpiY();
extern void C_ZNK12QPaintDevice11logicalDpiYEv(void* qthis); // 2
  // proto:  void QPaintDevice::~QPaintDevice();
extern void C_ZN12QPaintDeviceD2Ev(void* qthis); // 4
  // proto:  bool QPaintDevice::paintingActive();
extern void C_ZNK12QPaintDevice14paintingActiveEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// physicalDpiX()
func (this *QPaintDevice) physicalDpiX(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice12physicalDpiXEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiX", args)
  }

}

// physicalDpiY()
func (this *QPaintDevice) physicalDpiY(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice12physicalDpiYEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiY", args)
  }

}

// heightMM()
func (this *QPaintDevice) heightMM(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice8heightMMEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "heightMM", args)
  }

}

// devicePixelRatio()
func (this *QPaintDevice) devicePixelRatio(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice16devicePixelRatioEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "devicePixelRatio", args)
  }

}

// logicalDpiX()
func (this *QPaintDevice) logicalDpiX(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice11logicalDpiXEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiX", args)
  }

}

// depth()
func (this *QPaintDevice) depth(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice5depthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "depth", args)
  }

}

// height()
func (this *QPaintDevice) height(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "height", args)
  }

}

// colorCount()
func (this *QPaintDevice) colorCount(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice10colorCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "colorCount", args)
  }

}

// width()
func (this *QPaintDevice) width(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "width", args)
  }

}

// devType()
func (this *QPaintDevice) devType(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice7devTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "devType", args)
  }

}

// widthMM()
func (this *QPaintDevice) widthMM(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice7widthMMEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "widthMM", args)
  }

}

// logicalDpiY()
func (this *QPaintDevice) logicalDpiY(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice11logicalDpiYEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiY", args)
  }

}

// ~QPaintDevice()
func (this *QPaintDevice) FreeQPaintDevice(args ...interface{}) () {
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
    C.C_ZN12QPaintDeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPaintDevice", "~QPaintDevice", args)
  }

}

// paintingActive()
func (this *QPaintDevice) paintingActive(args ...interface{}) () {
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
    var ret = C.C_ZNK12QPaintDevice14paintingActiveEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintingActive", args)
  }

}

// <= body block end

