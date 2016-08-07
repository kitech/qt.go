package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  void QKeySequenceEdit::QKeySequenceEdit(QWidget * parent);
extern void* C_ZN16QKeySequenceEditC2EP7QWidget(void* arg0); // 3
  // proto:  void QKeySequenceEdit::QKeySequenceEdit(const QKeySequence & keySequence, QWidget * parent);
extern void* C_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QKeySequenceEdit::clear();
extern void C_ZN16QKeySequenceEdit5clearEv(void* qthis); // 4
  // proto:  const QMetaObject * QKeySequenceEdit::metaObject();
extern void C_ZNK16QKeySequenceEdit10metaObjectEv(void* qthis); // 4
  // proto:  QKeySequence QKeySequenceEdit::keySequence();
extern void* C_ZNK16QKeySequenceEdit11keySequenceEv(void* qthis); // 4
  // proto:  void QKeySequenceEdit::~QKeySequenceEdit();
extern void C_ZN16QKeySequenceEditD2Ev(void* qthis); // 4
  // proto:  void QKeySequenceEdit::setKeySequence(const QKeySequence & keySequence);
extern void C_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence(void* qthis, void* arg0); // 4
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

// class sizeof(QKeySequenceEdit)=1
type QKeySequenceEdit struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _editingFinished QKeySequenceEdit_editingFinished_signal;
//  _keySequenceChanged QKeySequenceEdit_keySequenceChanged_signal;
}

// QKeySequenceEdit(class QWidget *)
func NewQKeySequenceEdit(args ...interface{}) *QKeySequenceEdit {
  // QKeySequenceEdit(class QWidget *)
  // QKeySequenceEdit(const class QKeySequence &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QKeySequence{}) // "const QKeySequence &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEditC1EP7QWidget
    // invoke: void QKeySequenceEdit(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QKeySequenceEditC2EP7QWidget(arg0)
    return &QKeySequenceEdit{Qclsinst:qthis}
  case 1:
    // invoke: _ZN16QKeySequenceEditC1ERK12QKeySequenceP7QWidget
    // invoke: void QKeySequenceEdit(const class QKeySequence &, class QWidget *)
    var arg0 = args[0].(*qtgui.QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget(arg0, arg1)
    return &QKeySequenceEdit{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "QKeySequenceEdit", args)
  }

  return nil // QKeySequenceEdit{Qclsinst:qthis}
}

// clear()
func (this *QKeySequenceEdit) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit5clearEv
    // invoke: void clear()
    C.C_ZN16QKeySequenceEdit5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "clear", args)
  }

  return
}

// metaObject()
func (this *QKeySequenceEdit) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QKeySequenceEdit10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "metaObject", args)
  }

  return
}

// keySequence()
func (this *QKeySequenceEdit) Keysequence(args ...interface{}) (ret interface{}) {
  // keySequence()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QKeySequenceEdit11keySequenceEv
    // invoke: QKeySequence keySequence()
    var ret0 = C.C_ZNK16QKeySequenceEdit11keySequenceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QKeySequence{}) // "QKeySequence"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "keySequence", args)
  }

  return
}

// ~QKeySequenceEdit()
func (this *QKeySequenceEdit) Freeqkeysequenceedit(args ...interface{}) () {
  // ~QKeySequenceEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEditD0Ev
    // invoke: void ~QKeySequenceEdit()
    C.C_ZN16QKeySequenceEditD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "~QKeySequenceEdit", args)
  }

  return
}

// setKeySequence(const class QKeySequence &)
func (this *QKeySequenceEdit) Setkeysequence(args ...interface{}) () {
  // setKeySequence(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QKeySequence{}) // "const QKeySequence &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence
    // invoke: void setKeySequence(const class QKeySequence &)
    var arg0 = args[0].(*qtgui.QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeySequenceEdit", "setKeySequence", args)
  }

  return
}

// <= body block end

