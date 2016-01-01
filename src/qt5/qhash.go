package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qhash.h
// dst-file: /src/core/qhash.go
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
// class sizeof(QHashDummyValue)=1
type QHashDummyValue struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QHashData)=1
type QHashData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QHashData) hasShrunk(args ...interface{}) () {
  // hasShrunk()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QHashData9hasShrunkEv
  default:
    qtrt.ErrorResolve("QHashData", "hasShrunk", args)
 }

}


func (this *QHashData) allocateNode(args ...interface{}) () {
  // allocateNode(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QHashData12allocateNodeEi
  default:
    qtrt.ErrorResolve("QHashData", "allocateNode", args)
 }

}


func (this *QHashData) willGrow(args ...interface{}) () {
  // willGrow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8willGrowEv
  default:
    qtrt.ErrorResolve("QHashData", "willGrow", args)
 }

}


func (this *QHashData) rehash(args ...interface{}) () {
  // rehash(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QHashData6rehashEi
  default:
    qtrt.ErrorResolve("QHashData", "rehash", args)
 }

}


func (this *QHashData) freeNode(args ...interface{}) () {
  // freeNode(void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QHashData8freeNodeEPv
  default:
    qtrt.ErrorResolve("QHashData", "freeNode", args)
 }

}

// <= body block end

