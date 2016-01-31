package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qanimationgroup.h
// dst-file: /src/core/qanimationgroup.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QAnimationGroup::animationCount();
extern void C_ZNK15QAnimationGroup14animationCountEv(void* qthis); // 4
  // proto:  void QAnimationGroup::addAnimation(QAbstractAnimation * animation);
extern void C_ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  int QAnimationGroup::indexOfAnimation(QAbstractAnimation * animation);
extern void C_ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  void QAnimationGroup::~QAnimationGroup();
extern void C_ZN15QAnimationGroupD2Ev(void* qthis); // 4
  // proto:  void QAnimationGroup::clear();
extern void C_ZN15QAnimationGroup5clearEv(void* qthis); // 4
  // proto:  QAbstractAnimation * QAnimationGroup::animationAt(int index);
extern void C_ZNK15QAnimationGroup11animationAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QAnimationGroup::insertAnimation(int index, QAbstractAnimation * animation);
extern void C_ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QAnimationGroup::metaObject();
extern void C_ZNK15QAnimationGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QAnimationGroup::removeAnimation(QAbstractAnimation * animation);
extern void C_ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation(void* qthis, void* arg0); // 4
  // proto:  QAbstractAnimation * QAnimationGroup::takeAnimation(int index);
extern void C_ZN15QAnimationGroup13takeAnimationEi(void* qthis, int32_t arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QAnimationGroup)=1
type QAnimationGroup struct {
  /*qbase*/ QAbstractAnimation;
  qclsinst unsafe.Pointer /* *C.void */;
}

// animationCount()
func (this *QAnimationGroup) animationCount(args ...interface{}) () {
  // animationCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAnimationGroup14animationCountEv
    // invoke: int animationCount()
    var ret = C.C_ZNK15QAnimationGroup14animationCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAnimationGroup", "animationCount", args)
  }

}

// addAnimation(class QAbstractAnimation *)
func (this *QAnimationGroup) addAnimation(args ...interface{}) () {
  // addAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation
    // invoke: void addAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "addAnimation", args)
  }

}

// indexOfAnimation(class QAbstractAnimation *)
func (this *QAnimationGroup) indexOfAnimation(args ...interface{}) () {
  // indexOfAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation
    // invoke: int indexOfAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QAnimationGroup", "indexOfAnimation", args)
  }

}

// ~QAnimationGroup()
func (this *QAnimationGroup) FreeQAnimationGroup(args ...interface{}) () {
  // ~QAnimationGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroupD0Ev
    // invoke: void ~QAnimationGroup()
    C.C_ZN15QAnimationGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "~QAnimationGroup", args)
  }

}

// clear()
func (this *QAnimationGroup) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroup5clearEv
    // invoke: void clear()
    C.C_ZN15QAnimationGroup5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "clear", args)
  }

}

// animationAt(int)
func (this *QAnimationGroup) animationAt(args ...interface{}) () {
  // animationAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAnimationGroup11animationAtEi
    // invoke: QAbstractAnimation * animationAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK15QAnimationGroup11animationAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "animationAt", args)
  }

}

// insertAnimation(int, class QAbstractAnimation *)
func (this *QAnimationGroup) insertAnimation(args ...interface{}) () {
  // insertAnimation(int, class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation
    // invoke: void insertAnimation(int, class QAbstractAnimation *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "insertAnimation", args)
  }

}

// metaObject()
func (this *QAnimationGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAnimationGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QAnimationGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "metaObject", args)
  }

}

// removeAnimation(class QAbstractAnimation *)
func (this *QAnimationGroup) removeAnimation(args ...interface{}) () {
  // removeAnimation(class QAbstractAnimation *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractAnimation{}) // "QAbstractAnimation *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation
    // invoke: void removeAnimation(class QAbstractAnimation *)
    var arg0 = args[0].(QAbstractAnimation).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "removeAnimation", args)
  }

}

// takeAnimation(int)
func (this *QAnimationGroup) takeAnimation(args ...interface{}) () {
  // takeAnimation(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAnimationGroup13takeAnimationEi
    // invoke: QAbstractAnimation * takeAnimation(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QAnimationGroup13takeAnimationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "takeAnimation", args)
  }

}

// <= body block end

