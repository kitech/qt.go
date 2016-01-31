package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QByteArray QSplitter::saveState();
extern void C_ZNK9QSplitter9saveStateEv(void* qthis); // 4
  // proto:  Qt::Orientation QSplitter::orientation();
extern void C_ZNK9QSplitter11orientationEv(void* qthis); // 4
  // proto:  bool QSplitter::restoreState(const QByteArray & state);
extern void C_ZN9QSplitter12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QList<int> QSplitter::sizes();
extern void C_ZNK9QSplitter5sizesEv(void* qthis); // 4
  // proto:  QWidget * QSplitter::widget(int index);
extern void C_ZNK9QSplitter6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSplitter::getRange(int index, int * , int * );
extern void C_ZNK9QSplitter8getRangeEiPiS0_(void* qthis, int32_t arg0, int32_t* arg1, int32_t* arg2); // 4
  // proto:  void QSplitter::setOpaqueResize(bool opaque);
extern void C_ZN9QSplitter15setOpaqueResizeEb(void* qthis, bool arg0); // 4
  // proto:  void QSplitter::setHandleWidth(int );
extern void C_ZN9QSplitter14setHandleWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSplitter::handleWidth();
extern void C_ZNK9QSplitter11handleWidthEv(void* qthis); // 4
  // proto:  void QSplitter::setStretchFactor(int index, int stretch);
extern void C_ZN9QSplitter16setStretchFactorEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QSplitter::childrenCollapsible();
extern void C_ZNK9QSplitter19childrenCollapsibleEv(void* qthis); // 4
  // proto:  bool QSplitter::isCollapsible(int index);
extern void C_ZNK9QSplitter13isCollapsibleEi(void* qthis, int32_t arg0); // 4
  // proto:  QSplitterHandle * QSplitter::handle(int index);
extern void C_ZNK9QSplitter6handleEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSplitter::indexOf(QWidget * w);
extern void C_ZNK9QSplitter7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QSplitter::insertWidget(int index, QWidget * widget);
extern void C_ZN9QSplitter12insertWidgetEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QSplitter::QSplitter(QWidget * parent);
extern void C_ZN9QSplitterC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QSplitter::addWidget(QWidget * widget);
extern void C_ZN9QSplitter9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QSize QSplitter::sizeHint();
extern void C_ZNK9QSplitter8sizeHintEv(void* qthis); // 4
  // proto:  void QSplitter::setChildrenCollapsible(bool );
extern void C_ZN9QSplitter22setChildrenCollapsibleEb(void* qthis, bool arg0); // 4
  // proto:  int QSplitter::count();
extern void C_ZNK9QSplitter5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QSplitter::metaObject();
extern void C_ZNK9QSplitter10metaObjectEv(void* qthis); // 4
  // proto:  void QSplitter::setCollapsible(int index, bool );
extern void C_ZN9QSplitter14setCollapsibleEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  QSize QSplitter::minimumSizeHint();
extern void C_ZNK9QSplitter15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QSplitter::~QSplitter();
extern void C_ZN9QSplitterD2Ev(void* qthis); // 4
  // proto:  void QSplitter::refresh();
extern void C_ZN9QSplitter7refreshEv(void* qthis); // 4
  // proto:  bool QSplitter::opaqueResize();
extern void C_ZNK9QSplitter12opaqueResizeEv(void* qthis); // 4
  // proto:  Qt::Orientation QSplitterHandle::orientation();
extern void C_ZNK15QSplitterHandle11orientationEv(void* qthis); // 4
  // proto:  void QSplitterHandle::~QSplitterHandle();
extern void C_ZN15QSplitterHandleD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QSplitterHandle::metaObject();
extern void C_ZNK15QSplitterHandle10metaObjectEv(void* qthis); // 4
  // proto:  QSplitter * QSplitterHandle::splitter();
extern void C_ZNK15QSplitterHandle8splitterEv(void* qthis); // 4
  // proto:  QSize QSplitterHandle::sizeHint();
extern void C_ZNK15QSplitterHandle8sizeHintEv(void* qthis); // 4
  // proto:  bool QSplitterHandle::opaqueResize();
extern void C_ZNK15QSplitterHandle12opaqueResizeEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _splitterMoved QSplitter_splitterMoved_signal;
}

// class sizeof(QSplitterHandle)=1
type QSplitterHandle struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
}

// saveState()
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
    C.C_ZNK9QSplitter9saveStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "saveState", args)
  }

}

// orientation()
func (this *QSplitter) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSplitter11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK9QSplitter11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "orientation", args)
  }

}

// restoreState(const class QByteArray &)
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
    C.C_ZN9QSplitter12restoreStateERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "restoreState", args)
  }

}

// sizes()
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
    C.C_ZNK9QSplitter5sizesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "sizes", args)
  }

}

// widget(int)
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
    C.C_ZNK9QSplitter6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "widget", args)
  }

}

// getRange(int, int *, int *)
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
    C.C_ZNK9QSplitter8getRangeEiPiS0_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSplitter", "getRange", args)
  }

}

// setOpaqueResize(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSplitter15setOpaqueResizeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setOpaqueResize", args)
  }

}

// setHandleWidth(int)
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
    C.C_ZN9QSplitter14setHandleWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setHandleWidth", args)
  }

}

// handleWidth()
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
    C.C_ZNK9QSplitter11handleWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "handleWidth", args)
  }

}

// setStretchFactor(int, int)
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
    C.C_ZN9QSplitter16setStretchFactorEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setStretchFactor", args)
  }

}

// childrenCollapsible()
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
    C.C_ZNK9QSplitter19childrenCollapsibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "childrenCollapsible", args)
  }

}

// isCollapsible(int)
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
    C.C_ZNK9QSplitter13isCollapsibleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "isCollapsible", args)
  }

}

// handle(int)
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
    C.C_ZNK9QSplitter6handleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "handle", args)
  }

}

// indexOf(class QWidget *)
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
    C.C_ZNK9QSplitter7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "indexOf", args)
  }

}

// insertWidget(int, class QWidget *)
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
    C.C_ZN9QSplitter12insertWidgetEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "insertWidget", args)
  }

}

// QSplitter(class QWidget *)
func NewQSplitter(args ...interface{}) QSplitter {
  // QSplitter(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitterC1EP7QWidget
    // invoke: void QSplitter(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QSplitterC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "QSplitter", args)
  }

  return QSplitter{}
}

// addWidget(class QWidget *)
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
    C.C_ZN9QSplitter9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "addWidget", args)
  }

}

// sizeHint()
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
    C.C_ZNK9QSplitter8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "sizeHint", args)
  }

}

// setChildrenCollapsible(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSplitter22setChildrenCollapsibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setChildrenCollapsible", args)
  }

}

// count()
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
    C.C_ZNK9QSplitter5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "count", args)
  }

}

// metaObject()
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
    C.C_ZNK9QSplitter10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "metaObject", args)
  }

}

// setCollapsible(int, _Bool)
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
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QSplitter14setCollapsibleEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setCollapsible", args)
  }

}

// minimumSizeHint()
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
    C.C_ZNK9QSplitter15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "minimumSizeHint", args)
  }

}

// ~QSplitter()
func (this *QSplitter) FreeQSplitter(args ...interface{}) () {
  // ~QSplitter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSplitterD0Ev
    // invoke: void ~QSplitter()
    C.C_ZN9QSplitterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "~QSplitter", args)
  }

}

// refresh()
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
    C.C_ZN9QSplitter7refreshEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "refresh", args)
  }

}

// opaqueResize()
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
    C.C_ZNK9QSplitter12opaqueResizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "opaqueResize", args)
  }

}

// orientation()
func (this *QSplitterHandle) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSplitterHandle11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK15QSplitterHandle11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "orientation", args)
  }

}

// ~QSplitterHandle()
func (this *QSplitterHandle) FreeQSplitterHandle(args ...interface{}) () {
  // ~QSplitterHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSplitterHandleD0Ev
    // invoke: void ~QSplitterHandle()
    C.C_ZN15QSplitterHandleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "~QSplitterHandle", args)
  }

}

// metaObject()
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
    C.C_ZNK15QSplitterHandle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "metaObject", args)
  }

}

// splitter()
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
    C.C_ZNK15QSplitterHandle8splitterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "splitter", args)
  }

}

// sizeHint()
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
    C.C_ZNK15QSplitterHandle8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "sizeHint", args)
  }

}

// opaqueResize()
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
    C.C_ZNK15QSplitterHandle12opaqueResizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "opaqueResize", args)
  }

}

// <= body block end

