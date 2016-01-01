package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qdebug.h
// dst-file: /src/core/qdebug.go
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
// class sizeof(QNoDebug)=1
type QNoDebug struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDebugStateSaver)=1
type QDebugStateSaver struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDebug)=8
type QDebug struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QNoDebug) maybeQuote(args ...interface{}) () {
  // maybeQuote(const char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "const char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug10maybeQuoteEc
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeQuote", args)
  }

}


func (this *QNoDebug) quote(args ...interface{}) () {
  // quote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug5quoteEv
  default:
    qtrt.ErrorResolve("QNoDebug", "quote", args)
  }

}


func (this *QNoDebug) space(args ...interface{}) () {
  // space()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug5spaceEv
  default:
    qtrt.ErrorResolve("QNoDebug", "space", args)
  }

}


func (this *QNoDebug) nospace(args ...interface{}) () {
  // nospace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug7nospaceEv
  default:
    qtrt.ErrorResolve("QNoDebug", "nospace", args)
  }

}


func (this *QNoDebug) noquote(args ...interface{}) () {
  // noquote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug7noquoteEv
  default:
    qtrt.ErrorResolve("QNoDebug", "noquote", args)
  }

}


func (this *QNoDebug) maybeSpace(args ...interface{}) () {
  // maybeSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QNoDebug10maybeSpaceEv
  default:
    qtrt.ErrorResolve("QNoDebug", "maybeSpace", args)
  }

}


func NewQDebugStateSaver(args ...interface{}) QDebugStateSaver {
  return QDebugStateSaver{}
}


func (this *QDebugStateSaver) FreeQDebugStateSaver(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDebugStateSaver", "~QDebugStateSaver", args)
  }

}


func (this *QDebug) noquote(args ...interface{}) () {
  // noquote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug7noquoteEv
  default:
    qtrt.ErrorResolve("QDebug", "noquote", args)
  }

}


func (this *QDebug) FreeQDebug(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDebug", "~QDebug", args)
  }

}


func NewQDebug(args ...interface{}) QDebug {
  return QDebug{}
}


func (this *QDebug) space(args ...interface{}) () {
  // space()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug5spaceEv
  default:
    qtrt.ErrorResolve("QDebug", "space", args)
  }

}


func (this *QDebug) maybeSpace(args ...interface{}) () {
  // maybeSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug10maybeSpaceEv
  default:
    qtrt.ErrorResolve("QDebug", "maybeSpace", args)
  }

}


func (this *QDebug) resetFormat(args ...interface{}) () {
  // resetFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug11resetFormatEv
  default:
    qtrt.ErrorResolve("QDebug", "resetFormat", args)
  }

}


func (this *QDebug) setAutoInsertSpaces(args ...interface{}) () {
  // setAutoInsertSpaces(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug19setAutoInsertSpacesEb
  default:
    qtrt.ErrorResolve("QDebug", "setAutoInsertSpaces", args)
  }

}


func (this *QDebug) swap(args ...interface{}) () {
  // swap(class QDebug &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDebug{}) // "QDebug &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug4swapERS_
  default:
    qtrt.ErrorResolve("QDebug", "swap", args)
  }

}


func (this *QDebug) nospace(args ...interface{}) () {
  // nospace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug7nospaceEv
  default:
    qtrt.ErrorResolve("QDebug", "nospace", args)
  }

}


func (this *QDebug) autoInsertSpaces(args ...interface{}) () {
  // autoInsertSpaces()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QDebug16autoInsertSpacesEv
  default:
    qtrt.ErrorResolve("QDebug", "autoInsertSpaces", args)
  }

}


func (this *QDebug) quote(args ...interface{}) () {
  // quote()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug5quoteEv
  default:
    qtrt.ErrorResolve("QDebug", "quote", args)
  }

}


func (this *QDebug) maybeQuote(args ...interface{}) () {
  // maybeQuote(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QDebug10maybeQuoteEc
  default:
    qtrt.ErrorResolve("QDebug", "maybeQuote", args)
  }

}

// <= body block end

