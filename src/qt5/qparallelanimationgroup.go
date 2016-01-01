package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qparallelanimationgroup.h
// dst-file: /src/core/qparallelanimationgroup.go
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
// class sizeof(QParallelAnimationGroup)=1
type QParallelAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QParallelAnimationGroup) FreeQParallelAnimationGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "~QParallelAnimationGroup", args)
 }

}


func NewQParallelAnimationGroup(args ...interface{})() {
}


func (this *QParallelAnimationGroup) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QParallelAnimationGroup8durationEv
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "duration", args)
 }

}


func (this *QParallelAnimationGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QParallelAnimationGroup10metaObjectEv
  default:
    qtrt.ErrorResolve("QParallelAnimationGroup", "metaObject", args)
 }

}

// <= body block end

