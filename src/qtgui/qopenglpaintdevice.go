package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QOpenGLPaintDevice)=1
type QOpenGLPaintDevice struct {
  /*qbase*/ QPaintDevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setDotsPerMeterX(qreal)
func (this *QOpenGLPaintDevice) SetDotsPerMeterX(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice16setDotsPerMeterXEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterX", args)
  }

  return
}

// setDotsPerMeterY(qreal)
func (this *QOpenGLPaintDevice) SetDotsPerMeterY(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice16setDotsPerMeterYEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterY", args)
  }

  return
}

// ensureActiveTarget()
func (this *QOpenGLPaintDevice) EnsureActiveTarget(args ...interface{}) () {
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
    C.C_ZN18QOpenGLPaintDevice18ensureActiveTargetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "ensureActiveTarget", args)
  }

  return
}

// setDevicePixelRatio(qreal)
func (this *QOpenGLPaintDevice) SetDevicePixelRatio(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice19setDevicePixelRatioEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDevicePixelRatio", args)
  }

  return
}

// dotsPerMeterY()
func (this *QOpenGLPaintDevice) DotsPerMeterY(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterY", args)
  }

  return
}

// setSize(const class QSize &)
func (this *QOpenGLPaintDevice) SetSize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice7setSizeERK5QSize
    // invoke: void setSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLPaintDevice7setSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setSize", args)
  }

  return
}

// devType()
func (this *QOpenGLPaintDevice) DevType(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice7devTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
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
    C.C_ZNK18QOpenGLPaintDevice7contextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "context", args)
  }

  return
}

// ~QOpenGLPaintDevice()
func (this *QOpenGLPaintDevice) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QOpenGLPaintDeviceD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "~QOpenGLPaintDevice", args)
  }

  return
}

// paintFlipped()
func (this *QOpenGLPaintDevice) PaintFlipped(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice12paintFlippedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintFlipped", args)
  }

  return
}

// paintEngine()
func (this *QOpenGLPaintDevice) PaintEngine(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice11paintEngineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintEngine", args)
  }

  return
}

// setPaintFlipped(_Bool)
func (this *QOpenGLPaintDevice) SetPaintFlipped(args ...interface{}) () {
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
    C.C_ZN18QOpenGLPaintDevice15setPaintFlippedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setPaintFlipped", args)
  }

  return
}

// QOpenGLPaintDevice(const class QSize &)
func GcfreeQOpenGLPaintDevice(this *QOpenGLPaintDevice) {
  qtrt.UniverseFree(this)
}
func NewQOpenGLPaintDevice(args ...interface{}) *QOpenGLPaintDevice {
  // QOpenGLPaintDevice(const class QSize &)
  // QOpenGLPaintDevice()
  // QOpenGLPaintDevice(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
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
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2ERK5QSize(arg0)
    this := &QOpenGLPaintDevice{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQOpenGLPaintDevice)
    return this
  case 1:
    // invoke: _ZN18QOpenGLPaintDeviceC1Ev
    // invoke: void QOpenGLPaintDevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2Ev()
    this := &QOpenGLPaintDevice{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQOpenGLPaintDevice)
    return this
  case 2:
    // invoke: _ZN18QOpenGLPaintDeviceC1Eii
    // invoke: void QOpenGLPaintDevice(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLPaintDeviceC2Eii(arg0, arg1)
    this := &QOpenGLPaintDevice{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQOpenGLPaintDevice)
    return this
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "QOpenGLPaintDevice", args)
  }

  return nil // QOpenGLPaintDevice{Qclsinst:qthis}
}

// dotsPerMeterX()
func (this *QOpenGLPaintDevice) DotsPerMeterX(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
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
    var ret0 = C.C_ZNK18QOpenGLPaintDevice4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "size", args)
  }

  return
}

// <= body block end

