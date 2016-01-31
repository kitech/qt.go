package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qtabbar.h
// dst-file: /src/widgets/qtabbar.go
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
  // proto:  QString QTabBar::tabToolTip(int index);
extern void* C_ZNK7QTabBar10tabToolTipEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabBar::setIconSize(const QSize & size);
extern void C_ZN7QTabBar11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QTabBar::usesScrollButtons();
extern bool C_ZNK7QTabBar17usesScrollButtonsEv(void* qthis); // 4
  // proto:  bool QTabBar::isTabEnabled(int index);
extern bool C_ZNK7QTabBar12isTabEnabledEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabBar::expanding();
extern bool C_ZNK7QTabBar9expandingEv(void* qthis); // 4
  // proto:  QString QTabBar::tabWhatsThis(int index);
extern void* C_ZNK7QTabBar12tabWhatsThisEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabBar::QTabBar(QWidget * parent);
extern void* C_ZN7QTabBarC2EP7QWidget(void* arg0); // 3
  // proto:  int QTabBar::insertTab(int index, const QIcon & icon, const QString & text);
extern int32_t C_ZN7QTabBar9insertTabEiRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabBar::insertTab(int index, const QString & text);
extern int32_t C_ZN7QTabBar9insertTabEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTabBar::Shape QTabBar::shape();
extern void C_ZNK7QTabBar5shapeEv(void* qthis); // 4
  // proto:  void QTabBar::setCurrentIndex(int index);
extern void C_ZN7QTabBar15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabBar::setTabTextColor(int index, const QColor & color);
extern void C_ZN7QTabBar15setTabTextColorEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabBar::setMovable(bool movable);
extern void C_ZN7QTabBar10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  void QTabBar::setDrawBase(bool drawTheBase);
extern void C_ZN7QTabBar11setDrawBaseEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabBar::tabsClosable();
extern bool C_ZNK7QTabBar12tabsClosableEv(void* qthis); // 4
  // proto:  void QTabBar::setTabIcon(int index, const QIcon & icon);
extern void C_ZN7QTabBar10setTabIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  bool QTabBar::drawBase();
extern bool C_ZNK7QTabBar8drawBaseEv(void* qthis); // 4
  // proto:  void QTabBar::~QTabBar();
extern void C_ZN7QTabBarD2Ev(void* qthis); // 4
  // proto:  void QTabBar::setTabData(int index, const QVariant & data);
extern void C_ZN7QTabBar10setTabDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QIcon QTabBar::tabIcon(int index);
extern void* C_ZNK7QTabBar7tabIconEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabBar::autoHide();
extern bool C_ZNK7QTabBar8autoHideEv(void* qthis); // 4
  // proto:  QColor QTabBar::tabTextColor(int index);
extern void* C_ZNK7QTabBar12tabTextColorEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabBar::isMovable();
extern bool C_ZNK7QTabBar9isMovableEv(void* qthis); // 4
  // proto:  QRect QTabBar::tabRect(int index);
extern void* C_ZNK7QTabBar7tabRectEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabBar::documentMode();
extern bool C_ZNK7QTabBar12documentModeEv(void* qthis); // 4
  // proto:  void QTabBar::setTabsClosable(bool closable);
extern void C_ZN7QTabBar15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  void QTabBar::removeTab(int index);
extern void C_ZN7QTabBar9removeTabEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabBar::iconSize();
extern void* C_ZNK7QTabBar8iconSizeEv(void* qthis); // 4
  // proto:  void QTabBar::setTabText(int index, const QString & text);
extern void C_ZN7QTabBar10setTabTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  int QTabBar::tabAt(const QPoint & pos);
extern int32_t C_ZNK7QTabBar5tabAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QTabBar::changeCurrentOnDrag();
extern bool C_ZNK7QTabBar19changeCurrentOnDragEv(void* qthis); // 4
  // proto:  QString QTabBar::tabText(int index);
extern void* C_ZNK7QTabBar7tabTextEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTabBar::addTab(const QIcon & icon, const QString & text);
extern int32_t C_ZN7QTabBar6addTabERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  int QTabBar::addTab(const QString & text);
extern int32_t C_ZN7QTabBar6addTabERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTabBar::setAutoHide(bool hide);
extern void C_ZN7QTabBar11setAutoHideEb(void* qthis, bool arg0); // 4
  // proto:  void QTabBar::setDocumentMode(bool set);
extern void C_ZN7QTabBar15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  int QTabBar::count();
extern int32_t C_ZNK7QTabBar5countEv(void* qthis); // 4
  // proto:  void QTabBar::setUsesScrollButtons(bool useButtons);
extern void C_ZN7QTabBar20setUsesScrollButtonsEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTabBar::metaObject();
extern void C_ZNK7QTabBar10metaObjectEv(void* qthis); // 4
  // proto:  Qt::TextElideMode QTabBar::elideMode();
extern void C_ZNK7QTabBar9elideModeEv(void* qthis); // 4
  // proto:  QSize QTabBar::sizeHint();
extern void* C_ZNK7QTabBar8sizeHintEv(void* qthis); // 4
  // proto:  QSize QTabBar::minimumSizeHint();
extern void* C_ZNK7QTabBar15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QTabBar::setChangeCurrentOnDrag(bool change);
extern void C_ZN7QTabBar22setChangeCurrentOnDragEb(void* qthis, bool arg0); // 4
  // proto:  void QTabBar::setTabWhatsThis(int index, const QString & text);
extern void C_ZN7QTabBar15setTabWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabBar::setTabToolTip(int index, const QString & tip);
extern void C_ZN7QTabBar13setTabToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabBar::setExpanding(bool enabled);
extern void C_ZN7QTabBar12setExpandingEb(void* qthis, bool arg0); // 4
  // proto:  void QTabBar::setTabEnabled(int index, bool );
extern void C_ZN7QTabBar13setTabEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  QTabBar::SelectionBehavior QTabBar::selectionBehaviorOnRemove();
extern void C_ZNK7QTabBar25selectionBehaviorOnRemoveEv(void* qthis); // 4
  // proto:  QVariant QTabBar::tabData(int index);
extern void* C_ZNK7QTabBar7tabDataEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTabBar::currentIndex();
extern int32_t C_ZNK7QTabBar12currentIndexEv(void* qthis); // 4
  // proto:  void QTabBar::moveTab(int from, int to);
extern void C_ZN7QTabBar7moveTabEii(void* qthis, int32_t arg0, int32_t arg1); // 4
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

// class sizeof(QTabBar)=1
type QTabBar struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _tabCloseRequested QTabBar_tabCloseRequested_signal;
//  _tabBarDoubleClicked QTabBar_tabBarDoubleClicked_signal;
//  _tabMoved QTabBar_tabMoved_signal;
//  _tabBarClicked QTabBar_tabBarClicked_signal;
//  _currentChanged QTabBar_currentChanged_signal;
}

// tabToolTip(int)
func (this *QTabBar) Tabtooltip(args ...interface{}) (ret interface{}) {
  // tabToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar10tabToolTipEi
    // invoke: QString tabToolTip(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar10tabToolTipEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabToolTip", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QTabBar) Seticonsize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setIconSize", args)
  }

  return
}

// usesScrollButtons()
func (this *QTabBar) Usesscrollbuttons(args ...interface{}) (ret interface{}) {
  // usesScrollButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar17usesScrollButtonsEv
    // invoke: bool usesScrollButtons()
    var ret0 = C.C_ZNK7QTabBar17usesScrollButtonsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "usesScrollButtons", args)
  }

  return
}

// isTabEnabled(int)
func (this *QTabBar) Istabenabled(args ...interface{}) (ret interface{}) {
  // isTabEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12isTabEnabledEi
    // invoke: bool isTabEnabled(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar12isTabEnabledEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "isTabEnabled", args)
  }

  return
}

// expanding()
func (this *QTabBar) Expanding(args ...interface{}) (ret interface{}) {
  // expanding()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar9expandingEv
    // invoke: bool expanding()
    var ret0 = C.C_ZNK7QTabBar9expandingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "expanding", args)
  }

  return
}

// tabWhatsThis(int)
func (this *QTabBar) Tabwhatsthis(args ...interface{}) (ret interface{}) {
  // tabWhatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabWhatsThisEi
    // invoke: QString tabWhatsThis(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar12tabWhatsThisEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabWhatsThis", args)
  }

  return
}

// QTabBar(class QWidget *)
func NewQTabBar(args ...interface{}) *QTabBar {
  // QTabBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBarC1EP7QWidget
    // invoke: void QTabBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QTabBarC2EP7QWidget(arg0)
    return &QTabBar{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTabBar", "QTabBar", args)
  }

  return nil // QTabBar{qclsinst:qthis}
}

// insertTab(int, const class QIcon &, const class QString &)
func (this *QTabBar) Inserttab(args ...interface{}) (ret interface{}) {
  // insertTab(int, const class QIcon &, const class QString &)
  // insertTab(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar9insertTabEiRK5QIconRK7QString
    // invoke: int insertTab(int, const class QIcon &, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QTabBar9insertTabEiRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN7QTabBar9insertTabEiRK7QString
    // invoke: int insertTab(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QTabBar9insertTabEiRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "insertTab", args)
  }

  return
}

// shape()
func (this *QTabBar) Shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar5shapeEv
    // invoke: QTabBar::Shape shape()
    C.C_ZNK7QTabBar5shapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabBar", "shape", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QTabBar) Setcurrentindex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setCurrentIndex", args)
  }

  return
}

// setTabTextColor(int, const class QColor &)
func (this *QTabBar) Settabtextcolor(args ...interface{}) () {
  // setTabTextColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setTabTextColorEiRK6QColor
    // invoke: void setTabTextColor(int, const class QColor &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar15setTabTextColorEiRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabTextColor", args)
  }

  return
}

// setMovable(_Bool)
func (this *QTabBar) Setmovable(args ...interface{}) () {
  // setMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setMovableEb
    // invoke: void setMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar10setMovableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setMovable", args)
  }

  return
}

// setDrawBase(_Bool)
func (this *QTabBar) Setdrawbase(args ...interface{}) () {
  // setDrawBase(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setDrawBaseEb
    // invoke: void setDrawBase(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar11setDrawBaseEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setDrawBase", args)
  }

  return
}

// tabsClosable()
func (this *QTabBar) Tabsclosable(args ...interface{}) (ret interface{}) {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabsClosableEv
    // invoke: bool tabsClosable()
    var ret0 = C.C_ZNK7QTabBar12tabsClosableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabsClosable", args)
  }

  return
}

// setTabIcon(int, const class QIcon &)
func (this *QTabBar) Settabicon(args ...interface{}) () {
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
    // invoke: _ZN7QTabBar10setTabIconEiRK5QIcon
    // invoke: void setTabIcon(int, const class QIcon &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar10setTabIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabIcon", args)
  }

  return
}

// drawBase()
func (this *QTabBar) Drawbase(args ...interface{}) (ret interface{}) {
  // drawBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8drawBaseEv
    // invoke: bool drawBase()
    var ret0 = C.C_ZNK7QTabBar8drawBaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "drawBase", args)
  }

  return
}

// ~QTabBar()
func (this *QTabBar) Freeqtabbar(args ...interface{}) () {
  // ~QTabBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBarD0Ev
    // invoke: void ~QTabBar()
    C.C_ZN7QTabBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabBar", "~QTabBar", args)
  }

  return
}

// setTabData(int, const class QVariant &)
func (this *QTabBar) Settabdata(args ...interface{}) () {
  // setTabData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar10setTabDataEiRK8QVariant
    // invoke: void setTabData(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar10setTabDataEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabData", args)
  }

  return
}

// tabIcon(int)
func (this *QTabBar) Tabicon(args ...interface{}) (ret interface{}) {
  // tabIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabIconEi
    // invoke: QIcon tabIcon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar7tabIconEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabIcon", args)
  }

  return
}

// autoHide()
func (this *QTabBar) Autohide(args ...interface{}) (ret interface{}) {
  // autoHide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8autoHideEv
    // invoke: bool autoHide()
    var ret0 = C.C_ZNK7QTabBar8autoHideEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "autoHide", args)
  }

  return
}

// tabTextColor(int)
func (this *QTabBar) Tabtextcolor(args ...interface{}) (ret interface{}) {
  // tabTextColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12tabTextColorEi
    // invoke: QColor tabTextColor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar12tabTextColorEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabTextColor", args)
  }

  return
}

// isMovable()
func (this *QTabBar) Ismovable(args ...interface{}) (ret interface{}) {
  // isMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar9isMovableEv
    // invoke: bool isMovable()
    var ret0 = C.C_ZNK7QTabBar9isMovableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "isMovable", args)
  }

  return
}

// tabRect(int)
func (this *QTabBar) Tabrect(args ...interface{}) (ret interface{}) {
  // tabRect(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabRectEi
    // invoke: QRect tabRect(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar7tabRectEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabRect", args)
  }

  return
}

// documentMode()
func (this *QTabBar) Documentmode(args ...interface{}) (ret interface{}) {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12documentModeEv
    // invoke: bool documentMode()
    var ret0 = C.C_ZNK7QTabBar12documentModeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "documentMode", args)
  }

  return
}

// setTabsClosable(_Bool)
func (this *QTabBar) Settabsclosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setTabsClosableEb
    // invoke: void setTabsClosable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar15setTabsClosableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabsClosable", args)
  }

  return
}

// removeTab(int)
func (this *QTabBar) Removetab(args ...interface{}) () {
  // removeTab(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar9removeTabEi
    // invoke: void removeTab(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar9removeTabEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "removeTab", args)
  }

  return
}

// iconSize()
func (this *QTabBar) Iconsize(args ...interface{}) (ret interface{}) {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8iconSizeEv
    // invoke: QSize iconSize()
    var ret0 = C.C_ZNK7QTabBar8iconSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "iconSize", args)
  }

  return
}

// setTabText(int, const class QString &)
func (this *QTabBar) Settabtext(args ...interface{}) () {
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
    // invoke: _ZN7QTabBar10setTabTextEiRK7QString
    // invoke: void setTabText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar10setTabTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabText", args)
  }

  return
}

// tabAt(const class QPoint &)
func (this *QTabBar) Tabat(args ...interface{}) (ret interface{}) {
  // tabAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar5tabAtERK6QPoint
    // invoke: int tabAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar5tabAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabAt", args)
  }

  return
}

// changeCurrentOnDrag()
func (this *QTabBar) Changecurrentondrag(args ...interface{}) (ret interface{}) {
  // changeCurrentOnDrag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar19changeCurrentOnDragEv
    // invoke: bool changeCurrentOnDrag()
    var ret0 = C.C_ZNK7QTabBar19changeCurrentOnDragEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "changeCurrentOnDrag", args)
  }

  return
}

// tabText(int)
func (this *QTabBar) Tabtext(args ...interface{}) (ret interface{}) {
  // tabText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabTextEi
    // invoke: QString tabText(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar7tabTextEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabText", args)
  }

  return
}

// addTab(const class QIcon &, const class QString &)
func (this *QTabBar) Addtab(args ...interface{}) (ret interface{}) {
  // addTab(const class QIcon &, const class QString &)
  // addTab(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar6addTabERK5QIconRK7QString
    // invoke: int addTab(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QTabBar6addTabERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN7QTabBar6addTabERK7QString
    // invoke: int addTab(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QTabBar6addTabERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "addTab", args)
  }

  return
}

// setAutoHide(_Bool)
func (this *QTabBar) Setautohide(args ...interface{}) () {
  // setAutoHide(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar11setAutoHideEb
    // invoke: void setAutoHide(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar11setAutoHideEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setAutoHide", args)
  }

  return
}

// setDocumentMode(_Bool)
func (this *QTabBar) Setdocumentmode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar15setDocumentModeEb
    // invoke: void setDocumentMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar15setDocumentModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setDocumentMode", args)
  }

  return
}

// count()
func (this *QTabBar) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK7QTabBar5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "count", args)
  }

  return
}

// setUsesScrollButtons(_Bool)
func (this *QTabBar) Setusesscrollbuttons(args ...interface{}) () {
  // setUsesScrollButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar20setUsesScrollButtonsEb
    // invoke: void setUsesScrollButtons(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar20setUsesScrollButtonsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setUsesScrollButtons", args)
  }

  return
}

// metaObject()
func (this *QTabBar) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QTabBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabBar", "metaObject", args)
  }

  return
}

// elideMode()
func (this *QTabBar) Elidemode(args ...interface{}) () {
  // elideMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar9elideModeEv
    // invoke: Qt::TextElideMode elideMode()
    C.C_ZNK7QTabBar9elideModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabBar", "elideMode", args)
  }

  return
}

// sizeHint()
func (this *QTabBar) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK7QTabBar8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "sizeHint", args)
  }

  return
}

// minimumSizeHint()
func (this *QTabBar) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK7QTabBar15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "minimumSizeHint", args)
  }

  return
}

// setChangeCurrentOnDrag(_Bool)
func (this *QTabBar) Setchangecurrentondrag(args ...interface{}) () {
  // setChangeCurrentOnDrag(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar22setChangeCurrentOnDragEb
    // invoke: void setChangeCurrentOnDrag(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar22setChangeCurrentOnDragEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setChangeCurrentOnDrag", args)
  }

  return
}

// setTabWhatsThis(int, const class QString &)
func (this *QTabBar) Settabwhatsthis(args ...interface{}) () {
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
    // invoke: _ZN7QTabBar15setTabWhatsThisEiRK7QString
    // invoke: void setTabWhatsThis(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar15setTabWhatsThisEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabWhatsThis", args)
  }

  return
}

// setTabToolTip(int, const class QString &)
func (this *QTabBar) Settabtooltip(args ...interface{}) () {
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
    // invoke: _ZN7QTabBar13setTabToolTipEiRK7QString
    // invoke: void setTabToolTip(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar13setTabToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabToolTip", args)
  }

  return
}

// setExpanding(_Bool)
func (this *QTabBar) Setexpanding(args ...interface{}) () {
  // setExpanding(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar12setExpandingEb
    // invoke: void setExpanding(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QTabBar12setExpandingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabBar", "setExpanding", args)
  }

  return
}

// setTabEnabled(int, _Bool)
func (this *QTabBar) Settabenabled(args ...interface{}) () {
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
    // invoke: _ZN7QTabBar13setTabEnabledEib
    // invoke: void setTabEnabled(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar13setTabEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "setTabEnabled", args)
  }

  return
}

// selectionBehaviorOnRemove()
func (this *QTabBar) Selectionbehavioronremove(args ...interface{}) () {
  // selectionBehaviorOnRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar25selectionBehaviorOnRemoveEv
    // invoke: QTabBar::SelectionBehavior selectionBehaviorOnRemove()
    C.C_ZNK7QTabBar25selectionBehaviorOnRemoveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTabBar", "selectionBehaviorOnRemove", args)
  }

  return
}

// tabData(int)
func (this *QTabBar) Tabdata(args ...interface{}) (ret interface{}) {
  // tabData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar7tabDataEi
    // invoke: QVariant tabData(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QTabBar7tabDataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "tabData", args)
  }

  return
}

// currentIndex()
func (this *QTabBar) Currentindex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QTabBar12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK7QTabBar12currentIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTabBar", "currentIndex", args)
  }

  return
}

// moveTab(int, int)
func (this *QTabBar) Movetab(args ...interface{}) () {
  // moveTab(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QTabBar7moveTabEii
    // invoke: void moveTab(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QTabBar7moveTabEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabBar", "moveTab", args)
  }

  return
}

// <= body block end

