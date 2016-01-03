package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QRubberBand::resize(const QSize & s);
extern void demth_ZN11QRubberBand6resizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QRubberBand::setGeometry(int x, int y, int w, int h);
extern void demth_ZN11QRubberBand11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QRubberBand::move(const QPoint & p);
extern void demth_ZN11QRubberBand4moveERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRubberBand::~QRubberBand();
extern void _ZN11QRubberBandD0Ev(void* qthis);
  // proto:  void QRubberBand::move(int x, int y);
extern void demth_ZN11QRubberBand4moveEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  const QMetaObject * QRubberBand::metaObject();
extern void _ZNK11QRubberBand10metaObjectEv(void* qthis);
  // proto:  void QRubberBand::setGeometry(const QRect & r);
extern void _ZN11QRubberBand11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QRubberBand::resize(int w, int h);
extern void demth_ZN11QRubberBand6resizeEii(void* qthis, int32_t arg0, int32_t arg1);
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

  // proto:  void QRubberBand::resize(const QSize & s);
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
    C.demth_ZN11QRubberBand6resizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QRubberBand6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN11QRubberBand6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QRubberBand", "resize", args)
  }

}

  // proto:  void QRubberBand::setGeometry(int x, int y, int w, int h);
func (this *QRubberBand) setGeometry(args ...interface{}) () {
  // setGeometry(int, int, int, int)
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
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
    C.demth_ZN11QRubberBand11setGeometryEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN11QRubberBand11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QRubberBand11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRubberBand", "setGeometry", args)
  }

}

  // proto:  void QRubberBand::move(const QPoint & p);
func (this *QRubberBand) move_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRubberBand", "move", args)
  }

}

  // proto:  void QRubberBand::~QRubberBand();
func (this *QRubberBand) FreeQRubberBand(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRubberBand", "~QRubberBand", args)
  }

}

  // proto:  const QMetaObject * QRubberBand::metaObject();
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

// <= body block end

