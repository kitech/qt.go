package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QColor QPen::color();
extern void* C_ZNK4QPen5colorEv(void* qthis); // 4
  // proto:  QBrush QPen::brush();
extern void* C_ZNK4QPen5brushEv(void* qthis); // 4
  // proto:  void QPen::setWidth(int width);
extern void C_ZN4QPen8setWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::PenStyle QPen::style();
extern void C_ZNK4QPen5styleEv(void* qthis); // 4
  // proto:  bool QPen::isCosmetic();
extern bool C_ZNK4QPen10isCosmeticEv(void* qthis); // 4
  // proto:  void QPen::~QPen();
extern void C_ZN4QPenD2Ev(void* qthis); // 4
  // proto:  Qt::PenJoinStyle QPen::joinStyle();
extern void C_ZNK4QPen9joinStyleEv(void* qthis); // 4
  // proto:  int QPen::width();
extern int32_t C_ZNK4QPen5widthEv(void* qthis); // 4
  // proto:  void QPen::swap(QPen & other);
extern void C_ZN4QPen4swapERS_(void* qthis, void* arg0); // 2
  // proto:  qreal QPen::dashOffset();
extern double C_ZNK4QPen10dashOffsetEv(void* qthis); // 4
  // proto:  bool QPen::isDetached();
extern bool C_ZN4QPen10isDetachedEv(void* qthis); // 4
  // proto:  qreal QPen::miterLimit();
extern double C_ZNK4QPen10miterLimitEv(void* qthis); // 4
  // proto:  bool QPen::isSolid();
extern bool C_ZNK4QPen7isSolidEv(void* qthis); // 4
  // proto:  void QPen::QPen(const QPen & pen);
extern void* C_ZN4QPenC2ERKS_(void* arg0); // 3
  // proto:  void QPen::QPen(const QColor & color);
extern void* C_ZN4QPenC2ERK6QColor(void* arg0); // 3
  // proto:  void QPen::QPen();
extern void* C_ZN4QPenC2Ev(); // 3
  // proto:  qreal QPen::widthF();
extern double C_ZNK4QPen6widthFEv(void* qthis); // 4
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPen)=8
type QPen struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// color()
func (this *QPen) Color(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen5colorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "color", args)
  }

  return
}

// brush()
func (this *QPen) Brush(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen5brushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "brush", args)
  }

  return
}

// setWidth(int)
func (this *QPen) Setwidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setWidth", args)
  }

  return
}

// style()
func (this *QPen) Style(args ...interface{}) () {
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
    C.C_ZNK4QPen5styleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "style", args)
  }

  return
}

// isCosmetic()
func (this *QPen) Iscosmetic(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen10isCosmeticEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "isCosmetic", args)
  }

  return
}

// ~QPen()
func (this *QPen) Freeqpen(args ...interface{}) () {
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
    C.C_ZN4QPenD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "~QPen", args)
  }

  return
}

// joinStyle()
func (this *QPen) Joinstyle(args ...interface{}) () {
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
    C.C_ZNK4QPen9joinStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "joinStyle", args)
  }

  return
}

// width()
func (this *QPen) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "width", args)
  }

  return
}

// swap(class QPen &)
func (this *QPen) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "swap", args)
  }

  return
}

// dashOffset()
func (this *QPen) Dashoffset(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen10dashOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "dashOffset", args)
  }

  return
}

// isDetached()
func (this *QPen) Isdetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN4QPen10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "isDetached", args)
  }

  return
}

// miterLimit()
func (this *QPen) Miterlimit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen10miterLimitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "miterLimit", args)
  }

  return
}

// isSolid()
func (this *QPen) Issolid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen7isSolidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "isSolid", args)
  }

  return
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
    var arg0 = args[0].(*QPen).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2ERKS_(arg0)
    return &QPen{Qclsinst:qthis}
  case 1:
    // invoke: _ZN4QPenC1ERK6QColor
    // invoke: void QPen(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2ERK6QColor(arg0)
    return &QPen{Qclsinst:qthis}
  case 2:
    // invoke: _ZN4QPenC1Ev
    // invoke: void QPen()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QPenC2Ev()
    return &QPen{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPen", "QPen", args)
  }

  return nil // QPen{Qclsinst:qthis}
}

// widthF()
func (this *QPen) Widthf(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK4QPen6widthFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPen", "widthF", args)
  }

  return
}

// dashPattern()
func (this *QPen) Dashpattern(args ...interface{}) () {
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
    C.C_ZNK4QPen11dashPatternEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "dashPattern", args)
  }

  return
}

// setColor(const class QColor &)
func (this *QPen) Setcolor(args ...interface{}) () {
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
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setColor", args)
  }

  return
}

// setMiterLimit(qreal)
func (this *QPen) Setmiterlimit(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen13setMiterLimitEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setMiterLimit", args)
  }

  return
}

// setCosmetic(_Bool)
func (this *QPen) Setcosmetic(args ...interface{}) () {
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
    C.C_ZN4QPen11setCosmeticEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setCosmetic", args)
  }

  return
}

// setWidthF(qreal)
func (this *QPen) Setwidthf(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen9setWidthFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setWidthF", args)
  }

  return
}

// capStyle()
func (this *QPen) Capstyle(args ...interface{}) () {
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
    C.C_ZNK4QPen8capStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPen", "capStyle", args)
  }

  return
}

// setBrush(const class QBrush &)
func (this *QPen) Setbrush(args ...interface{}) () {
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
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen8setBrushERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setBrush", args)
  }

  return
}

// setDashOffset(qreal)
func (this *QPen) Setdashoffset(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN4QPen13setDashOffsetEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPen", "setDashOffset", args)
  }

  return
}

// <= body block end

