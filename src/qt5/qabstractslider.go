package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qabstractslider.h
// dst-file: /src/widgets/qabstractslider.go
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
  // proto:  void QAbstractSlider::setTracking(bool enable);
extern void C_ZN15QAbstractSlider11setTrackingEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Orientation QAbstractSlider::orientation();
extern void C_ZNK15QAbstractSlider11orientationEv(void* qthis); // 4
  // proto:  bool QAbstractSlider::invertedControls();
extern bool C_ZNK15QAbstractSlider16invertedControlsEv(void* qthis); // 4
  // proto:  int QAbstractSlider::singleStep();
extern int32_t C_ZNK15QAbstractSlider10singleStepEv(void* qthis); // 4
  // proto:  int QAbstractSlider::minimum();
extern int32_t C_ZNK15QAbstractSlider7minimumEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setSliderDown(bool );
extern void C_ZN15QAbstractSlider13setSliderDownEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSlider::~QAbstractSlider();
extern void C_ZN15QAbstractSliderD2Ev(void* qthis); // 4
  // proto:  int QAbstractSlider::pageStep();
extern int32_t C_ZNK15QAbstractSlider8pageStepEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setInvertedControls(bool );
extern void C_ZN15QAbstractSlider19setInvertedControlsEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSlider::setPageStep(int );
extern void C_ZN15QAbstractSlider11setPageStepEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::isSliderDown();
extern bool C_ZNK15QAbstractSlider12isSliderDownEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setSingleStep(int );
extern void C_ZN15QAbstractSlider13setSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::hasTracking();
extern bool C_ZNK15QAbstractSlider11hasTrackingEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setMaximum(int );
extern void C_ZN15QAbstractSlider10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractSlider::setValue(int );
extern void C_ZN15QAbstractSlider8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::invertedAppearance();
extern bool C_ZNK15QAbstractSlider18invertedAppearanceEv(void* qthis); // 4
  // proto:  int QAbstractSlider::sliderPosition();
extern int32_t C_ZNK15QAbstractSlider14sliderPositionEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setRange(int min, int max);
extern void C_ZN15QAbstractSlider8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QAbstractSlider::setInvertedAppearance(bool );
extern void C_ZN15QAbstractSlider21setInvertedAppearanceEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QAbstractSlider::metaObject();
extern void C_ZNK15QAbstractSlider10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractSlider::QAbstractSlider(QWidget * parent);
extern void* C_ZN15QAbstractSliderC2EP7QWidget(void* arg0); // 3
  // proto:  int QAbstractSlider::maximum();
extern int32_t C_ZNK15QAbstractSlider7maximumEv(void* qthis); // 4
  // proto:  int QAbstractSlider::value();
extern int32_t C_ZNK15QAbstractSlider5valueEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setSliderPosition(int );
extern void C_ZN15QAbstractSlider17setSliderPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractSlider::setMinimum(int );
extern void C_ZN15QAbstractSlider10setMinimumEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QAbstractSlider)=1
type QAbstractSlider struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _sliderReleased QAbstractSlider_sliderReleased_signal;
//  _rangeChanged QAbstractSlider_rangeChanged_signal;
//  _sliderPressed QAbstractSlider_sliderPressed_signal;
//  _actionTriggered QAbstractSlider_actionTriggered_signal;
//  _valueChanged QAbstractSlider_valueChanged_signal;
//  _sliderMoved QAbstractSlider_sliderMoved_signal;
}

// setTracking(_Bool)
func (this *QAbstractSlider) Settracking(args ...interface{}) () {
  // setTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider11setTrackingEb
    // invoke: void setTracking(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider11setTrackingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setTracking", args)
  }

  return
}

// orientation()
func (this *QAbstractSlider) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK15QAbstractSlider11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "orientation", args)
  }

  return
}

// invertedControls()
func (this *QAbstractSlider) Invertedcontrols(args ...interface{}) (ret interface{}) {
  // invertedControls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider16invertedControlsEv
    // invoke: bool invertedControls()
    var ret0 = C.C_ZNK15QAbstractSlider16invertedControlsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedControls", args)
  }

  return
}

// singleStep()
func (this *QAbstractSlider) Singlestep(args ...interface{}) (ret interface{}) {
  // singleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider10singleStepEv
    // invoke: int singleStep()
    var ret0 = C.C_ZNK15QAbstractSlider10singleStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "singleStep", args)
  }

  return
}

// minimum()
func (this *QAbstractSlider) Minimum(args ...interface{}) (ret interface{}) {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider7minimumEv
    // invoke: int minimum()
    var ret0 = C.C_ZNK15QAbstractSlider7minimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "minimum", args)
  }

  return
}

// setSliderDown(_Bool)
func (this *QAbstractSlider) Setsliderdown(args ...interface{}) () {
  // setSliderDown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider13setSliderDownEb
    // invoke: void setSliderDown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider13setSliderDownEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderDown", args)
  }

  return
}

// ~QAbstractSlider()
func (this *QAbstractSlider) Freeqabstractslider(args ...interface{}) () {
  // ~QAbstractSlider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSliderD0Ev
    // invoke: void ~QAbstractSlider()
    C.C_ZN15QAbstractSliderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "~QAbstractSlider", args)
  }

  return
}

// pageStep()
func (this *QAbstractSlider) Pagestep(args ...interface{}) (ret interface{}) {
  // pageStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider8pageStepEv
    // invoke: int pageStep()
    var ret0 = C.C_ZNK15QAbstractSlider8pageStepEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "pageStep", args)
  }

  return
}

// setInvertedControls(_Bool)
func (this *QAbstractSlider) Setinvertedcontrols(args ...interface{}) () {
  // setInvertedControls(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider19setInvertedControlsEb
    // invoke: void setInvertedControls(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider19setInvertedControlsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedControls", args)
  }

  return
}

// setPageStep(int)
func (this *QAbstractSlider) Setpagestep(args ...interface{}) () {
  // setPageStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider11setPageStepEi
    // invoke: void setPageStep(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider11setPageStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setPageStep", args)
  }

  return
}

// isSliderDown()
func (this *QAbstractSlider) Issliderdown(args ...interface{}) (ret interface{}) {
  // isSliderDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider12isSliderDownEv
    // invoke: bool isSliderDown()
    var ret0 = C.C_ZNK15QAbstractSlider12isSliderDownEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "isSliderDown", args)
  }

  return
}

// setSingleStep(int)
func (this *QAbstractSlider) Setsinglestep(args ...interface{}) () {
  // setSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider13setSingleStepEi
    // invoke: void setSingleStep(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider13setSingleStepEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSingleStep", args)
  }

  return
}

// hasTracking()
func (this *QAbstractSlider) Hastracking(args ...interface{}) (ret interface{}) {
  // hasTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider11hasTrackingEv
    // invoke: bool hasTracking()
    var ret0 = C.C_ZNK15QAbstractSlider11hasTrackingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "hasTracking", args)
  }

  return
}

// setMaximum(int)
func (this *QAbstractSlider) Setmaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider10setMaximumEi
    // invoke: void setMaximum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider10setMaximumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMaximum", args)
  }

  return
}

// setValue(int)
func (this *QAbstractSlider) Setvalue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider8setValueEi
    // invoke: void setValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider8setValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setValue", args)
  }

  return
}

// invertedAppearance()
func (this *QAbstractSlider) Invertedappearance(args ...interface{}) (ret interface{}) {
  // invertedAppearance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider18invertedAppearanceEv
    // invoke: bool invertedAppearance()
    var ret0 = C.C_ZNK15QAbstractSlider18invertedAppearanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedAppearance", args)
  }

  return
}

// sliderPosition()
func (this *QAbstractSlider) Sliderposition(args ...interface{}) (ret interface{}) {
  // sliderPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider14sliderPositionEv
    // invoke: int sliderPosition()
    var ret0 = C.C_ZNK15QAbstractSlider14sliderPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "sliderPosition", args)
  }

  return
}

// setRange(int, int)
func (this *QAbstractSlider) Setrange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider8setRangeEii
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QAbstractSlider8setRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setRange", args)
  }

  return
}

// setInvertedAppearance(_Bool)
func (this *QAbstractSlider) Setinvertedappearance(args ...interface{}) () {
  // setInvertedAppearance(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider21setInvertedAppearanceEb
    // invoke: void setInvertedAppearance(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider21setInvertedAppearanceEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedAppearance", args)
  }

  return
}

// metaObject()
func (this *QAbstractSlider) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QAbstractSlider10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "metaObject", args)
  }

  return
}

// QAbstractSlider(class QWidget *)
func NewQAbstractSlider(args ...interface{}) *QAbstractSlider {
  // QAbstractSlider(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSliderC1EP7QWidget
    // invoke: void QAbstractSlider(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QAbstractSliderC2EP7QWidget(arg0)
    return &QAbstractSlider{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "QAbstractSlider", args)
  }

  return nil // QAbstractSlider{Qclsinst:qthis}
}

// maximum()
func (this *QAbstractSlider) Maximum(args ...interface{}) (ret interface{}) {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider7maximumEv
    // invoke: int maximum()
    var ret0 = C.C_ZNK15QAbstractSlider7maximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "maximum", args)
  }

  return
}

// value()
func (this *QAbstractSlider) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider5valueEv
    // invoke: int value()
    var ret0 = C.C_ZNK15QAbstractSlider5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSlider", "value", args)
  }

  return
}

// setSliderPosition(int)
func (this *QAbstractSlider) Setsliderposition(args ...interface{}) () {
  // setSliderPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider17setSliderPositionEi
    // invoke: void setSliderPosition(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider17setSliderPositionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderPosition", args)
  }

  return
}

// setMinimum(int)
func (this *QAbstractSlider) Setminimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider10setMinimumEi
    // invoke: void setMinimum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider10setMinimumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMinimum", args)
  }

  return
}

// <= body block end

