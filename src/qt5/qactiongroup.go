package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QAction * QActionGroup::addAction(QAction * a);
extern void _ZN12QActionGroup9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QActionGroup::addAction(const QIcon & icon, const QString & text);
extern void _ZN12QActionGroup9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QActionGroup::addAction(const QString & text);
extern void _ZN12QActionGroup9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QActionGroup::actions();
extern void _ZNK12QActionGroup7actionsEv(void* qthis); // 4
  // proto:  QAction * QActionGroup::checkedAction();
extern void _ZNK12QActionGroup13checkedActionEv(void* qthis); // 4
  // proto:  void QActionGroup::setDisabled(bool b);
extern void _ZN12QActionGroup11setDisabledEb(void* qthis, bool arg0); // 2
  // proto:  void QActionGroup::QActionGroup(QObject * parent);
extern void _ZN12QActionGroupC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QActionGroup::setEnabled(bool );
extern void _ZN12QActionGroup10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QActionGroup::setExclusive(bool );
extern void _ZN12QActionGroup12setExclusiveEb(void* qthis, bool arg0); // 4
  // proto:  void QActionGroup::removeAction(QAction * a);
extern void _ZN12QActionGroup12removeActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QActionGroup::isExclusive();
extern void _ZNK12QActionGroup11isExclusiveEv(void* qthis); // 4
  // proto:  const QMetaObject * QActionGroup::metaObject();
extern void _ZNK12QActionGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QActionGroup::~QActionGroup();
extern void _ZN12QActionGroupD2Ev(void* qthis); // 4
  // proto:  bool QActionGroup::isVisible();
extern void _ZNK12QActionGroup9isVisibleEv(void* qthis); // 4
  // proto:  void QActionGroup::setVisible(bool );
extern void _ZN12QActionGroup10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  bool QActionGroup::isEnabled();
extern void _ZNK12QActionGroup9isEnabledEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _triggered QActionGroup_triggered_signal;
//  _hovered QActionGroup_hovered_signal;
}

// addAction(class QAction *)
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
    // invoke: QAction * addAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup9addActionEP7QAction(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QActionGroup9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN12QActionGroup9addActionERK5QIconRK7QString(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN12QActionGroup9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup9addActionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "addAction", args)
  }

}

// actions()
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
    // invoke: QList<QAction *> actions()
    C._ZNK12QActionGroup7actionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "actions", args)
  }

}

// checkedAction()
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
    // invoke: QAction * checkedAction()
    C._ZNK12QActionGroup13checkedActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "checkedAction", args)
  }

}

// setDisabled(_Bool)
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
    // invoke: void setDisabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup11setDisabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setDisabled", args)
  }

}

// QActionGroup(class QObject *)
func NewQActionGroup(args ...interface{}) QActionGroup {
  // QActionGroup(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroupC1EP7QObject
    // invoke: void QActionGroup(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QActionGroupC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "QActionGroup", args)
  }

  return QActionGroup{}
}

// setEnabled(_Bool)
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
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setEnabled", args)
  }

}

// setExclusive(_Bool)
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
    // invoke: void setExclusive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup12setExclusiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setExclusive", args)
  }

}

// removeAction(class QAction *)
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
    // invoke: void removeAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup12removeActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "removeAction", args)
  }

}

// isExclusive()
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
    // invoke: bool isExclusive()
    C._ZNK12QActionGroup11isExclusiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "isExclusive", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK12QActionGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "metaObject", args)
  }

}

// ~QActionGroup()
func (this *QActionGroup) FreeQActionGroup(args ...interface{}) () {
  // ~QActionGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QActionGroupD0Ev
    // invoke: void ~QActionGroup()
    C._ZN12QActionGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "~QActionGroup", args)
  }

}

// isVisible()
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
    // invoke: bool isVisible()
    C._ZNK12QActionGroup9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "isVisible", args)
  }

}

// setVisible(_Bool)
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN12QActionGroup10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setVisible", args)
  }

}

// isEnabled()
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
    // invoke: bool isEnabled()
    C._ZNK12QActionGroup9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "isEnabled", args)
  }

}

// <= body block end

