package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qtabwidget.h
// dst-file: /src/widgets/qtabwidget.go
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
  // proto:  void QTabWidget::setCurrentWidget(QWidget * widget);
extern void _ZN10QTabWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  int QTabWidget::count();
extern void _ZNK10QTabWidget5countEv(void* qthis);
  // proto:  void QTabWidget::setDocumentMode(bool set);
extern void _ZN10QTabWidget15setDocumentModeEb(void* qthis, bool arg0);
  // proto:  int QTabWidget::heightForWidth(int width);
extern void _ZNK10QTabWidget14heightForWidthEi(void* qthis, int arg0);
  // proto:  int QTabWidget::addTab(QWidget * widget, const QString & );
extern void _ZN10QTabWidget6addTabEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QString QTabWidget::tabText(int index);
extern void _ZNK10QTabWidget7tabTextEi(void* qthis, int arg0);
  // proto:  void QTabWidget::clear();
extern void _ZN10QTabWidget5clearEv(void* qthis);
  // proto:  bool QTabWidget::hasHeightForWidth();
extern void _ZNK10QTabWidget17hasHeightForWidthEv(void* qthis);
  // proto:  QTabBar * QTabWidget::tabBar();
extern void _ZNK10QTabWidget6tabBarEv(void* qthis);
  // proto:  bool QTabWidget::tabsClosable();
extern void _ZNK10QTabWidget12tabsClosableEv(void* qthis);
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QIcon & icon, const QString & label);
extern void _ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(void* qthis, int arg0, void* arg1, void* arg2, void* arg3);
  // proto:  int QTabWidget::addTab(QWidget * widget, const QIcon & icon, const QString & label);
extern void _ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QTabWidget::setUsesScrollButtons(bool useButtons);
extern void _ZN10QTabWidget20setUsesScrollButtonsEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QTabWidget::metaObject();
extern void _ZNK10QTabWidget10metaObjectEv(void* qthis);
  // proto:  QString QTabWidget::tabToolTip(int index);
extern void _ZNK10QTabWidget10tabToolTipEi(void* qthis, int arg0);
  // proto:  QWidget * QTabWidget::currentWidget();
extern void _ZNK10QTabWidget13currentWidgetEv(void* qthis);
  // proto:  void QTabWidget::setIconSize(const QSize & size);
extern void _ZN10QTabWidget11setIconSizeERK5QSize(void* qthis, void* arg0);
  // proto:  QWidget * QTabWidget::widget(int index);
extern void _ZNK10QTabWidget6widgetEi(void* qthis, int arg0);
  // proto:  void QTabWidget::setMovable(bool movable);
extern void _ZN10QTabWidget10setMovableEb(void* qthis, bool arg0);
  // proto:  bool QTabWidget::documentMode();
extern void _ZNK10QTabWidget12documentModeEv(void* qthis);
  // proto:  QString QTabWidget::tabWhatsThis(int index);
extern void _ZNK10QTabWidget12tabWhatsThisEi(void* qthis, int arg0);
  // proto:  void QTabWidget::setTabText(int index, const QString & );
extern void _ZN10QTabWidget10setTabTextEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QTabWidget::QTabWidget(const QTabWidget & );
extern void* dector_ZN10QTabWidgetC1ERKS_(void* arg0);
extern void _ZN10QTabWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTabWidget::QTabWidget(QWidget * parent);
extern void* dector_ZN10QTabWidgetC1EP7QWidget(void* arg0);
extern void _ZN10QTabWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QTabWidget::tabBarAutoHide();
extern void _ZNK10QTabWidget14tabBarAutoHideEv(void* qthis);
  // proto:  void QTabWidget::setTabIcon(int index, const QIcon & icon);
extern void _ZN10QTabWidget10setTabIconEiRK5QIcon(void* qthis, int arg0, void* arg1);
  // proto:  QIcon QTabWidget::tabIcon(int index);
extern void _ZNK10QTabWidget7tabIconEi(void* qthis, int arg0);
  // proto:  bool QTabWidget::isTabEnabled(int index);
extern void _ZNK10QTabWidget12isTabEnabledEi(void* qthis, int arg0);
  // proto:  void QTabWidget::setTabBarAutoHide(bool enabled);
extern void _ZN10QTabWidget17setTabBarAutoHideEb(void* qthis, bool arg0);
  // proto:  QSize QTabWidget::iconSize();
extern void _ZNK10QTabWidget8iconSizeEv(void* qthis);
  // proto:  void QTabWidget::setTabsClosable(bool closeable);
extern void _ZN10QTabWidget15setTabsClosableEb(void* qthis, bool arg0);
  // proto:  QSize QTabWidget::minimumSizeHint();
extern void _ZNK10QTabWidget15minimumSizeHintEv(void* qthis);
  // proto:  void QTabWidget::setCurrentIndex(int index);
extern void _ZN10QTabWidget15setCurrentIndexEi(void* qthis, int arg0);
  // proto:  void QTabWidget::~QTabWidget();
extern void _ZN10QTabWidgetD0Ev(void* qthis);
  // proto:  void QTabWidget::setTabWhatsThis(int index, const QString & text);
extern void _ZN10QTabWidget15setTabWhatsThisEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  QSize QTabWidget::sizeHint();
extern void _ZNK10QTabWidget8sizeHintEv(void* qthis);
  // proto:  int QTabWidget::indexOf(QWidget * widget);
extern void _ZNK10QTabWidget7indexOfEP7QWidget(void* qthis, void* arg0);
  // proto:  void QTabWidget::removeTab(int index);
extern void _ZN10QTabWidget9removeTabEi(void* qthis, int arg0);
  // proto:  void QTabWidget::setTabToolTip(int index, const QString & tip);
extern void _ZN10QTabWidget13setTabToolTipEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  bool QTabWidget::isMovable();
extern void _ZNK10QTabWidget9isMovableEv(void* qthis);
  // proto:  bool QTabWidget::usesScrollButtons();
extern void _ZNK10QTabWidget17usesScrollButtonsEv(void* qthis);
  // proto:  int QTabWidget::currentIndex();
extern void _ZNK10QTabWidget12currentIndexEv(void* qthis);
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QString & );
extern void _ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  void QTabWidget::setTabEnabled(int index, bool );
extern void _ZN10QTabWidget13setTabEnabledEib(void* qthis, int arg0, bool arg1);
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

// class sizeof(QTabWidget)=1
type QTabWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _tabCloseRequested QTabWidget_tabCloseRequested_signal;
//  _tabBarDoubleClicked QTabWidget_tabBarDoubleClicked_signal;
//  _tabBarClicked QTabWidget_tabBarClicked_signal;
//  _currentChanged QTabWidget_currentChanged_signal;
}

  // proto:  void QTabWidget::setCurrentWidget(QWidget * widget);
func (this *QTabWidget) setCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentWidget", args)
  }

}

  // proto:  int QTabWidget::count();
func (this *QTabWidget) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget5countEv
    // invoke: int count()
    C._ZNK10QTabWidget5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "count", args)
  }

}

  // proto:  void QTabWidget::setDocumentMode(bool set);
func (this *QTabWidget) setDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setDocumentModeEb
    // invoke: void setDocumentMode(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setDocumentMode", args)
  }

}

  // proto:  int QTabWidget::heightForWidth(int width);
func (this *QTabWidget) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "heightForWidth", args)
  }

}

  // proto:  int QTabWidget::addTab(QWidget * widget, const QString & );
func (this *QTabWidget) addTab(args ...interface{}) () {
  // addTab(class QWidget *, const class QString &)
  // addTab(class QWidget *, const class QIcon &, const class QString &)
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
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK7QString
    // invoke: int addTab(class QWidget *, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget6addTabEP7QWidgetRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString
    // invoke: int addTab(class QWidget *, const class QIcon &, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTabWidget", "addTab", args)
  }

}

  // proto:  QString QTabWidget::tabText(int index);
func (this *QTabWidget) tabText(args ...interface{}) () {
  // tabText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7tabTextEi
    // invoke: QString tabText(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget7tabTextEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabText", args)
  }

}

  // proto:  void QTabWidget::clear();
func (this *QTabWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget5clearEv
    // invoke: void clear()
    C._ZN10QTabWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "clear", args)
  }

}

  // proto:  bool QTabWidget::hasHeightForWidth();
func (this *QTabWidget) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C._ZNK10QTabWidget17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "hasHeightForWidth", args)
  }

}

  // proto:  QTabBar * QTabWidget::tabBar();
func (this *QTabWidget) tabBar(args ...interface{}) () {
  // tabBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget6tabBarEv
    // invoke: QTabBar * tabBar()
    C._ZNK10QTabWidget6tabBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBar", args)
  }

}

  // proto:  bool QTabWidget::tabsClosable();
func (this *QTabWidget) tabsClosable(args ...interface{}) () {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12tabsClosableEv
    // invoke: bool tabsClosable()
    C._ZNK10QTabWidget12tabsClosableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabsClosable", args)
  }

}

  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QIcon & icon, const QString & label);
func (this *QTabWidget) insertTab(args ...interface{}) () {
  // insertTab(int, class QWidget *, const class QIcon &, const class QString &)
  // insertTab(int, class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString
    // invoke: int insertTab(int, class QWidget *, const class QIcon &, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QIcon).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN10QTabWidget9insertTabEiP7QWidgetRK7QString
    // invoke: int insertTab(int, class QWidget *, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTabWidget", "insertTab", args)
  }

}

  // proto:  void QTabWidget::setUsesScrollButtons(bool useButtons);
func (this *QTabWidget) setUsesScrollButtons(args ...interface{}) () {
  // setUsesScrollButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget20setUsesScrollButtonsEb
    // invoke: void setUsesScrollButtons(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget20setUsesScrollButtonsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setUsesScrollButtons", args)
  }

}

  // proto:  const QMetaObject * QTabWidget::metaObject();
func (this *QTabWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QTabWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "metaObject", args)
  }

}

  // proto:  QString QTabWidget::tabToolTip(int index);
func (this *QTabWidget) tabToolTip(args ...interface{}) () {
  // tabToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget10tabToolTipEi
    // invoke: QString tabToolTip(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget10tabToolTipEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabToolTip", args)
  }

}

  // proto:  QWidget * QTabWidget::currentWidget();
func (this *QTabWidget) currentWidget(args ...interface{}) () {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget13currentWidgetEv
    // invoke: QWidget * currentWidget()
    C._ZNK10QTabWidget13currentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "currentWidget", args)
  }

}

  // proto:  void QTabWidget::setIconSize(const QSize & size);
func (this *QTabWidget) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setIconSize", args)
  }

}

  // proto:  QWidget * QTabWidget::widget(int index);
func (this *QTabWidget) widget(args ...interface{}) () {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget6widgetEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "widget", args)
  }

}

  // proto:  void QTabWidget::setMovable(bool movable);
func (this *QTabWidget) setMovable(args ...interface{}) () {
  // setMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setMovableEb
    // invoke: void setMovable(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget10setMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setMovable", args)
  }

}

  // proto:  bool QTabWidget::documentMode();
func (this *QTabWidget) documentMode(args ...interface{}) () {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12documentModeEv
    // invoke: bool documentMode()
    C._ZNK10QTabWidget12documentModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "documentMode", args)
  }

}

  // proto:  QString QTabWidget::tabWhatsThis(int index);
func (this *QTabWidget) tabWhatsThis(args ...interface{}) () {
  // tabWhatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12tabWhatsThisEi
    // invoke: QString tabWhatsThis(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget12tabWhatsThisEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabWhatsThis", args)
  }

}

  // proto:  void QTabWidget::setTabText(int index, const QString & );
func (this *QTabWidget) setTabText(args ...interface{}) () {
  // setTabText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setTabTextEiRK7QString
    // invoke: void setTabText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget10setTabTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabText", args)
  }

}

  // proto:  void QTabWidget::QTabWidget(const QTabWidget & );
func NewQTabWidget(args ...interface{}) QTabWidget {
  return QTabWidget{}
}

  // proto:  bool QTabWidget::tabBarAutoHide();
func (this *QTabWidget) tabBarAutoHide(args ...interface{}) () {
  // tabBarAutoHide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget14tabBarAutoHideEv
    // invoke: bool tabBarAutoHide()
    C._ZNK10QTabWidget14tabBarAutoHideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBarAutoHide", args)
  }

}

  // proto:  void QTabWidget::setTabIcon(int index, const QIcon & icon);
func (this *QTabWidget) setTabIcon(args ...interface{}) () {
  // setTabIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setTabIconEiRK5QIcon
    // invoke: void setTabIcon(int, const class QIcon &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget10setTabIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabIcon", args)
  }

}

  // proto:  QIcon QTabWidget::tabIcon(int index);
func (this *QTabWidget) tabIcon(args ...interface{}) () {
  // tabIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7tabIconEi
    // invoke: QIcon tabIcon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget7tabIconEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabIcon", args)
  }

}

  // proto:  bool QTabWidget::isTabEnabled(int index);
func (this *QTabWidget) isTabEnabled(args ...interface{}) () {
  // isTabEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12isTabEnabledEi
    // invoke: bool isTabEnabled(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget12isTabEnabledEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "isTabEnabled", args)
  }

}

  // proto:  void QTabWidget::setTabBarAutoHide(bool enabled);
func (this *QTabWidget) setTabBarAutoHide(args ...interface{}) () {
  // setTabBarAutoHide(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget17setTabBarAutoHideEb
    // invoke: void setTabBarAutoHide(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget17setTabBarAutoHideEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabBarAutoHide", args)
  }

}

  // proto:  QSize QTabWidget::iconSize();
func (this *QTabWidget) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8iconSizeEv
    // invoke: QSize iconSize()
    C._ZNK10QTabWidget8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "iconSize", args)
  }

}

  // proto:  void QTabWidget::setTabsClosable(bool closeable);
func (this *QTabWidget) setTabsClosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setTabsClosableEb
    // invoke: void setTabsClosable(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget15setTabsClosableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabsClosable", args)
  }

}

  // proto:  QSize QTabWidget::minimumSizeHint();
func (this *QTabWidget) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK10QTabWidget15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "minimumSizeHint", args)
  }

}

  // proto:  void QTabWidget::setCurrentIndex(int index);
func (this *QTabWidget) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentIndex", args)
  }

}

  // proto:  void QTabWidget::~QTabWidget();
func (this *QTabWidget) FreeQTabWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTabWidget", "~QTabWidget", args)
  }

}

  // proto:  void QTabWidget::setTabWhatsThis(int index, const QString & text);
func (this *QTabWidget) setTabWhatsThis(args ...interface{}) () {
  // setTabWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setTabWhatsThisEiRK7QString
    // invoke: void setTabWhatsThis(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget15setTabWhatsThisEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabWhatsThis", args)
  }

}

  // proto:  QSize QTabWidget::sizeHint();
func (this *QTabWidget) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK10QTabWidget8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "sizeHint", args)
  }

}

  // proto:  int QTabWidget::indexOf(QWidget * widget);
func (this *QTabWidget) indexOf(args ...interface{}) () {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTabWidget7indexOfEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "indexOf", args)
  }

}

  // proto:  void QTabWidget::removeTab(int index);
func (this *QTabWidget) removeTab(args ...interface{}) () {
  // removeTab(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget9removeTabEi
    // invoke: void removeTab(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTabWidget9removeTabEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "removeTab", args)
  }

}

  // proto:  void QTabWidget::setTabToolTip(int index, const QString & tip);
func (this *QTabWidget) setTabToolTip(args ...interface{}) () {
  // setTabToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget13setTabToolTipEiRK7QString
    // invoke: void setTabToolTip(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget13setTabToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabToolTip", args)
  }

}

  // proto:  bool QTabWidget::isMovable();
func (this *QTabWidget) isMovable(args ...interface{}) () {
  // isMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget9isMovableEv
    // invoke: bool isMovable()
    C._ZNK10QTabWidget9isMovableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "isMovable", args)
  }

}

  // proto:  bool QTabWidget::usesScrollButtons();
func (this *QTabWidget) usesScrollButtons(args ...interface{}) () {
  // usesScrollButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget17usesScrollButtonsEv
    // invoke: bool usesScrollButtons()
    C._ZNK10QTabWidget17usesScrollButtonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "usesScrollButtons", args)
  }

}

  // proto:  int QTabWidget::currentIndex();
func (this *QTabWidget) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12currentIndexEv
    // invoke: int currentIndex()
    C._ZNK10QTabWidget12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "currentIndex", args)
  }

}

  // proto:  void QTabWidget::setTabEnabled(int index, bool );
func (this *QTabWidget) setTabEnabled(args ...interface{}) () {
  // setTabEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget13setTabEnabledEib
    // invoke: void setTabEnabled(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
    C._ZN10QTabWidget13setTabEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabEnabled", args)
  }

}

// <= body block end

