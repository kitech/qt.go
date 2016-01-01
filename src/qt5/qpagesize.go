package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qpagesize.h
// dst-file: /src/gui/qpagesize.go
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
// class sizeof(QPageSize)=1
type QPageSize struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQPageSize(args ...interface{})() {
}


func (this *QPageSize) FreeQPageSize(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPageSize", "~QPageSize", args)
 }

}


func (this *QPageSize) key(args ...interface{}) () {
  // key()
  // key(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize3keyEv
  case 1:
    // invoke: _ZN9QPageSize3keyENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "key", args)
 }

}


func (this *QPageSize) name(args ...interface{}) () {
  // name()
  // name(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize4nameEv
  case 1:
    // invoke: _ZN9QPageSize4nameENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "name", args)
 }

}


func (this *QPageSize) definitionSize(args ...interface{}) () {
  // definitionSize()
  // definitionSize(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14definitionSizeEv
  case 1:
    // invoke: _ZN9QPageSize14definitionSizeENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "definitionSize", args)
 }

}


func (this *QPageSize) swap(args ...interface{}) () {
  // swap(class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize4swapERS_
  default:
    qtrt.ErrorResolve("QPageSize", "swap", args)
 }

}


func (this *QPageSize) windowsId(args ...interface{}) () {
  // windowsId(enum QPageSize::PageSizeId)
  // windowsId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize9windowsIdENS_10PageSizeIdE
  case 1:
    // invoke: _ZNK9QPageSize9windowsIdEv
  default:
    qtrt.ErrorResolve("QPageSize", "windowsId", args)
 }

}


func (this *QPageSize) sizePixels(args ...interface{}) () {
  // sizePixels(int)
  // sizePixels(enum QPageSize::PageSizeId, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10sizePixelsEi
  case 1:
    // invoke: _ZN9QPageSize10sizePixelsENS_10PageSizeIdEi
  default:
    qtrt.ErrorResolve("QPageSize", "sizePixels", args)
 }

}


func (this *QPageSize) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize7isValidEv
  default:
    qtrt.ErrorResolve("QPageSize", "isValid", args)
 }

}


func (this *QPageSize) rectPixels(args ...interface{}) () {
  // rectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPixelsEi
  default:
    qtrt.ErrorResolve("QPageSize", "rectPixels", args)
 }

}


func (this *QPageSize) rectPoints(args ...interface{}) () {
  // rectPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPointsEv
  default:
    qtrt.ErrorResolve("QPageSize", "rectPoints", args)
 }

}


func (this *QPageSize) isEquivalentTo(args ...interface{}) () {
  // isEquivalentTo(const class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14isEquivalentToERKS_
  default:
    qtrt.ErrorResolve("QPageSize", "isEquivalentTo", args)
 }

}


func (this *QPageSize) sizePoints(args ...interface{}) () {
  // sizePoints(enum QPageSize::PageSizeId)
  // sizePoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize10sizePointsENS_10PageSizeIdE
  case 1:
    // invoke: _ZNK9QPageSize10sizePointsEv
  default:
    qtrt.ErrorResolve("QPageSize", "sizePoints", args)
 }

}

// <= body block end

