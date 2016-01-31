package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QColor QPen::color();
extern void C_ZNK4QPen5colorEv(void* qthis); // 4
  // proto:  QBrush QPen::brush();
extern void C_ZNK4QPen5brushEv(void* qthis); // 4
  // proto:  void QPen::setWidth(int width);
extern void C_ZN4QPen8setWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::PenStyle QPen::style();
extern void C_ZNK4QPen5styleEv(void* qthis); // 4
  // proto:  bool QPen::isCosmetic();
extern void C_ZNK4QPen10isCosmeticEv(void* qthis); // 4
  // proto:  void QPen::~QPen();
extern void C_ZN4QPenD2Ev(void* qthis); // 4
  // proto:  Qt::PenJoinStyle QPen::joinStyle();
extern void C_ZNK4QPen9joinStyleEv(void* qthis); // 4
  // proto:  int QPen::width();
extern void C_ZNK4QPen5widthEv(void* qthis); // 4
  // proto:  void QPen::swap(QPen & other);
extern void C_ZN4QPen4swapERS_(void* qthis, void* arg0); // 2
  // proto:  qreal QPen::dashOffset();
extern void C_ZNK4QPen10dashOffsetEv(void* qthis); // 4
  // proto:  bool QPen::isDetached();
extern void C_ZN4QPen10isDetachedEv(void* qthis); // 4
  // proto:  qreal QPen::miterLimit();
extern void C_ZNK4QPen10miterLimitEv(void* qthis); // 4
  // proto:  bool QPen::isSolid();
extern void C_ZNK4QPen7isSolidEv(void* qthis); // 4
  // proto:  void QPen::QPen(const QPen & pen);
extern void* C_ZN4QPenC2ERKS_(void* arg0); // 3
  // proto:  void QPen::QPen(const QColor & color);
extern void* C_ZN4QPenC2ERK6QColor(void* arg0); // 3
  // proto:  void QPen::QPen();
extern void* C_ZN4QPenC2Ev(); // 3
  // proto:  qreal QPen::widthF();
extern void C_ZNK4QPen6widthFEv(void* qthis); // 4
  // proto:  QVector<qreal> QPen::dashPattern();
extern void C_ZNK4QPen11dashPatternEv(void* qthis); // 4
  // proto:  void QPen::setColor(const QColor & color);
extern void C_ZN4QPen8setColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QPen::setMiterLimit(qreal limit);
extern void C_ZN4QPen13setMiterLimitEd(void* qthis, double arg0); // 4
  // proto:  void QPen::setCosmetic(bool cosmetic);
extern void C_ZN4QPen11setCosmeticEb(void* qthis, bool arg0); // 4
  // proto:  void QPen::setWidthF(qreal width);
extern void C_ZN4QPen9setWidthFEd(void* qthis, double arg0); // 4
  // proto:  Qt::PenCapStyle QPen::capStyle();
extern void C_ZNK4QPen8capStyleEv(void* qthis); // 4
  // proto:  void QPen::setBrush(const QBrush & brush);
extern void C_ZN4QPen8setBrushERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QPen::setDashOffset(qreal doffset);
extern void C_ZN4QPen13setDashOffsetEd(void* qthis, double arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// color()
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
    // invoke: QColor color()
    var ret = C.C_ZNK4QPen5colorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "color", args)
  }

}

// brush()
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
    // invoke: QBrush brush()
    var ret = C.C_ZNK4QPen5brushEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "brush", args)
  }

}

// setWidth(int)
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
    // invoke: void setWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setWidth", args)
  }

}

// style()
func (this *QPen) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen5styleEv
    // invoke: Qt::PenStyle style()
    C.C_ZNK4QPen5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "style", args)
  }

}

// isCosmetic()
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
    // invoke: bool isCosmetic()
    var ret = C.C_ZNK4QPen10isCosmeticEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "isCosmetic", args)
  }

}

// ~QPen()
func (this *QPen) FreeQPen(args ...interface{}) () {
  // ~QPen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPenD0Ev
    // invoke: void ~QPen()
    C.C_ZN4QPenD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "~QPen", args)
  }

}

// joinStyle()
func (this *QPen) joinStyle(args ...interface{}) () {
  // joinStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen9joinStyleEv
    // invoke: Qt::PenJoinStyle joinStyle()
    C.C_ZNK4QPen9joinStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "joinStyle", args)
  }

}

// width()
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
    // invoke: int width()
    var ret = C.C_ZNK4QPen5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "width", args)
  }

}

// swap(class QPen &)
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
    // invoke: void swap(class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "swap", args)
  }

}

// dashOffset()
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
    // invoke: qreal dashOffset()
    var ret = C.C_ZNK4QPen10dashOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "dashOffset", args)
  }

}

// isDetached()
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
    // invoke: bool isDetached()
    var ret = C.C_ZN4QPen10isDetachedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "isDetached", args)
  }

}

// miterLimit()
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
    // invoke: qreal miterLimit()
    var ret = C.C_ZNK4QPen10miterLimitEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "miterLimit", args)
  }

}

// isSolid()
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
    // invoke: bool isSolid()
    var ret = C.C_ZNK4QPen7isSolidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "isSolid", args)
  }

}

// QPen(const class QPen &)
func NewQPen(args ...interface{}) *QPen {
  // QPen(const class QPen &)
  // QPen(const class QColor &)
  // QPen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QPenC1ERKS_
    // invoke: void QPen(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2ERKS_(arg0)
    return &QPen{qclsinst:qthis}
  case 1:
    // invoke: _ZN4QPenC1ERK6QColor
    // invoke: void QPen(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2ERK6QColor(arg0)
    return &QPen{qclsinst:qthis}
  case 2:
    // invoke: _ZN4QPenC1Ev
    // invoke: void QPen()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2Ev()
    return &QPen{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPen", "QPen", args)
  }

  return nil // QPen{qclsinst:qthis}
}

// widthF()
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
    // invoke: qreal widthF()
    var ret = C.C_ZNK4QPen6widthFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QPen", "widthF", args)
  }

}

// dashPattern()
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
    // invoke: QVector<qreal> dashPattern()
    C.C_ZNK4QPen11dashPatternEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "dashPattern", args)
  }

}

// setColor(const class QColor &)
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
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setColor", args)
  }

}

// setMiterLimit(qreal)
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
    // invoke: void setMiterLimit(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen13setMiterLimitEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setMiterLimit", args)
  }

}

// setCosmetic(_Bool)
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
    // invoke: void setCosmetic(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen11setCosmeticEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setCosmetic", args)
  }

}

// setWidthF(qreal)
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
    // invoke: void setWidthF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen9setWidthFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setWidthF", args)
  }

}

// capStyle()
func (this *QPen) capStyle(args ...interface{}) () {
  // capStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QPen8capStyleEv
    // invoke: Qt::PenCapStyle capStyle()
    C.C_ZNK4QPen8capStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "capStyle", args)
  }

}

// setBrush(const class QBrush &)
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
    // invoke: void setBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setBrush", args)
  }

}

// setDashOffset(qreal)
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
    // invoke: void setDashOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen13setDashOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setDashOffset", args)
  }

}

// <= body block end

