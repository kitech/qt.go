package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK9QColormap7colorAtEj(void* qthis, int32_t arg0); // 4
  // proto:  const QVector<QColor> QColormap::colormap();
extern void C_ZNK9QColormap8colormapEv(void* qthis); // 4
  // proto:  void QColormap::QColormap(const QColormap & colormap);
extern void* C_ZN9QColormapC2ERKS_(void* arg0); // 3
  // proto: static QColormap QColormap::instance(int screen);
extern void* C_ZN9QColormap8instanceEi(int32_t arg0); // 4
  // proto:  int QColormap::depth();
extern int32_t C_ZNK9QColormap5depthEv(void* qthis); // 4
  // proto: static void QColormap::cleanup();
extern void C_ZN9QColormap7cleanupEv(); // 4
  // proto:  void QColormap::~QColormap();
extern void C_ZN9QColormapD2Ev(void* qthis); // 4
  // proto:  QColormap::Mode QColormap::mode();
extern void C_ZNK9QColormap4modeEv(void* qthis); // 4
  // proto: static void QColormap::initialize();
extern void C_ZN9QColormap10initializeEv(); // 4
  // proto:  uint QColormap::pixel(const QColor & color);
extern int32_t C_ZNK9QColormap5pixelERK6QColor(void* qthis, void* arg0); // 4
  // proto:  int QColormap::size();
extern int32_t C_ZNK9QColormap4sizeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// colorAt(uint)
func (this *QColormap) Colorat(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QColormap7colorAtEj(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "const QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColormap", "colorAt", args)
  }

  return
}

// colormap()
func (this *QColormap) Colormap(args ...interface{}) () {
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
    C.C_ZNK9QColormap8colormapEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "colormap", args)
  }

  return
}

// QColormap(const class QColormap &)
func NewQColormap(args ...interface{}) *QColormap {
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
    var arg0 = args[0].(*QColormap).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QColormapC2ERKS_(arg0)
    return &QColormap{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QColormap", "QColormap", args)
  }

  return nil // QColormap{Qclsinst:qthis}
}

// instance(int)
func (this *QColormap) Instance_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QColormap8instanceEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColormap{}) // "QColormap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColormap", "instance", args)
  }

  return
}

// depth()
func (this *QColormap) Depth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QColormap5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColormap", "depth", args)
  }

  return
}

// cleanup()
func (this *QColormap) Cleanup_S(args ...interface{}) () {
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
    C.C_ZN9QColormap7cleanupEv()
  default:
    qtrt.ErrorResolve("QColormap", "cleanup", args)
  }

  return
}

// ~QColormap()
func (this *QColormap) Freeqcolormap(args ...interface{}) () {
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
    C.C_ZN9QColormapD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "~QColormap", args)
  }

  return
}

// mode()
func (this *QColormap) Mode(args ...interface{}) () {
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
    C.C_ZNK9QColormap4modeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColormap", "mode", args)
  }

  return
}

// initialize()
func (this *QColormap) Initialize_S(args ...interface{}) () {
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
    C.C_ZN9QColormap10initializeEv()
  default:
    qtrt.ErrorResolve("QColormap", "initialize", args)
  }

  return
}

// pixel(const class QColor &)
func (this *QColormap) Pixel(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QColormap5pixelERK6QColor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColormap", "pixel", args)
  }

  return
}

// size()
func (this *QColormap) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QColormap4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColormap", "size", args)
  }

  return
}

// <= body block end

