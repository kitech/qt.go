package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qscrollbar.h
// dst-file: /src/widgets/qscrollbar.go
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
  // proto:  bool QScrollBar::event(QEvent * event);
extern void _ZN10QScrollBar5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  void QScrollBar::QScrollBar(const QScrollBar & );
extern void* dector_ZN10QScrollBarC1ERKS_(void* arg0);
extern void _ZN10QScrollBarC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QScrollBar::metaObject();
extern void _ZNK10QScrollBar10metaObjectEv(void* qthis);
  // proto:  QSize QScrollBar::sizeHint();
extern void _ZNK10QScrollBar8sizeHintEv(void* qthis);
  // proto:  void QScrollBar::QScrollBar(QWidget * parent);
extern void* dector_ZN10QScrollBarC1EP7QWidget(void* arg0);
extern void _ZN10QScrollBarC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QScrollBar::~QScrollBar();
extern void _ZN10QScrollBarD0Ev(void* qthis);
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

// class sizeof(QScrollBar)=1
type QScrollBar struct {
  /*qbase*/ QAbstractSlider;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QScrollBar::event(QEvent * event);
func (this *QScrollBar) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QScrollBar5eventEP6QEvent
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QScrollBar", "event", args)
  }

}

  // proto:  void QScrollBar::QScrollBar(const QScrollBar & );
func NewQScrollBar(args ...interface{}) QScrollBar {
  return QScrollBar{}
}

  // proto:  const QMetaObject * QScrollBar::metaObject();
func (this *QScrollBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QScrollBar10metaObjectEv
  default:
    qtrt.ErrorResolve("QScrollBar", "metaObject", args)
  }

}

  // proto:  QSize QScrollBar::sizeHint();
func (this *QScrollBar) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QScrollBar8sizeHintEv
  default:
    qtrt.ErrorResolve("QScrollBar", "sizeHint", args)
  }

}

  // proto:  void QScrollBar::~QScrollBar();
func (this *QScrollBar) FreeQScrollBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollBar", "~QScrollBar", args)
  }

}

// <= body block end

