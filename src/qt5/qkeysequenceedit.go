package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qkeysequenceedit.h
// dst-file: /src/widgets/qkeysequenceedit.go
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
  // proto:  void QKeySequenceEdit::QKeySequenceEdit(const QKeySequenceEdit & );
extern void* dector_ZN16QKeySequenceEditC1ERKS_(void* arg0);
extern void _ZN16QKeySequenceEditC1ERKS_(void* qthis, void* arg0);
  // proto:  void QKeySequenceEdit::QKeySequenceEdit(const QKeySequence & keySequence, QWidget * parent);
extern void* dector_ZN16QKeySequenceEditC1ERK12QKeySequenceP7QWidget(void* arg0, void* arg1);
extern void _ZN16QKeySequenceEditC1ERK12QKeySequenceP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QKeySequenceEdit::clear();
extern void _ZN16QKeySequenceEdit5clearEv(void* qthis);
  // proto:  void QKeySequenceEdit::setKeySequence(const QKeySequence & keySequence);
extern void _ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence(void* qthis, void* arg0);
  // proto:  QKeySequence QKeySequenceEdit::keySequence();
extern void _ZNK16QKeySequenceEdit11keySequenceEv(void* qthis);
  // proto:  void QKeySequenceEdit::~QKeySequenceEdit();
extern void _ZN16QKeySequenceEditD0Ev(void* qthis);
  // proto:  void QKeySequenceEdit::QKeySequenceEdit(QWidget * parent);
extern void* dector_ZN16QKeySequenceEditC1EP7QWidget(void* arg0);
extern void _ZN16QKeySequenceEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QKeySequenceEdit::metaObject();
extern void _ZNK16QKeySequenceEdit10metaObjectEv(void* qthis);
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

// class sizeof(QKeySequenceEdit)=1
type QKeySequenceEdit struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _editingFinished QKeySequenceEdit_editingFinished_signal;
//  _keySequenceChanged QKeySequenceEdit_keySequenceChanged_signal;
}

  // proto:  void QKeySequenceEdit::QKeySequenceEdit(const QKeySequenceEdit & );
func NewQKeySequenceEdit(args ...interface{}) QKeySequenceEdit {
  return QKeySequenceEdit{}
}

  // proto:  void QKeySequenceEdit::clear();
func (this *QKeySequenceEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit5clearEv
    // invoke: void clear()
    C._ZN16QKeySequenceEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "clear", args)
  }

}

  // proto:  void QKeySequenceEdit::setKeySequence(const QKeySequence & keySequence);
func (this *QKeySequenceEdit) setKeySequence(args ...interface{}) () {
  // setKeySequence(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence
    // invoke: void setKeySequence(const class QKeySequence &)
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "setKeySequence", args)
  }

}

  // proto:  QKeySequence QKeySequenceEdit::keySequence();
func (this *QKeySequenceEdit) keySequence(args ...interface{}) () {
  // keySequence()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit11keySequenceEv
    // invoke: QKeySequence keySequence()
    C._ZNK16QKeySequenceEdit11keySequenceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "keySequence", args)
  }

}

  // proto:  void QKeySequenceEdit::~QKeySequenceEdit();
func (this *QKeySequenceEdit) FreeQKeySequenceEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "~QKeySequenceEdit", args)
  }

}

  // proto:  const QMetaObject * QKeySequenceEdit::metaObject();
func (this *QKeySequenceEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK16QKeySequenceEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "metaObject", args)
  }

}

// <= body block end

