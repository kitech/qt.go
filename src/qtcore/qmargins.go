package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern double C_ZNK9QMarginsF5rightEv(void* qthis); // 2
  // proto:  qreal QMarginsF::bottom();
extern double C_ZNK9QMarginsF6bottomEv(void* qthis); // 2
  // proto:  void QMarginsF::setBottom(qreal bottom);
extern void C_ZN9QMarginsF9setBottomEd(void* qthis, double arg0); // 2
  // proto:  qreal QMarginsF::top();
extern double C_ZNK9QMarginsF3topEv(void* qthis); // 2
  // proto:  QMargins QMarginsF::toMargins();
extern void* C_ZNK9QMarginsF9toMarginsEv(void* qthis); // 2
  // proto:  void QMarginsF::setRight(qreal right);
extern void C_ZN9QMarginsF8setRightEd(void* qthis, double arg0); // 2
  // proto:  void QMarginsF::setTop(qreal top);
extern void C_ZN9QMarginsF6setTopEd(void* qthis, double arg0); // 2
  // proto:  bool QMarginsF::isNull();
extern bool C_ZNK9QMarginsF6isNullEv(void* qthis); // 2
  // proto:  void QMarginsF::setLeft(qreal left);
extern void C_ZN9QMarginsF7setLeftEd(void* qthis, double arg0); // 2
  // proto:  void QMarginsF::QMarginsF();
extern void* C_ZN9QMarginsFC2Ev(); // 1
  // proto:  void QMarginsF::QMarginsF(qreal left, qreal top, qreal right, qreal bottom);
extern void* C_ZN9QMarginsFC2Edddd(double arg0, double arg1, double arg2, double arg3); // 1
  // proto:  void QMarginsF::QMarginsF(const QMargins & margins);
extern void* C_ZN9QMarginsFC2ERK8QMargins(void* arg0); // 1
  // proto:  qreal QMarginsF::left();
extern double C_ZNK9QMarginsF4leftEv(void* qthis); // 2
  // proto:  int QMargins::right();
extern int32_t C_ZNK8QMargins5rightEv(void* qthis); // 2
  // proto:  int QMargins::bottom();
extern int32_t C_ZNK8QMargins6bottomEv(void* qthis); // 2
  // proto:  void QMargins::setBottom(int bottom);
extern void C_ZN8QMargins9setBottomEi(void* qthis, int32_t arg0); // 2
  // proto:  void QMargins::setRight(int right);
extern void C_ZN8QMargins8setRightEi(void* qthis, int32_t arg0); // 2
  // proto:  void QMargins::setTop(int top);
extern void C_ZN8QMargins6setTopEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QMargins::isNull();
extern bool C_ZNK8QMargins6isNullEv(void* qthis); // 2
  // proto:  void QMargins::QMargins();
extern void* C_ZN8QMarginsC2Ev(); // 1
  // proto:  void QMargins::QMargins(int left, int top, int right, int bottom);
extern void* C_ZN8QMarginsC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QMargins::setLeft(int left);
extern void C_ZN8QMargins7setLeftEi(void* qthis, int32_t arg0); // 2
  // proto:  int QMargins::top();
extern int32_t C_ZNK8QMargins3topEv(void* qthis); // 2
  // proto:  int QMargins::left();
extern int32_t C_ZNK8QMargins4leftEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMargins)=16
type QMargins struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// right()
func (this *QMarginsF) Right(args ...interface{}) (ret interface{}) {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF5rightEv
    // invoke: qreal right()
    var ret0 = C.C_ZNK9QMarginsF5rightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "right", args)
  }

  return
}

// bottom()
func (this *QMarginsF) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF6bottomEv
    // invoke: qreal bottom()
    var ret0 = C.C_ZNK9QMarginsF6bottomEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "bottom", args)
  }

  return
}

// setBottom(qreal)
func (this *QMarginsF) Setbottom(args ...interface{}) () {
  // setBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF9setBottomEd
    // invoke: void setBottom(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMarginsF9setBottomEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setBottom", args)
  }

  return
}

// top()
func (this *QMarginsF) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF3topEv
    // invoke: qreal top()
    var ret0 = C.C_ZNK9QMarginsF3topEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "top", args)
  }

  return
}

// toMargins()
func (this *QMarginsF) Tomargins(args ...interface{}) (ret interface{}) {
  // toMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF9toMarginsEv
    // invoke: QMargins toMargins()
    var ret0 = C.C_ZNK9QMarginsF9toMarginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMargins{}) // "QMargins"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "toMargins", args)
  }

  return
}

// setRight(qreal)
func (this *QMarginsF) Setright(args ...interface{}) () {
  // setRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF8setRightEd
    // invoke: void setRight(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMarginsF8setRightEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setRight", args)
  }

  return
}

// setTop(qreal)
func (this *QMarginsF) Settop(args ...interface{}) () {
  // setTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF6setTopEd
    // invoke: void setTop(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMarginsF6setTopEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setTop", args)
  }

  return
}

// isNull()
func (this *QMarginsF) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK9QMarginsF6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "isNull", args)
  }

  return
}

// setLeft(qreal)
func (this *QMarginsF) Setleft(args ...interface{}) () {
  // setLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsF7setLeftEd
    // invoke: void setLeft(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN9QMarginsF7setLeftEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMarginsF", "setLeft", args)
  }

  return
}

// QMarginsF()
func NewQMarginsF(args ...interface{}) *QMarginsF {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMarginsFC1Ev
    // invoke: void QMarginsF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMarginsFC2Ev()
    return &QMarginsF{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QMarginsFC1Edddd
    // invoke: void QMarginsF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMarginsFC2Edddd(arg0, arg1, arg2, arg3)
    return &QMarginsF{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QMarginsFC1ERK8QMargins
    // invoke: void QMarginsF(const class QMargins &)
    var arg0 = args[0].(*QMargins).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMarginsFC2ERK8QMargins(arg0)
    return &QMarginsF{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMarginsF", "QMarginsF", args)
  }

  return nil // QMarginsF{Qclsinst:qthis}
}

// left()
func (this *QMarginsF) Left(args ...interface{}) (ret interface{}) {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMarginsF4leftEv
    // invoke: qreal left()
    var ret0 = C.C_ZNK9QMarginsF4leftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMarginsF", "left", args)
  }

  return
}

// right()
func (this *QMargins) Right(args ...interface{}) (ret interface{}) {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins5rightEv
    // invoke: int right()
    var ret0 = C.C_ZNK8QMargins5rightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMargins", "right", args)
  }

  return
}

// bottom()
func (this *QMargins) Bottom(args ...interface{}) (ret interface{}) {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins6bottomEv
    // invoke: int bottom()
    var ret0 = C.C_ZNK8QMargins6bottomEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMargins", "bottom", args)
  }

  return
}

// setBottom(int)
func (this *QMargins) Setbottom(args ...interface{}) () {
  // setBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins9setBottomEi
    // invoke: void setBottom(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMargins9setBottomEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setBottom", args)
  }

  return
}

// setRight(int)
func (this *QMargins) Setright(args ...interface{}) () {
  // setRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins8setRightEi
    // invoke: void setRight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMargins8setRightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setRight", args)
  }

  return
}

// setTop(int)
func (this *QMargins) Settop(args ...interface{}) () {
  // setTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins6setTopEi
    // invoke: void setTop(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMargins6setTopEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setTop", args)
  }

  return
}

// isNull()
func (this *QMargins) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK8QMargins6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMargins", "isNull", args)
  }

  return
}

// QMargins()
func NewQMargins(args ...interface{}) *QMargins {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMarginsC1Ev
    // invoke: void QMargins()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QMarginsC2Ev()
    return &QMargins{Qclsinst:qthis}
  case 1:
    // invoke: _ZN8QMarginsC1Eiiii
    // invoke: void QMargins(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QMarginsC2Eiiii(arg0, arg1, arg2, arg3)
    return &QMargins{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMargins", "QMargins", args)
  }

  return nil // QMargins{Qclsinst:qthis}
}

// setLeft(int)
func (this *QMargins) Setleft(args ...interface{}) () {
  // setLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMargins7setLeftEi
    // invoke: void setLeft(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMargins7setLeftEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMargins", "setLeft", args)
  }

  return
}

// top()
func (this *QMargins) Top(args ...interface{}) (ret interface{}) {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins3topEv
    // invoke: int top()
    var ret0 = C.C_ZNK8QMargins3topEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMargins", "top", args)
  }

  return
}

// left()
func (this *QMargins) Left(args ...interface{}) (ret interface{}) {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMargins4leftEv
    // invoke: int left()
    var ret0 = C.C_ZNK8QMargins4leftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMargins", "left", args)
  }

  return
}

// <= body block end

