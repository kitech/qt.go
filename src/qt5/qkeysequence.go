package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qkeysequence.h
// dst-file: /src/gui/qkeysequence.go
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
  // proto:  bool QKeySequence::isDetached();
extern void _ZNK12QKeySequence10isDetachedEv(void* qthis);
  // proto:  bool QKeySequence::isEmpty();
extern void _ZNK12QKeySequence7isEmptyEv(void* qthis);
  // proto:  void QKeySequence::QKeySequence(const QKeySequence & ks);
extern void* dector_ZN12QKeySequenceC1ERKS_(void* arg0);
extern void _ZN12QKeySequenceC1ERKS_(void* qthis, void* arg0);
  // proto:  void QKeySequence::QKeySequence(int k1, int k2, int k3, int k4);
extern void* dector_ZN12QKeySequenceC1Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
extern void _ZN12QKeySequenceC1Eiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  int QKeySequence::count();
extern void _ZNK12QKeySequence5countEv(void* qthis);
  // proto: static QKeySequence QKeySequence::mnemonic(const QString & text);
extern void _ZN12QKeySequence8mnemonicERK7QString(void* arg0);
  // proto:  void QKeySequence::QKeySequence();
extern void* dector_ZN12QKeySequenceC1Ev();
extern void _ZN12QKeySequenceC1Ev(void* qthis);
  // proto:  void QKeySequence::~QKeySequence();
extern void _ZN12QKeySequenceD0Ev(void* qthis);
  // proto:  void QKeySequence::swap(QKeySequence & other);
extern void demth_ZN12QKeySequence4swapERS_(void* qthis, void* arg0);
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

// class sizeof(QKeySequence)=8
type QKeySequence struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  bool QKeySequence::isDetached();
func (this *QKeySequence) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QKeySequence10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK12QKeySequence10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequence", "isDetached", args)
  }

}

  // proto:  bool QKeySequence::isEmpty();
func (this *QKeySequence) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QKeySequence7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK12QKeySequence7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequence", "isEmpty", args)
  }

}

  // proto:  void QKeySequence::QKeySequence(const QKeySequence & ks);
func NewQKeySequence(args ...interface{}) QKeySequence {
  return QKeySequence{}
}

  // proto:  int QKeySequence::count();
func (this *QKeySequence) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QKeySequence5countEv
    // invoke: int count()
    C._ZNK12QKeySequence5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequence", "count", args)
  }

}

  // proto: static QKeySequence QKeySequence::mnemonic(const QString & text);
func (this *QKeySequence) mnemonic_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeySequence", "mnemonic", args)
  }

}

  // proto:  void QKeySequence::~QKeySequence();
func (this *QKeySequence) FreeQKeySequence(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QKeySequence", "~QKeySequence", args)
  }

}

  // proto:  void QKeySequence::swap(QKeySequence & other);
func (this *QKeySequence) swap(args ...interface{}) () {
  // swap(class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QKeySequence4swapERS_
    // invoke: void swap(class QKeySequence &)
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN12QKeySequence4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeySequence", "swap", args)
  }

}

// <= body block end

