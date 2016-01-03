package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto: static QSurfaceFormat QSurfaceFormat::defaultFormat();
extern void _ZN14QSurfaceFormat13defaultFormatEv();
  // proto:  void QSurfaceFormat::setAlphaBufferSize(int size);
extern void _ZN14QSurfaceFormat18setAlphaBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  void QSurfaceFormat::setMinorVersion(int minorVersion);
extern void _ZN14QSurfaceFormat15setMinorVersionEi(void* qthis, int32_t arg0);
  // proto:  int QSurfaceFormat::stencilBufferSize();
extern void _ZNK14QSurfaceFormat17stencilBufferSizeEv(void* qthis);
  // proto:  void QSurfaceFormat::setRedBufferSize(int size);
extern void _ZN14QSurfaceFormat16setRedBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  void QSurfaceFormat::setDepthBufferSize(int size);
extern void _ZN14QSurfaceFormat18setDepthBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  int QSurfaceFormat::majorVersion();
extern void _ZNK14QSurfaceFormat12majorVersionEv(void* qthis);
  // proto:  void QSurfaceFormat::setSamples(int numSamples);
extern void _ZN14QSurfaceFormat10setSamplesEi(void* qthis, int32_t arg0);
  // proto:  void QSurfaceFormat::setMajorVersion(int majorVersion);
extern void _ZN14QSurfaceFormat15setMajorVersionEi(void* qthis, int32_t arg0);
  // proto: static void QSurfaceFormat::setDefaultFormat(const QSurfaceFormat & format);
extern void _ZN14QSurfaceFormat16setDefaultFormatERKS_(void* arg0);
  // proto:  int QSurfaceFormat::greenBufferSize();
extern void _ZNK14QSurfaceFormat15greenBufferSizeEv(void* qthis);
  // proto:  int QSurfaceFormat::minorVersion();
extern void _ZNK14QSurfaceFormat12minorVersionEv(void* qthis);
  // proto:  void QSurfaceFormat::setStencilBufferSize(int size);
extern void _ZN14QSurfaceFormat20setStencilBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  int QSurfaceFormat::swapInterval();
extern void _ZNK14QSurfaceFormat12swapIntervalEv(void* qthis);
  // proto:  void QSurfaceFormat::setVersion(int major, int minor);
extern void _ZN14QSurfaceFormat10setVersionEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  bool QSurfaceFormat::hasAlpha();
extern void _ZNK14QSurfaceFormat8hasAlphaEv(void* qthis);
  // proto:  void QSurfaceFormat::QSurfaceFormat(const QSurfaceFormat & other);
extern void* dector_ZN14QSurfaceFormatC1ERKS_(void* arg0);
extern void _ZN14QSurfaceFormatC1ERKS_(void* qthis, void* arg0);
  // proto:  QPair<int, int> QSurfaceFormat::version();
extern void _ZNK14QSurfaceFormat7versionEv(void* qthis);
  // proto:  int QSurfaceFormat::blueBufferSize();
extern void _ZNK14QSurfaceFormat14blueBufferSizeEv(void* qthis);
  // proto:  void QSurfaceFormat::QSurfaceFormat();
extern void* dector_ZN14QSurfaceFormatC1Ev();
extern void _ZN14QSurfaceFormatC1Ev(void* qthis);
  // proto:  int QSurfaceFormat::redBufferSize();
extern void _ZNK14QSurfaceFormat13redBufferSizeEv(void* qthis);
  // proto:  void QSurfaceFormat::~QSurfaceFormat();
extern void _ZN14QSurfaceFormatD0Ev(void* qthis);
  // proto:  void QSurfaceFormat::setGreenBufferSize(int size);
extern void _ZN14QSurfaceFormat18setGreenBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  int QSurfaceFormat::samples();
extern void _ZNK14QSurfaceFormat7samplesEv(void* qthis);
  // proto:  int QSurfaceFormat::depthBufferSize();
extern void _ZNK14QSurfaceFormat15depthBufferSizeEv(void* qthis);
  // proto:  void QSurfaceFormat::setBlueBufferSize(int size);
extern void _ZN14QSurfaceFormat17setBlueBufferSizeEi(void* qthis, int32_t arg0);
  // proto:  int QSurfaceFormat::alphaBufferSize();
extern void _ZNK14QSurfaceFormat15alphaBufferSizeEv(void* qthis);
  // proto:  bool QSurfaceFormat::stereo();
extern void demth_ZNK14QSurfaceFormat6stereoEv(void* qthis);
  // proto:  void QSurfaceFormat::setSwapInterval(int interval);
extern void _ZN14QSurfaceFormat15setSwapIntervalEi(void* qthis, int32_t arg0);
  // proto:  void QSurfaceFormat::setStereo(bool enable);
extern void _ZN14QSurfaceFormat9setStereoEb(void* qthis, bool arg0);
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

  // proto: static QSurfaceFormat QSurfaceFormat::defaultFormat();
func (this *QSurfaceFormat) defaultFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "defaultFormat", args)
  }

}

  // proto:  void QSurfaceFormat::setAlphaBufferSize(int size);
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
    C._ZN14QSurfaceFormat18setAlphaBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setAlphaBufferSize", args)
  }

}

  // proto:  void QSurfaceFormat::setMinorVersion(int minorVersion);
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
    C._ZN14QSurfaceFormat15setMinorVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMinorVersion", args)
  }

}

  // proto:  int QSurfaceFormat::stencilBufferSize();
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
    C._ZNK14QSurfaceFormat17stencilBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stencilBufferSize", args)
  }

}

  // proto:  void QSurfaceFormat::setRedBufferSize(int size);
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
    C._ZN14QSurfaceFormat16setRedBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setRedBufferSize", args)
  }

}

  // proto:  void QSurfaceFormat::setDepthBufferSize(int size);
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
    C._ZN14QSurfaceFormat18setDepthBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDepthBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::majorVersion();
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
    C._ZNK14QSurfaceFormat12majorVersionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "majorVersion", args)
  }

}

  // proto:  void QSurfaceFormat::setSamples(int numSamples);
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
    C._ZN14QSurfaceFormat10setSamplesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSamples", args)
  }

}

  // proto:  void QSurfaceFormat::setMajorVersion(int majorVersion);
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
    C._ZN14QSurfaceFormat15setMajorVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMajorVersion", args)
  }

}

  // proto: static void QSurfaceFormat::setDefaultFormat(const QSurfaceFormat & format);
func (this *QSurfaceFormat) setDefaultFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDefaultFormat", args)
  }

}

  // proto:  int QSurfaceFormat::greenBufferSize();
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
    C._ZNK14QSurfaceFormat15greenBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "greenBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::minorVersion();
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
    C._ZNK14QSurfaceFormat12minorVersionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "minorVersion", args)
  }

}

  // proto:  void QSurfaceFormat::setStencilBufferSize(int size);
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
    C._ZN14QSurfaceFormat20setStencilBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStencilBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::swapInterval();
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
    C._ZNK14QSurfaceFormat12swapIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "swapInterval", args)
  }

}

  // proto:  void QSurfaceFormat::setVersion(int major, int minor);
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
    C._ZN14QSurfaceFormat10setVersionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setVersion", args)
  }

}

  // proto:  bool QSurfaceFormat::hasAlpha();
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
    C._ZNK14QSurfaceFormat8hasAlphaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "hasAlpha", args)
  }

}

  // proto:  void QSurfaceFormat::QSurfaceFormat(const QSurfaceFormat & other);
func NewQSurfaceFormat(args ...interface{}) QSurfaceFormat {
  return QSurfaceFormat{}
}

  // proto:  QPair<int, int> QSurfaceFormat::version();
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
    C._ZNK14QSurfaceFormat7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "version", args)
  }

}

  // proto:  int QSurfaceFormat::blueBufferSize();
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
    C._ZNK14QSurfaceFormat14blueBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "blueBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::redBufferSize();
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
    C._ZNK14QSurfaceFormat13redBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "redBufferSize", args)
  }

}

  // proto:  void QSurfaceFormat::~QSurfaceFormat();
func (this *QSurfaceFormat) FreeQSurfaceFormat(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "~QSurfaceFormat", args)
  }

}

  // proto:  void QSurfaceFormat::setGreenBufferSize(int size);
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
    C._ZN14QSurfaceFormat18setGreenBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setGreenBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::samples();
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
    C._ZNK14QSurfaceFormat7samplesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "samples", args)
  }

}

  // proto:  int QSurfaceFormat::depthBufferSize();
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
    C._ZNK14QSurfaceFormat15depthBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "depthBufferSize", args)
  }

}

  // proto:  void QSurfaceFormat::setBlueBufferSize(int size);
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
    C._ZN14QSurfaceFormat17setBlueBufferSizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setBlueBufferSize", args)
  }

}

  // proto:  int QSurfaceFormat::alphaBufferSize();
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
    C._ZNK14QSurfaceFormat15alphaBufferSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "alphaBufferSize", args)
  }

}

  // proto:  bool QSurfaceFormat::stereo();
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
    C.demth_ZNK14QSurfaceFormat6stereoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stereo", args)
  }

}

  // proto:  void QSurfaceFormat::setSwapInterval(int interval);
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
    C._ZN14QSurfaceFormat15setSwapIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSwapInterval", args)
  }

}

  // proto:  void QSurfaceFormat::setStereo(bool enable);
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
    C._ZN14QSurfaceFormat9setStereoEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStereo", args)
  }

}

// <= body block end

