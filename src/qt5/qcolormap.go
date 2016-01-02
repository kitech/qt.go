package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  uint QColormap::pixel(const QColor & color);
extern void _ZNK9QColormap5pixelERK6QColor(void* qthis, void* arg0);
  // proto:  const QVector<QColor> QColormap::colormap();
extern void _ZNK9QColormap8colormapEv(void* qthis);
  // proto:  const QColor QColormap::colorAt(uint pixel);
extern void _ZNK9QColormap7colorAtEj(void* qthis, unsigned int arg0);
  // proto:  void QColormap::~QColormap();
extern void _ZN9QColormapD0Ev(void* qthis);
  // proto:  void QColormap::QColormap();
extern void* dector_ZN9QColormapC1Ev();
extern void _ZN9QColormapC1Ev(void* qthis);
  // proto: static QColormap QColormap::instance(int screen);
extern void _ZN9QColormap8instanceEi(int arg0);
  // proto:  int QColormap::size();
extern void _ZNK9QColormap4sizeEv(void* qthis);
  // proto:  void QColormap::QColormap(const QColormap & colormap);
extern void* dector_ZN9QColormapC1ERKS_(void* arg0);
extern void _ZN9QColormapC1ERKS_(void* qthis, void* arg0);
  // proto: static void QColormap::initialize();
extern void _ZN9QColormap10initializeEv();
  // proto:  int QColormap::depth();
extern void _ZNK9QColormap5depthEv(void* qthis);
  // proto: static void QColormap::cleanup();
extern void _ZN9QColormap7cleanupEv();
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  uint QColormap::pixel(const QColor & color);
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QColormap", "pixel", args)
  }

}

  // proto:  const QVector<QColor> QColormap::colormap();
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
  default:
    qtrt.ErrorResolve("QColormap", "colormap", args)
  }

}

  // proto:  const QColor QColormap::colorAt(uint pixel);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QColormap", "colorAt", args)
  }

}

  // proto:  void QColormap::~QColormap();
func (this *QColormap) FreeQColormap(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColormap", "~QColormap", args)
  }

}

  // proto:  void QColormap::QColormap();
func NewQColormap(args ...interface{}) QColormap {
  return QColormap{}
}

  // proto: static QColormap QColormap::instance(int screen);
func (this *QColormap) instance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColormap", "instance", args)
  }

}

  // proto:  int QColormap::size();
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
  default:
    qtrt.ErrorResolve("QColormap", "size", args)
  }

}

  // proto: static void QColormap::initialize();
func (this *QColormap) initialize_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColormap", "initialize", args)
  }

}

  // proto:  int QColormap::depth();
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
  default:
    qtrt.ErrorResolve("QColormap", "depth", args)
  }

}

  // proto: static void QColormap::cleanup();
func (this *QColormap) cleanup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColormap", "cleanup", args)
  }

}

// <= body block end

