package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qtextlist.h
// dst-file: /src/gui/qtextlist.go
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
// class sizeof(QTextList)=1
type QTextList struct {
  /*qbase*/ QTextBlockGroup;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextList) item(args ...interface{}) () {
  // item(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList4itemEi
  default:
    qtrt.ErrorResolve("QTextList", "item", args)
  }

}


func (this *QTextList) remove(args ...interface{}) () {
  // remove(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList6removeERK10QTextBlock
  default:
    qtrt.ErrorResolve("QTextList", "remove", args)
  }

}


func (this *QTextList) setFormat(args ...interface{}) () {
  // setFormat(const class QTextListFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList9setFormatERK15QTextListFormat
  default:
    qtrt.ErrorResolve("QTextList", "setFormat", args)
  }

}


func NewQTextList(args ...interface{}) QTextList {
  return QTextList{}
}


func (this *QTextList) add(args ...interface{}) () {
  // add(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList3addERK10QTextBlock
  default:
    qtrt.ErrorResolve("QTextList", "add", args)
  }

}


func (this *QTextList) itemText(args ...interface{}) () {
  // itemText(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList8itemTextERK10QTextBlock
  default:
    qtrt.ErrorResolve("QTextList", "itemText", args)
  }

}


func (this *QTextList) removeItem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList10removeItemEi
  default:
    qtrt.ErrorResolve("QTextList", "removeItem", args)
  }

}


func (this *QTextList) itemNumber(args ...interface{}) () {
  // itemNumber(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList10itemNumberERK10QTextBlock
  default:
    qtrt.ErrorResolve("QTextList", "itemNumber", args)
  }

}


func (this *QTextList) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList5countEv
  default:
    qtrt.ErrorResolve("QTextList", "count", args)
  }

}


func (this *QTextList) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList6formatEv
  default:
    qtrt.ErrorResolve("QTextList", "format", args)
  }

}


func (this *QTextList) FreeQTextList(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextList", "~QTextList", args)
  }

}


func (this *QTextList) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList7isEmptyEv
  default:
    qtrt.ErrorResolve("QTextList", "isEmpty", args)
  }

}


func (this *QTextList) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList10metaObjectEv
  default:
    qtrt.ErrorResolve("QTextList", "metaObject", args)
  }

}

// <= body block end

