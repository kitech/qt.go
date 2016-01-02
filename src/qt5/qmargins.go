package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QMargins QMarginsF::toMargins();
extern void _ZNK9QMarginsF9toMarginsEv(void* qthis);
  // proto:  qreal QMarginsF::right();
extern void _ZNK9QMarginsF5rightEv(void* qthis);
  // proto:  bool QMarginsF::isNull();
extern void _ZNK9QMarginsF6isNullEv(void* qthis);
  // proto:  void QMarginsF::setRight(qreal right);
extern void _ZN9QMarginsF8setRightEd(void* qthis, double arg0);
  // proto:  void QMarginsF::setTop(qreal top);
extern void _ZN9QMarginsF6setTopEd(void* qthis, double arg0);
  // proto:  qreal QMarginsF::left();
extern void _ZNK9QMarginsF4leftEv(void* qthis);
  // proto:  void QMarginsF::QMarginsF();
extern void* dector_ZN9QMarginsFC1Ev();
extern void _ZN9QMarginsFC1Ev(void* qthis);
  // proto:  void QMarginsF::QMarginsF(qreal left, qreal top, qreal right, qreal bottom);
extern void* dector_ZN9QMarginsFC1Edddd(double arg0, double arg1, double arg2, double arg3);
extern void _ZN9QMarginsFC1Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  qreal QMarginsF::bottom();
extern void _ZNK9QMarginsF6bottomEv(void* qthis);
  // proto:  void QMarginsF::QMarginsF(const QMargins & margins);
extern void* dector_ZN9QMarginsFC1ERK8QMargins(void* arg0);
extern void _ZN9QMarginsFC1ERK8QMargins(void* qthis, void* arg0);
  // proto:  void QMarginsF::setBottom(qreal bottom);
extern void _ZN9QMarginsF9setBottomEd(void* qthis, double arg0);
  // proto:  qreal QMarginsF::top();
extern void _ZNK9QMarginsF3topEv(void* qthis);
  // proto:  void QMarginsF::setLeft(qreal left);
extern void _ZN9QMarginsF7setLeftEd(void* qthis, double arg0);
  // proto:  void QMargins::setLeft(int left);
extern void _ZN8QMargins7setLeftEi(void* qthis, int arg0);
  // proto:  void QMargins::setRight(int right);
extern void _ZN8QMargins8setRightEi(void* qthis, int arg0);
  // proto:  int QMargins::left();
extern void _ZNK8QMargins4leftEv(void* qthis);
  // proto:  int QMargins::top();
extern void _ZNK8QMargins3topEv(void* qthis);
  // proto:  void QMargins::setTop(int top);
extern void _ZN8QMargins6setTopEi(void* qthis, int arg0);
  // proto:  void QMargins::setBottom(int bottom);
extern void _ZN8QMargins9setBottomEi(void* qthis, int arg0);
  // proto:  int QMargins::right();
extern void _ZNK8QMargins5rightEv(void* qthis);
  // proto:  int QMargins::bottom();
extern void _ZNK8QMargins6bottomEv(void* qthis);
  // proto:  bool QMargins::isNull();
extern void _ZNK8QMargins6isNullEv(void* qthis);
  // proto:  void QMargins::QMargins();
extern void* dector_ZN8QMarginsC1Ev();
extern void _ZN8QMarginsC1Ev(void* qthis);
  // proto:  void QMargins::QMargins(int left, int top, int right, int bottom);
extern void* dector_ZN8QMarginsC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN8QMarginsC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMargins)=16
type QMargins struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QMargins QMarginsF::toMargins();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "toMargins", args)
  }

}

  // proto:  qreal QMarginsF::right();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "right", args)
  }

}

  // proto:  bool QMarginsF::isNull();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "isNull", args)
  }

}

  // proto:  void QMarginsF::setRight(qreal right);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMarginsF", "setRight", args)
  }

}

  // proto:  void QMarginsF::setTop(qreal top);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMarginsF", "setTop", args)
  }

}

  // proto:  qreal QMarginsF::left();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "left", args)
  }

}

  // proto:  void QMarginsF::QMarginsF();
func NewQMarginsF(args ...interface{}) QMarginsF {
  return QMarginsF{}
}

  // proto:  qreal QMarginsF::bottom();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "bottom", args)
  }

}

  // proto:  void QMarginsF::setBottom(qreal bottom);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMarginsF", "setBottom", args)
  }

}

  // proto:  qreal QMarginsF::top();
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
  default:
    qtrt.ErrorResolve("QMarginsF", "top", args)
  }

}

  // proto:  void QMarginsF::setLeft(qreal left);
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMarginsF", "setLeft", args)
  }

}

  // proto:  void QMargins::setLeft(int left);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMargins", "setLeft", args)
  }

}

  // proto:  void QMargins::setRight(int right);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMargins", "setRight", args)
  }

}

  // proto:  int QMargins::left();
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
  default:
    qtrt.ErrorResolve("QMargins", "left", args)
  }

}

  // proto:  int QMargins::top();
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
  default:
    qtrt.ErrorResolve("QMargins", "top", args)
  }

}

  // proto:  void QMargins::setTop(int top);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMargins", "setTop", args)
  }

}

  // proto:  void QMargins::setBottom(int bottom);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMargins", "setBottom", args)
  }

}

  // proto:  int QMargins::right();
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
  default:
    qtrt.ErrorResolve("QMargins", "right", args)
  }

}

  // proto:  int QMargins::bottom();
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
  default:
    qtrt.ErrorResolve("QMargins", "bottom", args)
  }

}

  // proto:  bool QMargins::isNull();
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
  default:
    qtrt.ErrorResolve("QMargins", "isNull", args)
  }

}

  // proto:  void QMargins::QMargins();
func NewQMargins(args ...interface{}) QMargins {
  return QMargins{}
}

// <= body block end

