package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qlcdnumber.h
// dst-file: /src/widgets/qlcdnumber.go
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
// class sizeof(QLCDNumber)=1
type QLCDNumber struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
//  _overflow QLCDNumber_overflow_signal;
}


func (this *QLCDNumber) display(args ...interface{}) () {
  // display(int)
  // display(double)
  // display(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber7displayEi
  case 1:
    // invoke: _ZN10QLCDNumber7displayEd
  case 2:
    // invoke: _ZN10QLCDNumber7displayERK7QString
  default:
    qtrt.ErrorResolve("QLCDNumber", "display", args)
  }

}


func (this *QLCDNumber) setHexMode(args ...interface{}) () {
  // setHexMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setHexModeEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "setHexMode", args)
  }

}


func (this *QLCDNumber) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber10metaObjectEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "metaObject", args)
  }

}


func NewQLCDNumber(args ...interface{}) QLCDNumber {
  return QLCDNumber{}
}


func (this *QLCDNumber) digitCount(args ...interface{}) () {
  // digitCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber10digitCountEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "digitCount", args)
  }

}


func (this *QLCDNumber) FreeQLCDNumber(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLCDNumber", "~QLCDNumber", args)
  }

}


func (this *QLCDNumber) checkOverflow(args ...interface{}) () {
  // checkOverflow(int)
  // checkOverflow(double)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber13checkOverflowEi
  case 1:
    // invoke: _ZNK10QLCDNumber13checkOverflowEd
  default:
    qtrt.ErrorResolve("QLCDNumber", "checkOverflow", args)
  }

}


func (this *QLCDNumber) setDecMode(args ...interface{}) () {
  // setDecMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setDecModeEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDecMode", args)
  }

}


func (this *QLCDNumber) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber8sizeHintEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "sizeHint", args)
  }

}


func (this *QLCDNumber) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber5valueEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "value", args)
  }

}


func (this *QLCDNumber) setBinMode(args ...interface{}) () {
  // setBinMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setBinModeEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "setBinMode", args)
  }

}


func (this *QLCDNumber) intValue(args ...interface{}) () {
  // intValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber8intValueEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "intValue", args)
  }

}


func (this *QLCDNumber) setDigitCount(args ...interface{}) () {
  // setDigitCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber13setDigitCountEi
  default:
    qtrt.ErrorResolve("QLCDNumber", "setDigitCount", args)
  }

}


func (this *QLCDNumber) setSmallDecimalPoint(args ...interface{}) () {
  // setSmallDecimalPoint(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber20setSmallDecimalPointEb
  default:
    qtrt.ErrorResolve("QLCDNumber", "setSmallDecimalPoint", args)
  }

}


func (this *QLCDNumber) smallDecimalPoint(args ...interface{}) () {
  // smallDecimalPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QLCDNumber17smallDecimalPointEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "smallDecimalPoint", args)
  }

}


func (this *QLCDNumber) setOctMode(args ...interface{}) () {
  // setOctMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QLCDNumber10setOctModeEv
  default:
    qtrt.ErrorResolve("QLCDNumber", "setOctMode", args)
  }

}

// <= body block end

