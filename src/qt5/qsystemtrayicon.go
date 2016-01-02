package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qsystemtrayicon.h
// dst-file: /src/widgets/qsystemtrayicon.go
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
  // proto:  void QSystemTrayIcon::~QSystemTrayIcon();
extern void _ZN15QSystemTrayIconD0Ev(void* qthis);
  // proto:  void QSystemTrayIcon::setVisible(bool visible);
extern void _ZN15QSystemTrayIcon10setVisibleEb(void* qthis, bool arg0);
  // proto:  QString QSystemTrayIcon::toolTip();
extern void _ZNK15QSystemTrayIcon7toolTipEv(void* qthis);
  // proto:  void QSystemTrayIcon::QSystemTrayIcon(const QIcon & icon, QObject * parent);
extern void* dector_ZN15QSystemTrayIconC1ERK5QIconP7QObject(void* arg0, void* arg1);
extern void _ZN15QSystemTrayIconC1ERK5QIconP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QSystemTrayIcon::hide();
extern void demth_ZN15QSystemTrayIcon4hideEv(void* qthis);
  // proto:  const QMetaObject * QSystemTrayIcon::metaObject();
extern void _ZNK15QSystemTrayIcon10metaObjectEv(void* qthis);
  // proto:  void QSystemTrayIcon::QSystemTrayIcon(const QSystemTrayIcon & );
extern void* dector_ZN15QSystemTrayIconC1ERKS_(void* arg0);
extern void _ZN15QSystemTrayIconC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSystemTrayIcon::setIcon(const QIcon & icon);
extern void _ZN15QSystemTrayIcon7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  bool QSystemTrayIcon::isVisible();
extern void _ZNK15QSystemTrayIcon9isVisibleEv(void* qthis);
  // proto:  void QSystemTrayIcon::QSystemTrayIcon(QObject * parent);
extern void* dector_ZN15QSystemTrayIconC1EP7QObject(void* arg0);
extern void _ZN15QSystemTrayIconC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QSystemTrayIcon::show();
extern void demth_ZN15QSystemTrayIcon4showEv(void* qthis);
  // proto: static bool QSystemTrayIcon::supportsMessages();
extern void _ZN15QSystemTrayIcon16supportsMessagesEv();
  // proto:  void QSystemTrayIcon::setContextMenu(QMenu * menu);
extern void _ZN15QSystemTrayIcon14setContextMenuEP5QMenu(void* qthis, void* arg0);
  // proto:  QRect QSystemTrayIcon::geometry();
extern void _ZNK15QSystemTrayIcon8geometryEv(void* qthis);
  // proto:  void QSystemTrayIcon::setToolTip(const QString & tip);
extern void _ZN15QSystemTrayIcon10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  QIcon QSystemTrayIcon::icon();
extern void _ZNK15QSystemTrayIcon4iconEv(void* qthis);
  // proto:  QMenu * QSystemTrayIcon::contextMenu();
extern void _ZNK15QSystemTrayIcon11contextMenuEv(void* qthis);
  // proto: static bool QSystemTrayIcon::isSystemTrayAvailable();
extern void _ZN15QSystemTrayIcon21isSystemTrayAvailableEv();
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

// class sizeof(QSystemTrayIcon)=1
type QSystemTrayIcon struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QSystemTrayIcon_activated_signal;
//  _messageClicked QSystemTrayIcon_messageClicked_signal;
}

  // proto:  void QSystemTrayIcon::~QSystemTrayIcon();
func (this *QSystemTrayIcon) FreeQSystemTrayIcon(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "~QSystemTrayIcon", args)
  }

}

  // proto:  void QSystemTrayIcon::setVisible(bool visible);
func (this *QSystemTrayIcon) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN15QSystemTrayIcon10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setVisible", args)
  }

}

  // proto:  QString QSystemTrayIcon::toolTip();
func (this *QSystemTrayIcon) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon7toolTipEv
    // invoke: QString toolTip()
    C._ZNK15QSystemTrayIcon7toolTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "toolTip", args)
  }

}

  // proto:  void QSystemTrayIcon::QSystemTrayIcon(const QIcon & icon, QObject * parent);
func NewQSystemTrayIcon(args ...interface{}) QSystemTrayIcon {
  return QSystemTrayIcon{}
}

  // proto:  void QSystemTrayIcon::hide();
func (this *QSystemTrayIcon) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4hideEv
    // invoke: void hide()
    C.demth_ZN15QSystemTrayIcon4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "hide", args)
  }

}

  // proto:  const QMetaObject * QSystemTrayIcon::metaObject();
func (this *QSystemTrayIcon) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QSystemTrayIcon10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "metaObject", args)
  }

}

  // proto:  void QSystemTrayIcon::setIcon(const QIcon & icon);
func (this *QSystemTrayIcon) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QSystemTrayIcon7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setIcon", args)
  }

}

  // proto:  bool QSystemTrayIcon::isVisible();
func (this *QSystemTrayIcon) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon9isVisibleEv
    // invoke: bool isVisible()
    C._ZNK15QSystemTrayIcon9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isVisible", args)
  }

}

  // proto:  void QSystemTrayIcon::show();
func (this *QSystemTrayIcon) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4showEv
    // invoke: void show()
    C.demth_ZN15QSystemTrayIcon4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "show", args)
  }

}

  // proto: static bool QSystemTrayIcon::supportsMessages();
func (this *QSystemTrayIcon) supportsMessages_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "supportsMessages", args)
  }

}

  // proto:  void QSystemTrayIcon::setContextMenu(QMenu * menu);
func (this *QSystemTrayIcon) setContextMenu(args ...interface{}) () {
  // setContextMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon14setContextMenuEP5QMenu
    // invoke: void setContextMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QSystemTrayIcon14setContextMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setContextMenu", args)
  }

}

  // proto:  QRect QSystemTrayIcon::geometry();
func (this *QSystemTrayIcon) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon8geometryEv
    // invoke: QRect geometry()
    C._ZNK15QSystemTrayIcon8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "geometry", args)
  }

}

  // proto:  void QSystemTrayIcon::setToolTip(const QString & tip);
func (this *QSystemTrayIcon) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QSystemTrayIcon10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setToolTip", args)
  }

}

  // proto:  QIcon QSystemTrayIcon::icon();
func (this *QSystemTrayIcon) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon4iconEv
    // invoke: QIcon icon()
    C._ZNK15QSystemTrayIcon4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "icon", args)
  }

}

  // proto:  QMenu * QSystemTrayIcon::contextMenu();
func (this *QSystemTrayIcon) contextMenu(args ...interface{}) () {
  // contextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon11contextMenuEv
    // invoke: QMenu * contextMenu()
    C._ZNK15QSystemTrayIcon11contextMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "contextMenu", args)
  }

}

  // proto: static bool QSystemTrayIcon::isSystemTrayAvailable();
func (this *QSystemTrayIcon) isSystemTrayAvailable_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isSystemTrayAvailable", args)
  }

}

// <= body block end

