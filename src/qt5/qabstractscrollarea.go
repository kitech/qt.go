package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qabstractscrollarea.h
// dst-file: /src/widgets/qabstractscrollarea.go
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
  // proto:  QScrollBar * QAbstractScrollArea::horizontalScrollBar();
extern void _ZNK19QAbstractScrollArea19horizontalScrollBarEv(void* qthis);
  // proto:  QSize QAbstractScrollArea::maximumViewportSize();
extern void _ZNK19QAbstractScrollArea19maximumViewportSizeEv(void* qthis);
  // proto:  void QAbstractScrollArea::QAbstractScrollArea(QWidget * parent);
extern void* dector_ZN19QAbstractScrollAreaC1EP7QWidget(void* arg0);
extern void _ZN19QAbstractScrollAreaC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QAbstractScrollArea::setViewport(QWidget * widget);
extern void _ZN19QAbstractScrollArea11setViewportEP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QAbstractScrollArea::minimumSizeHint();
extern void _ZNK19QAbstractScrollArea15minimumSizeHintEv(void* qthis);
  // proto:  void QAbstractScrollArea::QAbstractScrollArea(const QAbstractScrollArea & );
extern void* dector_ZN19QAbstractScrollAreaC1ERKS_(void* arg0);
extern void _ZN19QAbstractScrollAreaC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractScrollArea::setCornerWidget(QWidget * widget);
extern void _ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractScrollArea::metaObject();
extern void _ZNK19QAbstractScrollArea10metaObjectEv(void* qthis);
  // proto:  void QAbstractScrollArea::setupViewport(QWidget * viewport);
extern void _ZN19QAbstractScrollArea13setupViewportEP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QAbstractScrollArea::cornerWidget();
extern void _ZNK19QAbstractScrollArea12cornerWidgetEv(void* qthis);
  // proto:  QScrollBar * QAbstractScrollArea::verticalScrollBar();
extern void _ZNK19QAbstractScrollArea17verticalScrollBarEv(void* qthis);
  // proto:  QWidget * QAbstractScrollArea::viewport();
extern void _ZNK19QAbstractScrollArea8viewportEv(void* qthis);
  // proto:  void QAbstractScrollArea::~QAbstractScrollArea();
extern void _ZN19QAbstractScrollAreaD0Ev(void* qthis);
  // proto:  QSize QAbstractScrollArea::sizeHint();
extern void _ZNK19QAbstractScrollArea8sizeHintEv(void* qthis);
  // proto:  void QAbstractScrollArea::setHorizontalScrollBar(QScrollBar * scrollbar);
extern void _ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar(void* qthis, void* arg0);
  // proto:  void QAbstractScrollArea::setVerticalScrollBar(QScrollBar * scrollbar);
extern void _ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar(void* qthis, void* arg0);
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

// class sizeof(QAbstractScrollArea)=1
type QAbstractScrollArea struct {
  /*qbase*/ QFrame;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QScrollBar * QAbstractScrollArea::horizontalScrollBar();
func (this *QAbstractScrollArea) horizontalScrollBar(args ...interface{}) () {
  // horizontalScrollBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea19horizontalScrollBarEv
    // invoke: QScrollBar * horizontalScrollBar()
    C._ZNK19QAbstractScrollArea19horizontalScrollBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "horizontalScrollBar", args)
  }

}

  // proto:  QSize QAbstractScrollArea::maximumViewportSize();
func (this *QAbstractScrollArea) maximumViewportSize(args ...interface{}) () {
  // maximumViewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea19maximumViewportSizeEv
    // invoke: QSize maximumViewportSize()
    C._ZNK19QAbstractScrollArea19maximumViewportSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "maximumViewportSize", args)
  }

}

  // proto:  void QAbstractScrollArea::QAbstractScrollArea(QWidget * parent);
func NewQAbstractScrollArea(args ...interface{}) QAbstractScrollArea {
  return QAbstractScrollArea{}
}

  // proto:  void QAbstractScrollArea::setViewport(QWidget * widget);
func (this *QAbstractScrollArea) setViewport(args ...interface{}) () {
  // setViewport(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea11setViewportEP7QWidget
    // invoke: void setViewport(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractScrollArea11setViewportEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setViewport", args)
  }

}

  // proto:  QSize QAbstractScrollArea::minimumSizeHint();
func (this *QAbstractScrollArea) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK19QAbstractScrollArea15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "minimumSizeHint", args)
  }

}

  // proto:  void QAbstractScrollArea::setCornerWidget(QWidget * widget);
func (this *QAbstractScrollArea) setCornerWidget(args ...interface{}) () {
  // setCornerWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget
    // invoke: void setCornerWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setCornerWidget", args)
  }

}

  // proto:  const QMetaObject * QAbstractScrollArea::metaObject();
func (this *QAbstractScrollArea) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK19QAbstractScrollArea10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "metaObject", args)
  }

}

  // proto:  void QAbstractScrollArea::setupViewport(QWidget * viewport);
func (this *QAbstractScrollArea) setupViewport(args ...interface{}) () {
  // setupViewport(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea13setupViewportEP7QWidget
    // invoke: void setupViewport(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractScrollArea13setupViewportEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setupViewport", args)
  }

}

  // proto:  QWidget * QAbstractScrollArea::cornerWidget();
func (this *QAbstractScrollArea) cornerWidget(args ...interface{}) () {
  // cornerWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea12cornerWidgetEv
    // invoke: QWidget * cornerWidget()
    C._ZNK19QAbstractScrollArea12cornerWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "cornerWidget", args)
  }

}

  // proto:  QScrollBar * QAbstractScrollArea::verticalScrollBar();
func (this *QAbstractScrollArea) verticalScrollBar(args ...interface{}) () {
  // verticalScrollBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea17verticalScrollBarEv
    // invoke: QScrollBar * verticalScrollBar()
    C._ZNK19QAbstractScrollArea17verticalScrollBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "verticalScrollBar", args)
  }

}

  // proto:  QWidget * QAbstractScrollArea::viewport();
func (this *QAbstractScrollArea) viewport(args ...interface{}) () {
  // viewport()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea8viewportEv
    // invoke: QWidget * viewport()
    C._ZNK19QAbstractScrollArea8viewportEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "viewport", args)
  }

}

  // proto:  void QAbstractScrollArea::~QAbstractScrollArea();
func (this *QAbstractScrollArea) FreeQAbstractScrollArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "~QAbstractScrollArea", args)
  }

}

  // proto:  QSize QAbstractScrollArea::sizeHint();
func (this *QAbstractScrollArea) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK19QAbstractScrollArea8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "sizeHint", args)
  }

}

  // proto:  void QAbstractScrollArea::setHorizontalScrollBar(QScrollBar * scrollbar);
func (this *QAbstractScrollArea) setHorizontalScrollBar(args ...interface{}) () {
  // setHorizontalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar
    // invoke: void setHorizontalScrollBar(class QScrollBar *)
    var arg0 = args[0].(QScrollBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setHorizontalScrollBar", args)
  }

}

  // proto:  void QAbstractScrollArea::setVerticalScrollBar(QScrollBar * scrollbar);
func (this *QAbstractScrollArea) setVerticalScrollBar(args ...interface{}) () {
  // setVerticalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar
    // invoke: void setVerticalScrollBar(class QScrollBar *)
    var arg0 = args[0].(QScrollBar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setVerticalScrollBar", args)
  }

}

// <= body block end

