package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGestureRecognizer::reset(QGesture * state);
extern void C_ZN18QGestureRecognizer5resetEP8QGesture(void* qthis, void* arg0); // 4
  // proto:  void QGestureRecognizer::~QGestureRecognizer();
extern void C_ZN18QGestureRecognizerD2Ev(void* qthis); // 4
  // proto:  QGesture * QGestureRecognizer::create(QObject * target);
extern void* C_ZN18QGestureRecognizer6createEP7QObject(void* qthis, void* arg0); // 4
  // proto: static Qt::GestureType QGestureRecognizer::registerRecognizer(QGestureRecognizer * recognizer);
extern void C_ZN18QGestureRecognizer18registerRecognizerEPS_(void* arg0); // 4
  // proto:  void QGestureRecognizer::QGestureRecognizer();
extern void* C_ZN18QGestureRecognizerC2Ev(); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGestureRecognizer)=8
type QGestureRecognizer struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// reset(class QGesture *)
func (this *QGestureRecognizer) Reset(args ...interface{}) () {
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
    var arg0 = args[0].(*QGesture).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGestureRecognizer5resetEP8QGesture(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "reset", args)
  }

  return
}

// ~QGestureRecognizer()
func (this *QGestureRecognizer) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QGestureRecognizerD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "~QGestureRecognizer", args)
  }

  return
}

// create(class QObject *)
func (this *QGestureRecognizer) Create(args ...interface{}) (ret interface{}) {
  // create(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer6createEP7QObject
    // invoke: QGesture * create(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QGestureRecognizer6createEP7QObject(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QGesture{}) // "QGesture *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "create", args)
  }

  return
}

// registerRecognizer(class QGestureRecognizer *)
func (this *QGestureRecognizer) RegisterRecognizer_s(args ...interface{}) () {
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
    var arg0 = args[0].(*QGestureRecognizer).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QGestureRecognizer18registerRecognizerEPS_(arg0)
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "registerRecognizer", args)
  }

  return
}

// QGestureRecognizer()
func GcfreeQGestureRecognizer(this *QGestureRecognizer) {
  qtrt.UniverseFree(this)
}
func NewQGestureRecognizer(args ...interface{}) *QGestureRecognizer {
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
    qthis = C.C_ZN18QGestureRecognizerC2Ev()
    this := &QGestureRecognizer{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGestureRecognizer)
    return this
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "QGestureRecognizer", args)
  }

  return nil // QGestureRecognizer{Qclsinst:qthis}
}

// <= body block end

