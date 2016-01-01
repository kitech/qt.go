package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qregexp.h
// dst-file: /src/core/qregexp.go
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
// class sizeof(QRegExp)=8
type QRegExp struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQRegExp(args ...interface{}) QRegExp {
  return QRegExp{}
}


func (this *QRegExp) capturedTexts(args ...interface{}) () {
  // capturedTexts()
  // capturedTexts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp13capturedTextsEv
  case 1:
    // invoke: _ZNK7QRegExp13capturedTextsEv
  default:
    qtrt.ErrorResolve("QRegExp", "capturedTexts", args)
  }

}


func (this *QRegExp) captureCount(args ...interface{}) () {
  // captureCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp12captureCountEv
  default:
    qtrt.ErrorResolve("QRegExp", "captureCount", args)
  }

}


func (this *QRegExp) escape_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegExp", "escape", args)
  }

}


func (this *QRegExp) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isEmptyEv
  default:
    qtrt.ErrorResolve("QRegExp", "isEmpty", args)
  }

}


func (this *QRegExp) isMinimal(args ...interface{}) () {
  // isMinimal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp9isMinimalEv
  default:
    qtrt.ErrorResolve("QRegExp", "isMinimal", args)
  }

}


func (this *QRegExp) matchedLength(args ...interface{}) () {
  // matchedLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp13matchedLengthEv
  default:
    qtrt.ErrorResolve("QRegExp", "matchedLength", args)
  }

}


func (this *QRegExp) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7patternEv
  default:
    qtrt.ErrorResolve("QRegExp", "pattern", args)
  }

}


func (this *QRegExp) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setPatternERK7QString
  default:
    qtrt.ErrorResolve("QRegExp", "setPattern", args)
  }

}


func (this *QRegExp) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp7isValidEv
  default:
    qtrt.ErrorResolve("QRegExp", "isValid", args)
  }

}


func (this *QRegExp) FreeQRegExp(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegExp", "~QRegExp", args)
  }

}


func (this *QRegExp) exactMatch(args ...interface{}) () {
  // exactMatch(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp10exactMatchERK7QString
  default:
    qtrt.ErrorResolve("QRegExp", "exactMatch", args)
  }

}


func (this *QRegExp) swap(args ...interface{}) () {
  // swap(class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp4swapERS_
  default:
    qtrt.ErrorResolve("QRegExp", "swap", args)
  }

}


func (this *QRegExp) pos(args ...interface{}) () {
  // pos(int)
  // pos(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp3posEi
  case 1:
    // invoke: _ZNK7QRegExp3posEi
  default:
    qtrt.ErrorResolve("QRegExp", "pos", args)
  }

}


func (this *QRegExp) cap(args ...interface{}) () {
  // cap(int)
  // cap(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp3capEi
  case 1:
    // invoke: _ZN7QRegExp3capEi
  default:
    qtrt.ErrorResolve("QRegExp", "cap", args)
  }

}


func (this *QRegExp) errorString(args ...interface{}) () {
  // errorString()
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegExp11errorStringEv
  case 1:
    // invoke: _ZN7QRegExp11errorStringEv
  default:
    qtrt.ErrorResolve("QRegExp", "errorString", args)
  }

}


func (this *QRegExp) setMinimal(args ...interface{}) () {
  // setMinimal(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegExp10setMinimalEb
  default:
    qtrt.ErrorResolve("QRegExp", "setMinimal", args)
  }

}

// <= body block end

