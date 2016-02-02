package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK9QSplitter9saveStateEv(void* qthis); // 4
  // proto:  Qt::Orientation QSplitter::orientation();
extern void C_ZNK9QSplitter11orientationEv(void* qthis); // 4
  // proto:  bool QSplitter::restoreState(const QByteArray & state);
extern bool C_ZN9QSplitter12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QList<int> QSplitter::sizes();
extern void C_ZNK9QSplitter5sizesEv(void* qthis); // 4
  // proto:  QWidget * QSplitter::widget(int index);
extern void* C_ZNK9QSplitter6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSplitter::getRange(int index, int * , int * );
extern void C_ZNK9QSplitter8getRangeEiPiS0_(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QSplitter::setOpaqueResize(bool opaque);
extern void C_ZN9QSplitter15setOpaqueResizeEb(void* qthis, bool arg0); // 4
  // proto:  void QSplitter::setHandleWidth(int );
extern void C_ZN9QSplitter14setHandleWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSplitter::handleWidth();
extern int32_t C_ZNK9QSplitter11handleWidthEv(void* qthis); // 4
  // proto:  void QSplitter::setStretchFactor(int index, int stretch);
extern void C_ZN9QSplitter16setStretchFactorEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QSplitter::childrenCollapsible();
extern bool C_ZNK9QSplitter19childrenCollapsibleEv(void* qthis); // 4
  // proto:  bool QSplitter::isCollapsible(int index);
extern bool C_ZNK9QSplitter13isCollapsibleEi(void* qthis, int32_t arg0); // 4
  // proto:  QSplitterHandle * QSplitter::handle(int index);
extern void* C_ZNK9QSplitter6handleEi(void* qthis, int32_t arg0); // 4
  // proto:  int QSplitter::indexOf(QWidget * w);
extern int32_t C_ZNK9QSplitter7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QSplitter::insertWidget(int index, QWidget * widget);
extern void C_ZN9QSplitter12insertWidgetEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QSplitter::QSplitter(QWidget * parent);
extern void* C_ZN9QSplitterC2EP7QWidget(void* arg0); // 3
  // proto:  void QSplitter::addWidget(QWidget * widget);
extern void C_ZN9QSplitter9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QSize QSplitter::sizeHint();
extern void* C_ZNK9QSplitter8sizeHintEv(void* qthis); // 4
  // proto:  void QSplitter::setChildrenCollapsible(bool );
extern void C_ZN9QSplitter22setChildrenCollapsibleEb(void* qthis, bool arg0); // 4
  // proto:  int QSplitter::count();
extern int32_t C_ZNK9QSplitter5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QSplitter::metaObject();
extern void C_ZNK9QSplitter10metaObjectEv(void* qthis); // 4
  // proto:  void QSplitter::setCollapsible(int index, bool );
extern void C_ZN9QSplitter14setCollapsibleEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  QSize QSplitter::minimumSizeHint();
extern void* C_ZNK9QSplitter15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QSplitter::~QSplitter();
extern void C_ZN9QSplitterD2Ev(void* qthis); // 4
  // proto:  void QSplitter::refresh();
extern void C_ZN9QSplitter7refreshEv(void* qthis); // 4
  // proto:  bool QSplitter::opaqueResize();
extern bool C_ZNK9QSplitter12opaqueResizeEv(void* qthis); // 4
  // proto:  Qt::Orientation QSplitterHandle::orientation();
extern void C_ZNK15QSplitterHandle11orientationEv(void* qthis); // 4
  // proto:  void QSplitterHandle::~QSplitterHandle();
extern void C_ZN15QSplitterHandleD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QSplitterHandle::metaObject();
extern void C_ZNK15QSplitterHandle10metaObjectEv(void* qthis); // 4
  // proto:  QSplitter * QSplitterHandle::splitter();
extern void* C_ZNK15QSplitterHandle8splitterEv(void* qthis); // 4
  // proto:  QSize QSplitterHandle::sizeHint();
extern void* C_ZNK15QSplitterHandle8sizeHintEv(void* qthis); // 4
  // proto:  bool QSplitterHandle::opaqueResize();
extern bool C_ZNK15QSplitterHandle12opaqueResizeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _splitterMoved QSplitter_splitterMoved_signal;
}

// class sizeof(QSplitterHandle)=1
type QSplitterHandle struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// saveState()
func (this *QSplitter) Savestate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter9saveStateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "saveState", args)
  }

  return
}

// orientation()
func (this *QSplitter) Orientation(args ...interface{}) () {
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
    C.C_ZNK9QSplitter11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "orientation", args)
  }

  return
}

// restoreState(const class QByteArray &)
func (this *QSplitter) Restorestate(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QSplitter12restoreStateERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "restoreState", args)
  }

  return
}

// sizes()
func (this *QSplitter) Sizes(args ...interface{}) () {
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
    C.C_ZNK9QSplitter5sizesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "sizes", args)
  }

  return
}

// widget(int)
func (this *QSplitter) Widget(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QSplitter6widgetEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "widget", args)
  }

  return
}

// getRange(int, int *, int *)
func (this *QSplitter) Getrange(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK9QSplitter8getRangeEiPiS0_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSplitter", "getRange", args)
  }

  return
}

// setOpaqueResize(_Bool)
func (this *QSplitter) Setopaqueresize(args ...interface{}) () {
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
    C.C_ZN9QSplitter15setOpaqueResizeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setOpaqueResize", args)
  }

  return
}

// setHandleWidth(int)
func (this *QSplitter) Sethandlewidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSplitter14setHandleWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setHandleWidth", args)
  }

  return
}

// handleWidth()
func (this *QSplitter) Handlewidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter11handleWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "handleWidth", args)
  }

  return
}

// setStretchFactor(int, int)
func (this *QSplitter) Setstretchfactor(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QSplitter16setStretchFactorEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setStretchFactor", args)
  }

  return
}

// childrenCollapsible()
func (this *QSplitter) Childrencollapsible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter19childrenCollapsibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "childrenCollapsible", args)
  }

  return
}

// isCollapsible(int)
func (this *QSplitter) Iscollapsible(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QSplitter13isCollapsibleEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "isCollapsible", args)
  }

  return
}

// handle(int)
func (this *QSplitter) Handle(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QSplitter6handleEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSplitterHandle{}) // "QSplitterHandle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "handle", args)
  }

  return
}

// indexOf(class QWidget *)
func (this *QSplitter) Indexof(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QSplitter7indexOfEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "indexOf", args)
  }

  return
}

// insertWidget(int, class QWidget *)
func (this *QSplitter) Insertwidget(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QSplitter12insertWidgetEiP7QWidget(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "insertWidget", args)
  }

  return
}

// QSplitter(class QWidget *)
func NewQSplitter(args ...interface{}) *QSplitter {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSplitterC2EP7QWidget(arg0)
    return &QSplitter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSplitter", "QSplitter", args)
  }

  return nil // QSplitter{Qclsinst:qthis}
}

// addWidget(class QWidget *)
func (this *QSplitter) Addwidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSplitter9addWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "addWidget", args)
  }

  return
}

// sizeHint()
func (this *QSplitter) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "sizeHint", args)
  }

  return
}

// setChildrenCollapsible(_Bool)
func (this *QSplitter) Setchildrencollapsible(args ...interface{}) () {
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
    C.C_ZN9QSplitter22setChildrenCollapsibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplitter", "setChildrenCollapsible", args)
  }

  return
}

// count()
func (this *QSplitter) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "count", args)
  }

  return
}

// metaObject()
func (this *QSplitter) Metaobject(args ...interface{}) () {
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
    C.C_ZNK9QSplitter10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "metaObject", args)
  }

  return
}

// setCollapsible(int, _Bool)
func (this *QSplitter) Setcollapsible(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN9QSplitter14setCollapsibleEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSplitter", "setCollapsible", args)
  }

  return
}

// minimumSizeHint()
func (this *QSplitter) Minimumsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "minimumSizeHint", args)
  }

  return
}

// ~QSplitter()
func (this *QSplitter) Freeqsplitter(args ...interface{}) () {
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
    C.C_ZN9QSplitterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "~QSplitter", args)
  }

  return
}

// refresh()
func (this *QSplitter) Refresh(args ...interface{}) () {
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
    C.C_ZN9QSplitter7refreshEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitter", "refresh", args)
  }

  return
}

// opaqueResize()
func (this *QSplitter) Opaqueresize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSplitter12opaqueResizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitter", "opaqueResize", args)
  }

  return
}

// orientation()
func (this *QSplitterHandle) Orientation(args ...interface{}) () {
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
    C.C_ZNK15QSplitterHandle11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "orientation", args)
  }

  return
}

// ~QSplitterHandle()
func (this *QSplitterHandle) Freeqsplitterhandle(args ...interface{}) () {
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
    C.C_ZN15QSplitterHandleD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "~QSplitterHandle", args)
  }

  return
}

// metaObject()
func (this *QSplitterHandle) Metaobject(args ...interface{}) () {
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
    C.C_ZNK15QSplitterHandle10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplitterHandle", "metaObject", args)
  }

  return
}

// splitter()
func (this *QSplitterHandle) Splitter(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QSplitterHandle8splitterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSplitter{}) // "QSplitter *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitterHandle", "splitter", args)
  }

  return
}

// sizeHint()
func (this *QSplitterHandle) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QSplitterHandle8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitterHandle", "sizeHint", args)
  }

  return
}

// opaqueResize()
func (this *QSplitterHandle) Opaqueresize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QSplitterHandle12opaqueResizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplitterHandle", "opaqueResize", args)
  }

  return
}

// <= body block end

