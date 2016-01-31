package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qopenglpaintdevice.h
// dst-file: /src/gui/qopenglpaintdevice.go
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
  // proto:  void QOpenGLPaintDevice::setDotsPerMeterX(qreal );
extern void C_ZN18QOpenGLPaintDevice16setDotsPerMeterXEd(void* qthis, double arg0); // 4
  // proto:  void QOpenGLPaintDevice::setDotsPerMeterY(qreal );
extern void C_ZN18QOpenGLPaintDevice16setDotsPerMeterYEd(void* qthis, double arg0); // 4
  // proto:  void QOpenGLPaintDevice::ensureActiveTarget();
extern void C_ZN18QOpenGLPaintDevice18ensureActiveTargetEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::setDevicePixelRatio(qreal devicePixelRatio);
extern void C_ZN18QOpenGLPaintDevice19setDevicePixelRatioEd(void* qthis, double arg0); // 4
  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterY();
extern double C_ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::setSize(const QSize & size);
extern void C_ZN18QOpenGLPaintDevice7setSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLPaintDevice::devType();
extern int32_t C_ZNK18QOpenGLPaintDevice7devTypeEv(void* qthis); // 2
  // proto:  QOpenGLContext * QOpenGLPaintDevice::context();
extern void C_ZNK18QOpenGLPaintDevice7contextEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::~QOpenGLPaintDevice();
extern void C_ZN18QOpenGLPaintDeviceD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLPaintDevice::paintFlipped();
extern bool C_ZNK18QOpenGLPaintDevice12paintFlippedEv(void* qthis); // 4
  // proto:  QPaintEngine * QOpenGLPaintDevice::paintEngine();
extern void* C_ZNK18QOpenGLPaintDevice11paintEngineEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::setPaintFlipped(bool flipped);
extern void C_ZN18QOpenGLPaintDevice15setPaintFlippedEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(const QSize & size);
extern void* C_ZN18QOpenGLPaintDeviceC2ERK5QSize(void* arg0); // 3
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice();
extern void* C_ZN18QOpenGLPaintDeviceC2Ev(); // 3
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(int width, int height);
extern void* C_ZN18QOpenGLPaintDeviceC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterX();
extern double C_ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(void* qthis); // 4
  // proto:  QSize QOpenGLPaintDevice::size();
extern void* C_ZNK18QOpenGLPaintDevice4sizeEv(void* qthis); // 4
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

// class sizeof(QOpenGLPaintDevice)=1
type QOpenGLPaintDevice struct {
  /*qbase*/ QPaintDevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setDotsPerMeterX(qreal)
func (this *QOpenGLPaintDevice) Setdotspermeterx(args ...interface{}) () {
  // setDotsPerMeterX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterXEd
    // invoke: void setDotsPerMeterX(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice16setDotsPerMeterXEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterX", args)
  }

  return
}

// setDotsPerMeterY(qreal)
func (this *QOpenGLPaintDevice) Setdotspermetery(args ...interface{}) () {
  // setDotsPerMeterY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterYEd
    // invoke: void setDotsPerMeterY(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice16setDotsPerMeterYEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterY", args)
  }

  return
}

// ensureActiveTarget()
func (this *QOpenGLPaintDevice) Ensureactivetarget(args ...interface{}) () {
  // ensureActiveTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice18ensureActiveTargetEv
    // invoke: void ensureActiveTarget()
    C.C_ZN18QOpenGLPaintDevice18ensureActiveTargetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "ensureActiveTarget", args)
  }

  return
}

// setDevicePixelRatio(qreal)
func (this *QOpenGLPaintDevice) Setdevicepixelratio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice19setDevicePixelRatioEd
    // invoke: void setDevicePixelRatio(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice19setDevicePixelRatioEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDevicePixelRatio", args)
  }

  return
}

// dotsPerMeterY()
func (this *QOpenGLPaintDevice) Dotspermetery(args ...interface{}) (ret interface{}) {
  // dotsPerMeterY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterYEv
    // invoke: qreal dotsPerMeterY()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterY", args)
  }

  return
}

// setSize(const class QSize &)
func (this *QOpenGLPaintDevice) Setsize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice7setSizeERK5QSize
    // invoke: void setSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice7setSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setSize", args)
  }

  return
}

// devType()
func (this *QOpenGLPaintDevice) Devtype(args ...interface{}) (ret interface{}) {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7devTypeEv
    // invoke: int devType()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice7devTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "devType", args)
  }

  return
}

// context()
func (this *QOpenGLPaintDevice) Context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7contextEv
    // invoke: QOpenGLContext * context()
    C.C_ZNK18QOpenGLPaintDevice7contextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "context", args)
  }

  return
}

// ~QOpenGLPaintDevice()
func (this *QOpenGLPaintDevice) Freeqopenglpaintdevice(args ...interface{}) () {
  // ~QOpenGLPaintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDeviceD0Ev
    // invoke: void ~QOpenGLPaintDevice()
    C.C_ZN18QOpenGLPaintDeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "~QOpenGLPaintDevice", args)
  }

  return
}

// paintFlipped()
func (this *QOpenGLPaintDevice) Paintflipped(args ...interface{}) (ret interface{}) {
  // paintFlipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice12paintFlippedEv
    // invoke: bool paintFlipped()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice12paintFlippedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintFlipped", args)
  }

  return
}

// paintEngine()
func (this *QOpenGLPaintDevice) Paintengine(args ...interface{}) (ret interface{}) {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice11paintEngineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintEngine", args)
  }

  return
}

// setPaintFlipped(_Bool)
func (this *QOpenGLPaintDevice) Setpaintflipped(args ...interface{}) () {
  // setPaintFlipped(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice15setPaintFlippedEb
    // invoke: void setPaintFlipped(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice15setPaintFlippedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setPaintFlipped", args)
  }

  return
}

// QOpenGLPaintDevice(const class QSize &)
func NewQOpenGLPaintDevice(args ...interface{}) *QOpenGLPaintDevice {
  // QOpenGLPaintDevice(const class QSize &)
  // QOpenGLPaintDevice()
  // QOpenGLPaintDevice(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDeviceC1ERK5QSize
    // invoke: void QOpenGLPaintDevice(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2ERK5QSize(arg0)
    return &QOpenGLPaintDevice{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QOpenGLPaintDeviceC1Ev
    // invoke: void QOpenGLPaintDevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2Ev()
    return &QOpenGLPaintDevice{qclsinst:qthis}
  case 2:
    // invoke: _ZN18QOpenGLPaintDeviceC1Eii
    // invoke: void QOpenGLPaintDevice(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2Eii(arg0, arg1)
    return &QOpenGLPaintDevice{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "QOpenGLPaintDevice", args)
  }

  return nil // QOpenGLPaintDevice{qclsinst:qthis}
}

// dotsPerMeterX()
func (this *QOpenGLPaintDevice) Dotspermeterx(args ...interface{}) (ret interface{}) {
  // dotsPerMeterX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterXEv
    // invoke: qreal dotsPerMeterX()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterX", args)
  }

  return
}

// size()
func (this *QOpenGLPaintDevice) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK18QOpenGLPaintDevice4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "size", args)
  }

  return
}

// <= body block end

