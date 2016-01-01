package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qtextboundaryfinder.h
// dst-file: /src/core/qtextboundaryfinder.go
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
// class sizeof(QTextBoundaryFinder)=48
type QTextBoundaryFinder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextBoundaryFinder) isAtBoundary(args ...interface{}) () {
  // isAtBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder12isAtBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isAtBoundary", args)
 }

}


func (this *QTextBoundaryFinder) toNextBoundary(args ...interface{}) () {
  // toNextBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder14toNextBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toNextBoundary", args)
 }

}


func (this *QTextBoundaryFinder) toEnd(args ...interface{}) () {
  // toEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder5toEndEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toEnd", args)
 }

}


func NewQTextBoundaryFinder(args ...interface{})() {
}


func (this *QTextBoundaryFinder) setPosition(args ...interface{}) () {
  // setPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder11setPositionEi
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "setPosition", args)
 }

}


func (this *QTextBoundaryFinder) toPreviousBoundary(args ...interface{}) () {
  // toPreviousBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder18toPreviousBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toPreviousBoundary", args)
 }

}


func (this *QTextBoundaryFinder) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder7isValidEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isValid", args)
 }

}


func (this *QTextBoundaryFinder) FreeQTextBoundaryFinder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "~QTextBoundaryFinder", args)
 }

}


func (this *QTextBoundaryFinder) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder6stringEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "string", args)
 }

}


func (this *QTextBoundaryFinder) toStart(args ...interface{}) () {
  // toStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder7toStartEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toStart", args)
 }

}


func (this *QTextBoundaryFinder) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder8positionEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "position", args)
 }

}

// <= body block end

