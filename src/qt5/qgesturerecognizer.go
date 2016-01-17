package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGestureRecognizer::reset(QGesture * state);
extern void _ZN18QGestureRecognizer5resetEP8QGesture(void* qthis, void* arg0); // 4
  // proto:  void QGestureRecognizer::~QGestureRecognizer();
extern void _ZN18QGestureRecognizerD2Ev(void* qthis); // 4
  // proto:  QGesture * QGestureRecognizer::create(QObject * target);
extern void _ZN18QGestureRecognizer6createEP7QObject(void* qthis, void* arg0); // 4
  // proto: static Qt::GestureType QGestureRecognizer::registerRecognizer(QGestureRecognizer * recognizer);
extern void _ZN18QGestureRecognizer18registerRecognizerEPS_(void* arg0); // 4
  // proto:  void QGestureRecognizer::QGestureRecognizer();
extern void _ZN18QGestureRecognizerC2Ev(void* qthis); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// reset(class QGesture *)
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

// ~QGestureRecognizer()
func (this *QGestureRecognizer) FreeQGestureRecognizer(args ...interface{}) () {
  // ~QGestureRecognizer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizerD0Ev
    // invoke: void ~QGestureRecognizer()
    C._ZN18QGestureRecognizerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "~QGestureRecognizer", args)
  }

}

// create(class QObject *)
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

// registerRecognizer(class QGestureRecognizer *)
func (this *QGestureRecognizer) registerRecognizer_s(args ...interface{}) () {
  // registerRecognizer(class QGestureRecognizer *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGestureRecognizer{}) // "QGestureRecognizer *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer18registerRecognizerEPS_
    // invoke: Qt::GestureType registerRecognizer(class QGestureRecognizer *)
    var arg0 = args[0].(QGestureRecognizer).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QGestureRecognizer18registerRecognizerEPS_(arg0)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "registerRecognizer", args)
  }

}

// QGestureRecognizer()
func NewQGestureRecognizer(args ...interface{}) QGestureRecognizer {
  // QGestureRecognizer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizerC1Ev
    // invoke: void QGestureRecognizer()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN18QGestureRecognizerC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "QGestureRecognizer", args)
  }

  return QGestureRecognizer{}
}

// <= body block end

