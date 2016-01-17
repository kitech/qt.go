package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qmargins.h
// dst-file: /src/core/qmargins.go
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
  // proto:  qreal QMarginsF::right();
extern void _ZNK9QMarginsF5rightEv(void* qthis); // 2
  // proto:  qreal QMarginsF::bottom();
extern void _ZNK9QMarginsF6bottomEv(void* qthis); // 2
  // proto:  void QMarginsF::setBottom(qreal bottom);
extern void _ZN9QMarginsF9setBottomEd(void* qthis, double arg0); // 2
  // proto:  qreal QMarginsF::top();
extern void _ZNK9QMarginsF3topEv(void* qthis); // 2
  // proto:  QMargins QMarginsF::toMargins();
extern void _ZNK9QMarginsF9toMarginsEv(void* qthis); // 2
  // proto:  void QMarginsF::setRight(qreal right);
extern void _ZN9QMarginsF8setRightEd(void* qthis, double arg0); // 2
  // proto:  void QMarginsF::setTop(qreal top);
extern void _ZN9QMarginsF6setTopEd(void* qthis, double arg0); // 2
  // proto:  bool QMarginsF::isNull();
extern void _ZNK9QMarginsF6isNullEv(void* qthis); // 2
  // proto:  void QMarginsF::setLeft(qreal left);
extern void _ZN9QMarginsF7setLeftEd(void* qthis, double arg0); // 2
  // proto:  void QMarginsF::QMarginsF();
extern void _ZN9QMarginsFC2Ev(void* qthis); // 1
  // proto:  void QMarginsF::QMarginsF(qreal left, qreal top, qreal right, qreal bottom);
extern void _ZN9QMarginsFC2Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QMarginsF::QMarginsF(const QMargins & margins);
extern void _ZN9QMarginsFC2ERK8QMargins(void* qthis, void* arg0); // 1
  // proto:  qreal QMarginsF::left();
extern void _ZNK9QMarginsF4leftEv(void* qthis); // 2
  // proto:  int QMargins::right();
extern void _ZNK8QMargins5rightEv(void* qthis); // 2
  // proto:  int QMargins::bottom();
extern void _ZNK8QMargins6bottomEv(void* qthis); // 2
  // proto:  void QMargins::setBottom(int bottom);
extern void _ZN8QMargins9setBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  void QMargins::setRight(int right);
extern void _ZN8QMargins8setRightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QMargins::setTop(int top);
extern void _ZN8QMargins6setTopEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QMargins::isNull();
extern void _ZNK8QMargins6isNullEv(void* qthis); // 2
  // proto:  void QMargins::QMargins();
extern void _ZN8QMarginsC2Ev(void* qthis); // 1
  // proto:  void QMargins::QMargins(int left, int top, int right, int bottom);
extern void _ZN8QMarginsC2Eiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QMargins::setLeft(int left);
extern void _ZN8QMargins7setLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  int QMargins::top();
extern void _ZNK8QMargins3topEv(void* qthis); // 2
  // proto:  int QMargins::left();
extern void _ZNK8QMargins4leftEv(void* qthis); // 2
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

// class sizeof(QMarginsF)=32
type QMarginsF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMargins)=16
type QMargins struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// right()
func (this *QMarginsF) right(args ...interface{}) () {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF5rightEv
    // invoke: qreal right()
    C._ZNK9QMarginsF5rightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "right", args)
  }

}

// bottom()
func (this *QMarginsF) bottom(args ...interface{}) () {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF6bottomEv
    // invoke: qreal bottom()
    C._ZNK9QMarginsF6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "bottom", args)
  }

}

// setBottom(qreal)
func (this *QMarginsF) setBottom(args ...interface{}) () {
  // setBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF9setBottomEd
    // invoke: void setBottom(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN9QMarginsF9setBottomEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setBottom", args)
  }

}

// top()
func (this *QMarginsF) top(args ...interface{}) () {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF3topEv
    // invoke: qreal top()
    C._ZNK9QMarginsF3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "top", args)
  }

}

// toMargins()
func (this *QMarginsF) toMargins(args ...interface{}) () {
  // toMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF9toMarginsEv
    // invoke: QMargins toMargins()
    C._ZNK9QMarginsF9toMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "toMargins", args)
  }

}

// setRight(qreal)
func (this *QMarginsF) setRight(args ...interface{}) () {
  // setRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF8setRightEd
    // invoke: void setRight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN9QMarginsF8setRightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setRight", args)
  }

}

// setTop(qreal)
func (this *QMarginsF) setTop(args ...interface{}) () {
  // setTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF6setTopEd
    // invoke: void setTop(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN9QMarginsF6setTopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setTop", args)
  }

}

// isNull()
func (this *QMarginsF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF6isNullEv
    // invoke: bool isNull()
    C._ZNK9QMarginsF6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "isNull", args)
  }

}

// setLeft(qreal)
func (this *QMarginsF) setLeft(args ...interface{}) () {
  // setLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF7setLeftEd
    // invoke: void setLeft(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN9QMarginsF7setLeftEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setLeft", args)
  }

}

// QMarginsF()
func NewQMarginsF(args ...interface{}) QMarginsF {
  // QMarginsF()
  // QMarginsF(qreal, qreal, qreal, qreal)
  // QMarginsF(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsFC1Ev
    // invoke: void QMarginsF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QMarginsFC2Ev(qthis)
  case 1:
    // invoke: _ZN9QMarginsFC1Edddd
    // invoke: void QMarginsF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QMarginsFC2Edddd(qthis, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN9QMarginsFC1ERK8QMargins
    // invoke: void QMarginsF(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QMarginsFC2ERK8QMargins(qthis, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "QMarginsF", args)
  }

  return QMarginsF{}
}

// left()
func (this *QMarginsF) left(args ...interface{}) () {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF4leftEv
    // invoke: qreal left()
    C._ZNK9QMarginsF4leftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMarginsF", "left", args)
  }

}

// right()
func (this *QMargins) right(args ...interface{}) () {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins5rightEv
    // invoke: int right()
    C._ZNK8QMargins5rightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMargins", "right", args)
  }

}

// bottom()
func (this *QMargins) bottom(args ...interface{}) () {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins6bottomEv
    // invoke: int bottom()
    C._ZNK8QMargins6bottomEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMargins", "bottom", args)
  }

}

// setBottom(int)
func (this *QMargins) setBottom(args ...interface{}) () {
  // setBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins9setBottomEi
    // invoke: void setBottom(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QMargins9setBottomEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setBottom", args)
  }

}

// setRight(int)
func (this *QMargins) setRight(args ...interface{}) () {
  // setRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins8setRightEi
    // invoke: void setRight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QMargins8setRightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setRight", args)
  }

}

// setTop(int)
func (this *QMargins) setTop(args ...interface{}) () {
  // setTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins6setTopEi
    // invoke: void setTop(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QMargins6setTopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setTop", args)
  }

}

// isNull()
func (this *QMargins) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins6isNullEv
    // invoke: bool isNull()
    C._ZNK8QMargins6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMargins", "isNull", args)
  }

}

// QMargins()
func NewQMargins(args ...interface{}) QMargins {
  // QMargins()
  // QMargins(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMarginsC1Ev
    // invoke: void QMargins()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QMarginsC2Ev(qthis)
  case 1:
    // invoke: _ZN8QMarginsC1Eiiii
    // invoke: void QMargins(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QMarginsC2Eiiii(qthis, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QMargins", "QMargins", args)
  }

  return QMargins{}
}

// setLeft(int)
func (this *QMargins) setLeft(args ...interface{}) () {
  // setLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins7setLeftEi
    // invoke: void setLeft(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QMargins7setLeftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setLeft", args)
  }

}

// top()
func (this *QMargins) top(args ...interface{}) () {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins3topEv
    // invoke: int top()
    C._ZNK8QMargins3topEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMargins", "top", args)
  }

}

// left()
func (this *QMargins) left(args ...interface{}) () {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins4leftEv
    // invoke: int left()
    C._ZNK8QMargins4leftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMargins", "left", args)
  }

}

// <= body block end

