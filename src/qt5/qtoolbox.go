package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qtoolbox.h
// dst-file: /src/widgets/qtoolbox.go
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
  // proto:  void QToolBox::removeItem(int index);
extern void _ZN8QToolBox10removeItemEi(void* qthis, int32_t arg0);
  // proto:  int QToolBox::insertItem(int index, QWidget * widget, const QString & text);
extern void demth_ZN8QToolBox10insertItemEiP7QWidgetRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2);
  // proto:  QString QToolBox::itemText(int index);
extern void _ZNK8QToolBox8itemTextEi(void* qthis, int32_t arg0);
  // proto:  int QToolBox::indexOf(QWidget * widget);
extern void _ZNK8QToolBox7indexOfEP7QWidget(void* qthis, void* arg0);
  // proto:  QString QToolBox::itemToolTip(int index);
extern void _ZNK8QToolBox11itemToolTipEi(void* qthis, int32_t arg0);
  // proto:  void QToolBox::QToolBox(const QToolBox & );
extern void* dector_ZN8QToolBoxC1ERKS_(void* arg0);
extern void _ZN8QToolBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QToolBox::setCurrentWidget(QWidget * widget);
extern void _ZN8QToolBox16setCurrentWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QToolBox::setCurrentIndex(int index);
extern void _ZN8QToolBox15setCurrentIndexEi(void* qthis, int32_t arg0);
  // proto:  void QToolBox::setItemIcon(int index, const QIcon & icon);
extern void _ZN8QToolBox11setItemIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QToolBox::setItemText(int index, const QString & text);
extern void _ZN8QToolBox11setItemTextEiRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  int QToolBox::count();
extern void _ZNK8QToolBox5countEv(void* qthis);
  // proto:  const QMetaObject * QToolBox::metaObject();
extern void _ZNK8QToolBox10metaObjectEv(void* qthis);
  // proto:  QWidget * QToolBox::widget(int index);
extern void _ZNK8QToolBox6widgetEi(void* qthis, int32_t arg0);
  // proto:  void QToolBox::setItemToolTip(int index, const QString & toolTip);
extern void _ZN8QToolBox14setItemToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  int QToolBox::currentIndex();
extern void _ZNK8QToolBox12currentIndexEv(void* qthis);
  // proto:  QWidget * QToolBox::currentWidget();
extern void _ZNK8QToolBox13currentWidgetEv(void* qthis);
  // proto:  int QToolBox::addItem(QWidget * widget, const QString & text);
extern void demth_ZN8QToolBox7addItemEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  bool QToolBox::isItemEnabled(int index);
extern void _ZNK8QToolBox13isItemEnabledEi(void* qthis, int32_t arg0);
  // proto:  void QToolBox::setItemEnabled(int index, bool enabled);
extern void _ZN8QToolBox14setItemEnabledEib(void* qthis, int32_t arg0, bool arg1);
  // proto:  QIcon QToolBox::itemIcon(int index);
extern void _ZNK8QToolBox8itemIconEi(void* qthis, int32_t arg0);
  // proto:  void QToolBox::~QToolBox();
extern void _ZN8QToolBoxD0Ev(void* qthis);
  // proto:  int QToolBox::addItem(QWidget * widget, const QIcon & icon, const QString & text);
extern void demth_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  int QToolBox::insertItem(int index, QWidget * widget, const QIcon & icon, const QString & text);
extern void _ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3);
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

// class sizeof(QToolBox)=1
type QToolBox struct {
  /*qbase*/ QFrame;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentChanged QToolBox_currentChanged_signal;
}

  // proto:  void QToolBox::removeItem(int index);
func (this *QToolBox) removeItem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox10removeItemEi
    // invoke: void removeItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QToolBox10removeItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "removeItem", args)
  }

}

  // proto:  int QToolBox::insertItem(int index, QWidget * widget, const QString & text);
func (this *QToolBox) insertItem(args ...interface{}) () {
  // insertItem(int, class QWidget *, const class QString &)
  // insertItem(int, class QWidget *, const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox10insertItemEiP7QWidgetRK7QString
    // invoke: int insertItem(int, class QWidget *, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QToolBox10insertItemEiP7QWidgetRK7QString(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString
    // invoke: int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QIcon).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QToolBox", "insertItem", args)
  }

}

  // proto:  QString QToolBox::itemText(int index);
func (this *QToolBox) itemText(args ...interface{}) () {
  // itemText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox8itemTextEi
    // invoke: QString itemText(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox8itemTextEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "itemText", args)
  }

}

  // proto:  int QToolBox::indexOf(QWidget * widget);
func (this *QToolBox) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "indexOf", args)
  }

}

  // proto:  QString QToolBox::itemToolTip(int index);
func (this *QToolBox) itemToolTip(args ...interface{}) () {
  // itemToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox11itemToolTipEi
    // invoke: QString itemToolTip(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox11itemToolTipEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "itemToolTip", args)
  }

}

  // proto:  void QToolBox::QToolBox(const QToolBox & );
func NewQToolBox(args ...interface{}) QToolBox {
  return QToolBox{}
}

  // proto:  void QToolBox::setCurrentWidget(QWidget * widget);
func (this *QToolBox) setCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QToolBox16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "setCurrentWidget", args)
  }

}

  // proto:  void QToolBox::setCurrentIndex(int index);
func (this *QToolBox) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QToolBox15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "setCurrentIndex", args)
  }

}

  // proto:  void QToolBox::setItemIcon(int index, const QIcon & icon);
func (this *QToolBox) setItemIcon(args ...interface{}) () {
  // setItemIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox11setItemIconEiRK5QIcon
    // invoke: void setItemIcon(int, const class QIcon &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QToolBox11setItemIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemIcon", args)
  }

}

  // proto:  void QToolBox::setItemText(int index, const QString & text);
func (this *QToolBox) setItemText(args ...interface{}) () {
  // setItemText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox11setItemTextEiRK7QString
    // invoke: void setItemText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QToolBox11setItemTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemText", args)
  }

}

  // proto:  int QToolBox::count();
func (this *QToolBox) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox5countEv
    // invoke: int count()
    C._ZNK8QToolBox5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBox", "count", args)
  }

}

  // proto:  const QMetaObject * QToolBox::metaObject();
func (this *QToolBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK8QToolBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBox", "metaObject", args)
  }

}

  // proto:  QWidget * QToolBox::widget(int index);
func (this *QToolBox) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "widget", args)
  }

}

  // proto:  void QToolBox::setItemToolTip(int index, const QString & toolTip);
func (this *QToolBox) setItemToolTip(args ...interface{}) () {
  // setItemToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox14setItemToolTipEiRK7QString
    // invoke: void setItemToolTip(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QToolBox14setItemToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemToolTip", args)
  }

}

  // proto:  int QToolBox::currentIndex();
func (this *QToolBox) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox12currentIndexEv
    // invoke: int currentIndex()
    C._ZNK8QToolBox12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBox", "currentIndex", args)
  }

}

  // proto:  QWidget * QToolBox::currentWidget();
func (this *QToolBox) currentWidget(args ...interface{}) () {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox13currentWidgetEv
    // invoke: QWidget * currentWidget()
    C._ZNK8QToolBox13currentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QToolBox", "currentWidget", args)
  }

}

  // proto:  int QToolBox::addItem(QWidget * widget, const QString & text);
func (this *QToolBox) addItem(args ...interface{}) () {
  // addItem(class QWidget *, const class QString &)
  // addItem(class QWidget *, const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox7addItemEP7QWidgetRK7QString
    // invoke: int addItem(class QWidget *, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QToolBox7addItemEP7QWidgetRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString
    // invoke: int addItem(class QWidget *, const class QIcon &, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QToolBox", "addItem", args)
  }

}

  // proto:  bool QToolBox::isItemEnabled(int index);
func (this *QToolBox) isItemEnabled(args ...interface{}) () {
  // isItemEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox13isItemEnabledEi
    // invoke: bool isItemEnabled(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox13isItemEnabledEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "isItemEnabled", args)
  }

}

  // proto:  void QToolBox::setItemEnabled(int index, bool enabled);
func (this *QToolBox) setItemEnabled(args ...interface{}) () {
  // setItemEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox14setItemEnabledEib
    // invoke: void setItemEnabled(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN8QToolBox14setItemEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemEnabled", args)
  }

}

  // proto:  QIcon QToolBox::itemIcon(int index);
func (this *QToolBox) itemIcon(args ...interface{}) () {
  // itemIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox8itemIconEi
    // invoke: QIcon itemIcon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK8QToolBox8itemIconEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "itemIcon", args)
  }

}

  // proto:  void QToolBox::~QToolBox();
func (this *QToolBox) FreeQToolBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolBox", "~QToolBox", args)
  }

}

// <= body block end

