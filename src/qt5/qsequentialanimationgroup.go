package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qsequentialanimationgroup.h
// dst-file: /src/core/qsequentialanimationgroup.go
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
// class sizeof(QSequentialAnimationGroup)=1
type QSequentialAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst uint64 /* *mut c_void*/;
//  _currentAnimationChanged QSequentialAnimationGroup_currentAnimationChanged_signal;
}


func (this *QSequentialAnimationGroup) insertPause(args ...interface{}) () {
  // insertPause(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup11insertPauseEii
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "insertPause", args)
  }

}


func NewQSequentialAnimationGroup(args ...interface{}) QSequentialAnimationGroup {
  return QSequentialAnimationGroup{}
}


func (this *QSequentialAnimationGroup) addPause(args ...interface{}) () {
  // addPause(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup8addPauseEi
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "addPause", args)
  }

}


func (this *QSequentialAnimationGroup) FreeQSequentialAnimationGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "~QSequentialAnimationGroup", args)
  }

}


func (this *QSequentialAnimationGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup10metaObjectEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "metaObject", args)
  }

}


func (this *QSequentialAnimationGroup) currentAnimation(args ...interface{}) () {
  // currentAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup16currentAnimationEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "currentAnimation", args)
  }

}


func (this *QSequentialAnimationGroup) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup8durationEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "duration", args)
  }

}

// <= body block end

