package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qcolormap.h
// dst-file: /src/widgets/qcolormap.go
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
  // proto:  const QColor QColormap::colorAt(uint pixel);
extern void _ZNK9QColormap7colorAtEj(void* qthis, int32_t arg0); // 4
  // proto:  const QVector<QColor> QColormap::colormap();
extern void _ZNK9QColormap8colormapEv(void* qthis); // 4
  // proto:  void QColormap::QColormap(const QColormap & colormap);
extern void _ZN9QColormapC2ERKS_(void* qthis, void* arg0); // 3
  // proto: static QColormap QColormap::instance(int screen);
extern void _ZN9QColormap8instanceEi(int32_t arg0); // 4
  // proto:  int QColormap::depth();
extern void _ZNK9QColormap5depthEv(void* qthis); // 4
  // proto: static void QColormap::cleanup();
extern void _ZN9QColormap7cleanupEv(); // 4
  // proto:  void QColormap::~QColormap();
extern void _ZN9QColormapD2Ev(void* qthis); // 4
  // proto:  QColormap::Mode QColormap::mode();
extern void _ZNK9QColormap4modeEv(void* qthis); // 4
  // proto: static void QColormap::initialize();
extern void _ZN9QColormap10initializeEv(); // 4
  // proto:  uint QColormap::pixel(const QColor & color);
extern void _ZNK9QColormap5pixelERK6QColor(void* qthis, void* arg0); // 4
  // proto:  int QColormap::size();
extern void _ZNK9QColormap4sizeEv(void* qthis); // 4
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

// class sizeof(QColormap)=8
type QColormap struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// colorAt(uint)
func (this *QColormap) colorAt(args ...interface{}) () {
  // colorAt(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap7colorAtEj
    // invoke: const QColor colorAt(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QColormap7colorAtEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColormap", "colorAt", args)
  }

}

// colormap()
func (this *QColormap) colormap(args ...interface{}) () {
  // colormap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap8colormapEv
    // invoke: const QVector<QColor> colormap()
    C._ZNK9QColormap8colormapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "colormap", args)
  }

}

// QColormap(const class QColormap &)
func NewQColormap(args ...interface{}) QColormap {
  // QColormap(const class QColormap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColormap{}) // "const QColormap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QColormapC1ERKS_
    // invoke: void QColormap(const class QColormap &)
    var arg0 = args[0].(QColormap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QColormapC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QColormap", "QColormap", args)
  }

  return QColormap{}
}

// instance(int)
func (this *QColormap) instance_s(args ...interface{}) () {
  // instance(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QColormap8instanceEi
    // invoke: QColormap instance(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QColormap8instanceEi(arg0)
  default:
    qtrt.ErrorResolve("QColormap", "instance", args)
  }

}

// depth()
func (this *QColormap) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap5depthEv
    // invoke: int depth()
    C._ZNK9QColormap5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "depth", args)
  }

}

// cleanup()
func (this *QColormap) cleanup_s(args ...interface{}) () {
  // cleanup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QColormap7cleanupEv
    // invoke: void cleanup()
    C._ZN9QColormap7cleanupEv()
  default:
    qtrt.ErrorResolve("QColormap", "cleanup", args)
  }

}

// ~QColormap()
func (this *QColormap) FreeQColormap(args ...interface{}) () {
  // ~QColormap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QColormapD0Ev
    // invoke: void ~QColormap()
    C._ZN9QColormapD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "~QColormap", args)
  }

}

// mode()
func (this *QColormap) mode(args ...interface{}) () {
  // mode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap4modeEv
    // invoke: QColormap::Mode mode()
    C._ZNK9QColormap4modeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "mode", args)
  }

}

// initialize()
func (this *QColormap) initialize_s(args ...interface{}) () {
  // initialize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QColormap10initializeEv
    // invoke: void initialize()
    C._ZN9QColormap10initializeEv()
  default:
    qtrt.ErrorResolve("QColormap", "initialize", args)
  }

}

// pixel(const class QColor &)
func (this *QColormap) pixel(args ...interface{}) () {
  // pixel(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap5pixelERK6QColor
    // invoke: uint pixel(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QColormap5pixelERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColormap", "pixel", args)
  }

}

// size()
func (this *QColormap) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QColormap4sizeEv
    // invoke: int size()
    C._ZNK9QColormap4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "size", args)
  }

}

// <= body block end

