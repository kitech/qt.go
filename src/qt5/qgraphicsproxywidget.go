package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qgraphicsproxywidget.h
// dst-file: /src/widgets/qgraphicsproxywidget.go
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
  // proto:  void QGraphicsProxyWidget::~QGraphicsProxyWidget();
extern void _ZN20QGraphicsProxyWidgetD0Ev(void* qthis);
  // proto:  QGraphicsProxyWidget * QGraphicsProxyWidget::createProxyForChildWidget(QWidget * child);
extern void _ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QGraphicsProxyWidget::setWidget(QWidget * widget);
extern void _ZN20QGraphicsProxyWidget9setWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsProxyWidget::subWidgetRect(const QWidget * widget);
extern void _ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget(void* qthis, void* arg0);
  // proto:  void QGraphicsProxyWidget::QGraphicsProxyWidget(const QGraphicsProxyWidget & );
extern void* dector_ZN20QGraphicsProxyWidgetC1ERKS_(void* arg0);
extern void _ZN20QGraphicsProxyWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsProxyWidget::metaObject();
extern void _ZNK20QGraphicsProxyWidget10metaObjectEv(void* qthis);
  // proto:  void QGraphicsProxyWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsProxyWidget::setGeometry(const QRectF & rect);
extern void _ZN20QGraphicsProxyWidget11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  QWidget * QGraphicsProxyWidget::widget();
extern void _ZNK20QGraphicsProxyWidget6widgetEv(void* qthis);
  // proto:  int QGraphicsProxyWidget::type();
extern void _ZNK20QGraphicsProxyWidget4typeEv(void* qthis);
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

// class sizeof(QGraphicsProxyWidget)=1
type QGraphicsProxyWidget struct {
  /*qbase*/ QGraphicsWidget;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsProxyWidget::~QGraphicsProxyWidget();
func (this *QGraphicsProxyWidget) FreeQGraphicsProxyWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "~QGraphicsProxyWidget", args)
  }

}

  // proto:  QGraphicsProxyWidget * QGraphicsProxyWidget::createProxyForChildWidget(QWidget * child);
func (this *QGraphicsProxyWidget) createProxyForChildWidget(args ...interface{}) () {
  // createProxyForChildWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget
    // invoke: QGraphicsProxyWidget * createProxyForChildWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "createProxyForChildWidget", args)
  }

}

  // proto:  void QGraphicsProxyWidget::setWidget(QWidget * widget);
func (this *QGraphicsProxyWidget) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsProxyWidget9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN20QGraphicsProxyWidget9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "setWidget", args)
  }

}

  // proto:  QRectF QGraphicsProxyWidget::subWidgetRect(const QWidget * widget);
func (this *QGraphicsProxyWidget) subWidgetRect(args ...interface{}) () {
  // subWidgetRect(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget
    // invoke: QRectF subWidgetRect(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "subWidgetRect", args)
  }

}

  // proto:  void QGraphicsProxyWidget::QGraphicsProxyWidget(const QGraphicsProxyWidget & );
func NewQGraphicsProxyWidget(args ...interface{}) QGraphicsProxyWidget {
  return QGraphicsProxyWidget{}
}

  // proto:  const QMetaObject * QGraphicsProxyWidget::metaObject();
func (this *QGraphicsProxyWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsProxyWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK20QGraphicsProxyWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "metaObject", args)
  }

}

  // proto:  void QGraphicsProxyWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
func (this *QGraphicsProxyWidget) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "paint", args)
  }

}

  // proto:  void QGraphicsProxyWidget::setGeometry(const QRectF & rect);
func (this *QGraphicsProxyWidget) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QGraphicsProxyWidget11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN20QGraphicsProxyWidget11setGeometryERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "setGeometry", args)
  }

}

  // proto:  QWidget * QGraphicsProxyWidget::widget();
func (this *QGraphicsProxyWidget) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QGraphicsProxyWidget6widgetEv
    // invoke: QWidget * widget()
    C._ZNK20QGraphicsProxyWidget6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "widget", args)
  }

}

  // proto:  int QGraphicsProxyWidget::type();
func (this *QGraphicsProxyWidget) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsProxyWidget", "type", args)
  }

}

// <= body block end

