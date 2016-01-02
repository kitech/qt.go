package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  default:
    qtrt.ErrorResolve("QStackedLayout", "widget", args)
  }

}

// <= body block end

