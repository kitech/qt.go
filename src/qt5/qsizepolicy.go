package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QSizePolicy::QSizePolicy();
extern void C_ZN11QSizePolicyC2Ev(void* qthis); // 1
  // proto:  void QSizePolicy::setHorizontalStretch(int stretchFactor);
extern void C_ZN11QSizePolicy20setHorizontalStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  Qt::Orientations QSizePolicy::expandingDirections();
extern void C_ZNK11QSizePolicy19expandingDirectionsEv(void* qthis); // 2
  // proto:  void QSizePolicy::setWidthForHeight(bool b);
extern void C_ZN11QSizePolicy17setWidthForHeightEb(void* qthis, bool arg0); // 2
  // proto:  bool QSizePolicy::hasWidthForHeight();
extern void C_ZNK11QSizePolicy17hasWidthForHeightEv(void* qthis); // 2
  // proto:  int QSizePolicy::verticalStretch();
extern void C_ZNK11QSizePolicy15verticalStretchEv(void* qthis); // 2
  // proto:  void QSizePolicy::transpose();
extern void C_ZN11QSizePolicy9transposeEv(void* qthis); // 2
  // proto:  bool QSizePolicy::retainSizeWhenHidden();
extern void C_ZNK11QSizePolicy20retainSizeWhenHiddenEv(void* qthis); // 2
  // proto:  void QSizePolicy::setVerticalStretch(int stretchFactor);
extern void C_ZN11QSizePolicy18setVerticalStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  QSizePolicy::Policy QSizePolicy::verticalPolicy();
extern void C_ZNK11QSizePolicy14verticalPolicyEv(void* qthis); // 2
  // proto:  void QSizePolicy::setHeightForWidth(bool b);
extern void C_ZN11QSizePolicy17setHeightForWidthEb(void* qthis, bool arg0); // 2
  // proto:  bool QSizePolicy::hasHeightForWidth();
extern void C_ZNK11QSizePolicy17hasHeightForWidthEv(void* qthis); // 2
  // proto:  int QSizePolicy::horizontalStretch();
extern void C_ZNK11QSizePolicy17horizontalStretchEv(void* qthis); // 2
  // proto:  void QSizePolicy::setRetainSizeWhenHidden(bool retainSize);
extern void C_ZN11QSizePolicy23setRetainSizeWhenHiddenEb(void* qthis, bool arg0); // 2
  // proto:  QSizePolicy::ControlType QSizePolicy::controlType();
extern void C_ZNK11QSizePolicy11controlTypeEv(void* qthis); // 4
  // proto:  QSizePolicy::Policy QSizePolicy::horizontalPolicy();
extern void C_ZNK11QSizePolicy16horizontalPolicyEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// QSizePolicy()
func NewQSizePolicy(args ...interface{}) QSizePolicy {
  // QSizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSizePolicyC1Ev
    // invoke: void QSizePolicy()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QSizePolicyC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QSizePolicy", "QSizePolicy", args)
  }

  return QSizePolicy{}
}

// setHorizontalStretch(int)
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
    // invoke: void setHorizontalStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy20setHorizontalStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHorizontalStretch", args)
  }

}

// expandingDirections()
func (this *QSizePolicy) expandingDirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C.C_ZNK11QSizePolicy19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "expandingDirections", args)
  }

}

// setWidthForHeight(_Bool)
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
    // invoke: void setWidthForHeight(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy17setWidthForHeightEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setWidthForHeight", args)
  }

}

// hasWidthForHeight()
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
    // invoke: bool hasWidthForHeight()
    C.C_ZNK11QSizePolicy17hasWidthForHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasWidthForHeight", args)
  }

}

// verticalStretch()
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
    // invoke: int verticalStretch()
    C.C_ZNK11QSizePolicy15verticalStretchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "verticalStretch", args)
  }

}

// transpose()
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
    // invoke: void transpose()
    C.C_ZN11QSizePolicy9transposeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "transpose", args)
  }

}

// retainSizeWhenHidden()
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
    // invoke: bool retainSizeWhenHidden()
    C.C_ZNK11QSizePolicy20retainSizeWhenHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "retainSizeWhenHidden", args)
  }

}

// setVerticalStretch(int)
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
    // invoke: void setVerticalStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy18setVerticalStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setVerticalStretch", args)
  }

}

// verticalPolicy()
func (this *QSizePolicy) verticalPolicy(args ...interface{}) () {
  // verticalPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy14verticalPolicyEv
    // invoke: QSizePolicy::Policy verticalPolicy()
    C.C_ZNK11QSizePolicy14verticalPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "verticalPolicy", args)
  }

}

// setHeightForWidth(_Bool)
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
    // invoke: void setHeightForWidth(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy17setHeightForWidthEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHeightForWidth", args)
  }

}

// hasHeightForWidth()
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
    // invoke: bool hasHeightForWidth()
    C.C_ZNK11QSizePolicy17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasHeightForWidth", args)
  }

}

// horizontalStretch()
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
    // invoke: int horizontalStretch()
    C.C_ZNK11QSizePolicy17horizontalStretchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "horizontalStretch", args)
  }

}

// setRetainSizeWhenHidden(_Bool)
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
    // invoke: void setRetainSizeWhenHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy23setRetainSizeWhenHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setRetainSizeWhenHidden", args)
  }

}

// controlType()
func (this *QSizePolicy) controlType(args ...interface{}) () {
  // controlType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy11controlTypeEv
    // invoke: QSizePolicy::ControlType controlType()
    C.C_ZNK11QSizePolicy11controlTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "controlType", args)
  }

}

// horizontalPolicy()
func (this *QSizePolicy) horizontalPolicy(args ...interface{}) () {
  // horizontalPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSizePolicy16horizontalPolicyEv
    // invoke: QSizePolicy::Policy horizontalPolicy()
    C.C_ZNK11QSizePolicy16horizontalPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "horizontalPolicy", args)
  }

}

// <= body block end

