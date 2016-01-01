package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qregularexpression.h
// dst-file: /src/core/qregularexpression.go
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
// class sizeof(QRegularExpressionMatchIterator)=1
type QRegularExpressionMatchIterator struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QRegularExpression)=1
type QRegularExpression struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QRegularExpressionMatch)=1
type QRegularExpressionMatch struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QRegularExpressionMatchIterator) hasNext(args ...interface{}) () {
  // hasNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator7hasNextEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "hasNext", args)
 }

}


func (this *QRegularExpressionMatchIterator) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator7isValidEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "isValid", args)
 }

}


func (this *QRegularExpressionMatchIterator) peekNext(args ...interface{}) () {
  // peekNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator8peekNextEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "peekNext", args)
 }

}


func NewQRegularExpressionMatchIterator(args ...interface{})() {
}


func (this *QRegularExpressionMatchIterator) regularExpression(args ...interface{}) () {
  // regularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK31QRegularExpressionMatchIterator17regularExpressionEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "regularExpression", args)
 }

}


func (this *QRegularExpressionMatchIterator) FreeQRegularExpressionMatchIterator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "~QRegularExpressionMatchIterator", args)
 }

}


func (this *QRegularExpressionMatchIterator) next(args ...interface{}) () {
  // next()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIterator4nextEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "next", args)
 }

}


func (this *QRegularExpressionMatchIterator) swap(args ...interface{}) () {
  // swap(class QRegularExpressionMatchIterator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatchIterator{}) // "QRegularExpressionMatchIterator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN31QRegularExpressionMatchIterator4swapERS_
  default:
    qtrt.ErrorResolve("QRegularExpressionMatchIterator", "swap", args)
 }

}


func (this *QRegularExpression) patternErrorOffset(args ...interface{}) () {
  // patternErrorOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression18patternErrorOffsetEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "patternErrorOffset", args)
 }

}


func (this *QRegularExpression) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression7patternEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "pattern", args)
 }

}


func (this *QRegularExpression) FreeQRegularExpression(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpression", "~QRegularExpression", args)
 }

}


func (this *QRegularExpression) optimize(args ...interface{}) () {
  // optimize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression8optimizeEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "optimize", args)
 }

}


func (this *QRegularExpression) escape_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpression", "escape", args)
 }

}


func NewQRegularExpression(args ...interface{})() {
}


func (this *QRegularExpression) swap(args ...interface{}) () {
  // swap(class QRegularExpression &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "QRegularExpression &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpression4swapERS_
  default:
    qtrt.ErrorResolve("QRegularExpression", "swap", args)
 }

}


func (this *QRegularExpression) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression11errorStringEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "errorString", args)
 }

}


func (this *QRegularExpression) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression7isValidEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "isValid", args)
 }

}


func (this *QRegularExpression) namedCaptureGroups(args ...interface{}) () {
  // namedCaptureGroups()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression18namedCaptureGroupsEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "namedCaptureGroups", args)
 }

}


func (this *QRegularExpression) captureCount(args ...interface{}) () {
  // captureCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QRegularExpression12captureCountEv
  default:
    qtrt.ErrorResolve("QRegularExpression", "captureCount", args)
 }

}


func (this *QRegularExpression) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QRegularExpression10setPatternERK7QString
  default:
    qtrt.ErrorResolve("QRegularExpression", "setPattern", args)
 }

}


func (this *QRegularExpressionMatch) lastCapturedIndex(args ...interface{}) () {
  // lastCapturedIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch17lastCapturedIndexEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "lastCapturedIndex", args)
 }

}


func NewQRegularExpressionMatch(args ...interface{})() {
}


func (this *QRegularExpressionMatch) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch7isValidEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "isValid", args)
 }

}


func (this *QRegularExpressionMatch) capturedLength(args ...interface{}) () {
  // capturedLength(int)
  // capturedLength(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch14capturedLengthEi
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch14capturedLengthERK7QString
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedLength", args)
 }

}


func (this *QRegularExpressionMatch) capturedRef(args ...interface{}) () {
  // capturedRef(int)
  // capturedRef(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch11capturedRefEi
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedRefERK7QString
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedRef", args)
 }

}


func (this *QRegularExpressionMatch) capturedEnd(args ...interface{}) () {
  // capturedEnd(const class QString &)
  // capturedEnd(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch11capturedEndERK7QString
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch11capturedEndEi
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedEnd", args)
 }

}


func (this *QRegularExpressionMatch) captured(args ...interface{}) () {
  // captured(const class QString &)
  // captured(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch8capturedERK7QString
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch8capturedEi
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "captured", args)
 }

}


func (this *QRegularExpressionMatch) capturedTexts(args ...interface{}) () {
  // capturedTexts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch13capturedTextsEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedTexts", args)
 }

}


func (this *QRegularExpressionMatch) swap(args ...interface{}) () {
  // swap(class QRegularExpressionMatch &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN23QRegularExpressionMatch4swapERS_
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "swap", args)
 }

}


func (this *QRegularExpressionMatch) FreeQRegularExpressionMatch(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "~QRegularExpressionMatch", args)
 }

}


func (this *QRegularExpressionMatch) hasMatch(args ...interface{}) () {
  // hasMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch8hasMatchEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasMatch", args)
 }

}


func (this *QRegularExpressionMatch) capturedStart(args ...interface{}) () {
  // capturedStart(const class QString &)
  // capturedStart(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch13capturedStartERK7QString
  case 1:
    // invoke: _ZNK23QRegularExpressionMatch13capturedStartEi
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "capturedStart", args)
 }

}


func (this *QRegularExpressionMatch) regularExpression(args ...interface{}) () {
  // regularExpression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch17regularExpressionEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "regularExpression", args)
 }

}


func (this *QRegularExpressionMatch) hasPartialMatch(args ...interface{}) () {
  // hasPartialMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QRegularExpressionMatch15hasPartialMatchEv
  default:
    qtrt.ErrorResolve("QRegularExpressionMatch", "hasPartialMatch", args)
 }

}

// <= body block end

