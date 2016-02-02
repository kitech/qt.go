package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN12QActionGroup9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QActionGroup::addAction(const QIcon & icon, const QString & text);
extern void* C_ZN12QActionGroup9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QAction * QActionGroup::addAction(const QString & text);
extern void* C_ZN12QActionGroup9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QActionGroup::actions();
extern void C_ZNK12QActionGroup7actionsEv(void* qthis); // 4
  // proto:  QAction * QActionGroup::checkedAction();
extern void* C_ZNK12QActionGroup13checkedActionEv(void* qthis); // 4
  // proto:  void QActionGroup::setDisabled(bool b);
extern void C_ZN12QActionGroup11setDisabledEb(void* qthis, bool arg0); // 2
  // proto:  void QActionGroup::QActionGroup(QObject * parent);
extern void* C_ZN12QActionGroupC2EP7QObject(void* arg0); // 3
  // proto:  void QActionGroup::setEnabled(bool );
extern void C_ZN12QActionGroup10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QActionGroup::setExclusive(bool );
extern void C_ZN12QActionGroup12setExclusiveEb(void* qthis, bool arg0); // 4
  // proto:  void QActionGroup::removeAction(QAction * a);
extern void C_ZN12QActionGroup12removeActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QActionGroup::isExclusive();
extern bool C_ZNK12QActionGroup11isExclusiveEv(void* qthis); // 4
  // proto:  const QMetaObject * QActionGroup::metaObject();
extern void C_ZNK12QActionGroup10metaObjectEv(void* qthis); // 4
  // proto:  void QActionGroup::~QActionGroup();
extern void C_ZN12QActionGroupD2Ev(void* qthis); // 4
  // proto:  bool QActionGroup::isVisible();
extern bool C_ZNK12QActionGroup9isVisibleEv(void* qthis); // 4
  // proto:  void QActionGroup::setVisible(bool );
extern void C_ZN12QActionGroup10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  bool QActionGroup::isEnabled();
extern bool C_ZNK12QActionGroup9isEnabledEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _triggered QActionGroup_triggered_signal;
//  _hovered QActionGroup_hovered_signal;
}

// addAction(class QAction *)
func (this *QActionGroup) Addaction(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QActionGroup9addActionEP7QAction(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QActionGroup9addActionERK5QIconRK7QString
    // invoke: QAction * addAction(const class QIcon &, const class QString &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QActionGroup9addActionERK5QIconRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN12QActionGroup9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QActionGroup9addActionERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QActionGroup", "addAction", args)
  }

  return
}

// actions()
func (this *QActionGroup) Actions(args ...interface{}) () {
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
    C.C_ZNK12QActionGroup7actionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "actions", args)
  }

  return
}

// checkedAction()
func (this *QActionGroup) Checkedaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QActionGroup13checkedActionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QActionGroup", "checkedAction", args)
  }

  return
}

// setDisabled(_Bool)
func (this *QActionGroup) Setdisabled(args ...interface{}) () {
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
    C.C_ZN12QActionGroup11setDisabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setDisabled", args)
  }

  return
}

// QActionGroup(class QObject *)
func NewQActionGroup(args ...interface{}) *QActionGroup {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QActionGroupC2EP7QObject(arg0)
    return &QActionGroup{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QActionGroup", "QActionGroup", args)
  }

  return nil // QActionGroup{Qclsinst:qthis}
}

// setEnabled(_Bool)
func (this *QActionGroup) Setenabled(args ...interface{}) () {
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
    C.C_ZN12QActionGroup10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setEnabled", args)
  }

  return
}

// setExclusive(_Bool)
func (this *QActionGroup) Setexclusive(args ...interface{}) () {
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
    C.C_ZN12QActionGroup12setExclusiveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setExclusive", args)
  }

  return
}

// removeAction(class QAction *)
func (this *QActionGroup) Removeaction(args ...interface{}) () {
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
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QActionGroup12removeActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "removeAction", args)
  }

  return
}

// isExclusive()
func (this *QActionGroup) Isexclusive(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QActionGroup11isExclusiveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QActionGroup", "isExclusive", args)
  }

  return
}

// metaObject()
func (this *QActionGroup) Metaobject(args ...interface{}) () {
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
    C.C_ZNK12QActionGroup10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "metaObject", args)
  }

  return
}

// ~QActionGroup()
func (this *QActionGroup) Freeqactiongroup(args ...interface{}) () {
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
    C.C_ZN12QActionGroupD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QActionGroup", "~QActionGroup", args)
  }

  return
}

// isVisible()
func (this *QActionGroup) Isvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QActionGroup9isVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QActionGroup", "isVisible", args)
  }

  return
}

// setVisible(_Bool)
func (this *QActionGroup) Setvisible(args ...interface{}) () {
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
    C.C_ZN12QActionGroup10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QActionGroup", "setVisible", args)
  }

  return
}

// isEnabled()
func (this *QActionGroup) Isenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QActionGroup9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QActionGroup", "isEnabled", args)
  }

  return
}

// <= body block end

