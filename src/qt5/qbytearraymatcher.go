package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qbytearraymatcher.h
// dst-file: /src/core/qbytearraymatcher.go
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
// class sizeof(QByteArrayMatcher)=1040
type QByteArrayMatcher struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QByteArrayMatcher) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QByteArrayMatcher7patternEv
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "pattern", args)
  }

}


func (this *QByteArrayMatcher) indexIn(args ...interface{}) () {
  // indexIn(const char *, int, int)
  // indexIn(const class QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QByteArrayMatcher7indexInEPKcii
  case 1:
    // invoke: _ZNK17QByteArrayMatcher7indexInERK10QByteArrayi
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "indexIn", args)
  }

}


func (this *QByteArrayMatcher) setPattern(args ...interface{}) () {
  // setPattern(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QByteArrayMatcher10setPatternERK10QByteArray
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "setPattern", args)
  }

}


func NewQByteArrayMatcher(args ...interface{}) QByteArrayMatcher {
  return QByteArrayMatcher{}
}


func (this *QByteArrayMatcher) FreeQByteArrayMatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArrayMatcher", "~QByteArrayMatcher", args)
  }

}

// <= body block end

