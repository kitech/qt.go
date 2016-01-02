package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QShortcut::setKey(const QKeySequence & key);
extern void _ZN9QShortcut6setKeyERK12QKeySequence(void* qthis, void* arg0);
  // proto:  const QMetaObject * QShortcut::metaObject();
extern void _ZNK9QShortcut10metaObjectEv(void* qthis);
  // proto:  QWidget * QShortcut::parentWidget();
extern void demth_ZNK9QShortcut12parentWidgetEv(void* qthis);
  // proto:  void QShortcut::setAutoRepeat(bool on);
extern void _ZN9QShortcut13setAutoRepeatEb(void* qthis, bool arg0);
  // proto:  bool QShortcut::isEnabled();
extern void _ZNK9QShortcut9isEnabledEv(void* qthis);
  // proto:  QKeySequence QShortcut::key();
extern void _ZNK9QShortcut3keyEv(void* qthis);
  // proto:  void QShortcut::~QShortcut();
extern void _ZN9QShortcutD0Ev(void* qthis);
  // proto:  void QShortcut::setWhatsThis(const QString & text);
extern void _ZN9QShortcut12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  void QShortcut::setEnabled(bool enable);
extern void _ZN9QShortcut10setEnabledEb(void* qthis, bool arg0);
  // proto:  int QShortcut::id();
extern void _ZNK9QShortcut2idEv(void* qthis);
  // proto:  QString QShortcut::whatsThis();
extern void _ZNK9QShortcut9whatsThisEv(void* qthis);
  // proto:  void QShortcut::QShortcut(QWidget * parent);
extern void* dector_ZN9QShortcutC1EP7QWidget(void* arg0);
extern void _ZN9QShortcutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QShortcut::autoRepeat();
extern void _ZNK9QShortcut10autoRepeatEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _activated QShortcut_activated_signal;
//  _activatedAmbiguously QShortcut_activatedAmbiguously_signal;
}

  // proto:  void QShortcut::setKey(const QKeySequence & key);
func (this *QShortcut) setKey(args ...interface{}) () {
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
    C._ZN9QShortcut6setKeyERK12QKeySequence(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setKey", args)
  }

}

  // proto:  const QMetaObject * QShortcut::metaObject();
func (this *QShortcut) metaObject(args ...interface{}) () {
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
    C._ZNK9QShortcut10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "metaObject", args)
  }

}

  // proto:  QWidget * QShortcut::parentWidget();
func (this *QShortcut) parentWidget(args ...interface{}) () {
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
    C.demth_ZNK9QShortcut12parentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "parentWidget", args)
  }

}

  // proto:  void QShortcut::setAutoRepeat(bool on);
func (this *QShortcut) setAutoRepeat(args ...interface{}) () {
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QShortcut13setAutoRepeatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setAutoRepeat", args)
  }

}

  // proto:  bool QShortcut::isEnabled();
func (this *QShortcut) isEnabled(args ...interface{}) () {
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
    C._ZNK9QShortcut9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "isEnabled", args)
  }

}

  // proto:  QKeySequence QShortcut::key();
func (this *QShortcut) key(args ...interface{}) () {
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
    C._ZNK9QShortcut3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "key", args)
  }

}

  // proto:  void QShortcut::~QShortcut();
func (this *QShortcut) FreeQShortcut(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QShortcut", "~QShortcut", args)
  }

}

  // proto:  void QShortcut::setWhatsThis(const QString & text);
func (this *QShortcut) setWhatsThis(args ...interface{}) () {
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
    C._ZN9QShortcut12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setWhatsThis", args)
  }

}

  // proto:  void QShortcut::setEnabled(bool enable);
func (this *QShortcut) setEnabled(args ...interface{}) () {
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QShortcut10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QShortcut", "setEnabled", args)
  }

}

  // proto:  int QShortcut::id();
func (this *QShortcut) id(args ...interface{}) () {
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
    C._ZNK9QShortcut2idEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "id", args)
  }

}

  // proto:  QString QShortcut::whatsThis();
func (this *QShortcut) whatsThis(args ...interface{}) () {
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
    C._ZNK9QShortcut9whatsThisEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "whatsThis", args)
  }

}

  // proto:  void QShortcut::QShortcut(QWidget * parent);
func NewQShortcut(args ...interface{}) QShortcut {
  return QShortcut{}
}

  // proto:  bool QShortcut::autoRepeat();
func (this *QShortcut) autoRepeat(args ...interface{}) () {
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
    C._ZNK9QShortcut10autoRepeatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QShortcut", "autoRepeat", args)
  }

}

// <= body block end

