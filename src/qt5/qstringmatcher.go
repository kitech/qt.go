package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qstringmatcher.h
// dst-file: /src/core/qstringmatcher.go
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
// class sizeof(QStringMatcher)=1048
type QStringMatcher struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQStringMatcher(args ...interface{}) QStringMatcher {
  return QStringMatcher{}
}


func (this *QStringMatcher) indexIn(args ...interface{}) () {
  // indexIn(const class QChar *, int, int)
  // indexIn(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStringMatcher7indexInEPK5QCharii
  case 1:
    // invoke: _ZNK14QStringMatcher7indexInERK7QStringi
  default:
    qtrt.ErrorResolve("QStringMatcher", "indexIn", args)
  }

}


func (this *QStringMatcher) setPattern(args ...interface{}) () {
  // setPattern(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringMatcher10setPatternERK7QString
  default:
    qtrt.ErrorResolve("QStringMatcher", "setPattern", args)
  }

}


func (this *QStringMatcher) pattern(args ...interface{}) () {
  // pattern()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStringMatcher7patternEv
  default:
    qtrt.ErrorResolve("QStringMatcher", "pattern", args)
  }

}


func (this *QStringMatcher) FreeQStringMatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStringMatcher", "~QStringMatcher", args)
  }

}

// <= body block end

