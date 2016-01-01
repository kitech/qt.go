package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qbitarray.h
// dst-file: /src/core/qbitarray.go
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
// class sizeof(QBitRef)=16
type QBitRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBitArray)=8
type QBitArray struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQBitRef(args ...interface{}) QBitRef {
  return QBitRef{}
}


func NewQBitArray(args ...interface{}) QBitArray {
  return QBitArray{}
}


func (this *QBitArray) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7isEmptyEv
  default:
    qtrt.ErrorResolve("QBitArray", "isEmpty", args)
  }

}


func (this *QBitArray) setBit(args ...interface{}) () {
  // setBit(int)
  // setBit(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6setBitEi
  case 1:
    // invoke: _ZN9QBitArray6setBitEib
  default:
    qtrt.ErrorResolve("QBitArray", "setBit", args)
  }

}


func (this *QBitArray) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray4sizeEv
  default:
    qtrt.ErrorResolve("QBitArray", "size", args)
  }

}


func (this *QBitArray) swap(args ...interface{}) () {
  // swap(class QBitArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitArray{}) // "QBitArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4swapERS_
  default:
    qtrt.ErrorResolve("QBitArray", "swap", args)
  }

}


func (this *QBitArray) count(args ...interface{}) () {
  // count()
  // count(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray5countEv
  case 1:
    // invoke: _ZNK9QBitArray5countEb
  default:
    qtrt.ErrorResolve("QBitArray", "count", args)
  }

}


func (this *QBitArray) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6detachEv
  default:
    qtrt.ErrorResolve("QBitArray", "detach", args)
  }

}


func (this *QBitArray) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray2atEi
  default:
    qtrt.ErrorResolve("QBitArray", "at", args)
  }

}


func (this *QBitArray) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray5clearEv
  default:
    qtrt.ErrorResolve("QBitArray", "clear", args)
  }

}


func (this *QBitArray) clearBit(args ...interface{}) () {
  // clearBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8clearBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "clearBit", args)
  }

}


func (this *QBitArray) testBit(args ...interface{}) () {
  // testBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7testBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "testBit", args)
  }

}


func (this *QBitArray) truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8truncateEi
  default:
    qtrt.ErrorResolve("QBitArray", "truncate", args)
  }

}


func (this *QBitArray) toggleBit(args ...interface{}) () {
  // toggleBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray9toggleBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "toggleBit", args)
  }

}


func (this *QBitArray) fill(args ...interface{}) () {
  // fill(_Bool, int, int)
  // fill(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4fillEbii
  case 1:
    // invoke: _ZN9QBitArray4fillEbi
  default:
    qtrt.ErrorResolve("QBitArray", "fill", args)
  }

}


func (this *QBitArray) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray6isNullEv
  default:
    qtrt.ErrorResolve("QBitArray", "isNull", args)
  }

}


func (this *QBitArray) resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6resizeEi
  default:
    qtrt.ErrorResolve("QBitArray", "resize", args)
  }

}


func (this *QBitArray) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray10isDetachedEv
  default:
    qtrt.ErrorResolve("QBitArray", "isDetached", args)
  }

}

// <= body block end

