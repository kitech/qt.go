package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  int QPaintDevice::physicalDpiY();
extern void _ZNK12QPaintDevice12physicalDpiYEv(void* qthis);
  // proto:  int QPaintDevice::heightMM();
extern void _ZNK12QPaintDevice8heightMMEv(void* qthis);
  // proto:  int QPaintDevice::colorCount();
extern void _ZNK12QPaintDevice10colorCountEv(void* qthis);
  // proto:  int QPaintDevice::physicalDpiX();
extern void _ZNK12QPaintDevice12physicalDpiXEv(void* qthis);
  // proto:  int QPaintDevice::widthMM();
extern void _ZNK12QPaintDevice7widthMMEv(void* qthis);
  // proto:  int QPaintDevice::devType();
extern void _ZNK12QPaintDevice7devTypeEv(void* qthis);
  // proto:  bool QPaintDevice::paintingActive();
extern void _ZNK12QPaintDevice14paintingActiveEv(void* qthis);
  // proto:  int QPaintDevice::width();
extern void _ZNK12QPaintDevice5widthEv(void* qthis);
  // proto:  void QPaintDevice::QPaintDevice(const QPaintDevice & );
extern void* dector_ZN12QPaintDeviceC1ERKS_(void* arg0);
extern void _ZN12QPaintDeviceC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPaintDevice::QPaintDevice();
extern void* dector_ZN12QPaintDeviceC1Ev();
extern void _ZN12QPaintDeviceC1Ev(void* qthis);
  // proto:  int QPaintDevice::devicePixelRatio();
extern void _ZNK12QPaintDevice16devicePixelRatioEv(void* qthis);
  // proto:  int QPaintDevice::height();
extern void _ZNK12QPaintDevice6heightEv(void* qthis);
  // proto:  int QPaintDevice::depth();
extern void _ZNK12QPaintDevice5depthEv(void* qthis);
  // proto:  QPaintEngine * QPaintDevice::paintEngine();
extern void _ZNK12QPaintDevice11paintEngineEv(void* qthis);
  // proto:  int QPaintDevice::logicalDpiY();
extern void _ZNK12QPaintDevice11logicalDpiYEv(void* qthis);
  // proto:  void QPaintDevice::~QPaintDevice();
extern void _ZN12QPaintDeviceD0Ev(void* qthis);
  // proto:  int QPaintDevice::logicalDpiX();
extern void _ZNK12QPaintDevice11logicalDpiXEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QPaintDevice::physicalDpiY();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiY", args)
  }

}

  // proto:  int QPaintDevice::heightMM();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "heightMM", args)
  }

}

  // proto:  int QPaintDevice::colorCount();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "colorCount", args)
  }

}

  // proto:  int QPaintDevice::physicalDpiX();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiX", args)
  }

}

  // proto:  int QPaintDevice::widthMM();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "widthMM", args)
  }

}

  // proto:  int QPaintDevice::devType();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "devType", args)
  }

}

  // proto:  bool QPaintDevice::paintingActive();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintingActive", args)
  }

}

  // proto:  int QPaintDevice::width();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "width", args)
  }

}

  // proto:  void QPaintDevice::QPaintDevice(const QPaintDevice & );
func NewQPaintDevice(args ...interface{}) QPaintDevice {
  return QPaintDevice{}
}

  // proto:  int QPaintDevice::devicePixelRatio();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "devicePixelRatio", args)
  }

}

  // proto:  int QPaintDevice::height();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "height", args)
  }

}

  // proto:  int QPaintDevice::depth();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "depth", args)
  }

}

  // proto:  QPaintEngine * QPaintDevice::paintEngine();
func (this *QPaintDevice) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11paintEngineEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintEngine", args)
  }

}

  // proto:  int QPaintDevice::logicalDpiY();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiY", args)
  }

}

  // proto:  void QPaintDevice::~QPaintDevice();
func (this *QPaintDevice) FreeQPaintDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintDevice", "~QPaintDevice", args)
  }

}

  // proto:  int QPaintDevice::logicalDpiX();
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
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiX", args)
  }

}

// <= body block end

