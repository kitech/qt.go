package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAction * QMenu::addAction(const QString & text);
extern void C_ZN5QMenu9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::addAction(const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void C_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, unsigned char* arg2, void* arg3); // 4
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
extern void C_ZN5QMenu9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void C_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, void* arg2, unsigned char* arg3, void* arg4); // 4
  // proto:  void QMenu::setTitle(const QString & title);
extern void C_ZN5QMenu8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::actionAt(const QPoint & );
extern void C_ZNK5QMenu8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QMenu::toolTipsVisible();
extern void C_ZNK5QMenu15toolTipsVisibleEv(void* qthis); // 4
  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
extern void C_ZN5QMenu24setSeparatorsCollapsibleEb(void* qthis, bool arg0); // 4
  // proto:  QAction * QMenu::menuAction();
extern void C_ZNK5QMenu10menuActionEv(void* qthis); // 4
  // proto:  QAction * QMenu::activeAction();
extern void C_ZNK5QMenu12activeActionEv(void* qthis); // 4
  // proto:  QAction * QMenu::addSeparator();
extern void C_ZN5QMenu12addSeparatorEv(void* qthis); // 4
  // proto:  void QMenu::setTearOffEnabled(bool );
extern void C_ZN5QMenu17setTearOffEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QMenu::setActiveAction(QAction * act);
extern void C_ZN5QMenu15setActiveActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::defaultAction();
extern void C_ZNK5QMenu13defaultActionEv(void* qthis); // 4
  // proto:  bool QMenu::isEmpty();
extern void C_ZNK5QMenu7isEmptyEv(void* qthis); // 4
  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
extern void C_ZN5QMenu13insertSectionEP7QActionRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::insertSection(QAction * before, const QIcon & icon, const QString & text);
extern void C_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QMenu::setIcon(const QIcon & icon);
extern void C_ZN5QMenu7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  bool QMenu::isTearOffEnabled();
extern void C_ZNK5QMenu16isTearOffEnabledEv(void* qthis); // 4
  // proto:  QMenu * QMenu::addMenu(const QString & title);
extern void C_ZN5QMenu7addMenuERK7QString(void* qthis, void* arg0); // 4
  // proto:  QMenu * QMenu::addMenu(const QIcon & icon, const QString & title);
extern void C_ZN5QMenu7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::addMenu(QMenu * menu);
extern void C_ZN5QMenu7addMenuEPS_(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
extern void C_ZN5QMenu4execERK6QPointP7QAction(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::exec();
extern void C_ZN5QMenu4execEv(void* qthis); // 4
  // proto:  QAction * QMenu::addSection(const QString & text);
extern void C_ZN5QMenu10addSectionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::addSection(const QIcon & icon, const QString & text);
extern void C_ZN5QMenu10addSectionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QMenu::title();
extern void C_ZNK5QMenu5titleEv(void* qthis); // 4
  // proto:  void QMenu::setDefaultAction(QAction * );
extern void C_ZN5QMenu16setDefaultActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
extern void C_ZN5QMenu10insertMenuEP7QActionPS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QMenu::isTearOffMenuVisible();
extern void C_ZNK5QMenu20isTearOffMenuVisibleEv(void* qthis); // 4
  // proto:  QIcon QMenu::icon();
extern void C_ZNK5QMenu4iconEv(void* qthis); // 4
  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
extern void C_ZN5QMenu14setNoReplayForEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QMenu::metaObject();
extern void C_ZNK5QMenu10metaObjectEv(void* qthis); // 4
  // proto:  QSize QMenu::sizeHint();
extern void C_ZNK5QMenu8sizeHintEv(void* qthis); // 4
  // proto:  void QMenu::clear();
extern void C_ZN5QMenu5clearEv(void* qthis); // 4
  // proto:  QAction * QMenu::insertSeparator(QAction * before);
extern void C_ZN5QMenu15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QMenu::popup(const QPoint & pos, QAction * at);
extern void C_ZN5QMenu5popupERK6QPointP7QAction(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMenu::setToolTipsVisible(bool visible);
extern void C_ZN5QMenu18setToolTipsVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QMenu::~QMenu();
extern void C_ZN5QMenuD2Ev(void* qthis); // 4
  // proto:  void QMenu::QMenu(QWidget * parent);
extern void* C_ZN5QMenuC2EP7QWidget(void* arg0); // 3
  // proto:  void QMenu::QMenu(const QString & title, QWidget * parent);
extern void* C_ZN5QMenuC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QPlatformMenu * QMenu::platformMenu();
extern void C_ZN5QMenu12platformMenuEv(void* qthis); // 4
  // proto:  bool QMenu::separatorsCollapsible();
extern void C_ZNK5QMenu21separatorsCollapsibleEv(void* qthis); // 4
  // proto:  QRect QMenu::actionGeometry(QAction * );
extern void C_ZNK5QMenu14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QMenu::hideTearOffMenu();
extern void C_ZN5QMenu15hideTearOffMenuEv(void* qthis); // 4
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

// addAction(const class QString &)
func (this *QMenu) addAction(args ...interface{}) () {
  // addAction(const class QString &)
  // addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1][3] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
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
    // invoke: _ZN5QMenu9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QMenu9addActionERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QKeySequence).qclsinst
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QMenu9addActionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 3:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[3].([]byte)).Pointer()))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QKeySequence).qclsinst
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "addAction", args)
  }

}

// setTitle(const class QString &)
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
    C.C_ZN5QMenu8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTitle", args)
  }

}

// actionAt(const class QPoint &)
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
    var ret = C.C_ZNK5QMenu8actionAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "actionAt", args)
  }

}

// toolTipsVisible()
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
    var ret = C.C_ZNK5QMenu15toolTipsVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "toolTipsVisible", args)
  }

}

// setSeparatorsCollapsible(_Bool)
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
    C.C_ZN5QMenu24setSeparatorsCollapsibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setSeparatorsCollapsible", args)
  }

}

// menuAction()
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
    var ret = C.C_ZNK5QMenu10menuActionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "menuAction", args)
  }

}

// activeAction()
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
    var ret = C.C_ZNK5QMenu12activeActionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "activeAction", args)
  }

}

// addSeparator()
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
    var ret = C.C_ZN5QMenu12addSeparatorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "addSeparator", args)
  }

}

// setTearOffEnabled(_Bool)
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
    C.C_ZN5QMenu17setTearOffEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTearOffEnabled", args)
  }

}

// setActiveAction(class QAction *)
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
    C.C_ZN5QMenu15setActiveActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setActiveAction", args)
  }

}

// defaultAction()
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
    var ret = C.C_ZNK5QMenu13defaultActionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "defaultAction", args)
  }

}

// isEmpty()
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
    var ret = C.C_ZNK5QMenu7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "isEmpty", args)
  }

}

// insertSection(class QAction *, const class QString &)
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
    var ret = C.C_ZN5QMenu13insertSectionEP7QActionRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString
    // invoke: QAction * insertSection(class QAction *, const class QIcon &, const class QString &)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "insertSection", args)
  }

}

// setIcon(const class QIcon &)
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
    C.C_ZN5QMenu7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setIcon", args)
  }

}

// isTearOffEnabled()
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
    var ret = C.C_ZNK5QMenu16isTearOffEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffEnabled", args)
  }

}

// addMenu(const class QString &)
func (this *QMenu) addMenu(args ...interface{}) () {
  // addMenu(const class QString &)
  // addMenu(const class QIcon &, const class QString &)
  // addMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu7addMenuERK7QString
    // invoke: QMenu * addMenu(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QMenu7addMenuERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QMenu7addMenuERK5QIconRK7QString
    // invoke: QMenu * addMenu(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QMenu7addMenuERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZN5QMenu7addMenuEPS_
    // invoke: QAction * addMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN5QMenu7addMenuEPS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "addMenu", args)
  }

}

// exec(const class QPoint &, class QAction *)
func (this *QMenu) exec(args ...interface{}) () {
  // exec(const class QPoint &, class QAction *)
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1] = make(map[int32]reflect.Type)

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
    var ret = C.C_ZN5QMenu4execERK6QPointP7QAction(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QMenu4execEv
    // invoke: QAction * exec()
    var ret = C.C_ZN5QMenu4execEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "exec", args)
  }

}

// addSection(const class QString &)
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
    var ret = C.C_ZN5QMenu10addSectionERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN5QMenu10addSectionERK5QIconRK7QString
    // invoke: QAction * addSection(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN5QMenu10addSectionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "addSection", args)
  }

}

// title()
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
    var ret = C.C_ZNK5QMenu5titleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "title", args)
  }

}

// setDefaultAction(class QAction *)
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
    C.C_ZN5QMenu16setDefaultActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setDefaultAction", args)
  }

}

// insertMenu(class QAction *, class QMenu *)
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
    var ret = C.C_ZN5QMenu10insertMenuEP7QActionPS_(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "insertMenu", args)
  }

}

// isTearOffMenuVisible()
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
    var ret = C.C_ZNK5QMenu20isTearOffMenuVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffMenuVisible", args)
  }

}

// icon()
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
    var ret = C.C_ZNK5QMenu4iconEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "icon", args)
  }

}

// setNoReplayFor(class QWidget *)
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
    C.C_ZN5QMenu14setNoReplayForEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setNoReplayFor", args)
  }

}

// metaObject()
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
    C.C_ZNK5QMenu10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "metaObject", args)
  }

}

// sizeHint()
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
    var ret = C.C_ZNK5QMenu8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "sizeHint", args)
  }

}

// clear()
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
    C.C_ZN5QMenu5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "clear", args)
  }

}

// insertSeparator(class QAction *)
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
    var ret = C.C_ZN5QMenu15insertSeparatorEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "insertSeparator", args)
  }

}

// popup(const class QPoint &, class QAction *)
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
    C.C_ZN5QMenu5popupERK6QPointP7QAction(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "popup", args)
  }

}

// setToolTipsVisible(_Bool)
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
    C.C_ZN5QMenu18setToolTipsVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setToolTipsVisible", args)
  }

}

// ~QMenu()
func (this *QMenu) FreeQMenu(args ...interface{}) () {
  // ~QMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenuD0Ev
    // invoke: void ~QMenu()
    C.C_ZN5QMenuD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "~QMenu", args)
  }

}

// QMenu(class QWidget *)
func NewQMenu(args ...interface{}) *QMenu {
  // QMenu(class QWidget *)
  // QMenu(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenuC1EP7QWidget
    // invoke: void QMenu(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QMenuC2EP7QWidget(arg0)
    return &QMenu{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QMenuC1ERK7QStringP7QWidget
    // invoke: void QMenu(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QMenuC2ERK7QStringP7QWidget(arg0, arg1)
    return &QMenu{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMenu", "QMenu", args)
  }

  return nil // QMenu{qclsinst:qthis}
}

// platformMenu()
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
    C.C_ZN5QMenu12platformMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "platformMenu", args)
  }

}

// separatorsCollapsible()
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
    var ret = C.C_ZNK5QMenu21separatorsCollapsibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "separatorsCollapsible", args)
  }

}

// actionGeometry(class QAction *)
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
    var ret = C.C_ZNK5QMenu14actionGeometryEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenu", "actionGeometry", args)
  }

}

// hideTearOffMenu()
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
    C.C_ZN5QMenu15hideTearOffMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "hideTearOffMenu", args)
  }

}

// <= body block end

