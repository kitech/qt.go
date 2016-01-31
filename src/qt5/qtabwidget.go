package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QTabWidget::tabToolTip(int index);
extern void C_ZNK10QTabWidget10tabToolTipEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabWidget::tabBarAutoHide();
extern void C_ZNK10QTabWidget14tabBarAutoHideEv(void* qthis); // 4
  // proto:  bool QTabWidget::isTabEnabled(int index);
extern void C_ZNK10QTabWidget12isTabEnabledEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QTabWidget::tabWhatsThis(int index);
extern void C_ZNK10QTabWidget12tabWhatsThisEi(void* qthis, int32_t arg0); // 4
  // proto:  QIcon QTabWidget::tabIcon(int index);
extern void C_ZNK10QTabWidget7tabIconEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::minimumSizeHint();
extern void C_ZNK10QTabWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QTabWidget::setCurrentIndex(int index);
extern void C_ZN10QTabWidget15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QTabWidget::currentWidget();
extern void C_ZNK10QTabWidget13currentWidgetEv(void* qthis); // 4
  // proto:  void QTabWidget::setMovable(bool movable);
extern void C_ZN10QTabWidget10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::tabsClosable();
extern void C_ZNK10QTabWidget12tabsClosableEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabIcon(int index, const QIcon & icon);
extern void C_ZN10QTabWidget10setTabIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::setCurrentWidget(QWidget * widget);
extern void C_ZN10QTabWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QTabWidget::~QTabWidget();
extern void C_ZN10QTabWidgetD2Ev(void* qthis); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QString & );
extern void C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QIcon & icon, const QString & label);
extern void C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QTabWidget::TabShape QTabWidget::tabShape();
extern void C_ZNK10QTabWidget8tabShapeEv(void* qthis); // 4
  // proto:  int QTabWidget::indexOf(QWidget * widget);
extern void C_ZNK10QTabWidget7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QTabBar * QTabWidget::tabBar();
extern void C_ZNK10QTabWidget6tabBarEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabToolTip(int index, const QString & tip);
extern void C_ZN10QTabWidget13setTabToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::QTabWidget(QWidget * parent);
extern void* C_ZN10QTabWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QTabWidget::setTabBarAutoHide(bool enabled);
extern void C_ZN10QTabWidget17setTabBarAutoHideEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::documentMode();
extern void C_ZNK10QTabWidget12documentModeEv(void* qthis); // 4
  // proto:  QWidget * QTabWidget::widget(int index);
extern void C_ZNK10QTabWidget6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::removeTab(int index);
extern void C_ZN10QTabWidget9removeTabEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::iconSize();
extern void C_ZNK10QTabWidget8iconSizeEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabText(int index, const QString & );
extern void C_ZN10QTabWidget10setTabTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QTabWidget::tabText(int index);
extern void C_ZNK10QTabWidget7tabTextEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabsClosable(bool closeable);
extern void C_ZN10QTabWidget15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  QSize QTabWidget::sizeHint();
extern void C_ZNK10QTabWidget8sizeHintEv(void* qthis); // 4
  // proto:  QTabWidget::TabPosition QTabWidget::tabPosition();
extern void C_ZNK10QTabWidget11tabPositionEv(void* qthis); // 4
  // proto:  bool QTabWidget::usesScrollButtons();
extern void C_ZNK10QTabWidget17usesScrollButtonsEv(void* qthis); // 4
  // proto:  void QTabWidget::setDocumentMode(bool set);
extern void C_ZN10QTabWidget15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  int QTabWidget::count();
extern void C_ZNK10QTabWidget5countEv(void* qthis); // 4
  // proto:  void QTabWidget::setUsesScrollButtons(bool useButtons);
extern void C_ZN10QTabWidget20setUsesScrollButtonsEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTabWidget::metaObject();
extern void C_ZNK10QTabWidget10metaObjectEv(void* qthis); // 4
  // proto:  Qt::TextElideMode QTabWidget::elideMode();
extern void C_ZNK10QTabWidget9elideModeEv(void* qthis); // 4
  // proto:  bool QTabWidget::hasHeightForWidth();
extern void C_ZNK10QTabWidget17hasHeightForWidthEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabWhatsThis(int index, const QString & text);
extern void C_ZN10QTabWidget15setTabWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::clear();
extern void C_ZN10QTabWidget5clearEv(void* qthis); // 4
  // proto:  void QTabWidget::setIconSize(const QSize & size);
extern void C_ZN10QTabWidget11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QTabWidget::isMovable();
extern void C_ZNK10QTabWidget9isMovableEv(void* qthis); // 4
  // proto:  int QTabWidget::heightForWidth(int width);
extern void C_ZNK10QTabWidget14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabEnabled(int index, bool );
extern void C_ZN10QTabWidget13setTabEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  int QTabWidget::currentIndex();
extern void C_ZNK10QTabWidget12currentIndexEv(void* qthis); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QIcon & icon, const QString & label);
extern void C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QString & );
extern void C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _tabCloseRequested QTabWidget_tabCloseRequested_signal;
//  _tabBarDoubleClicked QTabWidget_tabBarDoubleClicked_signal;
//  _tabBarClicked QTabWidget_tabBarClicked_signal;
//  _currentChanged QTabWidget_currentChanged_signal;
}

// tabToolTip(int)
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
    var ret = C.C_ZNK10QTabWidget10tabToolTipEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabToolTip", args)
  }

}

// tabBarAutoHide()
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
    var ret = C.C_ZNK10QTabWidget14tabBarAutoHideEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBarAutoHide", args)
  }

}

// isTabEnabled(int)
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
    var ret = C.C_ZNK10QTabWidget12isTabEnabledEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "isTabEnabled", args)
  }

}

// tabWhatsThis(int)
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
    var ret = C.C_ZNK10QTabWidget12tabWhatsThisEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabWhatsThis", args)
  }

}

// tabIcon(int)
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
    var ret = C.C_ZNK10QTabWidget7tabIconEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabIcon", args)
  }

}

// minimumSizeHint()
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
    var ret = C.C_ZNK10QTabWidget15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "minimumSizeHint", args)
  }

}

// setCurrentIndex(int)
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
    C.C_ZN10QTabWidget15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentIndex", args)
  }

}

// currentWidget()
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
    var ret = C.C_ZNK10QTabWidget13currentWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "currentWidget", args)
  }

}

// setMovable(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget10setMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setMovable", args)
  }

}

// tabsClosable()
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
    var ret = C.C_ZNK10QTabWidget12tabsClosableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabsClosable", args)
  }

}

// setTabIcon(int, const class QIcon &)
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
    C.C_ZN10QTabWidget10setTabIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabIcon", args)
  }

}

// setCurrentWidget(class QWidget *)
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
    C.C_ZN10QTabWidget16setCurrentWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentWidget", args)
  }

}

// ~QTabWidget()
func (this *QTabWidget) FreeQTabWidget(args ...interface{}) () {
  // ~QTabWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidgetD0Ev
    // invoke: void ~QTabWidget()
    C.C_ZN10QTabWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "~QTabWidget", args)
  }

}

// insertTab(int, class QWidget *, const class QString &)
func (this *QTabWidget) insertTab(args ...interface{}) () {
  // insertTab(int, class QWidget *, const class QString &)
  // insertTab(int, class QWidget *, const class QIcon &, const class QString &)
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
    // invoke: _ZN10QTabWidget9insertTabEiP7QWidgetRK7QString
    // invoke: int insertTab(int, class QWidget *, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  case 1:
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
    var ret = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "insertTab", args)
  }

}

// tabShape()
func (this *QTabWidget) tabShape(args ...interface{}) () {
  // tabShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8tabShapeEv
    // invoke: QTabWidget::TabShape tabShape()
    C.C_ZNK10QTabWidget8tabShapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabShape", args)
  }

}

// indexOf(class QWidget *)
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
    var ret = C.C_ZNK10QTabWidget7indexOfEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "indexOf", args)
  }

}

// tabBar()
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
    var ret = C.C_ZNK10QTabWidget6tabBarEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBar", args)
  }

}

// setTabToolTip(int, const class QString &)
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
    C.C_ZN10QTabWidget13setTabToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabToolTip", args)
  }

}

// QTabWidget(class QWidget *)
func NewQTabWidget(args ...interface{}) *QTabWidget {
  // QTabWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidgetC1EP7QWidget
    // invoke: void QTabWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTabWidgetC2EP7QWidget(arg0)
    return &QTabWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTabWidget", "QTabWidget", args)
  }

  return nil // QTabWidget{qclsinst:qthis}
}

// setTabBarAutoHide(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget17setTabBarAutoHideEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabBarAutoHide", args)
  }

}

// documentMode()
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
    var ret = C.C_ZNK10QTabWidget12documentModeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "documentMode", args)
  }

}

// widget(int)
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
    var ret = C.C_ZNK10QTabWidget6widgetEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "widget", args)
  }

}

// removeTab(int)
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
    C.C_ZN10QTabWidget9removeTabEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "removeTab", args)
  }

}

// iconSize()
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
    var ret = C.C_ZNK10QTabWidget8iconSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "iconSize", args)
  }

}

// setTabText(int, const class QString &)
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
    C.C_ZN10QTabWidget10setTabTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabText", args)
  }

}

// tabText(int)
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
    var ret = C.C_ZNK10QTabWidget7tabTextEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "tabText", args)
  }

}

// setTabsClosable(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget15setTabsClosableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabsClosable", args)
  }

}

// sizeHint()
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
    var ret = C.C_ZNK10QTabWidget8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "sizeHint", args)
  }

}

// tabPosition()
func (this *QTabWidget) tabPosition(args ...interface{}) () {
  // tabPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget11tabPositionEv
    // invoke: QTabWidget::TabPosition tabPosition()
    C.C_ZNK10QTabWidget11tabPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabPosition", args)
  }

}

// usesScrollButtons()
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
    var ret = C.C_ZNK10QTabWidget17usesScrollButtonsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "usesScrollButtons", args)
  }

}

// setDocumentMode(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setDocumentMode", args)
  }

}

// count()
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
    var ret = C.C_ZNK10QTabWidget5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "count", args)
  }

}

// setUsesScrollButtons(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget20setUsesScrollButtonsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setUsesScrollButtons", args)
  }

}

// metaObject()
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
    C.C_ZNK10QTabWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "metaObject", args)
  }

}

// elideMode()
func (this *QTabWidget) elideMode(args ...interface{}) () {
  // elideMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget9elideModeEv
    // invoke: Qt::TextElideMode elideMode()
    C.C_ZNK10QTabWidget9elideModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "elideMode", args)
  }

}

// hasHeightForWidth()
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
    var ret = C.C_ZNK10QTabWidget17hasHeightForWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "hasHeightForWidth", args)
  }

}

// setTabWhatsThis(int, const class QString &)
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
    C.C_ZN10QTabWidget15setTabWhatsThisEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabWhatsThis", args)
  }

}

// clear()
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
    C.C_ZN10QTabWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "clear", args)
  }

}

// setIconSize(const class QSize &)
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
    C.C_ZN10QTabWidget11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setIconSize", args)
  }

}

// isMovable()
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
    var ret = C.C_ZNK10QTabWidget9isMovableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "isMovable", args)
  }

}

// heightForWidth(int)
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
    var ret = C.C_ZNK10QTabWidget14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "heightForWidth", args)
  }

}

// setTabEnabled(int, _Bool)
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
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget13setTabEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabEnabled", args)
  }

}

// currentIndex()
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
    var ret = C.C_ZNK10QTabWidget12currentIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "currentIndex", args)
  }

}

// addTab(class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) addTab(args ...interface{}) () {
  // addTab(class QWidget *, const class QIcon &, const class QString &)
  // addTab(class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString
    // invoke: int addTab(class QWidget *, const class QIcon &, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK7QString
    // invoke: int addTab(class QWidget *, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTabWidget", "addTab", args)
  }

}

// <= body block end

