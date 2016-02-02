package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN5QMenu9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::addAction(const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void* C_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
extern void* C_ZN5QMenu9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void* C_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 4
  // proto:  void QMenu::setTitle(const QString & title);
extern void C_ZN5QMenu8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::actionAt(const QPoint & );
extern void* C_ZNK5QMenu8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QMenu::toolTipsVisible();
extern bool C_ZNK5QMenu15toolTipsVisibleEv(void* qthis); // 4
  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
extern void C_ZN5QMenu24setSeparatorsCollapsibleEb(void* qthis, bool arg0); // 4
  // proto:  QAction * QMenu::menuAction();
extern void* C_ZNK5QMenu10menuActionEv(void* qthis); // 4
  // proto:  QAction * QMenu::activeAction();
extern void* C_ZNK5QMenu12activeActionEv(void* qthis); // 4
  // proto:  QAction * QMenu::addSeparator();
extern void* C_ZN5QMenu12addSeparatorEv(void* qthis); // 4
  // proto:  void QMenu::setTearOffEnabled(bool );
extern void C_ZN5QMenu17setTearOffEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QMenu::setActiveAction(QAction * act);
extern void C_ZN5QMenu15setActiveActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::defaultAction();
extern void* C_ZNK5QMenu13defaultActionEv(void* qthis); // 4
  // proto:  bool QMenu::isEmpty();
extern bool C_ZNK5QMenu7isEmptyEv(void* qthis); // 4
  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
extern void* C_ZN5QMenu13insertSectionEP7QActionRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::insertSection(QAction * before, const QIcon & icon, const QString & text);
extern void* C_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QMenu::setIcon(const QIcon & icon);
extern void C_ZN5QMenu7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  bool QMenu::isTearOffEnabled();
extern bool C_ZNK5QMenu16isTearOffEnabledEv(void* qthis); // 4
  // proto:  QMenu * QMenu::addMenu(const QString & title);
extern void* C_ZN5QMenu7addMenuERK7QString(void* qthis, void* arg0); // 4
  // proto:  QMenu * QMenu::addMenu(const QIcon & icon, const QString & title);
extern void* C_ZN5QMenu7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::addMenu(QMenu * menu);
extern void* C_ZN5QMenu7addMenuEPS_(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
extern void* C_ZN5QMenu4execERK6QPointP7QAction(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QMenu::exec();
extern void* C_ZN5QMenu4execEv(void* qthis); // 4
  // proto:  QAction * QMenu::addSection(const QString & text);
extern void* C_ZN5QMenu10addSectionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::addSection(const QIcon & icon, const QString & text);
extern void* C_ZN5QMenu10addSectionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString QMenu::title();
extern void* C_ZNK5QMenu5titleEv(void* qthis); // 4
  // proto:  void QMenu::setDefaultAction(QAction * );
extern void C_ZN5QMenu16setDefaultActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
extern void* C_ZN5QMenu10insertMenuEP7QActionPS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QMenu::isTearOffMenuVisible();
extern bool C_ZNK5QMenu20isTearOffMenuVisibleEv(void* qthis); // 4
  // proto:  QIcon QMenu::icon();
extern void* C_ZNK5QMenu4iconEv(void* qthis); // 4
  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
extern void C_ZN5QMenu14setNoReplayForEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QMenu::metaObject();
extern void C_ZNK5QMenu10metaObjectEv(void* qthis); // 4
  // proto:  QSize QMenu::sizeHint();
extern void* C_ZNK5QMenu8sizeHintEv(void* qthis); // 4
  // proto:  void QMenu::clear();
extern void C_ZN5QMenu5clearEv(void* qthis); // 4
  // proto:  QAction * QMenu::insertSeparator(QAction * before);
extern void* C_ZN5QMenu15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
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
extern bool C_ZNK5QMenu21separatorsCollapsibleEv(void* qthis); // 4
  // proto:  QRect QMenu::actionGeometry(QAction * );
extern void* C_ZNK5QMenu14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToShow QMenu_aboutToShow_signal;
//  _triggered QMenu_triggered_signal;
//  _hovered QMenu_hovered_signal;
//  _aboutToHide QMenu_aboutToHide_signal;
}

// addAction(const class QString &)
func (this *QMenu) Addaction(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QMenu9addActionERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[1][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var arg3 = args[3].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu9addActionERK5QIconRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence
    // invoke: QAction * addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    argif3, free3 := qtrt.HandyConvert2c(args[3], vtys[3][3])
    var arg3 = argif3.(unsafe.Pointer)
    if false {fmt.Println(argif3, arg3)}
    if free3 {defer C.free(arg3)}
    var arg4 = args[4].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "addAction", args)
  }

  return
}

// setTitle(const class QString &)
func (this *QMenu) Settitle(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QMenu8setTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTitle", args)
  }

  return
}

// actionAt(const class QPoint &)
func (this *QMenu) Actionat(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QMenu8actionAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "actionAt", args)
  }

  return
}

// toolTipsVisible()
func (this *QMenu) Tooltipsvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu15toolTipsVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "toolTipsVisible", args)
  }

  return
}

// setSeparatorsCollapsible(_Bool)
func (this *QMenu) Setseparatorscollapsible(args ...interface{}) () {
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
    C.C_ZN5QMenu24setSeparatorsCollapsibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setSeparatorsCollapsible", args)
  }

  return
}

// menuAction()
func (this *QMenu) Menuaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu10menuActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "menuAction", args)
  }

  return
}

// activeAction()
func (this *QMenu) Activeaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu12activeActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "activeAction", args)
  }

  return
}

// addSeparator()
func (this *QMenu) Addseparator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN5QMenu12addSeparatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "addSeparator", args)
  }

  return
}

// setTearOffEnabled(_Bool)
func (this *QMenu) Settearoffenabled(args ...interface{}) () {
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
    C.C_ZN5QMenu17setTearOffEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setTearOffEnabled", args)
  }

  return
}

// setActiveAction(class QAction *)
func (this *QMenu) Setactiveaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QMenu15setActiveActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setActiveAction", args)
  }

  return
}

// defaultAction()
func (this *QMenu) Defaultaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu13defaultActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "defaultAction", args)
  }

  return
}

// isEmpty()
func (this *QMenu) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "isEmpty", args)
  }

  return
}

// insertSection(class QAction *, const class QString &)
func (this *QMenu) Insertsection(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu13insertSectionEP7QActionRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString
    // invoke: QAction * insertSection(class QAction *, const class QIcon &, const class QString &)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "insertSection", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QMenu) Seticon(args ...interface{}) () {
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
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QMenu7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setIcon", args)
  }

  return
}

// isTearOffEnabled()
func (this *QMenu) Istearoffenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu16isTearOffEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffEnabled", args)
  }

  return
}

// addMenu(const class QString &)
func (this *QMenu) Addmenu(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QMenu7addMenuERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QMenu7addMenuERK5QIconRK7QString
    // invoke: QMenu * addMenu(const class QIcon &, const class QString &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu7addMenuERK5QIconRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN5QMenu7addMenuEPS_
    // invoke: QAction * addMenu(class QMenu *)
    var arg0 = args[0].(*QMenu).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QMenu7addMenuEPS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "addMenu", args)
  }

  return
}

// exec(const class QPoint &, class QAction *)
func (this *QMenu) Exec(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAction).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu4execERK6QPointP7QAction(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QMenu4execEv
    // invoke: QAction * exec()
    var ret0 = C.C_ZN5QMenu4execEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "exec", args)
  }

  return
}

// addSection(const class QString &)
func (this *QMenu) Addsection(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QMenu10addSectionERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QMenu10addSectionERK5QIconRK7QString
    // invoke: QAction * addSection(const class QIcon &, const class QString &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu10addSectionERK5QIconRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "addSection", args)
  }

  return
}

// title()
func (this *QMenu) Title(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu5titleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "title", args)
  }

  return
}

// setDefaultAction(class QAction *)
func (this *QMenu) Setdefaultaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QMenu16setDefaultActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setDefaultAction", args)
  }

  return
}

// insertMenu(class QAction *, class QMenu *)
func (this *QMenu) Insertmenu(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMenu).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QMenu10insertMenuEP7QActionPS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "insertMenu", args)
  }

  return
}

// isTearOffMenuVisible()
func (this *QMenu) Istearoffmenuvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu20isTearOffMenuVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffMenuVisible", args)
  }

  return
}

// icon()
func (this *QMenu) Icon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "icon", args)
  }

  return
}

// setNoReplayFor(class QWidget *)
func (this *QMenu) Setnoreplayfor(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QMenu14setNoReplayForEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setNoReplayFor", args)
  }

  return
}

// metaObject()
func (this *QMenu) Metaobject(args ...interface{}) () {
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
    C.C_ZNK5QMenu10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "metaObject", args)
  }

  return
}

// sizeHint()
func (this *QMenu) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "sizeHint", args)
  }

  return
}

// clear()
func (this *QMenu) Clear(args ...interface{}) () {
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
    C.C_ZN5QMenu5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "clear", args)
  }

  return
}

// insertSeparator(class QAction *)
func (this *QMenu) Insertseparator(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QMenu15insertSeparatorEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "insertSeparator", args)
  }

  return
}

// popup(const class QPoint &, class QAction *)
func (this *QMenu) Popup(args ...interface{}) () {
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
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAction).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QMenu5popupERK6QPointP7QAction(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMenu", "popup", args)
  }

  return
}

// setToolTipsVisible(_Bool)
func (this *QMenu) Settooltipsvisible(args ...interface{}) () {
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
    C.C_ZN5QMenu18setToolTipsVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenu", "setToolTipsVisible", args)
  }

  return
}

// ~QMenu()
func (this *QMenu) Freeqmenu(args ...interface{}) () {
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
    C.C_ZN5QMenuD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "~QMenu", args)
  }

  return
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QMenuC2EP7QWidget(arg0)
    return &QMenu{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QMenuC1ERK7QStringP7QWidget
    // invoke: void QMenu(const class QString &, class QWidget *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QMenuC2ERK7QStringP7QWidget(arg0, arg1)
    return &QMenu{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMenu", "QMenu", args)
  }

  return nil // QMenu{Qclsinst:qthis}
}

// platformMenu()
func (this *QMenu) Platformmenu(args ...interface{}) () {
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
    C.C_ZN5QMenu12platformMenuEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "platformMenu", args)
  }

  return
}

// separatorsCollapsible()
func (this *QMenu) Separatorscollapsible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QMenu21separatorsCollapsibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "separatorsCollapsible", args)
  }

  return
}

// actionGeometry(class QAction *)
func (this *QMenu) Actiongeometry(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QMenu14actionGeometryEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMenu", "actionGeometry", args)
  }

  return
}

// hideTearOffMenu()
func (this *QMenu) Hidetearoffmenu(args ...interface{}) () {
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
    C.C_ZN5QMenu15hideTearOffMenuEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMenu", "hideTearOffMenu", args)
  }

  return
}

// <= body block end

