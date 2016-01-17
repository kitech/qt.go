package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qrubberband.h
// dst-file: /src/widgets/qrubberband.go
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
  // proto:  void QRubberBand::move(int x, int y);
extern void _ZN11QRubberBand4moveEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QRubberBand::move(const QPoint & p);
extern void _ZN11QRubberBand4moveERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QRubberBand::~QRubberBand();
extern void _ZN11QRubberBandD2Ev(void* qthis); // 4
  // proto:  void QRubberBand::setGeometry(const QRect & r);
extern void _ZN11QRubberBand11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QRubberBand::setGeometry(int x, int y, int w, int h);
extern void _ZN11QRubberBand11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  QRubberBand::Shape QRubberBand::shape();
extern void _ZNK11QRubberBand5shapeEv(void* qthis); // 4
  // proto:  const QMetaObject * QRubberBand::metaObject();
extern void _ZNK11QRubberBand10metaObjectEv(void* qthis); // 4
  // proto:  void QRubberBand::resize(const QSize & s);
extern void _ZN11QRubberBand6resizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QRubberBand::resize(int w, int h);
extern void _ZN11QRubberBand6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 2
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

// class sizeof(QRubberBand)=1
type QRubberBand struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
}

// move(int, int)
func (this *QRubberBand) move_(args ...interface{}) () {
  // move(int, int)
  // move(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QRubberBand4moveEii
    // invoke: void move(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QRubberBand4moveEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QRubberBand4moveERK6QPoint
    // invoke: void move(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QRubberBand4moveERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRubberBand", "move", args)
  }

}

// ~QRubberBand()
func (this *QRubberBand) FreeQRubberBand(args ...interface{}) () {
  // ~QRubberBand()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QRubberBandD0Ev
    // invoke: void ~QRubberBand()
    C._ZN11QRubberBandD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRubberBand", "~QRubberBand", args)
  }

}

// setGeometry(const class QRect &)
func (this *QRubberBand) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  // setGeometry(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QRubberBand11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QRubberBand11setGeometryERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QRubberBand11setGeometryEiiii
    // invoke: void setGeometry(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN11QRubberBand11setGeometryEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QRubberBand", "setGeometry", args)
  }

}

// shape()
func (this *QRubberBand) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QRubberBand5shapeEv
    // invoke: QRubberBand::Shape shape()
    C._ZNK11QRubberBand5shapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRubberBand", "shape", args)
  }

}

// metaObject()
func (this *QRubberBand) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QRubberBand10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QRubberBand10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRubberBand", "metaObject", args)
  }

}

// resize(const class QSize &)
func (this *QRubberBand) resize(args ...interface{}) () {
  // resize(const class QSize &)
  // resize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QRubberBand6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QRubberBand6resizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QRubberBand6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QRubberBand6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRubberBand", "resize", args)
  }

}

// <= body block end

