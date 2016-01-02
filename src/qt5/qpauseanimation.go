package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qpauseanimation.h
// dst-file: /src/core/qpauseanimation.go
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
  // proto:  void QPauseAnimation::QPauseAnimation(const QPauseAnimation & );
extern void* dector_ZN15QPauseAnimationC1ERKS_(void* arg0);
extern void _ZN15QPauseAnimationC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPauseAnimation::setDuration(int msecs);
extern void _ZN15QPauseAnimation11setDurationEi(void* qthis, int arg0);
  // proto:  void QPauseAnimation::QPauseAnimation(QObject * parent);
extern void* dector_ZN15QPauseAnimationC1EP7QObject(void* arg0);
extern void _ZN15QPauseAnimationC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QPauseAnimation::QPauseAnimation(int msecs, QObject * parent);
extern void* dector_ZN15QPauseAnimationC1EiP7QObject(int arg0, void* arg1);
extern void _ZN15QPauseAnimationC1EiP7QObject(void* qthis, int arg0, void* arg1);
  // proto:  int QPauseAnimation::duration();
extern void _ZNK15QPauseAnimation8durationEv(void* qthis);
  // proto:  const QMetaObject * QPauseAnimation::metaObject();
extern void _ZNK15QPauseAnimation10metaObjectEv(void* qthis);
  // proto:  void QPauseAnimation::~QPauseAnimation();
extern void _ZN15QPauseAnimationD0Ev(void* qthis);
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

// class sizeof(QPauseAnimation)=1
type QPauseAnimation struct {
  /*qbase*/ QAbstractAnimation;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPauseAnimation::QPauseAnimation(const QPauseAnimation & );
func NewQPauseAnimation(args ...interface{}) QPauseAnimation {
  return QPauseAnimation{}
}

  // proto:  void QPauseAnimation::setDuration(int msecs);
func (this *QPauseAnimation) setDuration(args ...interface{}) () {
  // setDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QPauseAnimation11setDurationEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPauseAnimation", "setDuration", args)
  }

}

  // proto:  int QPauseAnimation::duration();
func (this *QPauseAnimation) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QPauseAnimation8durationEv
  default:
    qtrt.ErrorResolve("QPauseAnimation", "duration", args)
  }

}

  // proto:  const QMetaObject * QPauseAnimation::metaObject();
func (this *QPauseAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QPauseAnimation10metaObjectEv
  default:
    qtrt.ErrorResolve("QPauseAnimation", "metaObject", args)
  }

}

  // proto:  void QPauseAnimation::~QPauseAnimation();
func (this *QPauseAnimation) FreeQPauseAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPauseAnimation", "~QPauseAnimation", args)
  }

}

// <= body block end

