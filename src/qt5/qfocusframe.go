package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qfocusframe.h
// dst-file: /src/widgets/qfocusframe.go
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
  // proto:  void QFocusFrame::~QFocusFrame();
extern void _ZN11QFocusFrameD2Ev(void* qthis); // 4
  // proto:  QWidget * QFocusFrame::widget();
extern void _ZNK11QFocusFrame6widgetEv(void* qthis); // 4
  // proto:  void QFocusFrame::setWidget(QWidget * widget);
extern void _ZN11QFocusFrame9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QFocusFrame::QFocusFrame(QWidget * parent);
extern void _ZN11QFocusFrameC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  const QMetaObject * QFocusFrame::metaObject();
extern void _ZNK11QFocusFrame10metaObjectEv(void* qthis); // 4
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

// class sizeof(QFocusFrame)=1
type QFocusFrame struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QFocusFrame()
func (this *QFocusFrame) FreeQFocusFrame(args ...interface{}) () {
  // ~QFocusFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFocusFrameD0Ev
    // invoke: void ~QFocusFrame()
    C._ZN11QFocusFrameD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusFrame", "~QFocusFrame", args)
  }

}

// widget()
func (this *QFocusFrame) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusFrame6widgetEv
    // invoke: QWidget * widget()
    C._ZNK11QFocusFrame6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusFrame", "widget", args)
  }

}

// setWidget(class QWidget *)
func (this *QFocusFrame) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFocusFrame9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFocusFrame9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFocusFrame", "setWidget", args)
  }

}

// QFocusFrame(class QWidget *)
func NewQFocusFrame(args ...interface{}) QFocusFrame {
  // QFocusFrame(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFocusFrameC1EP7QWidget
    // invoke: void QFocusFrame(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QFocusFrameC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFocusFrame", "QFocusFrame", args)
  }

  return QFocusFrame{}
}

// metaObject()
func (this *QFocusFrame) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFocusFrame10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QFocusFrame10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFocusFrame", "metaObject", args)
  }

}

// <= body block end

