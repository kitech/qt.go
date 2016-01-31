package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qsurfaceformat.h
// dst-file: /src/gui/qsurfaceformat.go
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
  // proto:  void QSurfaceFormat::QSurfaceFormat(const QSurfaceFormat & other);
extern void* C_ZN14QSurfaceFormatC2ERKS_(void* arg0); // 3
  // proto:  void QSurfaceFormat::QSurfaceFormat();
extern void* C_ZN14QSurfaceFormatC2Ev(); // 3
  // proto:  int QSurfaceFormat::swapInterval();
extern int32_t C_ZNK14QSurfaceFormat12swapIntervalEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::depthBufferSize();
extern int32_t C_ZNK14QSurfaceFormat15depthBufferSizeEv(void* qthis); // 4
  // proto: static QSurfaceFormat QSurfaceFormat::defaultFormat();
extern void* C_ZN14QSurfaceFormat13defaultFormatEv(); // 4
  // proto:  int QSurfaceFormat::alphaBufferSize();
extern int32_t C_ZNK14QSurfaceFormat15alphaBufferSizeEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setSamples(int numSamples);
extern void C_ZN14QSurfaceFormat10setSamplesEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setVersion(int major, int minor);
extern void C_ZN14QSurfaceFormat10setVersionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QSurfaceFormat::minorVersion();
extern int32_t C_ZNK14QSurfaceFormat12minorVersionEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setSwapInterval(int interval);
extern void C_ZN14QSurfaceFormat15setSwapIntervalEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setMinorVersion(int minorVersion);
extern void C_ZN14QSurfaceFormat15setMinorVersionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setBlueBufferSize(int size);
extern void C_ZN14QSurfaceFormat17setBlueBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setRedBufferSize(int size);
extern void C_ZN14QSurfaceFormat16setRedBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto: static void QSurfaceFormat::setDefaultFormat(const QSurfaceFormat & format);
extern void C_ZN14QSurfaceFormat16setDefaultFormatERKS_(void* arg0); // 4
  // proto:  QPair<int, int> QSurfaceFormat::version();
extern void C_ZNK14QSurfaceFormat7versionEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::samples();
extern int32_t C_ZNK14QSurfaceFormat7samplesEv(void* qthis); // 4
  // proto:  QSurfaceFormat::SwapBehavior QSurfaceFormat::swapBehavior();
extern void C_ZNK14QSurfaceFormat12swapBehaviorEv(void* qthis); // 4
  // proto:  QSurfaceFormat::OpenGLContextProfile QSurfaceFormat::profile();
extern void C_ZNK14QSurfaceFormat7profileEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::majorVersion();
extern int32_t C_ZNK14QSurfaceFormat12majorVersionEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setStencilBufferSize(int size);
extern void C_ZN14QSurfaceFormat20setStencilBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSurfaceFormat::hasAlpha();
extern bool C_ZNK14QSurfaceFormat8hasAlphaEv(void* qthis); // 4
  // proto:  bool QSurfaceFormat::stereo();
extern bool C_ZNK14QSurfaceFormat6stereoEv(void* qthis); // 2
  // proto:  void QSurfaceFormat::setStereo(bool enable);
extern void C_ZN14QSurfaceFormat9setStereoEb(void* qthis, bool arg0); // 4
  // proto:  int QSurfaceFormat::redBufferSize();
extern int32_t C_ZNK14QSurfaceFormat13redBufferSizeEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::greenBufferSize();
extern int32_t C_ZNK14QSurfaceFormat15greenBufferSizeEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setMajorVersion(int majorVersion);
extern void C_ZN14QSurfaceFormat15setMajorVersionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::~QSurfaceFormat();
extern void C_ZN14QSurfaceFormatD2Ev(void* qthis); // 4
  // proto:  void QSurfaceFormat::setAlphaBufferSize(int size);
extern void C_ZN14QSurfaceFormat18setAlphaBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSurfaceFormat::stencilBufferSize();
extern int32_t C_ZNK14QSurfaceFormat17stencilBufferSizeEv(void* qthis); // 4
  // proto:  QSurfaceFormat::FormatOptions QSurfaceFormat::options();
extern void C_ZNK14QSurfaceFormat7optionsEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::blueBufferSize();
extern int32_t C_ZNK14QSurfaceFormat14blueBufferSizeEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setDepthBufferSize(int size);
extern void C_ZN14QSurfaceFormat18setDepthBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setGreenBufferSize(int size);
extern void C_ZN14QSurfaceFormat18setGreenBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  QSurfaceFormat::RenderableType QSurfaceFormat::renderableType();
extern void C_ZNK14QSurfaceFormat14renderableTypeEv(void* qthis); // 4
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

// class sizeof(QSurfaceFormat)=8
type QSurfaceFormat struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QSurfaceFormat(const class QSurfaceFormat &)
func NewQSurfaceFormat(args ...interface{}) *QSurfaceFormat {
  // QSurfaceFormat(const class QSurfaceFormat &)
  // QSurfaceFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormatC1ERKS_
    // invoke: void QSurfaceFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QSurfaceFormatC2ERKS_(arg0)
    return &QSurfaceFormat{qclsinst:qthis}
  case 1:
    // invoke: _ZN14QSurfaceFormatC1Ev
    // invoke: void QSurfaceFormat()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QSurfaceFormatC2Ev()
    return &QSurfaceFormat{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "QSurfaceFormat", args)
  }

  return nil // QSurfaceFormat{qclsinst:qthis}
}

// swapInterval()
func (this *QSurfaceFormat) Swapinterval(args ...interface{}) (ret interface{}) {
  // swapInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12swapIntervalEv
    // invoke: int swapInterval()
    var ret0 = C.C_ZNK14QSurfaceFormat12swapIntervalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "swapInterval", args)
  }

  return
}

// depthBufferSize()
func (this *QSurfaceFormat) Depthbuffersize(args ...interface{}) (ret interface{}) {
  // depthBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15depthBufferSizeEv
    // invoke: int depthBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat15depthBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "depthBufferSize", args)
  }

  return
}

// defaultFormat()
func (this *QSurfaceFormat) Defaultformat_S(args ...interface{}) (ret interface{}) {
  // defaultFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat13defaultFormatEv
    // invoke: QSurfaceFormat defaultFormat()
    var ret0 = C.C_ZN14QSurfaceFormat13defaultFormatEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "defaultFormat", args)
  }

  return
}

// alphaBufferSize()
func (this *QSurfaceFormat) Alphabuffersize(args ...interface{}) (ret interface{}) {
  // alphaBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15alphaBufferSizeEv
    // invoke: int alphaBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat15alphaBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "alphaBufferSize", args)
  }

  return
}

// setSamples(int)
func (this *QSurfaceFormat) Setsamples(args ...interface{}) () {
  // setSamples(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat10setSamplesEi
    // invoke: void setSamples(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat10setSamplesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSamples", args)
  }

  return
}

// setVersion(int, int)
func (this *QSurfaceFormat) Setversion(args ...interface{}) () {
  // setVersion(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat10setVersionEii
    // invoke: void setVersion(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN14QSurfaceFormat10setVersionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setVersion", args)
  }

  return
}

// minorVersion()
func (this *QSurfaceFormat) Minorversion(args ...interface{}) (ret interface{}) {
  // minorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12minorVersionEv
    // invoke: int minorVersion()
    var ret0 = C.C_ZNK14QSurfaceFormat12minorVersionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "minorVersion", args)
  }

  return
}

// setSwapInterval(int)
func (this *QSurfaceFormat) Setswapinterval(args ...interface{}) () {
  // setSwapInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setSwapIntervalEi
    // invoke: void setSwapInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat15setSwapIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSwapInterval", args)
  }

  return
}

// setMinorVersion(int)
func (this *QSurfaceFormat) Setminorversion(args ...interface{}) () {
  // setMinorVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setMinorVersionEi
    // invoke: void setMinorVersion(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat15setMinorVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMinorVersion", args)
  }

  return
}

// setBlueBufferSize(int)
func (this *QSurfaceFormat) Setbluebuffersize(args ...interface{}) () {
  // setBlueBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat17setBlueBufferSizeEi
    // invoke: void setBlueBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat17setBlueBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setBlueBufferSize", args)
  }

  return
}

// setRedBufferSize(int)
func (this *QSurfaceFormat) Setredbuffersize(args ...interface{}) () {
  // setRedBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat16setRedBufferSizeEi
    // invoke: void setRedBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat16setRedBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setRedBufferSize", args)
  }

  return
}

// setDefaultFormat(const class QSurfaceFormat &)
func (this *QSurfaceFormat) Setdefaultformat_S(args ...interface{}) () {
  // setDefaultFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat16setDefaultFormatERKS_
    // invoke: void setDefaultFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat16setDefaultFormatERKS_(arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDefaultFormat", args)
  }

  return
}

// version()
func (this *QSurfaceFormat) Version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7versionEv
    // invoke: QPair<int, int> version()
    C.C_ZNK14QSurfaceFormat7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "version", args)
  }

  return
}

// samples()
func (this *QSurfaceFormat) Samples(args ...interface{}) (ret interface{}) {
  // samples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7samplesEv
    // invoke: int samples()
    var ret0 = C.C_ZNK14QSurfaceFormat7samplesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "samples", args)
  }

  return
}

// swapBehavior()
func (this *QSurfaceFormat) Swapbehavior(args ...interface{}) () {
  // swapBehavior()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12swapBehaviorEv
    // invoke: QSurfaceFormat::SwapBehavior swapBehavior()
    C.C_ZNK14QSurfaceFormat12swapBehaviorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "swapBehavior", args)
  }

  return
}

// profile()
func (this *QSurfaceFormat) Profile(args ...interface{}) () {
  // profile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7profileEv
    // invoke: QSurfaceFormat::OpenGLContextProfile profile()
    C.C_ZNK14QSurfaceFormat7profileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "profile", args)
  }

  return
}

// majorVersion()
func (this *QSurfaceFormat) Majorversion(args ...interface{}) (ret interface{}) {
  // majorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12majorVersionEv
    // invoke: int majorVersion()
    var ret0 = C.C_ZNK14QSurfaceFormat12majorVersionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "majorVersion", args)
  }

  return
}

// setStencilBufferSize(int)
func (this *QSurfaceFormat) Setstencilbuffersize(args ...interface{}) () {
  // setStencilBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat20setStencilBufferSizeEi
    // invoke: void setStencilBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat20setStencilBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStencilBufferSize", args)
  }

  return
}

// hasAlpha()
func (this *QSurfaceFormat) Hasalpha(args ...interface{}) (ret interface{}) {
  // hasAlpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat8hasAlphaEv
    // invoke: bool hasAlpha()
    var ret0 = C.C_ZNK14QSurfaceFormat8hasAlphaEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "hasAlpha", args)
  }

  return
}

// stereo()
func (this *QSurfaceFormat) Stereo(args ...interface{}) (ret interface{}) {
  // stereo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat6stereoEv
    // invoke: bool stereo()
    var ret0 = C.C_ZNK14QSurfaceFormat6stereoEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stereo", args)
  }

  return
}

// setStereo(_Bool)
func (this *QSurfaceFormat) Setstereo(args ...interface{}) () {
  // setStereo(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat9setStereoEb
    // invoke: void setStereo(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat9setStereoEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStereo", args)
  }

  return
}

// redBufferSize()
func (this *QSurfaceFormat) Redbuffersize(args ...interface{}) (ret interface{}) {
  // redBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat13redBufferSizeEv
    // invoke: int redBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat13redBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "redBufferSize", args)
  }

  return
}

// greenBufferSize()
func (this *QSurfaceFormat) Greenbuffersize(args ...interface{}) (ret interface{}) {
  // greenBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15greenBufferSizeEv
    // invoke: int greenBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat15greenBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "greenBufferSize", args)
  }

  return
}

// setMajorVersion(int)
func (this *QSurfaceFormat) Setmajorversion(args ...interface{}) () {
  // setMajorVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setMajorVersionEi
    // invoke: void setMajorVersion(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat15setMajorVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMajorVersion", args)
  }

  return
}

// ~QSurfaceFormat()
func (this *QSurfaceFormat) Freeqsurfaceformat(args ...interface{}) () {
  // ~QSurfaceFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormatD0Ev
    // invoke: void ~QSurfaceFormat()
    C.C_ZN14QSurfaceFormatD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "~QSurfaceFormat", args)
  }

  return
}

// setAlphaBufferSize(int)
func (this *QSurfaceFormat) Setalphabuffersize(args ...interface{}) () {
  // setAlphaBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setAlphaBufferSizeEi
    // invoke: void setAlphaBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat18setAlphaBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setAlphaBufferSize", args)
  }

  return
}

// stencilBufferSize()
func (this *QSurfaceFormat) Stencilbuffersize(args ...interface{}) (ret interface{}) {
  // stencilBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat17stencilBufferSizeEv
    // invoke: int stencilBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat17stencilBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stencilBufferSize", args)
  }

  return
}

// options()
func (this *QSurfaceFormat) Options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7optionsEv
    // invoke: QSurfaceFormat::FormatOptions options()
    C.C_ZNK14QSurfaceFormat7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "options", args)
  }

  return
}

// blueBufferSize()
func (this *QSurfaceFormat) Bluebuffersize(args ...interface{}) (ret interface{}) {
  // blueBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat14blueBufferSizeEv
    // invoke: int blueBufferSize()
    var ret0 = C.C_ZNK14QSurfaceFormat14blueBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "blueBufferSize", args)
  }

  return
}

// setDepthBufferSize(int)
func (this *QSurfaceFormat) Setdepthbuffersize(args ...interface{}) () {
  // setDepthBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setDepthBufferSizeEi
    // invoke: void setDepthBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat18setDepthBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDepthBufferSize", args)
  }

  return
}

// setGreenBufferSize(int)
func (this *QSurfaceFormat) Setgreenbuffersize(args ...interface{}) () {
  // setGreenBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setGreenBufferSizeEi
    // invoke: void setGreenBufferSize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QSurfaceFormat18setGreenBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setGreenBufferSize", args)
  }

  return
}

// renderableType()
func (this *QSurfaceFormat) Renderabletype(args ...interface{}) () {
  // renderableType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat14renderableTypeEv
    // invoke: QSurfaceFormat::RenderableType renderableType()
    C.C_ZNK14QSurfaceFormat14renderableTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "renderableType", args)
  }

  return
}

// <= body block end

