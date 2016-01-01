package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qlist.h
// dst-file: /src/core/qlist.go
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
// class sizeof(QListData)=8
type QListData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QListData) prepend(args ...interface{}) () {
  // prepend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData7prependEv
  default:
    qtrt.ErrorResolve("QListData", "prepend", args)
 }

}


func (this *QListData) realloc(args ...interface{}) () {
  // realloc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData7reallocEi
  default:
    qtrt.ErrorResolve("QListData", "realloc", args)
 }

}


func (this *QListData) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QListData3endEv
  default:
    qtrt.ErrorResolve("QListData", "end", args)
 }

}


func (this *QListData) remove(args ...interface{}) () {
  // remove(int, int)
  // remove(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData6removeEii
  case 1:
    // invoke: _ZN9QListData6removeEi
  default:
    qtrt.ErrorResolve("QListData", "remove", args)
 }

}


func (this *QListData) append(args ...interface{}) () {
  // append(const struct QListData &)
  // append()
  // append(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListData{}) // "const QListData &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData6appendERKS_
  case 1:
    // invoke: _ZN9QListData6appendEv
  case 2:
    // invoke: _ZN9QListData6appendEi
  default:
    qtrt.ErrorResolve("QListData", "append", args)
 }

}


func (this *QListData) dispose(args ...interface{}) () {
  // dispose(struct QListData::Data *)
  // dispose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  // vtys[0][0] = reflect.TypeOf(QListData::Data{}) // "QListData::Data *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData7disposeEPNS_4DataE
  case 1:
    // invoke: _ZN9QListData7disposeEv
  default:
    qtrt.ErrorResolve("QListData", "dispose", args)
 }

}


func (this *QListData) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QListData4sizeEv
  default:
    qtrt.ErrorResolve("QListData", "size", args)
 }

}


func (this *QListData) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QListData7isEmptyEv
  default:
    qtrt.ErrorResolve("QListData", "isEmpty", args)
 }

}


func (this *QListData) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QListData2atEi
  default:
    qtrt.ErrorResolve("QListData", "at", args)
 }

}


func (this *QListData) erase(args ...interface{}) () {
  // erase(void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData5eraseEPPv
  default:
    qtrt.ErrorResolve("QListData", "erase", args)
 }

}


func (this *QListData) move_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QListData", "move", args)
 }

}


func (this *QListData) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QListData5beginEv
  default:
    qtrt.ErrorResolve("QListData", "begin", args)
 }

}


func (this *QListData) insert(args ...interface{}) () {
  // insert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QListData6insertEi
  default:
    qtrt.ErrorResolve("QListData", "insert", args)
 }

}

// <= body block end

