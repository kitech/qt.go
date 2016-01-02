package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QList<QWidget *> QAction::associatedWidgets();
extern void _ZNK7QAction17associatedWidgetsEv(void* qthis);
  // proto:  void QAction::setAutoRepeat(bool );
extern void _ZN7QAction13setAutoRepeatEb(void* qthis, bool arg0);
  // proto:  QString QAction::whatsThis();
extern void _ZNK7QAction9whatsThisEv(void* qthis);
  // proto:  void QAction::QAction(const QString & text, QObject * parent);
extern void* dector_ZN7QActionC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN7QActionC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QAction::~QAction();
extern void _ZN7QActionD0Ev(void* qthis);
  // proto:  bool QAction::isVisible();
extern void _ZNK7QAction9isVisibleEv(void* qthis);
  // proto:  void QAction::setFont(const QFont & font);
extern void _ZN7QAction7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QAction::setData(const QVariant & var);
extern void _ZN7QAction7setDataERK8QVariant(void* qthis, void* arg0);
  // proto:  void QAction::setIcon(const QIcon & icon);
extern void _ZN7QAction7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QAction::QAction(QObject * parent);
extern void* dector_ZN7QActionC1EP7QObject(void* arg0);
extern void _ZN7QActionC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAction::metaObject();
extern void _ZNK7QAction10metaObjectEv(void* qthis);
  // proto:  void QAction::setText(const QString & text);
extern void _ZN7QAction7setTextERK7QString(void* qthis, void* arg0);
  // proto:  bool QAction::showStatusText(QWidget * widget);
extern void _ZN7QAction14showStatusTextEP7QWidget(void* qthis, void* arg0);
  // proto:  QString QAction::iconText();
extern void _ZNK7QAction8iconTextEv(void* qthis);
  // proto:  void QAction::setIconVisibleInMenu(bool visible);
extern void _ZN7QAction20setIconVisibleInMenuEb(void* qthis, bool arg0);
  // proto:  QString QAction::statusTip();
extern void _ZNK7QAction9statusTipEv(void* qthis);
  // proto:  bool QAction::isCheckable();
extern void _ZNK7QAction11isCheckableEv(void* qthis);
  // proto:  void QAction::setWhatsThis(const QString & what);
extern void _ZN7QAction12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  QMenu * QAction::menu();
extern void _ZNK7QAction4menuEv(void* qthis);
  // proto:  void QAction::trigger();
extern void _ZN7QAction7triggerEv(void* qthis);
  // proto:  QFont QAction::font();
extern void _ZNK7QAction4fontEv(void* qthis);
  // proto:  void QAction::QAction(const QIcon & icon, const QString & text, QObject * parent);
extern void* dector_ZN7QActionC1ERK5QIconRK7QStringP7QObject(void* arg0, void* arg1, void* arg2);
extern void _ZN7QActionC1ERK5QIconRK7QStringP7QObject(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QAction::setChecked(bool );
extern void _ZN7QAction10setCheckedEb(void* qthis, bool arg0);
  // proto:  void QAction::setDisabled(bool b);
extern void demth_ZN7QAction11setDisabledEb(void* qthis, bool arg0);
  // proto:  void QAction::setShortcut(const QKeySequence & shortcut);
extern void _ZN7QAction11setShortcutERK12QKeySequence(void* qthis, void* arg0);
  // proto:  void QAction::toggle();
extern void _ZN7QAction6toggleEv(void* qthis);
  // proto:  QKeySequence QAction::shortcut();
extern void _ZNK7QAction8shortcutEv(void* qthis);
  // proto:  QIcon QAction::icon();
extern void _ZNK7QAction4iconEv(void* qthis);
  // proto:  void QAction::setActionGroup(QActionGroup * group);
extern void _ZN7QAction14setActionGroupEP12QActionGroup(void* qthis, void* arg0);
  // proto:  void QAction::setMenu(QMenu * menu);
extern void _ZN7QAction7setMenuEP5QMenu(void* qthis, void* arg0);
  // proto:  QList<QKeySequence> QAction::shortcuts();
extern void _ZNK7QAction9shortcutsEv(void* qthis);
  // proto:  void QAction::setCheckable(bool );
extern void _ZN7QAction12setCheckableEb(void* qthis, bool arg0);
  // proto:  QString QAction::toolTip();
extern void _ZNK7QAction7toolTipEv(void* qthis);
  // proto:  QWidget * QAction::parentWidget();
extern void _ZNK7QAction12parentWidgetEv(void* qthis);
  // proto:  void QAction::setSeparator(bool b);
extern void _ZN7QAction12setSeparatorEb(void* qthis, bool arg0);
  // proto:  QList<QGraphicsWidget *> QAction::associatedGraphicsWidgets();
extern void _ZNK7QAction25associatedGraphicsWidgetsEv(void* qthis);
  // proto:  void QAction::setVisible(bool );
extern void _ZN7QAction10setVisibleEb(void* qthis, bool arg0);
  // proto:  bool QAction::isSeparator();
extern void _ZNK7QAction11isSeparatorEv(void* qthis);
  // proto:  void QAction::setIconText(const QString & text);
extern void _ZN7QAction11setIconTextERK7QString(void* qthis, void* arg0);
  // proto:  QActionGroup * QAction::actionGroup();
extern void _ZNK7QAction11actionGroupEv(void* qthis);
  // proto:  void QAction::setStatusTip(const QString & statusTip);
extern void _ZN7QAction12setStatusTipERK7QString(void* qthis, void* arg0);
  // proto:  QVariant QAction::data();
extern void _ZNK7QAction4dataEv(void* qthis);
  // proto:  void QAction::QAction(const QAction & );
extern void* dector_ZN7QActionC1ERKS_(void* arg0);
extern void _ZN7QActionC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QAction::isChecked();
extern void _ZNK7QAction9isCheckedEv(void* qthis);
  // proto:  bool QAction::autoRepeat();
extern void _ZNK7QAction10autoRepeatEv(void* qthis);
  // proto:  bool QAction::isEnabled();
extern void _ZNK7QAction9isEnabledEv(void* qthis);
  // proto:  QString QAction::text();
extern void _ZNK7QAction4textEv(void* qthis);
  // proto:  void QAction::hover();
extern void _ZN7QAction5hoverEv(void* qthis);
  // proto:  bool QAction::isIconVisibleInMenu();
extern void _ZNK7QAction19isIconVisibleInMenuEv(void* qthis);
  // proto:  void QAction::setToolTip(const QString & tip);
extern void _ZN7QAction10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  void QAction::setEnabled(bool );
extern void _ZN7QAction10setEnabledEb(void* qthis, bool arg0);
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

// class sizeof(QAction)=1
type QAction struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _changed QAction_changed_signal;
//  _hovered QAction_hovered_signal;
//  _triggered QAction_triggered_signal;
//  _toggled QAction_toggled_signal;
}

  // proto:  QList<QWidget *> QAction::associatedWidgets();
func (this *QAction) associatedWidgets(args ...interface{}) () {
  // associatedWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction17associatedWidgetsEv
  default:
    qtrt.ErrorResolve("QAction", "associatedWidgets", args)
  }

}

  // proto:  void QAction::setAutoRepeat(bool );
func (this *QAction) setAutoRepeat(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setAutoRepeat", args)
  }

}

  // proto:  QString QAction::whatsThis();
func (this *QAction) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9whatsThisEv
  default:
    qtrt.ErrorResolve("QAction", "whatsThis", args)
  }

}

  // proto:  void QAction::QAction(const QString & text, QObject * parent);
func NewQAction(args ...interface{}) QAction {
  return QAction{}
}

  // proto:  void QAction::~QAction();
func (this *QAction) FreeQAction(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAction", "~QAction", args)
  }

}

  // proto:  bool QAction::isVisible();
func (this *QAction) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isVisibleEv
  default:
    qtrt.ErrorResolve("QAction", "isVisible", args)
  }

}

  // proto:  void QAction::setFont(const QFont & font);
func (this *QAction) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setFontERK5QFont
  default:
    qtrt.ErrorResolve("QAction", "setFont", args)
  }

}

  // proto:  void QAction::setData(const QVariant & var);
func (this *QAction) setData(args ...interface{}) () {
  // setData(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setDataERK8QVariant
  default:
    qtrt.ErrorResolve("QAction", "setData", args)
  }

}

  // proto:  void QAction::setIcon(const QIcon & icon);
func (this *QAction) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QAction", "setIcon", args)
  }

}

  // proto:  const QMetaObject * QAction::metaObject();
func (this *QAction) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction10metaObjectEv
  default:
    qtrt.ErrorResolve("QAction", "metaObject", args)
  }

}

  // proto:  void QAction::setText(const QString & text);
func (this *QAction) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7setTextERK7QString
  default:
    qtrt.ErrorResolve("QAction", "setText", args)
  }

}

  // proto:  bool QAction::showStatusText(QWidget * widget);
func (this *QAction) showStatusText(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "showStatusText", args)
  }

}

  // proto:  QString QAction::iconText();
func (this *QAction) iconText(args ...interface{}) () {
  // iconText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8iconTextEv
  default:
    qtrt.ErrorResolve("QAction", "iconText", args)
  }

}

  // proto:  void QAction::setIconVisibleInMenu(bool visible);
func (this *QAction) setIconVisibleInMenu(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setIconVisibleInMenu", args)
  }

}

  // proto:  QString QAction::statusTip();
func (this *QAction) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9statusTipEv
  default:
    qtrt.ErrorResolve("QAction", "statusTip", args)
  }

}

  // proto:  bool QAction::isCheckable();
func (this *QAction) isCheckable(args ...interface{}) () {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11isCheckableEv
  default:
    qtrt.ErrorResolve("QAction", "isCheckable", args)
  }

}

  // proto:  void QAction::setWhatsThis(const QString & what);
func (this *QAction) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setWhatsThisERK7QString
  default:
    qtrt.ErrorResolve("QAction", "setWhatsThis", args)
  }

}

  // proto:  QMenu * QAction::menu();
func (this *QAction) menu(args ...interface{}) () {
  // menu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4menuEv
  default:
    qtrt.ErrorResolve("QAction", "menu", args)
  }

}

  // proto:  void QAction::trigger();
func (this *QAction) trigger(args ...interface{}) () {
  // trigger()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction7triggerEv
  default:
    qtrt.ErrorResolve("QAction", "trigger", args)
  }

}

  // proto:  QFont QAction::font();
func (this *QAction) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4fontEv
  default:
    qtrt.ErrorResolve("QAction", "font", args)
  }

}

  // proto:  void QAction::setChecked(bool );
func (this *QAction) setChecked(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setChecked", args)
  }

}

  // proto:  void QAction::setDisabled(bool b);
func (this *QAction) setDisabled(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setDisabled", args)
  }

}

  // proto:  void QAction::setShortcut(const QKeySequence & shortcut);
func (this *QAction) setShortcut(args ...interface{}) () {
  // setShortcut(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction11setShortcutERK12QKeySequence
  default:
    qtrt.ErrorResolve("QAction", "setShortcut", args)
  }

}

  // proto:  void QAction::toggle();
func (this *QAction) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction6toggleEv
  default:
    qtrt.ErrorResolve("QAction", "toggle", args)
  }

}

  // proto:  QKeySequence QAction::shortcut();
func (this *QAction) shortcut(args ...interface{}) () {
  // shortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction8shortcutEv
  default:
    qtrt.ErrorResolve("QAction", "shortcut", args)
  }

}

  // proto:  QIcon QAction::icon();
func (this *QAction) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4iconEv
  default:
    qtrt.ErrorResolve("QAction", "icon", args)
  }

}

  // proto:  void QAction::setActionGroup(QActionGroup * group);
func (this *QAction) setActionGroup(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setActionGroup", args)
  }

}

  // proto:  void QAction::setMenu(QMenu * menu);
func (this *QAction) setMenu(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setMenu", args)
  }

}

  // proto:  QList<QKeySequence> QAction::shortcuts();
func (this *QAction) shortcuts(args ...interface{}) () {
  // shortcuts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9shortcutsEv
  default:
    qtrt.ErrorResolve("QAction", "shortcuts", args)
  }

}

  // proto:  void QAction::setCheckable(bool );
func (this *QAction) setCheckable(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setCheckable", args)
  }

}

  // proto:  QString QAction::toolTip();
func (this *QAction) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction7toolTipEv
  default:
    qtrt.ErrorResolve("QAction", "toolTip", args)
  }

}

  // proto:  QWidget * QAction::parentWidget();
func (this *QAction) parentWidget(args ...interface{}) () {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction12parentWidgetEv
  default:
    qtrt.ErrorResolve("QAction", "parentWidget", args)
  }

}

  // proto:  void QAction::setSeparator(bool b);
func (this *QAction) setSeparator(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setSeparator", args)
  }

}

  // proto:  QList<QGraphicsWidget *> QAction::associatedGraphicsWidgets();
func (this *QAction) associatedGraphicsWidgets(args ...interface{}) () {
  // associatedGraphicsWidgets()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction25associatedGraphicsWidgetsEv
  default:
    qtrt.ErrorResolve("QAction", "associatedGraphicsWidgets", args)
  }

}

  // proto:  void QAction::setVisible(bool );
func (this *QAction) setVisible(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setVisible", args)
  }

}

  // proto:  bool QAction::isSeparator();
func (this *QAction) isSeparator(args ...interface{}) () {
  // isSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11isSeparatorEv
  default:
    qtrt.ErrorResolve("QAction", "isSeparator", args)
  }

}

  // proto:  void QAction::setIconText(const QString & text);
func (this *QAction) setIconText(args ...interface{}) () {
  // setIconText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction11setIconTextERK7QString
  default:
    qtrt.ErrorResolve("QAction", "setIconText", args)
  }

}

  // proto:  QActionGroup * QAction::actionGroup();
func (this *QAction) actionGroup(args ...interface{}) () {
  // actionGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction11actionGroupEv
  default:
    qtrt.ErrorResolve("QAction", "actionGroup", args)
  }

}

  // proto:  void QAction::setStatusTip(const QString & statusTip);
func (this *QAction) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction12setStatusTipERK7QString
  default:
    qtrt.ErrorResolve("QAction", "setStatusTip", args)
  }

}

  // proto:  QVariant QAction::data();
func (this *QAction) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4dataEv
  default:
    qtrt.ErrorResolve("QAction", "data", args)
  }

}

  // proto:  bool QAction::isChecked();
func (this *QAction) isChecked(args ...interface{}) () {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isCheckedEv
  default:
    qtrt.ErrorResolve("QAction", "isChecked", args)
  }

}

  // proto:  bool QAction::autoRepeat();
func (this *QAction) autoRepeat(args ...interface{}) () {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction10autoRepeatEv
  default:
    qtrt.ErrorResolve("QAction", "autoRepeat", args)
  }

}

  // proto:  bool QAction::isEnabled();
func (this *QAction) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction9isEnabledEv
  default:
    qtrt.ErrorResolve("QAction", "isEnabled", args)
  }

}

  // proto:  QString QAction::text();
func (this *QAction) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction4textEv
  default:
    qtrt.ErrorResolve("QAction", "text", args)
  }

}

  // proto:  void QAction::hover();
func (this *QAction) hover(args ...interface{}) () {
  // hover()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction5hoverEv
  default:
    qtrt.ErrorResolve("QAction", "hover", args)
  }

}

  // proto:  bool QAction::isIconVisibleInMenu();
func (this *QAction) isIconVisibleInMenu(args ...interface{}) () {
  // isIconVisibleInMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QAction19isIconVisibleInMenuEv
  default:
    qtrt.ErrorResolve("QAction", "isIconVisibleInMenu", args)
  }

}

  // proto:  void QAction::setToolTip(const QString & tip);
func (this *QAction) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QAction10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QAction", "setToolTip", args)
  }

}

  // proto:  void QAction::setEnabled(bool );
func (this *QAction) setEnabled(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QAction", "setEnabled", args)
  }

}

// <= body block end

