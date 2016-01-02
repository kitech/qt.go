package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qpen.h
// dst-file: /src/gui/qpen.go
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
  // proto:  void QPen::~QPen();
extern void _ZN4QPenD0Ev(void* qthis);
  // proto:  qreal QPen::dashOffset();
extern void _ZNK4QPen10dashOffsetEv(void* qthis);
  // proto:  void QPen::QPen(const QColor & color);
extern void* dector_ZN4QPenC1ERK6QColor(void* arg0);
extern void _ZN4QPenC1ERK6QColor(void* qthis, void* arg0);
  // proto:  qreal QPen::miterLimit();
extern void _ZNK4QPen10miterLimitEv(void* qthis);
  // proto:  void QPen::setWidthF(qreal width);
extern void _ZN4QPen9setWidthFEd(void* qthis, double arg0);
  // proto:  void QPen::setBrush(const QBrush & brush);
extern void _ZN4QPen8setBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  QColor QPen::color();
extern void _ZNK4QPen5colorEv(void* qthis);
  // proto:  void QPen::setWidth(int width);
extern void _ZN4QPen8setWidthEi(void* qthis, int arg0);
  // proto:  qreal QPen::widthF();
extern void _ZNK4QPen6widthFEv(void* qthis);
  // proto:  void QPen::setCosmetic(bool cosmetic);
extern void _ZN4QPen11setCosmeticEb(void* qthis, bool arg0);
  // proto:  bool QPen::isSolid();
extern void _ZNK4QPen7isSolidEv(void* qthis);
  // proto:  void QPen::setColor(const QColor & color);
extern void _ZN4QPen8setColorERK6QColor(void* qthis, void* arg0);
  // proto:  QVector<qreal> QPen::dashPattern();
extern void _ZNK4QPen11dashPatternEv(void* qthis);
  // proto:  bool QPen::isDetached();
extern void _ZN4QPen10isDetachedEv(void* qthis);
  // proto:  void QPen::QPen(const QPen & pen);
extern void* dector_ZN4QPenC1ERKS_(void* arg0);
extern void _ZN4QPenC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPen::setMiterLimit(qreal limit);
extern void _ZN4QPen13setMiterLimitEd(void* qthis, double arg0);
  // proto:  void QPen::QPen();
extern void* dector_ZN4QPenC1Ev();
extern void _ZN4QPenC1Ev(void* qthis);
  // proto:  int QPen::width();
extern void _ZNK4QPen5widthEv(void* qthis);
  // proto:  void QPen::swap(QPen & other);
extern void _ZN4QPen4swapERS_(void* qthis, void* arg0);
  // proto:  QBrush QPen::brush();
extern void _ZNK4QPen5brushEv(void* qthis);
  // proto:  bool QPen::isCosmetic();
extern void _ZNK4QPen10isCosmeticEv(void* qthis);
  // proto:  void QPen::setDashOffset(qreal doffset);
extern void _ZN4QPen13setDashOffsetEd(void* qthis, double arg0);
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

// class sizeof(QPen)=8
type QPen struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPen::~QPen();
func (this *QPen) FreeQPen(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPen", "~QPen", args)
  }

}

  // proto:  qreal QPen::dashOffset();
func (this *QPen) dashOffset(args ...interface{}) () {
  // dashOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen10dashOffsetEv
  default:
    qtrt.ErrorResolve("QPen", "dashOffset", args)
  }

}

  // proto:  void QPen::QPen(const QColor & color);
func NewQPen(args ...interface{}) QPen {
  return QPen{}
}

  // proto:  qreal QPen::miterLimit();
func (this *QPen) miterLimit(args ...interface{}) () {
  // miterLimit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen10miterLimitEv
  default:
    qtrt.ErrorResolve("QPen", "miterLimit", args)
  }

}

  // proto:  void QPen::setWidthF(qreal width);
func (this *QPen) setWidthF(args ...interface{}) () {
  // setWidthF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen9setWidthFEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setWidthF", args)
  }

}

  // proto:  void QPen::setBrush(const QBrush & brush);
func (this *QPen) setBrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen8setBrushERK6QBrush
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setBrush", args)
  }

}

  // proto:  QColor QPen::color();
func (this *QPen) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen5colorEv
  default:
    qtrt.ErrorResolve("QPen", "color", args)
  }

}

  // proto:  void QPen::setWidth(int width);
func (this *QPen) setWidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen8setWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setWidth", args)
  }

}

  // proto:  qreal QPen::widthF();
func (this *QPen) widthF(args ...interface{}) () {
  // widthF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen6widthFEv
  default:
    qtrt.ErrorResolve("QPen", "widthF", args)
  }

}

  // proto:  void QPen::setCosmetic(bool cosmetic);
func (this *QPen) setCosmetic(args ...interface{}) () {
  // setCosmetic(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen11setCosmeticEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setCosmetic", args)
  }

}

  // proto:  bool QPen::isSolid();
func (this *QPen) isSolid(args ...interface{}) () {
  // isSolid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen7isSolidEv
  default:
    qtrt.ErrorResolve("QPen", "isSolid", args)
  }

}

  // proto:  void QPen::setColor(const QColor & color);
func (this *QPen) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen8setColorERK6QColor
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setColor", args)
  }

}

  // proto:  QVector<qreal> QPen::dashPattern();
func (this *QPen) dashPattern(args ...interface{}) () {
  // dashPattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen11dashPatternEv
  default:
    qtrt.ErrorResolve("QPen", "dashPattern", args)
  }

}

  // proto:  bool QPen::isDetached();
func (this *QPen) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen10isDetachedEv
  default:
    qtrt.ErrorResolve("QPen", "isDetached", args)
  }

}

  // proto:  void QPen::setMiterLimit(qreal limit);
func (this *QPen) setMiterLimit(args ...interface{}) () {
  // setMiterLimit(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen13setMiterLimitEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setMiterLimit", args)
  }

}

  // proto:  int QPen::width();
func (this *QPen) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen5widthEv
  default:
    qtrt.ErrorResolve("QPen", "width", args)
  }

}

  // proto:  void QPen::swap(QPen & other);
func (this *QPen) swap(args ...interface{}) () {
  // swap(class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen4swapERS_
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "swap", args)
  }

}

  // proto:  QBrush QPen::brush();
func (this *QPen) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen5brushEv
  default:
    qtrt.ErrorResolve("QPen", "brush", args)
  }

}

  // proto:  bool QPen::isCosmetic();
func (this *QPen) isCosmetic(args ...interface{}) () {
  // isCosmetic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen10isCosmeticEv
  default:
    qtrt.ErrorResolve("QPen", "isCosmetic", args)
  }

}

  // proto:  void QPen::setDashOffset(qreal doffset);
func (this *QPen) setDashOffset(args ...interface{}) () {
  // setDashOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPen13setDashOffsetEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPen", "setDashOffset", args)
  }

}

// <= body block end

