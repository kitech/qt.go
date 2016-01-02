package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAbstractAnimation * QAnimationGroup::animationAt(int index);
extern void _ZNK15QAnimationGroup11animationAtEi(void* qthis, int arg0);
  // proto:  void QAnimationGroup::~QAnimationGroup();
extern void _ZN15QAnimationGroupD0Ev(void* qthis);
  // proto:  void QAnimationGroup::QAnimationGroup(const QAnimationGroup & );
extern void* dector_ZN15QAnimationGroupC1ERKS_(void* arg0);
extern void _ZN15QAnimationGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAnimationGroup::removeAnimation(QAbstractAnimation * animation);
extern void _ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  void QAnimationGroup::QAnimationGroup(QObject * parent);
extern void* dector_ZN15QAnimationGroupC1EP7QObject(void* arg0);
extern void _ZN15QAnimationGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  int QAnimationGroup::animationCount();
extern void _ZNK15QAnimationGroup14animationCountEv(void* qthis);
  // proto:  void QAnimationGroup::addAnimation(QAbstractAnimation * animation);
extern void _ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
  // proto:  void QAnimationGroup::clear();
extern void _ZN15QAnimationGroup5clearEv(void* qthis);
  // proto:  QAbstractAnimation * QAnimationGroup::takeAnimation(int index);
extern void _ZN15QAnimationGroup13takeAnimationEi(void* qthis, int arg0);
  // proto:  void QAnimationGroup::insertAnimation(int index, QAbstractAnimation * animation);
extern void _ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation(void* qthis, int arg0, void* arg1);
  // proto:  const QMetaObject * QAnimationGroup::metaObject();
extern void _ZNK15QAnimationGroup10metaObjectEv(void* qthis);
  // proto:  int QAnimationGroup::indexOfAnimation(QAbstractAnimation * animation);
extern void _ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QAbstractAnimation * QAnimationGroup::animationAt(int index);
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
    C._ZNK15QAnimationGroup11animationAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "animationAt", args)
  }

}

  // proto:  void QAnimationGroup::~QAnimationGroup();
func (this *QAnimationGroup) FreeQAnimationGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAnimationGroup", "~QAnimationGroup", args)
  }

}

  // proto:  void QAnimationGroup::QAnimationGroup(const QAnimationGroup & );
func NewQAnimationGroup(args ...interface{}) QAnimationGroup {
  return QAnimationGroup{}
}

  // proto:  void QAnimationGroup::removeAnimation(QAbstractAnimation * animation);
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
    C._ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "removeAnimation", args)
  }

}

  // proto:  int QAnimationGroup::animationCount();
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
    C._ZNK15QAnimationGroup14animationCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "animationCount", args)
  }

}

  // proto:  void QAnimationGroup::addAnimation(QAbstractAnimation * animation);
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
    C._ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "addAnimation", args)
  }

}

  // proto:  void QAnimationGroup::clear();
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
    C._ZN15QAnimationGroup5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "clear", args)
  }

}

  // proto:  QAbstractAnimation * QAnimationGroup::takeAnimation(int index);
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
    C._ZN15QAnimationGroup13takeAnimationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "takeAnimation", args)
  }

}

  // proto:  void QAnimationGroup::insertAnimation(int index, QAbstractAnimation * animation);
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
    C._ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "insertAnimation", args)
  }

}

  // proto:  const QMetaObject * QAnimationGroup::metaObject();
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
    C._ZNK15QAnimationGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "metaObject", args)
  }

}

  // proto:  int QAnimationGroup::indexOfAnimation(QAbstractAnimation * animation);
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
    C._ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAnimationGroup", "indexOfAnimation", args)
  }

}

// <= body block end

