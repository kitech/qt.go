package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qsplitter.h
// dst-file: /src/widgets/qsplitter.go
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
  // proto:  void QSplitter::insertWidget(int index, QWidget * widget);
extern void _ZN9QSplitter12insertWidgetEiP7QWidget(void* qthis, int arg0, void* arg1);
  // proto:  bool QSplitter::childrenCollapsible();
extern void _ZNK9QSplitter19childrenCollapsibleEv(void* qthis);
  // proto:  int QSplitter::count();
extern void _ZNK9QSplitter5countEv(void* qthis);
  // proto:  QByteArray QSplitter::saveState();
extern void _ZNK9QSplitter9saveStateEv(void* qthis);
  // proto:  const QMetaObject * QSplitter::metaObject();
extern void _ZNK9QSplitter10metaObjectEv(void* qthis);
  // proto:  bool QSplitter::opaqueResize();
extern void _ZNK9QSplitter12opaqueResizeEv(void* qthis);
  // proto:  void QSplitter::addWidget(QWidget * widget);
extern void _ZN9QSplitter9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QSplitter::QSplitter(const QSplitter & );
extern void* dector_ZN9QSplitterC1ERKS_(void* arg0);
extern void _ZN9QSplitterC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSplitter::QSplitter(QWidget * parent);
extern void* dector_ZN9QSplitterC1EP7QWidget(void* arg0);
extern void _ZN9QSplitterC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QSplitter::setHandleWidth(int );
extern void _ZN9QSplitter14setHandleWidthEi(void* qthis, int arg0);
  // proto:  void QSplitter::setStretchFactor(int index, int stretch);
extern void _ZN9QSplitter16setStretchFactorEii(void* qthis, int arg0, int arg1);
  // proto:  QSize QSplitter::minimumSizeHint();
extern void _ZNK9QSplitter15minimumSizeHintEv(void* qthis);
  // proto:  void QSplitter::setOpaqueResize(bool opaque);
extern void _ZN9QSplitter15setOpaqueResizeEb(void* qthis, bool arg0);
  // proto:  QWidget * QSplitter::widget(int index);
extern void _ZNK9QSplitter6widgetEi(void* qthis, int arg0);
  // proto:  QList<int> QSplitter::sizes();
extern void _ZNK9QSplitter5sizesEv(void* qthis);
  // proto:  bool QSplitter::isCollapsible(int index);
extern void _ZNK9QSplitter13isCollapsibleEi(void* qthis, int arg0);
  // proto:  void QSplitter::setCollapsible(int index, bool );
extern void _ZN9QSplitter14setCollapsibleEib(void* qthis, int arg0, bool arg1);
  // proto:  bool QSplitter::restoreState(const QByteArray & state);
extern void _ZN9QSplitter12restoreStateERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QSplitter::~QSplitter();
extern void _ZN9QSplitterD0Ev(void* qthis);
  // proto:  void QSplitter::setChildrenCollapsible(bool );
extern void _ZN9QSplitter22setChildrenCollapsibleEb(void* qthis, bool arg0);
  // proto:  int QSplitter::handleWidth();
extern void _ZNK9QSplitter11handleWidthEv(void* qthis);
  // proto:  void QSplitter::refresh();
extern void _ZN9QSplitter7refreshEv(void* qthis);
  // proto:  QSize QSplitter::sizeHint();
extern void _ZNK9QSplitter8sizeHintEv(void* qthis);
  // proto:  int QSplitter::indexOf(QWidget * w);
extern void _ZNK9QSplitter7indexOfEP7QWidget(void* qthis, void* arg0);
  // proto:  void QSplitter::getRange(int index, int * , int * );
extern void _ZNK9QSplitter8getRangeEiPiS0_(void* qthis, int arg0, int* arg1, int* arg2);
  // proto:  QSplitterHandle * QSplitter::handle(int index);
extern void _ZNK9QSplitter6handleEi(void* qthis, int arg0);
  // proto:  void QSplitterHandle::~QSplitterHandle();
extern void _ZN15QSplitterHandleD0Ev(void* qthis);
  // proto:  void QSplitterHandle::QSplitterHandle(const QSplitterHandle & );
extern void* dector_ZN15QSplitterHandleC1ERKS_(void* arg0);
extern void _ZN15QSplitterHandleC1ERKS_(void* qthis, void* arg0);
  // proto:  QSize QSplitterHandle::sizeHint();
extern void _ZNK15QSplitterHandle8sizeHintEv(void* qthis);
  // proto:  bool QSplitterHandle::opaqueResize();
extern void _ZNK15QSplitterHandle12opaqueResizeEv(void* qthis);
  // proto:  QSplitter * QSplitterHandle::splitter();
extern void _ZNK15QSplitterHandle8splitterEv(void* qthis);
  // proto:  const QMetaObject * QSplitterHandle::metaObject();
extern void _ZNK15QSplitterHandle10metaObjectEv(void* qthis);
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

// class sizeof(QSplitter)=1
type QSplitter struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
//  _splitterMoved QSplitter_splitterMoved_signal;
}

// class sizeof(QSplitterHandle)=1
type QSplitterHandle struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSplitter::insertWidget(int index, QWidget * widget);
func (this *QSplitter) insertWidget(args ...interface{}) () {
  // insertWidget(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter12insertWidgetEiP7QWidget
    // invoke: void insertWidget(int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QSplitter12insertWidgetEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "insertWidget", args)
  }

}

  // proto:  bool QSplitter::childrenCollapsible();
func (this *QSplitter) childrenCollapsible(args ...interface{}) () {
  // childrenCollapsible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter19childrenCollapsibleEv
    // invoke: bool childrenCollapsible()
    C._ZNK9QSplitter19childrenCollapsibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "childrenCollapsible", args)
  }

}

  // proto:  int QSplitter::count();
func (this *QSplitter) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter5countEv
    // invoke: int count()
    C._ZNK9QSplitter5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "count", args)
  }

}

  // proto:  QByteArray QSplitter::saveState();
func (this *QSplitter) saveState(args ...interface{}) () {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter9saveStateEv
    // invoke: QByteArray saveState()
    C._ZNK9QSplitter9saveStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "saveState", args)
  }

}

  // proto:  const QMetaObject * QSplitter::metaObject();
func (this *QSplitter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QSplitter10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "metaObject", args)
  }

}

  // proto:  bool QSplitter::opaqueResize();
func (this *QSplitter) opaqueResize(args ...interface{}) () {
  // opaqueResize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter12opaqueResizeEv
    // invoke: bool opaqueResize()
    C._ZNK9QSplitter12opaqueResizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "opaqueResize", args)
  }

}

  // proto:  void QSplitter::addWidget(QWidget * widget);
func (this *QSplitter) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter9addWidgetEP7QWidget
    // invoke: void addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QSplitter9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "addWidget", args)
  }

}

  // proto:  void QSplitter::QSplitter(const QSplitter & );
func NewQSplitter(args ...interface{}) QSplitter {
  return QSplitter{}
}

  // proto:  void QSplitter::setHandleWidth(int );
func (this *QSplitter) setHandleWidth(args ...interface{}) () {
  // setHandleWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter14setHandleWidthEi
    // invoke: void setHandleWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QSplitter14setHandleWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setHandleWidth", args)
  }

}

  // proto:  void QSplitter::setStretchFactor(int index, int stretch);
func (this *QSplitter) setStretchFactor(args ...interface{}) () {
  // setStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter16setStretchFactorEii
    // invoke: void setStretchFactor(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN9QSplitter16setStretchFactorEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setStretchFactor", args)
  }

}

  // proto:  QSize QSplitter::minimumSizeHint();
func (this *QSplitter) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK9QSplitter15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "minimumSizeHint", args)
  }

}

  // proto:  void QSplitter::setOpaqueResize(bool opaque);
func (this *QSplitter) setOpaqueResize(args ...interface{}) () {
  // setOpaqueResize(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter15setOpaqueResizeEb
    // invoke: void setOpaqueResize(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QSplitter15setOpaqueResizeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setOpaqueResize", args)
  }

}

  // proto:  QWidget * QSplitter::widget(int index);
func (this *QSplitter) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QSplitter6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "widget", args)
  }

}

  // proto:  QList<int> QSplitter::sizes();
func (this *QSplitter) sizes(args ...interface{}) () {
  // sizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter5sizesEv
    // invoke: QList<int> sizes()
    C._ZNK9QSplitter5sizesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "sizes", args)
  }

}

  // proto:  bool QSplitter::isCollapsible(int index);
func (this *QSplitter) isCollapsible(args ...interface{}) () {
  // isCollapsible(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter13isCollapsibleEi
    // invoke: bool isCollapsible(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QSplitter13isCollapsibleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "isCollapsible", args)
  }

}

  // proto:  void QSplitter::setCollapsible(int index, bool );
func (this *QSplitter) setCollapsible(args ...interface{}) () {
  // setCollapsible(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter14setCollapsibleEib
    // invoke: void setCollapsible(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
    C._ZN9QSplitter14setCollapsibleEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setCollapsible", args)
  }

}

  // proto:  bool QSplitter::restoreState(const QByteArray & state);
func (this *QSplitter) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter12restoreStateERK10QByteArray
    // invoke: bool restoreState(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QSplitter12restoreStateERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "restoreState", args)
  }

}

  // proto:  void QSplitter::~QSplitter();
func (this *QSplitter) FreeQSplitter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSplitter", "~QSplitter", args)
  }

}

  // proto:  void QSplitter::setChildrenCollapsible(bool );
func (this *QSplitter) setChildrenCollapsible(args ...interface{}) () {
  // setChildrenCollapsible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter22setChildrenCollapsibleEb
    // invoke: void setChildrenCollapsible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QSplitter22setChildrenCollapsibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setChildrenCollapsible", args)
  }

}

  // proto:  int QSplitter::handleWidth();
func (this *QSplitter) handleWidth(args ...interface{}) () {
  // handleWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter11handleWidthEv
    // invoke: int handleWidth()
    C._ZNK9QSplitter11handleWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "handleWidth", args)
  }

}

  // proto:  void QSplitter::refresh();
func (this *QSplitter) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitter7refreshEv
    // invoke: void refresh()
    C._ZN9QSplitter7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "refresh", args)
  }

}

  // proto:  QSize QSplitter::sizeHint();
func (this *QSplitter) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK9QSplitter8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "sizeHint", args)
  }

}

  // proto:  int QSplitter::indexOf(QWidget * w);
func (this *QSplitter) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QSplitter7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "indexOf", args)
  }

}

  // proto:  void QSplitter::getRange(int index, int * , int * );
func (this *QSplitter) getRange(args ...interface{}) () {
  // getRange(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter8getRangeEiPiS0_
    // invoke: void getRange(int, int *, int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C._ZNK9QSplitter8getRangeEiPiS0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSplitter", "getRange", args)
  }

}

  // proto:  QSplitterHandle * QSplitter::handle(int index);
func (this *QSplitter) handle(args ...interface{}) () {
  // handle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter6handleEi
    // invoke: QSplitterHandle * handle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QSplitter6handleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "handle", args)
  }

}

  // proto:  void QSplitterHandle::~QSplitterHandle();
func (this *QSplitterHandle) FreeQSplitterHandle(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSplitterHandle", "~QSplitterHandle", args)
  }

}

  // proto:  void QSplitterHandle::QSplitterHandle(const QSplitterHandle & );
func NewQSplitterHandle(args ...interface{}) QSplitterHandle {
  return QSplitterHandle{}
}

  // proto:  QSize QSplitterHandle::sizeHint();
func (this *QSplitterHandle) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK15QSplitterHandle8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "sizeHint", args)
  }

}

  // proto:  bool QSplitterHandle::opaqueResize();
func (this *QSplitterHandle) opaqueResize(args ...interface{}) () {
  // opaqueResize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle12opaqueResizeEv
    // invoke: bool opaqueResize()
    C._ZNK15QSplitterHandle12opaqueResizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "opaqueResize", args)
  }

}

  // proto:  QSplitter * QSplitterHandle::splitter();
func (this *QSplitterHandle) splitter(args ...interface{}) () {
  // splitter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle8splitterEv
    // invoke: QSplitter * splitter()
    C._ZNK15QSplitterHandle8splitterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "splitter", args)
  }

}

  // proto:  const QMetaObject * QSplitterHandle::metaObject();
func (this *QSplitterHandle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QSplitterHandle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "metaObject", args)
  }

}

// <= body block end

