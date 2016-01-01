package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qtextoption.h
// dst-file: /src/gui/qtextoption.go
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
// class sizeof(QTextOption)=32
type QTextOption struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQTextOption(args ...interface{})() {
}


func (this *QTextOption) tabStop(args ...interface{}) () {
  // tabStop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption7tabStopEv
  default:
    qtrt.ErrorResolve("QTextOption", "tabStop", args)
 }

}


func (this *QTextOption) setUseDesignMetrics(args ...interface{}) () {
  // setUseDesignMetrics(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextOption19setUseDesignMetricsEb
  default:
    qtrt.ErrorResolve("QTextOption", "setUseDesignMetrics", args)
 }

}


func (this *QTextOption) setTabStop(args ...interface{}) () {
  // setTabStop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextOption10setTabStopEd
  default:
    qtrt.ErrorResolve("QTextOption", "setTabStop", args)
 }

}


func (this *QTextOption) useDesignMetrics(args ...interface{}) () {
  // useDesignMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption16useDesignMetricsEv
  default:
    qtrt.ErrorResolve("QTextOption", "useDesignMetrics", args)
 }

}


func (this *QTextOption) tabArray(args ...interface{}) () {
  // tabArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption8tabArrayEv
  default:
    qtrt.ErrorResolve("QTextOption", "tabArray", args)
 }

}


func (this *QTextOption) FreeQTextOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextOption", "~QTextOption", args)
 }

}

// <= body block end

