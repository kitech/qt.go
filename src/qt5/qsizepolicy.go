package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qsizepolicy.h
// dst-file: /src/widgets/qsizepolicy.go
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
// class sizeof(QSizePolicy)=4
type QSizePolicy struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QSizePolicy", "setWidthForHeight", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QSizePolicy", "setVerticalStretch", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHeightForWidth", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QSizePolicy", "setRetainSizeWhenHidden", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QSizePolicy", "setHorizontalStretch", args)
 }

}


func NewQSizePolicy(args ...interface{})() {
}


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

