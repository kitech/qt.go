package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qshortcut.h
// dst-file: /src/widgets/qshortcut.go
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
  // proto:  void QShortcut::~QShortcut();
extern void C_ZN9QShortcutD2Ev(void* qthis); // 4
  // proto:  int QShortcut::id();
extern int32_t C_ZNK9QShortcut2idEv(void* qthis); // 4
  // proto:  bool QShortcut::isEnabled();
extern bool C_ZNK9QShortcut9isEnabledEv(void* qthis); // 4
  // proto:  QString QShortcut::whatsThis();
extern void* C_ZNK9QShortcut9whatsThisEv(void* qthis); // 4
  // proto:  void QShortcut::setKey(const QKeySequence & key);
extern void C_ZN9QShortcut6setKeyERK12QKeySequence(void* qthis, void* arg0); // 4
  // proto:  void QShortcut::setEnabled(bool enable);
extern void C_ZN9QShortcut10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QShortcut::setAutoRepeat(bool on);
extern void C_ZN9QShortcut13setAutoRepeatEb(void* qthis, bool arg0); // 4
  // proto:  void QShortcut::QShortcut(QWidget * parent);
extern void* C_ZN9QShortcutC2EP7QWidget(void* arg0); // 3
  // proto:  void QShortcut::setWhatsThis(const QString & text);
extern void C_ZN9QShortcut12setWhatsThisERK7QString(void* qthis, void* arg0); // 4
  // proto:  QKeySequence QShortcut::key();
extern void* C_ZNK9QShortcut3keyEv(void* qthis); // 4
  // proto:  const QMetaObject * QShortcut::metaObject();
extern void C_ZNK9QShortcut10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QShortcut::parentWidget();
extern void* C_ZNK9QShortcut12parentWidgetEv(void* qthis); // 2
  // proto:  Qt::ShortcutContext QShortcut::context();
extern void C_ZNK9QShortcut7contextEv(void* qthis); // 4
  // proto:  bool QShortcut::autoRepeat();
extern bool C_ZNK9QShortcut10autoRepeatEv(void* qthis); // 4
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

// class sizeof(QShortcut)=1
type QShortcut struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _activated QShortcut_activated_signal;
//  _activatedAmbiguously QShortcut_activatedAmbiguously_signal;
}

// ~QShortcut()
func (this *QShortcut) Freeqshortcut(args ...interface{}) () {
  // ~QShortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcutD0Ev
    // invoke: void ~QShortcut()
    C.C_ZN9QShortcutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "~QShortcut", args)
  }

  return
}

// id()
func (this *QShortcut) Id(args ...interface{}) (ret interface{}) {
  // id()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut2idEv
    // invoke: int id()
    var ret0 = C.C_ZNK9QShortcut2idEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "id", args)
  }

  return
}

// isEnabled()
func (this *QShortcut) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK9QShortcut9isEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "isEnabled", args)
  }

  return
}

// whatsThis()
func (this *QShortcut) Whatsthis(args ...interface{}) (ret interface{}) {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut9whatsThisEv
    // invoke: QString whatsThis()
    var ret0 = C.C_ZNK9QShortcut9whatsThisEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "whatsThis", args)
  }

  return
}

// setKey(const class QKeySequence &)
func (this *QShortcut) Setkey(args ...interface{}) () {
  // setKey(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcut6setKeyERK12QKeySequence
    // invoke: void setKey(const class QKeySequence &)
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QShortcut6setKeyERK12QKeySequence(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setKey", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QShortcut) Setenabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcut10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QShortcut10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setEnabled", args)
  }

  return
}

// setAutoRepeat(_Bool)
func (this *QShortcut) Setautorepeat(args ...interface{}) () {
  // setAutoRepeat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcut13setAutoRepeatEb
    // invoke: void setAutoRepeat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QShortcut13setAutoRepeatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setAutoRepeat", args)
  }

  return
}

// QShortcut(class QWidget *)
func NewQShortcut(args ...interface{}) *QShortcut {
  // QShortcut(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcutC1EP7QWidget
    // invoke: void QShortcut(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QShortcutC2EP7QWidget(arg0)
    return &QShortcut{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QShortcut", "QShortcut", args)
  }

  return nil // QShortcut{qclsinst:qthis}
}

// setWhatsThis(const class QString &)
func (this *QShortcut) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QShortcut12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QShortcut12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setWhatsThis", args)
  }

  return
}

// key()
func (this *QShortcut) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut3keyEv
    // invoke: QKeySequence key()
    var ret0 = C.C_ZNK9QShortcut3keyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QKeySequence{}) // "QKeySequence"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "key", args)
  }

  return
}

// metaObject()
func (this *QShortcut) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QShortcut10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "metaObject", args)
  }

  return
}

// parentWidget()
func (this *QShortcut) Parentwidget(args ...interface{}) (ret interface{}) {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut12parentWidgetEv
    // invoke: QWidget * parentWidget()
    var ret0 = C.C_ZNK9QShortcut12parentWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "parentWidget", args)
  }

  return
}

// context()
func (this *QShortcut) Context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut7contextEv
    // invoke: Qt::ShortcutContext context()
    C.C_ZNK9QShortcut7contextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "context", args)
  }

  return
}

// autoRepeat()
func (this *QShortcut) Autorepeat(args ...interface{}) (ret interface{}) {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QShortcut10autoRepeatEv
    // invoke: bool autoRepeat()
    var ret0 = C.C_ZNK9QShortcut10autoRepeatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QShortcut", "autoRepeat", args)
  }

  return
}

// <= body block end

