package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QKeySequence QKeySequence::mnemonic(const QString & text);
extern void* C_ZN12QKeySequence8mnemonicERK7QString(void* arg0); // 4
  // proto:  void QKeySequence::~QKeySequence();
extern void C_ZN12QKeySequenceD2Ev(void* qthis); // 4
  // proto:  bool QKeySequence::isEmpty();
extern bool C_ZNK12QKeySequence7isEmptyEv(void* qthis); // 4
  // proto:  void QKeySequence::swap(QKeySequence & other);
extern void C_ZN12QKeySequence4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QKeySequence::isDetached();
extern bool C_ZNK12QKeySequence10isDetachedEv(void* qthis); // 4
  // proto:  QKeySequence::SequenceMatch QKeySequence::matches(const QKeySequence & seq);
extern void C_ZNK12QKeySequence7matchesERKS_(void* qthis, void* arg0); // 4
  // proto:  void QKeySequence::QKeySequence();
extern void* C_ZN12QKeySequenceC2Ev(); // 3
  // proto:  void QKeySequence::QKeySequence(const QKeySequence & ks);
extern void* C_ZN12QKeySequenceC2ERKS_(void* arg0); // 3
  // proto:  void QKeySequence::QKeySequence(int k1, int k2, int k3, int k4);
extern void* C_ZN12QKeySequenceC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  int QKeySequence::count();
extern int32_t C_ZNK12QKeySequence5countEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QKeySequence)=8
type QKeySequence struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// mnemonic(const class QString &)
func (this *QKeySequence) Mnemonic_S(args ...interface{}) (ret interface{}) {
  // mnemonic(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QKeySequence8mnemonicERK7QString
    // invoke: QKeySequence mnemonic(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QKeySequence8mnemonicERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QKeySequence{}) // "QKeySequence"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeySequence", "mnemonic", args)
  }

  return
}

// ~QKeySequence()
func (this *QKeySequence) Freeqkeysequence(args ...interface{}) () {
  // ~QKeySequence()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QKeySequenceD0Ev
    // invoke: void ~QKeySequence()
    C.C_ZN12QKeySequenceD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeySequence", "~QKeySequence", args)
  }

  return
}

// isEmpty()
func (this *QKeySequence) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QKeySequence7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeySequence", "isEmpty", args)
  }

  return
}

// swap(class QKeySequence &)
func (this *QKeySequence) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QKeySequence4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeySequence", "swap", args)
  }

  return
}

// isDetached()
func (this *QKeySequence) Isdetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QKeySequence10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeySequence", "isDetached", args)
  }

  return
}

// matches(const class QKeySequence &)
func (this *QKeySequence) Matches(args ...interface{}) () {
  // matches(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QKeySequence7matchesERKS_
    // invoke: QKeySequence::SequenceMatch matches(const class QKeySequence &)
    var arg0 = args[0].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK12QKeySequence7matchesERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeySequence", "matches", args)
  }

  return
}

// QKeySequence()
func NewQKeySequence(args ...interface{}) *QKeySequence {
  // QKeySequence()
  // QKeySequence(const class QKeySequence &)
  // QKeySequence(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QKeySequenceC1Ev
    // invoke: void QKeySequence()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QKeySequenceC2Ev()
    return &QKeySequence{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QKeySequenceC1ERKS_
    // invoke: void QKeySequence(const class QKeySequence &)
    var arg0 = args[0].(*QKeySequence).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QKeySequenceC2ERKS_(arg0)
    return &QKeySequence{Qclsinst:qthis}
  case 2:
    // invoke: _ZN12QKeySequenceC1Eiiii
    // invoke: void QKeySequence(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QKeySequenceC2Eiiii(arg0, arg1, arg2, arg3)
    return &QKeySequence{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QKeySequence", "QKeySequence", args)
  }

  return nil // QKeySequence{Qclsinst:qthis}
}

// count()
func (this *QKeySequence) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QKeySequence5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeySequence", "count", args)
  }

  return
}

// <= body block end

