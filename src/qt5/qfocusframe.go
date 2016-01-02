package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QFocusFrame::~QFocusFrame();
extern void _ZN11QFocusFrameD0Ev(void* qthis);
  // proto:  const QMetaObject * QFocusFrame::metaObject();
extern void _ZNK11QFocusFrame10metaObjectEv(void* qthis);
  // proto:  void QFocusFrame::QFocusFrame(const QFocusFrame & );
extern void* dector_ZN11QFocusFrameC1ERKS_(void* arg0);
extern void _ZN11QFocusFrameC1ERKS_(void* qthis, void* arg0);
  // proto:  QWidget * QFocusFrame::widget();
extern void _ZNK11QFocusFrame6widgetEv(void* qthis);
  // proto:  void QFocusFrame::QFocusFrame(QWidget * parent);
extern void* dector_ZN11QFocusFrameC1EP7QWidget(void* arg0);
extern void _ZN11QFocusFrameC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QFocusFrame::setWidget(QWidget * widget);
extern void _ZN11QFocusFrame9setWidgetEP7QWidget(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QFocusFrame::~QFocusFrame();
func (this *QFocusFrame) FreeQFocusFrame(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFocusFrame", "~QFocusFrame", args)
  }

}

  // proto:  const QMetaObject * QFocusFrame::metaObject();
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
  default:
    qtrt.ErrorResolve("QFocusFrame", "metaObject", args)
  }

}

  // proto:  void QFocusFrame::QFocusFrame(const QFocusFrame & );
func NewQFocusFrame(args ...interface{}) QFocusFrame {
  return QFocusFrame{}
}

  // proto:  QWidget * QFocusFrame::widget();
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
  default:
    qtrt.ErrorResolve("QFocusFrame", "widget", args)
  }

}

  // proto:  void QFocusFrame::setWidget(QWidget * widget);
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
  default:
    qtrt.ErrorResolve("QFocusFrame", "setWidget", args)
  }

}

// <= body block end

