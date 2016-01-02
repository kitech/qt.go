package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  int QStackedLayout::insertWidget(int index, QWidget * w);
extern void _ZN14QStackedLayout12insertWidgetEiP7QWidget(void* qthis, int arg0, void* arg1);
  // proto:  void QStackedLayout::QStackedLayout(QLayout * parentLayout);
extern void* dector_ZN14QStackedLayoutC1EP7QLayout(void* arg0);
extern void _ZN14QStackedLayoutC1EP7QLayout(void* qthis, void* arg0);
  // proto:  void QStackedLayout::QStackedLayout(QWidget * parent);
extern void* dector_ZN14QStackedLayoutC1EP7QWidget(void* arg0);
extern void _ZN14QStackedLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QStackedLayout::setGeometry(const QRect & rect);
extern void _ZN14QStackedLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  QWidget * QStackedLayout::currentWidget();
extern void _ZNK14QStackedLayout13currentWidgetEv(void* qthis);
  // proto:  QLayoutItem * QStackedLayout::takeAt(int );
extern void _ZN14QStackedLayout6takeAtEi(void* qthis, int arg0);
  // proto:  QSize QStackedLayout::minimumSize();
extern void _ZNK14QStackedLayout11minimumSizeEv(void* qthis);
  // proto:  QSize QStackedLayout::sizeHint();
extern void _ZNK14QStackedLayout8sizeHintEv(void* qthis);
  // proto:  void QStackedLayout::QStackedLayout(const QStackedLayout & );
extern void* dector_ZN14QStackedLayoutC1ERKS_(void* arg0);
extern void _ZN14QStackedLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  int QStackedLayout::currentIndex();
extern void _ZNK14QStackedLayout12currentIndexEv(void* qthis);
  // proto:  int QStackedLayout::count();
extern void _ZNK14QStackedLayout5countEv(void* qthis);
  // proto:  void QStackedLayout::addItem(QLayoutItem * item);
extern void _ZN14QStackedLayout7addItemEP11QLayoutItem(void* qthis, void* arg0);
  // proto:  void QStackedLayout::setCurrentWidget(QWidget * w);
extern void _ZN14QStackedLayout16setCurrentWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QStackedLayout::metaObject();
extern void _ZNK14QStackedLayout10metaObjectEv(void* qthis);
  // proto:  void QStackedLayout::setCurrentIndex(int index);
extern void _ZN14QStackedLayout15setCurrentIndexEi(void* qthis, int arg0);
  // proto:  QLayoutItem * QStackedLayout::itemAt(int );
extern void _ZNK14QStackedLayout6itemAtEi(void* qthis, int arg0);
  // proto:  void QStackedLayout::~QStackedLayout();
extern void _ZN14QStackedLayoutD0Ev(void* qthis);
  // proto:  int QStackedLayout::addWidget(QWidget * w);
extern void _ZN14QStackedLayout9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  int QStackedLayout::heightForWidth(int width);
extern void _ZNK14QStackedLayout14heightForWidthEi(void* qthis, int arg0);
  // proto:  bool QStackedLayout::hasHeightForWidth();
extern void _ZNK14QStackedLayout17hasHeightForWidthEv(void* qthis);
  // proto:  void QStackedLayout::QStackedLayout();
extern void* dector_ZN14QStackedLayoutC1Ev();
extern void _ZN14QStackedLayoutC1Ev(void* qthis);
  // proto:  QWidget * QStackedLayout::widget(int );
extern void _ZNK14QStackedLayout6widgetEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _widgetRemoved QStackedLayout_widgetRemoved_signal;
//  _currentChanged QStackedLayout_currentChanged_signal;
}

  // proto:  int QStackedLayout::insertWidget(int index, QWidget * w);
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
    C._ZN14QStackedLayout12insertWidgetEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStackedLayout", "insertWidget", args)
  }

}

  // proto:  void QStackedLayout::QStackedLayout(QLayout * parentLayout);
func NewQStackedLayout(args ...interface{}) QStackedLayout {
  return QStackedLayout{}
}

  // proto:  void QStackedLayout::setGeometry(const QRect & rect);
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
    C._ZN14QStackedLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setGeometry", args)
  }

}

  // proto:  QWidget * QStackedLayout::currentWidget();
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
    C._ZNK14QStackedLayout13currentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentWidget", args)
  }

}

  // proto:  QLayoutItem * QStackedLayout::takeAt(int );
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
    C._ZN14QStackedLayout6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "takeAt", args)
  }

}

  // proto:  QSize QStackedLayout::minimumSize();
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
    C._ZNK14QStackedLayout11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "minimumSize", args)
  }

}

  // proto:  QSize QStackedLayout::sizeHint();
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
    C._ZNK14QStackedLayout8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "sizeHint", args)
  }

}

  // proto:  int QStackedLayout::currentIndex();
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
    C._ZNK14QStackedLayout12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "currentIndex", args)
  }

}

  // proto:  int QStackedLayout::count();
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
    C._ZNK14QStackedLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "count", args)
  }

}

  // proto:  void QStackedLayout::addItem(QLayoutItem * item);
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
    C._ZN14QStackedLayout7addItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "addItem", args)
  }

}

  // proto:  void QStackedLayout::setCurrentWidget(QWidget * w);
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
    C._ZN14QStackedLayout16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentWidget", args)
  }

}

  // proto:  const QMetaObject * QStackedLayout::metaObject();
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
    C._ZNK14QStackedLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "metaObject", args)
  }

}

  // proto:  void QStackedLayout::setCurrentIndex(int index);
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
    C._ZN14QStackedLayout15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "setCurrentIndex", args)
  }

}

  // proto:  QLayoutItem * QStackedLayout::itemAt(int );
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
    C._ZNK14QStackedLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "itemAt", args)
  }

}

  // proto:  void QStackedLayout::~QStackedLayout();
func (this *QStackedLayout) FreeQStackedLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStackedLayout", "~QStackedLayout", args)
  }

}

  // proto:  int QStackedLayout::addWidget(QWidget * w);
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
    C._ZN14QStackedLayout9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "addWidget", args)
  }

}

  // proto:  int QStackedLayout::heightForWidth(int width);
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
    C._ZNK14QStackedLayout14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "heightForWidth", args)
  }

}

  // proto:  bool QStackedLayout::hasHeightForWidth();
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
    C._ZNK14QStackedLayout17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedLayout", "hasHeightForWidth", args)
  }

}

  // proto:  QWidget * QStackedLayout::widget(int );
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
    C._ZNK14QStackedLayout6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedLayout", "widget", args)
  }

}

// <= body block end

