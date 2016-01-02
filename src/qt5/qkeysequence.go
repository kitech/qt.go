package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
extern void* dector_ZN12QKeySequenceC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN12QKeySequenceC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
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
  qclsinst uint64 /* *mut c_void*/;
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
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QKeySequence", "swap", args)
  }

}

// <= body block end

