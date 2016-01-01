package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qabstractslider.h
// dst-file: /src/widgets/qabstractslider.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QAbstractSlider)=1
type QAbstractSlider struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _sliderReleased QAbstractSlider_sliderReleased_signal;
//  _rangeChanged QAbstractSlider_rangeChanged_signal;
//  _sliderPressed QAbstractSlider_sliderPressed_signal;
//  _actionTriggered QAbstractSlider_actionTriggered_signal;
//  _valueChanged QAbstractSlider_valueChanged_signal;
//  _sliderMoved QAbstractSlider_sliderMoved_signal;
}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderPosition", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "isSliderDown", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "value", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedControls", args)
  }

}


func NewQAbstractSlider(args ...interface{}) QAbstractSlider {
  return QAbstractSlider{}
}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "minimum", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "singleStep", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "pageStep", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMaximum", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedControls", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setValue", args)
  }

}


func (this *QAbstractSlider) FreeQAbstractSlider(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractSlider", "~QAbstractSlider", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setPageStep", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderDown", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "maximum", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSingleStep", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedAppearance", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "hasTracking", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedAppearance", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "sliderPosition", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setTracking", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setRange", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMinimum", args)
  }

}

// <= body block end

