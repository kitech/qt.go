package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qsizepolicy.h
// dst-file: /src/widgets/qsizepolicy.go
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
  // proto:  bool QSizePolicy::hasHeightForWidth();
extern void _ZNK11QSizePolicy17hasHeightForWidthEv(void* qthis);
  // proto:  bool QSizePolicy::retainSizeWhenHidden();
extern void _ZNK11QSizePolicy20retainSizeWhenHiddenEv(void* qthis);
  // proto:  bool QSizePolicy::hasWidthForHeight();
extern void _ZNK11QSizePolicy17hasWidthForHeightEv(void* qthis);
  // proto:  void QSizePolicy::transpose();
extern void _ZN11QSizePolicy9transposeEv(void* qthis);
  // proto:  void QSizePolicy::setWidthForHeight(bool b);
extern void _ZN11QSizePolicy17setWidthForHeightEb(void* qthis, bool arg0);
  // proto:  void QSizePolicy::setVerticalStretch(int stretchFactor);
extern void _ZN11QSizePolicy18setVerticalStretchEi(void* qthis, int arg0);
  // proto:  void QSizePolicy::setHeightForWidth(bool b);
extern void _ZN11QSizePolicy17setHeightForWidthEb(void* qthis, bool arg0);
  // proto:  void QSizePolicy::setRetainSizeWhenHidden(bool retainSize);
extern void _ZN11QSizePolicy23setRetainSizeWhenHiddenEb(void* qthis, bool arg0);
  // proto:  int QSizePolicy::horizontalStretch();
extern void _ZNK11QSizePolicy17horizontalStretchEv(void* qthis);
  // proto:  void QSizePolicy::setHorizontalStretch(int stretchFactor);
extern void _ZN11QSizePolicy20setHorizontalStretchEi(void* qthis, int arg0);
  // proto:  void QSizePolicy::QSizePolicy(int i);
extern void* dector_ZN11QSizePolicyC1Ei(int arg0);
extern void _ZN11QSizePolicyC1Ei(void* qthis, int arg0);
  // proto:  void QSizePolicy::QSizePolicy();
extern void* dector_ZN11QSizePolicyC1Ev();
extern void _ZN11QSizePolicyC1Ev(void* qthis);
  // proto:  int QSizePolicy::verticalStretch();
extern void _ZNK11QSizePolicy15verticalStretchEv(void* qthis);
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

// class sizeof(QSizePolicy)=4
type QSizePolicy struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QSizePolicy::hasHeightForWidth();
func (this *QSizePolicy) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasHeightForWidth", args)
  }

}

  // proto:  bool QSizePolicy::retainSizeWhenHidden();
func (this *QSizePolicy) retainSizeWhenHidden(args ...interface{}) () {
  // retainSizeWhenHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy20retainSizeWhenHiddenEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "retainSizeWhenHidden", args)
  }

}

  // proto:  bool QSizePolicy::hasWidthForHeight();
func (this *QSizePolicy) hasWidthForHeight(args ...interface{}) () {
  // hasWidthForHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy17hasWidthForHeightEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasWidthForHeight", args)
  }

}

  // proto:  void QSizePolicy::transpose();
func (this *QSizePolicy) transpose(args ...interface{}) () {
  // transpose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy9transposeEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "transpose", args)
  }

}

  // proto:  void QSizePolicy::setWidthForHeight(bool b);
func (this *QSizePolicy) setWidthForHeight(args ...interface{}) () {
  // setWidthForHeight(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy17setWidthForHeightEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizePolicy", "setWidthForHeight", args)
  }

}

  // proto:  void QSizePolicy::setVerticalStretch(int stretchFactor);
func (this *QSizePolicy) setVerticalStretch(args ...interface{}) () {
  // setVerticalStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy18setVerticalStretchEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizePolicy", "setVerticalStretch", args)
  }

}

  // proto:  void QSizePolicy::setHeightForWidth(bool b);
func (this *QSizePolicy) setHeightForWidth(args ...interface{}) () {
  // setHeightForWidth(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy17setHeightForWidthEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHeightForWidth", args)
  }

}

  // proto:  void QSizePolicy::setRetainSizeWhenHidden(bool retainSize);
func (this *QSizePolicy) setRetainSizeWhenHidden(args ...interface{}) () {
  // setRetainSizeWhenHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy23setRetainSizeWhenHiddenEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizePolicy", "setRetainSizeWhenHidden", args)
  }

}

  // proto:  int QSizePolicy::horizontalStretch();
func (this *QSizePolicy) horizontalStretch(args ...interface{}) () {
  // horizontalStretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy17horizontalStretchEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "horizontalStretch", args)
  }

}

  // proto:  void QSizePolicy::setHorizontalStretch(int stretchFactor);
func (this *QSizePolicy) setHorizontalStretch(args ...interface{}) () {
  // setHorizontalStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicy20setHorizontalStretchEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHorizontalStretch", args)
  }

}

  // proto:  void QSizePolicy::QSizePolicy(int i);
func NewQSizePolicy(args ...interface{}) QSizePolicy {
  return QSizePolicy{}
}

  // proto:  int QSizePolicy::verticalStretch();
func (this *QSizePolicy) verticalStretch(args ...interface{}) () {
  // verticalStretch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy15verticalStretchEv
  default:
    qtrt.ErrorResolve("QSizePolicy", "verticalStretch", args)
  }

}

// <= body block end

