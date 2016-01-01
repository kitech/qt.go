package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qchar.h
// dst-file: /src/core/qchar.go
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
// class sizeof(QLatin1Char)=1
type QLatin1Char struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QChar)=2
type QChar struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QLatin1Char) unicode(args ...interface{}) () {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLatin1Char7unicodeEv
  default:
    qtrt.ErrorResolve("QLatin1Char", "unicode", args)
  }

}


func NewQLatin1Char(args ...interface{}) QLatin1Char {
  return QLatin1Char{}
}


func (this *QLatin1Char) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLatin1Char8toLatin1Ev
  default:
    qtrt.ErrorResolve("QLatin1Char", "toLatin1", args)
  }

}


func (this *QChar) toUpper_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

}


func (this *QChar) hasMirrored(args ...interface{}) () {
  // hasMirrored()
  // hasMirrored(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11hasMirroredEv
  case 1:
    // invoke: _ZN5QChar11hasMirroredEj
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

}


func (this *QChar) lowSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "lowSurrogate", args)
  }

}


func (this *QChar) isSymbol_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

}


func (this *QChar) cell(args ...interface{}) () {
  // cell()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar4cellEv
  default:
    qtrt.ErrorResolve("QChar", "cell", args)
  }

}


func (this *QChar) surrogateToUcs4_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "surrogateToUcs4", args)
  }

}


func (this *QChar) isTitleCase_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

}


func (this *QChar) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar6isNullEv
  default:
    qtrt.ErrorResolve("QChar", "isNull", args)
  }

}


func (this *QChar) digitValue(args ...interface{}) () {
  // digitValue()
  // digitValue(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar10digitValueEv
  case 1:
    // invoke: _ZN5QChar10digitValueEj
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

}


func (this *QChar) toTitleCase(args ...interface{}) () {
  // toTitleCase()
  // toTitleCase(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11toTitleCaseEv
  case 1:
    // invoke: _ZN5QChar11toTitleCaseEj
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

}


func (this *QChar) isLower_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

}


func (this *QChar) isLowSurrogate(args ...interface{}) () {
  // isLowSurrogate()
  // isLowSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14isLowSurrogateEv
  case 1:
    // invoke: _ZN5QChar14isLowSurrogateEj
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

}


func (this *QChar) isPrint_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

}


func (this *QChar) isSymbol(args ...interface{}) () {
  // isSymbol(uint)
  // isSymbol()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar8isSymbolEj
  case 1:
    // invoke: _ZNK5QChar8isSymbolEv
  default:
    qtrt.ErrorResolve("QChar", "isSymbol", args)
  }

}


func (this *QChar) isLower(args ...interface{}) () {
  // isLower(uint)
  // isLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isLowerEj
  case 1:
    // invoke: _ZNK5QChar7isLowerEv
  default:
    qtrt.ErrorResolve("QChar", "isLower", args)
  }

}


func (this *QChar) mirroredChar(args ...interface{}) () {
  // mirroredChar()
  // mirroredChar(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar12mirroredCharEv
  case 1:
    // invoke: _ZN5QChar12mirroredCharEj
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

}


func (this *QChar) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar3rowEv
  default:
    qtrt.ErrorResolve("QChar", "row", args)
  }

}


func (this *QChar) isDigit(args ...interface{}) () {
  // isDigit()
  // isDigit(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isDigitEv
  case 1:
    // invoke: _ZN5QChar7isDigitEj
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

}


func (this *QChar) toTitleCase_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toTitleCase", args)
  }

}


func (this *QChar) isSurrogate(args ...interface{}) () {
  // isSurrogate()
  // isSurrogate(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar11isSurrogateEv
  case 1:
    // invoke: _ZN5QChar11isSurrogateEj
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

}


func (this *QChar) hasMirrored_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "hasMirrored", args)
  }

}


func (this *QChar) isNumber(args ...interface{}) () {
  // isNumber()
  // isNumber(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8isNumberEv
  case 1:
    // invoke: _ZN5QChar8isNumberEj
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

}


func (this *QChar) isHighSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

}


func (this *QChar) toCaseFolded_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

}


func (this *QChar) isMark(args ...interface{}) () {
  // isMark()
  // isMark(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar6isMarkEv
  case 1:
    // invoke: _ZN5QChar6isMarkEj
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

}


func (this *QChar) toLower_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

}


func (this *QChar) mirroredChar_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "mirroredChar", args)
  }

}


func (this *QChar) setRow(args ...interface{}) () {
  // setRow(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar6setRowEh
  default:
    qtrt.ErrorResolve("QChar", "setRow", args)
  }

}


func (this *QChar) decomposition_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

}


func (this *QChar) digitValue_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "digitValue", args)
  }

}


func (this *QChar) setCell(args ...interface{}) () {
  // setCell(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7setCellEh
  default:
    qtrt.ErrorResolve("QChar", "setCell", args)
  }

}


func (this *QChar) isUpper_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

}


func NewQChar(args ...interface{}) QChar {
  return QChar{}
}


func (this *QChar) toCaseFolded(args ...interface{}) () {
  // toCaseFolded(uint)
  // toCaseFolded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar12toCaseFoldedEj
  case 1:
    // invoke: _ZNK5QChar12toCaseFoldedEv
  default:
    qtrt.ErrorResolve("QChar", "toCaseFolded", args)
  }

}


func (this *QChar) isPrint(args ...interface{}) () {
  // isPrint(uint)
  // isPrint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isPrintEj
  case 1:
    // invoke: _ZNK5QChar7isPrintEv
  default:
    qtrt.ErrorResolve("QChar", "isPrint", args)
  }

}


func (this *QChar) isPunct(args ...interface{}) () {
  // isPunct()
  // isPunct(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isPunctEv
  case 1:
    // invoke: _ZN5QChar7isPunctEj
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

}


func (this *QChar) decomposition(args ...interface{}) () {
  // decomposition(uint)
  // decomposition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar13decompositionEj
  case 1:
    // invoke: _ZNK5QChar13decompositionEv
  default:
    qtrt.ErrorResolve("QChar", "decomposition", args)
  }

}


func (this *QChar) isSpace(args ...interface{}) () {
  // isSpace()
  // isSpace(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7isSpaceEv
  case 1:
    // invoke: _ZN5QChar7isSpaceEj
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

}


func (this *QChar) isUpper(args ...interface{}) () {
  // isUpper(uint)
  // isUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7isUpperEj
  case 1:
    // invoke: _ZNK5QChar7isUpperEv
  default:
    qtrt.ErrorResolve("QChar", "isUpper", args)
  }

}


func (this *QChar) combiningClass_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

}


func (this *QChar) isNonCharacter(args ...interface{}) () {
  // isNonCharacter()
  // isNonCharacter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar14isNonCharacterEv
  case 1:
    // invoke: _ZN5QChar14isNonCharacterEj
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

}


func (this *QChar) isLetterOrNumber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

}


func (this *QChar) isDigit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isDigit", args)
  }

}


func (this *QChar) isPunct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isPunct", args)
  }

}


func (this *QChar) isTitleCase(args ...interface{}) () {
  // isTitleCase(uint)
  // isTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar11isTitleCaseEj
  case 1:
    // invoke: _ZNK5QChar11isTitleCaseEv
  default:
    qtrt.ErrorResolve("QChar", "isTitleCase", args)
  }

}


func (this *QChar) isLetter(args ...interface{}) () {
  // isLetter()
  // isLetter(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8isLetterEv
  case 1:
    // invoke: _ZN5QChar8isLetterEj
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

}


func (this *QChar) combiningClass(args ...interface{}) () {
  // combiningClass(uint)
  // combiningClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar14combiningClassEj
  case 1:
    // invoke: _ZNK5QChar14combiningClassEv
  default:
    qtrt.ErrorResolve("QChar", "combiningClass", args)
  }

}


func (this *QChar) isHighSurrogate(args ...interface{}) () {
  // isHighSurrogate(uint)
  // isHighSurrogate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar15isHighSurrogateEj
  case 1:
    // invoke: _ZNK5QChar15isHighSurrogateEv
  default:
    qtrt.ErrorResolve("QChar", "isHighSurrogate", args)
  }

}


func (this *QChar) highSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "highSurrogate", args)
  }

}


func (this *QChar) requiresSurrogates_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "requiresSurrogates", args)
  }

}


func (this *QChar) isLetterOrNumber(args ...interface{}) () {
  // isLetterOrNumber(uint)
  // isLetterOrNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar16isLetterOrNumberEj
  case 1:
    // invoke: _ZNK5QChar16isLetterOrNumberEv
  default:
    qtrt.ErrorResolve("QChar", "isLetterOrNumber", args)
  }

}


func (this *QChar) unicode(args ...interface{}) () {
  // unicode()
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar7unicodeEv
  case 1:
    // invoke: _ZN5QChar7unicodeEv
  default:
    qtrt.ErrorResolve("QChar", "unicode", args)
  }

}


func (this *QChar) isLowSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLowSurrogate", args)
  }

}


func (this *QChar) isNumber_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isNumber", args)
  }

}


func (this *QChar) toLower(args ...interface{}) () {
  // toLower(uint)
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7toLowerEj
  case 1:
    // invoke: _ZNK5QChar7toLowerEv
  default:
    qtrt.ErrorResolve("QChar", "toLower", args)
  }

}


func (this *QChar) isLetter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isLetter", args)
  }

}


func (this *QChar) toUpper(args ...interface{}) () {
  // toUpper(uint)
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QChar7toUpperEj
  case 1:
    // invoke: _ZNK5QChar7toUpperEv
  default:
    qtrt.ErrorResolve("QChar", "toUpper", args)
  }

}


func (this *QChar) isSpace_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSpace", args)
  }

}


func (this *QChar) fromLatin1_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "fromLatin1", args)
  }

}


func (this *QChar) isSurrogate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isSurrogate", args)
  }

}


func (this *QChar) isMark_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isMark", args)
  }

}


func (this *QChar) isNonCharacter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChar", "isNonCharacter", args)
  }

}


func (this *QChar) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QChar8toLatin1Ev
  default:
    qtrt.ErrorResolve("QChar", "toLatin1", args)
  }

}

// <= body block end

