package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qmimetype.h
// dst-file: /src/core/qmimetype.go
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
// class sizeof(QMimeType)=1
type QMimeType struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QMimeType) FreeQMimeType(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMimeType", "~QMimeType", args)
  }

}


func (this *QMimeType) comment(args ...interface{}) () {
  // comment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7commentEv
  default:
    qtrt.ErrorResolve("QMimeType", "comment", args)
  }

}


func (this *QMimeType) aliases(args ...interface{}) () {
  // aliases()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7aliasesEv
  default:
    qtrt.ErrorResolve("QMimeType", "aliases", args)
  }

}


func (this *QMimeType) filterString(args ...interface{}) () {
  // filterString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12filterStringEv
  default:
    qtrt.ErrorResolve("QMimeType", "filterString", args)
  }

}


func (this *QMimeType) parentMimeTypes(args ...interface{}) () {
  // parentMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15parentMimeTypesEv
  default:
    qtrt.ErrorResolve("QMimeType", "parentMimeTypes", args)
  }

}


func NewQMimeType(args ...interface{}) QMimeType {
  return QMimeType{}
}


func (this *QMimeType) inherits(args ...interface{}) () {
  // inherits(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8inheritsERK7QString
  default:
    qtrt.ErrorResolve("QMimeType", "inherits", args)
  }

}


func (this *QMimeType) isDefault(args ...interface{}) () {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType9isDefaultEv
  default:
    qtrt.ErrorResolve("QMimeType", "isDefault", args)
  }

}


func (this *QMimeType) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7isValidEv
  default:
    qtrt.ErrorResolve("QMimeType", "isValid", args)
  }

}


func (this *QMimeType) swap(args ...interface{}) () {
  // swap(class QMimeType &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeType{}) // "QMimeType &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeType4swapERS_
  default:
    qtrt.ErrorResolve("QMimeType", "swap", args)
  }

}


func (this *QMimeType) suffixes(args ...interface{}) () {
  // suffixes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8suffixesEv
  default:
    qtrt.ErrorResolve("QMimeType", "suffixes", args)
  }

}


func (this *QMimeType) genericIconName(args ...interface{}) () {
  // genericIconName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15genericIconNameEv
  default:
    qtrt.ErrorResolve("QMimeType", "genericIconName", args)
  }

}


func (this *QMimeType) iconName(args ...interface{}) () {
  // iconName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8iconNameEv
  default:
    qtrt.ErrorResolve("QMimeType", "iconName", args)
  }

}


func (this *QMimeType) allAncestors(args ...interface{}) () {
  // allAncestors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12allAncestorsEv
  default:
    qtrt.ErrorResolve("QMimeType", "allAncestors", args)
  }

}


func (this *QMimeType) globPatterns(args ...interface{}) () {
  // globPatterns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12globPatternsEv
  default:
    qtrt.ErrorResolve("QMimeType", "globPatterns", args)
  }

}


func (this *QMimeType) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType4nameEv
  default:
    qtrt.ErrorResolve("QMimeType", "name", args)
  }

}


func (this *QMimeType) preferredSuffix(args ...interface{}) () {
  // preferredSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15preferredSuffixEv
  default:
    qtrt.ErrorResolve("QMimeType", "preferredSuffix", args)
  }

}

// <= body block end

