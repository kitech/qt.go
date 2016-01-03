package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qmenu.h
// dst-file: /src/widgets/qmenu.go
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
  // proto:  bool QMenu::isTearOffEnabled();
extern void _ZNK5QMenu16isTearOffEnabledEv(void* qthis);
  // proto:  bool QMenu::toolTipsVisible();
extern void _ZNK5QMenu15toolTipsVisibleEv(void* qthis);
  // proto:  QAction * QMenu::menuAction();
extern void _ZNK5QMenu10menuActionEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
extern void _ZN5QMenu9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QMenu::setTearOffEnabled(bool );
extern void _ZN5QMenu17setTearOffEnabledEb(void* qthis, bool arg0);
  // proto:  QAction * QMenu::addSection(const QString & text);
extern void _ZN5QMenu10addSectionERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QMenu::metaObject();
extern void _ZNK5QMenu10metaObjectEv(void* qthis);
  // proto:  void QMenu::clear();
extern void _ZN5QMenu5clearEv(void* qthis);
  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
extern void _ZN5QMenu10insertMenuEP7QActionPS_(void* qthis, void* arg0, void* arg1);
  // proto:  QIcon QMenu::icon();
extern void _ZNK5QMenu4iconEv(void* qthis);
  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
extern void _ZN5QMenu13insertSectionEP7QActionRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QPlatformMenu * QMenu::platformMenu();
extern void _ZN5QMenu12platformMenuEv(void* qthis);
  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
extern void _ZN5QMenu14setNoReplayForEP7QWidget(void* qthis, void* arg0);
  // proto:  void QMenu::setIcon(const QIcon & icon);
extern void _ZN5QMenu7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
extern void _ZN5QMenu4execERK6QPointP7QAction(void* qthis, void* arg0, void* arg1);
  // proto:  bool QMenu::separatorsCollapsible();
extern void _ZNK5QMenu21separatorsCollapsibleEv(void* qthis);
  // proto:  QMenu * QMenu::addMenu(const QString & title);
extern void _ZN5QMenu7addMenuERK7QString(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addSeparator();
extern void _ZN5QMenu12addSeparatorEv(void* qthis);
  // proto:  void QMenu::hideTearOffMenu();
extern void _ZN5QMenu15hideTearOffMenuEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, unsigned char* arg2, void* arg3);
  // proto:  void QMenu::QMenu(QWidget * parent);
extern void* dector_ZN5QMenuC1EP7QWidget(void* arg0);
extern void _ZN5QMenuC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QMenu::setActiveAction(QAction * act);
extern void _ZN5QMenu15setActiveActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
extern void _ZN5QMenu24setSeparatorsCollapsibleEb(void* qthis, bool arg0);
  // proto:  void QMenu::QMenu(const QMenu & );
extern void* dector_ZN5QMenuC1ERKS_(void* arg0);
extern void _ZN5QMenuC1ERKS_(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addAction(const QString & text);
extern void _ZN5QMenu9addActionERK7QString(void* qthis, void* arg0);
  // proto:  QAction * QMenu::activeAction();
extern void _ZNK5QMenu12activeActionEv(void* qthis);
  // proto:  bool QMenu::isEmpty();
extern void _ZNK5QMenu7isEmptyEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, void* arg2, unsigned char* arg3, void* arg4);
  // proto:  QRect QMenu::actionGeometry(QAction * );
extern void _ZNK5QMenu14actionGeometryEP7QAction(void* qthis, void* arg0);
  // proto:  void QMenu::QMenu(const QString & title, QWidget * parent);
extern void* dector_ZN5QMenuC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN5QMenuC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  QAction * QMenu::insertSeparator(QAction * before);
extern void _ZN5QMenu15insertSeparatorEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addSection(const QIcon & icon, const QString & text);
extern void _ZN5QMenu10addSectionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  bool QMenu::isTearOffMenuVisible();
extern void _ZNK5QMenu20isTearOffMenuVisibleEv(void* qthis);
  // proto:  void QMenu::~QMenu();
extern void _ZN5QMenuD0Ev(void* qthis);
  // proto:  QString QMenu::title();
extern void _ZNK5QMenu5titleEv(void* qthis);
  // proto:  QAction * QMenu::defaultAction();
extern void _ZNK5QMenu13defaultActionEv(void* qthis);
  // proto:  QAction * QMenu::addMenu(QMenu * menu);
extern void _ZN5QMenu7addMenuEPS_(void* qthis, void* arg0);
  // proto:  QSize QMenu::sizeHint();
extern void _ZNK5QMenu8sizeHintEv(void* qthis);
  // proto:  void QMenu::setDefaultAction(QAction * );
extern void _ZN5QMenu16setDefaultActionEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenu::actionAt(const QPoint & );
extern void _ZNK5QMenu8actionAtERK6QPoint(void* qthis, void* arg0);
  // proto:  QAction * QMenu::insertSection(QAction * before, const QIcon & icon, const QString & text);
extern void _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QMenu::popup(const QPoint & pos, QAction * at);
extern void _ZN5QMenu5popupERK6QPointP7QAction(void* qthis, void* arg0, void* arg1);
  // proto:  void QMenu::setToolTipsVisible(bool visible);
extern void _ZN5QMenu18setToolTipsVisibleEb(void* qthis, bool arg0);
  // proto:  void QMenu::setTitle(const QString & title);
extern void _ZN5QMenu8setTitleERK7QString(void* qthis, void* arg0);
  // proto:  QMenu * QMenu::addMenu(const QIcon & icon, const QString & title);
extern void _ZN5QMenu7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QAction * QMenu::exec();
extern void _ZN5QMenu4execEv(void* qthis);
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

// class sizeof(QMenu)=1
type QMenu struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToShow QMenu_aboutToShow_signal;
//  _triggered QMenu_triggered_signal;
//  _hovered QMenu_hovered_signal;
//  _aboutToHide QMenu_aboutToHide_signal;
}

  // proto:  bool QMenu::isTearOffEnabled();
func (this *QMenu) isTearOffEnabled(args ...interface{}) () {
  // isTearOffEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu16isTearOffEnabledEv
    // invoke: bool isTearOffEnabled()
    C._ZNK5QMenu16isTearOffEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffEnabled", args)
  }

}

  // proto:  bool QMenu::toolTipsVisible();
func (this *QMenu) toolTipsVisible(args ...interface{}) () {
  // toolTipsVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu15toolTipsVisibleEv
    // invoke: bool toolTipsVisible()
    C._ZNK5QMenu15toolTipsVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "toolTipsVisible", args)
  }

}

  // proto:  QAction * QMenu::menuAction();
func (this *QMenu) menuAction(args ...interface{}) () {
  // menuAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu10menuActionEv
    // invoke: QAction * menuAction()
    C._ZNK5QMenu10menuActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "menuAction", args)
  }

}

  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
func (this *QMenu) addAction(args ...interface{}) () {
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
  // addAction(const class QString &)
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1][3] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][3] = qtrt.ByteTy(true) // "const char *"
  vtys[3][4] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu9addActionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QKeySequence).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN5QMenu9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu9addActionERK7QString(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[3].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QKeySequence).qclsinst
    if false {fmt.Println(arg4)}
    C._ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QMenu", "addAction", args)
  }

}

  // proto:  void QMenu::setTearOffEnabled(bool );
func (this *QMenu) setTearOffEnabled(args ...interface{}) () {
  // setTearOffEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu17setTearOffEnabledEb
    // invoke: void setTearOffEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN5QMenu17setTearOffEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTearOffEnabled", args)
  }

}

  // proto:  QAction * QMenu::addSection(const QString & text);
func (this *QMenu) addSection(args ...interface{}) () {
  // addSection(const class QString &)
  // addSection(const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu10addSectionERK7QString
    // invoke: QAction * addSection(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu10addSectionERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QMenu10addSectionERK5QIconRK7QString
    // invoke: QAction * addSection(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu10addSectionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "addSection", args)
  }

}

  // proto:  const QMetaObject * QMenu::metaObject();
func (this *QMenu) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK5QMenu10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "metaObject", args)
  }

}

  // proto:  void QMenu::clear();
func (this *QMenu) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu5clearEv
    // invoke: void clear()
    C._ZN5QMenu5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "clear", args)
  }

}

  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
func (this *QMenu) insertMenu(args ...interface{}) () {
  // insertMenu(class QAction *, class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu10insertMenuEP7QActionPS_
    // invoke: QAction * insertMenu(class QAction *, class QMenu *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMenu).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu10insertMenuEP7QActionPS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "insertMenu", args)
  }

}

  // proto:  QIcon QMenu::icon();
func (this *QMenu) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu4iconEv
    // invoke: QIcon icon()
    C._ZNK5QMenu4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "icon", args)
  }

}

  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
func (this *QMenu) insertSection(args ...interface{}) () {
  // insertSection(class QAction *, const class QString &)
  // insertSection(class QAction *, const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK7QString
    // invoke: QAction * insertSection(class QAction *, const class QString &)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu13insertSectionEP7QActionRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString
    // invoke: QAction * insertSection(class QAction *, const class QIcon &, const class QString &)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QMenu", "insertSection", args)
  }

}

  // proto:  QPlatformMenu * QMenu::platformMenu();
func (this *QMenu) platformMenu(args ...interface{}) () {
  // platformMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu12platformMenuEv
    // invoke: QPlatformMenu * platformMenu()
    C._ZN5QMenu12platformMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "platformMenu", args)
  }

}

  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
func (this *QMenu) setNoReplayFor(args ...interface{}) () {
  // setNoReplayFor(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu14setNoReplayForEP7QWidget
    // invoke: void setNoReplayFor(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu14setNoReplayForEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setNoReplayFor", args)
  }

}

  // proto:  void QMenu::setIcon(const QIcon & icon);
func (this *QMenu) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setIcon", args)
  }

}

  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
func (this *QMenu) exec(args ...interface{}) () {
  // exec(const class QPoint &, class QAction *)
  // exec(QList<class QAction *>, const class QPoint &, class QAction *, class QWidget *)
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "QList<QAction *>"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][2] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1][3] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu4execERK6QPointP7QAction
    // invoke: QAction * exec(const class QPoint &, class QAction *)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu4execERK6QPointP7QAction(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN5QMenu4execEv
    // invoke: QAction * exec()
    C._ZN5QMenu4execEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "exec", args)
  }

}

  // proto:  bool QMenu::separatorsCollapsible();
func (this *QMenu) separatorsCollapsible(args ...interface{}) () {
  // separatorsCollapsible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu21separatorsCollapsibleEv
    // invoke: bool separatorsCollapsible()
    C._ZNK5QMenu21separatorsCollapsibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "separatorsCollapsible", args)
  }

}

  // proto:  QMenu * QMenu::addMenu(const QString & title);
func (this *QMenu) addMenu(args ...interface{}) () {
  // addMenu(const class QString &)
  // addMenu(class QMenu *)
  // addMenu(const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu7addMenuERK7QString
    // invoke: QMenu * addMenu(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu7addMenuERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN5QMenu7addMenuEPS_
    // invoke: QAction * addMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu7addMenuEPS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN5QMenu7addMenuERK5QIconRK7QString
    // invoke: QMenu * addMenu(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu7addMenuERK5QIconRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "addMenu", args)
  }

}

  // proto:  QAction * QMenu::addSeparator();
func (this *QMenu) addSeparator(args ...interface{}) () {
  // addSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu12addSeparatorEv
    // invoke: QAction * addSeparator()
    C._ZN5QMenu12addSeparatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "addSeparator", args)
  }

}

  // proto:  void QMenu::hideTearOffMenu();
func (this *QMenu) hideTearOffMenu(args ...interface{}) () {
  // hideTearOffMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15hideTearOffMenuEv
    // invoke: void hideTearOffMenu()
    C._ZN5QMenu15hideTearOffMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "hideTearOffMenu", args)
  }

}

  // proto:  void QMenu::QMenu(QWidget * parent);
func NewQMenu(args ...interface{}) QMenu {
  return QMenu{}
}

  // proto:  void QMenu::setActiveAction(QAction * act);
func (this *QMenu) setActiveAction(args ...interface{}) () {
  // setActiveAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15setActiveActionEP7QAction
    // invoke: void setActiveAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu15setActiveActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setActiveAction", args)
  }

}

  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
func (this *QMenu) setSeparatorsCollapsible(args ...interface{}) () {
  // setSeparatorsCollapsible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu24setSeparatorsCollapsibleEb
    // invoke: void setSeparatorsCollapsible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN5QMenu24setSeparatorsCollapsibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setSeparatorsCollapsible", args)
  }

}

  // proto:  QAction * QMenu::activeAction();
func (this *QMenu) activeAction(args ...interface{}) () {
  // activeAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu12activeActionEv
    // invoke: QAction * activeAction()
    C._ZNK5QMenu12activeActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "activeAction", args)
  }

}

  // proto:  bool QMenu::isEmpty();
func (this *QMenu) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK5QMenu7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "isEmpty", args)
  }

}

  // proto:  QRect QMenu::actionGeometry(QAction * );
func (this *QMenu) actionGeometry(args ...interface{}) () {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu14actionGeometryEP7QAction
    // invoke: QRect actionGeometry(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK5QMenu14actionGeometryEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "actionGeometry", args)
  }

}

  // proto:  QAction * QMenu::insertSeparator(QAction * before);
func (this *QMenu) insertSeparator(args ...interface{}) () {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15insertSeparatorEP7QAction
    // invoke: QAction * insertSeparator(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu15insertSeparatorEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "insertSeparator", args)
  }

}

  // proto:  bool QMenu::isTearOffMenuVisible();
func (this *QMenu) isTearOffMenuVisible(args ...interface{}) () {
  // isTearOffMenuVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu20isTearOffMenuVisibleEv
    // invoke: bool isTearOffMenuVisible()
    C._ZNK5QMenu20isTearOffMenuVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffMenuVisible", args)
  }

}

  // proto:  void QMenu::~QMenu();
func (this *QMenu) FreeQMenu(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMenu", "~QMenu", args)
  }

}

  // proto:  QString QMenu::title();
func (this *QMenu) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu5titleEv
    // invoke: QString title()
    C._ZNK5QMenu5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "title", args)
  }

}

  // proto:  QAction * QMenu::defaultAction();
func (this *QMenu) defaultAction(args ...interface{}) () {
  // defaultAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu13defaultActionEv
    // invoke: QAction * defaultAction()
    C._ZNK5QMenu13defaultActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "defaultAction", args)
  }

}

  // proto:  QSize QMenu::sizeHint();
func (this *QMenu) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK5QMenu8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "sizeHint", args)
  }

}

  // proto:  void QMenu::setDefaultAction(QAction * );
func (this *QMenu) setDefaultAction(args ...interface{}) () {
  // setDefaultAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu16setDefaultActionEP7QAction
    // invoke: void setDefaultAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu16setDefaultActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setDefaultAction", args)
  }

}

  // proto:  QAction * QMenu::actionAt(const QPoint & );
func (this *QMenu) actionAt(args ...interface{}) () {
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu8actionAtERK6QPoint
    // invoke: QAction * actionAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK5QMenu8actionAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "actionAt", args)
  }

}

  // proto:  void QMenu::popup(const QPoint & pos, QAction * at);
func (this *QMenu) popup(args ...interface{}) () {
  // popup(const class QPoint &, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu5popupERK6QPointP7QAction
    // invoke: void popup(const class QPoint &, class QAction *)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QMenu5popupERK6QPointP7QAction(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "popup", args)
  }

}

  // proto:  void QMenu::setToolTipsVisible(bool visible);
func (this *QMenu) setToolTipsVisible(args ...interface{}) () {
  // setToolTipsVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu18setToolTipsVisibleEb
    // invoke: void setToolTipsVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN5QMenu18setToolTipsVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setToolTipsVisible", args)
  }

}

  // proto:  void QMenu::setTitle(const QString & title);
func (this *QMenu) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QMenu8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTitle", args)
  }

}

// <= body block end

