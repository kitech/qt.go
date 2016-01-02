package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qgesturerecognizer.h
// dst-file: /src/widgets/qgesturerecognizer.go
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
  // proto:  void QGestureRecognizer::~QGestureRecognizer();
extern void _ZN18QGestureRecognizerD0Ev(void* qthis);
  // proto:  void QGestureRecognizer::QGestureRecognizer();
extern void* dector_ZN18QGestureRecognizerC1Ev();
extern void _ZN18QGestureRecognizerC1Ev(void* qthis);
  // proto:  void QGestureRecognizer::reset(QGesture * state);
extern void _ZN18QGestureRecognizer5resetEP8QGesture(void* qthis, void* arg0);
  // proto:  QGesture * QGestureRecognizer::create(QObject * target);
extern void _ZN18QGestureRecognizer6createEP7QObject(void* qthis, void* arg0);
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

// class sizeof(QGestureRecognizer)=8
type QGestureRecognizer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGestureRecognizer::~QGestureRecognizer();
func (this *QGestureRecognizer) FreeQGestureRecognizer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "~QGestureRecognizer", args)
  }

}

  // proto:  void QGestureRecognizer::QGestureRecognizer();
func NewQGestureRecognizer(args ...interface{}) QGestureRecognizer {
  return QGestureRecognizer{}
}

  // proto:  void QGestureRecognizer::reset(QGesture * state);
func (this *QGestureRecognizer) reset(args ...interface{}) () {
  // reset(class QGesture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer5resetEP8QGesture
    // invoke: void reset(class QGesture *)
    var arg0 = args[0].(QGesture).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QGestureRecognizer5resetEP8QGesture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "reset", args)
  }

}

  // proto:  QGesture * QGestureRecognizer::create(QObject * target);
func (this *QGestureRecognizer) create(args ...interface{}) () {
  // create(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer6createEP7QObject
    // invoke: QGesture * create(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QGestureRecognizer6createEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "create", args)
  }

}

// <= body block end

