package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qloggingcategory.h
// dst-file: /src/core/qloggingcategory.go
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
// class sizeof(QLoggingCategory)=24
type QLoggingCategory struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQLoggingCategory(args ...interface{})() {
}


func (this *QLoggingCategory) isDebugEnabled(args ...interface{}) () {
  // isDebugEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory14isDebugEnabledEv
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isDebugEnabled", args)
 }

}


func (this *QLoggingCategory) FreeQLoggingCategory(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "~QLoggingCategory", args)
 }

}


func (this *QLoggingCategory) setEnabled(args ...interface{}) () {
  // setEnabled(enum QtMsgType, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QtMsgType"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategory10setEnabledE9QtMsgTypeb
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setEnabled", args)
 }

}


func (this *QLoggingCategory) isEnabled(args ...interface{}) () {
  // isEnabled(enum QtMsgType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QtMsgType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory9isEnabledE9QtMsgType
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isEnabled", args)
 }

}


func (this *QLoggingCategory) isWarningEnabled(args ...interface{}) () {
  // isWarningEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory16isWarningEnabledEv
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isWarningEnabled", args)
 }

}


func (this *QLoggingCategory) isInfoEnabled(args ...interface{}) () {
  // isInfoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory13isInfoEnabledEv
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isInfoEnabled", args)
 }

}


func (this *QLoggingCategory) categoryName(args ...interface{}) () {
  // categoryName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory12categoryNameEv
  default:
    qtrt.ErrorResolve("QLoggingCategory", "categoryName", args)
 }

}


func (this *QLoggingCategory) isCriticalEnabled(args ...interface{}) () {
  // isCriticalEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory17isCriticalEnabledEv
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isCriticalEnabled", args)
 }

}


func (this *QLoggingCategory) defaultCategory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "defaultCategory", args)
 }

}


func (this *QLoggingCategory) setFilterRules_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setFilterRules", args)
 }

}

// <= body block end

