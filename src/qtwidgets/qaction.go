package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qaction.h
// dst-file: /src/widgets/qaction.go
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QList<QWidget *> QAction::associatedWidgets();
extern void C_ZNK7QAction17associatedWidgetsEv(void* qthis); // 4
  // proto:  void QAction::setIconVisibleInMenu(bool visible);
extern void C_ZN7QAction20setIconVisibleInMenuEb(void* qthis, bool arg0); // 4
  // proto:  QList<QKeySequence> QAction::shortcuts();
extern void C_ZNK7QAction9shortcutsEv(void* qthis); // 4
  // proto:  void QAction::setChecked(bool );
extern void C_ZN7QAction10setCheckedEb(void* qthis, bool arg0); // 4
  // proto:  void QAction::setMenu(QMenu * menu);
extern void C_ZN7QAction7setMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QAction::setStatusTip(const QString & statusTip);
extern void C_ZN7QAction12setStatusTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QAction::toggle();
extern void C_ZN7QAction6toggleEv(void* qthis); // 4
  // proto:  bool QAction::isCheckable();
extern bool C_ZNK7QAction11isCheckableEv(void* qthis); // 4
  // proto:  bool QAction::isSeparator();
extern bool C_ZNK7QAction11isSeparatorEv(void* qthis); // 4
  // proto:  void QAction::~QAction();
extern void C_ZN7QActionD2Ev(void* qthis); // 4
  // proto:  bool QAction::showStatusText(QWidget * widget);
extern bool C_ZN7QAction14showStatusTextEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QFont QAction::font();
extern void* C_ZNK7QAction4fontEv(void* qthis); // 4
  // proto:  void QAction::setActionGroup(QActionGroup * group);
extern void C_ZN7QAction14setActionGroupEP12QActionGroup(void* qthis, void* arg0); // 4
  // proto:  QMenu * QAction::menu();
extern void* C_ZNK7QAction4menuEv(void* qthis); // 4
  // proto:  QKeySequence QAction::shortcut();
extern void* C_ZNK7QAction8shortcutEv(void* qthis); // 4
  // proto:  QString QAction::whatsThis();
extern void* C_ZNK7QAction9whatsThisEv(void* qthis); // 4
  // proto:  bool QAction::isChecked();
extern bool C_ZNK7QAction9isCheckedEv(void* qthis); // 4
  // proto:  void QAction::trigger();
extern void C_ZN7QAction7triggerEv(void* qthis); // 2
  // proto:  void QAction::setCheckable(bool );
extern void C_ZN7QAction12setCheckableEb(void* qthis, bool arg0); // 4
  // proto:  bool QAction::isIconVisibleInMenu();
extern bool C_ZNK7QAction19isIconVisibleInMenuEv(void* qthis); // 4
  // proto:  QIcon QAction::icon();
extern void* C_ZNK7QAction4iconEv(void* qthis); // 4
  // proto:  void QAction::setIcon(const QIcon & icon);
extern void C_ZN7QAction7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  void QAction::setData(const QVariant & var);
extern void C_ZN7QAction7setDataERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  void QAction::setEnabled(bool );
extern void C_ZN7QAction10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QActionGroup * QAction::actionGroup();
extern void* C_ZNK7QAction11actionGroupEv(void* qthis); // 4
  // proto:  void QAction::hover();
extern void C_ZN7QAction5hoverEv(void* qthis); // 2
  // proto:  void QAction::setIconText(const QString & text);
extern void C_ZN7QAction11setIconTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QAction::statusTip();
extern void* C_ZNK7QAction9statusTipEv(void* qthis); // 4
  // proto:  void QAction::setAutoRepeat(bool );
extern void C_ZN7QAction13setAutoRepeatEb(void* qthis, bool arg0); // 4
  // proto:  void QAction::setShortcut(const QKeySequence & shortcut);
extern void C_ZN7QAction11setShortcutERK12QKeySequence(void* qthis, void* arg0); // 4
  // proto:  QString QAction::toolTip();
extern void* C_ZNK7QAction7toolTipEv(void* qthis); // 4
  // proto:  bool QAction::isEnabled();
extern bool C_ZNK7QAction9isEnabledEv(void* qthis); // 4
  // proto:  void QAction::setWhatsThis(const QString & what);
extern void C_ZN7QAction12setWhatsThisERK7QString(void* qthis, void* arg0); // 4
  // proto:  QVariant QAction::data();
extern void* C_ZNK7QAction4dataEv(void* qthis); // 4
  // proto:  void QAction::setToolTip(const QString & tip);
extern void C_ZN7QAction10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QAction::metaObject();
extern void C_ZNK7QAction10metaObjectEv(void* qthis); // 4
  // proto:  void QAction::setSeparator(bool b);
extern void C_ZN7QAction12setSeparatorEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QAction::parentWidget();
extern void* C_ZNK7QAction12parentWidgetEv(void* qthis); // 4
  // proto:  void QAction::setText(const QString & text);
extern void C_ZN7QAction7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QAction::QAction(QObject * parent);
extern void* C_ZN7QActionC2EP7QObject(void* arg0); // 3
  // proto:  void QAction::QAction(const QIcon & icon, const QString & text, QObject * parent);
extern void* C_ZN7QActionC2ERK5QIconRK7QStringP7QObject(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QAction::QAction(const QString & text, QObject * parent);
extern void* C_ZN7QActionC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  QString QAction::iconText();
extern void* C_ZNK7QAction8iconTextEv(void* qthis); // 4
  // proto:  Qt::ShortcutContext QAction::shortcutContext();
extern void C_ZNK7QAction15shortcutContextEv(void* qthis); // 4
  // proto:  void QAction::setDisabled(bool b);
extern void C_ZN7QAction11setDisabledEb(void* qthis, bool arg0); // 2
  // proto:  QString QAction::text();
extern void* C_ZNK7QAction4textEv(void* qthis); // 4
  // proto:  bool QAction::autoRepeat();
extern bool C_ZNK7QAction10autoRepeatEv(void* qthis); // 4
  // proto:  bool QAction::isVisible();
extern bool C_ZNK7QAction9isVisibleEv(void* qthis); // 4
  // proto:  void QAction::setFont(const QFont & font);
extern void C_ZN7QAction7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  QAction::MenuRole QAction::menuRole();
extern void C_ZNK7QAction8menuRoleEv(void* qthis); // 4
  // proto:  void QAction::setVisible(bool );
extern void C_ZN7QAction10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QList<QGraphicsWidget *> QAction::associatedGraphicsWidgets();
extern void C_ZNK7QAction25associatedGraphicsWidgetsEv(void* qthis); // 4
  // proto:  QAction::Priority QAction::priority();
extern void C_ZNK7QAction8priorityEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QAction)=1
type QAction struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _changed QAction_changed_signal;
//  _hovered QAction_hovered_signal;
//  _triggered QAction_triggered_signal;
//  _toggled QAction_toggled_signal;
}

// associatedWidgets()
func (this *QAction) Associatedwidgets(args ...interface{}) () {
  // associatedWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction17associatedWidgetsEv
    // invoke: QList<QWidget *> associatedWidgets()
    C.C_ZNK7QAction17associatedWidgetsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "associatedWidgets", args)
  }

  return
}

// setIconVisibleInMenu(_Bool)
func (this *QAction) Seticonvisibleinmenu(args ...interface{}) () {
  // setIconVisibleInMenu(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction20setIconVisibleInMenuEb
    // invoke: void setIconVisibleInMenu(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction20setIconVisibleInMenuEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setIconVisibleInMenu", args)
  }

  return
}

// shortcuts()
func (this *QAction) Shortcuts(args ...interface{}) () {
  // shortcuts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9shortcutsEv
    // invoke: QList<QKeySequence> shortcuts()
    C.C_ZNK7QAction9shortcutsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "shortcuts", args)
  }

  return
}

// setChecked(_Bool)
func (this *QAction) Setchecked(args ...interface{}) () {
  // setChecked(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction10setCheckedEb
    // invoke: void setChecked(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction10setCheckedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setChecked", args)
  }

  return
}

// setMenu(class QMenu *)
func (this *QAction) Setmenu(args ...interface{}) () {
  // setMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setMenuEP5QMenu
    // invoke: void setMenu(class QMenu *)
    var arg0 = args[0].(*QMenu).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction7setMenuEP5QMenu(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setMenu", args)
  }

  return
}

// setStatusTip(const class QString &)
func (this *QAction) Setstatustip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction12setStatusTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setStatusTip", args)
  }

  return
}

// toggle()
func (this *QAction) Toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction6toggleEv
    // invoke: void toggle()
    C.C_ZN7QAction6toggleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "toggle", args)
  }

  return
}

// isCheckable()
func (this *QAction) Ischeckable(args ...interface{}) (ret interface{}) {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11isCheckableEv
    // invoke: bool isCheckable()
    var ret0 = C.C_ZNK7QAction11isCheckableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isCheckable", args)
  }

  return
}

// isSeparator()
func (this *QAction) Isseparator(args ...interface{}) (ret interface{}) {
  // isSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11isSeparatorEv
    // invoke: bool isSeparator()
    var ret0 = C.C_ZNK7QAction11isSeparatorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isSeparator", args)
  }

  return
}

// ~QAction()
func (this *QAction) Freeqaction(args ...interface{}) () {
  // ~QAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QActionD0Ev
    // invoke: void ~QAction()
    C.C_ZN7QActionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "~QAction", args)
  }

  return
}

// showStatusText(class QWidget *)
func (this *QAction) Showstatustext(args ...interface{}) (ret interface{}) {
  // showStatusText(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction14showStatusTextEP7QWidget
    // invoke: bool showStatusText(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QAction14showStatusTextEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "showStatusText", args)
  }

  return
}

// font()
func (this *QAction) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK7QAction4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "font", args)
  }

  return
}

// setActionGroup(class QActionGroup *)
func (this *QAction) Setactiongroup(args ...interface{}) () {
  // setActionGroup(class QActionGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QActionGroup{}) // "QActionGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction14setActionGroupEP12QActionGroup
    // invoke: void setActionGroup(class QActionGroup *)
    var arg0 = args[0].(*QActionGroup).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction14setActionGroupEP12QActionGroup(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setActionGroup", args)
  }

  return
}

// menu()
func (this *QAction) Menu(args ...interface{}) (ret interface{}) {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4menuEv
    // invoke: QMenu * menu()
    var ret0 = C.C_ZNK7QAction4menuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "menu", args)
  }

  return
}

// shortcut()
func (this *QAction) Shortcut(args ...interface{}) (ret interface{}) {
  // shortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8shortcutEv
    // invoke: QKeySequence shortcut()
    var ret0 = C.C_ZNK7QAction8shortcutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QKeySequence{}) // "QKeySequence"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "shortcut", args)
  }

  return
}

// whatsThis()
func (this *QAction) Whatsthis(args ...interface{}) (ret interface{}) {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9whatsThisEv
    // invoke: QString whatsThis()
    var ret0 = C.C_ZNK7QAction9whatsThisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "whatsThis", args)
  }

  return
}

// isChecked()
func (this *QAction) Ischecked(args ...interface{}) (ret interface{}) {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isCheckedEv
    // invoke: bool isChecked()
    var ret0 = C.C_ZNK7QAction9isCheckedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isChecked", args)
  }

  return
}

// trigger()
func (this *QAction) Trigger(args ...interface{}) () {
  // trigger()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7triggerEv
    // invoke: void trigger()
    C.C_ZN7QAction7triggerEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "trigger", args)
  }

  return
}

// setCheckable(_Bool)
func (this *QAction) Setcheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setCheckableEb
    // invoke: void setCheckable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction12setCheckableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setCheckable", args)
  }

  return
}

// isIconVisibleInMenu()
func (this *QAction) Isiconvisibleinmenu(args ...interface{}) (ret interface{}) {
  // isIconVisibleInMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction19isIconVisibleInMenuEv
    // invoke: bool isIconVisibleInMenu()
    var ret0 = C.C_ZNK7QAction19isIconVisibleInMenuEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isIconVisibleInMenu", args)
  }

  return
}

// icon()
func (this *QAction) Icon(args ...interface{}) (ret interface{}) {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4iconEv
    // invoke: QIcon icon()
    var ret0 = C.C_ZNK7QAction4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "icon", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QAction) Seticon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setIcon", args)
  }

  return
}

// setData(const class QVariant &)
func (this *QAction) Setdata(args ...interface{}) () {
  // setData(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setDataERK8QVariant
    // invoke: void setData(const class QVariant &)
    var arg0 = args[0].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction7setDataERK8QVariant(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setData", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QAction) Setenabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setEnabled", args)
  }

  return
}

// actionGroup()
func (this *QAction) Actiongroup(args ...interface{}) (ret interface{}) {
  // actionGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11actionGroupEv
    // invoke: QActionGroup * actionGroup()
    var ret0 = C.C_ZNK7QAction11actionGroupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QActionGroup{}) // "QActionGroup *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "actionGroup", args)
  }

  return
}

// hover()
func (this *QAction) Hover(args ...interface{}) () {
  // hover()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction5hoverEv
    // invoke: void hover()
    C.C_ZN7QAction5hoverEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "hover", args)
  }

  return
}

// setIconText(const class QString &)
func (this *QAction) Seticontext(args ...interface{}) () {
  // setIconText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction11setIconTextERK7QString
    // invoke: void setIconText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction11setIconTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setIconText", args)
  }

  return
}

// statusTip()
func (this *QAction) Statustip(args ...interface{}) (ret interface{}) {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9statusTipEv
    // invoke: QString statusTip()
    var ret0 = C.C_ZNK7QAction9statusTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "statusTip", args)
  }

  return
}

// setAutoRepeat(_Bool)
func (this *QAction) Setautorepeat(args ...interface{}) () {
  // setAutoRepeat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction13setAutoRepeatEb
    // invoke: void setAutoRepeat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction13setAutoRepeatEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setAutoRepeat", args)
  }

  return
}

// setShortcut(const class QKeySequence &)
func (this *QAction) Setshortcut(args ...interface{}) () {
  // setShortcut(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction11setShortcutERK12QKeySequence
    // invoke: void setShortcut(const class QKeySequence &)
    var arg0 = args[0].(*qtgui.QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction11setShortcutERK12QKeySequence(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setShortcut", args)
  }

  return
}

// toolTip()
func (this *QAction) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK7QAction7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "toolTip", args)
  }

  return
}

// isEnabled()
func (this *QAction) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK7QAction9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isEnabled", args)
  }

  return
}

// setWhatsThis(const class QString &)
func (this *QAction) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction12setWhatsThisERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setWhatsThis", args)
  }

  return
}

// data()
func (this *QAction) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4dataEv
    // invoke: QVariant data()
    var ret0 = C.C_ZNK7QAction4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "data", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QAction) Settooltip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setToolTip", args)
  }

  return
}

// metaObject()
func (this *QAction) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QAction10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "metaObject", args)
  }

  return
}

// setSeparator(_Bool)
func (this *QAction) Setseparator(args ...interface{}) () {
  // setSeparator(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setSeparatorEb
    // invoke: void setSeparator(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction12setSeparatorEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setSeparator", args)
  }

  return
}

// parentWidget()
func (this *QAction) Parentwidget(args ...interface{}) (ret interface{}) {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction12parentWidgetEv
    // invoke: QWidget * parentWidget()
    var ret0 = C.C_ZNK7QAction12parentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "parentWidget", args)
  }

  return
}

// setText(const class QString &)
func (this *QAction) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setText", args)
  }

  return
}

// QAction(class QObject *)
func NewQAction(args ...interface{}) *QAction {
  // QAction(class QObject *)
  // QAction(const class QIcon &, const class QString &, class QObject *)
  // QAction(const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QActionC1EP7QObject
    // invoke: void QAction(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QActionC2EP7QObject(arg0)
    return &QAction{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QActionC1ERK5QIconRK7QStringP7QObject
    // invoke: void QAction(const class QIcon &, const class QString &, class QObject *)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QActionC2ERK5QIconRK7QStringP7QObject(arg0, arg1, arg2)
    return &QAction{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QActionC1ERK7QStringP7QObject
    // invoke: void QAction(const class QString &, class QObject *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QActionC2ERK7QStringP7QObject(arg0, arg1)
    return &QAction{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAction", "QAction", args)
  }

  return nil // QAction{Qclsinst:qthis}
}

// iconText()
func (this *QAction) Icontext(args ...interface{}) (ret interface{}) {
  // iconText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8iconTextEv
    // invoke: QString iconText()
    var ret0 = C.C_ZNK7QAction8iconTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "iconText", args)
  }

  return
}

// shortcutContext()
func (this *QAction) Shortcutcontext(args ...interface{}) () {
  // shortcutContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction15shortcutContextEv
    // invoke: Qt::ShortcutContext shortcutContext()
    C.C_ZNK7QAction15shortcutContextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "shortcutContext", args)
  }

  return
}

// setDisabled(_Bool)
func (this *QAction) Setdisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction11setDisabledEb
    // invoke: void setDisabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction11setDisabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setDisabled", args)
  }

  return
}

// text()
func (this *QAction) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK7QAction4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "text", args)
  }

  return
}

// autoRepeat()
func (this *QAction) Autorepeat(args ...interface{}) (ret interface{}) {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction10autoRepeatEv
    // invoke: bool autoRepeat()
    var ret0 = C.C_ZNK7QAction10autoRepeatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "autoRepeat", args)
  }

  return
}

// isVisible()
func (this *QAction) Isvisible(args ...interface{}) (ret interface{}) {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isVisibleEv
    // invoke: bool isVisible()
    var ret0 = C.C_ZNK7QAction9isVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAction", "isVisible", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QAction) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setFont", args)
  }

  return
}

// menuRole()
func (this *QAction) Menurole(args ...interface{}) () {
  // menuRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8menuRoleEv
    // invoke: QAction::MenuRole menuRole()
    C.C_ZNK7QAction8menuRoleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "menuRole", args)
  }

  return
}

// setVisible(_Bool)
func (this *QAction) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QAction10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAction", "setVisible", args)
  }

  return
}

// associatedGraphicsWidgets()
func (this *QAction) Associatedgraphicswidgets(args ...interface{}) () {
  // associatedGraphicsWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction25associatedGraphicsWidgetsEv
    // invoke: QList<QGraphicsWidget *> associatedGraphicsWidgets()
    C.C_ZNK7QAction25associatedGraphicsWidgetsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "associatedGraphicsWidgets", args)
  }

  return
}

// priority()
func (this *QAction) Priority(args ...interface{}) () {
  // priority()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8priorityEv
    // invoke: QAction::Priority priority()
    C.C_ZNK7QAction8priorityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAction", "priority", args)
  }

  return
}

// <= body block end

