package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void C_ZN11QFocusFrameD2Ev(void* qthis); // 4
  // proto:  QWidget * QFocusFrame::widget();
extern void* C_ZNK11QFocusFrame6widgetEv(void* qthis); // 4
  // proto:  void QFocusFrame::setWidget(QWidget * widget);
extern void C_ZN11QFocusFrame9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QFocusFrame::QFocusFrame(QWidget * parent);
extern void* C_ZN11QFocusFrameC2EP7QWidget(void* arg0); // 3
  // proto:  const QMetaObject * QFocusFrame::metaObject();
extern void C_ZNK11QFocusFrame10metaObjectEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QFocusFrame()
func (this *QFocusFrame) Freeqfocusframe(args ...interface{}) () {
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
    C.C_ZN11QFocusFrameD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFocusFrame", "~QFocusFrame", args)
  }

  return
}

// widget()
func (this *QFocusFrame) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QFocusFrame6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFocusFrame", "widget", args)
  }

  return
}

// setWidget(class QWidget *)
func (this *QFocusFrame) Setwidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFocusFrame9setWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFocusFrame", "setWidget", args)
  }

  return
}

// QFocusFrame(class QWidget *)
func NewQFocusFrame(args ...interface{}) *QFocusFrame {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QFocusFrameC2EP7QWidget(arg0)
    return &QFocusFrame{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFocusFrame", "QFocusFrame", args)
  }

  return nil // QFocusFrame{Qclsinst:qthis}
}

// metaObject()
func (this *QFocusFrame) Metaobject(args ...interface{}) () {
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
    C.C_ZNK11QFocusFrame10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFocusFrame", "metaObject", args)
  }

  return
}

// <= body block end

