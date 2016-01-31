package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK14QSurfaceFormat12swapIntervalEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::depthBufferSize();
extern void C_ZNK14QSurfaceFormat15depthBufferSizeEv(void* qthis); // 4
  // proto: static QSurfaceFormat QSurfaceFormat::defaultFormat();
extern void C_ZN14QSurfaceFormat13defaultFormatEv(); // 4
  // proto:  int QSurfaceFormat::alphaBufferSize();
extern void C_ZNK14QSurfaceFormat15alphaBufferSizeEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setSamples(int numSamples);
extern void C_ZN14QSurfaceFormat10setSamplesEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::setVersion(int major, int minor);
extern void C_ZN14QSurfaceFormat10setVersionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QSurfaceFormat::minorVersion();
extern void C_ZNK14QSurfaceFormat12minorVersionEv(void* qthis); // 4
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
extern void C_ZNK14QSurfaceFormat7samplesEv(void* qthis); // 4
  // proto:  QSurfaceFormat::SwapBehavior QSurfaceFormat::swapBehavior();
extern void C_ZNK14QSurfaceFormat12swapBehaviorEv(void* qthis); // 4
  // proto:  QSurfaceFormat::OpenGLContextProfile QSurfaceFormat::profile();
extern void C_ZNK14QSurfaceFormat7profileEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::majorVersion();
extern void C_ZNK14QSurfaceFormat12majorVersionEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setStencilBufferSize(int size);
extern void C_ZN14QSurfaceFormat20setStencilBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSurfaceFormat::hasAlpha();
extern void C_ZNK14QSurfaceFormat8hasAlphaEv(void* qthis); // 4
  // proto:  bool QSurfaceFormat::stereo();
extern void C_ZNK14QSurfaceFormat6stereoEv(void* qthis); // 2
  // proto:  void QSurfaceFormat::setStereo(bool enable);
extern void C_ZN14QSurfaceFormat9setStereoEb(void* qthis, bool arg0); // 4
  // proto:  int QSurfaceFormat::redBufferSize();
extern void C_ZNK14QSurfaceFormat13redBufferSizeEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::greenBufferSize();
extern void C_ZNK14QSurfaceFormat15greenBufferSizeEv(void* qthis); // 4
  // proto:  void QSurfaceFormat::setMajorVersion(int majorVersion);
extern void C_ZN14QSurfaceFormat15setMajorVersionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSurfaceFormat::~QSurfaceFormat();
extern void C_ZN14QSurfaceFormatD2Ev(void* qthis); // 4
  // proto:  void QSurfaceFormat::setAlphaBufferSize(int size);
extern void C_ZN14QSurfaceFormat18setAlphaBufferSizeEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSurfaceFormat::stencilBufferSize();
extern void C_ZNK14QSurfaceFormat17stencilBufferSizeEv(void* qthis); // 4
  // proto:  QSurfaceFormat::FormatOptions QSurfaceFormat::options();
extern void C_ZNK14QSurfaceFormat7optionsEv(void* qthis); // 4
  // proto:  int QSurfaceFormat::blueBufferSize();
extern void C_ZNK14QSurfaceFormat14blueBufferSizeEv(void* qthis); // 4
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
func (this *QSurfaceFormat) swapInterval(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat12swapIntervalEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "swapInterval", args)
  }

}

// depthBufferSize()
func (this *QSurfaceFormat) depthBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat15depthBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "depthBufferSize", args)
  }

}

// defaultFormat()
func (this *QSurfaceFormat) defaultFormat_s(args ...interface{}) () {
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
    var ret = C.C_ZN14QSurfaceFormat13defaultFormatEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "defaultFormat", args)
  }

}

// alphaBufferSize()
func (this *QSurfaceFormat) alphaBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat15alphaBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "alphaBufferSize", args)
  }

}

// setSamples(int)
func (this *QSurfaceFormat) setSamples(args ...interface{}) () {
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

}

// setVersion(int, int)
func (this *QSurfaceFormat) setVersion(args ...interface{}) () {
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

}

// minorVersion()
func (this *QSurfaceFormat) minorVersion(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat12minorVersionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "minorVersion", args)
  }

}

// setSwapInterval(int)
func (this *QSurfaceFormat) setSwapInterval(args ...interface{}) () {
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

}

// setMinorVersion(int)
func (this *QSurfaceFormat) setMinorVersion(args ...interface{}) () {
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

}

// setBlueBufferSize(int)
func (this *QSurfaceFormat) setBlueBufferSize(args ...interface{}) () {
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

}

// setRedBufferSize(int)
func (this *QSurfaceFormat) setRedBufferSize(args ...interface{}) () {
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

}

// setDefaultFormat(const class QSurfaceFormat &)
func (this *QSurfaceFormat) setDefaultFormat_s(args ...interface{}) () {
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

}

// version()
func (this *QSurfaceFormat) version(args ...interface{}) () {
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

}

// samples()
func (this *QSurfaceFormat) samples(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat7samplesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "samples", args)
  }

}

// swapBehavior()
func (this *QSurfaceFormat) swapBehavior(args ...interface{}) () {
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

}

// profile()
func (this *QSurfaceFormat) profile(args ...interface{}) () {
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

}

// majorVersion()
func (this *QSurfaceFormat) majorVersion(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat12majorVersionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "majorVersion", args)
  }

}

// setStencilBufferSize(int)
func (this *QSurfaceFormat) setStencilBufferSize(args ...interface{}) () {
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

}

// hasAlpha()
func (this *QSurfaceFormat) hasAlpha(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat8hasAlphaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "hasAlpha", args)
  }

}

// stereo()
func (this *QSurfaceFormat) stereo(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat6stereoEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stereo", args)
  }

}

// setStereo(_Bool)
func (this *QSurfaceFormat) setStereo(args ...interface{}) () {
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

}

// redBufferSize()
func (this *QSurfaceFormat) redBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat13redBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "redBufferSize", args)
  }

}

// greenBufferSize()
func (this *QSurfaceFormat) greenBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat15greenBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "greenBufferSize", args)
  }

}

// setMajorVersion(int)
func (this *QSurfaceFormat) setMajorVersion(args ...interface{}) () {
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

}

// ~QSurfaceFormat()
func (this *QSurfaceFormat) FreeQSurfaceFormat(args ...interface{}) () {
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

}

// setAlphaBufferSize(int)
func (this *QSurfaceFormat) setAlphaBufferSize(args ...interface{}) () {
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

}

// stencilBufferSize()
func (this *QSurfaceFormat) stencilBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat17stencilBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stencilBufferSize", args)
  }

}

// options()
func (this *QSurfaceFormat) options(args ...interface{}) () {
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

}

// blueBufferSize()
func (this *QSurfaceFormat) blueBufferSize(args ...interface{}) () {
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
    var ret = C.C_ZNK14QSurfaceFormat14blueBufferSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "blueBufferSize", args)
  }

}

// setDepthBufferSize(int)
func (this *QSurfaceFormat) setDepthBufferSize(args ...interface{}) () {
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

}

// setGreenBufferSize(int)
func (this *QSurfaceFormat) setGreenBufferSize(args ...interface{}) () {
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

}

// renderableType()
func (this *QSurfaceFormat) renderableType(args ...interface{}) () {
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

}

// <= body block end

