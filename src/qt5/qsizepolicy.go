package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN11QSizePolicyC2Ev(); // 1
  // proto:  void QSizePolicy::setHorizontalStretch(int stretchFactor);
extern void C_ZN11QSizePolicy20setHorizontalStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  Qt::Orientations QSizePolicy::expandingDirections();
extern void C_ZNK11QSizePolicy19expandingDirectionsEv(void* qthis); // 2
  // proto:  void QSizePolicy::setWidthForHeight(bool b);
extern void C_ZN11QSizePolicy17setWidthForHeightEb(void* qthis, bool arg0); // 2
  // proto:  bool QSizePolicy::hasWidthForHeight();
extern bool C_ZNK11QSizePolicy17hasWidthForHeightEv(void* qthis); // 2
  // proto:  int QSizePolicy::verticalStretch();
extern int32_t C_ZNK11QSizePolicy15verticalStretchEv(void* qthis); // 2
  // proto:  void QSizePolicy::transpose();
extern void C_ZN11QSizePolicy9transposeEv(void* qthis); // 2
  // proto:  bool QSizePolicy::retainSizeWhenHidden();
extern bool C_ZNK11QSizePolicy20retainSizeWhenHiddenEv(void* qthis); // 2
  // proto:  void QSizePolicy::setVerticalStretch(int stretchFactor);
extern void C_ZN11QSizePolicy18setVerticalStretchEi(void* qthis, int32_t arg0); // 2
  // proto:  QSizePolicy::Policy QSizePolicy::verticalPolicy();
extern void C_ZNK11QSizePolicy14verticalPolicyEv(void* qthis); // 2
  // proto:  void QSizePolicy::setHeightForWidth(bool b);
extern void C_ZN11QSizePolicy17setHeightForWidthEb(void* qthis, bool arg0); // 2
  // proto:  bool QSizePolicy::hasHeightForWidth();
extern bool C_ZNK11QSizePolicy17hasHeightForWidthEv(void* qthis); // 2
  // proto:  int QSizePolicy::horizontalStretch();
extern int32_t C_ZNK11QSizePolicy17horizontalStretchEv(void* qthis); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QSizePolicy()
func NewQSizePolicy(args ...interface{}) *QSizePolicy {
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
    qthis = C.C_ZN11QSizePolicyC2Ev()
    return &QSizePolicy{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSizePolicy", "QSizePolicy", args)
  }

  return nil // QSizePolicy{Qclsinst:qthis}
}

// setHorizontalStretch(int)
func (this *QSizePolicy) Sethorizontalstretch(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy20setHorizontalStretchEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHorizontalStretch", args)
  }

  return
}

// expandingDirections()
func (this *QSizePolicy) Expandingdirections(args ...interface{}) () {
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
    C.C_ZNK11QSizePolicy19expandingDirectionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "expandingDirections", args)
  }

  return
}

// setWidthForHeight(_Bool)
func (this *QSizePolicy) Setwidthforheight(args ...interface{}) () {
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
    C.C_ZN11QSizePolicy17setWidthForHeightEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setWidthForHeight", args)
  }

  return
}

// hasWidthForHeight()
func (this *QSizePolicy) Haswidthforheight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSizePolicy17hasWidthForHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasWidthForHeight", args)
  }

  return
}

// verticalStretch()
func (this *QSizePolicy) Verticalstretch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSizePolicy15verticalStretchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizePolicy", "verticalStretch", args)
  }

  return
}

// transpose()
func (this *QSizePolicy) Transpose(args ...interface{}) () {
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
    C.C_ZN11QSizePolicy9transposeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "transpose", args)
  }

  return
}

// retainSizeWhenHidden()
func (this *QSizePolicy) Retainsizewhenhidden(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSizePolicy20retainSizeWhenHiddenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizePolicy", "retainSizeWhenHidden", args)
  }

  return
}

// setVerticalStretch(int)
func (this *QSizePolicy) Setverticalstretch(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QSizePolicy18setVerticalStretchEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setVerticalStretch", args)
  }

  return
}

// verticalPolicy()
func (this *QSizePolicy) Verticalpolicy(args ...interface{}) () {
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
    C.C_ZNK11QSizePolicy14verticalPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "verticalPolicy", args)
  }

  return
}

// setHeightForWidth(_Bool)
func (this *QSizePolicy) Setheightforwidth(args ...interface{}) () {
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
    C.C_ZN11QSizePolicy17setHeightForWidthEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHeightForWidth", args)
  }

  return
}

// hasHeightForWidth()
func (this *QSizePolicy) Hasheightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSizePolicy17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizePolicy", "hasHeightForWidth", args)
  }

  return
}

// horizontalStretch()
func (this *QSizePolicy) Horizontalstretch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QSizePolicy17horizontalStretchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSizePolicy", "horizontalStretch", args)
  }

  return
}

// setRetainSizeWhenHidden(_Bool)
func (this *QSizePolicy) Setretainsizewhenhidden(args ...interface{}) () {
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
    C.C_ZN11QSizePolicy23setRetainSizeWhenHiddenEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizePolicy", "setRetainSizeWhenHidden", args)
  }

  return
}

// controlType()
func (this *QSizePolicy) Controltype(args ...interface{}) () {
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
    C.C_ZNK11QSizePolicy11controlTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "controlType", args)
  }

  return
}

// horizontalPolicy()
func (this *QSizePolicy) Horizontalpolicy(args ...interface{}) () {
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
    C.C_ZNK11QSizePolicy16horizontalPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSizePolicy", "horizontalPolicy", args)
  }

  return
}

// <= body block end

