package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qactiongroup.h
// dst-file: /src/widgets/qactiongroup.go
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
  // proto:  void QActionGroup::QActionGroup(QObject * parent);
extern void* dector_ZN12QActionGroupC1EP7QObject(void* arg0);
extern void _ZN12QActionGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  QList<QAction *> QActionGroup::actions();
extern void _ZNK12QActionGroup7actionsEv(void* qthis);
  // proto:  void QActionGroup::setDisabled(bool b);
extern void demth_ZN12QActionGroup11setDisabledEb(void* qthis, bool arg0);
  // proto:  void QActionGroup::setEnabled(bool );
extern void _ZN12QActionGroup10setEnabledEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QActionGroup::metaObject();
extern void _ZNK12QActionGroup10metaObjectEv(void* qthis);
  // proto:  QAction * QActionGroup::addAction(QAction * a);
extern void _ZN12QActionGroup9addActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QActionGroup::~QActionGroup();
extern void _ZN12QActionGroupD0Ev(void* qthis);
  // proto:  QAction * QActionGroup::checkedAction();
extern void _ZNK12QActionGroup13checkedActionEv(void* qthis);
  // proto:  QAction * QActionGroup::addAction(const QIcon & icon, const QString & text);
extern void _ZN12QActionGroup9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QActionGroup::setVisible(bool );
extern void _ZN12QActionGroup10setVisibleEb(void* qthis, bool arg0);
  // proto:  bool QActionGroup::isVisible();
extern void _ZNK12QActionGroup9isVisibleEv(void* qthis);
  // proto:  void QActionGroup::setExclusive(bool );
extern void _ZN12QActionGroup12setExclusiveEb(void* qthis, bool arg0);
  // proto:  QAction * QActionGroup::addAction(const QString & text);
extern void _ZN12QActionGroup9addActionERK7QString(void* qthis, void* arg0);
  // proto:  bool QActionGroup::isEnabled();
extern void _ZNK12QActionGroup9isEnabledEv(void* qthis);
  // proto:  bool QActionGroup::isExclusive();
extern void _ZNK12QActionGroup11isExclusiveEv(void* qthis);
  // proto:  void QActionGroup::removeAction(QAction * a);
extern void _ZN12QActionGroup12removeActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QActionGroup::QActionGroup(const QActionGroup & );
extern void* dector_ZN12QActionGroupC1ERKS_(void* arg0);
extern void _ZN12QActionGroupC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QActionGroup)=1
type QActionGroup struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _triggered QActionGroup_triggered_signal;
//  _hovered QActionGroup_hovered_signal;
}

  // proto:  void QActionGroup::QActionGroup(QObject * parent);
func NewQActionGroup(args ...interface{}) QActionGroup {
  return QActionGroup{}
}

  // proto:  QList<QAction *> QActionGroup::actions();
func (this *QActionGroup) actions(args ...interface{}) () {
  // actions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup7actionsEv
  default:
    qtrt.ErrorResolve("QActionGroup", "actions", args)
  }

}

  // proto:  void QActionGroup::setDisabled(bool b);
func (this *QActionGroup) setDisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup11setDisabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "setDisabled", args)
  }

}

  // proto:  void QActionGroup::setEnabled(bool );
func (this *QActionGroup) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup10setEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "setEnabled", args)
  }

}

  // proto:  const QMetaObject * QActionGroup::metaObject();
func (this *QActionGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup10metaObjectEv
  default:
    qtrt.ErrorResolve("QActionGroup", "metaObject", args)
  }

}

  // proto:  QAction * QActionGroup::addAction(QAction * a);
func (this *QActionGroup) addAction(args ...interface{}) () {
  // addAction(class QAction *)
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup9addActionEP7QAction
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN12QActionGroup9addActionERK5QIconRK7QString
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZN12QActionGroup9addActionERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "addAction", args)
  }

}

  // proto:  void QActionGroup::~QActionGroup();
func (this *QActionGroup) FreeQActionGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QActionGroup", "~QActionGroup", args)
  }

}

  // proto:  QAction * QActionGroup::checkedAction();
func (this *QActionGroup) checkedAction(args ...interface{}) () {
  // checkedAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup13checkedActionEv
  default:
    qtrt.ErrorResolve("QActionGroup", "checkedAction", args)
  }

}

  // proto:  void QActionGroup::setVisible(bool );
func (this *QActionGroup) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup10setVisibleEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "setVisible", args)
  }

}

  // proto:  bool QActionGroup::isVisible();
func (this *QActionGroup) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup9isVisibleEv
  default:
    qtrt.ErrorResolve("QActionGroup", "isVisible", args)
  }

}

  // proto:  void QActionGroup::setExclusive(bool );
func (this *QActionGroup) setExclusive(args ...interface{}) () {
  // setExclusive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup12setExclusiveEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "setExclusive", args)
  }

}

  // proto:  bool QActionGroup::isEnabled();
func (this *QActionGroup) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup9isEnabledEv
  default:
    qtrt.ErrorResolve("QActionGroup", "isEnabled", args)
  }

}

  // proto:  bool QActionGroup::isExclusive();
func (this *QActionGroup) isExclusive(args ...interface{}) () {
  // isExclusive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QActionGroup11isExclusiveEv
  default:
    qtrt.ErrorResolve("QActionGroup", "isExclusive", args)
  }

}

  // proto:  void QActionGroup::removeAction(QAction * a);
func (this *QActionGroup) removeAction(args ...interface{}) () {
  // removeAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroup12removeActionEP7QAction
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QActionGroup", "removeAction", args)
  }

}

// <= body block end

