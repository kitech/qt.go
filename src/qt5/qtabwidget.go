package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK10QTabWidget10tabToolTipEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabWidget::tabBarAutoHide();
extern bool C_ZNK10QTabWidget14tabBarAutoHideEv(void* qthis); // 4
  // proto:  bool QTabWidget::isTabEnabled(int index);
extern bool C_ZNK10QTabWidget12isTabEnabledEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QTabWidget::tabWhatsThis(int index);
extern void* C_ZNK10QTabWidget12tabWhatsThisEi(void* qthis, int32_t arg0); // 4
  // proto:  QIcon QTabWidget::tabIcon(int index);
extern void* C_ZNK10QTabWidget7tabIconEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::minimumSizeHint();
extern void* C_ZNK10QTabWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QTabWidget::setCurrentIndex(int index);
extern void C_ZN10QTabWidget15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QTabWidget::currentWidget();
extern void* C_ZNK10QTabWidget13currentWidgetEv(void* qthis); // 4
  // proto:  void QTabWidget::setMovable(bool movable);
extern void C_ZN10QTabWidget10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::tabsClosable();
extern bool C_ZNK10QTabWidget12tabsClosableEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabIcon(int index, const QIcon & icon);
extern void C_ZN10QTabWidget10setTabIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::setCurrentWidget(QWidget * widget);
extern void C_ZN10QTabWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QTabWidget::~QTabWidget();
extern void C_ZN10QTabWidgetD2Ev(void* qthis); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QString & );
extern int32_t C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QIcon & icon, const QString & label);
extern int32_t C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QTabWidget::TabShape QTabWidget::tabShape();
extern void C_ZNK10QTabWidget8tabShapeEv(void* qthis); // 4
  // proto:  int QTabWidget::indexOf(QWidget * widget);
extern int32_t C_ZNK10QTabWidget7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QTabBar * QTabWidget::tabBar();
extern void* C_ZNK10QTabWidget6tabBarEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabToolTip(int index, const QString & tip);
extern void C_ZN10QTabWidget13setTabToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::QTabWidget(QWidget * parent);
extern void* C_ZN10QTabWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QTabWidget::setTabBarAutoHide(bool enabled);
extern void C_ZN10QTabWidget17setTabBarAutoHideEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::documentMode();
extern bool C_ZNK10QTabWidget12documentModeEv(void* qthis); // 4
  // proto:  QWidget * QTabWidget::widget(int index);
extern void* C_ZNK10QTabWidget6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::removeTab(int index);
extern void C_ZN10QTabWidget9removeTabEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::iconSize();
extern void* C_ZNK10QTabWidget8iconSizeEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabText(int index, const QString & );
extern void C_ZN10QTabWidget10setTabTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QTabWidget::tabText(int index);
extern void* C_ZNK10QTabWidget7tabTextEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabsClosable(bool closeable);
extern void C_ZN10QTabWidget15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  QSize QTabWidget::sizeHint();
extern void* C_ZNK10QTabWidget8sizeHintEv(void* qthis); // 4
  // proto:  QTabWidget::TabPosition QTabWidget::tabPosition();
extern void C_ZNK10QTabWidget11tabPositionEv(void* qthis); // 4
  // proto:  bool QTabWidget::usesScrollButtons();
extern bool C_ZNK10QTabWidget17usesScrollButtonsEv(void* qthis); // 4
  // proto:  void QTabWidget::setDocumentMode(bool set);
extern void C_ZN10QTabWidget15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  int QTabWidget::count();
extern int32_t C_ZNK10QTabWidget5countEv(void* qthis); // 4
  // proto:  void QTabWidget::setUsesScrollButtons(bool useButtons);
extern void C_ZN10QTabWidget20setUsesScrollButtonsEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTabWidget::metaObject();
extern void C_ZNK10QTabWidget10metaObjectEv(void* qthis); // 4
  // proto:  Qt::TextElideMode QTabWidget::elideMode();
extern void C_ZNK10QTabWidget9elideModeEv(void* qthis); // 4
  // proto:  bool QTabWidget::hasHeightForWidth();
extern bool C_ZNK10QTabWidget17hasHeightForWidthEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabWhatsThis(int index, const QString & text);
extern void C_ZN10QTabWidget15setTabWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::clear();
extern void C_ZN10QTabWidget5clearEv(void* qthis); // 4
  // proto:  void QTabWidget::setIconSize(const QSize & size);
extern void C_ZN10QTabWidget11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QTabWidget::isMovable();
extern bool C_ZNK10QTabWidget9isMovableEv(void* qthis); // 4
  // proto:  int QTabWidget::heightForWidth(int width);
extern int32_t C_ZNK10QTabWidget14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabEnabled(int index, bool );
extern void C_ZN10QTabWidget13setTabEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  int QTabWidget::currentIndex();
extern int32_t C_ZNK10QTabWidget12currentIndexEv(void* qthis); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QIcon & icon, const QString & label);
extern int32_t C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QString & );
extern int32_t C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1); // 4
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
func (this *QTabWidget) Tabtooltip(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget10tabToolTipEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabToolTip", args)
  }

  return
}

// tabBarAutoHide()
func (this *QTabWidget) Tabbarautohide(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget14tabBarAutoHideEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBarAutoHide", args)
  }

  return
}

// isTabEnabled(int)
func (this *QTabWidget) Istabenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget12isTabEnabledEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "isTabEnabled", args)
  }

  return
}

// tabWhatsThis(int)
func (this *QTabWidget) Tabwhatsthis(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget12tabWhatsThisEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabWhatsThis", args)
  }

  return
}

// tabIcon(int)
func (this *QTabWidget) Tabicon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget7tabIconEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabIcon", args)
  }

  return
}

// minimumSizeHint()
func (this *QTabWidget) Minimumsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "minimumSizeHint", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QTabWidget) Setcurrentindex(args ...interface{}) () {
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

  return
}

// currentWidget()
func (this *QTabWidget) Currentwidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget13currentWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "currentWidget", args)
  }

  return
}

// setMovable(_Bool)
func (this *QTabWidget) Setmovable(args ...interface{}) () {
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

  return
}

// tabsClosable()
func (this *QTabWidget) Tabsclosable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget12tabsClosableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabsClosable", args)
  }

  return
}

// setTabIcon(int, const class QIcon &)
func (this *QTabWidget) Settabicon(args ...interface{}) () {
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

  return
}

// setCurrentWidget(class QWidget *)
func (this *QTabWidget) Setcurrentwidget(args ...interface{}) () {
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

  return
}

// ~QTabWidget()
func (this *QTabWidget) Freeqtabwidget(args ...interface{}) () {
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

  return
}

// insertTab(int, class QWidget *, const class QString &)
func (this *QTabWidget) Inserttab(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
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
    var ret0 = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "insertTab", args)
  }

  return
}

// tabShape()
func (this *QTabWidget) Tabshape(args ...interface{}) () {
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

  return
}

// indexOf(class QWidget *)
func (this *QTabWidget) Indexof(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget7indexOfEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "indexOf", args)
  }

  return
}

// tabBar()
func (this *QTabWidget) Tabbar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget6tabBarEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTabBar{}) // "QTabBar *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBar", args)
  }

  return
}

// setTabToolTip(int, const class QString &)
func (this *QTabWidget) Settabtooltip(args ...interface{}) () {
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

  return
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
func (this *QTabWidget) Settabbarautohide(args ...interface{}) () {
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

  return
}

// documentMode()
func (this *QTabWidget) Documentmode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget12documentModeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "documentMode", args)
  }

  return
}

// widget(int)
func (this *QTabWidget) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget6widgetEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "widget", args)
  }

  return
}

// removeTab(int)
func (this *QTabWidget) Removetab(args ...interface{}) () {
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

  return
}

// iconSize()
func (this *QTabWidget) Iconsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget8iconSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "iconSize", args)
  }

  return
}

// setTabText(int, const class QString &)
func (this *QTabWidget) Settabtext(args ...interface{}) () {
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

  return
}

// tabText(int)
func (this *QTabWidget) Tabtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget7tabTextEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "tabText", args)
  }

  return
}

// setTabsClosable(_Bool)
func (this *QTabWidget) Settabsclosable(args ...interface{}) () {
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

  return
}

// sizeHint()
func (this *QTabWidget) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "sizeHint", args)
  }

  return
}

// tabPosition()
func (this *QTabWidget) Tabposition(args ...interface{}) () {
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

  return
}

// usesScrollButtons()
func (this *QTabWidget) Usesscrollbuttons(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget17usesScrollButtonsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "usesScrollButtons", args)
  }

  return
}

// setDocumentMode(_Bool)
func (this *QTabWidget) Setdocumentmode(args ...interface{}) () {
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

  return
}

// count()
func (this *QTabWidget) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "count", args)
  }

  return
}

// setUsesScrollButtons(_Bool)
func (this *QTabWidget) Setusesscrollbuttons(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QTabWidget) Metaobject(args ...interface{}) () {
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

  return
}

// elideMode()
func (this *QTabWidget) Elidemode(args ...interface{}) () {
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

  return
}

// hasHeightForWidth()
func (this *QTabWidget) Hasheightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget17hasHeightForWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "hasHeightForWidth", args)
  }

  return
}

// setTabWhatsThis(int, const class QString &)
func (this *QTabWidget) Settabwhatsthis(args ...interface{}) () {
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

  return
}

// clear()
func (this *QTabWidget) Clear(args ...interface{}) () {
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

  return
}

// setIconSize(const class QSize &)
func (this *QTabWidget) Seticonsize(args ...interface{}) () {
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

  return
}

// isMovable()
func (this *QTabWidget) Ismovable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget9isMovableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "isMovable", args)
  }

  return
}

// heightForWidth(int)
func (this *QTabWidget) Heightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "heightForWidth", args)
  }

  return
}

// setTabEnabled(int, _Bool)
func (this *QTabWidget) Settabenabled(args ...interface{}) () {
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

  return
}

// currentIndex()
func (this *QTabWidget) Currentindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK10QTabWidget12currentIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "currentIndex", args)
  }

  return
}

// addTab(class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) Addtab(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK7QString
    // invoke: int addTab(class QWidget *, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabWidget", "addTab", args)
  }

  return
}

// <= body block end

