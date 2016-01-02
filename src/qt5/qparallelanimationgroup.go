package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qparallelanimationgroup.h
// dst-file: /src/core/qparallelanimationgroup.go
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
  // proto:  void QParallelAnimationGroup::~QParallelAnimationGroup();
extern void _ZN23QParallelAnimationGroupD0Ev(void* qthis);
  // proto:  void QParallelAnimationGroup::QParallelAnimationGroup(const QParallelAnimationGroup & );
extern void* dector_ZN23QParallelAnimationGroupC1ERKS_(void* arg0);
extern void _ZN23QParallelAnimationGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  int QParallelAnimationGroup::duration();
extern void _ZNK23QParallelAnimationGroup8durationEv(void* qthis);
  // proto:  void QParallelAnimationGroup::QParallelAnimationGroup(QObject * parent);
extern void* dector_ZN23QParallelAnimationGroupC1EP7QObject(void* arg0);
extern void _ZN23QParallelAnimationGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QParallelAnimationGroup::metaObject();
extern void _ZNK23QParallelAnimationGroup10metaObjectEv(void* qthis);
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

// class sizeof(QParallelAnimationGroup)=1
type QParallelAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QParallelAnimationGroup::~QParallelAnimationGroup();
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

  // proto:  void QParallelAnimationGroup::QParallelAnimationGroup(const QParallelAnimationGroup & );
func NewQParallelAnimationGroup(args ...interface{}) QParallelAnimationGroup {
  return QParallelAnimationGroup{}
}

  // proto:  int QParallelAnimationGroup::duration();
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

  // proto:  const QMetaObject * QParallelAnimationGroup::metaObject();
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

