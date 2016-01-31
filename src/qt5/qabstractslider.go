package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK15QAbstractSlider16invertedControlsEv(void* qthis); // 4
  // proto:  int QAbstractSlider::singleStep();
extern void C_ZNK15QAbstractSlider10singleStepEv(void* qthis); // 4
  // proto:  int QAbstractSlider::minimum();
extern void C_ZNK15QAbstractSlider7minimumEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setSliderDown(bool );
extern void C_ZN15QAbstractSlider13setSliderDownEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSlider::~QAbstractSlider();
extern void C_ZN15QAbstractSliderD2Ev(void* qthis); // 4
  // proto:  int QAbstractSlider::pageStep();
extern void C_ZNK15QAbstractSlider8pageStepEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setInvertedControls(bool );
extern void C_ZN15QAbstractSlider19setInvertedControlsEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSlider::setPageStep(int );
extern void C_ZN15QAbstractSlider11setPageStepEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::isSliderDown();
extern void C_ZNK15QAbstractSlider12isSliderDownEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setSingleStep(int );
extern void C_ZN15QAbstractSlider13setSingleStepEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::hasTracking();
extern void C_ZNK15QAbstractSlider11hasTrackingEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setMaximum(int );
extern void C_ZN15QAbstractSlider10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAbstractSlider::setValue(int );
extern void C_ZN15QAbstractSlider8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QAbstractSlider::invertedAppearance();
extern void C_ZNK15QAbstractSlider18invertedAppearanceEv(void* qthis); // 4
  // proto:  int QAbstractSlider::sliderPosition();
extern void C_ZNK15QAbstractSlider14sliderPositionEv(void* qthis); // 4
  // proto:  void QAbstractSlider::setRange(int min, int max);
extern void C_ZN15QAbstractSlider8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QAbstractSlider::setInvertedAppearance(bool );
extern void C_ZN15QAbstractSlider21setInvertedAppearanceEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QAbstractSlider::metaObject();
extern void C_ZNK15QAbstractSlider10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractSlider::QAbstractSlider(QWidget * parent);
extern void* C_ZN15QAbstractSliderC2EP7QWidget(void* arg0); // 3
  // proto:  int QAbstractSlider::maximum();
extern void C_ZNK15QAbstractSlider7maximumEv(void* qthis); // 4
  // proto:  int QAbstractSlider::value();
extern void C_ZNK15QAbstractSlider5valueEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _sliderReleased QAbstractSlider_sliderReleased_signal;
//  _rangeChanged QAbstractSlider_rangeChanged_signal;
//  _sliderPressed QAbstractSlider_sliderPressed_signal;
//  _actionTriggered QAbstractSlider_actionTriggered_signal;
//  _valueChanged QAbstractSlider_valueChanged_signal;
//  _sliderMoved QAbstractSlider_sliderMoved_signal;
}

// setTracking(_Bool)
func (this *QAbstractSlider) setTracking(args ...interface{}) () {
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
    C.C_ZN15QAbstractSlider11setTrackingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setTracking", args)
  }

}

// orientation()
func (this *QAbstractSlider) orientation(args ...interface{}) () {
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
    C.C_ZNK15QAbstractSlider11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "orientation", args)
  }

}

// invertedControls()
func (this *QAbstractSlider) invertedControls(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider16invertedControlsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedControls", args)
  }

}

// singleStep()
func (this *QAbstractSlider) singleStep(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider10singleStepEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "singleStep", args)
  }

}

// minimum()
func (this *QAbstractSlider) minimum(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider7minimumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "minimum", args)
  }

}

// setSliderDown(_Bool)
func (this *QAbstractSlider) setSliderDown(args ...interface{}) () {
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
    C.C_ZN15QAbstractSlider13setSliderDownEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderDown", args)
  }

}

// ~QAbstractSlider()
func (this *QAbstractSlider) FreeQAbstractSlider(args ...interface{}) () {
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
    C.C_ZN15QAbstractSliderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "~QAbstractSlider", args)
  }

}

// pageStep()
func (this *QAbstractSlider) pageStep(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider8pageStepEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "pageStep", args)
  }

}

// setInvertedControls(_Bool)
func (this *QAbstractSlider) setInvertedControls(args ...interface{}) () {
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
    C.C_ZN15QAbstractSlider19setInvertedControlsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedControls", args)
  }

}

// setPageStep(int)
func (this *QAbstractSlider) setPageStep(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider11setPageStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setPageStep", args)
  }

}

// isSliderDown()
func (this *QAbstractSlider) isSliderDown(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider12isSliderDownEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "isSliderDown", args)
  }

}

// setSingleStep(int)
func (this *QAbstractSlider) setSingleStep(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider13setSingleStepEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSingleStep", args)
  }

}

// hasTracking()
func (this *QAbstractSlider) hasTracking(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider11hasTrackingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "hasTracking", args)
  }

}

// setMaximum(int)
func (this *QAbstractSlider) setMaximum(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider10setMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMaximum", args)
  }

}

// setValue(int)
func (this *QAbstractSlider) setValue(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider8setValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setValue", args)
  }

}

// invertedAppearance()
func (this *QAbstractSlider) invertedAppearance(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider18invertedAppearanceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedAppearance", args)
  }

}

// sliderPosition()
func (this *QAbstractSlider) sliderPosition(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider14sliderPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "sliderPosition", args)
  }

}

// setRange(int, int)
func (this *QAbstractSlider) setRange(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QAbstractSlider8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setRange", args)
  }

}

// setInvertedAppearance(_Bool)
func (this *QAbstractSlider) setInvertedAppearance(args ...interface{}) () {
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
    C.C_ZN15QAbstractSlider21setInvertedAppearanceEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedAppearance", args)
  }

}

// metaObject()
func (this *QAbstractSlider) metaObject(args ...interface{}) () {
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
    C.C_ZNK15QAbstractSlider10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "metaObject", args)
  }

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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QAbstractSliderC2EP7QWidget(arg0)
    return &QAbstractSlider{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "QAbstractSlider", args)
  }

  return nil // QAbstractSlider{qclsinst:qthis}
}

// maximum()
func (this *QAbstractSlider) maximum(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider7maximumEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "maximum", args)
  }

}

// value()
func (this *QAbstractSlider) value(args ...interface{}) () {
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
    var ret = C.C_ZNK15QAbstractSlider5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAbstractSlider", "value", args)
  }

}

// setSliderPosition(int)
func (this *QAbstractSlider) setSliderPosition(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider17setSliderPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderPosition", args)
  }

}

// setMinimum(int)
func (this *QAbstractSlider) setMinimum(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAbstractSlider10setMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMinimum", args)
  }

}

// <= body block end

