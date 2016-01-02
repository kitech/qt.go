package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qstackedwidget.h
// dst-file: /src/widgets/qstackedwidget.go
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
  // proto:  void QStackedWidget::setCurrentIndex(int index);
extern void _ZN14QStackedWidget15setCurrentIndexEi(void* qthis, int arg0);
  // proto:  void QStackedWidget::QStackedWidget(QWidget * parent);
extern void* dector_ZN14QStackedWidgetC1EP7QWidget(void* arg0);
extern void _ZN14QStackedWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QStackedWidget::currentWidget();
extern void _ZNK14QStackedWidget13currentWidgetEv(void* qthis);
  // proto:  int QStackedWidget::insertWidget(int index, QWidget * w);
extern void _ZN14QStackedWidget12insertWidgetEiP7QWidget(void* qthis, int arg0, void* arg1);
  // proto:  int QStackedWidget::indexOf(QWidget * );
extern void _ZNK14QStackedWidget7indexOfEP7QWidget(void* qthis, void* arg0);
  // proto:  void QStackedWidget::removeWidget(QWidget * w);
extern void _ZN14QStackedWidget12removeWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QStackedWidget::widget(int );
extern void _ZNK14QStackedWidget6widgetEi(void* qthis, int arg0);
  // proto:  int QStackedWidget::addWidget(QWidget * w);
extern void _ZN14QStackedWidget9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  int QStackedWidget::currentIndex();
extern void _ZNK14QStackedWidget12currentIndexEv(void* qthis);
  // proto:  int QStackedWidget::count();
extern void _ZNK14QStackedWidget5countEv(void* qthis);
  // proto:  void QStackedWidget::setCurrentWidget(QWidget * w);
extern void _ZN14QStackedWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QStackedWidget::~QStackedWidget();
extern void _ZN14QStackedWidgetD0Ev(void* qthis);
  // proto:  void QStackedWidget::QStackedWidget(const QStackedWidget & );
extern void* dector_ZN14QStackedWidgetC1ERKS_(void* arg0);
extern void _ZN14QStackedWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QStackedWidget::metaObject();
extern void _ZNK14QStackedWidget10metaObjectEv(void* qthis);
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

// class sizeof(QStackedWidget)=1
type QStackedWidget struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
//  _widgetRemoved QStackedWidget_widgetRemoved_signal;
//  _currentChanged QStackedWidget_currentChanged_signal;
}

  // proto:  void QStackedWidget::setCurrentIndex(int index);
func (this *QStackedWidget) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QStackedWidget15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "setCurrentIndex", args)
  }

}

  // proto:  void QStackedWidget::QStackedWidget(QWidget * parent);
func NewQStackedWidget(args ...interface{}) QStackedWidget {
  return QStackedWidget{}
}

  // proto:  QWidget * QStackedWidget::currentWidget();
func (this *QStackedWidget) currentWidget(args ...interface{}) () {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget13currentWidgetEv
    // invoke: QWidget * currentWidget()
    C._ZNK14QStackedWidget13currentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedWidget", "currentWidget", args)
  }

}

  // proto:  int QStackedWidget::insertWidget(int index, QWidget * w);
func (this *QStackedWidget) insertWidget(args ...interface{}) () {
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
    // invoke: _ZN14QStackedWidget12insertWidgetEiP7QWidget
    // invoke: int insertWidget(int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QStackedWidget12insertWidgetEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStackedWidget", "insertWidget", args)
  }

}

  // proto:  int QStackedWidget::indexOf(QWidget * );
func (this *QStackedWidget) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QStackedWidget7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "indexOf", args)
  }

}

  // proto:  void QStackedWidget::removeWidget(QWidget * w);
func (this *QStackedWidget) removeWidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QStackedWidget12removeWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "removeWidget", args)
  }

}

  // proto:  QWidget * QStackedWidget::widget(int );
func (this *QStackedWidget) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK14QStackedWidget6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "widget", args)
  }

}

  // proto:  int QStackedWidget::addWidget(QWidget * w);
func (this *QStackedWidget) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget9addWidgetEP7QWidget
    // invoke: int addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QStackedWidget9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "addWidget", args)
  }

}

  // proto:  int QStackedWidget::currentIndex();
func (this *QStackedWidget) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget12currentIndexEv
    // invoke: int currentIndex()
    C._ZNK14QStackedWidget12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedWidget", "currentIndex", args)
  }

}

  // proto:  int QStackedWidget::count();
func (this *QStackedWidget) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget5countEv
    // invoke: int count()
    C._ZNK14QStackedWidget5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedWidget", "count", args)
  }

}

  // proto:  void QStackedWidget::setCurrentWidget(QWidget * w);
func (this *QStackedWidget) setCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QStackedWidget16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "setCurrentWidget", args)
  }

}

  // proto:  void QStackedWidget::~QStackedWidget();
func (this *QStackedWidget) FreeQStackedWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStackedWidget", "~QStackedWidget", args)
  }

}

  // proto:  const QMetaObject * QStackedWidget::metaObject();
func (this *QStackedWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QStackedWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStackedWidget", "metaObject", args)
  }

}

// <= body block end

