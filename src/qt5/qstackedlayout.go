package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qstackedlayout.h
// dst-file: /src/widgets/qstackedlayout.go
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
  // proto:  int QStackedLayout::addWidget(QWidget * w);
extern void C_ZN14QStackedLayout9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStackedLayout::setCurrentIndex(int index);
extern void C_ZN14QStackedLayout15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QStackedLayout::currentWidget();
extern void C_ZNK14QStackedLayout13currentWidgetEv(void* qthis); // 4
  // proto:  void QStackedLayout::setGeometry(const QRect & rect);
extern void C_ZN14QStackedLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QStackedLayout::~QStackedLayout();
extern void C_ZN14QStackedLayoutD2Ev(void* qthis); // 4
  // proto:  void QStackedLayout::setCurrentWidget(QWidget * w);
extern void C_ZN14QStackedLayout16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QWidget * QStackedLayout::widget(int );
extern void C_ZNK14QStackedLayout6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStackedLayout::addItem(QLayoutItem * item);
extern void C_ZN14QStackedLayout7addItemEP11QLayoutItem(void* qthis, void* arg0); // 4
  // proto:  QLayoutItem * QStackedLayout::takeAt(int );
extern void C_ZN14QStackedLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStackedLayout::insertWidget(int index, QWidget * w);
extern void C_ZN14QStackedLayout12insertWidgetEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QStackedLayout::QStackedLayout();
extern void C_ZN14QStackedLayoutC2Ev(void* qthis); // 3
  // proto:  void QStackedLayout::QStackedLayout(QLayout * parentLayout);
extern void C_ZN14QStackedLayoutC2EP7QLayout(void* qthis, void* arg0); // 3
  // proto:  void QStackedLayout::QStackedLayout(QWidget * parent);
extern void C_ZN14QStackedLayoutC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QSize QStackedLayout::sizeHint();
extern void C_ZNK14QStackedLayout8sizeHintEv(void* qthis); // 4
  // proto:  QStackedLayout::StackingMode QStackedLayout::stackingMode();
extern void C_ZNK14QStackedLayout12stackingModeEv(void* qthis); // 4
  // proto:  int QStackedLayout::count();
extern void C_ZNK14QStackedLayout5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QStackedLayout::metaObject();
extern void C_ZNK14QStackedLayout10metaObjectEv(void* qthis); // 4
  // proto:  bool QStackedLayout::hasHeightForWidth();
extern void C_ZNK14QStackedLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QStackedLayout::itemAt(int );
extern void C_ZNK14QStackedLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QStackedLayout::minimumSize();
extern void C_ZNK14QStackedLayout11minimumSizeEv(void* qthis); // 4
  // proto:  int QStackedLayout::heightForWidth(int width);
extern void C_ZNK14QStackedLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStackedLayout::currentIndex();
extern void C_ZNK14QStackedLayout12currentIndexEv(void* qthis); // 4
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

// class sizeof(QStackedLayout)=1
type QStackedLayout struct {
  /*qbase*/ QLayout;
  qclsinst unsafe.Pointer /* *C.void */;
//  _widgetRemoved QStackedLayout_widgetRemoved_signal;
//  _currentChanged QStackedLayout_currentChanged_signal;
}

// addWidget(class QWidget *)
func (this *QStackedLayout) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout9addWidgetEP7QWidget
    // invoke: int addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "addWidget", args)
  }

}

// setCurrentIndex(int)
func (this *QStackedLayout) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentIndex", args)
  }

}

// currentWidget()
func (this *QStackedLayout) currentWidget(args ...interface{}) () {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout13currentWidgetEv
    // invoke: QWidget * currentWidget()
    C.C_ZNK14QStackedLayout13currentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentWidget", args)
  }

}

// setGeometry(const class QRect &)
func (this *QStackedLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setGeometry", args)
  }

}

// ~QStackedLayout()
func (this *QStackedLayout) FreeQStackedLayout(args ...interface{}) () {
  // ~QStackedLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayoutD0Ev
    // invoke: void ~QStackedLayout()
    C.C_ZN14QStackedLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "~QStackedLayout", args)
  }

}

// setCurrentWidget(class QWidget *)
func (this *QStackedLayout) setCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentWidget", args)
  }

}

// widget(int)
func (this *QStackedLayout) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QStackedLayout6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "widget", args)
  }

}

// addItem(class QLayoutItem *)
func (this *QStackedLayout) addItem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout7addItemEP11QLayoutItem
    // invoke: void addItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout7addItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "addItem", args)
  }

}

// takeAt(int)
func (this *QStackedLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedLayout6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "takeAt", args)
  }

}

// insertWidget(int, class QWidget *)
func (this *QStackedLayout) insertWidget(args ...interface{}) () {
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
    // invoke: _ZN14QStackedLayout12insertWidgetEiP7QWidget
    // invoke: int insertWidget(int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN14QStackedLayout12insertWidgetEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStackedLayout", "insertWidget", args)
  }

}

// QStackedLayout()
func NewQStackedLayout(args ...interface{}) QStackedLayout {
  // QStackedLayout()
  // QStackedLayout(class QLayout *)
  // QStackedLayout(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedLayoutC1Ev
    // invoke: void QStackedLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN14QStackedLayoutC2Ev(qthis)
  case 1:
    // invoke: _ZN14QStackedLayoutC1EP7QLayout
    // invoke: void QStackedLayout(class QLayout *)
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN14QStackedLayoutC2EP7QLayout(qthis, arg0)
  case 2:
    // invoke: _ZN14QStackedLayoutC1EP7QWidget
    // invoke: void QStackedLayout(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN14QStackedLayoutC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "QStackedLayout", args)
  }

  return QStackedLayout{}
}

// sizeHint()
func (this *QStackedLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK14QStackedLayout8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "sizeHint", args)
  }

}

// stackingMode()
func (this *QStackedLayout) stackingMode(args ...interface{}) () {
  // stackingMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout12stackingModeEv
    // invoke: QStackedLayout::StackingMode stackingMode()
    C.C_ZNK14QStackedLayout12stackingModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "stackingMode", args)
  }

}

// count()
func (this *QStackedLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout5countEv
    // invoke: int count()
    C.C_ZNK14QStackedLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "count", args)
  }

}

// metaObject()
func (this *QStackedLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QStackedLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "metaObject", args)
  }

}

// hasHeightForWidth()
func (this *QStackedLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C.C_ZNK14QStackedLayout17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "hasHeightForWidth", args)
  }

}

// itemAt(int)
func (this *QStackedLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QStackedLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "itemAt", args)
  }

}

// minimumSize()
func (this *QStackedLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    C.C_ZNK14QStackedLayout11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "minimumSize", args)
  }

}

// heightForWidth(int)
func (this *QStackedLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK14QStackedLayout14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "heightForWidth", args)
  }

}

// currentIndex()
func (this *QStackedLayout) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedLayout12currentIndexEv
    // invoke: int currentIndex()
    C.C_ZNK14QStackedLayout12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentIndex", args)
  }

}

// <= body block end

