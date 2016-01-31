package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::setSize(const QSize & size);
extern void C_ZN18QOpenGLPaintDevice7setSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLPaintDevice::devType();
extern void C_ZNK18QOpenGLPaintDevice7devTypeEv(void* qthis); // 2
  // proto:  QOpenGLContext * QOpenGLPaintDevice::context();
extern void C_ZNK18QOpenGLPaintDevice7contextEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::~QOpenGLPaintDevice();
extern void C_ZN18QOpenGLPaintDeviceD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLPaintDevice::paintFlipped();
extern void C_ZNK18QOpenGLPaintDevice12paintFlippedEv(void* qthis); // 4
  // proto:  QPaintEngine * QOpenGLPaintDevice::paintEngine();
extern void C_ZNK18QOpenGLPaintDevice11paintEngineEv(void* qthis); // 4
  // proto:  void QOpenGLPaintDevice::setPaintFlipped(bool flipped);
extern void C_ZN18QOpenGLPaintDevice15setPaintFlippedEb(void* qthis, bool arg0); // 4
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(const QSize & size);
extern void C_ZN18QOpenGLPaintDeviceC2ERK5QSize(void* qthis, void* arg0); // 3
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice();
extern void C_ZN18QOpenGLPaintDeviceC2Ev(void* qthis); // 3
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(int width, int height);
extern void C_ZN18QOpenGLPaintDeviceC2Eii(void* qthis, int32_t arg0, int32_t arg1); // 3
  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterX();
extern void C_ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(void* qthis); // 4
  // proto:  QSize QOpenGLPaintDevice::size();
extern void C_ZNK18QOpenGLPaintDevice4sizeEv(void* qthis); // 4
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
func (this *QOpenGLPaintDevice) setDotsPerMeterX(args ...interface{}) () {
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

}

// setDotsPerMeterY(qreal)
func (this *QOpenGLPaintDevice) setDotsPerMeterY(args ...interface{}) () {
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

}

// ensureActiveTarget()
func (this *QOpenGLPaintDevice) ensureActiveTarget(args ...interface{}) () {
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

}

// setDevicePixelRatio(qreal)
func (this *QOpenGLPaintDevice) setDevicePixelRatio(args ...interface{}) () {
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

}

// dotsPerMeterY()
func (this *QOpenGLPaintDevice) dotsPerMeterY(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterY", args)
  }

}

// setSize(const class QSize &)
func (this *QOpenGLPaintDevice) setSize(args ...interface{}) () {
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

}

// devType()
func (this *QOpenGLPaintDevice) devType(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "devType", args)
  }

}

// context()
func (this *QOpenGLPaintDevice) context(args ...interface{}) () {
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

}

// ~QOpenGLPaintDevice()
func (this *QOpenGLPaintDevice) FreeQOpenGLPaintDevice(args ...interface{}) () {
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

}

// paintFlipped()
func (this *QOpenGLPaintDevice) paintFlipped(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice12paintFlippedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintFlipped", args)
  }

}

// paintEngine()
func (this *QOpenGLPaintDevice) paintEngine(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintEngine", args)
  }

}

// setPaintFlipped(_Bool)
func (this *QOpenGLPaintDevice) setPaintFlipped(args ...interface{}) () {
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

}

// QOpenGLPaintDevice(const class QSize &)
func NewQOpenGLPaintDevice(args ...interface{}) QOpenGLPaintDevice {
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
    C.C_ZN18QOpenGLPaintDeviceC2ERK5QSize(qthis, arg0)
  case 1:
    // invoke: _ZN18QOpenGLPaintDeviceC1Ev
    // invoke: void QOpenGLPaintDevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QOpenGLPaintDeviceC2Ev(qthis)
  case 2:
    // invoke: _ZN18QOpenGLPaintDeviceC1Eii
    // invoke: void QOpenGLPaintDevice(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QOpenGLPaintDeviceC2Eii(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "QOpenGLPaintDevice", args)
  }

  return QOpenGLPaintDevice{}
}

// dotsPerMeterX()
func (this *QOpenGLPaintDevice) dotsPerMeterX(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterX", args)
  }

}

// size()
func (this *QOpenGLPaintDevice) size(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLPaintDevice4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "size", args)
  }

}

// <= body block end

